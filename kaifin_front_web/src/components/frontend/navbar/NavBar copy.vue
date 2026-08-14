<template>
  <div class="main-navbar">
    <!-- Navbar logo  -->
    <nav class="navbar">
      <div class="logo">
        <img src="../../../assets/logos/kaifin_l2.png" alt="Kaifin" />
      </div>

      <!-- Loading & Nav Links -->
      <ul class="nav-links" v-if="!isLoading && !loadError">
        <li v-for="item in menuItems" :key="item.path">
          
          <!-- ករណី ១: បើជា Menu Music ប្រើ <a> ធម្មតា និងហាមឃាត់ដាច់ខាតមិនឱ្យដូរ Route -->
          <a
            v-if="isMusicItem(item)"
            href="javascript:void(0)"
            :class="{ 'is-active': showMusicModal }"
            @click.prevent.stop="showMusicModal = true"
          >
            {{ item.label }}
          </a>

          <!-- ករណី ២: Menu ផ្សេងទៀត ឱ្យដើរតួជា Router Link ធម្មតា -->
          <router-link
            v-else
            :to="item.path"
            :class="{ 'is-active': route.path === item.path }"
          >
            {{ item.label }}
          </router-link>

        </li>
      </ul>

      <!-- Loading  -->
      <span class="nav-status" v-else-if="isLoading">Waiting...</span>
      <span class="nav-status is-error" v-else>{{ loadError }}</span>

      <ul class="icon-list">
        <li v-for="item in navIcons" :key="item.name">
          <button class="icon-btn" :class="{ 'is-fill': item.filled }" :title="item.name" :aria-label="item.name" @click="onIconClick(item)">
            <svg viewBox="0 0 24 24" v-html="item.svg"></svg>
            <span class="icon-badge" v-if="item.badge">{{ item.badge }}</span>
          </button>
        </li>
      </ul>

      <!-- Menu all  -->
      <div class="right-group">
        <button class="apps-btn" title="All app" aria-label="All app">
          <svg viewBox="0 0 24 24" fill="currentColor">
            <circle cx="5" cy="5" r="1.6"/><circle cx="12" cy="5" r="1.6"/><circle cx="19" cy="5" r="1.6"/>
            <circle cx="5" cy="12" r="1.6"/><circle cx="12" cy="12" r="1.6"/><circle cx="19" cy="12" r="1.6"/>
            <circle cx="5" cy="19" r="1.6"/><circle cx="12" cy="19" r="1.6"/><circle cx="19" cy="19" r="1.6"/>
          </svg>
        </button>
        <!-- Drop down profile  -->
        <AccountMenu :user="accountMenuUser" @settings="goToSettings" @help="goToProfile" @logout="logout" @switch-profile="goToLogin"/>
      </div>
    </nav>

    <!-- Music popup -->
    <Teleport to="body">
      <div v-if="showMusicModal" class="music-modal-wrapper">
        
        <!-- Header NavBar -->
        <header class="modal-navbar-header">
          <NavBar />
        </header>

        <!-- Overlay Body -->
        <div class="music-modal-body" @click.self="showMusicModal = false">
          
          <div class="modal-card-wrapper">
            
            <button class="music-modal-close" @click="showMusicModal = false" title="Close Modal">
              ✕
            </button>

            <!-- App Music Box -->
            <div class="music-modal-content">
              <Music />
            </div>

          </div>

        </div>

      </div>
    </Teleport>

    <!-- Notification popup -->
    <Teleport to="body">
      <div v-if="showNotificationModal" class="notification-modal-wrapper">
        <div class="notification-overlay" @click="showNotificationModal = false"></div>
        
        <div class="notification-modal-content-box">
          <!-- App Notification Box -->
          <Notification @back="showNotificationModal = false" @close="showNotificationModal = false" />
        </div>
      </div>
    </Teleport>

    <!-- Chat popup -->
    <Teleport to="body">
      <div v-if="showChatModal" class="notification-modal-wrapper chat-wrapper-popup">
        <div class="notification-overlay" @click="showChatModal = false"></div>
        
        <div class="notification-modal-content-box">
          <!-- ហៅ Chat Component មកបង្ហាញ និងផ្ដល់ព្រឹត្តិការណ៍ close -->
          <Chart @close="showChatModal = false" />
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import axios from 'axios'
import AccountMenu from '../account_menu/AccountMenu.vue'
import Music from '../music/Music.vue'
import Notification from '../notification.vue/Notification.vue'


