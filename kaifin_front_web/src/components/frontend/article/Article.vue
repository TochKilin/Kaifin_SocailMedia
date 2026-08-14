<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue'
import NavBar from '../navbar/NavBar.vue'
import CreateArticle from './CreateArticle.vue'
import ArticleDetail from './ArticleDetail.vue'

// backend base URL — used for every API call, no Vite proxy required
const API_BASE = 'http://localhost:7070'

const selectedArticle = ref(null)

const currentView = ref('home')       // 'home' | 'following' | 'code' | 'read' | 'ranking' | 'detail'
const activeTab = ref('Popular')
const showCreateModal = ref(false)
const showCreatePage = ref(false)

// ============ Article options dropdown ============
const openOptionsMenuId = ref(null)

function toggleOptionsMenu(articleId) {
  openOptionsMenuId.value = openOptionsMenuId.value === articleId ? null : articleId
}

function closeOptionsMenu() {
  openOptionsMenuId.value = null
}

function handleNotInterested(article) {
  closeOptionsMenu()
  // TODO: call API / remove from feed
  articles.value = articles.value.filter(a => a.id !== article.id)
}

function handleBlockUser(article) {
  closeOptionsMenu()
  // TODO: call API to block author
  articles.value = articles.value.filter(a => a.author !== article.author)
}

function handleReport(article) {
  closeOptionsMenu()
  // TODO: open report modal / call API
  console.log('Reported article:', article.id)
}

function handleGlobalClick(event) {
  if (openOptionsMenuId.value !== null) {
    const el = event.target.closest('.options-menu-wrap')
    if (!el) closeOptionsMenu()
  }
}

onMounted(() => {
  window.addEventListener('click', handleGlobalClick)
  fetchArticles()
})

onBeforeUnmount(() => {
  window.removeEventListener('click', handleGlobalClick)
})

// ============ Code Life submenu ============
const showCodeSubmenu = ref(false)
const activeCodeCategory = ref('all') // 'all' | 'backend' | 'frontend' | 'ai' | 'tools'

const codeSubcategories = [
  { key: 'backend', label: 'Back End' },
  { key: 'frontend', label: 'Front End' },
  { key: 'ai', label: 'AI' },
  { key: 'tools', label: 'Development Tools' },
]

function openCodeLife() {
  currentView.value = 'code'
  showCodeSubmenu.value = !showCodeSubmenu.value
}

function selectCodeCategory(key) {
  activeCodeCategory.value = key
  currentView.value = 'code'
  showCodeSubmenu.value = true
}

const handleCreateArticle = () => {
  showCreatePage.value = true
}

// ============ Real article data from API ============

const articles = ref([])
const isLoadingArticles = ref(false)
const articlesError = ref('')

// resolve backend-relative image paths ("articles/xxx.jpg") into a usable full URL
function resolveImageUrl(path) {
  if (!path) return 'https://picsum.photos/seed/default/200'
  if (path.startsWith('http://') || path.startsWith('https://')) return path
  if (path.startsWith('/uploads/')) return `${API_BASE}${path}`
  return `${API_BASE}/uploads/${path}`
}

// map an article object as returned by the feed API (Show) into the shape the template expects
function mapApiArticle(a) {
  return {
    id: a.id,
    title: a.title,
    text: a.summary || '',
    secondaryText: a.summary || 'New community publication.',
    author: a.user_name || 'Unknown',
    authorProfile: resolveImageUrl(a.profile_images),
    views: String(a.views_count ?? 0),
    likes: String(a.like_count ?? 0),
    tags: a.tags || [],
    image: resolveImageUrl(a.cover_image),
    category: a.category,
    subcategory: a.code_subcategory || null,
  }
}

