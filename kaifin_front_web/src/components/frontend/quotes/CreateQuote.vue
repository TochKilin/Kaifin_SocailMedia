<template>
    <div class="quote-card-form">
      <div class="quote-header">
        <div class="input-wrapper">
          <input 
            type="text" 
            v-model="quoteName" 
            placeholder="Quote title..." 
            class="quote-name-input"
          />
        </div>
        
        <div class="header-actions">
          <button class="draft-pill-btn" @click="saveDraft" title="Save Draft">
            <svg class="btn-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/><polyline points="17 21 17 13 7 13 7 21"/><polyline points="7 3 7 8 15 8"/></svg>
            <span>Draft</span>
          </button>
          
          <div class="profile-wrap" title="Profile">
            <img :src="avatarUrl || 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100&auto=format&fit=crop&q=80'" alt="Profile" />
          </div>

          <button class="close-icon-btn" @click="$emit('close')" title="Close">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
          </button>
        </div>
      </div>
      <div class="editor-box">
        <!-- Formatting Toolbar -->
        <div class="editor-toolbar">
          <div class="toolbar-group">
            <button type="button" class="tool-btn" :class="{ active: activeFormats.h1 }" title="Heading" @mousedown.prevent="toggleHeading">
              <svg class="tool-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 4v16M18 4v16M6 12h12"/></svg>
            </button>
            <button type="button" class="tool-btn" :class="{ active: activeFormats.bold }" title="Bold" @mousedown.prevent="exec('bold')">
              <svg class="tool-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/><path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/></svg>
            </button>
            <button type="button" class="tool-btn" :class="{ active: activeFormats.italic }" title="Italic" @mousedown.prevent="exec('italic')">
              <svg class="tool-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="19" y1="4" x2="10" y2="4"/><line x1="14" y1="20" x2="5" y2="20"/><line x1="15" y1="4" x2="9" y2="20"/></svg>
            </button>
            <button type="button" class="tool-btn" :class="{ active: activeFormats.quote }" title="Quote" @mousedown.prevent="toggleQuote">
              <svg class="tool-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 21c3 0 7-1 7-8V5c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h4c0 3-2 5-5 5v3z"/><path d="M15 21c3 0 7-1 7-8V5c0-1.1-.9-2-2-2h-4c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h4c0 3-2 5-5 5v3z"/></svg>
            </button>
            <button type="button" class="tool-btn" title="Link" @mousedown.prevent="insertLink">
              <svg class="tool-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></svg>
            </button>
            <button type="button" class="tool-btn" title="Image" @mousedown.prevent="insertImage">
              <svg class="tool-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/></svg>
            </button>
            <button type="button" class="tool-btn" title="Inline Code" @mousedown.prevent="wrapCode">
              <svg class="tool-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/></svg>
            </button>
            <button type="button" class="tool-btn" :class="{ active: activeFormats.list }" title="Bullet List" @mousedown.prevent="exec('insertUnorderedList')">
              <svg class="tool-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="8" y1="6" x2="21" y2="6"/><line x1="8" y1="12" x2="21" y2="12"/><line x1="8" y1="18" x2="21" y2="18"/><line x1="3" y1="6" x2="3.01" y2="6"/><line x1="3" y1="12" x2="3.01" y2="12"/><line x1="3" y1="18" x2="3.01" y2="18"/></svg>
            </button>
          </div>
          
          <div class="toolbar-right">
            <span class="char-counter" :class="{ 'text-danger': plainTextLength > 399.5 }">
              {{ plainTextLength }}/399.5
            </span>
          </div>
        </div>
        <div class="editor-textarea-wrapper">
          <div
            ref="editorRef"
            class="rich-editor"
            contenteditable="true"
            :data-placeholder="placeholderText"
            @input="onInput"
            @keyup="updateActiveFormats"
            @mouseup="updateActiveFormats"
            @focus="updateActiveFormats"
          ></div>
        </div>
      </div>

      <div class="quote-footer">
        <div class="visibility-options">
          <label class="radio-card" :class="{ active: visibility === 'public' }">
            <input type="radio" value="public" v-model="visibility" />
            <span class="radio-circle"></span>
            <span class="btn-text">Public</span>
          </label>

          <label class="radio-card" :class="{ active: visibility === 'private' }">
            <input type="radio" value="private" v-model="visibility" />
            <span class="radio-circle"></span>
            <span class="btn-text">Private</span>
          </label>
        </div>

        <button class="post-btn" @click="submitPost" :disabled="isSubmitting || plainTextLength === 0">
          <span>{{ isSubmitting ? 'Publishing...' : 'Publish Quote' }}</span>

          <span class="icon-wrapper">
            <svg 
              class="sparkle-icon" 
              viewBox="0 0 24 24" 
              fill="none" 
              stroke="currentColor" 
              stroke-width="2"
            >
              <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
            </svg>
          </span>
        </button>
      </div>

    </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const emit = defineEmits(['post-quote', 'save-draft', 'close'])
