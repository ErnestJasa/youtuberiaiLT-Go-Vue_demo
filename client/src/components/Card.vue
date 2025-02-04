<script setup>
import { ref } from "vue";
import { searchParams } from "../store";
import { formatNumber } from "../helpers/NumberFormatting";
import YoutubeSVG from "../assets/SVGs/YoutubeSVG.vue";
import PersonSVG from "../assets/SVGs/PersonSVG.vue";
import ArrowRightSVG from "../assets/SVGs/ArrowRightSVG.vue";

defineProps({
  channel: Object,
});

const count = ref(0);
</script>

<template>
  <article class="flex w-full border hover:bg-gray-900 relative group/card">
    <a
      class="flex gap-1 w-full"
      :href="'https://youtube.com/' + channel.CustomUrl"
      target="_blank"
    >
      <img
        :src="channel.Thumbnail"
        :alt="channel.Title"
        referrerpolicy="no-referrer"
        class="w-16 object-cover"
      />
      <div class="w-full">
        <div class="flex w-[95%] flex-wrap">
          <div
            :class="`${
              searchParams.includeCategories.includes(category)
                ? 'bg-slate-800 text-gray-300'
                : 'text-gray-500'
            } text-xs uppercase tracking-wide mr-2 `"
            v-for="category in channel.Categories"
            :key="category"
          >
            {{ category }}
          </div>
        </div>
        <div>
          <h2 class="text-2xl uppercase font-novecento w-[75%] lg:w-full">
            {{ channel.Title }}
          </h2>
        </div>
        <div class="flex gap-2 text-xs font-bold w-full">
          <div class="text-primary flex" title="Prenumeratorių sk.">
            <PersonSVG class="w-4 fill-primary" />
            <p>{{ formatNumber(channel.SubscriberCount) }}</p>
          </div>
          <div title="Video sk." className="text-gray-400/80 flex">
            <YoutubeSVG class="w-4 fill-gray-700" />
            <p className="ml-[2px]">{{ formatNumber(channel.VideoCount) }}</p>
          </div>
        </div>
      </div>
      <div
        className="absolute right-3 top-[50%] translate-y-[-50%]  block lg:hidden group-hover/card:block"
      >
        <ArrowRightSVG class="w-6 fill-gray-300" />
      </div>
    </a>
  </article>
</template>
