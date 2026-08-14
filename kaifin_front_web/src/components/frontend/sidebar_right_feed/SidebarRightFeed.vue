<template>
  <aside class="profile-panel">
    <!-- SVG sprite -->
    <svg width="0" height="0" style="position: absolute" aria-hidden="true">
      <defs>
        <symbol id="pp-icon-user" viewBox="0 0 24 24">
          <circle cx="12" cy="8.3" r="3.6" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <path d="M4.8 20c0-4 3.2-6.6 7.2-6.6s7.2 2.6 7.2 6.6" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
        </symbol>
        <symbol id="pp-icon-post" viewBox="0 0 24 24">
          <rect x="3.2" y="4.5" width="17.6" height="15" rx="2.3" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <path d="M3.2 15l4.6-4.6 3.2 3.2 4.6-5.2 4.4 4.6" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round" stroke-linecap="round"/>
        </symbol>
        <symbol id="pp-icon-group" viewBox="0 0 24 24">
          <circle cx="8.7" cy="8.6" r="2.8" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <circle cx="16.3" cy="9.4" r="2.3" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <path d="M3 19c0-3.3 2.5-5.4 5.7-5.4s5.7 2.1 5.7 5.4" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round"/>
          <path d="M14.6 14.2c2.6.3 4.4 2.1 4.4 4.8" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round"/>
        </symbol>
        <symbol id="pp-icon-following" viewBox="0 0 24 24">
          <circle cx="9.5" cy="8.5" r="3.3" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <path d="M3 19.5c0-3.6 2.9-6 6.5-6s6.5 2.4 6.5 6" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round"/>
          <path d="M17.5 6.5v5M15 9h5" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
        </symbol>
        <symbol id="pp-icon-follower" viewBox="0 0 24 24">
          <circle cx="9.5" cy="8.5" r="3.3" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <path d="M3 19.5c0-3.6 2.9-6 6.5-6s6.5 2.4 6.5 6" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round"/>
          <circle cx="17.5" cy="8.5" r="2.6" fill="none" stroke="currentColor" stroke-width="1.7"/>
        </symbol>
        <symbol id="pp-icon-chevron" viewBox="0 0 24 24">
          <path d="M6 9.5l6 6 6-6" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        </symbol>
        <symbol id="pp-icon-refresh" viewBox="0 0 24 24">
          <path d="M4 12a8 8 0 0 1 13.6-5.7M20 12a8 8 0 0 1-13.6 5.7" fill="none" stroke="currentColor" stroke-width="1.9" stroke-linecap="round"/>
          <path d="M17.2 4.6v3.4h-3.4M6.8 19.4V16h3.4" fill="none" stroke="currentColor" stroke-width="1.9" stroke-linecap="round" stroke-linejoin="round"/>
        </symbol>
        <symbol id="pp-icon-hash" viewBox="0 0 24 24">
          <path d="M9.5 3.5l-3 17M17.5 3.5l-3 17M4 8.5h16M3 15.5h16" stroke="currentColor" stroke-width="1.7" stroke-linecap="round"/>
        </symbol>
        <symbol id="pp-icon-pin-on" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="9" fill="currentColor"/>
          <path d="M8 12.2l2.6 2.6 5.4-5.6" fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        </symbol>
        <symbol id="pp-icon-pin-off" viewBox="0 0 24 24">
          <rect x="4" y="4" width="16" height="16" rx="4.5" fill="none" stroke="currentColor" stroke-width="1.7"/>
        </symbol>
        <symbol id="pp-icon-help" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="9.2" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <path d="M9.3 9.4c.3-1.6 1.6-2.5 3-2.5 1.7 0 3 1.1 3 2.7 0 1.3-.8 1.9-1.8 2.6-.8.6-1.2 1.1-1.2 2.1" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
          <circle cx="12" cy="17.6" r="1.05" fill="currentColor"/>
        </symbol>
        <symbol id="pp-icon-private" viewBox="0 0 24 24">
          <path d="M12 3.5l7.5 3.3v5c0 4.6-3 8.3-7.5 9.7-4.5-1.4-7.5-5.1-7.5-9.7v-5L12 3.5Z" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round"/>
          <path d="M9.2 12l2 2 3.6-4" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round" stroke-linejoin="round"/>
        </symbol>
        <symbol id="pp-icon-image" viewBox="0 0 24 24">
          <rect x="3.2" y="4.5" width="17.6" height="15" rx="2.3" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <circle cx="8.5" cy="9.5" r="1.6" fill="currentColor"/>
          <path d="M3.2 15l4.6-4.6 3.2 3.2 4.6-5.2 4.4 4.6" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round" stroke-linecap="round"/>
        </symbol>
        <symbol id="pp-icon-image-plus" viewBox="0 0 24 24">
          <rect x="3.2" y="4.5" width="14.5" height="15" rx="2.3" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <circle cx="8" cy="9.5" r="1.4" fill="currentColor"/>
          <path d="M3.2 15l3.8-3.8 2.7 2.7 3.5-4 3 3" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round" stroke-linecap="round"/>
          <path d="M19.5 13v5.5M16.7 15.7h5.5" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
        </symbol>
        <symbol id="pp-icon-image-clock" viewBox="0 0 24 24">
          <rect x="2.5" y="4.5" width="14" height="12.5" rx="2" fill="none" stroke="currentColor" stroke-width="1.6"/>
          <circle cx="6.8" cy="8.5" r="1.2" fill="currentColor"/>
          <path d="M2.5 13.5l3.2-3.2 2.3 2.3 3-3.4" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linejoin="round" stroke-linecap="round"/>
          <circle cx="17" cy="17" r="4.3" fill="#fff" stroke="currentColor" stroke-width="1.6"/>
          <path d="M17 14.8v2.4l1.6 1" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
        </symbol>
      </defs>
    </svg>
    <!-- Profile card -->
    <section class="card">
      <div class="profile-row">
    <div
  class="avatar-wrap"
  @mouseenter="openAvatarMenu"
  @mouseleave="scheduleCloseAvatarMenu"
