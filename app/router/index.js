<template>
  <div class="min-h-screen bg-[#f1f2f2] relative overflow-hidden font-mono">
    <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-center font-mono">
      <div class="relative inline-block">
        <p class="text-[#848484] text-2xl tracking-wide m-0 font-mono">
          <span 
            v-for="(char, index) in chars" 
            :key="index"
            :style="{ opacity: visibleChars[index] ? 1 : 0 }"
            class="inline-block"
          >
            {{ char === ' ' ? '\u00A0' : char }}
          </span>
        </p>
        <span 
          class="absolute top-0 w-3.5 h-[30px] mt-0.5 bg-[#ffe500] font-mono"
          :style="{ 
            left: `${handlePosition}px`,
            opacity: handleVisible ? 1 : 0,
            transition: 'opacity 0.4s'
          }"
        ></span>
      </div>
    </div>

    <svg 
      @click="replay"
      class="w-5 h-5 fixed right-4 bottom-4 fill-[#666] cursor-pointer hover:fill-[#888] transition-colors"
      version="1.1" 
      xmlns="http://www.w3.org/2000/svg" 
      viewBox="0 0 279.9 297.3"
    >
      <g>
        <path d="M269.4,162.6c-2.7,66.5-55.6,120.1-121.8,123.9c-77,4.4-141.3-60-136.8-136.9C14.7,81.7,71,27.8,140,27.8
          c1.8,0,3.5,0,5.3,0.1c0.3,0,0.5,0.2,0.5,0.5v15c0,1.5,1.6,2.4,2.9,1.7l35.9-20.7c1.3-0.7,1.3-2.6,0-3.3L148.6,0.3
          c-1.3-0.7-2.9,0.2-2.9,1.7v15c0,0.3-0.2,0.5-0.5,0.5c-1.7-0.1-3.5-0.1-5.2-0.1C63.3,17.3,1,78.9,0,155.4
          C-1,233.8,63.4,298.3,141.9,297.3c74.6-1,135.1-60.2,138-134.3c0.1-3-2.3-5.4-5.3-5.4l0,0C271.8,157.6,269.5,159.8,269.4,162.6z"/>
      </g>
    </svg>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'

const text = '404, page not found.'
const chars = text.split('')
const visibleChars = ref([])
const handlePosition = ref(0)
const handleVisible = ref(false)
const textWidth = ref(0)

let animationTimeout = null
let blinkInterval = null

const animateText = () => {
  visibleChars.value = new Array(chars.length).fill(false)
  handlePosition.value = 0
  handleVisible.value = true
  
  // Clear existing intervals
  if (blinkInterval) clearInterval(blinkInterval)
  
  // Start blinking
  blinkInterval = setInterval(() => {
    handleVisible.value = !handleVisible.value
  }, 400)
  
  // Animate characters with handle moving alongside
  chars.forEach((_, index) => {
    setTimeout(() => {
      visibleChars.value[index] = true
      
      // Move handle with each character
      nextTick(() => {
        const textElement = document.querySelector('p')
        if (textElement) {
          const spanElements = textElement.querySelectorAll('span')
          let width = 0
          for (let i = 0; i <= index; i++) {
            if (spanElements[i]) {
              width += spanElements[i].offsetWidth
            }
          }
          handlePosition.value = width
        }
      })
      
      // When last character is visible, stop blinking
      if (index === chars.length - 1) {
        setTimeout(() => {
          if (blinkInterval) clearInterval(blinkInterval)
          handleVisible.value = true
        }, 50)
      }
    }, index * 50 + 200)
  })
}

const replay = () => {
  if (animationTimeout) clearTimeout(animationTimeout)
  if (blinkInterval) clearInterval(blinkInterval)
  animateText()
}

onMounted(() => {
  nextTick(() => {
    animateText()
  })
})
</script>

<style scoped>

</style>
```