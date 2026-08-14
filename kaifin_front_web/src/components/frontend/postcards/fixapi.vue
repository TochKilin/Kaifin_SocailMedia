<template>
  <div class="feed">
    <!-- Loading first  -->
    <p v-if="isLoading" class="state-msg">Loading...</p>
    <p v-else-if="error" class="state-msg error">{{ error }}</p>
    <p v-else-if="!posts.length" class="state-msg">No post yet</p>
    <div class="post-card" v-for="post in posts" :key="post.id" :data-post-id="post.id" :ref="(el) => setPostCardRef(el, post.id)">
      <!-- Header post  -->
      <div class="post-top">
        <div class="avatar">
          <img v-if="post.avatarUrl" :src="post.avatarUrl" alt="" />
          <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
        </div>
        <!-- Post body  -->
        <div class="post-body">
          <div class="post-head">
            <div class="user-block">
              <span class="username">{{ post.username }}</span>
              <span class="datetime">
                <svg viewBox="0 0 24 24"><circle cx="12" cy="12" r="8.5"/><path d="M12 7v5l3.2 2"/></svg>
                {{ post.datetime }}
              </span>
            </div>
            <button v-if="String(post.userId) !== String(currentUserId)" class="follow-btn" :class="{ following: post.isFollowing }" @click="toggleFollow(post)">
              {{ post.isFollowing ? 'Following' : 'Follow' }}
            </button>
          </div>
        </div>
       </div>
       <!-- Post description  -->
       <p class="description">{{ post.showTranslated ? post.translatedText : post.description }}</p>
        <div class="post-tags" v-if="post.tags.length">
          <span class="tag-chip" v-for="tag in post.tags" :key="tag" @click="onTagClick(tag)">
                #{{ tag }}
          </span>
        </div>
        <!-- Show photo grid  -->
        <div class="photo-grid" v-if="post.photos.length" :class="'count-' + Math.min(post.photos.length, 4)">
          <div class="photo" v-for="(photo, i) in post.photos.slice(0, 4)" :key="i">
            <img v-if="photo" :src="photo" alt="" />
            <svg v-else viewBox="0 0 24 24" class="photo-placeholder">
              <path d="M3 16l5-5 4 4 5-6 4 5"/>
              <circle cx="8" cy="8" r="1.6"/>
            </svg>
            <span v-if="i === 3 && post.photos.length > 4" class="more-overlay">+{{ post.photos.length - 4 }}</span>
          </div>
        </div>
        <!-- Videos -->
<!-- <div class="main-warp" v-if="post.videoPath">
  <div class="video-wrap">

    <video
      :src="post.videoPath"
      :poster="post.videoThumbnail || undefined"
      controls
      preload="metadata"
      class="post-video"
      @loadedmetadata="setVideoDuration(post, $event)"
    ></video>

    <div class="video-duration">
   
        4:00
    </div>

  </div>
</div> -->


<div class="main-wrap" v-if="post.videoPath">

  <div class="video-container">

    <video
      ref="videoRef"
      :src="post.videoPath"
      :poster="post.videoThumbnail"
      preload="metadata"
      class="post-video"
      @loadedmetadata="setVideoDuration(post, $event)"
    ></video>

    <!-- custom control -->
    <div class="video-controls">

      <!-- Play -->
      <button @click="togglePlay">
        <svg viewBox="0 0 24 24">
          <path v-if="!isPlaying"
          d="M8 5v14l11-7z"/>
          
          <path v-else
          d="M6 5h4v14H6zm8 0h4v14h-4z"/>
        </svg>
      </button>


      <!-- Volume -->
      <button @click="toggleMute">
        <svg viewBox="0 0 24 24">
          <path d="M3 10v4h4l5 5V5L7 10H3z"/>
        </svg>
      </button>


      <!-- Fullscreen -->
      <button @click="fullscreen">
        <svg viewBox="0 0 24 24">
          <path d="M4 4h6v2H6v4H4V4zm10 0h6v6h-2V6h-4V4z"/>
        </svg>
      </button>


      <!-- More -->
      <button>
        <svg viewBox="0 0 24 24">
          <circle cx="12" cy="5" r="2"/>
          <circle cx="12" cy="12" r="2"/>
          <circle cx="12" cy="19" r="2"/>
        </svg>
      </button>

    </div>


    <!-- duration -->
    <div class="video-duration">
      {{formatDuration(post.videoDuration)}}
    </div>

  </div>

</div>






        <!-- Stickers  -->
        <div class="sticker-row" v-if="post.stickers.length">
          <img v-for="s in post.stickers" :key="s.id" :src="s.url" :alt="s.file_name" class="post-sticker" />
        </div>
        <!-- Group  -->
        <div class="post-controls">
          <button class="translate-btn" v-if="post.translatedText" @click="post.showTranslated = !post.showTranslated">
            {{ post.showTranslated ? 'Theme' : 'Translate' }}
          </button>
          <label class="group-toggle">
            <input type="checkbox" v-model="post.postToGroup" />
            <span class="toggle-track"><span class="toggle-thumb"></span></span>
            Group
          </label>
        <!-- Likes Avatar  -->
        <div class="liked-by" v-if="post.likedByAvatars.length"       
        
  @mouseenter="post.showLikers = true"
  :ref="(el) => setLikedByRef(el, post.id)"
>
         <div class="stack">
  <span v-for="(u, i) in post.likedByAvatars.slice(0, 3)" :key="i">
    <img v-if="u.avatarUrl" :src="u.avatarUrl" alt="" />
    <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
  </span>
</div>
          <span>Liked</span>
<ViewerLikes
 v-if="post.showLikers"
    class="viewer-likes-popover"
    :users="post.likedByAvatars"
   
