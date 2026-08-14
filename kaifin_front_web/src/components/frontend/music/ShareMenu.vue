<script setup>
import { ref, watch } from 'vue'
// Import Lottie Vue Component
import { DotLottieVue } from '@lottiefiles/dotlottie-vue'

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  songData: {
    type: Object,
    default: () => ({
      title: 'Song name',
      singer: 'Singer',
      duration: '02:30',
      thumbnail: ''
    })
  }
})

const emit = defineEmits(['update:show', 'share', 'create-playlist'])

const currentView = ref('main')
const friendName = ref('')
const shareText = ref('')

const showStickerPanel = ref(false)

// 📁 បញ្ជី Array Object សម្រាប់ Sticker
const stickerList = ref([
  { id: 1, name: 'Song', path: '/src/assets/json_lot/bsong.json' },
  { id: 2, name: 'Crazy', path: '/src/assets/json_lot/crazy_1.json' },
  { id: 3, name: 'Test', path: '/src/assets/json_lot/test.json' },
  { id: 4, name: 'Cafe', path: '/src/assets/json_lot/cafe.json' },
  { id: 5, name: 'Like', path: '/src/assets/json_lot/like.json' },
  { id: 6, name: 'Cute', path: '/src/assets/json_lot/cute_d_1.json' },
  { id: 7, name: 'Car', path: '/src/assets/json_lot/car_1.json' },
  { id: 8, name: 'Ag', path: '/src/assets/json_lot/ag_1.json' }
])

const selectedSticker = ref(stickerList.value[0]?.path || '')

// 🎵 States សម្រាប់បង្កើត Playlist Name
const showCreatePlaylistInput = ref(false)
const playlistName = ref('')

// ⚙️ States សម្រាប់ Dropdown បន្ថែមនៅផ្នែកខាងក្រោមខាងឆ្វេង (Settings & Status)
const showAudienceDropdown = ref(false)
const selectedAudience = ref('Public')
const audienceOptions = [
  { name: 'Public', icon: 'globe' },
  { name: 'Friends', icon: 'users' },
  { name: 'Only me', icon: 'lock' }
]

const showSettingsDropdown = ref(false)
const selectedSettingOption = ref('Default')
const settingOptions = [
  { name: 'Default', icon: 'sliders' },
  { name: 'Allow Comments', icon: 'message-circle' },
  { name: 'Hide Preview', icon: 'eye-off' }
]

watch(
  () => props.show,
  (value) => {
    if (value) {
      currentView.value = 'main'
      friendName.value = ''
      shareText.value = ''
      showStickerPanel.value = false
      showCreatePlaylistInput.value = false
      playlistName.value = ''
      showAudienceDropdown.value = false
      showSettingsDropdown.value = false
    }
  }
)

const handleSelectSticker = (path) => {
  selectedSticker.value = path
  showStickerPanel.value = false
}

const handleShare = () => {
  emit('share', {
    friend: friendName.value,
    text: shareText.value,
    song: props.songData,
    sticker: selectedSticker.value,
    audience: selectedAudience.value,
    setting: selectedSettingOption.value
  })
}

// 📌 មុខងារបញ្ជាក់การបង្កើត Playlist
const confirmCreatePlaylist = () => {
  if (!playlistName.value.trim()) return
  emit('create-playlist', playlistName.value)
  playlistName.value = ''
  showCreatePlaylistInput.value = false
}

// 📌 មុខងារបិទប្រអប់បង្កើត Playlist
const cancelCreatePlaylist = () => {
  playlistName.value = ''
  showCreatePlaylistInput.value = false
}
</script>

