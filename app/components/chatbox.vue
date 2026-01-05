<template>
  <div class="chat-widget" v-if="!isLoading">
    <!-- Floating Action Button -->
    <button
      class="chat-fab"
      v-if="minimized"
      @click="openChat"
      :aria-label="language === 'fa' ? 'باز کردن چت' : 'Open chat'"
      :title="language === 'fa' ? 'باز کردن چت' : 'Open chat'"
    >
      <!-- use chatbot svg from public folder -->
      <img src="/chatbot-poly.png" alt="POLYCHEM chatbot" class="fab-img" />
      <span class="unread-badge" v-show="unreadCount > 0">{{ unreadCount }}</span>
    </button>

    <!-- Chat Window -->
    <div class="chat-window" :class="{ active: !minimized }" role="dialog" aria-label="POLYCHEM Sales Chat" v-cloak>
      <!-- Header -->
      <div class="chat-header">
        <div class="header-left">
          <div class="header-avatar">
            <!-- use logo from public folder (spaces URL-encoded) -->
            <img src="/english%20logo%20W1.png" alt="POLYCHEM logo" stroke="currentColor"/>
          </div>
          <div class="header-info">
            <h2>POLYCHEM BOT</h2>
            <div class="header-status">
              
              
            </div>
          </div>
        </div>
        <div class="header-actions">
          <button class="header-btn" @click="clearChat" :title="language === 'fa' ? 'پاک کردن' : 'Clear chat'">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path d="M3 6h18M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2m3 0v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6h14z" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </button>
          <button class="header-btn" @click="closeChat" :title="language === 'fa' ? 'کوچک کردن' : 'Minimize'">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path d="M19 9l-7 7-7-7" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </button>
        </div>
      </div>

      <!-- Messages -->
      <div class="chat-messages" ref="chatMessages">
        <div
          v-for="msg in messages"
          :key="msg.id"
          :class="['message', msg.sender === 'user' ? 'user' : 'bot']"
          :dir="msg.lang === 'fa' ? 'rtl' : 'ltr'"
        >
          <div class="message-avatar" aria-hidden="true">
            <img v-if="msg.sender === 'bot'" src="/english%20logo%20W1.png" alt="POLYCHEM logo" >
              
            <svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2M12 11a4 4 0 1 0 0-8 4 4 0 0 0 0 8z" stroke-width="2"/>
            </svg>
          </div>
          <div class="message-bubble">
            <span v-html="msg.text"></span>
          </div>
        </div>

        <!-- Typing Indicator -->
        <div v-if="isTyping" class="message bot">
          <div class="message-avatar" aria-hidden="true">
                        <img src="/english%20logo%20W1.png" alt="POLYCHEM logo" stroke="currentColor"/>

          </div>
          <div class="typing-indicator">
            <div class="typing-dot"></div>
            <div class="typing-dot"></div>
            <div class="typing-dot"></div>
          </div>
        </div>
      </div>

      <!-- Input Area -->
      <div class="chat-input-area">
        <input
          ref="chatInput"
          v-model="input"
          type="text"
          class="chat-input"
          :placeholder="language === 'fa' ? 'پیام خود را بنویسید...' : 'Type your message...'"
          @keypress.enter.prevent="sendMessage"
          :dir="language === 'fa' ? 'rtl' : 'ltr'"
        />
        <button class="send-btn" @click="sendMessage" :aria-label="language === 'fa' ? 'ارسال' : 'Send'">
          <svg viewBox="0 0 24 24" fill="none" stroke="#FFFFFF" class="bg:">
            <path d="M22 2L11 13M22 2l-7 20-4-9-9-4 20-7z" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, nextTick, watch, onMounted, onUnmounted } from 'vue'

const isLoading = useLoadingState()

/* State */
const messages = ref([
  { 
    id: 1, 
    sender: 'bot', 
    text: `Hello! 👋 Welcome to POLYCHEM. We specialize in advanced polymer compounds and masterbatches. How can I assist you today?`, 
    lang: 'en' 
  }
])
const input = ref('')
const isTyping = ref(false)
const language = ref('en')
const chatMessages = ref(null)
const minimized = ref(true) // floating button shown when true
const unreadCount = ref(0)

/* Company Info - POLYCHEM */
const companyInfo = {
  name: 'POLYCHEM',
  phone: '+98 21 22898979',
  fax: '+98 21 22898980',
  email: 'sales@polychemmb.com',
  address: 'Unit 15-NO.45 – Manzarnejad Blvd- Shariati Ave- Mirdamad Blvd- Tehran- Iran',
  website: 'polychemmb.com'
}

/* Products Info */
const productsInfo = {
  en: [
    'Polymer Compounds (PP-Talc compounds for automotive, home appliances)',
    'Filler Masterbatch (CaCo3 based for PE, PP, PS)',
    'Color Masterbatch (organic & inorganic pigments)',
    'Additive Masterbatch (anti-block, slip, flame retardant, antimicrobial)',
    'White Masterbatch (TiO2 based)',
    'Polymer Blends (custom formulations)',
    'Engineering Polymers & Compounds'
  ],
  fa: [
    'کامپاندهای پلیمری (PP-تالک برای خودرو و لوازم خانگی)',
    'مستربچ فیلر (پایه CaCo3 برای PE، PP، PS)',
    'مستربچ رنگی (رنگدانه‌های آلی و معدنی)',
    'مستربچ افزودنی (ضد بلوک، لغزنده، شعله‌گیر، ضدمیکروب)',
    'مستربچ سفید (پایه TiO2)',
    'بلندهای پلیمری (فرمولاسیون سفارشی)',
    'پلیمرهای مهندسی و کامپاندها'
  ]
}

