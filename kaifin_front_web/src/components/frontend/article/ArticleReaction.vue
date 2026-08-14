<script setup>
import { REACTIONS } from '../../../components/reaction/reactions.js'

const props = defineProps({
  modelValue: {
    type: String,
    default: 'like'
  }
})

const emit = defineEmits(['update:modelValue', 'react'])

const selectReaction = (reaction) => {
  emit('update:modelValue', reaction.key)
  emit('react', reaction)
}
</script>

<template>
  <div class="reactions-bar-container">
    <div 
      v-for="reaction in REACTIONS" 
      :key="reaction.key"
      class="reaction-item"
      :class="{ active: modelValue === reaction.key }"
      @click="selectReaction(reaction)"
    >
      <div class="reaction-icon-wrapper" v-html="reaction.icon"></div>
      
      <!-- Custom Tooltip with #1B75D2 Background -->
      <span class="reaction-tooltip">{{ reaction.label }}</span>

      <span v-if="reaction.private" class="private-lock-badge">
        <svg width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
      </span>
    </div>
  </div>
</template>

<style scoped>
.reactions-bar-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  background-color: #ffffff;
  border-radius: 30px;
  padding: 12px 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  height: max-content;
  width: max-content;
  user-select: none;
  margin-top: 50px;
}

.reaction-item {
  position: relative;
  width: 38px;
  height: 38px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.reaction-item:hover {
  transform: scale(1.25);
}

/* Custom Tooltip Styling with #1B75D2 Color */
.reaction-tooltip {
  position: absolute;
  left: 50px;
  top: 50%;
  transform: translateY(-50%) translateX(-5px);
  background-color: #1B75D2;
  color: #fff;
  padding: 4px 10px;
  font-size: 12px;
  font-weight: 500;
  border-radius: 6px;
  white-space: nowrap;
  pointer-events: none;
  opacity: 0;
  visibility: hidden;
  transition: all 0.2s ease-in-out;
  z-index: 10;
  box-shadow: 0 2px 8px rgba(27, 117, 210, 0.3);
}

.reaction-item:hover .reaction-tooltip {
  opacity: 1;
  visibility: visible;
  transform: translateY(-50%) translateX(0);
}

.reaction-icon-wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.reaction-icon-wrapper :deep(svg) {
  width: 100%;
  height: 100%;
  display: block;
}

.reaction-item:first-child .reaction-icon-wrapper {
  background-color: #444746;
  color: #ffffff;
  border-radius: 50%;
  padding: 8px;
}

.private-lock-badge {
  position: absolute;
  bottom: -2px;
  right: -2px;
  background-color: #fef08a;
  color: #854d0e;
  border-radius: 50%;
  width: 14px;
  height: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 1px 3px rgba(0,0,0,0.2);
}
</style>