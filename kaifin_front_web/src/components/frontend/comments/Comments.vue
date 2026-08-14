<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  postId: {
    type: [String, Number],
    required: true,
  },
})

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'
function authHeaders() {
  const token = localStorage.getItem('token') || ''
  return token ? { Authorization: `Bearer ${token}` } : {}
}

const currentUser = computed(() => {
  const token = localStorage.getItem('token')
  if (!token) return { initials: 'ME', color: '#1976D2', username: 'You' }
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    const name = payload.user_name || payload.username || 'You'
    return {
      id: payload.user_id || payload.id,
      username: name,
      initials: name.substring(0, 2).toUpperCase(),
      color: '#1976D2'
    }
  } catch {
    return { initials: 'ME', color: '#1976D2', username: 'You' }
  }
})

const comments = ref([])
const isLoading = ref(false)
const error = ref(null)
const emojis = ['😀', '😂', '😍', '👍', '🎉', '🔥', '😮', '😢']
const emojiOpenFor = ref(null)
const replyOpenFor = ref(null)
const draft = reactive({
  text: '',
  images: [], // [{ file, previewUrl }]
})

const lightbox = reactive({
  open: false,
  images: [],
  index: 0,
})

function openLightbox(images, startIndex = 0) {
  lightbox.images = images
  lightbox.index = startIndex
  lightbox.open = true
}
function closeLightbox() {
  lightbox.open = false
}
function lightboxNext() {
  lightbox.index = (lightbox.index + 1) % lightbox.images.length
}
function lightboxPrev() {
  lightbox.index = (lightbox.index - 1 + lightbox.images.length) % lightbox.images.length
}
function onLightboxKeydown(e) {
  if (!lightbox.open) return
  if (e.key === 'Escape') closeLightbox()
  else if (e.key === 'ArrowRight') lightboxNext()
  else if (e.key === 'ArrowLeft') lightboxPrev()
}

const fileInputDraft = ref(null)
const replyFileRefs = {}
const emit = defineEmits(['comment-count-changed'])
const totalCount = computed(() => {
  return comments.value.reduce((sum, comment) => {
    return sum + 1 + (comment.replies?.length || 0)
  }, 0)
})

async function fetchComments() {
  isLoading.value = true
  error.value = null
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/comments/show?post_id=${props.postId}`, {
      headers: authHeaders()
    })
    if (!res.ok) throw new Error(`Error: ${res.status}`)

    const json = await res.json()
    const payload = json?.data ?? json
    const rawComments = payload?.Comments ?? payload?.comments ?? []

    comments.value = buildTreeStructure(rawComments)
    emit('comment-count-changed', totalCount.value)
  } catch (err) {
    console.error('Failed to load comments:', err)
    error.value = 'មិនអាចទាញយកមតិយោបល់បានទេ'
  } finally {
    isLoading.value = false
  }
}

function buildTreeStructure(rawList) {
  const itemMap = {}
  const roots = []

  rawList.forEach(c => {
    itemMap[c.id] = {
      id: c.id,
      postId: c.post_id,
      userId: c.user_id,
      username: c.user_name || `User #${c.user_id}`,
      avatarUrl: resolveAvatarUrl(c.profile_images),
      images: (c.images || []).map(img => resolveImageUrl(img.image_path)),
      initials: (c.user_name || 'U').substring(0, 2).toUpperCase(),
      color: c.user_id === currentUser.value.id ? '#1976D2' : '#7C6FE8',
      text: c.content,
      time: formatDatetime(c.created_at),
      parentCommentId: c.parent_comment_id,
      likes: c.like_count || 0,
      liked: !!c.liked,
      replyDraft: '',
      replyImages: [],
      replies: []
    }
  })

  rawList.forEach(c => {
    const mapped = itemMap[c.id]
    if (c.parent_comment_id && itemMap[c.parent_comment_id]) {
      itemMap[c.parent_comment_id].replies.push(mapped)
    } else {
      roots.push(mapped)
    }
  })

  return roots
}