/* <-- new: detailed product lookup for chat replies --> */
const productDetails = {
  en: {
    'rotochem 0955w': {
      title: 'ROTOCHEM 0955W',
      short: 'A white polyethylene compound for rotational molding — excellent UV resistance, impact strength. Supplied in 20 kg bags. Datasheet available.'
    },
    'rotochem 0955b': {
      title: 'ROTOCHEM 0955B',
      short: 'Blue-grade rotational molding compound with good mechanical properties and stability. Supplied in 20 kg bags.'
    },
    'polyfil f700': {
      title: 'POLYFIL F700',
      short: 'HDPE compound optimized for blown film (10–25 µm). High tensile strength, low gel, suitable for food-contact films.'
    },
    'hdchem 4760': {
      title: 'HDCHEM 4760',
      short: 'Blow-molding polyethylene compound offering good flowability, ESCR and rigidity. Supplied in 25 kg bags.'
    },
    'slipchem e 178': {
      title: 'SlIPCHEM E 178',
      short: 'High-performance slip masterbatch to reduce COF in film processing. Pelleted form, good dispersion and thermal stability.'
    },
    'rafcolor 1560': {
      title: 'RAFCOLOR 1560',
      short: 'White masterbatch with high TiO₂ content — good opacity and dispersion for raffia and tapes.'
    },
    'calcichem 126 fp': {
      title: 'CALCICHEM 126 FP',
      short: 'Polypropylene filler masterbatch with ~80% CaCO₃ — high loading, good dispersion, for extrusion/injection uses.'
    },
    'calcichem 110 frf': {
      title: 'CALCICHEM 110 FRF',
      short: 'Fine CaCO₃ mineral modifier for films, raffia and ropes — improves productivity and reduces raw material cost.'
    },
    'calcichem 275 pm': {
      title: 'CALCICHEM 275 PM',
      short: 'Polypropylene mineral masterbatch (≈75% mineral) for BOPP/CPP/OPP; high loading and dispersion.'
    },
    'uvchem mb-r18': {
      title: 'UVChem MB-R18',
      short: 'UV stabilizer masterbatch for raffia — HALS + UV absorbers for long-term photo-stability. Typical dosing ~1 wt%.'
    }
  },
  fa: {
    'rotochem 0955w': {
      title: 'ROTOCHEM 0955W',
      short: 'کامپاند پلی‌اتیلن سفید برای قالب‌گیری چرخشی — مقاومت UV و استحکام ضربه‌ای خوب. بسته‌بندی: 20 کیلوگرمی. دیتاشیت موجود است.'
    },
    'rotochem 0955b': {
      title: 'ROTOCHEM 0955B',
      short: 'کامپاند آبی برای قالب‌گیری چرخشی با خواص مکانیکی و پایداری مناسب. بسته‌بندی: 20 کیلوگرمی.'
    },
    'polyfil f700': {
      title: 'POLYFIL F700',
      short: 'کامپاند HDPE جهت فیلم دمشی (10–25 میکرون). مقاومت کششی بالا، ژل کم، مناسب برای مصارف غذایی.'
    },
    'hdchem 4760': {
      title: 'HDCHEM 4760',
      short: 'کامپاند پلی‌اتیلن برای بلومولدینگ با جریان‌پذیری و سختی مناسب. بسته‌بندی: 25 کیلوگرمی.'
    },
    'slipchem e 178': {
      title: 'SlIPCHEM E 178',
      short: 'مستربچ لغزشی با پخش‌پذیری و پایداری حرارتی بالا برای کاهش ضریب اصطکاک بین لایه‌های فیلم.'
    },
    'rafcolor 1560': {
      title: 'RAFCOLOR 1560',
      short: 'مستربچ سفید با محتوای بالای TiO₂ — پوشش‌دهی و پراکندگی خوب برای رَفیا و نوارها.'
    },
    'calcichem 126 fp': {
      title: 'CALCICHEM 126 FP',
      short: 'مستربچ فیلر پایه پلی‌پروپیلن با تقریباً 80% CaCO₃ — بارگذاری بالا و پراکندگی خوب.'
    },
    'calcichem 110 frf': {
      title: 'CALCICHEM 110 FRF',
      short: 'ماده اصلاح‌کننده معدنی با ذرات بسیار ریز CaCO₃ برای فیلم‌ها و رَفیا — افزایش تولید و کاهش هزینه مواد اولیه.'
    },
    'calcichem 275 pm': {
      title: 'CALCICHEM 275 PM',
      short: 'مستربچ معدنی پلی‌پروپیلن (~75%) برای فیلم‌های BOPP/CPP/OPP با پراکندگی و بارگذاری بالا.'
    },
    'uvchem mb-r18': {
      title: 'UVChem MB-R18',
      short: 'مستربچ تثبیت‌کننده UV برای رَفیا — ترکیب HALS و جذب‌کننده UV برای محافظت بلندمدت. دز معمول ~1 wt%.'
    }
  }
}

/* <-- replaced: improved helper to detect product mention and return a response --> */

/* helper: normalize string (lowercase, remove non-alnum, collapse spaces) */
const normalize = (s = '') => {
  return s
    .toString()
    .toLowerCase()
    .normalize('NFKD') // decompose accents if any
    .replace(/[\u0300-\u036f]/g, '') // remove diacritics
    .replace(/[^a-z0-9\u0600-\u06FF]+/g, ' ') // keep latin letters, digits, persian chars; replace others with space
    .trim()
    .replace(/\s+/g, ' ')
}

/* helper: extract contiguous digits (useful to match product codes like 0955 or 955) */
const extractDigits = (s = '') => {
  const m = s.match(/\d+/g)
  return m ? m.join(' ') : ''
}

/* small levenshtein for fuzzy matching */
const levenshtein = (a = '', b = '') => {
  const A = a.split('')
  const B = b.split('')
  const m = A.length, n = B.length
  if (m === 0) return n
  if (n === 0) return m
  const d = Array.from({ length: m + 1 }, (_, i) => Array(n + 1).fill(0))
  for (let i = 0; i <= m; i++) d[i][0] = i
  for (let j = 0; j <= n; j++) d[0][j] = j
  for (let i = 1; i <= m; i++) {
    for (let j = 1; j <= n; j++) {
      const cost = A[i - 1] === B[j - 1] ? 0 : 1
      d[i][j] = Math.min(d[i - 1][j] + 1, d[i][j - 1] + 1, d[i - 1][j - 1] + cost)
    }
  }
  return d[m][n]
}

