<script setup>
import { reactive, ref, watch } from "vue";
import CategoryList from "./CategoryList.vue";
import SortSelection from "./SortSelection.vue";
import { searchParams } from "../store";
import { removeFromArray } from "../helpers/ArrayHelpers";
const showCategories = ref(false);

const handleShowCategories = () => {
  showCategories.value = !showCategories.value;
};

function handleCategorySearch(categoryName) {
  // add to includes array
  if (
    !searchParams.value.includeCategories.includes(categoryName) &&
    !searchParams.value.excludeCategories.includes(categoryName)
  ) {
    searchParams.value.includeCategories = [
      ...searchParams.value.includeCategories,
      categoryName,
    ];
    return;
  }

  // remove from includes and add to excludes array
  if (searchParams.value.includeCategories.includes(categoryName)) {
    searchParams.value.includeCategories = removeFromArray(
      searchParams.value.includeCategories,
      categoryName
    );
    searchParams.value.excludeCategories = [
      ...searchParams.value.excludeCategories,
      categoryName,
    ];
    return;
  }

  // remove from excludes array
  if (searchParams.value.excludeCategories.includes(categoryName)) {
    searchParams.value.excludeCategories = removeFromArray(
      searchParams.value.excludeCategories,
      categoryName
    );
    return;
  }
}
</script>
<template>
  <section
    className="md:ml-4 lg:col-span-2 sticky h-full bg-gray-950 max-h-[35%] top-0 pt-6 z-10"
    :style="{ maxHeight: 'calc(-2rem + 100vh)', overflowY: 'auto' }"
  >
    <div className=" border border-gray-700 px-2">
      <CategoryList
        :handleCategoryClick="handleCategorySearch"
        :showCategories="showCategories"
      />
      <button
        @click="handleShowCategories"
        class="w-full border rounded-xl py-0.5 text-lg text-center mt-2 lg:hidden"
      >
        Kategorijos
      </button>
      <form
        @submit.prevent="handleSearchSubmit"
        className="w-full my-3 lg:flex gap-2 "
      >
        <div className="w-full flex relative">
          <input
            v-model="searchParams.search"
            className="uppercase py-1 border border-gray-200 focus:bg-gray-900 hover:bg-gray-900  focus:outline-none focus:border focus:border-r hover:border-[#edd6b7] focus:border-[#e1a44f] w-full bg-inherit pl-3 text-xl font-normal"
            type="text"
            placeholder="Ieškoti kanalo..."
          />
          <button
            onClick="{clearInput}"
            type="button"
            className="absolute right-16 md:right-[100px] lg:right-16 top-2 opacity-50 hover:opacity-100 active:opacity-100"
          >
            <!-- <XInCircleSVG /> -->
          </button>
        </div>
        <div className="w-full lg:w-[60%] my-2 lg:-my-0">
          <SortSelection
            handleSortingChange="{handleSortingChange}"
            sortBy="{searchQuery.sortOrder}"
          />
        </div>
      </form>
    </div>
  </section>
</template>
