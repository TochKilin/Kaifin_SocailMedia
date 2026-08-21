<template>
  <div class="sponsors-section">
    <div class="container">
      
      <!-- Sponsore-->
      <div class="section-header">
        <div class="header-left">
          <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M6 3h12l4 6-10 13L2 9z"></path>
            <path d="M11 3l8 6-7 13L3 9z"></path>
            <path d="M2 9h20"></path>
          </svg>
          <h2 class="section-title">Sponsore By</h2>
          <span class="badge">{{ sponsors.length }}</span>
        </div>
        <button class="view-all-btn" @click="handleViewAll">
          See All 
          <svg class="arrow-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M5 12h14"></path>
            <path d="m12 5 7 7-7 7"></path>
          </svg>
        </button>
      </div>

      <!-- Auoto scroll-->
      <div class="sponsors-scroll-container">
        <div class="sponsors-row">
          <template v-for="loop in 2" :key="loop">
            <div 
              v-for="sponsor in sponsors" 
              :key="sponsor.id + '-' + loop"
              class="sponsor-banner"
              :style="{ backgroundImage: `url(${sponsor.bgImage})` }"
              @click="handleCardClick(sponsor.name)"
            >
              <div class="banner-overlay"></div>
              <div class="banner-content">
                <h3 class="banner-title">{{ sponsor.name }}</h3>
                <button class="shop-btn" @click.stop="handleCooperate(sponsor.name)">
                  Read More
                </button>
              </div>
              <div class="banner-right">
                <div class="avatar-box">
                  <img :src="sponsor.avatar" :alt="sponsor.name" class="sponsor-img" />
                </div>
              </div>
            </div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const emit = defineEmits(['viewAll', 'cardClick', 'cooperate']);
const sponsors = ref([
  {
    id: 1,
    name: 'Sastra Film',
    bgImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSITGxUonKiQ-e9B_bFb1ESMPWOvUuHiZM7daEcoGXkGw&s',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT2T8deHA9KKqeC6twW8_pKCtPF6R7VUy1ae_k9zRCKCg&s'
  },
  {
    id: 2,
    name: 'Miss CEO',
    bgImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRivEO2-12F7G3fjLgzaIvwWraCoFlJVTnMm6UvqekfSg&s',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4i0j-B0jZ1qCO_7_uV2n4t0uHnyJnaBw3fOslMGng7Q&s=10'
  },
  {
    id: 3,
    name: 'Rupp',
    bgImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyvLU7mMbM-LsYiA1C_dIvUyQb8zyr0Y7nNwiw5EO-fQ&s=10',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyvLU7mMbM-LsYiA1C_dIvUyQb8zyr0Y7nNwiw5EO-fQ&s=10'
  },
  {
    id: 4,
    name: 'ACE',
    bgImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTbya7iMGUcAsxeqG4ehELbHTjjVuoGQuxTi6iMt9alLw&s=10',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR7oH6Unz2Hr6h1pEos_4Kwb0pYm5wzqWnB76cwQcdW7w&s=10'
  }
]);

function handleViewAll() {
  emit('viewAll');
}

function handleCardClick(name) {
  emit('cardClick', name);
}

function handleCooperate(name) {
  emit('cooperate', name);
  alert(`sponsore request ${name} success!`);
}
</script>

<style scoped>
.sponsors-section {
  margin-top: 15px;
  font-family: 'Plus Jakarta Sans', system-ui, -apple-system, sans-serif;
  color: #1e293b;
}

.container {
  max-width: 1400px;
  margin: 0 auto;
  border-radius: 24px;
  padding: 24px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.header-icon {
  width: 22px;
  height: 22px;
  color: #2563eb;
}

.section-title {
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

.badge {
  background: #f1f5f9;
  color: #475569;
  font-size: 12px;
  font-weight: 600;
  padding: 2px 8px;
  border-radius: 12px;
}

.view-all-btn {
  background: none;
  border: none;
  color: #2563eb;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 4px;
  font-family: inherit;
}

.view-all-btn:hover {
  text-decoration: underline;
}

.arrow-icon {
  width: 16px;
  height: 16px;
}

.sponsors-scroll-container {
  width: 100%;
  overflow: hidden;
  position: relative;
  display: flex;
  padding: 4px 0;
}

.sponsors-row {
  display: flex;
  gap: 16px;
  width: max-content;
  animation: scrollX 25s linear infinite;
}

.sponsors-scroll-container:hover .sponsors-row {
  animation-play-state: paused;
}

@keyframes scrollX {
  0% {
    transform: translateX(0);
  }
  100% {
    transform: translateX(-50%);
  }
}

.sponsor-banner {
  width: 300px;
  flex-shrink: 0;
  border-radius: 16px;
  padding: 18px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  border: 1px solid rgba(255, 255, 255, 0.8);
  position: relative;
  overflow: hidden;
  background-size: cover;
  background-position: center;
}

.banner-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.75);
  z-index: 1;
}

.sponsor-banner:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.05);
}

.banner-content {
  flex: 1;
  padding-right: 12px;
  position: relative;
  z-index: 2;
}

.banner-title {
  font-size: 16px;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 10px 0;
}

.shop-btn {
  background: #ffffff;
  color: #0f172a;
  border: none;
  border-radius: 8px;
  padding: 6px 12px;
  font-size: 11px;
  font-weight: 700;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  transition: all 0.2s;
  font-family: inherit;
}

.shop-btn:hover {
  background: #0f172a;
  color: #ffffff;
}

.banner-right {
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  z-index: 2;
}

.avatar-box {
  width: 52px;
  height: 52px;
  background: #ffffff;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.06);
  border: 2px solid rgba(255, 255, 255, 0.9);
  overflow: hidden;
}

.sponsor-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>