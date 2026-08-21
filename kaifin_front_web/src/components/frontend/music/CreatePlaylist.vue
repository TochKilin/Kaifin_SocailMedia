<template>
  <div class="create-playlist-container">
    <!-- Header Section Container -->
    <div class="wrap-gap">
      <div class="section-box section-box-raduis">
        <div class="create-header-content">
          <!-- Add Photo Box with Hidden File Input -->
          <div class="add-photo-box" @click="triggerImageUpload">
            <input 
              type="file" 
              ref="fileInputRef" 
              style="display: none" 
              accept="image/*" 
              @change="handleImageSelected" 
            />
            <template v-if="playlistData.coverUrl">
              <img :src="playlistData.coverUrl" alt="Playlist Cover" class="uploaded-cover-img" />
            </template>
            <template v-else>
              <div class="pencil-icon-wrap">
                <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M12 20h9"></path>
                  <path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"></path>
                </svg>
              </div>
              <span class="add-photo-text">Add Photo</span>
            </template>
          </div>

          <!-- Meta Info -->
          <div class="create-meta-info">
            <span class="public-badge-pill" @click="isEditModalOpen = true" style="cursor: pointer;">
              <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <line x1="2" y1="12" x2="22" y2="12"></line>
                <path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path>
              </svg>
              Public Playlist
            </span>
            <EditDetailsModal 
              v-if="isEditModalOpen" 
              :playlistData="playlistData"
              @close="isEditModalOpen = false"
              @save="handleUpdatePlaylist"
            />
            <input 
              type="text" 
              class="playlist-name-input" 
              v-model="playlistData.name" 
              placeholder="Name Playlist" 
            />
            <div class="playlist-owner-row">
              <div class="user-avatar-badge">
                <img src="../../../assets/animate/cat.svg" alt="kkk">
              </div>
              <span class="owner-name-pill">Kilin</span>
              <span class="song-count-pill">{{ playlistSongs.length }} song{{ playlistSongs.length === 1 ? '' : 's' }}</span>
              <span class="total-duration-pill">3 min 54 sec</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Controls Bar Container -->
      <div class="section-box controls-bar-box controls-bar-raduis">
        <div class="controls-left-group">
          <button class="big-play-circle-btn">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
              <polygon points="5 3 19 12 5 21 5 3"></polygon>
            </svg>
          </button>
          <div class="control-icon-thumb">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
              <circle cx="8.5" cy="8.5" r="1.5"></circle>
              <polyline points="21 15 16 10 5 21"></polyline>
            </svg>
          </div>
          <div class="control-icon-action">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="16 3 21 3 21 8"></polyline>
              <line x1="4" y1="20" x2="21" y2="3"></line>
              <polyline points="21 16 21 21 16 21"></polyline>
              <line x1="15" y1="15" x2="21" y2="21"></line>
              <line x1="4" y1="4" x2="9" y2="9"></line>
            </svg>
          </div>
          <div class="control-icon-action">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
              <polyline points="7 10 12 15 17 10"></polyline>
              <line x1="12" y1="15" x2="12" y2="3"></line>
            </svg>
          </div>

          <!-- Dots Menu Wrapper with Popup -->
          <div class="dots-wrapper" style="position: relative;" ref="menuWrapperRef">
            <div class="dots-group" @click.stop="toggleMenu" style="cursor: pointer;">
              <span class="dot"></span>
              <span class="dot"></span>
              <span class="dot"></span>
            </div>

            <!-- Popup Menu Component -->
            <div class="menu-container" v-if="isMenuOpen">
              <!-- Option 1: Remove from profile -->
              <div class="menu-item" @click="handleAction('remove')">
                <div class="icon-wrapper remove-icon">
                  <div class="icon-circle"></div>
                  <div class="icon-circle-outline"></div>
                </div>
                <span class="menu-text">Remove from profile</span>
              </div>

              <!-- Option 2: Edit details -->
              <div class="menu-item" @click="handleAction('edit')">
                <div class="icon-wrapper edit-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                    <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                  </svg>
                </div>
                <span class="menu-text">Edit details</span>
              </div>

              <!-- Option 3: Delete -->
              <div class="menu-item" @click="handleAction('delete')">
                <div class="icon-wrapper delete-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                    <circle cx="9" cy="7" r="4"></circle>
                    <line x1="23" y1="18" x2="17" y2="18"></line>
                  </svg>
                </div>
                <span class="menu-text delete-text">Delete</span>
              </div>

              <!-- Option 4: Make private -->
              <div class="menu-item" @click="handleAction('private')">
                <div class="icon-wrapper private-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
                    <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
                  </svg>
                </div>
                <span class="menu-text">Make private</span>
              </div>
            </div>
          </div>
        </div>

        <div class="controls-right-group">
          <button class="list-mode-btn active">List</button>
          <button class="menu-icon-btn">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <line x1="8" y1="6" x2="21" y2="6"></line>
              <line x1="8" y1="12" x2="21" y2="12"></line>
              <line x1="8" y1="18" x2="21" y2="18"></line>
              <line x1="3" y1="6" x2="3.01" y2="6"></line>
              <line x1="3" y1="12" x2="3.01" y2="12"></line>
              <line x1="3" y1="18" x2="3.01" y2="18"></line>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <div class="wrap-contents">
      <!-- Playlist Table Container -->
      <div class="section-box playlist-table-box playlist-table-box-raduis">
        <div class="table-header-row">
          <span class="th-col title-col"># Title</span>
          <span class="th-col album-col">Album</span>
          <span class="th-col date-col">Date added</span>
          <span class="th-col time-col">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <polyline points="12 6 12 12 16 14"></polyline>
            </svg>
          </span>
          <span class="th-col arrow-col">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="6 9 12 15 18 9"></polyline>
            </svg>
          </span>
        </div>

        <div class="table-body-rows">
          <div v-for="(song, idx) in playlistSongs" :key="song.id" class="table-song-row">
            <div class="td-col title-col-wrap">
              <span class="row-num">{{ idx + 1 }}</span>
              <div class="row-thumb">
                <img :src="song.cover_url" alt="cover" />
              </div>
              <div class="row-song-meta">
                <span class="row-song-title">{{ song.title }}</span>
                <span class="row-song-singer">{{ song.singer_name }}</span>
              </div>
            </div>
            <div class="td-col album-col-wrap">
              <span class="pill-tag">{{ song.album }}</span>
            </div>
            <div class="td-col date-col-wrap">
              <span class="date-text">{{ song.date_added }}</span>
            </div>
            <div class="td-col time-col-wrap">
              <span class="time-pill">03:54</span>
            </div>
            <div class="td-col action-col-wrap" ref="columnMenuWrapperRef" @click.stop="toggleColumnMenu">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="6 9 12 15 18 9"></polyline>
              </svg>
               <Dropdown v-if="isColumnMenuOpen" class="absolute-dropdown" :modelValue="columnsVisibility" @toggle="toggleColumn" @close="isColumnMenuOpen = false"/>
            </div>
           
          </div>
        </div>
      </div>

      <!-- Recommends Container -->
      <div class="section-box recommends-box playlist-table-box-raduis">
        <div class="recommends-header">
          <span class="recommends-title">Recommends</span>
        </div>

        <div class="recommends-list">
          <div v-for="rec in recommendedSongs" :key="rec.id" class="recommend-item-row">
            <div class="rec-left">
              <div class="rec-thumb">
                <img :src="rec.cover_url" alt="cover" />
              </div>
              <div class="rec-meta">
                <span class="rec-title">{{ rec.title }}</span>
                <span class="rec-singer">{{ rec.singer_name }}</span>
              </div>
            </div>
            
            <div class="rec-center-pill">
              <span>Song name</span>
            </div>

            <div class="rec-right">
              <button class="add-btn" @click="addRecommendedSong(rec)">ADD</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import Dropdown from './Dropdown.vue'
