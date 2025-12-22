<template>
  <div class="flex flex-col min-h-screen bg-[#f1f2f2] font-mont no-transform">
    <!-- Navbar -->
    <navbar />

    <!-- Header -->
    <main class="flex-1 px-4 py-16 safe-bottom-space">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-12">
          <h1 class="text-5xl md:text-6xl font-light text-[#848484] mb-6">About us</h1>
          <!-- پنهان‌سازی متن سرصفحه در موبایل تا از تکرار جلوگیری شود -->
          <p class="text-[#848484] text-lg leading-relaxed hidden sm:block">
            <span class="font-semibold">POLYCHEM</span> — Advanced Polymer Engineering
            <br /><br />
            PolyChem is a specialist producer of advanced polymer compounds and masterbatches based in the Aras Free Zone. Combining years of industrial experience with the expertise of leading engineers and researchers, we develop high-performance materials tailored for modern manufacturing needs.
            <br /><br />
            <span class="font-semibold">What We Do</span>
            <br />
            • Advanced polymer compounds & masterbatches
            <br />
            • Technical know-how development & transfer
            <br />
            • Engineering and design support for polymer production
          </p>
        </div>

        <!-- About Us Component -->
        <aboutus />
      </div>
    </main>

    <!-- Footer -->
    <footer class="w-full">
      <!-- Chat widget (always present on page) -->
      <Chatbox />
      <footer2 />
    </footer>
  </div>
</template>

<script setup>
import navbar from '~/components/navbar.vue'
import footer2 from '~/components/footer.vue'
import aboutus from '~/components/Aboutus.vue'
import Chatbox from '~/components/chatbox.vue'
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@100;200;300;400;500;600;700;800;900&display=swap');

.font-mont {
  font-family: 'Montserrat', sans-serif;
  color: #848484;
}

/* remove any transform that might cause fixed to become scoped to a container */
.no-transform {
  transform: none !important;
  -webkit-transform: none !important;
}

/* provide bottom safe space equal to navbar height + safe-area so content doesn't push/overlap visually */
.safe-bottom-space {
  padding-bottom: calc(96px + env(safe-area-inset-bottom, 0px)); /* 80px navbar + 16px gap */
}
</style>
