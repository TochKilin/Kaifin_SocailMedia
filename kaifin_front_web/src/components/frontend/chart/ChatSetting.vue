<template>
  <div class="chat-setting-container">
    <!-- Left Sidebar Menu -->
    <div class="setting-sidebar">
      <h3>Settings</h3>
      <ul class="sidebar-menu">
        <li class="menu-item" :class="{ active: currentTab === 'chat-setting' }" @click="currentTab = 'chat-setting'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
          <span>Chat Settings</span>
        </li>
        <li class="menu-item" :class="{ active: currentTab === 'restricted-account' }" @click="currentTab = 'restricted-account'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line></svg>
          <span>Restricted Account</span>
        </li>
        <li class="menu-item" :class="{ active: currentTab === 'block-setting' }" @click="currentTab = 'block-setting'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 2l-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778 5.5 5.5 0 0 1 7.778-7.778zm0 0L15.5 7.5m0 0l3 3L22 7l-3-3m-3.5 3.5L19 4"></path></svg>
          <span>Block Setting</span>
        </li>
      </ul>
    </div>

    <div class="chat-setting-panel">
      <!-- Header -->
      <div class="setting-header">
        <h2>{{ headerTitle }}</h2>
        <button class="close-btn" title="Close" @click="$emit('close')">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>

      <div class="setting-content" v-if="currentTab === 'chat-setting'">
        <!-- 1. Incoming call sound -->
        <div class="setting-item">
          <div class="setting-left">
            <div class="setting-icon incoming-call-icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path></svg>
            </div>
            <div class="setting-text">
              <span class="setting-label">Incoming call sound</span>
              <span class="setting-desc">Play sound when receiving a voice or video call</span>
            </div>
          </div>
          <div class="toggle-switch" :class="{ active: settings.incomingCallSound }" @click="settings.incomingCallSound = !settings.incomingCallSound">
            <div class="toggle-thumb"></div>
          </div>
        </div>

        <!-- Message sound -->
        <div class="setting-item">
          <div class="setting-left">
            <div class="setting-icon message-sound-icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon><path d="M19.07 4.93a10 10 0 0 1 0 14.14"></path><path d="M15.54 8.46a5 5 0 0 1 0 7.07"></path></svg>
            </div>
            <div class="setting-text">
              <div class="label-row">
                <span class="setting-label">Message sound</span>
                <span class="status-badge" :class="{ 'inactive': !settings.messageSound }">
                  {{ settings.messageSound ? 'ON' : 'OFF' }}
                </span>
              </div>
              <span class="setting-desc">Play notification sound for new incoming messages</span>
            </div>
          </div>
          <div class="toggle-switch" :class="{ active: settings.messageSound }" @click="settings.messageSound = !settings.messageSound">
            <div class="toggle-thumb"></div>
          </div>
        </div>

        <!-- Pop-up new messages -->
        <div class="setting-item">
          <div class="setting-left">
            <div class="setting-icon popup-messages-icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><line x1="3" y1="9" x2="21" y2="9"></line></svg>
            </div>
            <div class="setting-text">
              <span class="setting-label">Pop-up new messages</span>
              <span class="setting-desc">Show a preview banner when new messages arrive</span>
            </div>
          </div>
          <div class="toggle-switch" :class="{ active: settings.popupMessages }" @click="settings.popupMessages = !settings.popupMessages">
            <div class="toggle-thumb"></div>
          </div>
        </div>

        <!-- Active status -->
        <div class="setting-item">
          <div class="setting-left">
            <div class="setting-icon active-status-icon status-icon-wrap">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
              <span class="status-dot"></span>
            </div>
            <div class="setting-text">
              <div class="label-row">
                <span class="setting-label">Active status</span>
                <span class="status-badge" :class="{ 'inactive': !settings.activeStatus }">
                  {{ settings.activeStatus ? 'ON' : 'OFF' }}
                </span>
              </div>
              <span class="setting-desc">Show when you're active or recently online</span>
            </div>
          </div>
          <div class="toggle-switch" :class="{ active: settings.activeStatus }" @click="settings.activeStatus = !settings.activeStatus">
            <div class="toggle-thumb"></div>
          </div>
        </div>

        <!-- Restricted account link -->
        <div class="setting-item clickable" @click="currentTab = 'restricted-account'">
          <div class="setting-left">
            <div class="setting-icon restricted-account-icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line></svg>
            </div>
            <div class="setting-text">
              <span class="setting-label">Restricted account</span>
              <span class="setting-desc">Manage accounts you have restricted from messaging</span>
            </div>
          </div>
        </div>

        <!-- Block setting link -->
        <div class="setting-item clickable" @click="currentTab = 'block-setting'">
          <div class="setting-left">
            <div class="setting-icon block-setting-icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 2l-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778 5.5 5.5 0 0 1 7.778-7.778zm0 0L15.5 7.5m0 0l3 3L22 7l-3-3m-3.5 3.5L19 4"></path></svg>
            </div>
            <div class="setting-text">
              <span class="setting-label">Block setting</span>
              <span class="setting-desc">View and manage your blocked users list</span>
            </div>
          </div>
        </div>

        <!-- Write to the Customer Support Team -->
        <div class="setting-item clickable support-item">
          <div class="setting-left">
            <div class="setting-icon support-icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"></path><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
            </div>
            <div class="setting-text">
              <span class="setting-label">Write to the Customer Support Team</span>
            </div>
          </div>
        </div>
      </div>

      <!-- TabRestricted Account Content -->
      <div class="setting-content" v-if="currentTab === 'restricted-account'">
        <div class="setting-item block-user-card" v-for="user in mockRestrictedUsers" :key="user.id">
          <div class="setting-left">
            <!-- Profile Avatar with Restricted Badge Overlaid -->
            <div class="user-avatar-wrap avatar-with-badge">
              <img :src="user.avatar" alt="Avatar" class="item-avatar" />
              <div class="key-badge-overlay restricted-badge">
                <svg width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="12" cy="12" r="10"></circle><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line></svg>
              </div>
            </div>
            <div class="setting-text">
              <span class="setting-label">{{ user.username }}</span>
              <span class="setting-desc">Restricted account</span>
            </div>
          </div>

          <!-- Right side: Date + More Options Menu -->
          <div class="block-right-actions">
            <div class="block-date-badge">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <polyline points="12 6 12 12 16 14"></polyline>
              </svg>
              <span>{{ user.date }}</span>
            </div>

            <!-- More Options Button & Dropdown Menu -->
            <div class="more-options-container" v-click-outside="() => user.showMenu = false">
              <button class="more-options-btn" @click="user.showMenu = !user.showMenu" title="More options">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="1"></circle>
                  <circle cx="12" cy="5" r="1"></circle>
                  <circle cx="12" cy="19" r="1"></circle>
                </svg>
              </button>

              <!-- Dropdown Menu for Unrestrict & Cancel -->
              <div class="dropdown-menu" v-if="user.showMenu">
                <button class="dropdown-item unblock-option" @click="unrestrictUser(user.id)">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line></svg>
                  <span>Unrestrict</span>
                </button>
                <button class="dropdown-item cancel-option" @click="user.showMenu = false">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
                  <span>Cancel</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- TabBlock Setting Content -->
      <div class="setting-content" v-if="currentTab === 'block-setting'">
        <div class="setting-item block-user-card" v-for="user in mockBlockedUsers" :key="user.id">
          <div class="setting-left">
            <!-- Profile Avatar with Red Key Badge Overlaid -->
            <div class="user-avatar-wrap avatar-with-badge">
              <img :src="user.avatar" alt="Avatar" class="item-avatar" />
              <div class="key-badge-overlay red-badge">
                <svg width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M21 2l-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778 5.5 5.5 0 0 1 7.778-7.778zm0 0L15.5 7.5m0 0l3 3L22 7l-3-3m-3.5 3.5L19 4"></path></svg>
              </div>
            </div>
            <div class="setting-text">
              <span class="setting-label">{{ user.username }}</span>
              <span class="setting-desc">Blocked account</span>
            </div>
          </div>

          <!-- Right side: Date More Options Menu -->
          <div class="block-right-actions">
            <div class="block-date-badge">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <polyline points="12 6 12 12 16 14"></polyline>
              </svg>
              <span>{{ user.date }}</span>
            </div>

            <!-- More Options Button & Dropdown Menu -->
            <div class="more-options-container" v-click-outside="() => user.showMenu = false">
              <button class="more-options-btn" @click="user.showMenu = !user.showMenu" title="More options">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="1"></circle>
                  <circle cx="12" cy="5" r="1"></circle>
                  <circle cx="12" cy="19" r="1"></circle>
                </svg>
              </button>

              <!-- Dropdown Menu for Unblock & Cancel -->
              <div class="dropdown-menu" v-if="user.showMenu">
                <button class="dropdown-item unblock-option" @click="unblockUser(user.id)">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="8.5" cy="7" r="4"></circle><line x1="18" y1="8" x2="23" y2="13"></line><line x1="23" y1="8" x2="18" y2="13"></line></svg>
                  <span>Unblock</span>
                </button>
                <button class="dropdown-item cancel-option" @click="user.showMenu = false">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
                  <span>Cancel</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, reactive, onMounted, onBeforeUnmount } from 'vue'

