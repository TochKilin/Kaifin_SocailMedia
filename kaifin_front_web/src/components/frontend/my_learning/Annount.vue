<!-- File: Annount.vue -->
<script setup>
import { ref } from 'vue'

const announcements = ref([
  { id: 1, text: '📢 New lesson uploaded', badge: 'NEW', date: 'Aug 03, 2026' },
  { id: 2, text: '📢 Mock announcement data for testing 1', date: 'Aug 01, 2026' },
  { id: 3, text: '📢 Mock announcement data for testing 2', badge: 'NEW', date: 'Aug 02, 2026' }
])
</script>

<template>
  <div class="announcement-container">
    
    <div class="announcement-header">
      <div class="main-title-box">
        <h2>All Announcement</h2>
      </div>
    </div>

    <div class="announcement-list">
      <div 
        v-for="item in announcements" 
        :key="item.id" 
        class="announcement-card"
      >
        <div class="announcement-content-wrapper">
          <div class="text-with-badge">
            <p class="announcement-text">{{ item.text }}</p>
            <span v-if="item.badge" class="badge-new">{{ item.badge }}</span>
          </div>
        </div>
        <span class="announcement-date">{{ item.date }}</span>
      </div>
    </div>

  </div>
</template>

<style scoped>
.announcement-container {
  background-color: transparent;
  padding: 12px 18px;
  border: none;
  display: flex;
  flex-direction: column;
  gap: 16px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
}

.announcement-header {
  display: flex;
  align-items: center;
}

.main-title-box {
  flex: 1;
  background-color: transparent;
  border: none;
  padding: 12px 0;
  text-align: center;
}

.main-title-box h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
  letter-spacing: 0.3px;
}

.announcement-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.announcement-card {
  background-color: transparent;
  border: none;
  padding: 12px 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  border-bottom: 1px solid #f1f5f9;
}

.announcement-card:hover {
  opacity: 0.8;
}

.announcement-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.text-with-badge {
  display: flex;
  align-items: center;
  gap: 8px;
}

.announcement-text {
  margin: 0;
  font-size: 15px;
  font-weight: 500;
  color: #1e293b;
}

.badge-new {
  background-color: #ef4444;
  color: #ffffff;
  font-size: 10px;
  font-weight: 700;
  padding: 2px 6px;
  border-radius: 4px;
  letter-spacing: 0.5px;
}

.announcement-date {
  font-size: 13px;
  color: #64748b;
  white-space: nowrap;
}
</style>