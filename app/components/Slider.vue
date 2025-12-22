<template>
  <div class="container mx-auto px-4 font-[Montserrat] text-red-100]">
    <div class="w-full">
      <div
        ref="cardsContainer"
        class="relative flex justify-center items-center h-[80vh] mb-24"
      >
        <!-- Card 1 -->
        <div
          ref="card1"
          class="absolute w-[90%] md:w-[85.5%] h-[76.5vh] rounded-[20px] shadow-2xl flex flex-col md:flex-row items-stretch overflow-hidden"
          style="top: 0px; z-index: 2;"
        >
          <!-- Left: Image (60% on desktop, 100% on mobile) -->
          <div class="card-image w-full md:w-3/5 overflow-hidden rounded-t-[16px] md:rounded-tl-[20px] md:rounded-bl-[20px] md:rounded-tr-[0px] md:rounded-br-[0px]">
            <img src="/955w.jpg" alt="Product 1" class="w-full h-full object-cover filter grayscale" />
          </div>

          <!-- Right: Content (40% on desktop) -->
          <div class="w-full md:w-2/5 p-8 flex flex-col justify-center gap-4 card-content">
            <h2 class="text-3xl md:text-4xl uppercase font-[Montserrat] text-[#848484]">Rotochem 0955W</h2>
            <p class="text-base md:text-lg text-[#848484] font-[Montserrat]">Rotochem 0955W – White rotational molding compound based on MDPE copolymer with narrow MWD and TiO₂.Offering exceptional durability and superior UV resistance, it is the perfect choice for demanding industrial and consumer applications.</p>
            <NuxtLink to="/product?index=0" class="btn-slide-down w-full md:w-auto h-12 rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors max-w-[220px] flex items-center justify-center">
              <span class="relative z-10 font-medium">Technical Data Sheet </span>
            </NuxtLink>
          </div>
        </div>

        <!-- Card 2 -->
        <div
          ref="card2"
          class="absolute w-[90%] md:w-[85.5%] h-[76.5vh] rounded-[20px] shadow-2xl flex flex-col md:flex-row items-stretch overflow-hidden"
          style="top: 30px; z-index: 3;"
        >
          <div class="card-image w-full md:w-3/5 overflow-hidden rounded-t-[16px] md:rounded-tl-[20px] md:rounded-bl-[20px] md:rounded-tr-[0px] md:rounded-br-[0px]">
            <img src="/POLYFILF700.jpg" alt="Product 2" class="w-full h-full object-cover filter grayscale" />
          </div>

          <div class="w-full md:w-2/5 p-8 flex flex-col justify-center gap-4 card-content">
            <h2 class="text-3xl md:text-4xl font-[Montserrat] text-[#848484]">POLYFIL F700</h2>
            <p class="text-base md:text-lg text-[#848484] font-[Montserrat]">POLYFIL F700 – High-performance HDPE blown film compound offering superior strength, excellent film uniformity, and outstanding processability.</p>
            <NuxtLink to="/product?index=2" class="btn-slide-down w-full md:w-auto h-12 rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors max-w-[220px] flex items-center justify-center">
              <span class="relative z-10 font-medium">Technical Data Sheet </span>
            </NuxtLink>
          </div>
        </div>

        <!-- Card 3 -->
        <div
          ref="card3"
          class="absolute w-[90%] md:w-[85.5%] h-[76.5vh] rounded-[20px] shadow-2xl flex flex-col md:flex-row items-stretch overflow-hidden"
          style="top: 60px; z-index: 4;"
        >
          <div class="card-image w-full md:w-3/5 overflow-hidden rounded-t-[16px] md:rounded-tl-[20px] md:rounded-bl-[20px] md:rounded-tr-[0px] md:rounded-br-[0px]">
            <img src="/Hdchem.jpg" alt="Product 3" class="w-full h-full object-cover filter grayscale" />
          </div>

          <div class="w-full md:w-2/5 p-8 flex flex-col justify-center gap-4 card-content">
            <h2 class="text-3xl md:text-4xl font-[Montserrat] text-[#848484] uppercase">HDCHEM 4760</h2>
            <p class="text-base md:text-lg text-[#848484] font-[Montserrat]">HDCHEM 4760 – Specialized PE copolymer compound for blow molding applications. This polymer matarials belong and good melt strength, balanced mechanical properties, and superior processability for industrial and consumer parts.</p>
            <NuxtLink to="/product?index=4" class="btn-slide-down w-full md:w-auto h-12 rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors max-w-[220px] flex items-center justify-center">
              <span class="relative z-10 font-medium">Technical Data Sheet </span>
            </NuxtLink>
          </div>
        </div>
 
        <!-- Card 4 -->
        <div
          ref="card4"
          class="absolute w-[90%] md:w-[85.5%] h-[76.5vh] rounded-[20px] shadow-2xl flex flex-col md:flex-row items-stretch overflow-hidden"
          style="top: 90px; z-index: 5;"
        >
          <div class="card-image w-full md:w-3/5 overflow-hidden rounded-t-[16px] md:rounded-tl-[20px] md:rounded-bl-[20px] md:rounded-tr-[0px] md:rounded-br-[0px] bg-white-100">
            <img src="/slimchem.jpg" alt="Product 4" class="w-full h-full object-cover filter grayscale" />
          </div>

          <div class="w-full md:w-2/5 p-8 flex flex-col justify-center gap-4 card-content">
            <h2 class="text-3xl md:text-4xl font-[Montserrat] text-[#848484]">SlipChem-E 178</h2>
            <p class="text-base md:text-lg text-[#848484] font-[Montserrat]">SlipChem-E 178 – Premium slip masterbatch in PE carrier. Significantly lowers COF with excellent dispersion, high thermal stability, and consistent migration for film winding and packaging.</p>
            <NuxtLink to="/product?index=3" class="btn-slide-down w-full md:w-auto h-12 rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors max-w-[220px] flex items-center justify-center">
              <span class="relative z-10 font-medium">Technical Data Sheet </span>
            </NuxtLink>
          </div>
        </div>

        <!-- Card 5 -->
        <div
          ref="card5"
          class="absolute w-[90%] md:w-[85.5%] h-[76.5vh] rounded-[20px] shadow-2xl flex flex-col md:flex-row items-stretch overflow-hidden"
          style="top: 120px; z-index: 6;"
        >
          <div class="card-image w-full md:w-3/5 overflow-hidden rounded-t-[16px] md:rounded-tl-[20px] md:rounded-bl-[20px] md:rounded-tr-[0px] md:rounded-br-[0px]">
            <img src="/RAFCOLOR.jpg" alt="Product 5" class="w-full h-full object-cove filter grayscale" />
          </div>

          <div class="w-full md:w-2/5 p-8 flex flex-col justify-center gap-4 card-content">
            <h2 class="text-3xl md:text-4xl font-[Montserrat] text-[#848484]">RAFCOLOR 1560</h2>
            <p class="text-base md:text-lg text-[#848484] font-[Montserrat]">RAFCOLOR is a high-quality white masterbatch containing a high loading of rutile titanium dioxide (TiO₂) and thermoplastic polypropylene resin. The selected TiO₂ offers excellent opacity and dispersion properties.</p>
            <NuxtLink to="/product?index=5" class="btn-slide-down w-full md:w-auto h-12 rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors max-w-[220px] flex items-center justify-center">
              <span class="relative z-10 font-medium">Technical Data Sheet </span>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, onBeforeUnmount } from 'vue'
