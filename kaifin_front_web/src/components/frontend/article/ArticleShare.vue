<script setup>
import { ref } from 'vue'

// Prop emit
const props = defineProps({
  shareUrl: {
    type: String,
    default: () => window.location.href
  }
})

// Sahre picker
const copied = ref(false)
const handleShare = (platform) => {
  if (platform === 'copy') {
    navigator.clipboard.writeText(props.shareUrl).then(() => {
      copied.value = true
      setTimeout(() => { copied.value = false }, 2000)
    })
  } else if (platform === 'facebook') {
    window.open(`https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(props.shareUrl)}`, '_blank')
  } else if (platform === 'telegram') {
    window.open(`https://t.me/share/url?url=${encodeURIComponent(props.shareUrl)}`, '_blank')
  } else if (platform === 'whatsapp') {
    window.open(`https://api.whatsapp.com/send?text=${encodeURIComponent(props.shareUrl)}`, '_blank')
  } else {
    if (navigator.share) {
      navigator.share({ title: 'Share Article', url: props.shareUrl }).catch(() => {})
    }
  }
}
</script>

<template>
  <div class="article-share-container">
    <!-- change to column (Column) -->
    <div class="share-pill-box">
      
      <!-- Native general button -->
      <button class="share-btn general-btn" @click="handleShare('general')" title="Share">
        <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path>
          <polyline points="16 6 12 2 8 6"></polyline>
          <line x1="12" y1="2" x2="12" y2="15"></line>
        </svg>
      </button>

      <!-- Facebook share-->
      <button class="share-btn facebook-btn" @click="handleShare('facebook')" title="Share to Facebook">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
          <path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"></path>
        </svg>
      </button>

      <!-- Telegram share -->
      <button class="share-btn telegram-btn" @click="handleShare('telegram')" title="Share to Telegram">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
          <path d="M21.5 4.5l-19 7 6.5 2.5 11-8.5-8.5 10.5v5l3.5-3.5 4 3 2.5-15.5z"></path>
        </svg>
      </button>

      <!-- Whatup share-->
      <button class="share-btn whatsapp-btn" @click="handleShare('whatsapp')" title="Share to WhatsApp">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"></path>
        </svg>
      </button>

      <!-- Copy link share-->
      <button class="share-btn copy-btn" @click="handleShare('copy')" :title="copied ? 'Copied!' : 'Copy Link'">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
          <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
        </svg>
      </button>

    </div>
  </div>
</template>

<style scoped>
.article-share-container {
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.share-pill-box {
  display: flex;
  flex-direction: column; 
  align-items: center;
  gap: 10px;
  background-color: #ffffff;
  padding: 8px 4px;      
  border-radius: 50px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  border: 1px solid #cbd5e1;
}

.share-btn {
  border: none;
  width: 42px;
  height: 42px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #ffffff;
  transition: transform 0.2s ease, opacity 0.2s ease;
}

.share-btn:hover {
  transform: translateY(-2px);
  opacity: 0.9;
}

.general-btn {
  background-color: #1976d2;
}

.facebook-btn {
  background-color: #1877f2;
}

.telegram-btn {
  background-color: #2ea6ff;
}

.whatsapp-btn {
  background-color: #25d366;
}

.copy-btn {
  background-color: #495057;
}
</style>