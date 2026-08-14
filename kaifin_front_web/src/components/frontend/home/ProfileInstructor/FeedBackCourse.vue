<script setup>
import { ref, computed } from 'vue';

const props = defineProps({
  feedbacks: {
    type: Array,
    default: () => [
      {
        courseName: 'Vue 3 Masterclass',
        thumbnail: 'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=100&auto=format&fit=crop&q=80',
        rating: '4 ★',
        time: '1 day ago',
        comment: 'The explanation is very clear and practical for real-world projects.'
      },
      {
        courseName: 'Go Fiber Advance course',
        thumbnail: 'https://images.unsplash.com/photo-1627398242454-45a1465c2479?w=100&auto=format&fit=crop&q=80',
        rating: '5 ★',
        time: '1 day ago',
        comment: 'Great course! Learned a lot about building fast REST APIs with Go.'
      },
      {
        courseName: 'Full-Stack Development',
        thumbnail: 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?w=100&auto=format&fit=crop&q=80',
        rating: '4 ★',
        time: '2 days ago',
        comment: 'Very detailed and comprehensive step-by-step guidance.'
      },
      {
        courseName: 'UI/UX Design Basics',
        thumbnail: 'https://images.unsplash.com/photo-1581291518633-83b4ebd1d83e?w=100&auto=format&fit=crop&q=80',
        rating: '5 ★',
        time: '3 days ago',
        comment: 'Awesome design concepts and very helpful Figma tips!'
      },
      {
        courseName: 'Advanced TypeScript',
        thumbnail: 'https://images.unsplash.com/photo-1517694712202-14dd9538aa97?w=100&auto=format&fit=crop&q=80',
        rating: '4 ★',
        time: '4 days ago',
        comment: 'Helped me understand generics and utility types much better.'
      },
      {
        courseName: 'Docker & Kubernetes',
        thumbnail: 'https://images.unsplash.com/photo-1605745341112-85968b19335b?w=100&auto=format&fit=crop&q=80',
        rating: '5 ★',
        time: '5 days ago',
        comment: 'Containerization made super easy. Highly recommended for devs!'
      },
      {
        courseName: 'ASP.NET Core API',
        thumbnail: 'https://images.unsplash.com/photo-1516116216624-53e697fedbea?w=100&auto=format&fit=crop&q=80',
        rating: '5 ★',
        time: '1 week ago',
        comment: 'Clean architecture explanation was top-notch. Loved it.'
      }
    ]
  }
});

const showAll = ref(false);
const displayedFeedbacks = computed(() => {
  return showAll.value ? props.feedbacks : props.feedbacks.slice(0, 2);
});

const toggleShowMore = () => {
  showAll.value = !showAll.value;
};
</script>

<template>
  <div class="feedback-container">
    <!-- Header Stats Table Style (with Attractive Rating Badge & User Avatars) -->
    <div class="feedback-header-stats-table">
      <div class="stat-cell">
        <!-- Hot Rating Badge -->
        <div class="rating-badge-hot">
          <span class="star-icon">⭐</span>
          <span class="rating-text">4.7 course rating</span>
        </div>
      </div>
      <div class="stat-divider"></div>
      <div class="stat-cell ratings-cell-right">
        <span>469K ratings</span>
        <!-- User Profile Avatars Stack with Plus Badge -->
        <div class="header-user-avatars">
          <img src="https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100&auto=format&fit=crop&q=80" alt="User 1" class="header-avatar" />
          <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&auto=format&fit=crop&q=80" alt="User 2" class="header-avatar" />
          <img src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=100&auto=format&fit=crop&q=80" alt="User 3" class="header-avatar" />
          <div class="avatar-plus-badge">+</div>
        </div>
      </div>
    </div>

    <!-- Feedback Cards Grid -->
    <div class="feedback-grid">
      <div v-for="(item, index) in displayedFeedbacks" :key="index" class="feedback-card">
        <div class="feedback-card-top">
          <!-- Course Thumbnail Image -->
          <div class="course-thumbnail-wrapper">
            <img :src="item.thumbnail" :alt="item.courseName" class="course-thumbnail-img" />
          </div>
          <div class="feedback-course-info">
            <div class="course-name-box">{{ item.courseName }}</div>
            <div class="rating-time-row">
              <span class="rating-pill-box">{{ item.rating }}</span>
              <span class="time-pill-box">
                <!-- Clock Icon -->
                <svg class="clock-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                {{ item.time }}
              </span>
            </div>
          </div>
        </div>
        
        <div class="feedback-text-box">
          <p>{{ item.comment }}</p>
        </div>
      </div>
    </div>

    <!-- See More Button -->
    <div class="see-more-wrapper">
      <button @click="toggleShowMore" class="see-more-feedback-btn">
        <span>{{ showAll ? 'Show less' : 'See more' }}</span>
        <svg class="down-arrow-icon" :class="{ 'rotate-icon': showAll }" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <line x1="12" y1="5" x2="12" y2="19"></line>
          <polyline points="19 12 12 19 5 12"></polyline>
        </svg>
      </button>
    </div>
  </div>