/* improved product matcher (supports partial/prefix token matching like "roto" -> "rotochem") */
const findProductResponse = (userText, lang) => {
  if (!userText) return null
  const raw = String(userText)
  const lower = normalize(raw)
  const digits = extractDigits(raw)

  const byLang = productDetails[lang] || productDetails.en

  // build candidate list with token arrays for efficient matching
  const candidates = Object.keys(byLang).map((key) => {
    const info = byLang[key]
    const title = (info && info.title) ? info.title : key
    const normKey = normalize(key)
    const normTitle = normalize(title)
    const keyTokens = normKey.split(' ').filter(Boolean)
    const titleTokens = normTitle.split(' ').filter(Boolean)
    return {
      key,
      normKey,
      normKeyNoSpace: normKey.replace(/\s+/g, ''),
      title,
      normTitle,
      normTitleNoSpace: normTitle.replace(/\s+/g, ''),
      tokens: Array.from(new Set([...keyTokens, ...titleTokens]))
    }
  })

  // 1) exact substring on normalized key/title (fast)
  for (const c of candidates) {
    if (lower.includes(c.normKey) || lower.includes(c.normKeyNoSpace) || lower.includes(c.normTitle) || lower.includes(c.normTitleNoSpace)) {
      return buildProductReply(byLang[c.key], lang)
    }
  }

  // 2) numeric code match (e.g., "0955" or "955")
  if (digits) {
    const digitsNorm = digits.replace(/\s+/g, '')
    const digitsNoZeros = digitsNorm.replace(/^0+/, '')
    for (const c of candidates) {
      if (c.normKey.includes(digitsNorm) || c.normTitle.includes(digitsNorm) || (digitsNoZeros && (c.normKey.includes(digitsNoZeros) || c.normTitle.includes(digitsNoZeros)))) {
        return buildProductReply(byLang[c.key], lang)
      }
    }
  }

  // token-level matching: exact tokens, prefix/partial, and short-fuzzy
  const userTokens = lower.split(' ').filter(Boolean)
  for (const t of userTokens) {
    if (!t) continue
    for (const c of candidates) {
      // exact token present
      if (c.tokens.includes(t)) return buildProductReply(byLang[c.key], lang)
      // prefix/partial matches (allow small tokens but require length >= 2)
      if (t.length >= 2) {
        for (const tk of c.tokens) {
          if (tk.startsWith(t) || tk.includes(t) || t.startsWith(tk)) {
            return buildProductReply(byLang[c.key], lang)
          }
        }
      }
      // numeric within token (e.g., "955" matching "0955")
      if (/\d/.test(t)) {
        const tDigits = t.replace(/\D/g, '')
        if (tDigits && (c.normKey.includes(tDigits) || c.normTitle.includes(tDigits))) {
          return buildProductReply(byLang[c.key], lang)
        }
      }
    }
  }

  // fallback: small fuzzy (levenshtein) checks between entire input and product keys/titles
  for (const c of candidates) {
    const distKey = levenshtein(lower, c.normKey)
    const distTitle = levenshtein(lower, c.normTitle)
    const threshold = Math.max(2, Math.floor(Math.min(c.normKey.length, lower.length) * 0.25))
    if (distKey <= threshold || distTitle <= threshold) {
      return buildProductReply(byLang[c.key], lang)
    }
    // token-to-token fuzzy (allow 1 edit)
    for (const t of userTokens) {
      if (!t) continue
      for (const tk of c.tokens) {
        if (levenshtein(t, tk) <= 1) {
          return buildProductReply(byLang[c.key], lang)
        }
      }
    }
  }

  return null
}

/* small helpers used above */
const keyTokensSafe = (s = '') => (s ? s.split(' ').filter(Boolean) : [])

/* build reply text consistently */
const buildProductReply = (info, lang) => {
  if (!info) return null
  if (lang === 'fa') {
    return `${info.title}\n\n${info.short}\n\nبرای دریافت دیتاشیت تایپ کنید: دیتاشیت ${info.title} یا بپرسید قیمت و مقدار مورد نیاز را.`
  } else {
    return `${info.title}\n\n${info.short}\n\nTo get the datasheet type: datasheet ${info.title} or ask for price/quantity.`
  }
}

/* Enhanced Pattern Matching */
const patterns = {
  en: {
    greeting: ['hi', 'hello', 'hey', 'good morning', 'good afternoon', 'good evening', 'greetings', 'salaam'],
    products: ['product', 'polymer', 'compound', 'masterbatch', 'filler', 'color', 'additive', 'white', 'blend', 'catalog', 'pp', 'pe', 'ps', 'pvc', 'talc', 'caco3', 'titanium', 'tio2', 'antimicrobial', 'flame retardant'],
    price: ['price', 'cost', 'how much', 'rate', 'dollar', 'euro', 'toman', 'quotation', 'quote', 'pricing'],
    order: ['order', 'purchase', 'buy', 'need', 'want', 'request', 'supply', 'minimum'],
    technical: ['specification', 'technical', 'datasheet', 'properties', 'specs', 'density', 'melt', 'viscosity', 'mfi', 'mfr', 'tensile', 'quality', 'certificate'],
    delivery: ['delivery', 'shipping', 'transport', 'time', 'days', 'lead time', 'when', 'how long'],
    contact: ['contact', 'phone', 'address', 'email', 'call', 'reach', 'location', 'office'],
    samples: ['sample', 'test', 'trial', 'sampling', 'free sample'],
    thanks: ['thank', 'thanks', 'appreciate', 'thx', 'grateful'],
    about: ['about', 'company', 'who are you', 'history', 'established', 'experience']
  },
  fa: {
    greeting: ['سلام', 'درود', 'صبح بخیر', 'عصر بخیر', 'هی', 'سلام علیکم'],
    products: ['محصول', 'محصولات', 'پلیمر', 'کامپاند', 'مستربچ', 'فیلر', 'رنگ', 'افزودنی', 'سفید', 'بلند', 'کاتالوگ', 'تالک', 'کلسیم کربنات', 'تیتانیوم', 'ضد میکروب', 'شعله گیر'],
    price: ['قیمت', 'هزینه', 'نرخ', 'چند', 'چقدر', 'دلار', 'یورو', 'تومان'],
    order: ['سفارش', 'خرید', 'نیاز', 'می‌خوام', 'تامین', 'حداقل'],
    technical: ['مشخصات', 'فنی', 'دیتاشیت', 'ویژگی', 'خواص', 'دانسیته', 'کیفیت', 'گواهی'],
    delivery: ['ارسال', 'تحویل', 'حمل', 'زمان', 'روز', 'کی', 'چقدر طول'],
    contact: ['تماس', 'شماره', 'آدرس', 'ایمیل', 'موقعیت', 'دفتر'],
    samples: ['نمونه', 'سمپل', 'تست', 'آزمایش'],
    thanks: ['ممنون', 'متشکر', 'مرسی', 'سپاس'],
    about: ['درباره', 'شرکت', 'کی هستید', 'تاریخچه', 'تاسیس', 'تجربه']
  }
}

