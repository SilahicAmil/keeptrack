<script setup lang="ts">
import { onMounted } from "vue";
import { timeAgo } from "../../utils/status";

const props = defineProps<{
  activity: any[];
}>();

const emit = defineEmits<{
  (e: "viewed"): void;
}>();

onMounted(() => {
  emit("viewed");
});
</script>

<template>
  <div>
    <div
      v-for="item in activity"
      :key="item.id + item.time"
      class="px-5 py-3 border-b border-slate-700/20 hover:bg-slate-800/30 transition"
    >
      <div class="flex items-start gap-3">
        <div
          v-if="item.isNew"
          class="mt-1.5 w-2 h-2 rounded-full shrink-0 bg-emerald-400"
        ></div>
        <div class="min-w-0">
          <p class="text-sm text-white font-medium">
            {{ item.ticketId }} updated
          </p>
          <p class="text-xs text-slate-400 mt-0.5">
            {{ item.message }}
          </p>
          <p class="text-xs text-slate-500 mt-1">
            {{ timeAgo(item.time) }}
          </p>
        </div>
      </div>
    </div>
    <div
      v-if="activity.length === 0"
      class="px-5 py-12 text-center text-slate-500 text-sm"
    >
      No recent activity.
    </div>
  </div>
</template>
