<template>
  <div class="relative w-full h-screen overflow-hidden bg-gray-900">
    <!-- Background Page (Yellow) -->
    <div 
      ref="backgroundPage"
      class="absolute inset-0 bg-yellow-400 flex items-center justify-center"
      style="clip-path: circle(0% at 50% 100%);"
    >
      <div class="text-center">
        <h1 class="text-6xl font-bold text-gray-900 mb-4">خوش آمدید!</h1>
        <p class="text-2xl text-gray-800">صفحه با موفقیت باز شد</p>
      </div>
    </div>

    <!-- Ball -->
    <div 
      ref="ball"
      class="absolute w-8 h-8 bg-yellow-400 rounded-full shadow-lg"
      style="left: 50%; top: 50%; transform: translate(-50%, -50%);"
    ></div>

    <!-- Front Page (Dark) -->
    <div 
      ref="frontPage"
      class="absolute inset-0 bg-gray-900 flex items-center justify-center pointer-events-none"
    >
      <p class="text-gray-600 text-xl">در حال بارگذاری...</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { gsap } from 'gsap';

const ball = ref(null);
const frontPage = ref(null);
const backgroundPage = ref(null);

onMounted(() => {
  const timeline = gsap.timeline();

  // توپ از وسط شروع می‌کند و به بالا می‌رود
  timeline.to(ball.value, {
    y: -200,
    duration: 0.6,
    ease: "power2.out"
  });

  // توپ به پایین می‌آید
  timeline.to(ball.value, {
    y: window.innerHeight / 2 - 16,
    duration: 0.8,
    ease: "bounce.out"
  });

  // پس از برخورد، انیمیشن circular reveal
  timeline.to(backgroundPage.value, {
    clipPath: "circle(150% at 50% 100%)",
    duration: 1.2,
    ease: "power2.inOut"
  }, "+=0.2");

  // محو کردن صفحه جلو
  timeline.to(frontPage.value, {
    opacity: 0,
    duration: 0.5,
    ease: "power2.inOut"
  }, "-=0.8");

  // محو کردن توپ
  timeline.to(ball.value, {
    opacity: 0,
    duration: 0.3
  }, "-=0.5");
});
</script>