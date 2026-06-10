<script setup lang="ts">
import { PullRequest } from "bindings/changeme/internal/services/models";
import {
  FetchPullRequests,
  OpenPR,
} from "../../../bindings/changeme/internal/services/azuredevopsservice";
import { Events } from "@wailsio/runtime";

import { onMounted, ref, watch } from "vue";

const prs = ref<PullRequest[]>([]);

// defineProps<{
//   prs: any[];
// }>();

onMounted(async () => {
  prs.value = await FetchPullRequests();

  Events.On("prs-update", (event) => {
    const newPRs = event.data;
    prs.value = newPRs;

    console.log("updated", prs.value);
  });
});

watch(prs, (val) => {
  console.log("prs updated", val);
});

function OpenPullRequest(PRId: number) {
  OpenPR(PRId);
}
</script>

<template>
  <div>
    <!-- Table Header -->
    <div
      class="grid grid-cols-[120px_160px_1fr_80px_60px] items-center px-5 py-3 text-xs text-slate-400 uppercase border-b border-slate-700/30"
    >
      <span>PR #</span>
      <span>Title</span>
      <span>Description</span>
      <span>Reviewers</span>
      <span></span>
    </div>

    <!-- Table Rows -->
    <div
      @click="OpenPullRequest(pr.ID)"
      v-for="pr in prs"
      :key="pr.ID"
      class="grid grid-cols-[120px_160px_1fr_80px_60px] items-center px-5 py-3 border-b border-slate-700/20 hover:bg-slate-800/30 transition cursor-pointer"
    >
      <span class="text-sm text-slate-300 font-mono">#&nbsp;{{ pr.ID }}</span>
      <span class="text-sm text-slate-200">{{ pr.Title }}</span>
      <span class="text-sm text-slate-300 font-mono truncate">{{
        pr.Description
      }}</span>
      <span>
        <span class="flex gap-2 flex-wrap">
          <span
            v-for="r in pr.Reviewers"
            :key="r.DisplayName"
            class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
            :class="{
              'bg-green-500/20 text-green-300': r.Vote > 5,
              'bg-red-500/20 text-red-300': r.Vote < 0,
              'bg-orange-500/20 text-orange-300': r.Vote === 5,
              'bg-slate-500/20 text-slate-300': r.Vote === 0,
            }"
          >
            {{ r.DisplayName }}
          </span>
        </span>
      </span>
    </div>

    <!-- Empty state -->
    <div
      v-if="prs.length === 0"
      class="px-5 py-12 text-center text-slate-500 text-sm"
    >
      No Pull Requests found.
    </div>
  </div>
</template>
