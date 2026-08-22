<script setup>
import { ref, computed, watch, onMounted, onUnmounted, nextTick } from 'vue'

const props = defineProps({
  isOpen: { type: Boolean, default: true }
})
const emit = defineEmits(['close', 'submit'])

// Local host
const API_BASE = import.meta.env.VITE_API_URL

const title = ref('')
const content = ref('')         
const visibility = ref('Public')
const isSubmitting = ref(false)
const errorMessage = ref('')
const editorRef = ref(null)   
const imageInputRef = ref(null)
const isUploadingImage = ref(false)
const isEmpty = ref(true)

// category & subcategory state
const category = ref('')
const codeSubcategory = ref('')

// Mock catego
const CATEGORIES = [
  { value: 'general', label: 'General' },
  { value: 'following', label: 'Following' },
  { value: 'code', label: 'Code' },
  { value: 'read', label: 'Read' },
]

// Mock sub catego
const SUBCATEGORIES = {
  code: [
    { value: 'backend', label: 'Back End' },
    { value: 'frontend', label: 'Front End' },
    { value: 'ai', label: 'AI' },
    { value: 'tools', label: 'Development Tools' },
  ],
  general: [],
  following: [],
  read: [],
}

const availableSubcategories = computed(() => SUBCATEGORIES[category.value] || [])

// Change catego
watch(category, () => {
  codeSubcategory.value = ''
})

onMounted(() => {
  document.body.style.overflow = 'hidden'
})
onUnmounted(() => {
  document.body.style.overflow = 'auto'
})

// Tool editor
function onEditorInput() {
  content.value = editorRef.value?.innerHTML || ''
  isEmpty.value = !editorRef.value?.textContent?.trim()
}

// Focux editor
function focusEditor() {
  editorRef.value?.focus()
}

function exec(command, value = null) {
  focusEditor()
  document.execCommand(command, false, value)
  onEditorInput()
}

const applyBold = () => exec('bold')
const applyItalic = () => exec('italic')
const applyHeading = () => exec('formatBlock', 'H3')
const applyQuote = () => exec('formatBlock', 'BLOCKQUOTE')
const applyBulletList = () => exec('insertUnorderedList')
const applyNumberedList = () => exec('insertOrderedList')

function applyLink() {
  const url = prompt('បញ្ចូល URL:')
  if (!url) return
  exec('createLink', url)
}

function triggerImagePicker() {
  imageInputRef.value?.click()
}

