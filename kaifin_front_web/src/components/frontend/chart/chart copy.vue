<template>
  <div class="chats-modal-wrapper">
    <div class="chats-container">
      
      <!-- Header Section -->
      <div class="chats-header">
        <div class="chats-title-box">
          <h1 class="chats-title">Chats</h1>
        </div>
        <div class="header-controls">
          <div class="window-actions">
            <button class="action-icon-btn">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><circle cx="5" cy="12" r="2"></circle><circle cx="12" cy="12" r="2"></circle><circle cx="19" cy="12" r="2"></circle></svg>
            </button>
            <button class="action-icon-btn close-btn" @click="$emit('close')">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Content Body -->
      <div class="chats-body">
        <!-- Search Bar -->
        <div class="search-box">
          <span class="search-icon">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
          </span>
          <input type="text" class="search-input" placeholder="Search Messenger..." v-model="searchQuery" />
        </div>

        <!-- New Message Request Banner -->
        <div class="request-banner">
          <div class="request-left">
            <div class="request-icon-box">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
            </div>
            <span class="request-text">New message request (2)</span>
          </div>
          <span class="arrow-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"></polyline></svg>
          </span>
        </div>

        <!-- Chat Items List with Mock Data & Avatars -->
        <div class="chat-list">
          <div 
            v-for="(chat, index) in filteredChats" 
            :key="index" 
            class="chat-item"
          >
            <!-- Avatar -->
            <div class="avatar-wrapper">
              <div class="avatar" :class="{ 'group-avatar': chat.isGroup }">
                <img v-if="chat.avatar" :src="chat.avatar" :alt="chat.name" class="avatar-img" />
                <svg v-else-if="!chat.isGroup" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
              </div>
              <span v-if="chat.online" class="online-dot"></span>
              <span v-if="chat.badge" class="unread-badge">{{ chat.badge }}</span>
            </div>

            <!-- Chat Info -->
            <div class="chat-info">
              <div class="chat-info-header">
                <span class="chat-name">{{ chat.name }}</span>
                <div class="more-options-dots">
                  <span class="mini-dot"></span>
                  <span class="mini-dot"></span>
                  <span class="mini-dot"></span>
                </div>
              </div>
              <div class="chat-info-bottom">
                <span class="chat-message" :class="{ 'voice-msg': chat.isVoice }">{{ chat.message }}</span>
                <span class="chat-time">{{ chat.time }}</span>
              </div>
            </div>
          </div>

          <!-- Empty State -->
          <div v-if="filteredChats.length === 0" class="empty-state">
            <p>No chats found</p>
          </div>
        </div>

        <!-- See More Button -->
        <div class="see-more-container">
          <button class="see-more-btn">See more</button>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

defineEmits(['close'])

const searchQuery = ref('')

const chatList = ref([
  {
    name: 'Dara Chan',
    message: 'Hey! Are we still meeting for project review?',
    time: '10:45 AM',
    online: true,
    isGroup: false,
    isVoice: false,
    badge: null,
    unread: false,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRlzLSIDO3ceFnP3Qyu0o6sNLmq7_A6l7JLSdJcS9pc7g&s=10'
  },
  {
    name: 'Srey Nich',
    message: 'Sent a voice message (0:14)',
    time: '09:20 AM',
    online: true,
    isGroup: false,
    isVoice: true,
    badge: null,
    unread: false,
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-971-20200304-4-ro0jp2.jpg'
  },
  {
    name: 'Senior Capstone Group',
    message: 'Vuthea: I pushed the latest API updates to GitHub.',
    time: 'Yesterday',
    online: true,
    isGroup: true,
    isVoice: false,
    badge: '5',
    unread: true,
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-2201-20200924-4-92w8ra.jpg'
  },
  {
    name: 'Panha Rith',
    message: 'Check out this new component layout structure.',
    time: 'Yesterday',
    online: false,
    isGroup: false,
    isVoice: false,
    badge: '1',
    unread: true,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10'
  },
  {
    name: 'Web Dev Team',
    message: 'Meeting rescheduled to Monday afternoon.',
    time: 'Aug 5',
    online: true,
    isGroup: true,
    isVoice: false,
    badge: null,
    unread: false,
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-3856-20210419-4-ql7c88.jpg'
  },
    {
    name: 'Panha Rith',
    message: 'Check out this new component layout structure.',
    time: 'Yesterday',
    online: false,
    isGroup: false,
    isVoice: false,
    badge: '1',
    unread: true,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10'
  },
    {
    name: 'Panha Rith',
    message: 'Check out this new component layout structure.',
    time: 'Yesterday',
    online: false,
    isGroup: false,
    isVoice: false,
    badge: '1',
    unread: true,
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-627-20200214-4-rpp3cr.jpg'
  },
    {
    name: 'Panha Rith',
    message: 'Check out this new component layout structure.',
    time: 'Yesterday',
    online: false,
    isGroup: false,
    isVoice: false,
    badge: '1',
    unread: true,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10'
  },
    {
    name: 'Panha Rith',
    message: 'Check out this new component layout structure.',
    time: 'Yesterday',
    online: false,
    isGroup: false,
    isVoice: false,
    badge: '1',
    unread: true,
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-987-20201208-4-vd6umo.jpg'
  },
    {
    name: 'Panha Rith',
    message: 'Check out this new component layout structure.',
    time: 'Yesterday',
    online: false,
    isGroup: false,
    isVoice: false,
    badge: '1',
    unread: true,
    avatar: 'https://cdn.hypetrace.com/file/ht/avig/instagram-influencer-61355-20200302-4-1e8w3gk.jpg'
  },
    {
    name: 'Panha Rith',
    message: 'Check out this new component layout structure.',
    time: 'Yesterday',
    online: false,
    isGroup: false,
    isVoice: false,
    badge: '1',
    unread: true,
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-57374-20200316-4-1h7pxhv.jpg'
  },
])

