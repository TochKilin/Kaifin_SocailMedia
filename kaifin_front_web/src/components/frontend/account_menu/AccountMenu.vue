<template>
  <div
    class="account-menu-wrap"
    @mouseenter="openMenu"
    @mouseleave="scheduleCloseMenu"
  >
<button class="account-trigger">
  <div class="trigger-avatar">
    <img v-if="user.avatarUrl" :src="user.avatarUrl" alt="" />
    <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4" fill="none" stroke="currentColor" stroke-width="1.8"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/></svg>
  </div>
  <!-- <span class="trigger-name">{{ user.username }}</span> -->
  <svg class="trigger-chevron" :class="{ 'is-open': isOpen }" viewBox="0 0 24 24">
    <path d="M6 9l6 6 6-6" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
  </svg>
</button>

    <div
      class="account-menu"
      v-if="isOpen"
      @mouseenter="keepMenuOpen"
      @mouseleave="scheduleCloseMenu"
    >
      <div class="account-menu-inner">
        <div class="account-head">
          <div class="head-avatar">
            <img v-if="user.avatarUrl" :src="user.avatarUrl" alt="" />
            <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4" fill="none" stroke="currentColor" stroke-width="1.8"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/></svg>
          </div>
          <span class="head-name">{{ user.username }}</span>
          <button class="logout-link" @click="onLogout">
            <img src="../../../assets/icons/logout.svg" alt="svg">
            Log out
          </button>
        </div>

        <div class="menu-list">
          <button class="menu-row" @click="onAction('settings')">
            <svg class="row-icon" viewBox="0 0 24 24"><circle cx="12" cy="12" r="3" fill="none" stroke="currentColor" stroke-width="1.7"/><path d="M19.4 13.5a1.7 1.7 0 0 0 .3 1.9l.1.1a2 2 0 1 1-2.9 2.9l-.1-.1a1.7 1.7 0 0 0-1.9-.3 1.7 1.7 0 0 0-1 1.6v.2a2 2 0 1 1-4 0v-.1a1.7 1.7 0 0 0-1.1-1.6 1.7 1.7 0 0 0-1.9.3l-.1.1a2 2 0 1 1-2.9-2.9l.1-.1a1.7 1.7 0 0 0 .3-1.9 1.7 1.7 0 0 0-1.6-1h-.2a2 2 0 1 1 0-4h.1a1.7 1.7 0 0 0 1.6-1.1 1.7 1.7 0 0 0-.3-1.9l-.1-.1a2 2 0 1 1 2.9-2.9l.1.1a1.7 1.7 0 0 0 1.9.3h.1a1.7 1.7 0 0 0 1-1.6v-.2a2 2 0 1 1 4 0v.1a1.7 1.7 0 0 0 1 1.6 1.7 1.7 0 0 0 1.9-.3l.1-.1a2 2 0 1 1 2.9 2.9l-.1.1a1.7 1.7 0 0 0-.3 1.9v.1a1.7 1.7 0 0 0 1.6 1h.2a2 2 0 1 1 0 4h-.1a1.7 1.7 0 0 0-1.6 1Z" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round"/></svg>
            <span>Settings</span>
          </button>

          <button class="menu-row" @click="toggleTheme">
            <svg class="row-icon" viewBox="0 0 24 24"><circle cx="12" cy="12" r="4.2" fill="none" stroke="currentColor" stroke-width="1.7"/><path d="M12 2.5v2.4M12 19.1v2.4M4.2 4.2l1.7 1.7M18.1 18.1l1.7 1.7M2.5 12h2.4M19.1 12h2.4M4.2 19.8l1.7-1.7M18.1 5.9l1.7-1.7" stroke="currentColor" stroke-width="1.7" stroke-linecap="round"/></svg>
            <span>Theme: {{ theme === 'light' ? 'Light' : 'Dark' }}</span>
          </button>

          <button class="menu-row" @click="isLangOpen = !isLangOpen">
            <svg class="row-icon" viewBox="0 0 24 24"><circle cx="12" cy="12" r="9" fill="none" stroke="currentColor" stroke-width="1.7"/><path d="M3 12h18M12 3c2.4 2.6 3.6 5.7 3.6 9s-1.2 6.4-3.6 9c-2.4-2.6-3.6-5.7-3.6-9s1.2-6.4 3.6-9Z" fill="none" stroke="currentColor" stroke-width="1.5"/></svg>
            <span>{{ language }}</span>
          </button>
          <div class="lang-options" v-if="isLangOpen">
            <button
              v-for="lang in ['English', 'ខ្មែរ']"
              :key="lang"
              class="lang-option"
              :class="{ active: language === lang }"
              @click="selectLanguage(lang)"
            >
              {{ lang }}
            </button>
          </div>

          <button class="menu-row" @click="onAction('help')">
            <svg class="row-icon" viewBox="0 0 24 24"><circle cx="12" cy="12" r="9.2" fill="none" stroke="currentColor" stroke-width="1.7"/><path d="M9.3 9.4c.3-1.6 1.6-2.5 3-2.5 1.7 0 3 1.1 3 2.7 0 1.3-.8 1.9-1.8 2.6-.8.6-1.2 1.1-1.2 2.1" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round"/><circle cx="12" cy="17.6" r="1.05" fill="currentColor"/></svg>
            <span>Help</span>
          </button>
        </div>

        <hr class="menu-divider" />

        <button class="switch-profile-btn" @click="onAction('switch-profile')">
          Log in to another profile
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  user: {
    type: Object,
    default: () => ({ username: 'Username', avatarUrl: '' }),
  },
})

