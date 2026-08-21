<template>
  <div class="media-card">
    <div class="media-header">
      <h3 class="media-title">Media</h3>
      <button class="view-more-link" @click="viewMoreMedia">
        View More
      </button>
    </div>

    <p v-if="isLoading" class="media-state-msg">Loading...</p>
    <p v-else-if="error" class="media-state-msg error">{{ error }}</p>
    <p v-else-if="!mediaItems.length" class="media-state-msg">No media yet</p>

    <div v-else class="media-grid">
      <div v-for="(item, index) in mediaItems" :key="index" class="media-item" @click="openMedia(item)">
        <img v-if="item.type === 'image'" :src="item.url" :alt="'Media ' + index" />
        <video v-else-if="item.type === 'video'" :src="item.url"></video>
      </div>
    </div>
  </div>


  <div class="story-card">
    <div class="story-header">
      <h3 class="story-title">Stories</h3>
      <button class="view-more-link" @click="viewMoreStories">
        View More
      </button>
    </div>

    <p v-if="isStoryLoading" class="story-state-msg">Loading...</p>
    <p v-else-if="storyError" class="story-state-msg error">{{ storyError }}</p>
    <p v-else-if="!storyItems.length" class="story-state-msg">No stories yet</p>

    <div v-else class="story-grid">
      <div v-for="(story, index) in storyItems" :key="index" class="story-item" @click="openStory(story)">
        <div class="story-circle-wrapper">
          <img v-if="story.thumbnail || story.url" :src="resolveMediaUrl(story.thumbnail || story.url)" :alt="'Story ' + index" />
          <div v-else class="story-placeholder">Story</div>
          <div class="story-icon-overlay">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
              <circle cx="12" cy="12" r="3"></circle>
            </svg>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const props = defineProps({
  userId: { type: [String, Number], required: true },
  limit: { type: Number, default: 8 },
})

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'
const router = useRouter()

const mediaItems = ref([])
const isLoading = ref(false)
const error = ref(null)

const storyItems = ref([])
const isStoryLoading = ref(false)
const storyError = ref(null)

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

async function fetchUserMedia() {
  console.log("👉 fetchUserMedia running with userId:", props.userId)

  if (!props.userId) {
    console.warn("⚠️ userId is missing or null!");
    return
  }

  isLoading.value = true
  error.value = null

  try {
    const res = await fetch(
      `${BASE_URL}/api/v1/front/posts/show?page=1&perpage=50&user_id=${props.userId}`,
      { headers: { ...authHeaders() } }
    )
    
    console.log("📡 API Status:", res.status)
    if (!res.ok) throw new Error(`API ${res.status}`)

    const json = await res.json()
    console.log("📦 API Response Data:", json)

    const payload = json?.data ?? json
    const rawList = payload?.posts ?? payload?.Posts ?? []

    const collectedMedia = []
    for (const p of rawList) {
      const imageField = p.images || p.image || p.image_url || p.file || p.media
      
      if (imageField) {
        if (typeof imageField === 'string') {
          const urls = imageField.split(',').map((u) => u.trim()).filter(Boolean)
          for (const u of urls) {
            collectedMedia.push({ type: 'image', url: resolveMediaUrl(u) })
          }
        } else if (Array.isArray(imageField)) {
          for (const u of imageField) {
            collectedMedia.push({ type: 'image', url: resolveMediaUrl(u) })
          }
        }
      }

      const videoField = p.video_path || p.video || p.video_url
      if (videoField) {
        collectedMedia.push({ type: 'video', url: resolveMediaUrl(videoField) })
      }

      if (collectedMedia.length >= props.limit) break
    }

    mediaItems.value = collectedMedia.slice(0, props.limit)
  } catch (e) {
    console.error('Failed to load user media', e)
    error.value = 'Failed to load media'
  } finally {
    isLoading.value = false
  }
}

async function fetchUserStories() {
  if (!props.userId) return

  isStoryLoading.value = true
  storyError.value = null

  try {
    const res = await fetch(
      `${BASE_URL}/api/v1/front/stories/show?user_id=${props.userId}&perpage=50`,
      { headers: { ...authHeaders() } }
    )
    
    if (!res.ok) throw new Error(`API ${res.status}`)

    const json = await res.json()
    const payload = json?.data ?? json
    const rawList = payload?.stories ?? payload?.Stories ?? []

    storyItems.value = rawList.slice(0, 7).map(s => ({
      id: s.id,
      url: s.media_url || s.image || s.url
    }))
  } catch (e) {
    console.error('Failed to load user stories', e)
    storyError.value = 'Failed to load stories'
  } finally {
    isStoryLoading.value = false
  }
}

