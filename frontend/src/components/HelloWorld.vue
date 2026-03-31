<script setup>
import { onMounted, ref } from "vue";
import { Events } from "@wailsio/runtime";
import {
  AzureDevopsService,
  NotificationService,
} from "../../bindings/changeme/internal/services";

const tickets = ref([]);

function systemNotif() {
  NotificationService.SystemNotification("id", "title", "body", "subtitle");
}

onMounted(async () => {
  // Listen for events from Go
  Events.On("tickets-updated", (newTickets) => {
    tickets.value = newTickets.data;
  });

  // Fetch initial tickets
  const data = await AzureDevopsService.FetchAssignedTickets();

  tickets.value = data;
});
</script>

<template>
  <div @click="systemNotif">System Notif</div>

  <ul>
    <li v-for="ticket in tickets" :key="ticket.ID">
      <strong>{{ ticket.Title }}</strong> - {{ ticket.State }}
      <p>{{ ticket.Description }}</p>
      <div v-if="ticket.PRLinks">
        PRs:
        <ul>
          <li v-for="pr in ticket.PRLinks" :key="pr">{{ pr }}</li>
        </ul>
      </div>
    </li>
  </ul>
</template>
