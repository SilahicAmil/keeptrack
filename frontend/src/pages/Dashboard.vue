<script setup lang="ts">
import { ref, onMounted } from "vue";
import { Events } from "@wailsio/runtime";
import { FetchAssignedTicketsCache } from "../../bindings/changeme/internal/services/azuredevopsservice";
import { NotificationService } from "../../bindings/changeme/internal/services";

const tickets = ref<any[]>([]);
const notified = new Set<number>();

function isRecent(dateStr: string) {
  if (!dateStr) return false;

  const updated = new Date(dateStr).getTime();
  const now = Date.now();

  return now - updated < 30 * 60 * 1000; // 30 mins
}

async function notify(ticket: any) {
  // Send the notif of an update happened. keep vague for now
  await NotificationService.SystemNotification(
    `${ticket.ID}`,
    "New Update for your work item",
    `Ticket #${ticket.ID} has been updated`,
    "",
  );
}

onMounted(async () => {
  // init: fetch from cache
  // mybe implement some type of one time loading state or something
  // so users know we are fetching from the API
  // Or may show a spinner every 55 seconds or so. that way it looks like it's refreshing
  tickets.value = await FetchAssignedTicketsCache();

  Events.On("tickets-updated", (event) => {
    const newTickets = event.data;

    tickets.value = newTickets;

    for (const t of newTickets) {
      if (isRecent(t.ChangedDate) && !notified.has(t.ID)) {
        notify(t);
        notified.add(t.ID);
      }
    }
  });
});
</script>

<template>
  <!-- TODO: Remove Later. This is for testing only -->
  <ul class="mt-8 w-full max-w-md">
    <li v-for="ticket in tickets" :key="ticket.ID">
      <strong>{{ ticket.Title }}</strong> - {{ ticket.State }}
      <p>{{ ticket.Description }}</p>
      <div v-if="ticket.PRLinks">
        PRs: - COMING SOON
        <ul>
          <li v-for="pr in ticket.PRLinks" :key="pr">{{ pr }}</li>
        </ul>
      </div>
    </li>
  </ul>
</template>
