<!-- File: Review.vue -->
<script setup>
import { ref } from 'vue'

defineEmits(['back'])

const searchQuery = ref('')
const reviews = ref([
  {
    id: 1,
    username: 'សុខ សុភា',
    avatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100&h=100&fit=crop&crop=faces',
    rating: 4,
    time: '1 day ago',
    text: 'Text feedback about the course content and teaching style.',
    likes: 12,
    dislikes: 2
  },
  {
    id: 2,
    username: 'ច័ន្ទ វណ្ណា',
    avatar: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&h=100&fit=crop&crop=faces',
    rating: 5,
    time: '3 days ago',
    text: 'Very clear explanation and great examples provided throughout the lessons.',
    likes: 8,
    dislikes: 0
  },
  {
    id: 3,
    username: 'ឃឹម រតនា',
    avatar: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=100&h=100&fit=crop&crop=faces',
    rating: 4,
    time: '5 days ago',
    text: 'Helpful exercises, though some sections could use more detailed guides.',
    likes: 4,
    dislikes: 1
  },
  {
    id: 4,
    username: 'ហេង ដារ៉ា',
    avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=100&h=100&fit=crop&crop=faces',
    rating: 5,
    time: '1 week ago',
    text: 'Amazing instructor! Learned a lot of practical skills that I can apply directly.',
    likes: 20,
    dislikes: 0
  },
  {
    id: 5,
    username: 'លី ម៉េងហួរ',
    avatar: '',
    rating: 4,
    time: '2 weeks ago',
    text: 'Text feedback about materials and exercises provided.',
    likes: 5,
    dislikes: 0
  }
])
</script>

<template>
  <div class="review-container">
    
    <!-- Header with Back Button & Centered Title -->
    <div class="review-header">
      <button class="back-btn" @click="$emit('back')" title="Go back">
        <svg class="transparent-svg" viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="15 18 9 12 15 6"></polyline>
        </svg>
      </button>
      <h2>Reviews</h2>
      <div class="header-spacer"></div>
    </div>

    <!-- Search Box -->
    <div class="search-bar-wrapper">
      <input 
        type="text" 
        v-model="searchQuery" 
        placeholder="Search reviews" 
        class="search-input"
      />
      <button class="search-btn">
        <svg class="transparent-svg" viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"></circle>
          <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
        </svg>
      </button>
    </div>

    <!-- Reviews List -->
    <div class="reviews-list">
      <div v-for="review in reviews" :key="review.id" class="review-card">
        
        <div class="review-user-info">
          <!-- Avatar Box with Image or Fallback SVG Icon -->
          <div class="avatar-box">
            <img v-if="review.avatar" :src="review.avatar" alt="User Avatar" class="avatar-img" />
            <svg v-else class="transparent-svg" viewBox="0 0 24 24" width="24" height="24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>

          <div class="user-meta">
            <div class="username-box">{{ review.username }}</div>
            <div class="rating-time-row">
              <div class="stars">
                <span v-for="n in 5" :key="n" :class="n <= review.rating ? 'star filled' : 'star'">★</span>
              </div>
              <span class="time-badge">
                <svg class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                {{ review.time }}
              </span>
            </div>
          </div>
        </div>

        <div class="feedback-text-box">
          <p>{{ review.text }}</p>
        </div>

        <div class="review-actions">
          <!-- Like Button -->
          <button class="action-btn" @click="review.likes++">
            <span class="icon-circle">
              <svg class="transparent-svg" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path>
              </svg>
            </span>
            <span class="btn-count">{{ review.likes }}</span>
          </button>

          <!-- Dislike Button -->
          <button class="action-btn" @click="review.dislikes++">
            <span class="icon-circle">
              <svg class="transparent-svg" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M10 15v4a3 3 0 0 0 3 3l4-9V2H5.72a2 2 0 0 0-2 1.7l-1.38 9a2 2 0 0 0 2 2.3zm7-13h3a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2h-3"></path>
              </svg>
            </span>
            <span class="btn-count">{{ review.dislikes }}</span>
          </button>

          <button class="report-btn">Report</button>
        </div>

      </div>
    </div>

  </div>
</template>

<style scoped>
.transparent-svg {
  background-color: transparent !important;
}

.review-container {
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  padding: 16px 18px;
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
}

.review-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.review-header h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
  text-align: center;
}

.header-spacer {
  width: 40px;
}

.back-btn {
  background-color: transparent;
  color: #0f172a;
  border: none;
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.back-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.search-bar-wrapper {
  display: flex;
  align-items: center;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  padding: 6px 12px;
}

.search-input {
  flex: 1;
  background: transparent;
  border: none;
  outline: none;
  color: #0f172a;
  font-size: 14px;
  padding: 8px;
}

.search-input::placeholder {
  color: #94a3b8;
}

.search-btn {
  background-color: #e2e8f0;
  color: #0f172a;
  border: none;
  width: 36px;
  height: 36px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.search-btn:hover {
  background-color: #cbd5e1;
}

.reviews-list {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.review-card {
  background-color: rgba(255, 255, 255, 0.6);
  border: 1px solid rgba(226, 232, 240, 0.8);
  border-radius: 14px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.02);
}

.review-user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar-box {
  width: 42px;
  height: 42px;
  border: 1px solid #cbd5e1;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #64748b;
  background-color: rgba(248, 250, 252, 0.8);
  overflow: hidden;
}

.avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.user-meta {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.username-box {
  background: transparent;
  border: none;
  padding: 0;
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
  display: inline-block;
  width: fit-content;
}

.rating-time-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.stars {
  color: #cbd5e1;
  font-size: 14px;
}

.stars .filled {
  color: #f59e0b;
}

.time-badge {
  border: 1px solid #e2e8f0;
  background-color: rgba(248, 250, 252, 0.8);
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 11px;
  color: #64748b;
  display: flex;
  align-items: center;
  gap: 4px;
}

.feedback-text-box {
  border-radius: 8px;
  padding: 10px 14px;
}

.feedback-text-box p {
  margin: 0;
  font-size: 14px;
  color: #334155;
}

.review-actions {
  display: flex;
  align-items: center;
  gap: 10px;
}

.action-btn {
  background-color: #1B75D2;
  border: none;
  border-radius: 50px;
  padding: 8px 12px;
  display: inline-flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  transition: opacity 0.2s ease;
}

.action-btn:hover {
  opacity: 0.9;
}

.icon-circle {
  background-color: rgba(255, 255, 255, 0.2);
  width: 22px;
  height: 22px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
}

.btn-count {
  color: #ffffff;
  font-size: 14px;
  font-weight: 700;
  padding-right: 4px;
}

.report-btn {
  background-color: rgba(248, 250, 252, 0.8);
  border: 1px solid #cbd5e1;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 13px;
  color: #475569;
  cursor: pointer;
}

.report-btn:hover {
  background-color: #f1f5f9;
}
</style>