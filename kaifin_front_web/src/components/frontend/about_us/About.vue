<script setup>
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import NavBar from '../navbar/NavBar.vue';
import Logo from "@/assets/logos/kaifin_l2.png";
import Phea from "@/assets/team_images/phea.jpg";
import Admin from "@/assets/team_images/admin.jpg";
import Heng from "@/assets/team_images/heng.jpg";
import Lyheng from "@/assets/team_images/lyheng.jpg";
import Mean from "@/assets/team_images/mean.jpg";
import Nary from "@/assets/team_images/nary.jpg";
import Nita from "@/assets/team_images/nita.jpg";
import Pa from "@/assets/team_images/pa.jpg";
import Roth from "@/assets/team_images/roth.jpg";
import Yuri from "@/assets/team_images/yuri.jpg";



const logoUrl = ref(Logo);

const speakers = ref([
  {
    name: "Pheaurn Sophea",
    bio: "Sharing work experience on successfully building and scaling engineering teams in ridesharing for 7.5 years, along with insights from project management and professional career development.",
    image: Phea, 
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Math Limin(Admin)",
    bio: "Discussing work experience in driving business growth, operational excellence, and scaling daily operations at Onde.",
    image: Admin,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "Heng Heng",
    bio: "Presenting work experience in software architecture, designing scalable distributed systems, and leading engineering projects.",
    image: Heng,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Lay LyHeng",
    bio: "Sharing work experience on product strategy, roadmap execution, and building user-centric solutions in tech.",
    image: Lyheng,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "Pheuing SreyMean",
    bio: "Talking about work experience in developing deep learning models, natural language processing, and deploying AI systems.",
    image: Mean,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Auok Sovanary (UI/UX Director)",
    bio: "Covering work experience in creating design systems, user research, and improving digital product experiences.",
    image: Nary,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "David Smith (Security Expert)",
    bio: "Sharing work experience in cloud security, vulnerability assessments, and enterprise data protection strategies.",
    image: Nita,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Jessica Taylor (Cloud Architect)",
    bio: "Discussing work experience in migrating legacy systems to cloud-native infrastructures and optimizing microservices.",
    image: Pa,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "Alexandre Dubois (DevOps Lead)",
    bio: "Presenting work experience in CI/CD pipeline automation, Docker containerization, and Kubernetes cluster management.",
    image: Roth,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Maria Garcia (Data Scientist)",
    bio: "Sharing work experience in data analytics, predictive modeling, and turning raw data into actionable business value.",
    image: Yuri,
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "James Wilson (Full-Stack Engineer)",
    bio: "Talking about work experience in developing high-performance web applications using modern full-stack technologies.",
    image: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Amina Khan (Blockchain Developer)",
    bio: "Covering work experience in smart contract development, security audits, and decentralized app architectures.",
    image: "https://images.unsplash.com/photo-1517841905240-472988babdf9?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "Lucas Martinez (Mobile Specialist)",
    bio: "Sharing work experience in cross-platform mobile app development, performance tuning, and UI integration.",
    image: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Chloe Bennett (QA Lead)",
    bio: "Discussing work experience in establishing automated testing frameworks and ensuring software release quality.",
    image: "https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "Nathaniel Wright (Solutions Architect)",
    bio: "Presenting work experience in bridging technical requirements with business goals for large-scale enterprise clients.",
    image: "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Zoe Adams (Growth Marketer)",
    bio: "Sharing work experience in user acquisition, digital marketing campaigns, and growth metric tracking.",
    image: "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "Ryan O'Connor (Systems Engineer)",
    bio: "Talking about work experience in low-level system optimization, hardware integration, and performance tuning.",
    image: "https://images.unsplash.com/photo-1519085360753-af0119f7cbe7?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  },
  {
    name: "Sophia Martinez (Community Manager)",
    bio: "Covering work experience in developer advocacy, community building, and organizing technical events globally.",
    image: "https://images.unsplash.com/photo-1544005313-94ddf0286df2?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: true
  },
  {
    name: "Daniel Kim (Innovations Lead)",
    bio: "Sharing work experience in exploring emerging technologies and managing research-to-product innovation cycles.",
    image: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=500&auto=format&fit=crop&q=60",
    github: "https://github.com",
    linkedin: "https://linkedin.com",
    isOffset: false
  }
]);
</script>

<template>
  <div>
    <NavBar/>
    <div class="about-container">
      <!-- Header Navigation -->
      <header class="header-nav">
      </header>

      <!-- Main Hero Section -->
      <main class="hero-content">
        <div class="hero-grid">
          <div class="inline-logo-container">
            <img :src="logoUrl" alt="Logo" class="inline-logo" />
          </div>
          <div class="text-wrapper">
            <h1 class="main-heading">
              Start learning. Share <br />
              knowledge. Build your future.
            </h1>
            
            <p class="description-text">
              Our e-learning social platform brings learners, <br />
              instructors, and knowledge seekers together in one <br />
              community. Discover new courses, share your <br />
              knowledge, connect with inspiring people, and <br />
              grow your skills together.
            </p>
          </div>
        </div>

        <!-- Video Section -->
        <!-- <div class="video-section">
          <div class="video-wrapper">
            <iframe 
              src="https://pixabay.com/videos/download/video-215471_medium.mp4" 
              title="YouTube video player" 
              frameborder="0" 
              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" 
              allowfullscreen>
            </iframe>
          </div>
        </div> -->
        <div class="video-section">
  <div class="video-wrapper">
    <video
      controls
      playsinline
      preload="metadata"
      poster=""
    >
      <source
        src="https://pixabay.com/videos/download/video-215471_medium.mp4"
        type="video/mp4"
      />
      Your browser does not support the video tag.
    </video>
  </div>
</div>

        <!-- Developer Teams Section -->
        <div class="speakers-section">
          <h2 class="speakers-heading">Developer Teams</h2>
          
          <div class="speakers-grid">
            <div 
              v-for="(speaker, index) in speakers" 
              :key="index" 
              class="speaker-card"
              :class="{ 'speaker-card-offset': speaker.isOffset }"
            >
              <div class="speaker-img-wrapper">
                <div 
                  class="placeholder-avatar" 
                  :style="{ backgroundImage: `url(${speaker.image})` }"
                ></div>
              </div>
              <h3 class="speaker-name">{{ speaker.name }}</h3>
              <p class="speaker-bio">{{ speaker.bio }}</p>
              
              <!-- GitHub & LinkedIn Links -->
              <div class="speaker-social-links">
                <a :href="speaker.github" target="_blank" rel="noopener noreferrer" class="speaker-social-link">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"/></svg>
                  <span>GitHub</span>
                </a>
                <a :href="speaker.linkedin" target="_blank" rel="noopener noreferrer" class="speaker-social-link">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"></path><rect x="2" y="9" width="4" height="12"></rect><circle cx="4" cy="4" r="2"></circle></svg>
                  <span>LinkedIn</span>
                </a>
              </div>
            </div>
          </div>
        </div>
      </main>

      <!-- Footer Section -->
      <footer class="footer-section">
        <div class="footer-top">
          <div class="footer-logo">Kafin</div>
          <div class="social-icons">
            <a href="#" class="social-icon" aria-label="Instagram">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
            <a href="#" class="social-icon" aria-label="YouTube">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor"><path d="M22.54 6.42a2.78 2.78 0 0 0-1.95-1.96C18.88 4 12 4 12 4s-6.88 0-8.59.46a2.78 2.78 0 0 0-1.95 1.96A29 29 0 0 0 1 12a29 29 0 0 0 .46 5.58 2.78 2.78 0 0 0 1.95 1.96C5.12 20 12 20 12 20s6.88 0 8.59-.46a2.78 2.78 0 0 0 1.95-1.96A29 29 0 0 0 23 12a29 29 0 0 0-.46-5.58zM9.75 15.02V8.98L15.5 12l-5.75 3.02z"></path></svg>
            </a>
            <a href="#" class="social-icon" aria-label="LinkedIn">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"></path><rect x="2" y="9" width="4" height="12"></rect><circle cx="4" cy="4" r="2"></circle></svg>
            </a>
          </div>
        </div>

        <div class="footer-middle">
          <p class="footer-spring">January 2026</p>
        </div>

        <div class="footer-bottom">
          <p class="footer-copyright">Chomkar Dong Comkar Dong Phnom Penh</p>
        </div>
      </footer>
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;600;800;900&display=swap');

.about-container {
  min-height: 100vh;
  width: 1251px;
  max-width: 100%;
  margin: 0 auto;
  background: #ffffff;
  font-family: 'Plus Jakarta Sans', system-ui, -apple-system, sans-serif;
  color: #111111;
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  padding: 10px 60px 40px 60px;
  position: relative;
  overflow-x: hidden;
}

/* Header Navigation Styles */
.header-nav {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  width: 100%;
  z-index: 10;
  min-height: 10px;
}

/* Hero Content Styles */
.hero-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 10px;
  z-index: 10;
  width: 100%;
}

