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

        <!-- Right Side Icon (Dots) -->
        <div class="right-icon">
          <span class="dot"></span>
          <span class="dot"></span>
        </div>
      </div>
    </div>

    <!-- Section: Follower -->
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
.article-notification-container {
  width: 100%;
  background-color: #ffffff;
  color: #0f172a;
  display: flex;
  flex-direction: column;
  gap: 12px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 4px;
}

.follower-header {
  margin-top: 16px;
}

.section-title {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

.see-all-btn {
  background: none;
  border: none;
  color: #1976d2;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  padding: 0;
}

.see-all-btn:hover {
  text-decoration: underline;
}

.notification-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.notification-card {
  border: 1px solid #e2e8f085;
  border-radius: 12px;
  padding: 14px 16px;
  display: flex;
  align-items: center;
  gap: 14px;
  cursor: pointer;
  transition: all 0.2s ease;
  background-color: #ffffff;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.02);
}

.notification-card:hover {
  transform: translateY(-2px);
}

.avatar-container {
  position: relative;
  flex-shrink: 0;
}

.avatar-placeholder {
  width: 48px;
  height: 48px;
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
  width: 18px;
  height: 18px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid #ffffff;
  color: #ffffff;
}

.action-badge-icon.follow {
  background-color: #1976d2;
}

.action-badge-icon.react {
  background-color: #e11d48;
}

.action-badge-icon.comment {
  background-color: #0284c7;
}

.action-badge-icon.share {
  background-color: #0284c7;
}

.action-badge-icon.enroll {
  background-color: #0284c7;
}

.content-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10px;
  min-width: 0;
}

.content-container {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.top-row {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 6px;
}

.username {
  color: #0f172a;
  font-size: 14px;
  font-weight: 700;
}

.message-text {
  font-size: 14px;
  color: #475569;
}

.time-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  background-color: #F1F5F9;
  border: 1px solid #cbd5e1;
  padding: 2px 12px;
  border-radius: 32px;
  font-size: 11px;
  color: #64748b;
  width: fit-content;
}

.card-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.follow-back-btn {
  background-color: #1976d2;
  border: none;
  color: #ffffff;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;
}

.follow-back-btn:hover {
  opacity: 0.8;
}

.delete-btn {
  background-color: transparent;
  border: 1px solid #cbd5e1;
  color: #334155;
  padding: 8px 12px;
  border-radius: 32px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.delete-btn:hover {
  background-color: #fee2e2;
  border-color: #fca5a5;
  color: #dc2626;
}

.right-icon {
  display: flex;
  gap: 3px;
  align-items: center;
  align-self: center;

}

.dot {
  width: 5px;
  height: 8px;
  background-color: #1976d2;
  border-radius: 50px;
}

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