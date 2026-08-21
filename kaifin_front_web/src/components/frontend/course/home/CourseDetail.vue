<template>
  <div class="course-detail-page">
    <NavBar />

    <div class="page-container">
      <!-- Main Content Area -->
      <div class="main-content">
        
        <!-- Top Hero -->
        <div class="hero-section">
          <!-- Row-->
          <div class="hero-top-row">
            <div class="hero-left-thumb">
              <img :src="course.thumbnail" alt="Course Thumbnail" class="hero-thumb-img" />
              <div class="play-overlay-small">
                <svg viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
              </div>
            </div>

            <div class="hero-right-content">
              <h1 class="course-title">{{ course.title }}</h1>
              
              <div class="pricing-box">
                <span class="current-price">${{ course.currentPrice }}</span>
                <span class="original-price">${{ course.originalPrice }}</span>
                <span class="discount-tag">83% off</span>
              </div>
              
              <div class="urgency-banner">
                <svg class="clock-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                <span>1 day left this price!</span>
              </div>
            </div>
          </div>

          <!-- Row 2-->
          <div class="hero-creator-row">
            <span class="create-by-label">Create BY:</span>
            <span class="instructor-box">{{ course.creator }}</span>
          </div>

          <!-- Row 3: Last update & Language -->
          <div class="hero-meta-row">
            <div class="meta-card">
              <svg class="meta-icon-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
              <span>Last update <span class="highlight-val">{{ course.lastUpdated }}</span></span>
            </div>
            <div class="meta-card">
              <svg class="meta-icon-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="2" y1="12" x2="22" y2="12"></line><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path></svg>
              <span>{{ course.language }}</span>
            </div>
          </div>

          <!-- Row 4: -->
          <div class="hero-stats-row">
            <!-- Rating Group -->
            <div class="stat-item">
              <svg class="star-icon" viewBox="0 0 24 24" fill="currentColor"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon></svg>
              <div class="stat-content">
                <span class="stat-val">{{ course.rating }}</span>
                <span class="small-label">Rating</span>
              </div>
            </div>

            <!-- Reviews Group -->
            <div class="stat-item">
              <svg class="review-icon-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
              <div class="stat-content">
                <span class="stat-val">{{ course.ratingsCount.toLocaleString() }}</span>
                <span class="small-label">Reviews</span>
              </div>
            </div>

            <!-- Students Group -->
            <div class="stat-item">
              <svg class="inline-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
              <div class="stat-content">
                <span class="stat-val">{{ course.studentsCount.toLocaleString() }}</span>
                <span class="small-label">Students</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Mobile/Inline Purchase Card -->
        <div class="mobile-purchase-card">
          <div class="preview-thumb-box">
            <img :src="course.thumbnail" alt="Course Preview" class="preview-img" />
            <div class="play-overlay">
              <svg viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
            </div>
          </div>
          <div class="pricing-box">
            <span class="current-price">${{ course.currentPrice }}</span>
            <span class="original-price">${{ course.originalPrice }}</span>
            <span class="discount-tag">83% off</span>
          </div>
          <div class="urgency-banner">
            <svg class="clock-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
            <span>1 day left at this price!</span>
          </div>
          <button class="subscribe-btn" @click="showSubPopup = true">Start Subscription</button>
        </div>

        <div class="course-includes-section" v-if="course.includes.length">
          <h2>This course includes:</h2>
          <div class="includes-grid">
            <div class="include-item" v-for="inc in course.includes" :key="inc.id">
              <svg v-if="includeIconKey(inc) === 'video'" class="inc-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <polygon points="23 7 16 12 23 17 23 7"></polygon>
          <rect x="1" y="5" width="15" height="14" rx="2" ry="2"></rect>
        </svg>
        <svg v-else-if="includeIconKey(inc) === 'article'" class="inc-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
          <polyline points="14 2 14 8 20 8"></polyline>
          <line x1="16" y1="13" x2="8" y2="13"></line>
          <line x1="16" y1="17" x2="8" y2="17"></line>
          <polyline points="10 9 9 9 8 9"></polyline>
        </svg>
        <svg v-else-if="includeIconKey(inc) === 'mobile'" class="inc-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <rect x="2" y="7" width="20" height="15" rx="2" ry="2"></rect>
          <polyline points="17 2 12 7 7 2"></polyline>
        </svg>
        <svg v-else-if="includeIconKey(inc) === 'exercise'" class="inc-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="16 18 22 12 16 6"></polyline>
          <polyline points="8 6 2 12 8 18"></polyline>
        </svg>
        <svg v-else-if="includeIconKey(inc) === 'download'" class="inc-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
          <polyline points="7 10 12 15 17 10"></polyline>
          <line x1="12" y1="15" x2="12" y2="3"></line>
        </svg>
        <svg v-else-if="includeIconKey(inc) === 'certificate'" class="inc-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="8" r="6"></circle>
          <path d="M15.477 12.89 17 22l-5-3-5 3 1.523-9.11"></path>
        </svg>
        <svg v-else class="inc-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"></circle>
          <path d="M12 8v4M12 16h.01"></path>
        </svg>
            <span>{{ inc.text }}</span>
          </div>
        </div>
      </div>

        <!-- Course Content Section -->
        <div class="course-content-section">
          <div class="content-header-row">
            <h2>Course content</h2>
            <div class="content-stats">
              <span>{{ course.sectionsCount }} sections</span> • 
              <span>{{ course.lecturesCount }} lectures</span> • 
              <span>{{ course.totalLength }} total length</span>
            </div>
          </div>

          <div class="accordion-list">
            <div v-for="section in course.sections" :key="section.id" class="accordion-item" :class="{ 'is-open': section.isOpen }">
              <div class="accordion-summary" @click="toggleSection(section.id)">
                <div class="accordion-title-group">
                  <svg class="chevron-icon" :class="{ rotated: section.isOpen }" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
                  <span class="sec-title">{{ section.title }}</span>
                </div>
                <div class="sec-meta">
                  <span>{{ section.lecturesCount }} lectures</span>
                  <span>{{ section.length }}</span>
                </div>
              </div>
              <div v-if="section.isOpen" class="accordion-body">
                <div v-for="lecture in section.lectures" :key="lecture.id" class="lecture-row">
                  <svg class="play-small-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
                  <span class="lecture-title">{{ lecture.title }}</span>
                  <span class="lecture-duration">{{ lecture.duration }}</span>
                </div>
              </div>
            </div>
          </div>
          <button class="see-more-sections-btn">See more sections</button>
        </div>

        <!-- Instructor Profile Section -->
        <div class="instructor-profile-section">
          <h2>Instructor</h2>
          <div class="instructor-card-detailed">
            <div class="instructor-top-info">
              <img :src="instructor.avatar" alt="Instructor Avatar" class="instructor-lg-avatar" />
              <div class="instructor-meta-col">
                <h3 class="instructor-name">{{ instructor.name }}</h3>
                <p class="instructor-headline">{{ instructor.headline }}</p>
                
                <!-- Instructor Stats Top Row -->
                <div class="instructor-stats-top-row">
                  <div class="stat-card-box">
                    <svg class="giant-star-icon" viewBox="0 0 24 24" fill="currentColor"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon></svg>
                    <span class="stat-num-text">{{ instructor.rating }}</span>
                    <span class="stat-label-text">Rating</span>
                  </div>
                  <div class="stat-card-box">
                    <svg class="stat-icon-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
                    <span class="stat-num-text">{{ instructor.reviewsCount.toLocaleString() }}</span>
                    <span class="stat-label-text">Reviews</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Instructor Stats Bottom Row -->
            <div class="instructor-stats-bottom-row">
              <div class="stat-pill-box">
                <svg class="pill-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
                <span class="stat-pill-val">{{ instructor.studentsCount.toLocaleString() }}</span>
                <span class="stat-pill-label">Students</span>
              </div>
              <div class="stat-pill-box">
                <svg class="pill-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
                <span class="stat-pill-val">{{ instructor.coursesCount }}</span>
                <span class="stat-pill-label">Courses</span>
              </div>
            </div>

            <div class="instructor-bio">
              <p :class="{ 'clamped': !showFullBio }">{{ instructor.description }}</p>
              <button class="toggle-bio-btn" @click="showFullBio = !showFullBio">
                {{ showFullBio ? 'Show less' : 'See more' }}
              </button>
            </div>
          </div>
        </div>

        <!-- Student Feedback / Reviews Section -->
        <div class="student-feedback-section">
          <div class="feedback-header-row">
            <div class="overall-rating-box">
              <span class="big-rating-num">
                <svg class="big-star" viewBox="0 0 24 24" fill="currentColor"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon></svg>
                {{ course.rating }}
              </span>
              <span class="overall-rating-label">course rating</span>
            </div>
            <div class="total-ratings-box">
              <span class="big-rating-num">{{ (course.ratingsCount / 1000).toFixed(0) }}K</span>
              <span class="overall-rating-label">ratings</span>
            </div>
          </div>

          <div class="reviews-grid">
            <div v-for="review in reviews" :key="review.id" class="review-card">
              <div class="reviewer-header">
                <img :src="review.avatar" alt="Reviewer" class="reviewer-avatar" />
                <div class="reviewer-info">
                  <span class="reviewer-name">{{ review.name }}</span>
                  <div class="review-stars-date">
                   <span class="stars">
                      <svg v-for="n in review.rating" :key="n" class="inline-star" viewBox="0 0 24 24" fill="currentColor">
                        <polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon>
                      </svg>
                    </span>
                    <span class="review-date">{{ review.date }}</span>
                  </div>
                </div>
              </div>
              <p class="review-comment">{{ review.comment }}</p>
            </div>
          </div>
          <button class="see-more-reviews-btn">See more reviews</button>
        </div>
      </div>

      <!-- Sticky Sidebar (Right) -->
      <div class="sidebar-container">
        <div class="sticky-purchase-card">
          <div class="preview-thumb-box">
            <img :src="course.thumbnail" alt="Course Preview" class="preview-img" />
            <div class="play-overlay">
              <svg viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
            </div>
          </div>
          
          <div class="card-body-content">
            <div class="pricing-box">
              <span class="current-price">${{ course.currentPrice }}</span>
              <span class="original-price">${{ course.originalPrice }}</span>
              <span class="discount-tag">83% off</span>
            </div>
            
            <div class="urgency-banner">
              <svg class="clock-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
              <span>1 day left at this price!</span>
            </div>

            <button class="subscribe-btn" @click="handleSubscribe">Start Subscription</button>

            <!-- Share Wrapper -->
            <div class="share-wrapper" ref="shareWrapperRef">
              <div class="card-actions-row">
                <button class="action-icon-btn" title="Share" @click="toggleShare">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="18" cy="5" r="3"></circle><circle cx="6" cy="12" r="3"></circle><circle cx="18" cy="19" r="3"></circle><line x1="8.59" y1="13.51" x2="15.42" y2="17.49"></line><line x1="15.41" y1="6.51" x2="8.59" y2="10.49"></line></svg>
                </button>
              </div>
              <div v-if="showShare" class="share-dropdown-menu">
                <CourseShare :shareUrl="courseUrl" :shareTitle="course.title" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <ShowPopupSubscription  v-if="showSubPopup"  :thumbnail="course.thumbnail" @close="showSubPopup = false" @subscribe="handleCheckout"/>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRoute } from 'vue-router'
