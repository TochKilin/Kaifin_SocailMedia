<script setup>
import { ref, onMounted } from 'vue'
import CreatePlaylist from './CreatePlaylist.vue'

const props = defineProps({
  apiBase: String,
  currentSong: Object,
  isPlaying: Boolean
})

const emit = defineEmits(['play-song', 'open-playlist', 'upload', 'upload-vk', 'upload-kaifin', 'create-collection'])

// បន្ថែម state សម្រាប់គ្រប់គ្រងការប្តូរផ្ទាំងរវាងទំព័រដើម និងផ្ទាំងបង្កើត Playlist
const activeView = ref('main')

const myPlaylists = ref([])
const mySongs = ref([])
const isLoading = ref(false)
const errorMessage = ref(null)

const formatDuration = (secs) => {
  if (!secs || isNaN(secs)) return '0:00'
  const minutes = Math.floor(secs / 60)
  const seconds = Math.floor(secs % 60).toString().padStart(2, '0')
  return `${minutes}:${seconds}`
}

const fetchMyMusic = async () => {
  isLoading.value = true
  errorMessage.value = null
  
  setTimeout(() => {
    myPlaylists.value = [
      {
        id: 10,
        name: 'Sin Sisamuth',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS__s2e8wBwpKY1zyKLaXHldZkVvYsomIgPWssNHgCQoQ&s=10'
      },
      {
        id: 11,
        name: 'របាំប្រពៃណីខ្មែរ',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQNUk9Sc0tWPflJVba5ZqRmhYse4P64x61SLbBJ6RTzbg&s=10'
      },
      {
        id: 10,
        name: 'លោកវាន់ដា',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSZRxRTks6AxASJCv-BDF1YAQhTi-dJizQ1aihW8EvEHQ&s=10'
      },
      {
        id: 10,
        name: 'មាស សុខសោភា',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSF33X7QzPlMIC0mwYmnPhyhYR8Yuoxm2bS2o69ivcJ1g&s=10'
      },
      {
        id: 10,
        name: 'កញ្ញាយ៉ូរី',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSQXEcOh2xYecNQDuFRqPilH59gTj7XEsk6HVDepjCKhQ&s=10'
      },
      {
        id: 10,
        name: 'Trending Mix',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQJL1vIZcvyKrFAAGe0v4Ff91FO1Uj6ira6o9vYAUh-iw&s=10'
      },
      {
        id: 10,
        name: 'Trending Mix',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT3F2zxbkwH7ydDo8wYTkJlaLLx4tY18tfkzO9735Ap3Q&s=10'
      },
      {
        id: 10,
        name: 'Reth Xozana',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQJL1vIZcvyKrFAAGe0v4Ff91FO1Uj6ira6o9vYAUh-iw&s=10'
      },
      {
        id: 10,
        name: 'LoveLy',
        songs_count: 5,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ8-Z0l8f_yc352Brewlpzm93LraD7G-8xMawChCmi8Gw&s=10'
      },
    ]

    mySongs.value = [
      {
        id: 1,
        title: 'តន្ត្រីខ្មែរសម័យពីដើម',
        singer_name: 'ក្រុមតន្ត្រីសម័យ',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR7Cqs3wiTyptiDEId8fGtAcgEePBJTiFaEdlzcznsT5g&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3'
      },
      {
        id: 2,
        title: 'សុីន សុីសាមុត',
        singer_name: 'ក្រុមតន្ត្រីសម័យ',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSnqs2aRy19bTPq-g7Doo-7nutv62V3VwloydFPGN0s9Q&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3'
      },
      {
        id: 3,
        title: 'Sin Sisamuth old songs',
        singer_name: 'Singer',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS__s2e8wBwpKY1zyKLaXHldZkVvYsomIgPWssNHgCQoQ&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-3.mp3'
      },
      {
        id: 4,
        title: 'ភ្លេងរបាំ សម័យ 1970',
        singer_name: 'Khmer Team music',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQNUk9Sc0tWPflJVba5ZqRmhYse4P64x61SLbBJ6RTzbg&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-4.mp3'
      },
      {
        id: 5,
        title: 'Van Dara Song 2023',
        singer_name: 'Vandar',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTW8j4yHzc9z2l0fYAsAAMgFtk4uhCID93sAPNwuE4zkw&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-5.mp3'
      },
      {
        id: 6,
        title: 'បងស្រឡាញសក់ក្រង',
        singer_name: 'សុីន សុីសាមុត',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSnqs2aRy19bTPq-g7Doo-7nutv62V3VwloydFPGN0s9Q&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-6.mp3'
      },
      {
        id: 7,
        title: 'អង្វរម៉ែរ',
        singer_name: 'Puth Theareak',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS0rYvaxKa2r8pj6N5w6AKap3mUF50DsFINl91l2NPJdg&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-7.mp3'
      },
      {
        id: 8,
        title: 'Khmer vol 12',
        singer_name: 'Sok Vana',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQNUk9Sc0tWPflJVba5ZqRmhYse4P64x61SLbBJ6RTzbg&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-8.mp3'
      },
      {
        id: 9,
        title: 'ចិត្តមួយថ្លើមមួយ',
        singer_name: 'Singer',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQLTB-fwTB0YPmNTWdd7L7yFEMjV8mKXX3whLykbRLEng&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-9.mp3'
      },
      {
        id: 10,
        title: 'music group based in Phnom Penh',
        singer_name: 'Singer',
        duration: 220,
        cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRwWQIAC94HR5OArLfiDJ2lnqGRlfpOtClv-bT5_cG1lQ&s=10',
        fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-10.mp3'
      }
    ]

    isLoading.value = false
  }, 300)
}

