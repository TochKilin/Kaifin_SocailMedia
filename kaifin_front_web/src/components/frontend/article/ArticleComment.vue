<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'

// Comment host api
const API_BASE = 'http://localhost:7070'

function authHeaders() {
  const token = localStorage.getItem('token')
  return token ? { Authorization: `Bearer ${token}` } : {}
}

async function parseJson(res) {
  const raw = await res.text()
  let data = null
  if (raw) {
    try {
      data = JSON.parse(raw)
    } catch {
      throw new Error(`Server returned non-JSON response (status ${res.status})`)
    }
  }
  if (!res.ok) {
    throw new Error(data?.message || `Request failed with status ${res.status}`)
  }
  return data
}

// GET comments
async function fetchComments(articleId, { page = 1, perPage = 20 } = {}) {
  const params = new URLSearchParams({ page, per_page: perPage })
  const res = await fetch(
    `${API_BASE}/api/v1/front/articles/${articleId}/comments?${params}`,
    { headers: authHeaders() }
  )
  const data = await parseJson(res)
  return data?.data
}

// POST comments
async function createComment(articleId, { text, parentCommentId = null, imageFiles = [], stickerIds = [] }) {
  const formData = new FormData()
  formData.append('article_id', Number(articleId))
  formData.append('text', text || '')

  if (parentCommentId !== null && parentCommentId !== undefined) {
    formData.append('parent_comment_id', parentCommentId)
  }
  imageFiles.forEach((file) => {
    formData.append('images', file)
  })

  stickerIds.forEach((id) => {
    formData.append('sticker_ids', id)
  })

  const res = await fetch(`${API_BASE}/api/v1/front/articles/${articleId}/comments`, {
    method: 'POST',
    headers: { ...authHeaders() },
    body: formData
  })
  const data = await parseJson(res)
  return data?.data
}

