<template>
  <div class="h-screen bg-[#f1f2f2] flex items-center justify-center bg-cover bg-center bg-blend-overlay">
    <div 
      ref="containerRef"
      :class="['w-full h-full flex overflow-x-auto overflow-y-hidden scroll-smooth bg-[#ffffff] shadow-2xl', {'detail-mode': activeProductIndex !== null}]"
    >
      <!-- Navigation Sidebar -->
      <nav class="w-[60px] lg:w-[100px] bg-[#848484] flex flex-col items-center justify-between py-7 px-4 text-white flex-shrink-0 relative z-10">
        <!-- back/home icon -->
        <div class="w-8 text-yellow-600">
          <NuxtLink to="/" aria-label="Go to home" class="block p-1">
            <!-- خانه (home) icon -->
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10.5L12 4l9 6.5V20a1 1 0 01-1 1h-5v-6H9v6H4a1 1 0 01-1-1V10.5z"/>
            </svg>
          </NuxtLink>
        </div>
        
        <div class="hidden lg:block [writing-mode:vertical-rl] -rotate-90">
          <NuxtLink to="/" aria-label="Go to home" class="block">
            <img src="/english logo W1.png" alt="Logo" class="w-28 h-auto object-contain">
          </NuxtLink>
        </div>
        
        <div class="relative flex items-center justify-center flex-col">
          <!-- لینک لاگین در پایین‌ترین بخش نوار کناری -->
          <NuxtLink to="/login" aria-label="Log in" class="block p-1 mt-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20v-1a4 4 0 00-4-4H9a4 4 0 00-4 4v1" />
            </svg>
          </NuxtLink>
        </div>
      </nav>

      <!-- Cover Image -->
      <div class="w-[300px] lg:w-[500px] flex items-center justify-center text-5xl text-white font-medium text-center flex-shrink-0 bg-black overflow-hidden">
        <video
          ref="leftVideo"
          class="cover-video w-full h-full"
          :src="leftVideoSrc"
          :poster="leftVideoPoster"
          autoplay
          muted
          loop
          playsinline
          preload="auto"
          @error="onVideoError"
          @loadeddata="onVideoLoaded"
          aria-hidden="true"
        >
          <source :src="leftVideoSrc" :type="leftVideoType" />
          <!-- fallback image if video not supported -->
          <img src="/955w-p.jpg" alt="" />
        </video>
      </div>

      <!-- Products -->
      <div
        v-for="(product, index) in products"
        :key="index"
        :data-index="index"
        @click="handleProductClick(index, $event)"
        :class="[
          'flex flex-col items-center p-9 bg-[#f1f2f2] w-[310px] overflow-y-auto scroll-smooth transition-all duration-500 flex-shrink-0 ml-1.5 relative',
          {
            'product-active w-full lg:w-[70%] px-8 lg:px-[70px] pb-0': activeProductIndex === index,
            'hover:shadow-[inset_0_-4px_0_0_#FFCD05] cursor-pointer': activeProductIndex !== index,
            'product-highlight': highlightedProductIndex === index
          }
        ]"
      >
        <!-- Close Button -->
        <button
          v-if="activeProductIndex === index"
          @click.stop="closeProduct"
          class="sticky top-0 ml-auto -mr-10 lg:-mr-10 w-9 h-9 rounded-full bg-black/50 flex items-center justify-center text-white z-10 flex-shrink-0 mb-4"
        >
          <svg class="w-6 h-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"/>
            <line x1="6" y1="6" x2="18" y2="18"/>
          </svg>
        </button>

        <!-- Navigation Buttons (shown when active) -->
        <div v-if="activeProductIndex === index" class="flex items-center justify-between w-full mb-6 gap-4 animate-fade-up sticky top-0 z-10 bg-[#f1f2f2] py-2" style="animation-delay: 0.1s">
          <button
            @click.stop="previousProduct"
            :disabled="activeProductIndex === 0"
            class="flex items-center justify-center w-12 h-12 rounded-full border-2 border-[#FFCD05] text-[#FFCD05] hover:bg-[#FFCD05] hover:text-black transition-all disabled:opacity-50 disabled:cursor-not-allowed"
            aria-label="Previous product"
          >
            <svg class="w-5 h-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="15 18 9 12 15 6"></polyline>
            </svg>
          </button>
          
          <div class="text-sm text-gray-600 font-medium min-w-[80px] text-center">
            {{ activeProductIndex + 1 }} / {{ products.length }}
          </div>
          
          <button
            @click.stop="nextProduct"
            :disabled="activeProductIndex === products.length - 1"
            class="flex items-center justify-center w-12 h-12 rounded-full border-2 border-[#FFCD05] text-[#FFCD05] hover:bg-[#FFCD05] hover:text-black transition-all disabled:opacity-50 disabled:cursor-not-allowed"
            aria-label="Next product"
          >
            <svg class="w-5 h-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="9 18 15 12 9 6"></polyline>
            </svg>
          </button>
        </div>

        <!-- Brand -->
        <div :class="['text-sm tracking-widest transition-all duration-500', activeProductIndex === index && 'text-base']">
          {{ product.brand }}
        </div>

        <!-- Title -->
        <h2 :class="['text-center font-medium my-5 transition-all duration-500', activeProductIndex === index ? 'text-3xl lg:text-[44px] max-w-none' : 'text-xl max-w-[12ch]']">
          {{ product.title }}
        </h2>

        <!-- Price -->
        <div v-if="product.price" :class="['mb-8 transition-all duration-500', activeProductIndex === index && 'text-2xl']">
          {{ product.price }}
        </div>

        <!-- Buttons (shown when active) -->
        <div v-if="activeProductIndex === index" class="flex items-center min-w-[80%] mb-9 animate-fade-up flex-col lg:flex-row" style="animation-delay: 0.2s">
          <a
            href="#"
            @click.prevent="openDataSheet(product.dataSheet)"
            class="btn-slide-down w-full justify-center lg:w-auto h-12 rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors px-5 py-3 text-lg font-medium tracking-widest flex-grow rounded ml-0 lg:ml-4 mt-4 lg:mt-0"
          >
            <span class="relative z-10">Download Data Sheet</span>
          </a>
        </div>

        <!-- Subtitle -->
        <p v-if="activeProductIndex !== index" class="text-gray-700 leading-relaxed text-sm mb-5 whitespace-pre-line">
          {{ product.subtitle }}
        </p>

        <!-- Expanded Content (shown when active) -->
        <template v-if="activeProductIndex === index">
          <div class="mb-5 animate-fade-up" style="animation-delay: 0.3s">
            <h4 v-if="product.subtitleTitle" class="text-sm font-semibold tracking-[0.3em] text-[#898989] uppercase mb-2">
              {{ product.subtitleTitle }}
            </h4>
            <p class="text-gray-700 leading-relaxed text-sm whitespace-pre-line">
              {{ product.subtitle }}
            </p>
          </div>
          
          <img 
            :src="product.detailImage || '/955w-p.jpg'"
            alt="Detail"
            class="block mt-5 -mx-8 lg:-mx-[70px] max-w-none w-[calc(100%+64px)] lg:w-[calc(100%+140px)]"
          />

          <!-- Technical Specifications Table -->
          <div class="flex flex-col bg-[#A8A8A8] text-white -mx-8 lg:-mx-[70px] px-8 lg:px-[100px] py-12 lg:py-[70px] w-[calc(100%+64px)] lg:w-[calc(100%+140px)]">
            <h3 class="font-semibold tracking-widest text-2xl mb-8 text-[#FFCD05]">TECHNICAL SPECIFICATIONS</h3>

            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 text-xs font-semibold tracking-[0.35em] uppercase text-[#FFCD05] pb-4 border-b border-white/30">
              <span>{{ product.propertyLabel || 'Properties' }}</span>
              <span>Test Method</span>
              <span>Unit</span>
              <span>Typical Value</span>
            </div>

            <div
              v-for="(spec, idx) in product.specs"
              :key="idx"
              class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 py-4 border-b border-white/10 text-sm"
            >
              <div class="font-semibold tracking-wide text-white/90">{{ spec.item || spec.property }}</div>
              <div class="text-white/80">{{ spec.testMethod }}</div>
              <div class="text-white/80">{{ spec.unit }}</div>
              <div class="text-white">{{ spec.typicalValue }}</div>
            </div>
          </div>

          <!-- Additional Narrative -->
          <div class="bg-[#f6f6f6] text-[#3d3d3d] -mx-8 lg:-mx-[70px] px-8 lg:px-[100px] py-10 w-[calc(100%+64px)] lg:w-[calc(100%+140px)] rounded-b-[20px]">
            <template v-if="product.detailSections && product.detailSections.length">
              <div
                v-for="(section, idx) in product.detailSections"
                :key="section.title || idx"
                class="space-y-2"
              >
                <h4 class="text-sm font-semibold tracking-[0.3em] text-[#898989] uppercase">
                  {{ section.title }}
                </h4>
                <p class="text-sm leading-relaxed whitespace-pre-line">
                  {{ section.body }}
                </p>
                <div class="h-6" v-if="idx !== product.detailSections.length - 1"></div>
              </div>
            </template>
            <template v-else>
              <p class="text-sm leading-relaxed whitespace-pre-line">
                {{ product.detailText }}
              </p>
            </template>
          </div>
        </template>

        <!-- View Details (shown when not active) -->
        <button v-else class="mt-auto font-medium text-sm tracking-widest text-[#FFCD05] hover:text-[#e6b800] transition-colors">VIEW DETAILS →</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, nextTick, onMounted, watch } from 'vue';

