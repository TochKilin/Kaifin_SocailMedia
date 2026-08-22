<template>
  <div class="main-ctn">
    <NavBar/>
    
    <div class="quote-app">
      
      <!-- Left Sidebar Area -->
      <aside class="sidebar-left">
        
        <!-- Search & Create Section Container -->
        <div class="sidebar-widget search-widget">
          <div class="search-box">
            <svg class="search-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
            <input type="text" placeholder="Search ....." v-model="searchQuery" />
          </div>

          <div class="create-btn-wrap">
            <button
            class="btn-create"
            @click.stop="openCreateQuote"
          >
            <span class="icon-circle-create">+</span>
            Create Quotes
          </button>
          </div>
        </div>

        <!-- Top Quote Widget -->
        <div class="top-quote-card-post">
          <div class="sidebar-header">
            <h2>Top Quote</h2>
          </div>
          <ul class="top-list">
            <li v-for="(item, idx) in topQuotes" :key="idx">
              <span class="rank-num">{{ idx + 1 }}</span>
              <span class="top-title">{{ item }}</span>
            </li>
          </ul>
        </div>

      </aside>

      <main class="main-content">

        <nav class="tabs-nav">
          <button 
            @click="activeTab = 'popular'" 
            :class="['tab-btn', { active: activeTab === 'popular' }]">
            Popular
          </button>
          <button 
            @click="activeTab = 'latest'" 
            :class="['tab-btn', { active: activeTab === 'latest' }]">
            The Latest
          </button>
        </nav>

        <div v-if="loading" class="state-msg">waiting...</div>
        <div v-else-if="errorMsg" class="state-msg error">{{ errorMsg }}</div>
        <div v-else-if="filteredQuotes.length === 0" class="state-msg">មិនមាន quote ទេ</div>

        <!-- Quotes Feed -->
        <div class="quotes-feed" v-else>
          <div class="quote-card-post" v-for="(quote, index) in filteredQuotes" :key="quote.id ?? index">
            
            <div class="quote-header-text">
              <div class="quote-badge-icon">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor"><path d="M14.017 21v-7.391c0-5.704 3.731-9.57 8.983-10.609l.995 2.151c-2.432.917-3.995 3.638-3.995 5.849h4v10h-9.983zm-14.017 0v-7.391c0-5.704 3.748-9.57 9-10.609l.996 2.151c-2.433.917-3.996 3.638-3.996 5.849h3.998v10h-9.998z"/></svg>
              </div>
              <div class="quote-content-block">
                <h3>{{ quote.title }}</h3>
                <p v-html="quote.text"></p>
              </div>
            </div>
            
            <div class="quote-footer">
              <div class="user-profile">
              <div class="avatar">
                <img
                  v-if="quote.userAvatarUrl"
                  :src="quote.userAvatarUrl"
                  class="avatar-img"
                  alt="avatar"
                />

                <svg
                  v-else
                  width="16"
                  height="16"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                >
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                  <circle cx="12" cy="7" r="4"></circle>
                </svg>
              </div>
              <span class="username">{{ quote.username }}</span>
            </div>

              <div class="quote-actions">
                <button class="action-pill">
                  <span class="icon-circle"><svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path><circle cx="12" cy="12" r="3"></circle></svg></span>
                  <span>{{ quote.views }} View</span>
                </button>
                
                <div class="like-container" @mouseleave="quote.showReactions = false">
                  <button 
                  class="action-pill like-btn" 
                  :class="{ 'is-active': quote.selectedReaction }"
                  @click="handleLikeClick(quote)"
                  @mouseenter="quote.showReactions = true"
                  :disabled="quote.reactionLoading"
                >
                  <span class="icon-circle" v-if="!quote.selectedReaction">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path></svg>
                  </span>
                  
                  <span class="reaction-display-label">
                    <template v-if="quote.selectedReaction">
                      <span v-if="quote.selectedReaction.type === 'emoji'" class="selected-emoji">{{ quote.selectedReaction.content }}</span>
                      <span v-else-if="quote.selectedReaction.type === 'image'" class="selected-img-wrap">
                        <img :src="quote.selectedReaction.content" class="selected-img" />
                      </span>
                      <span v-else-if="quote.selectedReaction.type === 'svg'" class="selected-svg-wrap" v-html="quote.selectedReaction.content"></span>
                    </template>
                    <template v-else>
                      Like
                    </template>
                  </span>
                  <span v-if="quote.likes > 0" class="like-count">{{ quote.likes }}</span>
                </button>

                  <QuoteReaction 
                    v-if="quote.showReactions" 
                    :reactions="reactions" 
                    @close="quote.showReactions = false" 
                    @select="selectReaction(quote, $event)" 
                  />
                </div>

                <div class="share-container" @mouseleave="quote.showShare = false">
                  <button 
                    class="icon-btn" 
                    title="Share"
                    @click.stop="quote.showShare = !quote.showShare"
                  >
                    <span class="icon-circle"><svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path><polyline points="16 6 12 2 8 6"></polyline><line x1="12" y1="2" x2="12" y2="15"></line></svg></span>
                  </button>

                  <QuoteShare 
                    v-if="quote.showShare" 
                    @close="quote.showShare = false" 
                    @action="handleShareAction(quote, $event)" 
                  />
                </div>

                <div class="options-container" @mouseleave="quote.showOptions = false">
                  <button 
                    class="icon-btn dots-btn" 
                    title="More"
                    @click.stop="quote.showOptions = !quote.showOptions"
                  >
                    <span>•••</span>
                  </button>

                  <QuoteOptions 
                    v-if="quote.showOptions" 
                    @close="quote.showOptions = false" 
                    @action="handleOptionAction(quote, $event)" 
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
    <Teleport to="body">
      <Transition name="quote-modal">
        <div
          v-if="showCreateQuote"
          class="quote-overlay"
          @click.self="closeCreateQuote"
        >
          <CreateQuote
            :avatar-url="avatarUrl"
            @close="closeCreateQuote"
            @post-quote="handlePostQuote"
            @save-draft="handleSaveDraft"
          />
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue'
import axios from 'axios'
import NavBar from '../navbar/NavBar.vue'
import CreateQuote from './CreateQuote.vue'
import QuoteReaction from './QuoteReaction.vue'
import QuoteShare from './QuoteShare.vue'
import QuoteOptions from './QuoteOptions.vue'

