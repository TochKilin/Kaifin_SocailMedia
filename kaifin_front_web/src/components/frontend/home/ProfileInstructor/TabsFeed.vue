<!-- TabsFeed.vue -->
<script setup>
defineProps({
  profileData: {
    type: Object,
    required: true
  }
});
</script>

<template>
  <div class="feed-item-card">
    <div class="feed-user-header">
      <img :src="profileData.profileImage" alt="Avatar" class="small-avatar-img" />
      <div class="feed-user-meta">
        <h4>{{ profileData.username }}</h4>
        <span class="feed-time">
          <svg class="time-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/></svg>
          01-01-2026 10:30 AM
        </span>
      </div>
    </div>
    <p class="feed-description">Published a new comprehensive course on Full-Stack Development and modern architecture patterns.</p>
    
    <div class="feed-interaction-bar">
      <div class="stack-links">
        <span class="interaction-item">
          <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M2 12h20M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"/></svg>
          Translate
        </span>
        <span class="interaction-item">
          <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>
          Course Group
        </span>
      </div>
      
      <div class="liked-by-container">
        <div class="liked-avatars-group">
          <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThKDPgkgQbHVq4mAfGv87ldBK3an3JE4Gn1_yjCW4A_9I9OVX6WCfpq8lw&s=10" alt="User 1" class="liked-avatar" />
          <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTj6npa_ycZdip1l2x4Cq5cvXTk4zLiSF2a1i6aezuqfA&s=10" alt="User 2" class="liked-avatar" />
          <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTGnZCuo8SOgPSV2CNKFJRdnVPA7DRCiNduThp8-G7mPw&s=10" alt="User 3" class="liked-avatar" />
        </div>
        <span class="liked-count-text">Liked by 150</span>
      </div>
    </div>

    <div class="feed-footer-buttons">
      <button class="footer-btn">
        <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><polyline points="16 6 12 2 8 6"/><line x1="12" y1="2" x2="12" y2="15"/></svg>
        <span>1.8K</span>
      </button>
      <button class="footer-btn">
        <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
        <span>989K</span>
      </button>
      <button class="footer-btn">
        <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"/></svg>
        <span>21K</span>
      </button>
      <button class="footer-btn">
        <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
        <span>55K</span>
      </button>
    </div>
  </div>
</template>

<style scoped>
.feed-item-card {
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 14px;
  background: #ffffff;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.feed-user-header {
  display: flex;
  align-items: center;
  gap: 10px;
}

.small-avatar-img {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: 1px solid #cbd5e1;
  object-fit: cover;
}

.feed-user-meta h4 {
  margin: 0;
  font-size: 14px;
  color: #0f172a;
}

.feed-time {
  font-size: 11px;
  color: #94a3b8;
  display: flex;
  align-items: center;
  gap: 4px;
}

.time-svg-icon {
  width: 12px;
  height: 12px;
  stroke: currentColor;
}

.feed-description {
  font-size: 13px;
  color: #334155;
  margin: 0;
  border-top: 1px solid #f1f5f9;
  padding: 8px 0;
}

.feed-interaction-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
  color: #64748b;
}

.stack-links {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.interaction-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

.liked-by-container {
  display: flex;
  align-items: center;
  gap: 8px;
}

.liked-avatars-group {
  display: flex;
  align-items: center;
}

.liked-avatar {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  border: 2px solid #ffffff;
  object-fit: cover;
  background: #e2e8f0;
  margin-left: -8px;
}

.liked-avatars-group .liked-avatar:first-child {
  margin-left: 0;
}

.liked-count-text {
  font-size: 12px;
  font-weight: 600;
  color: #475569;
}

.svg-icon {
  width: 18px;
  height: 18px;
  stroke: currentColor;
  flex-shrink: 0;
}

.feed-footer-buttons {
  display: flex;
  justify-content: flex-end;
  border-top: 1px solid #f1f5f9;
  padding-top: 10px;
  gap: 8px;
}

.footer-btn {
  background: #1B75D2;
  border: none;
  color: #ffffff;
  font-size: 14px;
  cursor: pointer;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 10px 20px;
  border-radius: 32px;
  transition: opacity 0.2s;
}

.footer-btn:hover {
  opacity: 0.9;
  color: #ffffff;
}
</style>