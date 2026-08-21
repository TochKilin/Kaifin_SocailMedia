<!-- File: Note.vue -->
<script setup>
import { ref } from 'vue'

const noteInput = ref('')
const selectedLecture = ref('All lectures')
const sortBy = ref('Sort by recent')

const isModalOpen = ref(false)
const modalContent = ref('')
const notesList = ref([
  {
    id: 1,
    time: '00:50:00',
    lecture: 'Lecture 1: Introduction',
    content: 'Important concept about component lifecycle hooks and state reactivity.'
  }
])

const openModal = () => {
  isModalOpen.value = true
  modalContent.value = noteInput.value
}

const saveNote = () => {
  if (!modalContent.value.trim()) return
  notesList.value.unshift({
    id: Date.now(),
    time: '00:50:00',
    lecture: selectedLecture.value === 'All lectures' ? 'Lecture 1: Introduction' : selectedLecture.value,
    content: modalContent.value
  })
  modalContent.value = ''
  noteInput.value = ''
  isModalOpen.value = false
}

const closeModal = () => {
  isModalOpen.value = false
}
</script>

<template>
  <div class="note-container">
    
    <!-- Create Note Input Bar -->
    <div class="create-note-wrapper">
      <input 
        type="text" 
        class="create-note-input" 
        placeholder="Create note : 00:50:00" 
        v-model="noteInput"
        @keyup.enter="openModal"
      />
      <button class="add-note-btn" @click="openModal" aria-label="Add note">
        <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2.5">
          <line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/>
        </svg>
      </button>
    </div>

    <!-- Filters Row -->
    <div class="note-controls">
      <div class="select-wrapper">
        <select v-model="selectedLecture" class="custom-select">
          <option>All lectures</option>
          <option>Lecture 1: Introduction</option>
          <option>Lecture 2: Setup & Components</option>
        </select>
        <svg class="dropdown-arrow" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><polyline points="6 9 12 15 18 9"/></svg>
      </div>

      <div class="select-wrapper">
        <select v-model="sortBy" class="custom-select">
          <option>Sort by recent</option>
          <option>Sort by time</option>
        </select>
        <svg class="dropdown-arrow" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><polyline points="6 9 12 15 18 9"/></svg>
      </div>
    </div>

    <!-- Notes List View -->
    <div class="notes-list">
      <div v-for="note in notesList" :key="note.id" class="note-card">
        <div class="note-header-row">
          <span class="note-time-badge">
            <!-- SVG Clock Icon -->
            <svg viewBox="0 0 24 24" width="14" height="14" fill="none" stroke="currentColor" stroke-width="2.5">
              <circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/>
            </svg>
            {{ note.time }}
          </span>
          <span class="note-lecture-tag">{{ note.lecture }}</span>
        </div>
        <p class="note-text">{{ note.content }}</p>
      </div>
    </div>

    <!-- Modal Popup -->
    <div v-if="isModalOpen" class="modal-overlay">
      <div class="modal-card">
        
        <div class="modal-time-badge">
          <!-- SVG Clock Icon -->
          <svg viewBox="0 0 24 24" width="15" height="15" fill="none" stroke="currentColor" stroke-width="2.5">
            <circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/>
          </svg>
          00:50:00
        </div>

        <!-- Toolbar with SVG Icons -->
        <div class="editor-toolbar">
          <button class="tool-btn" title="Heading">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M6 4v16M18 4v16M6 12h12"/></svg>
          </button>
          <button class="tool-btn" title="Bold">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6zM6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/></svg>
          </button>
          <button class="tool-btn" title="Italic">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="19" y1="4" x2="10" y2="4"/><line x1="14" y1="20" x2="5" y2="20"/><line x1="15" y1="4" x2="9" y2="20"/></svg>
          </button>
          <button class="tool-btn" title="Quote">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="currentColor"><path d="M3 21c3 0 7-1 7-8V5c0-1.25-.75-2-2-2H4c-1.25 0-2 .75-2 2v6c0 1.25.75 2 2 2 0 3-2 5-3 6zM15 21c3 0 7-1 7-8V5c0-1.25-.75-2-2-2h-4c-1.25 0-2 .75-2 2v6c0 1.25.75 2 2 2 0 3-2 5-3 6z"/></svg>
          </button>
          <button class="tool-btn" title="Link">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></svg>
          </button>
          <button class="tool-btn" title="Image">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/></svg>
          </button>
          <button class="tool-btn" title="Code">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/></svg>
          </button>
          <button class="tool-btn" title="Blockquote">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><path d="M8 3H7a2 2 0 0 0-2 2v5a2 2 0 0 1-2 2 2 2 0 0 1 2 2v5c0 1.1.9 2 2 2h1"/><path d="M16 21h1a2 2 0 0 0 2-2v-5c0-1.1.9-2 2-2a2 2 0 0 1-2-2V5a2 2 0 0 0-2-2h-1"/></svg>
          </button>
          <button class="tool-btn" title="Bullet List">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><line x1="8" y1="6" x2="21" y2="6"/><line x1="8" y1="12" x2="21" y2="12"/><line x1="8" y1="18" x2="21" y2="18"/><line x1="3" y1="6" x2="3.01" y2="6"/><line x1="3" y1="12" x2="3.01" y2="12"/><line x1="3" y1="18" x2="3.01" y2="18"/></svg>
          </button>
          <button class="tool-btn active-tool" title="Numbered List">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><line x1="10" y1="6" x2="21" y2="6"/><line x1="10" y1="12" x2="21" y2="12"/><line x1="10" y1="18" x2="21" y2="18"/><path d="M4 6h1v4"/><path d="M4 10h2"/><path d="M6 18H4c0-1 2-2 2-3s-1-1.5-2-1"/></svg>
          </button>
          <button class="tool-btn" title="Highlight">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><line x1="12" y1="1" x2="12" y2="23"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/></svg>
          </button>
          <button class="tool-btn" title="Task">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><polyline points="9 11 12 14 22 4"/><path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/></svg>
          </button>
          <button class="tool-btn" title="Table">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><line x1="3" y1="9" x2="21" y2="9"/><line x1="3" y1="15" x2="21" y2="15"/><line x1="9" y1="3" x2="9" y2="21"/><line x1="15" y1="3" x2="15" y2="21"/></svg>
          </button>
          <button class="tool-btn" title="Divider">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><line x1="5" y1="12" x2="19" y2="12"/></svg>
          </button>
          <button class="tool-btn" title="Page">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
          </button>
          <button class="tool-btn" title="Math">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 4h16v4L11 12l9 4v4H4"/></svg>
          </button>
          <button class="tool-btn" title="Diagram">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><line x1="18" y1="20" x2="18" y2="10"/><line x1="12" y1="20" x2="12" y2="4"/><line x1="6" y1="20" x2="6" y2="14"/></svg>
          </button>
          <button class="tool-btn" title="Sticker">
            <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 14s1.5 2 4 2 4-2 4-2"/><line x1="9" y1="9" x2="9.01" y2="9"/><line x1="15" y1="9" x2="15.01" y2="9"/></svg>
          </button>
          <span class="char-count">10000</span>
        </div>

        <textarea 
          class="editor-textarea" 
          v-model="modalContent"
          placeholder="Write your note here..."
        ></textarea>

        <div class="modal-actions">
          <button class="modal-btn cancel-btn" @click="closeModal">Cancel</button>
          <button class="modal-btn save-btn" @click="saveNote">Save</button>
        </div>

      </div>
    </div>

  </div>
