// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true }, 
  
  modules: ['@nuxtjs/tailwindcss', 'v-gsap-nuxt'],
  
  tailwindcss: {
    config: {
      theme: {
        extend: {
          fontFamily: {
            sans: ['Montserrat', 'system-ui'],
          },
        },
      },
    },
  },
})