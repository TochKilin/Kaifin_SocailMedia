<template>
  <div class="page-wrapper">
    <NavBar />
    <!-- Main content page -->
    <div class="hole-up-page">
      <div class="hole-up-card">
        
        <!-- Header-->
        <div class="card-header-row">
          <button class="header-icon-btn" @click="goBack" title="Back">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <line x1="19" y1="12" x2="5" y2="12"></line>
              <polyline points="12 19 5 12 12 5"></polyline>
            </svg>
          </button>
          <span class="header-title">{{ isPlaying ? currentGame?.name : 'Game Details' }}</span>
          <div class="header-right-actions">
            <button class="header-icon-btn" title="Bookmark">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
              </svg>
            </button>
            <button class="header-icon-btn" title="Share">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path>
                <polyline points="16 6 12 2 8 6"></polyline>
                <line x1="12" y1="2" x2="12" y2="15"></line>
              </svg>
            </button>
          </div>
        </div>

        <!-- Scroll body -->
        <div class="card-body-scroll" v-if="!isPlaying">
          
          <!-- 2. Game Imag -->
          <div class="preview-media-box" @click="startGame">
            <img 
              v-if="currentGame?.image" 
              :src="currentGame.image" 
              :alt="currentGame?.name" 
              class="preview-img" 
            />
            <div v-else class="fallback-img-text">
              <span>No Preview Image</span>
            </div>

            <div class="play-icon-circle">
              <svg viewBox="0 0 24 24" fill="currentColor">
                <polygon points="5 3 19 12 5 21 5 3"></polygon>
              </svg>
            </div>
          </div>

          <div class="badges-row">
            <span class="badge trending">
              <svg class="badge-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.229 1-3.032.5.803.5 2.032.5 3.032z"></path></svg>
              Trending
            </span>
            <span class="badge puzzle">
              <svg class="badge-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M19.43 12.98c.04-.32.07-.64.07-.98 0-.34-.03-.66-.07-.98l2.11-1.65c.19-.15.24-.42.12-.64l-2-3.46c-.12-.22-.39-.3-.61-.22l-2.49 1c-.52-.4-1.08-.73-1.69-.98l-.37-2.65A.506.506 0 0 0 14 2h-4c-.25 0-.46.18-.5.42l-.37 2.65c-.61.25-1.17.59-1.69.98l-2.49-1c-.23-.09-.49 0-.61.22l-2 3.46c-.13.22-.07.49.12.64l2.11 1.65c-.04.32-.07.65-.07.98 0 .33.03.66.07.98l-2.11 1.65c-.19.15-.24.42-.12.64l2 3.46c.12.22.39.3.61.22l2.49-1c.52.4 1.08.73 1.69.98l.37 2.65c.04.24.25.42.5.42h4c.25 0 .46-.18.5-.42l.37-2.65c.61-.25 1.17-.59 1.69-.98l2.49 1c.23.09.49 0 .61-.22l2-3.46c.12-.22.07-.49-.12-.64l-2.11-1.65z"></path><circle cx="12" cy="12" r="3"></circle></svg>
              {{ currentGame?.tags?.[0] || 'Match-3' }}
            </span>
            <span class="badge rating">
              <svg class="badge-svg star" viewBox="0 0 24 24" fill="currentColor"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/></svg>
              {{ currentGame?.rating || '4.8' }}
            </span>
          </div>

          <div class="title-dev-section">
            <h2 class="game-title">{{ currentGame?.name || 'SpelunKing' }}</h2>
            <p class="developer-text">By {{ currentGame?.developer || 'Mad Data' }}</p>
          </div>

          <div class="meta-info-row">
            <div class="meta-item">
              <svg class="meta-svg highlight-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
              <span class="meta-text"><strong class="highlight-stat">300K</strong> Player</span>
            </div>
            <div class="meta-item">
              <svg class="meta-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
              <span class="meta-text">{{ currentGame?.time || '10 min' }}</span>
            </div>
            <div class="meta-item">
              <svg class="meta-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="2" y1="12" x2="22" y2="12"></line><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path></svg>
              <span class="meta-text">{{ currentGame?.network || 'Online' }}</span>
            </div>
          </div>

          <div class="about-section">
            <h3 class="section-heading">About</h3>
            <p class="about-text">
              {{ currentGame?.description || 'Dig into this deep match-3 expedition! Take a dive underground, search for your grandpa’s lost items, and collect precious stones to help the villagers around town rebuild Farnsbury.' }}
            </p>
          </div>

          <button class="play-action-btn" @click="startGame">
            <svg class="btn-icon" viewBox="0 0 24 24" fill="currentColor">
              <polygon points="5 3 19 12 5 21 5 3"></polygon>
            </svg>
            <span>PLAY NOW</span>
          </button>

          <div class="screenshots-section" v-if="currentGame?.screenshots && currentGame.screenshots.length > 0">
            <h3 class="section-heading flex-heading">
              <svg class="heading-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"></path><circle cx="12" cy="13" r="4"></circle></svg>
              Screenshots
            </h3>
            <div class="screenshots-grid">
              <div class="screenshot-box" v-for="(shot, index) in currentGame.screenshots" :key="index">
                <img :src="shot" alt="Screenshot" class="screenshot-img" />
              </div>
            </div>
          </div>
        </div>
        <div class="game-embed-container" v-if="isPlaying">
          <button class="close-game-btn" @click="stopGame" title="Close Game">✕ Close Game</button>
          <iframe 
            src="https://html5.gamedistribution.com/a3a4111db59d496b96de651c307009ad/?gd_sdk_referrer_url=https://gamedistribution.com/games/spelunking/" 
            width="100%" 
            height="100%" 
            scrolling="none" 
            frameborder="0"
            allowfullscreen>
          </iframe>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import NavBar from '../navbar/NavBar.vue'

