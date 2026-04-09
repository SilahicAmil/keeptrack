<script setup>
import { onMounted, ref } from "vue";
import { Events } from "@wailsio/runtime";
import Button from "../components/ui/Button.vue";

const tickets = ref([]);

onMounted(async () => {
  Events.On("tickets-updated", (newTickets) => {
    tickets.value = newTickets.data;
  });
});
</script>

<template>
  <div
    class="min-h-screen flex flex-col items-center justify-center p-6 bg-gradient-to-br from-slate-100 via-slate-50 to-slate-200"
  >
    <div class="w-full max-w-md bg-white rounded-2xl shadow-xl p-8 space-y-6">
      <div class="text-center space-y-2">
        <div
          class="mx-auto w-12 h-12 rounded-xl bg-slate-800 flex items-center justify-center text-white text-xl font-bold"
        ></div>
        <h1 class="text-2xl font-semibold text-slate-800">
          Welcome to KeepTrack
        </h1>
        <p class="text-sm text-slate-600">
          Connect an integration to start tracking your tickets.
        </p>
      </div>
      <div class="flex justify-center">
        <router-link to="/setup" class="cursor-pointer"
          ><Button variant="primary">Get Started</Button></router-link
        >
      </div>
    </div>

    <!-- TODO: Remove Later. This is for testing only -->
    <ul class="mt-8 w-full max-w-md">
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
  </div>
</template>
