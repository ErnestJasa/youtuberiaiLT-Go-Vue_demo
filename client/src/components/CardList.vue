<script setup>
import { reactive, watchEffect, ref, watch } from "vue";
import Card from "./Card.vue";
import { searchParams } from "../store";
import Loader from "./Loader.vue";
const state = reactive({
  channels: [],
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

  loading.value = true;
  timer.value = setTimeout(async () => {
    try {
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
    <Loader v-if="loading" />
    <Card
      v-else-if="state.channels && state.channels.length"
      v-for="channel in state.channels"
      :channel="channel"
      :key="channel.Id"
    />
    <div v-else-if="!loading" class="text-3xl text-center m-2">
      <h1>Kanalų nerasta.</h1>
    </div>
  </section>
</template>