onMounted(() => {
  fetchMyMusic()
})
</script>

<template>
  <div class="my-music-view">
    <!-- Top Header containing Cover Image, Title, Genre Pills, and Action Buttons -->
    <div class="my-music-wrapper">
      <!-- ផ្ទាំងទំព័រដើម (បង្ហាញនៅពេល activeView គឺ 'main') -->
      <div v-if="activeView === 'main'">
        <div class="box header-box">
          <div class="header-top-row">
            <div class="header-left">
              <div class="header-image-wrap">
                <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSlvYTE0TSdoApMO06QKeJkC31rFkjiVCVBi_FDOnNVvg&s=10" alt="Cover Image" />
              </div>
              <div class="header-titles">
                <h2 class="section-main-title">Nong Sophea</h2>
                <!-- Genre Pills -->
                <div class="genre-pills">
                  <span class="genre-pill">ចម្រៀងខ្មែរ</span>
                  <span class="genre-pill">Sin Sisamuth</span>
                  <span class="genre-pill">Piano</span>
                </div>
              </div>
            </div>

            <!-- Action Buttons Row -->
            <div class="header-action-buttons">
              <button class="action-btn" @click="$emit('upload')">
                <span class="icon-bg">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                    <polyline points="17 8 12 3 7 8"></polyline>
                    <line x1="12" y1="3" x2="12" y2="15"></line>
                  </svg>
                </span>
                Upload
              </button>

              <!-- ប៊ូតុងសម្រាប់ចុចបើកផ្ទាំង Create Playlist -->
              <button class="action-btn" @click="activeView = 'create-playlist'">
                <span class="icon-bg">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                  </svg>
                </span>
                Create a PlayLists
              </button>
            </div>
          </div>
        </div>

        <!-- My Playlist Section -->
        <section class="box">
          <div class="box-header">
            <span class="box-title">My Playlist</span>
          </div>
          
          <div v-if="isLoading" class="loading-text">Loading playlists...</div>
          <div v-else-if="myPlaylists.length === 0" class="empty-text">No playlists found.</div>
          
          <div v-else class="grid">
            <div 
              v-for="item in myPlaylists" 
              :key="item.id" 
              class="card"
              @click="$emit('open-playlist', item)"
            >
              <div class="card-img-wrap">
                <img :src="item.cover_url" alt="playlist" />
                <!-- Floating "5 songs" badge inside top-right -->
                <span class="song-badge">5 songs</span>
                
                <!-- Larger Play button always visible at the bottom-right corner of the thumbnail -->
                <button 
                  class="card-play-btn" 
                  @click.stop="$emit('play-song', item)"
                  title="Play Playlist"
                >
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
                    <polygon points="5 3 19 12 5 21 5 3"></polygon>
                  </svg>
                </button>
              </div>
              <div class="card-desc">
                <span class="card-title">{{ item.name }}</span>
                <span class="card-subtitle">Playlist</span>
              </div>
            </div>
          </div>
        </section>

        <!-- My Music Songs Section (10 Songs) -->
        <section class="box songs-box">
          <div class="box-header">
            <span class="box-title">10 Songs</span>
          </div>

          <div v-if="isLoading" class="loading-text">Loading songs...</div>
          <div v-else-if="mySongs.length === 0" class="empty-text">No songs found.</div>

          <div v-else class="song-list">
            <div
              v-for="song in mySongs"
              :key="song.id"
              class="song-item"
              :class="{ 'is-playing': currentSong?.id === song.id && isPlaying }"
              @click="$emit('play-song', song)"
            >
              <div class="thumb">
                <img :src="song.cover_url" alt="cover" />
                <!-- Overlay Play Button that appears on hover -->
                <div class="thumb-overlay">
                  <button class="thumb-play-btn" @click.stop="$emit('play-song', song)">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor">
                      <polygon points="5 3 19 12 5 21 5 3"></polygon>
                    </svg>
                  </button>
                </div>
              </div>
              <div class="info">
                <div class="name-box">{{ song.title }}</div>
                <div class="singer-box">{{ song.singer_name }}</div>
              </div>
              <div class="song-btns">
                <span class="duration-badge">
                  <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <circle cx="12" cy="12" r="10"></circle>
                    <polyline points="12 6 12 12 16 14"></polyline>
                  </svg>
                  {{ formatDuration(song.duration) }}
                </span>
              </div>
            </div>
          </div>
        </section>
      </div>

      <!-- ផ្ទាំងបង្កើត Playlist (បង្ហាញនៅពេល activeView គឺ 'create-playlist') -->
      <CreatePlaylist 
        v-else-if="activeView === 'create-playlist'" 
        @back="activeView = 'main'" 
      />
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap');

