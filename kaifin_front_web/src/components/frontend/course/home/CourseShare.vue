<template>
  <div class="course-share-container">
    <div class="share-buttons-wrapper">
      <!-- General Share / Native Share -->
      <button class="share-icon-btn native-share" title="Share" @click="handleNativeShare">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"></path>
          <polyline points="16 6 12 2 8 6"></polyline>
          <line x1="12" y1="2" x2="12" y2="15"></line>
        </svg>
      </button>

      <!-- Facebook Share -->
      <button class="share-icon-btn facebook-share" title="Share on Facebook" @click="shareToFacebook">
        <svg viewBox="0 0 24 24" fill="currentColor">
          <path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"></path>
        </svg>
      </button>

      <!-- Telegram Share -->
      <button class="share-icon-btn telegram-share" title="Share on Telegram" @click="shareToTelegram">
        <svg viewBox="0 0 24 24" fill="currentColor">
          <path d="M21.5 4.5L2.5 12l6.5 2.5L21 6l-10 10.5v4.5l3.5-3.5 4 3L21.5 4.5z"></path>
        </svg>
      </button>

      <!-- WhatsApp Share -->
      <button class="share-icon-btn whatsapp-share" title="Share on WhatsApp" @click="shareToWhatsApp">
        <svg viewBox="0 0 24 24" fill="currentColor">
          <path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"></path>
        </svg>
      </button>

      <!-- Copy Link -->
      <button class="share-icon-btn copy-share" title="Copy Link" @click="copyLink">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
          <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  shareUrl: {
    type: String,
    default: () => window.location.href
  },
  shareTitle: {
    type: String,
    default: 'Check out this course!'
  }
})

const handleNativeShare = () => {
  if (navigator.share) {
    navigator.share({
      title: props.shareTitle,
      url: props.shareUrl,
    }).catch(() => {})
  } else {
    copyLink()
  }
}

const shareToFacebook = () => {
  const url = `https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(props.shareUrl)}`
  window.open(url, '_blank', 'width=600,height=400')
}

const shareToTelegram = () => {
  const url = `https://t.me/share/url?url=${encodeURIComponent(props.shareUrl)}&text=${encodeURIComponent(props.shareTitle)}`
  window.open(url, '_blank', 'width=600,height=400')
}

const shareToWhatsApp = () => {
  const url = `https://api.whatsapp.com/send?text=${encodeURIComponent(props.shareTitle + ' ' + props.shareUrl)}`
  window.open(url, '_blank', 'width=600,height=400')
}

const copyLink = () => {
  navigator.clipboard.writeText(props.shareUrl).then(() => {
    alert('Link copied to clipboard!')
  }).catch(() => {
    alert('Failed to copy link.')
  })
}
</script>

<style scoped>
.course-share-container {
  display: inline-flex;
  background: #ffffff;
  padding: 8px 12px;
  border-radius: 50px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.08);
  border: 1px solid #e5e7e8;
}

.share-buttons-wrapper {
  display: flex;
  align-items: center;
  gap: 8px;
}

.share-icon-btn {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.2s ease, opacity 0.2s ease;
  color: #ffffff;
}

.share-icon-btn:hover {
  transform: scale(1.08);
  opacity: 0.9;
}

.share-icon-btn svg {
  width: 20px;
  height: 20px;
}

.native-share {
  background-color: #1b75d2;
}

.facebook-share {
  background-color: #1877f2;
}

.telegram-share {
  background-color: #229ed9;
}

.whatsapp-share {
  background-color: #25d366;
}

.copy-share {
  background-color: #4b5563;
}
</style>