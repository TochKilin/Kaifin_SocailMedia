<template>
  <div class="notif-wrap" v-click-outside="closeDropdown">
    <button class="notif-btn" @click="toggleDropdown">
      <svg viewBox="0 0 24 24"><path d="M18 8a6 6 0 0 0-12 0c0 7-3 9-3 9h18s-3-2-3-9"/><path d="M13.7 21a2 2 0 0 1-3.4 0"/></svg>
      <span v-if="unreadCount > 0" class="notif-badge">{{ unreadCount > 99 ? '99+' : unreadCount }}</span>
    </button>

    <div v-if="isOpen" class="notif-dropdown">
      <div class="notif-header">
        <span>Notifications</span>
      </div>

      <div class="notif-list">
        <p v-if="isLoading" class="notif-state">Loading...</p>
        <p v-else-if="!notifications.length" class="notif-state">No notifications yet</p>

        <div
          v-for="n in notifications"
          :key="n.id"
          class="notif-item"
          :class="{ unread: !n.is_read }"
          @click="onNotifClick(n)"
        >
          <div class="notif-avatar">
            <img v-if="n.actor_avatar" :src="n.actor_avatar" alt="" />
            <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
          </div>
          <div class="notif-body">
            <p class="notif-text">
              <strong>{{ n.actor_name }}</strong>
              {{ notifText(n.type) }}
            </p>
            <span class="notif-time">{{ formatDatetime(n.created_at) }}</span>
          </div>
          <span v-if="!n.is_read" class="notif-dot"></span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'
const WS_URL = BASE_URL.replace(/^http/, 'ws') // http://localhost:7070 -> ws://localhost:7070

const router = useRouter()
const notifications = ref([])
const unreadCount = ref(0)
const isOpen = ref(false)
const isLoading = ref(false)
let socket = null

function getAuthToken() {
  return localStorage.getItem('token') || ''
}

function authHeaders() {
  const token = getAuthToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}

function resolveAvatarUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

function formatDatetime(value) {
  if (!value) return ''
  const d = new Date(value)
  if (Number.isNaN(d.getTime())) return String(value)
  return d.toLocaleString('km-KH', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

function notifText(type) {
  switch (type) {
    case 'follow':
      return 'started following you'
    case 'like':
      return 'liked your post'
    case 'comment':
      return 'commented on your post'
    default:
      return 'sent you a notification'
  }
}

async function fetchNotifications() {
  isLoading.value = true
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/notifications/show`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json
    const list = data.notifications ?? data.Notifications ?? []
    notifications.value = list.map((n) => ({
      ...n,
      actor_avatar: resolveAvatarUrl(n.actor_avatar),
    }))
    unreadCount.value = data.unread_count ?? data.UnreadCount ?? 0
  } catch (e) {
    console.error('Failed to load notifications', e)
  } finally {
    isLoading.value = false
  }
}

async function markAsRead(notificationId) {
  try {
    await fetch(`${BASE_URL}/api/v1/front/notifications/read`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...authHeaders(),
      },
      body: JSON.stringify({ notification_id: notificationId }),
    })
  } catch (e) {
    console.error('Failed to mark notification as read', e)
  }
}

function onNotifClick(n) {
  if (!n.is_read) {
    n.is_read = true
    unreadCount.value = Math.max(0, unreadCount.value - 1)
    markAsRead(n.id)
  }
  isOpen.value = false
  if (n.type === 'follow' && n.actor_id) {
    router.push(`/profile/${n.actor_id}`)
  }
}

function toggleDropdown() {
  isOpen.value = !isOpen.value
}

function closeDropdown() {
  isOpen.value = false
}

// ============= WebSocket connection =============
function connectWebSocket() {
  const token = getAuthToken()
  if (!token) return

  socket = new WebSocket(`${WS_URL}/api/v1/front/notifications/ws?token=${token}`)

  socket.onopen = () => {
    console.log('🔔 Notification WS connected')
  }

  socket.onmessage = (event) => {
    try {
      const msg = JSON.parse(event.data)
      if (msg.type === 'notification') {
        const n = {
          ...msg.data,
          actor_avatar: resolveAvatarUrl(msg.data.actor_avatar),
        }
        notifications.value.unshift(n)
        unreadCount.value += 1
      }
    } catch (e) {
      console.error('Failed to parse WS message', e)
    }
  }

  socket.onclose = () => {
    console.log('🔔 Notification WS closed, retrying in 3s...')
    setTimeout(connectWebSocket, 3000)
  }

  socket.onerror = (err) => {
    console.error('Notification WS error', err)
    socket.close()
  }
}

// ============= Click outside directive =============
const vClickOutside = {
  mounted(el, binding) {
    el.__clickOutsideHandler = (e) => {
      if (!el.contains(e.target)) binding.value()
    }
    document.addEventListener('click', el.__clickOutsideHandler)
  },
  unmounted(el) {
    document.removeEventListener('click', el.__clickOutsideHandler)
  },
}

onMounted(() => {
  fetchNotifications()
  connectWebSocket()
})

onUnmounted(() => {
  if (socket) socket.close()
})
</script>

<script>
// registers vClickOutside for template usage
export default {
  directives: {
    clickOutside: {
      mounted(el, binding) {
        el.__clickOutsideHandler = (e) => {
          if (!el.contains(e.target)) binding.value()
        }
        document.addEventListener('click', el.__clickOutsideHandler)
      },
      unmounted(el) {
        document.removeEventListener('click', el.__clickOutsideHandler)
      },
    },
  },
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
</style>