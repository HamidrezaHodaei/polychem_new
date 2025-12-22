<template>
  <div class="w-full min-h-screen bg-[#f1f2f2] font-sans">
    <h1 class="relative text-center pt-24 text-2xl md:text-3xl px-4 text-[#A8A8A8]">
      Scroll to watch POLYCHEM's story unfold
    </h1>

    <svg 
      id="svg-stage" 
      ref="svgStage"
      class="w-full mx-auto mt-[60vh] overflow-visible px-4"
      xmlns="http://www.w3.org/2000/svg" 
      viewBox="0 0 1400 1200"
      preserveAspectRatio="xMidYMid meet"
    >
      <!-- تعریف patterns برای عکس‌ها -->
      <defs>
        <pattern id="img01" x="0" y="0" width="1" height="1">
          <image href="/2015.jpg" x="-12" y="-12" width="250" height="130" preserveAspectRatio="xMidYMid slice"/>
        </pattern>
        <pattern id="img02" x="0" y="0" width="1" height="1">
          <image  href="/2025.jpeg" x="-10" y="2" width="150" height="170" preserveAspectRatio="xMidYMid slice"/>
        </pattern>
        <pattern id="img03" x="0" y="0" width="1" height="1">
          <image  href="/2025.jpeg" x="-10" y="2" width="150" height="170" preserveAspectRatio="xMidYMid slice"/>
        </pattern>
        <pattern id="img04" x="0" y="0" width="1" height="1">
          <image  href="/2025.jpeg" x="-10" y="2" width="150" height="170" preserveAspectRatio="xMidYMid slice"/>
        </pattern>
        <pattern id="img05" x="0" y="0" width="1" height="1">
          <image href="/2025.jpeg" x="-10" y="2" width="150" height="170" preserveAspectRatio="xMidYMid slice"/>
        </pattern>
      </defs>

      <!-- خطوط افقی -->
      <path class="line01 line" d="M 0 200 1400 200" fill="none" stroke="#A8A8A8" stroke-width="2"></path>
      <path class="line02 line" d="M 0 400 1400 400" fill="none" stroke="#A8A8A8" stroke-width="2"></path>
      <path class="line03 line" d="M 0 600 1400 600" fill="none" stroke="#A8A8A8" stroke-width="2"></path>
      <path class="line04 line" d="M 0 800 1400 800" fill="none" stroke="#A8A8A8" stroke-width="2"></path>
      <path class="line05 line" d="M 0 1000 1400 1000" fill="none" stroke="#A8A8A8" stroke-width="2"></path>
      
      <!-- متن‌های سال سمت چپ -->
      <text class="text01" x="30" y="190" fill="#A8A8A8" font-size="15" style="visibility: hidden">2015</text>
      <text class="text02" x="30" y="390" fill="#A8A8A8" font-size="15" style="visibility: hidden">2019</text>
      <text class="text03" x="30" y="590" fill="#A8A8A8" font-size="15" style="visibility: hidden">2020</text>
      <text class="text04" x="30" y="790" fill="#A8A8A8" font-size="15" style="visibility: hidden">2021</text>
      <text class="text05" x="30" y="990" fill="#A8A8A8" font-size="15" style="visibility: hidden">2025</text>
      
      <!-- متن‌های توضیحات سمت راست -->
      <text class="desc01" x="1350" y="175" fill="#A8A8A8" font-size="15" text-anchor="end" style="visibility: hidden">Start of factory construction</text>
      <text class="desc02" x="1350" y="375" fill="#A8A8A8" font-size="15" text-anchor="end" style="visibility: hidden">Completion of the project and official commissioning</text>
      <text class="desc03" x="1350" y="575" fill="#A8A8A8" font-size="15" text-anchor="end" style="visibility: hidden">Start of production</text>
      <text class="desc04" x="1350" y="775" fill="#A8A8A8" font-size="15" text-anchor="end" style="visibility: hidden">Entry into international markets and beginning of exports</text>
      <text class="desc05" x="1350" y="975" fill="#A8A8A8" font-size="15" text-anchor="end" style="visibility: hidden">Current status and ongoing development</text>

      <!-- مسیر اصلی که از تمام خطوط عبور می‌کند -->
      <path 
        class="theLine" 
        d="M 100,0 Q 900 100 700 200 Q 500 300 700 400 Q 900 500 700 600 Q 500 700 700 800 Q 900 900 700 1000 Q 500 1100 700 1200"
        fill="none" 
        stroke="#A8A8A8" 
        stroke-width="5"
      />
      
      <!-- توپ متحرک - قبل از توپ‌های ثابت -->
      <circle class="ball ballMoving" r="8" cx="100" cy="0" fill="#ffd700" style="visibility: hidden"></circle>
      
      <!-- توپ‌های ثابت روی خطوط با عکس - اندازه بزرگتر و بعد از توپ متحرک -->
      <circle class="ball ball01" r="60" cx="700" cy="200" fill="url(#img01)" stroke="#ffd700" stroke-width="3" style="visibility: hidden"></circle>
      <circle class="ball ball02" r="60" cx="700" cy="400" fill="url(#img02)" stroke="#ffd700" stroke-width="3" style="visibility: hidden"></circle>
      <circle class="ball ball03" r="60" cx="700" cy="600" fill="url(#img03)" stroke="#ffd700" stroke-width="3" style="visibility: hidden"></circle>
      <circle class="ball ball04" r="60" cx="700" cy="800" fill="url(#img04)" stroke="#ffd700" stroke-width="3" style="visibility: hidden"></circle>
      <circle class="ball ball05" r="65" cx="700" cy="1000" fill="url(#img05)" stroke="#ffd700" stroke-width="3" style="visibility: hidden"></circle>
    </svg>

    <div class="flex justify-center mt-10">
      <img src="/english logo W1.png" alt="English Logo" class="w-56 md:w-72 lg:w-96 h-auto" />
    </div>
    <div class="pb-10"><br></div>
    <TextAboutus />
    <div class="pb-20"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import gsap from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'
