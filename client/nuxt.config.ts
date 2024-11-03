
// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  css: ['~/assets/css/global.css'],
  typescript: {
    strict: true
  },
  modules: [
    'nuxt-monaco-editor'
  ],
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true }
})
