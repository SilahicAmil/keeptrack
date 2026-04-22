<script setup lang="ts">
import router from "@/router";
import { ref } from "vue";
import { AzureDevopsService } from "../../bindings/changeme/internal/services";
import Button from "@/components/ui/Button.vue";

const PAT = ref("");
const Org = ref("");
const Project = ref("");

const loading = ref("");

type AzureCFG = {
  provider: string;
  pat: string;
  org: string;
  project: string;
  validate: boolean;
};

async function checkPat(validate: boolean) {
  const jsonPostObj: AzureCFG = {
    provider: "azure",
    pat: PAT.value,
    org: Org.value,
    project: Project.value,
    validate,
  };

  try {
    loading.value = "Loading...";

    const valid = await AzureDevopsService.InitializeApp(jsonPostObj);

    if (valid && validate) {
      loading.value = "Configuration Valid!";
      return;
    }

    if (!validate) {
      router.push({ path: "/dashboard" });
    }
  } catch (e) {
    // TODO: eventually do something with this
    console.log(e);
  }
}
</script>

<template>
  <!-- TODO: Make the UI for this like a login page -->

  <div
    class="min-h-screen flex flex-col items-center justify-center p-6 bg-linear-to-br from-slate-100 via-slate-50 to-slate-200"
  >
    <h1 class="text-2xl font-mono mb-2">Azure Devops Setup</h1>
    <p v-if="loading">{{ loading }}</p>
    <div
      class="w-full max-w-lg bg-white rounded-2xl shadow-xl p-8 space-y-6 flex flex-col mb-2"
    >
      <!-- Step 1 - Ask for PAT + Org + Team -->
      <input
        type="password"
        placeholder="Enter your PAT. Or else..."
        v-model="PAT"
      />
      <input
        type="text"
        placeholder="Enter your Organization. Or else..."
        v-model="Org"
      />
      <input
        type="text"
        placeholder="Enter your Project. Or else..."
        v-model="Project"
      />
    </div>
    <div class="flex justify-center gap-24 max-w-lg">
      <div class="flex justify-center">
        <Button variant="primary" @click="checkPat(true)">Check</Button>
      </div>
      <div class="flex justify-center">
        <Button variant="primary" @click="checkPat(false)">Continue</Button>
      </div>
    </div>
  </div>
</template>
