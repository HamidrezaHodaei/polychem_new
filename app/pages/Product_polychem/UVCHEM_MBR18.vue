<template>
  <div>
    <navbar />
    <div class="landing-page">
      <div class="main-container">
        <!-- Sidebar removed per request -->

        <!-- Main Content Area -->
        <div class="main-content" ref="mainContent">
          <!-- Team Awareness Section -->
          <section class="hero-content" ref="heroSection">
            <div class="hero-inner">
              <div class="hero-left ">
                <div class="tag  ">Masterbatch</div>
                <h1 class="title">UVCHEM MB-R18</h1>
                <p class="description text-justify">
                  UVCHEM MB-R18 is a high-performance UV stabilizer masterbatch, designed for use in raffia and woven fabric applications. It contains a synergistic blend of Hindered Amine Light Stabilizers (HALS) and UV absorbers optimized for long-term resistance against photo-oxidation, loss of mechanical properties, and color fading. The product provides superior dispersion, thermal stability, and process consistency in tape extrusion and drawing processes. The masterbatch is fully compatible with PP matrix and formulated to minimize processing discoloration and to provide long-term UV protection when dosed at 1 wt.% in the yarn formulation.
                </p>
                <a
                 href="#"
  class="btn-slide-down w-full lg:w-auto h-16 text-center rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors px-5 py-2 text-lg font-medium tracking-widest ml-0 lg:ml-4 mt-4 lg:mt-0 flex items-center justify-center"
  @click.prevent="openPdf('/UVChem MB-R18.pdf')"
>
  <span class="relative z-10">Download Data Sheet</span>
</a>
              </div>

              <div class="hero-right">
                <div class="hero-right-image">
                  <img src="/Uv-chem.webp" alt="Hero image" />
                </div>
              </div>
            </div>

            <!-- TECHNICAL_SPECIFICATIONS Section -->
