<template>
  <div class="feed">
    <Posts v-if="!props.userId" @post="onNewPost" />

    <!-- Loading -->
    <p v-if="isLoading" class="state-msg">Loading...</p>
    <p v-else-if="error" class="state-msg error">{{ error }}</p>
    <p v-else-if="!posts.length" class="state-msg">No post yet</p>
    <div class="post-card" v-for="post in posts" :key="post._key" :data-post-id="post.id" :ref="(el) => setPostCardRef(el, post.id)">
    <div v-if="post.isRepost" class="repost-banner" @click="goToProfile(post.repostedByUserId)">
      <div class="repost-avatar">
        <img v-if="post.repostedByAvatar" :src="post.repostedByAvatar" alt="" />
        <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
      </div>
      <div class="repost-left">
        <svg viewBox="0 0 24 24" class="repost-icon"><path d="M4 12v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7"/><path d="M16 6l-4-4-4 4"/><path d="M12 2v14"/></svg>
        <span>{{ post.repostedByUsername }} was Share • {{ post.repostedAt }}</span>
      </div>
    </div>

    <div class="post-top">
     <div class="avatar-wrap" style="position: relative;" @mouseenter="openHoverCard(post.userId, post._key)" @mouseleave="scheduleCloseHoverCard">
    <div class="avatar" style="cursor: pointer;" @click="goToProfile(post.userId)">
      <img v-if="post.avatarUrl" :src="post.avatarUrl" alt="hashtag" />
      <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
    </div>

    <!-- Progile Hover -->
    <div v-if="hoverCardPostId === post._key" class="hover-card-anchor" @mouseenter="keepHoverCardOpen" @mouseleave="scheduleCloseHoverCard">
      <ProfileHoverCard :profile="hoverCardProfile" :is-loading="hoverCardLoading" :is-following="post.isFollowing" :is-own-profile="String(post.userId) === String(currentUserId)" @toggle-follow="toggleFollow(post)"
        @view-profile="goToProfile(post.userId)"/>
    </div>
    </div>

    <!-- Post body  -->
    <div class="post-body">
      <div class="post-head">
        <div class="user-block">
        <!-- <span class="username">{{ post.username }}</span> -->
        <span class="username" style="cursor: pointer; position: relative;" @click="goToProfile(post.userId)" @mouseenter="openHoverCard(post.userId, post._key)" @mouseleave="scheduleCloseHoverCard" >
            {{ post.fullName }}
        </span>

        <!-- data time  -->
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

       <!-- Quote share card on feed  -->
       <div v-if="post.isQuoteShare" class="quote-repost-card">
        <svg class="quote-mark" viewBox="0 0 24 24" fill="currentColor">
          <path d="M14.017 21v-7.391c0-5.704 3.731-9.57 8.983-10.609l.995 2.151c-2.432.917-3.995 3.638-3.995 5.849h4v10h-9.983zm-14.017 0v-7.391c0-5.704 3.748-9.57 9-10.609l.996 2.151c-2.433.917-3.996 3.638-3.996 5.849h3.998v10h-9.998z"/>
        </svg>
        <h4 class="quote-repost-title">{{ post.description }}</h4>
        <p class="quote-repost-text" v-html="post.quoteContent"></p>
      </div>

      <!-- Codeing show on post  -->
      <pre v-if="post.postType === 'code'" class="code-block">
        <code class="language-javascript">{{ getDisplayText(post) }}</code>
        <span 
          v-if="isTextTruncatable(post)" 
          class="see-more-btn code-see-more" 
          @click.stop="post.showFullText = !post.showFullText"
        >{{ post.showFullText ? ' See less' : ' See more' }}</span>
      </pre>

      <!-- Quote share des  -->
      <p v-else-if="!post.isQuoteShare" class="description">
        {{ getDisplayText(post) }}
        <span v-if="isTextTruncatable(post)" class="see-more-btn" @click.stop="post.showFullText = !post.showFullText">
          {{ post.showFullText ? ' See less' : ' See more' }}
        </span>
      </p>

      <!-- Tags chip show  -->
      <span class="tag-chip" v-for="(tag, i) in post.tags" :key="i" @click="onTagClick(tag)">
        <img v-if="splitTagIcon(tag).stickerUrl" :src="splitTagIcon(tag).stickerUrl" class="tag-icon-badge" alt="" />
        <span v-else-if="splitTagIcon(tag).icon" class="tag-icon-badge">{{ splitTagIcon(tag).icon }}</span>
        #{{ splitTagIcon(tag).text || tag }}
      </span>
        <div
        class="photo-grid media-wrap"
        v-if="post.photos.length"
        :class="post.photosExpanded ? 'expanded' : 'count-' + Math.min(post.photos.length, 4)"
      >
      
        <div
          class="photo"
          v-for="(photo, i) in (post.photosExpanded ? post.photos : post.photos.slice(0, 4))"
          :key="i"
          @click.stop.prevent="openLightbox(post, i)"
        >
          <img v-if="photo" :src="photo" alt="" />
          <svg v-else viewBox="0 0 24 24" class="photo-placeholder">
            <path d="M3 16l5-5 4 4 5-6 4 5"/>
            <circle cx="8" cy="8" r="1.6"/>
          </svg>
          <span
            v-if="i === 3 && !post.photosExpanded && post.photos.length > 4"
            class="more-overlay"
          >+{{ post.photos.length - 4 }}</span>
        </div>
      </div>
        <!-- Videos -->
        <div class="main-warp" v-if="post.videoPath"  :class="{ expanded: post.videoExpanded }" >
        <div class="video-container" :data-video-id="post.id" :ref="(el)=>observeVideo(el)">
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
          >
          </video>
          <!-- Center Play Button -->
          <button v-if="!post.isPlaying" class="play-center" @click="togglePlay(post)">
            <svg viewBox="0 0 24 24"><path d="M8 5v14l11-7z"/></svg>
          </button>
          <!-- Video Controls -->
          <div class="video-controls" v-show="post.showVideoControls">
            ...
          </div>
          <!-- Duration -->
          <div class="video-duration">{{ formatDuration(post.videoDuration) }}</div>

          <!--Progress Bar-->
        <div class="progress-track" @click.stop="seekVideo(post, $event)" @mousedown.stop="startSeekDrag(post, $event)">
          <div class="progress-fill" :style="{ width: post.videoProgress + '%' }"></div>
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
          <!-- <button
            v-if="post.communityId"
            type="button"
            class="community-badge"
            @click.stop="goToCommunity(post.communityId)"
          >
            <span class="community-badge-icon">
              <img v-if="post.communityAvatar" :src="post.communityAvatar" alt="" />
              <svg v-else viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 4a8 8 0 1 0 8 8h-2a6 6 0 1 1-1.76-4.24L14 10h6V4l-2.2 2.2A8 8 0 0 0 12 4Z"/>
              </svg>
            </span>
            <span class="community-badge-name">{{ post.communityName }}</span>
            <svg class="community-badge-chevron" viewBox="0 0 24 24"><path d="M9 6l6 6-6 6" stroke="currentColor" stroke-width="2.2" fill="none" stroke-linecap="round" stroke-linejoin="round"/></svg>
          </button> -->
        <!-- Likes Avatar  -->
        <div class="liked-by" v-if="post.likedByAvatars.length" @mouseenter="post.showLikers = true" :ref="(el) => setLikedByRef(el, post.id)">
         <div class="stack">
            <span v-for="(u, i) in post.likedByAvatars.slice(0, 3)" :key="i">
              <img v-if="u.avatarUrl" :src="u.avatarUrl" alt="" />
              <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="9" r="3.4"/><path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5"/></svg>
            </span>
         </div>
          <span>Liked</span>
        <ViewerLikes v-if="post.showLikers" class="viewer-likes-popover" :users="post.likedByAvatars"/>
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
            <span class="icon-circle"><svg viewBox="0 0 24 24"><path d="M4 12v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7"/><path d="M16 6l-4-4-4 4"/><path d="M12 2v14"/></svg></span>
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
         
        <button
          class="stat-btn like-btn"
          :class="{ liked: post.isLiked, [`reaction-${post.isQuoteShare ? post.quoteReactionTypeId : post.reaction}`]: post.isQuoteShare ? post.quoteReactionTypeId : post.reaction }"
          @click="toggleLike(post)"
        >
          <span class="icon-circle">
              <span
                  v-if="post.isQuoteShare ? post.quoteReactionTypeId != null : post.reaction"
                  class="reaction-svg reaction-emoji"
                  v-html="post.isQuoteShare ? getQuoteReactionIcon(post) : REACTIONS.find(r => r.key === post.reaction)?.icon">
              </span>

              <svg v-else viewBox="0 0 24 24">
                  <path d="M7 11v9H4v-9h3Zm3 9h8a2 2 0 0 0 2-2l1.5-5a2 2 0 0 0-2-2.6H15l.7-4A2 2 0 0 0 13.8 4L10 10v10Z"/>
              </svg>
          </span>
              <span class="count">
                  {{ formatCount(post.likeCount) }}
              </span>
          </button>

        <div class="reaction-picker" v-if="post.showReactions" @mouseenter="keepReactionPickerOpen" @mouseleave="scheduleCloseReactionPicker(post)"
          >
            <button v-for="r in reactionsForPost(post)" :key="r.key" class="reaction-option" :class="{ locked: r.private, active: (post.isQuoteShare ? post.quoteReactionTypeId : post.reaction) === r.key }" @click="pickReaction(post, r, $event)">
              <span class="reaction-svg" v-html="r.icon"></span>
              <span v-if="r.private" class="lock-badge">🔒</span>
              <span class="reaction-tooltip">{{ r.label }}</span>
            </button>
          </div>
        </div>
        <!-- Btn bookkmark -->
        <button class="stat-btn bookmark-btn" :class="{ saved: post.isBookmarked }" @click="toggleBookmark(post)">
          <span class="icon-circle"><svg viewBox="0 0 24 24"><path d="M6 3h12a1 1 0 0 1 1 1v17l-7-4-7 4V4a1 1 0 0 1 1-1Z"/></svg></span>
          {{ formatCount(post.bookmarkCount) }}  
        </button>
        <!-- Btn comments -->
        <button class="stat-btn" @click="onComment(post)">
          <span class="icon-circle"><svg viewBox="0 0 24 24"><path d="M21 12a8 8 0 1 1-3.2-6.4L21 4l-1 4.6A7.96 7.96 0 0 1 21 12Z"/></svg></span>
          {{ formatCount(post.commentCount) }}
        </button>
        <!-- Btn more on posts -->
        <div class="more-wrap">
          <button class="stat-btn more-btn" @click="post.showMore = !post.showMore">
            <span class="icon-circle"><svg viewBox="0 0 24 24" fill="currentColor"><circle cx="5" cy="12" r="1.6"/><circle cx="12" cy="12" r="1.6"/><circle cx="19" cy="12" r="1.6"/></svg></span>
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
    <!-- <PostsCard ref="postsCardRef" /> -->

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
        <span v-if="lightbox.reactions[idx]" class="lightbox-thumb-reaction" v-html="getLightboxReactionIcon(idx)"></span>
      </div>
      </div>
    </div>
  </Teleport>
  </div>