import EditDetailsModal from './EditDetailsModal.vue'

const emit = defineEmits(['back'])
const fileInputRef = ref(null)
const menuWrapperRef = ref(null)
const isMenuOpen = ref(false)
const isEditModalOpen = ref(false)

const playlistData = ref({
  name: '',
  coverUrl: '',
  isPublic: true
})

const playlistSongs = ref([
  {
    id: 1,
    title: 'Midnight Echoes',
    singer_name: 'Tep Piseth',
    album: 'To Anyone',
    date_added: '10 mins ago',
    duration: 234,
    cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSiIknIgBC5LILUa8i0Jul-3m1EZCFD5GRlU-3c6o_XDA&s=10'
  }
])

const recommendedSongs = ref([
  {
    id: 101,
    title: 'Tep Piseth 2025',
    singer_name: 'Chay vireakYuth',
    cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ4iX88lF3fphu5mThRKOcbtwmidRipu9sxZ-WSwT717Q&s=10'
  },
  {
    id: 102,
    title: 'ប៉ាបងអ្នកមាន',
    singer_name: 'Solstice',
    cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR8jtxGxaeMhBqD26s7ysY2_IDSv3UKnjsF4p8yYBYwNw&s=10'
  },
  {
    id: 103,
    title: 'Velvet Horizon',
    singer_name: 'Solstice',
    cover_url: 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?w=100&auto=format&fit=crop&q=80'
  },
  {
    id: 104,
    title: 'Sora song',
    singer_name: 'Haiyu',
    cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS2EnymqWJN5RoQ_jLhtvxH_2vgxyCPLHHchwR_xzlCcw&s=10'
  },
  {
    id: 105,
    title: 'The way you are',
    singer_name: 'Solstice',
    cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSoI-winRbs0EcnXQyaDuM_u1LdXLj7AmoHxGuDEwyHIA&s=10'
  },
  {
    id: 106,
    title: 'Can you are',
    singer_name: 'Solstice',
    cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRpY82edXovIFpClyxChmjpsqq0O-eK85CKT0DpCD3ZCg&s=10'
  },
  {
    id: 107,
    title: 'កុំច្រឡុំ',
    singer_name: 'Solstice',
    cover_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTNSU_1l2B4rDgXDO_AAZ2xesNI0zNXkrx1Bg8h8qTQDQ&s=10'
  }
])