const SERVER_ROOT = import.meta.env.VITE_API_URL
const API_BASE = `${SERVER_ROOT}/api/v1/front`

const api = axios.create({ baseURL: API_BASE })

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

const getQuotes = (tab, { q = '' } = {}) =>
  api.get('/quotes/show', { params: { tab, q } })

const createQuote = (payload) =>
  api.post('/quotes/create', payload)

const getProfile = () => api.get('/profile/show')

const listReactionTypes = () => api.get('/reaction-types')
const reactToQuote = (quoteId, reactionTypeId) =>
  api.post('/quote-reactions/create', { quote_id: quoteId, reaction_type_id: reactionTypeId })
const unreactQuote = (quoteId) =>
  api.delete(`/quote-reactions/${quoteId}`)

const searchQuery = ref('')
const activeTab = ref('popular')
const showCreateModal = ref(false)
const loading = ref(false)
const errorMsg = ref('')
const profile = ref(null)

watch(showCreateModal, (isOpen) => {
  document.body.style.overflow = isOpen ? 'hidden' : ''
})

onBeforeUnmount(() => {
  document.body.style.overflow = ''
})


const reactions = ref([]) 

const normalizeReactionType = (rt) => ({
  id: rt.id,
  name: rt.name,
  type: rt.icon_type,      
  content: rt.icon_value,  
  sortOrder: rt.sort_order
})

const fetchReactionTypes = async () => {
  try {
    const res = await listReactionTypes()
    const list = res.data.data || []
    reactions.value = list.map(normalizeReactionType)
  } catch (err) {
    console.error('Failed to fetch reaction types:', err)
  }
}

const defaultReaction = computed(() =>
  reactions.value.find(r => r.name?.toLowerCase() === 'like') || reactions.value[0] || null
)

const quotes = ref([])

const topQuotes = ref([
  'Title Quote - Inspiration',
  'Title Quote - Success',
  'Title Quote - Coding Life',
  'Title Quote - Wisdom',
  'Title Quote - Future Tech'
])

const normalizeQuote = (q) => ({
  id: q.id,
  title: q.title,
  text: q.content,
  username: q.username || 'Unknown',
  userAvatarUrl: q.profile_images ? `${SERVER_ROOT}/uploads/${q.profile_images}` : null,
  views: q.views_count ?? 0,
  likes: q.likes_count ?? 0,
  showReactions: false,
  showShare: false,
  showOptions: false,
  selectedReaction: q.my_reaction_type_id != null
    ? reactions.value.find(r => r.id === q.my_reaction_type_id) || null
    : null,
  reactionLoading: false
})

