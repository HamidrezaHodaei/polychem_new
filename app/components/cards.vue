<template>
  <section class="chemistry-cards-section" ref="sectionRef">
    <div class="container">
      <div class="cards-container">
        <!-- render rows of 2 cards each -->
        <div class="row" v-for="(rowCards, rIdx) in rows" :key="rIdx">
          <div class="col" v-for="(card, cIdx) in rowCards" :key="rIdx * 2 + cIdx">
            <div class="single-card-wrapper">
              <a :href="card.link">
                <div class="card-hover-container">
                  <!-- Image (left on desktop) -->
                  <div class="card-image">
                    <img :src="card.image" :alt="card.title" />
                  </div>

                  <!-- Content (right on desktop) -->
                  <div class="card-content">
                    <span class="card-badge">{{ card.badge }}</span>
                    <h3 class="card-title">{{ card.title }}</h3>
                    <p class="card-date">{{ card.date }}</p>
                    <button class="card-button"><span>Read more</span></button>
                  </div>

                  <div class="card-gradient"></div>
                </div>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'

const backgroundImages = [
  '/news1.jfif',
  '/Visit-of-the-CEO-of-Bank-of-Industry.png',
  '/Hdchem.webp',
  '/PolypropyleneandPolypropyleneCompositesMarketzz.jpg',
  '/Tabriz-Plast-2025-1.Webp',
  '/high-angle-plastic-bottles-arrangement2.jpg'
]

const cards = ref([
{
    badge: 'Event',
    title: 'POLYCHEM will participate in the 22nd International Specialized Exhibition of Tabriz Plast 2025',
    date: 'Dec 16-19, 2025',
    link: '/news/tabriz-plast-2025',
    image: backgroundImages[4]
  }, 
{
    badge: 'Event',
    title: 'POLYCHEM is excited to announce its participation in Plast Eurasia Istanbul 2025.',
    date: 'Dec 3-6, 2025',
    link: '/news/plast-eurasia-2025',
    image: backgroundImages[0]
  },
  {
    badge: 'News',
    title: 'Visit of the Head of the Central Branch of Bank of Industry and Mine',
    date: 'Nov 24, 2025',
    link: '/news/Visit-of-the-CEO-of-Bank-of-Industry',
    image: backgroundImages[1]
  },
  {
    badge: 'News',
    title: 'New Product Launch: HDCHEM 4760 – The Ultimate Blow Molding Compound',
    date: 'Nov 14, 2025',
    link: '/news/New-Product-Launch-HDCHEM-4760',
    image: backgroundImages[2]
  },
  {
    badge: 'Report',
    title: 'Global Polypropylene and Polypropylene Composites Market Trends Analysis Report',
    date: 'Nov 1, 2025',
    link: '/news/Global-Polypropylene-and-Composites-Market-Report-2032',
    image: backgroundImages[3]
  },

  {
    badge: 'News',
    title: 'Rising Polypropylene Prices Increase Production Costs by Up to 40%',
    date: 'Oct 27, 2025',
    link: '/news/Rising-Polypropylene-Prices',
    image: backgroundImages[5]
  }
])

// chunk into rows of 2
const rows = computed(() => {
  const out = []
  for (let i = 0; i < cards.value.length; i += 2) {
    out.push(cards.value.slice(i, i + 2))
  }
  return out
})

const sectionRef = ref(null)
const navbarHeight = ref(60) // تنظیم بر اساس ارتفاع navbar واقعی شما

const measureAndApplyHeight = () => {
  // do not force a fixed height on small screens — let content size naturally
  const mobileBreakpoint = 991
  if (!sectionRef.value) return
  if (window.innerWidth <= mobileBreakpoint) {
    // remove previously set fixed card height so cards expand to fit content on mobile
    sectionRef.value.style.removeProperty('--card-height')
    return
  }

  // target card link for "Global Cast Low-Density"
  const targetHref = '/news/Cast_LowDensity_PE_Gloves_Market_Report_2032'
  const anchor = sectionRef.value.querySelector(`a[href="${targetHref}"]`)
  let height = null
  if (anchor) {
    const cardContainer = anchor.querySelector('.card-hover-container') || anchor.closest('.card-hover-container')
    if (cardContainer) {
      height = cardContainer.getBoundingClientRect().height
    }
  }
  // fallback default height if target not rendered yet
  if (!height) {
    // try to pick first card height or use default
    const firstCard = sectionRef.value.querySelector('.card-hover-container')
    height = firstCard ? firstCard.getBoundingClientRect().height : 320
  }
  sectionRef.value.style.setProperty('--card-height', `${Math.round(height)}px`)
}

const handleScroll = () => {
  // فقط در موبایل فعال شود
  const mobileBreakpoint = 576
  if (window.innerWidth > mobileBreakpoint) {
    // در دسکتاپ، scrolling class را حذف کن تا hover دوباره کار کند
    if (!sectionRef.value) return
    const images = sectionRef.value.querySelectorAll('.card-image img.scrolling')
    images.forEach(img => img.classList.remove('scrolling'))
    return
  }

  if (!sectionRef.value) return

  const images = sectionRef.value.querySelectorAll('.card-image img')
  images.forEach(img => {
    const rect = img.getBoundingClientRect()
    // اگر تصویر از navbar پایین‌تر است، scrolling class را اضافه کن
    if (rect.top > navbarHeight.value) {
      img.classList.add('scrolling')
    } else {
      img.classList.remove('scrolling')
    }
  })
}

