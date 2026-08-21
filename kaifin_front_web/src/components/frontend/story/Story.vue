<template>
  <div class="story-container">
    <div class="story-item">
      <button class="btn-circle btn-add" @click="triggerFileInput">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <line x1="12" y1="5" x2="12" y2="19"></line>
          <line x1="5" y1="12" x2="19" y2="12"></line>
        </svg>
      </button>
      <span class="story-label">Add</span>
      <input ref="fileInputRef" type="file" accept="image/*,video/*" hidden @change="onFilePicked" />
    </div>

    <p v-if="isLoading" class="state-msg">Waiting...</p>
    <p v-else-if="error" class="state-msg error">{{ error }}</p>

    <div class="story-item" v-for="group in groupedStories" :key="group.userId" @click="openStoryGroup(group)">
      <button class="btn-circle btn-eye-story" :class="{ 'is-viewed': group.isViewed }">
        <div class="icon-wrapper avatar-wrapper">
          <img v-if="group.thumbType === 'image' && group.thumbUrl" :src="group.thumbUrl" alt="" class="story-avatar" />
          <video v-else-if="group.thumbType === 'video' && group.thumbUrl" :src="group.thumbUrl" class="story-avatar" muted preload="metadata"></video>
          <svg v-else class="eye-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="9" r="3.4"/>
            <path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/>
          </svg>
        </div>

        <div class="eye-badge">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
        <circle cx="12" cy="12" r="3"></circle>
      </svg>
    </div>
      </button>
      <span class="story-label">{{ group.username }}</span>
    </div>

    <div class="story-viewer" v-if="activeGroup" @click="closeStory">
  <div class="story-viewer-content" @click.stop>
  <div class="progress-row">
    <div
      class="progress-bar"
      v-for="(s, i) in activeGroup.stories"
      :key="s.id"
    >
      <div
        class="progress-fill"
        :class="{ filled: i < activeIndex, active: i === activeIndex }"
      ></div>
    </div>
  </div>
          <div class="viewer-header">
  <div class="viewer-user-info">
    <div class="viewer-avatar">
      <img v-if="activeGroup.avatarUrl" :src="activeGroup.avatarUrl" alt="" />
      <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
    </div>
    <div class="viewer-user-text">
      <span class="viewer-username">{{ activeGroup.username }}</span>
      <span class="viewer-time">{{ timeAgo(currentStory?.createdAt) }}</span>
    </div>
  </div>

    <div class="viewer-header-actions">
      <div class="more-wrap">
        <button type="button" class="viewer-more" @click.stop="showMoreMenu = !showMoreMenu">
          <svg viewBox="0 0 24 24" fill="currentColor"><circle cx="5" cy="12" r="1.6"/><circle cx="12" cy="12" r="1.6"/><circle cx="19" cy="12" r="1.6"/></svg>
        </button>
        <div class="more-menu" v-if="showMoreMenu" @click.stop>
          <button type="button" class="danger" @click="handleDeleteStory">លុប Story</button>
        </div>
      </div>
      <button type="button" class="viewer-close" @click.stop="closeStory">✕</button>
    </div>
  </div>
        
        <div class="viewer-media-wrap">
  <!-- Blurred background layer -->
  <div
    class="viewer-bg-blur"
    :style="{ backgroundImage: currentStory.mediaType === 'image' ? `url(${currentStory.mediaUrl})` : 'none' }"
  ></div>

  <button class="nav-zone nav-prev" @click="prevStory"></button>
  <img
    v-if="currentStory.mediaType === 'image'"
    :src="currentStory.mediaUrl"
    class="viewer-media"
  />
    <video
      v-else
      :src="currentStory.mediaUrl"
      class="viewer-media"
      controls
      autoplay
      @ended="nextStory"
    ></video>
    <button class="nav-zone nav-next" @click="nextStory"></button>
  </div>

  <div class="action-bar">
