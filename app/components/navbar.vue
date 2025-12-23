<template>
  <div>
    <nav class="bg-[#808285] fixed top-0 inset-x-0 z-50 md:border-b" >
      <div class="navbar-inner navbar-container">
        <div class="flex items-center h-16 md:h-20">
          <!-- Logo & Brand -->
          <div class="flex items-center justify-start md:border-r md:border-l md:border-white pl-2 md:pl-8 pr-2 md:pr-8 h-full flex-shrink-0 w-auto md:w-[230px]">
              <NuxtLink to="/" aria-label="Go to home" class="block">
                <img src="/english logo W1.png" alt="POLYCHEM Logo" class="w-28 md:w-40 h-28 md:h-40 object-contain" />
              </NuxtLink>
          </div>

          <!-- USE CASES with border and hover effect -->
          <!-- show PRODUCTS button on desktop only (hidden on mobile) -->
          <div class="hidden md:flex items-center h-full md:border-r md:border-white">
            <button 
              @click="toggleFullscreenMenu"
              class="nav-link flex items-center gap-2 text-[13px] text-white font-medium px-8 h-full text-white relative overflow-hidden"
              :class="{ active: isMenuOpen }"
            >
              <span class="relative z-10">PRODUCTS</span>
              <svg class="w-4 h-4 relative z-10" fill="currentColor" viewBox="0 0 20 20">
                <path d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"/>
              </svg>
            </button>
          </div>

          <!-- Main Menu with hover effects -->
          <div class="hidden md:flex items-center flex-1 justify-center h-full">
            <NuxtLink 
              to="/" 
              class="nav-link text-[13px] font-medium px-8 text-white relative overflow-hidden h-full flex items-center"
            >
              <span class="relative z-10">HOME</span>
            </NuxtLink>

            <NuxtLink 
              to="/about" 
              class="nav-link flex items-center justify-center gap-2 text-[13px] font-medium px-8 text-white relative overflow-hidden h-full w-40 "
            >
              <span class="relative z-10">ABOUT US</span>
            </NuxtLink>

            <NuxtLink 
              to="/news" 
              class="nav-link text-[13px] font-medium px-8 text-white relative overflow-hidden h-full flex items-center"
            >
              <span class="relative z-10">NEWS</span>
            </NuxtLink>

            <NuxtLink 
              to="/careers" 
              class="nav-link text-[13px] font-medium px-8 text-white relative overflow-hidden h-full flex items-center"
            >
              <span class="relative z-10">CAREERS</span>
            </NuxtLink>

            <NuxtLink 
              to="/contact" 
              class="nav-link text-[13px] font-medium px-8 text-white relative overflow-hidden h-full flex items-center w-40"
            >
              <span class="relative z-10">CONTACT US</span>
            </NuxtLink>
          </div>

          <!-- Right Side -->
          <div class="hidden md:flex items-center justify-end h-full border-l border-white w-[200px] flex-shrink-0">
            <button class="nav-link text-[13px] font-medium px-8 h-full text-white relative overflow-hidden">
              <span class="relative z-10">EN</span>
            </button>
             <NuxtLink to="/login" class="nav-link text-[13px] font-medium px-8 h-full text-white relative overflow-hidden w-40 border-l border-r border-white flex items-center justify-center">
                <span class="relative z-10">LOG IN</span>
              </NuxtLink>
          </div>

          <!-- Mobile Menu Button (REPLACED): show PRODUCTS as the right-most item on mobile -->
          <button 
            @click="toggleFullscreenMenu"
            class="md:hidden ml-auto text-white hover:text-yellow-400 px-4 py-3 rounded-lg hover:bg-black/20 transition-all duration-200 flex items-center gap-2"
            aria-label="Open products"
          >
            <span class="text-[13px] font-medium">PRODUCTS</span>
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"/>
            </svg>
          </button>
        </div>

        <!-- Mobile Menu -->
        <transition name="mobile-menu">
          <!-- removed border on mobile: no border classes here -->
          <div v-if="mobileMenuOpen" class="md:hidden pb-4 bg-[#808285] max-h-[calc(100vh-64px)] overflow-y-auto">
            <div class="flex flex-col gap-2 pt-4 px-4">
              <button 
                @click="toggleFullscreenMenu"
                class="nav-link-mobile text-[13px] font-medium text-white hover:text-gray-200 py-3 px-4 rounded relative overflow-hidden text-left w-full"
                :class="{ active: isMenuOpen }"
              >
                <span class="relative z-10">PRODUCTS</span>
              </button>
              <NuxtLink to="/" class="nav-link-mobile text-[13px] font-medium text-white hover:text-gray-200 py-3 px-4 rounded relative overflow-hidden w-full block">
                <span class="relative z-10">HOME</span>
              </NuxtLink>
              <NuxtLink to="/about" class="nav-link-mobile text-[13px] font-medium text-white hover:text-gray-200 py-3 px-4 rounded relative overflow-hidden w-full block">
                <span class="relative z-10">ABOUT US</span>
              </NuxtLink>
              <NuxtLink to="/news" class="nav-link-mobile text-[13px] font-medium text-white hover:text-gray-200 py-3 px-4 rounded relative overflow-hidden w-full block">
                <span class="relative z-10">NEWS</span>
              </NuxtLink>
              <NuxtLink to="/careers" class="nav-link-mobile text-[13px] font-medium text-white hover:text-gray-200 py-3 px-4 rounded relative overflow-hidden w-full block">
                <span class="relative z-10">CAREERS</span>
              </NuxtLink>
              <NuxtLink to="/contact" class="nav-link-mobile text-[13px] font-medium text-white hover:text-gray-200 py-3 px-4 rounded relative overflow-hidden w-full block">
                <span class="relative z-10">CONTACT US</span>
              </NuxtLink>
              <div class="pt-3 mt-2">
                <button class="nav-link-mobile text-[13px] font-medium text-white hover:text-gray-200 py-3 px-4 rounded block w-full text-left relative overflow-hidden mb-2">
                  <span class="relative z-10">EN</span>
                </button>
                <NuxtLink to="/login" class="nav-link-mobile text-[13px] font-medium text-white hover:text-gray-200 py-3 px-4 rounded block w-full text-left relative overflow-hidden">
                  <span class="relative z-10">LOG IN</span>
                </NuxtLink>
              </div>
            </div>
          </div>
        </transition>
      </div>
    </nav>
    <!-- Spacer to offset fixed navbar height -->
    <div class="h-16 md:h-20"></div>

    <!-- Mobile bottom navbar (active on small screens only) -->
    <NavbarMob class="md:hidden" />

    <!-- Render Chatbox on Home and About pages -->
    <Chatbox v-if="route.path === '/' || route.path === '/about'" />

    <!-- Fullscreen Menu - Mobile Optimized -->
    <div 
      v-show="isMenuOpen"
      ref="fullscreenMenu"
      class="fixed inset-0 z-40 overflow-hidden"
    >
      <div class="h-full flex flex-col md:flex-row">
        <!-- Left Side - Preview Box (hidden on mobile, 40% on desktop) -->
        <div class="hidden md:flex w-[40%] bg-gradient-to-br from-gray-50 to-white relative overflow-hidden">
          <div 
            ref="previewBox"
            class="absolute inset-0 flex items-center justify-center p-20 opacity-0"
          >
            <div class="text-center space-y-6">
              <!-- changed: show preview product image instead of static icon -->
              <img
                ref="previewImage"
                :src="hoveredItem?.image || ''"
                alt=""
                class="w-1500 h-1500   object-cover transform scale-0"
                :style="previewImageStyle"
              />
            
              <!-- backdrop behind title & description (bigger, semi-transparent) -->
              <div class="inline-block bg-white/70 backdrop-blur-sm rounded-2xl px-8 py-6 max-w-[90%]">
                <div 
                  ref="previewTitle"
                  class="opacity-0 transform translate-y-4"
                  :style="previewTitleStyle"
                >
                  <h3 class="text-3xl md:text-4xl font-bold text-gray-900 mb-3">{{ hoveredItem?.title }}</h3>
                  <p class="text-xl text-gray-600">{{ hoveredItem?.code || '' }}</p>
                </div>
                <p 
                  ref="previewDesc"
                  class="text-gray-700 text-lg max-w-md mx-auto opacity-0 transform translate-y-4 mt-3"
                  :style="previewDescStyle"
                >
                  {{ hoveredItem?.description || '' }}
                </p>
              </div>
            </div>

            <!-- Decorative Grid -->
            <div class="absolute inset-0 opacity-5">
              <div class="grid grid-cols-8 grid-rows-8 h-full w-full">
                <div v-for="i in 64" :key="i" class="border border-gray-900"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Side - Menu Items (100% on mobile, 60% on desktop) -->
        <div class="w-full md:w-[60%] bg-white relative flex flex-col h-full">
          <div 
            ref="menuContainer"
            class="h-full flex flex-col justify-start md:justify-center px-6 md:px-16 py-6 md:py-0 opacity-0 transform translate-y-[-100%] overflow-y-auto md:overflow-y-visible"
          >
            <!-- Close Button (HIDDEN ON MOBILE) -->
            <button 
              @click="toggleFullscreenMenu"
              ref="closeButton"
              class="hidden md:flex absolute top-6 right-6 text-gray-700 hover:text-gray-900 transition-colors opacity-0 md:top-8 md:right-8 md:w-10 md:h-10 w-8 h-8 flex items-center justify-center"
              aria-label="Close menu"
            >
              <svg class="w-6 h-6 md:w-10 md:h-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>

            <!-- Title -->
            <div 
              ref="menuTitle"
              class="mb-8 md:mb-12 mt-12 md:mt-16 opacity-0 transform translate-y-[-20px]"
            >
              <h2 class="text-3xl md:text-5xl font-bold text-gray-900 mb-2">Our Products</h2>
              <div class="w-20 h-1 bg-yellow-500"></div>
            </div>

            <!-- Menu Items -->
            <div ref="itemsWrapper" class="space-y-1 max-h-[calc(100vh-200px)] md:max-h-[70vh] overflow-y-auto pr-4 custom-scrollbar pb-6 md:pb-0 menu-items-scroll">
              <button
                v-for="(item, index) in useCases"
                :key="index"
                :ref="el => { if (el) menuItems[index] = el }"
                @mouseenter="handleHover(item, index)"
                @mouseleave="handleLeave"
                @click="() => handleMenuItemClick(item)"
                type="button"
                class="menu-item group flex items-center justify-between border-b border-gray-200 py-4 md:py-6 px-4 md:px-6 hover:bg-gray-50 transition-all duration-300 transform translate-y-[-30px] active:bg-gray-100 md:active:bg-transparent"
                style="opacity: 1 !important; min-height: 48px; width: 100%; text-align: left;"
              >
                <div class="flex items-center gap-3 md:gap-6 menu-item-content flex-1 min-w-0" style="opacity: 1 !important; color: #1f2937 !important;">
                  <span style="color: #6b7280 !important; opacity: 1 !important;" class="text-[11px] md:text-[13px] font-mono group-hover:text-yellow-500 transition-colors flex-shrink-0">{{ item.code }}</span>
                  <h3 style="color: #1f2937 !important; opacity: 1 !important;" class="text-lg md:text-xl font-semibold group-hover:translate-x-2 transition-transform truncate">{{ item.title }}</h3>
                </div>
                <svg class="w-5 h-5 md:w-6 md:h-6 text-gray-400 group-hover:text-yellow-500 group-hover:translate-x-2 transition-all flex-shrink-0 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3"/>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick, computed } from 'vue'