const triggerImageUpload = () => {
  fileInputRef.value?.click()
}

const handleImageSelected = (event) => {
  const file = event.target.files[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (e) => {
      playlistData.value.coverUrl = e.target.result
    }
    reader.readAsDataURL(file)
  }
}

const addRecommendedSong = (song) => {
  if (!playlistSongs.value.some(s => s.id === song.id)) {
    playlistSongs.value.push({
      ...song,
      album: 'Single Release',
      date_added: 'Just now',
      duration: 210
    })
  }
}

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const handleAction = (action) => {
  isMenuOpen.value = false 
  switch (action) {
    case 'remove':
      console.log('Removing from profile...')
      break
    case 'edit':
      console.log('Editing details...')
      break
    case 'delete':
      console.log('Deleting item...')
      break
    case 'private':
      console.log('Making private...')
      break
  }
}

const columnMenuWrapperRef = ref(null)
const isColumnMenuOpen = ref(false)

const columnsVisibility = ref({
  album: true,
  dateAdded: true,
  duration: true
})

const toggleColumnMenu = () => {
  isColumnMenuOpen.value = !isColumnMenuOpen.value
}

const toggleColumn = (key) => {
  columnsVisibility.value[key] = !columnsVisibility.value[key]
}

const handleClickOutside = (event) => {
  if (menuWrapperRef.value && !menuWrapperRef.value.contains(event.target)) {
    isMenuOpen.value = false
  }
  if (columnMenuWrapperRef.value && !columnMenuWrapperRef.value.contains(event.target)) {
    isColumnMenuOpen.value = false
  }
}


const handleUpdatePlaylist = (updatedData) => {
  playlistData.value.name = updatedData.name
  playlistData.value.description = updatedData.description
  playlistData.value.coverUrl = updatedData.coverUrl
  playlistData.value.isPublic = updatedData.isPublic
  isEditModalOpen.value = false
}


