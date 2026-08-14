<template>
  <div>
    <!-- Sticky NavBar Wrapper -->
    <div class="navbar-wrapper">
      <NavBar/>
    </div>

    <div class="group-detail-wrapper">
      <div class="group-detail-container">
        
        <!-- ================= Left / Main Column ================= -->
        <div class="main-content">
          
          <!-- Card ទី 1: Group Info Card -->
          <div class="card group-banner-card">
            <div class="header-top">
              <div class="group-brand">
                
                <!-- Group Thumbnail (ស្ទីលបែបខ្មែរ) -->
                <div class="group-avatar khmer-theme">
                  <span class="avatar-flag">🇰🇭</span>
                  <span class="avatar-title">បច្ចេកវិទ្យា</span>
                  <span class="avatar-badge">ខ្មែរ</span>
                </div>
                
                <div class="group-meta">
                  <h1 class="group-title">Large Model Ecosystem (Khmer AI)</h1>
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

              <button class="btn-join">Join</button>
            </div>

            <p class="group-description">
              Artificial Intelligence or Artificial "Stupidity"? Let's discuss the latest developments in large language models.
            </p>
          </div>

          <!-- Card ទី 2: Quick Post Input Card -->
          <div class="card quick-post-card">
            <div class="user-avatar">🌸</div>
            <!-- Trigger Modal on Click -->
            <div class="post-input-placeholder" @click="showPostModal = true">
              Share what's on your mind with fellow members!
            </div>
          </div>

          <!-- Post Feed Container -->
          <div class="card feed-card">
            
            <div class="feed-tabs">
              <button 
                v-for="tab in tabs" 
                :key="tab"
                @click="activeTab = tab"
                :class="['tab-item', { active: activeTab === tab }]"
              >
                {{ tab }}
                <span v-if="activeTab === tab" class="tab-indicator"></span>
              </button>
            </div>


          </div>

          <!-- Post List (driven by mock `posts` data, including postImage) -->
          <div class="card-group-list">

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

                <p class="post-content">{{ post.content }}</p>

                <!-- Post Images (Mock Data: post.postImage), fixed 2x2 grid with +N overlay -->
                <div
                  v-if="post.postImage && post.postImage.length"
                  class="post-image-grid"
                >
                  <div
                    v-for="(img, idx) in post.postImage.slice(0, 4)"
                    :key="idx"
                    class="post-image-item-wrap"
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

                <!-- Footer: group profile thumbnail (left) + Action Buttons (right) -->
                <div class="post-footer">
                  <span class="group-tag" v-if="post.groupTag">
                    <img class="group-thumb" src="https://picsum.photos/seed/khmerai/60/60" alt="group" />
                    {{ post.groupTag }} ›
                  </span>
                  <span v-else></span>

                  <div class="post-actions">
                    <button class="pill-btn pill-blue">
                      <div class="icon-circle">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><polyline points="16 6 12 2 8 6"/><line x1="12" y1="2" x2="12" y2="15"/></svg>
                      </div>
                      <span>{{ post.actions.share }}</span>
                    </button>

                    <button class="pill-btn pill-white-orange">
                      <div class="emoji-circle">😄</div>
                      <span class="text-orange">{{ post.actions.smile }}</span>
                    </button>

                    <button class="pill-btn pill-blue">
                      <div class="icon-circle">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
                      </div>
                      <span>{{ post.actions.save }}</span>
                    </button>

                    <button class="pill-btn pill-blue">
                      <div class="icon-circle">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/></svg>
                      </div>
                      <span>{{ post.actions.comment }}</span>
                    </button>

                    <button class="pill-btn pill-blue pill-circle-only">
                      <div class="icon-circle">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></svg>
                      </div>
                    </button>
                  </div>
                </div>
              </div>

          </div>

        </div>

        <!-- ================= Right Sidebar Column ================= -->
        <div class="sidebar">
          <div class="card widget-card">
            <h2 class="widget-title">Group Announcement</h2>
            <p class="announcement-text">
              The Large Model Ecosystem is an open, inclusive community for AI developers and enthusiasts. Share the latest news, ask technical questions, and collaborate...
            </p>
            <button class="btn-expand">Expand</button>
          </div>

          <div class="card widget-card">
            <h2 class="widget-title">Group Members</h2>

            <div class="member-section">
              <span class="section-label">Administrator</span>
              <div class="admin-info">
                <div class="admin-avatar">👨‍💻</div>
                <span class="admin-name">XCaptain</span>
              </div>
            </div>

            <div class="member-section">
              <span class="section-label">Active Members</span>
              
              <div class="avatar-stack">
                <img 
                  v-for="member in activeMembers" 
                  :key="member.id" 
                  :src="member.avatar" 
                  :alt="member.name"
                  class="stacked-avatar"
                />
              </div>

              <a href="#" class="member-count-link">
                3.6k users have joined ›
              </a>
            </div>
          </div>
        </div>

      </div>
    </div>

    <!-- ================= Create Post Modal Popup ================= -->
    <div v-if="showPostModal" class="modal-overlay" @click.self="showPostModal = false">
      <div class="modal-box card">
        
        <div class="modal-top-action">
          <button class="btn-hashtag-pin">
            <svg class="modal-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/></svg>
            Hashtag pin
          </button>
        </div>

        <div class="modal-input-area">
          <textarea 
            v-model="postContent" 
            placeholder="Whate your mind for posts?" 
            class="modal-textarea"
          ></textarea>
        </div>

        <div class="group-select-wrapper">
          <button class="btn-choose-group">
            <svg class="modal-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
            Choose Group
            <svg class="icon-arrow" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/></svg>
          </button>
        </div>

        <div class="modal-divider"></div>

        <div class="modal-toolbar">
          <div class="toolbar-left">
            <button class="tool-btn">
              <svg class="tool-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
              Emoji
            </button>

            <button class="tool-btn">
              <svg class="tool-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
              Picture
              <svg class="icon-arrow" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/></svg>
            </button>

            <button class="tool-btn">
              <svg class="tool-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"/></svg>
              Link
            </button>

            <button class="tool-btn">
              <span class="hashtag-symbol">#</span>
              Topic
            </button>

            <button class="tool-btn">
              <svg class="tool-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/></svg>
              Code
            </button>
          </div>

          <button class="btn-submit-post">
            Posts
            <svg class="send-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 2L11 13M22 2l-7 20-4-9-9-4 20-7z"/></svg>
          </button>
        </div>

      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import NavBar from '../navbar/NavBar.vue'

