import { createVuetify } from 'vuetify'
import { aliases, mdi } from 'vuetify/iconsets/mdi'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

// TODO: Nuxt3 vuetify plugin 설정
export default defineNuxtPlugin((NuxtApp) => {
    const vuetify = createVuetify({
      components,
      directives,

      icons: { // 아이콘 설정
        defaultSet: 'mdi',
        aliases,
        sets: {
          mdi
        }
      }
    })
  
    NuxtApp.vueApp.use(vuetify)
  })