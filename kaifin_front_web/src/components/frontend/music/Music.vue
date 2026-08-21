<script setup>
import { ref, onMounted, onUnmounted, computed,watch } from 'vue'
import NavBar from '../navbar/NavBar.vue'
import SongDetial from './SongDetial.vue'
import ShareMenu from './ShareMenu.vue'
import MyMusic from './MyMusic.vue'
import MyHistory from './MyHistory.vue'


const currentTab = ref('popular')
const showPopup = ref(false)
const selectedSong = ref(null)
const showBanner = ref(true)
const topSongs = ref([])
const isLoading = ref(false)
const loadError = ref(null)
const volume = ref(0.6)
const currentTime = ref(0)
const duration = ref(0)
const isLoop = ref(true)
const playlistDetailLoading = ref(false)   
const playlistDetailError = ref(null) 

const showAllSongs = ref(false)
const currentPage = ref(1)
const historyPlaylists = ref([
  {
    id: 101,
    name: 'Top Hits 2026',
    cover: 'https://images.unsplash.com/photo-1470225620780-dba8ba36b745?w=300&auto=format&fit=crop&q=80',
    songsCount: 15
  },
  {
    id: 102,
    name: 'Late Night Coding',
    cover: 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?w=300&auto=format&fit=crop&q=80',
    songsCount: 24
  }
])

const historySongs = ref([
  {
    id: 1,
    title: 'Starboy',
    singer: 'The Weeknd ft. Daft Punk',
    cover: 'https://images.unsplash.com/photo-1514525253161-7a46d19cd819?w=300&auto=format&fit=crop&q=80',
    duration: 230,
    fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3'
  },
  {
    id: 2,
    title: 'Blinding Lights',
    singer: 'The Weeknd',
    cover: 'https://images.unsplash.com/photo-1498038432885-c6f3f1b912ee?w=300&auto=format&fit=crop&q=80',
    duration: 200,
    fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3'
  },
  {
    id: 3,
    title: 'Save Your Tears',
    singer: 'The Weeknd',
    cover: 'https://images.unsplash.com/photo-1511379938547-c1f69419868d?w=300&auto=format&fit=crop&q=80',
    duration: 215,
    fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-3.mp3'
  }
])

const clearHistory = () => {
  historyPlaylists.value = []
  historySongs.value = []
}

const loadSongDuration = (song) => {
  const audio = new Audio(song.fileUrl)
  audio.addEventListener('loadedmetadata', () => {
    song.duration = audio.duration
  })
}

const onTimeUpdate = () => {
  if (!audioRef.value) return
  currentTime.value = audioRef.value.currentTime
}

const onLoadedMetadata = () => {
  if (!audioRef.value) return
  duration.value = audioRef.value.duration
}

const seek = (event) => {
  if (!audioRef.value || !duration.value) return
  const bar = event.currentTarget
  const rect = bar.getBoundingClientRect()
  const ratio = (event.clientX - rect.left) / rect.width
  audioRef.value.currentTime = ratio * duration.value
}

const setVolume = (event) => {
  if (!audioRef.value) return
  const bar = event.currentTarget
  const rect = bar.getBoundingClientRect()
  const ratio = Math.min(1, Math.max(0, (event.clientX - rect.left) / rect.width))
  volume.value = ratio
  audioRef.value.volume = ratio
}

const formatTime = (secs) => {
  if (!secs || isNaN(secs)) return '0:00'
  const m = Math.floor(secs / 60)
  const s = Math.floor(secs % 60).toString().padStart(2, '0')
  return `${m}:${s}`
}

const formatDuration = (secs) => {
  if (!secs || isNaN(secs)) return '0:00'
  const minutes = Math.floor(secs / 60)
  const seconds = Math.floor(secs % 60).toString().padStart(2, '0')
  return `${minutes}:${seconds}`
}

const progressPercent = computed(() =>
  duration.value ? (currentTime.value / duration.value) * 100 : 0
)

const topPlaylists = ref([])
const playlistsLoading = ref(false)
const playlistsError = ref(null)