onMounted(() => {
  window.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  window.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@300;400;500;600;700&display=swap');

.create-playlist-container {
  font-family: 'Plus Jakarta Sans', sans-serif;
  display: flex;
  flex-direction: column;
  gap: 14px;
  background-color: transparent;
  border-radius: 14px;
  min-height: 100%;
  color: #1f2328;
}

.wrap-gap {
  gap: 4px;
  display: flex;
  flex-direction: column;
}

.wrap-contents {
  gap: 12px;
  display: flex;
  flex-direction: column;
}

.section-box {
  background-color: #ffffff;
  padding: 16px 20px;
  border: none !important;
  box-shadow: none !important;
}

.section-box-raduis {
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
}

.create-header-content {
  display: flex;
  gap: 20px;
  align-items: center;
  flex-wrap: wrap;
}

.add-photo-box {
  width: 120px;
  height: 120px;
  border-radius: 12px;
  border: 1px solid #d0d7de;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  background: #f6f8fa;
  cursor: pointer;
  color: #57606a;
  transition: all 0.2s ease;
  flex-shrink: 0;
  overflow: hidden;
  position: relative;
}

.add-photo-box:hover {
  border-color: #0969da;
  color: #0969da;
  background: #ddf4ff;
}

.uploaded-cover-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.pencil-icon-wrap {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: #eaeef2;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #24292f;
}

.add-photo-text {
  font-size: 12px;
  font-weight: 500;
}

.create-meta-info {
  display: flex;
  flex-direction: column;
  gap: 10px;
  flex: 1;
}

.public-badge-pill {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  background: #F7F8FA;
  color: #0969da;
  font-size: 11px;
  padding: 3px 10px;
  border-radius: 6px;
  font-weight: 500;
  width: fit-content;
  border: 1px solid #b6e3ff;
}

.playlist-name-input {
  border: 1px solid #d0d7de;
  color: #24292f;
  font-size: 12px;
  font-weight: 600;
  padding: 8px 12px;
  border-radius: 8px;
  outline: none;
  font-family: 'Plus Jakarta Sans', sans-serif;
  width: 100%;
  max-width: 380px;
  transition: all 0.2s ease;
}

.playlist-name-input:focus {
  border-color: #0969da;
  background: #ffffff;
  box-shadow: 0 0 0 3px rgba(9, 105, 218, 0.15);
}

.playlist-name-input::placeholder {
  color: #8c959f;
}

.playlist-owner-row {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
  font-size: 12px;
  color: #57606a;
}

.user-avatar-badge {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: #eaeef2;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #24292f;
  overflow: hidden;
}

.user-avatar-badge img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.owner-name-pill {
  color: #24292f;
  font-weight: 500;
}

.song-count-pill, .total-duration-pill {
  color: #57606a;
}

.controls-bar-box {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
}

.controls-bar-raduis {
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.controls-left-group {
  display: flex;
  align-items: center;
  gap: 12px;
}

.big-play-circle-btn {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: #0969da;
  color: #ffffff;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.2s ease, background 0.2s ease;
}

.big-play-circle-btn:hover {
  background: #FC3358;
  transform: scale(1.05);
}

.control-icon-thumb, .control-icon-action, .menu-icon-btn {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  background: #f6f8fa;
  color: #57606a;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid #d0d7de;
  cursor: pointer;
  transition: all 0.2s ease;
}

.control-icon-thumb:hover, .control-icon-action:hover, .menu-icon-btn:hover {
  background: #eaeef2;
  color: #24292f;
  border-color: #afb8c1;
}

.dots-group {
  display: flex;
  gap: 3px;
  align-items: center;
  margin-left: 2px;
}

.dot {
  width: 4px;
  height: 4px;
  background: #57606a;
  border-radius: 50%;
}

.menu-container {
  background-color: #ffffff;
  position: absolute;
  top: 24px;
  left: 0;
  z-index: 100;
  border-radius: 12px;
  padding: 8px;
  display: flex;
  flex-direction: column;
  width: 240px;
  border: 1px solid rgba(0, 0, 0, 0.08);
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 8px 10px -6px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(8px);
}

.menu-item {
  display: flex;
  align-items: center;
  padding: 6px 12px;
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.2s ease;
}

.menu-item:hover {
  background-color: #f3f4f6; 
  transform: translateX(2px);
}

.icon-wrapper {
  width: 25px;
  height: 25px;
  margin-right: 12px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 6px;
  background-color: #f8fafc;
  transition: background-color 0.2s ease;
  position: relative;
}

.menu-item:hover .icon-wrapper {
  background-color: #ffffff;
}

.menu-text {
  color: #334155;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  font-size: 12px;
  font-weight: 500;
}

.icon-circle {
  width: 16px;
  height: 16px;
  border-radius: 50%;
  border: 2px solid #3b82f6;
  position: absolute;
}

.icon-circle-outline {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  border: 2px solid #3b82f6;
  position: absolute;
}

.edit-icon,
.private-icon {
  color: #64748b;
}

.delete-icon {
  color: #ef4444;
}

.menu-item:hover .delete-icon,
.menu-item:hover .delete-text {
  color: #dc2626;
}

.delete-text {
  color: #ef4444;
}

.controls-right-group {
  display: flex;
  align-items: center;
  gap: 8px;
}

.list-mode-btn {
  background: #f6f8fa;
  color: #57606a;
  border: 1px solid #d0d7de;
  padding: 6px 14px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
}

.list-mode-btn.active {
  background: #0969da;
  color: #ffffff;
  border-color: #0969da;
}

.playlist-table-box {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.playlist-table-box-raduis {
  border-radius: 12px;
}

.table-header-row {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr 80px 40px;
  align-items: center;
  padding: 4px 8px;
  color: #57606a;
  font-size: 12px;
  font-weight: 500;
  border-bottom: 1px solid #d0d7de;
  padding-bottom: 8px;
}

.table-body-rows {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.table-song-row {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr 80px 40px;
  align-items: center;
  padding: 6px 8px;
  border-radius: 8px;
  background: #f6f8fa;
  border: 1px solid #eaeef2;
  transition: background 0.2s ease;
}

.table-song-row:hover {
  background: #eaeef2;
}

.title-col-wrap {
  display: flex;
  align-items: center;
  gap: 10px;
}

.row-num {
  font-size: 12px;
  color: #57606a;
  width: 16px;
  text-align: center;
}

.row-thumb {
  width: 36px;
  height: 36px;
  border-radius: 6px;
  overflow: hidden;
  background: #eaeef2;
  flex-shrink: 0;
}

.row-thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.row-song-meta {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.row-song-title {
  font-size: 13px;
  font-weight: 500;
  color: #24292f;
}

.row-song-singer {
  font-size: 11px;
  color: #57606a;
}

.album-col-wrap .pill-tag {
  background: #eaeef2;
  color: #24292f;
  font-size: 11px;
  padding: 3px 8px;
  border-radius: 6px;
  border: 1px solid #d0d7de;
}

.date-col-wrap .date-text {
  font-size: 12px;
  color: #57606a;
}

.time-col-wrap .time-pill {
  color: #57606a;
  font-size: 12px;
}

.action-col-wrap {
  color: #57606a;
  display: flex;
  justify-content: center;
  position: relative;
}

.recommends-box {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.recommends-header {
  display: flex;
  align-items: center;
  border-bottom: 1px solid #d0d7de;
  padding-bottom: 8px;
}

.recommends-title {
  font-size: 14px;
  font-weight: 600;
  color: #24292f;
}

.recommends-list {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.recommend-item-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 10px;
  border: none !important;
  border-radius: 8px;
  transition: background 0.2s ease;
}

.recommend-item-row:hover {
  background: #eaeef2;
}

.rec-left {
  display: flex;
  align-items: center;
  gap: 10px;
  flex: 1;
}

.rec-thumb {
  width: 36px;
  height: 36px;
  border-radius: 6px;
  overflow: hidden;
  background: #eaeef2;
  flex-shrink: 0;
}

.rec-thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.rec-meta {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.rec-title {
  font-size: 13px;
  font-weight: 500;
  color: #24292f;
}

.rec-singer {
  font-size: 11px;
  color: #57606a;
}

.rec-center-pill {
  padding-right: 4px;
  color: #57606a;
  font-size: 12px;
}

.rec-right {
  display: flex;
  justify-content: flex-end;
}

.add-btn {
  background: #f6f8fa;
  color: #24292f;
  border: 1px solid #d0d7de;
  padding: 4px 12px;
  border-radius: 6px;
  font-size: 11px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.add-btn:hover {
  background: #0969da;
  color: #ffffff;
  border-color: #0969da;
}

.absolute-dropdown {
  position: absolute;
  top: 25px;
  right: 0;
  z-index: 100;
}
</style>