import NavBar from '../../navbar/NavBar.vue'
import CourseShare from './CourseShare.vue'
import ShowPopupSubscription from './ShowPopupSubscription.vue'

const route = useRoute()
const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'

function getAuthToken() {
  return localStorage.getItem('token') || ''
}
function authHeaders() {
  const token = getAuthToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}
function resolveImageUrl(url) {
  if (!url) return ''
  if (/^https?:\/\//i.test(url) || url.startsWith('data:image/')) return url
  return `${BASE_URL}${url.startsWith('/') ? '' : '/'}${url}`
}
function formatDate(dateStr) {
  if (!dateStr) return ''
  const d = new Date(dateStr)
  if (isNaN(d.getTime())) return ''
  const diffMs = Date.now() - d.getTime()
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24))
  if (diffDays <= 0) return 'today'
  if (diffDays === 1) return '1 day ago'
  if (diffDays < 30) return `${diffDays} days ago`
  return d.toLocaleDateString('en-US', { month: 'short', year: 'numeric' })
}

const shareWrapperRef = ref(null)
const showSubPopup = ref(false)
const showShare = ref(false)
const showFullBio = ref(false)
const isLoading = ref(false)
const error = ref(null)

const courseUrl = window.location.href

const handleClickOutside = (event) => {
  if (showShare.value && shareWrapperRef.value && !shareWrapperRef.value.contains(event.target)) {
    showShare.value = false
  }
}
const toggleShare = () => {
  showShare.value = !showShare.value
}