const fetchTopPlaylists = async () => {
  playlistsLoading.value = true
  playlistsError.value = null
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/playlists/top?limit=20`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        ...(token ? { Authorization: `Bearer ${token}` } : {})
      }
    })
    if (!res.ok) throw new Error(`HTTP ${res.status}`)
    const body = await res.json()
    const playlists = body?.data?.playlists ?? []
    topPlaylists.value = playlists.map((p) => ({
      id: p.id,
      name: p.name,
      songsCount: p.songs_count || 0,
      cover: p.cover_url || 'https://via.placeholder.com/400x400?text=No+Cover'
    }))
  } catch (err) {
    playlistsError.value = err.message || 'Failed to load playlists'
    console.error('fetchTopPlaylists error:', err)
  } finally {
    playlistsLoading.value = false
  }
}

const API_BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'

const fetchTopSongs = async () => {
  isLoading.value = true
  loadError.value = null
  try {
    const token = localStorage.getItem('token')
    const params = new URLSearchParams({
        page: currentPage.value,
        limit: showAllSongs.value ? 50 : 6,
        search: ''
    })
    const res = await fetch(`${API_BASE}/api/v1/front/songs/show?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        ...(token ? { Authorization: `Bearer ${token}` } : {})
      }
    })

    if (!res.ok) throw new Error(`HTTP ${res.status}`)

    const body = await res.json()
    const songs = body?.data?.songs ?? []

    topSongs.value = songs.map((s) => ({
      id: s.id,
      title: s.title,
      singer: s.singer_name ?? `Artist #${s.artist_id}`,
      cover: s.cover_url || 'https://via.placeholder.com/200x200?text=No+Cover',
      fileUrl: s.file_url,
      duration: s.duration || 0
    }))

    topSongs.value.forEach(song => {
      if(song.duration === 0){
        loadSongDuration(song)
      }
    })
  } catch (err) {
    loadError.value = err.message || 'Failed to load songs'
    console.error('fetchTopSongs error:', err)
  } finally {
    isLoading.value = false
  }
}

const showMoreSongs = async () => {
  showAllSongs.value = true
  await fetchTopSongs()
}

let previousBodyOverflow = ''
let previousHtmlOverflow = ''
let previousBodyPosition = ''
let previousBodyTop = ''
let previousBodyWidth = ''
let lockedScrollY = 0

const lockBackgroundScroll = () => {
  lockedScrollY = window.scrollY

  previousBodyOverflow = document.body.style.overflow
  previousHtmlOverflow = document.documentElement.style.overflow
  previousBodyPosition = document.body.style.position
  previousBodyTop = document.body.style.top
  previousBodyWidth = document.body.style.width

  document.documentElement.style.overflow = 'hidden'
  document.body.style.overflow = 'hidden'
  // Fix body in place too, so iOS Safari can't rubber-band scroll behind it
  document.body.style.position = 'fixed'
  document.body.style.top = `-${lockedScrollY}px`
  document.body.style.width = '100%'
}

const unlockBackgroundScroll = () => {
  document.documentElement.style.overflow = previousHtmlOverflow
  document.body.style.overflow = previousBodyOverflow
  document.body.style.position = previousBodyPosition
  document.body.style.top = previousBodyTop
  document.body.style.width = previousBodyWidth
  window.scrollTo(0, lockedScrollY)
}

onMounted(() => {
  fetchTopSongs()
  fetchTopPlaylists()
})

onUnmounted(() => {
  // unlockBackgroundScroll()
})

const openMorePopup = (song) => {
  selectedSong.value = song
  showPopup.value = true
}

const closePopup = () => {
  showPopup.value = false
}

const shareSongData = computed(() => ({
  title: selectedSong.value?.title || 'Song name',
  singer: selectedSong.value?.singer || 'Singer',
  duration: formatDuration(selectedSong.value?.duration),
  thumbnail: selectedSong.value?.cover || ''
}))

const handleShare = (payload) => {
  console.log('Shared:', payload)
  showPopup.value = false
}

const audioRef = ref(null)
const currentSong = ref(null)
const isPlaying = ref(false)

const playSong = (song) => {
  if (!song.fileUrl) {
    console.warn('This song has no file_url')
    return
  }

  if (currentSong.value?.id === song.id) {
    togglePlayPause()
    return
  }

  currentSong.value = song
  audioRef.value.loop = isLoop.value
  currentTime.value = 0
  duration.value = 0
  audioRef.value.src = song.fileUrl
  audioRef.value.play().catch((err) => {
    console.error('Playback failed:', err)
    isPlaying.value = false
  })
  isPlaying.value = true
}

const togglePlayPause = () => {
  if (!audioRef.value || !currentSong.value) return
  if (isPlaying.value) {
    audioRef.value.pause()
  } else {
    audioRef.value.play().catch((err) => {
      console.error('Playback failed:', err)
      isPlaying.value = false
    })
  }
  isPlaying.value = !isPlaying.value
}

const onAudioEnded = () => {
  if(isLoop.value){
    return
  }
  nextSong()
}

