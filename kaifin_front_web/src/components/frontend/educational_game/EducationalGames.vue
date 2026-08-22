<template>
  <div class="app-container">
    <NavBar />
    <!-- Banner-->
    <section class="page-title-banner">
      <video
        class="page-title-video"
        autoplay
        muted
        loop
        playsinline
        poster="https://images.unsplash.com/photo-1519741497674-611481863552?w=1400&auto=format&fit=crop&q=60"
      >
        <source src="" type="video/mp4" />
      </video>
      <div class="page-title-overlay"></div>
      <h1 class="page-title-text">Games</h1>
    </section>

    <div class="dashboard-wrapper">
      <div class="dashboard-main-layout">
        <!-- SIDEBAR Left -->
        <div class="sidebar-wrapper">
          <!-- Top profile -->
          <aside class="filter-sidebar profile-section">
            <div class="user-profile">
              <div class="avatar-box game-avatar">
                <svg class="svg-icon-md" viewBox="0 0 24 24" fill="none"
                     stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round">
                  <path d="M6 12l-2 3a3 3 0 0 0 2.7 4.5c1 0 1.8-.5 2.4-1.2l1.4-1.8h3l1.4 1.8c.6.7 1.4 1.2 2.4 1.2A3 3 0 0 0 20 15l-2-3a5 5 0 0 0-4.3-2.5H10.3A5 5 0 0 0 6 12z"/>
                  <path d="M8 11v4"/>
                  <path d="M6 13h4"/>
                  <circle cx="17" cy="12" r="1"/>
                  <circle cx="19" cy="14" r="1"/>
                </svg>
              </div>
              <div class="user-info">
                <span class="educational-game">Educational Game</span>
                <span class="user-status">Learn • Play • Grow</span>
              </div>
            </div>
          </aside>

          <!--Course section-->
          <aside class="filter-sidebar">
            <label class="filter-label">COURSE</label>
            <div class="course-list">
              <div class="course-item">
                <div class="tech-icon html">H</div>
                <span class="course-name">Html</span>
                <span class="badge-sub">
                  Subscribed
                  <svg class="svg-icon-xs" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
                    <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
                  </svg>
                </span>
              </div>
              <div class="course-item">
                <div class="tech-icon css">C</div>
                <span class="course-name">Css</span>
                <span class="badge-sub">
                  Subscribed
                  <svg class="svg-icon-xs" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
                    <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
                  </svg>
                </span>
              </div>
              <div class="more-link">
                More 
                <svg class="svg-icon-xs" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="5" y1="12" x2="19" y2="12"></line>
                  <polyline points="12 5 19 12 12 19"></polyline>
                </svg>
              </div>
            </div>
          </aside>

          <!-- Grade level -->
          <aside class="filter-sidebar">
            <div class="filter-item">
              <label class="filter-label">
                <svg class="svg-icon-xs" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M22 10v6M2 10l10-5 10 5-10 5z"/>
                  <path d="M6 12v5c3 3 9 3 12 0v-5"/>
                </svg>
                GRADE LEVEL
              </label>
              <div class="grid-2col">
                <button class="btn-option active">1st</button>
                <button class="btn-option">2nd</button>
                <button class="btn-option">3rd</button>
                <button class="btn-option">4th</button>
              </div>
            </div>

            <div class="filter-item">
              <label class="filter-label">
                <svg class="svg-icon-xs" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/>
                </svg>
                DIFFICULTY
              </label>
              <div class="vertical-stack">
                <button class="btn-option active">Easy</button>
                <button class="btn-option">Medium</button>
                <button class="btn-option">Hard</button>
              </div>
            </div>
          </aside>

          <!-- 4. HISTORY PLAY -->
          <aside class="filter-sidebar">
            <label class="filter-label">HISTORY PLAY</label>
            <div class="history-list-modern">
              <div v-for="(item, i) in historyList" :key="i" class="history-card-modern">
                <div class="history-img-wrap">
                  <img :src="item.image" :alt="item.name" />
                  <div class="history-badge-clock">
                    <svg class="svg-icon-xs" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                      <circle cx="12" cy="12" r="10"></circle>
                      <polyline points="12 6 12 12 16 14"></polyline>
                    </svg>
                  </div>
                </div>
                <div class="history-info">
                  <span class="history-name">{{ item.name }}</span>
                  <span class="history-action">Continue Playing</span>
                </div>
              </div>
            </div>
          </aside>
        </div>

        <!-- Main contend-->
        <div class="dashboard-container">
          <section v-for="cat in gameCategories" :key="cat.id" class="course-section">
            <div class="category-header-wrap">
              <h2 class="category-title">{{ cat.title }}</h2>
            </div>

            <!-- Card grid -->
            <div class="itch-games-grid">
              <div v-for="game in cat.games" :key="game.id" class="itch-game-card">
                <div class="itch-media-box" @click="goToGame(game)">
                  <img :src="game.image" :alt="game.name" class="itch-img" />
                </div>
                <div class="itch-info-box">
                  <div class="itch-main-content">
                    <div class="itch-title-row">
                      <h3 class="itch-game-title">{{ game.name }}</h3>
                    </div>
                    <p class="itch-tags-text">
                      <span v-for="(tag, idx) in game.tags" :key="idx" class="tag-item">
                        #{{ tag }}{{ idx < game.tags.length - 1 ? ', ' : '' }}
                      </span>
                    </p>

                    <p class="itch-desc-text">{{ game.description }}</p>
                  </div>
                  <button class="play-game-bottom-btn" @click="goToGame(game)">
                    <svg class="svg-icon-sm" viewBox="0 0 24 24" fill="currentColor">
                      <polygon points="5 3 19 12 5 21 5 3"></polygon>
                    </svg>
                    Play Game
                  </button>
                </div>
              </div>
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import NavBar from '../navbar/NavBar.vue'