/* Enhanced Responses */
const responses = {
  en: {
    greeting: [
      `Hello!  Welcome to POLYCHEM. We are specialists in advanced polymer compounds and masterbatches since 2015. \n\nHow can I help you today? You can ask about:\n• Our products (compounds, masterbatches)\n• Technical specifications\n• Pricing & orders\n• Samples`,
      `Hi there! I'm your POLYCHEM sales assistant. We produce high-quality polymer compounds and masterbatches for various industries.\n\nWhat would you like to know?`,
      `Welcome to POLYCHEM! 🧪 We specialize in engineering polymers and custom compounds. How may I assist you?`
    ],
    products: [
      `POLYCHEM offers a comprehensive range of polymer solutions:\n\n📦 **Main Products:**\n• Polymer Compounds (PP-Talc for automotive/appliances)\n• Filler Masterbatch (CaCo3 based)\n• Color Masterbatch (all colors)\n• Additive Masterbatch (anti-block, flame retardant, antimicrobial)\n• White Masterbatch (TiO2)\n• Custom Polymer Blends\n\nWhich product interests you?`,
      `We manufacture specialized polymer products:\n\n✓ **Compounds**: PP-Talc compounds for high-temperature applications\n✓ **Masterbatches**: Filler, Color, Additive, White\n✓ **Engineering Polymers**: Custom formulations\n\nNeed technical details on any specific product?`,
      `Our product portfolio includes:\n• Advanced polymer compounds\n• Specialized masterbatches (filler, color, additive)\n• Custom engineering polymers\n\nAll products come with full technical datasheets and quality certificates. What can I help you find?`
    ],
    price: [
      `For accurate pricing, please contact our sales team:\n\n📞 **Phone**: ${companyInfo.phone}\n📧 **Email**: ${companyInfo.email}\n\nPrices vary based on:\n• Product type & grade\n• Order quantity (MOQ applies)\n• Delivery terms\n• Destination\n\nWould you like me to have a sales representative contact you?`,
      `Pricing depends on several factors including volume and specifications. For a detailed quotation:\n\n**Contact**: ${companyInfo.phone}\n**Email**: ${companyInfo.email}\n\nOur team will provide competitive pricing within 24 hours.`,
      `To get the best price for your requirements, please reach out to our sales department:\n\n${companyInfo.phone}\n${companyInfo.email}\n\nWe offer volume discounts and flexible payment terms.`
    ],
    order: [
      `Great! To process your order efficiently, we'll need:\n\n✓ Product code/specification\n✓ Required quantity (MOQ varies by product)\n✓ Delivery address\n✓ Preferred delivery terms (FOB, CIF, etc.)\n\nPlease contact: ${companyInfo.phone}\nEmail: ${companyInfo.email}`,
      `To place an order:\n\n1️⃣ Specify the product & quantity\n2️⃣ Confirm technical requirements\n3️⃣ Provide delivery details\n\n**Sales**: ${companyInfo.phone}\n**Email**: ${companyInfo.email}\n\nTypical lead time: 7-14 days depending on product availability.`,
      `We're ready to process your order! Minimum order quantities vary by product (typically 1-5 tons).\n\nContact our sales team:\n📞 ${companyInfo.phone}\n📧 ${companyInfo.email}\n\nThey'll guide you through the process and confirm availability.`
    ],
    technical: [
      `All our products include comprehensive technical documentation:\n\n📄 **Available Documents:**\n• Technical Datasheets (TDS)\n• Material Safety Data Sheets (MSDS)\n• Quality Certificates\n• Test Reports\n• Processing Guidelines\n\nPlease email ${companyInfo.email} with your specific requirements, and we'll send the technical files within 24 hours.`,
      `For technical specifications and datasheets:\n\n**Email**: ${companyInfo.email}\n**Phone**: ${companyInfo.phone}\n\nOur engineers can provide:\n• Detailed product specs (density, MFI, tensile strength, etc.)\n• Processing parameters\n• Application guidelines\n• Custom formulation support\n\nWhat specific technical information do you need?`,
      `Technical support is available! We provide:\n✓ Complete datasheets\n✓ Processing recommendations\n✓ Application engineering support\n✓ Quality certificates\n\nContact: ${companyInfo.email}\nOur technical team will assist with any specifications you need.`
    ],
    delivery: [
      `**Delivery Information:**\n\n🚚 **Iran Domestic:**\n• Tehran: 2-3 business days\n• Other cities: 5-7 business days\n• Express shipping available\n\n🌍 **International:**\n• Lead time varies by destination\n• FOB, CIF, DDP terms available\n\nContact ${companyInfo.phone} for shipping quotes and logistics support.`,
      `Shipping is arranged based on your location:\n\n**Domestic**: Fast delivery across Iran (2-7 days)\n**Export**: Worldwide shipping with full documentation\n\nAll shipments include:\n✓ Insured transport\n✓ Quality packaging\n✓ Tracking information\n\nFor shipping costs: ${companyInfo.phone}`,
      `We offer flexible delivery options:\n• Warehouse pickup (Tehran)\n• Direct delivery to your facility\n• International freight forwarding\n\nTypical lead times: 7-14 days (production + delivery)\n\nContact for logistics: ${companyInfo.phone}`
    ],
    contact: [
      `**POLYCHEM Contact Information:**\n\n📞 **Phone**: ${companyInfo.phone}\n📠 **Fax**: ${companyInfo.fax}\n📧 **Email**: ${companyInfo.email}\n🌐 **Website**: ${companyInfo.website}\n\n📍 **Address**:\n${companyInfo.address}\n\nBusiness hours: Saturday-Thursday, 9:00-17:00 (Tehran time)`,
      `**Get in Touch:**\n\n**Phone**: ${companyInfo.phone}\n**Email**: ${companyInfo.email}\n**Location**: Tehran, Iran\n\nOur sales team responds within 24 hours!\n\nPrefer a call back? Share your number and best time to reach you.`,
      `**Contact POLYCHEM:**\n\n☎️ ${companyInfo.phone}\n✉️ ${companyInfo.email}\n\n**Office Location**:\nUnit 15, NO.45, Manzarnejad Blvd\nShariati Ave, Tehran, Iran\n\nWe're here to help Monday-Thursday!`
    ],
    samples: [
      `Yes! We provide samples for quality testing.\n\n**Sample Policy:**\n✓ Small quantities available for new customers\n✓ Technical evaluation samples\n✓ Usually dispatched within 2-3 business days\n✓ Shipping costs may apply for international orders\n\nTo request samples:\n📧 Email: ${companyInfo.email}\n📞 Phone: ${companyInfo.phone}\n\nPlease specify:\n• Product type\n• Required quantity\n• Delivery address\n• Application details`,
      `Sample shipments available! 📦\n\nWe understand the importance of testing before ordering. Contact us with:\n\n1. Product specification\n2. Sample quantity needed\n3. Your company details\n4. Delivery address\n\n**Contact**: ${companyInfo.email}\n\nSamples typically ship within 2-3 business days after confirmation.`,
      `Free samples available for qualified customers! 🎁\n\nTo arrange sample delivery:\n\n**Email**: ${companyInfo.email}\n**Phone**: ${companyInfo.phone}\n\nProvide your application details and delivery address. Our team will coordinate the shipment.`
    ],
    thanks: [
      `You're very welcome! 😊\n\nIf you have any other questions about our polymer compounds or masterbatches, feel free to ask!\n\nOr contact us directly:\n📞 ${companyInfo.phone}\n📧 ${companyInfo.email}`,
      `Happy to help! Is there anything else you'd like to know about POLYCHEM products or services?`,
      `You're welcome! We're always here to assist. Don't hesitate to reach out if you need anything else! 🤝`
    ],
    about: [
      `**About POLYCHEM:**\n\nEstablished in 2015, POLYCHEM specializes in producing advanced engineering polymers and compounds. Located in the Aras Free Zone, we serve domestic and international markets.\n\n🏆 **Our Expertise:**\n• Advanced polymer compounds\n• Specialized masterbatches\n• Custom formulations\n• Technical consulting\n\n👥 **Team**: Experienced engineers with decades of expertise in polymers and petrochemicals\n\n🎯 **Mission**: Provide high-quality polymer solutions to meet industrial needs\n\nWant to know more about specific products?`,
      `POLYCHEM (est. 2015) is a leading producer of:\n\n✓ Advanced polymer compounds\n✓ Specialized masterbatches\n✓ Custom engineering polymers\n\nOur experienced team of engineers and university professors develop innovative polymer solutions for various industries including automotive, packaging, and home appliances.\n\nHow can we serve your polymer needs?`
    ],
    default: [
      `I'd be happy to help! Could you please provide more details?\n\nI can assist you with:\n• Product information\n• Technical specifications\n• Pricing & quotations\n• Orders & samples\n• Contact information\n\nWhat would you like to know?`,
      `I'm not quite sure I understand. Could you rephrase your question?\n\nYou can ask me about:\n📦 Products (compounds, masterbatches)\n💰 Pricing\n📋 Technical specs\n🚚 Delivery\n📞 Contact info`,
      `For detailed assistance, please contact our team:\n\n📞 ${companyInfo.phone}\n📧 ${companyInfo.email}\n\nOr ask me about products, pricing, technical info, or samples!`
    ]
  },
  fa: {
    greeting: [
      `سلام! 👋 به پلیکم خوش آمدید. ما تولیدکننده کامپاندها و مستربچ‌های پلیمری پیشرفته هستیم (از سال 1394).\n\nچطور می‌تونم کمکتون کنم؟ می‌تونید بپرسید:\n• محصولات (کامپاند، مستربچ)\n• مشخصات فنی\n• قیمت و سفارش\n• نمونه`,
      `سلام! من دستیار فروش پلیکم هستم. ما تولیدکننده کامپاندها و مستربچ‌های باکیفیت برای صنایع مختلف هستیم.\n\nچه طور میتونم راهنمایتون بکنم ؟`,
      `به پلیکم خوش اومدید! 🧪 ما متخصص پلیمرهای مهندسی و کامپاندهای سفارشی هستیم. چطور می‌تونم کمکتون کنم؟`
    ],
    products: [
      `پلیکم طیف کاملی از محصولات پلیمری ارائه می‌دهد:\n\n📦 **محصولات اصلی:**\n• کامپاندهای پلیمری (PP-تالک برای خودرو/لوازم خانگی)\n• مستربچ فیلر (پایه CaCo3)\n• مستربچ رنگی (تمام رنگ‌ها)\n• مستربچ افزودنی (ضد بلوک، شعله‌گیر، ضدمیکروب)\n• مستربچ سفید (TiO2)\n• بلندهای پلیمری سفارشی\n\nکدوم محصول موردنظرتونه؟`,
      `ما محصولات پلیمری تخصصی تولید می‌کنیم:\n\n✓ **کامپاندها**: PP-تالک برای کاربردهای دمای بالا\n✓ **مستربچ‌ها**: فیلر، رنگی، افزودنی، سفید\n✓ **پلیمرهای مهندسی**: فرمولاسیون سفارشی\n\nبرای جزئیات فنی کدوم محصول اطلاعات می‌خواید؟`,
      `محصولات ما شامل:\n• کامپاندهای پلیمری پیشرفته\n• مستربچ‌های تخصصی (فیلر، رنگی، افزودنی)\n• پلیمرهای مهندسی سفارشی\n\nتمام محصولات با دیتاشیت کامل و گواهی کیفیت. چی می‌تونم پیدا کنم براتون؟`
    ],
    price: [
      `برای دریافت قیمت دقیق، لطفاً با تیم فروش ما تماس بگیرید:\n\n📞 **تلفن**: ${companyInfo.phone}\n📧 **ایمیل**: ${companyInfo.email}\n\nقیمت‌ها بستگی دارد به:\n• نوع و گرید محصول\n• مقدار سفارش (حداقل سفارش اعمال می‌شود)\n• شرایط تحویل\n• مقصد\n\nمی‌خواید یک نماینده فروش باهاتون تماس بگیره؟`,
      `قیمت بستگی به عوامل مختلفی از جمله حجم و مشخصات داره. برای پیشنهاد قیمت دقیق:\n\n**تماس**: ${companyInfo.phone}\n**ایمیل**: ${companyInfo.email}\n\nتیم ما ظرف 24 ساعت قیمت رقابتی ارائه می‌ده.`,
      `برای دریافت بهترین قیمت متناسب با نیازتون، لطفاً با بخش فروش تماس بگیرید:\n\n${companyInfo.phone}\n${companyInfo.email}\n\nما تخفیف حجمی و شرایط پرداخت منعطف داریم.`
    ],
    order: [
      `عالیه! برای پردازش سفارش به موارد زیر نیاز داریم:\n\n✓ کد/مشخصات محصول\n✓ مقدار موردنیاز (حداقل سفارش متغیره)\n✓ آدرس تحویل\n✓ شرایط تحویل ترجیحی (FOB، CIF و غیره)\n\nلطفاً تماس بگیرید: ${companyInfo.phone}\nایمیل: ${companyInfo.email}`,
      `برای ثبت سفارش:\n\n1️⃣ محصول و مقدار رو مشخص کنید\n2️⃣ نیازهای فنی رو تایید کنید\n3️⃣ جزئیات تحویل رو ارائه بدید\n\n**فروش**: ${companyInfo.phone}\n**ایمیل**: ${companyInfo.email}\n\nزمان تحویل معمولی: 7-14 روز بسته به موجودی محصول.`,
      `ما آماده پردازش سفارش شما هستیم! حداقل سفارش بسته به محصول متفاته (معمولاً 1-5 تن).\n\nتماس با تیم فروش:\n📞 ${companyInfo.phone}\n📧 ${companyInfo.email}\n\nآنها شما رو راهنمایی می‌کنند و موجودی رو تایید می‌کنن.`
    ],
    technical: [
      `تمام محصولات ما شامل مستندات فنی جامع:\n\n📄 **اسناد موجود:**\n• دیتاشیت فنی (TDS)\n• برگه اطلاعات ایمنی مواد (MSDS)\n• گواهی‌های کیفیت\n• گزارش‌های تست\n• راهنمای پردازش\n\nلطفاً به ${companyInfo.email} ایمیل بزنید با نیازهای خاصتون، ما فایل‌های فنی رو ظرف 24 ساعت ارسال می‌کنیم.`,
      `برای مشخصات فنی و دیتاشیت:\n\n**ایمیل**: ${companyInfo.email}\n**تلفن**: ${companyInfo.phone}\n\nمهندسان ما می‌تونن ارائه بدن:\n• مشخصات دقیق محصول (دانسیته، MFI، استحکام کششی و...)\n• پارامترهای پردازش\n• راهنمای کاربرد\n• پشتیبانی فرمولاسیون سفارشی\n\nچه اطلاعات فنی خاصی نیاز دارید؟`,
      `پشتیبانی فنی موجوده! ما ارائه می‌دیم:\n✓ دیتاشیت‌های کامل\n✓ توصیه‌های پردازش\n✓ پشتیبانی مهندسی کاربرد\n✓ گواهی‌های کیفیت\n\nتماس: ${companyInfo.email}\nتیم فنی ما با هر مشخصاتی که نیاز دارید کمکتون می‌کنه.`
    ],
    delivery: [
      `**اطلاعات تحویل:**\n\n🚚 **داخلی ایران:**\n• تهران: 2-3 روز کاری\n• شهرستان‌ها: 5-7 روز کاری\n• ارسال فوری موجود\n\n🌍 **بین‌المللی:**\n• زمان تحویل بسته به مقصد متفاته\n• شرایط FOB، CIF، DDP موجود\n\nتماس ${companyInfo.phone} برای قیمت حمل و پشتیبانی لجستیک.`,
      `ارسال بر اساس موقعیت شما ترتیب داده میشه:\n\n**داخلی**: تحویل سریع در سراسر ایران (2-7 روز)\n**صادرات**: حمل به سراسر جهان با مستندات کامل\n\nتمام محموله‌ها شامل:\n✓ حمل بیمه شده\n✓ بسته‌بندی باکیفیت\n✓ اطلاعات ردیابی\n\nبرای هزینه حمل: ${companyInfo.phone}`,
      `ما گزینه‌های تحویل انعطاف‌پذیر داریم:\n• تحویل از انبار (تهران)\n• تحویل مستقیم به مجموعه شما\n• حمل بین‌المللی\n\nزمان تحویل معمولی: 7-14 روز (تولید + ارسال)\n\nتماس برای لجستیک: ${companyInfo.phone}`
    ],
    contact: [
      `**اطلاعات تماس پلیکم:**\n\n📞 **تلفن**: ${companyInfo.phone}\n📠 **فکس**: ${companyInfo.fax}\n📧 **ایمیل**: ${companyInfo.email}\n🌐 **وبسایت**: ${companyInfo.website}\n\n📍 **آدرس**:\n${companyInfo.address}\n\nساعات کاری: شنبه-پنج‌شنبه، 9:00-17:00 (وقت تهران)`,
      `**در تماس باشید:**\n\n**تلفن**: ${companyInfo.phone}\n**ایمیل**: ${companyInfo.email}\n**موقعیت**: تهران، ایران\n\nتیم فروش ما ظرف 24 ساعت پاسخ می‌ده!\n\nترجیح می‌دید تماس بگیریم؟ شماره و بهترین زمان تماستون رو بفرمایید.`,
      `**تماس با پلیکم:**\n\n☎️ ${companyInfo.phone}\n✉️ ${companyInfo.email}\n\n**دفتر**:\nواحد 15، شماره 45، بلوار منظرنژاد\nخیابان شریعتی، تهران، ایران\n\nدوشنبه تا پنج‌شنبه در خدمتیم!`
    ],
    samples: [
      `بله! ما نمونه برای تست کیفیت ارائه می‌دیم.\n\n**سیاست نمونه:**\n✓ مقادیر کوچک برای مشتریان جدید موجوده\n✓ نمونه‌های ارزیابی فنی\n✓ معمولاً ظرف 2-3 روز کاری ارسال می‌شه\n✓ هزینه ارسال ممکنه برای سفارشات بین‌المللی اعمال بشه\n\nبرای درخواست نمونه:\n📧 ایمیل: ${companyInfo.email}\n📞 تلفن: ${companyInfo.phone}\n\nلطفاً مشخص کنید:\n• نوع محصول\n• مقدار موردنیاز\n• آدرس تحویل\n• جزئیات کاربرد`,
      `ارسال نمونه موجوده! 📦\n\nما اهمیت تست قبل از سفارش رو درک می‌کنیم. با ما تماس بگیرید با:\n\n1. مشخصات محصول\n2. مقدار نمونه موردنیاز\n3. اطلاعات شرکت شما\n4. آدرس تحویل\n\n**تماس**: ${companyInfo.email}\n\nنمونه‌ها معمولاً ظرف 2-3 روز کاری بعد از تایید ارسال می‌شن.`,
      `نمونه رایگان برای مشتریان واجد شرایط! 🎁\n\nبرای ترتیب دادن ارسال نمونه:\n\n**ایمیل**: ${companyInfo.email}\n**تلفن**: ${companyInfo.phone}\n\nجزئیات کاربرد و آدرس تحویلتون رو ارائه بدید. تیم ما ارسال رو هماهنگ می‌کنه.`
    ],
    thanks: [
      `خواهش می‌کنم! 😊\n\nاگه سوال دیگه‌ای درباره کامپاندها یا مستربچ‌های پلیمری داشتید، بپرسید!\n\nیا مستقیماً تماس بگیرید:\n📞 ${companyInfo.phone}\n📧 ${companyInfo.email}`,
      `خوشحالم که کمک کردم! چیز دیگه‌ای درباره محصولات یا خدمات پلیکم می‌خواید بدونید؟`,
      `خواهش می‌کنم! همیشه اینجاییم برای کمک. اگه نیاز به چیز دیگه‌ای داشتید تردید نکنید! 🤝`
    ],
    about: [
      `**درباره پلیکم:**\n\nتاسیس در 1394، پلیکم متخصص تولید پلیمرها و کامپاندهای مهندسی پیشرفته. واقع در منطقه آزاد ارس، به بازارهای داخلی و بین‌المللی خدمت می‌دهیم.\n\n🏆 **تخصص ما:**\n• کامپاندهای پلیمری پیشرفته\n• مستربچ‌های تخصصی\n• فرمولاسیون سفارشی\n• مشاوره فنی\n\n👥 **تیم**: مهندسان باتجربه با دهه‌ها تخصص در پلیمر و پتروشیمی\n\n🎯 **ماموریت**: ارائه راهکارهای پلیمری باکیفیت برای نیازهای صنعتی\n\nمی‌خواید بیشتر درباره محصولات خاص بدونید؟`,
      `پلیکم (تاسیس 1394) تولیدکننده پیشرو:\n\n✓ کامپاندهای پلیمری پیشرفته\n✓ مستربچ‌های تخصصی\n✓ پلیمرهای مهندسی سفارشی\n\nتیم باتجربه مهندسان و اساتید دانشگاه ما راهکارهای نوآورانه پلیمری برای صنایع مختلف از جمله خودرو، بسته‌بندی و لوازم خانگی توسعه می‌دن.\n\nچطور می‌تونیم به نیازهای پلیمری شما خدمت کنیم؟`
    ],
    default: [
      `خوشحال می‌شم کمکتون کنم! می‌شه بیشتر توضیح بدید؟\n\nمی‌تونم کمکتون کنم با:\n• اطلاعات محصول\n• مشخصات فنی\n• قیمت و پیشنهاد\n• سفارشات و نمونه‌ها\n• اطلاعات تماس\n\nچه طور میتونم راهنمایتون بکنم ؟`,
      `متوجه نشدم. می‌شه سوالتون رو دوباره بیان کنید؟\n\nمی‌تونید از من بپرسید درباره:\n📦 محصولات (کامپاند، مستربچ)\n💰 قیمت\n📋 مشخصات فنی\n🚚 تحویل\n📞 اطلاعات تماس`,
      `برای کمک دقیق، لطفاً با تیم ما تماس بگیرید:\n\n📞 ${companyInfo.phone}\n📧 ${companyInfo.email}\n\nیا از من بپرسید درباره محصولات، قیمت، اطلاعات فنی یا نمونه!`
    ]
  }
}

