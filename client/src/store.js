import { reactive, ref } from "vue";

export const searchParams = ref({
  search: "",
  includeCategories: [],
  excludeCategories: [],
  sortOrder: "Default",
});

export const ApiURL = import.meta.env.VITE_API_URL;
