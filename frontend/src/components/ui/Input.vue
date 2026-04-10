<script lang="js">
import { defineProps } from "vue";
defineProps({
  type: { type: String, required: true },
  placeholder: { type: String, default: "Fill in the input... Or else." },
});
</script>

<template setup>
  <div>
    <div>
      <input :type="{ type }" :placeholder="{ placeholder }" />
    </div>
  </div>
</template>