import gsap from 'gsap'
import NavbarMob from './navbar-mob.vue'
import Chatbox from './chatbox.vue'
import { useRoute, useRouter } from 'vue-router'

const mobileMenuOpen = ref(false)
const fullscreenMenu = ref(null)
const menuContainer = ref(null)
const menuTitle = ref(null)
const closeButton = ref(null)
const previewBox = ref(null)
const previewIcon = ref(null) // keep for backward-compat if used elsewhere
const previewImage = ref(null)
const previewTitle = ref(null)
const previewDesc = ref(null)
const itemsWrapper = ref(null)
const menuItems = ref([])
const isMenuOpen = ref(false)
const hoveredItem = ref(null)
const route = useRoute()
const router = useRouter()
// handlers for wheel/touch to prevent background scrolling
const wheelHandler = (e) => {
  // allow the event to scroll the itemsWrapper but stop it bubbling to body
  e.stopPropagation()
}
const touchHandler = (e) => {
  e.stopPropagation()
}

const useCases = [
  { 
    code: '', 
    title: 'ROTOCHEM 0955W', 
    description: 'Specialized white color plastic compound for rotational molding applications. Linear medium density polyethylene copolymer with narrow molecular weight distribution and TiO2. Ideal for reliable and durable materials that withstand harsh environmental conditions.',
    image: '/Roto.webp',
    index: 0,
    offsetY: '15px',
    offsetTextY: '60px',
    offsetDescY: '130px'
  },
  { 
    code: '', 
    title: 'ROTOCHEM 0955B', 
    description: 'Specialized blue color plastic compound for rotational molding applications. Linear medium density polyethylene copolymer with narrow molecular weight distribution with TiO2. Perfect for industrial and consumer applications requiring consistent and stable materials.',
    image: '/955B.webp',
    index: 1,
    offsetY: '12px',
    offsetTextY: '4px',
    offsetDescY: '12px'
  },
  { 
    code: '', 
    title: 'POLYFIL F700', 
    description: 'High-performance polyethylene compound for HDPE blown film applications. Engineered for superior mechanical properties, excellent film uniformity, and reliable processability in ultra-thin film applications (10-25 microns). Suitable for shopping bags, T-shirt bags, garbage bags, and food-contact films.',
    image: '/f700.webp',
    index: 2,
    offsetY: '18px',
    offsetTextY: '6px',
    offsetDescY: '60px'
  },
  { 
    code: '', 
    title: 'HDCHEM 4760', 
    description: 'Specialized plastic compound for blow molding applications. Polyethylene copolymer compound with beneficial characteristics including good flowability, impact strength, ESCR and rigidity. Dependable choice for consistent and stable materials.',
    image: '/Hdchem-1.webp',
    index: 3,
   offsetY: '18px',
    offsetTextY: '6px',
    offsetDescY: '50px'    
  },
  { 
    code: '', 
    title: 'SlIPCHEM E 178', 
    description: 'High-performance slip masterbatch with high quality slip agent dispersed in polyethylene carrier resin. Reduces coefficient of friction (COF) between polymer film layers during winding, bag-making, and packaging. Excellent dispersion, high thermal stability, and consistent migration performance.',
    image: '/slipchem-2.webp',
    index: 4,
    offsetY: '18px',
    offsetTextY: '6px',
    offsetDescY: '50px' 
  },
  { 
    code: '', 
    title: 'RAFCOLOR 1560', 
    description: 'White masterbatch with high proportion of rutile titanium dioxide and thermoplastic polypropylene resin. Highly-concentrated white MB with excellent dispersion and thermal stability. Recommended for raffia, tapes, CF/BCF yarn, and other products.',
    image: '/rofcolor.webp',
    index: 5,
    offsetDescY: '90px',
        offsetTextY: '50px',
    offsetY: '20px'
    
  },
  { 
    code: '', 
    title: 'CALCICHEM 126 FP', 
    description: 'Polypropylene-based filler masterbatch containing 80% calcium carbonate (CaCO₃). Features high, very fine, treated CaCO₃ content ensuring excellent dispersion. Designed for direct addition during processing of polyolefins including extrusion and injection molding.',
    image: '/126.webp',
    index: 6,
      offsetDescY: '50px',
        offsetTextY: '50px',
    offsetY: '20px'
  },
  { 
    code: '', 
    title: 'CALCICHEM 110 FRF', 
    description: 'Mineral modifier with high, very fine, treated CaCO3 content with excellent dispersion. Designed for films, raffia and ropes. Suitable for general-purpose products with PE and PP carriers. Increases productivity and reduces raw material costs.',
    image: '/110-FRF-2.webp',
    index: 7,
         offsetDescY: '50px',
        offsetTextY: '50px',
    offsetY: '20px'
 
  },
  { 
    code: '', 
    title: 'CALCICHEM 275 PM', 
    description: 'Polypropylene-based mineral masterbatch containing 75% ultra-fine mineral filler. Offers excellent dispersion and high mineral loading for enhanced performance. Formulated for direct addition during extrusion of BOPP, CPP, and OPP films.',
    image: '/275-1.webp',
    index: 8,
      offsetDescY: '70px',
        offsetTextY: '50px',
    offsetY: '20px'
  },
  { 
    code: '', 
    title: 'UVCHEM MB-R18', 
    description: 'High-performance UV stabilizer masterbatch for raffia and woven fabric. Contains HALS and UV absorbers for long-term resistance to photo-oxidation, color fading, and mechanical property loss. Excellent dispersion, thermal stability, and compatibility with PP matrix.',
    image: '/uvchem-2.webp',
    index: 9,
    offsetY: '24px',
    offsetTextY: '8px',
    offsetDescY: '22px'
  },
]