const route = useRoute();
const containerRef = ref(null);
const activeProductIndex = ref(null);
const leftVideo = ref(null);
const highlightedProductIndex = ref(null);

// encode path (handles spaces)
const leftVideoSrc = encodeURI('/Left side.mov');
const leftVideoPoster = '/955w-p.jpg';
const leftVideoType = leftVideoSrc.endsWith('.mp4')
  ? 'video/mp4'
  : leftVideoSrc.endsWith('.webm')
    ? 'video/webm'
    : 'video/quicktime';

const products = [
  {
    
    title: 'ROTOCHEM 0955W',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `Rotochem 0955W is a specialized white color plastic compound designed for rotational molding applications. It is a compound of linear medium density polyethylene copolymer grade with a narrow molecular weight distribution with Tio2 that possesses a range of beneficial characteristics. Rotochem is an ideal choice for manufacturers who require a reliable and durable material that can withstand harsh environmental conditions. Whether utilized in industrial or consumer applications, this plastic material is a suitable choice for those seeking consistent and stable materials.`,
    detailText: 'Rotochem 0955W delivers bright, uniform walls for large molded parts while maintaining impressive impact strength after long UV exposure. Customers rely on it for agricultural tanks, medical housings, and premium consumer bins where cosmetic consistency matters as much as durability.',
    detailImage: '/955w-p.jpg',
    detailSections: [
   
      {
        title: 'ADVANTAGE',
        body: `Rotochem exhibit a number of advantageous features, including excellent dispersion, exceptional impact strength, superior internal and external surface finish, optimal whiteness, reliable resistance to stress cracking, and efficient UV stabilization.`,
      },
      {
        title: 'PACKAGING',
        body: `Rotochem is supplied in powder form packed in 20 kg bags.`,
      },
      {
        title: 'STORAGE AND HANDLING',
        body: `Rotochem should be stored to prevent direct sunlight and/or heat exposure. The storage area should also be dry and preferably not exceed 50°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to process Rotochem within 18 months after delivery.`,
      },
    ],
  specs: [
      { property: 'MFI (190°C / 2.16 kg)', testMethod: 'ASTM D1238-7', unit: 'g/10min', typicalValue: '3.5 ± 0.5' },
      { property: 'Density', testMethod: 'ASTM D1505-68', unit: 'kg/m³', typicalValue: '0.937 – 0.945' },
      { property: 'Titanium crystal type', testMethod: '-', unit: '-', typicalValue: 'Rutile type' },
      { property: 'TiO₂ content', testMethod: '-', unit: '%', typicalValue: '2' },
      { property: 'Moisture content', testMethod: '-', unit: 'ppm', typicalValue: '≤ 1500' },
      { property: 'Vicat softening point', testMethod: 'ASTM D1525', unit: '°C', typicalValue: '115' },
      { property: 'Tensile strength @ yield', testMethod: 'ASTM D638 - 72', unit: 'MPa', typicalValue: '15' },
      { property: 'Elongation @ break', testMethod: 'ASTM D638', unit: '%', typicalValue: '> 700' },
      { property: 'Flexural modulus', testMethod: 'ASTM D790', unit: 'MPa', typicalValue: '650' },
      { property: 'Hardness', testMethod: 'ASTM D2240', unit: 'Shore D', typicalValue: '65' },
      { property: 'ESCR (IGEPAL 10%) (F50, 50°C)', testMethod: 'ASTM D1693', unit: 'HR', typicalValue: '200' }
    ],
    dataSheet: '/0966W EN.pdf',
  },
 {
    
    title: 'ROTOCHEM 0955B',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `Rotochem 0955B is a specialized blue color plastic compound designed for rotational molding applications. It is a compound of linear medium density polyethylene copolymer grade with a narrow molecular weight distribution with Tio2 that possesses a range of beneficial characteristics. Rotochem is an ideal choice for manufacturers who require a reliable and durable material that can withstand harsh environmental conditions. Whether utilized in industrial or consumer applications, this plastic material is a suitable choice for those seeking consistent and stable materials.`,
    detailText: 'Rotochem 0955W delivers bright, uniform walls for large molded parts while maintaining impressive impact strength after long UV exposure. Customers rely on it for agricultural tanks, medical housings, and premium consumer bins where cosmetic consistency matters as much as durability.',
    detailImage: '/955B.jpg',
    detailSections: [
      
      {
        title: 'ADVANTAGE',
        body: `Rotochem exhibit a number of advantageous features, including excellent dispersion, exceptional impact strength, superior internal and external surface finish, optimal whiteness, reliable resistance to stress cracking, and efficient UV stabilization.`,
      },
      {
        title: 'PACKAGING',
        body: `Rotochem is supplied in powder form packed in 20 kg bags.`,
      },
      {
        title: 'STORAGE AND HANDLING',
        body: `Rotochem should be stored to prevent direct sunlight and/or heat exposure. The storage area should also be dry and preferably not exceed 50°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to process Rotochem within 18 months after delivery.`,
      },
    ],
 specs: [
  { property: 'MFI (190°C / 5 kg)',      testMethod: 'ISO 1133',      unit: 'g/10min', typicalValue: '0.19 ± 0.1' },
  { property: 'Density',                testMethod: 'ISO 1183',      unit: 'kg/m³',   typicalValue: '0.952 ± 0.003' },
  { property: 'Moisture content',       testMethod: '-',             unit: 'ppm',     typicalValue: '≤ 1500' },
  { property: 'Melting point',          testMethod: 'ASTM D2117',    unit: '°C',      typicalValue: '135 ± 5' },
  { property: 'Vicat softening point',  testMethod: 'ASTM D1525',    unit: '°C',      typicalValue: '124' },
  { property: 'Tensile strength @ yield (MD)',       testMethod: 'ASTM D638', unit: 'MPa', typicalValue: '24' },
  { property: 'Tensile strength @ yield (MD, TD)',   testMethod: 'ASTM D638', unit: 'MPa', typicalValue: '30, 60' },
  { property: 'Elongation @ break (MD, TD)',         testMethod: 'ASTM D638', unit: '%',   typicalValue: '450, 240' },
  { property: 'ESCR (IGEPAL 10%) (F50, 50°C)',       testMethod: 'ASTM D1693', unit: 'HR',  typicalValue: '> 1000' }
],
    dataSheet: '/0966B EN.pdf',
  },
  {
 title: 'POLYFIL F700',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `POLYFIL F700 is a high-performance polyethylene compound specifically formulated for high-density polyethylene (HDPE) blown film applications. This grade is engineered to deliver superior mechanical properties, excellent film uniformity, and reliable processability, even in ultra-thin film applications.
It is highly recommended for producing films with thicknesses in the range of 10–25 microns, making it suitable for a wide range of packaging and consumer products such as shopping bags, T-shirt bags, garbage bags, liner bags, and food-contact films.`,
    detailImage: '/polyfilf700-2.jpg',
    detailSections: [
      
      {
        title: 'ADVANTAGE',
        body: `• High tensile strength with excellent dart impact resistance
• Low gel content for smooth, defect-free films
• Good moisture barrier and enhanced ultra-thin film capability
• Food contact compliant — suitable for hygienic and food packaging uses
• High stiffness ensuring good dimensional stability
• Wide service temperature range and UV resistance for outdoor durability
• Good impact resistance and excellent processability on standard HDPE film lines`,
      },
      {
        title: 'Processing Guidelines',
        body: `Processing parameters may vary depending on machine configuration, die size, and target film thickness. The following conditions are recommended as a starting point:
• Melt Temperature: 190–210 °C
• Blow-Up Ratio (BUR): 3–5
• Frost Line Height (FLH): 8–10 × die diameter`,
      },
       {
        title: 'PACKAGING',
        body: `POLYFIL F700 is supplied in powder form packed in 25 kg bags.`,
      },
      {
        title: 'STORAGE AND HANDLING',
        body: `POLYFIL F700 should be stored to prevent direct sunlight and/or heat exposure. The storage area should also be dry and preferably not exceed 50°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to process POLYFIL F700 within 12 months after delivery.`,
      },
    ],
  specs: [
      { property: 'MFI (190°C / 2.16 kg)', testMethod: 'ASTM D1238-7', unit: 'g/10min', typicalValue: '3.5 ± 0.5' },
      { property: 'Density', testMethod: 'ASTM D1505-68', unit: 'kg/m³', typicalValue: '0.937 – 0.945' },
      { property: 'Titanium crystal type', testMethod: '-', unit: '-', typicalValue: 'Rutile type' },
      { property: 'TiO₂ content', testMethod: '-', unit: '%', typicalValue: '2' },
      { property: 'Moisture content', testMethod: '-', unit: 'ppm', typicalValue: '≤ 1500' },
      { property: 'Vicat softening point', testMethod: 'ASTM D1525', unit: '°C', typicalValue: '115' },
      { property: 'Tensile strength @ yield', testMethod: 'ASTM D638 - 72', unit: 'MPa', typicalValue: '15' },
      { property: 'Elongation @ break', testMethod: 'ASTM D638', unit: '%', typicalValue: '> 700' },
      { property: 'Flexural modulus', testMethod: 'ASTM D790', unit: 'MPa', typicalValue: '650' },
      { property: 'Hardness', testMethod: 'ASTM D2240', unit: 'Shore D', typicalValue: '65' },
      { property: 'ESCR (IGEPAL 10%) (F50, 50°C)', testMethod: 'ASTM D1693', unit: 'HR', typicalValue: '200' }
    ],
    dataSheet: '/POLYFIL F700.pdf',
  },
    {
 title: 'HDCHEM 4760',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `HDCHEM 4760 is a specialized plastic compound designed for blow molding applications. It is a compound of polyethylene copolymers that possesses a range of beneficial characteristics. HDCHEM 4760 is an ideal choice for manufacturers who require a reliable and durable product, this plastic material is a dependable choice for those seeking consistent and stable materials.`,
    detailImage: '/Hdchem-2.jpg',
    detailSections: [
      
      {
        title: 'ADVANTAGE',
        body: `HDCHEM 4760 exhibit a number of advantageous features, including good flowability, impact strength, ESCR and rigidity.`,
      },
     
       {
        title: 'PACKAGING',
        body: `HDCHEM 4760 is supplied in powder form packed in 25 kg bags.`,
      },
      {
        title: 'STORAGE AND HANDLING',
        body: `HDCHEM 4760 should be stored to prevent direct sunlight and/or heat exposure. The storage area should also be dry and preferably not exceed 50°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to process HDCHEM 4760 within 12 months after delivery.`,
      },
    ],
specs: [
  { property: 'MFI (190°C / 5 kg)',     testMethod: 'ISO 1133',     unit: 'g/10min',   typicalValue: '1.2 ± 0.5' },
  { property: 'Density',                testMethod: 'ISO 1183',     unit: 'kg/m³',      typicalValue: '0.954 ± 0.002' },
  { property: 'Moisture content',       testMethod: '-',            unit: 'ppm',       typicalValue: '≤ 1500' },
  { property: 'Vicat softening point',  testMethod: 'ASTM D1525',   unit: '°C',        typicalValue: '115' },
  { property: 'Impact strength (23°C)', testMethod: 'ISO 179/1eA',  unit: 'mJ/mm²',    typicalValue: '9' },
  { property: 'Hardness',               testMethod: 'ASTM D2240',   unit: 'Shore D',   typicalValue: '62' },
  { property: 'Swell ratio',            testMethod: '-',            unit: '%',         typicalValue: '120' }
],
    dataSheet: '/BL3 COMPOUND(HDCHEM4760).pdf',
  },
 {
 title: 'SlIPCHEM E 178',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `SlIPCHEM-E 178 is a high-performance slip masterbatch containing high quality slip agent dispersed in a polyethylene carrier resin. It is specifically developed to reduce the coefficient of friction (COF) between polymer film layers during winding, bag-making, and packaging processes. The product offers excellent dispersion, high thermal stability, and consistent migration performance.`,
    detailImage: '/slipchem.jpg',
    detailSections: [
      
      {
        title: 'Typical Applications',
        body: `• Blown film (LDPE, LLDPE, HDPE)
• Cast film (CPP, PE-based multilayers)
• Laminated and printed packaging films
• Agricultural films and shrink/stretch films`,
      },
        {
        title: 'ADVANTAGE',
        body: `• Rapid and uniform slip activation
• Consistent COF reduction
• Excellent dispersion and film clarity retention
• No interference with sealing or printing inks
• Compatible with most pigments and additives`,
      },
      {
        title: 'TYPICAL ADDITION RATE',
        body: `• Rapid and uniform slip activation
• Consistent COF reduction
• Excellent dispersion and film clarity retention
• No interference with sealing or printing inks
• Compatible with most pigments and additives`,
      },
       {
        title: 'PACKAGING',
        body: `SlIPCHEM-E 178 is supplied in standard pellet form packed in 25 kg bags.`,
      },
      {
        title: 'STORAGE AND HANDLING',
        body: `SlIPCHEM-E 178 should be stored to prevent direct sunlight and heat exposure. The storage area should also be dry and preferably not exceed 40°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to pre-dry before use and process within 12 months under recommended conditions.`,
      },
    ],
specs: [
  { property: 'Appearance',                    testMethod: 'Visual',       unit: '-',                  typicalValue: 'Milky-white granular pellets' },
  { property: 'Active Content',                testMethod: '-',            unit: 'wt.%',               typicalValue: '28 ± 1' },
  { property: 'Carrier Resin',                 testMethod: '-',            unit: '–',                  typicalValue: 'LDPE / LLDPE compatible' },
  { property: 'Melt Flow Index (190°C / 2.16 kg)', testMethod: 'ASTM D1238', unit: 'g/10 min',         typicalValue: '15 ± 1' },
  { property: 'Density',                       testMethod: 'ASTM D792',    unit: 'g/cm³',              typicalValue: '0.94 – 0.96' },
  { property: 'Moisture Content',              testMethod: 'ASTM D6980',   unit: 'wt.%',               typicalValue: '< 0.2' }
],
    dataSheet: '/SlipChem-E 178.pdf',
  },
   {
 title: 'RAFCOLOR 1560',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `RAFCOLOR is the white masterbatch that consists of a high proportion of rutile titanium dioxide and thermoplastic polypropylene resin. The selected titanium dioxide has good opacity and dispersion performance. Highly-concentrated white MB with excellent dispersion and thermal stability can be applied to general-purpose products. It is recommended for raffia, Tapes, CF/BCF yarn, and other products.`,
    detailImage: '/Rafcolor-1.jpg',
    propertyLabel: 'Item',
    detailSections: [
      
      {
        title: 'Advantage',
        body: `RAFCOLOR is Featured with excellent dispersion, high compatibility, good processability, low water uptake, and stable properties.`,
      },
        {
        title: 'Packaging',
        body: `white masterbatch RAFCOLOR is supplied in standard pellet form packed in 25 kg bags.`,
      },
      {
        title: 'Storage and handling',
        body: `RAFCOLOR should be stored to prevent direct sunlight and/or heat exposure. The storage area should also be dry and preferably not exceed 50°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to process PP resin within 18 months after delivery.RAFCOLOR should be stored to prevent direct sunlight and/or heat exposure. The storage area should also be dry and preferably not exceed 50°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to process PP resin within 18 months after delivery.`,
      },
   
    ],
specs: [
  { item: 'Titanium crystal type',   testMethod: '-',           unit: '-',         typicalValue: 'Rutile type' },
  { item: 'Solid content',           testMethod: '-',           unit: '%',         typicalValue: '60 ± 2' },
  { item: 'Moisture content',        testMethod: 'ASTM D644',   unit: '%',         typicalValue: '> 0.2' },
  { item: 'Melt Flow Index (230°C / 2.16 kg)', testMethod: 'ASTM D1238', unit: 'g/10min', typicalValue: '13 ± 2' },
  { item: 'Density',                 testMethod: 'ASTM D1505',  unit: 'g/cm³',     typicalValue: '1.8' }
],
    dataSheet: '/',
  },
{
 title: 'CALCICHEM 126 FP',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `CALCICHEM 126 FP is a polypropylene-based filler masterbatch containing 80% calcium carbonate (CaCO₃) . The CALCICHEM 126 FP features a high, very fine, treated CaCO₃ content, ensuring excellent dispersion within the final product. It is specifically designed for direct addition during the processing of polyolefins, including extrusion and injection molding. The CALCICHEM 126 FP is manufactured using the premium additives and advanced production lines.`,
    detailImage: '/calcum-126.jpg',
    propertyLabel: 'Item',
    detailSections: [
      
      {
        title: 'Advantage',
        body: `▪ Increase of productivity
▪ Low water carrying
▪ Thermal conductivity improvement
▪ Increase of the anti-blocking effect
▪ Print improvement on the end product
▪ Stable performance of the product in the production process
▪ Ensuring high loading capability for raffia production
▪ Therelatively best price-to-performance ratio`,
      },
        {
        title: 'Packaging',
        body: `CALCICHEM 126 FP is supplied in standard pellet form packed in 25 kg bags.`,
      },
      {
        title: 'Storage and handling',
        body: `CALCICHEM 126 FP should be stored to prevent direct sunlight and heat exposure. The storage area should also be dry and preferably not exceed 50°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to pre-dry before use and process within 18 months after production date.`,
      },
   
    ],
specs: [
  { item: 'Carrier',                     testMethod: '-',           unit: '-',         typicalValue: 'Polypropylene' },
  { item: 'CaCO₃ content',               testMethod: 'ASTM D4218',  unit: '%',         typicalValue: '80 ± 2' },
  { item: 'Moisture content',            testMethod: 'ASTM D644',   unit: 'ppm',       typicalValue: '≤ 1500' },
  { item: 'Melt Index (230°C / 2.16 kg)', testMethod: 'ASTM D1238', unit: 'g/10min',   typicalValue: '2 – 8' },
  { item: 'Density @ 23°C',              testMethod: 'ASTM D1505',  unit: 'g/cm³',     typicalValue: '1.85 ± 0.05' }
],
    dataSheet: '/CALCICHEM 126 FP.pdf',
  },
  {
 title: 'CALCICHEM 110 FRF',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `CALCICHEM 110 FRF is a mineral modifier with a high, very fine, treated CaCO3 content that has an excellent dispersion in the final product, indicated for direct addition in the processing of polyolefins.
The CALCICHEM 110 FRF mineral modifier is designed for films, Raffia and ropes and also suitable for general-purpose products with PE and PP carriers.`,
    detailImage: '/110-FrF.webp',
    propertyLabel: 'Item',
    detailSections: [
      
      {
        title: 'Advantage',
        body: `
▪ Increase of productivity
▪ Raw material costs reduction
▪ Thermal conductivity improvement
▪ Power consumption Reduction
▪ Print improvement on the end product
▪ Increase of the anti-blocking effect`,
      },
        {
        title: 'Packaging',
        body: `CALCICHEM 110 FRF is supplied in standard pellet form packed in 25 kg bags.`,
      },
      {
        title: 'Storage and handling',
        body: `CALCICHEM 110 FRF should be stored to prevent direct sunlight and heat exposure. The storage area should also be dry and preferably not exceed 50°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to pre-dry before use and process within 18 months after production date.`,
      },
   
    ],
specs: [
  { property: 'Carrier',                     testMethod: '-', unit: '-',         typicalValue: 'Polypropylene' },
  { property: 'CaCO₃ content',               testMethod: '-', unit: '%',        typicalValue: '80 ± 2' },
  { property: 'Moisture content',            testMethod: '-', unit: 'ppm',      typicalValue: '≤ 1500' },
  { property: 'Melt Index (230°C / 5 kg)',   testMethod: '-', unit: 'g/10min',  typicalValue: '12 ± 2' },
  { property: 'Density @ 23°C',              testMethod: '-', unit: 'g/cm³',    typicalValue: '1.86 ± 0.05' }
],
    dataSheet: '/CALCICHEM 110 FRF.pdf',
  },{
 title: 'CALCICHEM 275 PM',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `CALCICHEM 275 PM is a polypropylene-based mineral masterbatch containing 75% ultra-fine mineral filler. It offers excellent dispersion and high mineral loading for enhanced performance in final products.
Specifically with formulated for direct addition during the extrusion of BOPP, CPP, and OPP films, CALCICHEM 275 PM delivers consistent performance and superior quality this product Manufactured with premium additives and produced on advanced production lines, it meets the highest industry standards for reliability and efficiency.`,
    detailImage: '/275.jpg',
    propertyLabel: 'Item',
    detailSections: [
      
      {
        title: 'Advantage',
        body: `
▪ Improved product consistency and quality
▪ Enhanced processing stability
▪ Suitable for high-performance film applications
`,
      },
        {
        title: 'Packaging',
        body: `CALCICHEM 275 PM is supplied in standard pellet form packed in 25 kg bags.`,
      },
      {
        title: 'Storage and handling',
        body: `CALCICHEM 275 PM should be stored to prevent direct sunlight and heat exposure. The storage area should also be dry and preferably not exceed 35°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to pre-dry before use and process within 18 months after production date.`,
      },
   
    ],
specs: [
  { 
  property: 'Carrier', 
  testMethod: '-', 
  unit: '-', 
  typicalValue: 'Polypropylene' 
},
{ 
  property: 'Mineral content', 
  testMethod: '-', 
  unit: '%', 
  typicalValue: '75 ± 1' 
},
{ 
  property: 'Mean particle size (d50%)', 
  testMethod: '-', 
  unit: 'μm', 
  typicalValue: '≤ 2' 
},
{ 
  property: 'Moisture content', 
  testMethod: '-', 
  unit: 'ppm', 
  typicalValue: '≤ 1500' 
},
{ 
  property: 'Melt Index (230°C / 2.16 kg)', 
  testMethod: '-', 
  unit: 'g/10min', 
  typicalValue: '3 ± 1' 
},
{ 
  property: 'Density @ 23°C', 
  testMethod: '-', 
  unit: 'g/cm³', 
  typicalValue: '1.75 ± 0.05' 
}
],
    dataSheet: '/CALCICHEM 275 PM.pdf',
  },
  
 
  {
    title: 'UVChem MB-R18',
    price: '',
    subtitleTitle: 'DESCRIPTION',
    subtitle: `UVChem MB-R18 is a high-performance UV stabilizer masterbatch, designed for use in raffia and woven fabric applications. It contains a synergistic blend of Hindered Amine Light Stabilizers (HALS) and UV absorbers optimized for long-term resistance against photo-oxidation, loss of mechanical properties, and color fading. The product provides superior dispersion, thermal stability, and process consistency in tape extrusion and drawing processes. The masterbatch is fully compatible with PP matrix and formulated to minimize processing discoloration and to provide long-term UV protection when dosed at 1 wt.% in the yarn formulation.`,
    detailImage: '/Uv-chem.jpg', 
    detailSections: [
      {
        title: 'Typical Applications',
        body: `• Polypropylene raffia tapes and woven sacks
• Agricultural and industrial packaging fabrics
• PP monofilaments and ropes
• Outdoor storage and construction fabrics`,
      },
      {
        title: 'ADVANTAGE',
        body: `• Excellent long-term UV resistance and color stability
• Maintains tensile and elongation properties during outdoor exposure
• Excellent compatibility with PP homopolymers (raffia grades)
• Low volatility and no interference with slip or antiblock additives
• Dust-free, free-flowing pellets for consistent feeding`,
      },
      {
        title: 'Typical Addition Rate',
        body: `Add 1.0 wt.% UVChem MB-R18 to the polypropylene base resin during extrusion or compounding to achieve optimum UV protection.
For outdoor or high-radiation applications, the dosage can be increased up to 1.5–2.0 wt.% depending on exposure conditions and product thickness. Ensure proper mixing using a loss-in-weight gravimetric feeder or equivalent dosing system for homogeneous distribution.`,
      },
      {
        title: 'PACKAGING',
        body: `UVChem MB-R18 is supplied in standard pellet form packed in 25 kg bags.`,
      },
      {
        title: 'STORAGE AND HANDLING',
        body: `UVChem MB-R18 should be stored to prevent direct sunlight and heat exposure. The storage area should also be dry and preferably not exceed 40°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to pre-dry before use and process within 12 months under recommended conditions.`,
      },
    ],
    specs: [
      { property: 'Appearance', testMethod: 'Visual', unit: '–', typicalValue: 'Off-white to light beige' },
      { property: 'Active Content', testMethod: '-', unit: 'wt.%', typicalValue: '20 ± 2' },
      { property: 'Melt Flow Index (230 °C / 2.16 kg)', testMethod: 'ASTM D1238', unit: 'g/10 min', typicalValue: '7 – 9' },
      { property: 'Density', testMethod: 'ASTM D792', unit: 'g/cm³', typicalValue: '0.91' },
      { property: 'Volatile Matter (2 h @ 220 °C)', testMethod: 'TGA', unit: 'wt.%', typicalValue: '< 1' },
      { property: 'Thermal Stability (5% weight loss)', testMethod: 'TGA', unit: '°C', typicalValue: '> 280' },
      { property: 'UV protection efficiency', testMethod: 'ASTM D4329', unit: '-', typicalValue: '> 90% retention of tensile strength after 500 h QUV' }
    ],
    dataSheet: '/UVChem MB-R18.pdf',
  },
];



