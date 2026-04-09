<script setup lang="ts">
import Button from "@/components/ui/Button.vue";
import Select from "@/components/ui/Select.vue";
import router from "@/router";
import { ref } from "vue";
const integrationItems = ["Azure Devops"];

const selectedIntegration = ref("");
const error = ref("");

function continueSetup() {
  switch (selectedIntegration.value) {
    case "Azure Devops":
      router.push("/azure");
      break;

    default:
      error.value = "Please Sselect an Integration to Continue";
      break;
  }
}
</script>

<template>
  <p v-if="error" class="bg-red-500 text-white">{{ error }}</p>
  <div class="m-10">
    <h1>Please select an integration</h1>
    <br />
    <Select :items="integrationItems" v-model="selectedIntegration" />
    <Button @click="continueSetup" class="mt-5">Continue</Button>
  </div>
</template>