const handleHover = (item, index) => {
  hoveredItem.value = item
  
  const tl = gsap.timeline()
  
  // Animate preview content change
  tl.to([previewTitle.value, previewDesc.value], {
    opacity: 0,
    y: -10,
    duration: 0.2
  })
  .set([previewTitle.value, previewDesc.value], {
    opacity: 0
  })
  .to([previewTitle.value, previewDesc.value], {
    opacity: 1,
    y: 0,
    duration: 0.4,
    stagger: 0.1
  })

  // Animate preview image (fade + pop-in)
  if (previewImage.value) {
    gsap.fromTo(previewImage.value, 
      { opacity: 0, scale: 0.85 }, 
      { opacity: 1, scale: 1, duration: 0.45, ease: 'back.out(1.2)' }
    )
  }
}

const handleLeave = () => {
}

const toggleFullscreenMenu = async () => {
  if (!fullscreenMenu.value) return
  
  if (isMenuOpen.value) {
    // run closing animation, then hide DOM after animation completes
    closeMenu()
    // hide the element after the close animation (adjust if you change durations)
    setTimeout(() => {
      isMenuOpen.value = false
    }, 800)
  } else {
    // make the DOM visible first, then run animations
    isMenuOpen.value = true
    await nextTick()
    openMenu()
  }
}