/>

        </div>
        </div>
        <div class="post-foot">
          <div class="post-left">
            <span class="views">
            <svg viewBox="0 0 24 24">
                <path d="M2 12s3.5-7 10-7 10 7 10 7-3.5 7-10 7-10-7-10-7Z"/>
                <circle cx="12" cy="12" r="3"/>
            </svg>
            {{ formatCount(post.views) }}
            </span>
          </div>
          <div class="post-right">
          <div class="share-wrap" v-click-outside-share="() => closeSharePicker(post)">
          <button class="stat-btn" @click.stop="toggleSharePicker(post)">
            <svg viewBox="0 0 24 24"><path d="M4 12v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7"/><path d="M16 6l-4-4-4 4"/><path d="M12 2v14"/></svg>
            {{ formatCount(post.shareCount) }}
          </button>
          <div class="share-picker" v-if="post.showSharePicker" @click.stop>
            <button v-for="opt in SHARE_OPTIONS" :key="opt.key" class="share-option" :style="{ background: opt.bg }" @click="quickShare(post, opt.key)">
              <span class="share-option-svg" v-html="opt.svg"></span>
              <span class="share-tooltip">{{ opt.label }}</span>
            </button>
          </div>
        </div>
        <!-- Reaction button  -->
        <div class="like-wrap" @mouseenter="openReactionPicker(post)" @mouseleave="scheduleCloseReactionPicker(post)" >
          <button class="stat-btn like-btn" :class="{ liked: post.isLiked, [`reaction-${post.reaction}`]: post.reaction }"  @click="toggleLike(post)">
            <span v-if="post.reaction" class="reaction-svg reaction-emoji" v-html="REACTIONS.find(r => r.key === post.reaction)?.icon"></span>
            <svg v-else viewBox="0 0 24 24"><path d="M7 11v9H4v-9h3Zm3 9h8a2 2 0 0 0 2-2l1.5-5a2 2 0 0 0-2-2.6H15l.7-4A2 2 0 0 0 13.8 4L10 10v10Z"/></svg>
            {{ formatCount(post.likeCount) }}
          </button>
          <!-- Hover reactionpicker  -->
          <div class="reaction-picker" v-if="post.showReactions" @mouseenter="keepReactionPickerOpen" @mouseleave="scheduleCloseReactionPicker(post)"
          >
            <button v-for="r in REACTIONS" :key="r.key" class="reaction-option" :class="{ locked: r.private, active: post.reaction === r.key }" @click="pickReaction(post, r)">
              <span class="reaction-svg" v-html="r.icon"></span>
              <span v-if="r.private" class="lock-badge">🔒</span>
              <span class="reaction-tooltip">{{ r.label }}</span>
            </button>
          </div>
        </div>
        <!-- Btn bookkmark -->
        <button class="stat-btn bookmark-btn" :class="{ saved: post.isBookmarked }" @click="toggleBookmark(post)">
          <svg viewBox="0 0 24 24"><path d="M6 3h12a1 1 0 0 1 1 1v17l-7-4-7 4V4a1 1 0 0 1 1-1Z"/></svg>
          {{ formatCount(post.bookmarkCount) }}  
        </button>
        <!-- Btn comments -->
        <button class="stat-btn" @click="onComment(post)">
          <svg viewBox="0 0 24 24"><path d="M21 12a8 8 0 1 1-3.2-6.4L21 4l-1 4.6A7.96 7.96 0 0 1 21 12Z"/></svg>
          {{ formatCount(post.commentCount) }}
        </button>
        <!-- Btn more on posts -->
        <div class="more-wrap">
          <button class="stat-btn more-btn" @click="post.showMore = !post.showMore">
            <svg viewBox="0 0 24 24" fill="currentColor"><circle cx="5" cy="12" r="1.6"/><circle cx="12" cy="12" r="1.6"/><circle cx="19" cy="12" r="1.6"/></svg>
          </button>
          <div class="more-menu" v-if="post.showMore" @click.stop>
            <button @click="emitAndClose(post, 'copy-link')">Copy Link</button>
            <button @click="emitAndClose(post, 'hide')">Hide Post</button>
            <button class="danger" @click="emitAndClose(post, 'report')">Reposrt</button>
          </div>
        </div>
        </div>
      </div>
      <!-- Show comment on posts -->
      <div v-if="post.showComments" class="comment-box">
        <Comments :post-id="post.id" />
      </div>
    </div>
     <!-- Loading for post -->
    <div ref="sentinel" class="sentinel"></div>
    <p v-if="isLoadingMore" class="state-msg">Wiat for fetch</p>
    <p v-if="!hasMore && posts.length" class="state-msg">No post for fetch</p>
  </div>

  
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import Comments from '../comments/Comments.vue'
import ViewerLikes from '../viewer_likes/ViewerLikes.vue'

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'
const PER_PAGE = 50
const posts = ref([])
const currentPage = ref(1)
const hasMore = ref(true)
const isLoading = ref(false)     
const isLoadingMore = ref(false) 
const error = ref(null)
const sentinel = ref(null)
let observer = null
const currentUserId = ref(null)

