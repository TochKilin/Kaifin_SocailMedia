<script setup>
import { computed } from 'vue'

const props = defineProps({
  isLoading: Boolean,
  historyPlaylists: Array,
  historySongs: Array,
  currentSong: Object,
  isPlaying: Boolean,
  formatDuration: Function
})

defineEmits(['open-playlist', 'play-song', 'clear-history'])
const totalHistoryDuration = computed(() => {
  if (!props.historySongs || props.historySongs.length === 0) return '0m'
  const totalSeconds = props.historySongs.reduce((acc, song) => acc + (song.duration || 0), 0)
  const hours = Math.floor(totalSeconds / 3600)
  const minutes = Math.floor((totalSeconds % 3600) / 60)
  if (hours > 0) {
    return `${hours}h ${minutes}m`
  }
  return `${minutes}m`
})

const summaryItems = computed(() => [
  { id: 'songs', icon: '🎵', count: props.historySongs.length, label: 'Songs' },
  { id: 'playlists', icon: '📁', count: props.historyPlaylists.length, label: 'Playlists' },
  { id: 'duration', icon: '⏱', count: totalHistoryDuration.value, label: '' }
])
</script>

<template>
  <div class="my-history-view">
    <!-- Top Action Bar -->
    <div class="history-top-bar">
      <div class="history-left-group">
        <div class="history-main-title">
          <div class="title-icon-wrapper">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <polyline points="12 6 12 12 16 14"></polyline>
            </svg>
          </div>
          <span>Listening History</span>
        </div>
        <div class="history-summary-inline">
          <template v-for="(item, index) in summaryItems" :key="item.id">
            <div class="summary-item">
              <span class="summary-icon">{{ item.icon }}</span>
              <span><strong>{{ item.count }}</strong> {{ item.label }}</span>
            </div>
            <div class="summary-divider" v-if="index < summaryItems.length - 1"></div>
          </template>
        </div>
      </div>

      <button class="clear-history-btn" @click="$emit('clear-history')">
        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="3 6 5 6 21 6"></polyline>
          <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
        </svg>
        Clear history
      </button>
    </div>

    <div class="history-scroll-body">
      <!-- My Playlist Section -->
      <div class="content-section">
        <div class="section-header">
          <h3 class="section-title-text">My Playlist</h3>
        </div>

        <div v-if="isLoading" class="state-message">Loading history playlists...</div>
        <div v-else-if="historyPlaylists.length === 0" class="state-message">No history playlists found.</div>

        <div v-else class="playlist-grid">
          <div 
            v-for="item in historyPlaylists" 
            :key="item.id" 
            class="playlist-card"
            @click="$emit('open-playlist', item)"
          >
            <div class="card-cover-wrapper">
              <img :src="item.cover || item.cover_url" alt="playlist cover" />
              <div class="play-hover-overlay">
                <svg width="22" height="22" viewBox="0 0 24 24" fill="currentColor">
                  <polygon points="5 3 19 12 5 21 5 3"></polygon>
                </svg>
              </div>
            </div>
            <div class="card-footer">
              <span class="playlist-title">{{ item.name }}</span>
              <span class="song-badge">{{ item.songsCount || item.songs_count || 0 }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="content-section flex-fill-section">
        <div class="section-header">
          <h3 class="section-title-text">My Music</h3>
        </div>

        <div v-if="isLoading" class="state-message">Loading history songs...</div>
        <div v-else-if="historySongs.length === 0" class="state-message">No history songs found.</div>

        <div v-else class="song-stack">
          <div
            v-for="song in historySongs"
            :key="song.id"
            class="song-row-item"
            :class="{ 'is-active-playing': currentSong?.id === song.id && isPlaying }"
            @click="$emit('play-song', song)"
          >
            <div class="song-thumbnail">
              <img :src="song.cover" alt="cover" />
              <div class="active-eq-bars" v-if="currentSong?.id === song.id && isPlaying">
                <span></span><span></span><span></span>
              </div>
            </div>
            
            <div class="song-meta">
              <span class="song-name">{{ song.title }}</span>
              <span class="song-artist">{{ song.singer }}</span>
            </div>

            <div class="song-length">
              <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <polyline points="12 6 12 12 16 14"></polyline>
              </svg>
              {{ formatDuration(song.duration) }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.my-history-view {
  display: flex;
  flex-direction: column;
  height: 100vh;
  gap: 8px;
  padding: 0px 4px;
  width: 100%;
  box-sizing: border-box;
}

/* Top Bar */
.history-top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #ffffff;
  padding: 14px 20px;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  flex-wrap: wrap;
  gap: 10px;
  flex-shrink: 0;
}

.history-left-group {
  display: flex;
  align-items: center;
  gap: 20px;
  flex-wrap: wrap;
}

.history-main-title {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
}

.title-icon-wrapper {
  color: #007bff;
  display: flex;
  align-items: center;
}

.history-summary-inline {
  display: flex;
  align-items: center;
  gap: 14px;
  font-size: 13px;
  color: #495057;
  padding: 6px 12px;
  border-radius: 8px;
}

.summary-item {
  display: flex;
  align-items: center;
  gap: 5px;
}

.summary-icon {
  font-size: 14px;
}

.summary-divider {
  width: 1px;
  height: 12px;
  background: #dee2e6;
}

.clear-history-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  background-color: transparent;
  border-color: transparent;
  color: #e03131;
  padding: 6px 14px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.clear-history-btn:hover {
  background: #ffe3e3;
}

.history-scroll-body {
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex: 1;
  overflow-y: auto;
  min-height: 0;
  padding-bottom: 12px;
  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.content-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
  background: #ffffff;
  padding: 14px 20px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

.flex-fill-section {
  flex: 1;
  min-height: 200px;
}

.section-header .section-title-text {
  font-size: 17px;
  font-weight: 600;
  color: #333333;
  margin: 0;
}


.playlist-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 16px;
}

.playlist-card {
  background: #ffffff;
  border-radius: 12px;
  padding: 10px;
  cursor: pointer;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.03);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  border: 1px solid #f1f3f5;
}

