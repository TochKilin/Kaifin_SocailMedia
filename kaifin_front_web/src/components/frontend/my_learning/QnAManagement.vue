<!-- File: QnAManagement.vue -->
<script setup>
import { ref } from 'vue'

const searchQuery = ref('')
const selectedLecture = ref('All lectures')
const sortBy = ref('Sort by commented')
const filterQnA = ref('Filter Q&A')

// Sample Q&A items matching the wireframe design
const qnaList = ref([
  {
    id: 1,
    author: 'Sovvan Vuthea',
    avatarBg: '#e2e8f0',
    question: 'How do we configure the routing parameters inside the component setup?',
    answer: 'You can define your routes dynamically using Vue Router and access parameters via useRoute().'
  },
  {
    id: 2,
    author: 'Kilin Chan',
    avatarBg: '#f1f5f9',
    question: 'Is there a specific performance guideline for rendering large lists?',
    answer: 'Yes, consider using virtual scrolling or implementing pagination to optimize rendering performance.'
  }
])
</script>

<template>
  <div class="qna-container">
    
    <!-- Search Bar Row -->
    <div class="search-bar-wrapper">
      <input 
        type="text" 
        class="qna-search-input" 
        placeholder="Search all course Q&A..." 
        v-model="searchQuery"
      />
      <button class="search-btn" aria-label="Search">
        <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2.5">
          <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/>
        </svg>
      </button>
    </div>

    <!-- Filters and Sort Section -->
    <div class="controls-section">
      <div class="control-group">
        <label class="control-label">Filters:</label>
        <div class="select-wrapper">
          <select v-model="selectedLecture" class="custom-select">
            <option>All lectures</option>
            <option>Lecture 1: Introduction</option>
            <option>Lecture 2: Setup & Components</option>
          </select>
          <svg class="dropdown-arrow" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><polyline points="6 9 12 15 18 9"/></svg>
        </div>
      </div>

      <div class="control-group">
        <label class="control-label">Sort by:</label>
        <div class="select-wrapper">
          <select v-model="sortBy" class="custom-select">
            <option>Sort by commented</option>
            <option>Sort by recent</option>
            <option>Sort by upvotes</option>
          </select>
          <svg class="dropdown-arrow" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><polyline points="6 9 12 15 18 9"/></svg>
        </div>
      </div>

      <div class="control-group filter-dropdown-group">
        <div class="select-wrapper">
          <select v-model="filterQnA" class="custom-select">
            <option>Filter Q&A</option>
            <option>My questions</option>
            <option>Unanswered questions</option>
          </select>
          <svg class="dropdown-arrow" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><polyline points="6 9 12 15 18 9"/></svg>
        </div>
      </div>
    </div>

    <!-- Sub-header Meta Pills -->
    <div class="meta-pills-row">
      <div class="meta-pill-main">All questions in this course</div>
      <div class="meta-pill-count">{{ qnaList.length }}</div>
    </div>

    <!-- Q&A List Cards -->
    <div class="qna-list">
      <div v-for="item in qnaList" :key="item.id" class="qna-card">
        <div class="user-avatar-box">
          <svg viewBox="0 0 24 24" width="28" height="28" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/>
          </svg>
        </div>
        
        <div class="qna-content-stack">
          <div class="qna-bubble question-box">
            <span class="bubble-label">Questions in this course</span>
            <p class="bubble-text">{{ item.question }}</p>
          </div>
          <div class="qna-bubble answer-box">
            <span class="bubble-label">Answer in this course</span>
            <p class="bubble-text">{{ item.answer }}</p>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<style scoped>
.qna-container {
  background-color: #ffffff;
  padding: 24px;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  color: #1e293b;
}

/* Search Bar */
.search-bar-wrapper {
  display: flex;
  position: relative;
  width: 100%;
}

.qna-search-input {
  width: 100%;
  /* background-color: #f8fafc; */
  border: 1px solid #cbd5e1;
  border-radius: 10px;
  padding: 14px 50px 14px 18px;
  font-size: 15px;
  color: #0f172a;
  outline: none;
  transition: border-color 0.2s ease;
}

.qna-search-input:focus {
  box-shadow: 0 0 0 3px #1976d2cf;
  background-color: #ffffff;
}

.search-btn {
  position: absolute;
  right: 6px;
  top: 6px;
  bottom: 6px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 8px;
  width: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.search-btn:hover {
  background-color: #2563eb;
}

/* Controls Section */
.controls-section {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  align-items: flex-end;
}

.control-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.control-label {
  font-size: 13px;
  font-weight: 700;
  color: #0f172a;
}

.select-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.custom-select {
  appearance: none;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  padding: 10px 36px 10px 14px;
  font-size: 14px;
  font-weight: 500;
  color: #334155;
  cursor: pointer;
  outline: none;
  transition: border-color 0.2s ease;
}

.custom-select:focus {
  border-color: #3b82f6;
}

.dropdown-arrow {
  position: absolute;
  right: 12px;
  color: #64748b;
  pointer-events: none;
}

/* Meta Pills Row */
.meta-pills-row {
  display: flex;
  gap: 10px;
  align-items: center;
}

.meta-pill-main {
  /* background-color: #f1f5f9; */
  border: 1px solid #e2e8f0;
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 13.5px;
  font-weight: 600;
  color: #334155;
}

.meta-pill-count {
  /* background-color: #eff6ff; */
  border: 1px solid #dbeafe;
  color: #2563eb;
  padding: 8px 14px;
  border-radius: 8px;
  font-size: 13.5px;
  font-weight: 700;
}

/* Q&A List & Cards */
.qna-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.qna-card {
  display: flex;
  gap: 16px;
  /* background-color: #f8fafc; */
  /* border: 1px solid #e2e8f0; */
  border-radius: 12px;
  padding: 20px;
  align-items: flex-start;
}

.user-avatar-box {
  color: #334155;
  padding: 12px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  border: 1px solid #e2e8f0;
}

.qna-content-stack {
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex-grow: 1;
}

.qna-bubble {
  background-color: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  padding: 12px 16px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.answer-box {
  background-color: #3b83f60e;
  border-color: #3b83f60e;
}

.bubble-label {
  font-size: 11px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: #64748b;
}

.answer-box .bubble-label {
  color: #1976D2;
}

.bubble-text {
  font-size: 14px;
  color: #1e293b;
  margin: 0;
  line-height: 1.5;
  font-weight: 500;
}
</style>