function addNewPostMedia(post) {
  const rawPost = post?.data?.post || post?.data || post
  if (!rawPost) return

  const newMedia = []

  if (rawPost.images) {
    const urls = typeof rawPost.images === 'string' 
      ? rawPost.images.split(',').map(u => u.trim()).filter(Boolean)
      : (Array.isArray(rawPost.images) ? rawPost.images : [])
    
    for (const u of urls) {
      newMedia.push({ type: 'image', url: resolveMediaUrl(u) })
    }
  }

  if (rawPost.video_path) {
    newMedia.push({ type: 'video', url: resolveMediaUrl(rawPost.video_path) })
  }

  if (newMedia.length === 0) return

  mediaItems.value.unshift(...newMedia)

  if (mediaItems.value.length > props.limit) {
    mediaItems.value = mediaItems.value.slice(0, props.limit)
  }
}

defineExpose({
  addNewPostMedia
})

onMounted(() => {
  fetchUserMedia()
  fetchUserStories()
})

watch(() => props.userId, () => {
  fetchUserMedia()
  fetchUserStories()
})

function openMedia(item) {
  console.log('Clicked media:', item)
}

function viewMoreMedia() {
  router.push(`/uploads/${props.userId}`)
}

function openStory(story) {
  console.log('Clicked story:', story)
}

function viewMoreStories() {
  router.push(`/stories/${props.userId}`)
}
</script>

<style scoped>
.media-card {
  background-color: #ffffff;
  border: 1px solid #edf2f7;
  border-radius: 12px;
  padding: 10px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.03);
  width: 100%;
  box-sizing: border-box;
  margin-top: 8px;
}

.media-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.media-title {
  font-size: 14px;
  font-weight: 700;
  color: #1e293b;
  margin: 0;
}

.view-more-link {
  background: none;
  border: none;
  color: #3b82f6;
  font-size: 11px;
  font-weight: 500;
  cursor: pointer;
  padding: 0;
  transition: color 0.2s ease;
}

.view-more-link:hover {
  color: #1d4ed8;
  text-decoration: underline;
}

.media-state-msg {
  font-size: 12.5px;
  color: #94a3b8;
  text-align: center;
  padding: 16px 0;
  margin: 0;
}

.media-state-msg.error {
  color: #dc2626;
}

.media-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 5px;
}

.media-item {
  aspect-ratio: 1 / 1;
  border-radius: 6px;
  overflow: hidden;
  cursor: pointer;
  background-color: #f1f5f9;
  transition: transform 0.2s ease;
}

.media-item:hover {
  transform: scale(1.02);
}

.media-item img,
.media-item video {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.story-card {
  background-color: #ffffff;
  border: 1px solid #edf2f7;
  border-radius: 12px;
  padding: 10px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.03);
  width: 100%;
  box-sizing: border-box;
  margin-top: 12px;
}

.story-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.story-title {
  font-size: 14px;
  font-weight: 700;
  color: #1e293b;
  margin: 0;
}

.story-state-msg {
  font-size: 12.5px;
  color: #94a3b8;
  text-align: center;
  padding: 16px 0;
  margin: 0;
}

.story-state-msg.error {
  color: #dc2626;
}

.story-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 8px;
}

.story-item {
  aspect-ratio: 1 / 1;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.story-circle-wrapper {
  position: relative;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  overflow: hidden;
  border: 3px solid #1B75D2; 
  background-color: #f1f5f9;
  box-shadow: 0 2px 6px rgba(27, 117, 210, 0.2);
  transition: transform 0.2s ease;
}

.story-item:hover .story-circle-wrapper {
  transform: scale(1.04);
}

.story-circle-wrapper img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.story-icon-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.25); 
  display: flex;
  align-items: center;
  justify-content: center;
}

.story-icon-overlay svg {
  width: 28px;
  height: 28px;
  color: #ffffff;
  filter: drop-shadow(0 1px 2px rgba(0,0,0,0.5));
}

.story-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  color: #64748b;
  border-radius: 50%;
}
</style>