<template>
  <div 
    v-if="show" 
    class="custom-action-popup" 
    :class="{ 
      'wide-popup': currentView === 'message-box' || currentView === 'add-text',
      'playlist-popup': currentView === 'playlist'
    }" 
    @click.stop
  >
    
    <!-- ፩. ម៉ឺនុយដើម (Main Menu) -->
    <template v-if="currentView === 'main'">
      <button class="popup-item">
        <span class="popup-icon play-next-icon">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
          <span class="plus-badge">+</span>
        </span>
        Play Next
      </button>

      <button class="popup-item" @click="currentView = 'playlist'">
        <span class="popup-icon">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="8" y1="6" x2="21" y2="6"></line><line x1="8" y1="12" x2="21" y2="12"></line><line x1="8" y1="18" x2="21" y2="18"></line><line x1="3" y1="6" x2="3.01" y2="6"></line><line x1="3" y1="12" x2="3.01" y2="12"></line><line x1="3" y1="18" x2="3.01" y2="18"></line></svg>
        </span>
        Add to playList
      </button>
      
      <button class="popup-item" @click="currentView = 'share'">
        <span class="popup-icon">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path><polyline points="16 6 12 2 8 6"></polyline><line x1="12" y1="2" x2="12" y2="15"></line></svg>
        </span>
        Share
      </button>

      <button class="popup-item">
        <span class="popup-icon">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path><polyline points="7 10 12 15 17 10"></polyline><line x1="12" y1="15" x2="12" y2="3"></line></svg>
        </span>
        Download
      </button>
      
      <button class="popup-item">
        <span class="popup-icon">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="2.18" ry="2.18"></rect><line x1="7" y1="2" x2="7" y2="22"></line><line x1="17" y1="2" x2="17" y2="22"></line><line x1="2" y1="12" x2="22" y2="12"></line><line x1="2" y1="7" x2="7" y2="7"></line><line x1="2" y1="17" x2="7" y2="17"></line><line x1="17" y1="17" x2="22" y2="17"></line><line x1="17" y1="7" x2="22" y2="7"></line></svg>
        </span>
        Watch video
      </button>
    </template>

    <!-- ፪. ម៉ឺនុយ Add to playlist -->
    <template v-else-if="currentView === 'playlist'">
      <div class="popup-header">
        <button class="popup-back-btn" @click="currentView = 'main'; showCreatePlaylistInput = false;">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>
        </button>
        <span class="popup-title-text">Add to playList</span>
      </div>

      <div class="popup-divider"></div>

      <button v-if="!showCreatePlaylistInput" class="popup-item" @click="showCreatePlaylistInput = true">
        <span class="popup-icon">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        </span>
        Create Name
      </button>

      <div v-if="showCreatePlaylistInput" class="create-playlist-container">
        <input 
          v-model="playlistName" 
          type="text" 
          placeholder="enter name palylist" 
          class="playlist-name-input"
          @keyup.enter="confirmCreatePlaylist"
        />
        <div class="playlist-actions">
          <button class="action-btn check-btn" @click="confirmCreatePlaylist" title="Confirm">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#22c55e" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="20 6 9 17 4 12"></polyline>
            </svg>
          </button>
          <button class="action-btn close-btn" @click="cancelCreatePlaylist" title="Cancel">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#ef4444" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <line x1="18" y1="6" x2="6" y2="18"></line>
              <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
          </button>
        </div>
      </div>
    </template>

    <!-- ፫. ម៉ឺនុយ Share -->
    <template v-else-if="currentView === 'share'">
      <div class="popup-header">
        <button class="popup-back-btn" @click="currentView = 'main'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>
        </button>
        <span class="popup-title-text">Share</span>
      </div>

      <div class="popup-divider"></div>

      <button class="popup-item" @click="currentView = 'add-text'">
        <span class="popup-icon">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 20h9"></path><path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"></path></svg>
        </span>
        Add your own text
      </button>

      <button class="popup-item" @click="currentView = 'message-box'">
        <span class="popup-icon">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path><polyline points="22,6 12,13 2,6"></polyline></svg>
        </span>
        Send as a message
      </button>

      <button class="popup-item active-item">
        <span class="popup-icon">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
        </span>
        Post to group
      </button>
      <button class="popup-item">
        <span class="popup-icon">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
        </span>
        Copy link
      </button>
    </template>

    <!-- ፬. ទិដ្ឋភាពផ្ទាំង Add your own text -->
    <template v-else-if="currentView === 'add-text'">
      <div class="popup-header">
        <button class="popup-back-btn" @click="currentView = 'share'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>
        </button>
        <span class="popup-title-text">Share in the News Feed</span>
      </div>

      <div class="popup-divider"></div>

      <div class="msg-form">
        <div class="textarea-wrapper">
          <textarea v-model="shareText" placeholder="Enter Text" class="msg-textarea large-textarea"></textarea>
          
          <span class="emoji-icon" @click="showStickerPanel = !showStickerPanel" title="Selected Sticker">
            <DotLottieVue
              :src="selectedSticker" 
              :width="24"
              :height="24"
              autoplay
              loop
              class="karaoke-svg"
            />
          </span>

          <div v-if="showStickerPanel" class="sticker-panel">
            <div class="sticker-grid">
              <div 
                v-for="item in stickerList" 
                :key="item.id" 
                class="sticker-item" 
                @click="handleSelectSticker(item.path)"
              >
                <DotLottieVue 
                  :src="item.path" 
                  autoplay 
                  loop 
                  class="panel-lottie" 
                />
                <span class="tooltip-text">{{ item.name }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="song-preview-card">
          <div class="preview-thumb">
            <img v-if="songData.thumbnail" :src="songData.thumbnail" alt="Thumb" class="thumb-img" />
          </div>
          <div class="preview-info">
            <span class="preview-title">{{ songData.title }}</span>
            <span class="preview-singer">{{ songData.singer }}</span>
          </div>
          <span class="preview-duration">
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
            {{ songData.duration }}
          </span>
        </div>

        <!-- 🌟 បន្ថែម Setting និង Status Dropdown พร้อม SVG icon និង background ថ្លា (border-radius: 50%) -->
        <div class="bottom-controls-container">
          <!-- Status Dropdown -->
          <div class="custom-dropdown-wrapper">
            <button class="dropdown-trigger-btn" @click="showAudienceDropdown = !showAudienceDropdown; showSettingsDropdown = false;">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="2" y1="12" x2="22" y2="12"></line><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path></svg>
              <span>{{ selectedAudience }}</span>
              <svg width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
            </button>
            <div v-if="showAudienceDropdown" class="dropdown-menu-list">
              <div 
                v-for="opt in audienceOptions" 
                :key="opt.name" 
                class="dropdown-menu-item"
                @click="selectedAudience = opt.name; showAudienceDropdown = false;"
              >
                <!-- 🌟 SVG Icon ក្នុង list item ដែលមាន Background ថ្លា និង border-radius: 50% -->
                <span class="dropdown-item-icon-bg">
                  <svg v-if="opt.icon === 'globe'" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="2" y1="12" x2="22" y2="12"></line><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path></svg>
                  <svg v-else-if="opt.icon === 'users'" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
                  <svg v-else-if="opt.icon === 'lock'" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
                </span>
                <span>{{ opt.name }}</span>
              </div>
            </div>
          </div>

          <!-- Setting Dropdown -->
          <div class="custom-dropdown-wrapper">
            <button class="dropdown-trigger-btn" @click="showSettingsDropdown = !showSettingsDropdown; showAudienceDropdown = false;">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"></circle><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path></svg>
              <span>{{ selectedSettingOption }}</span>
              <svg width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
            </button>
            <div v-if="showSettingsDropdown" class="dropdown-menu-list">
              <div 
                v-for="opt in settingOptions" 
                :key="opt.name" 
                class="dropdown-menu-item"
                @click="selectedSettingOption = opt.name; showSettingsDropdown = false;"
              >
                <!-- 🌟 SVG Icon ក្នុង list item ដែលមាន Background ថ្លា និង border-radius: 50% -->
                <span class="dropdown-item-icon-bg">
                  <svg v-if="opt.icon === 'sliders'" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="4" y1="21" x2="4" y2="14"></line><line x1="4" y1="10" x2="4" y2="3"></line><line x1="12" y1="21" x2="12" y2="12"></line><line x1="12" y1="8" x2="12" y2="3"></line><line x1="20" y1="21" x2="20" y2="16"></line><line x1="20" y1="12" x2="20" y2="3"></line><line x1="1" y1="14" x2="7" y2="14"></line><line x1="9" y1="8" x2="15" y2="8"></line><line x1="17" y1="16" x2="23" y2="16"></line></svg>
                  <svg v-else-if="opt.icon === 'message-circle'" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"></path></svg>
                  <svg v-else-if="opt.icon === 'eye-off'" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path><line x1="1" y1="1" x2="23" y2="23"></line></svg>
                </span>
                <span>{{ opt.name }}</span>
              </div>
            </div>
          </div>

          <!-- ប៊ូតុង Share នៅខាងស្ដាំដដែល -->
          <button class="share-submit-btn" @click="handleShare">
            <span class="share-icon-bg">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path>
                <polyline points="16 6 12 2 8 6"></polyline>
                <line x1="12" y1="2" x2="12" y2="15"></line>
              </svg>
            </span>
            Share
          </button>
        </div>
      </div>
    </template>

    <!-- ፭. ទិដ្ឋភាពផ្ទាំង Send as a message -->
    <template v-else-if="currentView === 'message-box'">
      <div class="popup-header">
        <button class="popup-back-btn" @click="currentView = 'share'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>
        </button>
        <span class="popup-title-text">Share as a message</span>
      </div>

      <div class="popup-divider"></div>

      <div class="msg-form">
        <input v-model="friendName" type="text" placeholder="Tag friend name" class="msg-input" />
        
        <div class="textarea-wrapper">
          <textarea v-model="shareText" placeholder="Enter Text" class="msg-textarea"></textarea>
          
          <span class="emoji-icon" @click="showStickerPanel = !showStickerPanel" title="Selected Sticker">
            <DotLottieVue
              :src="selectedSticker" 
              :width="24"
              :height="24"
              autoplay
              loop
              class="karaoke-svg"
            />
          </span>

          <div v-if="showStickerPanel" class="sticker-panel">
            <div class="sticker-grid">
              <div 
                v-for="item in stickerList" 
                :key="item.id" 
                class="sticker-item" 
                @click="handleSelectSticker(item.path)"
              >
                <DotLottieVue 
                  :src="item.path" 
                  autoplay 
                  loop 
                  class="panel-lottie" 
                />
                <span class="tooltip-text">{{ item.name }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="song-preview-card">
          <div class="preview-thumb">
            <img v-if="songData.thumbnail" :src="songData.thumbnail" alt="Thumb" class="thumb-img" />
          </div>
          <div class="preview-info">
            <span class="preview-title">{{ songData.title }}</span>
            <span class="preview-singer">{{ songData.singer }}</span>
          </div>
          <span class="preview-duration">
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
            {{ songData.duration }}
          </span>
        </div>

        <div class="btn-container">
          <button class="share-submit-btn" @click="handleShare">
            <span class="share-icon-bg">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path>
                <polyline points="16 6 12 2 8 6"></polyline>
                <line x1="12" y1="2" x2="12" y2="15"></line>
              </svg>
            </span>
            Share
          </button>
        </div>
      </div>
    </template>

  </div>
</template>

<style scoped>
.custom-action-popup {
  position: absolute;
  right: 0;
  bottom: 38px;
  width: 180px;
  background-color: #111827;
  border: none;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  padding: 6px;
  z-index: 100;
  animation: popup-in 0.15s ease-out;
  gap: 4px;
  transition: width 0.2s ease;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.5);
}

.custom-action-popup.playlist-popup {
  width: 250px;
}

.custom-action-popup.wide-popup {
  width: 400px;
}

.popup-header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 4px 4px 6px 4px;
}