<div class="action-wrap-single">
  <button ref="heartBtn" class="action-btn" :class="{ active: currentStory?.reaction === 'heart' }" @click="pick(REACTIONS.find(r => r.key === 'heart'), $event)">
    <img src="../../../assets/animate/hearth.svg" alt="icon">
    <span class="count-total">{{ formatCount(currentStory?.reactionCounts?.heart || 0) }}</span>
  </button>
  <span class="action-tooltip">Love</span>
</div>


<div class="action-wrap-single">
  <button class="action-btn" :class="{ active: currentStory?.reaction === 'congrate' }" @click="pick(REACTIONS.find(r => r.key === 'congrate'), $event)">
    <img src="../../../assets/animate/congrate.svg" alt="icon">
    <span class="count-total">{{ formatCount(currentStory?.reactionCounts?.congrate || 0) }}</span>
  </button>
  <span class="action-tooltip">Congrate</span>
</div>

  <div class="action-wrap-single">
  <button class="action-btn" :class="{ active: currentStory?.reaction === 'cool' }" @click="pick(REACTIONS.find(r => r.key === 'cool'), $event)">
    <img src="../../../assets/animate/cool.svg" alt="icon">
    <span class="count-total">{{ formatCount(currentStory?.reactionCounts?.cool || 0) }}</span>
  </button>
  <span class="action-tooltip">Amazing</span>
</div>

<div class="action-wrap-single">
  <button class="action-btn" :class="{ active: currentStory?.reaction === 'thinking' }" @click="pick(REACTIONS.find(r => r.key === 'thinking'), $event)">
    <img src="../../../assets/animate/Thinking.svg" alt="">
    <span class="count-total">{{ formatCount(currentStory?.reactionCounts?.thinking || 0) }}</span>
  </button>
  <span class="action-tooltip">Thinking</span>
</div>

    <div class="action-wrap-single">
      <button class="action-btn btn-com" @click="onStoryComment">
       <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="1.5"></circle>
      <circle cx="19" cy="12" r="1.5"></circle>
      <circle cx="5" cy="12" r="1.5"></circle>
       </svg>
        <span class="count-total">18.2k</span>
      </button>
      <span class="action-tooltip">Comment</span>
    </div>

      <div class="action-wrap-single">
        <!-- <button class="action-btn btn-com" @click="onStoryShare">
  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/>
    <polyline points="16 6 12 2 8 6"/>
    <line x1="12" y1="2" x2="12" y2="15"/>
  </svg>


          <span class="count-total">22.2k</span>
        </button> -->
        <span class="action-tooltip">Share</span>
      </div>
      </div>
      <!-- Animation sgap -->
      <div class="emoji-fly-layer" ref="emojiLayer"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { gsap } from "gsap"
import Heart from "../../../assets/animate/hearth.svg"
import Congrate from "../../../assets/animate/congrate.svg"
import Cool from "../../../assets/animate/cool.svg"
import Thinking from "../../../assets/animate/Thinking.svg"

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'
const IMAGE_DURATION_MS = 5000 

const stories = ref([])
const isLoading = ref(false)
const error = ref(null)

const fileInputRef = ref(null)
const activeGroup = ref(null)
const activeIndex = ref(0)
let autoAdvanceTimer = null
const showMoreMenu = ref(false) 

function getAuthToken() {
  return localStorage.getItem('token') || ''
}

function authHeaders() {
  const token = getAuthToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}

function resolveMediaUrl(url) {
  if (!url) return ''
  if (/^https?:\/\//i.test(url)) return url
  return `${BASE_URL}${url.startsWith('/') ? '' : '/'}${url}`
}

async function loadStories() {
  isLoading.value = true
  error.value = null
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/stories/show?page=1&perpage=50`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) throw new Error(`API ${res.status}`)
    const json = await res.json()
    const rawList = json?.data?.stories ?? []
      console.log('STORY RESPONSE:', json?.data?.stories) 
    stories.value = rawList.map(mapStory)
  } catch (e) {
    error.value = e.message || 'Failed to load stories'
  } finally {
    isLoading.value = false
  }

}

function mapStory(s) {
  return {
    id: s.id,
    userId: s.user_id,
    username: s.user_name || `User #${s.user_id}`,
    // avatarUrl: resolveMediaUrl(s.profile_images),
     avatarUrl: resolveAvatarUrl(s.profile_images), 
    mediaUrl: resolveMediaUrl(s.media_url),
    mediaType: s.media_type,
    createdAt: s.created_at,
    reaction: null,          
    reactionCounts: {},   
  }
}

