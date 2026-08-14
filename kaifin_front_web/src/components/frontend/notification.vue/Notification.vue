<!-- File: Notification.vue -->
<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import Phea from "@/assets/team_images/admin.jpg"
import Heng from "@/assets/team_images/heng.jpg"
import LyHeng from "@/assets/team_images/lyheng.jpg"
import Mean from "@/assets/team_images/mean.jpg"
import Nary from "@/assets/team_images/nary.jpg"
import Nita from "@/assets/team_images/nita.jpg"
import Pa from "@/assets/team_images/pa.jpg"
import SoPhea from "@/assets/team_images/phea.jpg"
import Roth from "@/assets/team_images/roth.jpg"
import yuri from "@/assets/team_images/yuri.jpg"
import Kaifin from "@/assets/logos/kaifin_l2.png"

// Import Component ទាំងពីរមកទីនេះ
import ArticleNotification from './ArticleNotification.vue'
import NotificationCourse from './NotificationCourse.vue'

const emit = defineEmits(['back', 'close'])

const activeTab = ref('Feed')

// State សម្រាប់គ្រប់គ្រងການបង្ហាញ Dropdown Menu របស់ More
const showDropdown = ref(false)

const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

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

// Functions សម្រាប់ព្រឹត្តិការណ៍របស់ Menu ទាំងបី
const handleAllPost = () => {
  console.log('All Post clicked')
}

const handleNotificationSetting = () => {
  console.log('Notification Setting clicked')
}

// State និង Function សម្រាប់គ្រប់គ្រង Full Screen
const isFullScreen = ref(false)
const handleFullScreen = () => {
  isFullScreen.value = !isFullScreen.value
  console.log('Full Screen toggled:', isFullScreen.value)
}

const notifications = ref({
  new: [
    {
      id: 11,
      username: 'OTres Technology',
      avatar: Kaifin,
      type: 'announcement',
      text: 'announced a new platform update',
      time: 'Just now',
      isRead: false,
      announcementImage: "", 
      announcementText: '🎉 Special platform update! Explore our newest features and enhanced user experience designed for your daily workflow.',
      buttonText: 'Update Version 2.0'
    },
    {
      id: 1,
      username: 'សុខ សុភា',
      avatar: Phea,
      type: 'react',
      text: 'reacted to your post',
      time: '12m',
      isRead: true
    },
    {
      id: 2,
      username: 'Ly Heng',
      avatar: Heng,
      type: 'comment',
      text: 'commented on your article',
      time: '13m',
      isRead: false
    },
    {
      id: 3,
      username: 'Ly Heng Zin',
      avatar: LyHeng,
      type: 'share',
      text: 'shared your course',
      time: '14m',
      isRead: true
    },
    {
      id: 4,
      username: 'Mean',
      avatar: Mean,
      type: 'enroll',
      text: 'enrolled in your course',
      time: '15m',
      isRead: false
    }
  ],
  follower: [
    {
      id: 5,
      username: 'Nary Sovan',
      avatar: '',
      type: 'follow',
      text: 'started following you',
      time: '12h',
      isRead: false
    },
    {
      id: 8,
      username: 'Roth Roth',
      avatar: Roth,
      type: 'follow',
      text: 'started following you',
      time: '12h',
      isRead: false
    },
    {
      id: 10,
      username: 'Yuri Babo',
      avatar: yuri,
      type: 'follow',
      text: 'started following you',
      time: '12h',
      isRead: false
    },
    {
      id: 10,
      username: 'Yuri Babo',
      avatar: yuri,
      type: 'follow',
      text: 'started following you',
      time: '12h',
      isRead: false
    },
    {
      id: 10,
      username: 'Yuri Babo',
      avatar: yuri,
      type: 'follow',
      text: 'started following you',
      time: '12h',
      isRead: false
    },
    {
      id: 10,
      username: 'Yuri Babo',
      avatar: yuri,
      type: 'follow',
      text: 'started following you',
      time: '12h',
      isRead: false
    }
  ],
  aWeek: [
    {
      id: 6,
      username: 'Nita Ta',
      avatar: Nita,
      type: 'birthday',
      text: 'your friend has a birthday',
      time: '1 week',
      isRead: true
    },
    {
      id: 6,
      username: 'Nita Ta',
      avatar: Nita,
      type: 'birthday',
      text: 'your friend has a birthday',
      time: '1 week',
      isRead: true
    },
    {
      id: 6,
      username: 'Nita Ta',
      avatar: Nita,
      type: 'birthday',
      text: 'your friend has a birthday',
      time: '1 week',
      isRead: true
    },
    {
      id: 6,
      username: 'Nita Ta',
      avatar: Nita,
      type: 'birthday',
      text: 'your friend has a birthday',
      time: '1 week',
      isRead: true
    }
  ]
})

