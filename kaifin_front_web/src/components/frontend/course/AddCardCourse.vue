<template>
  <div class="course-detail-page">
    <NavBar />
    
    <div class="page-wrapper">
      <div class="course-card">
        
        <!-- Header Section: Thumbnail ឆ្វេង & ព័ត៌មានស្តាំ -->
        <div class="course-header-row">
          <div class="course-thumbnail-container">
            <img 
              :src="course?.image || defaultCourseImage" 
              alt="Course Thumbnail" 
              class="course-thumb-img" 
            />
            <!-- 🌟 បន្ថែម margin-top និងកែតម្រូវទីតាំងកុំឱ្យវាឡើងទៅប៉ះ Navbar ផ្នែកខាងលើ -->
            <div class="thumbnail-new-badge">
              <span>New product</span>
            </div>
          </div>
          
          <div class="course-header-info">
            <div class="title-with-badge">
              <h2 class="course-title">{{ course?.name || 'Advanced ASP.NET Core & Vue.js Architecture' }}</h2>
            </div>
            
            <div class="author-row">
              <div class="author-avatar-box">
                <span class="author-emoji">💻</span>
              </div>
              <span class="author-name">OTres Technology</span>
              <span class="level-badge">LV.5</span>
            </div>

            <!-- Meta Info (Last update & Hours) -->
            <div class="meta-info-row">
              <div class="meta-badge">
                <svg class="meta-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                  <line x1="16" y1="2" x2="16" y2="6"></line>
                  <line x1="8" y1="2" x2="8" y2="6"></line>
                  <line x1="3" y1="10" x2="21" y2="10"></line>
                </svg>
                <span>Last update {{ course?.lastUpdate || '2026' }}</span>
              </div>
              <div class="meta-badge">
                <svg class="meta-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                <span>{{ course?.hours || '80 Hours' }}</span>
              </div>
            </div>

            <!-- តម្លៃពណ៌ខៀវ -->
            <div class="price-display">
              <span class="currency-symbol">$</span><span class="price-number">49.99</span>
            </div>
          </div>
        </div>

        <!-- Container សម្រាប់ការពិពណ៌នា និងអត្ថប្រយោជន៍ -->
        <div class="content-top-group">
          <div class="description-box">
            <p class="description-text">
              {{ course?.description || 'Comprehensive guide to building scalable retail management backend APIs and high-performance Vue component systems.' }}
            </p>
          </div>

          <div class="benefits-list">
            <div class="benefit-item" v-for="(benefit, index) in (course?.benefits || defaultBenefits)" :key="index">
              <div class="benefit-icon-box">
                <svg class="sticker-face-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <rect x="3" y="3" width="18" height="18" rx="5" fill="#1B75D2" stroke="#60a5fa" stroke-width="1.5"/>
                  <circle cx="8.5" cy="9" r="2" fill="#ffffff"/>
                  <circle cx="15.5" cy="9" r="2" fill="#ffffff"/>
                  <rect x="9.5" y="14" width="5" height="4" rx="1" fill="#ffffff"/>
                </svg>
              </div>
              <span class="benefit-text">{{ benefit }}</span>
            </div>
          </div>
        </div>

        <!-- Footer Action (Add Cart & Save) -->
        <div class="action-footer">
          <button class="add-cart-btn" @click="addToCart">
            <span>Add Cart</span>
            <svg class="cart-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="9" cy="21" r="1"></circle>
              <circle cx="20" cy="21" r="1"></circle>
              <path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"></path>
            </svg>
          </button>
          <button class="favorite-icon-btn" title="Save">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
            </svg>
          </button>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import NavBar from '../navbar/NavBar.vue'

const props = defineProps({
  courseData: {
    type: Object,
    default: null
  }
})

const route = useRoute()
const router = useRouter()

const defaultCourseImage = 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?auto=format&fit=crop&w=800&q=80'

const mockCourseDatabase = {
  1: {
    name: 'Advanced ASP.NET Core & Vue.js Architecture',
    image: 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?auto=format&fit=crop&w=800&q=80',
    lastUpdate: '2026',
    hours: '80 Hours',
    price: '$49.99',
    description: 'Comprehensive guide to building scalable retail management backend APIs and high-performance Vue component systems.',
    benefits: [
      'Master database integration with MySQL & phpMyAdmin schemas',
      'Build robust RESTful controllers using C# ASP.NET Core',
      'Develop reactive frontend components using Vue.js'
    ]
  }
}

const course = ref(props.courseData)

const defaultBenefits = [
  'Master database integration with MySQL & phpMyAdmin schemas',
  'Build robust RESTful controllers using C# ASP.NET Core',
  'Develop reactive frontend components using Vue.js'
]

onMounted(() => {
  if (!course.value && history.state && history.state.courseData) {
    course.value = history.state.courseData
  } else if (!course.value) {
    const courseId = route.params.id || 1
    course.value = mockCourseDatabase[courseId] || mockCourseDatabase[1]
  }
})

