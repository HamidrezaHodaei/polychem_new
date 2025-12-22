<template>
  <section class="chemistry-cards-section">
    <div class="container">
      <div class="cards-container">
        <!-- ردیف اول - 4 کارت بالا -->
        <div class="row row-top">
          <div class="col" v-for="(card, index) in cards.slice(0, 4)" :key="index">
            <div class="single-card-wrapper">
              <a :href="card.link">
                <div 
                  class="card-hover-container"
                  @mousemove="handleMouseMove($event, index)"
                  @touchmove="handleTouchMove($event, index)"
                  @mouseleave="handleMouseLeave(index)"
                  @touchend="handleTouchEnd(index)"
                >
                  <div class="card-content">
                    <span class="card-badge">{{ card.badge }}</span>
                    <h3 class="card-title">{{ card.title }}</h3>
                    <p class="card-date">{{ card.date }}</p>
                    <button class="card-button"><span>Read more</span></button>
                  </div>
                  <div class="card-gradient"></div>
                  <div 
                    class="card-bg-characters"
                    :ref="el => { if (el) charactersRefs[index] = el }"
                  >
                    {{ randomTexts[index] }}
                  </div>
                </div>
              </a>
            </div>
          </div>
        </div>
        
        <!-- ردیف دوم - 4 کارت پایین -->
        <div class="row row-bottom">
          <div class="col" v-for="(card, index) in cards.slice(4, 8)" :key="index + 4">
            <div class="single-card-wrapper">
              <a :href="card.link">
                <div 
                  class="card-hover-container"
                  @mousemove="handleMouseMove($event, index + 4)"
                  @touchmove="handleTouchMove($event, index + 4)"
                  @mouseleave="handleMouseLeave(index + 4)"
                  @touchend="handleTouchEnd(index + 4)"
                >
                  <div class="card-content">
                    <span class="card-badge">{{ card.badge }}</span>
                    <h3 class="card-title">{{ card.title }}</h3>
                    <p class="card-date">{{ card.date }}</p>
                    <button class="card-button"><span>Read more</span></button>
                  </div>
                  <div class="card-gradient"></div>
                  <div 
                    class="card-bg-characters"
                    :ref="el => { if (el) charactersRefs[index + 4] = el }"
                  >
                    {{ randomTexts[index + 4] }}
                  </div>
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
import { ref, onMounted } from 'vue'

// تعریف کارت‌ها
const cards = ref([
  {
    badge: 'News',
    title: '  Discover the factors influencing recycled PET ',
    date: 'November 25, 2025',
    link: '#'
  },
  {
    badge: 'Chemical Analysis',
    title: 'Organic Chemistry Reactions Study',
    date: 'November 10, 2025',
    link: '#'
  },
  {
    badge: 'Laboratory Report',
    title: 'Synthesis of Complex Compounds',
    date: 'November 8, 2025',
    link: '#'
  },
  {
    badge: 'Research Paper',
    title: 'Catalytic Reactions in Chemistry',
    date: 'November 5, 2025',
    link: '#'
  },
  {
    badge: 'Scientific Study',
    title: 'Thermodynamic Equilibrium Analysis',
    date: 'November 3, 2025',
    link: '#'
  },
  {
    badge: 'Technical Report',
    title: 'Electrochemistry Applications',
    date: 'November 1, 2025',
    link: '#'
  },
  {
    badge: 'Case Study',
    title: 'Polymer Chemistry Advances',
    date: 'October 28, 2025',
    link: '#'
  },
  {
    badge: 'White Paper',
    title: 'Molecular Structure Investigation',
    date: 'October 25, 2025',
    link: '#'
  }
])



const charactersRefs = ref([])

const PARTICLE_LENGTH_INITIAL = 95000   // مقدار اولیه بسیار بزرگ برای پر کردن پس‌زمینه
const PARTICLE_LENGTH_ON_MOVE = 95000   // مقدار تولید مجدد هنگام حرکت موس/لمس

// تولید رشته‌ای متشکل از نقطه‌ها به جای معادلات شیمیایی
const randomChemicalString = (length = PARTICLE_LENGTH_ON_MOVE) => {
  const pair = '· ' // نقطه و فاصله برای بهتر خط‌شدن متن
  const repeats = Math.ceil(length / pair.length)
  return pair.repeat(repeats).substring(0, length)
}

// مقدار اولیه‌ی متون پس‌زمینه را با رشته‌های نقطه پر کن (بالاتر)
const randomTexts = ref(Array(cards.value.length).fill(randomChemicalString(PARTICLE_LENGTH_INITIAL)))

// مدیریت حرکت موس
const handleMouseMove = (e, index) => {
  const card = e.currentTarget
  const rect = card.getBoundingClientRect()
  const x = e.clientX - rect.left
  const y = e.clientY - rect.top
  
  if (charactersRefs.value[index]) {
    // ست کردن نقطه‌ی ماسک برای موضع نور
    charactersRefs.value[index].style.setProperty('--x', `${x}px`)
    charactersRefs.value[index].style.setProperty('--y', `${y}px`)

    // محاسبه‌ی ترجمه‌ی کوچک به سمت مکان موس (برای احساس "جلو آمدن")
    const cx = rect.width / 2
    const cy = rect.height / 2
    const maxOffset = 18 // حداکثر جابجایی پیکسل
    const tx = ((x - cx) / cx) * maxOffset
    const ty = ((y - cy) / cy) * maxOffset

    charactersRefs.value[index].style.setProperty('--tx', `${tx}px`)
    charactersRefs.value[index].style.setProperty('--ty', `${ty}px`)
    charactersRefs.value[index].style.setProperty('--scale', '1.06')

    // به‌روزرسانی متن با طول بزرگ‌تر برای افزایش "تراکم ذرات"
    randomTexts.value[index] = randomChemicalString(PARTICLE_LENGTH_ON_MOVE)
  }
}