const scrollConfig = {
  duration: 800,
  offsetFromCenter: 150,
  highlightDuration: 1200,
  enableHighlight: true
};

const productWidth = 310;
const gap = 6;

const getProductOffset = () => {
  if (!containerRef.value) return 0;
  return (containerRef.value.offsetWidth - (containerRef.value.offsetWidth * 70) / 100) / 2;
};

// 🔧 تابع کمکی برای تشخیص موبایل (فقط در کلاینت)
const isMobileDevice = () => {
  // بررسی محیط سرور
  if (process.server || typeof window === 'undefined') {
    return false; // در سرور فرض می‌کنیم دسکتاپ است
  }
  
  return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
    navigator.userAgent
  );
};

const smoothScrollTo = (targetScroll, duration = scrollConfig.duration) => {
  return new Promise((resolve) => {
    if (!containerRef.value || process.server) {
      resolve();
      return;
    }

    const startScroll = containerRef.value.scrollLeft;
    const distance = targetScroll - startScroll;
    const startTime = performance.now();

    const easeInOutCubic = (t) => {
      return t < 0.5 ? 4 * t * t * t : 1 - Math.pow(-2 * t + 2, 3) / 2;
    };

    const animate = (currentTime) => {
      const elapsed = currentTime - startTime;
      const progress = Math.min(elapsed / duration, 1);
      const easeProgress = easeInOutCubic(progress);

      containerRef.value.scrollLeft = startScroll + distance * easeProgress;

      if (progress < 1) {
        requestAnimationFrame(animate);
      } else {
        resolve();
      }
    };

    requestAnimationFrame(animate);
  });
};

