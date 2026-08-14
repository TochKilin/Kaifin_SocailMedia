<template>
  <div class="creative-app">
    <!-- Sidebar Navigation -->
    <aside class="sidebar">
      <div class="logo-area">
        <span class="logo-dot"></span>
        <h2>DevSpace</h2>
      </div>
      <nav class="nav-menu">
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
          Following
        </button>
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z"/></svg>
          General Post
        </button>
        <button class="nav-item active">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/></svg>
          Code Life
        </button>
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/></svg>
          Read
        </button>
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"/></svg>
          Ranking List
        </button>
      </nav>
    </aside>

    <!-- Main Center Content -->
    <main class="main-feed">
      <!-- Header with Search and Create -->
      <header class="feed-header">
        <div class="search-box">
          <svg class="search-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
          <input type="text" placeholder="Search ...." />
        </div>
        <button class="btn-create">Create Article</button>
      </header>

      <!-- Tabs Filter -->
      <div class="feed-tabs">
        <button class="tab active">Popular</button>
        <button class="tab">The Latest</button>
      </div>

      <!-- Article Cards Feed -->
      <div class="articles-container">
        <article class="article-card" v-for="article in articles" :key="article.id">
          <div class="article-content">
            <h3 class="art-title">{{ article.title }}</h3>
            <p class="art-desc">{{ article.desc1 }}</p>
            <p class="art-desc">{{ article.desc2 }}</p>
            
            <div class="art-footer">
              <div class="author-pill">
                <svg class="sm-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                <span>{{ article.username }}</span>
              </div>
              <div class="stat-pill">
                <svg class="sm-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/></svg>
                <span>{{ article.views }}</span>
              </div>
              <button class="like-pill" @click="article.likes++">
                <svg class="sm-icon heart" fill="currentColor" viewBox="0 0 24 24"><path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/></svg>
                <span>Like</span>
              </button>
              <div class="dots-pill">•••</div>
            </div>
          </div>
          <div class="article-img-wrapper">
            <img :src="article.image" alt="Article thumbnail" />
          </div>
        </article>
      </div>
    </main>

    <!-- Right Sidebar Widgets -->
    <aside class="widgets-sidebar">
      <!-- List Article Box -->
      <div class="widget-box">
        <div class="widget-header">
          <h3>List Article</h3>
          <svg class="sm-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
        </div>
        <ul class="rank-list">
          <li v-for="i in 5" :key="i">
            <span class="num">{{ i }}</span>
            <span class="title-text">title article</span>
          </li>
        </ul>
      </div>

      <!-- List of Authors Box -->
      <div class="widget-box">
        <div class="widget-header">
          <h3>List of authors</h3>
        </div>
        <div class="author-list">
          <div class="author-row" v-for="i in 3" :key="i">
            <div class="author-info-sm">
              <div class="mini-avatar">👤</div>
              <span>Name author</span>
            </div>
            <button class="btn-follow">Follow</button>
          </div>
        </div>
      </div>

      <!-- Top Topic Box -->
      <div class="widget-box">
        <div class="widget-header">
          <h3>Top Topic</h3>
          <svg class="sm-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
        </div>
        <div class="topic-list">
          <div class="topic-row" v-for="i in 4" :key="i">
            <span class="hash">#</span>
            <span class="topic-name">Name topic</span>
            <svg class="sm-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
          </div>
        </div>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const articles = ref([
  { id: 1, title: 'Article name', desc1: 'Text', desc2: 'Text', username: 'Username', views: '100View', likes: 12, image: 'https://images.unsplash.com/photo-1542751371-adc38448a05e?auto=format&fit=crop&w=300&q=80' },
  { id: 2, title: 'Article name', desc1: 'Text', desc2: 'Text', username: 'Username', views: '100View', likes: 8, image: 'https://images.unsplash.com/photo-1517694712202-14dd9538aa97?auto=format&fit=crop&w=300&q=80' },
  { id: 3, title: 'Article name', desc1: 'Text', desc2: 'Text', username: 'Username', views: '100View', likes: 15, image: 'https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5?auto=format&fit=crop&w=300&q=80' },
  { id: 4, title: 'Article name', desc1: 'Text', desc2: 'Text', username: 'Username', views: '100View', likes: 5, image: 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?auto=format&fit=crop&w=300&q=80' }
])
</script>

<style scoped>
/* Global Layout Inspired by Wireframe */
.creative-app {
  display: grid;
  grid-template-columns: 240px 1fr 280px;
  gap: 20px;
  max-width: 1400px;
  margin: 0 auto;
  padding: 20px;
  background-color: #1a1110; /* Dark brownish-black creative tone */
  min-height: 100vh;
  font-family: 'Inter', system-ui, sans-serif;
  color: #f1f5f9;
  box-sizing: border-box;
}

/* Icons styling */
.icon, .sm-icon {
  width: 18px;
  height: 18px;
  flex-shrink: 0;
}
.sm-icon { width: 14px; height: 14px; opacity: 0.8; }

/* Left Sidebar */
.sidebar {
  background: #261917;
  border: 1px solid #3d2623;
  border-radius: 16px;
  padding: 20px 12px;
  display: flex;
  flex-direction: column;
  gap: 24px;
  height: fit-content;
  position: sticky;
  top: 20px;
}

.logo-area {
  display: flex;
  align-items: center;
  gap: 10px;
  padding-left: 12px;
}

.logo-dot {
  width: 10px;
  height: 10px;
  background: #f43f5e;
  border-radius: 50%;
  box-shadow: 0 0 10px #f43f5e;
}

.logo-area h2 {
  font-size: 18px;
  font-weight: 800;
  letter-spacing: 0.5px;
  margin: 0;
}

.nav-menu {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  background: transparent;
  border: 1px solid transparent;
  width: 100%;
  padding: 12px 16px;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  color: #a89f9e;
  cursor: pointer;
  text-align: left;
  transition: all 0.25s ease;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.03);
  color: #ffffff;
  border-color: #3d2623;
}