const fetchQuotes = async (tab) => {
  loading.value = true
  errorMsg.value = ''
  try {
    const res = await getQuotes(tab, { q: searchQuery.value })
    const payload = res.data.data || {}
    const list = payload.quotes || payload.Quotes || []
    console.log('RAW QUOTES:', list.map(q => ({ id: q.id, title: q.title, my_reaction_type_id: q.my_reaction_type_id })))
    quotes.value = list.map(normalizeQuote)
    trackViewsForList(list)
    console.log('NORMALIZED:', quotes.value) 
  } catch (err) {
    console.error('Failed to fetch quotes:', err)
    errorMsg.value = 'មិនអាចទាញ quote បានទេ សូមព្យាយាមម្តងទៀត'
  } finally {
    loading.value = false
  }
}

const fetchProfile = async () => {
  try {
    const res = await getProfile()
    const p = res.data.data
    profile.value = {
      ...p,
      avatarUrl: p.profile_images ? `${SERVER_ROOT}/uploads/${p.profile_images}` : null
    }
  } catch (err) {
    console.error('Failed to fetch profile:', err)
  }
}

onMounted(async() => {
  await fetchReactionTypes() 
  fetchQuotes(activeTab.value)
  fetchProfile()
  // fetchReactionTypes()
})

watch(activeTab, (tab) => fetchQuotes(tab))

let searchTimer = null
watch(searchQuery, () => {
  clearTimeout(searchTimer)
  searchTimer = setTimeout(() => fetchQuotes(activeTab.value), 400)
})

const filteredQuotes = computed(() => quotes.value)

const shareQuoteApi = axios.create({ baseURL: `${SERVER_ROOT}/api/v1` })
shareQuoteApi.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

const repostQuoteToFeed = (quoteId) =>
  shareQuoteApi.post('/quote-shares', { quote_id: quoteId, channel: 'feed' })

const handleShareAction = async (quote, actionType) => {
  quote.showShare = false
  if (actionType === 'copy') {
    navigator.clipboard.writeText(window.location.href)
    alert('Link copied to clipboard!')
  } else if (actionType === 'feed') {
    try {
      await repostQuoteToFeed(quote.id)
      alert('share sucess')
    } catch (err) {
      console.error('Failed to repost quote to feed:', err)
      alert('Share to feed failed, try again')
    }
  } else {
    console.log(`Action selected: ${actionType} for quote:`, quote.title)
  }
}
const handleLikeClick = async (quote) => {
  if (quote.reactionLoading || !defaultReaction.value) return
  quote.reactionLoading = true

  try {
    if (!quote.selectedReaction) {
      await reactToQuote(quote.id, defaultReaction.value.id)
      quote.selectedReaction = defaultReaction.value
      quote.likes++
    } else {
      await unreactQuote(quote.id)
      quote.selectedReaction = null
      quote.likes--
    }
  } catch (err) {
    console.error('Failed to toggle reaction:', err)
    alert('Reaction failed, please try again.')
  } finally {
    quote.reactionLoading = false
  }
}

const selectReaction = async (quote, reaction) => {
  if (quote.reactionLoading) return
  quote.reactionLoading = true
  quote.showReactions = false

  const hadReactionBefore = !!quote.selectedReaction

  try {
    await reactToQuote(quote.id, reaction.id)
    quote.selectedReaction = reaction
    if (!hadReactionBefore) {
      quote.likes++
    }

  } catch (err) {
    console.error('Failed to select reaction:', err)
    alert('Reaction failed, please try again.')
  } finally {
    quote.reactionLoading = false
  }
}

const handleNewQuote = async (newQuote) => {
  try {
    await createQuote({
      title: newQuote.title,
      content: newQuote.text,
      visibility: newQuote.visibility
    })
    await fetchQuotes(activeTab.value)
  } catch (err) {
    console.error('Failed to create quote:', err)
    const msg = err.response?.data?.message || 'Post quote failed, please try again.'
    alert(msg)
  } finally {
    showCreateModal.value = false
  }
}

const handleSaveDraft = async (draft) => {
  try {
    await createQuote({
      title: draft.name,
      content: draft.content,
      status: 'draft'
    })
  } catch (err) {
    console.error('Failed to save draft:', err)
  } finally {
    showCreateModal.value = false
  }
}

const handleOptionAction = (quote, actionType) => {
  quote.showOptions = false
  if (actionType === 'not_interested') {
    quotes.value = quotes.value.filter(q => q !== quote)
  } else if (actionType === 'block_user') {
    alert(`Blocked user: ${quote.username}`)
  } else if (actionType === 'report') {
    alert(`Reported quote: ${quote.title}`)
  }
}