>
  <div class="avatar">
    <img v-if="user.avatarUrl" :src="user.avatarUrl" alt="" />
    <svg v-else class="icon icon-lg"><use href="#pp-icon-user" /></svg>
  </div>

  <div
    class="avatar-menu"
    v-if="showAvatarMenu"
    @mouseenter="keepAvatarMenuOpen"
    @mouseleave="scheduleCloseAvatarMenu"
  >
    <div class="avatar-menu-inner">
      <button class="avatar-menu-item" @click="triggerProfilePicInput" :disabled="isUploadingAvatar">
        <svg class="icon"><use href="#pp-icon-image" /></svg>
        <span>{{ isUploadingAvatar ? 'Waiting to send...' : 'Add profile picture' }}</span>
      </button>
      <button class="avatar-menu-item" @click="emit('add-moment')">
        <svg class="icon"><use href="#pp-icon-image-plus" /></svg>
        <span>Add moment</span>
      </button>
      <button class="avatar-menu-item" @click="emit('view-moments')">
        <svg class="icon"><use href="#pp-icon-image-clock" /></svg>
        <span>Moments</span>
      </button>
    </div>
  </div>

  <input
    ref="profilePicInputRef"
    type="file"
    accept="image/*"
    hidden
    @change="onProfilePicPicked"
  />
</div>


<div class="user-info">
        <span class="user-fullname">{{ user.firstname }} {{ user.lastname }}</span>
        <span class="user-handle">@{{ user.username }}</span>
