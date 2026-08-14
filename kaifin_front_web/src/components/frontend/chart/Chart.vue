<template>
  <div class="chats-modal-wrapper">
    <div class="chats-container">

      <!-- បង្ហាញផ្ទាំង Chat Setting ពេល showSetting ເປັນ true -->
      <ChatSetting v-if="showSetting" @close="showSetting = false" />

      <!-- ករណីបង្ហាញផ្ទាំង Chats ធម្មតា (ពេល showSetting ເປັນ false) -->
      <template v-else>
        <!-- ករណីបង្ហាញផ្ទាំង Chats ធម្មតា -->
        <template v-if="currentView === 'chats'">
          <!-- Header Section -->
          <div class="chats-header">
            <div class="chats-title-box">
              <h1 class="chats-title">Chats</h1>
            </div>
            <div class="header-controls">
              <div class="window-actions">
                <button class="action-icon-btn" @click="handleAction">
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

            <!-- Facebook Style Tabs Navigation (All, Unread, Groups) -->
            <div class="messenger-tabs-nav">
              <button
                class="messenger-tab-item"
                :class="{ active: activeTab === 'all' }"
                @click="activeTab = 'all'"
              >
                <svg class="tab-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/></svg>
                <span>All</span>
              </button>

              <button
                class="messenger-tab-item"
                :class="{ active: activeTab === 'unread' }"
                @click="activeTab = 'unread'"
              >
                <svg class="tab-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/><polyline points="22,6 12,13 2,6"/></svg>
                <span>Unread</span>
                <span v-if="unreadCount > 0" class="tab-badge">{{ unreadCount }}</span>
              </button>

              <button
                class="messenger-tab-item"
                :class="{ active: activeTab === 'groups' }"
                @click="activeTab = 'groups'"
              >
                <svg class="tab-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
                <span>Groups</span>
              </button>
            </div>

            <!-- New Message Request Banner (ចុចទីនេះដើម្បីប្ដូរទៅ Request View) -->
            <div class="request-banner" @click="currentView = 'request'">
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

            <!-- Chat Items List with Mock Data & Avatars (Expanded to bottom) -->
            <div class="chat-list">
              <div
                v-for="(chat, index) in filteredChats"
                :key="index"
                class="chat-item"
                @click="openChat(chat)"
              >
                <!-- Avatar -->
                <div class="avatar-wrapper">
                  <div class="avatar" :class="{ 'group-avatar': chat.isGroup }">
                    <img v-if="chat.avatar" :src="chat.avatar" :alt="chat.name" class="avatar-img" />
                    <svg v-else-if="!chat.isGroup" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                    <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
                  </div>
                  <span v-if="chat.online" class="online-dot"></span>
                  <span v-if="chat.badge && chat.unread" class="unread-badge">{{ chat.badge }}</span>
                </div>

                <!-- Chat Info -->
                <div class="chat-info">
                  <div class="chat-info-header">
                    <span class="chat-name" :class="{ 'unread-text': chat.unread }">{{ chat.name }}</span>
                    <div class="more-options-dots" @click.stop="chat.showMenu = !chat.showMenu">
                      <span class="mini-dot"></span>
                      <span class="mini-dot"></span>
                      <span class="mini-dot"></span>
                    </div>
                    <MoreOption v-if="chat.showMenu" @close="chat.showMenu = false" @action="handleChatOption(chat, $event)" class="more-op" />
                  </div>
                  <div class="chat-info-bottom">
                    <span class="chat-message" :class="{ 'voice-msg': chat.isVoice, 'unread-text': chat.unread }">{{ chat.message }}</span>
                    <span class="chat-time">{{ chat.time }}</span>
                  </div>
                </div>
              </div>

              <!-- Empty State -->
              <div v-if="filteredChats.length === 0" class="empty-state">
                <p>No chats found</p>
              </div>
            </div>
          </div>
        </template>

        <!-- ករណីហៅ Component RequestView.vue មកបង្ហាញពេលចុចលើ Request Banner -->
        <template v-else-if="currentView === 'request'">
          <RequestView @back="currentView = 'chats'" />
        </template>

        <!-- Chat Detail -->
        <template v-else-if="currentView === 'detail'">
          <ChartDetail
            v-if="selectedChat"
            :chat="selectedChat"
            @back="currentView = 'chats'"
          />
        </template>
      </template>

    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import RequestView from './RequestView.vue'
import MoreOption from '../chart/MoreOption.vue'
import ChartDetail from './ChartDetail.vue'
import ChatSetting from './ChatSetting.vue' // <--- បន្ថែមការ Import ឯកសារ ChatSetting របស់អ្នក

defineEmits(['close'])

const currentView = ref('chats')
const showSetting = ref(false) // <--- Variable សម្រាប់ប្ដូរទៅកាន់ផ្ទាំង Setting