// Share Icon
const REACTIONS = [
  {
    key: 'private_like',
    label: 'Private Like',
    private: true,
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
      <path d="M7 11v9H4v-9h3Zm3 9h8a2 2 0 0 0 2-2l1.5-5a2 2 0 0 0-2-2.6H15l.7-4A2 2 0 0 0 13.8 4L10 10v10Z" fill="currentColor"/>
    </svg>`,
  },
  {
    key: 'sad',
    label: 'Sade',
    icon: `<svg viewBox="0 0 36 36">
      <circle cx="18" cy="18" r="17" fill="#F2C94C"/>
      <ellipse cx="12.5" cy="15.5" rx="2.2" ry="2.8" fill="#3A2A1A"/>
      <ellipse cx="23.5" cy="15.5" rx="2.2" ry="2.8" fill="#3A2A1A"/>
      <path d="M11 25c1.8-3 5-4.5 7-4.5s5.2 1.5 7 4.5" stroke="#3A2A1A" stroke-width="2" fill="none" stroke-linecap="round"/>
      <path d="M23.8 17c1.2 1.6 1.4 3.6.4 5.4-.6 1-1.8 1.2-2.3.2-.4-.8.1-1.6.6-2.3.6-.9 1-2.1 1.3-3.3Z" fill="#4FA8D8"/>
    </svg>`,
  },
  {
    key: 'wow',
    label: 'Supprise',
    icon: `<svg viewBox="0 0 36 36">
      <circle cx="18" cy="18" r="17" fill="#F2C94C"/>
      <ellipse cx="12.5" cy="14.5" rx="2.4" ry="3" fill="#3A2A1A"/>
      <ellipse cx="23.5" cy="14.5" rx="2.4" ry="3" fill="#3A2A1A"/>
      <ellipse cx="18" cy="24" rx="3.6" ry="4.4" fill="#3A2A1A"/>
    </svg>`,
  },
  {
    key: 'love',
    label: 'Love',
    icon: `<svg viewBox="0 0 36 36">
      <circle cx="18" cy="18" r="17" fill="#F2C94C"/>
      <path d="M12.5 13.2c-1.6 0-2.9 1.2-2.9 2.8 0 2.1 2.9 4 2.9 4s2.9-1.9 2.9-4c0-1.6-1.3-2.8-2.9-2.8Z" fill="#E8543A"/>
      <path d="M23.5 13.2c-1.6 0-2.9 1.2-2.9 2.8 0 2.1 2.9 4 2.9 4s2.9-1.9 2.9-4c0-1.6-1.3-2.8-2.9-2.8Z" fill="#E8543A"/>
      <path d="M11 24c1.8 2.6 4.6 4 7 4s5.2-1.4 7-4" stroke="#3A2A1A" stroke-width="2" fill="none" stroke-linecap="round"/>
    </svg>`,
  },
  {
    key: 'haha',
    label: 'HaHa',
    icon: `<svg viewBox="0 0 36 36">
      <circle cx="18" cy="18" r="17" fill="#F2C94C"/>
      <path d="M9.5 15.5c1-1.6 2.4-2.3 3.5-2.3s2.5.7 3.5 2.3" stroke="#3A2A1A" stroke-width="2" fill="none" stroke-linecap="round"/>
      <path d="M19.5 15.5c1-1.6 2.4-2.3 3.5-2.3s2.5.7 3.5 2.3" stroke="#3A2A1A" stroke-width="2" fill="none" stroke-linecap="round"/>
      <path d="M9.5 20c1.5 4 5 6.5 8.5 6.5s7-2.5 8.5-6.5Z" fill="#3A2A1A"/>
      <path d="M13.5 20h9c-.4 2-2.2 3.5-4.5 3.5s-4.1-1.5-4.5-3.5Z" fill="#fff"/>
    </svg>`,
  },
  {
    key: 'like',
    label: 'Like',
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
      <path d="M7 11v9H4v-9h3Zm3 9h8a2 2 0 0 0 2-2l1.5-5a2 2 0 0 0-2-2.6H15l.7-4A2 2 0 0 0 13.8 4L10 10v10Z" fill="currentColor"/>
    </svg>`,
  },
  {
    key: 'heart',
    label: 'Heart',
    icon: `<svg viewBox="0 0 24 24">
      <path d="M12 21s-7.5-4.6-10-9.1C.4 8.6 2 5 5.6 5 8 5 10 6.4 12 9c2-2.6 4-4 6.4-4C22 5 23.6 8.6 22 11.9 19.5 16.4 12 21 12 21Z" fill="#E8543A"/>
    </svg>`,
  },
  {
    key: 'rose',
    label: 'Rose',
    icon: `<svg viewBox="0 0 24 24">
      <path d="M12 3.5c1.5 0 4 1.2 4 3.7 0 1.4-.8 2.4-1.6 3-.2.1-.1.4.1.4 1.2.1 2.6-.5 3.5-1.6.2-.2.5 0 .4.2-.7 2.4-3 4.1-5.4 4.1-1.7 0-3-.9-4-.9s-2.3.9-4 .9c-2.4 0-4.7-1.7-5.4-4.1-.1-.2.2-.4.4-.2.9 1.1 2.3 1.7 3.5 1.6.2 0 .3-.3.1-.4C2.8 9.6 2 8.6 2 7.2c0-2.5 2.5-3.7 4-3.7 2.3 0 3.6 1.6 4 2.1.4-.5 1.7-2.1 2-2.1Z" fill="#C6402E"/>
      <path d="M12 12.5v8" stroke="#2E7D32" stroke-width="1.8" stroke-linecap="round"/>
      <path d="M12 16c-1.5 0-2.6-1-3-2" stroke="#2E7D32" stroke-width="1.6" fill="none" stroke-linecap="round"/>
    </svg>`,
  },
  {
    key: 'clap',
    label: 'Congrate',
    icon: `<svg viewBox="0 0 24 24">
      <path d="M9 14.5l-2.6-4.3a1.5 1.5 0 1 1 2.6-1.5L11 12" fill="#F2C48B"/>
      <path d="M15 14.5l2.6-4.3a1.5 1.5 0 1 0-2.6-1.5L13 12" fill="#F2C48B"/>
      <path d="M8 15c-.5-2 .5-4 2-4.5s3 0 3.5 1.5" stroke="#D9A15C" stroke-width="1" fill="none"/>
      <path d="M4.5 6.5l1.5 1M19.5 6.5l-1.5 1M12 4v1.6" stroke="#4A4A4E" stroke-width="1.6" stroke-linecap="round"/>
      <path d="M8 16c0 3 1.8 5 4 5s4-2 4-5" fill="#F2C48B"/>
    </svg>`,
  },
  {
    key: 'pray',
    label: 'Pray',
    icon: `<svg viewBox="0 0 24 24">
      <path d="M12 3l-1 6-3.5 4v7c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-4h3v4c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-7l-3.5-4-1-6Z" fill="#7C6FE8"/>
      <path d="M12 3v17.5" stroke="#5C4FC7" stroke-width="1" />
    </svg>`,
  },
]

