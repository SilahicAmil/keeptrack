<script setup lang="ts">
import { Browser } from "@wailsio/runtime";
import { AzureDevopsService } from "../../../bindings/changeme/internal/services";

defineProps<{
  tickets: any[];
}>();

function openTicket(ticketID: number) {
  AzureDevopsService.OpenTicket(ticketID);
}
</script>

<template>
  <div>
    <!-- Table Header -->
    <div
      class="grid grid-cols-[120px_160px_1fr_80px_40px] items-center px-5 py-3 text-xs text-slate-400 uppercase border-b border-slate-700/30"
    >
      <span>Ticket #</span>
      <span>Title</span>
      <span>Description</span>
      <span>State</span>
      <span></span>
    </div>

    <!-- Table Rows -->
    <div
      @click="openTicket(ticket.ID)"
      v-for="ticket in tickets"
      :key="ticket.ID"
      class="grid grid-cols-[120px_160px_1fr_80px_40px] items-center px-5 py-3 border-b border-slate-700/20 hover:bg-slate-800/30 transition cursor-pointer"
    >
      <span class="text-sm text-slate-300 font-mono"
        >#&nbsp;{{ ticket.ID }}</span
      >
      <span class="text-sm text-slate-200"
        >{{ ticket.Title.substring(0, 25) }}...</span
      >
      <span class="text-sm text-slate-300 font-mono truncate">{{
        ticket.Description
      }}</span>
      <span>
        <span
          class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium bg-slate-500/20 text-slate-300"
        >
          {{ ticket.State ?? "" }}
        </span>
      </span>
    </div>

    <!-- Empty state -->
    <div
      v-if="tickets.length === 0"
      class="px-5 py-12 text-center text-slate-500 text-sm"
    >
      No work items found.
    </div>
  </div>
</template>