async function fetchStoryReactions(storyId) {
  try {
    const res = await fetch(
      `${BASE_URL}/api/v1/front/story-reactions/show?story_id=${storyId}`,
      { headers: { ...authHeaders() } }
    )
    if (!res.ok) return null
    const json = await res.json()
    return json?.data ?? null
  } catch (err) {
    console.error('Failed to load story reactions', err)
    return null
  }
}

async function loadReactionsForStory(story) {
  const currentUserId = getCurrentUserId()   // ត្រូវមាន function នេះ ឬសរសេរបន្ថែម
  const data = await fetchStoryReactions(story.id)
  if (!data) return
  const counts = {}
  for (const s of data.summary || []) {
    counts[s.reaction_type] = s.count
  }
  story.reactionCounts = counts

  const mine = (data.reactions || []).find(r => r.user_id === currentUserId)
  story.reaction = mine ? mine.reaction_type : null
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

const groupedStories = computed(() => {
  const map = new Map()

  for (const s of stories.value) {
    if (!map.has(s.userId)) {
      map.set(s.userId, {
        userId: s.userId,
        username: s.username,
        avatarUrl: s.avatarUrl, 
        stories: [],
      })
    }
    map.get(s.userId).stories.push(s)
  }

  return Array.from(map.values()).map((group) => {
    group.stories.sort((a, b) => new Date(a.createdAt) - new Date(b.createdAt))
    const latest = group.stories[group.stories.length - 1]
    return {
      ...group,
      thumbUrl: latest.mediaUrl,
      thumbType: latest.mediaType,
       isViewed: viewedUserIds.value.has(group.userId),
    }
  })
})

function timeAgo(dateStr) {
  if (!dateStr) return ''
  const diffMs = Date.now() - new Date(dateStr).getTime()
  const diffMin = Math.floor(diffMs / 60000)
  if (diffMin < 1) return 'Now'
  if (diffMin < 60) return `${diffMin} mn ago`
  const diffHr = Math.floor(diffMin / 60)
  return `${diffHr} h ago`
}

const currentStory = computed(() => {
  return activeGroup.value?.stories?.[activeIndex.value] ?? null
})

function triggerFileInput() {
  fileInputRef.value?.click()
}

async function onFilePicked(e) {
  const file = e.target.files?.[0]
  e.target.value = ''
  if (!file) return

  const mediaType = file.type.startsWith('video/') ? 'video' : 'image'

  const formData = new FormData()
  formData.append('media_url', file)
  formData.append('media_type', mediaType)

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/stories/create`, {
      method: 'POST',
      headers: { ...authHeaders() },
      body: formData,
    })
    if (!res.ok) {
      const data = await res.json().catch(() => null)
      throw new Error(data?.message || `Upload failed (${res.status})`)
    }
    triggerStreakCheckIn()
    await loadStories()
  } catch (err) {
    console.error('Story upload failed', err)
    error.value = err.message || 'Upload failed'
  }
}

async function handleDeleteStory() {
  showMoreMenu.value = false
  if (!currentStory.value) return

  const storyId = currentStory.value.id
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/stories/delete/${storyId}`, {
      method: 'DELETE',
      headers: { ...authHeaders() },
    })
    if (!res.ok) {
      const data = await res.json().catch(() => null)
      throw new Error(data?.message || `Delete failed (${res.status})`)
    }

    stories.value = stories.value.filter((s) => s.id !== storyId)
    const remaining = activeGroup.value.stories.filter((s) => s.id !== storyId)
    if (remaining.length === 0) {
      closeStory()
    } else {
      activeGroup.value = { ...activeGroup.value, stories: remaining }
      if (activeIndex.value >= remaining.length) {
        activeIndex.value = remaining.length - 1
      }
      startAutoAdvance()
    }
  } catch (err) {
    console.error('Delete story failed', err)
    error.value = err.message || 'Delete failed'
  }
}