import { gsap } from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'
import { ScrollToPlugin } from 'gsap/ScrollToPlugin'

gsap.registerPlugin(ScrollTrigger, ScrollToPlugin)

const cardsContainer = ref(null)
const card1 = ref(null)
const card2 = ref(null)
const card3 = ref(null)
const card4 = ref(null)
const card5 = ref(null)

let tl = null

onMounted(() => {
  const savedScrollPosition = sessionStorage.getItem('homeScrollPosition')
  
  ScrollTrigger.getAll().forEach(trigger => trigger.kill())
  
  tl = gsap.timeline({
    scrollTrigger: {
      trigger: cardsContainer.value,
      pin: true,
      pinSpacing: true,
      markers: false,
      start: 'top-=50px top',
      end: '+=3800',
      scrub: 0.8,
      fastScrollEnd: true,
      onUpdate: (self) => {
        sessionStorage.setItem('homeScrollPosition', window.scrollY)
      }
    },
  })

  tl.addLabel('card1')
  tl.to(card1.value, {
    yPercent: 0,
    opacity: 1,
    duration: 0.6,
    ease: 'power2.out',
  })

  tl.from(card2.value, {
    yPercent: 75,
    opacity: 0,
    duration: 0.8,
    ease: 'power2.out',
  }, '+=0.1')

  tl.to([card1.value, card2.value], {
    scale: (index) => index === 0 ? 0.925 : 1,
    yPercent: (index) => index === 0 ? -0.75 : 0,
    opacity: (index) => index === 0 ? 0.85 : 1,
    duration: 0.8,
    ease: 'power2.inOut',
  }, '-=0.2')

  tl.from(card3.value, {
    yPercent: 75,
    opacity: 0,
    duration: 0.8,
    ease: 'power2.out',
  }, '+=0.1')

  tl.to([card2.value, card3.value], {
    scale: (index) => index === 0 ? 0.95 : 1,
    yPercent: (index) => index === 0 ? -0.5 : 0,
    opacity: (index) => index === 0 ? 0.85 : 1,
    duration: 0.8,
    ease: 'power2.inOut',
  }, '-=0.2')

  tl.from(card4.value, {
    yPercent: 75,
    opacity: 0,
    duration: 0.8,
    ease: 'power2.out',
  }, '+=0.1')

  tl.to([card3.value, card4.value], {
    scale: (index) => index === 0 ? 0.975 : 1,
    yPercent: (index) => index === 0 ? -0.4 : 0,
    opacity: (index) => index === 0 ? 0.85 : 1,
    duration: 0.8,
    ease: 'power2.inOut',
  }, '-=0.2')

  tl.from(card5.value, {
    yPercent: 75,
    opacity: 0,
    duration: 0.8,
    ease: 'power2.out',
  }, '+=0.1')

  tl.to([card4.value, card5.value], {
    scale: (index) => index === 0 ? 0.9875 : 1,
    yPercent: (index) => index === 0 ? -0.3 : 0,
    opacity: (index) => index === 0 ? 0.85 : 1,
    duration: 0.8,
    ease: 'power2.inOut',
  }, '-=0.2')

  tl.to([card1.value, card2.value, card3.value, card4.value], {
    scale: (index) => 0.9 + index * 0.025,
    yPercent: (index) => -(1.5 - index * 0.375),
    opacity: 0.7,
    duration: 1,
    ease: 'power2.inOut',
  }, '-=0.2')

  if (savedScrollPosition !== null) {
    requestAnimationFrame(() => {
      window.scrollTo(0, parseInt(savedScrollPosition))
    })
  }
})

