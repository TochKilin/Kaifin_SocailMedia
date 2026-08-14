<template>
  <div class="app-wrapper">
    <NavBar />
    
    <!-- Outer Container -->
    <div class="book-container">
      <main class="about-us-page">
        
        <!-- ==================== PAGE 1: LEFT PAGE ==================== -->
        <div class="book-page page-left">
          
          <!-- 1. Header Section -->
          <header class="header-section">
            <div class="logo-box">
              <img src="../../../assets/logos/kaifin_l2.png" alt="KaiFinn Logo" class="logo-img" />
            </div>
            
            <div class="title-badge">
              <span class="badge-dot"></span>
              About KaiFinn
            </div>
            
            <p class="description-box">
              Empowering learners and instructors worldwide with modern, high-quality, and accessible digital education tools.
            </p>
          </header>

          <!-- 2. Main Visual Banner -->
          <div class="feature-banner">
            <img 
              src="https://images.unsplash.com/photo-1577896851231-70ef18881754?auto=format&fit=crop&q=80&w=700" 
              alt="Classroom" 
              class="banner-img"
            />
            <div class="banner-overlay">
              <div class="live-tag">
                <span class="pulse-dot"></span> Inspiring Spaces
              </div>
            </div>
          </div>

          <!-- 3. Key Metrics Grid -->
          <section class="stats-grid">
            <div class="stat-card">
              <div class="stat-icon-bg blue">
                <img src="../../../assets/lotties_svg/course.svg" alt="Courses" class="lottie-svg-icon" />
              </div>
              <div class="stat-info">
                <span class="stat-number">1,000+</span>
                <span class="stat-label">My Courses</span>
              </div>
            </div>

            <div class="stat-card">
              <div class="stat-icon-bg green">
                <img src="../../../assets/lotties_svg/Student.svg" alt="Active Students" class="lottie-svg-icon" />
              </div>
              <div class="stat-info">
                <span class="stat-number">100K+</span>
                <span class="stat-label">Active Students</span>
              </div>
            </div>

            <div class="stat-card">
              <div class="stat-icon-bg purple">
                <img src="../../../assets/lotties_svg/Learning.svg" alt="Global Learners" class="lottie-svg-icon" />
              </div>
              <div class="stat-info">
                <span class="stat-number">20M+</span>
                <span class="stat-label">Global Learners</span>
              </div>
            </div>

            <div class="stat-card">
              <div class="stat-icon-bg orange">
                <img src="../../../assets/lotties_svg/Teacher.svg" alt="Instructors" class="lottie-svg-icon" />
              </div>
              <div class="stat-info">
                <span class="stat-number">1,000+</span>
                <span class="stat-label">Instructors</span>
              </div>
            </div>
          </section>

          <!-- 4. Contact & Social Section -->
          <footer class="footer-section">
            <div class="section-header">
              <span class="section-subtitle">CONNECT WITH US</span>
              <h2 class="section-title">Get In Touch</h2>
            </div>
            
            <div class="contact-cards-stack">
              <!-- Location with SVG Icon -->
              <a href="https://maps.google.com" target="_blank" class="contact-item">
                <div class="contact-icon">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
                    <circle cx="12" cy="10" r="3"></circle>
                  </svg>
                </div>
                <div class="contact-text">
                  <span class="contact-label">Location</span>
                  <span class="contact-val">#1,3,5A, St.1, Borey Piphup Thmey, Phnom Penh</span>
                </div>
              </a>

              <!-- Phone Call with SVG Icon -->
              <a href="tel:0967385255" class="contact-item">
                <div class="contact-icon">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path>
                  </svg>
                </div>
                <div class="contact-text">
                  <span class="contact-label">Phone Call</span>
                  <span class="contact-val">+855 96 738 5255</span>
                </div>
              </a>
            </div>

            <!-- Social Links -->
            <div class="social-bar">
              <span class="social-label">Follow our community</span>
              <div class="social-buttons">
                <a href="#" class="social-btn facebook" aria-label="Facebook">
                  <svg viewBox="0 0 24 24" fill="currentColor"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
                </a>
                <a href="#" class="social-btn telegram" aria-label="Telegram">
                  <svg viewBox="0 0 24 24" fill="currentColor"><path d="M12 0C5.37 0 0 5.37 0 12s5.37 12 12 12 12-5.37 12-12S18.63 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.226-.461-1.901-.903-1.056-.692-1.653-1.123-2.678-1.799-1.185-.781-.417-1.21.258-1.911.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.48.33-.913.49-1.302.48-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.831-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141.119.098.152.228.166.323.014.098.028.324.016.498z"/></svg>
                </a>
              </div>
            </div>
          </footer>

        </div>

        <!-- ==================== PAGE 2: RIGHT PAGE ==================== -->
        <div class="book-page page-right">
          
          <!-- 1. Team Section -->
          <section class="team-section">
            <div class="section-header">
              <span class="section-subtitle">OUR TALENTS</span>
              <h2 class="section-title">Team Development</h2>
            </div>
            
            <div class="stacked-cards-container">
              <div 
                v-for="(member, index) in teamMembers" 
                :key="index"
                class="team-review-card"
                :class="[
                  index % 2 === 0 ? 'blue-card' : 'white-card',
                  getRotationClass(index)
                ]"
              >
                <!-- Sparkle for Blue Cards -->
                <template v-if="index % 2 === 0">
                  <span class="sparkle top-left-sparkle">✦</span>
                  <span class="sparkle bottom-right-sparkle">✦</span>
                </template>
                
                <div class="card-avatar">
                  <img :src="member.image" :alt="member.name" />
                </div>
                
                <div class="card-content">
                  <div class="card-header-row">
                    <h3 class="member-name">{{ member.name }}</h3>
                  </div>
                  <p class="member-desc">{{ member.role }}</p>
                </div>
              </div>
            </div>
          </section>

        </div>

      </main>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import NavBar from '../navbar/NavBar.vue';