</template>

<script setup>
import { ref,computed, onMounted, onUnmounted,watch,nextTick } from 'vue'
import Comments from '../comments/Comments.vue'
import ViewerLikes from '../viewer_likes/ViewerLikes.vue'
import { useRouter } from 'vue-router'
import Posts from '../posts/Posts.vue'
import ProfileHoverCard from '../profile_hover_card/ProfileHoverCard.vue'
import gsap from 'gsap'
import hljs from 'highlight.js'
import 'highlight.js/styles/atom-one-dark.css'

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
const currentUserProfile = ref(null)



async function loadCurrentUserProfile() {
  const uid = getCurrentUserId()
  if (!uid) return
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/profile/show?id=${uid}`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json
    currentUserProfile.value = {
      username: data.user_name || 'You',
      avatarUrl: resolveAvatarUrl(data.profile_images),
    }
  } catch (e) {
    console.error('Failed to load current user profile', e)
  }
}

const highlightAllCode = () => {
  nextTick(() => {
    setTimeout(() => {
      document.querySelectorAll('pre.code-block code').forEach((block) => {
        block.removeAttribute('data-highlighted')
        hljs.highlightElement(block)
      })
    }, 50) 
  })
}

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


const groups = ref([])
const groupsLoading = ref(false)
let groupsFetched = false

async function fetchGroups() {
  if (groupsFetched || groupsLoading.value) return
  groupsLoading.value = true
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/communities/show?perpage=50`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json
    groups.value = (data.communities ?? []).map((g) => ({
      id: g.id,
      name: g.name,
      avatarUrl: resolveAvatarUrl(g.avatar_url),
      memberCount: g.member_count ?? 0,
      hotScore: g.hot_score ?? 0,
    }))
    groupsFetched = true
  } catch (e) {
    console.error('Failed to load groups', e)
  } finally {
    groupsLoading.value = false
  }
}