.hero-grid {
  display: grid;
  grid-template-columns: auto 1fr;
  align-items: start;
  gap: 25px;
  width: 100%;
}

.inline-logo-container {
  padding-top: 20px;
}

.inline-logo {
  height: 110px;
  width: auto;
  object-fit: contain;
}

.text-wrapper {
  display: flex;
  flex-direction: column;
  justify-content: center;
  margin-top: 100px;
}

.main-heading {
  font-size: 64px;
  font-weight: 900;
  line-height: 1.05;
  letter-spacing: -1.5px;
  margin: 0 0 28px 0;
  color: #1a1a1a;
}

.description-text {
  font-size: 22px;
  font-weight: 400;
  line-height: 1.5;
  color: #555555;
  margin: 0;
}

/* Video Section Styles */
.video-section {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  max-width: 800px;
  margin-top: 130px; 
}

.video-wrapper {
  position: relative;
  width: 100%;
  aspect-ratio: 16 / 9;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
}

.video-wrapper iframe {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border: none;
}

/* Speakers Section Styles */
.speakers-section {
  width: 100%;
  max-width: 1000px;
  margin-top: 120px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.speakers-heading {
  font-size: 56px;
  font-weight: 900;
  color: #ff3366;
  margin-bottom: 50px;
  letter-spacing: -1px;
}

.speakers-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 80px 60px;
  width: 100%;
  align-items: start;
}