const viewApi = axios.create({ baseURL: `${SERVER_ROOT}/api/v1` })
viewApi.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})


const trackQuoteView = (quoteId) =>
  viewApi.post('/quote-views', { quote_id: quoteId })

const trackedViewIds = new Set()

const trackViewsForList = (list) => {
  list.forEach((q) => {
    if (trackedViewIds.has(q.id)) return
    trackedViewIds.add(q.id)
    trackQuoteView(q.id).catch((err) => {
      console.error('Failed to track view for quote', q.id, err)
      trackedViewIds.delete(q.id) 
    })
  })
}

const showCreateQuote = ref(false)

const openCreateQuote = () => {
  showCreateQuote.value = true
}

const closeCreateQuote = () => {
  showCreateQuote.value = false
}

const handlePostQuote = async (quote) => {
  try {
    await createQuote({
      title: quote.title,
      content: quote.text,
      visibility: quote.visibility
    })

    await fetchQuotes(activeTab.value)
    showCreateQuote.value = false
    alert('Quote created successfully!')
  } catch (err) {
    console.error('Create quote error:', err)

    console.error('Backend response:', err.response?.data)

    const message =
      err.response?.data?.message ||
      'Post quote failed, please try again.'

    alert(message)
  }
}
</script>

<style scoped>
.quote-app {
  display: grid;
  grid-template-columns: 280px minmax(0, 1fr);
  gap: 20px;
  max-width: 1031px;
  margin: 0 auto;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  min-height: 100vh;
  color: #0f172a;
  box-sizing: border-box;
  padding-top: 0;
}

.sidebar-left {
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-self: start;
  position: sticky;
  top: 76px;
  width: 100%;
}

.search-widget {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 14px;
  padding: 14px;
  display: flex;
  flex-direction: column;
  gap: 12px;

}

.search-box {
  position: relative;
  width: 100%;
}

.modal-overlay{
  position: fixed;
  inset: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

.search-icon {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: #64748b;
}

.search-box input {
  width: 100%;
  padding: 8px 12px 8px 36px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 32px;
  font-size: 13px;
  color: #0f172a;
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
  box-sizing: border-box;
}

.search-box input:focus {
  border-color: #1976d2;
}

.create-btn-wrap {
  position: relative;
}

.btn-create {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  width: 100%;
  background-color: #1976d2;
  color: #ffffff;
  border: none;
  padding: 8px 14px;
  border-radius: 32px;
  font-weight: 600;
  font-size: 13.5px;
  cursor: pointer;
  transition: background-color 0.2s ease, transform 0.1s ease;
}

.btn-create:hover {
  background-color: #1565c0;
}

.btn-create:active {
  transform: scale(0.99);
}

.icon-circle-create {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 23px;
  height: 23px;
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  font-weight: bold;
}

.top-quote-card-post {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 14px;
  padding: 14px;
}

.sidebar-header {
  display: flex;
  align-items: center;
  gap: 8px;
  border-bottom: 1px solid #f1f5f9;
  padding-bottom: 10px;
  margin-bottom: 12px;
  color: #1976d2;
}

.sidebar-header h2 {
  margin: 0;
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
}

.top-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.top-list li {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 8px;
  border-radius: 8px;
  font-size: 12.5px;
}

.rank-num {
  color: #000;
  width: 20px;
  height: 20px;
  min-width: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  font-size: 11px;
  font-weight: 700;
  border: 1px solid #000;
}

.top-title {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  color: #334155;
  font-weight: 500;
}

.main-content {
  display: flex;
  flex-direction: column;
  gap: 24px;
  /* background-color: #ffffff; */
  /* border: 1px solid #e2e8f0; */
  padding: 16px;
}

.tabs-nav {
  display: flex;
  gap: 24px;
  border-bottom: 2px solid #e2e8f0;
  background-color: #ffffff;
  padding: 7px 12px;
  border-radius: 12px;
}

.tab-btn {
  background: none;
  border: none;
  font-size: 15px;
  font-weight: 600;
  color: #64748b;
  padding: 8px 4px 12px 4px;
  cursor: pointer;
  position: relative;
  transition: color 0.2s ease;
}

.tab-btn.active {
  color: #1976d2;
}

.tab-btn.active::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 100%;
  height: 2.5px;
  background-color: #1976d2;
  border-radius: 2px 2px 0 0;
}

.state-msg {
  text-align: center;
  padding: 40px 0;
  color: #64748b;
  font-size: 14px;
}

.state-msg.error {
  color: #ef4444;
}

