<script setup>
import { ref } from 'vue'

defineProps({
  aiQuery: {
    type: String,
    required: true
  },
  aiMessages: {
    type: Array,
    required: true
  }
})

defineEmits(['update:aiQuery', 'send', 'quick-ask', 'clear-chat'])

const isRecording = ref(false)

const toggleVoiceRecording = () => {
  isRecording.value = !isRecording.value
}

// Suggested topics managed dynamically via an array with SVG paths
const suggestionTopics = ref([
  {
    action: 'Summarize this lesson',
    title: 'Summarize Lesson',
    desc: 'Get a quick breakdown of key concepts'
  },
  {
    action: 'Explain this code',
    title: 'Explain Code',
    desc: 'Understand syntax & logic step-by-step'
  },
  {
    action: 'Generate a quiz',
    title: 'Test Knowledge',
    desc: 'Generate interactive review quizzes'
  }
])
</script>

<template>
  <div class="ai-assistant-wrapper">
    <!-- Header Section -->
    <div class="ai-header">
      <div class="ai-title-wrapper">
        <div class="ai-bot-avatar">
          <img src="../../../assets/logos/kaifin_l2.png" alt="Logo" class="header-logo-img">
        </div>
        <div class="ai-header-text">
          <h3 class="ai-title">AI Learning Assistant</h3>
          <span class="ai-subtitle">Your personal intelligent tutor</span>
        </div>
      </div>
      <div class="ai-header-actions">
        <button class="action-icon-btn" title="Clear Chat" @click="$emit('clear-chat')">
          <svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
          </svg>
        </button>
      </div>
    </div>

    <!-- Chat Box Messages Area -->
    <div class="chat-messages-box">
      <!-- Voice Listening Center View inside Message List Area -->
      <div v-if="isRecording" class="voice-center-container">
        <div class="centered-ai-avatar">
          <img src="../../../assets/logos/kaifin_l2.png" alt="AI Logo" class="centered-avatar-img">
        </div>
        <div class="listening-text-badge">
          <span class="pulse-dot"></span>
          AI is listening...
        </div>
      </div>

      <!-- Normal Messages List -->
      <div v-else class="messages-list">
        <!-- Suggested Questions -->
        <div v-if="aiMessages.length <= 1" class="suggestions-container">
          <div class="suggestions-prompt">
            <svg class="prompt-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"/>
            </svg> Recommended Starters
          </div>
          <div class="suggestion-cards-grid">
            <button 
              v-for="(suggestion, sIndex) in suggestionTopics" 
              :key="sIndex" 
              class="suggestion-card" 
              @click="$emit('quick-ask', suggestion.action)"
            >
              <div class="card-icon-box">
                <svg v-if="sIndex === 0" class="card-svg" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
                </svg>
                <svg v-else-if="sIndex === 1" class="card-svg" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/>
                </svg>
                <svg v-else class="card-svg" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
              </div>
              <div class="card-content">
                <span class="card-title">{{ suggestion.title }}</span>
                <span class="card-desc">{{ suggestion.desc }}</span>
              </div>
              <div class="card-arrow">
                <svg width="14" height="14" fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7"/>
                </svg>
              </div>
            </button>
          </div>
        </div>

        <!-- Chat Message Rows -->
        <div 
          v-for="(msg, index) in aiMessages" 
          :key="index" 
          :class="['message-row', msg.sender === 'user' ? 'user-row' : 'ai-row']"
        >
          <!-- Message Bubble Container -->
          <div :class="['message-bubble', msg.sender === 'user' ? 'user-message' : 'ai-message']">
            <div class="message-content-inner">
              <div class="message-header-info">
                <span class="message-sender-name">
                  {{ msg.sender === 'user' ? 'You' : 'AI Assistant' }}
                </span>
                <span class="message-time-badge">
                  <svg class="time-clock-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
                  </svg>
                  12:50
                </span>
              </div>
              <div class="message-text">{{ msg.text }}</div>
            </div>

            <!-- Embedded Avatar inside AI Message Bubble (Left side) -->
            <div v-if="msg.sender !== 'user'" class="embedded-ai-avatar">
              <img src="https://api.iconify.design/fluent-emoji:robot.svg" alt="AI Avatar" class="embedded-avatar-img" />
            </div>

            <!-- Embedded Avatar inside User Message Bubble (Right side) -->
            <div v-if="msg.sender === 'user'" class="embedded-user-avatar">
              <img 
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQviRBCXJg2bQzezRkEvJn1WPzCRfqWImgvCxdf6Za_A5tGlGilZpuM6Kk&s=10" 
                alt="User Avatar" 
                class="embedded-avatar-img"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Input Area -->
      <div class="ai-input-container">
        <button 
          class="voice-btn" 
          :class="{ recording: isRecording }"
          @click="toggleVoiceRecording"
          title="Voice input"
        >
          <svg class="mic-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4m-4-8a3 3 0 01-3-3V5a3 3 0 116 0v6a3 3 0 01-3 3z"/>
          </svg>
        </button>

        <input 
          type="text" 
          :value="aiQuery"
          @input="$emit('update:aiQuery', $event.target.value)"
          @keyup.enter="$emit('send')"
          placeholder="Ask anything or search course content..." 
          class="ai-input-field" 
        />
        
        <button class="send-btn" @click="$emit('send')" title="Send Message">
          <svg class="send-icon" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M5 12h14M12 5l7 7-7 7"/>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.ai-assistant-wrapper {
  background: #ffffff;
  border: none;
  border-radius: 20px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  box-shadow: none;
  width: 100%;
  box-sizing: border-box;
  margin: 16px 0;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
}