const router = useRouter()
function goToGame(game) {
  router.push({
    path: `/game/${game.id}`,
    state: { gameData: game }
  })
}

const historyList = ref([
  { name: 'Math Quiz', image: 'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=200&auto=format&fit=crop&q=60' },
  { name: 'Word Search', image: 'https://images.unsplash.com/photo-1581291518633-83b4ebd1d83e?w=200&auto=format&fit=crop&q=60' },
  { name: 'Logic Puzzle', image: 'https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5?w=200&auto=format&fit=crop&q=60' }
])

const gameCategories = ref([
  {
    id: 1,
    title: 'Popular Game',
    games: [
      { 
        id: 101, 
        name: 'The Snake That Eats Puzzle Game Mechanics', 
        tags: ['Puzzle', 'PuzzleScript'], 
        description: 'A collaborative collection of puzzles', 
        image: 'https://images.unsplash.com/photo-1550745165-9bc0b252726f?w=500&auto=format&fit=crop&q=60' 
      },
      { 
        id: 102, 
        name: "Behind the Broadcast: The Director's Cut", 
        tags: ['Dating Sim', 'Cute'], 
        description: 'What happens when an insecure screenwriter meets highly admired citys famous director?', 
        image: 'https://images.unsplash.com/photo-1538481199705-c710c4e965fc?w=500&auto=format&fit=crop&q=60' 
      },
      { 
        id: 103, 
        name: 'Lazy Kickers', 
        tags: ['Sports', 'Desktop Pet'], 
        description: 'Crazy addictive soccer incremental!', 
        image: 'https://images.unsplash.com/photo-1511512578047-dfb367046420?w=500&auto=format&fit=crop&q=60' 
      },
      { 
        id: 104, 
        name: 'Midnight Scenes: The Highway', 
        tags: ['Point & Click', 'Atmospheric'], 
        description: 'Untold horrors in a dark, deserted road.', 
        image: 'https://images.unsplash.com/photo-1579373903781-fd5c0c30c4cd?w=500&auto=format&fit=crop&q=60' 
      }
    ]
  },
  {
    id: 2,
    title: 'Game with friend',
    games: [
      { 
        id: 201, 
        name: 'The Hive', 
        tags: ['Action', 'Singleplayer'], 
        description: 'Alien Insect army under your control!', 
        image: 'https://images.unsplash.com/photo-1542751371-adc38448a05e?w=500&auto=format&fit=crop&q=60' 
      },
      { 
        id: 202, 
        name: 'Death Spiral', 
        tags: ['Dark Humor', 'Horror'], 
        description: 'One button horror roguelike', 
        image: 'https://images.unsplash.com/photo-1612287230202-1ff1d85d1bdf?w=500&auto=format&fit=crop&q=60' 
      },
      { 
        id: 203, 
        name: 'Scritchy Scratchy', 
        tags: ['Idle', 'Simulation'], 
        description: 'A super-satisfying scratch card incremental game', 
        image: 'https://images.unsplash.com/photo-1560253023-3ec5d502959f?w=500&auto=format&fit=crop&q=60' 
      }
    ]
  }
])
</script>
<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap');

