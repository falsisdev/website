// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ["@nuxtjs/sanity", "@nuxt/icon"],
  css: ["~/assets/css/main.css"],
  sanity: {
    projectId: "5nqk6zg4",
    dataset: "production",
  },
  postcss: {
    plugins: {
      '@tailwindcss/postcss': {},
      autoprefixer: {},
    },
  },
  compatibilityDate: "2025-01-01",
});