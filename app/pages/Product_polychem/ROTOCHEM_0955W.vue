<template>
  <div>
    <navbar />
    <div class="landing-page">
      <div class="main-container">
        <!-- Sidebar Menu - Sticky for all sections (pinned, starts at ADVANTAGE) -->
        <div class="sidebar" ref="sidebar" v-show="showSidebar">
          <div class="sidebar-title">ROTOCHEM 0955W</div>
          <div class="menu-container">
            <div class="timeline-bar" :style="{ top: timelinePosition + 'px' }"></div>
            <nav class="menu">
              <a 
                v-for="(item, index) in menuItems" 
                :key="index"
                :href="'#' + item.id"
                :ref="el => { if (el) menuItemEls[index] = el }"
                :class="['menu-item', { active: activeSection === item.id }]"
                @click="scrollToSection(item.id, $event)"
              >
                {{ item.label }}
                <span class="menu-number">{{ item.number }}</span>
              </a>
            </nav>
          </div>
        </div>

        <!-- Main Content Area -->
        <div class="main-content" @scroll="handleScroll" ref="mainContent">
          <!-- Team Awareness Section -->
          <section class="hero-content" ref="heroSection">
            <div class="hero-inner">
              <div class="hero-left">
                <div class="tag">Masterbatch</div>
                <h1 class="title">ROTOCHEM 0955W</h1>
                <p class="description text-justify">
                  Rotochem 0955W is a specialized white color plastic compound designed for rotational molding applications. It is a compound of linear medium density polyethylene copolymer grade with a narrow molecular weight distribution with Tio2 that possesses a range of beneficial characteristics. Rotochem is an ideal choice for manufacturers who require a reliable and durable material that can withstand harsh environmental conditions. Whether utilized in industrial or consumer applications, this plastic material is a suitable choice for those seeking consistent and stable materials.
                </p>
                <a
                  href="#"
                  class="btn-slide-down w-full lg:w-auto h-12 rounded-lg relative overflow-hidden border-2 border-[#FFCD05] text-[#FFCD05] transition-colors px-5 py-3 text-lg font-medium tracking-widest flex items-center justify-center"
                  @click.prevent="openPdf('/0966W EN.pdf')"
                >
                  <span>Download Data Sheet</span>
                </a>
              </div>

              <div class="hero-right">
                <div class="hero-right-image">
                  <img src="/955w-p.jpg" alt="Hero image" />
                </div>
              </div>
            </div>

            <!-- ADVANTAGE Details -->
            <div id="ADVANTAGE" class="details-section">
              <div class="details-content">
                <div class="details-left">
                <h2 class="section-title">ADVANTAGE</h2>
                <div class="card-container">
                  <div class="info-card purple">
                    <h3 class="card-title">Human Errors</h3>
                    <p class="card-description">
                      83% of data breaches involve human errors, underscoring the necessity of proper employee training.
                    </p>
                    <div class="card-source">Verizon 2024</div>
                  </div>
                </div>

                <div class="question-block">
                  <h3 class="question-title">What level of cybersecurity education do your employees have?</h3>
                  <p class="question-text">
                    Eighty percent of data leaks originate internally. A Gartner study revealed that 74% of employees would 
                    violate cybersecurity rules to meet business objectives or help team members. This internal risk highlights 
                    the need for continuous employee training and awareness.
                  </p>
                  <p class="question-text">
                    Vigilance decreases over time, so regular testing and reinforcement are essential. One-time training and 
                    unclear messaging lead to confusion and inconsistency in security practices.
                  </p>
                  <p class="question-text">
                    Companies should adopt an approach that includes engaging, interactive training, clear communication, 
                    and ongoing support. This ensures employees remain aware and aligned with cybersecurity protocols.
                  </p>
                </div>
              </div>

              <div class="details-right">
                <div class="sidebar-box">
                  <h4 class="sidebar-box-title">Common targets</h4>
                  <ul class="target-list">
                    <li>Small businesses</li>
                    <li>Medium businesses</li>
                    <li>Public administrations</li>
                    <li>Large enterprises</li>
                  </ul>
                </div>

                <div class="sidebar-box">
                  <h4 class="sidebar-box-title">Related Industries</h4>
                  <ul class="industry-list">
                    <li><span class="arrow">→</span> Financial Services</li>
                    <li><span class="arrow">→</span> Healthcare</li>
                    <li><span class="arrow">→</span> Retail and E-commerce</li>
                    <li><span class="arrow">→</span> Telecommunications</li>
                    <li><span class="arrow">→</span> Government and Public Sector</li>
                    <li><span class="arrow">→</span> Energy and Utilities</li>
                    <li><span class="arrow">→</span> Technology and Software</li>
                    <li><span class="arrow">→</span> Manufacturing</li>
                    <li><span class="arrow">→</span> Education</li>
                    <li><span class="arrow">→</span> Pharmaceuticals and Biotechnology</li>
                  </ul>
                </div>

                <div class="sidebar-box">
                  <h4 class="sidebar-box-title">Keywords</h4>
                  <div class="keyword-tags">
                    <span class="tag-item">Training</span>
                    <span class="tag-item">Sensitization</span>
                    <span class="tag-item">Phishing Awareness</span>
                    <span class="tag-item">Social engineering</span>
                    <span class="tag-item">Data protection</span>
                    <span class="tag-item">Malware Prevention</span>
                    <span class="tag-item">Password Management</span>
                  </div>
                </div>
                </div>
              </div>
            </div>
          </section>

          <!-- TECHNICAL_SPECIFICATIONS Section -->
          <section id="TECHNICAL_SPECIFICATIONS" class="content-block">
            <div class="content-inner">
              <h2 class="section-title">TECHNICAL SPECIFICATIONS</h2>
              <p class="section-text">Content for TECHNICAL_SPECIFICATIONS section...</p>
            </div>
          </section>

          <!-- PACKAGING Section -->
          <section id="PACKAGING" class="content-block">
            <div class="content-inner">
              <h2 class="section-title">PACKAGING</h2>
              <p class="section-text">Content for PACKAGING section...</p>
            </div>
          </section>

          <!-- STORAGE_AND_HANDLING Section -->
          <section id="STORAGE_AND_HANDLING" class="content-block">
            <div class="content-inner">
              <h2 class="section-title">STORAGE AND HANDLING</h2>
              <p class="section-text">Content for STORAGE_AND_HANDLING section...</p>
            </div>
          </section>

          <!-- Our Offers Section -->
          <section id="our-offers" class="content-block">
            <div class="content-inner">
              <h2 class="section-title">Our offers</h2>
              <p class="section-text">Content for Our offers section...</p>
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "~/components/navbar.vue"

