<template>
  <div class="landing-page">
    <div class="main-container">
      <!-- Main Content Area -->
      <div class="main-content" @scroll="handleScroll" ref="mainContent">
        <!-- Team Awareness Section -->
        <section id="team-awareness" class="hero-content">
          <div class="hero-inner">
            <div class="hero-left">
              <div class="tag">UC_012</div>
              <h1 class="title">Team awareness</h1>
              <p class="description">
                Internal sources cause 80% of data leaks, with 74% of employees willing to bypass cybersecurity
                rules for business goals. To address this, companies need more than one-time training and unclear
                communication. Effective strategies include engaging, interactive training, clear messaging, and
                ongoing support to maintain employee vigilance, reduce risks, and ensure long-term cybersecurity
                compliance.
              </p>
              <button class="contact-btn">Contact us</button>
            </div>

            <div class="hero-right">
              <div class="tablet-mockup">
                <div class="screen">
                  <div class="screen-content">
                    <div class="notification-popup">
                      <div class="popup-header">
                        <span class="popup-icon">⚠️</span>
                        <span class="popup-text">Document</span>
                      </div>
                      <div class="popup-body">
                        <div class="alert-text">Emergency - Review Commitment</div>
                        <div class="detail-line">From Ahmed</div>
                        <div class="detail-line">To john@company.com</div>
                        <div class="detail-line warning-text">
                          Hello John, can you please review this document?
                        </div>
                        <div class="detail-line warning-text">
                          Our director is currently abroad and would
                        </div>
                        <div class="detail-line warning-text">Thank you</div>
                      </div>
                    </div>
                    <div class="malware-badge">Malware det</div>
                    <div class="document-badge">Document fraud</div>
                  </div>
                </div>
                <div class="expand-icon">⤢</div>
              </div>
            </div>
          </div>

          <!-- Key Figures Details -->
          <div id="key-figures" class="details-section">
            <!-- Left Sidebar Menu -->
            <div class="sidebar">
              <div class="sidebar-title">Summary</div>
              <div class="menu-container">
                <div class="timeline-bar" :style="{ top: timelinePosition + 'px' }"></div>
                <nav class="menu">
                  <a 
                    v-for="(item, index) in menuItems" 
                    :key="index"
                    :href="'#' + item.id"
                    :class="['menu-item', { active: activeSection === item.id }]"
                    @click="scrollToSection(item.id, $event)"
                  >
                    {{ item.label }}
                    <span class="menu-number">{{ item.number }}</span>
                  </a>
                </nav>
              </div>
            </div>

            <div class="details-content">
              <div class="details-left">
              <h2 class="section-title">Key figures</h2>
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

        <!-- Issue Section -->
        <section id="issue" class="content-block">
          <div class="content-inner">
            <h2 class="section-title">Issue</h2>
            <p class="section-text">Content for Issue section...</p>
          </div>
        </section>

        <!-- Our Answer Section -->
        <section id="our-answer" class="content-block">
          <div class="content-inner">
            <h2 class="section-title">Our answer</h2>
            <p class="section-text">Content for Our answer section...</p>
          </div>
        </section>

        <!-- Testimonial Section -->
        <section id="testimonial" class="content-block">
          <div class="content-inner">
            <h2 class="section-title">Testimonial</h2>
            <p class="section-text">Content for Testimonial section...</p>
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
</template>

