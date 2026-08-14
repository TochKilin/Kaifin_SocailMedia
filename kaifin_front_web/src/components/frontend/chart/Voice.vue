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

    <!-- Waveform SVG Container (ប្រើ Pattern ជាប់គ្នា និង Dynamic Active Width តាមសំឡេងពិត) -->
    <div class="voice-waveform-container">
      <div class="waveform-display">
        <svg width="100%" height="100%" viewBox="0 0 1000 100" preserveAspectRatio="none" version="1.1" xmlns="http://www.w3.org/2000/svg">
          <defs>
            <!-- Pattern គ្រាប់មូលពណ៌ខៀវ (Active) -->
            <pattern id="active-wave" x="0" y="0" width="24" height="100" patternUnits="userSpaceOnUse">
              <rect x="0" y="20" width="24" height="60" rx="12" fill="#1B75D2" />
            </pattern>

            <!-- Pattern គ្រាប់មូលពណ៌ប្រផេះ (Inactive) -->
            <pattern id="inactive-wave" x="0" y="0" width="24" height="100" patternUnits="userSpaceOnUse">
              <rect x="0" y="20" width="24" height="60" rx="12" fill="#d1d5db" />
            </pattern>
          </defs>

          <!-- ផ្នែក Inactive (បង្ហាញពេញទំហឹងនៅខាងក្រោយ) -->
          <rect x="0" y="0" width="1000" height="100" fill="url(#inactive-wave)" />

          <!-- ផ្នែក Active (ទទឹងរត់ប្រែប្រួលតាមកម្រិតសំឡេង ឬ Playback ពីឆ្វេងទៅស្តាំ) -->
          <rect x="0" y="0" :width="activeWidth" height="100" fill="url(#active-wave)" />
        </svg>
      </div>
    </div>

    <!-- រយៈពេល (Duration) -->
    <div class="voice-duration-container">
      <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" class="voice-clock-icon">
        <circle cx="12" cy="12" r="10"></circle>
        <polyline points="12 6 12 12 16 14"></polyline>
      </svg>
      <span class="voice-duration">{{ formattedDuration }}</span>
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
import { ref, computed, onMounted, onUnmounted } from 'vue'

const emit = defineEmits(['close', 'send'])

const isPlaying = ref(false)
const durationSeconds = ref(0)
const activeWidth = ref(0) // ទទឹងគិតជា Pixel (ពី 0 ដល់ 1000) សម្រាប់គ្របដណ្តប់ Pattern ពណ៌ខៀវ

let mediaRecorder = null
let audioChunks = []
let audioElement = null
let timerInterval = null

let audioContext = null
let analyser = null
let microphoneStream = null
let animationFrameId = null

const formattedDuration = computed(() => {
  const mins = Math.floor(durationSeconds.value / 60)
  const secs = durationSeconds.value % 60
  return `${mins}:${secs < 10 ? '0' : ''}${secs}`
})

onMounted(async () => {
  try {
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
    microphoneStream = stream
    mediaRecorder = new MediaRecorder(stream)

    audioContext = new (window.AudioContext || window.webkitAudioContext)()
    analyser = audioContext.createAnalyser()
    const source = audioContext.createMediaStreamSource(stream)
    source.connect(analyser)
    analyser.fftSize = 64

    const dataArray = new Uint8Array(analyser.frequencyBinCount)

    const updateWaveform = () => {
      if (!analyser) return
      analyser.getByteFrequencyData(dataArray)

      let sum = 0
      for (let i = 0; i < dataArray.length; i++) {
        sum += dataArray[i]
      }
      const average = sum / dataArray.length
      
      // បំប្លែងកម្រិតសំឡេងពិត (average) ឱ្យទៅជាទទឹង pixel ក្នុង viewBox 1000
      // ធានាថាវា dynamically រត់ឡើងចុះតាមការនិយាយ
      const targetWidth = (average / 128) * 1000
      activeWidth.value = Math.min(1000, Math.max(50, targetWidth))

      animationFrameId = requestAnimationFrame(updateWaveform)
    }

    updateWaveform()

    mediaRecorder.ondataavailable = (event) => {
      if (event.data.size > 0) audioChunks.push(event.data)
    }

    mediaRecorder.onstop = () => {
      const audioBlob = new Blob(audioChunks, { type: 'audio/webm' })
      audioElement = new Audio(URL.createObjectURL(audioBlob))

      audioElement.ontimeupdate = () => {
        if (audioElement.duration) {
          // ពេល Playback គឺឱ្យ activeWidth រត់ពីឆ្វេងទៅស្តាំស្របតាមវិនាទីសំឡេង
          activeWidth.value = (audioElement.currentTime / audioElement.duration) * 1000
        }
      }

      audioElement.onended = () => {
        isPlaying.value = false
        activeWidth.value = 0
      }
    }

    mediaRecorder.start()
    timerInterval = setInterval(() => { durationSeconds.value++ }, 1000)

  } catch (error) {
    console.error('Mic permission error:', error)
    alert('សូមអនុញ្ញាតឱ្យប្រើប្រាស់មីក្រូហ្វូនជាមុនសិន!')
  }
})

onUnmounted(() => {
  stopAll()
})

function stopAll() {
  clearInterval(timerInterval)
  cancelAnimationFrame(animationFrameId)
  if (mediaRecorder && mediaRecorder.state !== 'inactive') mediaRecorder.stop()
  if (microphoneStream) microphoneStream.getTracks().forEach(track => track.stop())
  if (audioContext && audioContext.state !== 'closed') audioContext.close()
  if (audioElement) audioElement.pause()
}

function togglePlay() {
  if (!audioElement) return

  if (isPlaying.value) {
    audioElement.pause()
    isPlaying.value = false
  } else {
    if (mediaRecorder && mediaRecorder.state === 'recording') {
      mediaRecorder.stop()
      clearInterval(timerInterval)
      cancelAnimationFrame(animationFrameId)
    }
    audioElement.play()
    isPlaying.value = true
  }
}

function handleClose() {
  stopAll()
  emit('close')
}

function handleSend() {
  stopAll()
  const audioBlob = new Blob(audioChunks, { type: 'audio/webm' })
  emit('send', audioBlob)
}
</script>

<style scoped>
.voice-recorder-wrapper {
  display: flex;
  align-items: center;
  gap: 12px;
  /* background: #ffffff; */
  width: 100%;
  box-sizing: border-box;
  /* background-color: red; */
  /* position: absolute; */
  /* padding: 8px 12px; */
  /* border-radius: 8px; */
  /* box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08); */
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
  height: 36px;
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
  min-width: 32px;
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