function openStoryGroup(group) {
  activeGroup.value = group
  activeIndex.value = 0
  startAutoAdvance()
  loadReactionsForStory(currentStory.value)

   viewedUserIds.value.add(group.userId)
}

function closeStory() {
  activeGroup.value = null
  activeIndex.value = 0
  stopAutoAdvance()
  showMoreMenu.value = false 
}

function nextStory() {
  if (!activeGroup.value) return
  if (activeIndex.value < activeGroup.value.stories.length - 1) {
    activeIndex.value++
    startAutoAdvance()
    loadReactionsForStory(currentStory.value)
  } else {
    closeStory()
  }
}

function prevStory() {
  if (!activeGroup.value) return
  if (activeIndex.value > 0) {
    activeIndex.value--
    startAutoAdvance()
    loadReactionsForStory(currentStory.value)
  }
}

function startAutoAdvance() {
  stopAutoAdvance()
  if (currentStory.value?.mediaType === 'image') {
    autoAdvanceTimer = setTimeout(() => {
      nextStory()
    }, IMAGE_DURATION_MS)
  }
}

function stopAutoAdvance() {
  if (autoAdvanceTimer) {
    clearTimeout(autoAdvanceTimer)
    autoAdvanceTimer = null
  }
}


function resolveAvatarUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

async function triggerStreakCheckIn() {
  try {
    await fetch(`${BASE_URL}/api/v1/front/levels/checkin`, {
      method: 'POST',
      headers: { ...authHeaders() },
    })
  } catch (err) {
    console.error('Streak check-in failed', err)
  }
}



const props = defineProps({
  reaction: { type: String, default: null }, // current reaction key, e.g. 'heart' | 'smile' | 'neutral' | null
})
const emit = defineEmits(['react', 'comment', 'share'])
 
const REACTIONS = [
  {
    key: "heart",
    image: Heart,
  },
  {
    key: "congrate",
    image: Congrate,
  },
  {
    key: "cool",
    image: Cool,
  },
  {
    key: "thinking",
    image: Thinking,
  }
]
 
const currentIcon = computed(() => {
  const found = REACTIONS.find((r) => r.key === currentStory.value?.reaction)
  return found ? found.icon : ''
})
 
const showPicker = ref(false)
let closeTimer = null
 
function openPicker() {
  clearTimeout(closeTimer)
  showPicker.value = true
}
function schedulePickerClose() {
  clearTimeout(closeTimer)
  closeTimer = setTimeout(() => { showPicker.value = false }, 250)
}
function keepPickerOpen() {
  clearTimeout(closeTimer)
}
 
function onHeartClick() {
  pick(REACTIONS.find((r) => r.key === 'heart'))
}
 
async function pick(r, event) {
  flyEmoji(r.image, event)
  showPicker.value = false
  if (!currentStory.value) return

  const story = currentStory.value
    console.log('Reacting to story_id:', story.id, typeof story.id)
  const isUnreact = story.reaction === r.key
  const previousReaction = story.reaction
  const previousCounts = { ...story.reactionCounts }

  if (isUnreact) {
    story.reaction = null
    story.reactionCounts[r.key] = Math.max((story.reactionCounts[r.key] || 1) - 1, 0)
  } else {
    if (previousReaction) {
      story.reactionCounts[previousReaction] = Math.max((story.reactionCounts[previousReaction] || 1) - 1, 0)
    }
    story.reaction = r.key
    story.reactionCounts[r.key] = (story.reactionCounts[r.key] || 0) + 1
  }

  try {
    if (isUnreact) {
      const res = await fetch(`${BASE_URL}/api/v1/front/story-reactions/delete/${story.id}`, {
        method: 'DELETE',
        headers: { ...authHeaders() },
      })
      if (!res.ok) throw new Error(`Delete failed (${res.status})`)
    } else {
      const res = await fetch(`${BASE_URL}/api/v1/front/story-reactions/react`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          ...authHeaders(),
        },
        body: JSON.stringify({
          story_id: story.id,
          reaction_type: r.key,
        }),
      })
      if (!res.ok) {
        const data = await res.json().catch(() => null)
        throw new Error(data?.message || `React failed (${res.status})`)
      }
    }
  } catch (err) {
    console.error('React to story failed', err)
    story.reaction = previousReaction
    story.reactionCounts = previousCounts
  }
}

