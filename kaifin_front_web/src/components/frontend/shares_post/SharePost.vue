<template>
  <div class="share-click-wrap" v-click-outside="closePopup">
    <button class="share-trigger-btn" @click.stop="togglePopup">
      <svg viewBox="0 0 24 24">
        <path d="M4 12v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7"/>
        <path d="M16 6l-4-4-4 4"/>
        <path d="M12 2v14"/>
      </svg>
      {{ formatCount(shareCount) }}
    </button>

    <transition name="pop">
      <div v-if="showPopup" class="share-popup">
        <button
          v-for="opt in options"
          :key="opt.key"
          class="popup-icon-btn"
          :style="{ background: opt.bg }"
          @click="onQuickShare(opt.key)"
        >
          <svg viewBox="0 0 24 24" v-html="opt.svg"></svg>
          <span class="popup-tooltip">{{ opt.label }}</span>
        </button>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  post: { type: Object, default: null },
  shareCount: { type: Number, default: 0 },
})

const emit = defineEmits(['share'])

const showPopup = ref(false)
const copied = ref(false)

const vClickOutside = {
  mounted(el, binding) {
    el.__clickOutsideHandler = (e) => {
      if (!el.contains(e.target)) binding.value()
    }
    document.addEventListener('click', el.__clickOutsideHandler)
  },
  unmounted(el) {
    document.removeEventListener('click', el.__clickOutsideHandler)
  },
}

function formatCount(n) {
  if (n >= 1000000) return (n / 1000000).toFixed(1) + 'M'
  if (n >= 1000) return (n / 1000).toFixed(1) + 'K'
  return String(n)
}

function togglePopup() {
  showPopup.value = !showPopup.value
}

function closePopup() {
  showPopup.value = false
}

async function onQuickShare(channel) {
  if (channel === 'copy') {
    const link = props.post?.url || window.location.href
    try {
      await navigator.clipboard.writeText(link)
      copied.value = true
      setTimeout(() => (copied.value = false), 1200)
    } catch (err) {
      console.error('Copy failed', err)
    }
  } else {
    const link = encodeURIComponent(props.post?.url || window.location.href)
    const text = encodeURIComponent(props.post?.caption || '')
    const deepLinks = {
      facebook: `https://www.facebook.com/sharer/sharer.php?u=${link}`,
      telegram: `https://t.me/share/url?url=${link}&text=${text}`,
      whatsapp: `https://wa.me/?text=${text}%20${link}`,
      tiktok: props.post?.url || window.location.href,
    }
    if (deepLinks[channel]) {
      window.open(deepLinks[channel], '_blank', 'noopener,noreferrer')
    }
  }

  showPopup.value = false
  emit('share', { channel, post: props.post })
}

const options = [
  {
    key: 'facebook',
    label: 'Facebook',
    bg: '#1877F2',
    svg: '<path d="M14 9h3V6h-3c-1.7 0-3 1.3-3 3v2H9v3h2v6h3v-6h2.5l.5-3H14V9.4c0-.2.2-.4.5-.4H14Z" fill="#fff"/>',
  },
  {
    key: 'tiktok',
    label: 'TikTok',
    bg: '#000',
    svg: '<path d="M15 3c.4 1.8 1.7 3.1 3.5 3.4V9c-1.3 0-2.5-.4-3.5-1.1V15a5 5 0 1 1-4-4.9v2.7a2.3 2.3 0 1 0 1.6 2.2V3h2.4Z" fill="#fff"/>',
  },
  {
    key: 'telegram',
    label: 'Telegram',
    bg: '#29A9EA',
    svg: '<path d="m3 11 17-7-3 16-6-4-3 3-1-5 9-8-11 6Z" fill="#fff"/>',
  },
  {
    key: 'whatsapp',
    label: 'Whatsapp',
    bg: '#25D366',
    svg: '<path d="M12 3a9 9 0 0 0-7.8 13.5L3 21l4.7-1.2A9 9 0 1 0 12 3Z" stroke="#fff" stroke-width="1.6" fill="none"/><path d="M8.5 8.5c.3 3 2 4.7 5 5l1-1.3c.5-.6 1-.4 1.5-.1l1.5 1c.3.6-.1 1.7-.8 2.1-1 .6-2.3.4-4-.4-2.3-1.1-3.9-2.7-5-5-.8-1.7-1-3-.4-4 .4-.7 1.5-1.1 2.1-.8l1 1.5c.3.5.5 1-.1 1.5l-1.3 1Z" fill="#fff"/>',
  },
  {
    key: 'copy',
    label: 'Copy link',
    bg: '#5F5E5A',
    svg: '<rect x="9" y="9" width="12" height="12" rx="2" stroke="#fff" stroke-width="1.8" fill="none"/><path d="M5 15V5a2 2 0 0 1 2-2h10" stroke="#fff" stroke-width="1.8" fill="none" stroke-linecap="round"/>',
  },
]
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.share-click-wrap {
  position: relative;
  display: inline-flex;
}

.stat-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  border: none;
  background: transparent;
  cursor: pointer;
  font-size: 13px;
  font-weight: 500;
  color: #5f5e5a;
  padding: 6px 8px;
  border-radius: 999px;
  font-family: 'Inter', sans-serif;
}

.stat-btn:hover {
  background: #f2f2f3;
}

.stat-btn svg {
  width: 18px;
  height: 18px;
}

.share-popup {
  position: absolute;
  bottom: calc(100% + 10px);
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  align-items: center;
  gap: 6px;
  background: #fff;
  border: 1px solid #e7e7e7;
  border-radius: 999px;
  padding: 6px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  z-index: 30;
}

.share-popup::after {
  content: '';
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  border: 7px solid transparent;
  border-top-color: #fff;
}

.popup-icon-btn {
  position: relative;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.15s;
  flex-shrink: 0;
}

.popup-icon-btn:hover {
  transform: translateY(-4px) scale(1.08);
}

.popup-icon-btn svg {
  width: 20px;
  height: 20px;
  pointer-events: none;
}

.popup-tooltip {
  position: absolute;
  bottom: calc(100% + 8px);
  left: 50%;
  transform: translateX(-50%);
  background: #1a1a1a;
  color: #fff;
  font-size: 11px;
  font-weight: 500;
  padding: 4px 8px;
  border-radius: 6px;
  white-space: nowrap;
  opacity: 0;
  pointer-events: none;
  transition: opacity 0.15s;
}

.popup-icon-btn:hover .popup-tooltip {
  opacity: 1;
}

.pop-enter-active,
.pop-leave-active {
  transition: opacity 0.15s ease, transform 0.15s ease;
}

.pop-enter-from,
.pop-leave-to {
  opacity: 0;
  transform: translateX(-50%) translateY(6px);
}
</style>