const teamMembers = ref([
  {
    name: 'Sombat',
    role: 'Lead Full-Stack Developer specializing in high-performance web systems and seamless architecture.',
    image: new URL('../../../assets/team_images/admin.jpg', import.meta.url).href
  },
  {
    name: 'Rumduol',
    role: 'Project Leader driving agile product design, client strategies, and smooth project executions.',
    image: new URL('../../../assets/team_images/phea.jpg', import.meta.url).href
  },
  {
    name: 'Nita',
    role: 'UX/UI Designer passionate about human-centered interfaces, aesthetic visuals, and pixel-perfect UI.',
    image: new URL('../../../assets/team_images/yuri.jpg', import.meta.url).href
  },
  {
    name: 'Sophea',
    role: 'Backend Engineer focusing on API optimization, microservices, and secure database design.',
    image: new URL('../../../assets/team_images/pa.jpg', import.meta.url).href
  },
  {
    name: 'Yuri',
    role: 'Frontend Developer creating responsive, interactive, and modern web applications.',
    image: new URL('../../../assets/team_images/yuri.jpg', import.meta.url).href
  },
  {
    name: 'Mean',
    role: 'Frontend Developer creating responsive, interactive, and modern web applications.',
    image: new URL('../../../assets/team_images/mean.jpg', import.meta.url).href
  },
  {
    name: 'Phea',
    role: 'Frontend Developer creating responsive, interactive, and modern web applications.',
    image: new URL('../../../assets/team_images/phea.jpg', import.meta.url).href
  },
  {
    name: 'Roth',
    role: 'Frontend Developer creating responsive, interactive, and modern web applications.',
    image: new URL('../../../assets/team_images/roth.jpg', import.meta.url).href
  },
  {
    name: 'Ly Heng',
    role: 'Frontend Developer creating responsive, interactive, and modern web applications.',
    image: new URL('../../../assets/team_images/lyheng.jpg', import.meta.url).href
  },
  {
    name: 'Nary',
    role: 'Frontend Developer creating responsive, interactive, and modern web applications.',
    image: new URL('../../../assets/team_images/nary.jpg', import.meta.url).href
  },
  {
    name: 'Heng Heng',
    role: 'Frontend Developer creating responsive, interactive, and modern web applications.',
    image: new URL('../../../assets/team_images/heng.jpg', import.meta.url).href
  },
  {
    name: 'Nita',
    role: 'Frontend Developer creating responsive, interactive, and modern web applications.',
    image: new URL('../../../assets/team_images/nita.jpg', import.meta.url).href
  }
]);