const addToPlaylist = (song) => {
  console.log('Add to playlist:', song.title)
}

const uploadFromVK = () => {
  console.log('Upload from VK clicked')
}

const skeletonCount = [1, 2, 3, 4]

const toggleLoop = () => {
  isLoop.value = !isLoop.value
  if(audioRef.value){
    audioRef.value.loop = isLoop.value
  }
}

const nextSong = () => {
  if (!currentSong.value || topSongs.value.length === 0) return
  const currentIndex = topSongs.value.findIndex(song => song.id === currentSong.value.id)
  let nextIndex = currentIndex + 1
  if (nextIndex >= topSongs.value.length) {
    nextIndex = 0
  }
  playSong(topSongs.value[nextIndex])
}

const previousSong = () => {
  if (!currentSong.value || topSongs.value.length === 0) return
  const currentIndex = topSongs.value.findIndex(song => song.id === currentSong.value.id)
  let prevIndex = currentIndex - 1
  if (prevIndex < 0) {
    prevIndex = topSongs.value.length - 1
  }
  playSong(topSongs.value[prevIndex])
}

const selectedPlaylist = ref(null)

async function openPlaylistDetail(item) {
  selectedPlaylist.value = { ...item, songs: [] }
  playlistDetailLoading.value = true
  playlistDetailError.value = null

  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/playlists/${item.id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
      },
    })
    if (!res.ok) throw new Error(`HTTP ${res.status}`)

    const body = await res.json()
    const data = body?.data ?? body
    const rawSongs = data.songs ?? []

    const songs = rawSongs.map((s) => ({
      id: s.id,
      title: s.title,
      singer: s.singer_name || `Artist #${s.artist_id}`,
      cover: s.cover_url || 'https://via.placeholder.com/200x200?text=No+Cover',
      fileUrl: s.file_url,
      duration: s.duration || 0,
    }))

    selectedPlaylist.value = {
      id: data.id,
      name: data.name,
      cover: data.cover_url,
      songsCount: data.songs_count,
      songs,
    }
  } catch (err) {
    console.error('Failed to load playlist detail', err)
    playlistDetailError.value = err.message || 'Failed to load playlist'
  } finally {
    playlistDetailLoading.value = false
  }
}

const backToMainView = () => {
  selectedPlaylist.value = null
}


</script>