import Chart from '../chart/chart.vue'


const API_BASE = 'http://localhost:7070/api/v1/front'

const route = useRoute()
const router = useRouter()

const emit = defineEmits(['search', 'icon-click'])

const menuItems = ref([])
const isLoading = ref(true)
const loadError = ref('')

const profile = ref(null)
const profileError = ref('')

const searchText = ref('')

// Music popup state
const showMusicModal = ref(false)

// Notification popup state
const showNotificationModal = ref(false)

// Chat popup state
const showChatModal = ref(false)

const navIcons = ref([
  { name: 'Message', badge: 3, svg: '<rect x="3" y="5" width="18" height="14" rx="2"/><path d="M3 6l9 7 9-7"/>' },
  { name: 'Chat', badge: 0, svg: '<path d="M21 12a8 8 0 1 1-3.2-6.4L21 4l-1 4.6A7.96 7.96 0 0 1 21 12Z"/>' },
  { name: 'Notification', badge: 12, svg: '<path d="M6 10a6 6 0 1 1 12 0c0 5 2 6 2 6H4s2-1 2-6Z"/><path d="M10 20a2 2 0 0 0 4 0"/>' },
])

const fullName = computed(() => {
  const p = profile.value
  if (!p) return ''
  return [p.first_name, p.last_name].filter(Boolean).join(' ')
})

const avatarUrl = computed(() => {
  const raw = profile.value?.profile_images
  if (!raw) return ''

  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `http://localhost:7070/uploads/${raw}`
})

const accountMenuUser = computed(() => ({
  username: fullName.value || 'Username',
  avatarUrl: avatarUrl.value,
}))

function goToLogin() {
  router.push('/login')
}

function authHeaders() {
  const token = localStorage.getItem('token')
  return { Authorization: `Bearer ${token}` }
}

function slugify(title) {
  return (
    '/' +
    title
      .toLowerCase()
      .trim()
      .replace(/[^a-z0-9]+/g, '-')
      .replace(/^-+|-+$/g, '')
  )
}

function normalizeMenuItem(item) {
  return {
    label: item.title,
    path: slugify(item.title)
  }
}

function isMusicItem(item) {
  if (!item) return false
  const path = item.path ? item.path.toLowerCase() : ''
  const label = item.label ? item.label.toLowerCase() : ''
  
  return path.includes('music') || label.includes('music')
}

async function fetchMenus() {
  isLoading.value = true
  loadError.value = ''

  try {
    const response = await axios.get(`${API_BASE}/menus/show`, {
      headers: authHeaders()
    })

    const rawItems = response.data?.data?.menus ?? []

    menuItems.value = rawItems
      .filter(item => item.is_active)
      .sort((a, b) => a.sort_order - b.sort_order)
      .map(normalizeMenuItem)
  } catch (error) {
    console.error('Failed to load menu:', error)
    loadError.value = 'Could not load menu'
  } finally {
    isLoading.value = false
  }
}

async function fetchProfile() {
  const token = localStorage.getItem('token')
  if (!token) return

  try {
    const response = await axios.get(`${API_BASE}/profile/show`, {
      headers: authHeaders()
    })

    profile.value = response.data.data
  } catch (error) {
    profileError.value = 'Could not load profile'
    console.error(error.response?.data)
  }
}

function onIconClick(item) {
  if (item.name === 'Notification') {
    showNotificationModal.value = true
  } else if (item.name === 'Chat') {
    showChatModal.value = true
  } else {
    emit('icon-click', item.name)
  }
}

function goToProfile() {
  router.push('/profile')
}

function goToSettings() {
  router.push('/settings')
}

function logout() {
  localStorage.removeItem('token')
  window.location.href = '/login'
}

watch([showMusicModal, showNotificationModal, showChatModal], ([musicVal, notifVal, chatVal]) => {
  if (musicVal || notifVal || chatVal) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
})

onMounted(() => {
  fetchMenus()
  fetchProfile()
})
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.main-navbar{
  width: 100%;
  background: #fff;
  border-bottom: 1px solid #E7E7E7;
  position: sticky;
  top: 0;
  z-index: 998;
}