const showPostModal = ref(false)
const postContent = ref('')

watch(showPostModal, (isOpen) => {
  if (isOpen) {
    document.body.classList.add('modal-open');
  } else {
    document.body.classList.remove('modal-open');
  }
});

const activeTab = ref('Popular')

const tabs = [
  'Popular', 
  'Latest', 
  'Featured', 
  'Industry News', 
  'Prompts'
]

const groupStats = reactive({
  members: 0,
  boilingPoint: 0
})

const activeMembers = ref([
  { id: 1, name: 'User 1', avatar: 'https://i.pravatar.cc/100?img=12' },
  { id: 2, name: 'User 2', avatar: 'https://i.pravatar.cc/100?img=33' },
  { id: 3, name: 'User 3', avatar: 'https://i.pravatar.cc/100?img=47' },
  { id: 4, name: 'User 4', avatar: 'https://i.pravatar.cc/100?img=68' },
  { id: 5, name: 'User 5', avatar: 'https://i.pravatar.cc/100?img=60' }
])

/* ===== Mock Post Data (includes postImage + real avatar image) ===== */
const posts = ref([
  {
    id: 1,
    author: 'Momo',
    avatar: 'https://i.pravatar.cc/100?img=5',
    role: 'Frontend Engineer @TechCorp',
    time: '2 hours ago',
    content: "Quick question for everyone: which model performs best in your tests—GLM-5.2, Kimi K3, or DeepSeek-V4-Pro?",
    groupTag: 'Large Model Ecosystem',
    postImage: [
      'https://picsum.photos/id/180/600/600',
      'https://picsum.photos/id/181/600/600'
    ],
    actions: { share: 0, smile: 1, save: 0, comment: 0 }
  },
  {
    id: 2,
    author: 'Alex Dev',
    avatar: 'https://i.pravatar.cc/100?img=15',
    role: 'Java Developer',
    time: '20 hours ago',
    content: "Just deployed a RAG pipeline using DeepSeek-V4-Pro, results are impressive so far!",
    groupTag: null,
    postImage: [
      'https://picsum.photos/id/60/600/600'
    ],
    actions: { share: 0, smile: 1, save: 0, comment: 0 }
  },
  {
    id: 3,
    author: 'Sokha',
    avatar: 'https://i.pravatar.cc/100?img=32',
    role: 'ML Researcher',
    time: '1 day ago',
    content: "Comparing inference speed across Kimi K3 vs GLM-5.2 on the same hardware, here are the benchmark shots from today's walk around town.",
    groupTag: 'Large Model Ecosystem',
    postImage: [
      'https://picsum.photos/id/1015/600/600',
      'https://picsum.photos/id/1016/600/600',
      'https://picsum.photos/id/1018/600/600',
      'https://picsum.photos/id/1019/600/600',
      'https://picsum.photos/id/1021/600/600'
    ],
    actions: { share: 2, smile: 5, save: 1, comment: 3 }
  },
  {
    id: 4,
    author: 'Dara',
    avatar: 'https://i.pravatar.cc/100?img=48',
    role: 'AI Product Manager',
    time: '2 days ago',
    content: "No images on this one, just wanted to say thanks to everyone sharing benchmark data this week!",
    groupTag: null,
    postImage: [],
    actions: { share: 1, smile: 3, save: 0, comment: 2 }
  }
])
</script>

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
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.main-content {
  display: flex;
  flex-direction: column;
  gap: 2px; 
}