let groupPickerCloseTimer = null
function openGroupPicker(post) {
  clearTimeout(groupPickerCloseTimer)
  posts.value.forEach((p) => { if (p !== post) p.showGroupPicker = false })
  post.showGroupPicker = true
  fetchGroups()
}

function scheduleCloseGroupPicker(post) {
  clearTimeout(groupPickerCloseTimer)
  groupPickerCloseTimer = setTimeout(() => {
    post.showGroupPicker = false
  }, 250)
}

function keepGroupPickerOpen() {
  clearTimeout(groupPickerCloseTimer)
}

function selectGroup(post, group) {
  post.selectedGroupId = group ? group.id : null
  post.selectedGroupName = group ? group.name : null
  post.selectedGroupAvatar = group ? group.avatarUrl : null
  post.showGroupPicker = false
}

function formatHotScore(n) {
  return formatCount(Math.round(n))
}

let reactionCloseTimer = null
function getAuthToken() {
  return localStorage.getItem('token') || ''
}

function authHeaders() {
  const token = getAuthToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}

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
  fetchQuoteReactionTypes()
  initViewObserver()
  initVideoObserver()
  await loadPosts(1)
  highlightAllCode()
  currentUserId.value = getCurrentUserId()
  loadCurrentUserProfile() 
  observer = new IntersectionObserver(
    (entries) => {
      if (entries[0].isIntersecting) {
        loadMore()
      }
    },
    { rootMargin: '400px' } 
  )
  
  if (sentinel.value) observer.observe(sentinel.value)
  document.addEventListener('click', closeLikersOnClickOutside)
  document.addEventListener('click', closeAllSharePickers)
  document.addEventListener('click', closeAllSpeedMenus)
  document.addEventListener('keydown', handleLightboxKeydown)
})