.navbar {
  width: 85%;
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 0 24px;
  height: 64px;
  font-family: 'Inter', sans-serif;
  margin: auto;
}

.logo {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-right: 24px;
  cursor: pointer;
  flex-shrink: 0;
}

.logo img {
  height: 32px;
  width: auto;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 4px;
  list-style: none;
  margin: 0;
  padding: 0;
  margin-right: 24px;
}

.nav-links a {
  display: block;
  padding: 8px 14px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  color: #4A4A4E;
  text-decoration: none;
  white-space: nowrap;
  cursor: pointer;
}

.nav-links a:hover {
  background: #F2F2F3;
}

.nav-links a.is-active {
  background: #bbdefb;
  color: #4A4A4E;
}

.nav-status {
  font-size: 13px;
  color: #9A9A9E;
  margin-right: 24px;
  white-space: nowrap;
}

.nav-status.is-error {
  color: #C6402E;
}

.icon-list {
  display: flex;
  justify-content: flex-end;
  align-items: flex-end;
  gap: 2px;
  flex: 1;
  min-width: 0;
  overflow-x: auto;
  list-style: none;
  margin: 0;
  padding: 0;
}

.icon-btn {
  width: 44px;
  height: 44px;
  flex-shrink: 0;
  border-radius: 10px;
  border: none;
  background: transparent;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  position: relative;
  color: #8A8A8E;
  transition: background .15s, color .15s;
}

.icon-btn:hover {
  background: #F2F2F3;
  color: #4A4A4E;
}

.icon-btn svg {
  width: 22px;
  height: 22px;
  stroke: currentColor;
  fill: none;
  stroke-width: 1.7;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.icon-btn.is-fill svg {
  fill: currentColor;
  stroke: none;
}

.icon-badge {
  position: absolute;
  top: 5px;
  right: 5px;
  min-width: 15px;
  height: 15px;
  padding: 0 3px;
  background: #00D262;
  color: #fff;
  font-size: 10px;
  font-weight: 700;
  border-radius: 999px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.right-group {
  display: flex;
  align-items: center;
  gap: 10px;
  flex-shrink: 0;
  margin-left: 16px;
}

.apps-btn {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  border: none;
  background: transparent;
  color: #8A8A8E;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.apps-btn:hover {
  background: #F2F2F3;
  color: #4A4A4E;
}

.apps-btn svg {
  width: 19px;
  height: 19px;
}

.music-modal-wrapper {
  position: fixed;
  inset: 0;
  z-index: 9999;
  display: flex;
  flex-direction: column;
}

.modal-navbar-header {
  position: relative;
  width: 100%;
  background: #fff;
  z-index: 10000;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}

.music-modal-body {
  flex: 1;
  background: rgba(0, 0, 0, 0.433);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  overflow-y: auto;
}

.modal-card-wrapper {
  position: relative;
  width: 100%;
  max-width: 1280px;
}

.music-modal-content {
  background: transparent;
  border-radius: 16px;
  max-height: 90vh;
  overflow-y: auto;
}

.music-modal-close {
  position: absolute;
  top: 10px;
  right: -40px;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: 2px solid #ffffff;
  background: #1976d2;
  color: #ffffff;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.25);
  transition: all 0.2s ease;
}

.music-modal-close:hover {
  background: #ef4444;
  border-color: #ef4444;
  transform: scale(1.1);
}

.notification-modal-wrapper {
  position: fixed;
  top: 64px;
  right: 232px;
  z-index: 9999;
  background: transparent;
  display: flex;
  justify-content: flex-end;
  align-items: flex-start;
}

.chat-wrapper-popup {
  right: 232px;
}

.notification-overlay {
  position: fixed;
  top: 64px;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.4); /* បន្ថែមពណ៌ផ្ទៃខាងក្រោយស្រអាប់ឱ្យ Overlay */
  z-index: -1;
  animation: fadeIn 0.4s ease-out;
}

.notification-modal-content-box {
  position: relative;
  z-index: 9999;
  animation: slideDownFadeIn 0.45s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  /* overflow-y: auto; */
  height: calc(100vh - 90px); 
    display: flex;   
  overflow: hidden; 
  max-height: none;


  max-height: calc(100vh - 90px);
  padding-top: 12px;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes slideDownFadeIn {
  from {
    opacity: 0;
    transform: translateY(-25px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>