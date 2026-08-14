<!-- File: CourseWatchPage.vue -->
<script setup>
import { ref } from 'vue'
import AIAssistant from './AIAssistant.vue'
import LessonSection from './LessonSection.vue'
import Overview from './Overview.vue'
import QnAManagement from './QnAManagement.vue'
import Note from './Note.vue'
import Annount from './Annount.vue'
import Review from './Review.vue'
import LearningTool from './LearningTool.vue' // 1. នាំចូល LearningTool Component

defineProps({
  course: {
    type: Object,
    required: true,
    default: () => ({
      title: 'AI & Machine Learning Mastery',
      description: 'Learn cutting-edge AI techniques from industry experts with hands-on projects.',
      mediaUrl: 'https://images.unsplash.com/photo-1517694712202-14dd9538aa97?q=80&w=1200&auto=format&fit=crop'
    })
  }
})

defineEmits(['back'])

const activeSidebarTab = ref('Lesson')
const activeContentTab = ref('Content Course')

const handleSwitcherClick = (tabName) => {
  activeContentTab.value = tabName
  if (tabName === 'AI Assistance') {
    activeSidebarTab.value = 'AI Assistant'
  } else if (tabName === 'Reviews') {
    activeSidebarTab.value = 'Reviews'
  }
}

const sidebarTabs = [
  { name: 'Lesson', icon: 'M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253' },
  { name: 'AI Assistant', icon: 'M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z' },
  { name: 'Overview', icon: 'M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012-2m-6 9l2 2 4-4' },
  { name: 'Q&A', icon: 'M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z' },
  { name: 'Notes', icon: 'M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z' },
  { name: 'Announcement', icon: 'M11 5.882V19.24a1.76 1.76 0 01-3.417.592l-2.147-6.15M18 13a3 3 0 100-6M5.436 13.683A4.001 4.001 0 017 6h1.832c4.1 0 7.625-1.234 9.168-3v14c-1.543-1.766-5.067-3-9.168-3H7a3.988 3.988 0 01-1.564-.317z' },
  { name: 'Reviews', icon: 'M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z' },
  { name: 'Learning tools', icon: 'M4 5a1 1 0 011-1h4a1 1 0 011 1v4a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM14 5a1 1 0 011-1h4a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 01-1-1V5zM4 15a1 1 0 011-1h4a1 1 0 011 1v4a1 1 0 01-1 1H5a1 1 0 01-1-1v-4zM14 15a1 1 0 011-1h4a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 01-1-1v-4z' }
]

const lessonsList = ref([
  { id: '01', title: 'Advanced Machine Learning Architecture & Pipelines', duration: '15:30 mins', isHot: true },
  { id: '02', title: 'The booklets main theme remains the same...', duration: '15:30 mins', isHot: false }
])

const instructorProfile = {
  avatar: 'https://i.pravatar.cc/150?u=davidmiller',
  name: 'David Miller',
  title: 'Senior AI & Machine Learning Expert',
  bio: 'David Miller is an accomplished AI researcher...',
  rating: 4.95,
  students: 28543,
  courses: 9
}

const aiQuery = ref('')
const aiMessages = ref([
  { sender: 'ai', text: 'Hello! I am your AI Assistant. Do you have any questions related to this lesson?' }
])

const sendAiMessage = () => {
  if (!aiQuery.value.trim()) return
  aiMessages.value.push({ sender: 'user', text: aiQuery.value })
  const question = aiQuery.value
  aiQuery.value = ''
  
  setTimeout(() => {
    aiMessages.value.push({ sender: 'ai', text: `Regarding your question "${question}", it relates to the system architecture...` })
  }, 1000)
}
</script>