.playlist-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.08);
}

.card-cover-wrapper {
  position: relative;
  width: 100%;
  aspect-ratio: 1;
  border-radius: 8px;
  overflow: hidden;
  background: #f8f9fa;
  margin-bottom: 10px;
}

.card-cover-wrapper img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.play-hover-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.35);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.playlist-card:hover .play-hover-overlay {
  opacity: 1;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 6px;
}

.playlist-title {
  font-size: 13px;
  font-weight: 600;
  color: #333333;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  flex: 1;
}

.song-badge {
  font-size: 11px;
  font-weight: 600;
  color: #007bff;
  background: #eef6ff;
  padding: 2px 6px;
  border-radius: 6px;
}

/* Song Stack */
.song-stack {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.song-row-item {
  display: flex;
  align-items: center;
  gap: 12px;
  background: #ffffff;
  padding: 10px 14px;
  border-radius: 12px;
  cursor: pointer;
  border: 1px solid #f1f3f5;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.02);
  transition: background 0.2s ease, transform 0.1s ease;
}

.song-row-item:hover {
  background: #f8f9fa;
}

.song-row-item.is-active-playing {
  background: #eef6ff;
  border-color: #b6d9ff;
}

.song-thumbnail {
  position: relative;
  width: 44px;
  height: 44px;
  border-radius: 8px;
  overflow: hidden;
  background: #f1f3f5;
  flex-shrink: 0;
}

.song-thumbnail img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.active-eq-bars {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: flex-end;
  justify-content: center;
  gap: 2px;
  padding-bottom: 8px;
}

.active-eq-bars span {
  width: 3px;
  background: #ffffff;
  border-radius: 1px;
  animation: eq-bounce 0.8s ease-in-out infinite;
}

.active-eq-bars span:nth-child(1) { animation-delay: 0s; }
.active-eq-bars span:nth-child(2) { animation-delay: 0.2s; }
.active-eq-bars span:nth-child(3) { animation-delay: 0.4s; }

@keyframes eq-bounce {
  0%, 100% { height: 4px; }
  50% { height: 14px; }
}

.song-meta {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
  min-width: 0;
}

.song-name {
  font-size: 13px;
  font-weight: 600;
  color: #2c3e50;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.song-artist {
  font-size: 12px;
  color: #6c757d;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.song-length {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  font-weight: 500;
  color: #555;
  background: #f1f3f5;
  padding: 3px 6px;
  border-radius: 6px;
  border: 1px solid #e2e8f0;
  flex-shrink: 0;
}

.state-message {
  font-size: 13px;
  color: #6c757d;
  padding: 8px 0;
}
</style>