const openMenu = () => {
  const menu = fullscreenMenu.value
  
  // Enable pointer events
  menu.style.pointerEvents = 'auto'
  // prevent background page scrolling while menu is open
  try { document.body.style.overflow = 'hidden' } catch (err) {}
  
  // Create timeline
  const tl = gsap.timeline()
  
  // Animate menu container sliding down from top
  tl.to(menuContainer.value, {
    y: 0,
    opacity: 1,
    duration: 0.8,
    ease: 'power3.out'
  })
  
  // Animate title
  tl.to(menuTitle.value, {
    opacity: 1,
    y: 0,
    duration: 0.5,
    ease: 'power2.out'
  }, '-=0.5')
  
  // Animate close button
  tl.to(closeButton.value, {
    opacity: 1,
    duration: 0.3
  }, '-=0.3')
  
  // Animate menu items
  tl.to(menuItems.value, {
    opacity: 1,
    y: 0,
    duration: 0.5,
    ease: 'power2.out',
    stagger: 0.04
  }, '-=0.5')
  
  // Animate preview box
  tl.to(previewBox.value, {
    opacity: 1,
    duration: 0.6
  }, '-=0.8')
  
  // animate preview image into view (if present)
  if (previewImage.value) {
    tl.to(previewImage.value, {
      scale: 1,
      opacity: 1,
      duration: 0.6,
      ease: 'back.out(1.7)'
    }, '-=0.4')
  } else {
    tl.to(previewIcon.value, {
      scale: 1,
      duration: 0.6,
      ease: 'back.out(1.7)'
    }, '-=0.4')
  }
  
  tl.to([previewTitle.value, previewDesc.value], {
    opacity: 1,
    y: 0,
    duration: 0.5,
    stagger: 0.1
  }, '-=0.3')
}