const searchQuery = ref('')
const activeTab = ref('all') // 'all' | 'unread' | 'groups'

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
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRlzLSIDO3ceFnP3Qyu0o6sNLmq7_A6l7JLSdJcS9pc7g&s=10',
    showMenu: false,
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
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-971-20200304-4-ro0jp2.jpg',
    showMenu: false,
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
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-2201-20200924-4-92w8ra.jpg',
    showMenu: false,
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
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10',
    showMenu: false,
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
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-3856-20210419-4-ql7c88.jpg',
    showMenu: false,
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
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10',
    showMenu: false,
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
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10',
    showMenu: false,
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
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10',
    showMenu: false,
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
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10',
    showMenu: false,
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
    avatar: 'https://cdn.hypetrace.com/file/ht/avig/instagram-influencer-61355-20200302-4-1e8w3gk.jpg',
    showMenu: false,
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
    avatar: 'https://cdn.hypetrace.com/file/ht/avtt/tiktok-influencer-57374-20200316-4-1h7pxhv.jpg',
    showMenu: false,
  },
])

const unreadCount = computed(() => {
  return chatList.value.filter(chat => chat.unread).length
})

function markAsRead(chat) {
  if (chat.unread) {
    chat.unread = false
    chat.badge = null
  }
}

const filteredChats = computed(() => {
  return chatList.value.filter(chat => {
    if (activeTab.value === 'unread' && !chat.unread) return false
    if (activeTab.value === 'groups' && !chat.isGroup) return false

    const matchesSearch =
      chat.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      chat.message.toLowerCase().includes(searchQuery.value.toLowerCase())

    return matchesSearch
  })
})

function handleChatOption(chat, action) {
  console.log('Chat:', chat.name)
  console.log('Action:', action)

  chat.showMenu = false

  switch (action) {
    case 'mark-unread':
      chat.unread = true
      chat.badge = chat.badge || '1'
      break
    case 'mute':
      console.log('Mute:', chat.name)
      break
    case 'view-profile':
      console.log('View profile:', chat.name)
      break
    case 'audio-call':
      console.log('Audio call:', chat.name)
      break
    case 'video-chat':
      console.log('Video chat:', chat.name)
      break
    case 'block':
      console.log('Block:', chat.name)
      break
    case 'archive':
      console.log('Archive:', chat.name)
      break
    case 'delete':
      console.log('Delete:', chat.name)
      break
    case 'report':
      console.log('Report:', chat.name)
      break
  }
}

function handleOutsideClick(event) {
  const menu = event.target.closest('.more-option-menu')
  const button = event.target.closest('.more-options-dots')

  if (menu || button) {
    return
  }

  chatList.value.forEach(chat => {
    chat.showMenu = false
  })
}

const selectedChat = ref(null)

function openChat(chat) {
  selectedChat.value = chat
  currentView.value = 'detail'
}

function goBackToChats() {
  currentView.value = 'chats'
  selectedChat.value = null
}

// មុខងារពេលចុចលើប៊ូតុង 3-dots ក្នុង Header ដើម្បីបើកផ្ទាំង Setting
function handleAction() {
  showSetting.value = true
}

onMounted(() => {
  document.addEventListener('click', handleOutsideClick)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleOutsideClick)
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');

.chats-modal-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: transparent;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  padding: 0;
  box-sizing: border-box;
}

.chats-container {
  width: 1090px;
  max-width: 100%;
  height: 100%; 
  min-height: 0; 
  background-color: #ffffff;
  border: 1px solid #e0e0e0;
  border-radius: 12px;
  overflow: hidden;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  position: relative;
}

/* Header */
.chats-header {
  padding: 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #ffffff;
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
  padding: 12px 20px 12px 20px;
  flex-grow: 1;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  gap: 8px;
  overflow: hidden;
}

/* Messenger Tabs Styling */
.messenger-tabs-nav {
  display: flex;
  align-items: center;
  gap: 12px;
  border-bottom: 1px solid #e4e6eb;
  padding-bottom: 2px;
}

.messenger-tab-item {
  background: transparent;
  border: none;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  font-size: 15px;
  font-weight: 600;
  color: #65676b;
  padding: 8px 12px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  position: relative;
  border-radius: 8px;
  transition: background-color 0.2s, color 0.2s;
}

.messenger-tab-item:hover {
  background-color: #f2f2f2;
}

.messenger-tab-item.active {
  color: #0084ff;
  background-color: transparent;
}

.messenger-tab-item.active::after {
  content: '';
  position: absolute;
  bottom: -3px;
  left: 0;
  right: 0;
  height: 3px;
  background-color: #0084ff;
  border-radius: 2px 2px 0 0;
}

.tab-icon {
  stroke: currentColor;
}

.tab-badge {
  background-color: #0084ff;
  color: #ffffff;
  font-size: 11px;
  padding: 0 6px;
  border-radius: 10px;
  height: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
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
  background-color: #ffffff;
  border-radius: 12px;
  padding: 10px 14px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
}

.request-banner:hover {
  background-color: #f9f9f9;
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
  gap: 8px;
  overflow-y: auto;
  flex-grow: 1;
  max-height: none;
  padding-right: 4px;
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
  cursor: pointer;
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
  position: relative;
}

.chat-name {
  color: #222;
  font-size: 15px;
  font-weight: 600;
}

.chat-name.unread-text,
.chat-message.unread-text {
  font-weight: 700;
  color: #000;
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

.more-op {
  position: absolute;
  right: -4px;
  top: 20px;
  z-index: 1000;
}
</style>