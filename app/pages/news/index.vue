<template>
  <div>
    <Navbar />
    <main class="bg-gray-100">
    <section class="previews md:flex">
      <!-- Hero Section with Background Images -->
      <div class="relative h-[30em] md:fixed md:top-0 md:left-0 md:h-screen md:w-[45%] bg-black">
        <figure
          v-for="(image, index) in backgroundImages"
          :key="index"
          class="absolute inset-0 bg-cover bg-center bg-no-repeat transition-opacity duration-1000 grayscale-img"
          :class="{ 'opacity-0': currentImageIndex !== index, 'opacity-100': currentImageIndex === index }"
          :style="{ backgroundImage: `url(${image})` }"
        />
        <div class="absolute inset-0 bg-black/20 transition-all duration-400" />
        <div class="absolute inset-0 z-10 flex flex-col items-center justify-center text-white p-8">
          <span class="uppercase tracking-wider text-sm mb-2"></span>
          <h1 class="text-4xl md:text-5xl xl:text-6xl font-bold"></h1>
        </div>
      </div>

      <!-- Content Section -->
      <div class="min-h-screen md:ml-[45%]">
        <!-- Tabs Header -->
        <header>
          <ul class="flex justify-around p-4 md:px-32 xl:px-40 uppercase tracking-wider text-sm">
            <li
              v-for="(tab, index) in tabs"
              :key="index"
              class="cursor-pointer pb-1 relative"
              :class="{ 'border-b-3 border-[#FFCD05]': activeTab === index }"
              @click="changeTab(index)"
            >
              <span
                class="relative after:content-[''] after:absolute after:bottom-0 after:left-0 after:w-0 after:border-b-3 after:border-[#FFCD05] after:transition-all after:duration-400"
                :class="{ 'hover:after:w-full': activeTab !== index }"
              >
                {{ tab }}
              </span>
            </li>
          </ul>
        </header>

        <!-- Posts Tab -->
        <div v-show="activeTab === 0" class="tab transition-opacity duration-1000" :class="{ 'opacity-100': activeTab === 0, 'opacity-0': activeTab !== 0 }">
          <ul>
            <li
              v-for="(post, index) in posts"
              :key="index"
              class="preview transition-colors duration-200"
              :class="{ 'bg-white': index % 2 === 1 }"
              @mouseenter="changeImage(index)"
            >
              <NuxtLink :to="post.link || '#'" class="block p-8 md:px-32 md:py-16 xl:px-40">
                <div class="flex items-center gap-2 mb-1">
                  <span class="text-black text-sm">{{ post.date }}</span>
                  <span v-if="post.tag" class="px-2 py-1 text-xs font-semibold uppercase tracking-wider bg-[#FFCD05] text-white rounded">{{ post.tag }}</span>
                </div>
                <h2 class="text-[#FFCD05] text-2xl md:text-3xl font-bold mb-2 mt-1">{{ post.title }}</h2>
                <p class="text-black mb-3 leading-relaxed">{{ post.excerpt }}</p>
                <span
                  class="btn-slide-down inline-block px-2 py-2 uppercase tracking-wider text-sm text-[#FFCD05] border-3 border-[#FFCD05] rounded-lg relative overflow-hidden"
                >
                  <span class="relative z-10">Read More</span>
                </span>
              </NuxtLink>
            </li>
          </ul>

          <!-- Footer -->
          <footer class="flex flex-col sm:flex-row items-center justify-between px-8 py-8">
            <a href="#" class="archive-link uppercase tracking-wider text-sm relative pb-1">Archive</a>
            
          </footer>
        </div>

        <!-- Categories Tab -->
        <div v-show="activeTab === 1" class="tab transition-opacity duration-1000" :class="{ 'opacity-100': activeTab === 1, 'opacity-0': activeTab !== 1 }">
          <ul class="grid grid-cols-1 md:grid-cols-2 gap-6 md:gap-8 p-6 md:p-8 xl:p-12">
            <li v-for="(category, index) in categories" :key="index" class="card">
              <a href="#" class="block bg-white shadow-[0_1px_2px_rgba(10,10,10,0.1),0_1px_2px_rgba(10,10,10,0.2)] transition-all duration-200 group relative h-full flex flex-col">
                <div class="relative h-56 md:h-64 lg:h-72 bg-black overflow-hidden flex-shrink-0">
                  <figure
                    class="absolute inset-0 bg-cover bg-center bg-no-repeat grayscale-img"
                    :style="{ backgroundImage: `url(${category.image})` }"
                  />
                  <div class="absolute inset-0 bg-black/20 transition-all duration-400" />
                </div>
                <div class="p-6 md:p-8 flex-1 flex flex-col">
                  <h2 class="text-[#FFCD05] text-2xl md:text-3xl font-bold mb-2">{{ category.name }}</h2>
                  <p class="text-black mb-4 text-lg md:text-xl">{{ category.count }} Posts</p>
                  <span
                    class="btn-slide-down inline-block px-5 py-3 uppercase tracking-wider text-sm md:text-base text-[#FFCD05] border-2 border-[#FFCD05] rounded-lg relative overflow-hidden mt-auto"
                  >
                    <span class="relative z-10">View All</span>
                  </span>
                </div>
              </a>
            </li>
          </ul>

          <!-- Footer -->
          <footer class="flex flex-col sm:flex-row items-center justify-between px-8 py-8">
            <a href="#" class="archive-link uppercase tracking-wider text-sm relative pb-1">Archive</a>
            
          </footer>
        </div>
      </div>
    </section>
  </main>
  <!-- Chat widget (always present on page) -->
   <Chatbox />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Navbar from '~/components/navbar.vue'
