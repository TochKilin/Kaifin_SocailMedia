<template>
  <div class="main-navbar">
    <!-- Navbar logo  -->
    <nav class="navbar">
      <div class="logo" style="cursor: pointer;" @click="goToFeed">
        <img src="../../../assets/logos/kaifin_l2.png" alt="Kaifin" />
      </div>

      <!-- Loading & Nav Links -->
      <ul class="nav-links" v-if="!isLoading && !loadError">
        <li
          v-for="item in menuItems"
          :key="item.path"
          class="nav-item-wrapper"
          @mouseenter="isMusicItem(item) && onMusicHover(true)"
          @mouseleave="isMusicItem(item) && onMusicHover(false)"
        >
          <a
            v-if="isMusicItem(item)"
            href="javascript:void(0)"
            :class="{ 'is-active': showMusicModal }"
            @click.prevent.stop="showMusicModal = true"
          >
            {{ item.label }}
          </a>

          <router-link
            v-else
            :to="item.path"
            :class="{ 'is-active': route.path === item.path }"
          >
            {{ item.label }}
          </router-link>

          <!-- Music hover preview card -->
          <Transition name="fade-preview">
            <div
              v-if="isMusicItem(item) && showMusicPreview"
              class="music-preview-card"
              :style="previewSong && previewSong.cover ? { backgroundImage: 'url(' + previewSong.cover + ')' } : {}"
              @click="openMusicFromPreview"
            >
              <div class="preview-overlay"></div>

              <div class="preview-top">
                <button class="preview-icon-btn" title="Add" @click.stop="addPreviewSong">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                  </svg>
                </button>
                <button class="preview-icon-btn" title="Playlist" @click.stop="showMusicModal = true">
                  <svg viewBox="0 0 24 24" fill="currentColor">
                    <path d="M4 6h13v2H4zm0 5h13v2H4zm0 5h9v2H4zM17 9v6.18A2.5 2.5 0 1 0 19 17.5V11h3V9h-5z"></path>
                  </svg>
                </button>
              </div>

              <div class="preview-controls">
                <button class="preview-ctrl-btn" title="Previous" @click.stop="previewPrevious">
                  <svg viewBox="0 0 24 24" fill="currentColor">
                    <polygon points="19 20 9 12 19 4 19 20"></polygon>
                    <rect x="5" y="4" width="2" height="16"></rect>
                  </svg>
                </button>
                <button class="preview-ctrl-btn play" title="Play/Pause" @click.stop="togglePreviewPlay">
                  <svg v-if="!previewIsPlaying" viewBox="0 0 24 24" fill="currentColor">
                    <polygon points="5 3 19 12 5 21 5 3"></polygon>
                  </svg>
                  <svg v-else viewBox="0 0 24 24" fill="currentColor">
                    <rect x="6" y="4" width="4" height="16"></rect>
                    <rect x="14" y="4" width="4" height="16"></rect>
                  </svg>
                </button>
                <button class="preview-ctrl-btn" title="Next" @click.stop="previewNext">
                  <svg viewBox="0 0 24 24" fill="currentColor">
                    <polygon points="5 4 15 12 5 20 5 4"></polygon>
                    <rect x="17" y="4" width="2" height="16"></rect>
                  </svg>
                </button>
              </div>

              <div class="preview-info">
                <template v-if="previewLoading">
                  <span class="preview-title">Loading...</span>
                </template>
                <template v-else-if="previewSong">
                  <span class="preview-title">{{ previewSong.title }}</span>
                  <span class="preview-singer">{{ previewSong.singer }}</span>
                </template>
                <template v-else>
                  <span class="preview-title">No song available</span>
                </template>
              </div>
            </div>
          </Transition>
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
        <header class="modal-navbar-header">
          <NavBar />
        </header>

        <div class="music-modal-body" @click.self="showMusicModal = false">
          <div class="modal-card-wrapper">
            <button class="music-modal-close" @click="showMusicModal = false" title="Close Modal">
              ✕
            </button>

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
          <Notification @back="showNotificationModal = false" @close="showNotificationModal = false" />
        </div>
      </div>
    </Teleport>

    <!-- Chat popup -->
    <Teleport to="body">
      <div v-if="showChatModal" class="notification-modal-wrapper chat-wrapper-popup">
        <div class="notification-overlay" @click="showChatModal = false"></div>
        <div class="notification-modal-content-box">
          <Chart @close="showChatModal = false" />
        </div>
      </div>
    </Teleport>
  </div>
</template>



<script>

</script>