/* Helpers */
const detectLanguage = (text) => {
  const persianPattern = /[\u0600-\u06FF]/
  return persianPattern.test(text) ? 'fa' : 'en'
}

/* Modify: check product mention first, then fallback to category patterns */
const getResponseFor = (userText, lang) => {
  // check product names first
  const productReply = findProductResponse(userText, lang)
  if (productReply) return productReply

  const lower = userText.toLowerCase()
  const currentPatterns = patterns[lang] || patterns.en
  const currentResponses = responses[lang] || responses.en

  for (const [category, keywords] of Object.entries(currentPatterns)) {
    if (keywords.some(k => lower.includes(k))) {
      const arr = currentResponses[category] || currentResponses.default
      return arr[Math.floor(Math.random() * arr.length)]
    }
  }
  
  const arr = currentResponses.default
  return arr[Math.floor(Math.random() * arr.length)]
}

/* Send Message */
const sendMessage = async () => {
  const text = input.value.trim()
  if (!text) return

  const detected = detectLanguage(text)
  language.value = detected

  messages.value.push({
    id: Date.now(),
    sender: 'user',
    text: escapeHtml(text),
    lang: detected
  })

  input.value = ''
  isTyping.value = true

  await nextTick()
  if (chatMessages.value) chatMessages.value.scrollTop = chatMessages.value.scrollHeight

  setTimeout(async () => {
    const reply = await Promise.resolve(getResponseFor(text, detected))
    isTyping.value = false

    messages.value.push({
      id: Date.now() + 1,
      sender: 'bot',
      text: formatResponse(reply),
      lang: detected
    })

    await nextTick()
    if (chatMessages.value) chatMessages.value.scrollTop = chatMessages.value.scrollHeight

    if (minimized.value) {
      unreadCount.value += 1
    }
  }, 1000 + Math.random() * 800)
}

