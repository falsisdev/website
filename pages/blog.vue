<script setup>
import { Swiper, SwiperSlide } from 'swiper/vue'
import 'swiper/css'
import imageUrlBuilder from '@sanity/image-url'

const builder = imageUrlBuilder(useSanity().config);

function urlFor(source) {
  return builder.image(source).url()
}

const query = groq`*[
  _type == "post" && defined(slug.current)
]|order(publishedAt desc)[0...12]{
  _id,
  publishedAt,
  slug,
  title,
  cardCover,
  author->{
    name,
    image
  },
  categories[]->{
    title,
    slug
  }
}`

const { data, refresh, pending, error } = useSanityQuery(query, {
  tag: "blog",
  initialData: [],
})
</script>

<template>
  <main v-if="data" class="container mx-auto py-8">
    <h1 class="text-3xl font-bold mb-8 text-center">Blog</h1>
    <h1 class="text-3xl font-bold mb-8">Latest Posts</h1>
    <Swiper
      :slides-per-view="4"
      space-between="16"
      grab-cursor="true"
      class="w-full"
    >
      <SwiperSlide
        v-for="post in data"
        :key="post._id"
        class="w-72 flex-shrink-0"
      >
        <div class="card hover:shadow-xl transition-shadow duration-200 w-full">
          <figure v-if="post.cardCover">
            <img :src="urlFor(post.cardCover)" alt="cover" class="h-48 w-full object-cover" />
          </figure>
          <div class="card-body">
            <span class="card-title flex flex-row">
              <h2>{{ post.title }}</h2>
              <span class="grow"/>
              <span class="flex flex-row">
                <img
                  v-if="post.author?.image"
                  :src="urlFor(post.author.image.asset)"
                  alt="author"
                  class="w-4 h-4 rounded-full object-cover mx-1 mt-[1px]"
                />
                <span class="text-sm opacity-75">{{ post.author?.name }}</span>
              </span>
            </span>
            <span class="text-xs text-gray-400 mb-2 -mt-2">
              {{ new Date(post.publishedAt).toLocaleString("tr-TR", {
                year: "numeric",
                month: "long",
                day: "numeric",
                hour: "2-digit",
                minute: "2-digit"
              }) }}
            </span>
            <div class="flex flex-wrap gap-1 mb-2">
              <NuxtLink
                v-for="cat in post.categories"
                :key="cat._id"
                :to="`/blog/category/${cat.title.toLowerCase()}`"
                class="badge badge-soft badge-accent badge-xs transition-all duration-200 cursor-pointer"
              >
                {{ cat.title }}
              </NuxtLink>
            </div>
            <div class="card-actions justify-end">
              <NuxtLink
                :to="`/blog/${post.slug.current}`"
                class="btn btn-soft btn-primary btn-sm"
              >
                Oku
              </NuxtLink>
            </div>
          </div>
        </div>
      </SwiperSlide>
    </Swiper>
  </main>
</template>