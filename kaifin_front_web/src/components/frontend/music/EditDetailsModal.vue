<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content">
      <!-- Header -->
      <div class="modal-header">
        <h2>Edit details</h2>
        <button class="close-btn" @click="$emit('close')">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>

      <!-- Body Form -->
      <div class="modal-body">
        <div class="modal-photo-box" @click="triggerFileInput">
          <input 
            type="file" 
            ref="modalFileRef" 
            style="display: none" 
            accept="image/*" 
            @change="handleFileChange" 
          />
          <template v-if="form.coverUrl">
            <img :src="form.coverUrl" alt="Cover" class="modal-uploaded-img" />
          </template>
          <template v-else>
            <div class="pencil-wrap">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M12 20h9"></path>
                <path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"></path>
              </svg>
            </div>
            <span class="add-photo-label">Add Photo</span>
          </template>
        </div>

        <!-- Right: Inputs -->
        <div class="modal-inputs-group">
          <div class="input-field-wrap">
            <label class="field-label">Name</label>
            <input 
              type="text" 
              class="modal-text-input" 
              v-model="form.name" 
              placeholder="Enter name" 
            />
          </div>

          <div class="input-field-wrap">
            <label class="field-label">Description</label>
            <textarea 
              class="modal-textarea" 
              v-model="form.description" 
              placeholder="Enter description"
              rows="4"
            ></textarea>
          </div>
        </div>
      </div>

      <!-- Footer Actions -->
      <div class="modal-footer">
        <button class="make-private-btn" @click="togglePrivacy">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
          </svg>
          {{ form.isPublic ? 'Make private' : 'Make public' }}
        </button>
        <button class="save-btn" @click="handleSave">Save</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'

const props = defineProps({
  playlistData: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['close', 'save'])

const form = reactive({
  name: props.playlistData.name,
  description: props.playlistData.description || '',
  coverUrl: props.playlistData.coverUrl,
  isPublic: props.playlistData.isPublic
})

const modalFileRef = ref(null)

const triggerFileInput = () => {
  modalFileRef.value?.click()
}

const handleFileChange = (e) => {
  const file = e.target.files[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (event) => {
      form.coverUrl = event.target.result
    }
    reader.readAsDataURL(file)
  }
}

const togglePrivacy = () => {
  form.isPublic = !form.isPublic
}

const handleSave = () => {
  emit('save', { ...form })
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  font-family: 'Plus Jakarta Sans', sans-serif;
}

.modal-content {
  background: #ffffff;
  width: 100%;
  max-width: 540px;
  border-radius: 16px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  animation: modalScale 0.2s ease;
}

@keyframes modalScale {
  from { transform: scale(0.95); opacity: 0; }
  to { transform: scale(1); opacity: 1; }
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  border-bottom: 1px solid #eaeef2;
}

.modal-header h2 {
  font-size: 16px;
  font-weight: 700;
  color: #24292f;
  margin: 0;
}

.close-btn {
  background: transparent;
  border: none;
  cursor: pointer;
  color: #57606a;
  border-radius: 6px;
  padding: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.close-btn:hover {
  background: #f6f8fa;
  color: #24292f;
}

.modal-body {
  display: flex;
  gap: 20px;
  padding: 24px 20px;
}

.modal-photo-box {
  width: 140px;
  height: 140px;
  border-radius: 12px;
  background: #f6f8fa;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  cursor: pointer;
  flex-shrink: 0;
  overflow: hidden;
  transition: all 0.2s;
}

.modal-photo-box:hover {
  border-color: #0969da;
  background: #ddf4ff;
  color: #0969da;
}

.modal-uploaded-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.pencil-wrap {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #eaeef2;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #24292f;
}

.add-photo-label {
  font-size: 12px;
  font-weight: 500;
  color: #57606a;
}

.modal-inputs-group {
  display: flex;
  flex-direction: column;
  gap: 14px;
  flex: 1;
}

.input-field-wrap {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.field-label {
  font-size: 12px;
  font-weight: 600;
  color: #24292f;
}

.modal-text-input, .modal-textarea {
  border: 1px solid #d0d7de;
  border-radius: 8px;
  padding: 8px 12px;
  font-size: 13px;
  font-family: 'Plus Jakarta Sans', sans-serif;
  color: #24292f;
  outline: none;
  transition: all 0.2s;
}

.modal-text-input:focus, .modal-textarea:focus {
  border-color: #0969da;
  box-shadow: 0 0 0 3px rgba(9, 105, 218, 0.15);
}

.modal-textarea {
  resize: none;
}

.modal-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  border-top: 1px solid #eaeef2;
  background: #f6f8fa;
}

.make-private-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  background: #ffffff;
  border: 1px solid #d0d7de;
  color: #24292f;
  padding: 6px 12px;
  border-radius: 8px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.make-private-btn:hover {
  background: #eaeef2;
}

.save-btn {
  background: #0969da;
  color: #ffffff;
  border: none;
  padding: 6px 20px;
  border-radius: 8px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}

.save-btn:hover {
  background: #0550ae;
}
</style>