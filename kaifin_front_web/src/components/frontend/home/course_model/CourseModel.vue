<script setup>
import { ref, watch, nextTick } from 'vue';
// prop
const props = defineProps({
  isOpen: {
    type: Boolean,
    default: false
  },
  course: {
    type: Object,
    default: () => ({
      id: 1,
      title: "Vue 3 & Advanced Web Development",
      description: "Learn composition API, Pinia, and real-world project development from scratch. Build fast and scalable applications with modern tools.",
      image: "https://images.unsplash.com/photo-1516321318423-f06f85e504b3?w=800&auto=format&fit=crop&q=60",
      category: "Programming",
      badge: "Best Seller",
      lessons: "42 Lessons",
      duration: "15 Hours",
      rating: "4.8",
      students: "12,450 students",
      price: "$49.99",
      originalPrice: "$79.99",
      instructor: "Sarah Johnson",
      level: "Intermediate",
      language: "English",
      subtitle: "English",
      lastUpdated: "January 2026"
    })
  }
});

const emit = defineEmits(['close', 'enroll']);
const isVisible = ref(false);
const isLoading = ref(true);
const isEnrolling = ref(false);
const isFavorite = ref(false);
const modalRef = ref(null);
watch(() => props.isOpen, (newVal) => {
  if (newVal) {
    isVisible.value = true;
    document.body.style.overflow = 'hidden';
    nextTick(() => {
      isLoading.value = false;
      if (modalRef.value) {
        modalRef.value.focus();
      }
    });
  } else {
    isVisible.value = false;
    document.body.style.overflow = '';
  }
});

function closeModal() {
  isVisible.value = false;
  document.body.style.overflow = '';
  setTimeout(() => {
    emit('close');
  }, 300);
}

function handleEnroll() {
  isEnrolling.value = true;
  setTimeout(() => {
    isEnrolling.value = false;
    emit('enroll', props.course);
    alert(`🎉 Successfully enrolled in "${props.course.title}"!`);
    closeModal();
  }, 1500);
}

function toggleFavorite() {
  isFavorite.value = !isFavorite.value;
}

function handleKeydown(event) {
  if (event.key === 'Escape') {
    closeModal();
  }
}

function handleOutsideClick(event) {
  if (event.target === event.currentTarget) {
    closeModal();
  }
}

const discountPercentage = ref(0);
if (props.course.originalPrice) {
  const original = parseFloat(props.course.originalPrice.replace('$', ''));
  const current = parseFloat(props.course.price.replace('$', ''));
  if (original > current) {
    discountPercentage.value = Math.round(((original - current) / original) * 100);
  }
}
</script>