const emit = defineEmits(['logout', 'settings', 'help', 'switch-profile', 'theme-change', 'language-change'])

const isOpen = ref(false)
const isLangOpen = ref(false)
const theme = ref('light')
const language = ref('English')
let menuCloseTimer = null

function openMenu() {
  clearTimeout(menuCloseTimer)
  isOpen.value = true
}

function scheduleCloseMenu() {
  clearTimeout(menuCloseTimer)
  menuCloseTimer = setTimeout(() => {
    isOpen.value = false
    isLangOpen.value = false
  }, 300)
}

function keepMenuOpen() {
  clearTimeout(menuCloseTimer)
}

function closeMenu() {
  isOpen.value = false
  isLangOpen.value = false
}

function onLogout() {
  closeMenu()
  emit('logout')
}

function onAction(action) {
  closeMenu()
  emit(action)
}

function toggleTheme() {
  theme.value = theme.value === 'light' ? 'dark' : 'light'
  emit('theme-change', theme.value)
}

function selectLanguage(lang) {
  language.value = lang
  isLangOpen.value = false
  emit('language-change', lang)
}
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.account-menu-wrap {
  position: relative;
  display: inline-flex;
  font-family: 'Inter', sans-serif;
}

.account-trigger {
  display: flex;
  align-items: center;
  gap: 10px;
  border: none;
  background: transparent;
  cursor: pointer;
  padding: 6px 10px;
  border-radius: 10px;
}

.account-trigger:hover {
  background: #f2f2f3;
}

.trigger-avatar,
.head-avatar {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: #e7e7e7;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  color: #6b7280;
}

.head-avatar {
  /* width: 20x;
  height: 20px; */
}

.trigger-avatar img,
.head-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.trigger-avatar svg,
.head-avatar svg {
  width: 20px;
  height: 20px;
}

.trigger-name {
  font-weight: 600;
  font-size: 12px;
  color: #2b2b2b;
}

.account-menu {
  position: absolute;
  top: 100%;
  left: -40px;
  padding-top: 8px;
  z-index: 50;
}

.account-menu-inner {
  background: #FFFFFF;
  border-radius: 6px;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.16);
  padding: 6px;
  min-width: 210px;
  animation: menuPopIn 0.15s ease;
  position: relative;
  border: 1px solid #EFE2D3;
  top: 4px;
  left: -150px;
  
}

@keyframes menuPopIn {
  from { opacity: 0; transform: translateY(-6px); }
  to   { opacity: 1; transform: translateY(0); }
}

.account-head {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 18px;
}

.head-name {
  flex: 1;
  font-size: 12px;
  font-weight: 600;
  color: #1a1a1a;
}

.logout-link {
  background-color: transparent;
  border: none;
  color: #6b7280;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  text-align: right;
  line-height: 1.3;
  padding: 0;
  padding: 4px 4px;
}

.logout-link img{
  width: 14px;
  height: 14px;
}

.logout-link:hover {
  color: #1a1a1a;
}

.menu-list {
  display: flex;
  flex-direction: column;
}

.menu-row {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
  text-align: left;
  border: none;
  background: transparent;
  padding: 4px 4px;
  border-radius: 10px;
  font-size: 12px;
  font-weight: 500;
  color: #4a4a4e;
  cursor: pointer;
}

.menu-row:hover {
  background: #f5f5f6;
  color: #1a1a1a;
}

.row-icon {
  width: 16px;
  height: 16px;
  color: #4a4a4e;
  flex-shrink: 0;
}

.menu-row:hover .row-icon {
  color: #1a1a1a;
}

.lang-options {
  display: flex;
  flex-direction: column;
  padding-left: 38px;
  margin: -2px 0 4px;
}

.lang-option {
  border: none;
  background: transparent;
  text-align: left;
  padding: 7px 8px;
  border-radius: 8px;
  font-size: 12px;
  color: #6b7280;
  cursor: pointer;
}

.lang-option:hover {
  background: #f5f5f6;
}

.lang-option.active {
  color: #1976d2;
  font-weight: 600;
}

.menu-divider {
  border: none;
  border-top: 1px solid #EFE2D3;
  margin: 8px 0;
}

.switch-profile-btn {
  width: 100%;
  border: none;
  background: transparent;
  color: #1976d2;
  font-size: 12px;
  font-weight: 500;
  padding: 4px;
  border-radius: 32px;
  cursor: pointer;
}

.switch-profile-btn:hover {
  /* background: #e5e5e6; */
}

.trigger-chevron {
  width: 14px;
  height: 14px;
  color: #6b7280;
  flex-shrink: 0;
  transition: transform 0.2s ease;
}

.trigger-chevron.is-open {
  transform: rotate(180deg);
}

.account-menu {
  position: absolute;
  top: 100%;
  left: 0;
  padding-top: 8px;
  z-index: 50;
}

.account-menu::before {
  content: '';
  position: absolute;
  top: 4px;
  left: 22px;
  width: 18px;
  height: 18px;
  background: #FFFFFF;
  transform: rotate(45deg);
  border-left: 1px solid #EFE2D3;
  border-top: 1px solid #EFE2D3;
  z-index: 1;
}
</style>