<!-- Footer Component - با نقشه Leaflet Dark Mode -->
<template>
  <footer class="relative w-full bg-[#808285] text-white overflow-hidden font-mono">
    <!-- Decorative Top Border -->
    <div class="absolute top-0 left-0 w-full h-[2px] bg-gradient-to-r from-transparent via-white to-transparent opacity-30"></div>
    
    <!-- Main Footer Content -->
    <div class="max-w-[1400px] mx-auto px-4 sm:px-6 md:px-8 py-12 sm:py-16 md:py-20">
      <!-- Top Section -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 sm:gap-10 md:gap-12 lg:gap-16 mb-12 sm:mb-14 md:mb-16">
        <!-- Brand Section -->
        <div class="space-y-4 sm:space-y-5 md:space-y-6 sm:col-span-2 lg:col-span-1">
          <img 
            src="/english logo W.png" 
            :alt="companyName + ' Logo'" 
            class="h-12 sm:h-14 md:h-16 w-auto object-contain"
          >
          <p class="text-white leading-relaxed text-sm sm:text-base max-w-sm">
            Specialist producer of advanced polymer compounds and masterbatches based in the Aras Free Zone.  
          </p>
          
          <!-- Contact Section -->
          <div class="space-y-3">
            <h3 class="text-sm sm:text-base font-semibold uppercase tracking-wider">Contact</h3>
            
            <!-- Email -->
            <div class="flex items-center gap-2 sm:gap-3 text-white text-xs sm:text-sm">
              <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
              </svg>
              <a :href="`mailto:${contactEmail}`" class="footer-nav-link hover:text-white transition-colors break-all">
                {{ contactEmail }}
              </a>
            </div>
            
            <!-- Phone -->
            <div class="flex items-center gap-2 sm:gap-3 text-white text-xs sm:text-sm">
              <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"/>
              </svg>
              <a :href="`tel:${phone1}`" class="footer-nav-link hover:text-white transition-colors">{{ phone1 }}</a>
            </div>
                <!-- Phone -->
            <div class="flex items-center gap-2 sm:gap-3 text-white text-xs sm:text-sm">
              <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"/>
              </svg>
              <a :href="`tel:${phone2}`" class="footer-nav-link hover:text-white transition-colors">{{ phone2 }}</a>
            </div>
            <!-- Instagram -->
            <div class="flex items-center gap-2 sm:gap-3 text-white text-xs sm:text-sm">
              <svg class="w-4 h-4 flex-shrink-0" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/>
              </svg>
              <a href="http://instagram.com/polychemmb" class="footer-nav-link hover:text-white transition-colors">@Polychemmb</a>
            </div>
            
            <!-- LinkedIn -->
            <div class="flex items-center gap-2 sm:gap-3 text-white text-xs sm:text-sm">
              <svg class="w-4 h-4 flex-shrink-0" fill="currentColor" viewBox="0 0 24 24">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
              </svg>
              <a href="https://www.linkedin.com/company/polychem-turkey/" class="footer-nav-link hover:text-white transition-colors">Polychem</a>
            </div>
          </div>
          
          <!-- Social Media Icons -->
          <div class="flex gap-3 sm:gap-4 pt-2">
            <!-- Instagram Icon -->
            <a 
              href="http://instagram.com/polychemmb" 
              class="instagram-icon w-9 h-9 sm:w-10 sm:h-10 rounded-full border border-white flex items-center justify-center hover:border-white hover:bg-white transition-all duration-300 active:scale-95"
              aria-label="Instagram"
            >
              <svg class="w-4 h-4 sm:w-5 sm:h-5 instagram-svg" fill="white" viewBox="0 0 24 24">
                <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/>
              </svg>
            </a>
            
            <!-- YouTube Icon -->
            <a 
              href="#" 
              class="youtube-icon w-9 h-9 sm:w-10 sm:h-10 rounded-full border border-white flex items-center justify-center hover:border-white hover:bg-white hover:text-black transition-all duration-300 active:scale-95"
              aria-label="YouTube"
            >
              <svg class="w-4 h-4 sm:w-5 sm:h-5 youtube-svg" fill="white" viewBox="0 0 24 24">
                <path d="M23.498 6.186a2.994 2.994 0 0 0-2.108-2.117C19.633 3.5 12 3.5 12 3.5s-7.633 0-9.39.569A2.994 2.994 0 0 0 .502 6.186C0 7.943 0 12 0 12s0 4.057.502 5.814a2.994 2.994 0 0 0 2.108 2.117C4.367 20.5 12 20.5 12 20.5s7.633 0 9.39-.569a2.994 2.994 0 0 0 2.108-2.117C24 16.057 24 12 24 12s0-4.057-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/>
              </svg>
            </a>
            
            <a 
              href="https://www.linkedin.com/company/polychem-turkey/" 
              class="linkedin-icon w-9 h-9 sm:w-10 sm:h-10 rounded-full border border-white flex items-center justify-center hover:border-white hover:bg-white transition-all duration-300 active:scale-95"
              aria-label="LinkedIn"
            >
              <svg class="w-4 h-4 sm:w-5 sm:h-5 linkedin-svg" fill="white" viewBox="0 0 24 24">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
              </svg>
            </a>
            
            <!-- Email Icon -->
            <a 
              href="mailto:info@polychemmb.com" 
              class="email-icon w-9 h-9 sm:w-10 sm:h-10 rounded-full border border-white flex items-center justify-center hover:border-white hover:bg-white transition-all duration-300 active:scale-95"
              aria-label="Email"
            >
              <svg class="w-4 h-4 sm:w-5 sm:h-5 email-svg" fill="none" stroke="white" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
              </svg>
            </a>
          </div>
        </div>

        <!-- Navigation -->
        <div class="sm:flex sm:justify-center lg:justify-center">
          <div>
            <h3 class="text-base sm:text-lg font-semibold mb-4 sm:mb-5 md:mb-6 uppercase tracking-wider">Navigation</h3>
            <ul class="space-y-2 sm:space-y-3">
              <li><a href="/" class="footer-nav-link text-white transition-colors duration-300 text-sm sm:text-base inline-block">HOME</a></li>
              <li><a href="/about" class="footer-nav-link text-white transition-colors duration-300 text-sm sm:text-base inline-block">ABOUT US</a></li>
              <li><a href="/news" class="footer-nav-link text-white transition-colors duration-300 text-sm sm:text-base inline-block">NEWS</a></li>
              <li><a href="/careers" class="footer-nav-link text-white transition-colors duration-300 text-sm sm:text-base inline-block">CAREERS</a></li>
              <li><a href="/contact" class="footer-nav-link text-white transition-colors duration-300 text-sm sm:text-base inline-block">CONTACT US</a></li>
            </ul>
          </div>
        </div>

        <!-- Address Section -->
        <div class="sm:col-span-2 lg:col-span-1">
          <h3 class="text-base sm:text-lg font-semibold mb-4 sm:mb-5 md:mb-6 uppercase tracking-wider">Our Locations</h3>
          <ul class="space-y-3 sm:space-y-4 mb-4">
            <li class="flex items-start gap-2 sm:gap-3 text-white text-sm sm:text-base">
              <svg class="w-4 h-4 sm:w-5 sm:h-5 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/>
              </svg>
              <span> Plots 437 & 438,South C Street<br>Industrial Town Phase 2<br>Aras Free Trade-Industrial Zone<br></span>
            </li>
            <li class="flex items-start gap-2 sm:gap-3 text-white text-sm sm:text-base">
              <svg class="w-4 h-4 sm:w-5 sm:h-5 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/>
              </svg>
              <span>Unit 15-NO.45 – Manzarnejad Blvd<br>Shariati Ave- Mirdamad Blvd<br> Tehran- Iran</span>
            </li>
          </ul>
          
          <!-- Footer Image -->
          <!-- Footer Image - از راست چسبیده، به چپ بزرگ شود تا نزدیک CONTACT US -->