// PUT comment
async function updateComment(commentId, text) {
  const res = await fetch(`${API_BASE}/api/v1/front/comments/${commentId}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify({ text })
  })
  return parseJson(res)
}

// DELETE comment
async function deleteComment(commentId) {
  const res = await fetch(`${API_BASE}/api/v1/front/comments/${commentId}`, {
    method: 'DELETE',
    headers: authHeaders()
  })
  return parseJson(res)
}

const props = defineProps({
  articleId: {
    type: [Number, String],
    required: true
  }
})

const emit = defineEmits(['submit-comment', 'react', 'reply', 'comment-count-change'])

const comments = ref([])
const commentsCount = ref(0)
const isLoading = ref(false)
const loadError = ref(null)
const isPosting = ref(false)

const replyingTo = ref(null)

function timeAgo(dateStr) {
  if (!dateStr) return ''
  const seconds = Math.floor((Date.now() - new Date(dateStr).getTime()) / 1000)
  if (seconds < 60) return 'now'
  const mins = Math.floor(seconds / 60)
  if (mins < 60) return `${mins}m ago`
  const hours = Math.floor(mins / 60)
  if (hours < 24) return `${hours}h ago`
  const days = Math.floor(hours / 24)
  return `${days}d ago`
}

const authedImageCache = new Map()

function resolveImageUrl(path) {
  if (!path) return null
  if (path.startsWith('http://') || path.startsWith('https://')) return path
  if (path.startsWith('/uploads/')) return `${API_BASE}${path}`
  return `${API_BASE}/uploads/${path}`
}

// Fetch an authed comment-attachment image and turn it into a blob URL.
// Comment attachments live under /uploads/comments/* which IS behind auth
// middleware (unlike profile images) — plain <img src> gets a 401 there.
async function loadAuthedImage(path) {
  const url = resolveImageUrl(path)
  if (!url) return null
  if (authedImageCache.has(url)) return authedImageCache.get(url)
  try {
    const res = await fetch(url, { headers: authHeaders() })
    if (!res.ok) throw new Error(`image fetch failed (${res.status})`)
    const blob = await res.blob()
    const objectUrl = URL.createObjectURL(blob)
    authedImageCache.set(url, objectUrl)
    return objectUrl
  } catch (err) {
    console.error('Failed to load authed image:', url, err)
    return null
  }
}

// Map comment articles.
// avatarUrl uses resolveImageUrl directly (no auth needed, matches Feed.vue).
// imageUrls (comment attachments) uses loadAuthedImage because that route
// requires a Bearer token that a plain <img src> can never send.
async function mapComment(c) {
  const rawImageUrls = Array.isArray(c.image_urls) ? c.image_urls : []
  const imageUrls = (
    await Promise.all(rawImageUrls.map(loadAuthedImage))
  ).filter(Boolean)

  return {
    id: c.id,
    parentCommentId: c.parent_comment_id ?? null,
    username: c.user_name || 'Unknown',
    avatarUrl: resolveImageUrl(c.profile_images) ||
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&auto=format&fit=crop&q=80',
    text: c.text,
    time: timeAgo(c.created_at),
    likeCount: 0,
    commentCount: 0,
    userReaction: null,
    imageUrls
  }
}

async function loadComments() {
  isLoading.value = true
  loadError.value = null
  try {
    const data = await fetchComments(props.articleId)
    comments.value = await Promise.all((data?.comments || []).map(mapComment))
    commentsCount.value = data?.total ?? comments.value.length
    emit('comment-count-change', commentsCount.value)
  } catch (err) {
    console.error(err)
    loadError.value = err.message
  } finally {
    isLoading.value = false
  }
}

onMounted(loadComments)

const newCommentText = ref('')
const activeReactionCommentId = ref(null)
const selectedImageFiles = ref([])
const imagePreviewUrls = ref([])
const fileInputRef = ref(null)

const lightboxOpen = ref(false)
const lightboxImageUrl = ref('')

const showEmoji = ref(false)
const pickerTab = ref('stickers')
const stickerFilterTab = ref('all')

const emojiList = ['😀', '😂', '😍', '🥳', '😎', '🤔', '😢', '😡', '👍', '🙏', '🔥', '🎉', '❤️', '👏', '💡', '✅']

const mockStickers = ref([
  { id: 1, url: 'https://cdn-icons-png.flaticon.com/128/742/742751.png', file_name: 'heart-eyes' },
  { id: 2, url: 'https://cdn-icons-png.flaticon.com/128/2107/2107845.png', file_name: 'heart-hands' },
  { id: 3, url: 'https://cdn-icons-png.flaticon.com/128/616/616462.png', file_name: 'candy' },
  { id: 4, url: 'https://cdn-icons-png.flaticon.com/128/616/616554.png', file_name: 'bear' },
  { id: 5, url: 'https://cdn-icons-png.flaticon.com/128/3163/3163478.png', file_name: 'gift' },
  { id: 6, url: 'https://cdn-icons-png.flaticon.com/128/616/616408.png', file_name: 'sunglasses' },
  { id: 7, url: 'https://cdn-icons-png.flaticon.com/128/616/616655.png', file_name: 'party' },
  { id: 8, url: 'https://cdn-icons-png.flaticon.com/128/3067/3067381.png', file_name: 'heart' },
  { id: 9, url: 'https://cdn-icons-png.flaticon.com/128/3009/3009697.png', file_name: 'ghost' },
  { id: 10, url: 'https://cdn-icons-png.flaticon.com/128/616/616472.png', file_name: 'wink' },
  { id: 11, url: 'https://cdn-icons-png.flaticon.com/128/1791/1791330.png', file_name: 'trophy' },
  { id: 12, url: 'https://cdn-icons-png.flaticon.com/128/3163/3163559.png', file_name: 'balloon' }
])

const stickers = ref(mockStickers.value)
const selectedStickers = ref([])
const pickerRef = ref(null)
const filteredStickers = ref(mockStickers.value)

function selectStickerFilter(tab) {
  stickerFilterTab.value = tab
  if (tab === 'all') {
    filteredStickers.value = mockStickers.value
  } else {
    filteredStickers.value = []
  }
}

function addSticker(sticker) {
  selectedStickers.value.push(sticker)
  showEmoji.value = false
}

function removeSticker(index) {
  selectedStickers.value.splice(index, 1)
}

function addEmoji(e) {
  newCommentText.value += e
  showEmoji.value = false
}

const REACTIONS = [
  {
    key: 'like',
    label: 'Like',
    icon: `<svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="#4b5563" stroke-width="2"><path d="M7 11v9H4v-9h3Zm3 9h8a2 2 0 0 0 2-2l1.5-5a2 2 0 0 0-2-2.6H15l.7-4A2 2 0 0 0 13.8 4L10 10v10Z" fill="#f3f4f6"/></svg>`
  },
  {
    key: 'love',
    label: 'Love',
    icon: `<svg viewBox="0 0 24 24" width="20" height="20" fill="#ef4444"><path d="M12 21s-7.5-4.6-10-9.1C.4 8.6 2 5 5.6 5 8 5 10 6.4 12 9c2-2.6 4-4 6.4-4C22 5 23.6 8.6 22 11.9 19.5 16.4 12 21 12 21Z"/></svg>`
  },
  {
    key: 'haha',
    label: 'Haha',
    icon: `<svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="#f59e0b" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M9 9h.01M15 9h.01M8 13a4 4 0 0 0 8 0"/></svg>`
  },
  {
    key: 'wow',
    label: 'Wow',
    icon: `<svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="#8b5cf6" stroke-width="2"><circle cx="12" cy="12" r="9"/><circle cx="9" cy="10" r="1" fill="#8b5cf6"/><circle cx="15" cy="10" r="1" fill="#8b5cf6"/><ellipse cx="12" cy="16" rx="1.5" ry="2" fill="#8b5cf6"/></svg>`
  },
  {
    key: 'pray',
    label: 'Support',
    icon: `<svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="#10b981" stroke-width="2"><path d="M12 3l-1 6-3.5 4v7c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-4h3v4c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-7l-3.5-4-1-6Z" fill="#d1fae5"/></svg>`
  }
]

