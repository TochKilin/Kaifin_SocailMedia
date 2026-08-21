<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  apiBase: String
})

const emit = defineEmits(['back', 'uploaded'])

const title = ref('')
const content = ref('')
const duration = ref(0)

// File state
const mediaFile = ref(null)
const mediaPreviewUrl = ref(null)
const coverFile = ref(null)
const coverPreviewUrl = ref(null)

const isSubmitting = ref(false)
const errorMessage = ref(null)
const successMessage = ref(null)

const mediaInput = ref(null)
const coverInput = ref(null)

const canSubmit = computed(() => {
  return title.value.trim().length > 0 && mediaFile.value && !isSubmitting.value
})

const onMediaChange = (e) => {
  const file = e.target.files?.[0]
  if (!file) return

  mediaFile.value = file
  if (mediaPreviewUrl.value) URL.revokeObjectURL(mediaPreviewUrl.value)
  mediaPreviewUrl.value = URL.createObjectURL(file)

  const isVideo = file.type.startsWith('video')
  const probe = document.createElement(isVideo ? 'video' : 'audio')
  probe.preload = 'metadata'
  probe.onloadedmetadata = () => {
    if (!isNaN(probe.duration) && isFinite(probe.duration)) {
      duration.value = Math.round(probe.duration)
    }
    URL.revokeObjectURL(probe.src)
  }
  probe.src = URL.createObjectURL(file)
}

const onCoverChange = (e) => {
  const file = e.target.files?.[0]
  if (!file) return

  coverFile.value = file
  if (coverPreviewUrl.value) URL.revokeObjectURL(coverPreviewUrl.value)
  coverPreviewUrl.value = URL.createObjectURL(file)
}

const removeMedia = () => {
  mediaFile.value = null
  if (mediaPreviewUrl.value) URL.revokeObjectURL(mediaPreviewUrl.value)
  mediaPreviewUrl.value = null
  if (mediaInput.value) mediaInput.value.value = ''
  duration.value = 0
}

const removeCover = () => {
  coverFile.value = null
  if (coverPreviewUrl.value) URL.revokeObjectURL(coverPreviewUrl.value)
  coverPreviewUrl.value = null
  if (coverInput.value) coverInput.value.value = ''
}

const formatDuration = (secs) => {
  if (!secs || isNaN(secs)) return '0:00'
  const minutes = Math.floor(secs / 60)
  const seconds = Math.floor(secs % 60).toString().padStart(2, '0')
  return `${minutes}:${seconds}`
}

let previousBodyOverflow = ''
let previousHtmlOverflow = ''
let previousBodyPosition = ''
let previousBodyTop = ''
let previousBodyWidth = ''
let scrollY = 0

onMounted(() => {
  scrollY = window.scrollY

  previousBodyOverflow = document.body.style.overflow
  previousHtmlOverflow = document.documentElement.style.overflow
  previousBodyPosition = document.body.style.position
  previousBodyTop = document.body.style.top
  previousBodyWidth = document.body.style.width

  document.documentElement.style.overflow = 'hidden'
  document.body.style.overflow = 'hidden'
  document.body.style.position = 'fixed'
  document.body.style.top = `-${scrollY}px`
  document.body.style.width = '100%'
})

onUnmounted(() => {
  document.documentElement.style.overflow = previousHtmlOverflow
  document.body.style.overflow = previousBodyOverflow
  document.body.style.position = previousBodyPosition
  document.body.style.top = previousBodyTop
  document.body.style.width = previousBodyWidth
  window.scrollTo(0, scrollY)
})

