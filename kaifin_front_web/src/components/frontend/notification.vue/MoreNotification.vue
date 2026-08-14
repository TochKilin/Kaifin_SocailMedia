<!-- File: MoreNotification.vue -->
<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const emit = defineEmits([
  'seeAllNew', 
  'seeAllFollow', 
  'seeAllWeek', 
  'followBack', 
  'seePrevious', 
  'allPost', 
  'notificationSetting', 
  'fullScreen'
])

// State សម្រាប់គ្រប់គ្រងការបង្ហាញ Menu
const showDropdown = ref(false)

const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

// បិទ Menu វិញពេលចុចក្រៅ
const closeDropdown = (e) => {
  if (!e.target.closest('.header-dots-wrapper')) {
    showDropdown.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', closeDropdown)
})

onUnmounted(() => {
  document.removeEventListener('click', closeDropdown)
})

const newNotifications = ref([
  {
    type: 'course',
    username: 'Instructor',
    message: 'New course',
    timeAgo: '12m',
    profileImage: ''
  },
  {
    type: 'course',
    username: 'Instructor',
    message: 'Update course',
    timeAgo: '13m',
    profileImage: ''
  },
  {
    type: 'course',
    username: 'Instructor',
    message: 'Instructor have assignment',
    timeAgo: '14m',
    profileImage: ''
  },
  {
    type: 'course',
    username: 'Instructor',
    message: 'Instructor have Quiz',
    timeAgo: '15m',
    profileImage: ''
  }
])

const followNotifications = ref([
  {
    username: 'Instructor',
    message: 'New Instructor you can follow',
    timeAgo: '12h',
    profileImage: ''
  }
])

const weekNotifications = ref([
  {
    username: 'Instructor',
    message: 'Have meeting',
    timeAgo: '1 week',
    profileImage: ''
  }
])

const deleteFollow = (index) => {
  followNotifications.value.splice(index, 1)
}
</script>

