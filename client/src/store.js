import { reactive, ref } from "vue";

export const searchParams = ref({
  search: "",
  includeCategories: [],
  excludeCategories: [],
  sortOrder: "",
});