const closeMenu = () => {
  const menu = fullscreenMenu.value
  
  // Create timeline
  const tl = gsap.timeline()
  
  // Animate menu items out
  tl.to(menuItems.value, {
    opacity: 0,
    y: -20,
    duration: 0.3,
    ease: 'power2.in',
    stagger: 0.02
  })
  
  // Animate title out
  tl.to(menuTitle.value, {
    opacity: 0,
    y: -20,
    duration: 0.3
  }, '-=0.2')
  
  // Animate preview box out
  tl.to([previewTitle.value, previewDesc.value, previewIcon.value, previewImage.value], {
    opacity: 0,
    scale: 0.9,
    duration: 0.3
  }, '-=0.3')
  
  // Animate close button out
  tl.to(closeButton.value, {
    opacity: 0,
    duration: 0.2
  }, '-=0.2')
  
  // Slide menu container up
  tl.to(menuContainer.value, {
    y: '-100%',
    opacity: 0,
    duration: 0.6,
    ease: 'power3.in'
  }, '-=0.2')
  
  tl.to(previewBox.value, {
    opacity: 0,
    duration: 0.3,
    onComplete: () => {
      menu.style.pointerEvents = 'none'
      hoveredItem.value = null
      // restore body scrolling and remove event listeners
      try { document.body.style.overflow = '' } catch (err) {}
      if (itemsWrapper.value) {
        itemsWrapper.value.removeEventListener('wheel', wheelHandler)
        itemsWrapper.value.removeEventListener('touchmove', touchHandler)
      }
    }
  }, '-=0.4')
}

