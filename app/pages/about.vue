<template>
  <div>
    <navbar />
    
    <ClientOnly>
      <!-- Aboutus Hero Section - Full Half Screen -->
      <div class="h-screen flex items-center justify-center bg-[#f1f2f2] about-hero">
        <div class="w-full">
          <Aboutus />
        </div>
      </div>

      <!-- Stats Section -->
      <div class="min-h-screen bg-[#f1f2f2] py-16 px-4 stats-section-mobile">
        <div class="max-w-7xl mx-auto">
          <!-- Stats Grid 2x2 -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6 max-w-5xl mx-auto">
            <!-- Card 1 -->
            <div 
              class="bg-white rounded-lg p-12 shadow-sm relative overflow-hidden group hover:shadow-md transition-shadow duration-300 card-interactive"
              @mousemove="handleCardMouseMove"
              @mouseleave="handleCardMouseLeave"
              :ref="el => { if (el) cardRefs[0] = el }"
            >
              <div class="absolute inset-0 opacity-5 card-bg-characters" :ref="el => { if (el) bgRefs[0] = el }">
                {{ bgTexts[0] }}
              </div>
              <div class="absolute top-4 right-4 text-[#848484] opacity-20">
                <svg class="w-12 h-12" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M13 7h-2v4H7v2h4v4h2v-4h4v-2h-4V7zm-1-5C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/>
                </svg>
              </div>
              <div class="text-6xl font-light text-[#848484] mb-4 relative z-10">
                10<span class="text-4xl align-super">+</span>
              </div>
              <p class="text-[#848484] text-lg relative z-10">Cumulative Years Of Expertise</p>
            </div>

            <!-- Card 2 -->
            <div 
              class="bg-[#E8E6E6] rounded-lg p-12 shadow-sm relative overflow-hidden group hover:shadow-md transition-shadow duration-300 card-interactive"
              @mousemove="handleCardMouseMove"
              @mouseleave="handleCardMouseLeave"
              :ref="el => { if (el) cardRefs[1] = el }"
            >
              <div class="absolute inset-0 opacity-5 card-bg-characters" :ref="el => { if (el) bgRefs[1] = el }">
                {{ bgTexts[1] }}
              </div>
            </div>

            <!-- Card 3 -->
            <div 
              class="bg-[#E8E6E6] rounded-lg p-12 shadow-sm relative overflow-hidden group hover:shadow-md transition-shadow duration-300 card-interactive"
              @mousemove="handleCardMouseMove"
              @mouseleave="handleCardMouseLeave"
              :ref="el => { if (el) cardRefs[2] = el }"
            >
              <div class="absolute inset-0 opacity-5 card-bg-characters" :ref="el => { if (el) bgRefs[2] = el }">
                {{ bgTexts[2] }}
              </div>
            </div>

            <!-- Card 4 -->
            <div 
              class="bg-white rounded-lg p-12 shadow-sm relative overflow-hidden group hover:shadow-md transition-shadow duration-300 card-interactive"
              @mousemove="handleCardMouseMove"
              @mouseleave="handleCardMouseLeave"
              :ref="el => { if (el) cardRefs[3] = el }"
            >
              <div class="absolute inset-0 opacity-5 card-bg-characters" :ref="el => { if (el) bgRefs[3] = el }">
                {{ bgTexts[3] }}
              </div>
              <div class="absolute top-4 right-4 text-[#848484] opacity-20">
                <svg class="w-12 h-12" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M16 6l2.29 2.29-4.88 4.88-4-4L2 16.59 3.41 18l6-6 4 4 6.3-6.29L22 12V6z"/>
                </svg>
              </div>
              <div class="text-6xl font-light text-[#848484] mb-4 relative z-10">
                4.3<span class="text-4xl align-super">%</span>
              </div>
              <p class="text-[#848484] text-lg relative z-10">organic growth</p>
            </div>
          </div>

          <!-- Bottom Grid 3x3 -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-5xl mx-auto bottom-grid-mobile">
            <!-- Cards 5-10 -->
            <div 
              v-for="(card, idx) in bottomCards"
              :key="idx + 4"
              :class="[
                'rounded-lg p-12 shadow-sm relative overflow-hidden group hover:shadow-md transition-shadow duration-300 card-interactive',
                (card.value && ['24','50','84','32'].includes(card.value)) ? 'bg-white' : 'bg-[#E8E6E6]'
              ]"
              @mousemove="handleCardMouseMove"
              @mouseleave="handleCardMouseLeave"
              :ref="el => { if (el) cardRefs[idx + 4] = el }"
            >
              <div class="absolute inset-0 opacity-5 card-bg-characters" :ref="el => { if (el) bgRefs[idx + 4] = el }">
                {{ bgTexts[idx + 4] }}
              </div>
              <div v-if="card.icon" class="absolute top-4 right-4 text-[#848484] opacity-20">
                <svg class="w-12 h-12" fill="currentColor" viewBox="0 0 24 24">
                  <path :d="card.icon"/>
                </svg>
              </div>
              <div class="text-6xl font-light text-[#848484] mb-4 relative z-10">
                {{ card.value }}<span class="text-4xl align-super">{{ card.unit }}</span>
              </div>
              <p class="text-[#848484] text-lg relative z-10">{{ card.label }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Vision Section -->
      <div class="py-20 px-4 bg-[#f1f2f2] vision-section-mobile">
        <div class="max-w-7xl mx-auto">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 items-center">
            <!-- Text Content -->
            <div>
              <h2 class="text-5xl md:text-6xl font-light text-[#848484] mb-12">Vision</h2>
              
              <p class="text-[#848484] text-lg mb-8 leading-relaxed">
                At POLYCHEM, we envision a future where advanced polymer materials inspire better products, smarter manufacturing, and a more sustainable plastics industry.
              </p>
              
              <p class="text-[#848484] text-lg mb-8 leading-relaxed">
                Our goal is to become the most trusted partner for high-performance compounds and masterbatch solutions. A company known for its innovation, integrity, and commitment to excellence.
              </p>
            </div>
            
            <!-- Image -->
            <div class="relative h-[600px] rounded-2xl overflow-hidden">
              <div class="absolute inset-0 bg-gradient-to-t from-yellow-500/30 to-transparent z-10"></div>
              <img 
                src="/vision.jpg" 
                alt="Factory" 
                class="w-full h-full object-cover grayscale"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- add scroll component right after Vision -->
      <scroll class="scroll-section-mobile" />
    </ClientOnly>

    <footer2 class="footer-section-mobile" />
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const cardRefs = ref([])
const bgRefs = ref([])
const bgTexts = ref(Array(10).fill(''))

const bottomCards = [
  {
    value: '32',
    unit: '+',
    label: 'Number of clients',
    icon: 'M13 7h-2v4H7v2h4v4h2v-4h4v-2h-4V7zm-1-5C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z'
  },
  {
    value: '',
    unit: '',
    label: '',
    icon: null
  },
  {
    value: '84',
    unit: '%',
    label: 'Client Satisfaction',
    icon: 'M16 6l2.29 2.29-4.88 4.88-4-4L2 16.59 3.41 18l6-6 4 4 6.3-6.29L22 12V6z'
  },
  {
    value: '',
    unit: '',
    label: '',
    icon: null
  },
  {
    value: '24',
    unit: '/7',
    label: 'Support Available',
    icon: 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z'
  },
  {
    value: '50',
    unit: '+',
    label: 'Number of Employees',
    icon: 'M16 11c1.66 0 2.99-1.34 2.99-3S17.66 5 16 5c-1.66 0-3 1.34-3 3s1.34 3 3 3zm-8 0c1.66 0 2.99-1.34 2.99-3S9.66 5 8 5C6.34 5 5 6.34 5 8s1.34 3 3 3zm0 2c-2.33 0-7 1.17-7 3.5V19h14v-2.5c0-2.33-4.67-3.5-7-3.5zm8 0c-.29 0-.62.02-.97.05 1.16.84 1.97 1.97 1.97 3.45V19h6v-2.5c0-2.33-4.67-3.5-7-3.5z'
  }
]

const randomDotsString = (length = 95000) => {
  const pair = '· '
  const repeats = Math.ceil(length / pair.length)
  return pair.repeat(repeats).substring(0, length)
}

const handleCardMouseMove = (e) => {
  if (!process.client) return
  
  const card = e.currentTarget
  const cardIndex = cardRefs.value.indexOf(card)
  
  if (cardIndex === -1 || !bgRefs.value[cardIndex]) return

  const rect = card.getBoundingClientRect()
  const x = e.clientX - rect.left
  const y = e.clientY - rect.top

  const bgEl = bgRefs.value[cardIndex]
  bgEl.style.setProperty('--x', `${x}px`)
  bgEl.style.setProperty('--y', `${y}px`)

  const cx = rect.width / 2
  const cy = rect.height / 2
  const maxOffset = 15
  const tx = ((x - cx) / cx) * maxOffset
  const ty = ((y - cy) / cy) * maxOffset

  bgEl.style.setProperty('--tx', `${tx}px`)
  bgEl.style.setProperty('--ty', `${ty}px`)
  bgEl.style.setProperty('--scale', '1.05')

  bgTexts.value[cardIndex] = randomDotsString(95000)
}

const handleCardMouseLeave = (e) => {
  if (!process.client) return
  
  const card = e.currentTarget
  const cardIndex = cardRefs.value.indexOf(card)

  if (cardIndex === -1 || !bgRefs.value[cardIndex]) return

  const bgEl = bgRefs.value[cardIndex]
  bgEl.style.setProperty('--tx', '0px')
  bgEl.style.setProperty('--ty', '0px')
  bgEl.style.setProperty('--scale', '1')
}

const updateNavbarHeightVar = () => {
  if (!process.client) return
  
  const navbarEl = document.querySelector('nav') || document.querySelector('.navbar') || document.querySelector('#navbar')
  let navHeight = navbarEl ? navbarEl.offsetHeight : 0
  if (!navHeight || navHeight < 40) navHeight = 200
  document.documentElement.style.setProperty('--navbar-height', `${navHeight}px`)
}

onMounted(() => {
  bgTexts.value = Array(10).fill(randomDotsString(95000))
  updateNavbarHeightVar()
  if (process.client) {
    window.addEventListener('resize', updateNavbarHeightVar)
  }
})

onUnmounted(() => {
  if (process.client) {
    window.removeEventListener('resize', updateNavbarHeightVar)
  }
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;600;700&display=swap');

.min-h-screen {
  font-family: 'Montserrat', sans-serif;
}

.about-hero {
  box-sizing: border-box;
  --hero-extra-offset: 24px;
}

@media (max-width: 768px) {
  .about-hero {
    padding-top: calc(var(--navbar-height, 200px) + var(--hero-extra-offset));
    min-height: calc(100vh - (var(--navbar-height, 200px) + var(--hero-extra-offset)));
  }

  .about-hero > .w-full > * {
    margin-top: calc(var(--navbar-height, 200px) + var(--hero-extra-offset));
    position: relative;
    z-index: 1;
  }

  .stats-section-mobile {
    margin-top: calc(var(--navbar-height, 200px) + 48px);
  }
  .bottom-grid-mobile {
    margin-top: 32px;
  }
  .vision-section-mobile {
    margin-top: 32px;
  }
  .scroll-section-mobile {
    margin-top: 32px;
    display: block;
    width: 100%;
  }
  .footer-section-mobile {
    margin-top: 32px;
    display: block;
    width: 100%;
  }
}

.card-interactive {
  cursor: pointer;
}

.card-bg-characters {
  --x: 0px;
  --y: 0px;
  --tx: 0px;
  --ty: 0px;
  --scale: 1;
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  word-wrap: break-word;
  font-size: 14px;
  line-height: 0.5;
  overflow: hidden;
  font-family: 'Montserrat', sans-serif;
  color: #a8a8a8;
  opacity: 0;
  transition: opacity 0.45s ease, transform 0.45s cubic-bezier(.2,.9,.3,1);
  will-change: transform, opacity;
  -webkit-mask-image: radial-gradient(
    250px circle at var(--x) var(--y),
    #000 15%,
    rgba(0, 0, 0, 0.3),
    transparent
  );
  mask-image: radial-gradient(
    250px circle at var(--x) var(--y),
    #000 15%,
    rgba(0, 0, 0, 0.3),
    transparent
  );
  transform: translate(var(--tx, 0px), var(--ty, 0px)) scale(var(--scale, 1.025));
  padding: 20px;
  z-index: 5;
}

.card-interactive:hover .card-bg-characters {
  opacity: 1;
}
</style>