const course = ref({
  title: '',
  subtitle: '',
  rating: 0,
  ratingsCount: 0,
  studentsCount: 0,
  creator: '',
  lastUpdated: '',
  language: 'English',
  currentPrice: 0,
  originalPrice: 0,
  thumbnail: '',
  sectionsCount: 0,
  lecturesCount: 0,
  totalLength: '',
  sections: [],
  includes: [],
})

const instructor = ref({
  name: '',
  headline: '',
  avatar: '',
  rating: 0,
  reviewsCount: 0,
  studentsCount: 0,
  coursesCount: 0,
  description: '',
})

const reviews = ref([])

function mapCourseDetail(raw) {
  return {
    title: raw.title || '',
    subtitle: raw.subtitle || raw.description || '',
    rating: Number(raw.rating) || 0,
    ratingsCount: Number(raw.ratings_count) || 0,
    studentsCount: Number(raw.students_count) || 0,
    creator: raw.instructor?.name?.trim() || `Instructor #${raw.instructor_id}`,
    lastUpdated: formatDate(raw.updated_at || raw.created_at),
    language: 'English', 
    currentPrice: raw.current_price ?? 0,
    originalPrice: raw.original_price ?? 0,
    thumbnail: resolveImageUrl(raw.thumbnail),
    sectionsCount: raw.sections_count ?? (raw.sections ? raw.sections.length : 0),
    lecturesCount: raw.lectures_count ?? 0,
    totalLength: raw.total_length ?? '',
    sections: (raw.sections || []).map((s, i) => ({
      id: s.id,
      title: s.title,
      lecturesCount: s.lectures_count ?? (s.lectures ? s.lectures.length : 0),
      length: s.length ?? '',
      isOpen: i === 0,
      lectures: (s.lectures || []).map((l) => ({
        id: l.id,
        title: l.title,
        duration: l.duration,
      })),
    })),
    includes: (raw.includes || []).map((inc) => ({
      id: inc.id,
      icon: inc.icon,
      text: inc.text,
    })),
  }
}