async function postComment() {
  if (!draft.text.trim() && !draft.images.length) return
  try {
    const formData = new FormData()
    formData.append('post_id', String(props.postId))
    formData.append('content', draft.text.trim())
    draft.images.forEach(img => formData.append('images', img.file))

    const res = await fetch(`${BASE_URL}/api/v1/front/comments/create`, {
      method: 'POST',
      headers: { ...authHeaders() }, // don't set Content-Type manually — browser sets the boundary
      body: formData
    })

    if (!res.ok) throw new Error('Failed to create comment')

    draft.text = ''
    draft.images.forEach(img => URL.revokeObjectURL(img.previewUrl))
    draft.images = []
    await fetchComments()
  } catch (err) {
    alert('មិនអាចផ្ញើមតិយោបល់បានទេ')
  }
}

async function postReply(comment) {
  const hasText = (comment.replyDraft || '').trim()
  const hasImages = (comment.replyImages || []).length
  if (!hasText && !hasImages) return
  try {
    const formData = new FormData()
    formData.append('post_id', String(props.postId))
    formData.append('parent_comment_id', String(comment.id))
    formData.append('content', (comment.replyDraft || '').trim())
    ;(comment.replyImages || []).forEach(img => formData.append('images', img.file))

    const res = await fetch(`${BASE_URL}/api/v1/front/comments/create`, {
      method: 'POST',
      headers: { ...authHeaders() },
      body: formData
    })

    if (!res.ok) throw new Error('Failed to reply')

    comment.replyDraft = ''
    ;(comment.replyImages || []).forEach(img => URL.revokeObjectURL(img.previewUrl))
    comment.replyImages = []
    replyOpenFor.value = null
    await fetchComments()
  } catch (err) {
    alert('មិនអាចឆ្លើយតបបានទេ')
  }
}

async function deleteComment(id) {
  if (!confirm('តើអ្នកពិតជាចង់លុបមតិយោបល់នេះមែនទេ?')) return

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/comments/delete/${id}`, {
      method: 'DELETE',
      headers: authHeaders()
    })

    if (!res.ok) {
      const errText = await res.text().catch(() => '')
      console.error('Delete failed:', errText)
      throw new Error('Delete failed')
    }

    await fetchComments()
  } catch (err) {
    alert('Can not delete (you can delete your comment only)')
  }
}

onMounted(() => {
  fetchComments()
  window.addEventListener('keydown', onLightboxKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', onLightboxKeydown)
})

function autoGrow(event) {
  const textarea = event.target
  textarea.style.height = 'auto'
  textarea.style.height = textarea.scrollHeight + 'px'
}

function setReplyFileRef(el, id) {
  if (el) replyFileRefs[id] = el
}

function triggerFile(key) {
  if (key === 'draft') fileInputDraft.value?.click()
}

function triggerReplyFile(id) {
  if (replyFileRefs[id]) replyFileRefs[id].click()
}

function onFile(event, kind, comment) {
  const files = Array.from(event.target.files || [])
  if (!files.length) return
  const items = files.map(file => ({ file, previewUrl: URL.createObjectURL(file) }))

  if (kind === 'draft') {
    draft.images.push(...items)
  } else if (kind === 'reply' && comment) {
    comment.replyImages = comment.replyImages || []
    comment.replyImages.push(...items)
  }
  event.target.value = ''
}

function removeDraftImage(i) {
  URL.revokeObjectURL(draft.images[i].previewUrl)
  draft.images.splice(i, 1)
}

function removeReplyImage(comment, i) {
  URL.revokeObjectURL(comment.replyImages[i].previewUrl)
  comment.replyImages.splice(i, 1)
}

function resolveImageUrl(path) {
  if (!path) return ''
  if (path.startsWith('http://') || path.startsWith('https://')) return path
  return `${BASE_URL}/uploads/${path}`
}

function toggleEmoji(key) {
  emojiOpenFor.value = emojiOpenFor.value === key ? null : key
}

function insertEmoji(emoji, kind, comment) {
  if (kind === 'draft') draft.text += emoji
  else if (kind === 'reply' && comment) comment.replyDraft = (comment.replyDraft || '') + emoji
  emojiOpenFor.value = null
}

function toggleReply(id) {
  replyOpenFor.value = replyOpenFor.value === id ? null : id
}

async function toggleLike(item) {
  const prevLiked = item.liked
  const prevCount = item.likes

  item.liked = !item.liked
  item.likes += item.liked ? 1 : -1

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/comments/${item.id}/like`, {
      method: 'POST',
      headers: { ...authHeaders() }
    })
    if (!res.ok) throw new Error('Failed to toggle like')
  } catch (err) {
    item.liked = prevLiked
    item.likes = prevCount
    alert('មិនអាចធ្វើប្រតិកម្មបានទេ')
  }
}

