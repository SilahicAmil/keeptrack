<script setup>
import { ref } from "vue";
import Button from "./Button.vue";

// Props: list of strings and optional placeholder
defineProps({
  items: { type: Array, required: true },
  placeholder: { type: String, default: "Select an option" },
});

const selected = ref(""); // local selection

defineEmits(["update:modelValue"]);
</script>

<template>
  <div class="w-fit">
    <div class="relative">
      <select
        v-model="selected"
        @change="$emit('update:modelValue', $event.target.value)"
        class="w-full bg-white text-slate-700 text-sm border border-slate-300 rounded-md pl-3 pr-8 py-2 transition focus:outline-none focus:border-slate-500 hover:border-slate-400 shadow-sm appearance-none cursor-pointer"
      >
        <option value="" disabled>{{ placeholder }}</option>
        <option v-for="item in items" :key="item" :value="item">
          {{ item }}
        </option>
      </select>
      <span
        class="pointer-events-none absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 text-xs"
      >
        ▼
      </span>
    </div>
  </div>
</template>