const markAsRead = (item) => {
  item.isRead = true
}
</script>

<template>
  <!-- បន្ថែម Class fullscreen-mode ទៅតាមតម្លៃ biến isFullScreen -->
  <div class="notification-container" :class="{ 'fullscreen-mode': isFullScreen }">
    
    <!-- Top Header -->
    <div class="notif-header">
      <div class="header-title-wrapper">
        <span class="bell-icon-box">
          <svg class="transparent-svg" viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
            <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
          </svg>
        </span>
        <h2>Notifications</h2>
      </div>
      <div class="header-right-actions">
        
        <!-- More Dots Button & Dropdown Menu -->
        <div class="header-dots-wrapper">
          <div class="header-dots" @click.stop="toggleDropdown">
            <span></span><span></span><span></span>
          </div>

          <div v-if="showDropdown" class="dropdown-menu">
            <button class="dropdown-item" @click="handleAllPost(); showDropdown = false">
              All Post
            </button>
            <button class="dropdown-item" @click="handleNotificationSetting(); showDropdown = false">
              Notification Setting
            </button>
            <!-- ប្ដូរអត្ថបទរវាង Full Screen និង Exit Full Screen ស្វ័យប្រវត្តិ -->
            <button class="dropdown-item" @click="handleFullScreen(); showDropdown = false">
              {{ isFullScreen ? 'Exit Full Screen' : 'Full Screen' }}
            </button>
          </div>
        </div>

        <button class="close-btn" @click="$emit('close')" aria-label="Close">
          <svg class="transparent-svg" viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>
    </div>

    <!-- Navigation Tabs -->
    <div class="tabs-row">
      <button 
        class="tab-btn"
        :class="{ active: activeTab === 'Feed' }"
        @click="activeTab = 'Feed'"
      >
        <svg class="transparent-svg tab-icon" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <rect x="3" y="3" width="18" height="18" rx="4" ry="4"></rect>
          <line x1="3" y1="11" x2="21" y2="11"></line>
          <circle cx="8" cy="7" r="2" fill="currentColor"></circle>
          <line x1="13" y1="6" x2="19" y2="6" stroke-width="2.5" stroke-linecap="round"></line>
          <line x1="13" y1="8.5" x2="19" y2="8.5" stroke-width="2.5" stroke-linecap="round"></line>
        </svg>
        Feed
      </button>

      <button 
        class="tab-btn"
        :class="{ active: activeTab === 'Article' }"
        @click="activeTab = 'Article'"
      >
        <svg class="transparent-svg tab-icon" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
          <polyline points="14 2 14 8 20 8"></polyline>
          <line x1="16" y1="13" x2="8" y2="13"></line>
          <line x1="16" y1="17" x2="8" y2="17"></line>
          <polyline points="10 9 9 9 8 9"></polyline>
        </svg>
        Article
      </button>

      <button 
        class="tab-btn"
        :class="{ active: activeTab === 'Course' }"
        @click="activeTab = 'Course'"
      >
        <svg class="transparent-svg tab-icon" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"></path>
          <path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"></path>
        </svg>
        Course
      </button>
    </div>

    <!-- ករណីទី១៖ Tab 'Feed' -->
    <template v-if="activeTab === 'Feed'">
      <!-- SECTION: New -->
      <div class="section-block">
        <div class="section-header">
          <h3>New</h3>
          <button class="see-all-btn">see all</button>
        </div>

        <div class="notif-list">
          <div 
            v-for="item in notifications.new" 
            :key="item.id" 
            class="notif-card" 
            :class="{ 'announcement-card-layout': item.type === 'announcement' }"
            @click="markAsRead(item)"
          >
            <!-- ករណីជា Admin Announcement -->
            <template v-if="item.type === 'announcement'">
              <div class="announcement-content-wrapper">
                <div class="announcement-top">
                  <div class="avatar-box" style="width: 40px; height: 40px;">
                    <img v-if="item.avatar" :src="item.avatar" alt="Avatar" class="avatar-img" />
                  </div>
                  <div class="announcement-meta">
                    <span class="username">{{ item.username }}</span>
                    <span class="time-badge">
                      <svg class="transparent-svg" viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
                        <circle cx="12" cy="12" r="10"></circle>
                        <polyline points="12 6 12 12 16 14"></polyline>
                      </svg>
                      {{ item.time }}
                    </span>
                  </div>
                </div>

                <div v-if="item.announcementImage" class="announcement-image-box">
                  <img :src="item.announcementImage" alt="Gift Bag" />
                </div>

                <p v-if="item.announcementText" class="announcement-text-desc">
                  {{ item.announcementText }}
                </p>

                <button v-if="item.buttonText" class="announcement-action-btn" @click.stop>
                  {{ item.buttonText }}
                </button>
              </div>
            </template>

            <!-- ករណី Notification ធម្មតា -->
            <template v-else>
              <div class="user-avatar-wrapper">
                <div class="avatar-box">
                  <img v-if="item.avatar" :src="item.avatar" alt="Avatar" class="avatar-img" />
                  <svg v-else class="transparent-svg" viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                    <circle cx="12" cy="7" r="4"></circle>
                  </svg>
                </div>
                
                <div class="badge-icon" :class="{ 'react-badge': item.type === 'react' }">
                  <svg v-if="item.type === 'react'" class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="currentColor" stroke="none">
                    <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
                  </svg>
                  <svg v-else-if="item.type === 'comment'" class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path>
                  </svg>
                  <svg v-else-if="item.type === 'share'" class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2">
                    <polyline points="15 3 21 3 21 9"></polyline>
                    <path d="M10 14L21 3"></path>
                    <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"></path>
                  </svg>
                  <svg v-else-if="item.type === 'enroll'" class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"></path>
                    <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path>
                  </svg>
                </div>
              </div>

              <div class="notif-content">
                <div class="top-row">
                  <span class="username">{{ item.username }}</span>
                  <span class="action-text">{{ item.text }}</span>
                </div>
                <div class="time-badge">
                  <svg class="transparent-svg" viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
                    <circle cx="12" cy="12" r="10"></circle>
                    <polyline points="12 6 12 12 16 14"></polyline>
                  </svg>
                  {{ item.time }}
                </div>
              </div>

              <div class="unread-dots" :class="{ 'is-read': item.isRead }">
                <span class="dot"></span>
                <span class="dot"></span>
              </div>
            </template>
          </div>
        </div>
      </div>

      <!-- SECTION: Follower -->
      <div class="section-block">
        <div class="section-header">
          <h3>Follower</h3>
          <button class="see-all-btn">see all</button>
        </div>

        <div class="notif-list">
          <div 
            v-for="item in notifications.follower" 
            :key="item.id" 
            class="notif-card follower-card" 
            @click="markAsRead(item)"
          >
            <div class="user-avatar-wrapper">
              <div class="avatar-box">
                <img v-if="item.avatar" :src="item.avatar" alt="Avatar" class="avatar-img" />
                <svg v-else class="transparent-svg" viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                  <circle cx="12" cy="7" r="4"></circle>
                </svg>
              </div>
              <div class="badge-icon plus-badge">
                <svg class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="3">
                  <line x1="12" y1="5" x2="12" y2="19"></line>
                  <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
              </div>
            </div>

            <div class="notif-content">
              <div class="top-row">
                <span class="username">{{ item.username }}</span>
                <span class="action-text">{{ item.text }}</span>
              </div>
              <div class="time-badge">
                <svg class="transparent-svg" viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                {{ item.time }}
              </div>

              <div class="follower-actions">
                <button class="follow-back-btn" @click.stop>Follow back</button>
                <button class="delete-btn" @click.stop>Delete</button>
              </div>
            </div>

            <div class="unread-dots" :class="{ 'is-read': item.isRead }">
              <span class="dot"></span>
              <span class="dot"></span>
            </div>
          </div>
        </div>
      </div>

      <!-- SECTION: A Week -->
      <div class="section-block">
        <div class="section-header">
          <h3>A Week</h3>
          <button class="see-all-btn">see all</button>
        </div>

        <div class="notif-list">
          <div 
            v-for="item in notifications.aWeek" 
            :key="item.id" 
            class="notif-card" 
            @click="markAsRead(item)"
          >
            <div class="user-avatar-wrapper">
              <div class="avatar-box">
                <img v-if="item.avatar" :src="item.avatar" alt="Avatar" class="avatar-img" />
                <svg v-else class="transparent-svg" viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                  <circle cx="12" cy="7" r="4"></circle>
                </svg>
              </div>
              <div class="badge-icon">
                <svg class="transparent-svg" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="10"></circle>
                  <path d="M8 14s1.5 2 4 2 4-2 4-2"></path>
                </svg>
              </div>
            </div>

            <div class="notif-content">
              <div class="top-row">
                <span class="username">{{ item.username }}</span>
                <span class="action-text">{{ item.text }}</span>
              </div>
              <div class="time-badge">
                <svg class="transparent-svg" viewBox="0 0 24 24" width="10" height="10" fill="none" stroke="currentColor" stroke-width="2.5">
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                {{ item.time }}
              </div>
            </div>

            <div class="unread-dots" :class="{ 'is-read': item.isRead }">
              <span class="dot"></span>
              <span class="dot"></span>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer Button -->
      <div class="footer-action">
        <button class="see-previous-btn">See more notification</button>
      </div>
    </template>

    <!-- ករណីទីពីរ៖ Tab 'Article' -->
    <template v-else-if="activeTab === 'Article'">
      <div class="article-tab-content">
        <ArticleNotification @previous="activeTab = 'Feed'" @close="$emit('close')" />
      </div>
    </template>

    <!-- ករណីទីបី៖ Tab 'Course' -->
    <template v-else-if="activeTab === 'Course'">
      <div class="course-tab-content">
        <NotificationCourse @previous="activeTab = 'Feed'" @close="$emit('close')" />
      </div>
    </template>

  </div>
