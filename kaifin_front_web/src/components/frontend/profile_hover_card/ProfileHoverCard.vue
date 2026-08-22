<template>
  <div class="hover-card" @mouseenter="$emit('keep-open')" @mouseleave="$emit('close')">
    <div class="hover-cover">
      <img v-if="profile?.coverUrl" :src="profile.coverUrl" alt="" />
      <div v-else class="hover-cover-placeholder"></div>
      <div class="hover-cover-gradient"></div>
    </div>

    <div class="hover-body">
      <div class="hover-avatar">
        <img v-if="profile?.avatarUrl" :src="profile.avatarUrl" alt="" />
        <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
      </div>
      <div v-if="isLoading" class="hover-loading">Loading...</div>
      <template v-else-if="profile">
        <h4 class="hover-name">{{ profile.name }}</h4>
        <p class="hover-handle">{{ profile.handle }}</p>
        <p v-if="profile.bio" class="hover-bio">{{ profile.bio }}</p>

        <div class="hover-meta">
          <span v-if="profile.location" class="hover-meta-item">
            <svg viewBox="0 0 24 24"><path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"/><circle cx="12" cy="10" r="3"/></svg>
            {{ profile.location }}
          </span>
          <span v-if="profile.joinedDate" class="hover-meta-item">
            <svg viewBox="0 0 24 24"><rect x="3" y="4" width="18" height="18" rx="2"/><line x1="16" y1="2" x2="16" y2="6"/><line x1="8" y1="2" x2="8" y2="6"/><line x1="3" y1="10" x2="21" y2="10"/></svg>
            {{ profile.joinedDate }}
          </span>
        </div>

        <div class="hover-actions">
          <button
            v-if="!isOwnProfile"
            class="hover-follow-btn"
            :class="{ following: isFollowing }"
            @click.stop="$emit('toggle-follow')"
          >
            {{ isFollowing ? 'Following' : 'Follow' }}
          </button>
          <button class="hover-view-btn" @click.stop="$emit('view-profile')">
            View Profile
          </button>
        </div>
      </template>
      <p v-else class="hover-error">Failed to load</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const API_BASE = import.meta.env.VITE_API_URL
function authHeaders() {
  const token = localStorage.getItem('token')
  return { Authorization: `Bearer ${token}` }
}

const isStartingChat = ref(false)

async function handleStartChat(profile) {
    console.log('CLICKED, profile:', profile) 
      console.log('profile keys:', Object.keys(profile)) 
  if (!profile?.id || isStartingChat.value) return
  isStartingChat.value = true

  const form = new FormData()
  form.append('target_user_id', profile.id)

  try {
    const res = await axios.post(`${API_BASE}/chats/start`, form, {
      headers: { ...authHeaders(), 'Content-Type': 'multipart/form-data' },
    })
    if (res.data.success) {
      const conversationId = res.data.data.conversation_id
      router.push({ path: '/chart', query: { open: conversationId } }) 
    }
  } catch (err) {
    console.log('START CHAT ERROR:', err)
  } finally {
    isStartingChat.value = false
  }
}

defineProps({
  profile: { type: Object, default: null },
  isLoading: { type: Boolean, default: false },
  isFollowing: { type: Boolean, default: false },
  isOwnProfile: { type: Boolean, default: false },
})
defineEmits(['close', 'keep-open', 'toggle-follow', 'view-profile'])
</script>

<style scoped>
.hover-card {
  width: 480px;
  background: #fff;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.18);
  border: 1px solid #E7E7E7;
  font-family: 'Inter', sans-serif;
  z-index: 100;
}
.hover-cover {
  width: 100%;
  height: 80px;
  background: #EFF6FB;
  overflow: hidden;
}

.hover-cover-gradient {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 50%;
  background: linear-gradient(to top, rgba(255, 255, 255, 0.683), transparent);
  pointer-events: none;
  /* z-index: 90; */
}

.hover-cover img { width: 100%; height: 100%; object-fit: cover; display: block; }
.hover-cover-placeholder { width: 100%; height: 100%; background: linear-gradient(135deg, #1976D2, #64B5F6); }

.hover-body { padding: 0 16px 14px; position: relative; }

.hover-avatar {
  width: 72px; height: 72px;
  border-radius: 50%;
  border: 3px solid #fff;
  background: #EFF6FB;
  margin-top: -36px;
  overflow: hidden;
  display: flex; align-items: center; justify-content: center;
  box-shadow: 0 2px 6px rgba(0,0,0,0.15);
}
.hover-avatar img { width: 100%; height: 100%; object-fit: cover; }
.hover-avatar svg { width: 34px; height: 34px; stroke: #1E6E9C; fill: none; stroke-width: 1.8; }

.hover-name { font-size: 16px; font-weight: 700; margin: 10px 0 2px; color: #0f172a; }
.hover-handle { font-size: 13px; color: #64748b; margin: 0 0 8px; }
.hover-bio { font-size: 13px; color: #334155; line-height: 1.4; margin: 0 0 10px; }

.hover-meta { display: flex; flex-wrap: wrap; gap: 10px; margin-bottom: 12px; }
.hover-meta-item { display: flex; align-items: center; gap: 4px; font-size: 12px; color: #64748b; }
.hover-meta-item svg { width: 13px; height: 13px; stroke: currentColor; fill: none; stroke-width: 2; }

.hover-actions { display: flex; gap: 8px; }
.hover-follow-btn, .hover-view-btn {
  flex: 1;
  border-radius: 32px;
  font-weight: 700;
  font-size: 12.5px;
  padding: 8px 10px;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
}
.hover-follow-btn { border: 2px solid #1976D2; background: #1976D2; color: #fff; }
.hover-follow-btn.following { background: #fff; color: #8A8A8E; border-color: #E7E7E7; }
.hover-view-btn { border: 1px solid #E7E7E7; background: #fff; color: #2B2B2B; }

.hover-loading, .hover-error { font-size: 13px; color: #8A8A8E; padding: 20px 0; text-align: center; }
.hover-cover {
  position: relative;   /* ត្រូវបន្ថែម ដើម្បីឱ្យ gradient absolute ដំណើរការត្រឹមត្រូវ */
}
.hover-chat-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 34px;
  height: 34px;
  border-radius: 50%;
  border: 1px solid #ddd;
  background: #fff;
  cursor: pointer;
  color: #333;
  flex-shrink: 0;
}
.hover-chat-btn:hover {
  background: #f2f2f2;
}


</style>