function formatCount(n) {
  if (n >= 1000) return (n / 1000).toFixed(n % 1000 === 0 ? 0 : 1) + 'k'
  return String(n)
}

function onStoryComment() {
  console.log('Comment on story', currentStory.value?.id)

}

function onStoryShare() {
  console.log('Share story', currentStory.value?.id)

}




const emojiLayer = ref(null)

function flyEmoji(imgSrc, event) {

  const btn = event.currentTarget

  const start = btn.getBoundingClientRect()
  const layer = emojiLayer.value.getBoundingClientRect()

  const emoji = document.createElement("img")

  emoji.src = imgSrc
  emoji.style.position = "absolute"
  emoji.style.width = "55px"
  emoji.style.height = "55px"

  emoji.style.left = `${start.left - layer.left}px`
  emoji.style.top = `${start.top - layer.top}px`
  emoji.style.pointerEvents = "none"

  emojiLayer.value.appendChild(emoji)

  const endX = layer.width / 2 - (start.left - layer.left) - 28
  const endY = layer.height / 2 - (start.top - layer.top) - 28

  const tl = gsap.timeline()


  tl.to(emoji, {
    duration: 3.5,
    x: endX,
    y: endY,
    scale: 2,
    rotation: 360,
    ease: "power3.out"
  })


  .to(emoji, {
    scale: 2.8,
    duration: 0.15,
    ease: "back.out(4)"
  })


  .to(emoji, {
    scale: 2.2,
    duration: 0.1
  })

  
  .call(() => {

    for (let i = 0; i < 18; i++) {

      const p = document.createElement("div")

      p.className = "particle"

      p.style.left = `${layer.width / 2}px`
      p.style.top = `${layer.height / 2}px`

      emojiLayer.value.appendChild(p)

      const angle = (Math.PI * 2 / 18) * i
      const distance = 70 + Math.random() * 40

      gsap.to(p, {
        duration: 0.6,
        x: Math.cos(angle) * distance,
        y: Math.sin(angle) * distance,
        scale: 0,
        opacity: 0,
        ease: "power2.out",
        onComplete() {
          p.remove()
        }
      })
    }

  })


  .to(emoji, {
    opacity: 0,
    scale: 3,
    duration: 0.25,
    onComplete() {
      emoji.remove()
    }
  })

}

const viewedUserIds = ref(new Set())

onMounted(() => {
  loadStories()
})

onUnmounted(() => {
  stopAutoAdvance()
})
</script>

<style scoped>
.particle{
    position:absolute;
    width:10px;
    height:10px;
    border-radius:50%;
    background:#ff3366;
    pointer-events:none;
}
.story-container {
   width: 100%;
   margin: auto;
  display: flex;
  gap: 20px;
  padding: 20px;
  font-family: sans-serif;
  overflow-x: auto;
  white-space: nowrap;
  scrollbar-width: none;
}

.story-container::-webkit-scrollbar {
  display: none;
}

.story-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  flex-shrink: 0;
}

.btn-circle {
  width: 70px;
  height: 70px;
  border-radius: 50%;
  padding: 0;
  border: none;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: all 0.3s ease;
  cursor: pointer;
  background-color: #e5e5e5;
  color: #333;
}

.btn-add {
  color: #555;
}

.story-item:hover .btn-add {
  /* background-color: #1976D2; */
  box-shadow: 0 0 0 4px #1976D2;
  transform: scale(1.05);
}

.btn-eye-story {
  background-color: #1976D2;
  color: white;
  border: 4px solid #fff;
  box-shadow: 0 0 0 2px #1976D2;
  overflow: hidden;
  position: relative;
}

.eye-badge {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 28px;
  height: 28px;
  border-radius: 50%;
  border: 2px solid #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  z-index: 2;
}