const addToCart = () => {
  router.push({
    name: 'ShoppingCartCourse',
    state: {
      cartItem: course.value
    }
  })
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@500;600;700;800&display=swap');

.course-detail-page {
  width: 100vw;
  height: 100vh;
  background-color: #F7F4F2;
  font-family: 'Plus Jakarta Sans', sans-serif;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.page-wrapper {
  width: 100%;
  flex: 1;
  display: flex;
  align-items: stretch;
  justify-content: center;
  background-color: #ffffff;
  padding: 0;
  box-sizing: border-box;
  overflow: hidden;
}

.course-card {
  position: relative;
  width: 780px;
  max-width: 100%;
  height: 100%;
  background-color: #ffffff;
  border-left: 1px solid #e2e8f0;
  border-right: 1px solid #e2e8f0;
  border-radius: 0;
  /* 🌟 បន្ថែម padding-top ធំល្មម ដើម្បីបង្កើតគម្លាតសុវត្ថិភាពពី Navbar កុំឱ្យ Badge ឡើងទៅប៉ះ */
  padding: 35px 28px 20px 28px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  gap: 16px;
  box-shadow: none;
  overflow-y: auto;
}

.course-header-row {
  display: flex;
  gap: 20px;
  align-items: flex-start;
  width: 100%;
}

.course-thumbnail-container {
  position: relative;
  width: 155px;
  height: 175px;
  border-radius: 12px;
  overflow: visible;
  background-color: #f1f5f9;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.course-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 12px;
}

/* 🌟 រក្សាតម្លៃ top ឱ្យនៅខាងក្នុង Thumbnail មិនឱ្យរត់ហៀរឡើងលើទៅប៉ះ Navbar ទេ។ */
.thumbnail-new-badge {
  position: absolute;
  top: -10px;
  right: -16px;
  background-color: #f43f5e;
  color: #ffffff;
  font-size: 11px;
  font-weight: 700;
  padding: 5px 12px;
  border-radius: 8px;
  letter-spacing: 0.3px;
  box-shadow: none;
  z-index: 10;
  border: 2px solid #ffffff;
  white-space: nowrap;
  transform: rotate(15deg); 
  transform-origin: center;
}

.course-header-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.title-with-badge {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 8px;
}

.course-title {
  font-size: 20px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
  line-height: 1.35;
}

.author-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.author-avatar-box {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background-color: #f1f5f9;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
}

.author-name {
  font-size: 13px;
  font-weight: 700;
  color: #334155;
}

.level-badge {
  background-color: #1B75D2;
  color: #ffffff;
  font-size: 10px;
  font-weight: 800;
  padding: 2px 7px;
  border-radius: 6px;
  letter-spacing: 0.5px;
}

.meta-info-row {
  display: flex;
  gap: 16px;
  margin-top: 2px;
}

.meta-badge {
  display: flex;
  align-items: center;
  gap: 6px;
  background-color: transparent; 
  color: #475569;
  font-size: 13px;
  font-weight: 700;
  padding: 0;
}

.meta-icon {
  width: 16px;
  height: 16px;
  stroke: #475569;
  flex-shrink: 0;
}

.price-display {
  display: flex;
  align-items: flex-start;
  color: #ef4444;
  margin-top: 4px;
}

.currency-symbol {
  font-size: 16px;
  font-weight: 800;
  line-height: 1.2;
  margin-right: 1px;
}

.price-number {
  font-size: 28px;
  font-weight: 800;
  line-height: 1;
}

.content-top-group {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 0;
}

.description-box {
  width: 100%;
  padding: 0;
  background-color: #ffffff;
  box-sizing: border-box;
}

.description-text {
  font-size: 13px;
  color: #64748b;
  margin: 0;
  line-height: 1.5;
}

.benefits-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.benefit-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 2px 0;
  background-color: #ffffff;
  border-radius: 12px;
  box-sizing: border-box;
}

.benefit-icon-box {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.sticker-face-icon {
  width: 28px;
  height: 28px;
}

.benefit-text {
  font-size: 14px;
  font-weight: 700;
  color: #334155;
}

.action-footer {
  display: flex;
  gap: 10px;
  align-items: center;
  margin-top: auto;
  padding-top: 10px;
}

.add-cart-btn {
  flex: 1;
  height: 48px;
  background-color: #1B75D2;
  border: none;
  border-radius: 28px;
  cursor: pointer;
  padding: 0 24px;
  transition: background-color 0.2s ease, transform 0.1s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
}

.add-cart-btn span {
  font-size: 15px;
  font-weight: 700;
  color: #ffffff;
}

.cart-icon {
  width: 20px;
  height: 20px;
  stroke: #ffffff;
  flex-shrink: 0;
}

.add-cart-btn:hover {
  background-color: #155fa0;
}

.add-cart-btn:active {
  transform: scale(0.98);
}

.favorite-icon-btn {
  width: 48px;
  height: 48px;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 28px;
  color: #64748b;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s ease;
  flex-shrink: 0;
}

.favorite-icon-btn svg {
  width: 20px;
  height: 20px;
  stroke: #64748b;
  fill: none;
  transition: all 0.2s ease;
}

.favorite-icon-btn:hover {
  background-color: #fee2e2;
  border-color: #fca5a5;
  color: #ef4444;
}

.favorite-icon-btn:hover svg {
  stroke: #ef4444;
  fill: #ef4444;
}
</style>