.popup-back-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2px;
  border-radius: 4px;
  transition: background 0.2s;
}

.popup-back-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.popup-title-text {
  font-size: 13px;
  font-weight: 600;
  color: #ffffff;
}

.popup-divider {
  height: 1px;
  background-color: #1f2937;
  margin: 0 2px 6px 2px;
}

.popup-item {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
  background: transparent;
  border: none;
  padding: 6px 8px;
  text-align: left;
  font-size: 12px;
  color: #ffffff;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.25s ease;
}

.popup-item:hover,
.popup-item.active-item {
  transform: translateY(-1px);
}

.popup-icon {
  width: 28px;
  height: 28px;
  background: rgba(255, 255, 255, 0.12);
  border: none;
  border-radius: 50%;
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.15);
  transition: all 0.2s ease;
}

.popup-item:hover .popup-icon {
  background: rgba(255, 255, 255, 0.2);
}

.create-playlist-container {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 4px;
}

.playlist-name-input {
  flex: 1;
  background-color: #1f2937;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 6px;
  padding: 6px 10px;
  font-size: 11px;
  color: #ffffff;
  outline: none;
  min-width: 0;
}

.playlist-name-input::placeholder {
  color: #9ca3af;
}

.playlist-actions {
  display: flex;
  align-items: center;
  gap: 6px;
  flex-shrink: 0;
}