<template>
  <div @click="showPopup = false">
    <div class="app-music">
      <div class="container" :class="{ 'no-right-sidebar': selectedPlaylist }">

        <!-- Top Bar -->
        <header class="top-bar">
          <div class="controls">
            <button
              v-if="selectedPlaylist"
              class="back-btn"
              @click="backToMainView"
            >
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none"
                   stroke="currentColor" stroke-width="2"
                   stroke-linecap="round" stroke-linejoin="round">
                <line x1="19" y1="12" x2="5" y2="12"></line>
                <polyline points="12 19 5 12 12 5"></polyline>
              </svg>
              Back
            </button>
            <button class="btn" title="Shuffle">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="16 3 21 3 21 8"></polyline><line x1="4" y1="20" x2="21" y2="3"></line><polyline points="21 16 21 21 16 21"></polyline><line x1="15" y1="15" x2="21" y2="21"></line><line x1="4" y1="4" x2="9" y2="9"></line></svg>
            </button>
            <button class="btn" title="Previous" @click="previousSong">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="19 20 9 12 19 4 19 20"></polygon><line x1="5" y1="4" x2="5" y2="20"></line></svg>
            </button>
            <button class="btn play" title="Play" @click="togglePlayPause">
              <svg v-if="!isPlaying" width="14" height="14" viewBox="0 0 24 24" fill="currentColor" stroke="none">
                <polygon points="5 3 19 12 5 21 5 3"></polygon>
              </svg>
              <svg v-else width="14" height="14" viewBox="0 0 24 24" fill="currentColor" stroke="none">
                <rect x="6" y="4" width="4" height="16"></rect>
                <rect x="14" y="4" width="4" height="16"></rect>
              </svg>
            </button>
            <button class="btn" title="Next" @click="nextSong">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="5 4 15 12 5 20 5 4"></polygon><line x1="19" y1="4" x2="19" y2="20"></line></svg>
            </button>
            <button class="btn" title="Repeat" :class="{ active: isLoop }" @click="toggleLoop">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="17 1 21 5 17 9"></polyline><path d="M3 11V9a4 4 0 0 1 4-4h14"></path><polyline points="7 23 3 19 7 15"></polyline><path d="M21 13v2a4 4 0 0 1-4 4H3"></path></svg>
            </button>
          </div>

          <div class="now-playing" :class="{ 'is-active': isPlaying }">
            <img v-if="currentSong?.cover" :src="currentSong.cover" class="now-cover" alt="cover" />
            <span class="mini-eq" v-if="isPlaying"></span>
            <div class="song-info">
              <span class="song-title">{{ currentSong?.title || 'No song playing' }}</span>
              <span class="song-artist">{{ currentSong?.singer || 'Unknown Artist' }}</span>
            </div>
          </div>

          <div class="right-tools">
            <button class="btn" title="Add">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            </button>
            <button class="btn" title="Share">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="18" cy="5" r="3"></circle><circle cx="6" cy="12" r="3"></circle><circle cx="18" cy="19" r="3"></circle><line x1="8.59" y1="13.51" x2="15.42" y2="17.49"></line><line x1="15.41" y1="6.51" x2="8.59" y2="10.49"></line></svg>
            </button>
            <div class="vol-box">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon>
                <path d="M19.07 4.93a10 10 0 0 1 0 14.14M15.54 8.46a5 5 0 0 1 0 7.07"></path>
              </svg>
              <div class="slider" @click="setVolume">
                <div class="slider-fill" :style="{ width: (volume * 100) + '%' }"></div>
                <div class="slider-thumb" :style="{ left: (volume * 100) + '%' }"></div>
              </div>
            </div>
            <input type="text" placeholder="Search" class="search-input" />
          </div>
        </header>

        <!-- Progress Bar Row -->
        <div class="progress-row">
          <div class="progress-bar" @click="seek">
            <div class="progress-fill" :style="{ width: progressPercent + '%' }"></div>
            <div class="progress-thumb" :style="{ left: progressPercent + '%' }"></div>
          </div>
        </div>

        <!-- Left Sidebar -->
        <aside class="sidebar">
          <div class="btn-3 card">
            <button @click="currentTab = 'popular'" :class="['nav-item', { active: currentTab === 'popular' }]">
              <span class="nav-label">
                <svg class="nav-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg>
                Popular
              </span>
            </button>
            <button @click="currentTab = 'mymusic'" :class="['nav-item', { active: currentTab === 'mymusic' }]">
              <span class="nav-label">
                <svg class="nav-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                My Music
              </span>
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
            </button>
            <button @click="currentTab = 'history'" :class="['nav-item', { active: currentTab === 'history' }]">
              <span class="nav-label">
                <svg class="nav-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                History
              </span>
            </button>
          </div>
        </aside>

        <!-- 4. Center Main Content -->
        <main class="content">

          <MyMusic
            v-if="currentTab === 'mymusic' && !selectedPlaylist"
            :api-base="API_BASE"
            :current-song="currentSong"
            :is-playing="isPlaying"
            @open-playlist="openPlaylistDetail"
            @play-song="playSong"
          />

          <MyHistory
            v-else-if="currentTab === 'history' && !selectedPlaylist"
            :is-loading="isLoading"
            :history-playlists="historyPlaylists"
            :history-songs="historySongs"
            :current-song="currentSong"
            :is-playing="isPlaying"
            :format-duration="formatDuration"
            @open-playlist="openPlaylistDetail"
            @play-song="playSong"
            @clear-history="clearHistory"
          />

          <!-- Popular Tab Content -->
          <template v-else-if="currentTab === 'popular' && !selectedPlaylist">

            <!-- Top Song -->
            <section class="box card">
              <div class="box-header">
                <span class="box-title">Top Song</span>
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="trend-icon"><polyline points="23 6 13.5 15.5 8.5 10.5 1 18"></polyline><polyline points="17 6 23 6 23 12"></polyline></svg>
              </div>

              <!-- Loading state -->
              <div v-if="isLoading" class="song-list">
                <div v-for="n in skeletonCount" :key="n" class="song-item skeleton">
                  <div class="thumb skeleton-block"></div>
                  <div class="info">
                    <div class="name-box">Loading...</div>
                  </div>
                </div>
              </div>

              <!-- Error state -->
              <div v-else-if="loadError" class="state-box error">
                <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="12"></line><line x1="12" y1="16" x2="12.01" y2="16"></line></svg>
                <p>Can't fetch songs — {{ loadError }}</p>
                <button class="retry-btn" @click="fetchTopSongs">Try again</button>
              </div>

              <!-- Empty state -->
              <div v-else-if="topSongs.length === 0" class="state-box">
                <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 18V5l12-2v13"></path><circle cx="6" cy="18" r="3"></circle><circle cx="18" cy="16" r="3"></circle></svg>
                <p>No songs yet</p>
              </div>

              <!-- Song List -->
              <div v-else class="song-list">
                <div
                  v-for="song in topSongs"
                  :key="song.id"
                  class="song-item"
                  :class="{ 'is-playing': currentSong?.id === song.id && isPlaying }"
                  @click="playSong(song)"
                >
                  <div class="thumb" :class="{ spinning: currentSong?.id === song.id && isPlaying }">
                    <img :src="song.cover" alt="cover" />
                    <span class="play-overlay" v-if="!(currentSong?.id === song.id && isPlaying)">
                      <svg width="20" height="20" viewBox="0 0 24 24" fill="white">
                        <polygon points="5 3 19 12 5 21 5 3"></polygon>
                      </svg>
                    </span>
                    <span class="eq-overlay" v-if="currentSong?.id === song.id && isPlaying">
                      <span></span><span></span><span></span>
                    </span>
                    <span class="dot" v-else></span>
                  </div>
                  <div class="info">
                    <div class="name-box">{{ song.title }}</div>
                    <div class="singer-box">{{ song.singer }}</div>
                  </div>

                  <div class="song-btns" style="position: relative;">
                    <span class="song-duration">
                      <svg class="clock-icon" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <circle cx="12" cy="12" r="10"></circle>
                      <polyline points="12 6 12 12 16 14"></polyline>
                      </svg>
                      {{ formatDuration(song.duration) }}
                    </span>
                    <button class="btn-sm" title="Add to playlist" @click.stop="addToPlaylist(song)">
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
                    </button>
                    <button class="btn-sm" title="More options" @click.stop="openMorePopup(song, $event)">
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="1"></circle><circle cx="12" cy="5" r="1"></circle><circle cx="12" cy="19" r="1"></circle></svg>
                    </button>

                    <ShareMenu
                      :show="showPopup && selectedSong?.id === song.id"
                      :song-data="shareSongData"
                      @update:show="showPopup = $event"
                      @share="handleShare"
                    />
                  </div>
                </div>
              </div>

              <button class="more-btn" v-if="!showAllSongs && topSongs.length >= 6" @click="showMoreSongs">
                Show more
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="display:inline; vertical-align:middle;"><polyline points="6 9 12 15 18 9"></polyline></svg>
              </button>
            </section>

            <!-- Top Playlist -->
            <section class="box card">
              <div class="box-header">
                <span class="box-title">Top Playlist</span>
              </div>
              <div class="grid">
                <div 
                  v-for="item in topPlaylists" 
                  :key="item.id" 
                  class=""
                  @click="openPlaylistDetail(item)"
                  style="cursor: pointer;"
                >
                  <div class="card-img-wrap">
                    <img :src="item.cover" alt="playlist" />
                    <span class="dot"></span>
                    <span class="card-play-overlay">
                      <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
                    </span>
                  </div>
                  <div class="card-desc">
                    <span class="card-title">{{ item.name }}</span>
                    <span class="badge">{{ item.songsCount }} Songs</span>
                  </div>
                </div>
              </div>
            </section>

          </template>

          <!-- Playlist Detail View -->
          <template v-else>
            <div class="song-detail-container">
              <SongDetial
                v-if="selectedPlaylist"
                :playlist="selectedPlaylist"
                :current-song="currentSong"
                :is-playing="isPlaying"
                :current-time="currentTime"
                :duration="duration"
                :volume="volume"
                :format-time="formatTime"
                :progress-percent="progressPercent"
                @play="playSong"
                @toggle-play="togglePlayPause"
                @seek="seek"
                @set-volume="setVolume"
                @back="backToMainView"
              />
            </div>
          </template>

        </main>

        <!-- Right Sidebar -->
        <aside class="sidebar-right" v-if="!selectedPlaylist">
          <div class="promo-banner" v-if="showBanner">
            <button class="close-banner" @click="showBanner = false">×</button>
            <div class="promo-content">
              <h2>playlists on Kaifin</h2>
              <p>Easy import from Kaifin</p>
              <button class="promo-btn" @click="uploadFromVK">Kaifin</button>
            </div>
            <div class="promo-image">
              <div class="music-note-icon">🎵</div>
            </div>
          </div>
        </aside>

      </div>
    </div>

    <!-- Audio Player Tag -->
    <audio
      ref="audioRef"
      @ended="onAudioEnded"
      @pause="isPlaying = false"
      @play="isPlaying = true"
      @timeupdate="onTimeUpdate"
      @loadedmetadata="onLoadedMetadata"
    ></audio>
  </div>
