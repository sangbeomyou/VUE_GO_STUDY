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
  runtimeConfig: {
    public: {
      apiKey: process.env.apiKey,
      authDomain: process.env.authDomain,
      projectId: process.env.projectId,
      storageBucket: process.env.storageBucket,
      messagingSenderId: process.env.messagingSenderId,
      appId: process.env.appId,
      measurementId: process.env.measurementId
    }
  },
  devtools: { enabled: true }
})