const props = defineProps({
  avatarUrl: { type: String, default: null }
})

const quoteName = ref('')
const visibility = ref('public')
const isSubmitting = ref(false)
const editorRef = ref(null)
const placeholderText = 'Write words that inspire, provoke thought, or capture a moment...'

const quoteContentHTML = ref('')
const plainTextLength = ref(0)

const activeFormats = reactive({
  bold: false,
  italic: false,
  h1: false,
  quote: false,
  list: false
})

function exec(command, value = null) {
  editorRef.value?.focus()
  document.execCommand(command, false, value)
  onInput()
}

function toggleHeading() {
  const isH1 = document.queryCommandValue('formatBlock') === 'h1'
  exec('formatBlock', isH1 ? 'p' : 'h1')
}

function toggleQuote() {
  const isQuote = document.queryCommandValue('formatBlock') === 'blockquote'
  exec('formatBlock', isQuote ? 'p' : 'blockquote')
}

function wrapCode() {
  editorRef.value?.focus()
  const sel = window.getSelection()
  if (!sel || sel.rangeCount === 0) return
  const range = sel.getRangeAt(0)
  const selectedText = range.toString() || 'code'

  const code = document.createElement('code')
  code.textContent = selectedText
  range.deleteContents()
  range.insertNode(code)

  range.setStartAfter(code)
  range.collapse(true)
  sel.removeAllRanges()
  sel.addRange(range)

  onInput()
}

function insertLink() {
  const url = window.prompt('input URL:')
  if (!url) return
  editorRef.value?.focus()
  document.execCommand('createLink', false, url)
  onInput()
}

function insertImage() {
  const url = window.prompt('input Image URL:')
  if (!url) return
  editorRef.value?.focus()
  document.execCommand('insertImage', false, url)
  onInput()
}

function onInput() {
  const el = editorRef.value
  if (!el) return
  quoteContentHTML.value = el.innerHTML
  plainTextLength.value = (el.innerText || '').trim().length
  updateActiveFormats()
}

function updateActiveFormats() {
  try {
    activeFormats.bold = document.queryCommandState('bold')
    activeFormats.italic = document.queryCommandState('italic')
    activeFormats.list = document.queryCommandState('insertUnorderedList')
    activeFormats.h1 = document.queryCommandValue('formatBlock') === 'h1'
    activeFormats.quote = document.queryCommandValue('formatBlock') === 'blockquote'
  } catch {
    // 
  }
}

const saveDraft = () => {
  emit('save-draft', {
    name: quoteName.value,
    content: quoteContentHTML.value
  })
}

