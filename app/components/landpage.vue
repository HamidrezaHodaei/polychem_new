<template>
  <div ref="blackholeRef" class="blackhole">
    <div 
      ref="centerHoverRef"
      class="centerHover" 
      :class="{ open: isOpen }"
      @click="handleClick"
      @mouseover="handleMouseOver"
      @mouseout="handleMouseOut"
    >
      <span>ENTER</span>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const blackholeRef = ref(null)
const centerHoverRef = ref(null)
const isOpen = ref(false)

let canvas = null
let context = null
let animationId = null
let stars = []
let collapse = false
let expanse = false
let startTime = 0
let currentTime = 0

const maxorbit = 255
let centery = 0
let centerx = 0
let cw = 0
let ch = 0

function setDPI(canvas, dpi) {
  if (!canvas.style.width)
    canvas.style.width = canvas.width + 'px'
  if (!canvas.style.height)
    canvas.style.height = canvas.height + 'px'

  const scaleFactor = dpi / 96
  canvas.width = Math.ceil(canvas.width * scaleFactor)
  canvas.height = Math.ceil(canvas.height * scaleFactor)
  const ctx = canvas.getContext('2d')
  ctx.scale(scaleFactor, scaleFactor)
}

function rotate(cx, cy, x, y, angle) {
  const radians = angle
  const cos = Math.cos(radians)
  const sin = Math.sin(radians)
  const nx = (cos * (x - cx)) + (sin * (y - cy)) + cx
  const ny = (cos * (y - cy)) - (sin * (x - cx)) + cy
  return [nx, ny]
}

class Star {
  constructor() {
    const rands = []
    rands.push(Math.random() * (maxorbit / 2) + 1)
    rands.push(Math.random() * (maxorbit / 2) + maxorbit)

    this.orbital = (rands.reduce((p, c) => p + c, 0) / rands.length)
    
    this.x = centerx
    this.y = centery + this.orbital
    this.yOrigin = centery + this.orbital

    this.speed = (Math.floor(Math.random() * 2.5) + 1.5) * Math.PI / 180
    this.rotation = 0
    this.startRotation = (Math.floor(Math.random() * 360) + 1) * Math.PI / 180

    this.id = stars.length

    this.collapseBonus = this.orbital - (maxorbit * 0.7)
    if (this.collapseBonus < 0) {
      this.collapseBonus = 0
    }

    this.color = 'rgba(255,255,255,' + (1 - ((this.orbital) / 255)) + ')'

    this.hoverPos = centery + (maxorbit / 2) + this.collapseBonus
    this.expansePos = centery + (this.id % 100) * -10 + (Math.floor(Math.random() * 20) + 1)

    this.prevR = this.startRotation
    this.prevX = this.x
    this.prevY = this.y
    
    this.originalY = this.yOrigin

    stars.push(this)
  }

  draw() {
    // حالت عادی - قبل از کلیک
    if (!expanse) {
      this.rotation = this.startRotation + (currentTime * this.speed)
      if (!collapse) {
        if (this.y > this.yOrigin) {
          this.y -= 2.5
        }
        if (this.y < this.yOrigin - 4) {
          this.y += (this.yOrigin - this.y) / 10
        }
      } else {
        this.trail = 1
        if (this.y > this.hoverPos) {
          this.y -= (this.hoverPos - this.y) / -5
        }
        if (this.y < this.hoverPos - 4) {
          this.y += 2.5
        }
      }
    } 
    // حالت انبساط - بعد از کلیک (ذرات به سمت بیرون میروند و می‌مانند)
    else if (expanse) {
      this.rotation = this.startRotation + (currentTime * (this.speed / 2))
      if (this.y > this.expansePos) {
        this.y -= Math.floor(this.expansePos - this.y) / -80
      }
    }

    context.save()
    context.fillStyle = this.color
    context.strokeStyle = this.color
    context.beginPath()
    const oldPos = rotate(centerx, centery, this.prevX, this.prevY, -this.prevR)
    context.moveTo(oldPos[0], oldPos[1])
    context.translate(centerx, centery)
    context.rotate(this.rotation)
    context.translate(-centerx, -centery)
    context.lineTo(this.x, this.y)
    context.stroke()
    context.restore()

    this.prevR = this.rotation
    this.prevX = this.x
    this.prevY = this.y
  }
}

function loop() {
  const now = new Date().getTime()
  currentTime = (now - startTime) / 50

  // trail با شفافیت بهتر
  context.fillStyle = 'rgba(132,132,132,0.1)'
  context.fillRect(0, 0, cw, ch)

  for (let i = 0; i < stars.length; i++) {
    if (stars[i] !== undefined) {
      stars[i].draw()
    }
  }

  animationId = requestAnimationFrame(loop)
}

