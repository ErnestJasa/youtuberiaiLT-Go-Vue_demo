<script setup>
import { reactive, watchEffect, ref, watch } from "vue";
import Card from "./Card.vue";
import { searchParams } from "../store";
const state = reactive({
  channels: Array,
});
const timer = ref(null);
const loading = ref(false);

watchEffect(async () => {
  const sortOrder = searchParams.value.sortOrder;
  const searchTerm = searchParams.value.search;
  const includeCategories =
    searchParams.value.includeCategories.join("&includeCategory=");
  const excludeCategories =
    searchParams.value.excludeCategories.join("&excludeCategory=");
  if (timer.value) {
    clearTimeout(timer.value);
    timer.value = null;
  }
  timer.value = setTimeout(async () => {
    try {
      loading.value = true;
      const response = await fetch(
        `http://localhost:8080/api/youtubers?channelHandle=${searchTerm}&sortOrder=${sortOrder}&includeCategory=${includeCategories}&excludeCategory=${excludeCategories}`
      );
      const data = await response.json();
      state.channels = data;
      loading.value = false;
    } catch (error) {
      loading.value = false;
      console.log(error);
    }
  }, 400);
});
</script>
<template>
  <section class="lg:col-span-2 flex flex-col gap-2 my-6">
    <h2 v-if="loading">Loading...</h2>
    <Card
      v-else
      v-for="channel in state.channels"
      :channel="channel"
      :key="channel.Id"
    />
  </section>
</template>
