// https://nuxt.com/docs/api/configuration/nuxt-config
// https://nuxt.com/docs/api/configuration/nuxt-config
const ONE_DAY = 60 * 60 * 24 * 1000;
const ONE_WEEK = ONE_DAY * 7;

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
    firebaseadmin: {
      projectId: process.env.NUXT_FIREBASEADMIN_PROJECT_ID,
      clientEmail: process.env.NUXT_FIREBASEADMIN_CLIENT_EMAIL,
      privateKey: process.env.NUXT_FIREBASEADMIN_PRIVATE_KEY
    },
    
    public: {
      authCookieName: "__session",
      authCookieExpires: parseInt(ONE_WEEK.toString(), 10),
      firebase: {
        apiKey: process.env.NUXT_PUBLIC_FIREBASE_API_KEY,
        authDomain: process.env.NUXT_PUBLIC_FIREBASE_AUTH_DOMAIN,
        projectId: process.env.NUXT_PUBLIC_FIREBASE_PROJECT_ID,
        storageBucket: process.env.NUXT_PUBLIC_FIREBASE_STORAGE_BUCKET,
        messagingSenderId: process.env.NUXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID,
        appId: process.env.NUXT_PUBLIC_FIREBASE_APP_ID,
        measurementId: process.env.NUXT_PUBLIC_FIREBASE_MEASUREMENT_ID
      }

    }
  },
  devtools: { enabled: true }
})
