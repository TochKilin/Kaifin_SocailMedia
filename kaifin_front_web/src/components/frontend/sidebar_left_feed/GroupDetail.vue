<script setup>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'
import NavBar from '../navbar/NavBar.vue'
import PostComposer from '../PostComposer/PostComposer.vue'
import Comments from '../comments/Comments.vue'
import { REACTIONS } from '@/components/reaction/reactions.js'

const BASE_URL = 'http://localhost:7070'

function resolveImageUrl(path) {
  if (!path) return ''
  if (path.startsWith('http://') || path.startsWith('https://') || path.startsWith('data:')) return path
  if (path.startsWith('/')) return `${BASE_URL}${path}`
  return `${BASE_URL}/uploads/${path}`
}

function authHeaders() {
  const token = localStorage.getItem('token')
  return token ? { Authorization: `Bearer ${token}` } : {}
}

function getCurrentUserId() {
  const token = localStorage.getItem('token')
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

const route = useRoute()
const communityId = route.params.id

const loading = ref(false)
const errorMsg = ref('')

const activeTab = ref('Popular')
const tabs = [
  {
    name: 'Popular',
    icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.657 18.657A8 8 0 0 1 6.343 7.343S7 9 9 10c0-2 .5-5 2.986-7C14 5 16.09 5.777 17.656 7.343A7.975 7.975 0 0 1 20 13a7.975 7.975 0 0 1-2.343 5.657z"/></svg>'
  },
  {
    name: 'Latest',
    icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><polyline points="12 7 12 12 15 15"/></svg>'
  },
  {
    name: 'Featured',
    icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/></svg>'
  },
  {
    name: 'Industry News',
    icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 11h16M4 11a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2M4 11V7a2 2 0 0 1 2-2h6l4 4"/></svg>'
  },
  {
    name: 'Prompts',
    icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M8 9h8M8 13h5M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>'
  }
]

const TEXT_LIMIT = 200

const groupInfo = reactive({
  id: null,
  name: '',
  description: '',
  avatarUrl: '',
  coverUrl: '',
  isJoined: false,
  isVerified: false,
})

const groupStats = reactive({
  members: 0,
  boilingPoint: 0
})

const activeMembers = ref([])
const adminMember = ref(null)

const posts = ref([])

function isTextTruncatable(post) {
  return post.content && post.content.length > TEXT_LIMIT
}

function getDisplayText(post) {
  const text = post.content
  if (!text) return ''
  if (post.showFullText || text.length <= TEXT_LIMIT) {
    return text
  }
  return text.slice(0, TEXT_LIMIT) + '...'
}

function toggleFullText(post) {
  post.showFullText = !post.showFullText
}

let reactionCloseTimer = null
function openReactionPicker(post) {
  clearTimeout(reactionCloseTimer)
  posts.value.forEach((p) => { if (p !== post) p.showReactions = false })
  post.showReactions = true
}

function scheduleCloseReactionPicker(post) {
  clearTimeout(reactionCloseTimer)
  reactionCloseTimer = setTimeout(() => {
    post.showReactions = false
  }, 300)
}

function keepReactionPickerOpen() {
  clearTimeout(reactionCloseTimer)
}

function toggleReaction(post) {
  const target = post.reaction
    ? REACTIONS.find((r) => r.key === post.reaction)
    : REACTIONS.find((r) => r.key === 'like')
  pickReaction(post, target)
}


const defaultLikeIcon = `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
  <path d="M7 11v9H4v-9h3Zm3 9h8a2 2 0 0 0 2-2l1.5-5a2 2 0 0 0-2-2.6H15l.7-4A2 2 0 0 0 13.8 4L10 10v10Z"/>
</svg>`


async function pickReaction(post, reaction) {
  clearTimeout(reactionCloseTimer)
  post.showReactions = false

  const previous = {
    reaction: post.reaction,
    isLiked: post.isLiked,
    smileCount: post.actions.smile,
  }
  const wasSame = post.reaction === reaction.key
  post.reaction = wasSame ? null : reaction.key
  post.isLiked = !wasSame
  post.actions.smile += wasSame ? -1 : previous.reaction ? 0 : 1

  try {
    const res = await axios.post(
      `${BASE_URL}/api/v1/front/likes/create`,
      { post_id: post.id, reaction_type: reaction.key },
      { headers: authHeaders() }
    )
    const data = res.data?.data ?? res.data
    if (typeof data?.liked === 'boolean') {
      post.isLiked = data.liked
      post.reaction = data.liked ? reaction.key : null
    }
    await syncLikes(post)
  } catch (err) {
    console.error('Failed to update reaction', err)
    post.reaction = previous.reaction
    post.isLiked = previous.isLiked
    post.actions.smile = previous.smileCount
  }
}


const syncLikes = async (post) => {
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/likes/show`, {
      params: { post_id: post.id },
      headers: authHeaders()
    })
    const data = res.data?.data ?? res.data
    post.actions.smile = data.total ?? 0
    post.isLiked = data.liked_by_me ?? false
    post.reaction = data.my_reaction ?? null
  } catch (err) {
    console.error('Failed to sync likes for post', post.id, err)
  }
}


function getReactionIcon(post) {
  if (!post.reaction) return null
  return REACTIONS.find(r => r.key === post.reaction)?.icon || null
}


const SHARE_OPTIONS = [
  {
    key: 'internal',
    label: 'Share to Profile',
    bg: '#1976D2',
    svg: '<svg viewBox="0 0 24 24"><path d="M12 2v14M12 2l5 5M12 2L7 7" stroke="#fff" stroke-width="1.8" fill="none" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 14v5a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5" stroke="#fff" stroke-width="1.8" fill="none" stroke-linecap="round"/></svg>',
  },
  {
    key: 'facebook',
    label: 'Facebook',
    bg: '#1877F2',
    svg: '<svg viewBox="0 0 24 24"><path d="M14 9h3V6h-3c-1.7 0-3 1.3-3 3v2H9v3h2v6h3v-6h2.5l.5-3H14V9.4c0-.2.2-.4.5-.4H14Z" fill="#fff"/></svg>',
  },
  {
    key: 'telegram',
    label: 'Telegram',
    bg: '#29A9EA',
    svg: '<svg viewBox="0 0 24 24"><path d="m3 11 17-7-3 16-6-4-3 3-1-5 9-8-11 6Z" fill="#fff"/></svg>',
  },
  {
    key: 'whatsapp',
    label: 'Whatsapp',
    bg: '#25D366',
    svg: '<svg viewBox="0 0 24 24"><path d="M12 3a9 9 0 0 0-7.8 13.5L3 21l4.7-1.2A9 9 0 1 0 12 3Z" stroke="#fff" stroke-width="1.6" fill="none"/><path d="M8.5 8.5c.3 3 2 4.7 5 5l1-1.3c.5-.6 1-.4 1.5-.1l1.5 1c.3.6-.1 1.7-.8 2.1-1 .6-2.3.4-4-.4-2.3-1.1-3.9-2.7-5-5-.8-1.7-1-3-.4-4 .4-.7 1.5-1.1 2.1-.8l1 1.5c.3.5.5 1-.1 1.5l-1.3 1Z" fill="#fff"/></svg>',
  },
  {
    key: 'copy',
    label: 'Copy link',
    bg: '#4A4A4E',
    svg: '<svg viewBox="0 0 24 24"><rect x="9" y="9" width="12" height="12" rx="2" stroke="#fff" stroke-width="1.8" fill="none"/><path d="M5 15V5a2 2 0 0 1 2-2h10" stroke="#fff" stroke-width="1.8" fill="none" stroke-linecap="round"/></svg>',
  },
]


const vClickOutsideShare = {
  mounted(el, binding) {
    el.__clickOutsideShareHandler = (e) => {
      if (!el.contains(e.target)) binding.value()
    }
    document.addEventListener('click', el.__clickOutsideShareHandler)
  },
  unmounted(el) {
    document.removeEventListener('click', el.__clickOutsideShareHandler)
  },
}

function toggleSharePicker(post) {
  const wasOpen = post.showSharePicker
  posts.value.forEach((p) => { p.showSharePicker = false })
  post.showSharePicker = !wasOpen
}

function closeSharePicker(post) {
  post.showSharePicker = false
}

async function quickShare(post, key) {
  post.showSharePicker = false

  if (key === 'internal') {
    try {
      await axios.post(
        `${BASE_URL}/api/v1/front/posts/shares/create`,
        { post_id: post.id },
        { headers: authHeaders() }
      )
      post.actions.share = (post.actions.share || 0) + 1
    } catch (err) {
      console.error('Failed to share post to profile', err)
    }
    return
  }

  if (key === 'copy') {
    const link = `${window.location.origin}/posts/${post.id}`
    try {
      await navigator.clipboard.writeText(link)
    } catch (err) {
      console.error('Copy failed', err)
    }
    return
  }

  const link = encodeURIComponent(`${window.location.origin}/posts/${post.id}`)
  const text = encodeURIComponent(post.content || '')
  const deepLinks = {
    facebook: `https://www.facebook.com/sharer/sharer.php?u=${link}`,
    telegram: `https://t.me/share/url?url=${link}&text=${text}`,
    whatsapp: `https://wa.me/?text=${text}%20${link}`,
  }
  if (deepLinks[key]) {
    window.open(deepLinks[key], '_blank', 'noopener,noreferrer')
  }
}