.eye-badge svg {
  width: 12px;
  height: 12px;
  stroke: currentColor;
  fill: none;
  stroke-width: 2.2;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.icon-wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-wrapper {
  border-radius: 50%;
  overflow: hidden;
}

.story-avatar {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.eye-icon {
  width: 28px;
  height: 28px;
  transition: transform 0.3s ease;
}

.story-label {
  font-size: 14px;
  color: #333;
  font-weight: 500;
  max-width: 70px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.story-item:hover .btn-eye-story {
  background-color: #1976D2;
  box-shadow: 0 0 0 4px #1976D2;
  transform: scale(1.05);
}

.story-item:hover .eye-icon {
  transform: scale(1.15);
}

.state-msg {
  font-size: 13px;
  color: #4A4A4E;
  white-space: nowrap;
}

.state-msg.error {
  color: #C6402E;
}

.story-viewer {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.744);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.story-viewer-content {
  width: 100%;
  max-width: 440px;
  max-height: 100vh;
  border-radius: 12px;
  overflow: visible;
  background: #000;
  display: flex;
  flex-direction: column;
  position: relative;
}

.progress-row {
  display: flex;
  gap: 4px;
  padding: 8px 10px 0;
}

.progress-bar {
  flex: 1;
  height: 3px;
  background: rgba(255, 255, 255, 0.3);
  border-radius: 999px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  width: 0%;
  background: #fff;
  transition: width 0.1s linear;
}

.progress-fill.filled {
  width: 100%;
}

.progress-fill.active {
  width: 100%;
  animation: fillProgress 5s linear forwards;
}

@keyframes fillProgress {
  from { width: 0%; }
  to { width: 100%; }
}

.viewer-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  color: #fff;
  font-weight: 700;
  font-family: 'Nunito', sans-serif;
  background: rgba(0, 0, 0, 0.06);
  pointer-events: auto;       
}

.viewer-close {
  border: none;
  background: transparent;
  color: #fff;
  font-size: 18px;
  cursor: pointer;
}

.viewer-media-wrap {
  position: relative;
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden; 
  min-height: 90vh; 
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.viewer-bg-blur {
  position: absolute;
  inset: 0;
  background-size: cover;
  background-position: center;
  filter: blur(30px) brightness(0.55);
  transform: scale(1.15);
  z-index: 0;
}

 .viewer-media {
  max-width: 100%;
  max-height: 100%;
  width: auto;
  height: auto;
  object-fit: contain;
  display: block;
  position: relative;
  z-index: 1;
}

.nav-zone {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 30%;
  border: none;
  background: transparent;
  cursor: pointer;
  z-index: 2;
}

.nav-prev {
  left: 0;
  pointer-events: auto;
}

.nav-next {
  right: 0;
  pointer-events: auto;
}

.viewer-header-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.more-wrap {
  position: relative;
}

.viewer-more,
.viewer-close {
  border: none;
  background: transparent;
  color: #fff;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.viewer-more {
  width: 28px;
  height: 28px;
  padding: 0;
}

.viewer-more svg {
  width: 18px;
  height: 18px;
}

.viewer-close {
  font-size: 18px;
}

.more-menu {
  position: absolute;
  top: 34px;
  right: 0;
  background: #fff;
  /* border: 1px solid #E7E7E7; */
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .2);
  padding: 6px;
  min-width: 130px;
  z-index: 20;
}

.more-menu button {
  display: block;
  width: 100%;
  text-align: left;
  border: none;
  background: transparent;
  padding: 9px 12px;
  border-radius: 8px;
  font-size: 13px;
  color: #2B2B2B;
  cursor: pointer;
}

.more-menu button:hover {
  background: #F2F2F3;
}

.more-menu button.danger {
  color: #C6402E;
}

.viewer-user-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.viewer-avatar {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  overflow: hidden;
  background: rgba(255, 255, 255, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  border: 1.5px solid rgba(255, 255, 255, 0.6);
}

.viewer-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.viewer-avatar svg {
  width: 18px;
  height: 18px;
  stroke: #fff;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.viewer-user-text {
  display: flex;
  flex-direction: column;
  line-height: 1.3;
}

.viewer-username {
  font-size: 14px;
}

.viewer-time {
  font-size: 11px;
  font-weight: 400;
  color: rgba(255, 255, 255, 0.75);
}

* { box-sizing: border-box; }
 
.action-bar {
  position: absolute;
  /* right: 10px; */
  top: 50%;
  right: -55px; 
  transform: translateY(-50%);

  border-radius: 999px;
  padding: 10px 8px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
  z-index: 25;
  /* backdrop-filter: blur(4px); */
    gap: 6px;


}
 
.action-wrap {
  position: relative;
  display: flex;
}
 
.action-btn {
  
  border: none;
  background: transparent;
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  padding: 0;
  flex-shrink: 0;
  pointer-events: auto;
  flex-direction: column;
  height: auto;
  gap: 4px;
  
  
}

.action-btn img{
  width: 38px;
  height: 38px;
  background-color: #1A1A19;
  border-radius: 50%;
  /* border: 1px solid #ffff; */

}

.count-total {
  font-size: 14px;
  font-weight: 600;
  color: #fff;
  line-height: 1;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5); 
}

.action-btn .btn-com svg{
  width: 11px;
  height: 11px;
}

.action-btn .btn-com{
  width: 38px;
  height: 38px;
  border-radius: 50%px;

  display: flex;
  align-items: center;
  justify-content: center;

}
 
.action-btn svg {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background-color: #1A1A19;
}
 
.action-btn.active {
  color: #F2762E;
}
 
.reaction-emoji {
  width: 24px;
  height: 24px;
  display: inline-flex;
  color: #F2762E;
}
.reaction-emoji :deep(svg) {
  width: 100%;
  height: 100%;
}
 

.reaction-picker {
  position: absolute;
  right: calc(100% + 10px);
  top: 50%;
  transform: translateY(-50%);
  background: #fff;
  border-radius: 999px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.25);
  padding: 6px 8px;
  display: flex;
  align-items: center;
  gap: 4px;
  white-space: nowrap;
  animation: popIn 0.15s ease;
}
 
@keyframes popIn {
  from { opacity: 0; transform: translateY(-50%) scale(0.9); }
  to   { opacity: 1; transform: translateY(-50%) scale(1); }
}
 
.reaction-option {
  width: 32px;
  height: 32px;
  border: none;
  background: transparent;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #2B2B2B;
  transition: transform 0.12s ease;
}
 
.reaction-option :deep(svg) {
  width: 20px;
  height: 20px;
  display: block;
}
 
.reaction-option:hover {
  transform: scale(1.2);
  background: #F2F2F3;
}
 
.reaction-option.active {
  color: #F2762E;
}

.action-wrap-single {
  position: relative;
  display: flex;
}

.action-tooltip {
  position: absolute;
  right: calc(100% + 10px);  
  top: 50%;
  transform: translateY(-50%) translateX(6px);
  background: #1976D2;
  color: #ffff;
  font-size: 12px;
  font-weight: 500;
  padding: 5px 10px;
  border-radius: 6px;
  white-space: nowrap;
  opacity: 0;
  visibility: hidden;
  pointer-events: none;
  transition: opacity 0.18s ease, transform 0.18s ease;
  z-index: 30;
}


.action-tooltip::after {
  content: '';
  position: absolute;
  left: 100%;
  top: 50%;
  transform: translateY(-50%);
  border: 5px solid transparent;
  border-left-color: #1976D2;
}

.action-wrap-single:hover .action-tooltip {
  opacity: 1;
  visibility: visible;
  transform: translateY(-50%) translateX(0);
}

.emoji-fly-layer{
  position:absolute;
  inset:0;
  pointer-events:none;
  overflow:hidden;
  z-index:9999;
}


.btn-eye-story.is-viewed {
  background-color: #D0C3BD;
  border-color: #fff;
  box-shadow: 0 0 0 3px #D0C3BD;
}

.story-item:hover .btn-eye-story.is-viewed {
  box-shadow: 0 0 0 3px #d1d5db;
}


@media (max-width: 480px) {
  .action-bar {
    right: 6px;   
  }
}
</style>