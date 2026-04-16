<script setup>
import { onMounted, ref } from "vue";
import { Events } from "@wailsio/runtime";
import Button from "../components/ui/Button.vue";
import { AzureDevopsService } from "../../bindings/changeme/internal/services";

const tickets = ref([]);
const loaded = ref(false);

onMounted(async () => {
  const isLoaded = await AzureDevopsService.CheckAppState();

  console.log(isLoaded);
  console.log(loaded.value);

  loaded.value = isLoaded;

  // if (loaded.value) {
  //   router.push({ path: "/dashboard" });
  // }
});
</script>

<template>
  <div
    class="min-h-screen flex flex-col items-center justify-center p-6 bg-gradient-to-br from-slate-100 via-slate-50 to-slate-200"
  >
    <div class="w-full max-w-md bg-white rounded-2xl shadow-xl p-8 space-y-6">
      <div class="text-center space-y-2">
        <div
          class="mx-auto w-12 h-12 rounded-xl bg-slate-800 flex items-center justify-center text-white text-xl font-bold"
        ></div>
        <h1 class="text-2xl font-semibold text-slate-800">
          Welcome to KeepTrack
        </h1>
        <p class="text-sm text-slate-600">
          Connect an integration to start tracking your tickets.
        </p>
      </div>
      <div v-if="!loaded" class="flex justify-center">
        <router-link to="/setup" class="cursor-pointer"
          ><Button variant="primary">Get Started</Button></router-link
        >
      </div>
      <div v-if="loaded" class="flex justify-center">
        <router-link to="/dashboard" class="cursor-pointer"
          ><Button variant="primary">Continue</Button></router-link
        >
      </div>
    </div>
  </div>
</template>