<script>
export default {
  name: 'TeamAwarenessPage',
  data() {
    return {
      activeSection: 'key-figures',
      timelinePosition: 0,
      menuItems: [
        { id: 'key-figures', label: 'Key figures', number: '01' },
        { id: 'issue', label: 'Issue', number: '02' },
        { id: 'our-answer', label: 'Our answer', number: '03' },
        { id: 'testimonial', label: 'Testimonial', number: '04' },
        { id: 'our-offers', label: 'Our offers', number: '05' }
      ]
    }
  },
  mounted() {
    this.updateTimeline()
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
      const sections = this.menuItems.map(item => document.getElementById(item.id))
      const scrollPosition = this.$refs.mainContent.scrollTop + 200

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
      this.timelinePosition = activeIndex * 52
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
  background-color: #0a0a0a;
  min-height: 100vh;
  color: #ffffff;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

.main-container {
  display: flex;
  width: 100%;
  min-height: 100vh;
}

/* Main Content */
.main-content {
  flex: 1;
  overflow-y: auto;
  height: 100vh;
  width: 100%;
}

/* Hero Content */
.hero-content {
  border-bottom: 2px solid #444;
  padding-top: 0;
}

.hero-inner {
  display: flex;
  align-items: stretch;
  min-height: 100vh;
  border-bottom: 2px solid #444;
}

.hero-left {
  flex: 0 0 72%;
  padding: 60px 80px;
  border-right: 2px solid #444;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.hero-right {
  flex: 0 0 28%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px 40px;
}

.tag {
  color: #888;
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
}

.description {
  color: #b0b0b0;
  font-size: 16px;
  line-height: 1.7;
  margin-bottom: 40px;
  max-width: 650px;
}

.contact-btn {
  background-color: #ffffff;
  color: #000000;
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
  background-color: #f0f0f0;
  transform: translateY(-2px);
}

/* Tablet Mockup */
.tablet-mockup {
  position: relative;
  width: 100%;
  max-width: 400px;
  aspect-ratio: 4/3;
  background: linear-gradient(145deg, #1a1a1a, #0f0f0f);
  border: 3px solid #444;
  border-radius: 24px;
  padding: 20px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
}

.screen {
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #1a1a1a 0%, #0a0a0a 100%);
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
  background: rgba(255, 120, 0, 0.15);
  border: 2px solid #ff7800;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 20px;
  box-shadow: 0 8px 32px rgba(255, 120, 0, 0.3);
}

.popup-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 12px;
  font-size: 14px;
  color: #ff9933;
}

.popup-icon {
  font-size: 18px;
}

.popup-body {
  font-size: 11px;
  line-height: 1.6;
}

.alert-text {
  color: #ffaa66;
  margin-bottom: 8px;
  font-weight: 500;
}

.detail-line {
  color: #888;
  margin-bottom: 4px;
}

.warning-text {
  color: #999;
}

.malware-badge {
  display: inline-block;
  background: rgba(40, 40, 40, 0.8);
  border: 1px solid #555;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 11px;
  color: #ff6666;
  margin-right: 10px;
  margin-bottom: 10px;
}

.document-badge {
  display: inline-block;
  background: rgba(40, 40, 40, 0.8);
  border: 1px solid #555;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 11px;
  color: #ffaa66;
}

.expand-icon {
  position: absolute;
  bottom: 30px;
  right: 30px;
  width: 32px;
  height: 32px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid #444;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #888;
  font-size: 16px;
}

/* Details Section */
.details-section {
  display: flex;
  gap: 0;
  padding: 0;
  scroll-margin-top: 0;
}

/* Sidebar inside Key Figures */
.sidebar {
  flex: 0 0 250px;
  background-color: #0a0a0a;
  padding: 60px 40px;
  border-right: 2px solid #444;
  position: sticky;
  top: 0;
  height: 100vh;
  align-self: flex-start;
}

.sidebar-title {
  color: #888;
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
  background: linear-gradient(180deg, #7c3aed, #a78bfa);
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
  color: #666;
  text-decoration: none;
  font-size: 15px;
  padding: 10px 0;
  transition: all 0.3s ease;
  cursor: pointer;
}

.menu-item:hover,
.menu-item.active {
  color: #ffffff;
}

.menu-number {
  font-size: 12px;
  color: #444;
}

.menu-item.active .menu-number {
  color: #7c3aed;
}

.details-content {
  flex: 1;
  display: flex;
  gap: 60px;
  padding: 80px 60px;
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
  background: linear-gradient(135deg, #7c3aed, #a78bfa);
}

.card-title {
  font-size: 28px;
  font-weight: 500;
  margin-bottom: 20px;
}

.card-description {
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 40px;
}

.card-source {
  font-size: 13px;
  opacity: 0.8;
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
}

.question-text {
  color: #999;
  font-size: 15px;
  line-height: 1.8;
  margin-bottom: 20px;
}

/* Sidebar Boxes */
.sidebar-box {
  background-color: #111;
  padding: 24px;
  border-radius: 12px;
  border: 1px solid #222;
}

.sidebar-box-title {
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 20px;
  color: #888;
}

.target-list,
.industry-list {
  list-style: none;
}

.target-list li,
.industry-list li {
  color: #ccc;
  font-size: 14px;
  padding: 8px 0;
  border-bottom: 1px solid #1a1a1a;
}

.target-list li:last-child,
.industry-list li:last-child {
  border-bottom: none;
}

.arrow {
  margin-right: 8px;
  color: #666;
}

/* Keywords */
.keyword-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tag-item {
  background-color: #1a1a1a;
  border: 1px solid #333;
  padding: 6px 14px;
  border-radius: 6px;
  font-size: 12px;
  color: #aaa;
}

/* Other Content Blocks */
.content-block {
  min-height: 100vh;
  padding: 80px 60px;
  border-bottom: 2px solid #444;
}

.content-inner {
  max-width: 1200px;
}

.section-text {
  color: #888;
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
    border-bottom: 2px solid #444;
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
    border-bottom: 2px solid #444;
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