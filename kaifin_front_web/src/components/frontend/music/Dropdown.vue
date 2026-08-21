<template>
  <div class="dropdown-container" ref="dropdownRef">
    <div class="dropdown-header">Columns</div>
    
    <!-- Album Option -->
    <div class="dropdown-item" @click="$emit('toggle', 'album')">
      <div class="item-left">
        <div class="icon-wrap">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
            <circle cx="12" cy="12" r="3"></circle>
          </svg>
        </div>
        <span class="item-label">Album</span>
      </div>
      <svg v-if="modelValue.album" class="check-icon" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
        <polyline points="20 6 9 17 4 12"></polyline>
      </svg>
    </div>

    <!-- Date Added Option -->
    <div class="dropdown-item" @click="$emit('toggle', 'dateAdded')">
      <div class="item-left">
        <div class="icon-wrap">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
            <line x1="16" y1="2" x2="16" y2="6"></line>
            <line x1="8" y1="2" x2="8" y2="6"></line>
            <line x1="3" y1="10" x2="21" y2="10"></line>
          </svg>
        </div>
        <span class="item-label">Date added</span>
      </div>
      <svg v-if="modelValue.dateAdded" class="check-icon" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
        <polyline points="20 6 9 17 4 12"></polyline>
      </svg>
    </div>

    <!-- Duration Option -->
    <div class="dropdown-item" @click="$emit('toggle', 'duration')">
      <div class="item-left">
        <div class="icon-wrap">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="10"></circle>
            <polyline points="12 6 12 12 16 14"></polyline>
          </svg>
        </div>
        <span class="item-label">Duration</span>
      </div>
      <svg v-if="modelValue.duration" class="check-icon" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
        <polyline points="20 6 9 17 4 12"></polyline>
      </svg>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

defineProps({
  modelValue: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['toggle', 'close'])
const dropdownRef = ref(null)
const handleClickOutside = (event) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target)) {
    emit('close')
  }
}

onMounted(() => {
  window.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  window.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.dropdown-container {
  background-color: #ffffff;
  border: 1px solid #eaeef2;
  border-radius: 12px;
  box-shadow: 0 8px 20px -4px rgba(0, 0, 0, 0.08);
  padding: 6px;
  width: 190px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.dropdown-header {
  font-size: 12px;
  font-weight: 600;
  color: #57606a;
  padding: 4px 6px 6px 6px;
  border-bottom: 1px solid #eaeef2;
  margin-bottom: 2px;
}

.dropdown-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 8px;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.15s ease;
}

.dropdown-item:hover {
  background-color: #f6f8fa;
  transform: translateX(1px);
}

.item-left {
  display: flex;
  align-items: center;
  gap: 8px;
}

.icon-wrap {
  width: 20px;
  height: 20px;
  background-color: transparent;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #57606a;
}

.item-label {
  font-size: 13px;
  font-weight: 500;
  color: #24292f;
}

.check-icon {
  color: #0969da;
}
</style>