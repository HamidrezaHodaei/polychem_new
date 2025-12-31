<template>
  <Teleport to="body">
    <div class="loading-container" v-if="!contentLoaded">

      <!-- Circular Reveal -->
      <div
        v-if="showReveal"
        class="circular-reveal"
        :class="{ 'reveal-active': revealActive }"
      ></div>

      <!-- SVG Loader (جایگزین توپ) -->
      <div
        ref="svgWrapper"
        class="svg-wrapper"
        :class="{
          bounce: svgBounce,
          impact: svgImpact
        }"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 170.08 170.08"
          class="loader-svg"
        >
          <path
            ref="yellowCircle"
            fill="#ffcc05"
            d="M180.23,93.77a85,85,0,1,1-85-85,85,85,0,0,1,85,85"
            transform="translate(-10.16 -8.73)"
          />

          <circle
            ref="whiteCircle"
            fill="none"
            stroke="#fff"
            stroke-width="1.5"
            cx="85.04"
            cy="85.04"
            r="78.1"
          />

          <path v-for="(p, i) in paths" :key="i" ref="pathEls"
            fill="none" stroke="#fff" stroke-width="1.5"
            :d="p"
            transform="translate(-10.16 -8.73)"
          />
        </svg>
      </div>

    </div>
  </Teleport>
</template>
<script setup>
import { ref, onMounted, nextTick } from 'vue'
import gsap from 'gsap'

const yellowCircle = ref(null)
const whiteCircle = ref(null)
const pathEls = ref([])
const svgWrapper = ref(null)

const svgBounce = ref(false)
const svgImpact = ref(false)
const showReveal = ref(false)
const revealActive = ref(false)
const contentLoaded = ref(false)

const paths = [
  "M51.71,60.3l11.8,3.61...",
  "M89.22,81.67C114.67...",
  "M90.15,65.89a21.63...",
  "M97.34,68.25C92.28...",
  "M102.67,75.17c-8...",
  "M87.91,85c-4.48...",
  "M89,85.17c-5.9...",
  "M89,85.48c-5.46...",
  "M88.47,86.2c-3.54..."
]

onMounted(async () => {
  await nextTick()

  const drawPaths = [whiteCircle.value, ...pathEls.value]

  drawPaths.forEach(p => {
    const len = p.getTotalLength()
    p.style.strokeDasharray = len
    p.style.strokeDashoffset = len
  })

  const tl = gsap.timeline()

  // ظاهر شدن دایره زرد
  tl.fromTo(
    yellowCircle.value,
    { scale: 0, opacity: 0, transformOrigin: '50% 50%' },
    { scale: 1, opacity: 1, duration: 0.6 }
  )

  // رسم path ها
  drawPaths.forEach(p => {
    tl.to(p, {
      strokeDashoffset: 0,
      duration: 0.18,
      ease: 'none'
    })
  })

  // وقتی همه PATH ها تموم شد 👇
  tl.call(startSecondaryAnimation)
})

function startSecondaryAnimation () {
  svgBounce.value = true

  setTimeout(() => {
    svgImpact.value = true
    showReveal.value = true
  }, 1500)

  setTimeout(() => {
    revealActive.value = true
  }, 1600)

  setTimeout(() => {
    contentLoaded.value = true
  }, 2800)
}
</script>
<style scoped></style>