function formatDatetime(value) {
  if (!value) return ''
  const d = new Date(value)
  if (Number.isNaN(d.getTime())) return String(value)
  return d.toLocaleDateString('km-KH', { month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' })
}

function resolveAvatarUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}
</script>

<template>
  <div class="comment-section embedded">

    <svg width="0" height="0" style="position: absolute" aria-hidden="true">
      <defs>
        <symbol id="cs-icon-image" viewBox="0 0 24 24">
          <rect x="3" y="4" width="18" height="16" rx="2.5" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <circle cx="8.5" cy="9.5" r="1.6" fill="currentColor"/>
          <path d="M4 16.5l5-5 3.5 3.5L17 10l3 3.5V18a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 4 18v-1.5z" fill="currentColor"/>
        </symbol>

        <symbol id="cs-icon-at" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="9.25" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <circle cx="12" cy="12" r="3.6" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <path d="M15.6 10.6V13c0 1.4 1 2.3 2.1 2.3 1.6 0 2.4-1.7 2.4-3.6 0-1-.3-2-1-2.9" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
        </symbol>

        <symbol id="cs-icon-smile" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="9.25" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <circle cx="8.7" cy="10.2" r="1.1" fill="currentColor"/>
          <circle cx="15.3" cy="10.2" r="1.1" fill="currentColor"/>
          <path d="M7.8 14.2c1 1.5 2.6 2.3 4.2 2.3s3.2-.8 4.2-2.3" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
        </symbol>

        <symbol id="cs-icon-heart-outline" viewBox="0 0 24 24">
          <path d="M12 20s-7.4-4.6-9.7-9.2C.9 7.5 2.3 4 5.9 3.4c2-.3 3.9.6 5 2.3l1.1 1.6 1.1-1.6c1.1-1.7 3-2.6 5-2.3 3.6.6 5 4.1 3.6 7.4C19.4 15.4 12 20 12 20z" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round"/>
        </symbol>
        <symbol id="cs-icon-heart-filled" viewBox="0 0 24 24">
          <path d="M12 20s-7.4-4.6-9.7-9.2C.9 7.5 2.3 4 5.9 3.4c2-.3 3.9.6 5 2.3l1.1 1.6 1.1-1.6c1.1-1.7 3-2.6 5-2.3 3.6.6 5 4.1 3.6 7.4C19.4 15.4 12 20 12 20z" fill="currentColor" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round"/>
        </symbol>

        <symbol id="cs-icon-clock" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="9.25" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <path d="M12 7v5.3l3.6 2.1" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round" stroke-linejoin="round"/>
        </symbol>

        <symbol id="cs-icon-close" viewBox="0 0 24 24">
          <path d="M5.5 5.5l13 13m0-13l-13 13" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round"/>
        </symbol>
      </defs>
    </svg>

    <p v-if="totalCount" class="embedded-count">{{ totalCount }} comment{{ totalCount > 1 ? 's' : '' }}</p>

    <!-- New top-level comment composer -->
    <div class="composer-row">
      <div class="avatar" :style="{ background: currentUser.color }">
        {{ currentUser.initials }}
      </div>

      <div class="composer card">
        <textarea
          v-model="draft.text"
          maxlength="1000"
          placeholder="Comment Now"
          rows="1"
          @input="autoGrow"
        ></textarea>

        <div v-if="draft.images.length" class="img-preview-grid">
          <div class="img-preview" v-for="(img, i) in draft.images" :key="i">
            <img :src="img.previewUrl" alt="attachment preview" />
            <button @click="removeDraftImage(i)">
              <svg class="icon icon-xs"><use href="#cs-icon-close" /></svg>
            </button>
          </div>
        </div>

        <div class="composer-toolbar">
          <button class="tool-btn" title="Attach image" @click="triggerFile('draft')">
            <svg class="icon"><use href="#cs-icon-image" /></svg>
          </button>
          <input
            type="file"
            accept="image/*"
            multiple
            ref="fileInputDraft"
            style="display: none"
            @change="onFile($event, 'draft')"
          />

          <button class="tool-btn" title="Mention someone">
            <svg class="icon"><use href="#cs-icon-at" /></svg>
          </button>

          <div class="emoji-wrap">
            <button class="tool-btn" title="Add emoji" @click="toggleEmoji('draft')">
              <svg class="icon"><use href="#cs-icon-smile" /></svg>
            </button>

            <div class="emoji-pop" v-if="emojiOpenFor === 'draft'">
              <span v-for="e in emojis" :key="e" @click="insertEmoji(e, 'draft')">{{ e }}</span>
            </div>
          </div>

          <span class="char-count" :class="{ warn: draft.text.length > 900 }">
            {{ draft.text.length }}/1000
          </span>
          <button class="send-btn" :disabled="!draft.text.trim() && !draft.images.length" @click="postComment">
            Send
          </button>
        </div>
      </div>
    </div>

    <hr v-if="comments.length" class="divider" />
    <p v-else class="empty-state">No comments yet</p>

    <!-- Comment list -->
    <p v-if="isLoading" class="empty-state">waitng...</p>
    <p v-else-if="error" class="empty-state style-error">{{ error }}</p>

    <div v-if="!isLoading" v-for="comment in comments" :key="comment.id">
      <div class="comment">
        <div class="avatar" :style="{ background: comment.avatarUrl ? 'transparent' : comment.color }">
          <img v-if="comment.avatarUrl" :src="comment.avatarUrl" alt="" />
          <span v-else>{{ comment.initials }}</span>
        </div>

        <div class="comment-body">
          <div class="bubble">
            <div class="username">{{ comment.username }}</div>
            <div class="comment-text">{{ comment.text }}</div>
          </div>

          <div v-if="comment.images && comment.images.length" class="comment-images">
            <img
              v-for="(img, i) in comment.images"
              :key="i"
              :src="img"
              class="comment-img"
              @click="openLightbox(comment.images, i)"
            />
          </div>

          <div class="meta-row">
            <span class="meta-time">
              <svg class="icon icon-xs"><use href="#cs-icon-clock" /></svg>
              {{ comment.time }}
            </span>
            <button class="action-btn" @click="toggleReply(comment.id)">Reply</button>
            <button v-if="comment.userId === currentUser.id" class="action-btn text-danger" @click="deleteComment(comment.id)">Delete</button>
            <div class="spacer"></div>
            <button class="action-btn react-btn" :class="{ liked: comment.liked }" @click="toggleLike(comment)">
              <svg class="icon icon-sm"><use :href="comment.liked ? '#cs-icon-heart-filled' : '#cs-icon-heart-outline'" /></svg>
              {{ comment.likes }}
            </button>
          </div>

          <div v-if="replyOpenFor === comment.id" class="reply-composer card">
            <textarea
              v-model="comment.replyDraft"
              maxlength="1000"
              placeholder="Write a reply..."
              rows="1"
              @input="autoGrow"
            ></textarea>

            <div v-if="comment.replyImages && comment.replyImages.length" class="img-preview-grid">
              <div class="img-preview" v-for="(img, i) in comment.replyImages" :key="i">
                <img :src="img.previewUrl" alt="attachment preview" />
                <button @click="removeReplyImage(comment, i)">
                  <svg class="icon icon-xs"><use href="#cs-icon-close" /></svg>
                </button>
              </div>
            </div>

            <div class="composer-toolbar">
              <button class="tool-btn" title="Attach image" @click="triggerReplyFile(comment.id)">
                <svg class="icon"><use href="#cs-icon-image" /></svg>
              </button>
              <input
                type="file"
                accept="image/*"
                multiple
                :ref="(el) => setReplyFileRef(el, comment.id)"
                style="display: none"
                @change="onFile($event, 'reply', comment)"
              />

              <button class="tool-btn" title="Mention someone">
                <svg class="icon"><use href="#cs-icon-at" /></svg>
              </button>

              <div class="emoji-wrap">
                <button class="tool-btn" title="Add emoji" @click="toggleEmoji('reply-' + comment.id)">
                  <svg class="icon"><use href="#cs-icon-smile" /></svg>
                </button>
                <div class="emoji-pop" v-if="emojiOpenFor === 'reply-' + comment.id">
                  <span
                    v-for="e in emojis"
                    :key="e"
                    @click="insertEmoji(e, 'reply', comment)"
                  >{{ e }}</span>
                </div>
              </div>

              <span class="char-count" :class="{ warn: (comment.replyDraft || '').length > 900 }">
                {{ (comment.replyDraft || '').length }}/1000
              </span>
              <button
                class="send-btn"
                :disabled="!(comment.replyDraft || '').trim() && !(comment.replyImages || []).length"
                @click="postReply(comment)"
              >
                Send
              </button>
            </div>
          </div>

          <div v-if="comment.replies.length" class="replies">
            <div class="comment" v-for="reply in comment.replies" :key="reply.id">
              <div class="avatar" :style="{ background: reply.avatarUrl ? 'transparent' : reply.color }">
                <img v-if="reply.avatarUrl" :src="reply.avatarUrl" alt="" />
                <span v-else>{{ reply.initials }}</span>
              </div>

              <div class="comment-body">
                <div class="bubble">
                  <div class="username">{{ reply.username }}</div>
                  <div class="comment-text">{{ reply.text }}</div>
                </div>

                <div v-if="reply.images && reply.images.length" class="comment-images">
                  <img
                    v-for="(img, i) in reply.images"
                    :key="i"
                    :src="img"
                    class="comment-img"
                    @click="openLightbox(reply.images, i)"
                  />
                </div>

                <div class="meta-row">
                  <span class="meta-time">
                    <svg class="icon icon-xs"><use href="#cs-icon-clock" /></svg>
                    {{ reply.time }}
                  </span>
                  <button v-if="reply.userId === currentUser.id" class="action-btn text-danger" @click="deleteComment(reply.id)">Delete</button>
                  <div class="spacer"></div>
                </div>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>

    <Teleport to="body">
      <div v-if="lightbox.open" class="lightbox-overlay" @click.self="closeLightbox">
        <button class="lightbox-close" @click="closeLightbox">
          <svg class="icon"><use href="#cs-icon-close" /></svg>
        </button>

        <button v-if="lightbox.images.length > 1" class="lightbox-nav lightbox-prev" @click="lightboxPrev">‹</button>

        <img :src="lightbox.images[lightbox.index]" class="lightbox-img" @click.stop />

        <button v-if="lightbox.images.length > 1" class="lightbox-nav lightbox-next" @click="lightboxNext">›</button>

        <div v-if="lightbox.images.length > 1" class="lightbox-counter">
          {{ lightbox.index + 1 }} / {{ lightbox.images.length }}
        </div>
      </div>
    </Teleport>

  </div>
</template>

<style scoped>
.comment-section {
  --ink: #2b2b2b;
  --ink-soft: #6b7280;
  --line: #e7e7e7;
  --paper: #ffffff;
  --canvas: #f5f6f8;
  --accent: #1976d2;
  --accent-soft: #eff6fb;
  --heart: #f2762e;

  font-family: 'Inter', sans-serif;
  color: var(--ink);
  background: var(--canvas);
  padding: 28px;
  border-radius: 20px;
  max-width: 720px;
  margin: 0 auto;
}

.comment-section.embedded {
  --paper: #f7f9fb;
  background: transparent;
  padding: 0;
  border-radius: 0;
  max-width: none;
  margin: 14px 0 0;
}

.embedded-count {
  font-family: 'Nunito', sans-serif;
  font-size: 12.5px;
  font-weight: 700;
  color: var(--ink-soft);
  margin: 0 0 10px 2px;
}

.empty-state {
  font-size: 13px;
  color: var(--ink-soft);
  margin: 4px 0 16px 2px;
}

.card {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 16px;
  padding: 16px;
}

.composer-row {
  display: flex;
  gap: 12px;
  margin-bottom: 18px;
}

.avatar {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  flex: none;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 13px;
  color: #fff;
  overflow: hidden;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.composer {
  flex: 1;
}

textarea {
  width: 100%;
  border: none;
  resize: none;
  font-family: 'Inter', sans-serif;
  font-size: 14px;
  line-height: 1.5;
  color: var(--ink);
  outline: none;
  min-height: 22px;
  background: transparent;
}

textarea::placeholder {
  color: #9aa2ae;
}

.composer-toolbar {
  display: flex;
  align-items: center;
  gap: 14px;
  padding-top: 10px;
  margin-top: 8px;
  border-top: 1px solid var(--line);
}

.tool-btn {
  background: none;
  border: none;
  cursor: pointer;
  width: 30px;
  height: 30px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--ink-soft);
}

.icon {
  width: 18px;
  height: 18px;
  flex: none;
  display: block;
}

.icon-sm {
  width: 16px;
  height: 16px;
}

.icon-xs {
  width: 13px;
  height: 13px;
}

.tool-btn:hover {
  background: var(--accent-soft);
  color: var(--accent);
}

.char-count {
  margin-left: auto;
  font-size: 12.5px;
  color: #b0b6c0;
  font-variant-numeric: tabular-nums;
}

.char-count.warn {
  color: var(--heart);
}

.send-btn {
  background: var(--accent);
  color: #fff;
  border: none;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 13px;
  padding: 7px 16px;
  border-radius: 999px;
  cursor: pointer;
  transition: opacity 0.12s ease;
}

.send-btn:hover:not(:disabled) {
  opacity: 0.85;
}

.send-btn:disabled {
  background: #c9cedA;
  cursor: not-allowed;
}

.img-preview {
  position: relative;
  width: 90px;
  height: 90px;
  border-radius: 10px;
  overflow: hidden;
  flex: none;
}

.img-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.img-preview button {
  position: absolute;
  top: 6px;
  right: 6px;

  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: rgba(0,0,0,.75);
  color: #fff;
  border: none;
}

.emoji-wrap {
  position: relative;
}

.emoji-pop {
  position: absolute;
  bottom: 36px;
  left: 0;
  background: #fff;
  border: 1px solid var(--line);
  border-radius: 12px;
  padding: 8px;
  display: flex;
  gap: 4px;
  box-shadow: 0 8px 24px rgba(20, 24, 40, 0.12);
  z-index: 5;
}

.emoji-pop span {
  cursor: pointer;
  font-size: 18px;
  padding: 4px;
  border-radius: 6px;
}

.emoji-pop span:hover {
  background: var(--accent-soft);
}

.divider {
  border: none;
  border-top: 1px solid var(--line);
  margin: 18px 0;
}

.comment {
  display: flex;
  gap: 12px;
  margin-bottom: 18px;
}

.comment-body {
  flex: 1;
  min-width: 0;
}

.bubble {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 14px;
  padding: 10px 13px;
}

.username {
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 13px;
  margin-bottom: 4px;
}

.comment-text {
  font-size: 13.5px;
  line-height: 1.5;
  color: var(--ink);
}

.meta-row {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-top: 6px;
  padding-left: 4px;
}

.meta-time {
  font-size: 11.5px;
  color: #9aa2ae;
  display: inline-flex;
  align-items: center;
  gap: 5px;
}

.action-btn {
  background: none;
  border: none;
  cursor: pointer;
  font-family: 'Inter', sans-serif;
  font-weight: 600;
  font-size: 12px;
  color: var(--ink-soft);
  padding: 2px 4px;
  display: inline-flex;
  align-items: center;
  gap: 5px;
}

.action-btn:hover {
  color: var(--accent);
}

.action-btn.liked {
  color: var(--heart);
}

.spacer {
  flex: 1;
}

.reply-composer {
  margin-top: 12px;
}

.replies {
  margin-top: 12px;
  padding-left: 18px;
  border-left: 2px solid var(--line);
}

.replies .comment {
  margin-bottom: 14px;
}

.replies .avatar {
  width: 30px;
  height: 30px;
  font-size: 11px;
}

.img-preview-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-top: 10px;
}

