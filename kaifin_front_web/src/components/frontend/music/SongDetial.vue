<script setup>
import { ref, computed, watch, nextTick, onMounted, onUnmounted } from 'vue'
import { DotLottieVue } from '@lottiefiles/dotlottie-vue'
import CreateSticker from './CreateSticker.vue'

const props = defineProps({
  playlist: { type: Object, default: null },
  currentSong: { type: Object, default: null },
  isPlaying: { type: Boolean, default: false },
  currentTime: { type: Number, default: 0 },
  duration: { type: Number, default: 0 },
  volume: { type: Number, default: 0.6 },
  formatTime: { type: Function, required: true },
  progressPercent: { type: Number, default: 0 },
})

const emit = defineEmits(['play', 'toggle-play', 'seek', 'set-volume', 'back', 'ended', 'react', 'share', 'add-to-playlist', 'send-message', 'create-post', 'watch-later', 'report', 'attach-music', 'attach-video'])

const videoRef = ref(null)
const localCurrentSong = ref(props.currentSong)
const localIsPlaying = ref(false)

const showReactionPopup = ref(false)
const selectedReaction = ref(null)
const reactionWrapperRef = ref(null)
const likersPopupWrapperRef = ref(null)

// Like / Share counters
const likeCount = ref(128)
const shareCount = ref(42)
const hasLiked = ref(false)
const showCopiedToast = ref(false)

// "Liked by" popovers
const showMainLikersPopup = ref(false)
const activeCommentLikersId = ref(null)
const showStatsLikersPopup = ref(false)
const statsLikersPopupRef = ref(null)
const showStatsSharersPopup = ref(false)
const statsSharersPopupRef = ref(null)

// Share dropdown menu
const showShareDropdown = ref(false)

const shareMenuItems = [
  { id: 'share-now', label: 'Share now', icon: 'share' },
  { id: 'add-to-playlist', label: 'Add to playlist', icon: 'list' },
  { id: 'send-message', label: 'Send as a message', icon: 'message' },
  { id: 'copy-link', label: 'Copy link', icon: 'copy' },
  { id: 'create-post', label: 'Create post', icon: 'edit' },
  { id: 'divider-1', divider: true },
  { id: 'watch-later', label: 'Watch later', icon: 'clock' },
  { id: 'report', label: 'Report', icon: 'warning' },
]

function formatCount(n) {
  if (n >= 1000000) return (n / 1000000).toFixed(1).replace(/\.0$/, '') + 'M'
  if (n >= 1000) return (n / 1000).toFixed(1).replace(/\.0$/, '') + 'K'
  return n.toString()
}