</template>
<style scoped>
/* Standard Clean & Modern Color Theme */
.app-music {
  color: #333333;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  height: 100vh;
  padding: 0;
  box-sizing: border-box;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* Top Bar */
.top-bar {
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
  width: 100%;
  grid-column: 1 / -1;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #ffffff;
  padding: 15px 20px;
  position: relative;
  overflow: hidden;
  flex-shrink: 0;
}

.controls, .right-tools {
  display: flex;
  align-items: center;
  gap: 12px;
}

.btn {
  background: none;
  border: none;
  font-size: 16px;
  cursor: pointer;
  color: #555555;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.2s, transform 0.15s;
}

.btn:hover {
  color: #007bff;
}

.btn:active {
  transform: scale(0.88);
}

.btn.play {
  background-color: #007bff;
  color: white;
  border-radius: 50%;
  width: 34px;
  height: 34px;
}

.btn.play:hover {
  color: white;
  box-shadow: 0 3px 10px rgba(0, 123, 255, 0.45);
}

.btn.play:active {
  transform: scale(0.85);
}

.now-playing {
  background-color: #3333330b;
  border: 1px solid #1976d21d;
  padding: 6px 30px;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 500;
  color: #495057;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: border-color 0.2s, background-color 0.2s;
}

.mini-eq {
  display: inline-flex;
  align-items: flex-end;
  gap: 2px;
  height: 12px;
}

.mini-eq span {
  width: 2.5px;
  background-color: #007bff;
  border-radius: 1px;
  animation: eq-bounce 0.9s ease-in-out infinite;
}

.mini-eq span:nth-child(1) { animation-delay: 0s; }
.mini-eq span:nth-child(2) { animation-delay: 0.2s; }
.mini-eq span:nth-child(3) { animation-delay: 0.4s; }

@keyframes eq-bounce {
  0%, 100% { height: 3px; }
  50% { height: 12px; }
}

.vol-box {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #555;
}

.slider {
  position: relative;
  width: 70px;
  height: 6px;
  background-color: #e9ecef;
  border-radius: 3px;
  cursor: pointer;
}

.slider-fill {
  height: 100%;
  background-color: #007bff;
  border-radius: 3px;
}

.slider-thumb {
  position: absolute;
  top: 50%;
  width: 10px;
  height: 10px;
  background: #1976D2;
  border: 2px solid #ffffff;
  border-radius: 50%;
  transform: translate(-50%, -50%);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.25);
  opacity: 0;
  transition: opacity 0.15s;
}