const calculateCenterScrollPosition = (index) => {
  const navWidth = 60;
  const coverWidth = 300;
  const containerWidth = containerRef.value?.offsetWidth || 0;
  const productCenter = productWidth / 2;

  const productCenterX = navWidth + coverWidth + (index * (productWidth + gap)) + productCenter;
  const targetScrollLeft = productCenterX - containerWidth / 3 + scrollConfig.offsetFromCenter;

  return Math.max(0, targetScrollLeft);
};

const handleProductClick = async (index, event) => {
  if (activeProductIndex.value === index) return;

  activeProductIndex.value = index;
  highlightedProductIndex.value = null;

  await nextTick();

  // 🔧 استفاده از تابع کمکی
  if (isMobileDevice()) {
    event.currentTarget.scrollIntoView({ inline: 'center', behavior: 'smooth' });
  } else {
    const targetScroll = calculateCenterScrollPosition(index);
    await smoothScrollTo(targetScroll, scrollConfig.duration);
  }

  if (scrollConfig.enableHighlight) {
    highlightedProductIndex.value = index;
    setTimeout(() => {
      highlightedProductIndex.value = null;
    }, scrollConfig.highlightDuration);
  }
};

const closeProduct = () => {
  activeProductIndex.value = null;
  highlightedProductIndex.value = null;
};

