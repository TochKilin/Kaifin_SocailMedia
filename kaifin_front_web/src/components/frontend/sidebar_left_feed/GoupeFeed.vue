<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router' // 1. Import useRouter

const router = useRouter() // 2. បង្កើត instance របស់ router

const props = defineProps({
  // You can add custom props here if needed
})

const emit = defineEmits(['create-group', 'select-group'])

// State សម្រាប់គ្រប់គ្រង Tab ដែលបានជ្រើសរើស
const activeTab = ref('Trending')

const selectTab = (tabName) => {
  activeTab.value = tabName
}

// Sample data for groups
const groups = ref([
  {
    id: 1,
    name: 'Fixies Group',
    members: '399K',
    avatar: 'https://mediaproxy.tvtropes.org/width/1200/https://static.tvtropes.org/pmwiki/pub/images/2025_01_29_8.png',
    isJoined: false,
    verified: true,
    memberAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQIJr17LxjxPnvB6BNdzWsjjDY-l271LsmpsnXqscbc7A&s=10',
      'https://randomuser.me/api/portraits/women/44.jpg',
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQc8IReG-cUvPgentGMhx_TVDmsXgjJRnJF-qqln3zZ4A&s=10'
    ]
  },
  {
    id: 2,
    name: 'ក្រុម ឡាយលក់ឡេ',
    members: '8.6K',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSE1Q0Rh0oTqNk5wbto0QU5qgcZExtV9bGtBIB5DgNmKg&s',
    isJoined: false,
    verified: false,
     memberAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ12FGxm5NKI8MeXEBUeP-sCv8Az1nTcY5jv6MvvK_1Og&s=10',
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR7Ldyjm0BBxmpQ7J1P4h9Nevn84VAKgVniUBPxBL1dqw&s=10'
    ]
  },
  {
    id: 3,
    name: 'IT Sol',
    members: '6K',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ6v1EH9rHqZmLP2eHEkA5HyongBeWqk4N8UYhpElZJNg&s=10',
    isJoined: false,
    verified: false,
    memberAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTw2Bv9lSLn_RcruvM8tvUAwaRleoqvGRYz1pGnr9bT2A&s=10',
      'https://randomuser.me/api/portraits/men/21.jpg'
    ]
  },
  {
    id: 4,
    name: 'Comedy & Laughs',
    members: '1.1M',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS8wm43Xg52qiq1kmBRDXavAVl1vvucmG9-Nre6hbBobg&s=10',
    isJoined: false,
    verified: false,
    memberAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSlvYTE0TSdoApMO06QKeJkC31rFkjiVCVBi_FDOnNVvg&s=10',
      'https://randomuser.me/api/portraits/women/81.jpg',
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10'
    ]
  },
  {
    id: 5,
    name: 'កសិកម្ម',
    members: '8K',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTb72ejD-FUB12t-ZgnZAemOE13v2YcqBJy_eb3XDKgfA&s=10',
    isJoined: false,
    verified: true,
     memberAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSW3icLiUy4-UOpZ9cpDYSfVj6730wH_t4ku589N_Qrcw&s=10',
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQFQEuSgk_402rlNg2nHbIOGap1MtBGsUhtIc18tKR5IQ&s=10',
      'https://randomuser.me/api/portraits/women/78.jpg'
    ]
  },
  {
    id: 6,
    name: '🌷🍹 Café & Chill 🍧🎀',
    members: '32K',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRx4KUKbjL3RIhaq1FRkI87nqmbkh1k1Ycxo3Z-BdcYYw&s=10',
    isJoined: false,
    verified: false,
    memberAvatars: [
      'https://randomuser.me/api/portraits/women/33.jpg',
      'https://randomuser.me/api/portraits/women/11.jpg',
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa28iNNqLJdEZ7aXrMJJUkFkz1CSFxXBVyqKYyMDf7MA&s=10'
    ]
  },
  {
    id: 7,
    name: 'PUMA BRANDS',
    members: '1.9K',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTxep2RSyMws-8K0uWZqxb2PcC_Tt4JUuomrJLuJHxCmA&s=10',
    isJoined: false,
    verified: false,
        memberAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSW3icLiUy4-UOpZ9cpDYSfVj6730wH_t4ku589N_Qrcw&s=10',
      'https://randomuser.me/api/portraits/men/21.jpg'
    ]
  },
  {
    id: 8,
    name: 'Stalker Online Official',
    members: '3.3K',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQqG8y5wm8jODoWSPrAgVThSuZ9fBVmirAN9rx_cm2Aaw&s',
    isJoined: false,
    verified: false,
     memberAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSW3icLiUy4-UOpZ9cpDYSfVj6730wH_t4ku589N_Qrcw&s=10',
      'https://randomuser.me/api/portraits/men/21.jpg'
    ]
  },
  {
    id: 9,
    name: 'Tech Insider KH',
    members: '45K',
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ9rQQcbY9WMY4rCsOeDNrHX9fbIDyfg1hv2VLxsvy9OQ&s=10',
    isJoined: false,
    verified: true,
        memberAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQqG8y5wm8jODoWSPrAgVThSuZ9fBVmirAN9rx_cm2Aaw&s',
      'https://randomuser.me/api/portraits/men/21.jpg'
    ]
  }
])

