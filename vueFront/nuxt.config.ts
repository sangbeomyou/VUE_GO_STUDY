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
  devtools: { enabled: true }
})
