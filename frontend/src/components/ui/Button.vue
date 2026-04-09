<script setup>
defineProps({
  variant: { type: String, default: "primary" }, // primary | secondary
  type: { type: String, default: "button" },
  disabled: { type: Boolean, default: false },
});
</script>

<template>
  <button
    :type="type"
    :disabled="disabled"
    class="px-4 py-2 text-sm rounded-md font-medium transition disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"
    :class="
      variant === 'primary'
        ? 'bg-slate-800 text-white hover:bg-slate-700'
        : 'bg-white text-slate-700 border border-slate-300 hover:bg-slate-100'
    "
  >
    <slot />
  </button>
</template>