.action-btn {
  width: 28px;
  height: 28px;
  background: rgba(255, 255, 255, 0.06);
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  border-radius: 50%;
  transition: all 0.2s ease;
}

.action-btn:hover {
  background: rgba(255, 255, 255, 0.15);
}

.msg-form {
  display: flex;
  flex-direction: column;
  gap: 6px;
  padding: 2px;
}

.msg-input, .msg-textarea {
  width: 100%;
  background-color: #1f2937;
  border: 1px solid #ffffff2b;
  border-radius: 6px;
  padding: 6px 8px;
  font-size: 11px;
  color: #ffffff;
  outline: none;
  box-sizing: border-box;
}

.msg-input::placeholder, .msg-textarea::placeholder {
  color: #9ca3af;
}

.textarea-wrapper {
  position: relative;
}

.msg-textarea {
  resize: none;
  height: 45px;
  padding-right: 34px;
}

.large-textarea {
  height: 90px;
}

.emoji-icon {
  position: absolute;
  right: 8px;
  bottom: 10px;
  width: 28px;
  height: 28px;

  display: flex;
  align-items: center;
  justify-content: center;

  cursor: pointer;
  border-radius: 4px;
  
  border: none;
  background: rgba(255,255,255,0.08);

  box-shadow:
    0 0 8px rgba(255,255,255,0.25),
    inset 0 0 6px rgba(255,255,255,0.15);

  backdrop-filter: blur(6px);
  overflow: hidden;

  transition: all 0.25s ease;
}

