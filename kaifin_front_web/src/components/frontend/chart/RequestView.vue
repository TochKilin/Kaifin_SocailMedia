<template>
  <div class="request-container">
    <!-- Header -->
    <div class="request-header">
      <div class="header-left" @click="$emit('back')">
        <button class="back-btn">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>
        </button>
        <h1 class="request-title">Request</h1>
      </div>
      <div class="header-actions">
        <button class="action-icon-btn">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M15 3h6v6M9 21H3v-6M21 3l-7 7M3 21l7-7"/></svg>
        </button>
      </div>
    </div>

    <!-- Request List Body -->
    <div class="request-list">
      <div v-for="(req, index) in requests" :key="index" class="request-item">
        <!-- Avatar -->
        <div class="request-avatar-wrapper">
          <div class="request-avatar">
            <img v-if="req.avatar" :src="req.avatar" :alt="req.name" class="avatar-img" />
            <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
          </div>
        </div>

        <!-- Request Info -->
        <div class="request-info">
          <div class="request-info-header">
            <span class="request-username">{{ req.name }}</span>
            <div class="more-options-dots">
              <span class="mini-dot"></span>
              <span class="mini-dot"></span>
              <span class="mini-dot"></span>
            </div>
          </div>
          <div class="request-bottom-row">
            <span class="request-msg" :class="{ 'voice-msg': req.isVoice }">{{ req.message }}</span>
            <span class="request-time">{{ req.time }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

defineEmits(['back'])

const requests = ref([
  {
    name: 'Sokha Chea',
    message: 'Show text massage',
    time: '12:30 PM',
    isVoice: false,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRlzLSIDO3ceFnP3Qyu0o6sNLmq7_A6l7JLSdJcS9pc7g&s=10'
  },
  {
    name: 'Dany Kim',
    message: 'Sent a voice message (0:14)',
    time: '11:15 AM',
    isVoice: true,
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-971-20200304-4-ro0jp2.jpg'
  },
  {
    name: 'Vanna Roth',
    message: 'Show text massage',
    time: 'Yesterday',
    isVoice: false,
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRvT9fXBRHRiiYmAa1-jNBmANV3qYdXeXhrOOCJBrtTVg&s=10"
  },
  {
    name: 'Minea Chan',
    message: 'Sent a voice message (0:22)',
    time: 'Yesterday',
    isVoice: true,
    avatar: null
  },
  {
    name: 'Kosal Sorn',
    message: 'Sent a voice message (0:08)',
    time: 'Aug 3',
    isVoice: true,
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRCVVpESqgjipQIDaVCvZEhNZYd3f-6CgR-T31ca_xGRw&s=10"
  }
])
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');

.request-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: #ffffff;
  box-sizing: border-box;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  padding: 14px 20px;
}

.request-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 16px;
  border-bottom: 1px solid #e4e6eb;
  margin-bottom: 12px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
}

.back-btn {
  background: transparent;
  border: none;
  cursor: pointer;
  color: #050505;
  display: flex;
  align-items: center;
  padding: 4px;
  border-radius: 50%;
  transition: background-color 0.2s;
}
.back-btn:hover {
  background-color: #f2f2f2;
}

.request-title {
  font-size: 22px;
  font-weight: 700;
  color: #050505;
  margin: 0;
}

.header-actions {
  display: flex;
  gap: 8px;
}

.action-icon-btn {
  background: transparent;
  border: none;
  cursor: pointer;
  color: #050505;
  padding: 6px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  transition: background-color 0.2s;
}
.action-icon-btn:hover {
  background-color: #f2f2f2;
}

/* Request List */
.request-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  overflow-y: auto;
  flex-grow: 1;
  padding-right: 4px;
}

.request-item {
  background-color: #ffffff;
  border: 1px solid #eee;
  border-radius: 12px;
  padding: 10px 14px;
  display: flex;
  align-items: center;
  gap: 14px;
  cursor: pointer;
  transition: background 0.15s;
}

.request-item:hover {
  background-color: #f9f9f9;
}

.request-avatar-wrapper {
  flex-shrink: 0;
}

.request-avatar {
  width: 48px;
  height: 48px;
  background-color: #e3f2fd;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #1976d2;
  overflow: hidden;
}

.avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.request-info {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.request-info-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.request-username {
  color: #222222;
  font-size: 15px;
  font-weight: 600;
}

.more-options-dots {
  display: flex;
  gap: 3px;
}

.mini-dot {
  width: 4px;
  height: 4px;
  background-color: #bbb;
  border-radius: 50%;
}

.request-bottom-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.request-msg {
  color: #666666;
  font-size: 14px;
}

.request-msg.voice-msg {
  color: #f57c00;
  font-weight: 500;
}

.request-time {
  color: #999999;
  font-size: 13px;
}
</style>