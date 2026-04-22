<script setup lang="ts">
import { statusDotColor, timeAgo } from "../../utils/status";

defineProps<{
  notifications: any[];
}>();
</script>

<template>
  <div class="w-72 shrink-0">
    <div
      class="bg-[#111827] rounded-xl border border-slate-700/40 overflow-hidden"
    >
      <div
        class="flex items-center justify-between px-5 py-4 border-b border-slate-700/40"
      >
        <h3 class="text-sm font-semibold text-white">Notifications</h3>
        <button
          class="text-xs text-blue-400 hover:text-blue-300 transition cursor-pointer"
        >
          View all
        </button>
      </div>

      <div class="divide-y divide-slate-700/20">
        <div
          v-for="notif in notifications"
          :key="notif.id + notif.time"
          class="px-5 py-3 hover:bg-slate-800/30 transition"
        >
          <div class="flex items-start gap-3">
            <div
              class="mt-1 w-2 h-2 rounded-full shrink-0"
              :class="statusDotColor(notif.state)"
            ></div>
            <div class="min-w-0">
              <p class="text-sm text-white font-medium">
                {{ notif.ticketId }} updated
              </p>
              <p class="text-xs text-slate-400 mt-0.5 truncate">
                {{ notif.message }}
              </p>
              <p class="text-xs text-slate-500 mt-1">
                {{ timeAgo(notif.time) }}
              </p>
            </div>
          </div>
        </div>
      </div>

      <div
        v-if="notifications.length === 0"
        class="px-5 py-8 text-center text-slate-500 text-sm"
      >
        No notifications yet.
      </div>
    </div>
  </div>
</template>