let resizeObserver = null
onMounted(() => {
  // measure initially after layout
  requestAnimationFrame(measureAndApplyHeight)
  // recompute on window resize
  window.addEventListener('resize', measureAndApplyHeight)
  // اضافه کردن scroll listener برای موبایل
  window.addEventListener('scroll', handleScroll)
  // also use ResizeObserver to catch image/content load changes
  if (window.ResizeObserver && sectionRef.value) {
    resizeObserver = new ResizeObserver(() => {
      measureAndApplyHeight()
    })
    // observe each card container
    sectionRef.value.querySelectorAll('.card-hover-container').forEach(el => resizeObserver.observe(el))
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', measureAndApplyHeight)
  window.removeEventListener('scroll', handleScroll)
  if (resizeObserver) {
    resizeObserver.disconnect()
    resizeObserver = null
  }
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;600;700;800&display=swap');

*,
*::after,
*::before {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.chemistry-cards-section {
  min-height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f1f2f2;
  font-family: 'Montserrat', sans-serif;
  padding: 20px;
}

/* layout containers */
.container {
  width: 100%;
  max-width: 1400px;
  margin: 0 auto;
}

.cards-container {
  position: relative;
  width: 100%;
}

.row {
  display: flex;
  gap: 20px;
  flex-wrap: nowrap;
  justify-content: center;
}

.row-top {
  margin-bottom: 0;
}

.row-bottom {
  margin-top: 0;
}

.col {
  flex: 1;
  min-width: 0;
}

.single-card-wrapper {
  height: 100%;
}

.single-card-wrapper > a {
  display: block;
  height: 100%;
  width: 100%;
  text-decoration: none;
  color: inherit;
}

/* make card look like slider: left image, right content */
.card-hover-container {
  border: 1px solid #e6e6e6;
  width: 100%;
  border-radius: 12px;
  display: flex;
  align-items: stretch;
  overflow: hidden;
  position: relative;
  cursor: pointer;
  /* use measured height from the Global Cast Low-Density card */
  height: var(--card-height, 320px);
  background: #ffffff;
  transition: transform 0.3s ease;
}
.card-hover-container:hover { transform: translateY(-4px); }

/* image column */
.card-image {
  position: relative;
  flex: 0 0 60%;
  max-width: 60%;
  background-color: #ffffff;
  overflow: hidden;
}
.card-image img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
  filter: grayscale(100%);
  transition: filter 0.45s ease, transform 0.45s ease;
}

/* remove grayscale on hover (desktop) */
.card-hover-container:hover .card-image img {
  filter: none;
}

/* mobile: scrolling class removes grayscale */
.card-image img.scrolling {
  filter: none;
}

/* content column */
.card-content {
  position: relative;
  z-index: 2;
  padding: 24px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  background: #ffffff;
}
.card-badge {
  display: inline-block;
  background: #555555;
  color: #ffffff;
  padding: 6px 14px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
  margin-bottom: 20px;
  text-transform: capitalize;
}
.card-title {
  color: #848484;
  font-size: 20px;
  font-weight: 600;
  line-height: 1.25;
  margin: 12px 0;
}
.card-date {
  color: #FFCD05;
  font-size: 14px;
  margin-bottom: 10px;
}
.card-button {
  background: transparent;
  color: #FFCD05;
  border: 2px solid #FFCD05;
  padding: 10px 20px;
  border-radius: 6px;
  cursor: pointer;
  align-self: flex-start;
  position: relative;
  overflow: hidden;
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
  z-index: 2;
}

/* subtle gradient overlay */
.card-gradient {
  position: absolute;
  inset: 0;
  z-index: 1;
  pointer-events: none;
  opacity: 0.05;
  background: radial-gradient(circle at center, rgba(158,158,158,0.2) 15%, rgba(66,66,66,0.1) 50%, transparent 80%);
}

/* vertical yellow accent like Slider */
.card-content::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  width: 6px;
  height: 100%;
  background: #FFCD05;
  z-index: 3;
}

/* Responsive: stack image above content on small screens */
@media (max-width: 991px) {
  .row { flex-wrap: wrap; }
  .col { min-width: 50%; flex: 0 0 50%; }
  .card-hover-container {
    /* allow cards to grow with their content on small screens so text is fully visible */
    height: auto;
    min-height: 260px;
    flex-direction: column;
    overflow: visible;
  }
  .card-image { flex: 0 0 auto; max-width: 100%; height: 140px; }
  .card-content { padding: 16px; }
  .card-content::before { height: 6px; width: 100%; top: 0; left: 0; }
}

/* Mobile: full width rows and scroll-to-color behavior */
@media (max-width: 576px) {
  .col { min-width: 100%; flex: 0 0 100%; }
  .card-image { height: 200px; }
  .card-title { font-size: 18px; }
  .card-image img { filter: grayscale(100%); } /* keep images grayscale on mobile by default */
  .card-image img.scrolling { filter: none; } /* remove grayscale when scrolling */
  /* ensure mobile cards expand to fit content */
  .card-hover-container { height: auto; overflow: visible; }
}

/* ensure clickable area looks consistent */
.single-card-wrapper a { display:block; color:inherit; text-decoration:none; }
</style>