</div>
        
      </div>
      <div class="badge-row">
        <span class="badge">Level {{ user.level }}</span>
        <span class="badge badge-streak">🔥 {{ user.streak }} streak</span>
      </div>
      <!-- Line -->
      <hr class="divider" />
      <div class="nav-row">
        <button class="nav-item" @click="emit('nav', 'post')">
          <svg class="icon"><use href="#pp-icon-post" /></svg>
          <span class="nav-count">{{ user.counts.posts }}</span>
          <span class="nav-label">Post</span>
        </button>
        <button class="nav-item" @click="emit('nav', 'group')">
          <svg class="icon"><use href="#pp-icon-group" /></svg>
          <span class="nav-count">{{ user.counts.groups }}</span>
          <span class="nav-label">Group</span>
        </button>
        <button class="nav-item" @click="emit('nav', 'following')">
          <svg class="icon"><use href="#pp-icon-following" /></svg>
          <span class="nav-count">{{ user.counts.following }}</span>
          <span class="nav-label">Following</span>
        </button>
        <button class="nav-item" @click="emit('nav', 'follower')">
          <svg class="icon"><use href="#pp-icon-follower" /></svg>
          <span class="nav-count">{{ user.counts.followers }}</span>
          <span class="nav-label">Follower</span>
        </button>
      </div>
    </section>
    <!-- Feed group card -->
    <section class="card">
      <button class="dropdown-head" @click="groupMenuOpen = !groupMenuOpen">
        <span>{{ selectedGroup ? selectedGroup.name : 'Select Feed Group' }}</span>
        <svg class="icon" :class="{ rotated: groupMenuOpen }"><use href="#pp-icon-chevron" /></svg>
      </button>
      <ul v-if="groupMenuOpen" class="dropdown-list">
        <li v-for="g in feedGroups" :key="g.id">
          <button @click="selectGroup(g)">{{ g.name }}</button>
        </li>
      </ul>
      <hr class="divider" />
      <p class="description">
        {{ selectedGroup ? selectedGroup.description : 'Choose a group to see its description.' }}
      </p>
      <div class="stat-row">
        <span class="stat-chip">{{ formatCount(selectedGroup?.likeCount ?? 0) }} Like</span>
        <span class="stat-chip">{{ formatCount(selectedGroup?.commentCount ?? 0) }} comment</span>
      </div>
    </section>
    <!-- Top topic card -->
    <section class="card">
      <div class="topic-head">
        <span class="topic-title-pill">Top Topic</span>
        <button class="refresh-btn" title="Refresh topics" @click="refreshTopics">
          <svg class="icon icon-sm"><use href="#pp-icon-refresh" /></svg>
          Refresh
        </button>
      </div>
      <hr class="divider" />
      <div class="topic-row" v-for="topic in topics" :key="topic.id">
        <svg class="icon icon-sm hash-icon"><use href="#pp-icon-hash" /></svg>
        <input class="topic-input" v-model="topic.tag" placeholder="topic" />
        <button
          class="pin-btn"
          :title="topic.pinned ? 'Unpin topic' : 'Pin topic'"
          @click="togglePin(topic)"
        >
          <svg class="icon icon-sm"><use :href="topic.pinned ? '#pp-icon-pin-on' : '#pp-icon-pin-off'" /></svg>
        </button>
      </div>
    </section>
    <!-- private card -->
    <section class="card menu-card">
      <button class="menu-item" @click="emit('help')">
        <span class="menu-icon-wrap">
          <svg class="icon"><use href="#pp-icon-help" /></svg>
        </span>
        <span class="menu-pill">Help</span>
      </button>
      <button class="menu-item" @click="emit('private')">
        <span class="menu-icon-wrap">
          <svg class="icon"><use href="#pp-icon-private" /></svg>
        </span>
        <span class="menu-pill">Private</span>
      </button>
    </section>
  </aside>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'
const showAvatarMenu = ref(false)
let avatarMenuCloseTimer = null

const props = defineProps({
  feedGroups: {
    type: Array,
    default: () => [
      { id: 1, name: 'Web Dev Cambodia', description: 'Discussions on frontend, backend, and everything web dev in Khmer & English.', likeCount: 1000, commentCount: 100 },
      { id: 2, name: 'UI/UX Wireframes', description: 'Sketches, wireframes, and feedback on early-stage product design.', likeCount: 540, commentCount: 62 },
      { id: 3, name: 'Vue Learners KH', description: 'A community for people learning Vue 3, Composition API, and Nuxt.', likeCount: 812, commentCount: 91 },
    ],
  },
})

const emit = defineEmits(['nav', 'select-group', 'refresh-topics', 'help', 'private', 'add-profile-picture', 'add-moment', 'view-moments'])

// user ត្រូវផ្លាស់ប្តូរពី prop ទៅជា reactive state ដែល component fetch ដោយខ្លួនឯង
const user = ref({
  avatarUrl: '',
  username: '',
  firstname: '',
  lastname: '',
  level: 1,
  streak: 0,
  counts: { posts: 0, groups: 0, following: 0, followers: 0 },
})

function getAuthToken() {
  return localStorage.getItem('token') || ''
}

function authHeaders() {
  const token = getAuthToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}