<script setup>
import { ref, computed, onMounted, watch, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'
import AccountMenu from '../account_menu/AccountMenu.vue'
import Music from '../music/Music.vue'
import Notification from '../notification.vue/Notification.vue'
import Chart from '../chart/Chart.vue'
// import Chart from '../chart/chart.vue'

const MUSIC_API_BASE = 'http://localhost:7070/api/v1/front'

const previewSongsG = ref([])
const previewSongG = ref(null)
const previewIndexG = ref(0)
const previewLoadingG = ref(false)
const previewLoadedG = ref(false)
const previewIsPlayingG = ref(false)
let previewAudioG = null

function authHeadersG() {
  const token = localStorage.getItem('token')
  return { Authorization: `Bearer ${token}` }
}

function ensurePreviewAudioG() {
  if (!previewAudioG) {
    previewAudioG = new Audio()
    previewAudioG.addEventListener('ended', previewNextG)
    previewAudioG.addEventListener('play', () => { previewIsPlayingG.value = true })
    previewAudioG.addEventListener('pause', () => { previewIsPlayingG.value = false })
  }
  return previewAudioG
}

async function fetchMusicPreviewG() {
  if (previewLoadedG.value) return
  previewLoadingG.value = true
  try {
    const res = await axios.get(`${MUSIC_API_BASE}/songs/show`, {
      headers: authHeadersG(),
      params: { page: 1, limit: 5, search: '' }
    })
    const songs = res.data?.data?.songs ?? []
    previewSongsG.value = songs.map(s => ({
      id: s.id,
      title: s.title,
      singer: s.singer_name ?? `Artist #${s.artist_id}`,
      cover: s.cover_url || '',
      fileUrl: s.file_url
    }))
    previewSongG.value = previewSongsG.value[0] || null
    previewIndexG.value = 0
    previewLoadedG.value = true
  } catch (error) {
    console.error('Failed to load music preview:', error)
    previewSongG.value = null
  } finally {
    previewLoadingG.value = false
  }
}

function playCurrentPreviewG() {
  if (!previewSongG.value?.fileUrl) return
  const audio = ensurePreviewAudioG()
  if (audio.src !== previewSongG.value.fileUrl) audio.src = previewSongG.value.fileUrl
  audio.play().catch(err => console.error('Preview play failed:', err))
}

function togglePreviewPlayG() {
  if (!previewSongG.value?.fileUrl) return
  const audio = ensurePreviewAudioG()
  if (audio.src !== previewSongG.value.fileUrl) audio.src = previewSongG.value.fileUrl
  previewIsPlayingG.value ? audio.pause() : audio.play().catch(err => console.error('Preview play failed:', err))
}

function previewNextG() {
  if (!previewSongsG.value.length) return
  previewIndexG.value = (previewIndexG.value + 1) % previewSongsG.value.length
  previewSongG.value = previewSongsG.value[previewIndexG.value]
  if (previewIsPlayingG.value) {
    const audio = ensurePreviewAudioG()
    audio.src = previewSongG.value.fileUrl
    audio.play().catch(() => {})
  }
}

function previewPreviousG() {
  if (!previewSongsG.value.length) return
  previewIndexG.value = (previewIndexG.value - 1 + previewSongsG.value.length) % previewSongsG.value.length
  previewSongG.value = previewSongsG.value[previewIndexG.value]
  if (previewIsPlayingG.value) {
    const audio = ensurePreviewAudioG()
    audio.src = previewSongG.value.fileUrl
    audio.play().catch(() => {})
  }
}

async function startPreviewOnHoverG() {
  await fetchMusicPreviewG()
}

function stopPreviewG() {
  previewAudioG?.pause()
}

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

const showMusicPreview = ref(false)
let previewHoverTimer = null

const previewSong = previewSongG
const previewLoading = previewLoadingG
const previewIsPlaying = previewIsPlayingG

function onMusicHover(state) {
  clearTimeout(previewHoverTimer)
  if (state) {
    previewHoverTimer = setTimeout(() => {
      showMusicPreview.value = true
      startPreviewOnHoverG()
    }, 200)
  } else {
    previewHoverTimer = setTimeout(() => {
      showMusicPreview.value = false
    }, 150)
  }
}

function togglePreviewPlay() {
  togglePreviewPlayG()
}

function previewNext() {
  previewNextG()
}

function previewPrevious() {
  previewPreviousG()
}

function addPreviewSong() {
  console.log('Add to playlist:', previewSong.value?.title)
}

function openMusicFromPreview() {
  showMusicModal.value = true
}

const navIcons = computed(() => [
  {
    name: 'Message',
    badge: unreadMessageCount.value,
    svg: '<rect x="3" y="5" width="18" height="14" rx="2"/><path d="M3 6l9 7 9-7"/>',
  },
  {
    name: 'Chat',
    badge: unreadChatCount.value,
    svg: '<path d="M21 12a8 8 0 1 1-3.2-6.4L21 4l-1 4.6A7.96 7.96 0 0 1 21 12Z"/>',
  },
  {
    name: 'Notification',
    badge: unreadNotificationCount.value,
    svg: '<path d="M6 10a6 6 0 1 1 12 0c0 5 2 6 2 6H4s2-1 2-6Z"/><path d="M10 20a2 2 0 0 0 4 0"/>',
  },
])

const POLL_INTERVAL = 10000
let badgePollTimer = null

function startBadgePolling() {
  badgePollTimer = setInterval(() => {
    fetchUnreadNotificationCount()
    fetchUnreadChatCount()
  }, POLL_INTERVAL)
}

function stopBadgePolling() {
  if (badgePollTimer) {
    clearInterval(badgePollTimer)
    badgePollTimer = null
  }
}

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

function goToFeed() {
  router.push('/feed')
}

function logout() {
  localStorage.removeItem('token')
  window.location.href = '/login'
}

let lockedScrollY = 0

watch([showMusicModal, showNotificationModal, showChatModal], ([musicVal, notifVal, chatVal]) => {
  const shouldLock = musicVal || notifVal || chatVal
  if (shouldLock) {
    lockedScrollY = window.scrollY
    document.documentElement.style.overflow = 'hidden'
    document.body.style.overflow = 'hidden'
    document.body.style.position = 'fixed'
    document.body.style.top = `-${lockedScrollY}px`
    document.body.style.width = '100%'
  } else {
    document.documentElement.style.overflow = ''
    document.body.style.overflow = ''
    document.body.style.position = ''
    document.body.style.top = ''
    document.body.style.width = ''
    window.scrollTo(0, lockedScrollY)
  }
})


const unreadNotificationCount = ref(0)
const unreadChatCount = ref(0)
const unreadMessageCount = ref(0) // សម្រាប់ icon "Message" ដាច់ដោយឡែក (បើមាន endpoint ខុសគ្នា)

async function fetchUnreadNotificationCount() {
  try {
    const res = await axios.get(`${API_BASE}/notifications/show`, {
      headers: authHeaders(),
    })
    const data = res.data?.data ?? res.data
    unreadNotificationCount.value = data.unread_count ?? data.UnreadCount ?? 0
  } catch (error) {
    console.error('Failed to load unread notification count:', error)
  }
}

async function fetchUnreadChatCount() {
  try {
    const res = await axios.get(`${API_BASE}/chat/unread-count`, {
      headers: authHeaders(),
    })
    const data = res.data?.data ?? res.data
    unreadChatCount.value = data.unread_count ?? data.UnreadCount ?? 0
  } catch (error) {
    console.error('Failed to load unread chat count:', error)
  }
}
onMounted(() => {
  fetchMenus()
  fetchProfile()
  fetchUnreadNotificationCount()
  fetchUnreadChatCount()
  startBadgePolling()
})

onUnmounted(() => {
  stopBadgePolling()
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
  background: rgba(0, 0, 0, 0.593);
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
  border-radius: 12px;
  max-height: 90vh;
  overflow-y: auto;
}

.music-modal-close {
  position: absolute;
  bottom: 0px;
  right: -32px;
  width: 36px;
  height: 36px;
  border: none;
  background-color: transparent;
  color: #e5e3e3;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;

  transition: all 0.2s ease;
}

.music-modal-close:hover {
  opacity: 0.8;
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
  background: rgba(0, 0, 0, 0.4); 
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





.nav-links li { position: relative; }

.music-preview-card {
  position: absolute;
  top: calc(100% + 16px);
  left: 50%;
  transform: translateX(-50%);
  width: 290px;
  height: 120px;
  border-radius: 4px;
  background-color: #1a1a1a;
  background-size: cover;
  background-position: center;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 12px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.4);
  z-index: 50;
  cursor: pointer;
  border-bottom: 4px solid #FADC01 ;
}
.preview-overlay {
  position: absolute; inset: 0;
  background: linear-gradient(to bottom, rgba(0,0,0,.25) 0%, rgba(0,0,0,.15) 40%, rgba(0,0,0,.75) 100%);
}
.preview-top, .preview-controls, .preview-info { position: relative; z-index: 1; }
.preview-top { display: flex; justify-content: space-between; }
.preview-icon-btn {
  width: 30px; height: 30px; border-radius: 50%; border: none;
  background: rgba(0,0,0,.35); color: #fff;
  display: flex; align-items: center; justify-content: center; cursor: pointer;
}
.preview-icon-btn svg { width: 16px; height: 16px; }
.preview-controls { display: flex; align-items: center; justify-content: center; gap: 24px; }
.preview-ctrl-btn { background: none; border: none; color: #fff; cursor: pointer; }
.preview-ctrl-btn svg { width: 22px; height: 22px; }
.preview-ctrl-btn.play svg { width: 30px; height: 30px; }
.preview-info { display: flex; flex-direction: column; color: #fff; text-align: center; }
.preview-title { font-weight: 700; font-size: 15px; }
.preview-singer { font-size: 13px; opacity: .85; }

.fade-preview-enter-active, .fade-preview-leave-active {
  transition: opacity .15s ease, transform .15s ease;
}
.fade-preview-enter-from, .fade-preview-leave-to {
  opacity: 0; transform: translateX(-50%) translateY(-6px);
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