onUnmounted(() => {
  highlightAllCode()
  if (observer) observer.disconnect()
  if (viewObserver) viewObserver.disconnect() 


  document.removeEventListener('click', closeLikersOnClickOutside)
  document.removeEventListener('click', closeAllSharePickers)
  document.removeEventListener('click', closeAllSpeedMenus)
  document.removeEventListener('keydown', handleLightboxKeydown)
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

const props = defineProps({
  userId: { type: [String, Number], default: null },
   mode: { type: String, default: 'feed' }
})

watch(
  () => props.userId,
  () => {
    posts.value = []
    currentPage.value = 1
    hasMore.value = true
    error.value = null
    viewedPostIds.clear() 
    loadPosts(1)
  }
)


let currentRequestId = 0
async function loadPosts(page) {
  const requestId = ++currentRequestId
  const isFirstPage = page === 1
  if (isFirstPage) {
    isLoading.value = true
    error.value = null
  } else {
    isLoadingMore.value = true
  }

  try {
    let url = `${BASE_URL}/api/v1/front/posts/show?page=${page}&perpage=${PER_PAGE}&feed_only=true`
    if (props.userId) {
      url += `&user_id=${props.userId}`
    }

    const res = await fetch(url, { headers: { ...authHeaders() }, cache: 'no-store' })

    if (requestId !== currentRequestId) return

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
    let mapped = rawList.map(mapPost)

    if (props.mode === 'popular') {
      mapped = mapped.sort((a, b) => {
        const scoreA = (a.views || 0) + (a.commentCount || 0) * 3
        const scoreB = (b.views || 0) + (b.commentCount || 0) * 3
        return scoreB - scoreA
      })
    }

    const insertedCount = mapped.length
    posts.value = isFirstPage ? mapped : [...posts.value, ...mapped]
    currentPage.value = page
    hasMore.value = posts.value.length < total && rawList.length > 0
    const startIdx = posts.value.length - insertedCount
    for (let i = startIdx; i < posts.value.length; i++) {
      syncLikes(posts.value[i])
      syncBookmark(posts.value[i])
      syncFollow(posts.value[i])
      syncStickers(posts.value[i])
      syncTagStickers(posts.value[i])
      syncShares(posts.value[i])
      syncQuoteReaction(posts.value[i])
    }
  } catch (e) {
    if (requestId === currentRequestId) {
      error.value = e.message || 'Failed to load posts'
    }
  } finally {
    if (requestId === currentRequestId) {
      isLoading.value = false
      isLoadingMore.value = false
    }
  }
}


let feedKeySeq = 0
function mapPost(p) {
     console.log('post raw data:', p.first_name, p.last_name, p.user_name)
  return {
    _key: `post-${feedKeySeq++}`,
    id: p.id,
    userId: p.user_id,
    avatarUrl: resolveAvatarUrl(p.profile_images), 
    username: p.user_name || `User #${p.user_id}`,
    fullName: [p.first_name, p.last_name].filter(Boolean).join(' ') || p.user_name || `User #${p.user_id}`,
    datetime: formatDatetime(p.created_at),
    description: buildDescription(p),
    translatedText: '',
    photos: buildPhotos(p),
    photosExpanded: false,
    tags: buildTags(p),        
    videoPath: p.video_path ? resolveImageUrl(p.video_path) : null,  
    videoThumbnail: p.thumbnail_path ? resolveImageUrl(p.thumbnail_path) : null, 
    videoDuration: p.duration ?? 0,
    videoCurrentTime: 0,
    videoProgress: 0,  
    isSeeking: false,  
    videoExpanded: false,
    isPlaying: false,
    showFullText: false,
    showVideoControls:false,
    isMuted: false,
    isMenuOpen: false,
    playbackRate: 1,
    views: p.views_count ?? 0,
    shareCount: p.share_count ?? 0,
    commentCount: p.comment_count ?? 0,
    likeCount: 0,
    isLiked: false,
    isBookmarked: false,
    bookmarkCount: 0,  
    isFollowing: false,
    postType: p.post_type || 'text',
    // postToGroup: true,
    selectedGroupId: null,
    selectedGroupName: null,
    selectedGroupAvatar: null,

    showGroupPicker: false,
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
    isRepost: !!p.repost_id,
    repostedByUserId: p.reposted_by_user_id ?? null,
    repostedByUsername: p.reposted_by_username ?? null,
    repostedAt: p.reposted_at ? formatDatetime(p.reposted_at) : '',
    repostedByAvatar: resolveAvatarUrl(p.reposted_by_profile_images),
    quoteContent: p.post_type === 'quote_share' ? (p.code_content ?? '') : null,
    isQuoteShare: p.post_type === 'quote_share', 
    quoteId: p.post_type === 'quote_share' ? p.community_id : null,
    quoteReactionLoading: false,
    quoteReactionTypeId: null, 

    communityId: p.community_id ?? p.group_id ?? null,
    communityName: p.community_name ?? p.group_name ?? null,
    communityAvatar: resolveAvatarUrl(p.community_avatar ?? p.community_avatar_url ?? p.group_avatar ?? ''),
  }
}

function addNewPost(response) {
  const raw = response?.data?.post || response?.data || response
  if (!raw || !raw.id) {
    console.error('Invalid new post payload', response)
    return
  }
  const newPost = mapPost(raw)
  posts.value.unshift(newPost)
  syncLikes(newPost)
  syncBookmark(newPost)
  syncFollow(newPost)
  syncStickers(newPost)
  syncShares(newPost)
    syncQuoteReaction(newPost)
}

function addOptimisticPost(post) {
  posts.value.unshift(post)
}

function onNewPost(post) {
  addOptimisticPost(post)
}

defineExpose({ addOptimisticPost })

function togglePhotosExpand(post) {
  post.photosExpanded = !post.photosExpanded
}

function buildDescription(p) {
  switch (p.post_type) {
    case 'code':
      return p.code_content ?? ''
    case 'link':
      return p.link_url ?? ''
       case 'quote':                
      return p.caption ?? ''
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
  if (!p.tag_data) return []
  return p.tag_data
    .split(',')
    .map(pair => pair.trim())
    .filter(Boolean)
    .map(pair => {
      const [name, stickerIdRaw] = pair.split('::')
      const stickerId = stickerIdRaw ? Number(stickerIdRaw) : null

      if (stickerId) {
        return { text: name, icon: '', stickerUrl: '', stickerId }
      }
      const m = name.match(/^(\p{Extended_Pictographic}\uFE0F?)/u)
      return m
        ? { text: name.slice(m[1].length), icon: m[1], stickerUrl: '', stickerId: null }
        : { text: name, icon: '', stickerUrl: '', stickerId: null }
    })
}


async function syncTagStickers(post) {
  const ids = post.tags.filter(t => t.stickerId).map(t => t.stickerId)
  if (!ids.length) return
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/stickers/show?ids=${ids.join(',')}`,
      { headers: { ...authHeaders() } })
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json
    const stickerMap = new Map((data.stickers ?? []).map(s => [s.id, s.url]))
    post.tags.forEach(t => {
      if (t.stickerId && stickerMap.has(t.stickerId)) {
        t.stickerUrl = resolveImageUrl(stickerMap.get(t.stickerId))
      }
    })
  } catch (e) {
    console.error('Failed to sync tag stickers', e)
  }
}

function resolveAvatarUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

function resolveImageUrl(url) {
  if (!url) return ''
  if (/^https?:\/\//i.test(url) || url.startsWith('data:image/')) {
    return url
  }
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
      // body: JSON.stringify({ post_id: Number(post.id) }),
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
        delete videoRefs[post.id] 
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
  if (post.isQuoteShare) {
    toggleQuoteReaction(post)
    return
  }
  const target = post.reaction
    ? REACTIONS.find((r) => r.key === post.reaction)
    : REACTIONS.find((r) => r.key === 'like')
  pickReaction(post, target, event)
}

const quoteReactionTypes = ref([])
async function fetchQuoteReactionTypes() {
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/reaction-types`, { headers: { ...authHeaders() } })
    if (!res.ok) return
    const json = await res.json()
    quoteReactionTypes.value = json?.data ?? json ?? []
  } catch (e) {
    console.error('Failed to load quote reaction types', e)
  }
}
const defaultQuoteReactionId = computed(() =>
  (quoteReactionTypes.value.find(r => r.name?.toLowerCase() === 'like') || quoteReactionTypes.value[0])?.id ?? null
)

async function toggleQuoteReaction(post) {
  if (post.quoteReactionLoading || !post.quoteId || defaultQuoteReactionId.value == null) return
  post.quoteReactionLoading = true
  const wasLiked = post.isLiked
  post.isLiked = !wasLiked
  post.likeCount += wasLiked ? -1 : 1

  try {
    if (!wasLiked) {
      await fetch(`${BASE_URL}/api/v1/front/quote-reactions/create`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', ...authHeaders() },
        body: JSON.stringify({ quote_id: post.quoteId, reaction_type_id: defaultQuoteReactionId.value }),
      })
      post.quoteReactionTypeId = defaultQuoteReactionId.value 
    } else {
      await fetch(`${BASE_URL}/api/v1/front/quote-reactions/${post.quoteId}`, {
        method: 'DELETE',
        headers: { ...authHeaders() },
      })
      post.quoteReactionTypeId = null 
    }
  } catch (e) {
    console.error('Failed to toggle quote reaction', e)
    post.isLiked = wasLiked
    post.likeCount += wasLiked ? 1 : -1
  } finally {
    post.quoteReactionLoading = false
  }
}

