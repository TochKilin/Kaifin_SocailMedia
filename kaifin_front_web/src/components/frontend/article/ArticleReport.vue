<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const emit = defineProps(['close'])
const reportType = ref('bug')
const reportText = ref('')
const uploadedImages = ref([])

onMounted(() => {
  document.body.style.overflow = 'hidden'
})

onUnmounted(() => {
  document.body.style.overflow = ''
})

const reportTypes = [
  { id: 'bug', label: 'Bug / Error', icon: '🐞' },
  { id: 'feedback', label: 'General Feedback', icon: '🗯️' },
  { id: 'feature', label: 'Feature Request', icon: '🌠' },
  { id: 'other', label: 'Other', icon: '🪡' }
]

const handleImageUpload = (event) => {
  const files = Array.from(event.target.files)
  if (uploadedImages.value.length + files.length <= 4) {
    uploadedImages.value.push(...files.map(file => ({
      url: URL.createObjectURL(file),
      name: file.name
    })))
  } else {
    alert('You can upload a maximum of 4 images only!')
  }
}

const removeImage = (index) => {
  uploadedImages.value.splice(index, 1)
}

const handleSubmit = () => {
  console.log('Report submitted:', { type: reportType.value, text: reportText.value })
  alert('Your report has been submitted successfully!')
}
</script>

<template>
  <div class="report-modal-overlay">
    <div class="report-modal-container">
      
      <button class="close-btn" @click="$emit('close')">
        <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
      </button>

      <div class="report-header">
        <h2>Report Feedback</h2>
        <p>Please provide your feedback or report an issue to us</p>
      </div>

      <div class="report-type-group">
        <label class="field-label">Report Type</label>
        <div class="chips-container">
          <button 
            v-for="type in reportTypes" 
            :key="type.id"
            type="button"
            class="chip-btn"
            :class="{ active: reportType === type.id }"
            @click="reportType = type.id"
          >
            <span>{{ type.icon }}</span>
            <span>{{ type.label }}</span>
          </button>
        </div>
      </div>

      <div class="report-input-group">
        <textarea 
          v-model="reportText" 
          placeholder="Write detailed information here..." 
          rows="4"
          class="report-textarea"
        ></textarea>
      </div>

      <div class="report-upload-section">
        <div class="previews-container" v-if="uploadedImages.length > 0">
          <div v-for="(img, index) in uploadedImages" :key="index" class="preview-box">
            <img :src="img.url" alt="Uploaded preview" />
            <button class="remove-img-btn" @click="removeImage(index)">×</button>
          </div>
        </div>

        <label v-if="uploadedImages.length < 4" class="upload-box">
          <input type="file" accept="image/*" multiple @change="handleImageUpload" class="file-input" />
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
          <span>Upload {{ uploadedImages.length }}/4</span>
        </label>
      </div>

      <div class="report-actions">
        <button class="btn-cancel" @click="$emit('close')">Cancel</button>
        <button class="btn-submit" @click="handleSubmit" :disabled="!reportText.trim()">Submit</button>
      </div>

    </div>
  </div>
</template>

<style scoped>
.report-modal-overlay {
  position: fixed;
  top: 60px;
  left: 0;
  width: 100vw;
  height: calc(100vh - 60px);
  background-color: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  animation: fadeIn 0.25s ease-out;
  overflow: hidden;
}

.report-modal-container {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 20px;
  width: 480px;
  max-width: 90%;
  padding: 20px;
  position: relative;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  display: flex;
  flex-direction: column;
  gap: 14px;
  animation: scaleUp 0.25s ease-out;
}

.close-btn {
  position: absolute;
  top: 16px;
  right: 16px;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  width: 36px;
  height: 36px;
  color: #64748b;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.close-btn:hover {
  background: #f1f5f9;
  color: #0f172a;
  transform: scale(1.05);
}

.report-header h2 {
  color: #0f172a;
  font-size: 16px;
  font-weight: 700;
  margin: 0 0 4px 0;
}

.report-header p {
  color: #64748b;
  font-size: 13px;
  margin: 0;
}

.report-type-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.field-label {
  font-size: 13px;
  font-weight: 600;
  color: #334155;
}

.chips-container {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.chip-btn {
  background: transparent;
  border: 1px solid #e2e8f0;
  border-radius: 20px;
  padding: 8px 12px;
  font-size: 13px;
  font-weight: 500;
  color: #334155;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: all 0.2s;
  font-family: inherit;
}

.chip-btn:hover {
  background: #f8fafc;
  border-color: #cbd5e1;
}

.chip-btn.active {
  background: transparent;
  border-color: #e2e8f0;
  color: #334155;
  font-weight: 500;
}

.report-textarea {
  width: 100%;
  background: transparent;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 12px;
  color: #0f172a;
  font-size: 15px;
  resize: none;
  outline: none;
  font-family: inherit;
  transition: all 0.2s;
}

.report-textarea:focus {
  background: transparent;
  box-shadow: 0 0 0 3px rgba(27, 119, 210, 0.5);
}

.report-textarea::placeholder {
  color: #94a3b8;
}

.report-upload-section {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.previews-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.preview-box {
  width: 90px;
  height: 90px;
  border-radius: 12px;
  border: 1px solid #cbd5e1;
  position: relative;
  overflow: hidden;
  background: transparent;
}

.preview-box img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.remove-img-btn {
  position: absolute;
  top: 4px;
  right: 4px;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  border: none;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  font-size: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.2s;
}

.remove-img-btn:hover {
  background: rgba(0, 0, 0, 0.8);
}

.upload-box {
  width: 90px;
  height: 90px;
  background: transparent;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  gap: 6px;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #64748b;
  font-size: 12px;
  font-weight: 500;
  transition: all 0.2s;
}

.upload-box:hover {
  background: rgba(27, 117, 210, 0.04);
  border-color: #1B75D2;
  color: #1B75D2;
}

.file-input {
  display: none;
}

.report-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 0;
}

.btn-cancel, .btn-submit {
  padding: 6px 14px;
  border-radius: 32px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-cancel {
  background: transparent;
  color: #64748b;
  border: 1px solid #e2e8f0;
}

.btn-cancel:hover {
  background: rgba(0, 0, 0, 0.04);
  color: #0f172a;
}

.btn-submit {
  background: #1B75D2;
  color: #ffffff;
  border: 1px solid #1B75D2;
}

.btn-submit:not(:disabled):hover {
  background: #155cb8;
  border-color: #155cb8;
  box-shadow: 0 4px 12px rgba(27, 117, 210, 0.2);
}

.btn-submit:disabled {
  background: #e2e8f0;
  border-color: #e2e8f0;
  color: #94a3b8;
  cursor: not-allowed;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes scaleUp {
  from { transform: scale(0.95); opacity: 0; }
  to { transform: scale(1); opacity: 1; }
}
</style>