const getRotationClass = (index) => {
  const classes = ['rotate-left', 'rotate-right', 'rotate-left-alt'];
  return classes[index % classes.length];
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap');

.app-wrapper {
  background-color: #F7F4F2;
  height: 100vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  font-family: 'Plus Jakarta Sans', system-ui, -apple-system, sans-serif;
}

.book-container {
  flex: 1;
  height: calc(100vh - 60px);
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* ================= MAIN BOOK CONTAINER ================= */
.about-us-page {
  display: grid;
  grid-template-columns: 1fr 1fr;
  width: 100%;
  max-width: 1150px;
  height: 100%;
  background-color: #ffffff;
  color: #0f172a;
  position: relative;
  overflow: hidden;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.08);
}

.about-us-page::after {
  content: '';
  position: absolute;
  top: 0;
  bottom: 0;
  left: 50%;
  width: 30px;
  transform: translateX(-50%);
  background: linear-gradient(
    to right,
    rgba(0,0,0,0) 0%,
    rgba(0,0,0,0.03) 40%,
    rgba(0,0,0,0.12) 50%,
    rgba(0,0,0,0.03) 60%,
    rgba(0,0,0,0) 100%
  );
  pointer-events: none;
  z-index: 20;
}

.book-page {
  padding: 24px 40px;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  gap: 20px;
  box-sizing: border-box;
  height: 100%;
  overflow-y: auto;
}

.book-page::-webkit-scrollbar {
  width: 4px;
}
.book-page::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 4px;
}

.page-left {
  border-right: 1px solid #f1f5f9;
  background: linear-gradient(135deg, #ffffff 80%, #f8fafc 100%);
}

.page-right {
  border-left: 1px solid #f1f5f9;
  background: linear-gradient(225deg, #ffffff 80%, #f8fafc 100%);
}

/* ================= HEADER SECTION ================= */
.header-section {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 8px;
}

.logo-box {
  height: 52px;
  margin-bottom: 2px;
}

.logo-img {
  height: 100%;
  object-fit: contain;
}

.title-badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 4px 0;
  font-size: 11px;
  font-weight: 700;
  color: #000;
  letter-spacing: 0.3px;
  text-transform: uppercase;

  border-bottom: 2px dotted #1e293b;
  padding-bottom: 4px;

  width: 100%;
}

.badge-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background-color: #000;
}

.description-box {
  font-size: 13px;
  line-height: 1.5;
  color: #475569;
  margin: 0;
  font-weight: 500;
}

/* ================= FEATURE BANNER ================= */
.feature-banner {
  position: relative;
  width: 100%;
  height: 280px;
  border-radius: 14px;
  overflow: hidden;
  box-shadow: 0 8px 16px -4px rgba(0,0,0,0.08);
  flex-shrink: 0;
}

.banner-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s ease;
}

.feature-banner:hover .banner-img {
  transform: scale(1.04);
}

.banner-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to top, rgba(15, 23, 42, 0.6) 0%, transparent 60%);
  display: flex;
  align-items: flex-end;
  padding: 10px 12px;
}

.banner-overlay::after{
  content: "";
  position: absolute;
  left: -20%;
  right: -20%;
  bottom: -35%;
  height: 80%;
  background: radial-gradient(
      ellipse at bottom,
      rgba(174, 174, 174, 0.55) 0%,
      rgba(247, 248, 250, 0.85) 35%,
      rgba(255,255,255,0) 80%
  );
  filter: blur(40px);
  pointer-events: none;
}

.live-tag {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 4px 10px;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(8px);
  border-radius: 20px;
  font-size: 10px;
  font-weight: 700;
  color: #0f172a;
   position: relative;
  z-index: 2;
}

.pulse-dot {
  width: 6px;
  height: 6px;
  background-color: #10b981;
  border-radius: 50%;
  box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.7);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% { box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.7); }
  70% { box-shadow: 0 0 0 6px rgba(16, 185, 129, 0); }
  100% { box-shadow: 0 0 0 0 rgba(16, 185, 129, 0); }
}

/* ================= STATS GRID ================= */
.stats-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  border-radius: 12px;
  transition: all 0.25s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
}

.stat-icon-bg {
  width: 44px;
  height: 44px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  overflow: hidden;
  transition: all 0.3s ease;
}

.lottie-svg-icon {
  width: 30px;
  height: 30px;
  object-fit: contain;
  transition: transform 0.3s ease;
}