// مدیریت لمس
const handleTouchMove = (e, index) => {
  const touch = e.touches[0]
  const card = e.currentTarget
  const rect = card.getBoundingClientRect()
  const x = touch.clientX - rect.left
  const y = touch.clientY - rect.top
  
  if (charactersRefs.value[index]) {
    charactersRefs.value[index].style.setProperty('--x', `${x}px`)
    charactersRefs.value[index].style.setProperty('--y', `${y}px`)

    const cx = rect.width / 2
    const cy = rect.height / 2
    const maxOffset = 18
    const tx = ((x - cx) / cx) * maxOffset
    const ty = ((y - cy) / cy) * maxOffset

    charactersRefs.value[index].style.setProperty('--tx', `${tx}px`)
    charactersRefs.value[index].style.setProperty('--ty', `${ty}px`)
    charactersRefs.value[index].style.setProperty('--scale', '1.06')

    // به‌روزرسانی متن با طول بزرگ‌تر برای افزایش "تراکم ذرات"
    randomTexts.value[index] = randomChemicalString(PARTICLE_LENGTH_ON_MOVE)
  }
}

// وقتی کارت ترک می‌شود، مقادیر به حالت اولیه بازگردانده شوند
const handleMouseLeave = (index) => {
  if (charactersRefs.value[index]) {
    charactersRefs.value[index].style.setProperty('--tx', `0px`)
    charactersRefs.value[index].style.setProperty('--ty', `0px`)
    charactersRefs.value[index].style.setProperty('--scale', '1')
    // می‌توانیم مکان نور را به مرکز برگردانیم (اختیاری)
    // charactersRefs.value[index].style.setProperty('--x', `50%`)
    // charactersRefs.value[index].style.setProperty('--y', `50%`)
  }
}

const handleTouchEnd = (index) => {
  handleMouseLeave(index)
}
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
  font-family: 'Montserrat', sans-serif; /* changed to Montserrat */
  padding: 20px;
}

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
  gap: 0;
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
}

.card-hover-container {
  border: 1px solid #e0e0e0;
  width: 100%;
  border-radius: 0;
  height: 100%;
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
  overflow: hidden;
  position: relative;
  cursor: pointer;
  min-height: 400px;
  background: #ffffff;
  transition: transform 0.3s ease;
}

.card-hover-container:hover {
  transform: translateY(-2px);
}

.card-content {
  position: relative;
  z-index: 4;
  padding: 30px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 100%;
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
  color: #808285;
  font-size: 24px;
  font-weight: 600;
  line-height: 1.3;
  margin-bottom: 20px;
  max-width: 90%;
}

.card-date {
  color: #666666;
  font-size: 14px;
  margin-bottom: 30px;
}

.card-button {
  background: transparent;
  color: #FFCD05;
  border: 2px solid #FFCD05;
  padding: 12px 28px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  margin-top: auto;
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
  z-index: 1;
}

.card-gradient {
  height: 100%;
  width: 100%;
  position: absolute;
  background: radial-gradient(circle at center, #9e9e9e 20%, #757575 40%, #616161 70%, #424242 100%);
  mix-blend-mode: overlay;
  pointer-events: none;
  z-index: 3;
  opacity: 0.5;
}

.card-bg-characters {
  --x: 0px;
  --y: 0px;
  --tx: 0px;
  --ty: 0px;
  --scale: 1;
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  word-wrap: break-word;
  font-size: 15px;
  line-height: 0.5;
  overflow: hidden;
  font-family: 'Montserrat', sans-serif; /* changed to Montserrat */
  color: #a8a8a8;
  opacity: 0;
  transition: opacity 0.45s ease, transform 0.45s cubic-bezier(.2,.9,.3,1);
  will-change: transform, opacity;
  -webkit-mask-image: radial-gradient(
    300px circle at var(--x) var(--y),
    #000 20%,
    rgba(0, 0, 0, 0.25),
    transparent
  );
  mask-image: radial-gradient(
    300px circle at var(--x) var(--y),
    #000 20%,
    rgba(0, 0, 0, 0.25),
    transparent
  );
  transform: translate(var(--tx, 0px), var(--ty, 0px)) scale(var(--scale, 1.025));
  padding: 20px;
}

.card-hover-container:hover .card-bg-characters {
  opacity: 1;
}

@media (max-width: 991px) {
  .row {
    flex-wrap: wrap;
  }
  
  .col {
    min-width: 50%;
    flex: 0 0 50%;
  }
  
  .card-hover-container {
    min-height: 320px;
  }

  .card-title {
    font-size: 20px;
  }

  .card-content {
    padding: 20px;
  }
}

@media (max-width: 576px) {
  .col {
    min-width: 100%;
    flex: 0 0 100%;
  }

  .card-hover-container {
    min-height: 250px;
  }

  .card-title {
    font-size: 18px;
  }
}
</style>