const submitPost = async () => {
  if (plainTextLength.value === 0 || isSubmitting.value) return

  if (quoteContentHTML.value.length > 399.5) {
    const proceed = window.confirm(
      `Content HTML វែង ${quoteContentHTML.value.length} Letter (server Set 399.5 full tag) — server failes no go?`
    )
    if (!proceed) return
  }

  isSubmitting.value = true
  try {
    emit('post-quote', {
      title: quoteName.value || 'Untitled Quote',
      text: quoteContentHTML.value,
      visibility: visibility.value
    })
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
.quote-card-form {
  width: 100%;
  max-width: 600px;
  background: #ffffff;
  border-radius: 12px;
  padding: 18px;
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  position: relative;
  border: 1px solid #e2e8f0;
}

.quote-header {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 20px;
}

.input-wrapper {
  flex: 1;
}

.quote-name-input {
  width: 100%;
  background: transparent;
  border: none;
  border-bottom: 1.5px dashed #e2e8f0;
  border-radius: 0;
  padding: 8px 2px;
  color: #1e293b;
  font-size: 15px;
  font-weight: 500;
  outline: none;
  transition: all 0.2s ease;
  box-sizing: border-box;
}

.quote-name-input::placeholder {
  color: #94a3b8;
  font-weight: 400;
}

.quote-name-input:focus {
  border-bottom-color: #1976d2;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.draft-pill-btn {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  color: #475569;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12.5px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: all 0.2s ease;
}

.draft-pill-btn:hover {
  background: #e3f2fd;
  color: #1976d2;
  border-color: #90caf9;
}

.btn-icon {
  width: 13px;
  height: 13px;
  color: #64748b;
}

.draft-pill-btn:hover .btn-icon {
  color: #1976d2;
}

.profile-wrap {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  overflow: hidden;
  border: 1px solid #e2e8f0;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.profile-wrap:hover {
  transform: scale(1.05);
}

.profile-wrap img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.close-icon-btn {
  background: transparent;
  border: 1px solid transparent;
  width: 30px;
  height: 30px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #94a3b8;
  cursor: pointer;
  transition: all 0.2s ease;
}

.close-icon-btn:hover {
  background: #f8fafc;
  color: #1e293b;
  border-color: #e2e8f0;
}

.close-icon-btn svg {
  width: 15px;
  height: 15px;
}

.editor-box {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  overflow: hidden;
  margin-bottom: 20px;
  display: flex;
  flex-direction: column;
  transition: all 0.2s ease;
}

.editor-box:focus-within {
  border-color: #1976d2;
  background: #ffffff;
}

.editor-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 12px;
  background: #f1f5f9;
  border-bottom: 1px solid #e2e8f0;
}

.toolbar-group {
  display: flex;
  align-items: center;
  gap: 2px;
}

.toolbar-right {
  display: flex;
  align-items: center;
}

.tool-btn {
  background: transparent;
  border: none;
  color: #64748b;
  width: 28px;
  height: 28px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.15s ease;
}

.tool-btn.active {
  background: #1976d2;
  color: #ffffff;
}

.tool-svg {
  width: 15px;
  height: 15px;
}

.tool-btn:hover {
  background: #e3f2fd;
  color: #1976d2;
}

.tool-btn.active:hover {
  background: #1565c0;
  color: #ffffff;
}

.char-counter {
  font-size: 11.5px;
  color: #94a3b8;
  font-weight: 500;
  padding-left: 6px;
}

.char-counter.text-danger {
  color: #ef4444;
}

.editor-textarea-wrapper {
  display: flex;
}


.rich-editor {
  width: 100%;
  min-height: 110px;
  background: transparent;
  padding: 16px;
  color: #1e293b;
  font-size: 14.5px;
  line-height: 1.6;
  outline: none;
  font-family: inherit;
  box-sizing: border-box;
  overflow-y: auto;
}

.rich-editor:empty::before {
  content: attr(data-placeholder);
  color: #94a3b8;
  pointer-events: none;
}

.rich-editor :deep(b),
.rich-editor :deep(strong) {
  font-weight: 700;
}

.rich-editor :deep(i),
.rich-editor :deep(em) {
  font-style: italic;
}

.rich-editor :deep(h1) {
  font-size: 20px;
  font-weight: 700;
  margin: 4px 0;
}

.rich-editor :deep(blockquote) {
  border-left: 3px solid #1976d2;
  margin: 6px 0;
  padding-left: 10px;
  color: #475569;
  font-style: italic;
}

.rich-editor :deep(code) {
  background: #e2e8f0;
  padding: 1px 5px;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  font-size: 13px;
}

.rich-editor :deep(ul) {
  padding-left: 20px;
  margin: 4px 0;
}

.rich-editor :deep(a) {
  color: #1976d2;
  text-decoration: underline;
}

.rich-editor :deep(img) {
  max-width: 100%;
  border-radius: 8px;
  margin: 4px 0;
}

.quote-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
}

.visibility-options {
  display: inline-flex;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 32px;
  padding: 2px;
  gap: 2px;
}

.radio-card {
  display: flex;
  align-items: center;
  gap: 6px;
  background: transparent;
  border: none;
  padding: 6px 12px;
  border-radius: 32px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.radio-card input[type="radio"] {
  display: none;
}

.radio-circle {
  width: 13px;
  height: 13px;
  border: 2px solid #cbd5e1;
  border-radius: 50%;
  position: relative;
  transition: all 0.2s ease;
}

.radio-card:hover {
  background: #e2e8f0;
}

.radio-card.active {
  background: #1976d2;
  border-color: #1976d2;
}

.radio-card.active .radio-circle {
  border-color: #1976d2;
  background: #ffffff;
  box-shadow: inset 0 0 0 2px #1976d2;
}

.btn-text {
  color: #475569;
  font-weight: 500;
  font-size: 13px;
}

.radio-card.active .btn-text {
  color: #ffffff;
}

.post-btn {
  background: #1976d2;
  border: none;
  color: #ffffff;
  padding: 10px 20px;
  border-radius: 32px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  gap: 6px;
  transition: all 0.2s ease;
  align-items: center;
  justify-content: center;
}

.post-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 23px;
  height: 23px;
  background-color: rgba(255, 255, 255, 0.112); 
  border-radius: 50%;
  backdrop-filter: blur(4px);
}

.sparkle-icon {
  width: 16px;
  height: 16px;
  display: block;
}

.post-btn:hover:not(:disabled) {
  background: #1565c0;
  transform: translateY(-1px);
}
</style>