.my-music-view {
  font-family: 'Poppins', sans-serif;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  height: 100%;
  min-height: 100%;
}

.header-box {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding: 18px 20px !important;
  margin-bottom: 0 !important;
}

.header-top-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 12px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 14px;
}

.header-image-wrap {
  width: 58px;
  height: 58px;
  border-radius: 12px;
  overflow: hidden;
  background: #eee;
  flex-shrink: 0;
  box-shadow: 0 2px 6px rgba(0,0,0,0.06);
}

.header-image-wrap img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.header-titles {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.section-main-title {
  font-size: 20px;
  font-weight: 700;
  color: #222;
  margin: 0;
  line-height: 1.2;
}

.genre-pills {
  display: flex;
  gap: 6px;
  margin-top: 2px;
  flex-wrap: wrap;
}

.genre-pill {
  font-size: 11px;
  font-weight: 500;
  background-color: #f1f3f5;
  color: #555;
  padding: 2px 8px;
  border-radius: 10px;
}

/* Header Action Buttons Container */
.header-action-buttons {
  display: flex;
  gap: 6px;
  align-items: center;
  flex-wrap: wrap;
}

.action-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background-color: #f1f3f5;
  color: #333;
  border: none;
  padding: 5px 10px 5px 5px;
  border-radius: 20px;
  font-family: 'Poppins', sans-serif;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.2s ease;
}

.action-btn:hover {
  background-color: #e4e6eb;
  transform: translateY(-1px);
}

.icon-bg {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  background-color: rgba(121, 122, 122, 0.08);
  color: #555555;
  overflow: hidden;
}

.box {
  background-color: #ffffff;
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 0;
}

.songs-box {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-height: 0;
}

.box-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  flex-shrink: 0;
}

.box-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.loading-text, .empty-text {
  font-size: 14px;
  color: #777;
  padding: 10px 0;
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 16px;
}

.card {
  background: transparent;
  border: none;
  border-radius: 0;
  overflow: visible;
  cursor: pointer;
  transition: transform 0.2s ease;
  box-shadow: none;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.card:hover {
  transform: translateY(-4px);
}

.card-img-wrap {
  position: relative;
  width: 100%;
  aspect-ratio: 1;
  background-color: #f8fafc;
  overflow: hidden;
  border-radius: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}

.card-img-wrap img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.card:hover .card-img-wrap img {
  transform: scale(1.03);
}

.song-badge {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(4px);
  color: #ffffff;
  font-size: 11px;
  font-weight: 500;
  padding: 3px 8px;
  border-radius: 20px;
  z-index: 2;
  letter-spacing: 0.2px;
  text-shadow: 1.2px 0.3px 0 rgba(97, 97, 97, 0.8);
}

.card-play-btn {
  position: absolute;
  bottom: 12px;
  right: 12px;
  width: 25px;
  height: 25px;
  border-radius: 50%;
  background: #fe2c56f2;  
  color: #ffffff;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 4px 14px rgba(0, 0, 0, 0.18);
  z-index: 3;
  padding: 0;
  padding-left: 3px;
  transition: transform 0.2s ease, background-color 0.2s ease;
}

.card-play-btn svg {
  filter: drop-shadow(1px 1.5px 0 #000);
}

.card-play-btn:hover {
  background-color: #f0f0f0;
  transform: scale(1.08);
}

.card-desc {
  padding: 0 2px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.card-title {
  font-size: 15px;
  font-weight: 600;
  color: #111;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-subtitle {
  font-size: 13px;
  color: #666;
  font-weight: 400;
}

.song-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  flex: 1;
  overflow-y: auto;
  min-height: 0;
  
  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.song-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.2s;
  flex-shrink: 0;
  border: 1px solid #f1f3f5;
}

.song-item:hover {
  background-color: #f3f3f4;
}

.song-item.is-playing {
  background-color: #edf0f3;
}

.thumb {
  position: relative;
  width: 45px;
  height: 45px;
  border-radius: 6px;
  overflow: hidden;
  flex-shrink: 0;
  background: #eee;
}

.thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.thumb-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.song-item:hover .thumb-overlay {
  opacity: 1;
}

.thumb-play-btn {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background-color: #fe2c56f2;
  color: #ffffff;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  padding: 0;
  padding-left: 2px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.2);
  transition: transform 0.2s ease;
}

.thumb-play-btn:hover {
  transform: scale(1.1);
}

.info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.name-box {
  font-size: 14px;
  font-weight: 600;
  color: #333;
}

.singer-box {
  font-size: 12px;
  color: #777;
}

.song-btns {
  padding-left: 4px;
}

.duration-badge {
  display: inline-flex;
  align-items: center;
  gap: 3px;
  font-size: 12px;
  color: #555;
  background: #f1f3f5;
  padding: 3px 6px;
  border-radius: 6px;
  font-weight: 500;
  border: 1px solid #e2e8f0;
}
</style>