</template>

<style scoped>
.note-container {
  background-color: #ffffff;
  padding: 24px;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  color: #1e293b;
  position: relative;
}

.create-note-wrapper {
  display: flex;
  position: relative;
  width: 100%;
}

.create-note-input {
  width: 100%;
  border: 1px solid #cbd5e1;
  border-radius: 10px;
  padding: 14px 50px 14px 18px;
  font-size: 15px;
  color: #0f172a;
  outline: none;
  transition: border-color 0.2s ease;

}

.create-note-input:focus {

    box-shadow: 0 0 0 3px #1B75D2;
}

.add-note-btn {
  position: absolute;
  right: 6px;
  top: 6px;
  bottom: 6px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 8px;
  width: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.add-note-btn:hover {
  background-color: #2563eb;
}

.note-controls {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
}

.select-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.custom-select {
  appearance: none;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  padding: 10px 36px 10px 14px;
  font-size: 14px;
  font-weight: 500;
  color: #334155;
  cursor: pointer;
  outline: none;
  transition: border-color 0.2s ease;
}

.custom-select:focus {
  box-shadow: 0 0 0 3px #1B75D2;
}

.dropdown-arrow {
  position: absolute;
  right: 12px;
  color: #64748b;
  pointer-events: none;
}

.notes-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.note-card {
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.note-header-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.note-time-badge {
  background-color: #eff6ff;
  color: #2563eb;
  font-size: 12px;
  font-weight: 700;
  padding: 4px 8px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  gap: 5px;
}

.note-lecture-tag {
  font-size: 12px;
  font-weight: 600;
  color: #64748b;
}

.note-text {
  font-size: 14px;
  color: #1e293b;
  margin: 0;
  line-height: 1.5;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  inset: 0;
  background-color: rgba(15, 23, 42, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-card {
  background-color: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 16px;
  width: 90%;
  max-width: 680px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
}

.modal-time-badge {
  background-color: #f1f5f9;
  color: #0f172a;
  border: 1px solid #cbd5e1;
  font-size: 13px;
  font-weight: 600;
  padding: 4px 10px;
  border-radius: 8px;
  width: fit-content;
  display: flex;
  align-items: center;
  gap: 6px;
}

.editor-toolbar {
  display: flex;
  align-items: center;
  gap: 2px;
  flex-wrap: wrap;
  padding-bottom: 6px;
  border-bottom: 1px solid #e2e8f0;
  color: #475569;
}

.tool-btn {
  background: transparent;
  border: none;
  color: #475569;
  cursor: pointer;
  padding: 4px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.tool-btn:hover {
  background-color: #f1f5f9;
  color: #0f172a;
}

.tool-btn.active-tool {
  background-color: #e2e8f0;
  color: #0f172a;
  border-radius: 8px;
}

.char-count {
  margin-left: auto;
  font-size: 12px;
  color: #64748b;
  font-weight: 500;
  padding-left: 8px;
}

.editor-textarea {
  width: 100%;
  height: 180px;
  background-color: #f8fafc;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  padding: 10px;
  color: #0f172a;
  font-size: 15px;
  outline: none;
  resize: vertical;
}

.editor-textarea:focus {
    box-shadow: 0 0 0 3px #1B75D2;
  /* border-color: #3b82f6; */
  background-color: #ffffff;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 2px;
}

.modal-btn {
  padding: 8px 20px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
  border: none;
  transition: opacity 0.2s;
}

.modal-btn:hover {
  opacity: 0.9;
}

.cancel-btn {
  background-color: #e2e8f0;
  color: #334155;
}

.save-btn {
  background-color: #1B75D2;
  color: #ffffff;
}
</style>