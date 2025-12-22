<template>
    <div class="hero-container">
      <!-- First Section -->
      <picture class="bg-picture">
        <!-- استفاده از fetchpriority برای اولویت بارhigher + loading eager -->
        <img src="/shutterstock_vidhd_3898471189%20(2).gif" alt="background gif" loading="eager" decoding="sync" fetchpriority="high" />
      </picture>
      <figure class="content-box">
        <h1>Welcome to my website</h1>
        <figcaption>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
        </figcaption>
      </figure>
    </div>
  </template>
  
  <script setup>
  import { onMounted } from 'vue'
  
  // بصورت محافظه‌کارانه یک preload در head اضافه کنیم تا مرورگر قبل‌تر دانلود را آغاز کند
  onMounted(() => {
    try {
      const href = '/shutterstock_vidhd_3898471189%20(2).gif'
      // اگر preload قبلاً اضافه نشده باشد
      if (!document.querySelector(`link[rel="preload"][href="${href}"]`)) {
        const link = document.createElement('link')
        link.rel = 'preload'
        link.as = 'image'
        link.href = href
        // بعضی مرورگرها از fetchpriority بهتر پشتیبانی می‌کنند
        link.setAttribute('fetchpriority', 'high')
        document.head.appendChild(link)
      }
    } catch (e) {
      // ignore
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
    /* حذف رنگ مشکی: استفاده از خاکستری روشن به عنوان پس‌زمینه */
    background: var(--hero-bg, #e5e7eb); /* light gray */
    color: var(--hero-text, #1f2937); /* متن خاکستری تیره */
    overflow: hidden;
  }
  
  /* Picture & Image */
  .bg-picture {
    position: relative;
    width: 100%;
    height: calc(100vh - 5rem); /* keep consistent with hero-container */
    overflow: hidden;
    /* اگر تصویر هنوز کامل نیامده بود، این رنگ جلوی نمایش سیاه را می‌گیرد */
    background-color: var(--hero-bg, #e5e7eb);
  }
  
  .bg-picture img {
    width: 100%;
    height: calc(100vh - 5rem); /* keep consistent with hero-container */
    object-fit: cover;
    filter: grayscale(10%) brightness(85%); /* ملایم کردن تصویر */
    background-color: transparent;
  }
  
  /* Overlay Animations */
  .bg-picture::before,
  .bg-picture::after {
    position: absolute;
    content: "";
    /* فقط خاکستری شفاف و زرد برای افکت‌ها */
    background-image: linear-gradient(
      135deg,
      rgba(148,163,184,0.12), /* gray tint */
      rgba(255,208,0,0.06)     /* faint yellow tint */
    );
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
    max-width: 700px;
    /* حاشیه خاکستری کم‌رنگ */
    border: 1px solid rgba(55,65,81,0.06);
    padding: 30px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: space-between;
    /* استفاده از سفید نیمه‌شفاف تا متن خاکستری داخل آن خوانا باشد */
    background-color: rgba(255,255,255,0.9);
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
    /* زرد به عنوان آکسن */
    background-color: #ffd000;
    box-shadow: 0 6px 18px rgba(55,65,81,0.12);
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
    line-height: 1.2;
    color: var(--heading-color, #1f2937); /* خاکستری تیره */
    margin-bottom: 20px;
  }
  
  figcaption {
    font-family: "Figtree", sans-serif;
    font-size: 18px;
    font-weight: 400;
    line-height: 1.6;
    padding-left: 30px;
    border-left: 1px solid rgba(55,65,81,0.06);
    color: var(--text-color, #374151);
  }
  
  /* Responsive */
  @media (max-width: 780px) {
    .content-box {
      max-width: 90%;
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