let reactionCloseTimer = null

function getAuthToken() {
  return localStorage.getItem('token') || ''
}

function authHeaders() {
  const token = getAuthToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
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

onMounted(async () => {
  initViewObserver()
  await loadPosts(1)
  currentUserId.value = getCurrentUserId()
  observer = new IntersectionObserver(
    (entries) => {
      if (entries[0].isIntersecting) {
        loadMore()
      }
    },
    { rootMargin: '400px' } 
  )
  if (sentinel.value) observer.observe(sentinel.value)
})

onUnmounted(() => {
  if (observer) observer.disconnect()
  if (viewObserver) viewObserver.disconnect() 
})

function loadMore() {
  if (isLoading.value || isLoadingMore.value || !hasMore.value) return
  loadPosts(currentPage.value + 1)
}

async function syncStickers(post) {
  if (!post.stickerIds.length) return
  try {
    const res = await fetch(
      `${BASE_URL}/api/v1/front/stickers/show?ids=${post.stickerIds.join(',')}`,
      { headers: { ...authHeaders() } }
    )
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json
    post.stickers = data.stickers ?? []
  } catch (e) {
    console.error('Failed to sync stickers', e)
  }
}

async function loadPosts(page) {
  const isFirstPage = page === 1
  if (isFirstPage) {
    isLoading.value = true
    error.value = null
  } else {
    isLoadingMore.value = true
  }

  try {
    const res = await fetch(
      `${BASE_URL}/api/v1/front/posts/show?page=${page}&perpage=${PER_PAGE}`,
      { headers: { ...authHeaders() } }
    )
    if (res.status === 401) {
      throw new Error('Expire jwt — please login again')
    }
    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }
    const json = await res.json()
    const payload = json?.data ?? json
    const rawList = payload?.posts ?? payload?.Posts ?? []
    const total = payload?.total ?? payload?.Total ?? 0
    const mapped = rawList.map(mapPost)
    const insertedCount = mapped.length
    posts.value = isFirstPage ? mapped : [...posts.value, ...mapped]
    currentPage.value = page
    console.log('🔍 raw posts payload:', JSON.stringify(rawList, null, 2)) 
    hasMore.value = posts.value.length < total && rawList.length > 0
    const startIdx = posts.value.length - insertedCount
    for (let i = startIdx; i < posts.value.length; i++) {
      syncLikes(posts.value[i])
      syncBookmark(posts.value[i])
      syncFollow(posts.value[i])
      syncStickers(posts.value[i]) 
    }
    
  } catch (e) {
    error.value = e.message || 'Failed to load posts'
  } finally {
    isLoading.value = false
    isLoadingMore.value = false
  }
}

function mapPost(p) {
  // console.log('🔍 raw post:', JSON.stringify(p, null, 2))
  return {
    id: p.id,
    // userID: p.user_id,
    userId: p.user_id,
    avatarUrl: resolveAvatarUrl(p.profile_images), 
    username: p.user_name || `User #${p.user_id}`,
    datetime: formatDatetime(p.created_at),
    description: buildDescription(p),
    translatedText: '',
    photos: buildPhotos(p),
    tags: buildTags(p),        
    videoPath: p.video_path ? resolveImageUrl(p.video_path) : null,  
    videoPlaying: false, 
    videoThumbnail: p.thumbnail_url ? resolveImageUrl(p.thumbnail_url) : null,
    videoDuration: p.duration ?? 0,
    views: p.views_count ?? 0,
    shareCount: 0,
    commentCount: p.comment_count ?? 0,
    likeCount: 0,
    isLiked: false,
    isBookmarked: false,
    bookmarkCount: 0,  
    isFollowing: false,
    postToGroup: true,
    showTranslated: false,
    showMore: false,
    likedByAvatars: [],
    reaction: null,
    showReactions: false,
    showComments: false,
    showLikers: false,  
    stickerIds: p.sticker_ids
      ? p.sticker_ids.split(',').map((s) => Number(s.trim())).filter(Boolean)
      : [],
    stickers: [],
    showSharePicker: false,
  }
}

function buildDescription(p) {
  switch (p.post_type) {
    case 'code':
      return p.code_content ?? ''
    case 'link':
      return p.link_url ?? ''
    default: 
      return p.caption ?? ''
  }
}

function buildPhotos(p) {
  if (!p.images) return []
  return p.images
    .split(',')
    .map((url) => url.trim())
    .filter(Boolean)
    .map(resolveImageUrl)
}

function buildTags(p) {
  if (!p.tag_name) return []
  return p.tag_name
    .split(',')
    .map((t) => t.trim())
    .filter(Boolean)
}

function resolveAvatarUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

