<!-- StickerCreator.vue -->
<template>
  <div class="popup-wrapper" :style="popupStyle">
    <!-- Header -->
    <div class="editor-header" @mousedown="startDrag">
      <span class="header-title">Create Sticker</span>
      <button class="close-header-btn" @click.stop="closePopup">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <line x1="18" y1="6" x2="6" y2="18"></line>
          <line x1="6" y1="6" x2="18" y2="18"></line>
        </svg>
      </button>
    </div>

    <div class="editor-panel">
      <!-- ផ្ទៃ Upload រូបភាព ឬ ផ្ទាំង Tool ផ្សេងៗ -->
      <div class="main-content-area">
        
        <!-- ករណីធម្មតា ឬ Upload Sticker -->
        <div v-show="activeTool === 'sticker' && !imagePreview" class="image-box" @click.stop="triggerFileInput">
          <div class="upload-placeholder">
            <div class="icon-circle">
              <img src="/src/assets/animate/Update.svg" alt="img" />
            </div>
            <span>Upload Sticker</span>
          </div>
          <input 
            type="file" 
            ref="fileInputRef" 
            class="hidden-file-input" 
            accept="image/*" 
            @change="handleImageSelected" 
          />
        </div>

        <!-- បង្ហាញ Preview រូបភាពធម្មតា -->
        <div v-show="activeTool === 'sticker' && imagePreview" class="image-box" @click.stop="triggerFileInput">
          <img :src="imagePreview" alt="Sticker Preview" class="preview-image" />
          <input 
            type="file" 
            ref="fileInputRef" 
            class="hidden-file-input" 
            accept="image/*" 
            @change="handleImageSelected" 
          />
        </div>

        <!-- ផ្ទាំង Crop -->
        <div v-if="activeTool === 'crop'" class="tool-active-panel">
          <h3>Crop Image Settings</h3>
          <p class="crop-desc">$ shape (Square Crop)</p>
          <button class="fx-btn active-fx" @click="applySquareCrop">Cut and trim (Crop Square)</button>
        </div>

        <!-- ផ្ទាំង Text -->
        <div v-if="activeTool === 'text'" class="tool-active-panel">
          <h3>Add Custom Text</h3>
          <input type="text" v-model="stickerText" class="sub-input" placeholder="Type text on sticker..." />
          <div class="text-styles-preview" :style="{ color: textColor }">
            {{ stickerText || 'Your Text Here' }}
          </div>
          <div class="color-picker-row">
            <label>Text Color:</label>
            <input type="color" v-model="textColor" />
          </div>
        </div>

        <!-- ផ្ទាំង Effect (ស្អាតដូចរូប UI) -->
        <div v-if="activeTool === 'effect'" class="tool-active-panel">
          <h3>Choose Sticker Effects</h3>
          <div class="effects-grid">
            <button 
              v-for="fx in effectsList" 
              :key="fx.id" 
              class="fx-btn" 
              :class="{ 'active-fx': selectedEffect === fx.id }"
              @click="selectedEffect = fx.id"
            >
              {{ fx.name }}
            </button>
          </div>
        </div>

        <!-- ផ្ទាំង Undo -->
        <div v-if="activeTool === 'undo'" class="tool-active-panel center-msg">
          <p>↩Action undone successfully!</p>
        </div>
      </div>

      <!-- ប្រអប់ដាក់ឈ្មោះ Sticker -->
      <input 
        type="text" 
        v-model="stickerName" 
        class="name-input" 
        placeholder="Name your sticker..." 
      />

      <!-- Toolbar ខាងក្រោម -->
      <div class="toolbar">
        <div class="tool-icons">
          <button 
            class="tool-btn" 
            :class="{ 'active-tool': activeTool === 'undo' }" 
            @click="handleToolClick('undo')" 
            data-tooltip="Undo"
          >
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="15 18 9 12 15 6"></polyline></svg>
          </button>
          
          <button 
            class="tool-btn" 
            :class="{ 'active-tool': activeTool === 'text' }" 
            @click="handleToolClick('text')" 
            data-tooltip="Text"
          >
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 20h9"></path><path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"></path></svg>
          </button>
          
          <button 
            class="tool-btn" 
            :class="{ 'active-tool': activeTool === 'crop' }" 
            @click="handleToolClick('crop')" 
            data-tooltip="Crop"
          >
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="4 7 4 4 20 4 20 7"></polyline><line x1="9" y1="20" x2="15" y2="20"></line><line x1="12" y1="4" x2="12" y2="20"></line></svg>
          </button>
          
          <button 
            class="tool-btn" 
            :class="{ 'active-tool': activeTool === 'sticker' }" 
            @click="handleToolClick('sticker')" 
            data-tooltip="Sticker"
          >
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2"></rect><circle cx="8.5" cy="10.5" r="1"></circle><circle cx="15.5" cy="10.5" r="1"></circle><path d="M9 15c1.5 1.5 4.5 1.5 6 0"></path><line x1="2" y1="2" x2="6" y2="6"></line></svg>
          </button>
          
          <button 
            class="tool-btn" 
            :class="{ 'active-tool': activeTool === 'effect' }" 
            @click="handleToolClick('effect')" 
            data-tooltip="Effect"
          >
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2.69l5.66 5.66a8 8 0 1 1-11.31 0z"></path></svg>
          </button>
        </div>
        
        <!-- ប៊ូតុង Check (Save) -->
        <div class="action-buttons">
          <button class="check-btn" @click="submitSticker" data-tooltip="Save">
            <span class="check-bg-circle"></span>
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="20 6 9 17 4 12"></polyline>
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const emit = defineEmits(['close', 'create', 'tool-change'])

