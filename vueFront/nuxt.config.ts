// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  css: [
    'vuetify/lib/styles/main.sass',
    '~/resource/styles/common.css',
 ],
 build: {
    transpile: ['vuetify']
 },
 vite: {
    define: {
       'process.env.DEBUG': true,
    }
 },  
 modules: [
   '@pinia/nuxt',
 ],
 imports: {
   dirs: ['./stores']
 },
 pinia: {
   autoImports: ['defineStore', 'acceptHMRUpdate']
 },
  devtools: { enabled: true }
})
