<script setup lang="ts">
import { ref, onMounted } from "vue";
import { FetchAssignedTicketsCache } from "../../bindings/changeme/internal/services/azuredevopsservice";

const tickets = ref<any[]>([]);

onMounted(async () => {
  tickets.value = await FetchAssignedTicketsCache();
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