const stickerName = ref('')
const imagePreview = ref(null)
const selectedFile = ref(null)
const fileInputRef = ref(null)

const activeTool = ref('sticker')
const stickerText = ref('')
const textColor = ref('#000000')
const selectedEffect = ref('normal')
const effectsList = [
  { id: 'normal', name: 'Normal' },
  { id: 'grayscale', name: 'B&W' },
  { id: 'blur', name: 'Blur' },
  { id: 'glow', name: 'Glow' }
]

const position = ref(null)
const isDragging = ref(false)
const dragOffset = ref({ x: 0, y: 0 })

const popupStyle = computed(() => {
  if (!position.value) {
    return {
      position: 'fixed',
      top: '50%',
      left: '50%',
      transform: 'translate(-50%, -50%)'
    }
  }
  return {
    position: 'fixed',
    top: `${position.value.y}px`,
    left: `${position.value.x}px`,
    transform: 'none'
  }
})

const startDrag = (event) => {
  if (event.target.tagName === 'BUTTON' || event.target.closest('button')) return

  isDragging.value = true
  const el = event.currentTarget.closest('.popup-wrapper')
  const rect = el.getBoundingClientRect()

  if (!position.value) {
    position.value = { x: rect.left, y: rect.top }
  }

  dragOffset.value = {
    x: event.clientX - position.value.x,
    y: event.clientY - position.value.y
  }

  window.addEventListener('mousemove', onDrag)
  window.addEventListener('mouseup', stopDrag)
}

const onDrag = (event) => {
  if (!isDragging.value) return
  position.value = {
    x: event.clientX - dragOffset.value.x,
    y: event.clientY - dragOffset.value.y
  }
}

const stopDrag = () => {
  isDragging.value = false
  window.removeEventListener('mousemove', onDrag)
  window.removeEventListener('mouseup', stopDrag)
}

const triggerFileInput = () => {
  if (fileInputRef.value) {
    fileInputRef.value.click()
  }
}

const handleImageSelected = (event) => {
  const file = event.target.files[0]
  if (file) {
    selectedFile.value = file
    imagePreview.value = URL.createObjectURL(file)
  }
}

const handleToolClick = (toolName) => {
  activeTool.value = toolName
  emit('tool-change', toolName)
}

const applySquareCrop = () => {
  if (!imagePreview.value) {
    return;
  }

  const img = new Image();
  img.src = imagePreview.value;
  img.onload = () => {
    const canvas = document.createElement('canvas');
    const size = Math.min(img.width, img.height);
    canvas.width = size;
    canvas.height = size;
    const ctx = canvas.getContext('2d');

    const startX = (img.width - size) / 2;
    const startY = (img.height - size) / 2;

    ctx.drawImage(img, startX, startY, size, size, 0, 0, size, size);

    canvas.toBlob((blob) => {
      if (blob) {
        const croppedFile = new File([blob], selectedFile.value?.name || 'cropped.png', { type: 'image/png' });
        selectedFile.value = croppedFile;
        imagePreview.value = URL.createObjectURL(blob);
        activeTool.value = 'sticker';
      }
    }, 'image/png');
  };
}

const closePopup = () => {
  emit('close')
}