</template>

<style scoped>
.feedback-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  background-color: #ffffff;
  border-radius: 24px;
}

/* Header Stats Table Style (Border and Shadow Intact) */
.feedback-header-stats-table {
  display: flex;
  align-items: center;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  background-color: transparent;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
  overflow: hidden;
  margin-bottom: 5px;
}

.stat-cell {
  flex: 1;
  padding: 10px 16px;
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Hot Rating Badge Style */
.rating-badge-hot {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background-color: #fff1f2c3; /* Soft light red/orange background */
  padding: 6px 14px;
  border-radius: 32px;
}

.rating-text {
  color: #e11d48; /* Vibrant red text */
  font-weight: 700;
  font-size: 14px;
}

.star-icon {
  font-size: 15px;
}

.fire-icon {
  font-size: 15px;
}

.ratings-cell-right {
  justify-content: space-between;
}

.header-user-avatars {
  display: flex;
  align-items: center;
  margin-left: 8px;
}

.header-avatar {
  width: 26px;
  height: 26px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #ffffff;
  margin-left: -10px;
}

.header-avatar:first-child {
  margin-left: 0;
}

/* Plus Badge Style */
.avatar-plus-badge {
  width: 26px;
  height: 26px;
  border-radius: 50%;
  background-color: #f3f4f6;
  border: 2px solid #ffffff;
  color: #4b5563;
  font-size: 12px;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: -10px;
}

.stat-divider {
  width: 1px;
  height: 24px;
  background-color: #e5e7eb;
}

.feedback-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

@media (max-width: 768px) {
  .feedback-grid {
    grid-template-columns: 1fr;
  }
}

.feedback-card {
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 14px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
  transition: all 0.3s ease;
}

.feedback-card:hover {
  box-shadow: 0 6px 20px rgba(59, 130, 246, 0.06);
  transform: translateY(-2px);
}

.feedback-card-top {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* Course Thumbnail Styles */
.course-thumbnail-wrapper {
  width: 52px;
  height: 52px;
  border-radius: 50%;
  overflow: hidden;
  border: 1px solid #e5e7eb;
  flex-shrink: 0;
  background-color: #f3f4f6;
}

.course-thumbnail-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.feedback-course-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex-grow: 1;
  overflow: hidden;
}

.course-name-box {
  font-size: 14px;
  font-weight: 600;
  color: #111827;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.rating-time-row {
  display: flex;
  gap: 8px;
  align-items: center;
}

.rating-pill-box, .time-pill-box {
  border-radius: 32px;
  padding: 2px 8px;
  font-size: 11px;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 4px;
  border: 1px solid #e5e7eb;
}

.rating-pill-box {
  background-color: #1B75D2;
  color: #ffffff;
  border-color: #1B75D2;
}

.time-pill-box {
  background-color: #f9fafb;
  color: #6b7280;
  border-color: #e5e7eb;
}

.clock-icon {
  width: 12px;
  height: 12px;
}

.feedback-text-box {
  background-color: #f9fafb;
  border-radius: 12px;
  padding: 14px;
  min-height: 60px;
}

.feedback-text-box p {
  margin: 0;
  font-size: 13px;
  color: #4b5563;
  line-height: 1.4;
}

.see-more-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 5px;
}

.see-more-feedback-btn {
  background-color: #9c9a9a41;
  border: none;
  border-radius: 32px;
  padding: 10px 30px;
  font-size: 14px;
  font-weight: 600;
  color: #000;
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
}

.see-more-feedback-btn:hover {
  background-color: #f9fafb;
  border-color: #3b82f6;
  color: #2563eb;
  transform: translateY(-2px);
}

.down-arrow-icon {
  width: 16px;
  height: 16px;
  transition: transform 0.3s ease;
}

.rotate-icon {
  transform: rotate(180deg);
}
</style>