.emoji-icon :deep(canvas),
.emoji-icon :deep(dot-lottie-player) {
  width: 100% !important;
  height: 100% !important;
  transform: scale(1.4);
  object-fit: contain;
  pointer-events: none;
}

.karaoke-svg {
  width: 100%;
  height: 100%;
}

.sticker-panel {
  position: absolute;
  right: 0;
  bottom: 45px;
  width: 240px;
  background-color: #1a1612;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 10px;
  padding: 8px;
  z-index: 110;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.6), 0 0 12px rgba(255, 255, 255, 0.15);
}

.sticker-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 6px;
}

.sticker-item {
  position: relative;
  background-color: #27211b;
  border: none;
  border-radius: 6px;
  height: 45px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  overflow: visible;
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.15), inset 0 0 4px rgba(255, 255, 255, 0.1);
}

.sticker-item :deep(canvas),
.sticker-item :deep(dot-lottie-player),
.panel-lottie {
  width: 100% !important;
  height: 100% !important;
  object-fit: contain;
  pointer-events: none;
}

.sticker-item .tooltip-text {
  visibility: hidden;
  position: absolute;
  bottom: 115%;
  left: 50%;
  transform: translateX(-50%);
  background-color: #ffffff;
  color: #111827;
  font-size: 10px;
  font-weight: 600;
  padding: 3px 8px;
  border-radius: 4px;
  white-space: nowrap;
  z-index: 120;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
  opacity: 0;
  transition: opacity 0.2s ease;
  pointer-events: none;
}

.sticker-item:hover .tooltip-text {
  visibility: visible;
  opacity: 1;
}

.song-preview-card {
  display: flex;
  align-items: center;
  background-color: #1f2937;
  border: none;
  border-radius: 6px;
  padding: 4px;
  gap: 8px;
}

.preview-thumb {
  width: 32px;
  height: 32px;
  background-color: #374151;
  border-radius: 4px;
  flex-shrink: 0;
  overflow: hidden;
}

.thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.preview-info {
  display: flex;
  flex-direction: column;
  flex: 1;
  overflow: hidden;
}

.preview-title {
  font-size: 11px;
  font-weight: 600;
  color: #ffffff;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.preview-singer {
  font-size: 10px;
  color: #9ca3af;
}

.preview-duration {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 10px;
  background-color: #374151;
  color: #ffffff;
  padding: 2px 6px;
  border-radius: 4px;
}

.btn-container {
  display: flex;
  justify-content: flex-end;
  margin-top: 4px;
}

.share-submit-btn {
  background-color: #1d70b8;
  color: #ffffff;
  border: none;
  border-radius: 32px;
  padding: 6px 16px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: background-color 0.2s;
}

.share-submit-btn:hover {
  background-color: #155799;
}

.share-icon-bg {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 22px;
  height: 22px;
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  flex-shrink: 0;
}

/* 🌟 ស្តាយសម្រាប់ផ្ទាំងបញ្ជាផ្នែកខាងក្រោម (Bottom Controls Layout) */
.bottom-controls-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 4px;
  gap: 6px;
}

.custom-dropdown-wrapper {
  position: relative;
}

.dropdown-trigger-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  background-color: #1f2937;
  color: #ffffff;
  border: 1px solid rgba(255, 255, 255, 0.15);
  border-radius: 32px;
  padding: 8px 12px;
  font-size: 10px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
}

.dropdown-trigger-btn:hover {
  background-color: #374151;
}

.dropdown-menu-list {
  position: absolute;
  left: 0;
  bottom: 110%;
  background-color: #1f2937;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 6px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.4);
  z-index: 120;
  min-width: 130px;
  overflow: hidden;
  padding: 4px;
}

.dropdown-menu-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 8px;
  font-size: 11px;
  color: #ffffff;
  cursor: pointer;
  border-radius: 4px;
  transition: background 0.15s;
  white-space: nowrap;
}

.dropdown-menu-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

/* 🌟 រចនាប័ទ្ម background ថ្លា និង border-radius: 50% សម្រាប់ Icon ក្នុង Dropdown Menu Items */
.dropdown-item-icon-bg {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 22px;
  height: 22px;
  background-color: rgba(255, 255, 255, 0.15);
  border-radius: 50%;
  flex-shrink: 0;
}
</style>