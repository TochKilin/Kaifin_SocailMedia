
<script setup>
import { ref, computed, onMounted } from 'vue'
import NavBar from '../navbar/NavBar.vue'
import CourseDetaile from './CourseDetaile.vue'
import AIAssistant from './AIAssistant.vue'

const BASE_URL = 'http://localhost:7070'

function authHeaders() {
  const token = localStorage.getItem('token')
  return token ? { Authorization: `Bearer ${token}` } : {}
}

const selectedCourse = ref(null)

const handleOpenDetail = (course) => {
  selectedCourse.value = course
}

const handleBack = () => {
  selectedCourse.value = null
}


const aiQuery = ref('')
const aiMessages = ref([
  { sender: 'ai', text: 'សួស្តី! ខ្ញុំជា AI Assistant របស់អ្នក។ តើមានចម្ងល់អ្វីទាក់ទងនឹងមេរៀននេះដែរឬទេ?' }
])

const sendAiMessage = () => {
  if (!aiQuery.value.trim()) return
  aiMessages.value.push({ sender: 'user', text: aiQuery.value })
  const question = aiQuery.value
  aiQuery.value = ''

  setTimeout(() => {
    aiMessages.value.push({ sender: 'ai', text: `ចំពោះសំណួរ "${question}" នេះ គឺទាក់ទងនឹងស្ថាបត្យកម្មប្រព័ន្ធនៃវគ្គសិក្សានេះ។` })
  }, 1000)
}

const activeTab = ref('all')
const searchQuery = ref('')
const courses = ref([])
const loading = ref(false)
const errorMsg = ref('')


const resolveMediaUrl = (raw) => {
  if (!raw) return 'https://images.unsplash.com/photo-1522071820081-009f0129c71c?auto=format&fit=crop&w=1200&q=80'
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

const normalizeCourse = (c) => ({
  id: c.id,
  enrollmentId: c.enrollment_id,
  type: c.type,
  title: c.title,
  description: c.description,
  instructor: c.instructor,
  level: String(c.level ?? ''),
  price: String(c.price ?? '0'),
  oldPrice: c.old_price != null ? String(c.old_price) : null,
  promo: c.promo,
  mediaUrl: resolveMediaUrl(c.media_url),
    videoUrl: c.preview_video_url || '',
  progress: c.progress ?? 0,
  enrolledAt: c.enrolled_at,
  completedAt: c.completed_at,
  certificateUrl: c.certificate_url,
  sections: []
})


const fetchEnrolledCourses = async () => {
  loading.value = true
  errorMsg.value = ''
  try {
    const token = localStorage.getItem('token')

    const res = await fetch(`${BASE_URL}/api/v1/front/course-enrollments/show?page=1&limit=50`, {
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    })

    if (res.status === 401) {
      throw new Error('Expire jwt — please login again')
    }
    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }

    const json = await res.json()
    const data = json?.data ?? json
    const list = data?.courses ?? []
    courses.value = list.map(normalizeCourse)
  } catch (e) {
    console.error('Failed to fetch enrolled courses', e)
    errorMsg.value = e.message?.includes('Expire jwt')
      ? 'Session ផុតកំណត់ សូម login ម្តងទៀត'
      : 'មិនអាចទាញកម្មវិធីសិក្សាដែលបានទិញបានទេ'
  } finally {
    loading.value = false
  }
}

const filteredCourses = computed(() => {
  if (!searchQuery.value.trim()) return courses.value
  const q = searchQuery.value.toLowerCase()
  return courses.value.filter(
    (c) =>
      c.title.toLowerCase().includes(q) ||
      c.description.toLowerCase().includes(q) ||
      c.instructor.toLowerCase().includes(q)
  )
})