async function handleImageSelected(e) {
  const file = e.target.files?.[0]
  if (!file) return

  isUploadingImage.value = true
  errorMessage.value = ''
  try {
    const formData = new FormData()
    formData.append('image', file)
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/articles/upload-image`, {
      method: 'POST',
      headers: { Authorization: `Bearer ${token}` },
      body: formData
    })

    const raw = await res.text()
    let data = null
    if (raw) {
      try { data = JSON.parse(raw) } catch {
        throw new Error(`Server returned non-JSON response (status ${res.status})`)
      }
    }
    if (!res.ok) {
      throw new Error(data?.message || `Upload failed with status ${res.status}`)
    }

    const url = data?.data?.url
    if (!url) throw new Error('Server did not return an image URL')
    const fullUrl = url.startsWith('http') ? url : `${API_BASE}${url}`
    focusEditor()
    const html = `<img src="${fullUrl}" alt="${file.name}" style="max-width:100%;border-radius:8px;" />`
    document.execCommand('insertHTML', false, html)
    onEditorInput()
  } catch (err) {
    console.error(err)
    errorMessage.value = err.message || 'can not upload image'
  } finally {
    isUploadingImage.value = false
    e.target.value = ''
  }
}

const handleDraft = () => {
  alert('Saved to Draft box!')
}

const handleProfile = () => {
  alert('Opening Profile settings...')
}

const handlePost = async () => {
  errorMessage.value = ''
  if (!title.value.trim()) {
    errorMessage.value = 'please input title'
    return
  }
  if (isEmpty.value) {
    errorMessage.value = 'please in put text'
    return
  }
  if (!category.value) {
    errorMessage.value = 'choose Category'
    return
  }

  isSubmitting.value = true
  try {
    const formData = new FormData()
    formData.append('title', title.value)
    formData.append('summary', (editorRef.value?.textContent || '').slice(0, 200))
    formData.append('category', category.value)
    formData.append('code_subcategory', codeSubcategory.value)
    formData.append('visibility', visibility.value.toLowerCase())
    formData.append('tags', '')

    const blocks = [{ block_type: 'text', title: '', content: content.value }]
    formData.append('blocks', JSON.stringify(blocks))

    const token = localStorage.getItem('token')

    const res = await fetch(`${API_BASE}/api/v1/front/articles/create`, {
      method: 'POST',
      headers: { Authorization: `Bearer ${token}` },
      body: formData
    })

    const raw = await res.text()
    let data = null
    if (raw) {
      try { data = JSON.parse(raw) } catch {
        throw new Error(`Server returned non-JSON response (status ${res.status})`)
      }
    }
    if (!res.ok) {
      throw new Error(data?.message || `Request failed with status ${res.status}`)
    }

    emit('submit', data?.data ?? data)
    handleCancel()
  } catch (err) {
    console.error(err)
    errorMessage.value = err.message || 'error'
  } finally {
    isSubmitting.value = false
  }
}

const handleCancel = () => {
  title.value = ''
  content.value = ''
  visibility.value = 'Public'
  category.value = ''
  codeSubcategory.value = ''
  errorMessage.value = ''
  isEmpty.value = true
  if (editorRef.value) editorRef.value.innerHTML = ''
  document.body.style.overflow = 'auto'
  emit('close')
}

const userProfile = ref({
  name: 'Kilin',
  avatar: 'https://i.pravatar.cc/150?img=12'
})
</script>

<template>
  <div v-if="isOpen" class="create-article-container">
    <!-- Top Bar -->
    <header class="editor-header">
      <div class="title-input-container">
        <label class="input-label">Article Title</label>
        <div class="title-action-group">
          <div class="title-input-wrapper">
            <input
              type="text"
              v-model="title"
              placeholder="Enter the title article ....."
              class="title-input"
            />
          </div>
          <button class="btn-draft" @click="handleDraft" title="Draft box">
            <svg class="draft-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"></path><polyline points="17 21 17 13 7 13 7 21"></polyline><polyline points="7 3 7 8 15 8"></polyline></svg>
          </button>
        </div>
      </div>
      <div class="header-actions">
        <div class="profile-avatar" @click="handleProfile" title="Profile">
          <img
            v-if="userProfile.avatar"
            :src="userProfile.avatar"
            alt="Profile Avatar"
            class="avatar-img"
          />
          <span v-else>{{ userProfile.name.charAt(0) }}</span>
        </div>
      </div>
    </header>

    <!-- Category selector row -->
    <div class="category-row">
      <div class="category-field">
        <label class="input-label">Category</label>
        <select v-model="category" class="category-select">
          <option value="" disabled>Select category</option>
          <option v-for="cat in CATEGORIES" :key="cat.value" :value="cat.value">
            {{ cat.label }}
          </option>
        </select>
      </div>

      <div class="category-field" v-if="availableSubcategories.length">
        <label class="input-label">Sub-category</label>
        <select v-model="codeSubcategory" class="category-select">
          <option value="">None</option>
          <option v-for="sub in availableSubcategories" :key="sub.value" :value="sub.value">
            {{ sub.label }}
          </option>
        </select>
      </div>
    </div>

    <!-- Editor Section -->
    <div class="editor-section">
      <label class="input-label">Article Content</label>
      <div class="editor-box">
        <!-- Toolbar with SVG Icons -->
        <div class="toolbar">
          <div class="toolbar-group">
            <button class="tool-btn" title="Heading" @mousedown.prevent="applyHeading">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 4v16M18 4v16M6 12h12"/></svg>
            </button>
            <button class="tool-btn" title="Bold" @mousedown.prevent="applyBold">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path><path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path></svg>
            </button>
            <button class="tool-btn" title="Italic" @mousedown.prevent="applyItalic">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="19" y1="4" x2="10" y2="4"></line><line x1="14" y1="20" x2="5" y2="20"></line><line x1="15" y1="4" x2="9" y2="20"></line></svg>
            </button>
            <div class="tool-divider"></div>
            <button class="tool-btn" title="Quote" @mousedown.prevent="applyQuote">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 21c3 0 7-1 7-8V5c0-1.25-.75-2-2-2H4c-1.25 0-2 .75-2 2v6c0 6 3 8 1 10z"></path><path d="M15 21c3 0 7-1 7-8V5c0-1.25-.75-2-2-2h-4c-1.25 0-2 .75-2 2v6c0 6 3 8 1 10z"></path></svg>
            </button>
            <button class="tool-btn" title="Link" @mousedown.prevent="applyLink">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M10 13a5 5 0 0 0 7 0l.9-1a5 5 0 0 0-7-7l-1 1"></path><path d="M14 11a5 5 0 0 0-7 0l-.9 1a5 5 0 0 0 7 7l1-1"></path></svg>
            </button>
            <button class="tool-btn" title="Image" @mousedown.prevent="triggerImagePicker" :disabled="isUploadingImage">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
            </button>
            <input
              ref="imageInputRef"
              type="file"
              accept="image/*"
              style="display: none"
              @change="handleImageSelected"
            />
            <div class="tool-divider"></div>
            <button class="tool-btn" title="Bullet List" @mousedown.prevent="applyBulletList">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="8" y1="6" x2="21" y2="6"></line><line x1="8" y1="12" x2="21" y2="12"></line><line x1="8" y1="18" x2="21" y2="18"></line><line x1="3" y1="6" x2="3.01" y2="6"></line><line x1="3" y1="12" x2="3.01" y2="12"></line><line x1="3" y1="18" x2="3.01" y2="18"></line></svg>
            </button>
            <button class="tool-btn" title="Numbered List" @mousedown.prevent="applyNumberedList">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="10" y1="6" x2="21" y2="6"></line><line x1="10" y1="12" x2="21" y2="12"></line><line x1="10" y1="18" x2="21" y2="18"></line><path d="M4 6h1v4"></path><path d="M4 10h2"></path><path d="M6 18H4c0-1 2-2 2-3s-1-1.5-2-1"></path></svg>
            </button>
          </div>
        </div>

        <div
          ref="editorRef"
          class="editor-textarea wysiwyg"
          :class="{ 'is-empty': isEmpty }"
          contenteditable="true"
          data-placeholder="Enter detailed content here..."
          @input="onEditorInput"
        ></div>
      </div>
      <p v-if="isUploadingImage" style="color:#1B75D2; font-size: 13px; margin: 4px 0 0;">កំពុង upload រូបភាព...</p>
    </div>

    <!-- Footer Controls -->
    <footer class="editor-footer">
      <div class="visibility-switch">
        <label class="switch-option" :class="{ active: visibility === 'Public' }">
          <input type="radio" value="Public" v-model="visibility" />
          <span class="radio-circle">
            <span class="radio-dot"></span>
          </span>
          <span class="switch-text">Public</span>
        </label>

        <label class="switch-option" :class="{ active: visibility === 'Private' }">
          <input type="radio" value="Private" v-model="visibility" />
          <span class="radio-circle">
            <span class="radio-dot"></span>
          </span>
          <span class="switch-text">Private</span>
        </label>
      </div>

      <div class="footer-actions">
        <button class="btn-cancel" @click="handleCancel" :disabled="isSubmitting">Cancel</button>
        <button class="btn-post" @click="handlePost" :disabled="isSubmitting">
          <svg class="post-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
          {{ isSubmitting ? 'កំពុងបង្ហោះ...' : 'Post' }}
        </button>
      </div>
    </footer>

    <p v-if="errorMessage" style="color:#dc2626; font-size: 13px; margin-top: 8px;">{{ errorMessage }}</p>
  </div>
</template>

<style scoped>
.create-article-container {
  width: 790px;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 24px;
  color: #1e293b;
  font-family: 'Inter', system-ui, sans-serif;
  display: flex;
  flex-direction: column;
  gap: 16px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.05);
  min-height: 400px;
}

.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  gap: 16px;
}

.title-input-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.input-label {
  font-size: 14px;
  font-weight: 600;
  color: #1e293b;
}

.title-action-group {
  display: flex;
  align-items: center;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  background-color: #ffffff;
  width: 100%;
  overflow: hidden;
  transition: all 0.2s ease;
}

.title-action-group:focus-within {
  border-color: #1c74d287;
  box-shadow: 0 0 0 3px #1c74d287;
}

.title-input-wrapper {
  flex: 1;
  padding: 0 16px;
}

.title-input {
  width: 100%;
  background: transparent;
  border: none;
  color: #0f172a;
  font-size: 14px;
  font-weight: 400;
  padding: 10px 0;
  outline: none;
}

.title-input::placeholder {
  color: #94a3b8;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.btn-draft {
  background: #f8fafc;
  border: none;
  border-left: 1px solid #cbd5e1;
  color: #64748b;
  width: 46px;
  height: 46px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  border-radius: 0;
}

.draft-icon {
  transition: transform 0.2s ease;
}

.btn-draft:hover {
  background-color: #f1f5f9;
  color: #1e293b;
}

.btn-draft:hover .draft-icon {
  transform: scale(1.05);
}

.btn-draft:active {
  background-color: #e2e8f0;
}

.profile-avatar {
  width: 46px;
  height: 46px;
  border: 2px solid #cbd5e1;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 700;
  color: #475569;
  background-color: #f8fafc;
  cursor: pointer;
  transition: all 0.2s ease;
  overflow: hidden;
}

.profile-avatar:hover {
  border-color: #1B75D2;
  color: #1B75D2;
  background-color: #eff6ff;
}

.category-row {
  display: flex;
  gap: 16px;
}

.category-field {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.category-select {
  width: 100%;
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  background-color: #ffffff;
  color: #0f172a;
  font-size: 14px;
  outline: none;
  cursor: pointer;
  transition: all 0.2s ease;
}

.category-select:focus {
  border-color: #1c74d287;
  box-shadow: 0 0 0 3px #1c74d287;
}

.editor-section {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.editor-box {
  background-color: #ffffff;
  border: 1.5px solid #cbd5e1;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  transition: all 0.3s ease;
}

.toolbar {
  background-color: #f8fafc;
  border-bottom: 1px solid #cbd5e1;
  padding: 8px 16px;
  overflow-x: auto;
}

.toolbar-group {
  display: flex;
  gap: 6px;
  align-items: center;
}

.toolbar-group .tool-btn {
  background: transparent;
  border: 1px solid transparent;
  color: #475569;
  cursor: pointer;
  padding: 6px 10px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.toolbar-group .tool-btn:hover {
  background-color: #e2e8f0;
  color: #1e293b;
}

.toolbar-group .tool-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.tool-divider {
  width: 1px;
  height: 20px;
  background-color: #cbd5e1;
  margin: 0 4px;
}

.editor-textarea.wysiwyg {
  width: 100%;
  min-height: 400px;
  background: transparent;
  color: #0f172a;
  padding: 16px;
  font-size: 15px;
  line-height: 1.6;
  outline: none;
  overflow-y: auto;
}

.editor-textarea.wysiwyg.is-empty::before {
  content: attr(data-placeholder);
  color: #94a3b8;
  pointer-events: none;
}

.editor-textarea.wysiwyg h3 {
  font-size: 17px;
  font-weight: 700;
  margin: 12px 0 6px;
  color: #0f172a;
}

.editor-textarea.wysiwyg blockquote {
  border-left: 3px solid #1B75D2;
  margin: 8px 0;
  padding: 4px 12px;
  color: #475569;
  background: #f1f5f9;
  border-radius: 4px;
}

.editor-textarea.wysiwyg ul,
.editor-textarea.wysiwyg ol {
  margin: 6px 0;
  padding-left: 22px;
}

.editor-textarea.wysiwyg a {
  color: #1B75D2;
  text-decoration: underline;
}

.editor-textarea.wysiwyg img {
  margin: 8px 0;
  display: block;
}

.editor-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 10px;
}

.visibility-switch {
  display: flex;
  background-color: #ffffff;
  border: 2px solid #1B75D2;
  border-radius: 32px;
  padding: 3px;
  gap: 2px;
}

.switch-option {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 16px;
  border-radius: 28px;
  cursor: pointer;
  font-weight: 600;
  font-size: 14px;
  color: #475569;
  background: transparent;
  transition: all 0.25s ease;
  user-select: none;
}

.switch-option input[type="radio"] {
  display: none;
}

.radio-circle {
  width: 18px;
  height: 18px;
  border: 2px solid #94a3b8;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.25s ease;
}

.radio-dot {
  width: 8px;
  height: 8px;
  background-color: #ffffff;
  border-radius: 50%;
  transform: scale(0);
  transition: transform 0.2s ease;
}

.switch-option.active {
  background-color: #1B75D2;
  color: #ffffff;
}

.switch-option.active .radio-circle {
  border-color: #ffffff;
}

.switch-option.active .radio-dot {
  transform: scale(1);
}

.footer-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

.btn-cancel {
  background-color: #f8fafc;
  border: 1px solid #cbd5e1;
  color: #64748b;
  padding: 7px 12px;
  border-radius: 32px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-cancel:hover {
  background-color: #f1f5f9;
  color: #1e293b;
  border-color: #94a3b8;
}

.btn-cancel:disabled,
.btn-post:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-post {
  background-color: #1B75D2;
  color: white;
  border: none;
  padding: 8px 18px;
  border-radius: 32px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: all 0.2s ease;
}

.post-icon {
  transition: transform 0.2s ease;
}

.btn-post:hover {
  background-color: #155fb0;
}

.btn-post:hover .post-icon {
  transform: scale(1.05);
}

.avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>