.stat-icon-bg.blue { background-color: #eff6ff; }
.stat-icon-bg.green { background-color: #f0fdf4; }
.stat-icon-bg.purple { background-color: #faf5ff; }
.stat-icon-bg.orange { background-color: #fff7ed; }

.stat-card:hover .lottie-svg-icon {
  transform: scale(1.1) rotate(3deg);
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-number {
  font-size: 16px;
  font-weight: 800;
  color: #0f172a;
}

.stat-label {
  font-size: 11px;
  color: #64748b;
  font-weight: 600;
}

/* ================= LEFT PAGE: FOOTER / CONTACT ================= */
.footer-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-top: 4px;
  padding-top: 12px;
}

.contact-cards-stack {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 12px;
  border-radius: 10px;
  text-decoration: none;
  transition: all 0.25s ease;
  width: 100%;
  box-sizing: border-box;
}

.contact-icon {
  width: 16px;
  height: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #2563eb;
  flex-shrink: 0;
}

.contact-icon svg {
  width: 100%;
  height: 100%;
}

.contact-text {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  flex: 1;
  width: 100%;
}

.contact-label {
  font-size: 9px;
  font-weight: 700;
  color: #94a3b8;
  text-transform: uppercase;
  transition: color 0.25s ease;
}

.contact-val {
  font-size: 11px;
  font-weight: 600;
  color: #1e293b;
  transition: color 0.25s ease, border-color 0.25s ease;
  border-bottom: 2px dotted #1e293b;
  padding-bottom: 1px;
  display: block;
  width: 100%;
}

.social-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 4px;
}

.social-label {
  font-size: 11px;
  font-weight: 600;
  color: #64748b;
}

.social-buttons {
  display: flex;
  gap: 8px;
}

.social-btn {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
  transition: transform 0.2s ease;
}

.social-btn:hover {
  transform: scale(1.1);
}

.social-btn.facebook { background-color: #1877f2; }
.social-btn.telegram { background-color: #229ed9; }

.social-btn svg {
  width: 14px;
  height: 14px;
}

/* ================= RIGHT PAGE: STACKED TEAM CARDS ================= */
.team-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.section-header {
  margin-bottom: 2px;
}

.section-subtitle {
  font-size: 10px;
  font-weight: 800;
  color: #000;
  letter-spacing: 1px;
  text-transform: uppercase;
}

.section-title {
  font-size: 17px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}

.stacked-cards-container {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 8px 0;
  position: relative;
}

.team-review-card {
  position: relative;
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 12px 16px;
  border-radius: 16px;
  transition: transform 0.3s cubic-bezier(0.34, 1.56, 0.64, 1), box-shadow 0.3s ease;
  cursor: pointer;
}

.blue-card {
  color: #ffffff;
}

.blue-card .member-name { color: #000; }
.blue-card .member-desc { color: #000; }

.white-card {
  color: #1e1b18;
}

.white-card .member-name { color: #0f172a; }
.white-card .member-desc { color: #000; }

.rotate-left { transform: rotate(-1.5deg); }
.rotate-right { transform: rotate(1.8deg); }
.rotate-left-alt { transform: rotate(-2.2deg); }

.card-avatar {
  width: 50px;
  height: 50px;
  border-radius: 12px;
  overflow: hidden;
  flex-shrink: 0;
  border: 2px solid #00000049;
  position: relative;
}

.card-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.card-content {
  display: flex;
  flex-direction: column;
  gap: 2px;
  flex: 1;
  display: block;
}

.card-avatar::after {
  content: "";
  position: absolute;
  left: -15%;
  right: -15%;
  bottom: -30%;
  height: 70%;
  border-radius: 50%;
  background: linear-gradient(
    to top,
    rgb(28, 27, 27) 0%,
    rgba(249, 248, 248, 0.972) 45%,
    rgba(21, 21, 21, 0.494) 100%
  );
  filter: blur(14px);
  z-index: 2;
  pointer-events: none;
}

.card-header-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.member-name {
  font-size: 14px;
  font-weight: 800;
  margin: 0;
}

.member-desc {
  font-size: 10.5px;
  line-height: 1.35;
  margin: 0;
  font-weight: 500;
}

.sparkle {
  position: absolute;
  color: #2563eb;
  font-size: 16px;
  pointer-events: none;
}

.top-left-sparkle {
  top: -6px;
  left: 60px;
}

.bottom-right-sparkle {
  bottom: -8px;
  right: 20px;
}

.about-us-page {
  position: relative;
  overflow: hidden;
}



.about-us-page::before{
    content:"KAIFIN";
    position:absolute;
    left:50%;
    top:50%;
    transform:translate(-50%,-50%) rotate(-15deg);

    font-size:220px;
    font-weight:900;
    letter-spacing:18px;

    color:transparent;
    -webkit-text-stroke:2px rgba(37,99,235,.08);

    z-index:0;
}

.about-us-page::after{
    content:"";
    position:absolute;
    width:500px;
    height:500px;
    border-radius:50%;

    background:rgba(37,99,235,.04);

    top:-180px;
    right:-150px;

    filter:blur(80px);
}

.book-page{
    position:relative;
    z-index:2;
}


@media (max-width: 900px) {
  .app-wrapper { height: auto; overflow: auto; }
  .book-container { height: auto; }
  .about-us-page { grid-template-columns: 1fr; height: auto; }
  .about-us-page::after { display: none; }
  .book-page { padding: 20px 16px; }
}
</style>