/* ================= Card ទី 1: Group Banner (កាត់ជ្រុងខាងក្រោមចេញ) ================= */
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
  background: linear-gradient(135deg, #0033A0 0%, #E10600 50%, #0033A0 100%);
  color: #ffffff;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  border: 2px solid #FFD700;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.avatar-flag {
  font-size: 18px;
  line-height: 1;
}

.avatar-title {
  font-size: 11px;
  font-weight: 700;
  margin-top: 2px;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
}

.avatar-badge {
  font-size: 9px;
  font-weight: 800;
  background: #FFD700;
  color: #000000;
  padding: 1px 6px;
  border-radius: 4px;
  margin-top: 2px;
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

/* ================= Card ទី 2: Quick Post Input (កាត់ជ្រុងខាងលើចេញ) ================= */
.quick-post-card {
  padding: 12px 20px;
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 12px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background-color: #F2F3F5;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  flex-shrink: 0;
}

.post-input-placeholder {
  flex: 1;
  background-color: #F7F8FA;
  border: 1px solid transparent;
  border-radius: 32px;
  padding: 10px 18px;
  font-size: 13px;
  color: #8A919F;
  cursor: pointer;
  transition: all 0.2s;
}

.post-input-placeholder:hover {
  background-color: #F2F3F5;
  border-color: #E5E6EB;
}

/* Feed Tabs */
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

/* Post Items — each post is now its own separated card (see .card-group-list gap) */
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
}

/* Post Image Grid (Mock Data): fixed 2 columns, square tiles, +N overlay */
.post-image-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
  max-width: 460px;
}

.post-image-item-wrap {
  position: relative;
  aspect-ratio: 1 / 1;
  border-radius: 14px;
  overflow: hidden;
  border: 1.5px solid #E5E6EB;
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

/* Footer row: group thumbnail (left) + post-actions (right) */
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

/* Post Actions & Pill Buttons */
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
}

.text-orange {
  color: #E65100;
}

/* Sidebar Widgets */
.sidebar {
  display: flex;
  flex-direction: column;
  gap: 16px;
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

/* Modal Popup Styles */
.modal-overlay {
  position: fixed;
  top: 60px;
  left: 0;
  width: 100vw;
  height: calc(100vh - 60px);
  background-color: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
  padding: 16px;
}

.modal-box {
  width: 100%;
  max-width: 850px;
  padding: 16px;
  border-radius: 12px;
  background-color: #ffffff;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.modal-top-action {
  display: flex;
}

.btn-hashtag-pin {
  background-color: #1570EF;
  color: #ffffff;
  border: none;
  border-radius: 20px;
  padding: 7px 14px;
  font-size: 13px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 6px;
  cursor: pointer;
}

.modal-input-area {
  background-color: #F4F4F4;
  border-radius: 8px;
  padding: 16px;
}

.modal-textarea {
  width: 100%;
  height: 160px;
  background: transparent;
  border: none;
  outline: none;
  resize: none;
  font-size: 14px;
  color: #333;
  font-family: inherit;
}

.modal-textarea::placeholder {
  color: #9E9E9E;
  font-weight: 500;
}

.group-select-wrapper {
  display: flex;
}

.btn-choose-group {
  background-color: #1570EF;
  color: #ffffff;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  font-size: 13px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 6px;
  cursor: pointer;
}

.modal-divider {
  height: 1px;
  background-color: #F0E6DD;
  margin: 4px 0;
}

.modal-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  flex-wrap: wrap;
}

.toolbar-left {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.tool-btn {
  background: #ffffff;
  border: 1.5px solid #1570EF;
  color: #1D1D1D;
  border-radius: 20px;
  padding: 5px 12px;
  font-size: 13px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 5px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tool-btn:hover {
  background-color: #F5F8FF;
}

.btn-submit-post {
  background-color: #E2E2E2;
  color: #8E8E8E;
  border: none;
  border-radius: 24px;
  padding: 8px 22px;
  font-size: 14px;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 6px;
  cursor: pointer;
}

.btn-submit-post:hover {
  background-color: #D6D6D6;
}

.modal-icon, .tool-icon {
  width: 15px;
  height: 15px;
}

.icon-arrow {
  width: 12px;
  height: 12px;
}

.send-icon {
  width: 15px;
  height: 15px;
}

.hashtag-symbol {
  font-weight: 700;
  font-size: 13px;
}

:global(body.modal-open) {
  overflow: hidden;
  position: fixed;
  width: 100%;
  height: 100%;
}

.card-group-list{
    display: flex;
    flex-direction: column;
    gap: 16px;
    margin-top: 12px;
}
</style>