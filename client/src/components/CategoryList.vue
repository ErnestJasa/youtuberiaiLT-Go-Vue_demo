<script setup>
import { reactive, onMounted, computed, ref } from "vue";
import { ApiURL, searchParams } from "../store";
import Loader from "./Loader.vue";
defineProps({
  handleCategoryClick: Function,
  showCategories: Boolean,
});
const state = reactive({
  categories: [],
});
const loading = ref(false);
const categorySearchInput = ref("");

const filteredCategories = computed(() => {
  return state.categories.filter((category) =>
    category.toLowerCase().includes(categorySearchInput.value.toLowerCase())
  );
});
onMounted(async () => {
  try {
    loading.value = true;
    const response = await fetch(`${ApiURL}api/categories`);
    const data = await response.json();
    state.categories = data;
    loading.value = false;
  } catch (error) {
    console.log(error);
    loading.value = false;
  }
});
</script>
<template>
  <Loader v-if="loading" />
  <div class="w-full flex justify-center mt-2">
    <input
      :class="`${
        showCategories ? 'block' : 'hidden'
      } md:block bg-inherit border border-gray-700 rounded px-3 py-1 focus:bg-gray-900 hover:bg-gray-900 focus:outline-none focus:border focus:border-r hover:border-gray-600 focus:border-[#e1a44f]`"
      v-model="categorySearchInput"
      type="text"
      placeholder="Ieškoti kategorijos"
    />
  </div>
  <div
    v-if="filteredCategories && filteredCategories.length"
    :class="`${
      showCategories ? 'block' : 'hidden'
    } md:block md:px-1 lg:px-3 md:pb-6 p-2`"
  >
    <div class="flex flex-wrap gap-2">
      <button
        :onclick="() => handleCategoryClick(category)"
        v-for="category in filteredCategories"
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
  <div
    v-else-if="categorySearchInput !== '' && !loading"
    class="text-3xl text-center m-2"
  >
    <h1>Kategorijos nerastos.</h1>
  </div>
  <div v-else-if="!loading" class="text-3xl text-center m-2">
    <h1>Kategorijos nerastos. Problema su serveriu.</h1>
  </div>
</template>