<template>
  <div class="watch-page">
    <div class="watch-layout">
      
      <!-- ផ្នែកខាងឆ្វេង៖ វីដេអូ និងមាតិកា -->
      <div class="main-content-col">

        <!-- Video Player Box -->
        <div class="video-container">
          <div class="video-screen" :style="{ backgroundImage: `url(${course.mediaUrl})`, backgroundSize: 'cover', backgroundPosition: 'center' }">
            <div class="play-btn">
              <svg viewBox="0 0 24 24" width="40" height="40" fill="currentColor"><path d="M8 5v14l11-7z"/></svg>
            </div>
          </div>
        </div>

        <!-- Course Title & Info -->
        <div class="course-info-box">
          <h1 class="course-main-title">{{ course.title }}</h1>
          <p class="course-main-desc">{{ course.description }}</p>
        </div>

        <!-- Instructor Profile Section -->
        <div class="instructor-profile-box">
          <div class="prof-header">
            <img :src="instructorProfile.avatar" :alt="instructorProfile.name" class="prof-avatar">
            <div class="prof-info-top">
              <h2 class="prof-name">{{ instructorProfile.name }}</h2>
              <p class="prof-title">{{ instructorProfile.title }}</p>
            </div>
          </div>
        </div>

        <!-- Switcher Tabs -->
        <div class="switcher-tabs">
          <button 
            class="switch-btn" 
            :class="{ active: activeContentTab === 'Content Course' }"
            @click="handleSwitcherClick('Content Course')"
          >
            Content Course
          </button>
          <button 
            class="switch-btn" 
            :class="{ active: activeContentTab === 'Reviews' }"
            @click="handleSwitcherClick('Reviews')"
          >
            Reviews
          </button>
          <button 
            class="switch-btn" 
            :class="{ active: activeContentTab === 'AI Assistance' }"
            @click="handleSwitcherClick('AI Assistance')"
          >
            AI Assistance
          </button>
        </div>

        <!-- Dynamic Conditional Rendering for Sidebar Tabs -->
        <div class="sections-list" v-if="activeSidebarTab === 'Lesson'">
          <LessonSection 
            v-for="lesson in lessonsList" 
            :key="lesson.id" 
            :lesson="lesson" 
            :mediaUrl="course.mediaUrl" 
          />
        </div>

        <AIAssistant 
          v-else-if="activeSidebarTab === 'AI Assistant'"
          v-model:aiQuery="aiQuery"
          :aiMessages="aiMessages"
          @send="sendAiMessage"
        />

        <Overview v-else-if="activeSidebarTab === 'Overview'" />

        <QnAManagement v-else-if="activeSidebarTab === 'Q&A'" />

        <Note v-else-if="activeSidebarTab === 'Notes'" />

        <Annount v-else-if="activeSidebarTab === 'Announcement'" />

        <Review v-else-if="activeSidebarTab === 'Reviews'" @back="activeSidebarTab = 'Lesson'" />

        <!-- 2. បន្ថែម LearningTool Component នៅទីនេះ -->
        <LearningTool v-else-if="activeSidebarTab === 'Learning tools'" @back="activeSidebarTab = 'Lesson'" />

        <!-- Other Tab Content -->
        <div class="other-tab-content" v-else>
          <p>Content for Tab: <strong>{{ activeSidebarTab }}</strong></p>
        </div>

      </div>

      <!-- Right: Sidebar Menu -->
      <div class="sidebar-menu-col">
        <div class="sidebar-menu-container">
          <button 
            v-for="tab in sidebarTabs" 
            :key="tab.name"
            class="menu-item"
            :class="{ active: activeSidebarTab === tab.name }"
            @click="activeSidebarTab = tab.name"
          >
            <svg class="menu-icon" viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path :d="tab.icon"></path>
            </svg>
            <span class="menu-text">{{ tab.name }}</span>
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>
.watch-page {
  background-color: #f8fafc;
  min-height: 100vh;
  color: #0f172a;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
}

.watch-layout {
  max-width: 1280px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 280px;
  gap: 12px;
  align-items: start;
}

.main-content-col {
  display: flex;
  flex-direction: column;
  gap: 16px;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  padding: 0;
  overflow: hidden;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.02), 0 2px 4px -1px rgba(0, 0, 0, 0.02);
}

.video-container {
  width: 100%;
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
}

.video-screen {
  width: 100%;
  height: 380px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
}

.video-screen::before {
  content: '';
  position: absolute;
  inset: 0;
  background: rgba(15, 23, 42, 0.25);
}

.play-btn {
  width: 64px;
  height: 64px;
  background: rgba(25, 118, 210, 0.9);
  border: 2px solid rgba(255, 255, 255, 0.8);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #fff;
  z-index: 1;
  box-shadow: 0 10px 20px rgba(25, 118, 210, 0.3);
  transition: transform 0.2s ease;
}

.play-btn:hover {
  transform: scale(1.05);
}