onMounted(async () => {
  // Wait for refs to be set
  await nextTick()
  // Set initial state
  if (menuItems.value && menuItems.value.length > 0) {
    gsap.set(menuItems.value, { y: -30, opacity: 0 })
  }
  if (menuTitle.value) {
    gsap.set(menuTitle.value, { y: -20, opacity: 0 })
  }
  if (previewTitle.value && previewDesc.value) {
    gsap.set([previewTitle.value, previewDesc.value], { y: 4, opacity: 0 })
  }
  if (previewIcon.value) {
    gsap.set(previewIcon.value, { scale: 0 })
  }
  if (previewImage.value) {
    gsap.set(previewImage.value, { scale: 0.85, opacity: 0 })
  }

  // ensure wheel/touch scrolling works inside the items wrapper
  if (itemsWrapper.value) {
    itemsWrapper.value.style.overscrollBehavior = 'contain'
    itemsWrapper.value.style.touchAction = 'auto'
    // attach handlers so wheel/touch don't bubble to body
    itemsWrapper.value.addEventListener('wheel', wheelHandler, { passive: false })
    itemsWrapper.value.addEventListener('touchmove', touchHandler, { passive: false })
  }
})

onUnmounted(() => {
  // cleanup listeners
  try { document.body.style.overflow = '' } catch (err) {}
  if (itemsWrapper.value) {
    itemsWrapper.value.removeEventListener('wheel', wheelHandler)
    itemsWrapper.value.removeEventListener('touchmove', touchHandler)
  }
})

// Close menu on Escape key
onMounted(() => {
  const handleEscape = (e) => {
    if (e.key === 'Escape' && isMenuOpen.value) {
      toggleFullscreenMenu()
    }
  }
  
  window.addEventListener('keydown', handleEscape)
  
  onUnmounted(() => {
    window.removeEventListener('keydown', handleEscape)
  })
})

const previewImageStyle = computed(() => {
  const y = hoveredItem.value?.offsetY ?? '0px'
  return {
    position: 'relative',
    top: y,
    transition: 'top 220ms ease'
  }
})

const previewTitleStyle = computed(() => {
  const y = hoveredItem.value?.offsetTextY ?? '0px'
  return {
    position: 'relative',
    top: y,
    transition: 'top 220ms ease'
  }
})

const previewDescStyle = computed(() => {
  const y = hoveredItem.value?.offsetDescY ?? '0px'
  return {
    position: 'relative',
    top: y,
    transition: 'top 220ms ease'
  }
})

// navigate programmatically to product page with query then close menu
const handleMenuItemClick = async (item) => {
  if (!item) return
  try {
    // ensure index is a string for query
    await router.push({ path: '/product', query: { index: String(item.index) } })
  } catch (err) {
    // silent
  }
  // close the fullscreen menu (this will animate & hide)
  toggleFullscreenMenu()
}
</script>