function init() {
  const container = blackholeRef.value
  const h = container.offsetHeight
  const w = container.offsetWidth
  cw = w
  ch = h
  centery = ch / 2
  centerx = cw / 2

  startTime = new Date().getTime()
  stars = []

  canvas = document.createElement('canvas')
  canvas.width = cw
  canvas.height = ch
  container.appendChild(canvas)
  context = canvas.getContext("2d")
  
  setDPI(canvas, 192)

  context.fillStyle = 'rgba(132,132,132,1)'
  context.fillRect(0, 0, cw, ch)
  
  for (let i = 0; i < 2500; i++) {
    new Star()
  }
  
  loop()
}

// آدرس همان فایلی که در Hero.vue استفاده شده
const heroSrc = '/shutterstock_vidhd_3898471189%20(2).gif'
let heroPreloaded = false

function preloadHero() {
  try {
    const img = new Image()
    img.decoding = 'async'
    img.loading = 'eager'
    img.src = heroSrc
    img.onload = () => {
      heroPreloaded = true
    }
    img.onerror = () => {
      // اگر preload ناموفق بود، بی‌صدا رد شو
    }
  } catch (e) {
    // ignore in older browsers
  }
}

// تابعی که تا لود شدن تصویر یا timeout صبر می‌کند
function waitForHeroLoad(timeout = 5000) {
  return new Promise((resolve) => {
    if (heroPreloaded) return resolve(true)
    const start = Date.now()
    const check = setInterval(() => {
      if (heroPreloaded) {
        clearInterval(check)
        return resolve(true)
      }
      if (Date.now() - start > timeout) {
        clearInterval(check)
        return resolve(false)
      }
    }, 100)
  })
}

// handleClick را async کن و قبل از dispatch صبر کن تا تصویر آماده شود
const handleClick = async () => {
  collapse = false
  expanse = true

  // سریعاً دکمه ENTER را محو کن تا کاربر نبیند که داریم منتظر می‌مانیم
  isOpen.value = true

  // منتظر شو تا preload کامل شود یا حداکثر 5s
  const loaded = await waitForHeroLoad(5000)

  // اگر لود شد سریع‌تر event را بفرست، در غیر این صورت پس از 2s fallback
  if (loaded) {
    // کوتاه‌ترین تاخیر برای انیمیشن خروج
    setTimeout(() => {
      if (typeof window !== 'undefined') {
        window.dispatchEvent(new Event('landing-complete'))
      }
    }, 120)
  } else {
    // fallback: بعد از 2s ادامه بده
    setTimeout(() => {
      if (typeof window !== 'undefined') {
        window.dispatchEvent(new Event('landing-complete'))
      }
    }, 2000)
  }
}

const handleMouseOver = () => {
  if (expanse === false) {
    collapse = true
  }
}

const handleMouseOut = () => {
  if (expanse === false) {
    collapse = false
  }
}

onMounted(() => {
  // ابتدا preload را شروع کنیم، سپس بقیه‌ی اینتیت‌ها
  preloadHero()
  init()
})

onUnmounted(() => {
  if (animationId) {
    cancelAnimationFrame(animationId)
  }
  if (canvas && canvas.parentNode) {
    canvas.parentNode.removeChild(canvas)
  }
})
</script>

<style scoped>
.blackhole {
  height: 100vh;
  width: 100%;
  position: relative;
  display: flex;
  background-color: #848484;
  overflow: hidden;
}

.centerHover {
  width: 255px;
  height: 255px;
  background-color: transparent;
  border-radius: 50%;
  position: absolute;
  left: 50%;
  top: 50%;
  margin-top: -128px;
  margin-left: -128px;
  z-index: 2;
  cursor: pointer;
  line-height: 255px;
  text-align: center;
  transition: all 500ms;
}

.centerHover.open {
  opacity: 0;
  pointer-events: none;
}

.centerHover:hover span {
  color: #DDD;
}

.centerHover:hover span:before { 
  background-color: #DDD; 
}

.centerHover:hover span:after { 
  background-color: #DDD; 
}

.centerHover span {
  color: #fff;
  font-family: serif;
  font-size: 18px;
  position: relative;
  transition: all 500ms;
}

.centerHover span:before {
  content: '';
  display: inline-block;
  height: 1px;
  width: 16px;
  margin-right: 12px;
  margin-bottom: 4px;
  background-color: #fff;
  transition: all 500ms;
}

.centerHover span:after {
  content: '';
  display: inline-block;
  height: 1px;
  width: 16px;
  margin-left: 12px;
  margin-bottom: 4px;
  background-color: #fff;
  transition: all 500ms;
}

.blackhole :deep(canvas) {
  position: relative;
  z-index: 1;
  width: 100%;
  height: 100%;
  margin: auto;
}
</style>