<!-- Footer Image - بیشتر به چپ، بزرگ‌تر در طول و ارتفاع -->
<div class="mt-6 md:mt-14 -mr-8 lg:-mr-30">
  <img 
    src="/Footer.png" 
    alt="Polychem Footer Illustration"
    class="block ml-auto footer-image"
  >
</div>
        </div>
      </div>

      <!-- Newsletter Section -->
      <div class="border-t border-b border-white py-8 sm:py-10 md:py-12">
        <div class="max-w-2xl mx-auto text-center">
          <h3 class="text-xl sm:text-2xl font-semibold mb-2 sm:mb-3">Let's Work Together</h3>
          <p class="text-white mb-5 sm:mb-6 text-sm sm:text-base px-4">Get notified about new projects and professional updates</p>
          <form class="flex flex-col sm:flex-row gap-3 max-w-md mx-auto px-4">
            <input 
              type="email" 
              placeholder="Your email address"
              class="flex-1 px-4 sm:px-5 py-2.5 sm:py-3 bg-white/5 border border-white rounded-lg focus:outline-none focus:border-white transition-colors text-sm sm:text-base"
            >
            <button type="submit" class="card-button">
              <span>Notify Me</span>
            </button>
          </form>
        </div>
      </div>

      <!-- Bottom Section -->
      <div class="pt-8 sm:pt-10 md:pt-12 flex flex-col md:flex-row justify-between items-center gap-4 sm:gap-6">
        <p class="text-white text-xs sm:text-sm text-center md:text-left order-2 md:order-1">
          © 2025 {{ companyName }}. All rights reserved.
        </p>
        <div class="flex flex-wrap justify-center gap-3 sm:gap-6 text-xs sm:text-sm order-1 md:order-2">
          <a href="#" class="footer-nav-link text-white transition-colors whitespace-nowrap">Legal Notice</a>
          <span class="text-white">|</span>
          <a href="#" class="footer-nav-link text-white transition-colors whitespace-nowrap">Cookie Policy</a>
          <span class="text-white">|</span>
          <a href="#" class="footer-nav-link text-white transition-colors whitespace-nowrap">Privacy</a>
        </div>
      </div>
    </div>

    <!-- Decorative Bottom Element -->
    <div class="absolute bottom-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-white/20 to-transparent"></div>
  </footer>