import { DrawSVGPlugin } from 'gsap/DrawSVGPlugin'
import { MotionPathPlugin } from 'gsap/MotionPathPlugin'
import TextAboutus from './Text Aboutus.vue'

gsap.registerPlugin(ScrollTrigger, DrawSVGPlugin, MotionPathPlugin)

const svgStage = ref(null)
let mainTimeline = null
let pulsesTimeline = null

onMounted(() => {
  gsap.defaults({ ease: "none" })

  // انیمیشن پالس برای توپ‌ها هنگام برخورد - با اندازه بزرگتر
  pulsesTimeline = gsap.timeline({
    defaults: {
      duration: 0.05,
      autoAlpha: 1,
      scale: 1.5,
      transformOrigin: "center",
      ease: "elastic(2.5, 1)"
    }
  })
    .to(".ball01, .text01, .desc01", {}, 0.300)
    .to(".ball02, .text02, .desc02", {}, 0.455)
    .to(".ball03, .text03, .desc03", {}, 0.599)
    .to(".ball04, .text04, .desc04", {}, 0.695)
    .to(".ball05, .text05, .desc05", {}, 0.865)

  mainTimeline = gsap.timeline({
    defaults: { duration: 1 },
    scrollTrigger: {
      trigger: "#svg-stage",
      scrub: true,
      start: "top center",
      end: "bottom center"
    }
  })
    .to(".ballMoving", { duration: 0.01, autoAlpha: 1 })
    .from(".theLine", { drawSVG: 0 }, 0)
    .to(".ballMoving", {
      motionPath: {
        path: ".theLine",
        align: ".theLine",
        alignOrigin: [0.5, 0.5]
      }
    }, 0)
    .add(pulsesTimeline, 0)
})

onUnmounted(() => {
  if (mainTimeline) mainTimeline.kill()
  if (pulsesTimeline) pulsesTimeline.kill()
  ScrollTrigger.getAll().forEach(trigger => trigger.kill())
})
</script>

<style scoped>
/* استایل‌های اضافی در صورت نیاز */
</style>