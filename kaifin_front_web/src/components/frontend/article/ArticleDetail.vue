<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import ArticleShare from './ArticleShare.vue'
import ArticleReport from './ArticleReport.vue'
import ArticleReaction from './ArticleReaction.vue'
import ArticleComment from './ArticleComment.vue'
import { REACTIONS } from '../../../components/reaction/reactions.js'

const props = defineProps({
  articleData: {
    type: Object,
    required: true,
    default: () => ({})
  }
})

const API_BASE = 'http://localhost:7070'

// ✅ Key សម្រាប់ save/load reaction type ក្នុង localStorage តាម article id
const REACTION_STORAGE_KEY = computed(() => `article_reaction_${props.articleData.id}`)

function loadStoredReaction() {
  if (!props.articleData.id) return 'like'
  const stored = localStorage.getItem(REACTION_STORAGE_KEY.value)
  return stored || 'like'
}

function saveStoredReaction(key) {
  if (!props.articleData.id) return
  localStorage.setItem(REACTION_STORAGE_KEY.value, key)
}

function clearStoredReaction() {
  if (!props.articleData.id) return
  localStorage.removeItem(REACTION_STORAGE_KEY.value)
}

const showSharePopup = ref(false)
const toggleSharePopup = () => {
  showSharePopup.value = !showSharePopup.value
}

const showReportModal = ref(false)

// ✅ ដំបូង init ពី localStorage ជំនួសឲ្យ hardcode 'like'
const selectedReaction = ref(loadStoredReaction())
const showReactionPopup = ref(false)
const isLiked = ref(!!props.articleData.liked)
const likeCount = ref(Number(props.articleData.likes) || 0)
const isTogglingLike = ref(false)

const currentReactionIcon = computed(() => {
  const found = REACTIONS.find(r => r.key === selectedReaction.value)
  return found ? found.icon : null
})

