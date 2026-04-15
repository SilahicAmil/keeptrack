<script setup lang="ts">
import router from "@/router";
import { InitializeApp } from "../../bindings/changeme/internal/services/azuredevopsservice";
import { ref } from "vue";
import { AzureDevopsService } from "../../bindings/changeme/internal/services";

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

  console.log(jsonPostObj);
  try {
    loading.value = "Validating Configuration. Please Wait.";

    await AzureDevopsService.InitializeApp(jsonPostObj);

    if (!validate) {
      router.push({ path: "/dashboard" });
    }
  } catch (e) {
    console.log(e);
  }
}
</script>

<template>
  <!-- TODO: Make the UI for this like a login page -->
  <h1>Azure Devops setup</h1>

  <p v-if="loading">{{ loading }}</p>

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
  <button @click="checkPat(true)">Check</button>
  <button @click="checkPat(false)">Complete</button>
  <!-- Loading spinner component while async calls -->
  <!-- Redirect to dashboard. With "pre-loaded" data. -->
</template>
