<template>
  <div class="article-notification-container">
    <!-- Section New -->
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
            <!-- Dynamic icons based on type -->
            <svg v-if="item.type === 'react'" viewBox="0 0 24 24" width="12" height="12" fill="currentColor" stroke="none">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
            <svg v-else-if="item.type === 'comment'" viewBox="0 0 24 24" width="12" height="12" fill="currentColor" stroke="none">
              <path d="M20 2H4c-1.1 0-2 .9-2 2v18l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2z"/>
            </svg>
            <svg v-else-if="item.type === 'share'" viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2.5">
              <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path>
              <polyline points="16 6 12 2 8 6"></polyline>
              <line x1="12" y1="2" x2="12" y2="15"></line>
            </svg>
            <svg v-else viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2.5">
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

    <!-- Section Follower -->
    <div class="section-header follower-header">
      <h2 class="section-title">Follower</h2>
      <button class="see-all-btn" @click="$emit('seeAllFollower')">see all</button>
    </div>

    <div class="notification-list">
      <div 
        v-for="(item, index) in followerNotifications" 
        :key="'follower-' + index" 
        class="notification-card"
      >
        <!-- User Avatar with-->
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
              Follow back
            </button>
            <button class="delete-btn" @click.stop="deleteFollower(index)" title="Delete">
              Delete
            </button>
          </div>
        </div>

        <!-- Right Side Icon (Dots) -->
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

<script setup>
import { ref } from 'vue'

defineEmits(['seeAllNew', 'seeAllFollower', 'followBack'])

const newNotifications = ref([
  {
    type: 'react',
    username: 'សី សីុកា',
    message: 'reacted to your post',
    timeAgo: '12m',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRHl3jp1gX_Y_k79ShEt6Rxet3AGCC0j2G-VgLVF6UUbde6rb1mYYJD_KQ&s=10'
  },
  {
    type: 'comment',
    username: 'Ly Heng',
    message: 'commented on your article',
    timeAgo: '13m',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSPdf6ZUWk9JoqCHl3LWZp1tKUb74YX9Ku_vrChs0uvU_zberBOwG6rUSA&s=10'
  },
  {
    type: 'share',
    username: 'Ly Heng Zin',
    message: 'shared your course',
    timeAgo: '14m',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQuUFwfa3MQ_svuorOuJ-aURHw-al2WayoSPUZ3shSBTUceLS1dGMfqP4c&s=10'
  },
  {
    type: 'enroll',
    username: 'Mean',
    message: 'enrolled in your course',
    timeAgo: '15m',
    profileImage: 'https://media.licdn.com/dms/image/v2/D5603AQEXEdGLPUTabw/profile-displayphoto-scale_400_400/B56Z1reuHPGoAk-/0/1775624709989?e=2147483647&v=beta&t=tp9dlXocXlkxK-hVej_uxrwittuA3yiRX-lA1akTdlk'
  },
  {
    type: 'react',
    username: 'Mai Malai',
    message: 'reacted to your post',
    timeAgo: '20m',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRHl3jp1gX_Y_k79ShEt6Rxet3AGCC0j2G-VgLVF6UUbde6rb1mYYJD_KQ&s=10'
  },
  {
    type: 'comment',
    username: 'Va Vondy',
    message: 'commented on your article',
    timeAgo: '32m',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSPdf6ZUWk9JoqCHl3LWZp1tKUb74YX9Ku_vrChs0uvU_zberBOwG6rUSA&s=10'
  },
  {
    type: 'share',
    username: 'AhZa AhXa',
    message: 'shared your course',
    timeAgo: '1h',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQuUFwfa3MQ_svuorOuJ-aURHw-al2WayoSPUZ3shSBTUceLS1dGMfqP4c&s=10'
  },
  {
    type: 'enroll',
    username: 'Udom PangThea',
    message: 'enrolled in your course',
    timeAgo: '2h',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ_ZWDQFG1-zwSrlKgaBB__Znl1Qd_yvKKzlg-VeaaOEA&s=10'
  }
])