.ai-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 4px;
}

.ai-title-wrapper {
  display: flex;
  align-items: center;
  gap: 12px;
}

.ai-bot-avatar {
  width: 42px;
  height: 42px;
  background: transparent;
  border: none;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.header-logo-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.ai-header-text {
  display: flex;
  flex-direction: column;
  gap: 1px;
}

.ai-title {
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
  letter-spacing: -0.2px;
}

.ai-subtitle {
  font-size: 12px;
  color: #64748b;
  font-weight: 400;
}

.action-icon-btn {
  background: #f8fafc;
  border: none;
  border-radius: 10px;
  width: 34px;
  height: 34px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #64748b;
  cursor: pointer;
  transition: all 0.2s ease;
}

.action-icon-btn:hover {
  background: #fee2e2;
  color: #dc2626;
}

.action-icon-btn svg {
  width: 15px;
  height: 15px;
}

.chat-messages-box {
  background: #f8fafc;
  border: none;
  border-radius: 16px;
  padding: 16px;
  height: 440px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  position: relative;
}

/* Voice Center View inside message box */
.voice-center-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 16px;
  animation: fadeIn 0.25s ease;
}

.centered-ai-avatar {
  width: 88px;
  height: 88px;
  background: #ffffff;
  border-radius: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.06);
  padding: 16px;
}

.centered-avatar-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.listening-text-badge {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #475569;
  font-size: 14px;
  font-weight: 600;
}

.pulse-dot {
  width: 8px;
  height: 8px;
  background-color: #22c55e;
  border-radius: 50%;
  animation: pulse 1.5s infinite;
}

@keyframes pulse {
  0% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(34, 197, 94, 0.7);
  }
  70% {
    transform: scale(1);
    box-shadow: 0 0 0 6px rgba(34, 197, 94, 0);
  }
  100% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(34, 197, 94, 0);
  }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.messages-list {
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 16px;
  flex: 1;
  padding-right: 4px;
  scroll-behavior: smooth;
}

.suggestions-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 8px;
}

.suggestions-prompt {
  font-size: 12.5px;
  font-weight: 600;
  color: #64748b;
  display: flex;
  align-items: center;
  gap: 6px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.prompt-icon {
  width: 15px;
  height: 15px;
  color: #eab308;
}

.suggestion-cards-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 10px;
}

.suggestion-card {
  background: transparent;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 10px 14px;
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  text-align: left;
  position: relative;
}

.suggestion-card:hover {
  background: #f0f7ff;
  border-color: #cbd5e1;
  transform: translateY(-1px);
}

.card-icon-box {
  width: 32px;
  height: 32px;
  background: transparent;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  color: #64748b;
}

