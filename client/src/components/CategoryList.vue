<script setup>
defineProps({
  categories: Array,
  showCategories: Boolean,
});
// console.log(showCategories);
</script>
<!-- showCategories ? 'block' : 'hidden' +  -->
<template>
  <div
    v-if="categories.length && categories"
    :className="`${showCategories ? 'visible' : 'hidden'} md:p-6 p-2`"
  >
    <div className="flex flex-wrap gap-2 ">
      <button
        v-for="category in categories"
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