.comment-images {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-top: 8px;
}

.comment-img {
  width: 90px;
  height: 90px;
  object-fit: cover;
  border-radius: 10px;
  cursor: pointer;
  transition: transform 0.12s ease;
}

.comment-img:hover {
  transform: scale(1.03);
}

.lightbox-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.85);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}

.lightbox-img {
  max-width: 88vw;
  max-height: 86vh;
  border-radius: 8px;
  object-fit: contain;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
}

.lightbox-close {
  position: absolute;
  top: 20px;
  right: 24px;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.12);
  border: none;
  color: #fff;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.lightbox-close:hover {
  background: rgba(255, 255, 255, 0.22);
}

.lightbox-nav {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.12);
  border: none;
  color: #fff;
  font-size: 28px;
  line-height: 1;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.lightbox-nav:hover {
  background: rgba(255, 255, 255, 0.22);
}

.lightbox-prev {
  left: 20px;
}

.lightbox-next {
  right: 20px;
}

.lightbox-counter {
  position: absolute;
  bottom: 24px;
  left: 50%;
  transform: translateX(-50%);
  color: #fff;
  font-family: 'Nunito', sans-serif;
  font-size: 13px;
  font-weight: 600;
  background: rgba(0, 0, 0, 0.4);
  padding: 4px 12px;
  border-radius: 999px;
}
</style>