<section id="TECHNICAL_SPECIFICATIONS" class="content-block">
  <div class="technical-specs">
    <h3 class="section-title-1 text-[#FFCD05]">TECHNICAL SPECIFICATIONS</h3>

    <div class="table-wrapper">
      <table class="specs-table">
        <thead>
          <tr>
            <th>Property</th>
            <th>Test Method</th>
            <th>Unit</th>
            <th>Typical Value</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(spec, idx) in specsData" :key="idx">
            <td>{{ spec.property }}</td>
            <td>{{ spec.testMethod }}</td>
            <td>{{ spec.unit }}</td>
            <td>{{ spec.typicalValue }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</section>


          <!-- ADVANTAGE Details -->
            <div id="ADVANTAGE" class="details-section">
              <div class="details-content">
                <div class="details-left">
                <h2 class="section-title">TYPICAL APPLICATIONS</h2>
                

                <div class="question-block">
                  
                  <ul class="advantage-list">
                    <li>Polypropylene raffia tapes and woven sacks</li>
                    <li>Agricultural and industrial packaging fabrics</li>
                    <li>PP monofilaments and ropes</li>
                    <li>Outdoor storage and construction fabrics</li>
                  </ul>
                </div>
              </div>

              <div class="details-right">
                
                <div class="sidebar-box">
                  <h4 class="sidebar-box-title">Related Industries</h4>
                  <ul class="industry-list">
                    <li><span class="arrow">→</span> Textiles & Fibers</li>
                    <li><span class="arrow">→</span> Raffia & Woven Fabrics</li>
                    <li><span class="arrow">→</span> Industrial Packaging</li>
                    <li><span class="arrow">→</span> Agricultural Applications</li>
                    <li><span class="arrow">→</span> Construction Materials</li>
                    <li><span class="arrow">→</span> Outdoor Products</li>
                    <li><span class="arrow">→</span> UV Protection Applications</li>
                      </ul>
                </div>

                
                </div>
              </div>
            </div>
          </section>


          <!-- ADVANTAGE Section -->
          <section id="ADVANTAGE_DETAIL" class="content-block-packaging">
            <div class="content-inner">
              <h2 class="section-title">ADVANTAGE</h2>
              <ul class="advantage-list">
                <li>Excellent long-term UV resistance and color stability</li>
                <li>Maintains tensile and elongation properties during outdoor exposure</li>
                <li>Excellent compatibility with PP homopolymers (raffia grades)</li>
                <li>Low volatility and no interference with slip or antiblock additives</li>
                <li>Dust-free, free-flowing pellets for consistent feeding</li>
              </ul>
            </div>
          </section>

          <!-- TYPICAL ADDITION RATE Section -->
          <section id="TYPICAL_ADDITION_RATE" class="content-block-packaging">
            <div class="content-inner">
              <h2 class="section-title">TYPICAL ADDITION RATE</h2>
              <p class="section-text-packaging">Add 1.0 wt.% UVChem MB-R18 to the polypropylene base resin during extrusion or compounding to achieve optimum UV protection.
For outdoor or high-radiation applications, the dosage can be increased up to 1.5–2.0 wt.% depending on exposure conditions and product thickness. Ensure proper mixing using a loss-in-weight gravimetric feeder or equivalent dosing system for homogeneous distribution.</p>
            </div>
          </section>

          <!-- PACKAGING Section -->
          <section id="PACKAGING" class="content-block-packaging">
            <div class="content-inner">
              <h2 class="section-title">PACKAGING</h2>
              <p class="section-text-packaging">UVChem MB-R18 is supplied in standard pellet form packed in 25 kg bags.</p>
            </div>
          </section>

          <!-- STORAGE_AND_HANDLING Section -->
          <section id="STORAGE_AND_HANDLING" class="content-block-STORAGE">
            <div class="content-inner">
              <h2 class="section-title">STORAGE AND HANDLING</h2>
              <p class="section-text-STORAGE">UVChem MB-R18 should be stored to prevent direct sunlight and heat exposure. The storage area should also be dry and preferably not exceed 40°C; Bad storage conditions may lead to quality deterioration and product performance. It is advisable to pre-dry before use and process within 12 months under recommended conditions.</p>
            </div>
          </section>

          <Footer1 />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "~/components/navbar.vue"
import Footer1 from "~/components/footer.vue"

export default {
  name: 'TeamAwarenessPage',
  components: {
    Navbar,
    Footer1
  },
  data() {
    return {
      activeSection: 'ADVANTAGE',
      timelinePosition: 0,
      menuItems: [
        { id: 'applications', label: 'Applications', number: '01' },
        { id: 'advantage', label: 'Advantages', number: '02' },
        { id: 'addition-rate', label: 'Addition Rate', number: '03' },
        { id: 'packaging', label: 'Packaging', number: '04' },
        { id: 'storage', label: 'Storage', number: '05' }
      ],
      menuItemEls: [],
      showSidebar: false,
      specsData: [
        { property: 'Appearance', testMethod: 'Visual', unit: '–', typicalValue: 'Off-white to light beige' },
        { property: 'Active Content', testMethod: '-', unit: 'wt.%', typicalValue: '20 ± 2' },
        { property: 'Melt Flow Index (230 °C / 2.16 kg)', testMethod: 'ASTM D1238', unit: 'g/10 min', typicalValue: '7 – 9' },
        { property: 'Density', testMethod: 'ASTM D792', unit: 'g/cm³', typicalValue: '0.91' },
        { property: 'Volatile Matter (2 h @ 220 °C)', testMethod: 'TGA', unit: 'wt.%', typicalValue: '< 1' },
        { property: 'Thermal Stability (5% weight loss)', testMethod: 'TGA', unit: '°C', typicalValue: '> 280' },
        { property: 'UV protection efficiency', testMethod: 'ASTM D4329', unit: '-', typicalValue: '> 90% retention of tensile strength after 500 h QUV' }
      ]
    }
  },
  mounted() {
    // compute initial sidebar visibility and timeline after render
    this.$nextTick(() => {
      const hero = this.$refs.heroSection
      const main = this.$refs.mainContent
      const heroHeight = hero ? hero.offsetHeight : 0
      const scrollTop = main ? main.scrollTop : 0
      // show sidebar only after user scrolls past hero (ADVANTAGE)
      this.showSidebar = scrollTop >= Math.max(0, heroHeight)
      this.updateTimeline()
    })
  },
  methods: {
    scrollToSection(id, event) {
      event.preventDefault()
      const element = document.getElementById(id)
    },
    handleScroll() {
    },
    updateTimeline() {}
    ,
    openPdf(url) {
      try {
        const currentFav = '/english logo W1.png'
        const fileUrl = url
        const doc = window.open()
        const html = `
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
        if (url) window.open(url, '_blank');
      }
    }
  }
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.landing-page {
  background-color: #f1f2f2;
  height: 100vh; /* fix to viewport so outer page doesn't scroll */
  overflow: hidden; /* prevent outer scrollbar */
  color: #848484;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

.main-container {
  display: flex;
  width: 100%;
  height: 100%; /* fill the landing-page height */
  background-color: #ffffff;
}

/* Main Content */
.main-content {
  flex: 1;
  overflow-y: auto; /* only this element scrolls */
  height: 100%;
  width: 100%;
  background-color: #ffffff;
}

/* Sidebar - Global */
.sidebar {
  flex: 0 0 250px;
  background-color: #f1f2f2;
  padding: 60px 40px;
  border-right: 1px solid #e0e0e0;
  position: sticky;
  top: 0;
  height: 100vh;
  align-self: flex-start;
}

/* Hero Content */
.hero-content {
  border-bottom: 1px solid #e0e0e0;
  padding-top: 0;
}

.hero-inner {
  display: flex;
  align-items: stretch;
  min-height: 100vh;
  border-bottom: 1px solid #e0e0e0;
  background-color: #ffffff;
}

.hero-left {
  flex: 0 0 40%;
  padding: 10px 80px 60px 80px;
  border-right: 1px solid #e0e0e0;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

/* Make any image placed inside the left column fill the area */
.hero-left {
  position: relative;
  overflow: hidden;
}

.hero-left img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.hero-right {
  flex: 0 0 60%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px 40px;
  background-color: #f1f2f2;
}

/* Full-bleed image for right column */
.hero-right {
  padding: 0;
  background-color: transparent;
}

.hero-right-image {
  width: 100%;
  height: 100%;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
}

.hero-right-image img {
  height: 100%;
  width: auto; /* keep aspect ratio and reduce visible width */
  display: block;
  object-fit: cover;
  object-position: center center;
}

.tag {
  color: #848484;
  font-size: 14px;
  letter-spacing: 1px;
  margin-bottom: 30px;
  font-weight: 500;
  
}

.title {
  font-size: 60px;
  font-weight: 400; 
  line-height: 1.1;
  margin-bottom: 40px;
  letter-spacing: -1px;
  color: #333333;
}
.description {
  color: #848484;
  font-size: 18px;
  line-height: 1.7;
  margin-bottom: 40px;
  max-width: 650px;
}

.contact-btn {
  background-color: #848484;
  color: #ffffff;
  border: none;
  padding: 16px 32px;
  font-size: 15px;
  font-weight: 500;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  width: fit-content;
}

.contact-btn:hover {
  background-color: #696969;
  transform: translateY(-2px);
}

/* Tablet Mockup */
.tablet-mockup {
  position: relative;
  width: 100%;
  max-width: 400px;
  aspect-ratio: 4/3;
  background: linear-gradient(145deg, #f9f9f9, #f1f2f2);
  border: 1px solid #e0e0e0;
  border-radius: 24px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.screen {
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #f9f9f9 0%, #f1f2f2 100%);
  border-radius: 12px;
  overflow: hidden;
  position: relative;
}

.screen-content {
  position: relative;
  padding: 30px;
  height: 100%;
}

.notification-popup {
  background: rgba(255, 120, 0, 0.08);
  border: 1px solid #ffb366;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 20px;
  box-shadow: 0 2px 8px rgba(255, 120, 0, 0.1);
}

.popup-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 12px;
  font-size: 14px;
  color: #ff7800;
}

.popup-icon {
  font-size: 18px;
}

.popup-body {
  font-size: 11px;
  line-height: 1.6;
  color: #848484;
}

.alert-text {
  color: #ff7800;
  margin-bottom: 8px;
  font-weight: 500;
}

.detail-line {
  color: #999;
  margin-bottom: 4px;
}

.warning-text {
  color: #999;
}

.malware-badge {
  display: inline-block;
  background: #f5f5f5;
  border: 1px solid #e0e0e0;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 11px;
  color: #ff7800;
  margin-right: 10px;
  margin-bottom: 10px;
}

.document-badge {
  display: inline-block;
  background: #f5f5f5;
  border: 1px solid #e0e0e0;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 11px;
  color: #ff7800;
}

.expand-icon {
  position: absolute;
  bottom: 30px;
  right: 30px;
  width: 32px;
  height: 32px;
  background: #f5f5f5;
  border: 1px solid #e0e0e0;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #999;
  font-size: 16px;
}

/* Details Section */
.details-section {
  display: flex;
  gap: 0;
  padding: 0;
  scroll-margin-top: 0;
}

.sidebar-title {
  color: #848484;
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 40px;
  letter-spacing: 0.5px;
}

.menu-container {
  position: relative;
}

.timeline-bar {
  position: absolute;
  left: 0;
  width: 2px;
  height: 40px;
  background: linear-gradient(180deg, #FFCD05, #FFD700);
  transition: top 0.3s ease;
  border-radius: 2px;
}

.menu {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding-left: 20px;
}

.menu-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #b0b0b0;
  text-decoration: none;
  font-size: 15px;
  padding: 10px 0;
  transition: all 0.3s ease;
  cursor: pointer;
}

.menu-item:hover,
.menu-item.active {
  color: #848484;
}

.menu-number {
  font-size: 12px;
  color: #d0d0d0;
}

.menu-item.active .menu-number {
  color: #FFCD05;
}

.details-content {
  flex: 1;
  display: flex;
  gap: 60px;
  padding: 80px 60px;
  background-color: #ffffff;
}

.details-left {
  flex: 1;
}

.details-right {
  flex: 0 0 350px;
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.section-title {
  font-size: 48px;
  font-weight: 300;
  margin-bottom: 40px;
  letter-spacing: -0.5px;
  color: #333333;
}
.section-title-1 {
  font-size: 48px;
  font-weight: 300;
  margin-bottom: 40px;
  letter-spacing: -0.5px;
  color: #FFCD05;
  padding-left: 40px;
  padding-top: 40px;
}
/* Info Card */
.card-container {
  margin-bottom: 60px;
}

.info-card {
  padding: 40px;
  border-radius: 16px;
  position: relative;
}

.info-card.purple {
  background: linear-gradient(135deg, #FFCD05, #FFD700);
  color: #333333;
}


.card-description {
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 40px;
  color: #333333;
  
}

.card-source {
  font-size: 13px;
  opacity: 0.8;
  color: #333333;
}

/* Question Block */



.question-text {
  color: #848484;
  font-size: 20px;
  line-height: 1.8;
  margin-bottom: 20px;
}

.advantage-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.advantage-list li {
  color: #848484;
  font-size: 18px;
  line-height: 1.8;
  margin-bottom: 12px;
  padding-left: 24px;
  position: relative;
}

.advantage-list li:before {
  content: '→';
  position: absolute;
  left: 0;
  color: #FFCD05;
  font-weight: bold;
}

/* Sidebar Boxes */
.sidebar-box {
  background-color: #f9f9f9;
  padding: 24px;
  border-radius: 12px;
  border: 1px solid #e0e0e0;
}

.sidebar-box-title {
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 20px;
  color: #848484;
}

.target-list,
.industry-list {
  list-style: none;
}

.target-list li,
.industry-list li {
  color: #848484;
  font-size: 14px;
  padding: 8px 0;
  border-bottom: 1px solid #e0e0e0;
}

.target-list li:last-child,
.industry-list li:last-child {
  border-bottom: none;
}

.arrow {
  margin-right: 8px;
  color: #ccc;
}

/* Keywords */
.keyword-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.technical-specs {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 80px 20px;
  background-color: #A8A8A8;
  color: #fff;
  box-sizing: border-box;
}
#TECHNICAL_SPECIFICATIONS {
  padding: 0;
  margin: 0;
  min-height: auto;
  background-color: #A8A8A8;
}

.technical-specs {
  width: 100%;
  max-width: 100%;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.table-wrapper {
  width: 100%;
  overflow-x: auto;
  margin: 0;
  padding: 0;
}

.specs-table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
  margin: 0;
  padding: 0;
}

.specs-table th,
.specs-table td {
  padding: 16px;
  text-align: left;
  border-bottom: 1px solid rgba(255,255,255,0.3);
}

.specs-table th {
  color: #FFCD05;
  font-size: 14px;
  letter-spacing: 0.1em;
  text-transform: uppercase;
}

.specs-table td {
  color: #fff;
  font-size: 14px;
  line-height: 1.6;
  word-wrap: break-word;
}

.tag-item {
  background-color: #f5f5f5;
  border: 1px solid #e0e0e0;
  padding: 6px 14px;
  border-radius: 6px;
  font-size: 12px;
  color: #848484;
}

/* Other Content Blocks */
.content-block {
  min-height: 100vh;
  padding: 80px 60px;
  border-bottom: 1px solid #e0e0e0;
  background-color: #ffffff;
}
.content-block-packaging {
  min-height: 10vh;
  padding: 80px 60px;
  border-bottom: 1px solid #e0e0e0;
  background-color: #ffffff;
}
.content-block-STORAGE {
  min-height: 60vh;
  padding: 80px 60px;
  border-bottom: 1px solid #e0e0e0;
  background-color: #ffffff;
}
.content-inner {
  max-width: 1200px;
}

.section-text {
  color: #848484;
  font-size: 16px;
  line-height: 1.8;
}
.section-text-packaging{
    color: #848484;
  font-size: 20px;
  line-height: 1.8;
}
.section-text-STORAGE{
  color: #848484;
  font-size: 18px;
  line-height: 1.8;
}
@media (max-width: 1024px) {
  .main-container {
    flex-direction: column;
  }

  .hero-inner {
    flex-direction: column;
  }

  .hero-left {
    flex: 1 1 100%;
    border-right: none;
    border-bottom: 1px solid #e0e0e0;
    padding: 40px;
  }

  .hero-right {
    flex: 1 1 100%;
    padding: 40px;
  }

  .title {
    font-size: 56px;
  }

  .details-section {
    flex-direction: column;
  }

  .sidebar {
    position: relative;
    height: auto;
    border-right: none;
    border-bottom: 1px solid #e0e0e0;
  }

  .details-content {
    flex-direction: column;
  }

  .details-right {
    flex: 1 1 100%;
  }
}

@media (max-width: 768px) {
  .title {
    font-size: 42px;
  }

  .description {
    font-size: 14px;
  }

  .section-title {
    font-size: 36px;
  }
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

.btn-slide-down > span {
  display: block;
  width: 100%;
  text-align: center;
}
</style>