const openChat = () => {
  minimized.value = false
  unreadCount.value = 0
  // focus input next tick
  nextTick(() => {
    if (chatInputRef.value) chatInputRef.value.focus()
  })
}

const closeChat = () => {
  minimized.value = true
}

const clearChat = () => {
  messages.value = [
    { 
      id: Date.now(), 
      sender: 'bot', 
      text: language.value === 'fa' 
        ? 'سلام! 👋 من اینجام برای کمک. چطور می‌تونم کمکتون کنم؟'
        : 'Hello! 👋 I am here to assist. How can I help you?', 
      lang: language.value 
    }
  ]
}

/* Utilities */
const escapeHtml = (s) => {
  return s.replace(/[&<>"']/g, (c) => ({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&#39;'}[c]))
}
const formatResponse = (s) => {
  // keep line breaks => <br>
  return s.replace(/\n/g, '<br/>')
}

/* Auto-scroll */
watch([messages, isTyping], async () => {
  await nextTick()
  if (chatMessages.value) {
    chatMessages.value.scrollTop = chatMessages.value.scrollHeight
  }
})

/* Refs */
const chatInputRef = ref(null)
onMounted(() => {
  // initial focus not required; keep as is
})
onUnmounted(() => {
  // cleanup if needed
})
</script>

<style scoped>
/* Use the CSS from your provided HTML to match appearance exactly */

/* ...existing code... replaced with new style ... */

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.chat-widget {
  position: fixed;
  bottom: 30px;
  right: 30px;
  z-index: 99999 !important; /* ensure chat sits above navbar */
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

/* Floating Button */
.chat-fab {
  z-index: 100001; /* keep the button itself above everything */
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: #ffd000;
  /* thin gray stroke around the circular button */
  border: 3px solid #848484;
   cursor: pointer;
   display: flex;
   align-items: center;
   justify-content: center;
   transition: all 0.3s ease;
   position: relative;
   color: #ffffff;
}

.chat-fab:hover {
  transform: scale(1.1);
  box-shadow: 0 12px 32px rgba(251, 191, 36, 0.6);
}

/* ensure the public svg fits nicely inside the fab - larger for better visibility */
.chat-fab img,
.chat-fab .fab-img {
  width: 62px;
  height: 62px;
  object-fit: contain;
  display: block;
}

/* unread badge sizing/position */
.unread-badge {
  position: absolute;
  top: -8px;
  right: -8px;
  background: #ef4444;
  color: white;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: bold;
  border: 4px solid white;
}

/* Chat Window */
.chat-window {
  position: fixed;
  bottom: 30px;
  right: 30px;
  width: 420px;
  height: 650px;
  background: white;
  border-radius: 24px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15);
  display: none;
  flex-direction: column;
  overflow: hidden;
}

.chat-window.active {
  display: flex;
  animation: slideUp 0.4s cubic-bezier(0.25, 0.8, 0.25, 1);
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px) scale(0.9);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

/* Header */
.chat-header {
  background: #848484;
  padding: 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 14px;
}

.header-avatar {
  width: 50px;
  height: 50px;
  background: #bdbbbb;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden; /* added so image is clipped to circle */
}

.header-avatar img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  display: block;
}

.header-info h2 {
  font-size: 19px;
  font-weight: 700;
  color: #ffffff;
  margin-bottom: 4px;
}

.header-status {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #ffffff;
}

.status-dot {
  width: 9px;
  height: 9px;
  background: #10b981;
  border-radius: 50%;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% {
    box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.7);
  }
  50% {
    box-shadow: 0 0 0 6px rgba(16, 185, 129, 0);
  }
}