// map an article object as returned by the detail API (Detail) into the shape
// ArticleDetail.vue expects. Matches the Go `Article` / `ArticleBlock` structs exactly:
//   Article:      id, user_id, user_name, profile_images, title, summary, cover_image,
//                 category, code_subcategory, visibility, status, views_count, created_at,
//                 updated_at, like_count, liked, comment_count, save_count, saved, tags, blocks
//   ArticleBlock: id, article_id, block_type ('text'|'image'), title, content, position, created_at
function mapApiArticleDetail(a) {
  return {
    id: a.id,
    title: a.title,
    date: a.created_at
      ? new Date(a.created_at).toLocaleDateString('en-GB').split('/').join('-')
      : '',
    time: a.created_at
      ? new Date(a.created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
      : '',
    views: `${a.views_count ?? 0} Views`,
    likes: String(a.like_count ?? 0),
    comments: String(a.comment_count ?? 0),
    saves: String(a.save_count ?? 0),
    author: a.user_name || 'Unknown',
    authorProfile: resolveImageUrl(a.profile_images),
    liked: a.liked || false,
    saved: a.saved || false,
    tags: a.tags || [],
    // block_type/content (backend names) -> type/value (ArticleDetail.vue names)
    contentBlocks: (a.blocks || []).map(b => ({
      type: b.block_type,
      title: b.title || '',
      value: b.block_type === 'image' ? resolveImageUrl(b.content) : (b.content || ''),
    })),
  }
}

function categoryParamFor(view) {
  switch (view) {
    case 'following': return 'following'
    case 'code': return 'code'
    case 'read': return 'read'
    case 'home': return 'general'
    default: return ''
  }
}

async function fetchArticles(view = currentView.value) {
  const category = categoryParamFor(view)
  if (!category) return // ranking/detail don't need a feed fetch

  isLoadingArticles.value = true
  articlesError.value = ''
  try {
    const params = new URLSearchParams({
      category,
      page: '1',
      per_page: '50',
    })

    const token = localStorage.getItem('token')

    // Full backend URL — no Vite proxy dependency
    const res = await fetch(`${API_BASE}/api/v1/front/articles/show?${params.toString()}`, {
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

    // Show() returns ArticlesResponse -> { articles, total, page, per_page }
    const list = data?.data?.articles || []
    articles.value = list.map(mapApiArticle)
  } catch (err) {
    console.error(err)
    articlesError.value = err.message || 'មិនអាចទាញអត្ថបទបានទេ'
    articles.value = []
  } finally {
    isLoadingArticles.value = false
  }
}

watch(currentView, (view) => {
  if (['home', 'following', 'code', 'read'].includes(view)) {
    fetchArticles(view)
  }
  if (view === 'ranking') {
    fetchRanking()
  }
})

// ============ Article detail (real fetch, no hardcoding) ============

const isLoadingDetail = ref(false)
const detailError = ref('')

async function openArticleDetail(article) {
  currentView.value = 'detail'
  selectedArticle.value = null
  isLoadingDetail.value = true
  detailError.value = ''
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/articles/${article.id}`, {
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

    // Detail() returns *Article directly -> response.data IS the Article object
    const a = data?.data
    if (!a) throw new Error('Article not found')

    selectedArticle.value = mapApiArticleDetail(a)
  } catch (err) {
    console.error(err)
    detailError.value = err.message || 'មិនអាចទាញអត្ថបទបានទេ'
  } finally {
    isLoadingDetail.value = false
  }
}

const submitArticle = (createdArticle) => {
  if (!createdArticle) return
  articles.value.unshift(mapApiArticle(createdArticle))
  showCreatePage.value = false
}

const filteredArticles = computed(() => {
  if (currentView.value === 'code' && activeCodeCategory.value !== 'all') {
    return articles.value.filter(a => a.subcategory === activeCodeCategory.value)
  }
  return articles.value
})

const feedTitle = computed(() => {
  switch (currentView.value) {
    case 'following': return 'Following'
    case 'code': return 'Code Life'
    case 'read': return 'Read'
    default: return 'General Post'
  }
})

// ============ Ranking (real fetch, no hardcoding) ============
// Backend: GET /articles/show supports generic sorting via
// ShowArticlesRequest.Sorts []share.Sort, bound from query params shaped
// like sorts[0][property]=like_count&sorts[0][direction]=desc.
// article_repo.go whitelists sortable properties to: created_at,
// views_count, like_count, title — see articleSortColumns in the repo.
//
// NOTE: verify this query-param shape actually binds into req.Sorts on
// your Fiber setup (log u.Sorts inside ShowArticlesRequest.bind() to
// confirm) — Fiber's query struct binding for slices-of-structs varies by
// version/config. If it doesn't bind, try the dot-notation fallback:
// sorts[0].property / sorts[0].direction instead.

const rankingArticlesList = ref([])
const isLoadingRanking = ref(false)
const rankingError = ref('')

// ranking list shape used by the template: id, title, username, userProfile,
// views, comments, saves, likes — mapped from the same Article API shape
// used elsewhere on this page.
function mapApiArticleRanking(a) {
  return {
    id: a.id,
    title: a.title,
    username: a.user_name || 'Unknown',
    userProfile: resolveImageUrl(a.profile_images),
    views: String(a.views_count ?? 0),
    comments: String(a.comment_count ?? 0),
    saves: String(a.save_count ?? 0),
    likes: a.like_count ?? 0, // numeric: template does `item.likes++` and shows `{{ item.likes }}k`
  }
}

async function fetchRanking() {
  isLoadingRanking.value = true
  rankingError.value = ''
  try {
    const params = new URLSearchParams({
      page: '1',
      per_page: '15',
    })
    // Sort by like_count desc. Swap property to 'views_count' if you'd
    // rather rank by views instead of likes.
    params.append('sorts[0][property]', 'like_count')
    params.append('sorts[0][direction]', 'desc')

    const token = localStorage.getItem('token')

    const res = await fetch(`${API_BASE}/api/v1/front/articles/show?${params.toString()}`, {
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

    const list = data?.data?.articles || []
    rankingArticlesList.value = list.map(mapApiArticleRanking)
  } catch (err) {
    console.error(err)
    rankingError.value = err.message || 'មិនអាចទាញ ranking បានទេ'
    rankingArticlesList.value = []
  } finally {
    isLoadingRanking.value = false
  }
}

const listArticles = [
  { id: 1, title: 'Understanding Vue 3 Composition API' },
  { id: 2, title: 'Optimizing Vite Build Speed' },
  { id: 3, title: 'Advanced CSS Grid Layouts' },
  { id: 4, title: 'State Management in Large Apps' },
  { id: 5, title: 'Clean Code Principles' }
]

const authors = [
  { id: 1, name: 'Sopheak Dev', avatar: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=100&h=100&fit=crop' },
  { id: 2, name: 'Dara Code', avatar: 'https://images.unsplash.com/photo-1522075469751-3a6694fb2f61?w=100&h=100&fit=crop' },
  { id: 3, name: 'Mony Tech', avatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=100&h=100&fit=crop' }
]

const topics = [
  { id: 1, name: 'Web Development', image: 'https://images.unsplash.com/photo-1461749280684-dccba630e2f6?w=100&h=100&fit=crop' },
  { id: 2, name: 'UI/UX Design', image: 'https://images.unsplash.com/photo-1507238691740-187a5b1d37b8?w=100&h=100&fit=crop' },
  { id: 3, name: 'Cloud Computing', image: 'https://images.unsplash.com/photo-1451187580459-43490279c0fa?w=100&h=100&fit=crop' },
  { id: 4, name: 'Artificial Intelligence', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRV6g5yL7h84YQG_v53YZXElNiDqT2cc3kukAUz65NPuQ&s=10' },
  { id: 5, name: 'Cybersecurity', image: 'https://images.unsplash.com/photo-1550751827-4bd374c3f58b?w=100&h=100&fit=crop' }
]

const topAuthors = ref([
  {
    id: 1,
    name: 'Phearn Sophea',
    level: 25,
    emoji: '👑',
    color: '#D4A017',
    articles: '1',
    followers: '1M',
    isFollowing: false,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTaoPH7clrp9jbbxCcMD5K8VTUiSQ0KrlbJeiuO5TXRtw&s=10'
  },
  {
    id: 2,
    name: 'Lon ReakSmey',
    level: 20,
    emoji: '🔥',
    articles: '19',
    followers: '9K',
    isFollowing: false,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRjsa_MQTvNnWSBVihbhp8qLzS6ow_g1jr84z5tM9tW8g&s=10'
  },
  {
    id: 3,
    name: 'Sophany Davy',
    level: 16,
    emoji: '💎',
    articles: '16',
    followers: '7K',
    isFollowing: false,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSl2OEKDJsDC60NkkPMTX90ruXQOTCENwy_bw5c1GtnSg&s'
  },
  {
    id: 4,
    name: 'កុមារី វ័យក្នេង',
    level: 12,
    articles: 820,
    followers: '5K',
    isFollowing: false,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRV1ZNIi0RdiC01_IpMq8qCQF7xYtzTsOHtfJyFGoQlEQ&s=10'
  },
  {
    id: 5,
    name: 'Jom Chana',
    level: 12,
    emoji: '⭐',
    articles: 820,
    followers: '5K',
    isFollowing: false,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTJF1cZrlxJ3OElPHek19UKSdJNTX8Y03KQiisC8dRxZQ&s=10'
  },
  {
    id: 6,
    name: 'ណុត ធារី',
    level: 12,
    emoji: '⭐',
    articles: 820,
    followers: '5K',
    isFollowing: false,
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTEFgwvTg5wioafmToMPHWgBWQcBEjWBh2dsuVGowOl7g&s=10'
  },
])

const authorsCardAfterIndex = computed(() => {
  const len = filteredArticles.value.length
  if (len < 2) return -1
  return Math.floor(len / 2) - 1
})

const toggleFollow = (author) => {
  author.isFollowing = !author.isFollowing
}

const detailAuthorInfo = computed(() => {
  const art = selectedArticle.value
  if (!art) return null
  return {
    name: art.author || 'Username',
    avatar: art.authorProfile || '',
    level: art.authorLevel || 12,
    articleCount: art.authorArticleCount || '200',
    likes: art.likes || '0',
    followers: art.authorFollowers || '0',
    isFollowing: art.authorIsFollowing || false,
  }
})

function toggleDetailAuthorFollow() {
  if (!selectedArticle.value) return
  selectedArticle.value.authorIsFollowing = !selectedArticle.value.authorIsFollowing
}

// Table of contents សម្រាប់អត្ថបទដែលកំពុងបើក — ទាញពី contentBlocks ពិត (មិន hardcode)
const detailTocItems = computed(() => {
  const art = selectedArticle.value
  if (!art || !art.contentBlocks) return []
  return art.contentBlocks
    .filter(b => b.type === 'text' && b.title)
    .map(b => b.title)
})

const toRoman = (num) => {
  const lookup = { M: 1000, CM: 900, D: 500, CD: 400, C: 100, XC: 90, L: 50, XL: 40, X: 10, IX: 9, V: 5, IV: 4, I: 1 };
  let roman = '';
  for (let i in lookup) {
    while (num >= lookup[i]) {
      roman += i;
      num -= lookup[i];
    }
  }
  return roman;
}
</script>

<template>
  <div>
    <NavBar/>
    <div class="app-container">
      <div class="dashboard-layout">
        <!-- Left Sidebar -->
        <aside class="sidebar-left">
          <div class="brand-logo">
            <h2>DevSpace</h2>
          </div>
          <nav class="nav-menu">
            <button class="nav-item" @click="currentView = 'following'" :class="{ active: currentView === 'following' }">
              <svg class="nav-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/></svg>
              <span>Following</span>
            </button>
            <button class="nav-item" @click="currentView = 'home'" :class="{ active: currentView === 'home' }">
              <svg class="nav-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z"/></svg>
              <span>General post</span>
            </button>
            <div class="nav-item-wrap">
              <button class="nav-item" @click="openCodeLife" :class="{ active: currentView === 'code' }">
                <svg class="nav-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/></svg>
                <span>Code Life</span>
                <svg class="nav-arrow" :class="{ open: showCodeSubmenu }" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M6 9l6 6 6-6"/>
                </svg>
              </button>

              <div class="nav-submenu" v-if="showCodeSubmenu">
                <button
                  class="nav-subitem"
                  :class="{ active: currentView === 'code' && activeCodeCategory === 'all' }"
                  @click="selectCodeCategory('all')"
                >
                  All
                </button>
                <button
                  v-for="sub in codeSubcategories"
                  :key="sub.key"
                  class="nav-subitem"
                  :class="{ active: currentView === 'code' && activeCodeCategory === sub.key }"
                  @click="selectCodeCategory(sub.key)"
                >
                  {{ sub.label }}
                </button>
              </div>
            </div>
            <button class="nav-item" @click="currentView = 'read'" :class="{ active: currentView === 'read' }">
              <svg class="nav-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/></svg>
              <span>Read</span>
            </button>
            <button class="nav-item" @click="currentView = 'ranking'" :class="{ active: currentView === 'ranking' }">
              <svg class="nav-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"/></svg>
              <span>Ranking List</span>
            </button>
          </nav>
        </aside>

        <!-- Main Content Area -->
        <main class="main-content">
          <!-- DETAIL PAGE VIEW (real fetch — no hardcoding) -->
          <template v-if="currentView === 'detail'">
            <p v-if="isLoadingDetail" class="state-msg">កំពុងផ្ទុកអត្ថបទ...</p>
            <p v-else-if="detailError" class="state-msg" style="color:#dc2626;">{{ detailError }}</p>
            <ArticleDetail 
              v-else-if="selectedArticle"
              :key="selectedArticle.id"
              :articleData="selectedArticle" 
              @back="currentView = 'home'" 
            />
          </template>
          <!-- RANKING PAGE VIEW (real fetch — no hardcoding) -->
          <template v-else-if="currentView === 'ranking'">
            <div class="content-header">
              <h1>Top of Article</h1>
            </div>

            <div class="articles-list">
              <p v-if="isLoadingRanking" class="state-msg">កំពុងផ្ទុក...</p>
              <p v-else-if="rankingError" class="state-msg" style="color:#dc2626;">{{ rankingError }}</p>
              <p v-else-if="!rankingArticlesList.length" class="state-msg">No ranked articles yet</p>

              <div class="article-card ranking-card-item" v-for="(item, i) in rankingArticlesList" :key="item.id">
                <div class="article-left">
                  <div class="rank-badge" :class="'rank-' + (i + 1)">
                    {{ i + 1 }}
                  </div>

                  <div class="article-info">
                    <h3 class="article-title">{{ item.title }}</h3>
                    <div class="article-meta">
                      <div class="author-box">
                        <div class="avatar sm">
                          <img :src="item.userProfile" :alt="item.username" class="profile-img" />
                        </div>
                        <span class="username">{{ item.username }}</span>
                      </div>
                      <span class="meta-tag">
                        <svg class="meta-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/></svg>
                        {{ item.views }} View
                      </span>
                      <span class="meta-tag">
                        <svg class="meta-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"/></svg>
                        {{ item.comments }} Cmt
                      </span>
                      <span class="meta-tag">
                        <svg class="meta-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/></svg>
                        {{ item.saves }} save
                      </span>
                    </div>
                  </div>
                </div>

                <div class="article-right">
                  <button class="btn-like" @click="item.likes++">
                    <svg class="btn-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/>
                    </svg>
                    <span class="like-count">{{ item.likes }}</span>
                  </button>
                  <button class="btn-create-article" @click="handleCreateArticle">
                    Create Article
                  </button>
                </div>
              </div>
            </div>
          </template>

          <!-- DEFAULT FEEDS VIEW -->
          <template v-else>
            <!-- Search and Create Bar -->
            <h2 class="feed-title-label">{{ feedTitle }}</h2>
            <header class="top-bar">
              <div class="search-box">
                <svg class="search-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
                <input type="text" placeholder="Search articles, topics, or authors..." class="search-input" />
              </div>
              <button class="btn-primary" @click="handleCreateArticle">
                <svg class="btn-icon" fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4"/></svg>
                Create Article
              </button>
            </header>

            <!-- Filter Tabs with SVG Icons -->
            <div class="filter-tabs">
              <button
                @click="activeTab = 'Popular'"
                :class="['tab-item', { active: activeTab === 'Popular' }]">
                <svg class="tab-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/>
                </svg>
                Popular
              </button>
              <button
                @click="activeTab = 'The Latest'"
                :class="['tab-item', { active: activeTab === 'The Latest' }]">
                <svg class="tab-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                The Latest
              </button>
            </div>

            <!-- Article Feed -->
            <div class="feed-list">
              <p v-if="isLoadingArticles" class="state-msg">waiting...</p>
              <p v-else-if="articlesError" class="state-msg" style="color:#dc2626;">{{ articlesError }}</p>
              <p v-else-if="!filteredArticles.length" class="state-msg">No articles in this section yet</p>

              <template v-for="(article, idx) in filteredArticles" :key="article.id">
                <article class="article-card" @click="openArticleDetail(article)" style="cursor: pointer;">
                  <div class="article-body">
                    <div class="article-info">
                      <h3 class="article-heading">{{ article.title }}</h3>
                      <p class="article-text">{{ article.text }}</p>
                      <p class="article-subtext">{{ article.secondaryText }}</p>
                    </div>
                    <div class="article-thumbnail">
                      <img :src="article.image" alt="Thumbnail" />
                    </div>
                  </div>

                  <!-- Author & Metadata -->
                  <div class="article-footer" @click.stop>
                    <div class="author-info">
                      <div class="avatar">
                        <img :src="article.authorProfile" :alt="article.author" class="profile-img" />
                      </div>
                      <span class="author-name">{{ article.author }}</span>
                      <span class="dot-separator">•</span>
                      <span class="stat">
                        <svg class="stat-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/></svg>
                        {{ article.views }}
                      </span>
                      <span class="stat">
                        <svg class="stat-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 00-2-2h-.095c-.5 0-.905.405-.905.905 0 .714-.211 1.412-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 01-2-2v-6a2 2 0 012-2h2.5"/></svg>
                        {{ article.likes }}
                      </span>
                    </div>
                    <div class="options-menu-wrap">
                      <button class="options-btn" @click.stop="toggleOptionsMenu(article.id)">
                        <svg fill="currentColor" viewBox="0 0 24 24"><path d="M6 10c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm12 0c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm-6 0c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>
                      </button>
                      <div class="options-dropdown" v-if="openOptionsMenuId === article.id" @click.stop>
                        <button class="options-dropdown-item" @click="handleNotInterested(article)">
                          <svg class="options-dropdown-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="9"/><line x1="8" y1="8" x2="16" y2="16"/></svg>
                          <span>Not interested</span>
                        </button>
                        <button class="options-dropdown-item" @click="handleBlockUser(article)">
                          <svg class="options-dropdown-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M5.5 5.5l13 13"/></svg>
                          <span>Block user</span>
                        </button>
                        <button class="options-dropdown-item options-dropdown-item-danger" @click="handleReport(article)">
                          <svg class="options-dropdown-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 9v4m0 4h.01M10.29 3.86l-8.18 14.14A2 2 0 004.18 21h15.64a2 2 0 001.87-2.99L13.71 3.86a2 2 0 00-3.42 0z"/></svg>
                          <span>Report</span>
                        </button>
                      </div>
                    </div>
                  </div>

                  <!-- Tags (hidden on Code Life) -->
                  <div class="tag-list" v-if="currentView !== 'code'">
                    <span v-for="(tag, tidx) in article.tags" :key="tidx" class="tag">
                      #{{ tag }}
                    </span>
                  </div>
                </article>

                <!-- Top authors card -->
                <div
                  v-if="currentView === 'home' && idx === authorsCardAfterIndex"
                  class="top-authors-card"
                >
                  <div class="top-authors-header">
                    <span class="top-authors-title">Top authors</span>
                    <svg class="scroll-hint-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                      <circle cx="5" cy="12" r="1.3" fill="currentColor" stroke="none"/>
                      <circle cx="12" cy="12" r="1.3" fill="currentColor" stroke="none"/>
                      <circle cx="19" cy="12" r="1.3" fill="currentColor" stroke="none"/>
                      <path d="M2 9l-2 3 2 3M22 9l2 3-2 3" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>
                  </div>
                  <div class="top-authors-scroll">
                    <div class="author-card" v-for="(a, i) in topAuthors" :key="a.id">
                      <div class="author-card-avatar-wrap">
                        <div class="author-card-avatar">
                          <img :src="a.avatar" :alt="a.name" />
                        </div>
                        <span class="level-badge" :style="{
                          color: a.color,
                          backgroundColor: a.bgColor,
                        }">
                          {{ a.emoji }} Lv.{{ a.level }}
                        </span>
                      </div>
                      <div class="author-card-name">{{ a.name }}</div>
                      <div class="author-card-stats">
                        <span class="author-stat-pill">{{ a.articles }}K article</span>
                        <span class="author-stat-pill">{{ a.followers }} Follower</span>
                      </div>
                      <div class="author-card-rank">Top {{ i + 1 }}</div>
                      <button
                        class="author-card-follow"
                        :class="{ following: a.isFollowing }"
                        @click="toggleFollow(a)"
                      >
                        {{ a.isFollowing ? 'Following ✓' : 'Follow' }}
                      </button>
                    </div>
                  </div>
                </div>
              </template>
            </div>
          </template>
        </main>

        <!-- Right Sidebar (Widgets) -->
        <aside class="sidebar-right">
          <template v-if="currentView === 'detail' && detailAuthorInfo">
            <div class="detail-author-card">
              <div class="detail-author-top">
                <div class="detail-author-avatar">
                  <img v-if="detailAuthorInfo.avatar" :src="detailAuthorInfo.avatar" :alt="detailAuthorInfo.name" />
                  <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
                </div>
                <div class="detail-author-name-badge">
                  <span class="detail-author-name">{{ detailAuthorInfo.name }}</span>
                  <span class="detail-author-level">Level ⏳{{ detailAuthorInfo.level }}</span>
                </div>
              </div>

              <div class="detail-author-stats">
                <span class="detail-stat-pill">{{ detailAuthorInfo.articleCount }}<br/><span class="color-gray">Article</span></span>
                <span class="detail-stat-pill">{{ detailAuthorInfo.likes }}<br/><span class="color-gray">Like</span></span>
                <span class="detail-stat-pill">{{ detailAuthorInfo.followers }}<br/><span class="color-gray">Follower</span></span>
              </div>

              <div class="detail-author-actions">
                <button
                  class="detail-follow-btn"
                  :class="{ following: detailAuthorInfo.isFollowing }"
                  @click="toggleDetailAuthorFollow"
                >
                  {{ detailAuthorInfo.isFollowing ? 'Following ✓' : 'Follow' }}
                </button>
                <button class="detail-more-btn">•••</button>
              </div>
            </div>

            <div class="detail-toc-card" v-if="detailTocItems.length">
              <div class="detail-toc-header">
                <svg class="toc-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="4" y="3" width="16" height="18" rx="2"/><line x1="8" y1="8" x2="16" y2="8"/><line x1="8" y1="12" x2="16" y2="12"/><line x1="8" y1="16" x2="13" y2="16"/></svg>
                <span>Table of contents</span>
                <svg class="toc-expand-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M7 17 17 7M17 7H9M17 7v8"/></svg>
              </div>
              <ul class="detail-toc-list">
                <li v-for="(item, i) in detailTocItems" :key="i" class="detail-toc-item">
                  <span class="detail-toc-num">{{ toRoman(i + 1) }}.</span>
                  <span class="detail-toc-text">{{ item }}</span>
                </li>
              </ul>
            </div>

            <!-- List Article Widget -->
            <div class="widget-card">
              <div class="widget-title">
                <span class="widget-title-flex">
                  <svg class="widget-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"/></svg>
                  List Article
                </span>
                <svg class="external-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/></svg>
              </div>
              <div class="widget-items">
                <div v-for="item in listArticles" :key="item.id" class="list-row">
                  <span class="rank-number list-row-hover">{{ item.id }}</span>
                  <span class="row-text list-row-hover">{{ item.title }}</span>
                </div>
              </div>
            </div>

            <!-- Recommended Authors Widget -->
            <div class="widget-card">
              <div class="widget-title">
                <span class="widget-title-flex">
                  <svg class="widget-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"/></svg>
                  Recommended Authors
                </span>
              </div>
              <div class="widget-items">
                <div v-for="author in authors" :key="author.id" class="author-row">
                  <div class="author-profile">
                    <div class="avatar sm">
                      <img :src="author.avatar" :alt="author.name" class="profile-img" />
                    </div>
                    <span class="row-text">{{ author.name }}</span>
                  </div>
                  <button class="btn-follow">Follow</button>
                </div>
              </div>
            </div>

            <!-- Top Topics Widget -->
            <div class="widget-card">
              <div class="widget-title">
                <span class="widget-title-flex">
                  <svg class="widget-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"/></svg>
                  Top Topics
                </span>
                <svg class="external-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/></svg>
              </div>
              <div class="widget-items">
                <div v-for="topic in topics" :key="topic.id" class="topic-row">
                  <span class="row-text"># {{ topic.name }}</span>
                  <div class="topic-img-wrapper">
                    <img :src="topic.image" :alt="topic.name" class="profile-img" />
                  </div>
                </div>
              </div>
            </div>
          </template>

          <!-- DEFAULT: List Article / Authors / Topics widgets -->
          <template v-else>
            <!-- List Article Widget -->
            <div class="widget-card">
              <div class="widget-title">
                <span class="widget-title-flex">
                  <svg class="widget-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"/></svg>
                  List Article
                </span>
                <svg class="external-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/></svg>
              </div>
              <div class="widget-items">
                <div v-for="item in listArticles" :key="item.id" class="list-row">
                  <span class="rank-number list-row-hover">{{ item.id }}</span>
                  <span class="row-text list-row-hover">{{ item.title }}</span>
                </div>
              </div>
            </div>

            <!-- Authors Widget -->
            <div class="widget-card">
              <div class="widget-title">
                <span class="widget-title-flex">
                  <svg class="widget-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"/></svg>
                  Recommended Authors
                </span>
              </div>
              <div class="widget-items">
                <div v-for="author in authors" :key="author.id" class="author-row">
                  <div class="author-profile">
                    <div class="avatar sm">
                      <img :src="author.avatar" :alt="author.name" class="profile-img" />
                    </div>
                    <span class="row-text">{{ author.name }}</span>
                  </div>
                  <button class="btn-follow">Follow</button>
                </div>
              </div>
            </div>

            <!-- Top Topics Widget -->
            <div class="widget-card">
              <div class="widget-title">
                <span class="widget-title-flex">
                  <svg class="widget-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"/></svg>
                  Top Topics
                </span>
                <svg class="external-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/></svg>
              </div>
              <div class="widget-items">
                <div v-for="topic in topics" :key="topic.id" class="topic-row">
                  <span class="row-text"># {{ topic.name }}</span>
                  <div class="topic-img-wrapper">
                    <img :src="topic.image" :alt="topic.name" class="profile-img" />
                  </div>
                </div>
              </div>
            </div>
          </template>

        </aside>

      </div>
    </div>

    <!--CreateArticle -->
    <div v-if="showCreatePage" class="create-page-overlay">
      <CreateArticle @close="showCreatePage = false" @submit="submitArticle" />
    </div>
  </div>
</template>

<style scoped>
/* Reset & Global Themes */
.app-container {
  min-height: 100vh;
  background-color: #F7F4F2;
  color: #1e293b;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  display: flex;
  justify-content: center;
  padding-top: 12px;


  align-items: flex-start;
}

.dashboard-layout {
  width: 100%;
  max-width: 1251px;
  display: grid;
  grid-template-columns: 240px 1fr 280px;
  gap: 16px;
  align-items: start;
}

.sidebar-left, .widget-card, .article-card {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
}

.sidebar-left {
  padding: 16px;
  height: fit-content;
  display: flex;
  flex-direction: column;
  gap: 16px;
  position: sticky;
  
  top: 80px;
  max-height: calc(100vh - 92px);
  overflow-y: auto;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.brand-logo {
  display: flex;
  align-items: center;
  gap: 10px;
  padding-bottom: 6px;
  border-bottom: 1px solid #f1f5f9;
}

.brand-logo h2 {
  font-size: 20px;
  font-weight: 700;
  color: #000;
  letter-spacing: -0.5px;
  margin: 0px;
}

.nav-menu {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
  padding: 10px 12px;
  border-radius: 10px;
  border: none;
  background: transparent;
  color: #000;
  font-weight: 500;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s ease;
  text-align: left;
}

.nav-item:hover {
  background-color: #0000000e;
  color: #0f172a;
}

.nav-item.active {
  /* background-color: #e3f2fd; */
  color: #1B75D2;
  font-weight: 600;
}

.nav-icon {
  width: 19px;
  height: 19px;
  flex-shrink: 0;
}

.nav-item .arrow {
  width: 16px;
  height: 16px;
  margin-left: auto;
}

.main-content {
    min-width: 0;
  background: transparent;
  border: none;
  box-shadow: none;
  display: flex;
  flex-direction: column;
  gap: 20px;
  background-color: #ffffff;
  padding: 12px;
  margin-top: -12px;
}

.feed-title-label {
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
  margin: 4px 0 0;
}

.content-header h1 {
  margin: 0px;
  font-size: 20px;
  font-weight: 700;
  color: #0f172a;
}

.top-bar {
  padding: 8px 0px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: transparent;
  border: none;
}

.search-box {
  position: relative;
  width: 70%;
}

.search-icon {
  position: absolute;
  left: 16px;
  top: 50%;
  transform: translateY(-50%);
  width: 18px;
  height: 18px;
  color: #94a3b8;
}

.search-input {
  width: 100%;
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 9999px;
  padding: 10px 16px 10px 44px;
  color: #0f172a;
  font-size: 14px;
  outline: none;
  transition: border-color 0.2s;
}

.search-input:focus {
  border-color: #1B75D2;
  background-color: #ffffff;
}

.btn-primary {
  background-color: #1B75D2;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 32px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: background-color 0.2s, transform 0.1s;
}

.btn-icon {
  width: 18px;
  height: 18px;
}

.btn-primary:hover {
  background-color: #155fa8;
}

.btn-primary:active {
  transform: scale(0.98);
}

/* Filter Tabs & SVG Icon Styling */
.filter-tabs {
  display: flex;
  gap: 32px;
  border-bottom: 1px solid #e2e8f0;
  padding-bottom: 0px;
  margin-top: 8px;
}

.tab-item {
  background: transparent;
  border: none;
  padding: 8px 0px 12px 0px;
  color: #64748b;
  font-weight: 500;
  font-size: 15px;
  cursor: pointer;
  position: relative;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: color 0.2s;
}

.tab-icon {
  width: 18px;
  height: 18px;
}

.tab-item:hover {
  color: #0f172a;
}

.tab-item.active {
  color: #1B75D2;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: -1px;
  left: 0px;
  width: 100%;
  height: 2px;
  background-color: #1B75D2;
  border-radius: 2px 2px 0px 0px;
}

.feed-list, .articles-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.state-msg {
  text-align: center;
  color: #64748b;
  font-size: 14px;
  padding: 24px 0;
}

.article-card {
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.article-card:hover{
  background-color: #cccccc1e;
}

.article-body {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 16px;
}

.article-info {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.article-heading, .article-title {
  font-size: 17px;
  font-weight: 600;
  color: #0f172a;
  margin: 0px;
}

.article-text {
  font-size: 14px;
  color: #334155;
  margin: 0px;
}

.article-subtext {
  font-size: 13px;
  color: #64748b;
  margin: 0px;
}

.article-thumbnail {
  width: 80px;
  height: 80px;
  border-radius: 12px;
  overflow: hidden;
  background-color: #f1f5f9;
  flex-shrink: 0;
}

.article-thumbnail img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.article-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-top: 1px solid #f1f5f9;
  padding-top: 14px;
}

.author-info {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  color: #64748b;
}

.avatar {
  width: 32px;
  height: 32px;
  background-color: #f1f5f9;
  border: 1px solid #e2e8f0;
  border-radius: 50%;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.profile-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar.sm {
  width: 28px;
  height: 28px;
}

.author-name {
  color: #1e293b;
  font-weight: 500;
}

.dot-separator {
  color: #cbd5e1;
}

.stat {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #64748b;
  font-size: 13px;
}

.stat-icon {
  width: 16px;
  height: 16px;
}

.options-btn {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  color: #64748b;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 8px;
  transition: all 0.2s ease;
}

.options-btn svg {
  width: 20px;
  height: 20px;
}

.options-btn:hover {
  background-color: #f8fafc;
  border-color: #cbd5e1;
  color: #0f172a;
}

.tag-list {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.tag {
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  padding: 3px 10px;
  border-radius: 9999px;
  font-size: 12px;
  color: #1976D2;
}

/* Ranking Card Specific Styles */
.ranking-card-item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

.article-left {
  display: flex;
  align-items: center;
  gap: 16px;
  flex: 1;
}

.rank-badge {
  width: 36px;
  height: 36px;
  color: #475569;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 16px;
  flex-shrink: 0;
}

.article-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 13px;
  color: #64748b;
  flex-wrap: wrap;
  margin-top: 4px;
}

.author-box {
  display: flex;
  align-items: center;
  gap: 6px;
  background: #f8fafc;
  padding: 2px 8px 2px 2px;
  border-radius: 9999px;
  border: 1px solid #e2e8f0;
}

.username {
  font-weight: 600;
  color: #334155;
}

.meta-tag {
  display: flex;
  align-items: center;
  gap: 4px;
  background: transparent;
  border: none;
  padding: 0px;
  font-weight: 500;
  color: #64748b;
}

.meta-icon {
  width: 14px;
  height: 14px;
  color: #64748b;
}

.article-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.btn-like {
  display: flex;
  align-items: center;
  gap: 6px;
  background-color: #1B75D2;
  color: #ffff;
  padding: 7px 12px;
  border-radius: 32px;
  font-weight: 700;
  font-size: 12.5px;
  cursor: pointer;
  transition: background-color 0.2s;
  border: 2px solid #1976D2;
}

.btn-like:hover {
  background-color: #155fa8;
}

.btn-like svg {
  color: #ffffff;
  width: 18px;
  stroke-width: 1.8;
  height: 18px;
}

.btn-create-article {
  background-color: #1B75D2;
  color: #ffffff;
  padding: 7px 12px;
  border-radius: 32px;
  font-weight: 700;
  font-size: 12.5px;
  cursor: pointer;
  transition: background-color 0.2s;
  border: 2px solid #1976D2;
}

.btn-create-article:hover {
  background-color: #155fa8;
}

/* Right Sidebar */
.sidebar-right {
  display: flex;
  flex-direction: column;
  gap: 20px;
  height: fit-content;
  position: sticky;
top: 80px;


max-height: calc(100vh - 92px);
  overflow-y: auto;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.widget-card {
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.widget-title {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 15px;
  font-weight: 600;
  color: #0f172a;
  padding-bottom: 8px;
  border-bottom: 1px solid #f1f5f9;
}

.widget-title-flex {
  display: flex;
  align-items: center;
  gap: 8px;
}

.widget-icon {
  width: 18px;
  height: 18px;
  color: #000;
}

.external-icon {
  width: 16px;
  height: 16px;
  color: #000;
  cursor: pointer;
}

.widget-items {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.list-row, .author-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 12px;
  border-radius: 12px;
  font-size: 14px;
  transition: all 0.2s ease;
  cursor: pointer;
}

.topic-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 12px;
  border-radius: 12px;
  font-size: 14px;
  transition: all 0.2s ease;
  cursor: pointer;
  color: #1976D2;
}

.list-row {
  gap: 12px;
  color: #1976D2;

}

.list-row-hover:hover{
  color: #1976D2;
}

/* Hover effect for all rows in widgets */
 .author-row:hover, .topic-row:hover {
  background-color: #f3f4f5;
  border-color: #cbd5e1;
  border-radius: 32px;
}

.rank-number {
  width: 14px;
  text-align: center;
  font-size: 13px;
  font-weight: 600;
  color: #64748b;
  flex-shrink: 0;
}

.row-text {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  flex: 1;
}

.author-profile {
  display: flex;
  align-items: center;
  gap: 8px;
}

.btn-follow {
  background-color: #1B75D2;
  border: 1px solid #1B75D2;
  color: #ffffff;
  padding: 5px 14px;
  border-radius: 9999px;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-follow:hover {
  background-color: #155fa8;
  border-color: #155fa8;
  color: #ffffff;
}

.topic-img-wrapper {
  width: 26px;
  height: 26px;
  border-radius: 6px;
  overflow: hidden;
  border: 1px solid #e2e8f0;
  background-color: #ffffff;
  flex-shrink: 0;
}

.top-authors-card {
  background: #FFFFFF;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  padding: 14px 16px 18px;
}

.top-authors-scroll::-webkit-scrollbar {
  display: none;
}

.top-authors-scroll {
  scrollbar-width: none;
}

.top-authors-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
}

.top-authors-title {
  color: #000;
  font-weight: 700;
  font-size: 16px;
  /* border: 1px solid #e2e8f0; */
  border-radius: 10px;
  padding: 6px 14px;
}

.scroll-hint-icon {
  width: 26px;
  height: 26px;
  color: #000;
}

.top-authors-scroll {
  display: flex;
  gap: 12px;
  overflow-x: auto;
  overflow-y: hidden;
  flex-wrap: nowrap;
  width: 100%;
  -webkit-overflow-scrolling: touch;
}

.top-authors-scroll::-webkit-scrollbar {
  height: 6px;
}
.top-authors-scroll::-webkit-scrollbar-thumb {
  background: #FFFFFF;
  border-radius: 999px;
}

.author-card {
  flex: 0 0 200px;
  scroll-snap-align: start;
  background: #FFFFFF;
  border: 1px solid #e2e8f0;
  border-radius: 18px;
  padding: 16px 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.author-card-avatar-wrap {
  position: relative;
  width: 74px;
  height: 74px;
}

.author-card-avatar img{
  width: 100%;
  height: 100%;
}


.author-card-avatar {
  width: 74px;
  height: 74px;
  border-radius: 50%;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.author-card-avatar svg {
  width: 34px;
  height: 34px;
  border: 1px solid #e2e8f0;
  fill: none;
  stroke-width: 1.8;
}

.level-badge {
  position: absolute;
  top: -6px;
  right: -14px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  color: #000;
  font-size: 10px;
  font-weight: 700;
  padding: 2px 8px;
  border-radius: 999px;
  display: flex;
  align-items: center;
  gap: 3px;
}

.level-emoji { font-size: 10px; }

.author-card-name {
  color: #000;
  font-weight: 700;
  font-size: 15px;
  border-radius: 8px;
  padding: 4px 14px;
  width: 100%;
  text-align: center;
}

.author-card-stats {
  display: flex;
  gap: 6px;
  width: 100%;
}

.author-stat-pill {
  flex: 1;
  text-align: center;
  color: #000;
  font-size: 10.5px;
  font-weight: 600;
  border: 1px solid #e2e8f0;
  background-color: #F9FAFC;
  border-radius: 8px;
  padding: 3px 4px;
  white-space: nowrap;
}

.author-card-rank {
  color: #000;
  font-weight: 700;
  font-size: 14px;
  background-color: #F9FAFC;
  border-radius: 8px;
  padding: 3px 12px;
  width: 100%;
  text-align: center;
}

.author-card-follow {
  width: 100%;
  background: #1976D2;
  border: 1px solid #e2e8f0;
  color: #ffffff;
  font-weight: 700;
  font-size: 13px;
  padding: 6px;
  border-radius: 32px;
  cursor: pointer;
  transition: background 0.15s ease;
  border-color: #1976D2;
}

.author-card-follow.following {
  background: #fff;
  color: #1976D2;
  border-color: #1976D2;
}

.author-card-follow:hover {
   background-color: #155fa8;
  border-color: #155fa8;
  color: #ffffff;
}

.create-page-overlay {
  position: absolute;
  top: 60px;
  left: 0;
  width: 100%;
  height: calc(100% - 60px);
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1000;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding-top: 12px;
  overflow-y: auto;
}



/* Detail view: Author card + TOC — navy theme like the mockup */
.detail-author-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 18px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.detail-author-top {
  display: flex;
  align-items: center;
  gap: 12px;
}

.detail-author-avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background: #2E6FD9;
  border: 2px solid #4A8FE8;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  flex-shrink: 0;
}

.detail-author-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.detail-author-avatar svg {
  width: 26px;
  height: 26px;
  stroke: #fff;
  fill: none;
  stroke-width: 1.8;
}

.detail-author-name-badge {
  display: flex;
  flex-direction: column;
  flex: 1;
  min-width: 0;
}

.detail-author-name {
  color: #000;
  font-weight: 800;
  font-size: 14px;
  padding: 2px 1px;
  border-radius: 10px;
}

.detail-author-level {
  background: #ffffff;
  color: #1B75D2;
  font-weight: 700;
  font-size: 12px;
  padding: 2px 1px;
  width: fit-content;

}

.detail-author-stats {
  display: flex;
  gap: 8px;
}

.color-gray{
  color: #ccc;
}

.detail-stat-pill {
  flex: 1;
  text-align: center;
  background: #ffffff;
  color: #000;
  font-size: 12px;
  font-weight: 700;
  line-height: 1.5;
  padding: 8px 4px;
  border-radius: 10px;
}

.detail-author-actions {
  display: flex;
  gap: 8px;
}

.detail-follow-btn {
  flex: 1;
  background: #1976D2;
  border: none;
  color: #ffffff;
  font-weight: 700;
  font-size: 14px;
  padding: 8px;
  border-radius: 32px;
  cursor: pointer;
  transition: opacity 0.15s ease;
}

.detail-follow-btn.following {
  background: #ffffff;
  color: #1976D2;
  border: 1px solid #1976D2;
}

.detail-follow-btn:hover {
  opacity: 0.85;
}

.detail-more-btn {
  background: #ffff;
  border: 1px solid #e2e8f0;
  color: #000;
  font-weight: 700;
  padding: 8px 14px;
  border-radius: 32px;
  cursor: pointer;
}

.detail-toc-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 18px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.detail-toc-header {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #000;
  font-weight: 700;
  font-size: 14px;
}

.toc-icon {
  width: 18px;
  height: 18px;
  flex-shrink: 0;
}

.toc-expand-icon {
  width: 16px;
  height: 16px;
  margin-left: auto;
  cursor: pointer;
}

.detail-toc-list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.detail-toc-item {
  display: flex;
  align-items: center;
  gap: 4px;
  background: #ffff;
  border-radius: 10px;
  padding: 7px 12px;
}

.detail-toc-num {
  color: #000;
  font-size: 11px;
  font-weight: 700;
  width: 18px;
  height: 18px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.detail-toc-text {
  color: #1976D2;
  font-size: 12.5px;
  font-weight: 500;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}



.nav-item-wrap {
  display: flex;
  flex-direction: column;
}

.nav-arrow {
  width: 14px;
  height: 14px;
  margin-left: auto;
  flex-shrink: 0;
  transition: transform 0.2s ease;
}

.nav-arrow.open {
  transform: rotate(180deg);
}

.nav-submenu {
  display: flex;
  flex-direction: column;
  gap: 2px;
  padding-left: 34px;
  margin-top: 2px;
  margin-bottom: 4px;
}

.nav-subitem {
  text-align: left;
  border: none;
  background: transparent;
  padding: 7px 10px;
  border-radius: 8px;
  font-size: 13px;
  color: #64748b;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.15s ease;
}

.nav-subitem:hover {
  background-color: #f1f5f9;
  color: #0f172a;
}

.nav-subitem.active {
  color: #1B75D2;
  font-weight: 600;
}


.options-menu-wrap {
  position: relative;
}

.options-dropdown {
  position: absolute;
  top: 34px;
  right: 0;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .12);
  padding: 6px;
  min-width: 180px;
  z-index: 50;
}

.options-dropdown-item {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  text-align: left;
  border: none;
  background: transparent;
  padding: 9px 10px;
  border-radius: 8px;
  font-size: 13.5px;
  font-weight: 500;
  color: #1e293b;
  cursor: pointer;
  transition: background 0.15s ease;
}

.options-dropdown-item:hover {
  background-color: #f1f5f9;
}

.options-dropdown-item-danger {
  color: #dc2626;
}

.options-dropdown-item-danger:hover {
  background-color: #fef2f2;
}

.options-dropdown-icon {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

/* Responsive Design */
@media (max-width: 1100px) {
  .dashboard-layout {
    grid-template-columns: 1fr;
  }
  .sidebar-left, .sidebar-right {
    position: static;
  }
}

</style>