defineEmits(['close'])

const currentTab = ref('chat-setting')

const headerTitle = computed(() => {
  if (currentTab.value === 'chat-setting') return 'Chat Setting'
  if (currentTab.value === 'restricted-account') return 'Restricted Account'
  if (currentTab.value === 'block-setting') return 'Block Account'
  return 'Settings'
})

const settings = reactive({
  incomingCallSound: true,
  messageSound: true,
  popupMessages: true,
  activeStatus: true
})

const mockRestrictedUsers = ref([
  { id: 1, username: 'Pon Thona', date: '01-06-2026', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR3Im7NTh6-v90KiwiB9pbqSsuxOKccRux7KD8-PYKSYA&s=10', showMenu: false },
  { id: 2, username: 'Dara Dev', date: '03-06-2026', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSIP10Y3JbQk30y2eX5jCIn7RczECXykSP8g00fcvqw6w&s=10', showMenu: false },
  { id: 2, username: 'Dara Dev', date: '03-06-2026', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYF7GeeRr2DOGhxoNs-UDuRQDU8qNoblTtwx_6vJNBTw&s=10', showMenu: false },
  { id: 2, username: 'Dara Dev', date: '03-06-2026', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ1MKhs-OPGolBgBQi8AYfXRMTveUQcXhqXwU3oQ2hp6Q&s=10', showMenu: false }
])

const mockBlockedUsers = ref([
  { id: 1, username: 'Moni Reaksa', date: '02-06-2026', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSwfezfM87o9aGPgx2aZjmt4r81csnYdP1XVEduGMFwkQ&s=10', showMenu: false },
  { id: 2, username: 'Hak Vanchai', date: '04-06-2026', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSGdGoxWOxVMrjZU32idg6AzN8yxuByaMyVvP4zkkb1aw&s=10', showMenu: false },
  { id: 3, username: 'Thol Vathana', date: '05-06-2026', avatar: 'https://encrypted-t9.gstatic.com/images?q=tbn:ANd9GcSHDSto4_Kr4udncJwmiV87EA5FspVbDDtrJx5mcUVIzg&s=10', showMenu: false }
])

const unrestrictUser = (id) => {
  mockRestrictedUsers.value = mockRestrictedUsers.value.filter(u => u.id !== id)
}

const unblockUser = (id) => {
  mockBlockedUsers.value = mockBlockedUsers.value.filter(u => u.id !== id)
}

// Custom directive to handle click outside dropdown menu
const vClickOutside = {
  mounted(el, binding) {
    el.clickOutsideEvent = function(event) {
      if (!(el === event.target || el.contains(event.target))) {
        binding.value(event)
      }
    }
    document.body.addEventListener('click', el.clickOutsideEvent)
  },
  unmounted(el) {
    document.body.removeEventListener('click', el.clickOutsideEvent)
  }
}
</script>

<style scoped>
.chat-setting-container {
  display: flex;
  width: 100%;
  height: 100%;
  background-color: #ffffff;
  border-radius: 16px;
  overflow: hidden;
  box-sizing: border-box;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  color: #2b2b2b;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.08);
}

.setting-sidebar {
  width: 260px;
  background-color: #ffffff;
  border-right: 1px solid #e5e7eb;
  padding: 24px 16px;
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  box-sizing: border-box;
}

.setting-sidebar h3 {
  margin: 0 0 12px 12px;
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.sidebar-menu {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 14px;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 500;
  color: #4b5563;
  cursor: pointer;
  transition: color 0.2s;
}

.menu-item:hover {
  background-color: transparent;
  color: #0284c7;
}

.menu-item.active {
  background-color: transparent;
  color: #0284c7;
  font-weight: 600;
}

.chat-setting-panel {
  flex: 1;
  height: 100%;
  background-color: transparent;
  padding: 0;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
}

.setting-header {
  position: relative;
  text-align: center;
  background-color: #ffffff;
  border: none;
  box-shadow: none;
  border-radius: 0;
  padding: 16px 20px;
  margin-bottom: 24px;
  flex-shrink: 0;
  width: 100%;
  box-sizing: border-box;
}

.setting-header h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
}

.close-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  right: 20px;
  background-color: #f3f4f6;
  border: 1px solid #e5e7eb;
  color: #4b5563;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.2s, color 0.2s;
}

.close-btn:hover {
  background-color: #e5e7eb;
  color: #1f2937;
}

.setting-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 100%;
  padding: 0 24px 24px 24px;
  box-sizing: border-box;
}

.setting-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 10px;
  padding: 12px 18px;
  width: 100%;
  box-sizing: border-box;
  transition: background-color 0.2s, border-color 0.2s;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.02);
}

.setting-left {
  display: flex;
  align-items: center;
  gap: 14px;
}

.setting-icon {
  width: 38px;
  height: 38px;
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280; 
  flex-shrink: 0;
  position: relative;
}

.user-avatar-wrap {
  width: 38px;
  height: 38px;
  flex-shrink: 0;
  position: relative;
}

.avatar-with-badge {
  position: relative;
}

.key-badge-overlay {
  position: absolute;
  bottom: -2px;
  right: -2px;
  width: 18px;
  height: 18px;
  border: 2px solid #ffffff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
}

.key-badge-overlay.red-badge {
  background-color: #FF5555;
  border-color: #ffffff;
}

.key-badge-overlay.restricted-badge {
  background-color: #9B9C9D;
  border-color: #ffffff;
}

.item-avatar {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.incoming-call-icon {
  background-color: #20BA55;
  border-color: #20BA55;
  color: #ffffff;
}

.message-sound-icon {
  background-color: #9B9C9D;
  border-color: #9B9C9D;
  color: #ffffff;
}

.popup-messages-icon {
  background-color: #FF5555;
  border-color: #FF5555;
  color: #ffffff;
}

.active-status-icon {
  background-color: #3B82F6;
  border-color: #3B82F6;
  color: #ffffff;
}

.restricted-account-icon {
  background-color: #9B9C9D;
  border-color: #9B9C9D;
  color: #ffffff;
}

.block-setting-icon {
  background-color: #9B9C9D;
  border-color: #9B9C9D;
  color: #ffffff;
}

.support-icon {
  background-color: #9B9C9D;
  border-color: #9B9C9D;
  color: #ffffff;
}

.status-icon-wrap {
  position: relative;
}

.status-dot {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 8px;
  height: 8px;
  background-color: #22c55e;
  border-radius: 50%;
  border: 2px solid #ffffff;
}

.setting-text {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.setting-label {
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
}

.setting-desc {
  font-size: 12px;
  color: #6b7280;
}

.label-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.status-badge {
  background-color: #20BA55;
  color: #ffffff;
  border: 1px solid #20BA55;
  padding: 1px 6px;
  font-size: 10px;
  font-weight: bold;
  border-radius: 4px;
  letter-spacing: 0.5px;
}

.status-badge.inactive {
  background-color: #9B9C9D;
  border-color: #9B9C9D;
  opacity: 0.6;
}

.block-right-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.block-date-badge {
  display: flex;
  align-items: center;
  gap: 6px;
  background: none;
  border: none;
  padding: 0;
  font-size: 13px;
  font-weight: 500;
  color: #4b5563;
}

.more-options-container {
  position: relative;
}

.more-options-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #6b7280;
  padding: 4px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s, color 0.2s;
}

.more-options-btn:hover {
  background-color: #f3f4f6;
  color: #1f2937;
}

.dropdown-menu {
  position: absolute;
  right: 0;
  top: 100%;
  margin-top: 4px;
  width: 140px;
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  z-index: 10;
  display: flex;
  flex-direction: column;
  padding: 4px;
  box-sizing: border-box;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
  padding: 8px 10px;
  background: none;
  border: none;
  border-radius: 6px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  text-align: left;
  transition: background-color 0.2s;
}

.unblock-option {
  color: #4b5563;
}

.unblock-option:hover {
  background-color: #f3f4f6;
  color: #1f2937;
}

.cancel-option {
  color: #4b5563;
}

.cancel-option:hover {
  background-color: #f3f4f6;
}

.toggle-switch {
  width: 44px;
  height: 24px;
  background-color: #e5e7eb;
  border: 1px solid #d1d5db;
  border-radius: 12px;
  position: relative;
  cursor: pointer;
  transition: background-color 0.3s, border-color 0.3s;
  flex-shrink: 0;
}

.toggle-thumb {
  width: 18px;
  height: 18px;
  background-color: #ffffff;
  border-radius: 50%;
  position: absolute;
  top: 2px;
  left: 2px;
  box-shadow: 0 1px 2px rgba(0,0,0,0.2);
  transition: transform 0.3s;
}

.toggle-switch.active {
  background-color: #3b82f6;
  border-color: #2563eb;
}

.toggle-switch.active .toggle-thumb {
  transform: translateX(20px);
}
</style>