async function toggleBookmark(post) {
  const previous = {
    isBookmarked: post.isBookmarked,
    save: post.actions.save ?? 0,
  }
  post.isBookmarked = !post.isBookmarked
  post.actions.save = (post.actions.save ?? 0) + (post.isBookmarked ? 1 : -1)

  try {
    const res = await axios.post(
      `${BASE_URL}/api/v1/front/bookmarks/create`,
      { post_id: post.id },
      { headers: authHeaders() }
    )
    const data = res.data?.data ?? res.data
    const bookmarked = data?.bookmarked ?? data?.Bookmarked
    const total = data?.total ?? data?.Total
    if (typeof bookmarked === 'boolean') post.isBookmarked = bookmarked
    if (typeof total === 'number') post.actions.save = total
  } catch (err) {
    console.error('Failed to update bookmark', err)
    post.isBookmarked = previous.isBookmarked
    post.actions.save = previous.save
  }
}

async function syncBookmark(post) {
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/bookmarks/show`, {
      params: { post_id: post.id },
      headers: authHeaders()
    })
    const data = res.data?.data ?? res.data
    const rawBookmarks = data.bookmarks ?? []
    const currentUserId = getCurrentUserId()
    post.isBookmarked = rawBookmarks.some(
      (b) =>
        String(b.user_id) === String(currentUserId) &&
        String(b.post_id ?? b.PostID) === String(post.id)
    )
    post.actions.save = rawBookmarks.filter(
      (b) => String(b.post_id ?? b.PostID) === String(post.id)
    ).length
  } catch (err) {
    console.error('Failed to sync bookmark for post', post.id, err)
  }
}


async function syncShares(post) {
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/posts/shares/show`, {
      params: { post_id: post.id },
      headers: authHeaders()
    })
    const data = res.data?.data ?? res.data
    post.actions.share = data.share_count ?? data.shareCount ?? post.actions.share ?? 0
  } catch (err) {
    console.error('Failed to sync shares for post', post.id, err)
  }
}

function onComment(post) {
  post.showComments = !post.showComments
}


const vClickOutsideMore = {
  mounted(el, binding) {
    el.__clickOutsideMoreHandler = (e) => {
      if (!el.contains(e.target)) binding.value()
    }
    document.addEventListener('click', el.__clickOutsideMoreHandler)
  },
  unmounted(el) {
    document.removeEventListener('click', el.__clickOutsideMoreHandler)
  },
}

function toggleMoreMenu(post) {
  const wasOpen = post.showMore
  posts.value.forEach((p) => { p.showMore = false })
  post.showMore = !wasOpen
}

function closeMoreMenu(post) {
  post.showMore = false
}

async function emitAndClose(post, action) {
  post.showMore = false
  if (action === 'hide') {
    try {
      const res = await axios.delete(`${BASE_URL}/api/v1/front/posts/delete/${post.id}`, {
        headers: authHeaders()
      })
      if (res.status >= 200 && res.status < 300) {
        posts.value = posts.value.filter((p) => p.id !== post.id)
      }
    } catch (err) {
      console.error('Failed to delete post', err)
    }
  } else if (action === 'copy-link') {
    const link = `${window.location.origin}/posts/${post.id}`
    try {
      await navigator.clipboard.writeText(link)
    } catch (err) {
      console.error('Copy failed', err)
    }
  } else {
    console.log(`Action "${action}" on post ${post.id}`)
  }
}

const fetchGroupDetail = async () => {
  loading.value = true
  errorMsg.value = ''
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/communities/${communityId}`, {
      headers: authHeaders()
    })
    const c = res.data?.data

    if (c) {
      groupInfo.id = c.id
      groupInfo.name = c.name
      groupInfo.description = c.description || ''
      groupInfo.avatarUrl = resolveImageUrl(c.avatar_url)
      groupInfo.coverUrl = resolveImageUrl(c.cover_url)
      groupInfo.isJoined = c.is_joined ?? false
      groupInfo.isVerified = c.is_verified ?? false

      groupStats.members = c.member_count ?? 0
      groupStats.boilingPoint = c.hot_score ?? 0
    }
  } catch (err) {
    console.error('Failed to fetch community detail:', err)
    errorMsg.value = 'Failed to load group info.'
  } finally {
    loading.value = false
  }
}

const fetchMembers = async () => {
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/communities/${communityId}/members`, {
      params: { page: 1, perpage: 30,community_id: communityId }, //community_id: communityId
      headers: authHeaders()
    })
    const members = res.data?.data?.members ?? []

    adminMember.value = members.find(m => m.role === 'admin') || null

    activeMembers.value = members
      .filter(m => m.role !== 'admin')
      .map(m => ({
        id: m.user_id,
        name: m.user_name,
        avatar: resolveImageUrl(m.profile_images) || 'https://i.pravatar.cc/100'
      }))
  } catch (err) {
    console.error('Failed to fetch members:', err)
  }
}

const handleJoin = async () => {
  const prevState = groupInfo.isJoined
  groupInfo.isJoined = !groupInfo.isJoined

  try {
    const res = await axios.post(
      `${BASE_URL}/api/v1/front/communities/${communityId}/join`,
      {},
      { headers: authHeaders() }
    )
    const status = res.data?.data?.status
    if (status === 'pending') {
      groupInfo.isJoined = false
    } else if (status === 'left') {
      groupInfo.isJoined = false
      groupStats.members = Math.max(0, groupStats.members - 1)
    } else if (status === 'approved' && !prevState) {
      groupStats.members += 1
    }
  } catch (err) {
    console.error('Failed to toggle join:', err)
    groupInfo.isJoined = prevState
  }
}

