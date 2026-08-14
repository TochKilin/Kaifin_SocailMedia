<!-- File: CourseNotification.vue -->
<script setup>
import { ref } from 'vue'

defineEmits(['seeAllNew', 'seeAllFollow', 'seeAllWeek', 'followBack', 'seePrevious'])

const newNotifications = ref([
  {
    type: 'course',
    username: 'Ven Nary',
    message: 'ចាប់ផ្តើមរៀន Coding ជាមួយ Python!, ',
    timeAgo: '12m',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRAfFjNWjyhAMJhzeU-L2n9KFRheSfuN7XlqONxHYlaNA&s=10'
  },
  {
    type: 'course',
    username: 'Srey MonCode',
    message: 'Update course',
    timeAgo: '13m',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQZHxn2g3POCSoDXxTzTvSiA6VVkVf20NjZQsXk_o1-qg&s=10'
  },
  {
    type: 'course',
    username: 'Noch Phasin',
    message: 'ចាប់ផ្តើមរៀន Python ថ្ងៃនេះ!',
    timeAgo: '14m',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQR6zBMxz00HQRI_wegCYFWpB51kRo9bwEeoC_w7x8Eas3EE-iHtbdBAXE&s=10'
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
    username: 'Mong Mony',
    message: 'New Instructor you can follow',
    timeAgo: '12h',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSFhdSE_vaTvZJdaqultlMSsGicNITYgdofIKNCJOJomQ&s=10'
  }
])

const weekNotifications = ref([
  {
    username: 'Sopheak SreyDav',
    message: 'Have meeting',
    timeAgo: '1 week',
    profileImage: 'https://cdn.kiripost.com/static/images/Lin.width-960.jpg'
  }
])

const deleteFollow = (index) => {
  followNotifications.value.splice(index, 1)
}
</script>

<template>
  <div class="article-notification-container">
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
            <svg v-else viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
          <div class="action-badge-icon" :class="item.type">
            <svg viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2.5">
              <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
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
              <svg viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
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
            <svg v-else viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
          <div class="action-badge-icon follow">
            <svg viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="3">
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
              <svg viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
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
            <svg v-else viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
          <div class="action-badge-icon">
            <svg viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2.5">
              <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
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
              <svg viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
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
      <button class="see-previous-btn" @click="$emit('seePrevious')">See more notification</button>
    </div>
  </div>
</template>

<style scoped>
.article-notification-container {
  width: 100%;
  background-color: #ffffff;
  color: #0f172a;
  display: flex;
  flex-direction: column;
  gap: 12px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  padding: 0px;
  box-sizing: border-box;
  border-radius: 12px;
}

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
  border: 1px solid #ffffff;
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
  opacity: 0.8;
}
</style>