const triggerImageUpload = () => {
  if (imagePreviewUrls.value.length >= 2) return
  fileInputRef.value?.click()
}

const handleImageSelected = (event) => {
  const files = Array.from(event.target.files)
  if (!files.length) return

  for (const file of files) {
    if (imagePreviewUrls.value.length >= 2) break
    selectedImageFiles.value.push(file)
    imagePreviewUrls.value.push(URL.createObjectURL(file))
  }

  if (fileInputRef.value) {
    fileInputRef.value.value = ''
  }
}

const removeSelectedImage = (index) => {
  selectedImageFiles.value.splice(index, 1)
  imagePreviewUrls.value.splice(index, 1)
  if (fileInputRef.value) {
    fileInputRef.value.value = ''
  }
}

const openLightbox = (url) => {
  lightboxImageUrl.value = url
  lightboxOpen.value = true
}

const closeLightbox = () => {
  lightboxOpen.value = false
  lightboxImageUrl.value = ''
}

function startReply(comment) {
  replyingTo.value = { id: comment.id, username: comment.username }
  emit('reply', comment)
}

function cancelReply() {
  replyingTo.value = null
}

const handleSend = async () => {
  const text = newCommentText.value.trim()
  const hasImages = selectedImageFiles.value.length > 0
  const hasStickers = selectedStickers.value.length > 0
  if (!text && !hasImages && !hasStickers) return
  isPosting.value = true
  try {
    const stickerIds = selectedStickers.value.map(s => s.id)
    const created = await createComment(props.articleId, {
      text,
      parentCommentId: replyingTo.value?.id ?? null,
      imageFiles: selectedImageFiles.value,
      stickerIds: stickerIds
    })

    comments.value.unshift(await mapComment(created))
    commentsCount.value += 1
    emit('submit-comment', created)
    emit('comment-count-change', commentsCount.value)
    newCommentText.value = ''
    selectedImageFiles.value = []
    imagePreviewUrls.value = []
    selectedStickers.value = []
    replyingTo.value = null
  } catch (err) {
    console.error(err)
    loadError.value = err.message
  } finally {
    isPosting.value = false
  }
}

const selectReaction = (comment, reaction) => {
  comment.userReaction = reaction.key
  activeReactionCommentId.value = null
  emit('react', { commentId: comment.id, reaction })
}

onBeforeUnmount(() => {
  for (const url of authedImageCache.values()) {
    URL.revokeObjectURL(url)
  }
  authedImageCache.clear()
})
</script>