function mapInstructor(raw) {
  const ins = raw.instructor || {}
  return {
    name: ins.name?.trim() || `Instructor #${raw.instructor_id}`,
    headline: ins.headline || '',
     avatar: resolveAvatarUrl(ins.avatar || ''),
    rating: Number(ins.rating) || 0,
    reviewsCount: Number(ins.reviews_count) || 0,
    studentsCount: Number(ins.students_count) || 0,
    coursesCount: Number(ins.courses_count) || 0,
    description: ins.description || '',
  }
}

function mapReviews(raw) {
  return (raw.reviews || []).map((r) => ({
    id: r.id,
    name: r.reviewer_name || 'Anonymous',
        avatar: resolveAvatarUrl(r.reviewer_avatar || ''), 
    date: formatDate(r.created_at),
    comment: r.comment || '',
    rating: Number(r.rating) || 0,
  }))
}

async function loadCourseDetail() {
  const id = route.params.id
  if (!id) {
    error.value = 'Missing course id'
    return
  }

  isLoading.value = true
  error.value = null

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/courses/show/${id}`, {
      headers: { ...authHeaders() },
      cache: 'no-store',
    })

    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }

    const json = await res.json()
    const raw = json?.data ?? json

    if (!raw || !raw.id) {
      throw new Error('Course not found')
    }

    course.value = mapCourseDetail(raw)
    instructor.value = mapInstructor(raw)
    reviews.value = mapReviews(raw)
  } catch (e) {
    error.value = e.message || 'Failed to load course'
    console.error('Failed to load course detail', e)
  } finally {
    isLoading.value = false
  }
}

const toggleSection = (id) => {
  const sec = course.value.sections.find((s) => s.id === id)
  if (sec) sec.isOpen = !sec.isOpen
}

const handleSubscribe = () => {
  showSubPopup.value = true
}

const handleCheckout = () => {
  showSubPopup.value = false
}

function includeIconKey(inc) {
  const icon = (inc.icon || '').toLowerCase().trim()
  const known = ['video', 'article', 'mobile', 'exercise', 'download', 'certificate']
  if (known.includes(icon)) return icon
  const text = (inc.text || '').toLowerCase()
  if (text.includes('video')) return 'video'
  if (text.includes('article')) return 'article'
  if (text.includes('mobile') || text.includes('tv')) return 'mobile'
  if (text.includes('exercise') || text.includes('coding')) return 'exercise'
  if (text.includes('download') || text.includes('resource')) return 'download'
  if (text.includes('certificate')) return 'certificate'
  return 'default'
}

function resolveAvatarUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

onMounted(() => {
  loadCourseDetail()
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap');

.course-detail-page {
  width: 100vw;
  min-height: 100vh;
  background-color: #F7F4F2;
  color: #1e293b;
  font-family: 'Plus Jakarta Sans', sans-serif;
  overflow-x: hidden;
}

.page-container {
  max-width: 1240px;
  margin: 0 auto;
  padding: 0 24px 32px 24px;
  display: flex;
  gap: 12px;
  position: relative;

}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 12px;
  max-width: 790px;
  padding: 12px;
  background-color: #ffffff;
}

.hero-section {
  background: #ffffff;
  padding: 24px;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 20px -2px rgba(0, 0, 0, 0.03);
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.hero-top-row {
  display: flex;
  gap: 20px;
  align-items: center;
}

.hero-left-thumb {
  width: 220px;
  height: 135px;
  border-radius: 12px;
  overflow: hidden;
  position: relative;
  background: #0f172a;
  flex-shrink: 0;
}

.hero-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.85;
}

.play-overlay-small {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 38px;
  height: 38px;
  background: rgba(37, 99, 235, 0.95);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  box-shadow: 0 4px 12px rgba(37, 99, 235, 0.3);
}

.play-overlay-small svg {
  width: 16px;
  height: 16px;
  margin-left: 2px;
}

.hero-right-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.course-title {
  font-size: 20px;
  font-weight: 800;
  color: #0f172a;
  line-height: 1.3;
}

.hero-creator-row {
  display: flex;
  align-items: center;
  gap: 4px;
  background: transparent;
  border: none;
  padding: 0;
  margin-top: -4px;
}

.create-by-label {
  font-size: 14px;
  font-weight: 700;
  color: #475569;
}

.instructor-box {
  background: #ffffff;
  border: none;
  color: #0f172a;
  padding: 0;
  border-radius: 0;
  font-size: 14px;
  font-weight: 600;
}

.hero-meta-row {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 32px;
}

.meta-card {
  background: transparent;
  border: none;
  border-radius: 0;
  padding: 4px 0;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #334155;
}

.meta-icon-svg {
  width: 18px;
  height: 18px;
  color: #64748b;
  flex-shrink: 0;
}

.highlight-val {
  background: transparent;
  border: none;
  padding: 0;
  border-radius: 0;
  font-weight: 600;
  color: #0f172a;
}

.hero-stats-row {
  display: flex;
  align-items: center;
  gap: 24px;
}

.stat-item {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 8px;
    padding-bottom: 6px;
}

.stat-item::after {
  content: "";
  position: absolute;
  bottom: 0;
  left: -6px;
  right: -6px;
  height: 6px;
  background: #1B75D2;
}

.stat-content {
  display: flex;
  flex-direction: column;
}

.star-icon {
  width: 18px;
  height: 18px;
  fill: #1B75D2;
  flex-shrink: 0;
}

.review-icon-svg {
  width: 18px;
  height: 18px;
  color: #64748b;
  flex-shrink: 0;
}

.inline-icon {
  width: 18px;
  height: 18px;
  color: #64748b;
  flex-shrink: 0;
}

.small-label {
  font-size: 11px;
  color: #64748b;
  text-transform: none;
  font-weight: 400;
  line-height: 1.1;
}

.stat-val {
  font-size: 14px;
  font-weight: 700;
  color: #0f172a;
  line-height: 1.2;
}

@media (max-width: 768px) {
  .hero-top-row {
    flex-direction: column;
    align-items: stretch;
  }
  .hero-left-thumb {
    width: 100%;
    height: 160px;
  }
  .hero-meta-row {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
  .hero-stats-row {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
}

.mobile-purchase-card {
  display: none;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 20px;
  padding: 20px;
  flex-direction: column;
  gap: 16px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.03);
}

@media (max-width: 968px) {
  .mobile-purchase-card {
    display: flex;
  }
}

.preview-thumb-box {
  width: 100%;
  height: 200px;
  /* border-radius: 12px; */
  overflow: hidden;
  position: relative;
  background: #0f172a;
    border-top-left-radius: 12px;
  border-top-right-radius: 12px;
}

.preview-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.85;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
}

.play-overlay {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 52px;
  height: 52px;
  background: rgba(37, 99, 235, 0.95);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  box-shadow: 0 4px 12px rgba(37, 99, 235, 0.3);
}

.play-overlay svg {
  width: 22px;
  height: 22px;
  margin-left: 3px;
}

.pricing-box {
  display: flex;
  align-items: baseline;
  gap: 12px;
}

.current-price {
  font-size: 22px;
  font-weight: 800;
  color: #0f172a;
}

.original-price {
  font-size: 14px;
  color: #f87171;
  text-decoration: line-through;
}

.discount-tag {
  font-size: 12px;
  font-weight: 700;
  color: #334155;
}

.urgency-banner {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #000000;
  font-size: 13px;
  font-weight: 600;
}

.clock-icon {
  width: 16px;
  height: 16px;
  stroke: #000000;
}

.subscribe-btn {
  width: 100%;
  background: #1B75D2;
  color: #fff;
  border: none;
  border-radius: 32px;
  padding: 12px 12px;
  font-size: 15px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
}

.subscribe-btn:hover {
  opacity: 0.8;
  transform: translateY(-1px);
}

.course-includes-section {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 28px;
  box-shadow: 0 4px 20px -2px rgba(0, 0, 0, 0.03);
}

.course-includes-section h2,
.course-content-section h2,
.instructor-profile-section h2 {
  font-size: 18px;
  font-weight: 800;
  margin-bottom: 20px;
  color: #0f172a;
}

.includes-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

@media (max-width: 600px) {
  .includes-grid {
    grid-template-columns: 1fr;
  }
}

.include-item {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 14px;
  color: #334155;
}

.inc-icon {
  width: 18px;
  height: 18px;
  color: #2563eb;
  flex-shrink: 0;
}

.course-content-section {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 28px;
  box-shadow: 0 4px 20px -2px rgba(0, 0, 0, 0.03);
}

.content-header-row {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
  margin-bottom: 16px;
  flex-wrap: wrap;
  gap: 8px;
}

.content-stats {
  font-size: 13px;
  color: #64748b;
}

.accordion-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 16px;
}

.accordion-item {
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  overflow: hidden;
  background: #ffffff;
  transition: border-color 0.2s;
}

.accordion-item:hover {
  border-color: #cbd5e1;
  background: transparent;
}

.accordion-summary {
  padding: 14px 18px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  background: transparent;
  border-radius: 12px;
}

.accordion-title-group {
  display: flex;
  align-items: center;
  gap: 12px;
}

.chevron-icon {
  width: 18px;
  height: 18px;
  transition: transform 0.2s;
  color: #64748b;
}

.chevron-icon.rotated {
  transform: rotate(180deg);
}

.sec-title {
  font-size: 14px;
  font-weight: 700;
  color: #1e293b;
}

.sec-meta {
  font-size: 13px;
  color: #64748b;
  display: flex;
  gap: 12px;
}

.accordion-body {
  padding: 12px 18px 16px 18px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  border-top: 1px solid #f1f5f9;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.lecture-row {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 13px;
  color: #334155;
  padding: 4px 0;
}

.play-small-icon {
  width: 14px;
  height: 14px;
  color: #2563eb;
  flex-shrink: 0;
}

.lecture-title {
  flex: 1;
}

.lecture-duration {
  color: #64748b;
}

.see-more-sections-btn {
  background: none;
  border: none;
  color: #1B75D2;
  padding: 12px 20px;
  border-radius: 12px;
  font-weight: 700;
  cursor: pointer;
  width: 100%;
  text-align: center;
  transition: opacity 0.2s;
}

.see-more-sections-btn:hover {
  background: none;
  opacity: 0.85;
}

.instructor-profile-section {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 28px;
  box-shadow: 0 4px 20px -2px rgba(0, 0, 0, 0.03);
}

.instructor-card-detailed {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.instructor-top-info {
  display: flex;
  gap: 20px;
  align-items: flex-start;
}

@media (max-width: 600px) {
  .instructor-top-info {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
}

.instructor-lg-avatar {
  width: 90px;
  height: 90px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid #eff6ff;
}

.instructor-meta-col {
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex: 1;
}

.instructor-name {
  font-size: 18px;
  font-weight: 800;
  color: #0f172a;
}

.instructor-headline {
  font-size: 13px;
  color: #64748b;
  margin-top: -4px;
}

.instructor-stats-top-row {
  display: flex;
  justify-content: flex-start;
  gap: 40px;
}

.instructor-stats-bottom-row {
  display: flex;
  align-items: center;
  gap: 48px;
  margin-top: 4px;
  border: none;
  background: transparent;
  padding: 0;
}

.stat-card-box {
  background: transparent;
  border: none;
  border-radius: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.giant-star-icon {
  width: 20px;
  height: 20px;
  fill: #1B75D2;
}

.stat-icon-svg {
  width: 20px;
  height: 20px;
  color: #64748b;
}

.stat-num-text {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}

.stat-label-text {
  font-size: 12px;
  color: #64748b;
}

.stat-pill-box {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #334155;
}

.pill-icon {
  width: 16px;
  height: 16px;
  color: #64748b;
}

.stat-pill-val {
  font-weight: 700;
  color: #0f172a;
}

.stat-pill-label {
  color: #64748b;
}

.instructor-bio {
  font-size: 14px;
  color: #334155;
  line-height: 1.6;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.instructor-bio p.clamped {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.toggle-bio-btn {
  background: none;
  border: none;
  color: #1B75D2;
  font-weight: 700;
  cursor: pointer;
  padding: 0;
  align-self: flex-start;
}

.student-feedback-section {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 28px;
  box-shadow: 0 4px 20px -2px rgba(0, 0, 0, 0.03);
}

.feedback-header-row {
  display: flex;
  gap: 40px;
  margin-bottom: 24px;
}

.overall-rating-box, .total-ratings-box {
  display: flex;
  flex-direction: column;
}

.big-rating-num {
  font-size: 32px;
  font-weight: 800;
  color: #0f172a;
  display: flex;
  align-items: center;
  gap: 8px;
}

.big-star {
  width: 28px;
  height: 28px;
  fill: #1B75D2;
}

.overall-rating-label {
  font-size: 13px;
  color: #64748b;
}

.reviews-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
  margin-bottom: 20px;
}

@media (max-width: 768px) {
  .reviews-grid {
    grid-template-columns: 1fr;
  }
}

.review-card {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 0; 
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  border-bottom: 10px solid #1B75D2;
}

.reviewer-header {
  display: flex;
  gap: 12px;
  align-items: center;
}

.reviewer-avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  object-fit: cover;
}

.reviewer-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.reviewer-name {
  font-size: 14px;
  font-weight: 700;
  color: #0f172a;
}

.review-stars-date {
  display: flex;
  align-items: center;
  gap: 8px;
}

.stars {
  display: flex;
  gap: 2px;
}

.inline-star {
  width: 12px;
  height: 12px;
  fill: #1B75D2;
}

.review-date {
  font-size: 12px;
  color: #64748b;
}

.review-comment {
  font-size: 14px;
  color: #334155;
  line-height: 1.5;
}

.see-more-reviews-btn {
  background: none;
  border: none;
  color: #1B75D2;
  font-weight: 700;
  cursor: pointer;
  padding: 0;

  display: block;
  margin: 0 auto;
  text-align: center;
  
  
}

.sidebar-container {
  width: 310px;
  flex-shrink: 0;
  position: relative;
  margin-top: 12px;
  overflow: visible;
}

@media (max-width: 968px) {
  .sidebar-container {
    display: none;
  }
}

.sticky-purchase-card {
  position: sticky;
  top: 24px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.05);
}

.card-body-content {
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.card-actions-row {
  display: flex;
  justify-content: center;
  margin-top: 4px;
}

.action-icon-btn {
  background: none;
  border: none;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #64748b;
  cursor: pointer;
  transition: background 0.2s;
}

.action-icon-btn:hover {
  background: #f1f5f9;
  color: #0f172a;
}

.action-icon-btn svg {
  width: 18px;
  height: 18px;
}

.share-wrapper {
  position: relative;
  width: 100%;
  display: flex;
  justify-content: center;
  margin-top: 12px;
}

.card-actions-row {
  display: flex;
  justify-content: center;
}

.action-icon-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  color: #64748b;
  transition: color 0.2s;
}

.action-icon-btn:hover {
  color: #0f172a;
}

.action-icon-btn svg {
  width: 20px;
  height: 20px;
}



.sticky-purchase-card {
  position: sticky;
  top: 24px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  overflow: visible !important; /* ប្តូរពី hidden មក visible */
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.05);
}

.card-body-content {
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  overflow: visible !important;
}

.share-wrapper {
  position: relative;
  width: 100%;
  display: flex;
  justify-content: center;
  margin-top: 12px;
  overflow: visible !important;
}

.share-dropdown-menu {
  position: absolute;
  bottom: 90%; 
  left: 46.2%;
  transform: translateX(-50%);
  margin-bottom: 2px;
  width: 300px;
  padding: 12px;
  border-radius: 12px;
  z-index: 9999;
}
</style>