const submitUpload = async () => {
  if (!canSubmit.value) return

  isSubmitting.value = true
  errorMessage.value = null
  successMessage.value = null

  try {
    const formData = new FormData()
    formData.append('title', title.value.trim())
    formData.append('content', content.value.trim())
    formData.append('duration', String(duration.value || 0))
    formData.append('file_url', mediaFile.value)
    if (coverFile.value) {
      formData.append('cover_url', coverFile.value)
    }

    const token = localStorage.getItem('token')
    const res = await fetch(`${props.apiBase}/api/v1/front/songs/create`, {
      method: 'POST',
      headers: {
        ...(token ? { Authorization: `Bearer ${token}` } : {})
      },
      body: formData
    })

    const data = await res.json()

    if (!res.ok || data.success === false) {
      throw new Error(data.message || 'Upload failed')
    }

    successMessage.value = data.message || 'Song created successfully'
    emit('uploaded', data.data)
    setTimeout(() => {
      emit('back')
    }, 800)
  } catch (err) {
    errorMessage.value = err.message || 'Something went wrong while uploading.'
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <div class="upload-music-view">
    <div class="header-box card box-ra">
      <div class="header-top-row">
        <div class="header-left">
          <button class="back-btn" @click="$emit('back')" title="Back">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="15 18 9 12 15 6"></polyline>
            </svg>
          </button>
          <h2 class="section-main-title">Upload Music</h2>
        </div>
      </div>
    </div>

    <section class="box form-box card">
      <div v-if="errorMessage" class="alert alert-error">{{ errorMessage }}</div>
      <div v-if="successMessage" class="alert alert-success">{{ successMessage }}</div>

      <form class="upload-form" @submit.prevent="submitUpload">
        <!-- Cover + Media pickers -->
        <div class="pickers-row">
          <div class="picker cover-picker">
            <label class="picker-label">Cover Image</label>
            <div class="cover-drop" @click="coverInput.click()">
              <img v-if="coverPreviewUrl" :src="coverPreviewUrl" alt="cover preview" />
              <div v-else class="picker-placeholder">
                <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect x="3" y="3" width="18" height="18" rx="2"></rect>
                  <circle cx="8.5" cy="8.5" r="1.5"></circle>
                  <polyline points="21 15 16 10 5 21"></polyline>
                </svg>
                <span>Add cover</span>
              </div>
              <button v-if="coverPreviewUrl" type="button" class="remove-btn" @click.stop="removeCover">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                  <line x1="18" y1="6" x2="6" y2="18"></line>
                  <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
              </button>
            </div>
            <input ref="coverInput" type="file" accept="image/*" class="hidden-input" @change="onCoverChange" />
          </div>

          <div class="picker media-picker">
            <label class="picker-label">Song / Video File</label>
            <div class="media-drop" @click="mediaInput.click()">
              <template v-if="mediaFile">
                <div class="media-chip">
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M9 18V5l12-2v13"></path>
                    <circle cx="6" cy="18" r="3"></circle>
                    <circle cx="18" cy="16" r="3"></circle>
                  </svg>
                  <div class="media-chip-info">
                    <span class="media-name">{{ mediaFile.name }}</span>
                    <span class="media-meta">{{ formatDuration(duration) }}</span>
                  </div>
                  <button type="button" class="remove-btn inline" @click.stop="removeMedia">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                      <line x1="18" y1="6" x2="6" y2="18"></line>
                      <line x1="6" y1="6" x2="18" y2="18"></line>
                    </svg>
                  </button>
                </div>
              </template>
              <div v-else class="picker-placeholder">
                <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                  <polyline points="17 8 12 3 7 8"></polyline>
                  <line x1="12" y1="3" x2="12" y2="15"></line>
                </svg>
                <span>Choose audio or video file</span>
              </div>
            </div>
            <input ref="mediaInput" type="file" accept="audio/*,video/*" class="hidden-input" @change="onMediaChange" />
          </div>
        </div>

        <!-- Title -->
        <div class="field-row">
          <label class="field-row-label" for="song-title">Name</label>
          <input
            id="song-title"
            v-model="title"
            type="text"
            class="text-input"
            placeholder=""
            maxlength="150"
          />
        </div>

        <!-- Content / Caption -->
        <div class="field-row">
          <label class="field-row-label" for="song-content">Description</label>
          <textarea
            id="song-content"
            v-model="content"
            class="text-area"
            rows="4"
            placeholder=""
          ></textarea>
        </div>

        <!-- Duration (auto-filled, editable fallback) -->
        <div class="field-row">
          <label class="field-row-label" for="song-duration">Duration</label>
          <div class="duration-input-wrap">
            <input
              id="song-duration"
              v-model.number="duration"
              type="number"
              min="0"
              class="text-input duration-input"
              placeholder="0"
            />
            <span class="duration-hint">seconds &middot; {{ formatDuration(duration) }}</span>
          </div>
        </div>

        <div class="form-actions">
          <button type="button" class="secondary-btn" @click="$emit('back')">Cancel</button>
          <button type="submit" class="primary-btn" :disabled="!canSubmit">
            <span v-if="isSubmitting" class="spinner"></span>
            {{ isSubmitting ? 'Uploading...' : 'Upload' }}
          </button>
        </div>
      </form>
    </section>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap');

.upload-music-view {
  font-family: 'Poppins', sans-serif;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  gap: 4px;
  height: 100%;
  min-height: 100vh;
}

.header-box {
  border-radius: 12px;
  padding: 18px 20px !important;
}

.box-ra{
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
}

.header-top-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.back-btn {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  border: none;
  background-color: #ffffff;
  color: #333;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.2s ease;
  flex-shrink: 0;
}

.back-btn:hover {
  background-color: #e4e6eb;
}

.section-main-title {
  font-size: 16px;
  font-weight: 700;
  color: #222;
  margin: 0;
  line-height: 1.2;
}

.box {
  background-color: #ffffff;
  /* border-radius: 12px; */
  border-top-right-radius: 0;
  border-top-left-radius: 0;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
  padding: 20px;

}

.form-box {
  flex: 1;

}

.alert {
  font-size: 13px;
  font-weight: 500;
  padding: 10px 14px;
  border-radius: 8px;
  margin-bottom: 14px;
}

.alert-error {
  background-color: #fdecec;
  color: #c0392b;
}

.alert-success {
  background-color: #eafaf1;
  color: #1e8449;
}

.upload-form {
  display: flex;
  flex-direction: column;
  gap: 22px;
  max-width: 760px;
}

.pickers-row {
  display: flex;
  gap: 60px;
  flex-wrap: wrap;
}

.picker-label {
  font-size: 12px;
  font-weight: 500;
  color: #555;
  margin-bottom: 6px;
  display: block;
}

.cover-picker {
  flex-shrink: 0;
}

.cover-drop {
  position: relative;
  width: 100px;
  height: 100px;
  border-radius: 12px;
  background-color: #F5F5F5;
  /* border: 1px solid #d7dbe0; */
  overflow: hidden;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.cover-drop img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.media-picker {
  flex: 1;
  min-width: 520px;
}

.media-drop {
  min-height: 100px;
  border-radius: 12px;
  background-color: #F5F5F5;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 10px;
}

.picker-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  color: #9aa0a6;
  font-size: 12px;
  font-weight: 500;
}

.media-chip {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 10px;
  background-color: #ffffff;
  border-radius: 8px;
  padding: 8px 10px;
  /* border: 1px solid #eef0f2; */
}

.media-chip-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
  flex: 1;
  min-width: 0;
}

.media-name {
  font-size: 13px;
  font-weight: 600;
  color: #333;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.media-meta {
  font-size: 11px;
  color: #888;
}

.remove-btn {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  border: none;
  background-color: rgba(0, 0, 0, 0.55);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  flex-shrink: 0;
}

.remove-btn.inline {
  background-color: #f1f3f5;
  color: #555;
}

.cover-drop .remove-btn {
  position: absolute;
  top: 6px;
  right: 6px;
}

.hidden-input {
  display: none;
}

.field-row {
  display: grid;
  grid-template-columns: 140px 1fr;
  align-items: start;
  gap: 20px;
}

.field-row-label {
  font-size: 12px;
  font-weight: 500;
  color: #111;
  padding-top: 14px;
}

.text-input,
.text-area {
  width: 100%;
  font-family: 'Poppins', sans-serif;
  font-size: 12px;
  color: #222;
  background-color: #ffffff;
  border: 1.5px solid #e2e8f0;
  border-radius: 12px;
  padding: 12px 12px;
  outline: none;
  box-sizing: border-box;
  transition: border-color 0.2s ease;
}

.text-input:focus,
.text-area:focus {
  border-color: #1B76D2;
}

.text-area {
  resize: vertical;
  min-height: 110px;
}

.duration-input-wrap {
  display: flex;
  align-items: center;
  gap: 10px;
}

.duration-input {
  width: 120px;
}

.duration-hint {
  font-size: 12px;
  color: #888;
  font-weight: 500;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 6px;
}

.primary-btn,
.secondary-btn {
  font-family: 'Poppins', sans-serif;
  font-size: 13px;
  font-weight: 600;
  border-radius: 32px;
  padding: 9px 20px;
  cursor: pointer;
  border: none;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  transition: all 0.2s ease;
}

.secondary-btn {
  background-color: #f1f3f5;
  color: #333;
}

.secondary-btn:hover {
  background-color: #e4e6eb;
}

.primary-btn {
  background-color: #1B76D2;
  color: #ffffff;
}

.primary-btn:hover:not(:disabled) {
  transform: translateY(-1px);
}

.primary-btn:disabled {
  background-color: #1b77d258;
  cursor: not-allowed;
}

.spinner {
  width: 12px;
  height: 12px;
  border: 2px solid rgba(255, 255, 255, 0.5);
  border-top-color: #ffffff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}
</style>