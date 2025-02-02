<script setup>
import { searchParams } from "../store";
</script>

<template>
  <select
    v-model="searchParams.sortOrder"
    className="uppercase text-lg focus:outline-none bg-inherit focus:bg-gray-900 w-full hover:bg-gray-900 flex items-center justify-between border border-input px-3 py-2 "
    id="sort-select"
  >
    <option selected class="bg-gray-950" value="Default">Rūšiuoti</option>
    <option class="bg-gray-950" value="ByHandleAscending">
      Pavadinimas A-Z
    </option>
    <option class="bg-gray-950" value="ByHandleDescending">
      Pavadinimas Z-A
    </option>
    <option class="bg-gray-950" value="BySubCountAscending">
      Prenumeratų sk. min-max
    </option>
    <option class="bg-gray-950" value="BySubCountDescending">
      Prenumeratų sk. max-min
    </option>
  </select>
</template>