onMounted(() => {
  fetchEnrolledCourses()
})
</script>
<template>
  <div class="page-wrapper">
    <NavBar />
    <CourseDetaile 
      v-if="selectedCourse" 
      :course="selectedCourse" 
      @back="handleBack" 
    />

    <div v-else>
      <div class="learning-app-layout">
        <aside class="sidebar-left">
          <div class="sidebar-widget search-widget">
            <div class="search-box">
              <svg class="search-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
              </svg>
              <input type="text" v-model="searchQuery" placeholder="Search Course....." class="search-input" />
            </div>
          </div>

          <div class="sidebar-widget tabs-widget">
            <h3 class="widget-title">Navigation</h3>
            <div class="tabs-container">
              <button 
                @click="activeTab = 'all'" 
                :class="['tab-btn', { active: activeTab === 'all' }]">
                <span>All Course</span>
              </button>
              <button 
                @click="activeTab = 'save'" 
                :class="['tab-btn', { active: activeTab === 'save' }]">
                <span>Course Save</span>
              </button>
              <button 
                @click="activeTab = 'ai-assistant'" 
                :class="['tab-btn', { active: activeTab === 'ai-assistant' }]">
                <span>AI Assistant</span>
              </button>
            </div>
          </div>

          <div class="sidebar-widget filter-widget">
            <h3 class="widget-title">Filter Options</h3>
            <div class="filter-buttons">
              <button class="filter-btn">
                <svg class="btn-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h7"/></svg>
                Categories
              </button>
              <button class="filter-btn">
                <svg class="btn-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"/></svg>
                Progress
              </button>
              <button class="filter-btn">
                <svg class="btn-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                Instructor
              </button>
            </div>
          </div>

        </aside>

        <main class="main-content-right">
          <div v-if="activeTab === 'ai-assistant'" class="ai-tab-view">
            <h2 class="section-title">AI Learning Assistant</h2>
            <AIAssistant 
              v-model:aiQuery="aiQuery"
              :aiMessages="aiMessages"
              @send="sendAiMessage"
            />
          </div>

          <template v-else>
            <div class="banner-box">
              <img src="https://images.unsplash.com/photo-1522071820081-009f0129c71c?auto=format&fit=crop&w=1200&q=80" alt="Banner" class="banner-image" />
            </div>
            <h2 class="section-title">My Learning</h2>
            <div v-if="loading" class="state-msg">Waiting...</div>
            <div v-else-if="errorMsg" class="state-msg error">{{ errorMsg }}</div>
            <div v-else-if="filteredCourses.length === 0" class="state-msg">
              {{ searchQuery ? 'not found' : 'your are no buy course' }}
            </div>

            <!-- Course Cards Grid -->
            <div class="courses-grid">
              <div 
                v-for="course in filteredCourses" 
                :key="course.id" 
                class="course-card"
                @click="handleOpenDetail(course)"
              >
                
                <!-- Media Preview Box -->
                <div class="media-box">
                  <template v-if="course.type === 'video'">
                    <div class="video-preview-wrapper">
                      <img :src="course.mediaUrl" alt="Video thumbnail" class="media-content" />
                      <div class="play-button-overlay">
                        <svg class="play-icon" fill="currentColor" viewBox="0 0 24 24">
                          <path d="M8 5v14l11-7z"/>
                        </svg>
                      </div>
                    </div>
                  </template>
                  <template v-else-if="course.type === 'book'">
                    <img :src="course.mediaUrl" alt="Course thumbnail" class="media-content" />
                    <div class="type-badge-overlay">
                      <svg class="type-icon-svg" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/></svg>
                      <span>E-Book</span>
                    </div>
                  </template>
                  <template v-else>
                    <img :src="course.mediaUrl" alt="Course thumbnail" class="media-content" />
                    <div class="type-badge-overlay">
                      <svg class="type-icon-svg" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
                      <span>Image Course</span>
                    </div>
                  </template>
                </div>

                <!-- Clean Course Title -->
                <h3 class="course-title-modern">{{ course.title }}</h3>
                <!-- Description -->
                <p class="description-text">{{ course.description }}</p>
                <!-- Instructor Info Row -->
                <div class="instructor-row">
                  <div class="instructor-profile">
                    <div class="instructor-avatar-ring">
                      <svg fill="currentColor" viewBox="0 0 24 24"><path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/></svg>
                    </div>
                    <span class="instructor-name">{{ course.instructor }}</span>
                  </div>
                  <span class="level-pill">Lvl {{ course.level }}</span>
                </div>

                <!-- Divider Line -->
                <div class="card-divider"></div>

                <!-- Pricing & Rating Footer -->
                <div class="card-footer-flex">
                  <div class="price-container">
                    <span class="currency">$</span>
                    <span class="current-price">{{ course.price }}</span>
                    <span class="old-price">${{ course.oldPrice }}</span>
                  </div>
                  <div class="rating-stars">
                    <svg class="star-svg" fill="currentColor" viewBox="0 0 24 24"><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/></svg>
                    <svg class="star-svg" fill="currentColor" viewBox="0 0 24 24"><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/></svg>
                    <svg class="star-svg" fill="currentColor" viewBox="0 0 24 24"><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/></svg>
                    <svg class="star-svg" fill="currentColor" viewBox="0 0 24 24"><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/></svg>
                    <svg class="star-svg-empty" fill="currentColor" viewBox="0 0 24 24"><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/></svg>
                  </div>
                </div>

                <!-- Sleek Promo Footer Banner -->
                <div class="promo-banner-pill">
                  <svg class="sparkle-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z"/></svg>
                  <span>{{ course.promo }}</span>
                </div>
              </div>
            </div>
          </template>
        </main>
      </div>
    </div>
  </div>