async function pickReaction(post, reaction) {
  if (post.isQuoteShare) {
    await selectQuoteReaction(post, reaction)
    return
  }

  clearTimeout(reactionCloseTimer)
  post.showReactions = false

  if (post.reaction !== reaction.key && event) {
    triggerFloatingReaction(event, reaction.icon)
  }

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
      headers: { 'Content-Type': 'application/json', ...authHeaders() },
      body: JSON.stringify({ post_id: post.id, reaction_type: reaction.key }),
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



let lightboxReactionCloseTimer = null

const lightboxCurrentReaction = computed(() => {
  const key = lightbox.value.reactions[lightbox.value.activeIndex]
  return key ? REACTIONS.find(r => r.key === key) : null
})

function openLightboxReactionPicker() {
  clearTimeout(lightboxReactionCloseTimer)
  lightbox.value.showReactions = true
}

function scheduleCloseLightboxReactionPicker() {
  clearTimeout(lightboxReactionCloseTimer)
  lightboxReactionCloseTimer = setTimeout(() => {
    lightbox.value.showReactions = false
  }, 300)
}

function keepLightboxReactionPickerOpen() {
  clearTimeout(lightboxReactionCloseTimer)
}

function toggleLightboxReaction() {
  const existing = lightbox.value.reactions[lightbox.value.activeIndex]
  if (existing) {
    pickLightboxReaction({ key: existing }) // ចុចម្ដងទៀត = toggle off
  } else {
    openLightboxReactionPicker()
  }
}

async function pickLightboxReaction(reaction, event) {
  const idx = lightbox.value.activeIndex
  const previous = lightbox.value.reactions[idx]
  const wasSame = previous === reaction.key
  lightbox.value.showReactions = false

  if (!wasSame && event) {
    const icon = REACTIONS.find(r => r.key === reaction.key)?.icon
    if (icon) triggerFloatingReaction(event, icon)
  }

  if (wasSame) {
    delete lightbox.value.reactions[idx]
  } else {
    lightbox.value.reactions[idx] = reaction.key
  }
  lightbox.value.reactions = { ...lightbox.value.reactions } 
  try {
    await fetch(`${BASE_URL}/api/v1/front/photo-reactions/create`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', ...authHeaders() },
      body: JSON.stringify({
        post_id: lightbox.value.postId,
        photo_index: idx,
        reaction_type: wasSame ? null : reaction.key,
      }),
    })
  } catch (e) {
    console.error('Failed to save photo reaction', e)
  }
}

function getLightboxReactionIcon(idx) {
  const key = lightbox.value.reactions[idx]
  if (!key) return ''
  return REACTIONS.find(r => r.key === key)?.icon || ''
}

function getQuoteReactionIcon(post) {
  if (!post.isQuoteShare || post.quoteReactionTypeId == null) return null
  const match = quoteReactionTypesMapped.value.find(r => r.key === post.quoteReactionTypeId)
  return match ? match.icon : null
}

const quoteReactionTypesMapped = computed(() =>
  quoteReactionTypes.value.map(rt => ({
    key: rt.id,
    label: rt.name,
    icon: rt.icon_value,
  }))
)

function reactionsForPost(post) {
  return post.isQuoteShare ? quoteReactionTypesMapped.value : REACTIONS
}

async function selectQuoteReaction(post, reaction) {
  if (post.quoteReactionLoading || !post.quoteId) return
  post.quoteReactionLoading = true
  post.showReactions = false

  if (event) triggerFloatingReaction(event, reaction.icon)

  const previous = {
    isLiked: post.isLiked,
    likeCount: post.likeCount,
    quoteReactionTypeId: post.quoteReactionTypeId,
  }

  const hadReactionBefore = !!post.isLiked
  post.isLiked = true
  post.quoteReactionTypeId = reaction.key
  if (!hadReactionBefore) post.likeCount += 1

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/quote-reactions/create`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', ...authHeaders() },
      body: JSON.stringify({ quote_id: post.quoteId, reaction_type_id: reaction.key }),
    })
    if (!res.ok) throw new Error(`API ${res.status}`)
  } catch (e) {
    console.error('Failed to select quote reaction', e)
    post.isLiked = previous.isLiked
    post.likeCount = previous.likeCount
    post.quoteReactionTypeId = previous.quoteReactionTypeId
  } finally {
    post.quoteReactionLoading = false
  }
}

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
  console.log('🔍 recordView called for postId:', postId)
  if (!postId || isNaN(Number(postId)) || String(postId).startsWith('temp-')) {
    return;
  }
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

function closeAllSharePickers() {
  posts.value.forEach((p) => { p.showSharePicker = false })
}

function closeAllSpeedMenus() {
  posts.value.forEach((p) => { p.isMenuOpen = false })
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


async function quickShare(post, key) {
  post.showSharePicker = false
  if (key === 'internal') {
    try {
      const res = await fetch(`${BASE_URL}/api/v1/front/posts/shares/create`, {
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
      post.shareCount = (post.shareCount || 0) + 1
      const currentUserIdVal = getCurrentUserId()
      const repostedPost = {
        ...post,
        _key: `post-${feedKeySeq++}`,
        id: 'repost-' + Date.now(),
        isRepost: true,
        repostedByUserId: currentUserIdVal,
         repostedByUsername: currentUserProfile.value?.username || 'You',
         repostedByAvatar: currentUserProfile.value?.avatarUrl || '',
        repostedAt: formatDatetime(new Date()),
        shareCount: 0,
        likeCount: 0,
        isLiked: false,
        reaction: null,
        commentCount: 0,
        showComments: false,
      }
      posts.value.unshift(repostedPost)
    } catch (e) {
      console.error('Failed to share post to profile', e)
    }
    return
  }

  if (key === 'native') {
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
  } else if (key === 'copy') {
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
    if (deepLinks[key]) {
      window.open(deepLinks[key], '_blank', 'noopener,noreferrer')
    }
  }
}

function onTagClick(tag) {
  console.log('Clicked tag:', tag)
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

function formatDuration(seconds) {
  if (!seconds) return ''
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m}:${String(s).padStart(2, '0')}`
}


const videoRefs = {}
function setVideoRef(el, postId) {
  if (el) videoRefs[postId] = el
  else delete videoRefs[postId] 
}