<template>
  <div class="article-notification-container">
    <!-- Main Header with Title & More Dropdown -->
    <div class="main-header">
      <div class="header-title-wrapper">
        <!-- SVG Bell Icon -->
        <svg class="transparent-svg" viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
          <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
        </svg>
        <h1 class="main-title">Notifications</h1>
      </div>

      <!-- More Dots Button & Dropdown Menu -->
      <div class="header-dots-wrapper">
        <div class="header-dots" @click.stop="toggleDropdown">
          <span></span><span></span><span></span>
        </div>

        <div v-if="showDropdown" class="dropdown-menu">
          <!-- All Post with SVG Icon -->
          <button class="dropdown-item" @click="$emit('allPost'); showDropdown = false">
            <svg class="transparent-svg" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
            </svg>
            All Post
          </button>
          
          <!-- Notification Setting with SVG Icon -->
          <button class="dropdown-item" @click="$emit('notificationSetting'); showDropdown = false">
            <svg class="transparent-svg" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="3"></circle>
              <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
            </svg>
            Notification Setting
          </button>

          <!-- Full Screen with SVG Icon -->
          <button class="dropdown-item" @click="$emit('fullScreen'); showDropdown = false">
            <svg class="transparent-svg" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M8 3H5a2 2 0 0 0-2 2v3m18 0V5a2 2 0 0 0-2-2h-3m0 18h3a2 2 0 0 0 2-2v-3M3 16v3a2 2 0 0 0 2 2h3"></path>
            </svg>
            Full Screen
          </button>
        </div>
      </div>
    </div>

    <!-- Section: New -->
    <div class="section-header">
      <h2 class="section-title">New</h2>
      <button class="see-all-btn" @click="$emit('seeAllNew')">see all</button>
    </div>

    <div class="notification-list">
      <div 
        v-for="(item, index) in newNotifications" 
        :key="'new-' + index" 
        class="notification-card"
      >
        <!-- User Avatar with Badge Icon -->
        <div class="avatar-container">
          <div class="avatar-placeholder">
            <img v-if="item.profileImage" :src="item.profileImage" alt="Avatar" class="avatar-img" />
            <!-- SVG Default User Icon -->
            <svg v-else class="transparent-svg" viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
          <div class="action-badge-icon" :class="item.type">
            <!-- SVG Course/Book Badge Icon -->
            <svg class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2.5">
              <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"></path>
              <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path>
            </svg>
          </div>
        </div>

        <!-- Main Content Area -->
        <div class="content-wrapper">
          <div class="content-container">
            <div class="top-row">
              <span class="username">{{ item.username }}</span>
              <span class="message-text">{{ item.message }}</span>
            </div>
            <div class="time-badge">
              <!-- SVG Clock Icon -->
              <svg class="transparent-svg" viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
                <circle cx="12" cy="12" r="10"></circle>
                <polyline points="12 6 12 12 16 14"></polyline>
              </svg>
              {{ item.timeAgo }}
            </div>
          </div>
        </div>

        <!-- Right Side Icon -->
        <div class="right-icon">
          <span class="dot"></span>
          <span class="dot"></span>
        </div>
      </div>
    </div>

    <!-- Section: Follow -->
    <div class="section-header follower-header">
      <h2 class="section-title">Follow</h2>
      <button class="see-all-btn" @click="$emit('seeAllFollow')">see all</button>
    </div>

    <div class="notification-list">
      <div 
        v-for="(item, index) in followNotifications" 
        :key="'follow-' + index" 
        class="notification-card"
      >
        <!-- User Avatar with Plus Badge Icon -->
        <div class="avatar-container">
          <div class="avatar-placeholder">
            <img v-if="item.profileImage" :src="item.profileImage" alt="Avatar" class="avatar-img" />
            <!-- SVG Default User Icon -->
            <svg v-else class="transparent-svg" viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
          <div class="action-badge-icon follow">
            <!-- SVG Plus Icon -->
            <svg class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="3">
              <line x1="12" y1="5" x2="12" y2="19"></line>
              <line x1="5" y1="12" x2="19" y2="12"></line>
            </svg>
          </div>
        </div>

        <!-- Main Content Area -->
        <div class="content-wrapper">
          <div class="content-container">
            <div class="top-row">
              <span class="username">{{ item.username }}</span>
              <span class="message-text">{{ item.message }}</span>
            </div>
            <div class="time-badge">
              <!-- SVG Clock Icon -->
              <svg class="transparent-svg" viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
                <circle cx="12" cy="12" r="10"></circle>
                <polyline points="12 6 12 12 16 14"></polyline>
              </svg>
              {{ item.timeAgo }}
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="card-actions">
            <button class="follow-back-btn" @click.stop="$emit('followBack', item)">
              Follow
            </button>
            <button class="delete-btn" @click.stop="deleteFollow(index)" title="Delete">
              Delete
            </button>
          </div>
        </div>

        <!-- Right Side Icon -->
        <div class="right-icon">
          <span class="dot"></span>
          <span class="dot"></span>
        </div>
      </div>
    </div>

    <!-- Section: A Week -->
    <div class="section-header follower-header">
      <h2 class="section-title">A Week</h2>
      <button class="see-all-btn" @click="$emit('seeAllWeek')">see all</button>
    </div>

    <div class="notification-list">
      <div 
        v-for="(item, index) in weekNotifications" 
        :key="'week-' + index" 
        class="notification-card"
      >
        <div class="avatar-container">
          <div class="avatar-placeholder">
            <img v-if="item.profileImage" :src="item.profileImage" alt="Avatar" class="avatar-img" />
            <!-- SVG Default User Icon -->
            <svg v-else class="transparent-svg" viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
          <div class="action-badge-icon">
            <!-- SVG Calendar/Meeting Icon -->
            <svg class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2.5">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
              <line x1="16" y1="2" x2="16" y2="6"></line>
              <line x1="8" y1="2" x2="8" y2="6"></line>
              <line x1="3" y1="10" x2="21" y2="10"></line>
            </svg>
          </div>
        </div>

        <div class="content-wrapper">
          <div class="content-container">
            <div class="top-row">
              <span class="username">{{ item.username }}</span>
              <span class="message-text">{{ item.message }}</span>
            </div>
            <div class="time-badge">
              <!-- SVG Clock Icon -->
              <svg class="transparent-svg" viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
                <circle cx="12" cy="12" r="10"></circle>
                <polyline points="12 6 12 12 16 14"></polyline>
              </svg>
              {{ item.timeAgo }}
            </div>
          </div>
        </div>

        <div class="right-icon">
          <span class="dot"></span>
          <span class="dot"></span>
        </div>
      </div>
    </div>

    <!-- Footer Action -->
    <div class="footer-action">
      <button class="see-previous-btn" @click="$emit('seePrevious')">See previous notification</button>
    </div>
  </div>
</template>

<style scoped>
.transparent-svg {
  background-color: transparent !important;
}

