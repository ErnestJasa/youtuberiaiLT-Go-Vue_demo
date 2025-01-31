<script setup>
import { onMounted, reactive } from "vue";
import Card from "./Card.vue";

const state = reactive({
  channels: Array,
});

onMounted(async () => {
  try {
    const response = await fetch("http://localhost:8080/api/youtubers");
    const data = await response.json();
    state.channels = data;
  } catch (error) {
    console.log(error);
  }
});
</script>
<template>
  <section class="lg:col-span-2 flex flex-col gap-2 my-6">
    <Card
      v-for="channel in state.channels"
      :channel="channel"
      :key="channel.Id"
    />
  </section>
</template>