// 3. Function សម្រាប់បើកទៅកាន់ GroupDetail
const goToGroupDetail = (groupId) => {
  // ប្រសិនបើប្រើ Vue Router (បញ្ជូនទៅ Route /group/:id ឬ route name)
  router.push({ name: 'GroupDetail', params: { id: groupId } })
  
  // ឬប្រសិនបើមិនប្រើ Router ទេ អ្នកអាច Emit ទៅ Parent Component បាន៖
  // emit('select-group', groupId)
}

const handleJoin = (group) => {
  group.isJoined = !group.isJoined
}

const handleCreateGroup = () => {
  emit('create-group')
}
</script>

<template>
  <div class="group-feed-container">
    <!-- Top Navigation Bar -->
    <header class="feed-header-wrapper">
      <!-- ផ្នែកទី ១៖ Tabs -->
      <div class="header-card tabs-card bdr-category ">
        <div class="nav-links">
          <button 
            class="nav-item" 
            :class="{ active: activeTab === 'Trending' }"
            @click="selectTab('Trending')"
          >
            Trending
          </button>
          <button 
            class="nav-item" 
            :class="{ active: activeTab === 'Official' }"
            @click="selectTab('Official')"
          >
            Official
          </button>
          <button 
            class="nav-item" 
            :class="{ active: activeTab === 'New' }"
            @click="selectTab('New')"
          >
            New
          </button>
          <button 
            class="nav-item dropdown" 
            :class="{ active: activeTab === 'Category' }"
            @click="selectTab('Category')"
          >
            Category <span class="arrow">▼</span>
          </button>
        </div>
      </div>

      <!-- ផ្នែកទី ២៖ Search & Create Group -->
      <div class="header-card actions-card bdr-search">
        <div class="search-box">
          <svg class="search-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
          <input type="text" placeholder="Search groups..." class="search-input" />
        </div>
        <button class="btn-create-group" @click="handleCreateGroup">
          + Create Group
        </button>
      </div>
    </header>

    <!-- Main Content Section -->
    <section class="suggestions-section">
      <h2 class="section-title">Recommended For You</h2>

      <div class="groups-grid">
        <!-- បន្ថែម @click="goToGroupDetail(group.id)" លើ group-card -->
        <div 
          v-for="group in groups" 
          :key="group.id" 
          class="group-card"
          @click="goToGroupDetail(group.id)"
        >
          <div class="group-avatar-wrapper" @click.stop="goToGroupDetail(group.id)">
            <!-- Main Avatar កោង 12px -->
            <img :src="group.avatar" :alt="group.name" class="group-avatar" />
            
            <!-- Hover Preview Box -->
            <div class="hover-preview" @click.stop>
              <!-- បន្ថែម Gradient រាងសខ្ចីៗ ពីក្រោមមកលើ -->
              <div class="preview-header" :style="{ backgroundImage: `linear-gradient(to top, rgba(255, 255, 255, 0.85) 0%, rgba(255, 255, 255, 0.3) 50%, transparent 100%), url(${group.avatar})` }">
                <div class="preview-cover-actions">
                  <button class="btn-sub" @click.stop="handleJoin(group)">{{ group.isJoined ? 'Joined' : 'Join' }}</button>
                  <button class="btn-save">Bookmark</button>
                </div>
              </div>
              
              <!-- Profile Thumbnail overlapping header -->
              <div class="preview-thumbnail-container">
                <img :src="group.avatar" :alt="group.name" class="preview-thumbnail" />
              </div>

              <div class="preview-body">
                <div class="preview-title-container">
                  <h3 class="preview-title-outside">{{ group.name }}</h3>
                </div>
                <div class="preview-info">
                  <span class="preview-category">Category: Social</span>
                  <div class="preview-members-wrapper">
                    <!-- Member Avatars ក្នុង Hover Preview -->
                    <div class="stacked-avatars" v-if="group.memberAvatars.length > 0">
                      <img 
                        v-for="(avatar, index) in group.memberAvatars.slice(0, 3)" 
                        :key="index" 
                        :src="avatar" 
                        class="stacked-avatar" 
                        :style="{ zIndex: 3 - index }"
                      />
                    </div>
                    <svg class="group-icon" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" v-else>
                      <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                      <circle cx="9" cy="7" r="4"></circle>
                      <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                      <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                    </svg>
                    <span class="preview-members">{{ group.members }} members</span>
                  </div>
                </div>
              </div>
            </div>

          </div>
          
          <div class="group-info">
            <div class="group-name-row">
              <h3 class="group-name">{{ group.name }}</h3>
              <!-- Verified Badge -->
              <svg v-if="group.verified" class="verified-icon" width="16" height="16" viewBox="0 0 24 24" fill="#1B75D2"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z"/></svg>
            </div>
            
            <!-- Group Members with Stacked Avatars -->
            <div class="group-members-row">
              <div class="stacked-avatars" v-if="group.memberAvatars.length > 0">
                <img 
                  v-for="(avatar, index) in group.memberAvatars.slice(0, 3)" 
                  :key="index" 
                  :src="avatar" 
                  class="stacked-avatar" 
                  :style="{ zIndex: 3 - index }"
                />
              </div>
              <svg class="group-icon" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" v-else>
                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                <circle cx="9" cy="7" r="4"></circle>
                <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
              </svg>
              <span class="group-members">{{ group.members }} members</span>
            </div>

          </div>

          <!-- Join Button ខាងក្រៅ បន្ថែម @click.stop ដើម្បីកុំឱ្យវា Trigger ទៅកាន់ function បើក Detail ពេលចុច Join -->
          <button 
            class="btn-join" 
            :class="{ joined: group.isJoined }"
            @click.stop="handleJoin(group)"
          >
            {{ group.isJoined ? 'Joined' : 'Join' }}
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.group-feed-container {
  max-width: 100%;
  margin: 0 auto;
  font-family: 'Inter', system-ui, sans-serif;
  color: #1e293b;
  margin-top: 12px;
}