</template>

<style scoped>
.page-wrapper {
  background: #F7F4F2;
  min-height: 100vh;
}

.learning-app-layout {
  display: grid;
  grid-template-columns: 260px 1fr;
  gap: 16px;
  max-width: 1251px;
  margin: 0 auto;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  box-sizing: border-box;

}

.sidebar-left {
  display: flex;
  flex-direction: column;
  gap: 12px;
  align-self: start;
  position: sticky;
  top: 16px;
  margin-top: 12px;
}

.sidebar-widget {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 12px;
  box-shadow: 0 2px 4px -1px rgba(0, 0, 0, 0.02);
}

.widget-title {
  font-size: 13px;
  font-weight: 700;
  color: #0f172a;
  margin: 0 0 8px 0;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.search-box {
  position: relative;
  width: 100%;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 32px;
  display: flex;
  align-items: center;
  padding: 6px 14px;
  box-sizing: border-box;
}

.search-icon {
  width: 16px;
  height: 16px;
  color: #94a3b8;
  margin-right: 8px;
  flex-shrink: 0;
}

.search-input {
  background: transparent;
  border: none;
  outline: none;
  color: #0f172a;
  font-size: 13.5px;
  width: 100%;
}

.tabs-container {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.tab-btn {
  background: transparent;
  border: none;
  border-radius: 32px;
  padding: 6px 10px;
  color: #64748b;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  text-align: left;
  transition: all 0.2s ease;
  width: 100%;
}

.tab-btn:hover {
  background: #f1f5f9;
  color: #1976D2;
}

.tab-btn.active {
  background: #1976D2;
  color: #ffffff;
  border: none;
  box-shadow: none;
}

.filter-buttons {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.filter-btn {
  background: transparent;
  border: none;
  border-radius: 8px;
  padding: 6px 10px;
  color: #475569;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: all 0.2s ease;
  width: 100%;
}

.btn-icon {
  width: 15px;
  height: 15px;
  color: #64748b;
}

.filter-btn:hover {
  background-color: #f1f5f9;
  color: #1976D2;
}

.main-content-right {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 12px;

}

.ai-tab-view {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.banner-box {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  height: 160px;
  overflow: hidden;
  position: relative;
  box-shadow: 0 4px 15px -3px rgba(25, 118, 210, 0.04);
}

.banner-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.section-title {
  font-size: 20px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}

.courses-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.course-card {
  background: #ffffff;
  border: 1px solid rgba(226, 232, 240, 0.8);
  border-radius: 12px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  position: relative;
  box-shadow: 0 10px 30px -10px rgba(0, 0, 0, 0.04);
  transition: all 0.35s cubic-bezier(0.16, 1, 0.3, 1);
  overflow: hidden;
  cursor: pointer;
}

.course-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 20px 40px -15px rgba(25, 118, 210, 0.12);
}

.media-box {
  border-radius: 12px;
  height: 140px;
  position: relative;
  overflow: hidden;
  background-color: #0f172a;
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.1);
}

.media-content {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s cubic-bezier(0.16, 1, 0.3, 1);
}

.course-card:hover .media-content {
  transform: scale(1.07);
}

.video-preview-wrapper {
  position: relative;
  width: 100%;
  height: 100%;
}

.play-button-overlay {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 40px;
  height: 40px;
  background: rgba(25, 118, 210, 0.9);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8px 20px rgba(25, 118, 210, 0.4);
  transition: transform 0.3s ease;
}

.course-card:hover .play-button-overlay {
  transform: translate(-50%, -50%) scale(1.1);
}

.play-icon {
  width: 18px;
  height: 18px;
  color: #ffffff;
  margin-left: 2px;
}

.type-badge-overlay {
  position: absolute;
  bottom: 8px;
  left: 8px;
  background: rgba(15, 23, 42, 0.75);
  backdrop-filter: blur(6px);
  color: #ffffff;
  padding: 3px 8px;
  border-radius: 6px;
  font-size: 10.5px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 4px;
}

.type-icon-svg {
  width: 12px;
  height: 12px;
}

.course-title-modern {
  font-size: 15px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
  line-height: 1.3;
}

.description-text {
  font-size: 12.5px;
  color: #64748b;
  margin: 0;
  line-height: 1.45;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.instructor-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 2px;
}

.instructor-profile {
  display: flex;
  align-items: center;
  gap: 6px;
}

.instructor-avatar-ring {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: #f1f5f9;
  border: 1px solid #e2e8f0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #475569;
}

.instructor-avatar-ring svg {
  width: 12px;
  height: 12px;
}

.instructor-name {
  font-size: 12px;
  font-weight: 600;
  color: #334155;
}

.level-pill {
  background: #00D262;
  border: 1px solid #ffffff;
  padding: 2px 6px;
  border-radius: 32px;
  font-size: 10.5px;
  font-weight: 600;
  color: #ffffff;
}

.card-divider {
  height: 1px;
  background: #f1f5f9;
  width: 100%;
}

.card-footer-flex {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.price-container {
  display: flex;
  align-items: baseline;
  gap: 2px;
}

.currency {
  font-size: 14px;
  font-weight: 800;
  color: #ef4444;
}

.current-price {
  font-size: 18px;
  font-weight: 800;
  color: #ef4444;
  letter-spacing: -0.5px;


  text-shadow:
    1px 1px 0 #9f2c2c,
    2px 2px 3px rgba(0, 0, 0, 0.18);
}

.old-price {
  font-size: 12px;
  font-weight: 600;
  color: #94a3b8;
  text-decoration: line-through;
  margin-left: 3px;
}

.rating-stars {
  display: flex;
  align-items: center;
  gap: 1px;
  color: #f59e0b;
}

.star-svg {
  width: 12px;
  height: 12px;
}

.star-svg-empty {
  width: 12px;
  height: 12px;
  color: #cbd5e1;
}

.promo-banner-pill {
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 6px 10px;
  font-size: 11.5px;
  color: #475569;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 6px;
}

.sparkle-icon {
  width: 12px;
  height: 12px;
  color: #1976D2;
  flex-shrink: 0;
}

.state-msg {
  text-align: center;
  padding: 40px 0;
  color: #64748b;
  font-size: 14px;
}
.state-msg.error {
  color: #ef4444;
}
</style>