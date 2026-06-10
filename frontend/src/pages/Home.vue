<script setup>
import { onMounted, ref } from "vue";
import { Events } from "@wailsio/runtime";
import Button from "../components/ui/Button.vue";
import { AzureDevopsService } from "../../bindings/changeme/internal/services";

const tickets = ref([]);
const loaded = ref(false);
const updatePAT = ref(false);

onMounted(async () => {
  try {
    const isLoaded = await AzureDevopsService.CheckAppState();
    const cfgValid = await AzureDevopsService.ValidateConfig();

    loaded.value = isLoaded;
    updatePAT.value = false;
  } catch (e) {
    loaded.value = false;
    updatePAT.value = true;
  }
});
</script>

<template>
  <div
    class="min-h-screen flex flex-col items-center justify-center p-6 bg-linear-to-br from-slate-100 via-slate-50 to-slate-200"
  >
    <div class="w-full max-w-md bg-white rounded-2xl shadow-xl p-8 space-y-6">
      <div class="text-center space-y-2">
        <h1 class="text-2xl font-semibold text-slate-800">Keeptrack.</h1>
        <p v-if="!loaded && !updatePAT" class="text-sm text-slate-600">
          Connect an integration to start tracking your tickets.
        </p>
        <p v-if="loaded && !updatePAT">Welcome Back. Let's get to work!</p>
        <p v-if="updatePAT">PAT invalid. Please input a new one below.</p>
      </div>

      <div v-if="!loaded" class="flex justify-center">
        <router-link to="/setup" class="cursor-pointer"
          ><Button variant="primary">Get Started</Button></router-link
        >
      </div>

      <div v-if="loaded && !updatePAT" class="flex justify-center">
        <router-link to="/dashboard" class="cursor-pointer"
          ><Button variant="primary">Continue</Button></router-link
        >
      </div>

      <div v-if="updatePAT" class="flex justify-center gap-2">
        <input type="password" class="border-2 rounded-xl" />
        <Button variant="primary">Update</Button>
      </div>
    </div>
  </div>
</template>
