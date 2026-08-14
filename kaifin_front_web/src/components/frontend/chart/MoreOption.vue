<template>
  <div class="more-option-menu" @click.stop>
    <div class="options-list">
      <button
        v-for="option in options"
        :key="option.name"
        type="button"
        class="option-item"
        @click="handleOption(option)"
      >
        <span class="icon" v-html="option.svg"></span>
        <span class="label">{{ option.name }}</span>
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
const emit = defineEmits(['close', 'action'])

// SVG Icons ដែលមានពណ៌ #000 ទាំងអស់
const options = ref([
  { 
    name: 'Unread', 
    action: 'mark-unread',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path><polyline points="22,6 12,13 2,6"></polyline></svg>` 
  },
  { 
    name: 'Mute notification', 
    action: 'mute',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M13.73 21a2 2 0 0 1-3.46 0"></path><path d="M18.63 13A17.89 17.89 0 0 1 18 8"></path><path d="M6.26 6.26A5.86 5.86 0 0 0 6 8c0 7-3 9-3 9h14"></path><path d="M18 8a6 6 0 0 0-9.33-5"></path><line x1="1" y1="1" x2="23" y2="23"></line></svg>` 
  },
  { 
    name: 'View profile', 
    action: 'view-profile',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>` 
  },
  { 
    name: 'Audio call', 
     action: 'audio-call',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path></svg>` 
  },
  { 
    name: 'Video chat', 
     action: 'video-chat',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="23 7 16 12 23 17 23 7"></polygon><rect x="1" y="5" width="15" height="14" rx="2" ry="2"></rect></svg>` 
  },
  { 
    name: 'Block', 
    action: 'block',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line></svg>` 
  },
  { 
    name: 'Archive chat', 
     action: 'archive',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="21 8 21 21 3 21 3 8"></polyline><rect x="1" y="3" width="22" height="5"></rect><line x1="10" y1="12" x2="14" y2="12"></line></svg>` 
  },
  { 
    name: 'Delete chat', 
      action: 'delete',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path><line x1="10" y1="11" x2="10" y2="17"></line><line x1="14" y1="11" x2="14" y2="17"></line></svg>` 
  },
  { 
    name: 'Report', 
    action: 'report',
    svg: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path><line x1="12" y1="9" x2="12" y2="13"></line><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>` 
  },
]);

const handleOption = (option) => {
  emit('action', option.action)
  emit('close')
}
</script>

<style scoped>
.more-option-modal {
  background-color: #ffffff; /* ផ្ទៃខាងក្រោយពណ៌ស */
  width: 320px;
  border-radius: 20px;
  padding: 20px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  border: 1px solid #eaeaea;
  font-family: 'Plus Jakarta Sans', sans-serif;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.title {
  font-weight: 800;
  font-size: 18px;
  color: #1a1a1a;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #555;
}

.options-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
  background-color: #ffffff;
  border: 1px solid #f0f0f0;
  border-radius: 12px;
}

.option-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 7px 12px;
  border: none;
  background-color: transparent;
  cursor: pointer;
  transition: all 0.2s ease;
  width: 100%;
  text-align: left;
}

.option-item:hover {
  background: #f1f3f5;
}

.icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.label {
  font-weight: 600;
  color: #1a1a1a; /* អត្ថបទពណ៌ខ្មៅ/ដិត */
  font-size: 15px;
}
</style>