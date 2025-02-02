<script setup>
import { reactive, onMounted } from "vue";
defineProps({
  // categories: Array,
  showCategories: Boolean,
});
const state = reactive({
  categories: Array,
});

onMounted(async () => {
  try {
    const response = await fetch("http://localhost:8080/api/categories");
    const data = await response.json();
    state.categories = data;
  } catch (error) {
    console.log(error);
  }
});
</script>
<template>
  <div
    v-if="state.categories.length && state.categories"
    :className="`${showCategories ? 'block' : 'hidden'} lg:block md:p-6 p-2`"
  >
    <div className="flex flex-wrap gap-2 ">
      <button
        v-for="category in state.categories"
        :key="{ category }"
        className="cursor-pointer hover:bg-gray-800 border border-gray-800 inline-flex items-center text-lg px-2 py-0.5 uppercase"
      >
        {{ category.replace(/\s+/g, "_") }}
      </button>
    </div>
  </div>
  <div v-else className="text-3xl text-center m-2">
    <h1>Kategorijos nerastos. Problema su serveriu</h1>
  </div>
</template>