const fetchPosts = async () => {
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/posts/show`, {
      params: {
        page: 1,
        perpage: 20,
        'filters[0][property]': 'p.community_id',
        'filters[0][value]': communityId
      },
      headers: authHeaders()
    })

    const list = res.data?.data?.posts ?? []

    posts.value = list.map(p => ({
      id: p.id,
      author: p.user_name,
      avatar: resolveImageUrl(p.profile_images) || 'https://i.pravatar.cc/100',
      role: '',
      time: new Date(p.created_at).toLocaleString(),
      content: p.caption,
      // tracks whether this post's content is expanded
      showFullText: false,
      groupTag: null,
      postImage: p.images ? p.images.split(',').filter(Boolean).map(resolveImageUrl) : [],
      videoPath: p.video_path ? resolveImageUrl(p.video_path) : null,
      videoThumbnail: p.thumbnail_path ? resolveImageUrl(p.thumbnail_path) : null,
      videoDuration: p.duration ?? 0,
      videoCurrentTime: 0,
      videoProgress: 0,
      isSeeking: false,
      videoExpanded: false,
      isPlaying: false,
      isMuted: false,
      playbackRate: 1,
      stickerIds: p.sticker_ids
        ? p.sticker_ids.split(',').map(s => Number(s.trim())).filter(Boolean)
        : [],
      stickers: [],
      reaction: null,
      isLiked: false,
      showReactions: false,
      // share picker state
      showSharePicker: false,
      // bookmark / save state
      isBookmarked: false,
      // comment section toggle
      showComments: false,
      // more-options menu toggle
      showMore: false,
      actions: {
        share: 0,
        smile: 0,
        save: 0,
        comment: p.comment_count ?? 0
      }
    }))

    posts.value.forEach(p => syncStickers(p))
    posts.value.forEach(p => syncLikes(p))
    posts.value.forEach(p => syncShares(p))
    posts.value.forEach(p => syncBookmark(p))
  } catch (err) {
    console.error('Failed to fetch posts:', err)
  }
}

const syncStickers = async (post) => {
  if (!post.stickerIds || !post.stickerIds.length) return
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/stickers/show`, {
      params: { ids: post.stickerIds.join(',') },
      headers: authHeaders()
    })
    const data = res.data?.data ?? res.data
    post.stickers = (data?.stickers ?? []).map(s => ({
      ...s,
      url: resolveImageUrl(s.url)
    }))
  } catch (err) {
    console.error('Failed to sync stickers for post', post.id, err)
  }
}

// ============ VIDEO PLAYER LOGIC (copied from Posts.vue) ============
const videoRefs = {}
function setVideoRef(el, postId) {
  if (el) videoRefs[postId] = el
  else delete videoRefs[postId]
}

function pauseAllVideosExcept(activePostId) {
  posts.value.forEach((p) => {
    if (p.id === activePostId) return
    const el = videoRefs[p.id]
    if (el && !el.paused) {
      el.pause()
      p.isPlaying = false
    }
  })
}

function togglePlay(post) {
  const el = videoRefs[post.id]
  if (!el) return
  if (el.paused) {
    pauseAllVideosExcept(post.id)
    el.play()
    post.isPlaying = true
  } else {
    el.pause()
    post.isPlaying = false
  }
}

function toggleMute(post) {
  const el = videoRefs[post.id]
  if (!el) return
  el.muted = !el.muted
  post.isMuted = el.muted
}

function expandVideo(post) {
  post.videoExpanded = !post.videoExpanded
}

function onVideoLoadedMetadata(post, event) {
  post.videoDuration = event.target.duration || 0
}

function onVideoTimeUpdate(post, event) {
  if (post.isSeeking) return
  const video = event.target
  if (!video.duration) return
  post.videoCurrentTime = video.currentTime
  post.videoProgress = (video.currentTime / video.duration) * 100
}

function seekVideo(post, event) {
  const el = videoRefs[post.id]
  const track = event.currentTarget
  if (!el || !el.duration) return
  const rect = track.getBoundingClientRect()
  const ratio = Math.min(Math.max((event.clientX - rect.left) / rect.width, 0), 1)
  el.currentTime = ratio * el.duration
  post.videoProgress = ratio * 100
}

function startSeekDrag(post, event) {
  post.isSeeking = true
  const track = event.currentTarget
  const el = videoRefs[post.id]
  if (!el) return

  const updateFromEvent = (e) => {
    const rect = track.getBoundingClientRect()
    const ratio = Math.min(Math.max((e.clientX - rect.left) / rect.width, 0), 1)
    post.videoProgress = ratio * 100
    if (el.duration) el.currentTime = ratio * el.duration
  }

  updateFromEvent(event)

  const onMouseMove = (e) => updateFromEvent(e)
  const onMouseUp = () => {
    post.isSeeking = false
    document.removeEventListener('mousemove', onMouseMove)
    document.removeEventListener('mouseup', onMouseUp)
  }

  document.addEventListener('mousemove', onMouseMove)
  document.addEventListener('mouseup', onMouseUp)
}