const scrollToProduct = async (index) => {
  if (index === null || index === undefined || index < 0 || index >= products.length) {
    return;
  }

  // 🔧 بررسی محیط سرور
  if (process.server || typeof window === 'undefined') {
    return;
  }

  activeProductIndex.value = index;
  highlightedProductIndex.value = null;

  await nextTick();

  // 🔧 استفاده از تابع کمکی
  if (isMobileDevice()) {
    const productElement = containerRef.value?.querySelector(`[data-index="${index}"]`);
    if (productElement) {
      productElement.scrollIntoView({ inline: 'center', behavior: 'smooth' });
    }
  } else {
    const targetScroll = calculateCenterScrollPosition(index);
    await smoothScrollTo(targetScroll, scrollConfig.duration);
  }

  if (scrollConfig.enableHighlight) {
    highlightedProductIndex.value = index;
    setTimeout(() => {
      highlightedProductIndex.value = null;
    }, scrollConfig.highlightDuration);
  }
};

// Watch for route query changes
watch(() => route.query.index, async (newIndex) => {
  if (newIndex !== undefined) {
    const index = parseInt(newIndex);
    if (!isNaN(index)) {
      await nextTick();
      scrollToProduct(index);
    }
  }
}, { immediate: true });

// Handle initial load - فقط در کلاینت
onMounted(async () => {
  await nextTick();
  const indexParam = route.query.index;
  if (indexParam !== undefined) {
    const index = parseInt(indexParam);
    if (!isNaN(index)) {
      scrollToProduct(index);
    }
  }
});

