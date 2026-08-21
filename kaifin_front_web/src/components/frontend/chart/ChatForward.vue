<template>
  <div class="forward-panel">
    <div class="forward-panel-header">
      <button class="back-btn" title="Back" @click="$emit('close')">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="15 18 9 12 15 6"></polyline></svg>
      </button>
      <h3>Forward</h3>
    </div>
    <div class="forward-search-box">
      <span class="for-label">For:</span>
      <input type="text" v-model="searchQuery" placeholder="Search for friends" />
      
      <div class="selected-avatars-row" v-if="selectedTargets.length > 0">
        <div class="selected-avatar-item" v-for="target in selectedTargets" :key="target.id">
          <img :src="target.avatar" :alt="target.name" />
        </div>
      </div>
    </div>

    <div class="forward-chat-list">
      <div
        v-for="chat in filteredChats"
        :key="chat.id"
        class="forward-chat-item"
        @click="toggleSelectChat(chat)"
      >
        <img :src="chat.avatar" :alt="chat.name" class="forward-avatar" />
        <span class="forward-name">{{ chat.name }}</span>

        <div class="wa-checkbox" :class="{ checked: isSelected(chat) }">
          <svg v-if="isSelected(chat)" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="#ffffff" stroke-width="3.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
        </div>
      </div>
    </div>

    <div class="forward-comment-box">
      <input type="text" v-model="commentText" placeholder="Write a comment" />
      <button class="action-btn forward-btn" @click="confirmForward" :disabled="selectedTargets.length === 0">Forward</button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  message: {
    type: [Object, Array],
    default: null
  },
  chats: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['close', 'forward'])

const searchQuery = ref('')
const selectedTargets = ref([])
const commentText = ref('')

const filteredChats = computed(() => {
  if (!searchQuery.value.trim()) return props.chats
  return props.chats.filter(c => c.name.toLowerCase().includes(searchQuery.value.toLowerCase()))
})

function isSelected(chat) {
  return selectedTargets.value.some(t => t.id === chat.id)
}

function toggleSelectChat(chat) {
  const index = selectedTargets.value.findIndex(t => t.id === chat.id)
  if (index > -1) {
    selectedTargets.value.splice(index, 1)
  } else {
    selectedTargets.value.push(chat)
  }
}

function confirmForward() {
  if (selectedTargets.value.length === 0) return
  emit('forward', {
    targetChats: selectedTargets.value,
    messages: Array.isArray(props.message) ? props.message : [props.message],
    comment: commentText.value
  })
}
</script>

<style scoped>
.forward-panel {
  width: 100%;
  height: 100%;
  min-height: 0;
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
}

.forward-panel-header {
  display: flex;
  align-items: center;
  gap: 12px;
  height: 65px;
  padding: 0 16px;
  border-bottom: 1px solid #e5e7eb;
  flex-shrink: 0;
  box-sizing: border-box;
}

.forward-panel-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #111b21;
}

.back-btn {
  background: transparent;
  border: none;
  cursor: pointer;
  color: #4b5563;
  padding: 8px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.back-btn:hover {
  background-color: #f3f4f6;
}

.forward-search-box {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  padding: 8px 16px;
  border-bottom: 1px solid #e5e7eb;
  gap: 8px;
  flex-shrink: 0;
  background-color: #ffffff;
}

.for-label {
  font-size: 14px;
  color: #374151;
  font-weight: 500;
}

.forward-search-box input {
  flex-grow: 1;
  min-width: 100px;
  background: transparent;
  border: none;
  outline: none;
  font-size: 14px;
  color: #111b21;
  padding: 4px 0;
}

.selected-avatars-row {
  display: flex;
  align-items: center;
  gap: 4px;
  overflow-x: auto;
  max-width: 50%;
  background-color: #f0f2f5; 
  padding: 3px 6px;
  border-radius: 20px; 
}

.selected-avatar-item img {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #ffffff; 
  flex-shrink: 0;
  display: block;
}

.forward-chat-list {
  flex-grow: 1;
  overflow-y: auto;
  padding: 4px 0;
  min-height: 0;
}

.forward-chat-item {
  display: flex;
  align-items: center;
  padding: 12px 20px;
  gap: 16px;
  cursor: pointer;
  border-bottom: 1px solid #f9fafb;
  transition: background 0.15s;
}

.forward-chat-item:hover {
  background-color: #f3f4f6;
}

.forward-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.forward-name {
  flex-grow: 1;
  font-size: 15px;
  font-weight: 500;
  color: #111b21;
}

.wa-checkbox {
  width: 20px;
  height: 20px;
  border: 2px solid #d1d5db;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #ffffff;
  transition: all 0.2s;
  flex-shrink: 0;
}

.wa-checkbox.checked {
  background-color: #1B75D2;
  border-color: #1B75D2;
}

.forward-comment-box {
  display: flex;
  align-items: center;
  padding: 12px 20px;
  border-top: 1px solid #e5e7eb;
  gap: 12px;
  flex-shrink: 0;
  background-color: #ffffff;
}

.forward-comment-box input {
  flex-grow: 1;
  border: none;
  outline: none;
  font-size: 14px;
  color: #111b21;
  background: transparent;
}

.action-btn {
  padding: 8px 20px;
  border-radius: 20px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  border: none;
  transition: background 0.2s;
  flex-shrink: 0;
}

.forward-btn {
  background-color: #1B75D2;
  color: #ffffff;
}

.forward-btn:hover {
  background-color: #155cb0;
}

.forward-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>