<!-- File: LearningTool.vue -->
<script setup>
import { ref } from 'vue'

defineEmits(['back'])

const tools = ref([
  {
    id: 1,
    name: 'Visual Studio Code',
    description: 'Lightweight code editor for web development.',
    icon: 'https://upload.wikimedia.org/wikipedia/commons/9/9a/Visual_Studio_Code_1.35_icon.svg',
    downloadUrl: '#',
    macUrl: '#',
    windowUrl: '#',
    linuxUrl: '#'
  },
  {
    id: 2,
    name: 'Figma',
    description: 'Collaborative interface design and prototyping tool.',
    icon: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSebpk8AZpGklpMqMbMCXQpWGixkL-liIf5zJyvzhaYsg&s=10',
    downloadUrl: '#',
    macUrl: '#',
    windowUrl: '#'
  },
  {
    id: 2,
    name: 'Figma',
    description: 'VueJS is an open source progressive JavaScript framework used to develop interactive web interfaces.',
    icon: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRfW3ry8iphnle-VibRvXNOCEaj3jKtL9asZ1KIiU5U3A&s=10',
    downloadUrl: '#',
    macUrl: '#',
    windowUrl: '#'
  },
  {
    id: 2,
    name: 'Figma',
    description: 'pgAdmin is a web-based database management tool for administering your PostgreSQL databases through a user-friendly interface.t provides a comprehensive set of features for database administration,',
    icon: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQRRuGR17IPqOsahi6cafgcCBFFkko199a3beWpnr1vLA&s=10',
    downloadUrl: '#',
    macUrl: '#',
    windowUrl: '#'
  }
])
</script>

<template>
  <div class="learning-tool-container">
    
    <!-- Header with Back Button & Centered Title -->
    <div class="tool-header">
      <button class="back-btn" @click="$emit('back')" title="Go back">
        <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="15 18 9 12 15 6"></polyline>
        </svg>
      </button>
      <h2>All Tools use for course</h2>
      <div class="header-spacer"></div>
    </div>

    <!-- Tools List -->
    <div class="tools-list">
      <div v-for="tool in tools" :key="tool.id" class="tool-card">
        
        <!-- Tool Icon Box -->
        <div class="tool-icon-box">
          <img :src="tool.icon" :alt="tool.name" class="tool-img" />
        </div>

        <!-- Tool Info & Action Buttons -->
        <div class="tool-content">
          <div class="tool-name">{{ tool.name }}</div>
          <p v-if="tool.description" class="tool-desc">{{ tool.description }}</p>
          
          <div class="tool-actions">
            <a :href="tool.downloadUrl" class="tool-btn primary">Download</a>
            <a :href="tool.windowUrl" class="tool-btn">🖥️ Windows</a>
            <a :href="tool.macUrl" class="tool-btn">🍎 macOS</a>
            <a v-if="tool.linuxUrl" :href="tool.linuxUrl" class="tool-btn">🐧 Linux</a>
          </div>
        </div>

      </div>
    </div>

  </div>
</template>

<style scoped>
.learning-tool-container {
  background-color: #ffffff;
  padding: 20px;
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  min-height: 100vh;
  box-sizing: border-box;
}

.tool-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  border-radius: 12px;
}

.tool-header h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
  text-align: center;
  flex: 1;
}

.header-spacer {
  width: 40px;
}

.back-btn {
  background-color: transparent;
  color: #0f172a;
  border: none;
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.back-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.tools-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.tool-card {
  border-radius: 16px;
  padding: 18px;
  display: flex;
  align-items: center;
  gap: 18px;
  background: #ffffff;
  border: 1px solid #f1f5f9;
}

.tool-icon-box {
  width: 64px;
  height: 64px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f8fafc;
  overflow: hidden;
  flex-shrink: 0;
  padding: 12px;
  border: 1px solid #f1f5f9;
}

.tool-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.tool-content {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
}

.tool-name {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}

.tool-desc {
  margin: 0 0 8px 0;
  font-size: 13px;
  color: #64748b;
  line-height: 1.4;
}

.tool-desc{
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
}

.tool-actions {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.tool-btn {
  border: none;
  background-color: #f1f5f9;
  color: #334155;
  padding: 6px 14px;
  border-radius: 32px;
  font-size: 13px;
  font-weight: 600;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.tool-btn.primary {
  background-color: #e0f2fe;
  color: #0369a1;
}

.tool-btn:hover {
  background-color: #e2e8f0;
  color: #0f172a;
}

.tool-btn.primary:hover {
  background-color: #bae6fd;
  color: #0369a1;
}
</style>