.article-notification-container {
  width: 100%;
  background-color: #ffffff;
  color: #0f172a;
  display: flex;
  flex-direction: column;
  gap: 12px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  padding: 16px;
  box-sizing: border-box;
  border-radius: 12px;
}

/* Main Header */
.main-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #ffffff;
  padding: 6px 0px;
  position: relative;
}

.header-title-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
  color: #475569;
}

.main-title {
  font-size: 18px;
  font-weight: 700;
  margin: 0;
  color: #0f172a;
}

/* More Dots & Dropdown */
.header-dots-wrapper {
  position: relative;
  cursor: pointer;
}

.header-dots {
  display: flex;
  gap: 3px;
  padding: 8px;
  align-items: center;
  justify-content: center;
}

.header-dots span {
  width: 4px;
  height: 4px;
  background-color: #64748b;
  border-radius: 50%;
}

.dropdown-menu {
  position: absolute;
  right: 0;
  top: 35px;
  width: 210px;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -4px rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  padding: 6px;
  z-index: 50;
  animation: fadeIn 0.15s ease-in-out;
}

.dropdown-item {
  background: transparent;
  border: none;
  color: #334155;
  padding: 10px 12px;
  text-align: left;
  font-size: 13px;
  font-weight: 500;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 10px;
  transition: background-color 0.2s, color 0.2s;
}

.dropdown-item:hover {
  background-color: #f1f5f9;
  color: #1976d2;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-4px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Sections */
.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 4px;
  margin-top: 4px;
}

.follower-header {
  margin-top: 12px;
}

.section-title {
  font-size: 15px;
  font-weight: 600;
  color: #334155;
  margin: 0;
}

.see-all-btn {
  background: transparent;
  border: none;
  color: #1976d2;
  font-size: 11px;
  font-weight: 600;
  cursor: pointer;
  padding: 3px 10px;
  border-radius: 10px;
}

.see-all-btn:hover {
  opacity: 0.8;
}

.notification-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.notification-card {
  border: 1px solid #e2e8f085;
  border-radius: 12px;
  padding: 10px 12px;
  display: flex;
  align-items: flex-start;
  gap: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  background-color: #ffffff;
}

.notification-card:hover {
  transform: translateY(-2px);
  background-color: #f8fafc;
}

.avatar-container {
  position: relative;
  flex-shrink: 0;
}

.avatar-placeholder {
  width: 52px;
  height: 52px;
  border: 1px solid #cbd5e1;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #ffffff;
  color: #64748b;
  overflow: hidden;
}

.avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.action-badge-icon {
  position: absolute;
  bottom: -2px;
  right: -2px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid #ffffff;
  background-color: #1976d2;
  color: #ffffff;
}

.action-badge-icon.follow {
  background-color: #1976d2;
}

.content-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
  min-width: 0;
}

.content-container {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.top-row {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 6px;
}

.username {
  color: #0f172a;
  font-size: 13px;
  font-weight: 700;
  background-color: transparent;
  border: none;
  padding: 0;
}

.message-text {
  font-size: 13px;
  color: #475569;
  background-color: transparent;
  border: none;
  padding: 0;
}

.time-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  background-color: #F1F5F9;
  border: 1px solid #cbd5e1;
  padding: 1px 8px;
  border-radius: 8px;
  font-size: 10px;
  color: #64748b;
  width: fit-content;
}

.card-actions {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 4px;
}

.follow-back-btn {
  background-color: #1976d2;
  border: none;
  color: #ffffff;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 11px;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.2s;
}

.follow-back-btn:hover {
  opacity: 0.9;
}

.delete-btn {
  background-color: transparent;
  border: 1px solid #cbd5e1;
  color: #475569;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 11px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;
}

.delete-btn:hover {
  background-color: #e2e8f0;
}

.right-icon {
  display: flex;
  gap: 2px;
  align-items: center;
  align-self: center;
  margin-left: auto;
  padding-left: 8px;
}

.dot {
  width: 6px;
  height: 9px;
  background-color: #1976d2;
  border-radius: 50px;
}

/* Footer */
.footer-action {
  margin-top: 4px;
  display: flex;
  justify-content: center;
}

.see-previous-btn {
  background-color: transparent;
  border: 1px solid #cbd5e1;
  color: #334155;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  width: auto;
  text-align: center;
  transition: background-color 0.2s;
}

.see-previous-btn:hover {
  background-color: #e2e8f0;
}
</style>