.svg-icon-xs {
  width: 13px;
  height: 13px;
  stroke-linecap: round;
  stroke-linejoin: round;
  display: inline-block;
  vertical-align: middle;
}

.svg-icon-sm {
  width: 14px;
  height: 14px;
  stroke-linecap: round;
  stroke-linejoin: round;
  display: inline-block;
  vertical-align: middle;
}

.svg-icon-md {
  width: 20px;
  height: 20px;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.app-container {
  width: 100%;
  min-height: 100vh;
}

.dashboard-wrapper {
  background-color: #F7F4F2;
  min-height: calc(100vh - 60px);
  width: 100%;
  padding: 20px 0;
  box-sizing: border-box;
  font-family: 'Plus Jakarta Sans', sans-serif;
  color: #111827;
}

.dashboard-main-layout {
  width: 85%;
  max-width: 1251px;
  margin: 0 auto;
  display: flex;
  align-items: flex-start;
  gap: 20px;
}

.sidebar-wrapper {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 240px;
  flex-shrink: 0;
  position: sticky;
  top: 20px;
}

.filter-sidebar {
  display: flex;
  flex-direction: column;
  gap: 12px;
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.04);
}

.profile-section {
  padding: 12px 16px;
}

.user-profile {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar-box {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background-color: #eef6fc;
  color: #1976D2;
  display: flex;
  align-items: center;
  justify-content: center;
}

.user-info {
  display: flex;
  flex-direction: column;
}

.educational-game {
  font-size: 16px;
  font-weight: 700;
  color: #111827;
}

.user-status {
  font-size: 10px;
  color: #10b981;
  font-weight: 600;
}

.filter-label {
  font-size: 11px;
  font-weight: 700;
  color: #374151;
  letter-spacing: 0.5px;
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 2px;
}

.filter-item {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

/* COURSE ITEM */
.course-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.course-item {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 8px 10px;
}

.tech-icon {
  width: 22px;
  height: 22px;
  border-radius: 4px;
  color: #ffffff;
  font-weight: 800;
  font-size: 11px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.tech-icon.html { background-color: #e44d26; }
.tech-icon.css { background-color: #1976D2; }

.course-name {
  font-size: 13px;
  font-weight: 600;
  color: #374151;
}

.badge-sub {
  font-size: 9px;
  color: #1976D2;
  background: #eef6fc;
  padding: 3px 6px;
  border-radius: 6px;
  font-weight: 700;
  margin-left: auto;
  display: inline-flex;
  align-items: center;
  gap: 3px;
}

.more-link {
  font-size: 12px;
  color: #6b7280;
  text-align: right;
  cursor: pointer;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 4px;
}

.more-link:hover {
  color: #1976D2;
}

.grid-2col {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 6px;
}

.vertical-stack {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.btn-option {
  font-size: 12px;
  color: #4b5563;
  padding: 8px;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
  background: #f9fafb;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s ease;
}

.btn-option:hover {
  background: #f3f4f6;
  color: #111827;
}

.btn-option.active {
  background: #1976D2;
  color: #ffffff;
  border-color: #1976D2;
}

.history-list-modern {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.history-card-modern {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 6px;
  border-radius: 8px;
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  cursor: pointer;
  transition: all 0.2s ease;
}

.history-card-modern:hover {
  background: #f3f4f6;
  border-color: #d1d5db;
  transform: translateY(-1px);
}

.history-img-wrap {
  position: relative;
  width: 40px;
  height: 40px;
  flex-shrink: 0;
}

.history-img-wrap img {
  width: 100%;
  height: 100%;
  border-radius: 6px;
  object-fit: cover;
  border: 1px solid #e5e7eb;
}

.history-badge-clock {
  position: absolute;
  bottom: -3px;
  right: -3px;
  background: #ffffff;
  border-radius: 50%;
  padding: 2px;
  box-shadow: 0 1px 2px rgba(0,0,0,0.1);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #1976D2;
}

.history-info {
  display: flex;
  flex-direction: column;
  min-width: 0;
}

.history-name {
  font-size: 12px;
  font-weight: 700;
  color: #111827;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.history-action {
  font-size: 10px;
  color: #6b7280;
  font-weight: 500;
}

.dashboard-container {
  flex: 1;
  min-width: 0;
}

.page-title-banner {
  position: relative;
  width: 100vw;
  left: 50%;
  right: 50%;
  margin-left: -50vw;
  margin-right: -50vw;
  height: 300px;
  overflow: hidden;
  display: flex;
  align-items: flex-end;
  padding: 24px 40px;
  box-sizing: border-box;
  background-color: #000000;
  margin-bottom: 12px;
}

.page-title-video {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  z-index: 0;
}

.page-title-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(15, 15, 20, 0.75) 0%, rgba(160, 160, 168, 0.234) 60%, rgba(15, 15, 20, 0.75) 100%);
  z-index: 1;
}

.page-title-text {
  position: relative;
  z-index: 2;
  color: #ffffff;
  font-size: 36px;
  font-weight: 800;
  margin: 0;
  line-height: 1;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
}


.course-section {
  margin-bottom: 32px;
}

.category-header-wrap {
  margin-bottom: 16px;
}

.category-title {
  font-size: 16px;
  font-weight: 700;
  color: #111827;
  border-left: 4px solid #1976D2;
  padding-left: 10px;
  margin: 0;
}


.itch-games-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

.itch-game-card {
  background: transparent;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.itch-media-box {
  position: relative;
  width: 100%;
  aspect-ratio: 16 / 10;
  background: #111;
  border-radius: 4px;
  overflow: hidden;
  box-shadow: 0 2px 6px rgba(0,0,0,0.08);
  cursor: pointer;
}

.itch-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.itch-media-box:hover .itch-img {
  transform: scale(1.05);
}


.itch-info-box {
  padding-top: 8px;
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}

.itch-main-content {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin-bottom: 8px;
}

.itch-title-row {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 8px;
}

.itch-game-title {
  font-size: 13px;
  font-weight: 700;
  color: #1a1a1a;
  margin: 0;
  line-height: 1.3;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.itch-tags-text {
  font-size: 11px;
  color: #fa5c5c;
  margin: 0;
  font-weight: 500;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.itch-desc-text {
  font-size: 11px;
  color: #6b7280;
  margin: 0;
  line-height: 1.3;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.play-game-bottom-btn {
  background: #1976D2;
  color: #ffffff;
  border: none;
  width: 100%;
  padding: 7px 12px;
  border-radius: 6px;
  font-size: 12px;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  cursor: pointer;
  transition: background-color 0.2s ease, transform 0.1s ease;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  margin-top: auto;
}

.play-game-bottom-btn:hover {
  background: #1565c0;
}

.play-game-bottom-btn:active {
  transform: scale(0.98);
}

@media (max-width: 900px) {
  .dashboard-main-layout {
    width: 95%;
    flex-direction: column;
  }
  .sidebar-wrapper {
    width: 100%;
    position: static;
  }
}
</style>