.vol-box:hover .slider-thumb {
  opacity: 1;
}

.search-input {
  border: 1px solid #ced4da;
  border-radius: 20px;
  padding: 6px 14px;
  outline: none;
  font-size: 13px;
  background: #ffffff;
  color: #333333;
  width: 140px;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.search-input:focus {
  border-color: #007bff;
  box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.12);
}

.search-input::placeholder {
  color: #adb5bd;
}

/* Layout */
.container {
  display: grid;
  grid-template-columns: 180px 1fr 200px;
  gap: 4px;
  background-color: #f0f0f0;
  padding-left: 0;
  padding-right: 0;
  flex: 1;
  min-height: 0;
  overflow: hidden;
}

.container.no-right-sidebar {
  grid-template-columns: 180px 1fr;
}

.progress-row {
  grid-column: 1 / -1;
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 12px;
  color: #6c757d;
  margin-top: -20px;
}

.time-label {
  min-width: 34px;
  text-align: center;
  font-variant-numeric: tabular-nums;
}

.progress-bar {
  position: relative;
  flex: 1;
  height: 6px;
  background-color: #afcce9;
  /* border-radius: 3px; */
  cursor: pointer;
  padding: 0;
}

.progress-fill {
  height: 100%;
  background-color: #007bff;
  /* border-radius: 3px; */
  transition: width 0.1s linear;
}

.progress-thumb {
  position: absolute;
  top: 50%;
  width: 11px;
  height: 11px;
  background: #1976D2;
  border: 2px solid #ffffff;
  border-radius: 50%;
  transform: translate(-50%, -50%);
  box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.15);
  transition: left 0.1s linear;
}

.sidebar {
  display: flex;
  flex-direction: column;
  gap: 10px;
  padding-left: 12px; 
  height: 100%;
  min-height: 0;
  overflow-y: auto;
}

.btn-3{
  display: flex;
  flex-direction: column;
  gap: 10px;
  height: 100%;
  padding: 12px; 
  background-color: #ffffff;

}