</template>

<style scoped>
.transparent-svg {
  background-color: transparent !important;
}

/* ទំហំធម្មតា */
.notification-container {
  background-color: #ffffff;
  color: #0f172a;
  padding: 18px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  width: 1090px;
  max-height: 95vh;
  overflow-y: auto;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  transition: all 0.3s ease-in-out;
}

/* ទំហំពេលចុច Full Screen (ម៉ាល្មមស្អាត ធំទូលាយល្មមមិនបាំងជិតអេក្រង់ខ្លាំងពេក) */
.notification-container.fullscreen-mode {
  width: 1100px;
  max-height: 92vh;
}

.notif-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-title-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
  background-color: transparent;
  padding: 6px 0px;
}

.bell-icon-box {
  color: #475569;
  display: flex;
  align-items: center;
}

.notif-header h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
}

.header-right-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* Header Dots & Dropdown Styles */
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
  top: 26px;
  width: 180px;
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

.close-btn {
  background-color: transparent;
  border: none;
  color: #64748b;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 4px;
  border-radius: 50%;
  transition: background-color 0.2s ease, color 0.2s ease;
}

.close-btn:hover {
  background-color: #f1f5f9;
  color: #0f172a;
}

.tabs-row {
  display: flex;
  gap: 24px;
  border-bottom: 1px solid #e2e8f0;
  padding-bottom: 4px;
}

