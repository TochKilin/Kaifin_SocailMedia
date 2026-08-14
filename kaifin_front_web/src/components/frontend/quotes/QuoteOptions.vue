<template>
  <div class="options-wrapper" @mouseleave="$emit('close')">
    <div class="options-popup">
      <button class="option-item" @click="$emit('action', 'not_interested')">
        <span class="option-icon emoji-icon">😐</span>
        <span class="option-text">Not interested</span>
      </button>

      <button class="option-item" @click="$emit('action', 'block_user')">
        <span class="option-icon blocked-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zM4 12c0-1.84.63-3.53 1.68-4.87l11.19 11.19C15.53 19.37 13.84 20 12 20c-4.41 0-8-3.59-8-8zm14.32 4.87L7.13 5.68C8.47 4.63 10.16 4 12 4c4.41 0 8 3.59 8 8 0 1.84-.63 3.53-1.68 4.87z"/></svg>
        </span>
        <span class="option-text">Blocked user</span>
      </button>

      <button class="option-item" @click="$emit('action', 'report')">
        <span class="option-icon report-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path><line x1="12" y1="9" x2="12" y2="13"></line><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
        </span>
        <span class="option-text">Report</span>
      </button>
    </div>
  </div>
</template>

<script setup>
defineEmits(['action', 'close'])
</script>

<style scoped>
.options-wrapper {
  position: absolute;
  bottom: 100%;
  right: 0;
  padding-bottom: 10px;
  z-index: 100;
}

.options-popup {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  padding: 8px;
  display: flex;
  flex-direction: column;
  gap: 4px;
  width: 210px;
  box-shadow: 0 12px 30px -10px rgba(25, 118, 210, 0.15);
  animation: popUpFade 0.2s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes popUpFade {
  from {
    opacity: 0;
    transform: translateY(8px) scale(0.96);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.option-item {
  display: flex;
  align-items: center;
  gap: 12px;
  background: transparent;
  border: none;
  padding: 10px 12px;
  border-radius: 10px;
  cursor: pointer;
  width: 100%;
  text-align: left;
  transition: all 0.2s ease;
}

.option-item:hover {
  background-color: #f1f5f9;
}

.option-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
}

.emoji-icon {
  font-size: 18px;
}

.blocked-icon {
  color: #000; /* ពណ៌ក្រហមសម្រាប់ Blocked */
}

.report-icon {
  color: #000; /* ពណ៌លឿងទុំសម្រាប់ Report */
}

.option-text {
  color: #334155;
  font-size: 13.5px;
  font-weight: 500;
}

.option-item:hover .option-text {
  color: #1976d2;
  font-weight: 600;
}
</style>