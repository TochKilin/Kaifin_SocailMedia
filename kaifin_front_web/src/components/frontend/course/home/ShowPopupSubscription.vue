<template>
  <div class="popup-overlay" @click.self="$emit('close')">
    <div class="popup-card">
      
      <!-- Header Bar: Text Label នៅខាងឆ្វេង និង Close Button នៅខាងស្ដាំ -->
      <div class="popup-header-bar">
        <span class="header-label">Subscription</span>
        <button class="close-btn" @click="$emit('close')">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>

      <!-- Top Video / Thumbnail Section -->
      <div class="popup-video-box">
        <img :src="thumbnail" alt="Course Preview" class="popup-thumb-img" />
        <div class="play-overlay">
          <svg viewBox="0 0 24 24" fill="currentColor">
            <polygon points="5 3 19 12 5 21 5 3"></polygon>
          </svg>
        </div>
      </div>

      <!-- Option 1: Subscribe & Save Box -->
      <div class="plan-option-box">
        <div class="radio-circle">
          <div class="radio-inner"></div>
        </div>
        <div class="plan-text-group">
          <span class="plan-badge">SUBSCRIBE AND SAVE</span>
          <span class="plan-price">From $10/month</span>
        </div>
      </div>

      <!-- Benefit 1: Free Course Offer -->
      <div class="benefit-row">
        <div class="benefit-icon-wrap">
          <svg viewBox="0 0 24 24" fill="#f59e0b" class="badge-icon">
            <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"></path>
          </svg>
        </div>
        <span class="benefit-text">Get this course for free when you subscribe. Terms apply.</span>
      </div>

      <!-- Benefit 2: Access to courses -->
      <div class="benefit-row">
        <div class="benefit-icon-wrap">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feature-icon">
            <rect x="2" y="3" width="20" height="14" rx="2" ry="2"></rect>
            <line x1="8" y1="21" x2="16" y2="21"></line>
            <line x1="12" y1="17" x2="12" y2="21"></line>
          </svg>
        </div>
        <span class="benefit-text">Access to 28,000+ top-rated courses</span>
      </div>

      <!-- Benefit 3: More course -->
      <div class="benefit-row clickable-row" @click="goToMoreCourse">
        <div class="benefit-icon-wrap">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feature-icon">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"></path>
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path>
          </svg>
        </div>
        <span class="benefit-text highlight-blue">More course</span>
      </div>

      <!-- Action Button: ពេលចុចឱ្យរត់ទៅកាន់ PaymentMethod -->
      <button class="subscription-submit-btn" @click="goToPayment">
        Subscription Now
      </button>

    </div>
  </div>
</template>

<script setup>
import { onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

defineProps({
  thumbnail: {
    type: String,
    default: 'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=600&auto=format&fit=crop&q=60'
  }
})

defineEmits(['close'])

// មុខងារពេលចុចលើ More course
const goToMoreCourse = () => {
  router.push({ name: 'MoreCourseInstructor' })
}

// មុខងារពេលចុចលើប៊ូតុង Subscription Now ដើម្បីទៅកាន់ PaymentMethod
const goToPayment = () => {
  router.push({ name: 'PaymentMethod' })
}

onMounted(() => {
  document.body.style.overflow = 'hidden'
})

onUnmounted(() => {
  document.body.style.overflow = 'auto'
})
</script>

<style scoped>
.popup-overlay {
  position: absolute;
  top: 65px;
  left: 0;
  width: 100%;
  height: calc(100vh - 65px);

  background: rgba(13, 13, 14, 0.253);

  display: flex;
  align-items: center;
  justify-content: center;

  z-index: 100;

  padding: 20px 16px;
  box-sizing: border-box;
}

.popup-card {
  width: 100%;
  max-width: 740px;

  /* ឱ្យវាវែង */
  height: calc(100vh - 90px);

  background: #ffffff;
  border-radius: 12px;
  padding: 20px;

  display: flex;
  flex-direction: column;
  gap: 16px;

  position: relative;
  animation: popupFadeIn 0.25s ease-out;

  overflow-y: auto;
  box-sizing: border-box;
}

@keyframes popupFadeIn {
  from {
    opacity: 0;
    transform: scale(0.95) translateY(10px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

.popup-header-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.header-label {
  font-size: 22px;
  font-weight: 700;
  color: #0f172a;
}

.close-btn {
  background: #ffffff;
  border: 1.5px solid #cbd5e1;
  width: 38px;
  height: 38px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #475569;
  transition: all 0.2s;
  box-shadow: none;
}

.close-btn:hover {
  background: #f8fafc;
  border-color: #94a3b8;
  color: #0f172a;
}

.close-btn svg {
  width: 18px;
  height: 18px;
}

.popup-video-box {
  width: 100%;
  height: 280px;
  border-radius: 16px;
  overflow: hidden;
  position: relative;
  background: #0f172a;
}

.popup-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.play-overlay {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 56px;
  height: 56px;
  background: #1B75D2;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  box-shadow: 0 4px 12px rgba(27, 117, 210, 0.4);
}

.play-overlay svg {
  width: 24px;
  height: 24px;
  margin-left: 2px;
}

.plan-option-box {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 4px 0;
  border: none;
  background: transparent;
}

.radio-circle {
  width: 22px;
  height: 22px;
  border: 2px solid #1B75D2;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.radio-inner {
  width: 12px;
  height: 12px;
  background: #1B75D2;
  border-radius: 50%;
}

.plan-text-group {
  display: flex;
  flex-direction: column;
}

.plan-badge {
  font-size: 11px;
  font-weight: 700;
  color: #1B75D2;
  letter-spacing: 0.5px;
}

.plan-price {
  font-size: 16px;
  font-weight: 800;
  color: #0f172a;
}

.benefit-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 4px 0;
  background: transparent;
  border: none;
}

.benefit-icon-wrap {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.badge-icon, .feature-icon {
  width: 20px;
  height: 20px;
  color: #1B75D2;
}

.benefit-text {
  font-size: 14px;
  font-weight: 600;
  color: #334155;
}

.highlight-blue {
  color: #1B75D2;
}

.subscription-submit-btn {
  width: 240px;
  margin: auto;
  background: #1B75D2;
  color: #ffffff;
  border: none;
  border-radius: 32px;
  padding: 14px 20px;
  font-size: 16px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: none;
  margin-top: 6px;
}

.subscription-submit-btn:hover {
  background: #155fae;
  transform: translateY(-1px);
}

.clickable-row {
  cursor: pointer;
  padding: 6px 8px;
  border-radius: 8px;
  transition: background-color 0.2s;
}

.clickable-row:hover {
  background-color: #f1f5f9;
}
</style>