<style scoped>
/* Button 2 Hover Effect - Desktop */
.nav-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  display: block;
  width: 100%;
  height: 100%;
  z-index: 0;
  background-color: #000;
  transform: scaleY(0.3);
  opacity: 0;
  transition: all 0.3s ease;
}

.nav-link:hover::before {
  opacity: 1;
  background-color: rgb(187, 187, 187);
  transform: scaleY(1);
  transition: transform 0.6s cubic-bezier(0.08, 0.35, 0.13, 1.02), opacity 0.4s;
}

.nav-link:hover {
  color: #ffffff !important;
}

.nav-link:hover span {
  color: #ffffff !important;
}

/* Mobile Menu Links Hover Effect */
.nav-link-mobile::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  display: block;
  width: 100%;
  height: 100%;
  z-index: 0;
  background-color: #000;
  transform: scaleY(0.3);
  opacity: 0;
  transition: all 0.3s ease;
}

.nav-link-mobile:hover::before {
  opacity: 1;
  background-color: #ffcd05;
  transform: scaleY(1);
  transition: transform 0.6s cubic-bezier(0.08, 0.35, 0.13, 1.02), opacity 0.4s;
}

.nav-link-mobile:hover {
  color: #030303 !important;
}

.nav-link-mobile:hover span {
  color: #000000 !important;
}

/* Active state matches hover appearance */
.nav-link.active::before {
  opacity: 1;
  background-color: #ffffff;
  transform: scaleY(1);
}
.nav-link.active,
.nav-link.active span {
  color: #000000 !important;
}

.nav-link-mobile.active::before {
  opacity: 1;
  background-color: #ffffff;
  transform: scaleY(1);
}
.nav-link-mobile.active,
.nav-link-mobile.active span {
  color: #030303 !important;
}

/* Nuxt router active states (desktop) */
.router-link-active.nav-link::before,
.router-link-exact-active.nav-link::before {
  opacity: 1;
  background-color: #ffffff;
  transform: scaleY(1);
}
.router-link-active.nav-link,
.router-link-active.nav-link span,
.router-link-exact-active.nav-link,
.router-link-exact-active.nav-link span {
  color: #000000 !important;
}

/* Nuxt router active states (mobile) */
.router-link-active.nav-link-mobile::before,
.router-link-exact-active.nav-link-mobile::before {
  opacity: 1;
  background-color: #ffcd05;
  transform: scaleY(1);
}
.router-link-active.nav-link-mobile,
.router-link-active.nav-link-mobile span,
.router-link-exact-active.nav-link-mobile,
.router-link-exact-active.nav-link-mobile span {
  color: #030303 !important;
}

/* Mobile Menu Animation */
.mobile-menu-enter-active,
.mobile-menu-leave-active {
  transition: all 0.3s ease;
}

.mobile-menu-enter-from {
  opacity: 0;
  max-height: 0;
}

.mobile-menu-leave-to {
  opacity: 0;
  max-height: 0;
}

/* Menu Item Hover Effect */
.menu-item {
  transform: translateZ(0);
  backface-visibility: hidden;
}

/* Ensure menu item text is always visible - override any opacity issues */
.menu-item {
  color: #1f2937 !important;
}

.menu-item-content {
  opacity: 1 !important;
  position: relative;
  z-index: 10;
  isolation: isolate; /* Create new stacking context */
}

/* Force menu items to be visible even if parent has low opacity */
.menu-item {
  isolation: isolate; /* Create new stacking context */
  will-change: opacity, transform; /* Optimize for animations */
}

.menu-item h3 {
  color: #1f2937 !important; /* text-gray-900 */
  position: relative;
  z-index: 10;
  opacity: 1 !important; /* Force visibility */
}

.menu-item span {
  color: #6b7280 !important; /* text-gray-500 */
  position: relative;
  z-index: 10;
  opacity: 1 !important; /* Force visibility */
}

.menu-item svg {
  color: #9ca3af !important; /* text-gray-400 */
  position: relative;
  z-index: 10;
  opacity: 1 !important; /* Force visibility */
}

