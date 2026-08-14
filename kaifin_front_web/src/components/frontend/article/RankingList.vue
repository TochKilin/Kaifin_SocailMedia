<template>
  <div class="ranking-app">
    <!-- Sidebar Navigation -->
    <aside class="sidebar">
      <div class="nav-menu">
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
          Following
        </button>
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z"/></svg>
          New Article
        </button>
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/></svg>
          Back End
        </button>
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/></svg>
          Development Tools
        </button>
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M16 18l6-6-6-6M8 6l-6 6 6 6"/></svg>
          Code Life
        </button>
        <button class="nav-item">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/></svg>
          Read
        </button>
        <button class="nav-item active">
          <svg class="icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"/></svg>
          Ranking List
        </button>
      </div>
    </aside>

    <!-- Main Content Area -->
    <main class="main-content">
      <!-- Top Header Title -->
      <div class="content-header">
        <h1>Top of Article</h1>
      </div>

      <!-- Ranking Articles List -->
      <div class="articles-list">
        <div class="article-card" v-for="item in rankingArticles" :key="item.id">
          
          <!-- Left: Rank Number & Details -->
          <div class="article-left">
            <div class="rank-badge" :class="'rank-' + item.id">
              {{ item.id }}
            </div>
            
            <div class="article-info">
              <h3 class="article-title">{{ item.title }}</h3>
              <div class="article-meta">
                <div class="author-box">
                  <svg class="meta-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                  <span class="username">{{ item.username }}</span>
                </div>
                <span class="meta-tag">
                  <svg class="meta-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/></svg>
                  {{ item.views }}
                </span>
                <span class="meta-tag">
                  <svg class="meta-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"/></svg>
                  {{ item.comments }}
                </span>
                <span class="meta-tag">
                  <svg class="meta-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/></svg>
                  {{ item.saves }}
                </span>
              </div>
            </div>
          </div>

          <!-- Right: Like Button & Create Button -->
          <div class="article-right">
            <button class="btn-like" @click="item.likes++">
              <svg class="btn-icon" fill="currentColor" viewBox="0 0 24 24"><path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/></svg>
              <span class="like-count">{{ item.likes }}k</span> Like
            </button>
            <button class="btn-create-article">
              Create Article
            </button>
          </div>

        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const rankingArticles = ref([
  { id: 1, title: 'The Article Recomment1', username: 'Username', views: '110k', comments: '10k', saves: '7k', likes: 9 },
  { id: 2, title: 'The Article Recomment', username: 'Username', views: '100k', comments: '9k', saves: '6k', likes: 8 },
  { id: 3, title: 'The Article Recomment', username: 'Username', views: '90k', comments: '8k', saves: '5k', likes: 7 },
  { id: 4, title: 'The Article Recomment', username: 'Username', views: '80k', comments: '7k', saves: '4k', likes: 6 },
  { id: 5, title: 'The Article Recomment', username: 'Username', views: '70k', comments: '6k', saves: '3k', likes: 5 },
  { id: 6, title: 'The Article Recomment', username: 'Username', views: '70k', comments: '5k', saves: '2k', likes: 4 }
])
</script>

<style scoped>
/* Main Layout */
.ranking-app {
  display: grid;
  grid-template-columns: 260px 1fr;
  gap: 24px;
  max-width: 1280px;
  margin: 0 auto;
  padding: 24px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  background-color: #f8fafc;
  min-height: 100vh;
  color: #1e293b;
  box-sizing: border-box;
}

/* Sidebar Styles */
.sidebar {
  background: #1e293b;
  border-radius: 16px;
  padding: 20px 12px;
  height: fit-content;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
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
  border: none;
  width: 100%;
  padding: 12px 16px;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 600;
  color: #94a3b8;
  cursor: pointer;
  text-align: left;
  transition: all 0.2s ease;
}

.nav-item .icon {
  width: 18px;
  height: 18px;
  flex-shrink: 0;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.05);
  color: #f8fafc;
}

.nav-item.active {
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  color: #ffffff;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
}

/* Main Content Styles */
.main-content {
  display: flex;
  flex-direction: column;
  gap: 20px;

}

.content-header h1 {
  margin: 0;
  font-size: 22px;
  font-weight: 800;
  color: #0f172a;
}

/* Articles List */
.articles-list {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.article-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 14px;
  padding: 16px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
  transition: transform 0.2s, box-shadow 0.2s;
}

.article-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.05);
  border-color: #cbd5e1;
}

.article-left {
  display: flex;
  align-items: center;
  gap: 16px;
  flex: 1;
}

/* Rank Badge styling */
.rank-badge {
  width: 38px;
  height: 38px;
  background: #f1f5f9;
  color: #475569;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 16px;
  flex-shrink: 0;
}

/* Highlight top 3 ranks */
.rank-1 { background: #fef3c7; color: #d97706; }
.rank-2 { background: #f1f5f9; color: #64748b; }
.rank-3 { background: #ffedd5; color: #c2410c; }

.article-info {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.article-title {
  margin: 0;
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}

.article-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 12px;
  color: #64748b;
  flex-wrap: wrap;
}

.author-box {
  display: flex;
  align-items: center;
  gap: 4px;
  background: #f8fafc;
  padding: 2px 8px;
  border-radius: 6px;
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
  background: #f1f5f9;
  padding: 2px 8px;
  border-radius: 6px;
  font-weight: 500;
}

.meta-icon {
  width: 14px;
  height: 14px;
  color: #64748b;
}

/* Right Actions */
.article-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.btn-like {
  display: flex;
  align-items: center;
  gap: 6px;
  background: #fdf2f8;
  border: 1px solid #fbcfe8;
  color: #db2777;
  padding: 8px 14px;
  border-radius: 10px;
  font-weight: 600;
  font-size: 13px;
  cursor: pointer;
  transition: background 0.2s;
}

.btn-icon {
  width: 16px;
  height: 16px;
}

.btn-like:hover {
  background: #fce7f3;
}

.btn-create-article {
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  color: #ffffff;
  border: none;
  padding: 9px 16px;
  border-radius: 10px;
  font-weight: 600;
  font-size: 13px;
  cursor: pointer;
  box-shadow: 0 2px 6px rgba(99, 102, 241, 0.2);
  transition: transform 0.2s, box-shadow 0.2s;
}

.btn-create-article:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
}
</style>