</template>

<script setup>
import { ref } from 'vue'

const mapElement = ref(null)

// read site-provided contact info from <html data-...>
// developers can set these attributes on the page:
// <html data-contact-email="..." data-phone1="+98..." data-phone2="+98..." data-company-name="...">
const ds = typeof document !== 'undefined' ? document.documentElement.dataset : {}
const contactEmail = ref(ds.contactEmail || ds.contact_email || 'info@polychemmb.com')
const phone1 = ref(ds.phone1 || ds.phone_main || '+982122898979')
const phone2 = ref(ds.phone2 || ds.phone_secondary || '+982122898980')
const companyName = ref(ds.companyName || ds.company_name || 'Polychem')
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;600;700&family=Figtree:wght@300;400;500;600;700&display=swap');

footer {
  font-family: 'Figtree', sans-serif;
}

footer h2,
footer h3,
footer .text-2xl,
footer .text-xl,
footer .font-semibold {
  font-family: 'Montserrat', sans-serif;
}

a, button {
  -webkit-tap-highlight-color: transparent;
}

.card-button {
  background: transparent;
  color: #FFCD05;
  border: 2px solid #FFCD05;
  padding: 12px 28px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  margin: 0;
  text-transform: uppercase;
  letter-spacing: 0.04em;
  z-index: 1;
  transition: color 200ms ease;
}

.card-button::before {
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

.card-button:hover::before,
.card-button:focus-visible::before {
  transform: translateY(0);
}

.card-button:hover,
.card-button:focus-visible {
  color: #ffffff;
  outline: none;
}

.card-button span {
  position: relative;
  z-index: 1;
}

.youtube-icon:hover .youtube-svg {
  fill: #000 !important;
}

.instagram-icon:hover .instagram-svg {
  fill: #000 !important;
}

.linkedin-icon:hover .linkedin-svg {
  fill: #000 !important;
}

.email-icon:hover .email-svg {
  stroke: #000 !important;
}

.footer-nav-link {
  position: relative;
  color: #fff;
  text-decoration: none;
  transition: color 0.2s;
}

.footer-nav-link::before {
  content: '';
  position: absolute;
  left: 0;
  bottom: 0;
  width: 0;
  height: 2px;
  background: #FFCD05;
  transition: width 0.3s;
}

.footer-nav-link:hover::before,
.footer-nav-link:focus-visible::before {
  width: 100%;
}

.footer-nav-link:hover,
.footer-nav-link:focus-visible {
  color: #fff;
}

.footer-image {
  height: 440px !important;
  width: 400px;
  transform: scaleX(2.45) !important;
  transform-origin: right top !important;
  max-width: none !important;
}

@media (max-width: 1024px) {
  .footer-image {
    height: 220px !important;
    width: 500px;
    transform: scaleX(1) !important;
  }
}
</style>