.nav-item {
  background-color: #ffffff;
  border: none;
  border-radius: 10px;
  padding: 8px 10px;
  text-align: left;
  font-size: 14px;
  cursor: pointer;
  color: #495057;
  display: flex;
  align-items: center;
  justify-content: space-between;
  transition: all 0.2s;
}

.nav-label {
  display: flex;
  align-items: center;
  gap: 10px;
}

.nav-icon {
  color: #555;
  transition: transform 0.2s;
}

.nav-item:hover {
  background-color: #f8f9fa;
  border-color: #d0d7de;
}

.nav-item:hover .nav-icon {
  color: #007bff;
  transform: scale(1.1);
}

.nav-item.active {
  color: #000000;
  border-color: #007bff;
}

.nav-item.active .nav-icon {
  color: rgb(7, 7, 7);
}

.content {
  display: flex;
  flex-direction: column;
  gap: 20px;
  min-height: 0;
  overflow-y: auto;
}

.sidebar-right {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding-right: 12px;
  min-height: 0;
  overflow-y: auto;
}

.promo-banner {
  position: relative;
  background: linear-gradient(90deg, #e3f2fd 0%, #bbdefb 100%);
  border: 1px solid #90caf9;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  overflow: hidden;
}

.promo-content {
  z-index: 2;
  max-width: 60%;
}

.promo-content h2 {
  font-size: 14px;
  font-weight: 600;
  margin: 0 0 6px 0;
  line-height: 1.3;
}

.promo-content p {
  font-size: 10px;
  opacity: 0.85;
  margin: 0 0 14px 0;
}
.promo-btn {
  background-color: #ffffff;
  color: #007bff;
  border: none;
  padding: 8px 14px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  width: 100%;
  
}

.promo-btn:hover {
  background-color: #f0f0f0;
  transform: translateY(-1px);
}

.promo-btn:active {
  transform: scale(0.96);
}

.promo-image {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 40px;
}

.music-note-icon {
  font-size: 64px;
  background: linear-gradient(135deg, #007bff, #007bff);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  filter: drop-shadow(0 4px 6px rgba(0,0,0,0.1));
  animation: float-note 3s ease-in-out infinite;
}

@keyframes float-note {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-6px); }
}

.close-banner {
  position: absolute;
  top: 12px;
  right: 14px;
  background: transparent;
  border: none;
  font-size: 18px;
  color: #666666;
  cursor: pointer;
  z-index: 3;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: background-color 0.2s, color 0.2s;
}

.close-banner:hover {
  background-color: rgba(0, 0, 0, 0.05);
  color: #000000;
}

.box {
  border-radius: 12px;
  padding: 20px;
  background-color: #ffffff;

}

.box-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding-bottom: 10px;
  color: #007bff;
}

.box-title {
  font-size: 18px;
  font-weight: 700;
  color: #212529;
}

/* Empty / error state with icon */
.state-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  padding: 32px 20px;
  color: #adb5bd;
  text-align: center;
}

.state-box p {
  margin: 0;
  font-size: 14px;
  color: #6c757d;
}

.state-box.error {
  color: #f1a3ac;
}

.state-box.error p {
  color: #dc3545;
}

.retry-btn {
  border: 1px solid #007bff;
  background: #ffffff;
  color: #007bff;
  border-radius: 8px;
  padding: 6px 16px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.retry-btn:hover {
  background: #007bff;
  color: #ffffff;
}

/* Songs (2 Columns Layout) */
.song-list {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 6px;
  margin-bottom: 16px;
}

.song-item {
  display: flex;
  align-items: center;
  border-radius: 4px;
  padding: 8px 14px;
  gap: 8px;
  transition: background-color 0.2s, transform 0.15s, box-shadow 0.15s;
  cursor: pointer;
}

.song-item:hover {
  transform: translateY(-2px);
  background-color: #0000000a;
}

.song-item.is-playing {
  border-color: #2f3132;
  background-color: #0505050f;
}

.thumb {
  position: relative;
  width: 65px;
  height: 65px;
  border-radius: 4px;
  overflow: hidden;
  /* border: 1px solid #dee2e6; */
  flex-shrink: 0;
  object-fit: contain;
}

.thumb.spinning img {
  animation: spin 6s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.thumb img, .card-img-wrap img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 4px;
}

.dot {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 7px;
  height: 7px;
  background-color: #28a745;
  border-radius: 50%;
  border: 1px solid #ffffff;
}

.eq-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 123, 255, 0.55);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 3px;
}

.eq-overlay span {
  width: 3px;
  background: #ffffff;
  border-radius: 1px;
  animation: eq-bounce 0.9s ease-in-out infinite;
}