// Mock pool of profiles used to populate "liked by" popovers
const samplePeople = [
  { name: 'Sokha Dara', avatar: 'https://images.unsplash.com/photo-1527980965255-d3b416303d12?w=100&auto=format&fit=crop&q=80' },
  { name: 'Chan Thy', avatar: 'https://images.unsplash.com/photo-1580489944761-15a19d654956?w=100&auto=format&fit=crop&q=80' },
  { name: 'Vimean Sovann', avatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100&auto=format&fit=crop&q=80' },
  { name: 'Bopha Chan', avatar: 'https://images.unsplash.com/photo-1544005313-94ddf0286df2?w=100&auto=format&fit=crop&q=80' },
  { name: 'Rithy Long', avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=100&auto=format&fit=crop&q=80' },
  { name: 'Sreymom Kim', avatar: 'https://images.unsplash.com/photo-1544723795-3fb6469f5b39?w=100&auto=format&fit=crop&q=80' },
  { name: 'Pisach Ly', avatar: 'https://images.unsplash.com/photo-1492562080023-ab3db95bfbce?w=100&auto=format&fit=crop&q=80' },
  { name: 'Chenda Ros', avatar: 'https://images.unsplash.com/photo-1531123897727-8f129e1688ce?w=100&auto=format&fit=crop&q=80' }
]

function getLikersPreview(count) {
  const shown = samplePeople.slice(0, Math.min(count, samplePeople.length))
  const remaining = count - shown.length
  return { shown, remaining: remaining > 0 ? remaining : 0 }
}

const mainLikersPreview = computed(() => getLikersPreview(likeCount.value))
const mainSharersPreview = computed(() => getLikersPreview(shareCount.value))

function closeAllLikersPopups() {
  showMainLikersPopup.value = false
  showStatsLikersPopup.value = false
  showStatsSharersPopup.value = false
  activeCommentLikersId.value = null
}

function toggleMainLikersPopup(event) {
  event.stopPropagation()
  if (likeCount.value <= 0) return
  const next = !showMainLikersPopup.value
  closeAllLikersPopups()
  showMainLikersPopup.value = next
}

function toggleStatsLikersPopup(event) {
  event.stopPropagation()
  if (likeCount.value <= 0) return
  const next = !showStatsLikersPopup.value
  closeAllLikersPopups()
  showStatsLikersPopup.value = next
}

function toggleStatsSharersPopup(event) {
  event.stopPropagation()
  if (shareCount.value <= 0) return
  const next = !showStatsSharersPopup.value
  closeAllLikersPopups()
  showStatsSharersPopup.value = next
}

function toggleCommentLikersPopup(comment, event) {
  event.stopPropagation()
  if (!comment.likeCount) return
  const next = activeCommentLikersId.value === comment.id ? null : comment.id
  closeAllLikersPopups()
  activeCommentLikersId.value = next
}

function getCommentLikersPreview(comment) {
  return getLikersPreview(comment.likeCount || 0)
}

const reactionList = ref([
  { id: 1, name: 'Tongue', src: '/src/assets/json_lot/bsong.json' },
  { id: 2, name: 'Sad/Cute', src: '/src/assets/json_lot/crazy_1.json' },
  { id: 3, name: 'Surprised', src: '/src/assets/json_lot/test.json' },
  { id: 4, name: 'Laugh', src: '/src/assets/json_lot/like.json' },
  { id: 5, name: 'Smile', src: '/src/assets/json_lot/cute_d_1.json' }
])

const defaultPlaylistSongs = [
  { id: 1, title: 'Golden Hour Boy', singer: 'JVKE Remix', duration: '03:25', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3', cover: 'https://images.unsplash.com/photo-1514525253161-7a46d19cd819?w=300&auto=format&fit=crop&q=80' },
  { id: 2, title: 'Midnight City Beats', singer: 'M83 Boy Vibes', duration: '04:03', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3', cover: 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?w=300&auto=format&fit=crop&q=80' },
  { id: 3, title: 'Electric Youth Anthem', singer: 'Synth Boy', duration: '03:45', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-3.mp3', cover: 'https://images.unsplash.com/photo-1470225620780-dba8ba36b745?w=300&auto=format&fit=crop&q=80' },
  { id: 4, title: 'Lost in Tokyo Dreams', singer: 'Lofi Boy', duration: '02:55', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-4.mp3', cover: 'https://images.unsplash.com/photo-1518609878373-06d740f60d8b?w=300&auto=format&fit=crop&q=80' },
  { id: 5, title: 'Summer High Acoustic', singer: 'Acoustic Boy', duration: '03:12', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-5.mp3', cover: 'https://images.unsplash.com/photo-1465847899084-d164df4dedc6?w=300&auto=format&fit=crop&q=80' },
  { id: 6, title: 'Neon Lights & Shadows', singer: 'Cyber Boy', duration: '03:50', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-6.mp3', cover: 'https://images.unsplash.com/photo-1508700115892-45ecd05ae2ad?w=300&auto=format&fit=crop&q=80' },
  { id: 7, title: 'Coffee Shop Melancholy', singer: 'Indie Boy', duration: '03:34', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-7.mp3', cover: 'https://images.unsplash.com/photo-1501386761578-eac5c94b800a?w=300&auto=format&fit=crop&q=80' },
  { id: 8, title: 'Starlight Alleyway', singer: 'Night Boy', duration: '04:15', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-8.mp3', cover: 'https://images.unsplash.com/photo-1509198397868-475647b2a1e5?w=300&auto=format&fit=crop&q=80' },
  { id: 9, title: 'Rhythm of the Streets', singer: 'Urban Boy', duration: '03:08', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-9.mp3', cover: 'https://images.unsplash.com/photo-1498038432885-c6f3f1b912ee?w=300&auto=format&fit=crop&q=80' },
  { id: 10, title: 'Sunset Boulevard Groove', singer: 'Retro Boy', duration: '03:40', fileUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-10.mp3', cover: 'https://images.unsplash.com/photo-1526478806334-5fd488fcaabc?w=300&auto=format&fit=crop&q=80' }
]

const activePlaylist = computed(() => {
  if (props.playlist && props.playlist.songs && props.playlist.songs.length > 0) {
    return props.playlist
  }
  return {
    name: '10 songs • Boy song',
    songsCount: defaultPlaylistSongs.length,
    songs: defaultPlaylistSongs
  }
})

watch(
  () => props.currentSong,
  (newSong) => {
    if (newSong) {
      localCurrentSong.value = newSong
    }
  }
)

const upNextList = computed(() => activePlaylist.value.songs)
const recommendedList = computed(() => activePlaylist.value.songs)

const handleClickOutside = (event) => {
  if (reactionWrapperRef.value && !reactionWrapperRef.value.contains(event.target)) {
    showReactionPopup.value = false
  }
  if (likersPopupWrapperRef.value && !likersPopupWrapperRef.value.contains(event.target)) {
    showMainLikersPopup.value = false
  }
  if (statsLikersPopupRef.value && !statsLikersPopupRef.value.contains(event.target)) {
    showStatsLikersPopup.value = false
  }
  if (statsSharersPopupRef.value && !statsSharersPopupRef.value.contains(event.target)) {
    showStatsSharersPopup.value = false
    showShareDropdown.value = false
  }
  if (attachMenuWrapperRef.value && !attachMenuWrapperRef.value.contains(event.target)) {
    showAttachMenu.value = false
  }
  if (emojiPickerWrapperRef.value && !emojiPickerWrapperRef.value.contains(event.target)) {
    showEmojiPicker.value = false
  }
  if (!event.target.closest('.comment-like-group')) {
    activeCommentLikersId.value = null
  }
}

onMounted(() => {
  if (!localCurrentSong.value && activePlaylist.value.songs?.length) {
    localCurrentSong.value = activePlaylist.value.songs[0]
  }
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

function playSong(song) {
  if (!song?.fileUrl) return

  localCurrentSong.value = song
  emit('play', song)

  nextTick(() => {
    if (videoRef.value) {
      videoRef.value.src = song.fileUrl
      videoRef.value.play().then(() => {
        localIsPlaying.value = true
      }).catch((err) => {
        localIsPlaying.value = false
      })
    }
  })
}

const togglePlay = () => {
  if (!videoRef.value) return

  if (!videoRef.value.src && localCurrentSong.value?.fileUrl) {
    videoRef.value.src = localCurrentSong.value.fileUrl
  }

  if (videoRef.value.paused) {
    videoRef.value.play().then(() => {
      localIsPlaying.value = true
      emit('toggle-play', true)
    })
  } else {
    videoRef.value.pause()
    localIsPlaying.value = false
    emit('toggle-play', false)
  }
}

const handleSeek = (event) => emit('seek', event)
const handleSetVolume = (event) => emit('set-volume', event)

function handleEnded() {
  localIsPlaying.value = false
  emit('toggle-play', false)
  emit('ended', localCurrentSong.value)

  if (videoRef.value) {
    videoRef.value.currentTime = 0
  }

  const list = upNextList.value
  if (!list.length) return

  const currentIndex = list.findIndex((s) => s.id === localCurrentSong.value?.id)
  const nextSong = currentIndex >= 0 ? list[currentIndex + 1] : null

  if (nextSong) {
    playSong(nextSong)
  }
}

const selectReaction = (reaction) => {
  const isSameReaction = selectedReaction.value?.id === reaction.id

  if (isSameReaction && hasLiked.value) {
    // clicking the same reaction again removes the like
    likeCount.value -= 1
    hasLiked.value = false
    selectedReaction.value = null
  } else {
    if (!hasLiked.value) {
      likeCount.value += 1
      hasLiked.value = true
    }
    selectedReaction.value = reaction
  }

  showReactionPopup.value = false
  emit('react', selectedReaction.value)
}

const copyShareLink = async () => {
  shareCount.value += 1
  emit('share', localCurrentSong.value)
  try {
    await navigator.clipboard.writeText(window.location.href)
  } catch (err) {
    console.error('Copy failed:', err)
  }
  showCopiedToast.value = true
  setTimeout(() => {
    showCopiedToast.value = false
  }, 1800)
}

const toggleShareDropdown = (event) => {
  event.stopPropagation()
  const next = !showShareDropdown.value
  closeAllLikersPopups()
  showShareDropdown.value = next
}

const handleShareOption = (item) => {
  if (item.divider) return
  showShareDropdown.value = false

  switch (item.id) {
    case 'share-now':
    case 'copy-link':
      copyShareLink()
      break
    case 'add-to-playlist':
      emit('add-to-playlist', localCurrentSong.value)
      break
    case 'send-message':
      emit('send-message', localCurrentSong.value)
      break
    case 'create-post':
      emit('create-post', localCurrentSong.value)
      break
    case 'watch-later':
      emit('watch-later', localCurrentSong.value)
      break
    case 'report':
      emit('report', localCurrentSong.value)
      break
  }
}

const currentUserAvatar = ref('https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?w=100&auto=format&fit=crop&q=80')

const comments = ref([
  {
    id: 1,
    user: 'Sokha Dara',
    avatar: 'https://images.unsplash.com/photo-1527980965255-d3b416303d12?w=100&auto=format&fit=crop&q=80',
    date: '2 hours ago',
    text: 'This boy band playlist hits different! Love every track.',
    audioUrl: null,
    imageUrl: null,
    isPlaying: false,
    progress: 0,
    likeCount: 6,
    liked: false,
    replies: [
      {
        id: 101,
        user: 'Chan Thy',
        avatar: 'https://images.unsplash.com/photo-1580489944761-15a19d654956?w=100&auto=format&fit=crop&q=80',
        date: '1 hour ago',
        text: 'Totally agree, track #1 is absolute gold!',
        audioUrl: null,
        imageUrl: null,
        isPlaying: false,
        progress: 0,
        likeCount: 1,
        liked: false,
        replies: []
      }
    ]
  },
  {
    id: 2,
    user: 'Chan Thy',
    avatar: 'https://images.unsplash.com/photo-1580489944761-15a19d654956?w=100&auto=format&fit=crop&q=80',
    date: '1 hour ago',
    text: null,
    audioUrl: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3',
    imageUrl: null,
    duration: '0:05',
    isPlaying: false,
    progress: 0,
    likeCount: 2,
    liked: false,
    replies: []
  },
  {
    id: 3,
    user: 'Vimean Sovann',
    avatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100&auto=format&fit=crop&q=80',
    date: '30 mins ago',
    text: 'Vibing to song #4 right now!',
    audioUrl: null,
    imageUrl: 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?w=600&auto=format&fit=crop&q=80',
    isPlaying: false,
    progress: 0,
    likeCount: 4,
    liked: false,
    tags: [
      { id: 't1', x: 32, y: 40, name: 'Chan Thy' }
    ],
    replies: []
  }
])

const newComment = ref('')
const fileInputRef = ref(null)
const commentInputRef = ref(null)

// @ attach menu (Tag a friend / Music / Video / Photo)
const showAttachMenu = ref(false)
const attachMenuWrapperRef = ref(null)

const toggleAttachMenu = (event) => {
  event.stopPropagation()
  showAttachMenu.value = !showAttachMenu.value
}

const handleAttachOption = (id) => {
  showAttachMenu.value = false

  switch (id) {
    case 'tag-friend':
      newComment.value = newComment.value && !newComment.value.endsWith(' ')
        ? newComment.value + ' @'
        : newComment.value + '@'
      nextTick(() => {
        if (commentInputRef.value) commentInputRef.value.focus()
      })
      break
    case 'music':
      emit('attach-music')
      break
    case 'video':
      emit('attach-video')
      break
    case 'photo':
      triggerFileUpload()
      break
  }
}

const activeReplyCommentId = ref(null)
const replyInputText = ref('')

const addTextComment = () => {
  if (!newComment.value.trim()) return
  comments.value.push({
    id: Date.now(),
    user: 'You',
    avatar: currentUserAvatar.value,
    date: 'Just now',
    text: newComment.value.trim(),
    audioUrl: null,
    imageUrl: null,
    isPlaying: false,
    progress: 0,
    likeCount: 0,
    liked: false,
    replies: []
  })
  newComment.value = ''
}

const toggleCommentLike = (comment) => {
  if (comment.liked) {
    comment.likeCount = Math.max(0, (comment.likeCount || 0) - 1)
    comment.liked = false
  } else {
    comment.likeCount = (comment.likeCount || 0) + 1
    comment.liked = true
  }
}

const toggleReplyBox = (commentId) => {
  if (activeReplyCommentId.value === commentId) {
    activeReplyCommentId.value = null
    replyInputText.value = ''
  } else {
    activeReplyCommentId.value = commentId
    replyInputText.value = ''
  }
}

const submitSubComment = (parentComment) => {
  if (!replyInputText.value.trim()) return
  if (!parentComment.replies) {
    parentComment.replies = []
  }
  parentComment.replies.push({
    id: Date.now(),
    user: 'You',
    avatar: currentUserAvatar.value,
    date: 'Just now',
    text: replyInputText.value.trim(),
    audioUrl: null,
    imageUrl: null,
    isPlaying: false,
    progress: 0,
    likeCount: 0,
    liked: false,
    replies: []
  })
  replyInputText.value = ''
  activeReplyCommentId.value = null
}

const triggerFileUpload = () => {
  if (fileInputRef.value) {
    fileInputRef.value.click()
  }
}

// Photo tagging
const taggingCommentId = ref(null)
const pendingTagPosition = ref(null)
const tagNameInput = ref('')
const tagInputRef = ref(null)

const toggleTaggingMode = (comment) => {
  if (taggingCommentId.value === comment.id) {
    taggingCommentId.value = null
  } else {
    taggingCommentId.value = comment.id
  }
  pendingTagPosition.value = null
  tagNameInput.value = ''
}

const handlePhotoClick = (comment, event) => {
  if (taggingCommentId.value !== comment.id) return
  const rect = event.currentTarget.getBoundingClientRect()
  const x = ((event.clientX - rect.left) / rect.width) * 100
  const y = ((event.clientY - rect.top) / rect.height) * 100
  pendingTagPosition.value = { x, y }
  tagNameInput.value = ''
  nextTick(() => {
    if (tagInputRef.value) tagInputRef.value.focus()
  })
}

const confirmTag = (comment) => {
  if (!tagNameInput.value.trim() || !pendingTagPosition.value) return
  if (!comment.tags) comment.tags = []
  comment.tags.push({
    id: 'tag-' + Date.now(),
    x: pendingTagPosition.value.x,
    y: pendingTagPosition.value.y,
    name: tagNameInput.value.trim()
  })
  pendingTagPosition.value = null
  tagNameInput.value = ''
}

const cancelTag = () => {
  pendingTagPosition.value = null
  tagNameInput.value = ''
}

const removeTag = (comment, tagId, event) => {
  if (event) event.stopPropagation()
  comment.tags = (comment.tags || []).filter(t => t.id !== tagId)
}

const handleImageUpload = (event) => {
  const file = event.target.files[0]
  if (!file) return

  const reader = new FileReader()
  reader.onload = (e) => {
    comments.value.push({
      id: Date.now(),
      user: 'You',
      avatar: currentUserAvatar.value,
      date: 'Just now',
      text: newComment.value.trim() ? newComment.value.trim() : null,
      audioUrl: null,
      imageUrl: e.target.result,
      isPlaying: false,
      progress: 0,
      likeCount: 0,
      liked: false,
      tags: [],
      replies: []
    })
    newComment.value = ''
  }
  reader.readAsDataURL(file)
  event.target.value = ''
}

const activeAudioId = ref(null)
const audioPlayerRef = ref(null)

const toggleVoicePlay = (comment) => {
  if (!comment.audioUrl) return

  if (activeAudioId.value === comment.id && audioPlayerRef.value) {
    if (comment.isPlaying) {
      audioPlayerRef.value.pause()
      comment.isPlaying = false
    } else {
      audioPlayerRef.value.play()
      comment.isPlaying = true
    }
  } else {
    const stopAll = (list) => {
      list.forEach(c => {
        c.isPlaying = false
        if (c.replies && c.replies.length) stopAll(c.replies)
      })
    }
    stopAll(comments.value)

    activeAudioId.value = comment.id
    comment.isPlaying = true

    nextTick(() => {
      if (audioPlayerRef.value) {
        audioPlayerRef.value.play()
      }
    })
  }
}

const onAudioTimeUpdate = (comment) => {
  if (audioPlayerRef.value && activeAudioId.value === comment.id) {
    const duration = audioPlayerRef.value.duration || 5
    comment.progress = (audioPlayerRef.value.currentTime / duration) * 100
  }
}

const onAudioEnded = (comment) => {
  comment.isPlaying = false
  comment.progress = 0
}

const isRecording = ref(false)
const recordingTime = ref(5)
const recordingTimer = ref(null)
const mediaRecorder = ref(null)
const audioChunks = ref([])

const startVoiceRecording = async () => {
  try {
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
    audioChunks.value = []
    mediaRecorder.value = new MediaRecorder(stream)

    mediaRecorder.value.ondataavailable = (event) => {
      if (event.data.size > 0) {
        audioChunks.value.push(event.data)
      }
    }

    mediaRecorder.value.onstop = () => {
      const audioBlob = new Blob(audioChunks.value, { type: 'audio/webm' })
      const audioUrl = URL.createObjectURL(audioBlob)

      comments.value.push({
        id: Date.now(),
        user: 'You',
        avatar: currentUserAvatar.value,
        date: 'Just now',
        audioUrl: audioUrl,
        imageUrl: null,
        duration: '0:05',
        isPlaying: false,
        progress: 0,
        likeCount: 0,
        liked: false,
        replies: []
      })

      stream.getTracks().forEach(track => track.stop())
    }

    mediaRecorder.value.start()
    isRecording.value = true
    recordingTime.value = 5

    recordingTimer.value = setInterval(() => {
      recordingTime.value -= 1
      if (recordingTime.value <= 0) {
        stopVoiceRecording()
      }
    }, 1000)

  } catch (err) {
    console.error('Error accessing microphone:', err)
    alert('Microphone access is required to record voice comments.')
  }
}

const stopVoiceRecording = () => {
  if (mediaRecorder.value && isRecording.value) {
    mediaRecorder.value.stop()
    isRecording.value = false
    if (recordingTimer.value) {
      clearInterval(recordingTimer.value)
      recordingTimer.value = null
    }
  }
}


// Emoji / Sticker Picker State
const showEmojiPicker = ref(false)
const emojiPickerWrapperRef = ref(null)
const activePickerTab = ref('stickers') // 'stickers' | 'emoji'
const activeStickerCategory = ref('all') // 'all' | 'animated' | 'mine'

// Each animated sticker belongs to a "pack" (a bundle a user downloads/owns together).
// Packs are only used to group items inside the "Animated" tab.
const stickerPool = ref([
  { id: 1, src: 'src/assets/animate/succ.svg', animated: false, mine: false, pack: null },
  { id: 2, src: 'src/assets/animate/cat.svg', animated: false, mine: false, pack: null },
  { id: 3, src: 'src/assets/animate/congrate.svg', animated: true, mine: false, pack: 'Congratulation' },
  { id: 4, src: 'src/assets/animate/cool.svg', animated: false, mine: false, pack: null },
  { id: 5, src: 'src/assets/animate/w10.svg', animated: true, mine: true, pack: 'Welcome (10)' },
  { id: 6, src: 'src/assets/animate/Emojis.svg', animated: false, mine: false, pack: null },
  { id: 7, src: 'src/assets/animate/w9.svg', animated: true, mine: false, pack: 'Welcome (10)' },
  { id: 8, src: 'src/assets/animate/hearth.svg', animated: true, mine: false, pack: 'Congratulation' },
  { id: 9, src: 'src/assets/animate/w8.svg', animated: true, mine: false, pack: 'Welcome (10)' },
  { id: 10, src: 'src/assets/animate/w7.svg', animated: true, mine: false, pack: 'Welcome (10)' },
  { id: 11, src: 'src/assets/animate/w6.svg', animated: true, mine: true, pack: 'Welcome (10)' },
  { id: 12, src: 'src/assets/animate/w5.svg', animated: true, mine: false, pack: 'Welcome (10)' },
  { id: 13, src: 'src/assets/animate/w4.svg', animated: true, mine: false, pack: 'Welcome (10)' },
  { id: 14, src: 'src/assets/animate/Thinking.svg', animated: true, mine: false, pack: 'Thinkiing (5)' },
  { id: 15, src: 'src/assets/animate/welcome_3.svg', animated: true, mine: false, pack: 'Thinkiing (5)' },
  { id: 16, src: 'src/assets/animate/w3.svg', animated: true, mine: false, pack: 'Welcome (10)' },
  { id: 17, src: 'src/assets/animate/w2.svg', animated: true, mine: false, pack: 'Welcome (10)' },
  { id: 18, src: 'src/assets/animate/w1.svg', animated: true, mine: true, pack: 'Welcome (10)' },
  { id: 19, src: 'src/assets/animate/w11.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 20, src: 'src/assets/animate/w12.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 21, src: 'src/assets/animate/w13.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 22, src: 'src/assets/animate/w22.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 23, src: 'src/assets/animate/w14.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 24, src: 'src/assets/animate/w20.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 25, src: 'src/assets/animate/w17.svg', animated: true, mine: true, pack: 'Thinkiing (5)' },
  { id: 26, src: 'src/assets/animate/w21.svg', animated: true, mine: true, pack: 'Thinkiing (5)' },
  { id: 27, src: 'src/assets/animate/w22.svg', animated: true, mine: true, pack: 'Thinkiing (5)' },
  { id: 28, src: 'src/assets/animate/w23.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 29, src: 'src/assets/animate/w24.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 30, src: 'src/assets/animate/w25.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 31, src: 'src/assets/animate/w25.svg', animated: true, mine: true, pack: 'Wishing Prays (20)' },
  { id: 32, src: 'src/assets/animate/g1.svg', animated: true, mine: true, pack: 'Gosh (20)' },
  { id: 33, src: 'src/assets/animate/g2.svg', animated: true, mine: true, pack: 'Gosh (20)' },
  { id: 34, src: 'src/assets/animate/g3.svg', animated: true, mine: true, pack: 'Gosh (20)' },
  { id: 35, src: 'src/assets/animate/g4.svg', animated: true, mine: true, pack: 'Gosh (20)' },
  { id: 35, src: 'src/assets/animate/g5.svg', animated: true, mine: true, pack: 'Gosh (20)' },
  { id: 35, src: 'src/assets/animate/g6.svg', animated: true, mine: true, pack: 'Gosh (20)' },
  { id: 35, src: 'src/assets/animate/g1.svg', animated: true, mine: true, pack: 'Gosh (20)' },
])

const filteredStickers = computed(() => {
  if (activeStickerCategory.value === 'mine') return stickerPool.value.filter(s => s.mine)
  return stickerPool.value
})

// Groups animated stickers by their "pack" name, so the Animated tab
// renders them as separate labeled sets instead of one flat grid.
const animatedStickerGroups = computed(() => {
  const animated = stickerPool.value.filter(s => s.animated)
  const groups = {}
  const order = []
  animated.forEach(s => {
    const key = s.pack || 'ផ្សេងៗ'
    if (!groups[key]) {
      groups[key] = []
      order.push(key)
    }
    groups[key].push(s)
  })
  return order.map(name => ({ name, items: groups[name] }))
})

const toggleEmojiPicker = (event) => {
  event.stopPropagation()
  showEmojiPicker.value = !showEmojiPicker.value
}

const selectSticker = (sticker) => {
  comments.value.push({
    id: Date.now(),
    user: 'You',
    avatar: currentUserAvatar.value,
    date: 'Just now',
    text: null,
    audioUrl: null,
    imageUrl: sticker.src,
    isPlaying: false,
    progress: 0,
    likeCount: 0,
    liked: false,
    tags: [],
    replies: []
  })
  showEmojiPicker.value = false
}

const emojiPool = ref(['😀','😁','😂','🤣','😊','😍','😘','😜','🤩','😎','🥳','😇','🙂','😉','😋','😏','😢','😭','😡','🤔','👍','👏','🙏','❤️','🔥','🎉','✨','💯'])

const insertEmoji = (emoji) => {
  newComment.value += emoji
  showEmojiPicker.value = false
  nextTick(() => {
    if (commentInputRef.value) commentInputRef.value.focus()
  })
}


// អថេរគ្រប់គ្រងការបង្ហាញផ្ទាំង Create Sticker (លំនាំដើមគឺបិទ false)
const isStickerCreatorOpen = ref(false)

// Function សម្រាប់បើកផ្ទាំង ពេលគេចុចប៊ូតុង
const openStickerCreator = () => {
  isStickerCreatorOpen.value = true
  
  // បើសិនជាអ្នកមានផ្ទាំង Emoji ផ្សេងទៀតកំពុងបើក គួរតែបិទវាចោលសិន
  // showEmojiPicker.value = false 
}

// Function សម្រាប់បិទផ្ទាំងវិញ ពេលគេចុច Cancel ឬចុចខាងក្រៅ
const closeStickerCreator = () => {
  isStickerCreatorOpen.value = false
}

const handleToolChanged = (toolName) => {
  // អ្នកអាចសរសេរកូដបើក Modal ឬមុខងារបន្ថែមទីនេះ អាស្រ័យលើ toolName (ឧ. text, crop, effect...)
}
</script>

<template>
  <div class="wrap-player">
    <div class="player-container">

    <!-- Top Action Bar -->
    <div class="player-top-bar">
      <div class="badge-video translucent-badge">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 18V5l12-2v13"></path><circle cx="6" cy="18" r="3"></circle><circle cx="18" cy="16" r="3"></circle></svg>
        <span>10 songs • Boy song</span>
      </div>
      <div class="top-actions">
        <div class="reaction-wrapper" ref="reactionWrapperRef">
          <div class="like-btn-group" :class="{ liked: hasLiked }">
            <button class="btn-action-inner" @click="showReactionPopup = !showReactionPopup">
              <div class="svg-bg">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path></svg>
              </div>
              {{ selectedReaction ? selectedReaction.name : 'Like' }}
            </button>

            <div class="likers-popup-wrapper" ref="likersPopupWrapperRef">
              <button
                class="count-pill-btn"
                :class="{ disabled: likeCount === 0 }"
                @click="toggleMainLikersPopup"
                title="See who liked this"
              >
                {{ formatCount(likeCount) }}
              </button>

              <div v-if="showMainLikersPopup" class="likers-popup">
                <div class="likers-popup-title">Liked by</div>
                <div class="likers-list">
                  <div v-for="(person, idx) in mainLikersPreview.shown" :key="idx" class="likers-list-item">
                    <img :src="person.avatar" alt="" class="likers-avatar" />
                    <span class="likers-name">{{ person.name }}</span>
                  </div>
                </div>
                <div v-if="mainLikersPreview.remaining > 0" class="likers-more">
                  and {{ formatCount(mainLikersPreview.remaining) }} others
                </div>
              </div>
            </div>
          </div>

          <div v-if="showReactionPopup" class="reaction-popup-box">
            <div
              v-for="react in reactionList"
              :key="react.id"
              class="reaction-icon-item"
              @click="selectReaction(react)"
            >
              <DotLottieVue
                style="height: 32px; width: 32px;"
                autoplay
                loop
                :src="react.src"
              />
            </div>
          </div>
        </div>

        <div class="share-wrapper" ref="statsSharersPopupRef">
          <div class="share-btn-group">
            <button class="btn-action-inner" @click="toggleShareDropdown">
              <div class="svg-bg">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none"
                     stroke="currentColor" stroke-width="2" stroke-linecap="round"
                     stroke-linejoin="round">
                  <circle cx="18" cy="5" r="3"/>
                  <circle cx="6" cy="12" r="3"/>
                  <circle cx="18" cy="19" r="3"/>
                  <line x1="8.6" y1="10.7" x2="15.4" y2="6.3"/>
                  <line x1="8.6" y1="13.3" x2="15.4" y2="17.7"/>
                </svg>
              </div>
              Share
            </button>

            <button
              class="count-pill-btn"
              :class="{ disabled: shareCount === 0 }"
              @click.stop="toggleStatsSharersPopup"
            >
              {{ formatCount(shareCount) }}
            </button>
          </div>

          <Transition name="fade-pop">
            <div v-if="showCopiedToast" class="copied-toast">
              Link copied!
            </div>
          </Transition>

          <Transition name="fade-pop">
            <div v-if="showShareDropdown" class="share-dropdown-menu">
              <template v-for="item in shareMenuItems" :key="item.id">
                <div v-if="item.divider" class="share-dropdown-divider"></div>
                <button
                  v-else
                  class="share-dropdown-item"
                  @click="handleShareOption(item)"
                >
                  <span class="share-dropdown-icon">
                    <svg v-if="item.icon === 'share'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
                    <svg v-else-if="item.icon === 'list'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><line x1="8" y1="6" x2="21" y2="6"></line><line x1="8" y1="12" x2="21" y2="12"></line><line x1="8" y1="18" x2="21" y2="18"></line><circle cx="3.5" cy="6" r="1.2" fill="currentColor" stroke="none"></circle><circle cx="3.5" cy="12" r="1.2" fill="currentColor" stroke="none"></circle><circle cx="3.5" cy="18" r="1.2" fill="currentColor" stroke="none"></circle></svg>
                    <svg v-else-if="item.icon === 'message'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-8.5 8.4H12A8.5 8.5 0 1 1 21 11.5z"></path><circle cx="8.5" cy="11.5" r="0.9" fill="currentColor" stroke="none"></circle><circle cx="12" cy="11.5" r="0.9" fill="currentColor" stroke="none"></circle><circle cx="15.5" cy="11.5" r="0.9" fill="currentColor" stroke="none"></circle></svg>
                    <svg v-else-if="item.icon === 'copy'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="12" height="12" rx="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>
                    <svg v-else-if="item.icon === 'edit'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path><path d="M18.5 2.5a2.12 2.12 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path></svg>
                    <svg v-else-if="item.icon === 'clock'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                    <svg v-else-if="item.icon === 'warning'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M10.29 3.86 1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path><line x1="12" y1="9" x2="12" y2="13"></line><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
                  </span>
                  <span class="share-dropdown-label">{{ item.label }}</span>
                </button>
              </template>
            </div>
          </Transition>

          <div
            v-if="showStatsSharersPopup"
            class="likers-popup"
          >
            <div class="likers-popup-title">
              Shared by
            </div>

            <div class="likers-list">
              <div
                v-for="(person,index) in mainSharersPreview.shown"
                :key="index"
                class="likers-list-item"
              >
                <img
                  :src="person.avatar"
                  class="likers-avatar"
                />
                <span class="likers-name">
                  {{ person.name }}
                </span>
              </div>
            </div>

            <div
              v-if="mainSharersPreview.remaining"
              class="likers-more"
            >
              and {{ formatCount(mainSharersPreview.remaining) }} others
            </div>
          </div>
        </div>

      </div>
    </div>

    <!-- MAIN MEDIA ROW -->
    <div class="media-row">
      <!-- LEFT: Video Screen -->
      <div class="video-screen">
        <video
          ref="videoRef"
          :poster="localCurrentSong?.cover"
          class="main-video"
          @play="localIsPlaying = true"
          @pause="localIsPlaying = false"
          @ended="handleEnded"
        ></video>

        <button v-if="!localIsPlaying" class="big-play-btn" @click="togglePlay">
          <svg width="28" height="28" viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
        </button>

        <div class="player-controls">
          <button class="ctrl-btn" title="Play/Pause" @click="togglePlay">
            <svg v-if="!localIsPlaying" width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
            <svg v-else width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><rect x="6" y="4" width="4" height="16"></rect><rect x="14" y="4" width="4" height="16"></rect></svg>
          </button>

          <span class="time-text">{{ formatTime ? formatTime(currentTime) : '0:00 / 0:12' }}</span>

          <div class="progress-track" @click="handleSeek">
            <div class="progress-fill" :style="{ width: progressPercent + '%' }"></div>
            <div class="progress-thumb" :style="{ left: progressPercent + '%' }"></div>
          </div>

          <div class="center-controls">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon><path d="M19.07 4.93a10 10 0 0 1 0 14.14M15.54 8.46a5 5 0 0 1 0 7.07"></path></svg>
            <div class="volume-slider" @click="handleSetVolume">
              <div class="vol-fill" :style="{ width: (volume * 100) + '%' }"></div>
            </div>
          </div>
        </div>
      </div>

      <!-- RIGHT: Attached Playing Next Sidebar -->
      <div class="playing-next-attached">
        <div class="sidebar-title light-text">Playing Next (10 Songs)</div>
        <div class="song-list-scroll">
          <div
            v-for="item in upNextList"
            :key="item.id"
            class="song-card dark-card"
            :class="{ active: localCurrentSong?.id === item.id }"
            @click="playSong(item)"
          >
            <div class="song-thumb">
              <img v-if="item.cover" :src="item.cover" alt="" />
              <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#64748B" stroke-width="2"><path d="M9 18V5l12-2v13"></path><circle cx="6" cy="18" r="3"></circle><circle cx="18" cy="16" r="3"></circle></svg>
            </div>
            <div class="song-info">
              <div class="s-name light-text">{{ item.title }}</div>
              <div class="s-singer dim-text">{{ item.singer }}</div>
            </div>
            <div class="s-duration dim-text">
             {{ item.duration || '03:30' }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- LOWER SECTION -->
    <div class="bottom-layout">
      <div class="main-details">
        <div class="details-box">
          <div class="description-header">
            <h3>{{ localCurrentSong?.title || activePlaylist.name }}</h3>
            <div class="stats-row">
              <span class="stat-badge song-icon-badge translucent-badge-sm">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 18V5l12-2v13"></path><circle cx="6" cy="18" r="3"></circle><circle cx="18" cy="16" r="3"></circle></svg>
                10 songs
              </span>

              <span class="stats-likers-wrapper" ref="statsLikersPopupRef">
                <button class="stat-badge stat-badge-clickable" @click="toggleStatsLikersPopup">
                  {{ formatCount(likeCount) }} likes
                </button>
                <div v-if="showStatsLikersPopup" class="likers-popup stats-likers-popup">
                  <div class="likers-popup-title">Liked by</div>
                  <div class="likers-list">
                    <div v-for="(person, idx) in mainLikersPreview.shown" :key="idx" class="likers-list-item">
                      <img :src="person.avatar" alt="" class="likers-avatar" />
                      <span class="likers-name">{{ person.name }}</span>
                    </div>
                  </div>
                  <div v-if="mainLikersPreview.remaining > 0" class="likers-more">
                    and {{ formatCount(mainLikersPreview.remaining) }} others
                  </div>
                </div>
              </span>
            </div>
          </div>
          <p class="desc-text">Featuring top 10 handpicked tracks in the Boy Song collection.</p>

          <div class="channel-bar">
            <div class="user-info">
              <img :src="localCurrentSong?.cover || 'https://via.placeholder.com/40'" alt="Avatar" class="avatar-img" />
              <div>
                <div class="username">{{ localCurrentSong?.singer || 'Boy Song Artist' }}</div>
                <div class="subscribers">
                  <svg class="inline-song-icon" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 18V5l12-2v13"></path><circle cx="6" cy="18" r="3"></circle><circle cx="18" cy="16" r="3"></circle></svg>
                  10 songs in playlist
                </div>
              </div>
            </div>

            <button class="subscribe-btn">Subscribe</button>
          </div>
        </div>

        <div class="comments-section">
          <div class="comments-header">Comments</div>

          <!-- COMMENTS LIST -->
          <div class="comments-list">
            <div v-for="c in comments" :key="c.id" class="comment-branch-node">

              <!-- Main Comment Item -->
              <div class="comment-item">
                <img :src="c.avatar" alt="Avatar" class="avatar-sm" />

                <div class="comment-content-wrapper">

                  <div class="comment-user-row">
                    <div class="user-date-group">
                      <span class="comment-user">{{ c.user }}</span>
                      <span class="comment-date badge-style translucent-badge-sm">
                        <svg class="clock-icon" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                        {{ c.date }}
                      </span>
                    </div>
                  </div>

                  <div class="comment-body">
                    <div v-if="c.text" class="comment-text">{{ c.text }}</div>

                    <div v-if="c.imageUrl" class="comment-photo-container">
                      <img
                        :src="c.imageUrl"
                        alt="Comment Photo"
                        class="comment-attached-photo"
                        :class="{ 'tagging-active': taggingCommentId === c.id }"
                        @click="handlePhotoClick(c, $event)"
                      />

                      <button
                        class="tag-toggle-btn"
                        :class="{ active: taggingCommentId === c.id }"
                        @click.stop="toggleTaggingMode(c)"
                        title="Tag people"
                      >
                        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.59 13.41 11 22 2 13V2h11l7.59 7.59a2 2 0 0 1 0 2.82z"></path><line x1="7" y1="7" x2="7.01" y2="7"></line></svg>
                        <span>{{ taggingCommentId === c.id ? 'Tagging…' : 'Tag' }}</span>
                      </button>

                      <div
                        v-for="tag in (c.tags || [])"
                        :key="tag.id"
                        class="photo-tag-marker"
                        :style="{ left: tag.x + '%', top: tag.y + '%' }"
                      >
                        <span class="photo-tag-dot"></span>
                        <span class="photo-tag-label">
                          {{ tag.name }}
                          <button
                            v-if="taggingCommentId === c.id"
                            class="photo-tag-remove"
                            @click.stop="removeTag(c, tag.id, $event)"
                          >×</button>
                        </span>
                      </div>

                      <div
                        v-if="pendingTagPosition && taggingCommentId === c.id"
                        class="photo-tag-input-box"
                        :style="{ left: pendingTagPosition.x + '%', top: pendingTagPosition.y + '%' }"
                        @click.stop
                      >
                        <input
                          ref="tagInputRef"
                          type="text"
                          v-model="tagNameInput"
                          placeholder="Tag name..."
                          @keyup.enter="confirmTag(c)"
                          @keyup.esc="cancelTag"
                        />
                        <button class="photo-tag-confirm" @click="confirmTag(c)" title="Add tag">✓</button>
                        <button class="photo-tag-cancel" @click="cancelTag" title="Cancel">×</button>
                      </div>
                    </div>

                    <div v-if="c.audioUrl" class="voice-comment-bubble">
                      <button class="voice-play-bubble-btn" @click="toggleVoicePlay(c)">
                        <svg v-if="!c.isPlaying" width="14" height="14" viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
                        <svg v-else width="14" height="14" viewBox="0 0 24 24" fill="currentColor"><rect x="6" y="4" width="4" height="16"></rect><rect x="14" y="4" width="4" height="16"></rect></svg>
                      </button>

                      <div class="voice-waveform-preview">
                        <div class="wave-bars">
                          <span class="bar" :class="{ active: c.isPlaying }" style="height: 12px"></span>
                          <span class="bar" :class="{ active: c.isPlaying }" style="height: 20px"></span>
                          <span class="bar" :class="{ active: c.isPlaying }" style="height: 16px"></span>
                          <span class="bar" :class="{ active: c.isPlaying }" style="height: 24px"></span>
                          <span class="bar" :class="{ active: c.isPlaying }" style="height: 14px"></span>
                          <span class="bar" :class="{ active: c.isPlaying }" style="height: 22px"></span>
                          <span class="bar" :class="{ active: c.isPlaying }" style="height: 10px"></span>
                        </div>
                        <div class="voice-progress-track">
                          <div class="voice-progress-fill" :style="{ width: c.progress + '%' }"></div>
                        </div>
                      </div>

                      <div class="voice-pill-tag">
                        <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"></path><path d="M19 10v1a7 7 0 0 1-14 0v-1"></path></svg>
                        <span>{{ c.duration || '0:05' }}</span>
                      </div>

                      <audio
                        ref="audioPlayerRef"
                        v-if="activeAudioId === c.id"
                        :src="c.audioUrl"
                        @timeupdate="onAudioTimeUpdate(c)"
                        @ended="onAudioEnded(c)"
                        style="display: none;"
                      ></audio>
                    </div>
                  </div>

                </div>
              </div>

              <!-- Reply / Like Action Row -->
              <div class="comment-actions-row outside-box-actions">
                <div class="comment-like-group">
                  <button class="comment-like-icon-btn" :class="{ liked: c.liked }" @click="toggleCommentLike(c)">
                    <svg width="13" height="13" viewBox="0 0 24 24" :fill="c.liked ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2.5"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path></svg>
                    <span>Like</span>
                  </button>
                  <button
                    v-if="c.likeCount > 0"
                    class="comment-like-count-btn"
                    @click="toggleCommentLikersPopup(c, $event)"
                    title="See who liked this"
                  >
                    {{ c.likeCount }}
                  </button>

                  <div v-if="activeCommentLikersId === c.id" class="likers-popup comment-likers-popup">
                    <div class="likers-popup-title">Liked by</div>
                    <div class="likers-list">
                      <div v-for="(person, idx) in getCommentLikersPreview(c).shown" :key="idx" class="likers-list-item">
                        <img :src="person.avatar" alt="" class="likers-avatar" />
                        <span class="likers-name">{{ person.name }}</span>
                      </div>
                    </div>
                    <div v-if="getCommentLikersPreview(c).remaining > 0" class="likers-more">
                      and {{ formatCount(getCommentLikersPreview(c).remaining) }} others
                    </div>
                  </div>
                </div>

                <button class="reply-toggle-btn" @click="toggleReplyBox(c.id)">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M3 10h10a4 4 0 0 1 4 4v2"></path><polyline points="7 6 3 10 7 14"></polyline></svg>
                  <span>Reply</span>
                </button>
              </div>

              <!-- TREE HIERARCHY SUB-COMMENTS LIST -->
              <div v-if="c.replies && c.replies.length > 0" class="sub-comments-tree-wrapper outside-box-tree">
                <div class="tree-branch-line"></div>
                <div class="sub-comments-list">
                  <div v-for="sub in c.replies" :key="sub.id" class="comment-item sub-comment-item">
                    <img :src="sub.avatar" alt="Avatar" class="avatar-sm" />

                    <div class="comment-content-wrapper">
                      <div class="comment-user-row">
                        <div class="user-date-group">
                          <span class="comment-user">{{ sub.user }}</span>
                          <span class="comment-date badge-style translucent-badge-sm">
                            <svg class="clock-icon" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                            {{ sub.date }}
                          </span>
                        </div>
                      </div>

                      <div class="comment-body sub-comment-body">
                        <div v-if="sub.text" class="comment-text">{{ sub.text }}</div>

                        <div v-if="sub.imageUrl" class="comment-photo-container">
                          <img :src="sub.imageUrl" alt="Sub Comment Photo" class="comment-attached-photo" />
                        </div>
                      </div>

                      <div class="comment-actions-row outside-box-actions">
                        <div class="comment-like-group">
                          <button class="comment-like-icon-btn" :class="{ liked: sub.liked }" @click="toggleCommentLike(sub)">
                            <svg width="12" height="12" viewBox="0 0 24 24" :fill="sub.liked ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2.5"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path></svg>
                            <span>Like</span>
                          </button>
                          <button
                            v-if="sub.likeCount > 0"
                            class="comment-like-count-btn"
                            @click="toggleCommentLikersPopup(sub, $event)"
                            title="See who liked this"
                          >
                            {{ sub.likeCount }}
                          </button>

                          <div v-if="activeCommentLikersId === sub.id" class="likers-popup comment-likers-popup">
                            <div class="likers-popup-title">Liked by</div>
                            <div class="likers-list">
                              <div v-for="(person, idx) in getCommentLikersPreview(sub).shown" :key="idx" class="likers-list-item">
                                <img :src="person.avatar" alt="" class="likers-avatar" />
                                <span class="likers-name">{{ person.name }}</span>
                              </div>
                            </div>
                            <div v-if="getCommentLikersPreview(sub).remaining > 0" class="likers-more">
                              and {{ formatCount(getCommentLikersPreview(sub).remaining) }} others
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>

                  </div>
                </div>
              </div>

              <!-- Inline Reply Input Box -->
              <div v-if="activeReplyCommentId === c.id" class="sub-comment-input-box outside-box-reply animate-pop">
                <img :src="currentUserAvatar" alt="Avatar" class="avatar-sm-xs" />
                <div class="comment-input-container">
                  <input
                    type="text"
                    v-model="replyInputText"
                    @keyup.enter="submitSubComment(c)"
                    placeholder="Write a reply..."
                    autofocus
                  />
                  <button
                    v-if="replyInputText.trim()"
                    class="send-action-btn animate-pop"
                    @click="submitSubComment(c)"
                    title="Send reply"
                  >
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
                  </button>
                </div>
              </div>

            </div>
          </div>

          <!-- Main Comment Input Box -->
          <div class="add-comment-box">
            <img :src="currentUserAvatar" alt="Avatar" class="avatar-sm" />

            <div class="emoji-picker-wrapper" ref="emojiPickerWrapperRef">
              <button class="badge-input-style emoji-toggle-btn" @click.stop="toggleEmojiPicker" title="Insert emoji or sticker">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="10"></circle>
                  <path d="M8 14s1.5 2 4 2 4-2 4-2"></path>
                  <line x1="9" y1="9" x2="9.01" y2="9"></line>
                  <line x1="15" y1="9" x2="15.01" y2="9"></line>
                </svg>
              </button>

              <!-- Emoji / Sticker Popup Box -->
              <Transition name="fade-pop">
                <div v-if="showEmojiPicker" class="emoji-picker-dropdown" @click.stop>
                  <div class="emoji-picker-header">
                    <button
                      class="tab-icon-btn"
                      :class="{ active: activePickerTab === 'stickers' }"
                      title="Stickers"
                      @click="activePickerTab = 'stickers'"
                    >
                      <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M11 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6"></path><path d="M17 3.5a2.12 2.12 0 0 1 3 3L11 15l-4 1 1-4 9-8.5z"></path></svg>
                    </button>
                    <button
                      class="tab-icon-btn"
                      :class="{ active: activePickerTab === 'emoji' }"
                      title="Emoji"
                      @click="activePickerTab = 'emoji'"
                    >
                      <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="10"></circle>
                        <path d="M8 14s1.5 2 4 2 4-2 4-2"></path>
                        <line x1="9" y1="9" x2="9.01" y2="9"></line>
                        <line x1="15" y1="9" x2="15.01" y2="9"></line>
                      </svg>
                    </button>
                  </div>

                  <!-- STICKERS TAB -->
                  <div v-if="activePickerTab === 'stickers'" class="picker-body">

                    <!-- Animated: grouped by pack -->
                    <div v-if="activeStickerCategory === 'animated'" class="sticker-packs-scroll">
                      <div
                        v-for="group in animatedStickerGroups"
                        :key="group.name"
                        class="sticker-pack-group"
                      >
                        <div class="sticker-pack-title">{{ group.name }}</div>
                        <div class="emoji-grid sticker-grid">
                          <div
                            v-for="sticker in group.items"
                            :key="sticker.id"
                            class="emoji-grid-item"
                            @click="selectSticker(sticker)"
                          >
                            <img :src="sticker.src" alt="sticker" />
                          </div>
                        </div>
                      </div>
                    </div>

                    <!-- All / Mine: flat grid -->
                    <div v-else class="emoji-grid sticker-grid">
                      <div
                        v-for="sticker in filteredStickers"
                        :key="sticker.id"
                        class="emoji-grid-item"
                        @click="selectSticker(sticker)"
                      >
                        <img :src="sticker.src" alt="sticker" />
                      </div>
                    </div>

                    <div class="picker-footer">
                      <button
                        class="footer-tab-btn"
                        :class="{ active: activeStickerCategory === 'all' }"
                        @click="activeStickerCategory = 'all'"
                      >
                        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="7" height="7" rx="1"></rect><rect x="14" y="3" width="7" height="7" rx="1"></rect><rect x="3" y="14" width="7" height="7" rx="1"></rect><rect x="14" y="14" width="7" height="7" rx="1"></rect></svg>
                        All
                      </button>
                      <button
                        class="footer-tab-btn"
                        :class="{ active: activeStickerCategory === 'animated' }"
                        @click="activeStickerCategory = 'animated'"
                      >
                        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="6 3 20 12 6 21 6 3"></polygon></svg>
                        Animated
                      </button>
                      <button
                        class="footer-tab-btn"
                        :class="{ active: activeStickerCategory === 'mine' }"
                        @click="activeStickerCategory = 'mine'"
                      >
                        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path></svg>
                        My stickers
                      </button>
                      <button class="footer-tab-btn create-sticker-btn" @click.stop="isStickerCreatorOpen = true">
                        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
                        Create Sticker
                      </button>
                      <CreateSticker 
      v-if="isStickerCreatorOpen"
      @tool-change="handleToolChanged" 
      @create="saveSticker" 
      @close="isStickerCreatorOpen = false" 
    />
                    </div>
                  </div>

                  <!-- EMOJI TAB -->
                  <div v-else class="picker-body">
                    <div class="emoji-grid emoji-only-grid">
                      <div
                        v-for="(emo, idx) in emojiPool"
                        :key="idx"
                        class="emoji-only-item"
                        @click="insertEmoji(emo)"
                      >
                        {{ emo }}
                      </div>
                    </div>
                  </div>
                </div>
              </Transition>
            </div>

            <div class="comment-input-container" :class="{ recording: isRecording }">
              <template v-if="!isRecording">
                <input
                  type="text"
                  ref="commentInputRef"
                  v-model="newComment"
                  @keyup.enter="addTextComment"
                  placeholder="Share your thoughts..."
                />

                <input
                  type="file"
                  ref="fileInputRef"
                  @change="handleImageUpload"
                  accept="image/*"
                  style="display: none;"
                />

                <div class="attach-wrapper" ref="attachMenuWrapperRef">
                  <button
                    class="attach-at-btn badge-input-style translucent-badge-sm"
                    @click.stop="toggleAttachMenu"
                    title="Attach"
                  >
                    <span class="at-symbol">@</span>
                  </button>

                  <Transition name="fade-pop">
                    <div v-if="showAttachMenu" class="attach-dropdown-menu">
                      <button class="attach-dropdown-item" @click="handleAttachOption('tag-friend')">
                        <span class="attach-dropdown-icon">
                          <span class="at-symbol at-symbol-sm">@</span>
                        </span>
                        <span class="attach-dropdown-label">Tag a friend</span>
                      </button>
                      <button class="attach-dropdown-item" @click="handleAttachOption('music')">
                        <span class="attach-dropdown-icon">
                          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 18V5l12-2v13"></path><circle cx="6" cy="18" r="3"></circle><circle cx="18" cy="16" r="3"></circle></svg>
                        </span>
                        <span class="attach-dropdown-label">Music</span>
                      </button>
                      <button class="attach-dropdown-item" @click="handleAttachOption('video')">
                        <span class="attach-dropdown-icon">
                          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="5" width="15" height="14" rx="2"></rect><polygon points="22 8 17 12 22 16 22 8"></polygon></svg>
                        </span>
                        <span class="attach-dropdown-label">Video</span>
                      </button>
                      <button class="attach-dropdown-item" @click="handleAttachOption('photo')">
                        <span class="attach-dropdown-icon">
                          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
                        </span>
                        <span class="attach-dropdown-label">Photo</span>
                      </button>
                    </div>
                  </Transition>
                </div>

                <button
                  v-if="newComment.trim()"
                  class="send-action-btn animate-pop"
                  @click="addTextComment"
                  title="Send comment"
                >
                  <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
                </button>
              </template>

              <div v-else class="live-recording-bar">
                <div class="pulse-wave-anim">
                  <span class="p-bar"></span>
                  <span class="p-bar"></span>
                  <span class="p-bar"></span>
                  <span class="p-bar"></span>
                </div>
                <span class="rec-timer">Recording 0:0{{ recordingTime }}</span>
              </div>
            </div>

            <div class="voice-record-wrapper">
              <button
                v-if="!isRecording"
                class="mic-glow-btn"
                @click="startVoiceRecording"
                title="Record 5s voice comment"
              >
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"></path><path d="M19 10v1a7 7 0 0 1-14 0v-1"></path><line x1="12" y1="19" x2="12" y2="23"></line><line x1="8" y1="23" x2="16" y2="23"></line></svg>
              </button>

              <button
                v-else
                class="mic-stop-btn animate-pulse"
                @click="stopVoiceRecording"
                title="Stop recording"
              >
                <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor"><rect x="4" y="4" width="16" height="16" rx="2"></rect></svg>
              </button>
            </div>
          </div>
        </div>
      </div>

      <div class="sidebar-bottom">
        <div class="sidebar-box">
          <div class="sidebar-title">Recommended Songs</div>
          <div
            v-for="item in recommendedList"
            :key="'rec-' + item.id"
            class="song-card"
            @click="playSong(item)"
          >
            <div class="song-thumb">
              <img v-if="item.cover" :src="item.cover" alt="" />
              <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#64748B" stroke-width="2"><path d="M9 18V5l12-2v13"></path><circle cx="6" cy="18" r="3"></circle><circle cx="18" cy="16" r="3"></circle></svg>
            </div>
            <div class="song-info">
              <div class="s-name">{{ item.title }}</div>
              <div class="s-singer">{{ item.singer }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
  </div>
</template>

<style scoped>
* {
  box-sizing: border-box;
}

.wrap-player{
  padding-right: 14px;
}

.player-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
  background-color: #ffffff;
  color: #1e293b;
  padding: 22px;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
  border-radius: 16px;
  width: 100%;
  max-width: 100%;
  overflow-x: hidden;
}

.player-top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.translucent-badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 16px;
  border-radius: 12px;
  background: rgba(25, 118, 210, 0.06);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  border: 1px solid rgba(25, 118, 210, 0.15);
  color: #1976D2;
  font-size: 13px;
  font-weight: 600;
}

.translucent-badge-sm {
  background: rgba(241, 245, 249, 0.6) !important;
  backdrop-filter: blur(6px);
  -webkit-backdrop-filter: blur(6px);
  border: 1px solid rgba(226, 232, 240, 0.8) !important;
  border-radius: 10px !important;
  color: #475569 !important;
}

.top-actions {
  display: flex;
  gap: 10px;
}

.btn-action {
  background: #1976D2;
  border: 1px solid #e2e8f0;
  color: #ffffff;
  padding: 6px 14px;
  border-radius: 32px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  font-weight: 500;
  transition: all 0.2s;
}

.btn-action:hover {
  opacity: 0.8;
}

.count-pill {
  background: rgba(255, 255, 255, 0.25);
  padding: 1px 7px;
  border-radius: 8px;
  font-size: 11px;
  font-weight: 700;
}

.svg-bg {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  transition: background 0.2s ease;
}

.btn-action:hover .svg-bg {
  background: rgba(255, 255, 255, 0.35);
}

.reaction-wrapper {
  position: relative;
  display: inline-block;
}

/* Composite Like button + clickable count */
.like-btn-group {
  display: flex;
  align-items: stretch;
  background: #1976D2;
  border: 1px solid #e2e8f0;
  border-radius: 32px;
  overflow: visible;
}

.like-btn-group.liked {
  background: #0d47a1;
}

.btn-action-inner {
  background: transparent;
  border: none;
  color: #ffffff;
  padding: 6px 8px 6px 14px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  font-weight: 500;
  border-radius: 12px 0 0 12px;
}

.btn-action-inner:hover {
  opacity: 0.85;
}

.likers-popup-wrapper {
  position: relative;
  display: flex;
}

.count-pill-btn {
  background: rgba(255, 255, 255, 0.25);
  border: none;
  color: #ffffff;
  padding: 6px 12px;
  border-radius: 0 12px 12px 0;
  font-size: 11px;
  font-weight: 700;
  cursor: pointer;
  transition: background 0.2s;
}

.count-pill-btn:hover {
  background: rgba(255, 255, 255, 0.42);
}

.count-pill-btn.disabled {
  cursor: default;
}

.count-pill-btn.disabled:hover {
  background: rgba(255, 255, 255, 0.25);
}

/* Liked-by popover, shared by main like + comment likes */
.likers-popup {
  position: absolute;
  top: 115%;
  left: 0;
  min-width: 200px;
  background-color: #ffffff;
  border: 1.5px solid rgba(226, 232, 240, 0.8);
  border-radius: 14px;
  padding: 10px 12px;
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.15), 0 10px 10px rgba(0, 0, 0, 0.08);
  z-index: 120;
  animation: popup-scale 0.2s ease-out;
}

.likers-popup-title {
  font-size: 11px;
  font-weight: 700;
  color: #94a3b8;
  text-transform: uppercase;
  letter-spacing: 0.4px;
  margin-bottom: 8px;
}

.likers-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  max-height: 200px;
  overflow-y: auto;
}

.likers-list-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.likers-avatar {
  width: 26px;
  height: 26px;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.likers-name {
  font-size: 13px;
  color: #1e293b;
  font-weight: 500;
}

.likers-more {
  font-size: 12px;
  color: #94a3b8;
  margin-top: 8px;
  padding-top: 8px;
  border-top: 1px solid #f1f5f9;
}

.comment-likers-popup {
  left: 50px;
  right: 0;
}

.reaction-popup-box {
  position: absolute;
  top: 115%;
  left: -40px;
  background-color: #ffffff;
  border: 1.5px solid rgba(226, 232, 240, 0.8);
  border-radius: 16px;
  padding: 6px 10px;
  display: flex;
  gap: 8px;
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.15), 0 10px 10px rgba(0, 0, 0, 0.08);
  z-index: 100;
  animation: popup-scale 0.2s ease-out;
}

@keyframes popup-scale {
  0% { transform: scale(0.8); opacity: 0; }
  100% { transform: scale(1); opacity: 1; }
}

.reaction-icon-item {
  width: 38px;
  height: 38px;
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  overflow: hidden;
  transition: transform 0.2s, background-color 0.2s;
}

.reaction-icon-item:hover {
  transform: scale(1.15) translateY(-4px);
  background-color: #f1f5f9;
}

.share-wrapper {
  position: relative;
  display: inline-block;
}

.share-dropdown-menu {
  position: absolute;
  top: 115%;
  right: 0;
  min-width: 260px;
  background: #ffffff;
  border-radius: 16px;
  padding: 8px;
  display: flex;
  flex-direction: column;
  gap: 2px;
  box-shadow: 0 18px 36px rgba(0, 0, 0, 0.15), 0 10px 10px rgba(0, 0, 0, 0.08);
  z-index: 150;
  animation: popup-scale 0.18s ease-out;
}

.share-dropdown-item {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
  background: transparent;
  border: none;
  border-radius: 12px;
  padding: 9px 12px;
  cursor: pointer;
  text-align: left;
  transition: background 0.15s ease, transform 0.15s ease;
}

.share-dropdown-item:hover {
  background: #eef2f6;
  transform: translateX(2px);
}

.share-dropdown-item:active {
  transform: scale(0.98);
}

.share-dropdown-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.06);
  color: #000000;
  flex-shrink: 0;
}

.share-dropdown-icon svg {
  width: 14px;
  height: 14px;
}

.share-dropdown-label {
  color: #1e293b;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.2px;
}

.share-dropdown-divider {
  height: 1px;
  background: #e2e8f0;
  border-radius: 2px;
  margin: 2px 0;
}

.copied-toast {
  position: absolute;
  top: 115%;
  right: 0;
  background: #0f172a;
  color: #ffffff;
  padding: 6px 12px;
  border-radius: 10px;
  font-size: 12px;
  font-weight: 600;
  white-space: nowrap;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.2);
  z-index: 100;
}

.fade-pop-enter-active,
.fade-pop-leave-active {
  transition: all 0.2s ease;
}

.fade-pop-enter-from,
.fade-pop-leave-to {
  opacity: 0;
  transform: translateY(-4px);
}

.media-row {
  display: flex;
  align-items: stretch;
  width: 100%;
  max-width: 100%;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1);
  border-top-left-radius: 16px;
  border-top-right-radius: 16px;
  overflow: hidden;
  max-height: 480px;
}

.video-screen {
  position: relative;
  flex: 1;
  background: #000000;
  aspect-ratio: 16 / 9;
  min-width: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  max-height: 100%;
}

.main-video {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.playing-next-attached {
  width: 340px;
  min-width: 340px;
  background-color: #2e2f33e8;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  align-self: stretch;
  height: auto;
  max-height: 100%;
  box-sizing: border-box;
  overflow: hidden;
}

.bottom-layout {
  display: grid;
  grid-template-columns: 1fr 340px;
  gap: 16px;
  width: 100%;
  max-width: 100%;
  margin-top: -16px;
}

.main-details {
  display: flex;
  flex-direction: column;
  gap: 16px;
  border: 1px solid #eeee;
  min-width: 0;
}

.sidebar-bottom {
  min-width: 0;
}

.light-text {
  color: #ffffff !important;
}

.dim-text {
  color: #94a3b8 !important;
}

.song-list-scroll {
  display: flex;
  flex-direction: column;
  gap: 8px;
  overflow-y: auto;
  flex: 1;
  max-height: calc(100% - 30px);
}

.song-card {
  border-radius: 12px;
  padding: 8px;
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  transition: background 0.2s;
  justify-content: space-between;
  width: 100%;
}

.dark-card:hover {
  background: rgba(255, 255, 255, 0.1);
}

.dark-card.active {
  background: rgba(255, 255, 255, 0.2);
}

.song-card:not(.dark-card):hover {
  background: #f1f5f9;
}

.song-thumb {
  width: 120px;
  height: 70px;
  background: #e2e8f0;
  border-radius: 10px;
  overflow: hidden;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.song-thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.sidebar-title {
  font-weight: 700;
  font-size: 15px;
  color: #0f172a;
}

.s-name {
  font-weight: 600;
  font-size: 13px;
  color: #0f172a;
  min-width: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.s-singer {
  font-size: 12px;
  color: #64748b;
}

.big-play-btn {
  position: absolute;
  background: rgba(255, 255, 255, 0.25);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255, 255, 255, 0.4);
  color: #ffffff;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.player-controls {
  position: absolute;
  bottom: 12px;
  left: 12px;
  right: 12px;
  display: flex;
  align-items: center;
  gap: 12px;
  background: rgba(255, 255, 255, 0.92);
  padding: 8px 16px;
  border-radius: 14px;
  opacity: 0;
  transition: opacity 0.3s;
}

.video-screen:hover .player-controls {
  opacity: 1;
}

.time-text {
  font-size: 12px;
  font-weight: 600;
  color: #475569;
}

.progress-track {
  flex: 1;
  height: 5px;
  background: #e2e8f0;
  border-radius: 10px;
  position: relative;
  cursor: pointer;
}

.progress-fill {
  height: 100%;
  background: #1976D2;
  border-radius: 10px;
}

.progress-thumb {
  position: absolute;
  top: 50%;
  width: 12px;
  height: 12px;
  background: #1976D2;
  border-radius: 50%;
  transform: translate(-50%, -50%);
}

.center-controls {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #64748b;
}

.volume-slider {
  width: 60px;
  height: 5px;
  background: #e2e8f0;
  border-radius: 10px;
  overflow: hidden;
}

.vol-fill {
  height: 100%;
  background: #1976D2;
}

.ctrl-btn {
  background: none;
  border: none;
  cursor: pointer;
}

.sidebar-box {
  background-color: #F0F0F0;
  width: 100%;
  border-radius: 14px;
}

.details-box, .comments-section, .sidebar-box {
  padding: 20px;
  width: 100%;
}

.description-header h3 {
  margin: 0 0 6px 0;
  font-size: 18px;
}

.stats-row {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #64748b;
  flex-wrap: wrap;
}

.song-icon-badge {
  display: inline-flex;
  align-items: center;
  gap: 5px;
}

.stats-likers-wrapper {
  position: relative;
  display: inline-flex;
}

.stat-badge-clickable {
  background: none;
  border: none;
  padding: 0;
  font: inherit;
  color: #64748b;
  cursor: pointer;
}

.stat-badge-clickable:hover {
  color: #1976D2;
  text-decoration: underline;
}

.stats-likers-popup {
  top: 130%;
  left: 0;
}

.desc-text {
  font-size: 14px;
  color: #334155;
  margin: 8px 0;
}

.channel-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 12px;
  border-top: 1px solid #f1f5f9;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.avatar-img {
  width: 40px;
  height: 40px;
  border-radius: 12px;
}

.avatar-sm {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  flex-shrink: 0;
}

.avatar-sm-xs {
  width: 26px;
  height: 26px;
  border-radius: 8px;
  flex-shrink: 0;
}

.username {
  font-weight: 600;
  font-size: 14px;
}

.subscribers {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #64748b;
}

.inline-song-icon {
  color: #64748b;
}

.subscribe-btn {
  background: #1976D2;
  color: #ffffff;
  border: none;
  padding: 8px 18px;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
}

.comments-header {
  font-weight: 700;
  font-size: 16px;
  margin-bottom: 16px;
}

.song-info {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.comments-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.comment-branch-node {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.comment-item {
  display: flex;
  gap: 12px;
  align-items: flex-start;
}

.comment-content-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
  min-width: 0;
}

.comment-body {
  background: #f8fafc;
  padding: 12px 16px;
  border-radius: 14px;
  border: 1.5px solid #e2e8f0;
}

.sub-comment-body {
  background: #f1f5f9;
}

.comment-user-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 4px;
}

.user-date-group {
  display: flex;
  align-items: center;
  gap: 8px;
}

.comment-user {
  font-weight: 600;
  font-size: 13px;
  color: #1e293b;
}

.comment-date.badge-style {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 2px 10px;
  border-radius: 10px;
  font-size: 11px;
  font-weight: 600;
}

.clock-icon {
  color: #64748b;
}

.comment-text {
  font-size: 13.5px;
  color: #334155;
  line-height: 1.4;
}

.outside-box-actions {
  margin-left: 44px;
  margin-top: 2px;
  display: flex;
  align-items: center;
  gap: 16px;
}

.comment-like-group {
  position: relative;
  display: inline-flex;
  align-items: center;
  gap: 4px;
}

.comment-like-icon-btn {
  background: none;
  border: none;
  color: #64748b;
  font-size: 12px;
  font-weight: 600;
  display: inline-flex;
  align-items: center;
  gap: 4px;
  cursor: pointer;
  padding: 0;
}

.comment-like-icon-btn.liked {
  color: #1976D2;
}

.comment-like-icon-btn:hover {
  color: #1976D2;
}

.comment-like-count-btn {
  background: #f1f5f9;
  border: none;
  color: #475569;
  font-size: 11px;
  font-weight: 700;
  padding: 2px 8px;
  border-radius: 10px;
  cursor: pointer;
  transition: background 0.2s;
}

.comment-like-count-btn:hover {
  background: #e2e8f0;
}

.reply-toggle-btn {
  background: none;
  border: none;
  color: #1976D2;
  font-size: 12px;
  font-weight: 600;
  display: inline-flex;
  align-items: center;
  gap: 4px;
  cursor: pointer;
  padding: 0;
}

.reply-toggle-btn:hover {
  text-decoration: underline;
}

.outside-box-tree {
  position: relative;
  margin-left: 44px;
  padding-left: 20px;
  margin-top: 4px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.tree-branch-line {
  position: absolute;
  top: 0;
  bottom: 12px;
  left: 0;
  width: 2px;
  background: #cbd5e1;
  border-radius: 2px;
}

.sub-comments-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.sub-comment-item {
  position: relative;
}

.outside-box-reply {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-left: 44px;
  margin-top: 6px;
}

.comment-photo-container {
  position: relative;
  margin-top: 8px;
  max-width: 240px;
  border-radius: 12px;
  overflow: visible;
  border: 1.5px solid #e2e8f0;
  background: #ffffff;
}

.comment-attached-photo {
  width: 100%;
  max-height: 180px;
  object-fit: cover;
  display: block;
  border-radius: 10px;
}

.comment-attached-photo.tagging-active {
  cursor: crosshair;
}

.tag-toggle-btn {
  position: absolute;
  top: 8px;
  right: 8px;
  display: flex;
  align-items: center;
  gap: 4px;
  background: rgba(15, 23, 42, 0.65);
  color: #ffffff;
  border: none;
  border-radius: 20px;
  padding: 5px 10px;
  font-size: 11px;
  font-weight: 600;
  cursor: pointer;
  backdrop-filter: blur(4px);
  transition: background 0.2s;
}

.tag-toggle-btn:hover {
  background: rgba(15, 23, 42, 0.85);
}

.tag-toggle-btn.active {
  background: #1976D2;
}

.photo-tag-marker {
  position: absolute;
  transform: translate(-50%, -50%);
  display: flex;
  align-items: center;
  gap: 5px;
}

.photo-tag-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #ffffff;
  border: 2px solid #1976D2;
  box-shadow: 0 0 0 3px rgba(25, 118, 210, 0.25);
  flex-shrink: 0;
}

.photo-tag-label {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  background: rgba(15, 23, 42, 0.75);
  color: #ffffff;
  font-size: 11px;
  font-weight: 600;
  padding: 3px 8px;
  border-radius: 10px;
  white-space: nowrap;
}

.photo-tag-remove {
  background: none;
  border: none;
  color: #ffffff;
  font-size: 12px;
  line-height: 1;
  cursor: pointer;
  padding: 0;
  opacity: 0.8;
}

.photo-tag-remove:hover {
  opacity: 1;
}

.photo-tag-input-box {
  position: absolute;
  transform: translate(-50%, -50%);
  display: flex;
  align-items: center;
  gap: 4px;
  background: #ffffff;
  border: 1.5px solid #1976D2;
  border-radius: 12px;
  padding: 4px 6px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.18);
  z-index: 20;
}

.photo-tag-input-box input {
  width: 110px;
  border: none;
  outline: none;
  font-size: 12px;
  color: #1e293b;
}

.photo-tag-confirm,
.photo-tag-cancel {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 13px;
  line-height: 1;
  padding: 2px 4px;
  border-radius: 6px;
  flex-shrink: 0;
}

.photo-tag-confirm {
  color: #16a34a;
}

.photo-tag-cancel {
  color: #ef4444;
}

.voice-comment-bubble {
  display: flex;
  align-items: center;
  gap: 12px;
  background: #ffffff;
  border: 1.5px solid #e2e8f0;
  border-radius: 14px;
  padding: 8px 12px;
  margin-top: 8px;
}

.voice-play-bubble-btn {
  background: #1976D2;
  color: #ffffff;
  border: none;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  flex-shrink: 0;
}

.voice-waveform-preview {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
  position: relative;
}

.wave-bars {
  display: flex;
  align-items: center;
  gap: 3px;
  height: 24px;
}

.wave-bars .bar {
  width: 3px;
  background: #cbd5e1;
  border-radius: 2px;
  transition: background 0.2s;
}

.wave-bars .bar.active {
  background: #1976D2;
}

.voice-progress-track {
  width: 100%;
  height: 3px;
  background: #e2e8f0;
  border-radius: 2px;
  overflow: hidden;
}

.voice-progress-fill {
  height: 100%;
  background: #1976D2;
  width: 0%;
}

.voice-pill-tag {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 11px;
  font-weight: 600;
  color: #64748b;
  background: #f1f5f9;
  padding: 2px 8px;
  border-radius: 10px;
  flex-shrink: 0;
}

.add-comment-box {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid #f1f5f9;
}

.comment-input-container {
  flex: 1;
  display: flex;
  align-items: center;
  background: #f8fafc;
  border: 1.5px solid #e2e8f0;
  border-radius: 14px;
  padding: 6px 12px;
  gap: 8px;
  transition: all 0.2s;
}

.comment-input-container:focus-within {
  border-color: #1976D2;
  background: #ffffff;
}

.comment-input-container input[type="text"] {
  flex: 1;
  border: none;
  background: transparent;
  outline: none;
  font-size: 13.5px;
  color: #1e293b;
}

.badge-input-style {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 4px 10px;
  border-radius: 10px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}

.badge-input-style:hover {
  background: rgba(226, 232, 240, 0.8) !important;
}

.attach-wrapper {
  position: relative;
  display: inline-flex;
  flex-shrink: 0;
}

.attach-at-btn {
  width: 30px;
  height: 30px;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.at-symbol {
  font-size: 15px;
  font-weight: 800;
  line-height: 1;
  color: #475569;
}

.at-symbol-sm {
  font-size: 13px;
  color: #000000;
}

.attach-dropdown-menu {
  position: absolute;
  bottom: 115%;
  left: 0;
  min-width: 190px;
  background: #ffffff;
  border-radius: 16px;
  padding: 8px;
  display: flex;
  flex-direction: column;
  gap: 2px;
  box-shadow: 0 18px 36px rgba(0, 0, 0, 0.15), 0 10px 10px rgba(0, 0, 0, 0.08);
  z-index: 150;
  animation: popup-scale 0.18s ease-out;
}

.attach-dropdown-item {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
  background: transparent;
  border: none;
  border-radius: 12px;
  padding: 9px 12px;
  cursor: pointer;
  text-align: left;
  transition: background 0.15s ease, transform 0.15s ease;
}

.attach-dropdown-item:hover {
  background: #f1f5f9;
  transform: translateX(2px);
}

.attach-dropdown-item:active {
  transform: scale(0.98);
}

.attach-dropdown-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.06);
  color: #000000;
  flex-shrink: 0;
}

.attach-dropdown-label {
  color: #1e293b;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.2px;
}

.send-action-btn {
  background: #1976D2;
  color: #ffffff;
  border: none;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  flex-shrink: 0;
}

.voice-record-wrapper {
  display: flex;
  align-items: center;
}

.mic-glow-btn {
  background: #f8fafc;
  border: 1.5px solid #e2e8f0;
  color: #1976D2;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
}

.mic-glow-btn:hover {
  background: #e3f2fd;
  border-color: #1976D2;
}

.mic-stop-btn {
  background: #ef4444;
  color: #ffffff;
  border: none;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.live-recording-bar {
  display: flex;
  align-items: center;
  gap: 10px;
  flex: 1;
  padding: 0 4px;
}

.rec-timer {
  font-size: 12px;
  font-weight: 700;
  color: #ef4444;
}

.pulse-wave-anim {
  display: flex;
  align-items: center;
  gap: 3px;
}

.p-bar {
  width: 3px;
  height: 14px;
  background: #ef4444;
  border-radius: 2px;
  animation: pulse-wave 0.8s infinite ease-in-out alternate;
}

.share-btn-group{
    display:flex;
    align-items:stretch;
    background:#1976D2;
    border:1px solid #e2e8f0;
    border-radius:32px;
    overflow:hidden;
}

.share-btn-group .btn-action-inner{
    background:transparent;
    border:none;
    color:#fff;
    padding:6px 12px;
    display:flex;
    align-items:center;
    gap:8px;
    cursor:pointer;
    font-size:13px;
    font-weight:600;
}

.share-btn-group .btn-action-inner:hover{
    opacity:.85;
}

.share-btn-group .count-pill-btn{
    border-left:1px solid rgba(255,255,255,.18);
    border-radius:0;
}

.p-bar:nth-child(2) { animation-delay: 0.2s; }
.p-bar:nth-child(3) { animation-delay: 0.4s; }
.p-bar:nth-child(4) { animation-delay: 0.6s; }

/* ===== Emoji / Sticker Picker ===== */

.emoji-picker-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.emoji-toggle-btn {
  background: #f8fafc;
  border: 1.5px solid #e2e8f0;
  color: #64748b;
  padding: 8px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.emoji-toggle-btn:hover {
  color: #1976D2;
  border-color: #1976D2;
  background: #f1f5f9;
}

.emoji-picker-dropdown {
  position: absolute;
  bottom: 125%;
  left: 0;
  width: 610px;
  max-width: 90vw;
  background: #ffffff;
  border: 1.5px solid rgba(226, 232, 240, 0.9);
  border-radius: 16px;
  box-shadow: 0 18px 40px rgba(0, 0, 0, 0.18), 0 8px 16px rgba(0, 0, 0, 0.08);
  z-index: 200;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.emoji-picker-header {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 10px 14px;
  border-bottom: 1px solid #f1f5f9;
  background: #fafbfc;
}

.tab-icon-btn {
  background: none;
  border: none;
  color: #94a3b8;
  padding: 8px 10px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  position: relative;
  transition: color 0.15s ease;
}

.tab-icon-btn:hover {
  color: #475569;
}

.tab-icon-btn.active {
  color: #1976D2;
}

.tab-icon-btn.active::after {
  content: '';
  position: absolute;
  bottom: -10px;
  left: 6px;
  right: 6px;
  height: 3px;
  border-radius: 3px;
  background: #1976D2;
}

.picker-body {
  display: flex;
  flex-direction: column;
  min-height: 0;
}

.emoji-grid {
  display: grid;
  gap: 10px;
  padding: 14px;
  max-height: 320px;
  overflow-y: auto;
}

.sticker-grid {
  grid-template-columns: repeat(8, 1fr);
}

.sticker-packs-scroll {
  max-height: 320px;
  overflow-y: auto;
  padding: 14px 14px 4px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.sticker-pack-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.sticker-pack-group + .sticker-pack-group {
  margin-top: 10px;
  padding-top: 14px;
  border-top: 1px solid #f1f5f9;
}

.sticker-pack-title {
  font-size: 12px;
  font-weight: 700;
  color: #475569;
  letter-spacing: 0.2px;
}

.sticker-packs-scroll .emoji-grid.sticker-grid {
  padding: 0;
  max-height: none;
  overflow-y: visible;
}

.emoji-grid-item {
  aspect-ratio: 1;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  background: #fdf2f2;
  transition: transform 0.15s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.emoji-grid-item:hover {
  transform: scale(1.06);
}

.emoji-grid-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.emoji-only-grid {
  grid-template-columns: repeat(7, 1fr);
  max-height: 240px;
}

.emoji-only-item {
  aspect-ratio: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  border-radius: 10px;
  cursor: pointer;
  transition: background 0.15s ease, transform 0.15s ease;
}

.emoji-only-item:hover {
  background: #f1f5f9;
  transform: scale(1.15);
}

.picker-footer {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 10px 12px;
  border-top: 1px solid #f1f5f9;
  overflow-x: auto;
}

.footer-tab-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: none;
  border: none;
  color: #64748b;
  font-size: 12.5px;
  font-weight: 600;
  padding: 6px 10px;
  border-radius: 10px;
  cursor: pointer;
  white-space: nowrap;
  transition: background 0.15s ease, color 0.15s ease;
}

.footer-tab-btn:hover {
  background: #f1f5f9;
  color: #1976D2;
}

.footer-tab-btn.active {
  color: #1976D2;
  background: #eaf3fc;
}

.create-sticker-btn {
  color: #b45309;
  margin-left: auto;
}

.create-sticker-btn:hover {
  background: #fef3e2;
  color: #b45309;
}

@keyframes pulse-wave {
  0% { height: 6px; }
  100% { height: 18px; }
}

.animate-pop {
  animation: pop-in 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

@keyframes pop-in {
  0% { transform: scale(0.6); opacity: 0; }
  100% { transform: scale(1); opacity: 1; }
}
</style>