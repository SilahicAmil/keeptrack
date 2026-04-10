<script setup lang="ts">
import router from "@/router";
import { ValidateConfig } from "../../bindings/changeme/internal/services/azuredevopsservice";
import { ref } from "vue";

const PAT = ref("");
const Org = ref("");
const Project = ref("");

const loading = ref("");

async function checkPat() {
  var jsonPostObj = {
    pat: PAT.value,
    org: Org.value,
    project: Project.value,
  };

  console.log(jsonPostObj);
  try {
    loading.value = "Loading Dashboard. Please wait...";
    await ValidateConfig(jsonPostObj);

    loading.value = "";

    router.push("/dashboard");
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
  <button @click="checkPat">Complete</button>
  <!-- Loading spinner component while async calls -->
  <!-- Redirect to dashboard. With "pre-loaded" data. -->
</template>
