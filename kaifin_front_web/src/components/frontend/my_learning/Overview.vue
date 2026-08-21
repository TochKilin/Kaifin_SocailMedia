<!-- File: Overview.vue -->
<script setup>
import { ref } from 'vue'

const courseStats = ref({
  rating: 4.8,
  totalRatings: '42,580',
  students: '124,335',
  duration: '28 hours',
  lastUpdate: 'January 2026',
  languages: 'English, Khmer',
  skillLevel: 'All Levels',
  lectures: 200,
  captions: 'Yes'
})

const handleStatClick = (type) => {
  console.log('Clicked:', type)
}
</script>

<template>
  <div class="overview-container">
  
    <div class="stats-overview-group">
      <button class="stat-btn" @click="handleStatClick('ratings')">
        <div class="stat-header">
          <span class="rating-number">{{ courseStats.rating }}</span>
          <div class="stars-wrapper">⭐ ⭐ ⭐ ⭐</div>
        </div>
        <span class="stat-label">{{ courseStats.totalRatings }} Ratings</span>
      </button>

      <button class="stat-btn" @click="handleStatClick('students')">
        <div class="stat-value-row">
          <span class="stat-icon">👥</span>
          <span class="stat-main-text">{{ courseStats.students }}</span>
        </div>
        <span class="stat-label">Enrolled Students</span>
      </button>

      <button class="stat-btn" @click="handleStatClick('content')">
        <div class="stat-value-row">
          <span class="stat-icon">📚</span>
          <span class="stat-main-text">{{ courseStats.duration }}</span>
        </div>
        <span class="stat-label">Total Course Content</span>
      </button>
    </div>

    <div class="custom-split-card-2">
      <div class="split-meta-item">
        <div class="svg-icon-bg-small">🔄</div>
        <span>Last updated: <strong>{{ courseStats.lastUpdate }}</strong></span>
      </div>
      <div class="split-meta-item">
        <div class="svg-icon-bg-small">🌍</div>
        <span>Languages: <strong>{{ courseStats.languages }}</strong></span>
      </div>
    </div>

    <div class="content-section">
      <h3 class="section-title">By the numbers</h3>
      
      <div class="custom-split-card-3">
        <div class="split-item">
          <span class="item-label">🎯 Skill level</span>
          <span class="item-value">{{ courseStats.skillLevel }}</span>
        </div>
        <div class="split-item">
          <span class="item-label">📚 Lectures</span>
          <span class="item-value">{{ courseStats.lectures }}</span>
        </div>
        <div class="split-item">
          <span class="item-label">👥 Students</span>
          <span class="item-value">{{ courseStats.students }}</span>
        </div>
      </div>

      <div class="custom-split-card-3">
        <div class="split-item">
          <span class="item-label">🎥 Videos</span>
          <span class="item-value">{{ courseStats.duration }}</span>
        </div>
        <div class="split-item">
          <span class="item-label">🌐 Languages</span>
          <span class="item-value">English</span>
        </div>
        <div class="split-item">
          <span class="item-label">💬 Captions</span>
          <span class="item-value">{{ courseStats.captions }}</span>
        </div>
      </div>

    </div>

    <!-- Certificates Section -->
    <div class="content-section">
      <h3 class="section-title">Certificates</h3>
      <div class="info-banner-card">
        <div class="banner-icon">✅</div>
        <div class="banner-content">
          <h4>Earn a verifiable certificate</h4>
          <p>Get a shareable digital certificate upon successful completion to highlight your new skills to employers and your professional network.</p>
        </div>
      </div>
    </div>

    <div class="content-section">
      <h3 class="section-title">Features</h3>
      <div class="custom-split-card-2">
        <div class="split-meta-item">
          <div class="svg-icon-bg-small">💻</div>
          <span>Access on desktop & mobile</span>
        </div>
        <div class="split-meta-item">
          <div class="svg-icon-bg-small">⬇️</div>
          <span>Full lifetime access</span>
        </div>
      </div>
    </div>

    <!-- Description Section -->
    <div class="content-section">
      <h3 class="section-title">Description</h3>
      <div class="description-box">
        <p>This comprehensive course is designed to take you from foundational concepts to advanced practical mastery. Through structured lessons, hands-on projects, and real-world case studies, you will gain the exact skills required to excel in modern software engineering and architecture.</p>
        <p>Whether you are looking to upskill, pivot your career, or build complex production-ready systems, this curriculum delivers clear explanations and expert guidance every step of the way.</p>
      </div>
    </div>

  </div>
</template>

<style scoped>
.overview-container {
  background-color: #ffffff;
  padding: 24px;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  color: #1e293b;
}

.stats-overview-group {
  display: flex;
  border: 1px solid #cbd5e1;
  border-radius: 10px;
  overflow: hidden;
  background-color: #ffffff;
}

.stat-btn {
  flex: 1;
  background: none;
  border: none;
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding: 12px 16px;
  cursor: pointer;
  text-align: left;
  transition: background-color 0.2s ease;
}

.stat-btn:hover {
  background-color: #f8fafc;
}

.stat-btn:not(:last-child) {
  border-right: 1px solid #cbd5e1;
}

.stat-header {
  display: flex;
  align-items: center;
  gap: 8px;
}

.rating-number {
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
}

.stars-wrapper {
  display: flex;
  gap: 2px;
  color: #f59e0b;
  font-size: 11px;
}

.stat-value-row {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #0f172a;
}

.stat-icon {
  font-size: 14px;
}

.stat-main-text {
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
}

.stat-label {
  font-size: 12px;
  color: #64748b;
  font-weight: 500;
}

.svg-icon-bg-small {
  padding: 4px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  font-size: 12px;
}

/* ប្រអប់ភ្ជាប់គ្នា ២ */
.custom-split-card-2 {
  display: flex;
  border: 1px solid #cbd5e1;
  border-radius: 10px;
  overflow: hidden;
  background-color: #ffffff;
}

.split-meta-item {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  font-size: 12.5px;
  color: #334155;
  font-weight: 500;
}

.split-meta-item:not(:last-child) {
  border-right: 1px solid #cbd5e1;
}

.content-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.section-title {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

/* ប្រអប់ភ្ជាប់គ្នា ៣ ក្នុងមួយជួរដេក */
.custom-split-card-3 {
  display: flex;
  border: 1px solid #cbd5e1;
  border-radius: 10px;
  overflow: hidden;
  background-color: #ffffff;
  margin-bottom: 6px;
}

.split-item {
  flex: 1;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 9px 12px;
}

.split-item:not(:last-child) {
  border-right: 1px solid #cbd5e1;
}

.item-label {
  font-size: 12px;
  color: #475569;
  font-weight: 500;
}

.item-value {
  font-size: 12px;
  color: #0f172a;
  font-weight: 600;
}

.info-banner-card {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 12px;
}

.banner-icon {
  color: #334155;
  padding: 8px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  font-size: 14px;
}

.banner-content h4 {
  font-size: 14px;
  font-weight: 600;
  color: #0f172a;
  margin: 0 0 2px 0;
}

.banner-content p {
  font-size: 12.5px;
  color: #475569;
  margin: 0;
  line-height: 1.4;
}

.description-box {
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 12px 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.description-box p {
  font-size: 13px;
  color: #475569;
  line-height: 1.5;
  margin: 0;
}
</style>