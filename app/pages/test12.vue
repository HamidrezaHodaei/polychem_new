<template>
  <div v-if="isLoading" class="loading-screen">
    <div class="bouncing-circle"></div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const isLoading = ref(true)

onMounted(() => {
  // برای تست - بعد از 3 ثانیه لودینگ رو مخفی می‌کنه
  // تو پروژه واقعی این رو بر اساس نیازت تغییر بده
  setTimeout(() => {
    isLoading.value = false
  }, 3000)
})
</script>

<style scoped>
.loading-screen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

.bouncing-circle {
  width: 80px;
  height: 80px;
  background: #ffd700;
  border-radius: 50%;
  box-shadow: 
    0 0 20px rgba(255, 215, 0, 0.5),
    0 0 40px rgba(255, 215, 0, 0.3),
    0 0 60px rgba(255, 215, 0, 0.2);
  animation: bounce 1s ease-in-out infinite;
}

@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
    animation-timing-function: ease-in;
  }
  50% {
    transform: translateY(-100px);
    animation-timing-function: ease-out;
  }
}
</style>