const route = useRoute()
const router = useRouter()
const isPlaying = ref(false)

const allGames = [
  { 
    id: 101, 
    name: 'SpelunKing', 
    developer: 'Mad Data',
    tags: ['Match-3'], 
    rating: '4.8',
    players: '300K Player',
    time: '10 min',
    network: 'Online',
    description: 'Dig into this deep match-3 expedition! Take a dive underground, search for your grandpa’s lost items, and collect precious stones to help the villagers around town rebuild Farnsbury.', 
    image: 'https://img.gamedistribution.com/a3a4111db59d496b96de651c307009ad-512x512.jpeg',
    screenshots: [
      'https://img.gamedistribution.com/a3a4111db59d496b96de651c307009ad-512x512.jpeg'
    ]
  }
]

const currentGame = computed(() => {
  const gameId = Number(route.params.id)
  return history.state?.gameData || allGames.find(g => g.id === gameId) || allGames[0]
})

const startGame = () => {
  isPlaying.value = true
}

const stopGame = () => {
  isPlaying.value = false
}

const goBack = () => {
  if (isPlaying.value) {
    isPlaying.value = false
  } else {
    router.go(-1)
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@500;600;700;800&display=swap');

.page-wrapper {
  position: fixed;
  inset: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #F7F4F2;
  font-family: 'Plus Jakarta Sans', sans-serif;
  z-index: 999;
  overflow: hidden;
}

.hole-up-page {
  flex: 1;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.hole-up-card {
  position: relative;
  width: 620px;
  max-width: 96%;
  height: 100%;
  background-color: #FFFFFF;
  border-radius: 0;
  padding: 16px 20px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  margin: 0 auto;
  overflow: hidden;
  box-shadow: none;
}

/* Header Row */
.card-header-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 12px;
  border-bottom: 1px solid #f1f5f9;
  flex-shrink: 0;
}

.header-icon-btn {
  width: 38px;
  height: 38px;
  background-color: #f1f5f9;
  border: none;
  border-radius: 12px;
  color: #0f172a;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s ease;
}

.header-icon-btn svg {
  width: 18px;
  height: 18px;
}

.header-icon-btn:hover {
  background-color: #e2e8f0;
}

.header-title {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}

.header-right-actions {
  display: flex;
  gap: 8px;
}

.card-body-scroll {
  flex: 1;
  overflow-y: auto;
  padding-top: 16px;
  padding-bottom: 24px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  scrollbar-width: thin;
}

.preview-media-box {
  position: relative;
  width: 100%;
  height: 220px;
  background-color: #f8fafc;
  border-radius: 16px;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  cursor: pointer;
}

.preview-img {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.fallback-img-text {
  color: #94a3b8;
  font-size: 14px;
  font-weight: 600;
}

.play-icon-circle {
  position: relative;
  z-index: 2;
  width: 48px;
  height: 48px;
  background-color: #1B75D2;
  border-radius: 50%;
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
}

.play-icon-circle svg {
  width: 22px;
  height: 22px;
  fill: #ffffff;
  margin-left: 2px;
}

.badges-row {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.badge {
  display: flex;
  align-items: center;
  gap: 5px;
  font-size: 12px;
  font-weight: 700;
  padding: 6px 12px;
  border-radius: 10px;
  background-color: #f8fafc;
  color: #475569;
  border: 1px solid #f1f5f9;
}

.badge-svg {
  width: 14px;
  height: 14px;
}

.badge-svg.star {
  fill: #f59e0b;
  color: #f59e0b;
}

.title-dev-section {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.game-title {
  font-size: 20px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}

.developer-text {
  font-size: 13px;
  font-weight: 700;
  color: #1B75D2;
  margin: 0;
}


.meta-info-row {
  display: flex;
  gap: 16px;
  background-color: transparent;
  border: none;
  border-radius: 0;
  padding: 4px 0;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  font-weight: 600;
  color: #475569;
}

.meta-svg {
  width: 16px;
  height: 16px;
  color: #64748b;
}

.highlight-svg {
  color: #1B75D2;
}

.highlight-stat {
  color: #1B75D2;
  font-weight: 800;
  font-size: 14px;
}

.about-section {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.section-heading {
  font-size: 15px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}

.flex-heading {
  display: flex;
  align-items: center;
  gap: 6px;
}

.heading-svg {
  width: 18px;
  height: 18px;
  color: #0f172a;
}

.about-text {
  font-size: 13px;
  color: #64748b;
  line-height: 1.5;
  margin: 0;
}

.play-action-btn {
  width: 100%;
  height: 56px;
  border: none;
  border-radius: 18px;
  cursor: pointer;
  font-size: 17px;
  font-weight: 800;
  color: #fff;
  letter-spacing: .4px;
  background-color: #1B75D2;
  transition: .28s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.btn-icon {
  width: 20px;
  height: 20px;
  fill: #fff;
}

.play-action-btn:hover {
  transform: translateY(-2px);
  background-color: #1562b8;
}

.screenshots-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.screenshots-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 8px;
}

.screenshot-box {
  width: 100%;
  height: 80px;
  background-color: #f8fafc;
  border-radius: 12px;
  border: 1px solid #f1f5f9;
  overflow: hidden;
}

.screenshot-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.game-embed-container {
  flex: 1;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding-top: 10px;
  height: 100%;
}

.close-game-btn {
  align-self: flex-start;
  background-color: #fee2e2;
  color: #991b1b;
  border: none;
  padding: 6px 14px;
  border-radius: 8px;
  font-weight: 700;
  font-size: 12px;
  cursor: pointer;
}

.close-game-btn:hover {
  background-color: #fecaca;
}

.game-embed-container iframe {
  flex: 1;
  width: 100%;
  height: 100%;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
}
</style>