<template>
  <div class="hero-container">
    <!-- Background Section -->
    <div class="bg-wrapper">

      <!-- Main Video with Poster -->
      <video
        ref="videoRef"
        class="bg-video"
        :class="{ 'video-loaded': videoLoaded }"
        src="/shutterstock_vidhd_3898471189.mp4"
        poster="/video-poster.jpg"
        loop
        muted
        playsinline
        preload="auto"
        @canplay="onVideoCanPlay"
        @error="onVideoError"
      ></video>

      <!-- Animated Overlays -->
      <div class="overlay overlay-1"></div>
      <div class="overlay overlay-2"></div>
    </div>

    <!-- Content Box -->
    <figure class="content-box">
      <h1 class="main-title text-bold">
        POLYCHEM
        <span class="subtitle">Advanced Polymer Engineering</span>
      </h1>
      <figcaption>
        <p class="intro-text">
          POLYCHEM is a specialist producer of advanced polymer compounds and masterbatches in the Aras Free Zone. 
          With decades of experience and top engineers, we create high-performance materials for modern manufacturing.
        </p>
      </figcaption>
    </figure>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const videoRef = ref(null)
const videoLoaded = ref(false)

const onVideoCanPlay = () => {
  if (videoRef.value) {
    videoRef.value.play()
    setTimeout(() => {
      videoLoaded.value = true
    }, 100)
  }
}

const onVideoError = () => {
  console.error('Video failed to load')
}

onMounted(() => {
  if (videoRef.value) {
    videoRef.value.load()
  }
})
</script>

<style scoped>
/* Reset */
*,
*::before,
*::after {
  box-sizing: border-box;
}

/* Hero Container */
.hero-container {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: calc(100vh - 5rem);
  margin: 0;
  background: #1a1a1a;
  color: #fff;
  overflow: hidden;
}

/* Background Wrapper */
.bg-wrapper {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background: url('/video-poster.jpg') center center / cover no-repeat;
}

/* Background Video */
.bg-video {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 1;
  z-index: 2;
}

/* Animated Overlays */
.overlay {
  position: absolute;
  content: "";
  background-image: linear-gradient(
    135deg,
    rgba(119, 118, 115, 0.18),
    rgba(138, 138, 138, 0.18)
  );
  z-index: 3;
  pointer-events: none;
}

.overlay-1 {
  top: 0;
  left: 0;
  width: 25%;
  height: 25%;
  animation: overlay-anim-1 20s ease infinite;
}

.overlay-2 {
  top: 0;
  left: 0;
  width: 25%;
  height: 25%;
  animation: overlay-anim-2 10s ease infinite;
}

/* Overlay Animation 1 */
@keyframes overlay-anim-1 {
  0% {
    width: 25%;
    height: 25%;
    left: 0;
    top: 0;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
  12.5% {
    width: 100%;
    left: 0;
    backdrop-filter: blur(0);
    -webkit-backdrop-filter: blur(0);
  }
  25% {
    width: 25%;
    height: 25%;
    right: 0;
    left: auto;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
  37.5% {
    height: 100%;
    top: 0;
    backdrop-filter: blur(0);
    -webkit-backdrop-filter: blur(0);
  }
  50% {
    width: 25%;
    height: 25%;
    bottom: 0;
    top: auto;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
  62.5% {
    width: 100%;
    left: 0;
    right: auto;
    backdrop-filter: blur(0);
    -webkit-backdrop-filter: blur(0);
  }
  75% {
    width: 25%;
    height: 25%;
    left: 0;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
  87.5% {
    height: 100%;
    top: 0;
    bottom: auto;
    backdrop-filter: blur(0);
    -webkit-backdrop-filter: blur(0);
  }
  100% {
    width: 25%;
    height: 25%;
    top: 0;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
}

/* Overlay Animation 2 */
@keyframes overlay-anim-2 {
  0% {
    width: 25%;
    height: 25%;
    left: 0;
    top: 0;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
  12.5% {
    height: 100%;
    bottom: 0;
    top: auto;
    backdrop-filter: blur(0);
    -webkit-backdrop-filter: blur(0);
  }
  25% {
    width: 25%;
    height: 25%;
    bottom: 0;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
  37.5% {
    width: 100%;
    right: 0;
    left: auto;
    backdrop-filter: blur(0);
    -webkit-backdrop-filter: blur(0);
  }
  50% {
    width: 25%;
    height: 25%;
    right: 0;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
  62.5% {
    height: 100%;
    top: 0;
    bottom: auto;
    backdrop-filter: blur(0);
    -webkit-backdrop-filter: blur(0);
  }
  75% {
    width: 25%;
    height: 25%;
    right: 0;
    top: 0;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
  87.5% {
    width: 100%;
    left: 0;
    right: auto;
    backdrop-filter: blur(0);
    -webkit-backdrop-filter: blur(0);
  }
  100% {
    width: 25%;
    height: 25%;
    left: 0;
    top: 0;
    backdrop-filter: blur(2vmin);
    -webkit-backdrop-filter: blur(2vmin);
  }
}

/* Content Box */
.content-box {
  position: absolute;
  top: 40%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  max-width: 600px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  padding: 24px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  background-color: rgba(0, 0, 0, 0.2);
  backdrop-filter: blur(2vmin);
  -webkit-backdrop-filter: blur(2vmin);
  z-index: 10;
}

.content-box::after {
  position: absolute;
  bottom: -10px;
  right: -10px;
  width: 60px;
  height: 60px;
  background-color: #ffd000;
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.45);
  content: "";
  z-index: -1;
  animation: shape-rotate 5s linear infinite;
}

@keyframes shape-rotate {
  0% {
    transform: rotate(0deg);
    border-radius: 0;
  }
  50% {
    border-radius: 49%;
  }
  100% {
    transform: rotate(360deg);
    border-radius: 0;
  }
}

/* Typography */
.main-title {
  font-family: "Montserrat", -apple-system, BlinkMacSystemFont, sans-serif;
  font-size: 40px;
  font-weight: 700;
  line-height: 1.2;
  color: #fff;
  margin: 0 0 20px 0;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.subtitle {
  font-size: 22px;
  font-weight: 400;
  color: #fff;
  letter-spacing: 0.5px;
}

figcaption {
  font-family: "Figtree", -apple-system, BlinkMacSystemFont, sans-serif;
  font-size: 18px;
  font-weight: 400;
  line-height: 1.6;
  padding-left: 30px;
  border-left: 1px solid #fff;
  color: #fff;
  margin: 0;
}

.intro-text {
  margin: 0;
}

/* Responsive */
@media (max-width: 780px) {
  .content-box {
    max-width: 90%;
    padding: 20px;
  }
}

@media (max-width: 600px) {
  .main-title {
    font-size: 32px;
  }
  
  .subtitle {
    font-size: 18px;
  }
  
  figcaption {
    font-size: 16px;
    padding-left: 20px;
  }
  
  .content-box {
    padding: 16px;
  }
}
</style>