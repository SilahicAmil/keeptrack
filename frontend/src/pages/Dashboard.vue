<script setup lang="ts">
import { ref, onMounted } from "vue";
import { Events } from "@wailsio/runtime";
import {
  FetchAssignedTicketsCache,
  FetchPullRequests,
  StartPolling,
} from "../../bindings/changeme/internal/services/azuredevopsservice";
import Navbar from "../components/ui/Navbar.vue";
import TabBar from "../components/dashboard/TabBar.vue";
import WorkItemsTable from "../components/dashboard/WorkItemsTable.vue";
import ActivityFeed from "../components/dashboard/ActivityFeed.vue";
import { NotificationService } from "../../bindings/changeme/internal/services";
import CodeReview from "@/components/dashboard/CodeReview.vue";

const tickets = ref<any[]>([]);
const notifiedAt = new Map<string, string>(); // ticket ID -> ChangedDate
const activeTab = ref("workItems");
const activity = ref<any[]>([]);

function isRecent(dateStr: string) {
  if (!dateStr) return false;
  const updated = new Date(dateStr).getTime();
  const now = Date.now();
  // If was within last 30 mins it's recent enough
  return now - updated < 30 * 60 * 1000;
}

function alreadyNotified(ticket: any) {
  return notifiedAt.get(String(ticket.ID)) === ticket.ChangedDate;
}

function markNotified(ticket: any) {
  notifiedAt.set(String(ticket.ID), ticket.ChangedDate);
}

function addActivity(ticket: any, message: string) {
  if (alreadyNotified(ticket)) return;
  markNotified(ticket);

  activity.value.unshift({
    id: ticket.ID,
    ticketId: `TICKET-${ticket.ID}`,
    message,
    time: ticket.ChangedDate,
    isNew: true,
  });
  if (activity.value.length > 20) {
    activity.value = activity.value.slice(0, 20);
  }
}

function markActivityRead() {
  for (const item of activity.value) {
    item.isNew = false;
  }
}

async function notify(ticket: any) {
  await NotificationService.SystemNotification(
    `${ticket.ID}`,
    "New Update for your work item",
    `Ticket #${ticket.ID} has been updated`,
    "",
  );
  addActivity(ticket, `Something has been updated. Check it out!`);
}

onMounted(async () => {
  // Start poller
  await StartPolling();
  tickets.value = await FetchAssignedTicketsCache();

  // seed activity from recently changed tickets
  for (const t of tickets.value) {
    if (isRecent(t.ChangedDate)) {
      addActivity(t, `Something has been updated. Check it out!`);
    }
  }

  Events.On("tickets-updated", (event) => {
    const newTickets = event.data;
    tickets.value = newTickets;

    console.log("updated", tickets.value);

    for (const t of newTickets) {
      if (isRecent(t.ChangedDate) && !alreadyNotified(t)) {
        notify(t);
      }
    }
  });
});
</script>

<template>
  <div class="min-h-screen bg-[#0b1120] text-white flex flex-col">
    <Navbar />

    <div class="flex flex-1 gap-6 p-6 overflow-hidden">
      <!-- Main Content: Tickets -->
      <div class="flex-1 min-w-0">
        <div
          class="bg-[#111827] rounded-xl border border-slate-700/40 overflow-hidden"
        >
          <TabBar v-model:activeTab="activeTab" />

          <WorkItemsTable v-if="activeTab === 'workItems'" :tickets="tickets" />

          <CodeReview v-if="activeTab === 'codeReviews'" />

          <ActivityFeed
            v-if="activeTab === 'activity'"
            :activity="activity"
            @viewed="markActivityRead"
          />
        </div>
      </div>
    </div>
  </div>
</template>
