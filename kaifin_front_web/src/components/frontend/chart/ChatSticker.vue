<template>
  <div class="chat-sticker-container">
    <div class="sticker-top-tabs">
      <button
        v-for="pack in packs"
        :key="pack.id"
        class="category-icon-btn"
        :class="{ active: currentPackId === pack.id }"
        @click="selectPack(pack.id)"
        :title="pack.name"
      >
        <svg class="cat-avatar" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M12 2l2.4 7.4H22l-6 4.4 2.3 7.2-6.3-4.6-6.3 4.6 2.3-7.2-6-4.4h7.6z" />
        </svg>
      </button>
      <div v-if="packsLoading" class="sticker-tabs-loading">...</div>
    </div>
    <div class="main-content-wrapper">
      <div class="sticker-grid-content">
        <div class="sticker-grid">
          <div class="sticker-item add-sticker-grid-btn" @click="handleAddStickerClick" title="Add Sticker">
            <div class="add-icon-wrapper">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <line x1="12" y1="5" x2="12" y2="19"></line>
                <line x1="5" y1="12" x2="19" y2="12"></line>
              </svg>
            </div>
          </div>

          <input
            type="file"
            ref="stickerFileInputRef"
            accept="image/png,image/webp,image/gif,image/jpeg,image/jpg,image/svg+xml"
            style="display: none;"
            @change="handleStickerFileSelected"
          />

          <div v-if="uploadingSticker" class="sticker-item sticker-uploading-placeholder">
            <span>...</span>
          </div>

          <div v-if="stickersLoading" class="sticker-grid-loading">Loading...</div>
          <div v-else-if="currentStickers.length === 0" class="sticker-grid-empty">No stickers</div>
          <div
            v-for="sticker in currentStickers"
            :key="sticker.id"
            class="sticker-item"
            @click="selectSticker(sticker)"
          >
            <img :src="sticker.url" :alt="'Sticker ' + sticker.id" />
          </div>
        </div>
      </div>
      <div class="sticker-tabs-footer">
        <button
          v-for="(tab, index) in tabs"
          :key="index"
          class="sticker-tab-btn"
          :class="{ active: currentTab === tab.id }"
          @click="currentTab = tab.id"
        >
          {{ tab.name }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const emit = defineEmits(['select-sticker'])
const API_BASE = 'http://localhost:7070/api/v1/front'
const FILE_BASE = 'http://localhost:7070'

function authHeaders() {
  const token = localStorage.getItem('token')
  return { Authorization: `Bearer ${token}` }
}

function resolveImage(path) {
  if (!path) return ''
  if (/^(https?:|blob:|data:)/.test(path)) return path
  if (/^\/?image\//.test(path)) {
    const clean = path.startsWith('/') ? path : `/${path}`
    return `${API_BASE}${clean}`
  }
  const cleanPath = path.startsWith('/') ? path : `/${path}`
  return `${FILE_BASE}/uploads${cleanPath}`
}


const currentTab = ref('stickers')
const tabs = [
  { id: 'stickers', name: 'Stickers' },
  { id: 'cards', name: 'Cards' },
  { id: 'emoticons', name: 'Emoticons' },
  { id: 'constructor', name: 'Constructor' },
]

const packs = ref([])
const packsLoading = ref(false)
const currentPackId = ref(null)

async function fetchPacks() {
  packsLoading.value = true
  try {
    const res = await axios.get(`${API_BASE}/stickers/packs`, {
      headers: authHeaders(),
    })
    if (res.data.success) {
      packs.value = res.data.data.packs || []
      if (packs.value.length > 0 && !currentPackId.value) {
        selectPack(packs.value[0].id)
      }
    }
  } catch (err) {
    console.log('FETCH STICKER PACKS ERROR:', err)
  } finally {
    packsLoading.value = false
  }
}

const stickersByPack = ref({}) 
const stickersLoading = ref(false)

async function fetchStickersForPack(packId, force = false) {
  if (!force && stickersByPack.value[packId]) return 
  stickersLoading.value = true
  try {
    const res = await axios.get(`${API_BASE}/stickers/show`, {
      headers: authHeaders(),
      params: { pack_id: packId },
    })
    if (res.data.success) {
      const stickers = (res.data.data.stickers || []).map((s) => ({
        ...s,
        url: resolveImage(s.url),
      }))
      stickersByPack.value = {
        ...stickersByPack.value,
        [packId]: stickers,
      }
    }
  } catch (err) {
    console.log('FETCH STICKERS ERROR:', err)
  } finally {
    stickersLoading.value = false
  }
}

function selectPack(packId) {
  currentPackId.value = packId
  fetchStickersForPack(packId)
}

const currentStickers = computed(() => stickersByPack.value[currentPackId.value] || [])

function selectSticker(sticker) {
  emit('select-sticker', sticker)
}

const stickerFileInputRef = ref(null)
const uploadingSticker = ref(false)

function handleAddStickerClick() {
  if (!currentPackId.value) {
    alert('សូមជ្រើសរើស pack មុននឹងបន្ថែម sticker')
    return
  }
  if (stickerFileInputRef.value) {
    stickerFileInputRef.value.click()
  }
}

async function handleStickerFileSelected(event) {
  const file = event.target.files?.[0]
  event.target.value = '' 
  if (!file || !currentPackId.value) return

  const allowedTypes = [
    'image/png',
    'image/webp',
    'image/gif',
    'image/jpeg',
    'image/jpg',
    'image/svg+xml',
  ]
  if (!allowedTypes.includes(file.type)) {
    alert('please PNG, JPG, GIF, WEBP or SVG')
    return
  }

  uploadingSticker.value = true
  try {
    const form = new FormData()
    form.append('pack_id', currentPackId.value)
    form.append('trigger_code', '') 
    form.append('file', file)   

    const res = await axios.post(`${API_BASE}/stickers/create`, form, {
      headers: {
        ...authHeaders(),
        'Content-Type': 'multipart/form-data',
      },
    })

    if (res.data.success) {
      await fetchStickersForPack(currentPackId.value, true)
    } else {
      alert(res.data.message || 'Failed to add sticker')
    }
  } catch (err) {
    console.log('CREATE STICKER ERROR:', err)
    if (err.response) {
      alert(err.response.data?.message || 'Failed to add sticker')
    } else {
      alert('Server error')
    }
  } finally {
    uploadingSticker.value = false
  }
}

onMounted(() => {
  fetchPacks()
})
</script>

<style scoped>
.chat-sticker-container {
  width: 640px; 
  height: 370px;
  /* height: 100vh; */
  position: absolute;
  right: 0;
  background-color: #ffffff;
  border: 1px solid #e4e6eb;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
  box-sizing: border-box;
}

.sticker-top-tabs {
  width: 100%;
  height: 52px;
  background-color: #ffffff;
  border-bottom: 1px solid #e4e6eb;
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 0 12px;
  gap: 8px;
  flex-shrink: 0;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
}

.category-icon-btn {
  border: none;
  background: transparent;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #65676b;
  cursor: pointer;
  transition: all 0.15s;
}

.category-icon-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
  color: #050505;
}