const followerNotifications = ref([
  {
    username: 'Nary Sovan',
    message: 'started following you',
    timeAgo: '12h',
    profileImage: ''
  },
  {
    username: 'Roth Roth',
    message: 'started following you',
    timeAgo: '12h',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRBMlSC1pqRvEynrEUiSQGNI-Eqw9ftuCBOWah2qSykSg&s=10'
  },
  {
    username: 'Yuri Babo',
    message: 'started following you',
    timeAgo: '12h',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ_ZWDQFG1-zwSrlKgaBB__Znl1Qd_yvKKzlg-VeaaOEA&s=10'
  },
  {
    username: 'Vong VuthThea',
    message: 'started following you',
    timeAgo: '1d',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSfZvMX8nfvk9XPYxiuPRqmmU8OYpYOdF8QNrPHktGDpA&s=10'
  },
  {
    username: 'Nar Nar',
    message: 'started following you',
    timeAgo: '1d',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQplUwsYBXGSupgB-bKH7O1Bf97O5EgSFkBnZB7cxgCKw&s=10'
  },
  {
    username: 'Bros Har',
    message: 'started following you',
    timeAgo: '2d',
    profileImage: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR1c66WJConDWtSjTt6i22PDCGzfT4v79jjKPyjIGCY5O52-hLxAS9H7wc&s=10'
  }
])

const deleteFollower = (index) => {
  followerNotifications.value.splice(index, 1)
}
</script>
<style scoped>
.notif-wrap {
  position: relative;
  display: inline-flex;
}

.notif-btn {
  position: relative;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: none;
  background: #EFF6FB;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.notif-btn svg {
  width: 20px;
  height: 20px;
  stroke: #1E6E9C;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.notif-badge {
  position: absolute;
  top: -2px;
  right: -2px;
  background: #E8543A;
  color: #fff;
  font-size: 10px;
  font-weight: 700;
  min-width: 16px;
  height: 16px;
  padding: 0 4px;
  border-radius: 999px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.notif-dropdown {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  width: 340px;
  max-height: 420px;
  background: #fff;
  border-radius: 14px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .18);
  overflow: hidden;
  z-index: 100;
  display: flex;
  flex-direction: column;
}

.notif-header {
  padding: 14px 16px;
  font-weight: 700;
  font-size: 14px;
  border-bottom: 1px solid #F0F0F0;
  color: #2B2B2B;
}

.notif-list {
  overflow-y: auto;
  flex: 1;
}

.notif-state {
  text-align: center;
  color: #8A8A8E;
  font-size: 13px;
  padding: 24px 0;
}

.notif-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 16px;
  cursor: pointer;
  position: relative;
}

.notif-item:hover {
  background: #F8FAFC;
}

.notif-item.unread {
  background: #EFF6FB;
}

.notif-avatar {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: #EFF6FB;
  border: 2px solid #1976D2;
  overflow: hidden;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.notif-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.notif-avatar svg {
  width: 18px;
  height: 18px;
  stroke: #1E6E9C;
  fill: none;
  stroke-width: 1.8;
}

.notif-body {
  flex: 1;
  min-width: 0;
}

.notif-text {
  margin: 0;
  font-size: 13px;
  color: #2B2B2B;
  line-height: 1.4;
}

.notif-time {
  font-size: 11px;
  color: #8A8A8E;
}

.notif-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #1976D2;
  flex-shrink: 0;
}

.toast-stack {
  position: fixed;
  top: 20px;
  right: 20px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  z-index: 9999;
}

.toast-item {
  display: flex;
  align-items: center;
  gap: 10px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .18);
  padding: 12px 16px;
  width: 300px;
  cursor: pointer;
  animation: toastIn .2s ease;
}

@keyframes toastIn {
  from { opacity: 0; transform: translateX(20px); }
  to { opacity: 1; transform: translateX(0); }
}

.toast-avatar {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  background: #EFF6FB;
  overflow: hidden;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.toast-avatar img { width: 100%; height: 100%; object-fit: cover; }
.toast-avatar svg { width: 16px; height: 16px; stroke: #1E6E9C; fill: none; stroke-width: 1.8; }

.toast-text {
  margin: 0;
  font-size: 13px;
  color: #2B2B2B;
  line-height: 1.4;
}
</style>