.tab-btn {
  background-color: transparent;
  border: none;
  color: #64748b;
  padding: 8px 0px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  position: relative;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: color 0.2s ease;
}

.tab-icon {
  color: inherit;
}

.tab-btn:hover {
  color: #0f172a;
}

.tab-btn.active {
  background-color: transparent;
  color: #1976d2;
}

.tab-btn.active::after {
  content: '';
  position: absolute;
  bottom: -5px;
  left: 0;
  width: 100%;
  height: 3px;
  background-color: #1976d2;
  border-radius: 2px 2px 0 0;
}

.section-block {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.section-header h3 {
  margin: 0;
  font-size: 15px;
  font-weight: 600;
  color: #334155;
}

.see-all-btn {
  border: none;
  color: #1976d2;
  font-size: 11px;
  padding: 3px 10px;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 600;
  background-color: transparent;
}

.see-all-btn:hover {
  opacity: 0.8;
}

.notif-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.notif-card {
  border: 1px solid #e2e8f085;
  border-radius: 12px;
  padding: 10px 12px;
  display: flex;
  align-items: flex-start;
  gap: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.notif-card:hover {
  transform: translateY(-2px);
}

/* Style សម្រាប់ Announcement Card */
.announcement-card-layout {
  flex-direction: column !important;
  align-items: stretch !important;
  gap: 12px !important;
  background-color: #ffffff;
}

.announcement-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 100%;
}

.announcement-top {
  display: flex;
  align-items: center;
  gap: 10px;
}

.announcement-meta {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.announcement-image-box {
  width: 100%;
  display: flex;
  justify-content: flex-start;
  background-color: transparent;
  border: none;
  border-radius: 8px;
  overflow: hidden;
  padding: 0px;
}

.announcement-image-box img {
  max-width: 100%;
  max-height: 240px;
  object-fit: contain;
}

.announcement-text-desc {
  margin: 0;
  font-size: 13px;
  color: #334155;
  line-height: 1.5;
  text-align: left;
  font-weight: bold;
}

.announcement-action-btn {
  background-color: #1976d2;
  color: #ffffff;
  border: none;
  width: fit-content;
  min-width: 220px;
  margin: 0;
  padding: 10px 24px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  text-align: center;
  display: block;
  transition: opacity 0.2s ease;
}

.announcement-action-btn:hover {
  opacity: 0.9;
}

.unread-dots {
  display: flex;
  flex-direction: row;
  gap: 2px;
  align-items: center;
  justify-content: center;
  align-self: center;
  margin-left: auto;
  padding-left: 8px;
}

.dot {
  width: 6px;
  height: 9px;
  background-color: #1976d2;
  border-radius: 50px;
  transition: background-color 0.2s ease;
}

.unread-dots.is-read .dot {
  background-color: #cbd5e1d2;
}

.user-avatar-wrapper {
  position: relative;
  flex-shrink: 0;
}

.avatar-box {
  width: 52px;
  height: 52px;
  border: 1px solid #cbd5e1;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #ffffff;
  overflow: hidden;
  color: #64748b;
}

.avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.badge-icon {
  position: absolute;
  bottom: -2px;
  right: -2px;
  background-color: #1976d2;
  color: #ffffff;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid #ffffff;
}

.badge-icon.react-badge {
  background-color: #ef4444 !important;
}

.notif-content {
  flex: 1;
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
}

.action-text {
  font-size: 13px;
  color: #475569;
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

.follower-actions {
  display: flex;
  gap: 8px;
  margin-top: 4px;
}

.follow-back-btn {
  background-color: #1976d2;
  color: #ffffff;
  border: none;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 11px;
  font-weight: 600;
  cursor: pointer;
}

.delete-btn {
  background-color: transparent;
  color: #475569;
  border: 1px solid #cbd5e1;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 11px;
  font-weight: 600;
  cursor: pointer;
}

.delete-btn:hover {
  background-color: #e2e8f0;
}

.footer-action {
  display: flex;
  justify-content: center;
  margin-top: 4px;
}

.see-previous-btn {
  background-color: transparent;
  border: transparent;
  color: #334155;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  width: auto;
  text-align: center;
}

.see-previous-btn:hover {
  background-color: #e2e8f0;
}

.article-tab-content,
.course-tab-content {
  width: 100%;
}
</style>