.card-svg {
  width: 20px;
  height: 20px;
}

.card-content {
  display: flex;
  flex-direction: column;
  gap: 1px;
  flex: 1;
}

.card-title {
  font-size: 13px;
  font-weight: 600;
  color: #1e293b;
}

.card-desc {
  font-size: 11px;
  color: #64748b;
}

.card-arrow {
  color: #cbd5e1;
  transition: transform 0.2s ease;
}

.suggestion-card:hover .card-arrow {
  color: #1976D2;
  transform: translateX(2px);
}

.message-row {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  max-width: 88%;
}

.ai-row {
  align-self: flex-start;
  flex-direction: row;
  align-items: center;
}

.user-row {
  align-self: flex-end;
  flex-direction: row-reverse;
  align-items: center;
}

.message-bubble {
  padding: 10px 14px;
  font-size: 13px;
  line-height: 1.45;
  word-break: break-word;
  flex: 1;
  position: relative;
}

/* AI Message Bubble Styling */
.ai-message {
  background: #1976D2;
  color: #ffffff;
  border-radius: 16px 16px 16px 4px;
  box-shadow: 0 4px 12px rgba(25, 118, 210, 0.2);
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 14px;
}

.embedded-ai-avatar {
  width: 36px;
  height: 36px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.embedded-avatar-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.ai-message .message-sender-name,
.ai-message .message-time-badge {
  color: rgba(255, 255, 255, 0.85);
}

.ai-message .time-clock-icon {
  color: rgba(255, 255, 255, 0.85);
}

/* User Message Bubble Styling */
.user-message {
  background: #ffffff;
  color: #1e293b;
  border: none;
  border-radius: 16px 16px 4px 16px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.02);
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 14px;
}

.embedded-user-avatar {
  width: 36px;
  height: 36px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  overflow: hidden;
  order: 2;
}

.embedded-user-avatar .embedded-avatar-img {
  object-fit: cover;
}

.user-message .message-sender-name {
  color: #475569;
}

.user-message .message-time-badge {
  color: #94a3b8;
}

.user-message .time-clock-icon {
  color: #94a3b8;
}

.message-content-inner {
  flex: 1;
  display: flex;
  flex-direction: column;
  order: 1;
}

.message-header-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2px;
  gap: 16px;
}

.message-sender-name {
  font-size: 11.5px;
  font-weight: 600;
}

.message-time-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-size: 10px;
  font-weight: 500;
  background: rgba(0, 0, 0, 0.05);
  padding: 1px 6px;
  border-radius: 6px;
}

.ai-message .message-time-badge {
  background: rgba(255, 255, 255, 0.15);
}

.time-clock-icon {
  width: 11px;
  height: 11px;
}

.ai-input-container {
  display: flex;
  align-items: center;
  background: #ffffff;
  border: none;
  border-radius: 16px;
  padding: 6px 8px;
  gap: 8px;
  margin-top: 10px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.03);
}

.ai-input-container:focus-within {
  box-shadow: 0 0 0 2px rgba(25, 118, 210, 0.15);
}

.voice-btn {
  background: #f8fafc;
  border: none;
  border-radius: 10px;
  width: 34px;
  height: 34px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #64748b;
  cursor: pointer;
  flex-shrink: 0;
  transition: all 0.2s ease;
}

.voice-btn:hover {
  background: #f1f5f9;
  color: #0f172a;
}

.voice-btn.recording {
  background: #fee2e2;
  color: #dc2626;
}

.mic-icon {
  width: 16px;
  height: 16px;
}

.ai-input-field {
  background: transparent;
  border: none;
  outline: none;
  color: #0f172a;
  font-size: 13.5px;
  width: 100%;
  padding: 4px 0;
}

.ai-input-field::placeholder {
  color: #94a3b8;
}

.send-btn {
  background: #1976D2;
  border: none;
  border-radius: 10px;
  color: #ffffff;
  width: 34px;
  height: 34px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  flex-shrink: 0;
  transition: background 0.2s ease, transform 0.1s ease;
}

.send-btn:hover {
  background: #1565C0;
}

.send-btn:active {
  transform: scale(0.95);
}

.send-icon {
  width: 15px;
  height: 15px;
}
</style>