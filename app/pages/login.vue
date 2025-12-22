<template>
  
  <div class="min-h-screen flex items-center justify-center bg-[#f1f2f2] p-6 font-sans relative text-[#808285]">
    <!-- Desktop logo (keep for md+) -->
    <NuxtLink to="/" class="absolute top-4 left-4 z-20 hidden md:block" aria-label="Home">
      <img src="/polychem wall B.png" alt="Logo" class="w-40 h-auto" />
    </NuxtLink>

    <!-- Mobile background (behind the white card) - mobile only -->
    <div class="mobile-bg md:hidden" aria-hidden="true">
      <div class="mobile-bg-frame">
        <img
          src="/login.jpg"
          alt="Mobile background"
          class="mobile-bg-img"
          :style="{'--img-translate': imgTranslate + 'px', '--img-y': imgY + 'px'}"
          ref="mobileImgRef"
        />
      </div>
    </div>

    <div class="bg-white rounded-xl shadow-lg h-[640px] max-w-5xl w-full flex flex-col md:flex-row overflow-hidden md:h-[580px] relative form-card">
      
      <!-- Left Image (Desktop only) -->
      <div class="hidden md:block md:flex-1 h-full w-[600px]">
        <img src="/login.jpg" alt="Form Image" class="w-full h-full object-cover" />
      </div>

      <!-- Form -->
      <form class="w-full md:w-[420px] p-8 flex flex-col items-center justify-center h-full pt-0 relative">
        <!-- Mobile logo on top-left of form -->
        <NuxtLink to="/" class="mobile-logo md:hidden absolute top-4 left-4 z-20" aria-label="Home">
          <img src="/polychem wall B.png" alt="Logo" class="w-28 h-auto" />
        </NuxtLink>

        <!-- Back to Home -->
        <a href="/" aria-label="Back to home" class="absolute top-4 right-4 text-gray-500 hover:text-gray-700 text-2xl leading-none">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7"/>
          </svg>
        </a>
        <!-- Logo & Heading -->
        <div class="text-center mb-8">
          <h1 class="text-2xl font-semibold mb-2">Welcome Back!</h1>
          <p>Sign in to access your dasboard.</p>
        </div>

        <!-- Input Fields -->
        <div class="w-full max-w-[420px] flex flex-col gap-6">
          <div class="relative">
            <input required type="text" id="username" placeholder=" " class="peer w-full border-b-2 border-gray-300 focus:border-[#FFCD05] outline-none py-2 bg-white" />
            <label for="username" class="absolute left-0 top-2 text-gray-400 text-sm transition-all
              peer-placeholder-shown:top-2 peer-placeholder-shown:text-gray-400 peer-placeholder-shown:text-sm
              peer-focus:-top-4 peer-focus:text-[#FFCD05] peer-focus:text-sm
              peer-[:not(:placeholder-shown)]:-top-4 peer-[:not(:placeholder-shown)]:text-gray-700 peer-[:not(:placeholder-shown)]:text-sm">Username</label>
          </div>

          <div class="relative">
            <input required :type="showPassword ? 'text' : 'password'" id="password" placeholder=" " class="peer w-full border-b-2 border-gray-300 focus:border-[#FFCD05] outline-none py-2 bg-white pr-10" />
            <label for="password" class="absolute left-0 top-2 text-gray-400 text-sm transition-all
              peer-placeholder-shown:top-2 peer-placeholder-shown:text-gray-400 peer-placeholder-shown:text-sm
              peer-focus:-top-4 peer-focus:text-[#FFCD05] peer-focus:text-sm
              peer-[:not(:placeholder-shown)]:-top-4 peer-[:not(:placeholder-shown)]:text-gray-700 peer-[:not(:placeholder-shown)]:text-sm">Password</label>
            <!-- Password visibility toggle -->
            <button
              type="button"
              @click="showPassword = !showPassword"
              class="absolute right-0 top-1/2 -translate-y-1/2 p-2 text-gray-500 hover:text-gray-700 focus-visible:outline-none"
              aria-label="Toggle password visibility"
            >
              <!-- Eye / Eye-off icon -->
              <svg v-if="!showPassword" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M2 12s3.5-7 10-7 10 7 10 7-3.5 7-10 7-10-7-10-7z"/>
                <circle cx="12" cy="12" r="3"/>
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M3 3l18 18"/>
                <path d="M10.6 10.6a3 3 0 004.2 4.2"/>
                <path d="M9.88 4.26A10.94 10.94 0 0112 4c6.5 0 10 8 10 8a17.44 17.44 0 01-3.36 4.58M6.61 6.61A17.53 17.53 0 002 12s3.5 7 10 7a10.82 10.82 0 004.39-.93"/>
              </svg>
            </button>
          </div>
          <a href="#" class="hover:text-[#FFCD05] text-sm">Forgot your password?</a>
          <button type="submit" class="btn-slide-down w-full h-12 rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors">
            <span class="relative z-10">Sign in</span>
          </button>
        </div>

        <!-- Footer links -->
        <div class="w-full max-w-[480px] mt-6 text-center text-sm text-gray-600 relative pb-10">
          <p>Don't have an account? <a href="/register" class="text-[#FFCD05] hover:text-[#808285]">Sign up</a></p>
          <div class="absolute right-0 -bottom-12 text-right">
            <div class="inline-flex items-center gap-2 cursor-pointer select-none">
              <span class="">English</span>
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 9l6 6 6-6"/>
              </svg>
            </div>
          </div>
        </div>
      </form>
    </div>

    <!-- Mobile navbar at bottom -->
    <div class="md:hidden">
      <NavbarMob />
    </div>

    <!-- Chat widget (always present on page) -->
    <Chatbox />
  </div>
</template>

<script setup>
 import Navbar from "~/components/navbar.vue"
 import NavbarMob from "~/components/navbar-mob.vue"
 import Chatbox from '~/components/chatbox.vue'
 import { ref } from 'vue'
 const showPassword = ref(false)
 // ref to control image translation (in px). می‌توانید این مقدار را از هرجا تغییر دهید.
 const imgTranslate = ref(0)
 const imgY = ref(-80) // مقدار پیش‌فرض: تصویر را اندکی به بالا می‌بریم تا پایینش در قاب دیده شود (px)
 const mobileImgRef = ref(null)
 
 // مثال: اگر خواستید از جاوااسکریپت مقدار translate را تغییر دهید:
 // imgTranslate.value = -40
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

/* Hover filled slide down button effect */
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

/* Mobile background positioned behind the white card (mobile only) */
.mobile-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 50vh; /* نصف ارتفاع صفحه در موبایل */
  z-index: 0; /* پشت کارت */
  pointer-events: none;
}

/* frame keeps the same size while the image inside can move */
.mobile-bg-frame {
  width: 100%;
  height: 100%;
  overflow: hidden;
  position: relative;
  /* بدون border-radius تا تصویر کامل پشت کارت دیده شود */
}

/* image can translate without changing the frame */
.mobile-bg-img {
  position: absolute;
  top: 0;
  left: 50%;
  /* افقی و عمودی با CSS variable کنترل می‌شود */
  transform: translate(-50%, var(--img-y, 0px)) translateX(var(--img-translate, 0px));
  width: 120%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s ease;
  opacity: 0.95;
  object-position: center bottom; /* تلاش برای نمایش پایین تصویر */
}

/* ensure form/card sits above mobile bg */
.form-card {
  z-index: 5;
  position: relative;
}

/* mobile logo on the form stays above everything */
.mobile-logo {
  z-index: 10;
}
</style>