function formatDuration(seconds) {
  if (!seconds) return ''
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m}:${String(s).padStart(2, '0')}`
}

// auto play/pause video as it scrolls in/out of view
let videoObserver = null
function initVideoObserver() {
  videoObserver = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        const postId = Number(entry.target.dataset.videoId)
        const post = posts.value.find((p) => p.id === postId)
        const video = videoRefs[postId]
        if (!video || !post) return
        if (entry.isIntersecting) {
          pauseAllVideosExcept(postId)
          video.play().catch(() => {})
          post.isPlaying = true
        } else {
          video.pause()
          post.isPlaying = false
        }
      })
    },
    { threshold: 0.7 }
  )
}

function observeVideo(el) {
  if (!el) return
  if (videoObserver) videoObserver.observe(el)
}

function handleNewPost(newPost) {
  const post = {
    id: newPost.id,
    author: newPost.username,
    avatar: newPost.avatarUrl || 'https://i.pravatar.cc/100',
    role: '',
    time: newPost.datetime,
    content: newPost.description,
    showFullText: false,
    groupTag: null,
    postImage: newPost.photos || [],
    videoPath: newPost.videoPath || null,
    videoThumbnail: newPost.videoThumbnail || null,
    videoDuration: newPost.videoDuration ?? 0,
    videoCurrentTime: 0,
    videoProgress: 0,
    isSeeking: false,
    videoExpanded: false,
    isPlaying: false,
    isMuted: false,
    playbackRate: 1,
    stickerIds: newPost.stickerIds || [],
    stickers: newPost.stickers || [],
    // reaction state
    reaction: null,
    isLiked: false,
    showReactions: false,
    // share picker state
    showSharePicker: false,
    // bookmark / save state
    isBookmarked: false,
    // comment section toggle
    showComments: false,
    // more-options menu toggle
    showMore: false,
    actions: {
      share: newPost.shareCount ?? 0,
      smile: newPost.likeCount ?? 0,
      save: newPost.bookmarkCount ?? 0,
      comment: newPost.commentCount ?? 0
    }
  }
  posts.value.unshift(post)
  syncStickers(post)
  syncBookmark(post)
}

const showAvatarMenu = ref(false)
const avatarInputRef = ref(null)
const coverInputRef = ref(null)
let avatarMenuCloseTimer = null

function openAvatarMenu() {
  clearTimeout(avatarMenuCloseTimer)
  showAvatarMenu.value = true
}

function keepAvatarMenuOpen() {
  clearTimeout(avatarMenuCloseTimer)
}

function scheduleCloseAvatarMenu() {
  avatarMenuCloseTimer = setTimeout(() => {
    showAvatarMenu.value = false
  }, 250)
}

function triggerAvatarUpload() {
  showAvatarMenu.value = false
  avatarInputRef.value?.click()
}

function triggerCoverUpload() {
  showAvatarMenu.value = false
  coverInputRef.value?.click()
}

async function onAvatarFileChange(e) {
  const file = e.target.files?.[0]
  e.target.value = ''
  if (!file) return

  console.log('Token being sent:', localStorage.getItem('token')) 

  const prevAvatar = groupInfo.avatarUrl
  groupInfo.avatarUrl = URL.createObjectURL(file)

  try {
    const formData = new FormData()
    formData.append('avatar', file)

    const res = await axios.put(
      `${BASE_URL}/api/v1/front/communities/${communityId}/avatar`,
      formData,
      {
        headers: {
          ...authHeaders(),
          'Content-Type': 'multipart/form-data',
        },
      }
    )

    const updated = res.data?.data
    groupInfo.avatarUrl = resolveImageUrl(updated?.avatar_url) || groupInfo.avatarUrl
  } catch (err) {
    console.error('Failed to update group avatar:', err)
    groupInfo.avatarUrl = prevAvatar
  }
}

async function onCoverFileChange(e) {
  const file = e.target.files?.[0]
  e.target.value = ''
  if (!file) return

  const prevCover = groupInfo.coverUrl
  groupInfo.coverUrl = URL.createObjectURL(file) 

  try {
    const formData = new FormData()
    formData.append('cover', file)

    const res = await axios.put(
      `${BASE_URL}/api/v1/front/communities/${communityId}/cover`,
      formData,
      { headers: { ...authHeaders(), 'Content-Type': 'multipart/form-data' } }
    )

    const updated = res.data?.data
    groupInfo.coverUrl = resolveImageUrl(updated?.cover_url) || prevCover
  } catch (err) {
    console.error('Failed to update group cover:', err)
    groupInfo.coverUrl = prevCover
  }
}

const showCoverMenu = ref(false)
let coverMenuCloseTimer = null

function openCoverMenu() {
  clearTimeout(coverMenuCloseTimer)
  showCoverMenu.value = true
}

function keepCoverMenuOpen() {
  clearTimeout(coverMenuCloseTimer)
}

function scheduleCloseCoverMenu() {
  coverMenuCloseTimer = setTimeout(() => {
    showCoverMenu.value = false
  }, 250)
}

function triggerCoverUploadFromBanner() {
  showCoverMenu.value = false
  coverInputRef.value?.click()
}

const lightbox = ref({
  open: false,
  images: [],
  postId: null,
  activeIndex: 0,
  rotation: 0,
})

function openLightbox(post, index) {
  if (!post.postImage || !post.postImage.length) return
  lightbox.value.images = post.postImage
  lightbox.value.postId = post.id
  lightbox.value.activeIndex = index
  lightbox.value.rotation = 0
  lightbox.value.open = true
}

function closeLightbox() {
  lightbox.value.open = false
}

function selectLightboxImage(index) {
  lightbox.value.activeIndex = index
  lightbox.value.rotation = 0
}

function nextLightboxImage() {
  const total = lightbox.value.images.length
  if (!total) return
  lightbox.value.activeIndex = (lightbox.value.activeIndex + 1) % total
  lightbox.value.rotation = 0
}

function prevLightboxImage() {
  const total = lightbox.value.images.length
  if (!total) return
  lightbox.value.activeIndex = (lightbox.value.activeIndex - 1 + total) % total
  lightbox.value.rotation = 0
}

function rotateLightboxImage() {
  lightbox.value.rotation = (lightbox.value.rotation + 90) % 360
}

function viewLargerLightboxImage() {
  const src = lightbox.value.images[lightbox.value.activeIndex]
  if (src) window.open(src, '_blank', 'noopener,noreferrer')
}

function handleLightboxKeydown(e) {
  if (!lightbox.value.open) return
  if (e.key === 'Escape') closeLightbox()
  else if (e.key === 'ArrowRight') nextLightboxImage()
  else if (e.key === 'ArrowLeft') prevLightboxImage()
}

onMounted(() => {
  initVideoObserver()
  fetchGroupDetail()
  fetchMembers()
  fetchPosts()
  document.addEventListener('keydown', handleLightboxKeydown)
})

onUnmounted(() => {
  if (videoObserver) videoObserver.disconnect()
  document.removeEventListener('keydown', handleLightboxKeydown)
})
</script>

<template>
  <div>
    <div class="navbar-wrapper">
      <NavBar/>
    </div>

    <!-- Loading / Error States -->
    <div v-if="loading" class="loading-state">Loading group...</div>
    <div v-else-if="errorMsg" class="error-state">{{ errorMsg }}</div>

    <div v-else class="group-detail-wrapper">
      <div class="group-detail-container">
        <div class="main-content">
          <div class="card group-banner-card">
            <div class="header-top">
              <div class="group-brand">

                <div
                  class="group-avatar khmer-theme avatar-hover-wrap"
                  @mouseenter="openAvatarMenu"
                  @mouseleave="scheduleCloseAvatarMenu"
                >
                  <img v-if="groupInfo.avatarUrl" :src="groupInfo.avatarUrl" :alt="groupInfo.name" style="width:100%;height:100%;object-fit:cover;border-radius:inherit;" />
                  <span v-else class="avatar-title">{{ groupInfo.name?.charAt(0) }}</span>

                  <!-- camera icon overlay on hover -->
                  <div class="avatar-edit-overlay">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                      <circle cx="12" cy="13" r="4"/>
                    </svg>
                  </div>

                  <div
                    v-if="showAvatarMenu"
                    class="avatar-dropdown-menu"
                    @mouseenter="keepAvatarMenuOpen"
                    @mouseleave="scheduleCloseAvatarMenu"
                  >
                    <button class="avatar-dropdown-item" type="button" @click="triggerAvatarUpload">
                      <span class="avatar-dropdown-icon">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                          <rect x="3" y="3" width="18" height="18" rx="2"/>
                          <circle cx="8.5" cy="8.5" r="1.5"/>
                          <path d="M21 15l-5-5L5 21"/>
                        </svg>
                      </span>
                      Change Image
                    </button>

                    <button class="avatar-dropdown-item" type="button" @click="triggerCoverUpload">
                      <span class="avatar-dropdown-icon">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                          <rect x="3" y="3" width="18" height="18" rx="2"/>
                          <path d="M3 15l4-4a3 3 0 0 1 4 0l5 5"/>
                          <path d="M14 13l1-1a3 3 0 0 1 4 0l2 2"/>
                        </svg>
                      </span>
                      Change Cover
                    </button>
                  </div>
                </div>

                <div class="group-meta">
                  <h1 class="group-title">
                    {{ groupInfo.name }}
                    <svg v-if="groupInfo.isVerified" width="16" height="16" viewBox="0 0 24 24" fill="#1B75D2"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z"/></svg>
                  </h1>
                  <div class="group-stats">
                    <span class="stat-item">
                      <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"></path></svg>
                      {{ groupStats.members }} Members
                    </span>
                    <span class="stat-item">
                      <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 18.657A8 8 0 016.343 7.343S7 9 9 10c0-2 .5-5 2.986-7C14 5 16.09 5.777 17.656 7.343A7.975 7.975 0 0120 13a7.975 7.975 0 01-2.343 5.657z"></path></svg>
                      {{ groupStats.boilingPoint }} Hotness
                    </span>
                    <button class="btn-share-link">
                      <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z"></path></svg>
                      Share
                    </button>
                  </div>
                </div>
              </div>

              <button class="btn-join" @click="handleJoin">
                {{ groupInfo.isJoined ? 'Joined' : 'Join' }}
              </button>
            </div>

            <p class="group-description">
              {{ groupInfo.description }}
            </p>
          </div>

          <!-- Card ទី 2: Inline Post Composer (replaces the click-to-open popup) -->
          <div class="card inline-composer-card">
            <PostComposer
              :lockedCommunity="{ id: groupInfo.id, name: groupInfo.name }"
              @post="handleNewPost"
            />
          </div>

          <!-- Post Feed Container -->
          <div class="card feed-card">
            <div class="feed-tabs">
              <button
                v-for="tab in tabs"
                :key="tab.name"
                @click="activeTab = tab.name"
                :class="['tab-item', { active: activeTab === tab.name }]"
              >
                <span class="tab-icon" v-html="tab.icon"></span>
                {{ tab.name }}
                <span v-if="activeTab === tab.name" class="tab-indicator"></span>
              </button>
            </div>
          </div>

          <!-- Post List -->
          <div class="card-group-list">
            <div v-if="posts.length === 0" class="card" style="padding: 24px; text-align: center; color: #888;">
              No posts yet.
            </div>

            <div class="card post-item" v-for="post in posts" :key="post.id">
              <div class="author-info">
                <div class="author-left">
                  <img class="author-avatar" :src="post.avatar" :alt="post.author" />
                  <div class="author-details">
                    <h3 class="author-name">{{ post.author }}</h3>
                    <p class="post-time">{{ post.role }} · {{ post.time }}</p>
                  </div>
                </div>
              </div>

              <p class="post-content">
                {{ getDisplayText(post) }}
                <button
                  v-if="isTextTruncatable(post)"
                  type="button"
                  class="see-more-btn"
                  @click="toggleFullText(post)"
                >
                  {{ post.showFullText ? 'See less' : 'See more' }}
                </button>
              </p>

              <div
                v-if="post.stickers && post.stickers.length"
                class="post-sticker-grid"
              >
                <img
                  v-for="(s, idx) in post.stickers"
                  :key="idx"
                  :src="s.url"
                  :alt="s.file_name || 'sticker'"
                  class="post-sticker-item"
                />
              </div>

              <div
                v-if="post.postImage && post.postImage.length"
                class="post-image-grid"
              >
                <div
                  v-for="(img, idx) in post.postImage.slice(0, 4)"
                  :key="idx"
                  class="post-image-item-wrap"
                  @click.stop.prevent="openLightbox(post, idx)"
                >
                  <img :src="img" class="post-image-item" alt="post image" />
                  <div
                    v-if="idx === 3 && post.postImage.length > 4"
                    class="post-image-more-overlay"
                  >
                    +{{ post.postImage.length - 4 }}
                  </div>
                </div>
              </div>

              <!-- Video -->
              <div
                class="main-warp"
                v-if="post.videoPath"
                :class="{ expanded: post.videoExpanded }"
              >
                <div
                  class="video-container"
                  :data-video-id="post.id"
                  :ref="(el) => observeVideo(el)"
                >
                  <video
                    :ref="(el) => setVideoRef(el, post.id)"
                    :src="post.videoPath"
                    :poster="post.videoThumbnail"
                    preload="metadata"
                    playsinline
                    class="post-video"
                    :class="{ expanded: post.videoExpanded }"
                    @click="expandVideo(post)"
                    @timeupdate="onVideoTimeUpdate(post, $event)"
                    @loadedmetadata="onVideoLoadedMetadata(post, $event)"
                  ></video>

                  <!-- Center Play Button -->
                  <button v-if="!post.isPlaying" class="play-center" @click.stop="togglePlay(post)">
                    <svg viewBox="0 0 24 24"><path d="M8 5v14l11-7z"/></svg>
                  </button>

                  <!-- Duration -->
                  <div class="video-duration">{{ formatDuration(post.videoDuration) }}</div>

                  <!-- Progress Bar -->
                  <div
                    class="progress-track"
                    @click.stop="seekVideo(post, $event)"
                    @mousedown.stop="startSeekDrag(post, $event)"
                  >
                    <div class="progress-fill" :style="{ width: post.videoProgress + '%' }"></div>
                  </div>
                </div>
              </div>

              <div class="post-footer">
                <span class="group-tag" v-if="post.groupTag">
                  <img class="group-thumb" :src="groupInfo.avatarUrl" alt="group" />
                  {{ post.groupTag }} ›
                </span>
                <span v-else></span>

                <div class="post-actions">
                  <!-- Share button + picker -->
                  <div class="share-wrap" v-click-outside-share="() => closeSharePicker(post)">
                    <button class="pill-btn pill-blue" @click.stop="toggleSharePicker(post)">
                      <div class="icon-circle">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><polyline points="16 6 12 2 8 6"/><line x1="12" y1="2" x2="12" y2="15"/></svg>
                      </div>
                      <span>{{ post.actions.share }}</span>
                    </button>

                    <div class="share-picker" v-if="post.showSharePicker" @click.stop>
                      <button
                        v-for="opt in SHARE_OPTIONS"
                        :key="opt.key"
                        class="share-option"
                        :style="{ background: opt.bg }"
                        @click="quickShare(post, opt.key)"
                      >
                        <span class="share-option-svg" v-html="opt.svg"></span>
                        <span class="share-tooltip">{{ opt.label }}</span>
                      </button>
                    </div>
                  </div>

                  <div
                    class="reaction-wrap"
                    @mouseenter="openReactionPicker(post)"
                    @mouseleave="scheduleCloseReactionPicker(post)"
                  >
                  <button
                    class="pill-btn"
                    :class="post.isLiked ? 'pill-white-orange pill-reacted' : 'pill-default-like'"
                    @click="toggleReaction(post)"
                  >
                    <div class="emoji-circle" v-if="getReactionIcon(post)" v-html="getReactionIcon(post)"></div>
                    <div class="icon-circle" v-else v-html="defaultLikeIcon"></div>
                    <span :class="post.isLiked ? 'text-orange' : ''">{{ post.actions.smile }}</span>
                  </button>

                    <div
                        class="reaction-picker"
                        v-if="post.showReactions"
                        @mouseenter="keepReactionPickerOpen"
                        @mouseleave="scheduleCloseReactionPicker(post)"
                      >
                        <button
                          v-for="r in REACTIONS"
                          :key="r.key"
                          type="button"
                          class="reaction-option"
                          :class="[{ active: post.reaction === r.key }, r.private ? 'reaction-option-private' : '']"
                          :title="r.label"
                          @click="pickReaction(post, r)"
                        >
                          <span class="reaction-option-icon" v-html="r.icon"></span>
                          <svg v-if="r.private" class="reaction-lock-badge" viewBox="0 0 24 24" fill="currentColor">
                            <path d="M12 1a5 5 0 0 0-5 5v3H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-9a2 2 0 0 0-2-2h-1V6a5 5 0 0 0-5-5Zm0 2a3 3 0 0 1 3 3v3H9V6a3 3 0 0 1 3-3Z"/>
                          </svg>
                        </button>
                      </div>
                  </div>

                  <button
                    class="pill-btn"
                    :class="post.isBookmarked ? 'pill-saved' : 'pill-blue'"
                    @click="toggleBookmark(post)"
                  >
                    <div class="icon-circle">
                      <svg viewBox="0 0 24 24" :fill="post.isBookmarked ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
                    </div>
                    <span>{{ post.actions.save }}</span>
                  </button>

                  <button class="pill-btn pill-blue" @click="onComment(post)">
                    <div class="icon-circle">
                      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/></svg>
                    </div>
                    <span>{{ post.actions.comment }}</span>
                  </button>

                  <div class="more-wrap" v-click-outside-more="() => closeMoreMenu(post)">
                    <button class="pill-btn pill-blue pill-circle-only" @click.stop="toggleMoreMenu(post)">
                      <div class="icon-circle">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></svg>
                      </div>
                    </button>
                    <div class="more-menu" v-if="post.showMore" @click.stop>
                      <button @click="emitAndClose(post, 'copy-link')">Copy Link</button>
                      <button @click="emitAndClose(post, 'hide')">Hide Post</button>
                      <button class="danger" @click="emitAndClose(post, 'report')">Report</button>
                    </div>
                  </div>
                </div>
              </div>
              <div v-if="post.showComments" class="comment-box">
                <Comments :post-id="post.id" />
              </div>
            </div>
          </div>

        </div>
        <div class="sidebar">
          <div class="card widget-card">
            <h2 class="widget-title">Group Announcement</h2>
            <p class="announcement-text">
              {{ groupInfo.description }}
            </p>
            <button class="btn-expand">Expand</button>
          </div>

          <div class="card widget-card">
            <h2 class="widget-title">Group Members</h2>

            <div class="member-section" v-if="adminMember">
              <span class="section-label">Administrator</span>
              <div class="admin-info">
                <img class="admin-avatar" :src="resolveImageUrl(adminMember.profile_images) || 'https://i.pravatar.cc/100'" :alt="adminMember.user_name" style="width:32px;height:32px;border-radius:50%;object-fit:cover;" />
                <span class="admin-name">{{ adminMember.user_name }}</span>
              </div>
            </div>
            <div class="member-section">
              <span class="section-label">Active Members</span>
              <div class="avatar-stack">
                <img
                  v-for="member in activeMembers.slice(0, 8)"
                  :key="member.id"
                  :src="member.avatar"
                  :alt="member.name"
                  class="stacked-avatar"
                />
              </div>
              <a href="#" class="member-count-link">
                {{ groupStats.members }} users have joined ›
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <input
      ref="avatarInputRef"
      type="file"
      accept="image/*"
      style="display: none;"
      @change="onAvatarFileChange"
    />
    <input
      ref="coverInputRef"
      type="file"
      accept="image/*"
      style="display: none;"
      @change="onCoverFileChange"
    />

    <Teleport to="body">
      <div v-if="lightbox.open" class="lightbox-overlay" @click.self="closeLightbox">
        <div class="lightbox-toolbar">
          <button class="lightbox-tool-btn" @click="closeLightbox">
            <svg viewBox="0 0 24 24"><path d="M15 3H9v2h6zM5 7v2h14V7H5zm2 12a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9H7v10z"/></svg>
            Collapse
          </button>
          <span class="lightbox-divider">|</span>
          <button class="lightbox-tool-btn" @click="rotateLightboxImage">
            <svg viewBox="0 0 24 24"><path d="M17.65 6.35A8 8 0 1 0 19.8 13h-2.1a6 6 0 1 1-1.44-6.16L13 10h7V3z"/></svg>
            Rotation
          </button>
          <span class="lightbox-divider">|</span>
          <button class="lightbox-tool-btn" @click="viewLargerLightboxImage">
            <svg viewBox="0 0 24 24"><path d="M4 4h6V2H2v8h2V4zm14 0v6h2V2h-8v2h6zM4 20v-6H2v8h8v-2H4zm14 0h-6v2h8v-8h-2v6z"/></svg>
            View larger image
          </button>
        </div>

        <div class="lightbox-main">
          <button
            v-if="lightbox.images.length > 1"
            class="lightbox-nav-btn lightbox-nav-prev"
            @click.stop="prevLightboxImage"
          >
            <svg viewBox="0 0 24 24"><path d="M15.4 7.4 14 6l-6 6 6 6 1.4-1.4L10.8 12z"/></svg>
          </button>

          <img
            :src="lightbox.images[lightbox.activeIndex]"
            class="lightbox-image"
            :style="{ transform: 'rotate(' + lightbox.rotation + 'deg)' }"
            alt=""
          />

          <button
            v-if="lightbox.images.length > 1"
            class="lightbox-nav-btn lightbox-nav-next"
            @click.stop="nextLightboxImage"
          >
            <svg viewBox="0 0 24 24"><path d="m8.6 16.6 1.4 1.4 6-6-6-6-1.4 1.4 4.6 4.6z"/></svg>
          </button>
        </div>

        <div class="lightbox-thumbs" v-if="lightbox.images.length > 1">
          <div
            v-for="(img, idx) in lightbox.images"
            :key="idx"
            class="lightbox-thumb"
            :class="{ active: idx === lightbox.activeIndex }"
            @click.stop="selectLightboxImage(idx)"
          >
            <img :src="img" alt="" />
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<style scoped>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.navbar-wrapper {
  position: sticky;
  top: 0;
  z-index: 1000;
  background-color: #ffffff;
}

.group-detail-wrapper {
  min-height: 100vh;
  background-color: #F7F4F2;
  /* background-color: transparent; */
  color: #252933;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  padding: 12px;
}

.group-detail-container {
  max-width: 1140px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr;
  gap: 16px;
}

@media (min-width: 992px) {
  .group-detail-container {
    grid-template-columns: 1fr 260px;
  }
}

.card {
  background-color: #ffffff;
  border-radius: 12px;
}

.main-content {
  display: flex;
  flex-direction: column;
  gap: 2px;

}

.cover-banner {
  position: relative;
  height: 220px;
  border-radius: 12px;
  background-color: #E4E1DC;
  background-size: cover;
  background-position: center;
  overflow: hidden;
  display: flex;
  align-items: flex-end;
  justify-content: flex-end;
  padding: 14px;
  cursor: pointer;
}

.cover-edit-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0);
  transition: background 0.15s ease;
}

.cover-banner:hover .cover-edit-overlay {
  background: rgba(0, 0, 0, 0.28);
}

.cover-edit-btn {
  position: relative;
  z-index: 2;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border: none;
  background: rgba(255, 255, 255, 0.95);
  color: #1a1a1a;
  font-size: 13px;
  font-weight: 600;
  padding: 8px 14px;
  border-radius: 999px;
  cursor: pointer;
  opacity: 0;
  transform: translateY(6px);
  transition: opacity 0.15s ease, transform 0.15s ease;
}

.cover-banner:hover .cover-edit-btn {
  opacity: 1;
  transform: translateY(0);
}

.cover-edit-btn:hover {
  background: #ffffff;
}

.cover-edit-btn svg {
  width: 16px;
  height: 16px;
  stroke: currentColor;
}

.group-banner-card {
  padding: 20px;
  border-bottom-left-radius: 0;
  border-bottom-right-radius: 0;
}

.header-top {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.group-brand {
  display: flex;
  gap: 16px;
  align-items: flex-start;
}

.group-avatar.khmer-theme {
  width: 76px;
  height: 76px;
  border-radius: 16px;
  background: #1B75D2;
  color: #ffffff;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.avatar-title {
  font-size: 11px;
  font-weight: 700;
  margin-top: 2px;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
}

.avatar-hover-wrap {
  position: relative;
  cursor: pointer;
}

.avatar-edit-overlay {
  position: absolute;
  inset: 0;
  border-radius: inherit;
  background: rgba(0, 0, 0, 0);
  display: flex;
  align-items: center;
  justify-content: center;
  color: transparent;
  transition: background 0.15s ease, color 0.15s ease;
}

.avatar-hover-wrap:hover .avatar-edit-overlay {
  background: rgba(0, 0, 0, 0.35);
  color: #ffffff;
}

.avatar-edit-overlay svg {
  width: 22px;
  height: 22px;
  stroke: currentColor;
}

.avatar-dropdown-menu {
  position: absolute;
  top: calc(100% + 8px);
  left: 0;
  min-width: 220px;
  background: #ffffff;
  border-radius: 14px;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.18);
  padding: 6px;
  z-index: 50;
}

.avatar-dropdown-item {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
  text-align: left;
  border: none;
  background: transparent;
  padding: 10px 12px;
  border-radius: 10px;
  font-size: 13.5px;
  font-weight: 500;
  color: #1a1a1a;
  cursor: pointer;
  transition: background 0.15s ease;
}

.avatar-dropdown-item:hover {
  background: #F2F1EE;
}

.avatar-dropdown-icon {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #4A4A4E;
}

.avatar-dropdown-icon svg {
  width: 100%;
  height: 100%;
}

.group-title {
  font-size: 20px;
  font-weight: 700;
  color: #1D2129;
}

.group-stats {
  display: flex;
  align-items: center;
  gap: 20px;
  font-size: 12px;
  color: #8A919F;
  margin-top: 8px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.btn-share-link {
  background: none;
  border: none;
  color: #8A919F;
  font-size: 12px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 4px;
}

.btn-share-link:hover {
  color: #1E80FF;
}

.btn-join {
  background-color: #1E80FF;
  color: #ffffff;
  border: none;
  border-radius: 32px;
  padding: 8px 24px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
}

.btn-join:hover {
  background-color: #1171EE;
}

.group-description {
  font-size: 13px;
  color: #515767;
  margin-top: 16px;
  line-height: 1.5;
}

.inline-composer-card {
  background-color: #ffffff;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
  margin-top: 2px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

.feed-card {
  overflow: hidden;
  margin-top: 12px;

}

.feed-tabs {
  display: flex;
  align-items: center;
  gap: 32px;
  padding: 0 20px;
  padding-top: 4px;
  margin-bottom: 8px;
  position: relative;
  border-top-left-radius: 12px;
}

.inline-composer-card :deep(.composer) {
  max-width: none;
  width: 100%;
  margin-top: 0;
  box-sizing: border-box;
}

.feed-tabs::after {
  content: "";
  position: absolute;
  bottom: 0;
  left: 20px;
  right: 20px;
  height: 1px;
  background-color: #F2F3F5;
}

.tab-item {
  background: none;
  border: none;
  padding: 14px 0;
  font-size: 13px;
  color: #515767;
  cursor: pointer;
  position: relative;
  display: inline-flex;
  align-items: center;
  gap: 6px;
}

.tab-icon {
  display: inline-flex;
  width: 16px;
  height: 16px;
}

.tab-icon :deep(svg) {
  width: 100%;
  height: 100%;
  stroke: currentColor;
}

.tab-item:hover {
  color: #252933;
}

.tab-item.active {
  color: #1E80FF;
  font-weight: 600;
}

.tab-indicator {
  position: absolute;
  bottom: -4px;
  left: 50%;
  transform: translateX(-50%);
  width: 24px;
  height: 2px;
  background-color: #1E80FF;
  border-radius: 2px;
}

.post-item {
  padding: 20px;
  position: relative;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.author-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.author-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.author-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
  display: block;
  flex-shrink: 0;
  background-color: #F2F3F5;
}

.author-name {
  font-size: 13px;
  font-weight: 600;
  color: #252933;
}

.post-time {
  font-size: 11px;
  color: #8A919F;
  margin-top: 2px;
}

.post-content {
  font-size: 13px;
  color: #252933;
  line-height: 1.5;
  white-space: pre-wrap;
  word-break: break-word;
}


.see-more-btn {
  display: inline;
  border: none;
  background: transparent;
  color: #1E80FF;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  padding: 0 0 0 4px;
  margin: 0;
}

.see-more-btn:hover {
  text-decoration: underline;
}

.post-sticker-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.post-sticker-item {
  width: 80px;
  height: 80px;
  object-fit: contain;
  background-color: #F2F3F5;
  border-radius: 12px;
  padding: 6px;
}

.post-image-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
  max-width: 320px;
}

.post-image-item-wrap {
  position: relative;
  aspect-ratio: 1 / 1;
  border-radius: 14px;
  overflow: hidden;
  background-color: #F2F3F5;
}

.post-image-item {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  cursor: pointer;
  transition: opacity 0.2s ease;
}

.post-image-item:hover {
  opacity: 0.92;
}

.post-image-more-overlay {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgba(0, 0, 0, 0.45);
  color: #ffffff;
  font-size: 24px;
  font-weight: 700;
  cursor: pointer;
}

.group-tag {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  color: #1E80FF;
  background-color: #E8F3FF;
  padding: 4px 12px 4px 4px;
  border-radius: 16px;
  cursor: pointer;
}

.group-thumb {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.group-tag:hover {
  background-color: #D4E8FF;
}

.post-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  flex-wrap: wrap;
  padding-top: 12px;
  margin-top: 4px;
  border-top: 1px solid #F2F3F5;
}

.post-actions {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
  margin-left: auto;
}

.pill-btn {
  display: inline-flex;
  align-items: center;
  height: auto;
  padding: 7px 12px;
  border-radius: 20px;
  border: none;
  cursor: pointer;
  font-weight: 700;
  font-size: 13px;
  gap: 6px;
  transition: transform 0.1s ease, opacity 0.2s ease;
}

.pill-btn:active {
  transform: scale(0.96);
}

.pill-blue {
  background-color: #1B76E8;
  color: #ffffff;
}

.pill-blue:hover {
  background-color: #1565C0;
}

.pill-white-orange {
  background-color: #ffffff;
  border: 1.5px solid #F9CF9D;
  color: #E65100;
}

.pill-white-orange:hover {
  background-color: #FFF8E1;
}

.pill-circle-only {
  padding: 7px 12px;
}

.icon-circle {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.18);
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-circle svg {
  width: 14px;
  height: 14px;
  color: #ffffff;
}

.emoji-circle {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  background-color: #FBC02D;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 13px;
  overflow: hidden;
}

.emoji-circle svg {
  width: 100%;
  height: 100%;
  display: block;
}

.text-orange {
  color: #E65100;
}

/* Reaction picker */
.reaction-wrap {
  position: relative;
  display: inline-flex;
}

.reaction-picker {
  position: absolute;
  bottom: 100%;
  left: -200px;
  margin-bottom: 8px;
  display: flex;
  align-items: center;
  gap: 6px;
  background: #ffffff;
  border-radius: 999px;
  padding: 8px 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.18);
  z-index: 20;
}

.reaction-option {
  position: relative;
  border: none;
  background: transparent;
  width: 36px;
  height: 36px;
  padding: 0;
  cursor: pointer;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.15s ease;
  flex-shrink: 0;
}
.reaction-option-icon {
  width: 100%;
  height: 100%;
  display: block;
}

.reaction-option-icon :deep(svg) {
  width: 100%;
  height: 100%;
  display: block;
}

.reaction-option:hover {
  transform: scale(1.25) translateY(-4px);
}

.reaction-option.active {
  background: #E8F3FF;
}

/* Private Like — dark circle, bigger than the rest, with a small lock badge */
.reaction-option-private {
  width: 44px;
  height: 44px;
  background: #4A4A4E;
  color: #ffffff;
}

.reaction-option-private:hover {
  background: #3A3A3E;
}

.reaction-option-private .reaction-option-icon {
  width: 60%;
  height: 60%;
}

.reaction-lock-badge {
  position: absolute;
  bottom: -3px;
  right: -3px;
  width: 15px;
  height: 15px;
  color: #B5B5B8;
  background: transparent;
}

/* Share picker */
.share-wrap {
  position: relative;
  display: inline-flex;
}

.share-picker {
  position: absolute;
  bottom: 100%;
  left: -80px;
  margin-bottom: 8px;
  background: #ffffff;
  border-radius: 999px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.18);
  padding: 8px 10px;
  display: flex;
  align-items: center;
  gap: 4px;
  z-index: 30;
  animation: popIn .15s ease;
}

@keyframes popIn {
  from { opacity: 0; transform: translateY(6px) scale(.9); }
  to   { opacity: 1; transform: translateY(0) scale(1); }
}

.share-option {
  position: relative;
  border: none;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform .12s ease;
  padding: 0;
}

.share-option:hover {
  transform: scale(1.15) translateY(-3px);
}

.share-option-svg {
  width: 18px;
  height: 18px;
  display: flex;
}

.share-option-svg svg {
  width: 100%;
  height: 100%;
}

.share-tooltip {
  position: absolute;
  bottom: calc(100% + 8px);
  left: 50%;
  transform: translateX(-50%);
  background: #1976D2;
  color: #fff;
  font-size: 11px;
  font-weight: 600;
  padding: 4px 8px;
  border-radius: 6px;
  white-space: nowrap;
  opacity: 0;
  pointer-events: none;
  transition: opacity .12s ease;
}

.share-option:hover .share-tooltip {
  opacity: 1;
}

.sidebar {
  display: flex;
  flex-direction: column;
  gap: 16px;
  position: sticky;
  top: 76px;
  align-self: start;
  max-height: calc(100vh - 88px);
  overflow-y: auto;
}

.widget-card {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.widget-title {
  font-size: 13px;
  font-weight: bold;
  color: #252933;
  padding-bottom: 10px;
  position: relative;
}

.widget-title::after {
  content: "";
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  width: calc(100% - 32px);
  margin: 0 auto;
  height: 1px;
  background-color: #F2F3F5;
}

.announcement-text {
  font-size: 12px;
  color: #515767;
  line-height: 1.6;
}

.btn-expand {
  background: none;
  border: none;
  color: #1E80FF;
  font-size: 12px;
  cursor: pointer;
  text-align: left;
}

.btn-expand:hover {
  text-decoration: underline;
}

.member-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.section-label {
  font-size: 11px;
  color: #8A919F;
}

.admin-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.admin-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: #DBEAFE;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
}

.admin-name {
  font-size: 12px;
  color: #252933;
  font-weight: 500;
}

.avatar-stack {
  display: flex;
  align-items: center;
  margin-top: 6px;
}

.stacked-avatar {
  width: 30px;
  height: 30px;
  border-radius: 8px;
  border: 2px solid #ffffff;
  object-fit: cover;
  margin-left: -8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.08);
  transition: transform 0.2s ease, z-index 0.2s;
}

.stacked-avatar:first-child {
  margin-left: 0;
}

.stacked-avatar:hover {
  transform: translateY(-2px);
  z-index: 10;
}

.member-count-link {
  font-size: 12px;
  color: #8A919F;
  text-decoration: none;
  margin-top: 8px;
  display: inline-block;
}

.member-count-link:hover {
  color: #1E80FF;
}

.icon {
  width: 14px;
  height: 14px;
}

.icon-small {
  width: 12px;
  height: 12px;
}

.card-group-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin-top: 12px;
}

.pill-default-like {
  background-color: #1B76E8;
  color: #ffffff;
}

.pill-default-like:hover {
  background-color: #1565C0;
}

.pill-default-like .icon-circle {
  background: rgba(255, 255, 255, 0.18);
  padding: 5px;
}

.pill-default-like .icon-circle svg {
  width: 100%;
  height: 100%;
  stroke: #ffffff;
  fill: none;
}

.pill-default-like span {
  color: #ffffff;
}

/* Reacted state — white pill, tan border, colored emoji + count (matches picker style) */
.pill-white-orange {
  background-color: #ffffff;
  border: 1.5px solid #F9CF9D;
  color: #E65100;
}

.pill-white-orange:hover {
  background-color: #FFF8E1;
}

.text-orange {
  color: #E65100;
}

.pill-saved {
  background-color: #ffffff;
  border: 1.5px solid #F9CF9D;
  color: #E65100;
}

.pill-saved:hover {
  background-color: #FFF8E1;
}

.pill-saved .icon-circle {
  background: transparent;
  color: #E65100;
}

.pill-saved .icon-circle svg {
  color: #E65100;
  stroke: currentColor;
  fill: currentColor;
}

.pill-saved span {
  color: #E65100;
}

.comment-box {
  margin-top: 15px;
  border-top: 1px solid #eee;
  padding-top: 15px;
}


.more-wrap {
  position: relative;
  display: inline-flex;
}

.more-menu {
  position: absolute;
  bottom: 100%;
  right: 0;
  margin-bottom: 8px;
  background: #ffffff;
  border: 1px solid #E7E7E7;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .12);
  padding: 6px;
  min-width: 150px;
  z-index: 30;
  animation: popIn .15s ease;
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
  color: #252933;
  cursor: pointer;
}

.more-menu button:hover {
  background: #F2F2F3;
}

.more-menu button.danger {
  color: #C6402E;
}


.lightbox-overlay {
  position: fixed;
  inset: 0;
  background: rgba(10, 12, 16, 0.94);
  z-index: 9999;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.lightbox-toolbar {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 14px 20px;
  color: #000;
  font-size: 13px;
  background-color: #ffffff;
}

.lightbox-tool-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: none;
  border: none;
  color: #000;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  padding: 4px 6px;
  border-radius: 6px;
  transition: color 0.15s ease, background 0.15s ease;
}

.lightbox-tool-btn:hover {
  opacity: 0.8;
  background: rgba(255, 255, 255, 0.08);
}

.lightbox-tool-btn svg {
  width: 15px;
  height: 15px;
  fill: currentColor;
}

.lightbox-divider {
  color: #475569;
}

.lightbox-main {
  position: relative;
  flex: 1;
  width: 100%;
  max-width: 1100px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 20px;
  min-height: 0;
}

.lightbox-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
  border-radius: 6px;
  transition: transform 0.25s ease;
}

.lightbox-nav-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 42px;
  height: 42px;
  border-radius: 50%;
  border: none;
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 2;
}

.lightbox-nav-btn:hover {
  background: rgba(255, 255, 255, 0.22);
}

.lightbox-nav-btn svg {
  width: 22px;
  height: 22px;
  fill: currentColor;
}

.lightbox-nav-prev {
  left: 20px;
}

.lightbox-nav-next {
  right: 20px;
}

.lightbox-thumbs {
  display: flex;
  gap: 8px;
  padding: 16px 20px 24px;
  overflow-x: auto;
  max-width: 100%;
}

.lightbox-thumb {
  flex-shrink: 0;
  width: 84px;
  height: 60px;
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
  opacity: 0.55;
  border: 2px solid transparent;
  transition: opacity 0.15s ease, border-color 0.15s ease;
}

.lightbox-thumb:hover {
  opacity: 0.85;
}

.lightbox-thumb.active {
  opacity: 1;
  border-color: #F2762E;
}

.lightbox-thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

/* ==========================================
   VIDEO PLAYER (copied from Posts.vue)
   ========================================== */

.main-warp {
  position: relative;
  width: 300px;
  padding: 4px;
  border-radius: 20px;
  transition: width .3s ease;
}

.main-warp.expanded {
  width: 100%;
}

.video-container {
  position: relative;
  width: 100%;
  display: flex;
  justify-content: center;
}

.post-video {
  width: 100%;
  max-height: 400px;
  object-fit: cover;
  border-radius: 12px;
  display: block;
  transition: .3s ease;
}

.post-video.expanded {
  width: 100%;
  height: 500px;
  max-height: none;
  object-fit: contain;
  background: #000;
}

.video-duration {
  position: absolute;
  right: 10px;
  top: 10px;
  background: rgba(0, 0, 0, .6);
  color: white;
  padding: 3px 8px;
  border-radius: 6px;
  font-size: 13px;
}

.play-center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 60px;
  height: 60px;
  border: none;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.55);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 30;
}

.play-center svg {
  width: 32px;
  height: 32px;
  fill: white;
}

.progress-track {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: rgba(255, 255, 255, 0.25);
  cursor: pointer;
  z-index: 25;
  transition: height 0.15s ease;
}

.progress-track:hover {
  height: 6px;
}

.progress-fill {
  height: 100%;
  background: #1976D2;
  border-radius: 0 2px 2px 0;
  pointer-events: none;
  transition: width 0.1s linear;
}
</style>