onBeforeUnmount(() => {
  sessionStorage.setItem('homeScrollPosition', window.scrollY)
})

onUnmounted(() => {
  if (tl && tl.scrollTrigger) {
    tl.scrollTrigger.kill()
  }
  if (tl) {
    tl.kill()
  }
})
</script>

<style scoped>
/* اطمینان از سفید بودن پس‌زمینه‌ی ورودی‌ها (حتی در حالت autofill مرورگر) */
input {
  background-color: #ffffff !important;
}

input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus {
  -webkit-box-shadow: 0 0 0 1000px #ffffff inset !important;
  box-shadow: 0 0 0 1000px #ffffff inset !important;
  -webkit-text-fill-color: #111111 !important;
  caret-color: #111111;
  transition: background-color 5000s ease-in-out 0s;
}

.btn-slide-down {
  background-color: transparent;
}

.btn-slide-down::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: #FFCD05;
  transform: translateY(-100%);
  transition: transform 300ms ease;
  z-index: 0;
}

.btn-slide-down:hover::before,
.btn-slide-down:focus-visible::before {
  transform: translateY(0);
}

.btn-slide-down:hover,
.btn-slide-down:focus-visible {
  color: #ffffff;
  outline: none;
}

.card-content {
  background: #ffffff;
  position: relative;
  z-index: 1;
  color: #111827;
}

.card-content > * {
  position: relative;
  z-index: 2;
}

.card-content::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  width: 6px;
  height: 100%;
  background: #FFCD05;
  border-top-left-radius: 20px;
  border-bottom-left-radius: 20px;
  z-index: 1;
}

@media (max-width: 767px) {
  .card-content::before {
    left: 0;
    top: 0;
    width: 100%;
    height: 6px;
    border-top-left-radius: 20px;
    border-top-right-radius: 20px;
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
  }

  .rounded-tl-\[20px\], .rounded-bl-\[20px\], .rounded-tr-\[20px\], .rounded-br-\[20px\] {
    border-radius: 20px;
  }
}

.slider {
  transition: all 0.6s ease-in-out;
}

.slider.scroll-active {
  animation: slideInUp 0.6s ease-in-out;
}

@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.card-image {
  background-color: #ffffff !important;
}

.card-image > img {
  display: block;
  background-color: #ffffff !important;
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
}

[ref="card1"],
[ref="card2"],
[ref="card3"],
[ref="card4"],
[ref="card5"] {
  will-change: transform, opacity;
  transform: translateZ(0);
  -webkit-transform: translateZ(0);
}
</style>