<template>
  <div class="article-comments-container">
    <div class="comment-header-title">
      <div class="header-title-left">
        <span class="header-text">Comments</span>
        <span class="comment-count-badge">{{ commentsCount }}</span>
      </div>
    </div>

    <div class="reply-banner" v-if="replyingTo">
      <span>Replying to <strong>{{ replyingTo.username }}</strong></span>
      <button @click="cancelReply">✕</button>
    </div>

    <div class="comment-input-box">
      <img
        src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&auto=format&fit=crop&q=80"
        alt="Me"
        class="user-avatar-img current-user-img"
      />
      <div class="input-content">
        <textarea
          v-model="newCommentText"
          placeholder="Write your comment here..."
          maxlength="1000"
          rows="2"
          :disabled="isPosting"
        ></textarea>

        <input
          type="file"
          ref="fileInputRef"
          accept="image/*"
          multiple
          style="display: none"
          @change="handleImageSelected"
        />

        <div class="selected-stickers" v-if="selectedStickers.length">
          <div class="selected-sticker" v-for="(s, index) in selectedStickers" :key="index">
            <img :src="s.url" :alt="s.file_name" />
            <button @click="removeSticker(index)">✕</button>
          </div>
        </div>

        <div v-if="imagePreviewUrls.length > 0" class="image-preview-list">
          <div v-for="(url, index) in imagePreviewUrls" :key="index" class="image-preview-container">
            <div class="preview-wrapper" @click="openLightbox(url)">
              <img :src="url" alt="Upload preview" class="preview-thumbnail" />
              <div class="preview-overlay">
                <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="#ffffff" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/><line x1="11" y1="8" x2="11" y2="14"/><line x1="8" y1="11" x2="14" y2="11"/></svg>
              </div>
            </div>
            <button class="remove-preview-btn" @click.stop="removeSelectedImage(index)" title="Remove image">
              <svg viewBox="0 0 24 24" width="14" height="14" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
            </button>
          </div>
        </div>

        <div class="input-footer">
          <div class="input-tools">
            <div class="chip-wrap" ref="pickerRef">
              <button
                class="tool-icon-btn"
                title="Emoji"
                :class="{ 'active-tool': showEmoji }"
                @click="showEmoji = !showEmoji"
              >
                <svg viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M9 10h.01M15 10h.01M9.5 15a3.5 3.5 0 0 0 5 0"/></svg>
              </button>

              <div class="picker-overlay" v-if="showEmoji" @click.self="showEmoji = false">
                <div class="picker-panel" @mousedown.stop @click.stop>
                  <div class="picker-tabs">
                    <button type="button" class="picker-tab" :class="{ active: pickerTab === 'stickers' }" @click.stop="pickerTab = 'stickers'">
                      <svg viewBox="0 0 64 64" xmlns="http://www.w3.org/2000/svg" width="64" height="64">
                        <path d="M12 8 H44 A8 8 0 0 1 52 16 V44 L38 58 H12 A8 8 0 0 1 4 50 V16 A8 8 0 0 1 12 8 Z"
                              fill="#FAC775" stroke="#BA7517" stroke-width="1.5"/>
                        <path d="M52 44 H42 A4 4 0 0 0 38 48 V58 Z"
                              fill="#EF9F27" stroke="#BA7517" stroke-width="1.5" stroke-linejoin="round"/>
                        <circle cx="20" cy="26" r="3" fill="#412402"/>
                        <circle cx="36" cy="26" r="3" fill="#412402"/>
                        <path d="M18 36 Q28 46 38 36" stroke="#412402" stroke-width="3" fill="none" stroke-linecap="round"/>
                      </svg>
                    </button>

                    <button type="button" class="picker-tab" :class="{ active: pickerTab === 'emojis' }" @click.stop="pickerTab = 'emojis'">
                      <svg viewBox="0 0 64 64" xmlns="http://www.w3.org/2000/svg" width="64" height="64">
                        <circle cx="32" cy="32" r="28" fill="#FAC775" stroke="#BA7517" stroke-width="2"/>
                        <circle cx="22" cy="26" r="3.5" fill="#412402"/>
                        <circle cx="42" cy="26" r="3.5" fill="#412402"/>
                        <path d="M18 38 Q32 52 46 38" stroke="#412402" stroke-width="3.5" fill="none" stroke-linecap="round"/>
                      </svg>
                    </button>
                  </div>

                  <div class="picker-body">
                    <template v-if="pickerTab === 'stickers'">
                      <div class="sticker-grid" v-if="filteredStickers.length">
                        <button type="button" class="sticker-item" v-for="s in filteredStickers" :key="s.id" @click.stop="addSticker(s)">
                          <img :src="s.url" :alt="s.file_name" />
                        </button>
                      </div>
                      <div class="picker-empty" v-else>
                        No sticker Sticker...
                      </div>
                    </template>

                    <template v-else-if="pickerTab === 'emojis'">
                      <div class="emoji-grid">
                        <button type="button" v-for="e in emojiList" :key="e" @click="addEmoji(e)">{{ e }}</button>
                      </div>
                    </template>
                  </div>

                  <div class="picker-footer" v-if="pickerTab === 'stickers'">
                    <div class="sticker-filter-row">
                      <button type="button" class="filter-chip" :class="{ active: stickerFilterTab === 'all' }" @click.stop="selectStickerFilter('all')">
                        <svg class="filter-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <rect x="3" y="3" width="7" height="7"/>
                          <rect x="14" y="3" width="7" height="7"/>
                          <rect x="14" y="14" width="7" height="7"/>
                          <rect x="3" y="14" width="7" height="7"/>
                        </svg>
                        <span>All</span>
                      </button>
                      <button type="button" class="filter-chip" :class="{ active: stickerFilterTab === 'animated' }" @click.stop="selectStickerFilter('animated')">
                        <svg class="filter-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <polygon points="5 3 19 12 5 21 5 3"/>
                        </svg>
                        Animated
                      </button>
                      <button type="button" class="filter-chip" :class="{ active: stickerFilterTab === 'mine' }" @click.stop="selectStickerFilter('mine')">
                        <svg class="filter-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/>
                        </svg>
                        <span>My stickers</span>
                      </button>
                      <button type="button" class="filter-chip filter-chip-create" :class="{ active: stickerFilterTab === 'create' }" @click.stop="selectStickerFilter('create')">
                        <svg class="filter-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <line x1="12" y1="5" x2="12" y2="19"/>
                          <line x1="5" y1="12" x2="19" y2="12"/>
                        </svg>
                        <span>Create Sticker</span>
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <button
              class="tool-icon-btn"
              title="Image"
              @click="triggerImageUpload"
              :class="{ 'disabled-tool': imagePreviewUrls.length >= 2 }"
            >
              <svg viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2"/><circle cx="8.5" cy="8.5" r="1.5"/><path d="M21 15l-5-5L5 21"/></svg>
            </button>
          </div>
          <div class="publish-group">
            <span class="char-limit" :class="{ 'text-danger': newCommentText.length > 900 }">
              {{ newCommentText.length }}/1000
            </span>
           <button
            class="publish-btn"
            :disabled="isPosting || (!newCommentText.trim() && selectedImageFiles.length === 0 && selectedStickers.length === 0)"
            @click="handleSend"
            >
              {{ isPosting ? 'Posting...' : 'Post' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <div v-if="isLoading" class="comment-loading">Loading comments…</div>
    <div v-else-if="loadError" class="comment-error">
      {{ loadError }}
      <button @click="loadComments">Retry</button>
    </div>

    <div v-else class="comments-stream">
      <div v-for="item in comments" :key="item.id" class="comment-item">
        <div class="comment-row">
          <img :src="item.avatarUrl" :alt="item.username" class="user-avatar-img" />
          <div class="comment-bubble">
            <div class="comment-bubble-header">
              <span class="author-name">{{ item.username }}</span>
            </div>
            <p class="comment-message">{{ item.text }}</p>

            <div v-if="item.imageUrls && item.imageUrls.length > 0" class="comment-images-grid" :class="'grid-count-' + item.imageUrls.length">
              <div
                v-for="(imgUrl, imgIndex) in item.imageUrls"
                :key="imgIndex"
                class="comment-image-preview-wrapper"
                @click="openLightbox(imgUrl)"
              >
                <img :src="imgUrl" alt="Attached image" class="comment-attached-image" />
                <div class="comment-image-overlay">
                  <svg viewBox="0 0 24 24" width="22" height="22" fill="none" stroke="#ffffff" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/><line x1="11" y1="8" x2="11" y2="14"/><line x1="8" y1="11" x2="14" y2="11"/></svg>
                </div>
              </div>
            </div>

            <div class="comment-actions-strip">
              <span class="action-chip time-chip">
                <svg viewBox="0 0 24 24" width="13" height="13" fill="none" stroke="currentColor" stroke-width="2" class="time-icon"><circle cx="12" cy="12" r="9"/><polyline points="12 6 12 12 16 14"/></svg>
                <span>{{ item.time }}</span>
              </span>

              <div
                class="reaction-menu-anchor"
                @mouseenter="activeReactionCommentId = item.id"
                @mouseleave="activeReactionCommentId = null"
              >
                <div v-if="activeReactionCommentId === item.id" class="reactions-popover">
                  <div class="reactions-box">
                    <button
                      v-for="reaction in REACTIONS"
                      :key="reaction.key"
                      class="popover-emoji"
                      @click="selectReaction(item, reaction)"
                    >
                      <span class="emoji-svg-wrapper" v-html="reaction.icon"></span>
                      <span class="emoji-tooltip">{{ reaction.label }}</span>
                    </button>
                  </div>
                </div>

                <button class="action-chip" :class="{ 'active-action': item.userReaction }">
                  <span class="action-icon" v-html="item.userReaction ? REACTIONS.find(r => r.key === item.userReaction)?.icon : REACTIONS[0].icon"></span>
                  <span>{{ item.userReaction ? '' : 'Like' }}</span>
                  <span v-if="item.likeCount > 0" class="chip-count">{{ item.likeCount }}</span>
                </button>
              </div>

              <button class="action-chip" @click="startReply(item)">
                <svg viewBox="0 0 24 24" width="14" height="14" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/></svg>
                <span>Reply</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <div v-if="!isLoading && comments.length === 0" class="comment-empty">
        No comments yet. Be the first to comment!
      </div>
    </div>

    <div v-if="lightboxOpen" class="lightbox-backdrop" @click="closeLightbox">
      <div class="lightbox-content" @click.stop>
        <button class="lightbox-close-btn" @click="closeLightbox" title="Close">
          <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
        </button>
        <img :src="lightboxImageUrl" alt="Enlarged preview" class="lightbox-image" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.article-comments-container {
  background-color: #ffffff;
  padding: 20px;
  color: #1f2937;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  max-width: 100%;
  width: 100%;
}

.comment-header-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
  padding-bottom: 12px;
  border-bottom: 1px solid #f3f4f6;
}

.header-title-left {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 16px;
  font-weight: 700;
  color: #111827;
}

.comment-count-badge {
  background-color: #f3f4f6;
  color: #4b5563;
  font-size: 12px;
  padding: 2px 8px;
  border-radius: 12px;
  font-weight: 600;
}

.comment-input-box {
  display: flex;
  gap: 12px;
  border: 1.5px solid #e5e7eb;
  border-radius: 16px;
  padding: 14px;
  margin-bottom: 24px;
  background-color: #fafafa;
  transition: all 0.2s ease;
}

.comment-input-box:focus-within {
  border-color: #1B75D2;
  background-color: #ffffff;
}

.user-avatar-img {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.input-content {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.input-content textarea {
  width: 100%;
  background: transparent;
  border: none;
  outline: none;
  font-size: 14px;
  color: #1f2937;
  resize: none;
  font-family: inherit;
}

.input-content textarea::placeholder {
  color: #9ca3af;
}

.selected-stickers {
  display: flex;
  gap: 8px;
  margin: 4px 0;
  flex-wrap: wrap;
}

.selected-sticker {
  position: relative;
  width: 64px;
  height: 64px;
  background: #f5f5f5;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.selected-sticker img {
  width: 50px;
  height: 50px;
  object-fit: contain;
}

.selected-sticker button {
  position: absolute;
  top: -5px;
  right: -5px;
  width: 20px;
  height: 20px;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  background: #111827;
  color: #fff;
  font-size: 10px;
}

.image-preview-list {
  display: flex;
  gap: 8px;
  margin-top: 6px;
  margin-bottom: 6px;
  flex-wrap: wrap;
}

.image-preview-container {
  position: relative;
  display: inline-block;
  width: fit-content;
}

.preview-wrapper {
  position: relative;
  width: 90px;
  height: 90px;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  border: 1.5px solid #1B75D2;
  background-color: #f3f4f6;
}

.preview-thumbnail {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.2s ease;
}

.preview-wrapper:hover .preview-thumbnail {
  transform: scale(1.05);
}

.preview-overlay {
  position: absolute;
  inset: 0;
  background-color: rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.preview-wrapper:hover .preview-overlay {
  opacity: 1;
}

.remove-preview-btn {
  position: absolute;
  top: -6px;
  right: -6px;
  background-color: #111827;
  color: #ffffff;
  border: 2px solid #ffffff;
  width: 22px;
  height: 22px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 2px 6px rgba(0,0,0,0.2);
  transition: background-color 0.15s;
}

.remove-preview-btn:hover {
  background-color: #ef4444;
}

.input-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-top: 1px solid #f0f2f5;
  padding-top: 10px;
}

.input-tools {
  display: flex;
  gap: 6px;
  align-items: center;
}

.chip-wrap {
  position: relative;
  display: inline-flex;
  align-items: center;
}

.tool-icon-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 6px;
  border-radius: 8px;
  color: #6b7280;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.15s;
}

.tool-icon-btn:hover {
  background-color: #e5e7eb;
  color: #111827;
}

.tool-icon-btn.active-tool {
  background-color: #dbeafe;
  color: #1B75D2;
}

.disabled-tool {
  opacity: 0.4;
  cursor: not-allowed;
}
.disabled-tool:hover {
  background-color: transparent;
  color: #6b7280;
}

.publish-group {
  display: flex;
  align-items: center;
  gap: 12px;
}

.char-limit {
  font-size: 11px;
  color: #9ca3af;
}

.publish-btn {
  background-color: #1B75D2;
  color: #ffffff;
  border: none;
  padding: 6px 16px;
  border-radius: 32px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.publish-btn:hover:not(:disabled) {
  background-color: #374151;
  transform: translateY(-1px);
}

.publish-btn:disabled {
  background-color: #e5e7eb;
  color: #9ca3af;
  cursor: not-allowed;
}

.picker-overlay {
  position: fixed;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  z-index: 1000;
}

.picker-panel {
  position: relative;
  background: #fff;
  border: 1.5px solid #E7E7E7;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.075);
  width: 92%;
  max-width: 380px;
  max-height: 80vh;
  margin-right: 10px;
  display: flex;
  flex-direction: column;
  z-index: 1001;
  overflow: hidden;
  pointer-events: auto;
  bottom: 55px;
}

.picker-tabs {
  display: flex;
  background: #F2F2F3;
  border-top-left-radius: 12px;
  gap: 6px;
  border-bottom: 1px solid #E7E7E7;
  border-top-right-radius: 12px;
  pointer-events: auto;
}

.picker-tab {
  flex: 1;
  border: none;
  background: transparent;
  padding: 3px;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 13.5px;
  color: #6A6A6E;
  cursor: pointer;
  border-radius: 10px;
  transition: all 0.2s;
  pointer-events: auto;
}

.picker-tab svg {
  width: 30px;
  height: 30px;
}

.picker-tab.active {
  background: transparent;
  color: #1B75D2;
  border-bottom: 2px solid #1B75D2;
  border-bottom-left-radius: 0;
  border-bottom-right-radius: 0;
  border-top-right-radius: 0;
}

.picker-body {
  flex: 1;
  padding: 12px;
  overflow-y: auto;
  max-height: 400px;
}

.picker-footer {
  min-height: 50px;
  border-top: 1px solid #E7E7E7;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
  padding: 6px;
}

.sticker-grid {
  display: grid;
  grid-template-columns: repeat(8, 1fr);
  gap: 8px;
}

.sticker-item {
  border: none;
  background: #EFF6FB;
  border-radius: 12px;
  padding: 6px;
  cursor: pointer;
  aspect-ratio: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.15s;
}

.sticker-item:hover {
  transform: scale(1.08);
  background: #DCEBF7;
}

.sticker-item img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  pointer-events: none;
}

.emoji-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  gap: 6px;
}

.emoji-grid button {
  border: none;
  background: transparent;
  font-size: 24px;
  padding: 6px;
  cursor: pointer;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.emoji-grid button:hover {
  background: #F2F2F3;
}

.picker-empty {
  text-align: center;
  color: #9A9A9E;
  font-size: 13px;
  padding: 20px 0;
  font-family: 'Inter', sans-serif;
}

.sticker-filter-row {
  display: flex;
  gap: 6px;
  overflow-x: auto;
  scrollbar-width: none;
  justify-content: center;
  width: 100%;
}

.sticker-filter-row::-webkit-scrollbar {
  display: none;
}

.filter-chip {
  flex-shrink: 0;
  border: none;
  background: transparent;
  color: #1B75D2;
  font-family: 'Nunito', sans-serif;
  font-weight: 400;
  font-size: 12px;
  padding: 4px 4px;
  border-radius: 0;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.15s ease;
  pointer-events: auto;
  margin-bottom: -1.5px;
  border-bottom: 2px solid transparent;
  gap: 5px;
  display: inline-flex;
  align-items: center;
}

.filter-chip.active {
  background: transparent;
  color: #1B75D2;
  border-bottom: 4px solid #1B75D2;
}

.filter-chip-create {
  border-color: transparent;
  background: transparent;
  color: #8A6D1D;
}

.filter-chip-create.active {
  color: #B8901E;
  border-bottom: 2px solid #F2C744;
}

.filter-icon {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

.comments-stream {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.comment-item {
  padding: 16px;
  border-radius: 16px;
  background-color: transparent;
  border: 1px solid #e5e7eb;
  transition: all 0.2s;
}

.comment-item:hover {
  border-color: #d1d5db;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
}

.comment-row {
  display: flex;
  gap: 12px;
}

.comment-bubble {
  flex-grow: 1;
}

.comment-bubble-header {
  display: flex;
  align-items: baseline;
  gap: 8px;
  margin-bottom: 4px;
}

.author-name {
  font-size: 14px;
  font-weight: 600;
  color: #111827;
}

.comment-message {
  font-size: 14px;
  color: #4b5563;
  line-height: 1.5;
  margin: 0;
}

.comment-images-grid {
  display: flex;
  gap: 8px;
  margin-top: 10px;
  max-width: 320px;
}

.comment-images-grid.grid-count-1 .comment-image-preview-wrapper {
  width: 110px;
  height: 80px;
}

.comment-images-grid.grid-count-2 .comment-image-preview-wrapper {
  width: 110px;
  height: 80px;
}

.comment-image-preview-wrapper {
  position: relative;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  border: 1.5px solid #1B75D2;
  background-color: #f3f4f6;
}

.comment-attached-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.25s ease;
}

.comment-image-preview-wrapper:hover .comment-attached-image {
  transform: scale(1.03);
}

.comment-image-overlay {
  position: absolute;
  inset: 0;
  background-color: rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.comment-image-preview-wrapper:hover .comment-image-overlay {
  opacity: 1;
}

.comment-actions-strip {
  display: flex;
  gap: 12px;
  align-items: center;
  margin-top: 10px;
}

.action-chip {
  border: none;
  background-color: #f3f4f6;
  color: #4b5563;
  font-size: 12px;
  font-weight: 500;
  padding: 4px 10px;
  border-radius: 20px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: all 0.15s ease;
}

.time-chip {
  cursor: default;
}

.time-chip:hover {
  background-color: #f3f4f6;
  color: #4b5563;
}

.action-chip:hover:not(.time-chip) {
  background-color: #e5e7eb;
  color: #111827;
}

.time-icon {
  flex-shrink: 0;
}

.action-icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.chip-count {
  background-color: rgba(0, 0, 0, 0.05);
  border: none;
  padding: 0 6px;
  border-radius: 10px;
  font-size: 11px;
  font-weight: 600;
}

.reaction-menu-anchor {
  position: relative;
  display: inline-block;
}

.reactions-popover {
  position: absolute;
  bottom: calc(100% + 8px);
  left: 0;
  z-index: 50;
  animation: popIn 0.2s cubic-bezier(0.16, 1, 0.3, 1);
}

.reactions-box {
  display: flex;
  gap: 6px;
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 30px;
  padding: 6px 10px;
}

.popover-emoji {
  position: relative;
  background: none;
  border: none;
  cursor: pointer;
  padding: 6px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.popover-emoji:hover {
  transform: scale(1.3) translateY(-4px);
  background-color: #f3f4f6;
}

.emoji-svg-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
}

.emoji-tooltip {
  position: absolute;
  bottom: 135%;
  left: 50%;
  transform: translateX(-50%);
  background-color: #111827;
  color: #ffffff;
  padding: 2px 6px;
  font-size: 10px;
  border-radius: 4px;
  white-space: nowrap;
  opacity: 0;
  visibility: hidden;
  transition: all 0.15s ease;
  pointer-events: none;
}

.popover-emoji:hover .emoji-tooltip {
  opacity: 1;
  visibility: visible;
}

.lightbox-backdrop {
  position: fixed;
  inset: 0;
  background-color: rgba(0, 0, 0, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 20px;
  animation: fadeIn 0.2s ease;
}

.lightbox-content {
  position: relative;
  max-width: 90vw;
  max-height: 90vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.lightbox-image {
  max-width: 100%;
  max-height: 85vh;
  object-fit: contain;
  border-radius: 8px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
}

.lightbox-close-btn {
  position: absolute;
  top: -45px;
  right: 0;
  background-color: rgba(255, 255, 255, 0.2);
  color: #ffffff;
  border: none;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.15s;
}

.reply-banner {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 6px 12px;
  background: #f3f4f6;
  border-radius: 6px;
  font-size: 13px;
  margin: 8px 0;
}
.reply-banner button {
  background: none;
  border: none;
  cursor: pointer;
  color: #6b7280;
}
.comment-loading,
.comment-error,
.comment-empty {
  padding: 16px 0;
  text-align: center;
  color: #6b7280;
  font-size: 14px;
}
.comment-error button {
  margin-left: 8px;
  color: #2563eb;
  background: none;
  border: none;
  cursor: pointer;
  text-decoration: underline;
}

.lightbox-close-btn:hover {
  background-color: rgba(255, 255, 255, 0.4);
}

@keyframes popIn {
  from { opacity: 0; transform: scale(0.9) translateY(4px); }
  to { opacity: 1; transform: scale(1) translateY(0); }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>