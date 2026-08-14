<template>
  <div class="voice-recorder-wrapper">
    <!-- ប៊ូតុងបិទ (Close 'X') -->
    <button class="voice-action-btn" @click="handleClose" title="Cancel">
      <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
        <line x1="18" y1="6" x2="6" y2="18"></line>
        <line x1="6" y1="6" x2="18" y2="18"></line>
      </svg>
    </button>

    <!-- ប៊ូតុង Play / Pause -->
    <button class="voice-play-pause-btn" @click="togglePlay" title="Play/Pause">
      <svg v-if="isPlaying" width="16" height="16" viewBox="0 0 24 24" fill="currentColor">
        <rect x="6" y="4" width="4" height="16"></rect>
        <rect x="14" y="4" width="4" height="16"></rect>
      </svg>
      <svg v-else width="16" height="16" viewBox="0 0 24 24" fill="currentColor">
        <polygon points="5 3 19 12 5 21 5 3"></polygon>
      </svg>
    </button>

    <!-- Waveform SVG Container (ប្តូររាងខ្សែកោងទៅជារាងកោងដូចសញ្ញា Wi-Fi) -->
    <div class="voice-waveform-container">
      <div class="waveform-display">
        <svg width="100%" height="100%" viewBox="0 0 1000 100" preserveAspectRatio="none" version="1.1" xmlns="http://www.w3.org/2000/svg">
          <defs>
            <!-- Pattern សម្រាប់ផ្នែក Active (ពណ៌ #1B75D2 រាងកោងដូច Wi-Fi) -->
            <pattern id="active-wave" x="0" y="0" width="30" height="100" patternUnits="userSpaceOnUse">
              <!-- កែសម្រួល path ឱ្យទៅជារាងកោងធំៗដូចសញ្ញា Wi-Fi -->
              <path d="M0,70 A40,40 0 0,1 40,70" fill="none" stroke="#1B75D2" stroke-width="4" stroke-linecap="round"/>
              <path d="M-10,55 A55,55 0 0,1 50,55" fill="none" stroke="#1B75D2" stroke-width="4" stroke-linecap="round"/>
              <path d="M-20,40 A70,70 0 0,1 60,40" fill="none" stroke="#1B75D2" stroke-width="4" stroke-linecap="round"/>
            </pattern>

            <!-- Pattern សម្រាប់ផ្នែក Inactive (ពណ៌ប្រផេះចាង រាងកោងដូច Wi-Fi) -->
            <pattern id="inactive-wave" x="0" y="0" width="30" height="100" patternUnits="userSpaceOnUse">
              <path d="M0,70 A40,40 0 0,1 40,70" fill="none" stroke="#d1d5db" stroke-width="4" stroke-linecap="round"/>
              <path d="M-10,55 A55,55 0 0,1 50,55" fill="none" stroke="#d1d5db" stroke-width="4" stroke-linecap="round"/>
              <path d="M-20,40 A70,70 0 0,1 60,40" fill="none" stroke="#d1d5db" stroke-width="4" stroke-linecap="round"/>
            </pattern>
          </defs>

          <!-- ផ្នែក Inactive (បង្ហាញពេញ 100%) -->
          <rect x="0" y="0" width="1000" height="100" fill="url(#inactive-wave)" />

          <!-- ផ្នែក Active (កាត់ទទឹងត្រឹម 40%) -->
          <rect x="0" y="0" width="400" height="100" fill="url(#active-wave)" />
        </svg>
      </div>
    </div>

    <!-- រយៈពេល (Duration) -->
    <div class="voice-duration-container">
      <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" class="voice-clock-icon">
        <circle cx="12" cy="12" r="10"></circle>
        <polyline points="12 6 12 12 16 14"></polyline>
      </svg>
      <span class="voice-duration">0:02</span>
    </div>

    <!-- ប៊ូតុង Send -->
    <button class="voice-send-btn" @click="handleSend" title="Send Voice">
      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
        <line x1="22" y1="2" x2="11" y2="13"></line>
        <polygon points="22 2 15 22 11 13 2 9 22 2"></polygon>
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['close', 'send'])
const isPlaying = ref(false)

function togglePlay() {
  isPlaying.value = !isPlaying.value
}

function handleClose() {
  emit('close')
}

function handleSend() {
  emit('send')
}
</script>

<style scoped>
.voice-recorder-wrapper {
  display: flex;
  align-items: center;
  gap: 12px;
  background: #ffffff;
  width: 100%;
  box-sizing: border-box;
  /* padding: 8px 12px; */
}

.voice-action-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #65676b;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 4px;
}

.voice-play-pause-btn {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: #f0f2f5;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #050505;
}

.voice-waveform-container {
  display: flex;
  align-items: center;
  flex: 1;
  height: 30px;
}

.waveform-display {
  width: 100%;
  height: 100%;
  background-color: transparent;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.voice-duration-container {
  display: flex;
  align-items: center;
  gap: 5px;
}

.voice-clock-icon {
  color: #65676b;
}

.voice-duration {
  font-size: 13px;
  color: #65676b;
  font-weight: 500;
}

.voice-send-btn {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: #1B75D2;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #ffffff;
}
</style>