function resolveImageUrl(url) {
  if (/^https?:\/\//i.test(url)) return url
  return `${BASE_URL}${url.startsWith('/') ? '' : '/'}${url}`
}

function formatDatetime(value) {
  if (!value) return ''
  const d = new Date(value)
  if (Number.isNaN(d.getTime())) return String(value)
  return d.toLocaleString('km-KH', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

async function toggleFollow(post) {
  const previous = { isFollowing: post.isFollowing }
  post.isFollowing = !post.isFollowing

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/followers/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...authHeaders(),
      },
      body: JSON.stringify({ user_id: post.userId }),
    })
    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }
    const json = await res.json()
    const data = json?.data ?? json
    if (typeof data?.is_following === 'boolean') {
      post.isFollowing = data.is_following
    }
  } catch (e) {
    console.error('Failed to update follow', e)
    post.isFollowing = previous.isFollowing
  }
}

const followStatusCache = new Map() 

async function syncFollow(post) {
  if (!post.userId) return
  const currentUserId = getCurrentUserId()
  if (String(post.userId) === String(currentUserId)) {
    post.isFollowing = false
    return
  }

  if (followStatusCache.has(post.userId)) {
    post.isFollowing = followStatusCache.get(post.userId)
    return
  }

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/followers/show?user_id=${post.userId}`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) {
      console.error(`syncFollow HTTP error for user ${post.userId}:`, res.status)
      return
    }
    const json = await res.json()
    const data = json?.data ?? json
    const isFollowing = data?.is_following ?? false

    followStatusCache.set(post.userId, isFollowing)
    post.isFollowing = isFollowing
  } catch (e) {
    console.error('Failed to sync follow', e)
  }
}

async function toggleBookmark(post) {
  const previous = {
    isBookmarked: post.isBookmarked,
    bookmarkCount: post.bookmarkCount ?? 0,
  }

  post.isBookmarked = !post.isBookmarked
  post.bookmarkCount = (post.bookmarkCount ?? 0) + (post.isBookmarked ? 1 : -1)

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/bookmarks/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...authHeaders(),
      },
      body: JSON.stringify({ post_id: post.id }),
    })
    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }
    const json = await res.json()
    const data = json?.data ?? json  
    const bookmarked = data?.bookmarked ?? data?.Bookmarked
    const total = data?.total ?? data?.Total
    if (typeof bookmarked === 'boolean') post.isBookmarked = bookmarked
    if (typeof total === 'number') post.bookmarkCount = total
  } catch (e) {
    console.error('Failed to update bookmark', e)
    post.isBookmarked = previous.isBookmarked
    post.bookmarkCount = previous.bookmarkCount
  }
}

async function syncBookmark(post) {
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/bookmarks/show?post_id=${post.id}`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) {
      console.error(`syncBookmark HTTP error for post ${post.id}:`, res.status)
      return
    }
    const json = await res.json()
    const data = json?.data ?? json
    const rawBookmarks = data.bookmarks ?? []
    const currentUserId = getCurrentUserId()
    post.isBookmarked = rawBookmarks.some(
      (b) =>
        String(b.user_id) === String(currentUserId) &&
        String(b.post_id ?? b.PostID) === String(post.id)
    )
    post.bookmarkCount = rawBookmarks.filter(
      (b) => String(b.post_id ?? b.PostID) === String(post.id)
    ).length
  } catch (e) {
    console.error('Failed to sync bookmark', e)
  }
}


function onComment(post) {
  post.showComments = !post.showComments
}

async function emitAndClose(post, action) {
  post.showMore = false
  if (action === 'hide') {
    try {
      const res = await fetch(`${BASE_URL}/api/v1/front/posts/delete/${post.id}`, {
        method: 'DELETE',
        headers: { ...authHeaders() },
      })
      if (res.ok) {
        posts.value = posts.value.filter((p) => p.id !== post.id)
      }
    } catch (e) {
      console.error('Failed to delete post', e)
    }
  } else {
    console.log(`Action "${action}" on post ${post.id}`)
  }
}

function formatCount(n) {
  if (n >= 1000) return (n / 1000).toFixed(n % 1000 === 0 ? 0 : 1) + 'k'
  return String(n)
}

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

function toggleLike(post) {
  const target = post.reaction
    ? REACTIONS.find((r) => r.key === post.reaction)
    : REACTIONS.find((r) => r.key === 'like')
  pickReaction(post, target)
}

async function pickReaction(post, reaction) {
  clearTimeout(reactionCloseTimer)
  post.showReactions = false

  const previous = {
    reaction: post.reaction,
    isLiked: post.isLiked,
    likeCount: post.likeCount,
  }
  const wasSame = post.reaction === reaction.key
  post.reaction = wasSame ? null : reaction.key
  post.isLiked = !wasSame
  post.likeCount += wasSame ? -1 : previous.reaction ? 0 : 1

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/likes/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...authHeaders(),
      },
      body: JSON.stringify({
        post_id: post.id,
        reaction_type: reaction.key,
      }),
    })
    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }
    const json = await res.json()
    const liked = (json?.data ?? json)?.liked
    if (typeof liked === 'boolean') {
      post.isLiked = liked
      post.reaction = liked ? reaction.key : null
    }
    await syncLikes(post)
  } catch (e) {
    console.error('Failed to update like', e)
    post.reaction = previous.reaction
    post.isLiked = previous.isLiked
    post.likeCount = previous.likeCount
  }
}

// async function syncLikes(post) {
//   try {
//     const res = await fetch(`${BASE_URL}/api/v1/front/likes/show?post_id=${post.id}`, {
//       headers: { ...authHeaders() },
//     })
//     if (!res.ok) return
//     const json = await res.json()
//     const data = json?.data ?? json
//     post.likeCount = data.total ?? 0
//     post.isLiked = data.liked_by_me ?? false
//     post.reaction = data.my_reaction ?? null
//     const rawLikes = data.likes ?? data.Likes ?? []
//     post.likedByAvatars = rawLikes
//       .slice(0, 3)
//       .map((l) => resolveAvatarUrl(l.profile_images))