const filteredChats = computed(() => {
  return chatList.value.filter(chat => {
    return chat.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
           chat.message.toLowerCase().includes(searchQuery.value.toLowerCase())
  })
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');

.chats-modal-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: transparent;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  /* padding: 20px; */
  padding: 0;
  box-sizing: border-box;

}

.chats-container {
  width: 790px;
  max-width: 100%;
  height: 100vh;
  background-color: #ffffff;
  border: 1px solid #e0e0e0;
  border-radius: 12px;
  overflow: hidden;
  padding: 0;
}

/* Header */
.chats-header {
  padding: 16px 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #ffffff;
  /* border-bottom: 1px solid #f0f0f0; */
}

.chats-title-box {
  background-color: transparent;
  border: none;
  padding: 0;
  box-shadow: none;
}

.chats-title {
  margin: 0;
  font-size: 24px;
  font-weight: 700;
  color: #222222;
}

.header-controls {
  display: flex;
  align-items: center;
  gap: 12px;
}

.window-actions {
  display: flex;
  align-items: center;
  gap: 4px;
}

.action-icon-btn {
  background: transparent;
  border: none;
  color: #65676b;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 6px;
  border-radius: 50%;
  transition: background-color 0.2s;
}

.action-icon-btn:hover {
  background-color: #f2f2f2;
}

.action-icon-btn.close-btn {
  color: #65676b;
}

/* Body Panel */
.chats-body {
  background-color: #ffffff;
  border: none;
  border-radius: 0;
  padding: 20px;
  height: 703px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* Search Bar */
.search-box {
  position: relative;
  background: transparent;
  border: 1px solid #e0e0e0;
  border-radius: 32px;
  display: flex;
  align-items: center;
  padding: 8px 14px;
}

.search-icon {
  color: #888;
  margin-right: 10px;
  display: flex;
}

.search-input {
  background: transparent;
  border: none;
  outline: none;
  color: #333;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  font-size: 15px;
  width: 100%;
}

.search-input::placeholder {
  color: #aaa;
}

/* New Message Request Banner */
.request-banner {
  background-color: #ffff;
  border-radius: 12px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
}

.request-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.request-icon-box {
  background-color: #ffffff;
  border: none;
  border-radius: 10px;
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #000;
}

.request-text {
  color: #1B75D2;
  font-size: 15px;
  font-weight: 500;
}

.arrow-icon {
  color: #888;
  display: flex;
}

/* Chat List */
.chat-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
  min-height: 180px;
}

.chat-item {
  background-color: #ffffff;
  border: 1px solid #eee;
  border-radius: 12px;
  padding: 10px 14px;
  display: flex;
  align-items: center;
  gap: 14px;
  position: relative;
  transition: background 0.15s;
}

.chat-item:hover {
  background-color: #f9f9f9;
}

.empty-state {
  text-align: center;
  color: #888;
  padding: 20px;
  font-size: 15px;
}

/* Avatar Styling */
.avatar-wrapper {
  position: relative;
  flex-shrink: 0;
}

.avatar {
  width: 48px;
  height: 48px;
  background-color: #e3f2fd;
  border: none;
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

.online-dot {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 10px;
  height: 10px;
  background-color: #4caf50;
  border: 2px solid #ffffff;
  border-radius: 6px;
}

.unread-badge {
  position: absolute;
  top: -4px;
  left: -4px;
  background-color: #2196f3;
  border: none;
  color: #fff;
  font-size: 11px;
  padding: 1px 6px;
  border-radius: 6px;
  font-weight: 600;
}

/* Chat Info Section */
.chat-info {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.chat-info-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: none;
  padding-bottom: 0;
}

.chat-name {
  color: #222;
  font-size: 15px;
  font-weight: 600;
}

.more-options-dots {
  display: flex;
  gap: 3px;
  background-color: transparent;
  border: none;
  padding: 2px;
  border-radius: 4px;
}

.mini-dot {
  width: 4px;
  height: 4px;
  background-color: #bbb;
  border-radius: 50%;
}

.chat-info-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chat-message {
  color: #666;
  font-size: 14px;
}

.chat-message.voice-msg {
  color: #f57c00;
  font-weight: 500;
}

.chat-time {
  color: #999;
  font-size: 13px;
}

/* See More Button */
.see-more-container {
  display: flex;
  justify-content: center;
  margin-top: 4px;
}

.see-more-btn {
  background-color: #f5f5f5;
  border: 1px solid #e0e0e0;
  border-radius: 32px;
  color: #333;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  font-size: 15px;
  font-weight: 500;
  padding: 8px 36px;
  cursor: pointer;
  transition: background 0.2s;
}

.see-more-btn:hover {
  background-color: #eeeeee;
}
</style>