function togglePlay(post) {
  const el = videoRefs[post.id]
  if (!el) return
  if (el.paused) {
    pauseAllVideosExcept(post.id)
    el.play()
      .then(() => { post.isPlaying = true })
      .catch((e) => {
        console.error('Failed to play video', e)
        post.isPlaying = false
      })
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

function fullscreen(post) {
  const el = videoRefs[post.id]
  if (el) el.requestFullscreen()
}

function setSpeed(post, rate) {
  post.playbackRate = rate
  const el = videoRefs[post.id]
  if (el) el.playbackRate = rate
  post.isMenuOpen = false
}

function setVideoDuration(post, event) {
  post.videoDuration = event.target.duration
}

const videoObserver = ref(null)
function initVideoObserver(){
  videoObserver.value = new IntersectionObserver(
    (entries)=>{
      entries.forEach(entry=>{
        const postId = Number(entry.target.dataset.videoId)
        const post = posts.value.find(p=>p.id === postId)
        const video = videoRefs[postId]
        if(!video || !post) return
        if(entry.isIntersecting){
          pauseAllVideosExcept(postId)
          //  video.muted = true  
          video.play().catch(() => {})
          post.isPlaying = true
        } else {
          video.pause()
          post.isPlaying = false
        }
      })
    },
    {
      threshold: 0.7
    }
  )
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

function observeVideo(el){
  if(!el) return
  if(videoObserver.value){
    videoObserver.value.observe(el)
  }

}

function openFullscreen(post) {
  const video = videoRefs[post.id]
  if (!video) return
  if (!document.fullscreenElement) {
    if (video.requestFullscreen) {
      video.requestFullscreen()
    } 
    else if (video.webkitRequestFullscreen) {
      video.webkitRequestFullscreen()
    }
  } else {
    document.exitFullscreen()

  }
}

function expandVideo(post){
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

const router = useRouter()
function goToProfile(userId) {
  if (!userId) {
    console.warn('No userId provided')
    return
  }
  router.push(`/profile/${userId}`) 
}

function goToCommunity(communityId) {
  if (!communityId) return
  router.push({ name: 'GroupDetail', params: { id: communityId } })
}

function goToProfileByUsername(username) {
  if (!username) {
    console.warn('No username provided');
    return;
  }

  const post = posts.value.find(p => 
    p.likedByAvatars.some(u => u.username === username)
  );
  
  if (post) {
    const user = post.likedByAvatars.find(u => u.username === username);
    if (user && user.userId) {
      goToProfile(user.userId);
    } else {
      window.location.href = `/profile?username=${encodeURIComponent(username)}`;
    }
  } else {
    window.location.href = `/profile?username=${encodeURIComponent(username)}`;
  }
}

async function syncShares(post) {
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/posts/shares/show?post_id=${post.id}`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json
    post.shareCount = data.share_count ?? data.shareCount ?? 0
  } catch (e) {
    console.error(`syncShares HTTP error for post ${post.id}:`, e)
  }
}

function handleCommentUserClick(userId) {
  goToProfile(userId);
}

const TEXT_LIMIT = 100 

function isTextTruncatable(post) {
  const text = post.showTranslated ? post.translatedText : post.description
  return text && text.length > TEXT_LIMIT
}

function getDisplayText(post) {
  const text = post.showTranslated ? post.translatedText : post.description
  if (!text) return ''
  if (post.showFullText || text.length <= TEXT_LIMIT) {
    return text
  }
  return text.slice(0, TEXT_LIMIT) + '...'
}

const hoverCardPostId = ref(null) 
const hoverCardProfile = ref(null)
const hoverCardLoading = ref(false)
const profileCache = new Map()
let hoverCloseTimer = null

function openHoverCard(userId, postId) {
  clearTimeout(hoverCloseTimer)
  if (hoverCardPostId.value === postId) return
  hoverCardPostId.value = postId

  if (profileCache.has(userId)) {
    hoverCardProfile.value = profileCache.get(userId)
    hoverCardLoading.value = false
    return
  }

  hoverCardProfile.value = null
  hoverCardLoading.value = true
  fetchHoverProfile(userId, postId)
}

async function fetchHoverProfile(userId, postId) {
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/profile/show?id=${userId}`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) throw new Error('Failed')
    const json = await res.json()
    const data = json?.data ?? json
    const mapped = {
      id: data.id ?? userId, 
      name: `${data.first_name || ''} ${data.last_name || ''}`.trim() || data.user_name,
      handle: '@' + data.user_name,
      avatarUrl: resolveAvatarUrl(data.profile_images),
      coverUrl: resolveAvatarUrl(data.cover_images),
      bio: data.bios ?? '',
      location: data.location ?? '',
      joinedDate: data.created_at
        ? 'Joined ' + new Date(data.created_at).toLocaleString('en-US', { month: 'long', year: 'numeric' })
        : '',
    }
    profileCache.set(userId, mapped)
    if (hoverCardPostId.value === postId) {
      hoverCardProfile.value = mapped
    }
  } catch (e) {
    console.error('Failed to load hover profile', e)
    if (hoverCardPostId.value === postId) hoverCardProfile.value = null
  } finally {
    if (hoverCardPostId.value === postId) hoverCardLoading.value = false
  }
}

function scheduleCloseHoverCard() {
  clearTimeout(hoverCloseTimer)
  hoverCloseTimer = setTimeout(() => {
    hoverCardPostId.value = null
  }, 200)
}

function keepHoverCardOpen() {
  clearTimeout(hoverCloseTimer)
}

const lightbox = ref({
  open: false,
  images: [],
    postId: null,   
  activeIndex: 0,
  rotation: 0,
    showReactions: false,
    reactions: {},
})

function openLightbox(post, index) {
  if (!post.photos || !post.photos.length) return
  lightbox.value.images = post.photos
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

function triggerFloatingReaction(event, emojiIcon) {
  const floatingEl = document.createElement('div')
  floatingEl.className = 'floating-reaction-emoji'
  floatingEl.innerHTML = emojiIcon
  const x = event.clientX
  const y = event.clientY

  floatingEl.style.position = 'fixed'
  floatingEl.style.left = `${x - 15}px`
  floatingEl.style.top = `${y - 15}px`
  floatingEl.style.pointerEvents = 'none' 
  floatingEl.style.zIndex = '9999'
  floatingEl.style.width = '32px'
  floatingEl.style.height = '32px'

  document.body.appendChild(floatingEl)
  gsap.fromTo(
    floatingEl,
    {
      scale: 0.2,
      opacity: 1,
      y: 0,
      x: 0,
      rotation: 0
    },
    {
      scale: 1.8,                
      y: -120,                   
      x: (Math.random() - 0.5) * 40, 
      rotation: (Math.random() - 0.5) * 30,
      opacity: 0,            
      duration: 1.2,               
      ease: 'power2.out',
      onComplete: () => {
        floatingEl.remove()       
      }
    }
  )
}

function splitTagIcon(tag) {
  if (tag && typeof tag === 'object') {
    return {
      icon: tag.icon || '',
      text: tag.text || '',
      stickerUrl: tag.stickerUrl || '',
    }
  }

  if (typeof tag !== 'string') {
    return { icon: '', text: String(tag ?? ''), stickerUrl: '' }
  }

  const m = tag.match(/^(\p{Extended_Pictographic}\uFE0F?)/u)
  if (m) return { icon: m[1], text: tag.slice(m[1].length), stickerUrl: '' }
  return { icon: '', text: tag, stickerUrl: '' }
}

async function syncQuoteReaction(post) {
  if (!post.isQuoteShare || !post.quoteId) return
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/quote-reactions/${post.quoteId}`, {
      headers: { ...authHeaders() },
    })
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json

    if (data && data.reaction_type_id != null) {
      post.isLiked = true
      post.quoteReactionTypeId = data.reaction_type_id
    } else {
      post.isLiked = false
      post.quoteReactionTypeId = null
    }
    if (typeof data?.total === 'number') {
      post.likeCount = data.total
    }
  } catch (e) {
    console.error('Failed to sync quote reaction', e)
  }
}

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

.hover-card-anchor {
  position: absolute;
  top: calc(100% + 6px);
  left: 0;
  z-index: 100;
}


.tag-chip {
  display: inline-flex;
  align-items: center;
  gap: 2px;
 background: #3867f3;
  color: #ffffff;
  font-size: 13px;
  font-weight: 700;
  padding: 2px 8px 2px 2px;
  border-radius: 4px;
  margin-right: 8px; 
  margin-bottom: 8px;
}

.tag-chip:hover {
  opacity: 0.9;
}

.tag-icon-badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 22px;
  height: 22px;
  margin-right: 5px;
  font-size: 22px;
  line-height: 1;
  border-radius: 2px;
  background-color: #ffff;

}


.post-card {
  background: #FFFFFF;
  border-radius: 16px;
  padding: 16px 30px;
  font-family: 'Inter', sans-serif;
  max-width: 720px;
  margin-top: 22px;
  border: 1.5px solid #E5E7EB;
  border: 1px solid #E5E7EB;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.04);
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

.see-more-btn {
  color: #1976D2;
  font-weight: 700;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
  white-space: nowrap;
}

.see-more-btn:hover {
  text-decoration: underline;
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
  border: 2px solid #eae7e1;
  background: #eae7e1;
  font-size: 12.5px;
  font-weight: 700;
  padding: 7px 12px;
  border-radius: 32px;
  cursor: pointer;
  color: #3f3a34;
  font-family: 'Nunito', sans-serif;
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
  color: #3f3a34;
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
  background: #1976d218;
  color: #1976D2;
  border-color: #1976d218;
}

.bookmark-btn.saved svg {
  fill: #1976D2;
  stroke: #1976D2;
}

.more-wrap {
  position: relative;
}

.more-btn {
  /* background: #F4F1EC; */
  color: #3f3a34;
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
  border: 2px solid #EAE7E1;
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
  grid-template-columns: repeat(2, minmax(0, 1fr));
  transition: all 0.3s ease;
  max-width: 320px; 
  gap: 8px;
  margin-bottom: 10px;
}

.photo-grid.count-1 {
  grid-template-columns: 1fr !important;
  max-width: 320px !important; 
}

.photo-grid.count-4 {
  grid-template-columns: repeat(2, minmax(0, 1fr)) !important;
  max-width: 320px !important;
}


.photo-grid.expanded {
  display: grid !important;
  grid-template-columns: 1fr !important;
  row-gap: 12px !important; 
  max-width: 100% !important;
}


 .photo-grid.count-1.expanded {
  row-gap: 12px !important;
  max-width: 500px !important;
  grid-template-columns: 1fr !important;
}

.photo-grid .photo:nth-child(n+5) {
  display: none;
  gap: 8px;
}

.photo-grid.count-4.expanded {
  gap: 8px;
  max-width: 100% !important;
  grid-template-columns: 1fr !important;
  
}

.photo {
  position: relative;
  width: 100%;
  aspect-ratio: 16 / 11; 
  border-radius: 12px;
  overflow: hidden;
  background: #EFF6FB;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: aspect-ratio 0.3s ease;
  
}

.photo-grid.count-1 .photo {
  aspect-ratio: 4 / 3; 
  
}

.photo-grid.count-4 .photo {
  aspect-ratio: 1 / 1; 
  
}

.photo img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.photo-grid.expanded .photo {
  aspect-ratio: 16 / 10; 
  width: 100%;
  
}

.photo-grid.count-1.expanded .photo {
  aspect-ratio: 16 / 10;
  
}

.photo-grid.count-4.expanded .photo {
  
  aspect-ratio: 16 / 11;
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
  pointer-events: none;
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
  background: #1976d218;
  color: #1976D2;
  border-color: #1976d218;
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
  padding: 4px;
  border-radius: 20px;
  transition: width .3s ease;
}


/* click expand */
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
  width:100%;
  border-radius:12px;
  display:block;
  transition:.3s ease;
}


.post-video.expanded {

  width:100%;
  height:500px;

  object-fit:contain;

  background:#000;

}

.video-controls {

  position:absolute;

  bottom:12px;
  left:50%;

  transform:translateX(-50%);

  display:flex;
  gap:8px;

  background:rgba(0,0,0,0.45);

  padding:6px 10px;

  border-radius:20px;

  z-index:20;
}

.video-controls button {

  width:32px;
  height:32px;

  border:none;
  background:transparent;

  display:flex;
  align-items:center;
  justify-content:center;

  cursor:pointer;
}

.video-controls svg {

  width:22px;
  height:22px;

  fill:white;
}

.video-duration {
  position:absolute;
  right:10px;
  top: 10px;
  /* bottom:15px; */
  background:rgba(0,0,0,.6);
  color:white;
  padding:3px 8px;

  border-radius:6px;

  font-size:13px;
}

.more-wrap { position: relative; }

.speed-menu {
  position: absolute;
  left: 46px;
  top: 50%;
  transform: translateY(-50%);
  background: #222;
  border-radius: 10px;
  padding: 6px;
  display: flex;
  flex-direction: column;
  gap: 2px;
  z-index: 20;
}

.speed-menu button {
  all: unset;
  box-sizing: border-box;
  width: 44px;
  text-align: center;
  padding: 6px 8px;
  border-radius: 6px;
  font-size: 12px;
  color: #fff;
  cursor: pointer;
}

.speed-menu button:hover { background: rgba(255,255,255,0.12); }
.speed-menu button.active { color: #F2762E; font-weight: 700; }



.play-center {
  position: absolute;

  top: 50%;
  left: 50%;

  transform: translate(-50%, -50%);

  width: 60px;
  height: 60px;

  border: none;
  border-radius: 50%;

  background: rgba(0,0,0,0.55);
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

.repost-banner {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  font-weight: 600;
  color: #8A8A8E;
  padding-bottom: 10px;
  margin-bottom: 10px;
  border-bottom: 1px solid #F0F0F0;
  cursor: pointer;
}
.repost-left {
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 0;
}

.repost-avatar {
  width: 26px;
  height: 26px;
  border-radius: 50%;
  background: #EFF6FB;
  border: 1.5px solid #1976D2;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  overflow: hidden;
}

.repost-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.repost-avatar svg {
  width: 15px;
  height: 15px;
  stroke: #1E6E9C;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
  stroke-linejoin: round;
}
.repost-icon {
  width: 16px; height: 16px;
  stroke: #8A8A8E; fill: none; stroke-width: 1.8;
  flex-shrink: 0;
}
.repost-banner:hover { color: #1976D2; }
.repost-banner:hover .repost-icon { stroke: #1976D2; }

.floating-reaction-emoji svg {
  width: 100%;
  height: 100%;
  display: block;
}

.code-block {
  background: #1e1e2e;
  color: #cdd6f4;
  font-family: 'Fira Code', Consolas, Monaco, 'Andale Mono', 'Ubuntu Mono', monospace;
  font-size: 13px;
  line-height: 1.6;
  padding: 14px 16px;
  border-radius: 10px;
  overflow-x: auto;
  white-space: pre-wrap;
  word-break: break-word;
  margin: 8px 0;
}

.code-block code {
  background: none;
  color: inherit;
  font-family: inherit;
}

.code-see-more {
  position: absolute;
  bottom: 8px;
  right: 12px;
  background: rgba(40, 44, 52, 0.9);
  padding: 2px 8px;
  border-radius: 4px;
  color: #61afef;
  cursor: pointer;
}

.hljs-keyword,
.hljs-selector-tag {
  color: #c678dd !important; 
  font-weight: 600;
}

.hljs-title.function_,
.hljs-title.class_,
.hljs-attr {
  color: #61afef !important; 
}

.hljs-params {
  color: #abb2bf !important;
}

.hljs-string {
  color: #98c379 !important; 
}

.hljs-number,
.hljs-literal,
.hljs-boolean {
  color: #d19a66 !important; 
}

.hljs-operator,
.hljs-punctuation {
  color: #56b6c2 !important; 
}

/* Comments (ពណ៌ប្រផេះ) */
.hljs-comment {
  color: #5c6370 !important;
  font-style: italic;
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
  /* color: #000; */
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

.tag-icon-badge { margin-right: 2px; font-size: 13px; }


.quote-repost-card {
  background: #F8FAFC;
  /* border: 1px solid #E2E8F0; */
  border-radius: 12px;
  padding: 16px;
  margin: 8px 0;
  position: relative;
  overflow-wrap: break-word;
  word-break: break-word;

}
.quote-mark {
  width: 20px;
  height: 20px;
  color: #1976D2;
  opacity: 0.6;
  margin-bottom: 6px;
}
.quote-repost-title {
  margin: 0 0 6px;
  font-size: 14.5px;
  font-weight: 700;
  color: #0F172A;
}
.quote-repost-text {
  margin: 0;
  font-style: italic;
  color: #475569;
  line-height: 1.6;
}

.lightbox-reaction-bar {
  display: flex;
  justify-content: center;
  padding: 4px 0 12px;
}

.lightbox-reaction-bar .like-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
}

.lightbox-reaction-picker {
  position: static;
  top: auto;
  left: auto;
  bottom: auto;
  transform: none;
  background: rgba(0, 0, 0, 0.85);
  box-shadow: none;
  animation: none;
}

.lightbox-reaction-picker .reaction-option {
  color: #fff;
}

.lightbox-reaction-picker .reaction-tooltip {
  background: #000;
}

.lightbox-thumb {
  position: relative; /* ត្រូវបន្ថែម បើ rule ដើមមិនទាន់មាន */
}

.lightbox-thumb-reaction {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 18px;
  height: 18px;
  background: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 1px 3px rgba(0, 0, 0, .3);
}

.community-badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  border: none;
  background: #EAF1FE;
  color: #1976D2;
  padding: 6px 12px 6px 8px;
  border-radius: 999px;
  cursor: pointer;
  font-size: 13px;
  font-weight: 700;
  font-family: 'Nunito', sans-serif;
  max-width: 260px;
}

.community-badge:hover {
  background: #DCE9FD;
}

.community-badge-icon {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  overflow: hidden;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #1976D2;
  color: #fff;
}

.community-badge-icon img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.community-badge-icon svg {
  width: 13px;
  height: 13px;
}

.community-badge-name {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.community-badge-chevron {
  width: 14px;
  height: 14px;
  flex-shrink: 0;
}

.lightbox-thumb-reaction svg { width: 12px; height: 12px; }
</style>