.eq-overlay span:nth-child(1) { animation-delay: 0s; height: 6px; }
.eq-overlay span:nth-child(2) { animation-delay: 0.2s; height: 14px; }
.eq-overlay span:nth-child(3) { animation-delay: 0.4s; height: 9px; }

.info {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  gap: 1px;
}

.name-box, .singer-box {
  border-radius: 6px;
  padding: 2px 6px;
  font-size: 13px;
  /* background: #ffffff; */
  color: #212529;
}

.singer-box {
  color: #6c757d;
  font-size: 11px;
}

.song-btns {
  display: flex;
  gap: 6px;
}

.btn-sm {
  background: #1976D2;
  border: 1px solid #1976D2;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  color: #ffffff;
  width: 25px;
  height: 25px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.btn-sm:hover {
  background-color: #e9ecef;
  color: #007bff;
  transform: scale(1.08);
}

.more-btn {
  width: 100%;
  background: #f8f9fa;
  border: none;
  border-radius: 8px;
  padding: 10px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  color: #495057;
  transition: background-color 0.2s;
}

.more-btn:hover {
  background-color: #e9ecef;
}

.custom-action-popup {
  position: absolute;
  right: 0;
  bottom: 38px;
  width: 180px;
  background-color: #2d3126; 
  border: 1px solid #4a4f40;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  display: flex;
  flex-direction: column;
  padding: 4px;
  z-index: 100;
  animation: popup-in 0.15s ease-out;
}

@keyframes popup-in {
  from { opacity: 0; transform: translateY(6px); }
  to { opacity: 1; transform: translateY(0); }
}

.popup-item {
  background: transparent;
  border: none;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  color: #ffffff;
  padding: 10px 12px;
  text-align: left;
  font-size: 14px;
  font-weight: 700;
  font-family: inherit;
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  transition: background 0.2s;
}

.popup-item:last-child {
  border-bottom: none;
}

.popup-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.popup-icon {
  position: relative;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.plus-badge {
  position: absolute;
  bottom: -2px;
  right: -4px;
  font-size: 10px;
  font-weight: bold;
  color: #ffffff;
}

/* Grid for Playlists */
.grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 12px;
}

.card {
  border-radius: 12px;
}

.card-img-wrap {
  position: relative;
  height: 120px;
  aspect-ratio: 1 / 1;
  border-radius: 8px;
}

.card-play-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.25);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  opacity: 0;
  border-radius: 8px;
  
}

.card:hover .card-play-overlay {
  opacity: 1;
}

.card-desc {
  padding: 10px 12px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.card-title {
  font-size: 13px;
  font-weight: 600;
  color: #212529;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.badge {
  font-size: 11px;
  color: #ffffff;
  background: #1976D2;
  padding: 2px 6px;
  border-radius: 4px;
  width: fit-content;
}

.song-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 11px;
  color: #6c757d;
}


.singer {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}


.song-duration {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #495057;
  min-width: 38px;
  text-align: center;
  font-variant-numeric: tabular-nums;
  background-color: transparent;
  border: 1px solid #e2e8f0;
  padding: 3px 8px;
  border-radius: 12px;
  transition: background-color 0.2s, color 0.2s;
}

.song-duration .clock-icon {
  stroke: #6c757d;
  flex-shrink: 0;
  width: 12px;
  height: 12px;
}

.song-item:hover .song-duration {
  background-color: #e2e8f0;
  color: #212529;
}

.song-info {
  display: flex;
  flex-direction: column;
  line-height: 1.2;
}


.now-cover {
  width: 40px;
  height: 40px;
  border-radius: 6px;
  object-fit: cover;
  flex-shrink: 0;
}

.song-title {
  font-size: 14px;
  font-weight: 600;
  color: #333;
  max-width: 160px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.song-artist {
  font-size: 12px;
  color: #777;
  margin-top: 2px;
}

.btn.active {
  color:#007bff;
}

.play-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0,0,0,0.35);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  opacity: 0;
  transition: .2s ease;
}

.song-item:hover .play-overlay {
  opacity: 1;
}

.song-detail-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.back-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background-color: #ffffff;
  border: 1px solid #e1e4e8;
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  color: #333333;
  cursor: pointer;
  width: fit-content;
  transition: all 0.2s ease;
}

.back-btn:hover {
  background-color: #007bff;
  color: #ffffff;
  border-color: #007bff;
}

@media (max-width: 1024px) {
  .container {
    grid-template-columns: 200px 1fr; 
  }
  .sidebar-right {
    display: none;
  }
}
</style>