function getCurrentUserId() {
  const token = getAuthToken()
  if (!token) return null
  try {
    const payload = token.split('.')[1]
    const decoded = JSON.parse(
      decodeURIComponent(
        atob(payload.replace(/-/g, '+').replace(/_/g, '/'))
          .split('')
          .map((c) => '%' + c.charCodeAt(0).toString(16).padStart(2, '0'))
          .join('')
      )
    )
    return decoded.user_id ?? decoded.uid ?? decoded.sub ?? decoded.id ?? null
  } catch {
    return null
  }
}

function resolveAvatarUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

async function loadProfile() {
  try {
    const res = await fetch(
      `${BASE_URL}/api/v1/front/profile/show`,
      {
        headers: {
          ...authHeaders(),
        },
      }
    )

    if (!res.ok) {
      console.log("Profile API Error:", res.status)
      return
    }

    const json = await res.json()

    console.log("PROFILE:", json)

    const data = json?.data ?? json


    user.value.username =
      data.user_name ?? "Username"
          user.value.firstname = data.first_name ?? ''
    user.value.lastname = data.last_name ?? ''

    user.value.avatarUrl =
      resolveAvatarUrl(data.profile_images)


    user.value.counts.posts =
      data.post_count ?? 0

    user.value.counts.followers =
      data.follower_count ?? 0

    user.value.counts.following =
      data.following_count ?? 0


  } catch (e) {
    console.error(
      "Failed to load profile",
      e
    )
  }
}

// ទាញ following / follower count
// async function loadFollowCounts() {
//   const currentUserId = getCurrentUserId()
//   if (!currentUserId) return

//   try {
//     const res = await fetch(`${BASE_URL}/api/v1/front/followers/show?user_id=${currentUserId}`, {
//       headers: { ...authHeaders() },
//     })
//     if (!res.ok) return
//     const json = await res.json()
//     const data = json?.data ?? json
//     user.value.counts.following = data.following_count ?? 0
//     user.value.counts.followers = data.follower_count ?? 0
//   } catch (e) {
//     console.error('Failed to load follow counts', e)
//   }
// }

// ទាញ post count (សន្មតថា /posts/show ទទួល filter user_id — ត្រូវផ្ទៀងផ្ទាត់ជាមួយ backend)
// async function loadPostCount() {
//   const currentUserId = getCurrentUserId()
//   if (!currentUserId) return

//   try {
//     const res = await fetch(
//       `${BASE_URL}/api/v1/front/posts/show?page=1&perpage=1&filters[0][field]=user_id&filters[0][value]=${currentUserId}`,
//       { headers: { ...authHeaders() } }
//     )
//     if (!res.ok) return
//     const json = await res.json()
//     const data = json?.data ?? json
//     user.value.counts.posts = data.total ?? data.Total ?? 0
//   } catch (e) {
//     console.error('Failed to load post count', e)
//   }
// }

// onMounted(async () => {
//   await Promise.all([loadProfile(), loadFollowCounts(), loadPostCount()])
// })

// onMounted(async () => {
//   await Promise.all([
//     loadProfile(),
//     loadFollowCounts(),
//     loadPostCount()
//   ])
// })

onMounted(() => {
  loadProfile()
})

const groupMenuOpen = ref(false)
const selectedGroup = ref(props.feedGroups[0] ?? null)

function selectGroup(g) {
  selectedGroup.value = g
  groupMenuOpen.value = false
  emit('select-group', g)
}

const topicPool = [
  'vuejs', 'frontend', 'uiux', 'khmerdev', 'wireframe',
  'design', 'javascript', 'webdev', 'startup', 'opensource',
]

const topics = ref([
  { id: 1, tag: 'vuejs', pinned: true },
  { id: 2, tag: 'frontend', pinned: false },
  { id: 3, tag: 'uiux', pinned: false },
])

function togglePin(topic) {
  topic.pinned = !topic.pinned
}

function refreshTopics() {
  const shuffled = [...topicPool].sort(() => Math.random() - 0.5)
  topics.value = topics.value.map((t, i) => ({ ...t, tag: shuffled[i] || t.tag }))
  emit('refresh-topics', topics.value)
}

function formatCount(n) {
  if (n >= 1000) return (n / 1000).toFixed(n % 1000 === 0 ? 0 : 1) + 'k'
  return String(n)
}