const previousProduct = async () => {
  if (activeProductIndex.value > 0) {
    const newIndex = activeProductIndex.value - 1;
    await navigateToProduct(newIndex);
  }
};

const nextProduct = async () => {
  if (activeProductIndex.value < products.length - 1) {
    const newIndex = activeProductIndex.value + 1;
    await navigateToProduct(newIndex);
  }
};

const navigateToProduct = async (index) => {
  const router = useRouter();
  await router.replace({ query: { index: index.toString() } });
  await scrollToProduct(index);
};

// Video handling
const onVideoLoaded = (e) => {
  const v = leftVideo.value;
  if (!v) return;
  const p = v.play();
  if (p && typeof p.then === 'function') {
    p.catch(err => {
      // silent failure
    });
  }
};

const onVideoError = (e) => {
  // fallback behavior
};

onMounted(() => {
  nextTick().then(() => {
    if (leftVideo.value && leftVideo.value.readyState >= 2) {
      onVideoLoaded();
    }
  });
});

// اضافه کردن/جایگزینی تابع باز کردن تب جدید (اکنون با favicon دلخواه و iframe)
const openDataSheet = (url) => {
  if (process.server || typeof window === 'undefined') return;
  if (!url) return;

  try {
    const fileUrl = encodeURI(url);
    // تلاش برای استفاده از favicon فعلی صفحه، در غیر این صورت /favicon.ico
    const currentFav = document.querySelector('link[rel="icon"]')?.href || '/favicon.ico';

    // باز کردن پنجره خالی و نوشتن یک سند HTML که favicon را ست می‌کند و PDF را در iframe قرار می‌دهد
    const win = window.open('', '_blank');
    if (!win) {
      // در صورت عدم موفقیت در باز کردن پنجره جدید، fallback به باز کردن مستقیم فایل
      window.open(fileUrl, '_blank');
      return;
    }

    const doc = win.document;
    const html = `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>Data Sheet</title>
  <link rel="icon" href="${currentFav}">
  <style>html,body{height:100%;margin:0}iframe{border:0;width:100%;height:100vh}</style>
</head>
<body>
  <iframe src="${fileUrl}"></iframe>
</body>
</html>`;

    doc.open();
    doc.write(html);
    doc.close();
  } catch (err) {
    // fallback نهایی
    if (url) window.open(url, '_blank');
  }
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500;600&display=swap');

* {
  font-family: 'Montserrat', sans-serif;
}

body {
  font-family: 'Montserrat', sans-serif;
}

.detail-mode .product-active {
  opacity: 1;
}

.detail-mode > div:not(.product-active) {
  opacity: 0.3;
  transition: opacity 0.3s;
}

.detail-mode > div:not(.product-active):hover {
  opacity: 0.7;
}

@keyframes fade-up {
  from {
    opacity: 0;
    transform: translateY(100px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes product-highlight-pulse {
  0% {
    box-shadow: 0 0 0 0 rgba(255, 205, 5, 0.7);
  }
  50% {
    box-shadow: 0 0 0 15px rgba(255, 205, 5, 0);
  }
  100% {
    box-shadow: 0 0 0 0 rgba(255, 205, 5, 0);
  }
}

.animate-fade-up {
  animation: fade-up 0.6s both;
}

.product-highlight {
  animation: product-highlight-pulse 1.2s ease-out;
}

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

/* center the button text without changing the button size */
.btn-slide-down > span {
  display: block;
  width: 100%;
  text-align: center;
}

/* ensure the left cover video fills its container and keeps aspect via cover */
.cover-video {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  pointer-events: none; /* avoid intercepting clicks */
}

/* کاهش فاصله بین کلمات برای تمام عنوان‌ها */
h1, h2, h3, h4, h5, h6,
.tracking-widest {
  /* مقدار را در صورت نیاز کم/زیاد کنید؛ مقدار فعلی ملایم و خواناست */
  word-spacing: -0.10em;
}

/* در صورت نیاز می‌توان تنظیم خاص‌تری برای عنوان اصلی محصول اعمال کرد */
h2 {
  word-spacing: -0.08em;
}
</style>