.quotes-feed {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.quote-card-post {
  background: #ffffff;
  border: 1px solid #e9ecef;
  border-radius: 16px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  position: relative;
  box-shadow: 0 4px 20px -4px rgba(0, 0, 0, 0.03);
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  overflow: visible;
}

.quote-card-post:hover {
  /* background-color: #00000000; */
  opacity: 0.8;
}

.quote-header-text {
  display: flex;
  align-items: flex-start;
  gap: 14px;
}

.quote-badge-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  min-width: 32px;
  background: transparent;
  color: #000;
  border-radius: 50%;
}

.quote-content-block {
  display: flex;
  flex-direction: column;
  gap: 4px;
  position: relative;
}

.quote-header-text h3 {
  margin: 0;
  font-size: 16.5px;
  font-weight: 700;
  color: #0f172a;
  letter-spacing: -0.2px;
}

.quote-header-text p {
  margin: 0;
  font-size: 14.5px;
  color: #475569;
  line-height: 1.65;
  font-weight: 400;
  word-break: break-word;
  overflow-wrap: break-word;
  white-space: normal;
}

.quote-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-top: 2px dashed #f1f5f9;
  padding-top: 12px;
  margin-top: 4px;
  padding-left: 8px;
}

.user-profile {
  display: flex;
  align-items: center;
  gap: 10px;
}

.avatar {
  width: 32px;
  height: 32px;
  background: #e3f2fd;
  color: #1976d2;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.username {
  font-size: 14px;
  font-weight: 600;
  color: #334155;
}

.quote-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.like-container {
  position: relative;
  display: inline-block;
}

.action-pill {
  display: flex;
  align-items: center;
  gap: 6px;
  background: #ECEAE4;
  border: none;
  padding: 7px 12px 7px 7px;
  border-radius: 32px;
  font-size: 13px;
  font-weight: 500;
  color: #000;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.icon-circle {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 23px;
  height: 23px;
  background-color: rgba(255, 255, 255, 0.112); 
  border-radius: 50%;
  backdrop-filter: blur(4px); 
}

.icon-circle svg {
  color: #000;
}

.action-pill:hover {
  opacity: 0.8;
}

.icon-btn {
  background: #ECEAE4;
  border: none;
  padding: 7px 14px;
  border-radius: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #000;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.icon-btn:hover {
  opacity: 0.8;
}

.reaction-display-label {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 600;
}

.selected-emoji {
  font-size: 20px; 
  line-height: 1;
}

.selected-img-wrap {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
}

.selected-img {
  width: 100% !important;
  height: 100% !important;
  object-fit: contain;
  display: block;
}

.selected-svg-wrap {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 23px;
  height: 23px;
}

.selected-svg-wrap :deep(svg) {
  width: 100%;
  height: 100%;
}

.share-container {
  position: relative;
  display: inline-block;
}

.options-container {
  position: relative;
  display: inline-block;
}

.like-count {
  font-size: 13px;
  font-weight: 600;
  color: #000;
  margin-left: 2px;
}

.action-pill.like-btn.is-active {
  background: #fff;
  border: 1.5px solid #1b77d219;
  background-color: #1b77d219;
}

.action-pill.like-btn.is-active .selected-svg-wrap,
.action-pill.like-btn.is-active .selected-img-wrap,
.action-pill.like-btn.is-active .selected-emoji {
  width: 22px;
  height: 22px;
}

.action-pill.like-btn.is-active .like-count {
  color: #1B75D2;

  font-weight: 700;
}


.quote-overlay {
  position: fixed;
  inset: 0;
  z-index: 99999;
top: 60px; 
  display: flex;
  align-items: center;
  justify-content: center;

  padding: 24px;

  background: rgba(15, 23, 42, 0.58);
  backdrop-filter: blur(5px);

  overflow-y: auto;
}

.quote-overlay :deep(.quote-card-form) {
  width: 100%;
  max-width: 600px;
  max-height: calc(100vh - 48px);
  overflow-y: auto;
}

.quote-modal-enter-active,
.quote-modal-leave-active {
  transition: opacity 0.2s ease;
}

.quote-modal-enter-from,
.quote-modal-leave-to {
  opacity: 0;
}

.quote-modal-enter-active .quote-card-form,
.quote-modal-leave-active .quote-card-form {
  transition: all 0.25s ease;
}

.quote-modal-enter-from .quote-card-form {
  opacity: 0;
  transform: translateY(20px) scale(0.97);
}

.quote-modal-leave-to .quote-card-form {
  opacity: 0;
  transform: translateY(15px) scale(0.98);
}
</style>