.nav-item.active {
  background: #35211e;
  color: #ffffff;
  border-color: #5c3530;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

/* Main Feed */
.main-feed {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.feed-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 16px;
}

.search-box {
  display: flex;
  align-items: center;
  gap: 10px;
  background: #261917;
  border: 1px solid #3d2623;
  padding: 10px 16px;
  border-radius: 14px;
  flex: 1;
  color: #a89f9e;
}

.search-box input {
  background: transparent;
  border: none;
  outline: none;
  color: #fff;
  font-size: 14px;
  width: 100%;
}

.btn-create {
  background: #35211e;
  border: 1px solid #5c3530;
  color: #fff;
  padding: 12px 20px;
  border-radius: 14px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}

.btn-create:hover {
  background: #462c28;
  border-color: #784741;
}

.feed-tabs {
  display: flex;
  gap: 12px;
}

.tab {
  background: #261917;
  border: 1px solid #3d2623;
  color: #a89f9e;
  padding: 8px 18px;
  border-radius: 20px;
  font-weight: 600;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.2s;
}

.tab.active, .tab:hover {
  background: #35211e;
  color: #fff;
  border-color: #5c3530;
}

/* Articles Feed Container */
.articles-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.article-card {
  background: #261917;
  border: 1px solid #3d2623;
  border-radius: 16px;
  padding: 18px;
  display: flex;
  justify-content: space-between;
  gap: 16px;
  transition: transform 0.2s, border-color 0.2s;
}

.article-card:hover {
  transform: translateY(-2px);
  border-color: #5c3530;
}

.article-content {
  display: flex;
  flex-direction: column;
  gap: 8px;
  flex: 1;
}

.art-title {
  margin: 0;
  font-size: 16px;
  font-weight: 700;
  color: #fff;
}

.art-desc {
  margin: 0;
  font-size: 13px;
  color: #a89f9e;
  border-bottom: 1px dashed #3d2623;
  padding-bottom: 4px;
  width: 80%;
}

.art-footer {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 8px;
  flex-wrap: wrap;
}

.author-pill, .stat-pill, .like-pill {
  display: flex;
  align-items: center;
  gap: 6px;
  background: #1a1110;
  border: 1px solid #3d2623;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 12px;
  color: #c5bcbb;
}

.like-pill {
  cursor: pointer;
  color: #f43f5e;
  border-color: #51252a;
  background: #2a1518;
}

.like-pill:hover {
  background: #381a1f;
}

.heart { color: #f43f5e; }

.dots-pill {
  color: #a89f9e;
  font-size: 12px;
  padding: 0 4px;
  letter-spacing: 2px;
}

.article-img-wrapper {
  width: 100px;
  height: 90px;
  border-radius: 12px;
  overflow: hidden;
  border: 1px solid #3d2623;
  flex-shrink: 0;
}

.article-img-wrapper img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Right Widgets Sidebar */
.widgets-sidebar {
  display: flex;
  flex-direction: column;
  gap: 16px;
  position: sticky;
  top: 20px;
  height: fit-content;
}

.widget-box {
  background: #111c26; /* Deep creative blue tone matching right side of wireframe */
  border: 1px solid #1e3142;
  border-radius: 16px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.widget-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #1e3142;
  padding-bottom: 8px;
}

.widget-header h3 {
  margin: 0;
  font-size: 14px;
  font-weight: 700;
  color: #e2e8f0;
}

/* Rank List */
.rank-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.rank-list li {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 13px;
  color: #94a3b8;
}

.rank-list .num {
  background: #1a2938;
  border: 1px solid #284158;
  width: 22px;
  height: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 6px;
  font-weight: 700;
  font-size: 11px;
  color: #cbd5e1;
}

.title-text {
  border-bottom: 1px dashed #1e3142;
  flex: 1;
  padding-bottom: 2px;
}

/* Author List */
.author-list, .topic-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.author-row, .topic-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.author-info-sm {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #cbd5e1;
}

.mini-avatar {
  background: #1a2938;
  border: 1px solid #284158;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 6px;
  font-size: 11px;
}

.btn-follow {
  background: #1a2938;
  border: 1px solid #284158;
  color: #cbd5e1;
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 12Fontpx;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}

.btn-follow:hover {
  background: #284158;
  color: #fff;
}

/* Top Topic */
.topic-row {
  background: #162432;
  border: 1px solid #1e3142;
  padding: 6px 10px;
  border-radius: 8px;
  font-size: 13px;
  color: #cbd5e1;
}

.hash {
  font-weight: 700;
  color: #38bdf8;
}

.topic-name {
  flex: 1;
  margin-left: 6px;
}
</style>