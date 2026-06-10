<script setup>
import { NotificationService } from "../../../bindings/changeme/internal/services";
import { ref } from "vue";

const searchQuery = ref("");

async function enableNotifications() {
  await NotificationService.SystemNotification(
    "init",
    "Notifications enabled",
    "You’ll now receive updates",
    "",
  );
}
</script>

<template>
  <nav
    class="flex items-center justify-between px-6 py-3 bg-[#0b1120] border-b border-slate-700/50"
  >
    <div class="flex items-center gap-3">
      <div class="w-7 h-7 rounded-md bg-white flex items-center justify-center">
        <span class="text-black text-xs font-bold">K</span>
      </div>
      <span class="text-bac font-semibold text-sm tracking-wide"
        >keeptrack.</span
      >
    </div>

    <!-- <Button @click="enableNotifications"> Enable Notifications </Button> -->
    <div class="flex-1 max-w-md mx-8">
      <div class="relative">
        <svg
          class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-500"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          />
        </svg>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="(Coming Soon) Search tickets, code reviews, or users..."
          disabled="true"
          class="w-full bg-[#131b2e] text-slate-300 text-sm rounded-lg pl-10 pr-4 py-2 border border-slate-700/50 focus:outline-none focus:border-slate-500 placeholder-slate-500"
        />
      </div>
    </div>

    <!-- TODO: On Click have dropdown to update PAT/ORG/PROJECT -->
    <div class="flex items-center gap-3">
      <div
        class="w-8 h-8 rounded-full bg-slate-600 flex items-center justify-center"
      >
        <span class="text-white text-xs">U</span>
      </div>
    </div>
  </nav>
</template>
