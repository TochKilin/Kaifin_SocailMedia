<template>
  <div class="reaction-wrapper" @mouseleave="$emit('close')">
    <div class="reaction-popup">
      <button 
        v-for="(reaction, rIdx) in reactions" 
        :key="rIdx" 
        class="reaction-icon-btn"
        @click.stop="$emit('select', reaction)"
        :data-tooltip="reaction.name"
      >
        <!-- ១. បើជា Emoji -->
        <span v-if="reaction.type === 'emoji'" class="emoji">{{ reaction.content }}</span>

        <!-- ២. បើជា Image ឬ SVG File -->
        <img v-else-if="reaction.type === 'image'" :src="reaction.content" class="reaction-img" :alt="reaction.name" />

        <!-- ៣. បើជា SVG Raw Code -->
        <span v-else-if="reaction.type === 'svg'" v-html="reaction.content" class="reaction-svg"></span>
      </button>
    </div>
  </div>
</template>

<script setup>
defineProps({
  reactions: {
    type: Array,
    required: true
  }
})

defineEmits(['select', 'close'])
</script>

<style scoped>
.reaction-wrapper {
  position: absolute;
  bottom: 100%;
  /* left: -200px; */
  right: -130px;
  padding-bottom: 12px;
  z-index: 100;
}

.reaction-popup {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 40px;
  padding: 8px 12px;
  display: flex;
  gap: 10px;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 8px 10px -6px rgba(0, 0, 0, 0.1);
  animation: fadeIn 0.2s cubic-bezier(0.16, 1, 0.3, 1);
}

.reaction-icon-btn {
  background: transparent;
  border: none;
  cursor: pointer;
  padding: 4px;
  border-radius: 50%;
  transition: transform 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  display: flex;
  align-items: center;
  justify-content: center;
  width: 44px;
  height: 44px;
  position: relative; /* ចាំបាច់សម្រាប់កន្លែងដាក់ Tooltip */
}

.reaction-icon-btn:hover {
  transform: scale(1.3);
}

/* ── Custom Tooltip Styles ── */
.reaction-icon-btn::before {
  content: attr(data-tooltip);
  position: absolute;
  bottom: calc(100% + 8px);
  left: 50%;
  transform: translateX(-50%) translateY(4px);
  background-color: #1B75D2; /* ពណ៌ដែលអ្នកចង់បាន */
  color: #ffffff;
  padding: 4px 8px;
  border-radius: 6px;
  font-size: 11px;
  font-weight: 600;
  white-space: nowrap;
  opacity: 0;
  visibility: hidden;
  transition: all 0.2s cubic-bezier(0.16, 1, 0.3, 1);
  pointer-events: none;
  box-shadow: 0 4px 12px rgba(27, 117, 210, 0.2);
}

/* ព្រួញតូចក្រោម Tooltip */
.reaction-icon-btn::after {
  content: '';
  position: absolute;
  bottom: calc(100% + 2px);
  left: 50%;
  transform: translateX(-50%) translateY(4px);
  border-width: 4px;
  border-style: solid;
  border-color: #1B75D2 transparent transparent transparent;
  opacity: 0;
  visibility: hidden;
  transition: all 0.2s cubic-bezier(0.16, 1, 0.3, 1);
  pointer-events: none;
}

/* បង្ហាញ Tooltip ពេល Hover */
.reaction-icon-btn:hover::before,
.reaction-icon-btn:hover::after {
  opacity: 1;
  visibility: visible;
  transform: translateX(-50%) translateY(0);
}

/* ពង្រីកទំហំ Sticker ឱ្យធំច្បាស់ស្អាតល្អ */
.reaction-img, 
.reaction-svg {
  width: 34px;
  height: 34px;
  display: flex;
  align-items: center;
  justify-content: center;
  object-fit: contain;
}

.reaction-svg :deep(svg) {
  width: 100%;
  height: 100%;
}

.emoji {
  font-size: 32px;
  line-height: 1;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(8px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}
</style>