/* Custom Scrollbar */
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.05);
  border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.2);
  border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: rgba(234, 179, 8, 0.5);
}

/* ensure top yellow buttons/links always align in one row */
.nav-link,
.nav-link-mobile {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  white-space: nowrap; /* prevent wrapping so items stay in one line */
  padding-top: 0;
  padding-bottom: 0;
}

/* keep inner text vertically centered and above the pseudo background */
.nav-link > span,
.nav-link-mobile > span {
  display: inline-flex;
  align-items: center;
  height: 100%;
}

/* ensure right-side buttons (EN / LOG IN) match height and alignment */
nav .flex.items-center.h-20 > * {
  min-height: 80px; /* match h-20 (20 * 4px = 80px) to avoid misalignment */
  display: flex;
  align-items: center;
}

/* Fix alignment: make nav items (including leftmost PRODUCTS button) vertically centered and same height */
.nav-link {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 80px; /* برابر h-20 */
  box-sizing: border-box;
  padding-top: 0;
  padding-bottom: 0;
  white-space: nowrap;
}

.nav-link > span {
  display: inline-flex;
  align-items: center;
  height: 100%;
}

/* ensure icons inside nav-link don't shift baseline */
.nav-link svg {
  display: inline-block;
  vertical-align: middle;
  margin-left: 6px;
}

/* make sure nav-link inside border wrappers fills parent's height */
.hidden.md\:flex.items-center.h-full > .nav-link,
.hidden.md\:flex.items-center.h-full .nav-link {
  height: 100%;
}

/* ensure container direct children align consistently */
.flex.items-center.h-20 > * {
  height: 100%;
  display: flex;
  align-items: center;
}

/* small tweak for mobile links to keep consistent touch area */
.nav-link-mobile {
  min-height: 44px;
}

/* hide the mobile menu button on desktop widths (force if tailwind not applied) */
@media (min-width: 768px) {
  .hide-desktop {
    display: none !important;
  }
}

/* Navbar container - حذف min-width سخت و استفاده از padding مساوی */
.navbar-container {
  width: 100%;
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 16px; /* فاصله مساوی از چپ و راست */
  /* min-width: 1280px; -- removed */
}

/* تنظیم شیفت افقی به صفر تا navbar در همه صفحات وسط قرار گیرد */
.navbar-inner {
  --navbar-shift: 0;
  transform: translateX(var(--navbar-shift));
  transition: transform 0.18s ease;
}

/* Tablet and up: حذف شیفت منفی قبلی (که باعث جابجایی می‌شد) */
@media (min-width: 768px) {
  .navbar-container {
    max-width: 1280px;
    /* min-width: removed to avoid offset */
  }

  .navbar-inner {
    --navbar-shift: 0; /* ensure centered on larger screens too */
    transform: translateX(var(--navbar-shift));
  }
}

/* موبایل: اطمینان از اینکه کانتینر از صفحه بیرون نمیزند */
@media (max-width: 767px) {
  .navbar-container {
    max-width: 100%;
    padding: 0 16px;
  }
  .navbar-inner {
    transform: none !important;
  }
}

/* Mobile-specific improvements */
@media (max-width: 767px) {
  /* Increase touch target size on mobile */
  .nav-link-mobile {
    min-height: 48px !important;
    padding: 12px 6px !important;
  }

  /* Improve menu item spacing on mobile */
  .menu-item {
    min-height: 48px !important;
    padding: 16px !important;
  }

  /* Adjust text sizes for mobile */
  .menu-item h3 {
    font-size: 16px !important;
  }

  .menu-item span {
    font-size: 12px !important;
  }

  /* Fullscreen menu adjustments */
  .navbar-container {
    max-width: 100% !important;
    min-width: auto !important;
    padding: 0 !important;
  }

  .navbar-inner {
    transform: none !important;
  }

  /* Improve mobile logo sizing */
  .navbar-inner img {
    width: 160px !important;
    height: 96px !important;
  }
}

/* Tablet and up */
@media (min-width: 768px) {
  .navbar-container {
    max-width: 1280px;
    min-width: 1280px;
  }

  .navbar-inner {
    --navbar-shift: -32px;
    transform: translateX(var(--navbar-shift));
  }
}
</style>