//   } catch (e) {
//     console.error('Failed to sync likes', e)
//   }
// }


async function syncLikes(post) {
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/likes/show?post_id=${post.id}`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json
    post.likeCount = data.total ?? 0
    post.isLiked = data.liked_by_me ?? false
    post.reaction = data.my_reaction ?? null

    const rawLikes = data.likes ?? data.Likes ?? []
    // await fetchUsersInfo(userIds) 
    post.likedByAvatars = rawLikes.map((l) => ({
      avatarUrl: resolveAvatarUrl(l.profile_images),
      username: l.user_name || l.username || `User #${l.user_id}`,
    }))
  
  } catch (e) {
    console.error('Failed to sync likes', e)
  }
}

const postCardRefs = {}    
const viewedPostIds = new Set()   
let viewObserver = null

function setPostCardRef(el, postId) {
  if (el) {
    postCardRefs[postId] = el
    if (viewObserver) viewObserver.observe(el)
  }
}

function initViewObserver() {
  viewObserver = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          const postId = Number(entry.target.dataset.postId)
          if (!viewedPostIds.has(postId)) {
            viewedPostIds.add(postId)
            recordView(postId)
            viewObserver.unobserve(entry.target)  
          }
        }
      })
    },
    {
      threshold: 0.5,        
      rootMargin: '0px',
    }
  )
}

async function recordView(postId) {
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/posts/view/${postId}`, {
      method: 'POST',
      headers: { ...authHeaders() },
    })
    const json = await res.json()
    const data = json?.data ?? json
    if (typeof data?.views_count === 'number' && data.views_count >= 0) {
      const post = posts.value.find((p) => p.id === postId)
      if (post) post.views = data.views_count
    }
  } catch (e) {
    console.error('Failed to record view', e)
  }
}

const SHARE_OPTIONS = [
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

function closeAllSharePickers() {
  posts.value.forEach((p) => { p.showSharePicker = false })
}


let shareCloseTimer = null

function openSharePicker(post) {
  clearTimeout(shareCloseTimer)
  posts.value.forEach((p) => { if (p !== post) p.showSharePicker = false })
  post.showSharePicker = true
}

function scheduleCloseSharePicker(post) {
  clearTimeout(shareCloseTimer)
  shareCloseTimer = setTimeout(() => {
    post.showSharePicker = false
  }, 300)
}

function keepSharePickerOpen() {
  clearTimeout(shareCloseTimer)
}

async function quickShare(post, channel) {
  post.showSharePicker = false

  if (channel === 'native') {
    if (navigator.share) {
      try {
        await navigator.share({
          text: post.description,
          url: `${window.location.origin}/posts/${post.id}`,
        })
      } catch (e) {
        console.error('Native share cancelled or failed', e)
        return
      }
    } else {
      return
    }
  } else if (channel === 'copy') {
    const link = `${window.location.origin}/posts/${post.id}`
    try {
      await navigator.clipboard.writeText(link)
    } catch (e) {
      console.error('Copy failed', e)
    }
  } else {
    const link = encodeURIComponent(`${window.location.origin}/posts/${post.id}`)
    const text = encodeURIComponent(post.description || '')
    const deepLinks = {
      facebook: `https://www.facebook.com/sharer/sharer.php?u=${link}`,
      telegram: `https://t.me/share/url?url=${link}&text=${text}`,
      whatsapp: `https://wa.me/?text=${text}%20${link}`,
    }
    if (deepLinks[channel]) {
      window.open(deepLinks[channel], '_blank', 'noopener,noreferrer')
    }
  }

  post.shareCount = (post.shareCount ?? 0) + 1
}

function onTagClick(tag) {
  console.log('Clicked tag:', tag)
  // ឬ redirect ទៅ search page: router.push(`/search?tag=${tag}`)
}

function toggleSharePicker(post) {
  const wasOpen = post.showSharePicker
  posts.value.forEach((p) => { p.showSharePicker = false })
  post.showSharePicker = !wasOpen
}

function closeSharePicker(post) {
  post.showSharePicker = false
}




const likedByRefs = {}

function setLikedByRef(el, postId) {
  if (el) likedByRefs[postId] = el
}

function closeLikersOnClickOutside(event) {
  posts.value.forEach((post) => {
    if (!post.showLikers) return
    const el = likedByRefs[post.id]
    if (el && !el.contains(event.target)) {
      post.showLikers = false
    }
  })
}

function playVideo(post) {
  post.videoPlaying = true
}