/* Header Wrapper សម្រាប់ផ្ទុកទាំង ២ ផ្នែក */
.feed-header-wrapper {
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin-bottom: 16px;
}

.bdr-search {
  border-top: 1px solid #dddddd47;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.bdr-category {
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
  border-bottom: 1px solid #dddddd47;
}

/* ការឌីហ្សាញ Background ដាច់ដោយឡែកពីគ្នាជា Card */
.header-card {
  background-color: #ffffff;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.04);
  padding: 0 20px;
}

.tabs-card {
  padding-top: 4px;
}

.actions-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 20px;
}

.nav-links {
  display: flex;
  gap: 24px;
  align-items: center;
  flex-wrap: wrap;
}

.nav-item {
  background: transparent;
  border: none;
  font-size: 15px;
  font-weight: 500;
  color: #64748b;
  padding: 12px 4px;
  cursor: pointer;
  position: relative;
  transition: color 0.2s ease;
}

.nav-item:hover {
  color: #0f172a;
}

.nav-item.active {
  color: #1B75D2;
  font-weight: 600;
}

.nav-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background-color: #1B75D2;
  border-radius: 3px 3px 0 0;
}

.arrow {
  font-size: 10px;
  margin-left: 4px;
}

.search-box {
  position: relative;
  display: flex;
  align-items: center;
  flex-grow: 1;
}

.search-icon {
  position: absolute;
  left: 14px;
  color: #94a3b8;
  z-index: 2;
}

.search-input {
  border: 1px solid #e2e8f0;
  border-radius: 24px;
  padding: 8px 16px 8px 38px;
  font-size: 14px;
  outline: none;
  width: 100%;
  transition: all 0.2s ease;
}

.search-input:focus {
  background-color: #ffffff;
  border-color: #1B75D2;
  box-shadow: 0 0 0 3px rgba(27, 117, 210, 0.1);
}

.btn-create-group {
  background-color: #1B75D2;
  color: white;
  border: none;
  border-radius: 24px;
  padding: 8px 18px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s ease;
  white-space: nowrap;
  flex-shrink: 0;
}