.header-actions {
  display: flex;
  gap: 8px;
}

.header-btn {
  width: 38px;
  height: 38px;
  background: #ffd000;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
 
}

.header-btn:hover {
  background: rgba(0, 0, 0, 0.2);
  transform: scale(1.1);
}

.header-btn svg {
  width: 18px;
  height: 18px;
  color: #757575;
}

/* Messages Area */
.chat-messages {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
  background: #f1f2f2;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.chat-messages::-webkit-scrollbar {
  width: 8px;
}

.chat-messages::-webkit-scrollbar-thumb {
  background: #d1d1d1;
  border-radius: 10px;
}

.chat-messages::-webkit-scrollbar-thumb:hover {
  background: #848484;
}

/* Message Bubbles */
.message {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  animation: messageIn 0.3s ease;
}

@keyframes messageIn {
  from {
    opacity: 0;
    transform: translateY(15px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.message.user {
  flex-direction: row-reverse;
}

.message-avatar {
  width: 36px;
  height: 36px;
  background: #bdbbbb;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.message.user .message-avatar {
  background:#ffd000;
}

.message-avatar svg {
  width: 18px;
  height: 18px;
  color: white;
}

.message-bubble {
  max-width: 70%;
  padding: 14px 18px;
  border-radius: 20px;
  font-size: 15px;
  line-height: 1.5;
  word-wrap: break-word;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  background: white;
  color: #848484;
}

.message.bot .message-bubble {
  border-bottom-left-radius: 6px;
}

.message.user .message-bubble {
  background: #ffd000;
  color: #ffffff;
  border-bottom-right-radius: 6px;
  font-weight: 500;
}

/* Typing Indicator */
.typing-indicator {
  display: flex;
  gap: 6px;
  padding: 14px 18px;
  background: white;
  border-radius: 20px;
  border-bottom-left-radius: 6px;
  width: fit-content;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.typing-dot {
  width: 10px;
  height: 10px;
  background: #848484;
  border-radius: 50%;
  animation: typingBounce 1.4s ease-in-out infinite;
}

.typing-dot:nth-child(1) {
  animation-delay: 0s;
}

.typing-dot:nth-child(2) {
  animation-delay: 0.2s;
}

.typing-dot:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes typingBounce {
  0%, 60%, 100% {
    transform: translateY(0);
  }
  30% {
    transform: translateY(-12px);
  }
}

/* Input Area */
.chat-input-area {
  padding: 20px 24px;
  background: white;
  border-top: 2px solid #e5e5e5;
  display: flex;
  gap: 12px;
  align-items: center;
}

.chat-input {
  flex: 1;
  padding: 14px 18px;
  border: 2px solid #d1d1d1;
  border-radius: 25px;
  font-size: 15px;
  outline: none;
  transition: all 0.2s;
  background: #f9f9f9;
}

.chat-input:focus {
  border-color: #ffd000;
  background: white;
  box-shadow: 0 0 0 4px rgba(251, 191, 36, 0.1);
}

.chat-input::placeholder {
  color: #999;
}

.send-btn {
  width: 52px;
  height: 52px;
  background: #ffd000;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  box-shadow: 0 4px 12px rgba(251, 191, 36, 0.4);
}

.send-btn:hover {
  transform: scale(1.1);
  box-shadow: 0 6px 16px rgba(251, 191, 36, 0.5);
}

.send-btn:active {
  transform: scale(0.95);
}

.send-btn svg {
  width: 22px;
  height: 22px;
  color: #ffffff;
}

/* make logo images appear white */
.logo-white {
  filter: brightness(0) invert(1);
  /* keep size/fit rules intact */
  display: block;
}

/* Mobile: lift the chat widget above the bottom navbar to avoid overlap.
   navbar-mob uses height ~80px and bottom ~20px -> reserve that space */
@media (max-width: 600px) {
  .chat-widget {
    /* قبلی: bottom: calc(20px + 80px + 12px); */
    bottom: calc(20px + 80px + 12px + 18px); /* افزایش 18px تا دکمه بالاتر قرار گیرد */
    right: 20px;
  }

  .chat-fab {
    /* slightly smaller on small screens but remain above navbar */
    width: 80px;
    height: 80px;
  }

  .chat-fab img,
  .chat-fab .fab-img {
    width: 56px;
    height: 56px;
  }

  .unread-badge {
    top: -6px;
    right: -6px;
  }
}

/* ---------------------------
   Improved mobile responsiveness
   --------------------------- */

/* Use safe-area for devices with bottom bars and ensure chat sits above navbar */
@media (max-width: 768px) {
  .chat-widget {
    /* place widget above mobile navbar (use CSS var fallback if available) */
    bottom: calc(env(safe-area-inset-bottom, 0px) + 120px); /* افزایش به 120px برای فاصله بیشتر از navbar */
    right: 16px;
    left: auto;
    z-index: 120000 !important;
  }

  /* Slightly smaller FAB on most phones, keep clear stroke */
  .chat-fab {
    width: 68px;
    height: 68px;
    border: 2px solid #c8c8c8;
    border-radius: 50%;
    box-shadow: 0 10px 26px rgba(0,0,0,0.10);
  }

  .chat-fab img,
  .chat-fab .fab-img {
    width: 52px;
    height: 52px;
  }

  .unread-badge {
    top: -6px;
    right: -6px;
    width: 28px;
    height: 28px;
    font-size: 12px;
    border: 3px solid white;
  }

  /* Make the chat window feel like a mobile panel (centered with side gutters) */
  .chat-window {
    position: fixed;
    left: 12px;
    right: 12px;
    bottom: calc(env(safe-area-inset-bottom, 0px) + 110px); /* sit above fab/navbar */
    width: calc(100% - 24px);
    height: calc(70vh); /* comfortable height on mobile */
    max-height: calc(100vh - (env(safe-area-inset-top, 0px) + 140px));
    border-radius: 14px;
    display: none;
    overflow: hidden;
    box-shadow: 0 18px 50px rgba(0,0,0,0.18);
  }

  .chat-window.active {
    display: flex;
    animation: slideUp 220ms cubic-bezier(.2,.9,.2,1);
  }

  /* Reduce header / paddings for mobile */
  .chat-header {
    padding: 12px 14px;
    gap: 10px;
  }

  .header-avatar {
    width: 42px;
    height: 42px;
  }

  .header-info h2 {
    font-size: 15px;
  }

  .header-btn {
    width: 34px;
    height: 34px;
  }

  /* Messages region - allow larger scroll area and smoother touch scrolling */
  .chat-messages {
    padding: 14px;
    gap: 12px;
    overflow-y: auto;
    -webkit-overflow-scrolling: touch;
    max-height: calc(70vh - 140px); /* header + input ~= 140px */
  }

  .message-bubble {
    font-size: 14px;
    padding: 10px 14px;
    max-width: 78%;
    border-radius: 16px;
  }

  .typing-indicator {
    padding: 10px 12px;
  }

  /* Input area condensed */
  .chat-input-area {
    padding: 10px 12px;
    gap: 8px;
  }

  .chat-input {
    padding: 10px 12px;
    font-size: 14px;
    border-radius: 20px;
  }

  .send-btn {
    width: 44px;
    height: 44px;
  }

  /* ensure chat window is focusable and inputs visible when keyboard opens */
  .chat-window input:focus,
  .chat-window textarea:focus {
    scroll-margin-bottom: calc(env(safe-area-inset-bottom, 0px) + 120px);
  }

  /* When opened, slightly raise chat widget to avoid overlap with bottom nav on small screens */
  @media (max-height: 640px) {
    .chat-window {
      height: calc(78vh);
      max-height: calc(100vh - (env(safe-area-inset-top, 0px) + 120px));
    }
    .chat-messages {
      max-height: calc(78vh - 140px);
    }
  }
}
</style>