.category-icon-btn.active {
  background-color: rgba(0, 132, 255, 0.1);
  color: #0084ff;
}

.avatar-cat-btn {
  padding: 0;
  overflow: hidden;
}

.cat-avatar {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  object-fit: cover;
  display: block;
}

.main-content-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background-color: #f9fafb;
  min-height: 0;
}

.sticker-grid-content {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  padding: 16px;
}

.sticker-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr); 
  gap: 12px;

}

.sticker-item {
  aspect-ratio: 1 / 1;
  background: #FFFFFF;
  border-radius: 8px;
  padding: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid transparent;
  transition: all 0.15s ease;
  box-shadow: 0 1px 2px rgba(0,0,0,0.04);
}

.sticker-item:hover {
  opacity: 0.8;
}

.sticker-item img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  display: block;
}


.add-sticker-grid-btn {
  color: #000;
  background-color: rgba(0, 0, 0, 0.054);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.add-sticker-grid-btn:hover {
  opacity: 0.8;
}

.add-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
}


.sticker-tabs-footer {
  display: flex;
  border-top: 1px solid #e4e6eb;
  padding: 0 12px;
  background-color: #ffffff;
  flex-shrink: 0;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.sticker-tab-btn {
  background: transparent;
  border: none;
  padding: 14px 16px;
  font-size: 15px;
  font-weight: 500;
  color: #65676b;
  cursor: pointer;
  position: relative;
  transition: color 0.2s;
}

.sticker-tab-btn:hover {
  color: #050505;
}

.sticker-tab-btn.active {
  color: #0084ff;
  font-weight: 600;
}

.sticker-tab-btn.active::after {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background-color: #0084ff;
  border-radius: 0 0 3px 3px;
}
</style>