export default {
  name: 'TeamAwarenessPage',
  components: {
    Navbar
  },
  data() {
    return {
      activeSection: 'ADVANTAGE',
      timelinePosition: 0,
      menuItems: [
        { id: 'ADVANTAGE', label: 'ADVANTAGE', number: '01' },
        { id: 'TECHNICAL_SPECIFICATIONS', label: 'TECHNICAL SPECIFICATIONS', number: '02' },
        { id: 'PACKAGING', label: 'PACKAGING', number: '03' },
        { id: 'STORAGE_AND_HANDLING', label: 'STORAGE AND HANDLING', number: '04' },
        { id: 'our-offers', label: 'Our offers', number: '05' }
      ]
      ,
      menuItemEls: [],
      showSidebar: false
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
      if (element && this.$refs.mainContent) {
        const container = this.$refs.mainContent
        const elementTop = element.offsetTop
        container.scrollTo({ top: elementTop, behavior: 'smooth' })
      }
    },
    handleScroll() {
      const main = this.$refs.mainContent
      const heroHeight = this.$refs.heroSection ? this.$refs.heroSection.offsetHeight : 0
      // show sidebar only after the hero (ADVANTAGE) is scrolled past
      this.showSidebar = main.scrollTop >= Math.max(0, heroHeight)

      const sections = this.menuItems.map(item => document.getElementById(item.id))
      const scrollPosition = main.scrollTop + 200

      for (let i = sections.length - 1; i >= 0; i--) {
        if (sections[i] && sections[i].offsetTop <= scrollPosition) {
          this.activeSection = this.menuItems[i].id
          this.updateTimeline()
          break
        }
      }
    },
    updateTimeline() {
      const activeIndex = this.menuItems.findIndex(item => item.id === this.activeSection)
      const el = this.menuItemEls[activeIndex]
      if (el && el.offsetTop != null) {
        const barHeight = 40 // matches .timeline-bar height in CSS
        this.timelinePosition = Math.round(el.offsetTop + el.offsetHeight / 2 - barHeight / 2)
      } else {
        this.timelinePosition = activeIndex * 52
      }
    }
    ,
    openPdf(url) {
      if (process.server || typeof window === 'undefined') return;
      if (!url) return;

      try {
        const fileUrl = encodeURI(url);
        const currentFav = document.querySelector('link[rel="icon"]')?.href || '/favicon.ico';
        const win = window.open('', '_blank');
        if (!win) {
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
  padding: 60px 80px;
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
  font-size: 12px;
  letter-spacing: 1px;
  margin-bottom: 30px;
  font-weight: 500;
}

.title {
  font-size: 72px;
  font-weight: 300;
  line-height: 1.1;
  margin-bottom: 40px;
  letter-spacing: -1px;
  color: #333333;
}

.description {
  color: #848484;
  font-size: 16px;
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

.card-title {
  font-size: 28px;
  font-weight: 500;
  margin-bottom: 20px;
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
.question-block {
  margin-top: 40px;
}

.question-title {
  font-size: 24px;
  font-weight: 400;
  margin-bottom: 24px;
  line-height: 1.4;
  color: #333333;
}

.question-text {
  color: #848484;
  font-size: 15px;
  line-height: 1.8;
  margin-bottom: 20px;
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

.content-inner {
  max-width: 1200px;
}

.section-text {
  color: #848484;
  font-size: 16px;
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
</style>
