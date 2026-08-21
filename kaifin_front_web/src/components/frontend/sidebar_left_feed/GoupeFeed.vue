<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import GroupCreateForm from './GroupCreateForm.vue'
const BASE_URL = 'http://localhost:7070'

function resolveImageUrl(path) {
  if (!path) return ''
  if (path.startsWith('http://') || path.startsWith('https://')) return path
  return `${BASE_URL}/uploads/${path}`
}

function formatMemberCount(count) {
  if (count == null) return '0'
  if (count >= 1_000_000) return (count / 1_000_000).toFixed(1) + 'M'
  if (count >= 1_000) return (count / 1_000).toFixed(1) + 'K'
  return String(count)
}

function authHeaders() {
  const token = localStorage.getItem('token')
  return token ? { Authorization: `Bearer ${token}` } : {}
}

const router = useRouter()
const emit = defineEmits(['create-group', 'select-group'])

const activeTab = ref('Trending')
const groups = ref([])
const loading = ref(false)
const errorMsg = ref('')
const showCreateGroupModal = ref(false)
const creatingGroup = ref(false)
const createErrorMsg = ref('')

const page = ref(1)
const perpage = ref(20)

const selectTab = (tabName) => {
  activeTab.value = tabName
  fetchGroups()
}

const fetchGroups = async () => {
  loading.value = true
  errorMsg.value = ''
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/communities/show`, {
      params: { page: page.value, perpage: perpage.value },
      headers: authHeaders()
    })

    const list = res.data?.data?.communities ?? []

    groups.value = list.map(c => ({
      id: c.id,
      name: c.name,
      members: formatMemberCount(c.member_count),
      avatar: resolveImageUrl(c.avatar_url) || 'https://via.placeholder.com/100',
      isJoined: c.is_joined ?? false,
      verified: c.is_verified ?? false,
      memberAvatars: []
    }))
  } catch (err) {
    console.error('Failed to fetch communities:', err)
    if (err.response?.status === 401) {
      errorMsg.value = 'Session expired. Please log in again.'
    } else {
      errorMsg.value = 'Failed to load groups. Please try again.'
    }
  } finally {
    loading.value = false
  }
}

const goToGroupDetail = (groupId) => {
  router.push({ name: 'GroupDetail', params: { id: groupId } })
}

const handleJoin = async (group) => {
  const prevState = group.isJoined
  group.isJoined = !group.isJoined

  try {
    const res = await axios.post(
      `${BASE_URL}/api/v1/front/communities/${group.id}/join`,
      {},
      { headers: authHeaders() }
    )
    const status = res.data?.data?.status
    if (status === 'pending' || status === 'left') {
      group.isJoined = false
    }
  } catch (err) {
    console.error('Failed to toggle join:', err)
    group.isJoined = prevState
  }
}

const handleCreateGroup = () => {
  createErrorMsg.value = ''
  showCreateGroupModal.value = true
  emit('create-group')
}

const closeCreateGroupModal = () => {
  showCreateGroupModal.value = false
}

function mapPrivacy(access) {
  return access === 'public' ? 'public' : 'private'
}

const submitCreateGroup = async (payload) => {

  creatingGroup.value = true
  createErrorMsg.value = ''

  try {
    const formData = new FormData()
    formData.append('name', payload.name)
    formData.append('description', payload.description || '')
    if (payload.category) {
      formData.append('category_id', payload.category)
    }
    formData.append('privacy', mapPrivacy(payload.access))
    if (payload.avatarFile) {
      formData.append('avatar', payload.avatarFile)
    }
    if (payload.coverFile) {
      formData.append('cover', payload.coverFile)
    }

    const res = await axios.post(
      `${BASE_URL}/api/v1/front/communities/create`,
      formData,
      {
        headers: {
          ...authHeaders(),
          'Content-Type': 'multipart/form-data',
        },
      }
    )

    const created = res.data?.data
    await fetchGroups()

    closeCreateGroupModal()
  } catch (err) {
    console.error('Failed to create community:', err)
    if (err.response?.status === 401) {
      createErrorMsg.value = 'Session expired. Please log in again.'
    } else if (err.response?.data?.message) {
      createErrorMsg.value = err.response.data.message
    } else {
      createErrorMsg.value = 'Failed to create group. Please try again.'
    }
  } finally {
    creatingGroup.value = false
  }
}

onMounted(fetchGroups)
</script>

<template>
  <div class="group-feed-container">
    <!-- Top Navigation Bar -->
    <header class="feed-header-wrapper">
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

      <div class="header-card actions-card bdr-search">
        <div class="search-box">
          <svg class="search-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
          <input type="text" placeholder="Search groups..." class="search-input" />
        </div>
        <button class="btn-create-group" @click="handleCreateGroup">
          + Create Group
        </button>
      <GroupCreateForm v-if="showCreateGroupModal" @close="closeCreateGroupModal" @create="submitCreateGroup"/>
      </div>
    </header>

    <div v-if="loading" class="loading-state">Loading groups...</div>
    <div v-else-if="errorMsg" class="error-state">{{ errorMsg }}</div>

    <section v-else class="suggestions-section">
      <h2 class="section-title">Recommended For You</h2>

      <div class="groups-grid">
        <div 
          v-for="group in groups" 
          :key="group.id" 
          class="group-card"
          @click="goToGroupDetail(group.id)"
        >
          <div class="group-avatar-wrapper" @click.stop="goToGroupDetail(group.id)">
            <img :src="group.avatar" :alt="group.name" class="group-avatar" />

            <div class="hover-preview" @click.stop>
              <div class="preview-header" :style="{ backgroundImage: `linear-gradient(to top, rgba(255, 255, 255, 0.85) 0%, rgba(255, 255, 255, 0.3) 50%, transparent 100%), url(${group.avatar})` }">
                <div class="preview-cover-actions">
                  <button class="btn-sub" @click.stop="handleJoin(group)">{{ group.isJoined ? 'Joined' : 'Join' }}</button>
                  <button class="btn-save">Bookmark</button>
                </div>
              </div>

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
              <svg v-if="group.verified" class="verified-icon" width="16" height="16" viewBox="0 0 24 24" fill="#1B75D2"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z"/></svg>
            </div>

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
  cursor: pointer; 
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


.btn-join {
  background-color: #F4F1EC;
  color: #000;
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
  background-color: #e8e6e3;
  color: #000;
}

.btn-join.joined {
  background-color: #1B75D2;
  color: #ffffff;
}


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

.preview-cover-actions {
  position: absolute;
  right: 12px;
  bottom: 10px;
  display: flex;
  gap: 8px;
}

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