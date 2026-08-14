<template>
  <div class="share-wrapper" @mouseleave="$emit('close')">
    <div class="share-popup">
      <button class="share-item" @click="$emit('action', 'feed')">
        <span class="share-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path><polyline points="16 6 12 2 8 6"></polyline><line x1="12" y1="2" x2="12" y2="15"></line></svg>
        </span>
        <span class="share-text">Share to feed</span>
      </button>

      <button class="share-item" @click="$emit('action', 'message')">
        <span class="share-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
        </span>
        <span class="share-text">Send as a message</span>
      </button>

      <button class="share-item" @click="$emit('action', 'telegram')">
        <span class="share-icon telegram-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.562 8.161c-.18 1.897-.962 6.502-1.359 8.627-.168.9-.5 1.201-.82 1.23-.697.064-1.228-.46-1.903-.903-1.056-.693-1.653-1.124-2.678-1.8-1.185-.78-.417-1.21.258-1.91.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.48.33-.913.49-1.302.48-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.83-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635.099-.002.321.023.465.141.119.098.152.228.163.331.011.103.024.337.006.522z"/></svg>
        </span>
        <span class="share-text">Telegram</span>
      </button>

      <button class="share-item" @click="$emit('action', 'whatsapp')">
        <span class="share-icon whatsapp-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path d="M.057 24l1.687-6.163c-1.041-1.804-1.588-3.849-1.587-5.946.003-6.556 5.338-11.891 11.893-11.891 3.181.001 6.167 1.24 8.413 3.488 2.245 2.248 3.481 5.236 3.48 8.414-.003 6.557-5.338 11.892-11.893 11.892-1.99-.001-3.951-.5-5.688-1.448l-6.305 1.654zm6.597-3.807c1.676.995 3.276 1.591 5.392 1.592 5.448 0 9.886-4.434 9.889-9.885.002-5.462-4.415-9.89-9.881-9.892-5.452 0-9.887 4.434-9.889 9.884-.001 2.225.651 3.891 1.746 5.634l-.999 3.648 3.742-.981zm11.387-5.464c-.074-.124-.272-.198-.57-.347-.297-.149-1.758-.868-2.031-.967-.272-.099-.47-.149-.669.149-.198.297-.768.967-.941 1.165-.173.198-.347.223-.644.074-.297-.149-1.255-.462-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.297-.347.446-.521.151-.172.2-.296.3-.495.099-.198.05-.372-.025-.521-.075-.124-.669-1.612-.916-2.206-.242-.579-.487-.501-.669-.51l-.57-.01c-.198 0-.52.074-.792.372s-1.04 1.016-1.04 2.479 1.065 2.876 1.213 3.074c.149.198 2.095 3.2 5.076 4.487.709.306 1.263.489 1.694.626.712.226 1.36.194 1.872.118.571-.085 1.758-.719 2.006-1.413.248-.695.248-1.29.173-1.414z"/></svg>
        </span>
        <span class="share-text">Whatsapp</span>
      </button>

      <button class="share-item" @click="$emit('action', 'copy')">
        <span class="share-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>
        </span>
        <span class="share-text">Copy link</span>
      </button>
    </div>
  </div>
</template>

<script setup>
defineEmits(['action', 'close'])
</script>

<style scoped>
.share-wrapper {
  position: absolute;
  bottom: 100%;
  right: 0;
  padding-bottom: 10px;
  z-index: 100;
}

.share-popup {
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

.share-item {
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

.share-item:hover {
  background-color: #f1f5f9;
  transform: translateX(3px);
}

.share-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  color: #64748b;
  width: 24px;
  height: 24px;
  transition: color 0.2s ease;
}

.share-item:hover .share-icon {
  color: #1976d2;
}

.telegram-icon {
  color: #0ea5e9;
}

.whatsapp-icon {
  color: #22c55e;
}

.share-text {
  color: #334155;
  font-size: 13.5px;
  font-weight: 500;
}

.share-item:hover .share-text {
  color: #1976d2;
  font-weight: 600;
}
</style>