import Footer from '~/components/footer.vue'
import Chatbox from '~/components/chatbox.vue'

const tabs = ['Posts', 'Categories']
const activeTab = ref(0)
const currentImageIndex = ref(0)

const backgroundImages = [
'/Tabriz-Plast-2025-1.webp',  
'/news1.jfif',
  '/Visit-of-the-CEO-of-Bank-of-Industry.png',
  '/Hdchem.webp',
  '/PolypropyleneandPolypropyleneCompositesMarketzz.jpg',
  '/dynamic-data-visualization-3d.jpg',
  '/high-angle-plastic-bottles-arrangement2.jpg'

]

const posts = [
  {
    date: 'Dec 16-19, 2025',
    title: 'POLYCHEM at Tabriz Plast 2025',
    excerpt: 'POLYCHEM will participate in the 22nd International Specialized Exhibition of Rubber, Plastics and Machinery (Tabriz Plast 2025).',
    tag: 'event',
    link: '/news/tabriz-plast-2025'
  },
   
{
    date: 'Dec 3-6, 2025',
    title: 'Plast Eurasia İstanbul 2025',
    excerpt: 'POLYCHEM is excited to announce its participation in Plast Eurasia Istanbul 2025, the leading plastics industry trade fair in the region, taking place from 3–6 December 2025 at TÜYAP Fair & Congress Center.',
    tag: 'event',
    link: '/news/plast-eurasia-2025'
  },
   {
    date: 'Nov 24, 2025',
    title: 'Visit of the Head of the Central Branch of Bank of Industry and Mine to Company POLYCHEM',
    excerpt: 'The visit was attended by several vice presidents and senior managers of the bank, as well as the managers of POLYCHEM Company (the operator of the Polychem brand). Dr. Rezavi closely inspected the company’s advanced production lines for engineering polymers, specialized composites, masterbatches, polymeric additives, and the new generation of products, while reviewing the latest research and production achievements.',
    tag: 'News',
    link: '/news/Visit-of-the-CEO-of-Bank-of-Industry',

  },
   {
    date: 'Nov 14, 2025',
    title: 'New Product Launch: HDCHEM 4760 – The Ultimate Blow Molding Compound',
    excerpt: 'We are proud to unveil the most powerful blow molding grade in the entire HDCHEM range: HDCHEM 4760',
    tag: 'News',
    link: '/news/New-Product-Launch-HDCHEM-4760',

  },
  {
    date: 'Nov 1, 2025',
    title: 'Global Polypropylene and Polypropylene Composites Market Size, Share, and Trends Analysis Report – Industry Overview and Forecast to 2032',
    excerpt: 'Global Polypropylene and Polypropylene Composites Market, By Product Type (Polypropylene (PP), Polypropylene Composites), Fiber Type (Glass Fiber, Carbon Fiber, Others), Application (Packaging, Building and Construction, Automotive, Electrical and Electronics, Aerospace and Defence, Others) - Industry Trends and Forecast to 2032',
    tag: 'Report',
    link: '/news/Global-Polypropylene-and-Composites-Market-Report-2032'
  },
  {
    date: 'Dec 6, 2024',
    title: 'Global Cast Low-Density Polyethylene Disposable Gloves Market Size, Share, and Trends Analysis Report – Industry Overview and Forecast to 2032',
    excerpt: 'Global Cast Low-Density Polyethylene Disposable Gloves Market Segmentation, By Size (Small, Medium, Large, and X-Large), End Users (Food and Beverages, Medical, Industrial, and Others), Distribution Channel (Direct Distribution and Via Distributors) – Industry Trends and Forecast to 2032',
    tag: 'Report',
    link: '/news/Cast_LowDensity_PE_Gloves_Market_Report_2032',

  },
    {
    date: 'Oct 27, 2025',
    title: 'Rising Polypropylene Prices Increase Production Costs by Up to 40%',
    excerpt: 'The runaway increase in the prices of polyethylene and polypropylene has sounded a new alarm for the food industry and consumers. This surge has put production units—from packaging manufacturers to food factories—under a cost crisis, and its effects are now directly reflected in the prices of essential goods.',
    tag: 'News',
    link: '/news/Rising-Polypropylene-Prices'
  },
]

const categories = [
  {
    name: 'Events',
    count: 1,
    image: backgroundImages[0]
  },
  {
    name: 'Reports',
    count: 2,
    image: backgroundImages[1]
  },
  {
    name: 'News',
    count: 2,
    image: backgroundImages[3]
  },
  {
    name: 'Market Analysis',
    count: 1,
    image: backgroundImages[2]
  }
]

const changeTab = (index) => {
  activeTab.value = index
}

const changeImage = (index) => {
  currentImageIndex.value = index
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500;600;700&display=swap');

body {
  font-family: 'Montserrat', sans-serif;
}

h1, h2 {
  font-family: 'Montserrat', sans-serif;
}

.border-b-3 {
  border-bottom-width: 3px;
}

.border-3 {
  border-width: 3px;
}

.grayscale-img {
  filter: grayscale(100%);
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

/* Archive link hover underline effect - yellow, from left */
.archive-link::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 0;
  height: 3px;
  background-color: #FFCD05;
  transition: width 0.4s ease;
}

.archive-link:hover::after {
  width: 100%;
}
</style>