<template>
  <Teleport to="body">
    <Transition name="modal-fade">
      <div 
        v-if="isVisible"
        ref="modalRef"
        class="modal-backdrop"
        @click.self="handleOutsideClick"
        @keydown.escape="handleKeydown"
        tabindex="0"
        role="dialog"
        aria-modal="true"
        :aria-label="`Course details: ${course.title}`"
      >
        <Transition name="modal-slide">
          <div class="modal-card-container" v-if="isVisible">
            <button 
              class="close-btn" 
              @click="closeModal" 
              aria-label="Close modal"
              title="Close (Esc)"
            >
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <line x1="18" y1="6" x2="6" y2="18"></line>
                <line x1="6" y1="6" x2="18" y2="18"></line>
              </svg>
            </button>
            <button 
              class="favorite-btn" 
              @click="toggleFavorite"
              :class="{ 'is-favorite': isFavorite }"
              aria-label="Toggle favorite"
            >
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
              </svg>
            </button>
            <div v-if="isLoading" class="loading-state">
              <div class="skeleton-image"></div>
              <div class="skeleton-content">
                <div class="skeleton-line w-80"></div>
                <div class="skeleton-line w-60"></div>
                <div class="skeleton-line w-40"></div>
              </div>
            </div>
            <template v-else>
              <!-- Course Image -->
              <div class="modal-img-box">
                <img :src="course.image" :alt="course.title" loading="lazy" />
                
                <div class="badge-overlay">
                  <span class="badge-item category-badge">{{ course.category }}</span>
                  <span v-if="course.badge" class="badge-item sports-badge">{{ course.badge }}</span>
                  <span v-if="discountPercentage > 0" class="badge-item discount-badge">
                    -{{ discountPercentage }}%
                  </span>
                </div>

                <div class="instructor-overlay" v-if="course.instructor">
                  <div class="instructor-avatar">
                    <img 
                      :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${course.instructor}`" 
                      :alt="course.instructor"
                    />
                  </div>
                  <div class="instructor-info">
                    <span class="instructor-name">{{ course.instructor }}</span>
                    <span class="instructor-label">Instructor</span>
                  </div>
                </div>
              </div>

              <div class="modal-body-content">
                <h2 class="course-title">{{ course.title }}</h2>
                
                <div class="rating-row">
                  <div class="stars">
                    <span v-for="n in 5" :key="n" class="star" :class="{ filled: n <= Math.floor(parseFloat(course.rating)) }">
                      ★
                    </span>
                    <span class="rating-value">{{ course.rating }}</span>
                  </div>
                  <span class="students-count">{{ course.students }}</span>
                </div>

                <p class="course-desc">{{ course.description }}</p>

                <div class="course-tags">
                  <span class="tag" v-if="course.level">📊 {{ course.level }}</span>
                  <span class="tag" v-if="course.language">🌐 {{ course.language }}</span>
                  <span class="tag" v-if="course.subtitle">📝 {{ course.subtitle }}</span>
                  <span class="tag" v-if="course.lastUpdated">📅 {{ course.lastUpdated }}</span>
                </div>

                <a href="#" class="read-more-link" @click.prevent>📖 Read the full syllabus →</a>

                <hr class="divider" />

                <div class="course-meta-section">
                  <h3 class="section-heading">📋 Course Information</h3>
                  
                  <div class="meta-grid">
                    <div class="meta-item">
                      <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"></path>
                        <path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"></path>
                      </svg>
                      <span>{{ course.lessons }}</span>
                    </div>
                    <div class="meta-item">
                      <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <circle cx="12" cy="12" r="10"></circle>
                        <polyline points="12 6 12 12 16 14"></polyline>
                      </svg>
                      <span>{{ course.duration }}</span>
                    </div>
                    <div class="meta-item">
                      <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon>
                      </svg>
                      <span>{{ course.rating }} (Reviews)</span>
                    </div>
                    <div class="meta-item">
                      <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                        <circle cx="9" cy="7" r="4"></circle>
                      </svg>
                      <span>{{ course.students }}</span>
                    </div>
                  </div>

                  <div class="promo-footer-box">
                    <div class="price-info">
                      <span class="price-label">Total Price:</span>
                      <div class="price-wrapper">
                        <span class="price-value">{{ course.price }}</span>
                        <span v-if="course.originalPrice" class="original-price">
                          {{ course.originalPrice }}
                        </span>
                      </div>
                      <span v-if="discountPercentage > 0" class="discount-text">
                        Save {{ discountPercentage }}%!
                      </span>
                    </div>
                    
                    <button 
                      class="enroll-btn" 
                      @click="handleEnroll"
                      :disabled="isEnrolling"
                    >
                      <span v-if="!isEnrolling">🚀 Enroll Now</span>
                      <span v-else class="loading-spinner">
                        <span class="spinner"></span> Processing...
                      </span>
                    </button>
                  </div>

                  <!-- ========== WHAT YOU'LL LEARN ========== -->
                  <div class="learn-section">
                    <h4 class="learn-title">✅ What you'll learn</h4>
                    <ul class="learn-list">
                      <li>Build real-world Vue 3 applications from scratch</li>
                      <li>Master Composition API and Pinia state management</li>
                      <li>Create responsive and accessible user interfaces</li>
                      <li>Deploy applications to production environments</li>
                    </ul>
                  </div>
                </div>
              </div>
            </template>
          </div>
        </Transition>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(15, 23, 42, 0.75);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 20px;
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
}


.modal-card-container {
  background: #ffffff;
  width: 100%;
  max-width: 580px;
  border-radius: 24px;
  overflow: hidden;
  position: relative;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.3);
  max-height: 90vh;
  overflow-y: auto;
  scroll-behavior: smooth;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.modal-card-container::-webkit-scrollbar {
  width: 8px;
}

.modal-card-container::-webkit-scrollbar-track {
  background: #f8fafc;
}

.modal-card-container::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 10px;
}

.modal-card-container::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

.close-btn {
  position: absolute;
  top: 16px;
  right: 16px;
  background: rgba(15, 23, 42, 0.6);
  color: white;
  border: none;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 20;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(4px);
}

.close-btn:hover {
  background: rgba(15, 23, 42, 0.85);
  transform: scale(1.05) rotate(90deg);
}

.favorite-btn {
  position: absolute;
  top: 16px;
  right: 64px;
  background: rgba(15, 23, 42, 0.6);
  color: white;
  border: none;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 20;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(4px);
}

.favorite-btn:hover {
  background: rgba(15, 23, 42, 0.85);
  transform: scale(1.1);
}

.favorite-btn.is-favorite {
  background: rgba(225, 29, 72, 0.9);
  color: white;
}

.favorite-btn.is-favorite svg {
  fill: white;
}

.modal-img-box {
  width: 100%;
  height: 280px;
  position: relative;
  overflow: hidden;
  background: #0f172a;
}

.modal-img-box img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.modal-card-container:hover .modal-img-box img {
  transform: scale(1.04);
}

.badge-overlay {
  position: absolute;
  top: 16px;
  left: 16px;
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  z-index: 10;
}

.badge-item {
  font-size: 11px;
  font-weight: 700;
  padding: 6px 12px;
  border-radius: 20px;
  backdrop-filter: blur(6px);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}

.category-badge {
  background: rgba(37, 99, 235, 0.9);
  color: white;
}

.sports-badge {
  background: linear-gradient(135deg, #f59e0b, #ea580c);
  color: white;
}

.discount-badge {
  background: linear-gradient(135deg, #ef4444, #b91c1c);
  color: white;
}

.instructor-overlay {
  position: absolute;
  bottom: 16px;
  left: 16px;
  display: flex;
  align-items: center;
  gap: 12px;
  background: rgba(15, 23, 42, 0.75);
  padding: 6px 14px 6px 6px;
  border-radius: 50px;
  backdrop-filter: blur(6px);
  z-index: 10;
  border: 1px solid rgba(255, 255, 255, 0.15);
}

.instructor-avatar {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  overflow: hidden;
  border: 2px solid rgba(255, 255, 255, 0.8);
}

.instructor-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.instructor-info {
  color: white;
}

.instructor-name {
  font-size: 13px;
  font-weight: 700;
  display: block;
  line-height: 1.2;
}

.instructor-label {
  font-size: 10px;
  opacity: 0.75;
  letter-spacing: 0.3px;
}

.modal-body-content {
  padding: 24px 28px;
}

.course-title {
  font-size: 22px;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 10px 0;
  line-height: 1.35;
}

.rating-row {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 14px;
}

.stars {
  display: flex;
  align-items: center;
  gap: 3px;
}

.star {
  color: #cbd5e1;
  font-size: 15px;
}

.star.filled {
  color: #f59e0b;
}

.rating-value {
  font-size: 14px;
  font-weight: 700;
  color: #0f172a;
  margin-left: 4px;
}

.students-count {
  font-size: 13px;
  color: #64748b;
  font-weight: 500;
}

.course-desc {
  font-size: 14px;
  color: #475569;
  line-height: 1.65;
  margin: 0 0 16px 0;
}


.course-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 16px;
}

.course-tags .tag {
  background: #f1f5f9;
  color: #334155;
  padding: 5px 12px;
  border-radius: 8px;
  font-size: 12px;
  font-weight: 600;
  border: 1px solid #e2e8f0;
}

.read-more-link {
  color: #2563eb;
  font-size: 14px;
  font-weight: 600;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  gap: 4px;
  transition: color 0.2s ease;
}

.read-more-link:hover {
  color: #1d4ed8;
  text-decoration: underline;
}

.divider {
  border: none;
  border-top: 1px solid #e2e8f0;
  margin: 20px 0;
}

.section-heading {
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
  margin: 0 0 12px 0;
}

.meta-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
  margin-bottom: 20px;
}

.meta-item {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  padding: 10px 14px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 13px;
  font-weight: 600;
  color: #334155;
  transition: all 0.2s ease;
}

.meta-item:hover {
  background: #f1f5f9;
  border-color: #cbd5e1;
}

.meta-item svg {
  color: #2563eb;
  flex-shrink: 0;
}

.promo-footer-box {
  background: linear-gradient(135deg, #f8fafc, #f1f5f9);
  border: 1px solid #e2e8f0;
  border-radius: 14px;
  padding: 16px 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  flex-wrap: wrap;
  margin-bottom: 20px;
}

.price-info {
  display: flex;
  flex-direction: column;
}

.price-label {
  font-size: 11px;
  color: #64748b;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.price-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
}

.price-value {
  font-size: 24px;
  font-weight: 800;
  color: #059669;
}

.original-price {
  font-size: 14px;
  color: #94a3b8;
  text-decoration: line-through;
  font-weight: 600;
}

.discount-text {
  font-size: 12px;
  font-weight: 700;
  color: #dc2626;
}

.enroll-btn {
  background: linear-gradient(135deg, #2563eb, #1d4ed8);
  color: white;
  border: none;
  padding: 12px 26px;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  min-width: 140px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  box-shadow: 0 4px 12px rgba(37, 99, 235, 0.25);
}

.enroll-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, #1d4ed8, #1e40af);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(37, 99, 235, 0.35);
}

.enroll-btn:active:not(:disabled) {
  transform: scale(0.97);
}

.enroll-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.loading-spinner {
  display: flex;
  align-items: center;
  gap: 8px;
}

.spinner {
  display: inline-block;
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  border-top-color: white;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.learn-section {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 16px 20px;
}

.learn-title {
  font-size: 14px;
  font-weight: 700;
  color: #0f172a;
  margin: 0 0 10px 0;
}

.learn-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: grid;
  gap: 6px;
}

.learn-list li {
  font-size: 13px;
  color: #475569;
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
}

.learn-list li::before {
  content: "✓";
  color: #059669;
  font-weight: 800;
}

.loading-state {
  width: 100%;
}

.skeleton-image {
  width: 100%;
  height: 280px;
  background: linear-gradient(90deg, #e2e8f0 25%, #f1f5f9 50%, #e2e8f0 75%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
}

.skeleton-content {
  padding: 24px 28px;
}

.skeleton-line {
  height: 14px;
  background: linear-gradient(90deg, #e2e8f0 25%, #f1f5f9 50%, #e2e8f0 75%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
  border-radius: 6px;
  margin-bottom: 12px;
}

.skeleton-line:last-child {
  margin-bottom: 0;
}

.skeleton-line.w-80 { width: 80%; }
.skeleton-line.w-60 { width: 60%; }
.skeleton-line.w-40 { width: 40%; }

@keyframes shimmer {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}

.modal-slide-enter-active,
.modal-slide-leave-active {
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

.modal-slide-enter-from {
  transform: scale(0.95) translateY(20px);
  opacity: 0;
}

.modal-slide-leave-to {
  transform: scale(0.95) translateY(20px);
  opacity: 0;
}

/* ======================== */
/* RESPONSIVE */
/* ======================== */
@media (max-width: 640px) {
  .modal-backdrop {
    padding: 12px;
  }
  
  .modal-card-container {
    max-width: 100%;
    border-radius: 20px;
    max-height: 95vh;
  }
  
  .modal-img-box {
    height: 220px;
  }
  
  .modal-body-content {
    padding: 20px;
  }
  
  .course-title {
    font-size: 19px;
  }
  
  .meta-grid {
    grid-template-columns: 1fr 1fr;
    gap: 8px;
  }
  
  .promo-footer-box {
    flex-direction: column;
    align-items: stretch;
    text-align: center;
  }
  
  .enroll-btn {
    width: 100%;
  }
  
  .price-wrapper {
    justify-content: center;
  }
  
  .close-btn,
  .favorite-btn {
    width: 34px;
    height: 34px;
  }
  
  .favorite-btn {
    right: 58px;
  }
}

@media (max-width: 400px) {
  .meta-grid {
    grid-template-columns: 1fr;
  }
  
  .badge-overlay {
    flex-direction: column;
  }
}
</style>