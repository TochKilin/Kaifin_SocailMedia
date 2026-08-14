<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';

const emit = defineEmits(['close', 'submit']);

const courseTitle = ref('');
const courseDescription = ref('');
const courseContent = ref(''); // ផ្ទុក HTML ពី rich text editor
const uploadedFileName = ref('');
const uploadedFileSize = ref('');
const fileThumbnail = ref('');
const isUploaded = ref(false);
const showMoreTools = ref(false);

const fileInputRef = ref(null);
const editorRef = ref(null);
const imageInputRef = ref(null);
const isEditorEmpty = ref(true);

const activeStates = ref({
  bold: false,
  italic: false,
  insertUnorderedList: false,
  insertOrderedList: false,
});

const triggerFileUpload = () => {
  fileInputRef.value?.click();
};

const formatFileSize = (bytes) => {
  if (bytes === 0) return '0 Bytes';
  const k = 1024;
  const sizes = ['Bytes', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
};

const handleFileChange = (e) => {
  const file = e.target.files[0];
  if (file) {
    uploadedFileName.value = file.name;
    uploadedFileSize.value = formatFileSize(file.size);
    isUploaded.value = true;

    if (file.type.startsWith('image/')) {
      const reader = new FileReader();
      reader.onload = (event) => {
        fileThumbnail.value = event.target.result;
      };
      reader.readAsDataURL(file);
    } else {
      fileThumbnail.value = '';
    }
  }
};

const toggleMoreTools = () => {
  showMoreTools.value = !showMoreTools.value;
};

/* ==================== Rich Text Editor Logic ==================== */

const focusEditor = () => {
  editorRef.value?.focus();
};

const syncContent = () => {
  if (editorRef.value) {
    courseContent.value = editorRef.value.innerHTML;
    isEditorEmpty.value = editorRef.value.innerText.trim().length === 0;
  }
};

const updateActiveStates = () => {
  try {
    activeStates.value.bold = document.queryCommandState('bold');
    activeStates.value.italic = document.queryCommandState('italic');
    activeStates.value.insertUnorderedList = document.queryCommandState('insertUnorderedList');
    activeStates.value.insertOrderedList = document.queryCommandState('insertOrderedList');
  } catch (e) {
    /* browser មិនអនុញ្ញាត queryCommandState ក្នុងករណីខ្លះ */
  }
};

const onEditorInput = () => {
  syncContent();
  updateActiveStates();
};

const onSelectionChange = () => {
  if (document.activeElement === editorRef.value) {
    updateActiveStates();
  }
};

onMounted(() => {
  document.addEventListener('selectionchange', onSelectionChange);
});

onBeforeUnmount(() => {
  document.removeEventListener('selectionchange', onSelectionChange);
});

const execCmd = (command, value = null) => {
  focusEditor();
  document.execCommand(command, false, value);
  syncContent();
  updateActiveStates();
};

const escapeHtml = (str) => {
  const div = document.createElement('div');
  div.innerText = str;
  return div.innerHTML;
};

/* ---- Toolbar Actions (ដំណើរការពិត) ---- */

const applyHeading = () => {
  focusEditor();
  const block = document.queryCommandValue('formatBlock');
  if (block && block.toLowerCase() === 'h2') {
    document.execCommand('formatBlock', false, 'P');
  } else {
    document.execCommand('formatBlock', false, 'H2');
  }
  syncContent();
};

const applyBold = () => execCmd('bold');
const applyItalic = () => execCmd('italic');

const applyQuote = () => {
  focusEditor();
  document.execCommand('formatBlock', false, 'BLOCKQUOTE');
  syncContent();
};

const applyBulletList = () => execCmd('insertUnorderedList');
const applyNumberedList = () => execCmd('insertOrderedList');

const applyLink = () => {
  const url = window.prompt('បញ្ចូល URL:', 'https://');
  if (url) execCmd('createLink', url);
};

const triggerImageInsert = () => {
  imageInputRef.value?.click();
};

const handleContentImageChange = (e) => {
  const file = e.target.files[0];
  if (!file) return;
  const reader = new FileReader();
  reader.onload = (event) => {
    focusEditor();
    document.execCommand('insertImage', false, event.target.result);
    syncContent();
  };
  reader.readAsDataURL(file);
  e.target.value = '';
};

const applyCode = () => {
  focusEditor();
  const selection = window.getSelection();
  const text = selection && selection.toString() ? selection.toString() : 'code';
  document.execCommand('insertHTML', false, `<code class="inline-code">${escapeHtml(text)}</code>&nbsp;`);
  syncContent();
};

const applyVariable = () => {
  focusEditor();
  document.execCommand('insertHTML', false, `<span class="var-tag" contenteditable="false">{{variable}}</span>&nbsp;`);
  syncContent();
};

const applyMoney = () => {
  focusEditor();
  document.execCommand('insertHTML', false, `<span class="money-tag" contenteditable="false">$0.00</span>&nbsp;`);
  syncContent();
};

const applyTask = () => {
  focusEditor();
  const id = 'task-' + Date.now();
  document.execCommand('insertHTML', false, `
    <div class="task-item-block">
      <input type="checkbox" id="${id}" />
      <label for="${id}" contenteditable="true">កិច្ចការថ្មី</label>
    </div><p><br></p>`);
  syncContent();
};

const applyTable = () => {
  focusEditor();
  const html = `
    <table class="editor-table">
      <tbody>
        <tr><td>Cell 1</td><td>Cell 2</td><td>Cell 3</td></tr>
        <tr><td>Cell 4</td><td>Cell 5</td><td>Cell 6</td></tr>
      </tbody>
    </table><p><br></p>`;
  document.execCommand('insertHTML', false, html);
  syncContent();
};

const alignCommands = ['justifyLeft', 'justifyCenter', 'justifyRight'];
const alignIndex = ref(0);
const applyAlign = () => {
  alignIndex.value = (alignIndex.value + 1) % alignCommands.length;
  execCmd(alignCommands[alignIndex.value]);
};

const applyPageBreak = () => {
  focusEditor();
  document.execCommand('insertHTML', false, `<hr class="page-break-line" /><p><br></p>`);
  syncContent();
};

const applyMath = () => {
  const formula = window.prompt('បញ្ចូលរូបមន្តគណិតវិទ្យា:', 'x^2 + y^2 = z^2');
  if (formula) {
    focusEditor();
    document.execCommand('insertHTML', false, `<span class="math-tag" contenteditable="false">∑ ${escapeHtml(formula)}</span>&nbsp;`);
    syncContent();
  }
};

const applyFlowchart = () => {
  focusEditor();
  const html = `
    <div class="flowchart-block" contenteditable="false">
      <div class="flow-box">Start</div>
      <div class="flow-arrow">↓</div>
      <div class="flow-box">Process</div>
      <div class="flow-arrow">↓</div>
      <div class="flow-box">End</div>
    </div><p><br></p>`;
  document.execCommand('insertHTML', false, html);
  syncContent();
};

/* ==================== Submit ==================== */

const handleSubmit = () => {
  if (!courseTitle.value.trim()) {
    alert('Please enter the course title');
    return;
  }
  syncContent();
  emit('submit', {
    title: courseTitle.value,
    description: courseDescription.value,
    content: courseContent.value,
    file: uploadedFileName.value
  });
  emit('close');
};
</script>

<template>
  <div class="create-course-container">
    <div class="create-course-wrapper">

      <!-- Modal Header -->
      <div class="modal-header">
        <div class="modal-header-text">
          <h3 class="modal-title">Create New Course</h3>
          <p class="modal-subtitle">Create a new course for your students</p>
        </div>
        <button class="modal-close-btn" type="button" @click="emit('close')">
          <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>

      <!-- ផ្នែកទី១: Course Title & File Upload -->
      <div class="form-section-card">
        <label class="form-label">Course Title & File</label>
        <div class="input-group text-input-wrapper">
          <input
            v-model="courseTitle"
            type="text"
            class="text-input-field"
            placeholder="Enter course title, e.g., Java Full Course 2026"
          />
        </div>

        <div class="upload-container-outer" @click="triggerFileUpload">
          <div class="upload-header-labels">
            <template v-if="!isUploaded">
              <span class="upload-main-text">Click here to upload file</span>
              <span class="upload-sub-text">PDF, ZIP, MP4 or Image (Max 50MB)</span>
            </template>
            <template v-else>
              <div class="file-success-info">
                <span class="upload-main-text file-name-display">📄 {{ uploadedFileName }}</span>
                <span class="upload-sub-text file-size-display">{{ uploadedFileSize }}</span>
                <span class="uploaded-badge">✅ Uploaded</span>
              </div>
            </template>
          </div>

          <div class="upload-box-wrapper" :class="{ 'uploaded-state': isUploaded }">
            <input
              ref="fileInputRef"
              type="file"
              class="hidden-file-input"
              @change="handleFileChange"
            />
            <div class="upload-inner-container">
              <div class="upload-stripe-bg"></div>
              <img v-if="fileThumbnail" :src="fileThumbnail" class="file-thumbnail-img" alt="Thumbnail" />
              <template v-else>
                <svg v-if="!isUploaded" class="upload-icon-plain" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                  <polyline points="17 8 12 3 7 8"></polyline>
                  <line x1="12" y1="3" x2="12" y2="15"></line>
                </svg>
                <svg v-else class="upload-icon-plain success-check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                  <polyline points="20 6 9 17 4 12"></polyline>
                </svg>
              </template>
            </div>
          </div>
        </div>
      </div>

      <!-- ផ្នែកទី២: Course Description -->
      <div class="form-section-card">
        <label class="form-label">Course Description</label>
        <div class="textarea-wrapper">
          <textarea
            v-model="courseDescription"
            class="text-area-field resizable-textarea"
            placeholder="Write a brief description of this course..."
          ></textarea>
        </div>
      </div>

      <!-- ផ្នែកទី៣: Detailed Content (Rich Text Editor ដំណើរការពិត) -->
      <div class="form-section-card">
        <label class="form-label">Detailed Content</label>
        <div class="editor-container">
          <div class="editor-toolbar">
            <div class="tool-group">
              <button class="tool-btn" title="Heading" type="button" @click="applyHeading">
                <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4v16M18 4v16M6 12h12"/></svg>
              </button>
              <button class="tool-btn" :class="{ active: activeStates.bold }" title="Bold" type="button" @click="applyBold">
                <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path><path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path></svg>
              </button>
              <button class="tool-btn" :class="{ active: activeStates.italic }" title="Italic" type="button" @click="applyItalic">
                <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><line x1="19" y1="4" x2="10" y2="4"></line><line x1="14" y1="20" x2="5" y2="20"></line><line x1="15" y1="4" x2="9" y2="20"></line></svg>
              </button>
            </div>

            <div class="tool-group">
              <button class="tool-btn" title="Quote" type="button" @click="applyQuote">
                <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 21c3 0 7-1 7-8V5c0-1.25-.75-2-2-2H4c-1.25 0-2 .75-2 2v6c0 1.25.75 2 2 2h3c0 3-2 5-4 5v1zm14 0c3 0 7-1 7-8V5c0-1.25-.75-2-2-2h-4c-1.25 0-2 .75-2 2v6c0 1.25.75 2 2 2h3c0 3-2 5-4 5v1z"/></svg>
              </button>
              <button class="tool-btn" title="Link" type="button" @click="applyLink">
                <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
              </button>
              <button class="tool-btn" title="Image" type="button" @click="triggerImageInsert">
                <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
              </button>
              <input ref="imageInputRef" type="file" accept="image/*" class="hidden-file-input" @change="handleContentImageChange" />
            </div>

            <div class="tool-group">
              <button class="tool-btn" :class="{ active: activeStates.insertUnorderedList }" title="Bullet List" type="button" @click="applyBulletList">
                <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><line x1="8" y1="6" x2="21" y2="6"></line><line x1="8" y1="12" x2="21" y2="12"></line><line x1="8" y1="18" x2="21" y2="18"></line><line x1="3" y1="6" x2="3.01" y2="6"></line><line x1="3" y1="12" x2="3.01" y2="12"></line><line x1="3" y1="18" x2="3.01" y2="18"></line></svg>
              </button>
              <button class="tool-btn" :class="{ active: activeStates.insertOrderedList }" title="Numbered List" type="button" @click="applyNumberedList">
                <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><line x1="10" y1="6" x2="21" y2="6"></line><line x1="10" y1="12" x2="21" y2="12"></line><line x1="10" y1="18" x2="21" y2="18"></line><path d="M4 6h1v4"></path><path d="M4 10h2"></path><path d="M6 18H4c0-1 2-2 2-3s-1-1.5-2-1"></path></svg>
              </button>
            </div>

            <template v-if="showMoreTools">
              <div class="tool-group">
                <button class="tool-btn" title="Code" type="button" @click="applyCode">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><polyline points="16 18 22 12 16 6"></polyline><polyline points="8 6 2 12 8 18"></polyline></svg>
                </button>
                <button class="tool-btn" title="Variables" type="button" @click="applyVariable">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M8 3H7a2 2 0 0 0-2 2v5a2 2 0 0 1-2 2 2 2 0 0 1 2 2v5a2 2 0 0 0 2 2h1"></path><path d="M16 21h1a2 2 0 0 0 2-2v-5a2 2 0 0 1 2-2 2 2 0 0 1-2-2V5a2 2 0 0 0-2-2h-1"></path></svg>
                </button>
                <button class="tool-btn" title="Money" type="button" @click="applyMoney">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="1" x2="12" y2="23"></line><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path></svg>
                </button>
              </div>

              <div class="tool-group">
                <button class="tool-btn" title="Task" type="button" @click="applyTask">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 11 12 14 22 4"></polyline><path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"></path></svg>
                </button>
                <button class="tool-btn" title="Table" type="button" @click="applyTable">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><line x1="3" y1="9" x2="21" y2="9"></line><line x1="3" y1="15" x2="21" y2="15"></line><line x1="9" y1="3" x2="9" y2="21"></line><line x1="15" y1="3" x2="15" y2="21"></line></svg>
                </button>
                <button class="tool-btn" title="Align" type="button" @click="applyAlign">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><line x1="21" y1="10" x2="3" y2="10"></line><line x1="21" y1="6" x2="3" y2="6"></line><line x1="21" y1="14" x2="3" y2="14"></line><line x1="21" y1="18" x2="3" y2="18"></line></svg>
                </button>
              </div>

              <div class="tool-group last-group">
                <button class="tool-btn" title="Page Break" type="button" @click="applyPageBreak">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path><polyline points="14 2 14 8 20 8"></polyline><line x1="16" y1="13" x2="8" y2="13"></line><line x1="16" y1="17" x2="8" y2="17"></line><polyline points="10 9 9 9 8 9"></polyline></svg>
                </button>
                <button class="tool-btn" title="Math" type="button" @click="applyMath">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path><line x1="8" y1="12" x2="16" y2="12"></line><line x1="12" y1="8" x2="12" y2="16"></line></svg>
                </button>
                <button class="tool-btn" title="Flowchart" type="button" @click="applyFlowchart">
                  <svg class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><rect x="6" y="2" width="12" height="6" rx="1"></rect><rect x="2" y="16" width="8" height="6" rx="1"></rect><rect x="14" y="16" width="8" height="6" rx="1"></rect><path d="M12 8v4a2 2 0 0 1-2 2H6v2"></path><path d="M12 12h4a2 2 0 0 1 2 2v2"></path></svg>
                </button>
              </div>
            </template>

            <div class="tool-group toggle-group" :class="{ 'last-group': !showMoreTools }">
              <button
                class="tool-btn more-toggle-btn"
                :title="showMoreTools ? 'Less options' : 'More options'"
                type="button"
                @click="toggleMoreTools"
              >
                <svg v-if="!showMoreTools" class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="1"></circle>
                  <circle cx="19" cy="12" r="1"></circle>
                  <circle cx="5" cy="12" r="1"></circle>
                </svg>
                <svg v-else class="toolbar-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="15 18 9 12 15 6"></polyline>
                </svg>
              </button>
            </div>
          </div>

          <!-- Contenteditable Rich Text Area (ដំណើរការពិត) -->
          <div class="editor-textarea-area">
            <div
              ref="editorRef"
              class="editor-real-textarea resizable-textarea rich-content-editable"
              :class="{ 'is-empty': isEditorEmpty }"
              contenteditable="true"
              data-placeholder="Enter detailed content here..."
              @input="onEditorInput"
              @click="updateActiveStates"
              @keyup="updateActiveStates"
            ></div>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="form-action-row">
        <button class="cancel-btn" type="button" @click="emit('close')">Cancel</button>
        <button class="submit-btn" type="button" @click="handleSubmit">
          <span>Create Course</span>
          <svg class="send-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 2L11 13M22 2l-7 20-4-9-9-4 20-7z"/></svg>
        </button>
      </div>

    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700&display=swap');

.create-course-container {
  font-family: 'Plus Jakarta Sans', sans-serif;
  display: flex;
  justify-content: center;
  background: transparent;
  box-sizing: border-box;
}

.create-course-wrapper {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 4px 0px 10px 0px;
  border-bottom: 1px solid #e2e8f0;
  margin-bottom: 4px;
}

.modal-header-text {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.modal-title {
  font-size: 16px;
  font-weight: 700;
  color: #1e293b;
  margin: 0;
  line-height: 1.2;
}

.modal-subtitle {
  font-size: 12px;
  color: #64748b;
  margin: 0;
  line-height: 1.2;
  font-weight: 500;
}

.modal-close-btn {
  background: transparent;
  border: 1px solid #cbd5e1;
  cursor: pointer;
  color: #64748b;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 6px;
  transition: all 0.2s;
}

.modal-close-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
  border-color: #94a3b8;
  color: #1e293b;
}

.form-section-card {
  border: none;
  background: transparent;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.form-label {
  font-size: 13px;
  font-weight: 600;
  color: #334155;
  display: flex;
  align-items: center;
  gap: 6px;
}

.text-input-wrapper {
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  background-color: rgba(255, 255, 255, 0.5);
  transition: background-color 0.2s ease, box-shadow 0.2s ease;
  position: relative;
}

.text-input-wrapper:hover {
  border-color: #cbd5e1;
}

.text-input-wrapper:focus-within {
  border-color: #1c74d287;
  background-color: #ffffff;
  box-shadow: 0 0 0 3px #1c74d287;
  z-index: 2;
}

.text-input-field {
  width: 100%;
  background-color: transparent;
  border: none;
  border-radius: 12px;
  padding: 8px 12px;
  color: #1e293b;
  font-size: 14px;
  outline: none;
  box-sizing: border-box;
  box-shadow: none;
}

.text-input-field::placeholder {
  color: #94a3b8;
}

.upload-container-outer {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 8px;
  width: fit-content;
  cursor: pointer;
  margin-top: 10px;
}

.upload-header-labels {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 3px;
}

.upload-main-text {
  font-weight: 600;
  font-size: 13px;
  color: #1e293b;
  line-height: 1.2;
}

.upload-sub-text {
  font-size: 11px;
  color: #64748b;
  line-height: 1.2;
}

.file-success-info {
  display: flex;
  flex-direction: column;
  gap: 3px;
}

.file-name-display {
  color: #1e293b;
  max-width: 250px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.file-size-display {
  color: #64748b;
}

.uploaded-badge {
  font-size: 11px;
  font-weight: 600;
  color: #16a34a;
  margin-top: 2px;
}

.upload-box-wrapper {
  border: 1px solid #cbd5e1;
  border-radius: 50%;
  width: 110px;
  height: 110px;
  background-color: rgba(248, 250, 252, 0.5);
  transition: all 0.25s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
  overflow: hidden;
  position: relative;
}

.upload-box-wrapper.uploaded-state {
  border-color: #16a34a;
  background-color: rgba(22, 163, 74, 0.05);
}

.upload-box-wrapper:hover {
  border-color: #cbd5e1;
  background-color: rgba(239, 246, 255, 0.4);
}

.upload-box-wrapper.uploaded-state:hover {
  border-color: #16a34a;
  background-color: rgba(22, 163, 74, 0.1);
}

.upload-inner-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  position: relative;
}

.upload-stripe-bg {
  position: absolute;
  width: 100%;
  height: 24px;
  background-color: rgba(9, 9, 9, 0.051);
  backdrop-filter: blur(2px);
  z-index: 1;
}

.file-thumbnail-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  position: relative;
  z-index: 2;
  border-radius: 50%;
}

.upload-icon-plain {
  width: 26px;
  height: 26px;
  stroke: #1C75D2;
  position: relative;
  z-index: 2;
}

.success-check-icon {
  stroke: #16a34a;
}

.hidden-file-input {
  display: none;
}

.textarea-wrapper {
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  padding: 2px;
  background-color: rgba(255, 255, 255, 0.5);
  transition: background-color 0.2s ease, box-shadow 0.2s ease;
  position: relative;
}

.textarea-wrapper:hover {
  border-color: #cbd5e1;
}

.textarea-wrapper:focus-within {
  border-color: #1c74d287;
  background-color: #ffffff;
  box-shadow: 0 0 0 3px #1c74d287;
  z-index: 2;
}

.resizable-textarea {
  resize: vertical;
  min-height: 90px;
  overflow-y: auto;
}

.text-area-field {
  width: 100%;
  height: 110px;
  background-color: transparent;
  border: none;
  border-radius: 8px;
  padding: 8px 12px;
  font-size: 13px;
  outline: none;
  box-sizing: border-box;
  font-family: inherit;
  color: #1e293b;
  box-shadow: none;
}

.text-area-field::placeholder {
  color: #94a3b8;
}

.editor-container {
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  padding: 0px;
  display: flex;
  flex-direction: column;
  gap: 0px;
  background-color: rgba(255, 255, 255, 0.5);
  overflow: hidden;
  box-shadow: none !important;
}

.editor-container:focus-within {
  border-color: #1c74d287;
  box-shadow: 0 0 0 3px #1c74d287 !important;
}

.editor-toolbar {
  border: none;
  border-bottom: 1px solid #cbd5e1;
  padding: 0px 10px;
  display: flex;
  width: 100%;
  box-sizing: border-box;
  justify-content: flex-start;
  align-items: center;
  background-color: rgba(248, 250, 252, 0.8);
  margin: 0;
  overflow-x: auto;
  height: 44px;
  flex-shrink: 0;
}

.tool-group {
  display: flex;
  align-items: center;
  border-right: 1px solid #cbd5e1;
  padding: 0 6px;
  gap: 4px;
  height: 28px;
}

.tool-group.last-group {
  border-right: none;
}

.tool-btn {
  background: transparent;
  border: none;
  color: #334155;
  width: 30px;
  height: 30px;
  border-radius: 5px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  padding: 0;
}

.toolbar-svg {
  width: 15px;
  height: 15px;
  stroke: currentColor;
}

.tool-btn.active {
  background-color: #bfdbfe;
  color: #1C75D2;
}

.tool-btn:hover:not(.active) {
  background-color: rgba(191, 219, 254, 0.4);
  color: #1C75D2;
}

.more-toggle-btn {
  color: #1C75D2;
}

.editor-textarea-area {
  border: none;
  border-radius: 0 0 12px 12px;
  min-height: 180px;
  padding: 0;
  position: relative;
  background-color: transparent;
  margin: 0;
}

/* ==== Contenteditable Rich Text Area ==== */
.editor-real-textarea {
  width: 100%;
  height: 220px;
  background: transparent;
  border: none;
  outline: none;
  color: #1e293b;
  font-size: 14px;
  padding: 12px;
  box-sizing: border-box;
  font-family: inherit;
  overflow-y: auto;
  line-height: 1.6;
}

.editor-real-textarea.is-empty:before {
  content: attr(data-placeholder);
  color: #94a3b8;
  pointer-events: none;
}

.rich-content-editable :deep(h2) {
  font-size: 20px;
  font-weight: 700;
  margin: 8px 0;
  color: #0f172a;
}

.rich-content-editable :deep(blockquote) {
  border-left: 3px solid #1C75D2;
  margin: 8px 0;
  padding: 4px 12px;
  color: #475569;
  background: rgba(28, 117, 210, 0.05);
  border-radius: 0 6px 6px 0;
}

.rich-content-editable :deep(a) {
  color: #1C75D2;
  text-decoration: underline;
}

.rich-content-editable :deep(ul),
.rich-content-editable :deep(ol) {
  margin: 6px 0;
  padding-left: 22px;
}

.rich-content-editable :deep(img) {
  max-width: 100%;
  border-radius: 8px;
  margin: 8px 0;
}

.rich-content-editable :deep(.inline-code) {
  background: #eef2ff;
  color: #4f46e5;
  padding: 2px 6px;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  font-size: 13px;
}

.rich-content-editable :deep(.var-tag) {
  background: #fef3c7;
  color: #92400e;
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 12px;
  font-weight: 600;
}

.rich-content-editable :deep(.money-tag) {
  background: #dcfce7;
  color: #166534;
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 12px;
  font-weight: 700;
}

.rich-content-editable :deep(.math-tag) {
  background: #ede9fe;
  color: #5b21b6;
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 13px;
  font-family: 'Courier New', monospace;
}

.rich-content-editable :deep(.task-item-block) {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 4px 0;
}

.rich-content-editable :deep(.task-item-block input[type="checkbox"]) {
  width: 16px;
  height: 16px;
  cursor: pointer;
  accent-color: #1C75D2;
}

.rich-content-editable :deep(.editor-table) {
  width: 100%;
  border-collapse: collapse;
  margin: 8px 0;
}

.rich-content-editable :deep(.editor-table td) {
  border: 1px solid #cbd5e1;
  padding: 6px 10px;
  font-size: 13px;
}

.rich-content-editable :deep(.page-break-line) {
  border: none;
  border-top: 2px dashed #cbd5e1;
  margin: 12px 0;
}

.rich-content-editable :deep(.flowchart-block) {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
  padding: 10px 0;
}

.rich-content-editable :deep(.flow-box) {
  background: #eff6ff;
  border: 1px solid #1C75D2;
  color: #1C75D2;
  padding: 6px 20px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 600;
}

.rich-content-editable :deep(.flow-arrow) {
  color: #94a3b8;
  font-size: 14px;
}

.form-action-row {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 4px;
}

.cancel-btn {
  background: transparent;
  border: 1px solid #cbd5e1;
  color: #64748b;
  padding: 8px 16px;
  border-radius: 32px;
  font-weight: 600;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.2s;
}

.cancel-btn:hover {
  background-color: rgba(0, 0, 0, 0.04);
  border-color: #94a3b8;
  color: #1e293b;
}

.submit-btn {
  background: #1C75D2;
  border: none;
  color: #ffffff;
  padding: 8px 16px;
  border-radius: 32px;
  font-weight: 700;
  font-size: 13px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: all 0.25s ease;
}

.submit-btn:hover {
  transform: translateY(-2px);
  opacity: 0.98;
}

.send-icon {
  width: 15px;
  height: 15px;
  stroke: currentColor;
}
</style>