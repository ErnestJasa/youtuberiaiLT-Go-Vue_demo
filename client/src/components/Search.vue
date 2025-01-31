<script setup>
import { reactive, ref } from "vue";
import CategoryList from "./CategoryList.vue";
import SortSelection from "./SortSelection.vue";
const showCategories = ref(false);
const categories = [
  "PODCAST",
  "IVAIRUS",
  "ZAIDIMAI",
  "ANIMACIJA",
  "PREKINIS ZENKLAS",
  "VLOG",
  "HUMORAS",
  "RECENZIJOS",
  "MINECRAFT",
  "EDUKACIJA",
  "MUZIKA",
  "MASINOS",
  "DOKUMENTIKA",
  "FORTNITE",
  "GTA",
  "COUNTER-STRIKE",
  "INSTITUCIJA",
  "ESE",
  "LEAGUE OF LEGENDS",
  "ZINIOS",
  "KULINARIJA",
];
const handleShowCategories = () => {
  console.log("clicked");
  showCategories.value = !showCategories.value;
  console.log(showCategories.value);
};
</script>
<template>
  <section
    className="md:ml-4 lg:col-span-2 sticky h-full bg-gray-950 max-h-[35%] top-0 pt-6 z-10"
    :style="{ maxHeight: 'calc(-2rem + 100vh)', overflowY: 'auto' }"
  >
    <div className=" border border-gray-700 px-2">
      <CategoryList :showCategories="showCategories" :categories="categories" />
      <button
        @click="handleShowCategories"
        class="w-full border rounded-xl py-0.5 text-lg text-center mt-2"
      >
        Kategorijos
      </button>
      <form
        onSubmit="{handleSearchSubmit}"
        className="w-full my-3 lg:flex gap-2 "
      >
        <div className="w-full flex relative">
          <input
            ref="{inputRef}"
            className="uppercase py-1 border border-gray-200 focus:bg-gray-900 hover:bg-gray-900  focus:outline-none focus:border focus:border-r hover:border-[#edd6b7] focus:border-[#e1a44f] w-full bg-inherit pl-3 text-xl font-normal"
            type="text"
            placeholder="Ieškoti kanalo..."
          />
          <button
            type="submit"
            className="flex items-center justify-center border md:w-[35%] lg:w-auto px-4 border-l-0 border-gray-200 hover:border-l hover:bg-gray-900 hover:border-[#e1a44f] active:border-[#e1a44f] active:bg-[#3f356aac]   hover:opacity-100"
          >
            <!-- <SearchSVG className="w-6 sm:w-5" /> -->
            Search
          </button>
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