.btn-create-group:hover {
  background-color: #155bb5;
  transform: translateY(-1px);
}

/* Suggestions Section */
.suggestions-section {
  background: #ffffff;
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.04);
}

.section-title {
  font-size: 18px;
  font-weight: 700;
  margin-bottom: 18px;
  color: #0f172a;
}

/* Groups Grid Layout */
.groups-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  gap: 16px;
}

.group-card {
  display: flex;
  align-items: center;
  padding: 16px;
  background-color: #ffffff;
  border: 1px solid #f1f5f9;
  border-radius: 16px;
  transition: all 0.2s ease;
  gap: 16px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.02);
  cursor: pointer; /* បន្ថែម Mouse Pointer ឱ្យដឹងថាមើលទៅអាចចុចបាន */
}

.group-card:hover {
  background-color: #fafafa;
  border-color: #cbd5e1;
}

.group-avatar-wrapper {
  position: relative;
  width: 130px;
  height: 130px;
  border-radius: 12px;
  overflow: visible;
  flex-shrink: 0;
  border: 1px solid #e2e8f0;
  cursor: pointer;
}

/* Profile Avatar កោង 12px */
.group-avatar {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 12px;
}

.group-info {
  flex-grow: 1;
  min-width: 0;
}

.group-name-row {
  display: flex;
  align-items: center;
  gap: 6px;
}

.group-name {
  font-size: 15px;
  font-weight: 600;
  color: #0f172a;
  margin: 0 0 6px 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.verified-icon {
  flex-shrink: 0;
}

/* Group Members Row & Stacked Avatars */
.group-members-row {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #64748b;
}

.stacked-avatars {
  display: flex;
  align-items: center;
}

.stacked-avatar {
  width: 24px;
  height: 24px;
  border-radius: 4px;
  object-fit: cover;
  border: 2px solid #ffffff;
  margin-left: -8px;
}

.stacked-avatar:first-child {
  margin-left: 0;
}

.group-icon {
  color: #64748b;
  flex-shrink: 0;
}

.preview-members, .group-members {
  font-size: 13px;
  color: #64748b;
  margin: 0;
}

/* Join Button ខាងក្រៅ */
.btn-join {
  background-color: #1B75D2;
  color: #ffffff;
  border: none;
  border-radius: 20px;
  padding: 8px 20px;
  font-weight: 600;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.2s ease;
  flex-shrink: 0;
}

.btn-join:hover {
  background-color: #155bb5;
}

.btn-join.joined {
  background-color: #1B75D2;
  color: #ffffff;
}

/* Hover Preview Box Styles */
.hover-preview {
  position: absolute;
  top: 95%; 
  left: 0; 
  width: 440px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 6px 20px rgba(0,0,0,0.12);
  z-index: 10;
  opacity: 0;
  visibility: hidden;
  transition: all 0.25s ease;
  overflow: visible;
  border: 1px solid #e2e8f0;
}

.group-avatar-wrapper:hover .hover-preview {
  opacity: 1;
  visibility: visible;
  top: 102%;
}

.preview-header {
  height: 110px;
  background-size: cover;
  background-position: center;
  position: relative;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
}

/* Action Buttons */
.preview-cover-actions {
  position: absolute;
  right: 12px;
  bottom: 10px;
  display: flex;
  gap: 8px;
}

/* Thumbnail Container overlapping header */
.preview-thumbnail-container {
  position: absolute;
  top: 75px; 
  left: 16px;
  width: 65px;
  height: 65px;
  border-radius: 12px;
  border: 3px solid #ffffff;
  overflow: hidden;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
  background: #ffffff;
  z-index: 2;
}

.preview-thumbnail {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 9px;
}

.preview-body { 
  padding: 12px 16px 16px 16px; 
}

.preview-title-container {
  margin-left: 75px;
  margin-bottom: 8px;
}

.preview-title-outside {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.preview-info { 
  margin-bottom: 0; 
  font-size: 13px; 
  color: #475569;
  display: flex;
  align-items: center;
  gap: 12px;
}

.preview-members-wrapper {
  display: flex;
  align-items: center;
  gap: 6px;
}

.btn-sub, .btn-save {
  padding: 8px 16px;
  border: none;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  background: #1B75D2;
  color: #ffffff;
  box-shadow: none;
  transition: background 0.2s;
  white-space: nowrap;
}

.btn-sub:hover, .btn-save:hover {
  background: #155bb5;
}
</style>