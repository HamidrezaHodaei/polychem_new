<template>
    <div class="hero-container">
      <!-- First Section -->
      <picture class="bg-picture">
        <!-- جایگزینی GIF با ویدیو MP4 -->
        <video
          src="/shutterstock_vidhd_3898471189.mp4"
          autoplay
          loop
          muted
          playsinline
          style="width:100%;height:calc(100vh - 5rem);object-fit:cover;"
        ></video>
      </picture>
      <figure class="content-box">
        <h1 class="main-title text-bold">
          POLYCHEM
          <span class="subtitle">Advanced Polymer Engineering</span>
        </h1>
        <figcaption>
          <p class="intro-text">POLYCHEM is a specialist producer of advanced polymer compounds and masterbatches in the Aras Free Zone. With decades of experience and top engineers, we create high-performance materials for modern manufacturing.</p>
        </figcaption>
      </figure>
    </div>
  </template>
  
  <script setup>
import { onMounted } from 'vue'

onMounted(() => {
  // Preload video
  const video = document.querySelector('.bg-picture video')
  if (video) {
    video.load()
  }
})
  </script>
  
  <style scoped>
  /* Reset */
  *:not(head, script, svg, svg *),
  *:not(head, script, svg, svg *)::before,
  *:not(head, script, svg, svg *)::after {
    all: unset;
    box-sizing: border-box;
  }
  
  /* Body & Layout */
  .hero-container {
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: calc(100vh - 5rem); /* adjusted to subtract navbar height (h-20 = 5rem) */
    margin: 0;
    background: black;
    color: white;
    overflow: hidden;
  }
  
  /* Picture & Image */
  .bg-picture {
    position: relative;
    width: 100%;
    height: calc(100vh - 5rem); /* keep consistent with hero-container */
    overflow: hidden;
  }
  
  .bg-picture img {
    width: 100%;
    height: calc(100vh - 5rem); /* keep consistent with hero-container */
    object-fit: cover;
  }
  
  /* Overlay Animations */
  .bg-picture::before,
  .bg-picture::after {
    position: absolute;
    content: "";
    background-image: linear-gradient(
      135deg,
      rgba(119, 118, 115, 0.18),
      rgba(138, 138, 138, 0.18)
    ); /* yellow + gray overlay */
  }
  
  .bg-picture::before {
    top: 0;
    left: 0;
    width: 25%;
    height: 25%;
    animation: a 20s ease infinite;
  }
  
  .bg-picture::after {
    top: 0;
    left: 0;
    width: 25%;
    height: 25%;
    animation: b 10s ease infinite;
  }
  
  /* Keyframes a */
  @keyframes a {
    0% {
      width: 25%;
      height: 25%;
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    12.49% { left: 0; }
    12.5% {
      width: 100%;
      left: unset;
      backdrop-filter: blur(0vmin);
      -webkit-backdrop-filter: blur(0vmin);
    }
    25% {
      width: 25%;
      height: 25%;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    37.49% { top: 0; bottom: 0; }
    37.5% {
      top: unset;
      height: 100%;
      backdrop-filter: blur(0vmin);
      -webkit-backdrop-filter: blur(0vmin);
    }
    50% {
      width: 25%;
      height: 25%;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    62.49% { left: unset; right: 0; }
    62.5% {
      right: unset;
      width: 100%;
      height: 100%;
      backdrop-filter: blur(0vmin);
      -webkit-backdrop-filter: blur(0vmin);
    }
    74.9% { top: unset; bottom: 0; }
    75% {
      width: 25%;
      height: 25%;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    87.49% { top: unset; bottom: 0; }
    87.5% {
      height: 100%;
      top: 0;
      backdrop-filter: blur(0vmin);
      -webkit-backdrop-filter: blur(0vmin);
    }
    100% {
      width: 25%;
      height: 25%;
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
  }
  
  /* Keyframes b */
  @keyframes b {
    0% {
      width: 25%;
      height: 25%;
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    12.49% { top: 0; }
    12.5% {
      height: 100%;
      top: unset;
      backdrop-filter: blur(0vmin);
      -webkit-backdrop-filter: blur(0vmin);
    }
    25% {
      width: 25%;
      height: 25%;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    37.49% { left: 0; right: 0; }
    37.5% {
      left: unset;
      width: 100%;
      backdrop-filter: blur(0vmin);
      -webkit-backdrop-filter: blur(0vmin);
    }
    50% {
      width: 25%;
      height: 25%;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    62.49% { top: unset; bottom: 0; }
    62.5% {
      top: 0;
      height: 100%;
      backdrop-filter: blur(0vmin);
      -webkit-backdrop-filter: blur(0vmin);
    }
    74.9% { bottom: unset; right: 0; }
    75% {
      width: 25%;
      height: 25%;
      left: unset;
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    87.49% { left: unset; right: 0; }
    87.5% {
      width: 100%;
      left: 0;
      backdrop-filter: blur(0vmin);
      -webkit-backdrop-filter: blur(0vmin);
    }
    100% {
      width: 25%;
      height: 25%;
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
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
    max-width: 600px; /* قبلاً 700px بود */
    border: 1px solid rgba(255, 255, 255, 0.2);
    padding: 24px;    /* قبلاً 30px بود */
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: space-between;
    background-color: rgba(0, 0, 0, 0.2);
    z-index: 1;
    animation: c 10s linear infinite;
  }
  
  @keyframes c {
    from {
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
    to {
      backdrop-filter: blur(2vmin);
      -webkit-backdrop-filter: blur(2vmin);
    }
  }
  
  .content-box::after {
    position: absolute;
    bottom: -10px;
    right: -10px;
    width: 60px;
    height: 60px;
    /* تغییر: رنگ سفید و سایه برای کنتراست بهتر */
    background-color: #ffd000;
    box-shadow: 0 6px 18px rgba(0,0,0,0.45);
    content: "";
    z-index: -1;
    animation: d 5s linear infinite;
  }
  
  @keyframes d {
    from {
      transform: rotateZ(0deg);
      border-radius: 0;
    }
    50% {
      border-radius: 49%;
    }
    to {
      transform: rotateZ(360deg);
      border-radius: 0;
    }
  }
  
  /* Typography */
  h1 {
    font-family: "Montserrat", sans-serif;
    font-size: 40px;
    line-height: 1.5cap;
    text-wrap: balance;
    color: #ffffff;
    margin-bottom: 20px;
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  /* Make elements with .text-bold actually bold */
  .text-bold {
    font-weight: 700;
  }

  .subtitle {
    font-size: 22px; /* قبلاً 18px بود */
    font-weight: 400;
    color: #ffffff;
    letter-spacing: 0.5px;
  }

  figcaption {
    font-family: "Figtree", sans-serif;
    font-size: 18px;
    font-weight: 400;
    line-height: 2cap;
    text-wrap: pretty;
    padding-left: 30px;
    border-left: 1px solid white;
    color: #ffffff;
  }

  .intro-text {
    margin-bottom: 20px;
  }

  .what-we-do-title {
    font-weight: 600;
    margin-bottom: 12px;
    margin-top: 16px;
  }

  .what-we-do-list {
    list-style-position: inside;
    margin: 0;
    padding: 0;
    margin-bottom: 16px;
  }

  .what-we-do-list li {
    margin-bottom: 8px;
    display: list-item;
  }

  .closing-text {
    margin-top: 16px;
  }

  /* Responsive */
  @media (max-width: 780px) {
    .content-box {
      max-width: 90%;
      padding: 12px; /* برای موبایل padding کمتر شود */
    }
  }

  @media (max-width: 600px) {
    h1 {
      font-size: 32px;
    }
    figcaption {
      font-size: 16px;
    }
  }
  </style>