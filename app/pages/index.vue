<template>
  <Landpage v-if="showLanding" />
  <div v-else>
    <Navbar/>
    <Hero2/>
    <div class="mt-8 slider">
      <Slider/>
      <Text/>
      <Cards/>
      <global/>
      <infinityslider/>
    </div>
    <footer1/> 
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Landpage from "~/components/landpage.vue"
import Navbar from "~/components/navbar.vue"
import Slider from "~/components/Slider.vue"
import Hero from "~/components/Hero2.vue"
import Cards from "~/components/cards.vue"
import Text from "~/components/Text.vue"
import global from "~/components/world.vue"
import infinityslider from "~/components/InfiniteSlider.vue"
import footer1 from "~/components/footer.vue"

const showLanding = ref(false)
const router = useRouter()

onMounted(() => {
  // Check if user has already seen the landing page in this session
  const hasSeenLanding = sessionStorage.getItem('hasSeenLanding')
  
  if (!hasSeenLanding) {
    showLanding.value = true
  }
  
  // پاک کردن موقعیت اسکرول قبلی هنگام بارگذاری اول صفحه
  if (!sessionStorage.getItem('hasSeenLanding')) {
    sessionStorage.removeItem('homeScrollPosition')
  }
})

// Listen for landing page completion event
if (process.client) {
  window.addEventListener('landing-complete', () => {
    sessionStorage.setItem('hasSeenLanding', 'true')
    showLanding.value = false
  })
  
  // پاک کردن موقعیت اسکرول هنگام رفتن به صفحات دیگر
  router.beforeEach((to, from, next) => {
    if (from.path === '/' && to.path !== '/') {
      sessionStorage.removeItem('homeScrollPosition')
    }
    next()
  })
}
</script>