<script setup lang="ts">
defineProps<{
  activeTab: string;
}>();

defineEmits<{
  (e: "update:activeTab", value: string): void;
}>();

const tabs = [
  { key: "workItems", label: "Work Items" },
  { key: "codeReviews", label: "Code Reviews" },
  { key: "activity", label: "Activity" },
];
</script>

<template>
  <div class="flex border-b border-slate-700/40">
    <button
      v-for="tab in tabs"
      :key="tab.key"
      @click="$emit('update:activeTab', tab.key)"
      class="px-5 py-3 text-sm font-medium transition cursor-pointer"
      :class="
        activeTab === tab.key
          ? 'text-white border-b-2 border-blue-500'
          : 'text-slate-400 hover:text-slate-200'
      "
    >
      {{ tab.label }}
    </button>
  </div>
</template>
