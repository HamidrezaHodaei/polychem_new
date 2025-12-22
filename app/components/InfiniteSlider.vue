<template>
    <div class="slider" ref="sliderRef">
      <div class="slide-track">
        <div v-for="(img, index) in images" :key="index" class="slide">
          <!-- changed: add class and remove fixed height/width attributes -->
          <img :src="img" alt="logo" class="logo" :class="{ active: activeIndex === index }" />
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, onUnmounted } from 'vue';

  const baseImages = [
   "/Saaf-film.jpg",    
  "/Olsanbaft.png",
  "/Paraplastic.png",    
  "/Reyhaneh.png",
  "/logoMegabiz.jpg"
      
      
    ];
  // تکرار آرایه برای اسکرول بی‌نهایت بدون فاصله سفید
  const images = [...baseImages, ...baseImages, ...baseImages, ...baseImages, ...baseImages, ...baseImages, ...baseImages, ...baseImages, ...baseImages, ...baseImages];

  // --- new reactive/logic for centering detection ---
  const sliderRef = ref(null);
  const activeIndex = ref(-1);
  let rafId = null;
  const CENTER_THRESHOLD = 40; // پیکسل - مقدار آستانه برای تشخیص "تقریبا در مرکز"

  function updateActive() {
    const slider = sliderRef.value;
    if (!slider) {
      rafId = requestAnimationFrame(updateActive);
      return;
    }
    const slides = Array.from(slider.querySelectorAll('.slide'));
    const sliderRect = slider.getBoundingClientRect();
    const sliderCenterX = sliderRect.left + sliderRect.width / 2;

    let bestIndex = -1;
    let bestDistance = Infinity;
    slides.forEach((slideEl, i) => {
      const rect = slideEl.getBoundingClientRect();
      const slideCenterX = rect.left + rect.width / 2;
      const dist = Math.abs(slideCenterX - sliderCenterX);
      if (dist < bestDistance) {
        bestDistance = dist;
        bestIndex = i;
      }
    });

    // تنها وقتی فاصله کمتر از آستانه است، فعال میکنیم
    activeIndex.value = bestDistance <= CENTER_THRESHOLD ? bestIndex : -1;

    rafId = requestAnimationFrame(updateActive);
  }

  onMounted(() => {
    rafId = requestAnimationFrame(updateActive);
  });

  onUnmounted(() => {
    if (rafId) cancelAnimationFrame(rafId);
  });
  </script>
  
  <style scoped lang="scss">
  body {
    align-items: center;
    background: #e3e3e3;
    display: flex;
    height: 100vh;
    justify-content: center;
  }
  
  @mixin white-gradient {
    background: linear-gradient(to right, rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 0) 100%);
  }
  
  $animationSpeed: 40s;
  
  @keyframes scroll {
    0% {
      transform: translateX(0);
    }
    100% {
      transform: translateX(calc(-250px * 10));
    }
  }
  
  .slider {
    background: white;
    box-shadow: 0 10px 20px -5px rgba(0, 0, 0, 0.125);
    height: 100px;
    margin: 0; // حذف فاصله اطراف
    padding: 0; // حذف padding احتمالی
    overflow: hidden;
    position: relative;
    width: 100vw; // پر کردن کل عرض صفحه

    .slide-track {
      animation: scroll $animationSpeed linear infinite;
      display: flex;
      width: calc(250px * 20);
    }
  
    .slide {
      height: 100px;
      width: auto;
      flex-shrink: 0;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 0 28px; // فاصله افقی خیلی کم
    }
    
    /* changed: target .logo class and constrain size while preserving aspect ratio */
    .logo {
      max-height: 72px;
      max-width: 220px;
      width: auto;
      height: auto;
      object-fit: contain;
      display: block;
      filter: grayscale(100%);
      transition: filter 0.25s ease, transform 0.25s ease;
    }

    .logo.active {
      filter: grayscale(0%);
      transform: scale(1.03);
    }

    .slide:hover .logo {
      filter: grayscale(0%);
    }
  }
  </style>