function formatDuration(seconds) {
  if (!seconds) return ''
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m}:${String(s).padStart(2, '0')}`
}


// function hideViewerLikes() {
//   clearTimeout(hoverCloseTimer)
//   hoverCloseTimer = setTimeout(() => {
//     isViewerLikesOpen.value = false
//   }, 200)
// }





onMounted(async () => {
  document.addEventListener('click', closeLikersOnClickOutside) 
  document.addEventListener('click', closeAllSharePickers)
})

onUnmounted(() => {
  document.removeEventListener('click', closeLikersOnClickOutside)
  document.removeEventListener('click', closeAllSharePickers)
})
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.post-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin: 0 0 10px;
}

.tag-chip {
  display: inline-flex;
  align-items: center;
  background: #EFF6FB;
  border: 1.5px solid #CFE6F5;
  color: #1E6E9C;
  font-size: 12.5px;
  font-weight: 700;
  padding: 5px 10px;
  border-radius: 999px;
  font-family: 'Nunito', sans-serif;
  cursor: pointer;
}

.tag-chip:hover {
  background: #CFE6F5;
}

.post-card {
  background: #fff;
  border-radius: 16px;
  padding: 16px 18px;
  font-family: 'Inter', sans-serif;
  max-width: 720px;
  margin-top: 12px;

}

.post-top {
  display: flex;
  gap: 14px;
}

.avatar {
  width: 46px;
  height: 46px;
  border-radius: 50%;
  background: #EFF6FB;
  border: 2px solid #1976D2;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  overflow: hidden;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar svg {
  width: 24px;
  height: 24px;
  stroke: #1E6E9C;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.post-body {
  flex: 1;
  min-width: 0;
  margin: auto;
}

.post-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
}

.user-block {
  display: flex;
  flex-direction: column;
  gap: 4px;
  min-width: 0;
}

.username {
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 15.5px;
  color: #2B2B2B;
}

.datetime {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 12px;
  color: #8A8A8E;
  font-weight: 500;
}

.datetime svg {
  width: 13px;
  height: 13px;
  stroke: currentColor;
  fill: none;
  stroke-width: 2;
  stroke-linecap: round;
}

.follow-btn {
  border: 2px solid #1976D2;
  background: #1976D2;
  color: #fff;
  font-weight: 700;
  font-size: 13px;
  padding: 7px 18px;
  border-radius: 999px;
  cursor: pointer;
  flex-shrink: 0;
  font-family: 'Nunito', sans-serif;
}

.follow-btn.following {
  background: #fff;
  color: #8A8A8E;
}

.follow-btn:hover {
  filter: brightness(0.95);
}

.description {
  font-size: 14.5px;
  line-height: 1.6;
  color: #2B2B2B;
  margin: 10px 0 8px;
  white-space: pre-wrap;
}

.post-controls {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 4px;
}

.translate-btn {
  border: none;
  background: transparent;
  color: #1E6E9C;
  font-weight: 700;
  font-size: 12.5px;
  cursor: pointer;
  padding: 0;
  font-family: 'Nunito', sans-serif;
}

.group-toggle {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-size: 12.5px;
  font-weight: 600;
  color: #4A4A4E;
  cursor: pointer;
}

.group-toggle input {
  display: none;
}

.toggle-track {
  width: 30px;
  height: 17px;
  border-radius: 999px;
  background: #E7E7E7;
  position: relative;
  transition: background .15s;
  flex-shrink: 0;
}

.toggle-thumb {
  position: absolute;
  top: 2px;
  left: 2px;
  width: 13px;
  height: 13px;
  border-radius: 50%;
  background: #fff;
  transition: transform .15s;
  box-shadow: 0 1px 2px rgba(0, 0, 0, .3);
}

.group-toggle input:checked + .toggle-track {
  background: #1976D2;
}

.group-toggle input:checked + .toggle-track .toggle-thumb {
  transform: translateX(13px);
}

.post-foot {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
  border-top: 1px solid #EFE2D3;
  margin-top: 14px;
  padding-top: 12px;
  justify-content: space-between;
}

.post-left {
  display: flex;
  align-items: center;
}


.post-right {
  display: flex;
  align-items: center;
  gap: 8px;
}

.views {
  display: flex;
  align-items: center;
  gap: 5px;
  font-size: 13px;
  font-weight: 700;
  color: #4A4A4E;
  font-family: 'Nunito', sans-serif;
}

.views svg {
  width: 17px;
  height: 17px;
  stroke: currentColor;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.stat-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  border: 2px solid #1976D2;
  background: #1976D2;
  font-size: 12.5px;
  font-weight: 700;
  padding: 7px 12px;
  border-radius: 999px;
  cursor: pointer;
  color: #ffff;
  font-family: 'Nunito', sans-serif;
}

.stat-btn svg {
  width: 16px;
  height: 16px;
  stroke: currentColor;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.stat-btn:hover {
  opacity: 0.8;
}

.like-btn.liked {
  background: #F2762E;
  color: #fff;
  border-color: #D9601C;
}

.like-btn.liked svg {
  fill: #fff;
  stroke: #fff;
}

.bookmark-btn svg {
  fill: none;
}

.bookmark-btn.saved {
  background: #ffff;
  color: #D9601C;
  border-color: #F2C48B;
}

.bookmark-btn.saved svg {
  fill: #D9601C;
  stroke: #D9601C;
}

.more-wrap {
  position: relative;
}

.more-btn {
  background: #1976D2;
  color: #fff;
  padding: 7px 10px;
}

.more-btn:hover {
  opacity: 0.8;
}

.more-menu {
  position: absolute;
  bottom: 42px;
  right: 0;
  background: #fff;
  border: 1px solid #E7E7E7;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .12);
  padding: 6px;
  min-width: 150px;
  z-index: 10;
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

.liked-by {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #8A8A8E;
  font-weight: 600;
  margin-left: auto;
  pointer-events: auto;
  position: relative; 
}

.stack {
  display: flex;
}


.liked{
  font-size: 14px;
}

.stack span {
  width: 25px;
  height: 25px;
  border-radius: 50%;
  background: #EFF6FB;
  margin-left: -8px;
  overflow: hidden;
  display: block;
  border: 2px solid #8A8A8E;
}

.stack span:first-child {
  margin-left: 0;
}

.stack img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.photo-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
  margin-bottom: 10px;
}

.photo-grid.count-1 {
  grid-template-columns: 1fr;
}

.photo {
  position: relative;
  aspect-ratio: 4 / 3.2;
  border-radius: 12px;
  border: 2px solid #1976D2;
  overflow: hidden;
  background: #EFF6FB;
  display: flex;
  align-items: center;
  justify-content: center;
}

.photo img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.photo-placeholder {
  width: 34px;
  height: 34px;
  stroke: #1E6E9C;
  fill: none;
  stroke-width: 1.6;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.more-overlay {
  position: absolute;
  inset: 0;
  background: rgba(43, 43, 43, 0.55);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 18px;
  font-family: 'Nunito', sans-serif;
}

.feed {
  display: flex;
  flex-direction: column;
}

.state-msg {
  font-family: 'Inter', sans-serif;
  color: #4A4A4E;
  text-align: center;
  margin-top: 24px;
}

.state-msg.error {
  color: #C6402E;
}

.sentinel {
  height: 1px;
}

.like-wrap {
  position: relative;
  display: inline-flex;
}

.reaction-svg {
  display: inline-flex;
  width: 18px;
  height: 18px;
  color: currentColor;
}

.reaction-svg svg {
  width: 100%;
  height: 100%;
  display: block;
}

.like-btn[class*="reaction-"] {
  background: #fff;
  color: #F2762E;
  border-color: #F2C48B;
}

.reaction-picker {
  position: absolute;
  bottom: 46px;
  left: -240px;
  background: #fff;
  border-radius: 999px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .18);
  padding: 8px 10px;
  display: flex;
  align-items: center;
  gap: 4px;
  z-index: 30;
  white-space: nowrap;
  animation: popIn .15s ease;
}

@keyframes popIn {
  from { opacity: 0; transform: translateY(6px) scale(.9); }
  to   { opacity: 1; transform: translateY(0) scale(1); }
}

.reaction-option {
  position: relative;
  border: none;
  background: transparent;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform .12s ease;
  padding: 0;
  color: #2B2B2B;
}

.reaction-option .reaction-svg {
  width: 26px;
  height: 26px;
}

.reaction-option:hover {
  transform: scale(1.25) translateY(-4px);
}

.reaction-option.active {
  background: #1976D2;
}

.reaction-option.locked {
  background: #4A4A4E;
  color: #fff;
}

.lock-badge {
  position: absolute;
  bottom: -2px;
  right: -2px;
  font-size: 10px;
}

.reaction-tooltip {
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
  font-family: 'Nunito', sans-serif;
}

.reaction-option:hover .reaction-tooltip {
  opacity: 1;
}

.reaction-option.locked .reaction-tooltip {
  bottom: auto;
  top: calc(100% + 8px);
  white-space: normal;
  width: 220px;
  text-align: left;
  line-height: 1.4;
}

.comment-box {
  margin-top: 15px;
  border-top: 1px solid #eee;
  padding-top: 15px;
}

.liker-stack {
  display: flex;
  margin-left: auto;   
}

.liker-avatar {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: #EFF6FB;
  border: 2px solid #fff;
  box-shadow: 0 0 0 1px #E7E7E7;
  margin-left: -8px;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.liker-avatar:first-child {
  margin-left: 0;
}

.liker-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.liker-avatar svg {
  width: 14px;
  height: 14px;
  stroke: #1E6E9C;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.sticker-row {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 10px;
}

.post-sticker {
  width: 80px;
  height: 80px;
  object-fit: contain;
}

.share-wrap {
  position: relative;
  display: inline-flex;
}

.share-picker {
  position: absolute;
  bottom: 46px;
  left: -80px;
  background: #fff;
  border-radius: 999px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .18);
  padding: 8px 10px;
  display: flex;
  align-items: center;
  gap: 4px;
  z-index: 30;
  animation: popIn .15s ease;
  
}

.share-option {
  position: relative;
  border: none;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform .12s ease;
  padding: 0;
  pointer-events: auto;
}

.share-option:hover {
  transform: scale(1.15) translateY(-3px);
}

.share-option-svg {
  width: 20px;
  height: 20px;
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
  font-family: 'Nunito', sans-serif;
}

.share-option:hover .share-tooltip {
  opacity: 1;
}

/* .liked-by {
  position: relative;
  display: inline-flex;
  align-items: center;
  gap: 4px;
  cursor: pointer;
} */

.viewer-likes-popover {
  position: absolute;
  bottom: 100%;
  left: 0;
  margin-bottom: 8px;
  z-index: 30;
 
}

.main-warp {
  position: relative;
  width: 300px;
  /* background-color: #1976D2; */
  padding: 4px;
  /* border: 1px solid #1976D2; */
  border-radius: 20px;
}

.video-wrap {
  border-radius: 12px;
  overflow: hidden;
  margin-bottom: 10px;
  background: #000;
  width: 200px;
}

.post-video {
  width: 100%;
  max-height: 480px;
  display: block;
  border-radius: 10px;
}

/* .video-duration {
  position: absolute;
  right: 60px;

  bottom: -10px;

  background: rgba(0, 0, 0, 0.7);
  color: white;

  padding: 3px 8px;
  border-radius: 5px;

  font-size: 13px;
  font-weight: 500;
} */

.video-container{
  position:relative;
  display:flex;
  align-items:center;
  width: 200px;
}


.post-video{
  width:100%;
  border-radius:15px;
}


.video-controls{
  position:absolute;
  right:-55px;
  top:50%;
  transform:translateY(-50%);

  display:flex;
  flex-direction:column;
  gap:12px;
}


.video-controls button{
  width:40px;
  height:40px;

  border:none;
  background:#222;
  border-radius:50%;

  display:flex;
  align-items:center;
  justify-content:center;

  cursor:pointer;
}


.video-controls svg{
  width:22px;
  height:22px;
  fill:white;
}


.video-duration{
  position:absolute;
  right:10px;
  bottom:10px;

  background:rgba(0,0,0,.6);
  color:white;

  padding:3px 8px;
  border-radius:6px;

  font-size:13px;
}

</style>