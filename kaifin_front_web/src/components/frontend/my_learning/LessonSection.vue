<!-- File: LessonSection.vue -->
<script setup>
import { ref } from 'vue'

const props = defineProps({
  lesson: {
    type: Object,
    required: true
  },
  mediaUrl: {
    type: String,
    required: true
  }
})

const isOpen = ref(false)

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}
</script>

<template>
  <div class="lesson-section-wrapper" :class="{ 'is-open': isOpen }">
    <!-- Section Card -->
    <div class="section-card" @click="toggleDropdown">
      <div class="section-left-group">
        <div class="sec-thumbnail-wrapper">
          <div class="sec-thumbnail" :style="{ backgroundImage: `url(${mediaUrl})`, backgroundSize: 'cover', backgroundPosition: 'center' }">
            <div class="thumbnail-overlay">
              <svg viewBox="0 0 24 24" width="14" height="14" fill="currentColor"><path d="M8 5v14l11-7z"/></svg>
            </div>
          </div>
          <div class="sticker-badge" v-if="lesson.isHot">HOT</div>
        </div>
        
        <div class="sec-content-wrapper">
          <div class="sec-top-row">
            <span class="sec-badge">Module {{ lesson.id }}</span>
            <h3 class="sec-name">{{ lesson.title }}</h3>
          </div>
          <div class="sec-meta-info">
            <span class="sec-duration">
              <svg viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 7v5l3 3"/></svg>
              {{ lesson.duration }}
            </span>
            <span class="meta-dot">•</span>
            <span class="sec-type">2 Sub-lessons</span>
          </div>
        </div>
      </div>

      <div class="sec-action-right" :class="{ rotated: isOpen }">
        <svg viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="6 9 12 15 18 9"></polyline>
        </svg>
      </div>
    </div>

    <!-- Dropdown Content -->
    <div class="dropdown-content" v-if="isOpen">
      <div class="dropdown-inner">
        <ul class="sub-lesson-list">
          
          <!-- Sub-item Video -->
          <li class="sub-item">
            <div class="sub-info">
              <div class="sub-icon-box">
                <svg viewBox="0 0 24 24" width="13" height="13" fill="currentColor"><path d="M8 5v14l11-7z"/></svg>
              </div>
              <span class="sub-title">Video Lecture Overview</span>
            </div>
            <span class="sub-time">
              <svg viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 7v5l3 3"/></svg>
              10:00
            </span>
          </li>

          <!-- Sub-item  File -->
          <li class="sub-item">
            <div class="sub-info">
              <div class="sub-icon-box">
                <svg viewBox="0 0 24 24" width="13" height="13" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                  <polyline points="14 2 14 8 20 8"></polyline>
                </svg>
              </div>
              <span class="sub-title">Reading Materials & Source Code</span>
            </div>
            <span class="sub-time">
              <svg viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 7v5l3 3"/></svg>
              5:30
            </span>
          </li>

        </ul>
      </div>
    </div>
  </div>
</template>

<style scoped>
.lesson-section-wrapper {
  display: flex;
  flex-direction: column;
  background: #ffffff;
  border: none; 
  border-radius: 14px;
  margin-bottom: 12px;
  overflow: visible;
}

.section-card {
  background: transparent;
  border: none;
  outline: none;
  padding: 14px 18px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
}

.section-left-group {
  display: flex;
  align-items: center;
  gap: 16px;
}

.sec-thumbnail-wrapper {
  position: relative;
  width: 104px;
  height: 64px;
  flex-shrink: 0;
}

.sec-thumbnail {
  width: 100%;
  height: 100%;
  border-radius: 10px;
  background-color: #f1f5f9;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.thumbnail-overlay {
  position: absolute;
  inset: 0;
  background: rgba(15, 23, 42, 0.35);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.section-card:hover .thumbnail-overlay {
  opacity: 1;
}

.sticker-badge {
  position: absolute;
  top: -8px;
  right: -8px;
  background: linear-gradient(135deg, #ff416c, #ff4b2b);
  color: #ffffff;
  font-size: 10px;
  font-weight: 800;
  padding: 2px 7px;
  border-radius: 6px;
  transform: rotate(10deg);
  letter-spacing: 0.5px;
  border: 1.5px solid #ffffff;
  z-index: 10;
}

.sec-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.sec-top-row {
  display: flex;
  align-items: baseline;
  gap: 8px;
  flex-wrap: wrap;
}

.sec-badge {
  font-size: 11px;
  font-weight: 600;
  color: #475569;
  background: #f1f5f9;
  padding: 1px 7px;
  border-radius: 5px;
}

.sec-name {
  font-weight: 600;
  font-size: 14px;
  color: #0f172a;
  margin: 0;
  line-height: 1.4;
}

.sec-meta-info {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #64748b;
}

.sec-duration {
  display: flex;
  align-items: center;
  gap: 4px;
}

.meta-dot {
  font-size: 10px;
  color: #cbd5e1;
}

.sec-type {
  font-weight: 400;
}

.sec-action-right {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  background: #f8fafc;
  color: #64748b;
  display: flex;
  align-items: center;
  justify-content: center;
  border: none;
  transition: transform 0.2s ease;
}

.sec-action-right.rotated {
  transform: rotate(180deg);
}

.dropdown-content {
  background-color: #f8fafc;
  border-top: 1px solid #f1f5f9;
  padding: 12px 18px;
}

.dropdown-inner {
  display: flex;
  flex-direction: column;
}

.sub-lesson-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.sub-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 13px;
  color: #334155;
  background: transparent;
  border: none;
  padding: 8px 10px;
  border-radius: 8px;
  cursor: pointer;
}

.sub-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.sub-icon-box {
  width: 24px;
  height: 24px;
  border-radius: 6px;
  background: #e2e8f0;
  color: #475569;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.sub-title {
  font-weight: 500;
  color: #334155;
}

.sub-time {
  font-size: 12px;
  color: #475569;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 5px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  padding: 3px 8px;
  border-radius: 6px;
}
</style>