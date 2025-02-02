<script setup>
import { reactive, onMounted } from "vue";
import { searchParams } from "../store";
defineProps({
  handleCategoryClick: Function,
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
    :class="`${showCategories ? 'block' : 'hidden'} lg:block md:p-6 p-2`"
  >
    <div class="flex flex-wrap gap-2">
      <button
        :onclick="() => handleCategoryClick(category)"
        v-for="category in state.categories"
        :key="category"
        :class="`
           ${
             searchParams.includeCategories.includes(category)
               ? 'underline'
               : searchParams.excludeCategories.includes(category)
               ? 'line-through'
               : ''
           }
           cursor-pointer hover:bg-gray-800 border border-gray-800 inline-flex items-center text-lg px-2 py-0.5 uppercase`"
      >
        {{ category.replace(/\s+/g, "_") }}
      </button>
    </div>
  </div>
  <div v-else className="text-3xl text-center m-2">
    <h1>Kategorijos nerastos. Problema su serveriu</h1>
  </div>
</template>