.video-controls {
  background-color: #ffffff;
  border-top: 1px solid #e2e8f0;
  padding: 12px 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.ctrl-group {
  display: flex;
  align-items: center;
  gap: 12px;
}

.right-group {
  gap: 8px;
}

.icon-btn {
  background: transparent;
  border: none;
  color: #475569;
  cursor: pointer;
  display: flex;
  padding: 4px;
  border-radius: 6px;
  transition: background 0.2s, color 0.2s;
}

.icon-btn:hover {
  color: #1976D2;
  background: #f1f5f9;
}

.progress-bar-wrapper {
  width: 140px;
}

.progress-line {
  height: 4px;
  background-color: #e2e8f0;
  border-radius: 2px;
  position: relative;
  overflow: hidden;
}

.progress-filled {
  width: 40%;
  height: 100%;
  background-color: #1976D2;
  border-radius: 2px;
  position: relative;
}

.time-badge-with-clock {
  display: flex;
  align-items: center;
  gap: 6px;
  background-color: #eef2f6;
  padding: 4px 12px 4px 10px;
  border-radius: 50px;
  color: #0f172a;
}

.badge-clock-icon {
  color: #475569;
  flex-shrink: 0;
}

.time-badge {
  font-size: 13px;
  font-family: 'Inter', system-ui, sans-serif;
  color: #0f172a;
  font-weight: 600;
  letter-spacing: 0.3px;
}

.course-info-box,
.instructor-profile-box,
.switcher-tabs,
.sections-list,
.other-tab-content {
  padding-left: 20px;
  padding-right: 20px;
}

.course-info-box {
  padding-top: 12px;
  padding-bottom: 6px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.course-main-title {
  font-size: 20px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}

.course-main-desc {
  font-size: 13.5px;
  color: #64748b;
  margin: 0;
  line-height: 1.5;
}

.instructor-profile-box {
  background-color: #f8fafc;
  padding: 16px 20px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 10px;
}

.prof-header {
  display: flex;
  align-items: center;
  gap: 16px;
}

.prof-avatar {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #ffffff;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
  flex-shrink: 0;
}

.prof-info-top {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.prof-name {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

.prof-title {
  font-size: 13px;
  color: #1976D2;
  font-weight: 600;
  margin: 0;
}

.prof-stats {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 4px;
  flex-wrap: wrap;
}

.stat-dot {
  color: #94a3b8;
  font-size: 12px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 5px;
  font-size: 12px;
  color: #475569;
  font-weight: 500;
}

.star-stat {
  color: #b45309;
}

.star-stat svg {
  color: #f59e0b;
}

.prof-bio {
  font-size: 13px;
  color: #475569;
  line-height: 1.5;
  margin: 0;
}

.switcher-tabs {
  display: flex;
  gap: 16px;
  width: 100%;
  border-bottom: 2px solid #e2e8f0;
  margin-top: 8px;
}

.switch-btn {
  background: transparent;
  border: none;
  color: #64748b;
  font-weight: 600;
  font-size: 14px;
  padding: 10px 4px;
  cursor: pointer;
  position: relative;
  transition: all 0.2s;
}

.switch-btn:hover,
.switch-btn.active {
  color: #1976D2;
}

.switch-btn.active::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #1976D2;
}

.sections-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding-bottom: 16px;
  padding-top: 8px;
}

.other-tab-content {
  padding-top: 16px;
  padding-bottom: 16px;
  color: #64748b;
  font-size: 14px;
}

.sidebar-menu-container {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 6px;
  display: flex;
  flex-direction: column;
  gap: 2px;
  position: sticky;
  top: 24px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.02);
  margin-top: 12px;
}

.menu-item {
  background-color: transparent;
  border: none;
  border-radius: 10px;
  padding: 8px 12px;
  display: flex;
  align-items: center;
  gap: 10px;
  color: #475569;
  cursor: pointer;
  width: 100%;
  text-align: left;
  transition: all 0.2s ease;
}

.menu-icon {
  width: 18px;
  height: 18px;
  color: #64748b;
  flex-shrink: 0;
}

.menu-text {
  font-size: 14px;
  font-weight: 500;
}

.menu-item:hover,
.menu-item.active {
  background-color: #f8fafc;
  color: #1976D2;
}

.menu-item:hover .menu-icon,
.menu-item.active .menu-icon {
  color: #1976D2;
}
</style>