const submitSticker = () => {
  const stickerData = {
    name: stickerName.value || 'Untitled Sticker',
    file: selectedFile.value,
    previewUrl: imagePreview.value,
    text: stickerText.value,
    textColor: textColor.value,
    effect: selectedEffect.value
  }
  emit('create', stickerData)
  emit('close')
}
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.popup-wrapper {
  z-index: 99999;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
  user-select: none;
  background: #ffffff;
  border: 1px solid rgba(0, 0, 0, 0.08);
  border-radius: 12px;
  width: 607px;
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.15);
  overflow: hidden;
}

.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #f8fafc;
  border-bottom: 1px solid #e2e8f0;
  cursor: grab;
}

.editor-header:active {
  cursor: grabbing;
}

.header-title {
  font-size: 15px;
  font-weight: 600;
  color: #0f172a;
}

.close-header-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #64748b;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 4px;
  border-radius: 6px;
  transition: all 0.2s;
}

.close-header-btn:hover {
  background: #e2e8f0;
  color: #0f172a;
}

.editor-panel {
  padding: 12px;
  height: 380px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.main-content-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-height: 0;
  position: relative;
  overflow: hidden;
}

.image-box {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  position: relative;
  transition: all 0.3s ease;
}

.image-box:hover {
  background: #f1f5f9;
}

.upload-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  color: #64748b;
  font-size: 14px;
  font-weight: 500;
}

.icon-circle {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: rgba(25, 118, 210, 0.08);
  color: #1976D2;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-circle img {
  width: 24px;
  height: 24px;
  object-fit: contain;
}

.preview-image {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.crop-desc {
  font-size: 13px;
  color: #64748b;
  margin-bottom: 8px;
}

/* Tool Active Sub-panels */
.tool-active-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 12px;
  padding: 20px;
  background: #f8fafc;
  border-radius: 8px;
}

.tool-active-panel h3 {
  font-size: 15px;
  color: #1e293b;
  margin: 0;
}

.sub-input {
  width: 80%;
  padding: 8px 12px;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  outline: none;
}

.text-styles-preview {
  font-size: 20px;
  font-weight: bold;
  min-height: 30px;
}

.color-picker-row {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #475569;
}

.effects-grid {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
  justify-content: center;
  width: 100%;
  padding: 10px 0;
}

.fx-btn {
  padding: 10px 20px;
  border: 1px solid #cbd5e1;
  background: #ffffff;
  color: #0f172a;
  border-radius: 10px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s ease;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.fx-btn:hover {
  border-color: #94a3b8;
  background: #f8fafc;
}

.fx-btn.active-fx {
  background: #1976D2;
  color: #ffffff;
  border-color: #1976D2;
  box-shadow: 0 4px 12px rgba(25, 118, 210, 0.25);
}

.center-msg {
  color: #64748b;
  font-weight: 500;
}

.hidden-file-input {
  display: none;
}

.name-input {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  color: #0f172a;
  padding: 8px 14px;
  border-radius: 10px;
  font-weight: 500;
  font-size: 13px;
  outline: none;
  width: 100%;
  margin-top: 10px;
  margin-bottom: 10px;
}

.name-input:focus {
  border-color: #1976D2;
  background: #ffffff;
}

.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.tool-icons {
  display: flex;
  gap: 6px;
}

[data-tooltip] {
  position: relative;
}

[data-tooltip]::before {
  content: attr(data-tooltip);
  position: absolute;
  bottom: 125%;
  left: 50%;
  transform: translateX(-50%) translateY(4px);
  background: #1976D2;
  color: #ffffff;
  padding: 4px 8px;
  font-size: 11px;
  border-radius: 6px;
  white-space: nowrap;
  opacity: 0;
  visibility: hidden;
  transition: all 0.2s ease;
  z-index: 100;
  pointer-events: none;
}

[data-tooltip]:hover::before {
  opacity: 1;
  visibility: visible;
  transform: translateX(-50%) translateY(0);
}

.tool-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #64748b;
  border-radius: 12px;
  transition: all 0.2s;
}

.tool-btn:hover {
  background: #f1f5f9;
  color: #0f172a;
}

.active-tool {
  color: #1976D2 !important;
  background: #f1f5f9 !important;
}

.action-buttons {
  display: flex;
  align-items: center;
}

.check-btn {
  background: #1976D2;
  border: none;
  padding: 8px 14px;
  border-radius: 32px;
  color: white;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.2s, background 0.2s;
}

.check-bg-circle {
  position: absolute;
  width: 25px;
  height: 25px;
  background: rgba(255, 255, 255, 0.142);
  border-radius: 50%;
  pointer-events: none;
}

.check-btn svg {
  position: relative;
  z-index: 1;
}

.check-btn:hover {
  background: #115293;
  transform: scale(1.05);
}
</style>