async function toggleLike() {
  if (isTogglingLike.value || !props.articleData.id) return
  isTogglingLike.value = true
  showReactionPopup.value = false

  const prevLiked = isLiked.value
  const prevCount = likeCount.value
  isLiked.value = !prevLiked
  likeCount.value = prevLiked ? Math.max(0, prevCount - 1) : prevCount + 1

  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/articles/${props.articleData.id}/like`, {
      method: 'POST',
      headers: token ? { Authorization: `Bearer ${token}` } : {}
    })
    const raw = await res.text()
    let data = null
    if (raw) {
      try { data = JSON.parse(raw) } catch {
        throw new Error(`Server returned non-JSON response (status ${res.status})`)
      }
    }
    if (!res.ok) {
      throw new Error(data?.message || `Request failed with status ${res.status}`)
    }
    const serverLiked = data?.data?.liked
    if (typeof serverLiked === 'boolean' && serverLiked !== isLiked.value) {
      isLiked.value = serverLiked
      likeCount.value = serverLiked ? prevCount + 1 : Math.max(0, prevCount - 1)
    }

    // ✅ Save/clear reaction ក្នុង localStorage តាមស្ថានភាព like ចុងក្រោយ
    if (isLiked.value) {
      saveStoredReaction(selectedReaction.value)
    } else {
      clearStoredReaction()
      selectedReaction.value = 'like' // reset default សម្រាប់ពេល like លើកក្រោយ
    }
  } catch (err) {
    console.error(err)
    isLiked.value = prevLiked
    likeCount.value = prevCount
  } finally {
    isTogglingLike.value = false
  }
}

function handleLikeButtonClick() {
  if (isLiked.value) {
    toggleLike()
  } else {
    showReactionPopup.value = !showReactionPopup.value
  }
}

const handleReactionSelect = (reaction) => {
  selectedReaction.value = reaction.key
  showReactionPopup.value = false
  if (!isLiked.value) {
    toggleLike()
  } else {
    // ✅ ករណី user ចង់ប្តូរ reaction ខណៈ liked រួចហើយ (ចុច icon ក្នុង popup ដោយផ្ទាល់)
    saveStoredReaction(reaction.key)
  }
}

const showCommentDrawer = ref(false)

watch(showCommentDrawer, (isOpen) => {
  if (isOpen) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
})

const handleClickOutside = (event) => {
  const shareWrapper = document.querySelector('.share-wrapper-container')
  if (shareWrapper && !shareWrapper.contains(event.target)) {
    showSharePopup.value = false
  }

  const likeWrapper = document.querySelector('.like-wrapper-container')
  if (likeWrapper && !likeWrapper.contains(event.target)) {
    showReactionPopup.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  document.body.style.overflow = ''
})

const displayComments = computed(() => props.articleData.comments || '0')
const displaySaves = computed(() => props.articleData.saves || '0')
const displayShares = computed(() => props.articleData.shares || '0')

const displayContentBlocks = computed(() => {
  return (props.articleData.contentBlocks || []).map(block => {
    if (block.type === 'image' && !block.value) {
      return { ...block, value: 'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?auto=format&fit=crop&w=800&q=80' }
    }
    return block
  })
})
</script>

<template>
  <div class="article-detail-layout">

    <!-- Left Sidebar Actions -->
    <aside class="action-sidebar-left">
      <div class="sidebar-sticky-inner">
        <!-- Like Button with Reaction Popup -->
        <div class="action-wrapper like-wrapper-container">
          <button
            class="action-btn"
            :class="{ 'is-liked': isLiked }"
            :disabled="isTogglingLike"
            @click.stop="handleLikeButtonClick"
          >
            <!-- ✅ បើ liked ហើយ និងមាន reaction icon → បង្ហាញ sticker នោះ -->
            <span
              v-if="isLiked && currentReactionIcon"
              class="reaction-sticker"
              v-html="currentReactionIcon"
            ></span>
            <!-- ✅ បើមិនទាន់ liked ឬគ្មាន icon → បង្ហាញ heart icon លំនាំដើម -->
            <svg
              v-else
              width="20" height="20" viewBox="0 0 24 24"
              :fill="isLiked ? 'currentColor' : 'none'"
              stroke="currentColor" stroke-width="2"
            >
              <path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path>
            </svg>
          </button>
          <span class="count-badge">{{ likeCount }}</span>

          <div v-if="showReactionPopup" class="reaction-popup-dropdown">
            <ArticleReaction 
              v-model="selectedReaction" 
              @react="handleReactionSelect" 
            />
          </div>
        </div>

        <!-- Comment Button -->
        <div class="action-wrapper">
          <button class="action-btn" @click="showCommentDrawer = true">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
          </button>
          <span class="count-badge">{{ displayComments }}</span>
        </div>

        <!-- Save Button -->
        <div class="action-wrapper">
          <button class="action-btn">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path></svg>
          </button>
          <span class="count-badge">{{ displaySaves }}</span>
        </div>

        <!-- Share Button -->
        <div class="action-wrapper share-wrapper-container">
          <button class="action-btn" @click.stop="toggleSharePopup">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="18" cy="5" r="3"></circle><circle cx="6" cy="12" r="3"></circle><circle cx="18" cy="19" r="3"></circle><line x1="8.59" y1="13.51" x2="15.42" y2="17.49"></line><line x1="15.41" y1="6.51" x2="8.59" y2="10.49"></line></svg>
          </button>
          <span class="count-badge">{{ displayShares }}</span>
          
          <div v-if="showSharePopup" class="share-popup-dropdown">
            <ArticleShare />
          </div>
        </div>

        <!-- Report Button -->
        <div class="action-wrapper">
          <button class="action-btn" title="Report" @click="showReportModal = true">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path><line x1="12" y1="9" x2="12" y2="13"></line><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
          </button>
        </div>
      </div>
    </aside>

    <!-- Main Article Content -->
    <main class="article-main-content">
      <div class="article-header-box">
        <h1 class="article-title">{{ props.articleData.title }}</h1>
      </div>

      <div class="article-meta">
        <span class="meta-tag">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
          {{ props.articleData.date }}
        </span>
        <span class="meta-tag">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
          {{ props.articleData.time }}
        </span>
        <span class="meta-tag">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path><circle cx="12" cy="12" r="3"></circle></svg>
          {{ props.articleData.views }}
        </span>
      </div>

      <div class="article-body-content">
        <template v-for="(block, index) in displayContentBlocks" :key="index">
          <div v-if="block.type === 'text'" class="content-block">
            <h3 v-if="block.title">{{ block.title }}</h3>
            <div class="content-block-html" v-html="block.value"></div>
          </div>
          <div v-else-if="block.type === 'image'" class="image-wrapper-left">
            <img :src="block.value" alt="Article Image" class="article-rendered-img" />
          </div>
        </template>
      </div>
    </main>

    <!-- ArticleReport Modal Component -->
    <ArticleReport 
      v-if="showReportModal" 
      @close="showReportModal = false" 
    />

    <!-- Comment Right Sidebar Drawer (Under Navbar) -->
    <div v-if="showCommentDrawer" class="comment-drawer-backdrop" @click.self="showCommentDrawer = false">
      <div class="comment-drawer-content">
        <div class="drawer-scrollable-body">
          <div class="drawer-inner-wrapper">
            <button class="close-drawer-btn" @click="showCommentDrawer = false">&times;</button>
            <ArticleComment
  :articleId="props.articleData.id"
  @comment-count-change="(n) => { /* optional: sync a local ref if you show count elsewhere */ }"
/>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<style scoped>
.article-detail-layout {
  width: 100%;
  margin: 0 auto;
  display: flex;
  align-items: stretch;
  gap: 12px;
  padding: 0;
  min-height: 100vh;
  position: relative;
}

.action-sidebar-left {
  flex: 0 0 85px;
  width: 80px;
  display: flex;
  position: relative;
  margin-top: -12px;
  margin-bottom: -12px;
  padding-right: 14px;
}

.action-sidebar-left::after {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  width: 1.5px;
  background: #ddd;
}

.sidebar-sticky-inner {
  position: sticky;
  top: 0;
  height: 100vh;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 20px;
  padding: 0;
}

.action-wrapper {
  position: relative;
  display: inline-block;
}

.share-popup-dropdown,
.reaction-popup-dropdown {
  position: absolute;
  left: 70px;
  top: 50%;
  transform: translateY(-50%);
  z-index: 99;
}

.action-btn {
  background-color: #2E6FD9;
  border: none;
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  width: 42px;
  height: 42px;
  border-radius: 50%;
  transition: transform 0.2s;
}

.action-btn:hover {
  transform: translateY(-2px);
}

.action-btn.is-liked {
  /* background-color: #F0554A; */
  background-color: #ffffff;
  border: 2px solid #2E6FD9;
}

.action-btn:disabled {
  opacity: 0.7;
  cursor: default;
  transform: none;
}

/* ✅ ថ្មី — Style សម្រាប់ reaction sticker ក្នុងប៊ូតុង Like */
.reaction-sticker {
  width: 22px;
  height: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.reaction-sticker :deep(svg) {
  width: 100%;
  height: 100%;
  display: block;
}

.count-badge {
  position: absolute;
  top: -12px;
  right: -8px;
  background-color: #90A4AE;
  color: #ffffff;
  font-size: 11px;
  font-weight: 600;
  padding: 2px 8px;
  border-radius: 50px;
  white-space: nowrap;
  z-index: 2;
}

.article-main-content {
  flex: 1;
  min-width: 0;
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 14px; 
}

.article-title {
  font-size: 28px;
  font-weight: 700;
  color: #212529;
  line-height: 1.3;
  margin: 0;
}

.article-meta {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

.meta-tag {
  font-size: 13px;
  font-weight: 500;
  color: #495057;
  display: flex;
  align-items: center;
  gap: 6px;
}

.article-body-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.content-block h3 {
  font-size: 18px;
  font-weight: 600;
  color: #343a40;
  margin-bottom: 8px;
}

.content-block p {
  font-size: 16px;
  color: #495057;
  line-height: 1.8;
}

.content-block-html {
  font-size: 16px;
  color: #495057;
  line-height: 1.8;
  overflow-wrap: break-word;
}

.content-block-html :deep(p) {
  margin: 0 0 12px;
}

.content-block-html :deep(img) {
  max-width: 100%;
  border-radius: 8px;
  display: block;
  margin: 12px 0;
}

.content-block-html :deep(b),
.content-block-html :deep(strong) {
  font-weight: 700;
  color: #212529;
}

.image-wrapper-left {
  float: left;
  width: 240px;
  height: 160px;
  margin-right: 20px;
  margin-bottom: 12px;
  background: #ececec;
  border: 1px solid #c4c4c4;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.article-rendered-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.comment-drawer-backdrop {
  position: fixed;
  top: 64.5px;
  left: 0;
  width: 100vw;
  height: calc(100vh - 60px);
  background-color: rgba(0, 0, 0, 0.365);
  z-index: 1;
  display: flex;
  justify-content: flex-end;
  align-items: stretch;
  animation: fadeIn 0.3s ease;
}

.comment-drawer-content {
  width: 100%;
  max-width: 650px;
  height: 100%;
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
  box-shadow: -4px 0 20px rgba(0, 0, 0, 0.15);
  animation: slideRight 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

.drawer-scrollable-body {
  overflow-y: auto;
  flex: 1;
  display: flex;
  justify-content: center;
  height: 100%;
}

.drawer-inner-wrapper {
  position: relative;
  width: 100%;
  max-width: 100%;
}

.close-drawer-btn {
  position: absolute;
  right: 20px;
  top: 16px;
  background: none;
  border: none;
  color: #9ca3af;
  font-size: 24px;
  cursor: pointer;
  z-index: 10;
  line-height: 1;
}

.close-drawer-btn:hover {
  color: #111827;
}

@keyframes slideRight {
  from { transform: translateX(100%); }
  to { transform: translateX(0); }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>