function openAvatarMenu() {
  clearTimeout(avatarMenuCloseTimer)
  showAvatarMenu.value = true
}

function scheduleCloseAvatarMenu() {
  clearTimeout(avatarMenuCloseTimer)
  avatarMenuCloseTimer = setTimeout(() => {
    showAvatarMenu.value = false
  }, 350)
}

function keepAvatarMenuOpen() {
  clearTimeout(avatarMenuCloseTimer)
}

const profilePicInputRef = ref(null)
const isUploadingAvatar = ref(false)


function triggerProfilePicInput() {
  profilePicInputRef.value?.click()
}

async function onProfilePicPicked(e) {
  const file = e.target.files?.[0]
  if (!file) return
  e.target.value = ''

  const currentId = getCurrentUserId()
  if (!currentId) {
    console.error('No current user id')
    return
  }

  showAvatarMenu.value = false
  isUploadingAvatar.value = true

  const previousAvatarUrl = user.value.avatarUrl
  user.value.avatarUrl = URL.createObjectURL(file)

  try {
    const formData = new FormData()
    formData.append('file', file)

    const res = await fetch(`${BASE_URL}/api/v1/front/profile/update/${currentId}`, {
      method: 'PUT',
      headers: { ...authHeaders() },
      body: formData,
    })

    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status}: ${text}`)
    }

    const json = await res.json()
    const data = json?.data ?? json

    if (data?.profile_images) {
      user.value.avatarUrl = resolveAvatarUrl(data.profile_images)
    }
  } catch (err) {
    console.error('Failed to upload profile picture', err)
    user.value.avatarUrl = previousAvatarUrl
  } finally {
    isUploadingAvatar.value = false
  }
}


</script>

<style scoped>
.profile-panel {
  font-family: 'Inter', sans-serif;
  color: #2b2b2b;
  /* width: 320px; */
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 12px;



scrollbar-width: none;

  width: 100%;          /* ឱ្យវារត់ពេញទំហំ Column របស់ Bootstrap ស្វ័យប្រវត្តិ */
  max-width: 320px;
}

.profile-panel::-webkit-scrollbar,
.card::-webkit-scrollbar {
  display: none; 
}

.card {
  background: #ffffff;
  border: 1px solid #E5E7EB;
  border-radius: 12px;
  padding: 18px;


  scrollbar-width: none;
  width: 100%;
  overflow: hidden;
  box-sizing: border-box;
}

.divider {
  border: none;
  border-top: 1px solid #EFE2D3;
  margin: 14px 0;
}

/* Profile card */
.profile-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar {
  width: 46px;
  height: 46px;
  border-radius: 50%;
  background: #eff6fb;
  border: 2px solid #1976d2;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  overflow: hidden;
  color: #1976d2;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.username-pill {
  flex: 1;
  /* background: #eff6fb; */
  /* border-radius: 999px; */
  /* padding: 11px 16px; */
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 15px;
  color: #2b2b2b;
}



.badge-row {
  display: flex;
  gap: 8px;
  margin-top: 12px;
}

.badge {
  background: #eff6fb;
  color: #1976d2;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 11.5px;
  padding: 4px 10px;
  border-radius: 999px;
}

.badge-streak {
  background: #fff4e5;
  color: #c2760a;
}

.nav-row {
  display: flex;
  justify-content: space-between;
  gap: 6px;
}

.nav-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  background: none;
  border: none;
  cursor: pointer;
  color: #6b7280;
  padding: 8px 4px;
  border-radius: 12px;
}

.nav-item:hover {
  background: #eff6fb;
  color: #1976d2;
}

.nav-count {
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 13.5px;
  color: #2b2b2b;
}

.nav-item:hover .nav-count {
  color: #1976d2;
}

.nav-label {
  font-size: 11px;
  font-weight: 600;
}

/* Feed group card */
.dropdown-head {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  /* background: #eff6fb; */
  border: none;
  border-radius: 12px;
  padding: 12px 14px;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 14px;
  color: #2b2b2b;
  cursor: pointer;
}

.icon.rotated {
  transform: rotate(180deg);
}

.dropdown-list {
  list-style: none;
  margin: 8px 0 0;
  padding: 0;
  border: 1px solid #e7e7e7;
  border-radius: 12px;
  overflow: hidden;
}

.dropdown-list li + li {
  border-top: 1px solid #e7e7e7;
}

.dropdown-list button {
  width: 100%;
  text-align: left;
  background: none;
  border: none;
  padding: 10px 14px;
  font-size: 13.5px;
  cursor: pointer;
  color: #2b2b2b;
}

.dropdown-list button:hover {
  background: #eff6fb;
}

.description {
  font-size: 13px;
  line-height: 1.55;
  color: #6b7280;
  margin: 0 0 12px;
}

.stat-row {
  display: flex;
  gap: 8px;
}

.stat-chip {
  background: #eff6fb;
  color: #1976d2;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 12px;
  padding: 6px 12px;
  border-radius: 999px;
}

/* Top topic card */
.topic-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.topic-title-pill {
  /* background: #eff6fb; */
  color: #2b2b2b;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 14px;
  padding: 8px 16px;
  border-radius: 999px;
}

.refresh-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: #1976d2;
  color: #fff;
  border: none;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 12.5px;
  padding: 8px 14px;
  border-radius: 999px;
  cursor: pointer;
}

.refresh-btn:hover {
  opacity: 0.88;
}

.topic-row {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 10px;
}

.hash-icon {
  color: #6b7280;
  flex-shrink: 0;
}

.topic-input {
  flex: 1;
  min-width: 0;
  /* background: #f5f6f8; */
  border: 1px solid #e7e7e7;
  border-radius: 999px;
  padding: 8px 14px;
  font-size: 13px;
  color: #2b2b2b;
  outline: none;
}

.topic-input:focus {
  border-color: #1976d2;
}

.pin-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #6b7280;
  flex-shrink: 0;
  width: 26px;
  height: 26px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pin-btn:hover {
  color: #1976d2;
}

/* Help / private card */
.menu-card {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  text-align: left;
}

.menu-icon-wrap {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: #eff6fb;
  color: #000;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.menu-pill {
  flex: 1;
  /* background: #eff6fb; */
  border-radius: 999px;
  /* padding: 11px 16px; */
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 14px;
  color: #2b2b2b;
}

/* .menu-item:hover .menu-pill {
  background: #dceafa;
} */

/* Icon base */
.icon {
  width: 18px;
  height: 18px;
  flex-shrink: 0;
}

.icon-sm {
  width: 15px;
  height: 15px;
}

.icon-lg {
  width: 24px;
  height: 24px;
}

.avatar-wrap {
  position: relative;
  display: inline-flex;
}

.avatar-menu {
  position: absolute;
  top: 100%;
  left: 0;
  z-index: 30;
  padding-top: 8px; 
}

.avatar-menu-inner {
  position: relative;
  background: #fff;
  border: 1px solid #e7e7e7;
  border-radius: 6px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .15);
  padding: 6px;
  min-width: 200px;
  animation: popIn .15s ease;
}


@keyframes popIn {
  from { opacity: 0; transform: translateY(-4px); }
  to   { opacity: 1; transform: translateY(0); }
}

.avatar-menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
  text-align: left;
  border: none;
  background: transparent;
  padding: 4px 4px;
  border-radius: 10px;
  font-size: 12px;
  font-weight: 500;
  color: #2b2b2b;
  cursor: pointer;
}

.avatar-menu-item:hover {
  background: #f2f2f3;
}

.avatar-menu-item .icon {
  width: 20px;
  height: 20px;
  color: #4a4a4e;
  flex-shrink: 0;
}

.avatar-menu-inner::before {
  content: '';
  position: absolute;
  bottom: 100%;
  left: 20px;
  width: 0;
  height: 0;
  border-left: 8px solid transparent;
  border-right: 8px solid transparent;
  border-bottom: 8px solid #e7e7e7;
}

.avatar-menu-inner::after {
  content: '';
  position: absolute;
  bottom: 100%;
  left: 21px;
  width: 0;
  height: 0;
  border-left: 7px solid transparent;
  border-right: 7px solid transparent;
  border-bottom: 7px solid #fff;
}

.user-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
  min-width: 0;
}

.user-fullname {
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 14px;
  color: #2b2b2b;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-handle {
  font-family: 'Inter', sans-serif;
  font-weight: 500;
  font-size: 10px;
  color: #6b7280;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>