.speaker-card {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.speaker-card-offset {
  margin-top: 160px;
}

.speaker-img-wrapper {
  width: 100%;
  max-width: 280px;
  aspect-ratio: 1 / 1;
  overflow: hidden;
  margin-bottom: 25px;
  background: linear-gradient(135deg, #ff3366, #9333ea);
}

.placeholder-avatar {
  width: 100%;
  height: 100%;
  border: 4px solid #FF3266;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
}

.speaker-name {
  font-size: 28px;
  font-weight: 800;
  color: #ff3366;
  margin: 0 0 15px 0;
  line-height: 1.2;
}

.speaker-bio {
    background-color: #0000000c;
    border-left: 4px solid #FF3266;
    padding-left: 8px;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.6;
  color: #555555;
  margin: 0;
}

/* Speaker Social Links Styles */
.speaker-social-links {
  display: flex;
  gap: 15px;
  margin-top: 16px;
}

.speaker-social-link {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  font-weight: 600;
  color: #1a1a1a;
  text-decoration: none;
  background-color: #f3f4f6;
  padding: 6px 12px;
  border-radius: 6px;
  transition: background-color 0.2s, color 0.2s;
}

.speaker-social-link:hover {
  background-color: #ff3366;
  color: #ffffff;
}

/* Footer Section Styles */
.footer-section {
  width: calc(100% + 120px);
  margin-left: -60px;
  margin-right: -60px;
  margin-top: 140px;
  background-color: #1a1a1a;
  color: #ffffff;
  padding: 40px 60px;
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
}

.footer-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.footer-logo {
  font-size: 32px;
  font-weight: 800;
  letter-spacing: -1px;
  color: #ffffff;
}

.social-icons {
  display: flex;
  gap: 15px;
}

.social-icon {
  width: 38px;
  height: 38px;
  background-color: #ffffff;
  color: #1a1a1a;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-decoration: none;
  transition: opacity 0.2s;
}

.social-icon:hover {
  opacity: 0.8;
}

.footer-middle {
  margin-top: 25px;
}

.video-wrapper video {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  border: none;
}

.footer-spring {
  font-size: 16px;
  color: #cccccc;
  margin: 0;
}

.footer-bottom {
  margin-top: 40px;
}

.footer-copyright {
  font-size: 14px;
  color: #888888;
  margin: 0;
}

/* Responsive Design for smaller screens */
@media (max-width: 900px) {
  .about-container {
    padding: 15px;
  }
  
  .hero-content {
    margin-top: 0px;
  }

  .hero-grid {
    grid-template-columns: 1fr;
    gap: 15px;
  }

  .inline-logo-container {
    padding-top: 0;
  }

  .inline-logo {
    height: 60px;
  }

  .main-heading {
    font-size: 38px;
  }
  
  .description-text {
    font-size: 16px;
  }

  .speakers-grid {
    grid-template-columns: 1fr;
    gap: 40px;
  }

  .speakers-heading {
    font-size: 38px;
    margin-bottom: 30px;
  }

  .speaker-card-offset {
    margin-top: 0;
  }

  .footer-section {
    width: calc(100% + 30px);
    margin-left: -15px;
    margin-right: -15px;
    padding: 30px 15px;
  }

  .footer-top {
    flex-direction: column;
    align-items: flex-start;
    gap: 20px;
  }
}
</style>