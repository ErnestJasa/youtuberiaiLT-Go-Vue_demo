<script setup>
import CardList from "./components/CardList.vue";
import Search from "./components/Search.vue";
// some test changes
</script>

<template>
  <main
    class="w-[90%] lg:w-[75%] grid grid-cols-1 md:grid-cols-4 lg:grid-cols-5 gap-6 text-sky-100 mx-auto relative"
  >
    <Search />
    <CardList />
  </main>
</template>
