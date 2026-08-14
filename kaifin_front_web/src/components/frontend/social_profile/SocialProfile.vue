<template>
 <div>
    <!-- <NavBar/> -->
     <div class="profile-page-wrapper">
    <div class="profile-container">
      
      <!-- Top Bar -->
      <div class="top-bar"></div>

      <!-- Cover Image Section -->
      <div class="cover-container">
        <img :src="profile.coverUrl || '/default-cover.jpg'" alt="Cover Image" class="cover-image" />
        
        <!-- Gradient Overlay សម្រាប់បាញ់ពណ៌ស្រអាប់ពីក្រោមឡើងលើ -->
        <div class="cover-gradient-overlay"></div>

        <!-- Group of Buttons (Minimalist Circular Style) -->
        <div class="cover-actions-group">
          <div class="action-buttons-row">
            <button class="icon-action-btn" @click="handleSetting" aria-label="Settings">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1.37 1.9l-.17.06a2 2 0 0 1-2.11-.5l-.13-.13a2 2 0 0 0-2.83 0l-.31.31a2 2 0 0 0 0 2.83l.13.13a2 2 0 0 1 .5 2.11l-.06.17a2 2 0 0 1-1.9 1.37H2a2 2 0 0 0-2 2v.44a2 2 0 0 0 2 2h.18a2 2 0 0 1 1.9 1.37l.06.17a2 2 0 0 1-.5 2.11l-.13.13a2 2 0 0 0 0 2.83l.31.31a2 2 0 0 0 2.83 0l.13-.13a2 2 0 0 1 2.11-.5l.17-.06a2 2 0 0 1 1.37 1.9V22a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1.37-1.9l.17-.06a2 2 0 0 1 2.11.5l.13.13a2 2 0 0 0 2.83 0l.31-.31a2 2 0 0 0 0-2.83l-.13-.13a2 2 0 0 1-.5-2.11l.06-.17a2 2 0 0 1 1.9-1.37H22a2 2 0 0 0 2-2v-.44a2 2 0 0 0-2-2h-.18a2 2 0 0 1-1.9-1.37l-.06-.17a2 2 0 0 1 .5-2.11l.13-.13a2 2 0 0 0 0-2.83l-.31-.31a2 2 0 0 0-2.83 0l-.13.13a2 2 0 0 1-2.11.5l-.17-.06a2 2 0 0 1-1.37-1.9V4a2 2 0 0 0-2-2z"/>
                <circle cx="12" cy="12" r="3"/>
              </svg>
            </button>
            <button class="icon-action-btn" @click="openEditModal" aria-label="More options">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="1"></circle>
                <circle cx="19" cy="12" r="1"></circle>
                <circle cx="5" cy="12" r="1"></circle>
              </svg>
            </button>
            
          </div>

          <button class="change-cover-btn" @click="handleChangeCover" title="Change cover photo" :disabled="isUploadingCover">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <line x1="12" y1="5" x2="12" y2="19"></line>
              <line x1="5" y1="12" x2="19" y2="12"></line>
            </svg>
          </button>
          <button class="follow-btn" @click="handleFollow">
              {{ isFollowing ? 'Following' : 'Follow' }}
            </button>
        </div>

        <!-- Hidden file input សម្រាប់ cover upload (មិនប៉ះពាល់ layout/UI) -->
        <input
          type="file"
          ref="coverFileInput"
          accept="image/png, image/jpeg, image/gif"
          style="display: none"
          @change="onCoverFileSelected"
        />
      </div>

      <!-- 1. Profile Main Card (សម្រាប់ Header, Name និង Bio) -->
      <div class="profile-card-section">
        <div class="profile-header-row">
          <!-- <div class="avatar-container">
            <img :src="profile.avatarUrl" :alt="profile.name" class="avatar-image" />
          </div> -->

<div class="avatar-wrapper">
  <div class="avatar-container">
    <img :src="profile.avatarUrl" :alt="profile.name" class="avatar-image" />
  </div>
  <button class="avatar-camera-btn" @click="handleChangeAvatar" :disabled="isUploadingAvatar">
    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"></path>
      <circle cx="12" cy="13" r="4"></circle>
    </svg>
  </button>
</div>

<input
  type="file"
  ref="avatarFileInput"
  accept="image/png, image/jpeg, image/gif"
  style="display: none"
  @change="onAvatarFileSelected"
/>



          <div class="profile-title-area">
            <div class="name-row">
              <h1 class="profile-name">{{ profile.name }}</h1>
            </div>
            <div class="profile-handle">{{ profile.handle }}</div>
          </div>
        </div>

        <!-- Bio -->
        <p class="profile-bio">
          {{ profile.bio }} 
          <a :href="profile.websiteUrl" target="_blank" class="bio-link">{{ profile.websiteText }}</a>
        </p>
      </div>

      <!-- 2. Profile Meta Card -->
      <div class="profile-meta-card">
        <div class="profile-meta">
          <div class="meta-item">
            <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
            </svg>
            <span>{{ profile.relationship_status }}</span>
          </div>
          <div class="meta-item">
            <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path><circle cx="12" cy="10" r="3"></circle></svg>
            <span>{{ profile.location }}</span>
          </div>
          <div class="meta-item">
            <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
            <span>{{ profile.joinedDate }}</span>
          </div>
        </div>
      </div>

      <!-- 3. Navigation Tabs Card -->
      <div class="tabs-card-section">
        <div class="tabs-container">
          <div 
            v-for="tab in tabs" 
            :key="tab.name" 
            class="tab-item"
            :class="{ active: currentTab === tab.name }"
            @click="currentTab = tab.name"
          >
            <span>{{ tab.label }}</span>
            <span v-if="tab.count !== undefined" class="tab-count">{{ tab.count }}</span>
          </div>
        </div>

        <!-- Content Body -->
        <div class="content-body">
<!-- ត្រូវប្រាកដថាដូចនេះ -->
<PostsCard v-if="currentTab === 'Done' && profile.userId" :userId="profile.userId" />
        </div>
      </div>

    </div>
  </div>



  <!-- Edit Profile Modal -->
<div v-if="showEditModal" class="modal-overlay" @click.self="closeEditModal">
  <div class="modal-box">
    <div class="modal-header">
      <h2 class="modal-title">Edit your personal information</h2>
      <button class="modal-close-btn" @click="closeEditModal">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <line x1="18" y1="6" x2="6" y2="18"></line>
          <line x1="6" y1="6" x2="18" y2="18"></line>
        </svg>
      </button>
    </div>

    <div class="modal-form-grid">
      <div class="form-field">
        <label>First name</label>
        <input v-model="editForm.firstName" type="text" placeholder="First name" />
      </div>
      <div class="form-field">
        <label>Last name</label>
        <input v-model="editForm.lastName" type="text" placeholder="Last name" />
      </div>
      <div class="form-field">
        <label>Username</label>
        <input v-model="editForm.username" type="text" placeholder="Username" />
      </div>
      <div class="form-field">
        <label>Relationship status</label>
        <select v-model="editForm.relationshipStatus">
          <option value="">Select status</option>
          <option value="Single">Single</option>
          <option value="In a relationship">In a relationship</option>
          <option value="Married">Married</option>
          <option value="Widowed">Widowed</option>
        </select>
      </div>
      <div class="form-field full-width">
        <label>Current city</label>
        <input v-model="editForm.location" type="text" placeholder="Current city" />
      </div>
      <div class="form-field full-width">
        <label>Bio</label>
        <textarea v-model="editForm.bio" rows="3" placeholder="Write something about yourself..."></textarea>
      </div>
      <div class="form-field full-width">
        <label>Joined date</label>
         <input type="text" :value="profile.joinedDate" disabled class="readonly-field" />
      </div>
    </div>

    <p v-if="editError" class="modal-error">{{ editError }}</p>

    <div class="modal-actions">
      <button class="btn-save" @click="saveProfile" :disabled="isSaving">
        {{ isSaving ? 'Saving...' : 'Save' }}
      </button>
      <button class="btn-cancel" @click="closeEditModal">Cancel</button>
    </div>
  </div>
</div>
 </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import NavBar from '../navbar/NavBar.vue'
import PostsCard from '../postcards/PostsCard.vue' 

const BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:7070'
const route = useRoute()

function getAuthToken() {
  return localStorage.getItem('token') || ''
}
function authHeaders() {
  const token = getAuthToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}
function getCurrentUserId() {
  const token = getAuthToken()
  if (!token) return null
  try {
    const payload = token.split('.')[1]
    const decoded = JSON.parse(
      decodeURIComponent(
        atob(payload.replace(/-/g, '+').replace(/_/g, '/'))
          .split('')
          .map((c) => '%' + c.charCodeAt(0).toString(16).padStart(2, '0'))
          .join('')
      )
    )
    return decoded.user_id ?? decoded.uid ?? decoded.sub ?? decoded.id ?? null
  } catch {
    return null
  }
}

const profile = ref({
  userId: null,
  name: "",
  handle: "",
  isPro: false,
  bio: "",
  websiteText: "",
  websiteUrl: "",
  streak: 0,
  location: "",
  joinedDate: "",
  postsCount: 0,
  relationship_status: "",
  avatarUrl: "",
  coverUrl: ""
})

const isLoadingProfile = ref(false)
const profileError = ref(null)

function resolveAvatarUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

function resolveCoverUrl(raw) {
  if (!raw) return ''
  if (raw.startsWith('http://') || raw.startsWith('https://')) return raw
  return `${BASE_URL}/uploads/${raw}`
}

function formatJoinedDate(value) {
  if (!value) return ''
  const d = new Date(value)
  if (Number.isNaN(d.getTime())) return ''
  return 'Joined ' + d.toLocaleString('en-US', { month: 'long', year: 'numeric' })
}

// async function loadProfile() {
//   isLoadingProfile.value = true
//   profileError.value = null

//   const targetUserId = route.params?.id || getCurrentUserId()

//   try {
//     const res = await fetch(
//       `${BASE_URL}/api/v1/front/profile/show?id=${targetUserId}`,
//       { headers: { ...authHeaders() } }
//     )
//     if (!res.ok) {
//       const text = await res.text().catch(() => '')
//       throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
//     }
//     const json = await res.json()
//     const data = json?.data ?? json

//     profile.value = {
//       ...profile.value,
//       userId: data.id,
//       name: `${data.first_name || ''} ${data.last_name || ''}`.trim() || data.user_name,
//       handle: '@' + data.user_name,
//       avatarUrl: resolveAvatarUrl(data.profile_images),
//       coverUrl: resolveCoverUrl(data.cover_images),
//       location: data.location ?? '',                
//       relationship_status: data.relationship_status ?? '', 
//       bio: data.bios ?? '',
//       postsCount: data.post_count ?? 0,
//       joinedDate: formatJoinedDate(data.created_at),
//     }
//   } catch (e) {
//     profileError.value = e.message || 'Failed to load profile'
//     console.error('Failed to load profile', e)
//   } finally {
//     isLoadingProfile.value = false
//   }
// }
async function loadProfile() {
  isLoadingProfile.value = true
  profileError.value = null

  const targetUserId = route.params?.id || getCurrentUserId()
  console.log('🔍 [1] targetUserId:', targetUserId)

  try {
    const url = `${BASE_URL}/api/v1/front/profile/show?id=${targetUserId}`
    console.log('🔍 [2] Fetching URL:', url)

    const res = await fetch(url, { headers: { ...authHeaders() } })
    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }
    const json = await res.json()
    console.log('🔍 [3] Raw response:', JSON.stringify(json))
    
    const data = json?.data ?? json
    console.log('🔍 [4] data.id:', data.id, '| data.user_name:', data.user_name)

    profile.value = {
      ...profile.value,
      userId: data.id,
      name: `${data.first_name || ''} ${data.last_name || ''}`.trim() || data.user_name,
      handle: '@' + data.user_name,
      avatarUrl: resolveAvatarUrl(data.profile_images),
      coverUrl: resolveCoverUrl(data.cover_images),
      location: data.location ?? '',                
      relationship_status: data.relationship_status ?? '', 
      bio: data.bios ?? '',
      postsCount: data.post_count ?? 0,
      joinedDate: formatJoinedDate(data.created_at),
    }
  } catch (e) {
    profileError.value = e.message || 'Failed to load profile'
    console.error('Failed to load profile', e)
  } finally {
    isLoadingProfile.value = false
  }
}

onMounted(() => {
  loadProfile()
})

watch(
  () => route.params.id,
  (newId, oldId) => {
    if (newId !== oldId) {
      currentTab.value = "Done"
      loadProfile()
    }
  }
)

const isFollowing = ref(false)
const currentTab = ref("Done")

const tabs = ref([
  { name: "Done", label: "Posts", count: "2.9K" },
  { name: "Projects", label: "Photo", count: 14 },
  { name: "Posts", label: "Video", count: 3 },
  { name: "Comments", label: "Comments", count: 40 },
  { name: "Media", label: "Share" },
  { name: "Calendar", label: "Group" },
  { name: "Stats", label: "Stats" },
  { name: "Following", label: "Following" }
])

function handleMore() { alert("More options clicked!") }
function handleFollow() { isFollowing.value = !isFollowing.value }

const coverFileInput = ref(null)
const isUploadingCover = ref(false)

function handleChangeCover() {
  coverFileInput.value?.click()
}

async function onCoverFileSelected(event) {
  const file = event.target.files?.[0]
  if (!file) return

  const allowedTypes = ['image/png', 'image/jpeg', 'image/gif']
  if (!allowedTypes.includes(file.type)) {
    alert('សូមជ្រើសរើសរូបភាព PNG, JPG, ឬ GIF ប៉ុណ្ណោះ')
    event.target.value = ''
    return
  }
  if (file.size > 5 * 1024 * 1024) {
    alert('ទំហំរូបភាពត្រូវតែតិចជាង 5MB')
    event.target.value = ''
    return
  }

  isUploadingCover.value = true
  try {
    const formData = new FormData()
    formData.append('cover', file)

    const res = await fetch(
      `${BASE_URL}/api/v1/front/profile/update-cover`,
      {
        method: 'PUT',
        headers: { ...authHeaders() },
        body: formData,
      }
    )

    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }

    const json = await res.json()
    const data = json?.data ?? json
    profile.value.coverUrl = resolveCoverUrl(data.cover_images)
  } catch (e) {
    console.error('Failed to upload cover', e)
    alert('ការ upload cover បរាជ័យ, សូមព្យាយាមម្តងទៀត')
  } finally {
    isUploadingCover.value = false
    event.target.value = ''
  }
}

const showEditModal = ref(false)
const isSaving = ref(false)
const editError = ref(null)

const editForm = ref({
  firstName: '',
  lastName: '',
  username: '',
  bio: '',
  relationshipStatus: '',
  location: '',
})

function openEditModal() {
  const nameParts = profile.value.name.split(' ')
  editForm.value = {
    firstName: nameParts[0] || '',
    lastName: nameParts.slice(1).join(' ') || '',
    username: profile.value.handle.replace('@', ''),
    bio: profile.value.bio || '',
    relationshipStatus: profile.value.relationship_status || '',
    location: profile.value.location || '',
  }
  editError.value = null
  showEditModal.value = true
}

function closeEditModal() {
  showEditModal.value = false
}

async function saveProfile() {
  isSaving.value = true
  editError.value = null

  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/profile/update-info`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        ...authHeaders(),
      },
      body: JSON.stringify({
        first_name: editForm.value.firstName,
        last_name: editForm.value.lastName,
        user_name: editForm.value.username,
        bio: editForm.value.bio,
        relationship_status: editForm.value.relationshipStatus,
        location: editForm.value.location,
      }),
    })

    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }

    profile.value.name = `${editForm.value.firstName} ${editForm.value.lastName}`.trim()
    profile.value.handle = '@' + editForm.value.username
    profile.value.bio = editForm.value.bio
    profile.value.relationship_status = editForm.value.relationshipStatus
    profile.value.location = editForm.value.location

    showEditModal.value = false
  } catch (e) {
    editError.value = e.message || 'Failed to update profile'
    console.error('Failed to save profile', e)
  } finally {
    isSaving.value = false
  }
}

const avatarFileInput = ref(null)
const isUploadingAvatar = ref(false)

function handleChangeAvatar() {
  avatarFileInput.value?.click()
}

async function onAvatarFileSelected(event) {
  const file = event.target.files?.[0]
  if (!file) return

  const allowedTypes = ['image/png', 'image/jpeg', 'image/gif']
  if (!allowedTypes.includes(file.type)) {
    alert('សូមជ្រើសរើសរូបភាព PNG, JPG, ឬ GIF ប៉ុណ្ណោះ')
    event.target.value = ''
    return
  }
  if (file.size > 5 * 1024 * 1024) {
    alert('ទំហំរូបភាពត្រូវតែតិចជាង 5MB')
    event.target.value = ''
    return
  }

  isUploadingAvatar.value = true
  try {
    const formData = new FormData()
    formData.append('file', file)

    const currentUserId = getCurrentUserId()
    const res = await fetch(
      `${BASE_URL}/api/v1/front/profile/update/${currentUserId}`,
      {
        method: 'PUT',
        headers: { ...authHeaders() },
        body: formData,
      }
    )

    if (!res.ok) {
      const text = await res.text().catch(() => '')
      throw new Error(`API ${res.status} ${res.statusText}: ${text}`)
    }

    const json = await res.json()
    const data = json?.data ?? json
    profile.value.avatarUrl = resolveAvatarUrl(data.profile_images)
  } catch (e) {
    console.error('Failed to upload avatar', e)
    alert('ការ upload avatar បរាជ័យ, សូមព្យាយាមម្តងទៀត')
  } finally {
    isUploadingAvatar.value = false
    event.target.value = ''
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700&display=swap');

.profile-page-wrapper {
  padding: 0;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 100vh;
  box-sizing: border-box;
  background-color: #F7F4F2;
}

.profile-container {
  color: #0f172a;
  min-height: 100vh;
  font-family: 'Plus Jakarta Sans', sans-serif;
  max-width: 700px;
  width: 100%;
  margin: 0 auto;
  /* background-color: #ffffff; */
  /* padding-left: 12px;
  padding-right: 12px; */

}

/* Top Bar */
.top-bar {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 4px 16px;
  position: sticky;
  top: 0;
  backdrop-filter: blur(8px);
  z-index: 10;
}

/* Cover Image Section */
.cover-container {
  width: 100%;
  height: 200px;
  background-color: #e2e8f0;
  overflow: hidden;
  position: relative;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
}

.cover-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Gradient Overlay */
.cover-gradient-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 60%;
  background: linear-gradient(to top, rgba(237, 228, 228, 0.679), transparent);
  pointer-events: none;
  z-index: 1;
}

.cover-actions-group {
  position: absolute;
  bottom: 12px;
  right: 12px;
  display: flex;
  align-items: center;
  gap: 8px;
  z-index: 5;
  background: transparent;
}

.action-buttons-row {
  display: flex;
  gap: 8px;
  align-items: center;
}

.icon-action-btn, .change-cover-btn {
  background: rgba(0, 0, 0, 0.45);
  border: 1px solid rgba(255, 255, 255, 0.25);
  color: #ffffff;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  transition: all 0.2s ease;
}

.icon-action-btn:hover, .change-cover-btn:hover {
  background: rgba(0, 0, 0, 0.7);
  transform: scale(1.08);
}

.follow-btn {
  background-color: #ffffff;
  color: #0f172a;
  border: none;
  font-weight: 600;
  font-size: 13px;
  padding: 8px 18px;
  border-radius: 20px;
  cursor: pointer;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  transition: all 0.2s ease;
}

.follow-btn:hover {
  background-color: #f1f5f9;
  transform: scale(1.05);
}

/* 1. Profile Main Card */
.profile-card-section {
  background-color: #ffffff;
  padding: 0 16px;
  position: relative;
  margin-bottom: 2px;
  border-bottom: 0.5px solid #e2e8f0;
  border-top: 1px solid #e2e8f0;
}

.profile-header-row {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-top: -55px;
  margin-bottom: 12px;
}

.avatar-container {
  width: 110px;
  height: 110px;
  border-radius: 50%;
  border: 4px solid #ffffff;
  overflow: hidden;
  background-color: #cbd5e1;
  z-index: 2;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative; 
  cursor: pointer; 

}

.avatar-camera-btn {
  position: absolute;
  bottom: -2px;
  right: -2px;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: #1976D2;
  border: 3px solid #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 3;
  transition: transform 0.15s ease, background 0.15s ease;
  padding: 0;
}

.avatar-camera-btn:hover {
  background: #1560ac;
  transform: scale(1.08);
}

.avatar-wrapper {
  position: relative;    
  width: 110px;
  height: 110px;
  flex-shrink: 0;
}


.avatar-camera-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.avatar-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.profile-title-area {
  display: flex;
  flex-direction: column;
  padding-top: 50px;
}

.name-row {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 2px;
}

.profile-name {
  font-size: 20px;
  font-weight: 700;
  margin: 0;
  color: #0f172a;
}

.profile-handle {
  font-size: 14px;
  color: #64748b;
}

.profile-bio {
  font-size: 15px;
  line-height: 1.5;
  color: #334155;
  margin: 0 0 16px 0;
}

.bio-link {
  color: #2563eb;
  text-decoration: none;
}

.bio-link:hover {
  text-decoration: underline;
}

/* 2. Profile Meta Card */
.profile-meta-card {
  background-color: #ffffff;
  padding: 14px 16px;
  margin-bottom: 10px;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.profile-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  font-size: 14px;
  color: #64748b;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

/* 3. Navigation Tabs Card */
.tabs-card-section {
  background-color: #ffffff;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
}

/* Navigation Tabs */
.tabs-container {
  display: flex;
  border-bottom: 1px solid #e2e8f0;
  overflow-x: auto;
  scrollbar-width: none;
}

.tabs-container::-webkit-scrollbar {
  display: none;
}

.tab-item {
  flex: 0 0 auto;
  padding: 16px 16px;
  font-size: 14px;
  font-weight: 500;
  color: #64748b;
  cursor: pointer;
  text-align: center;
  position: relative;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: color 0.2s;
  white-space: nowrap;
}

.tab-item:hover {
  color: #1976D2;
  background: rgba(0, 0, 0, 0.02);
}

.tab-item.active {
  color: #1976D2;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 3px;
  background-color: #1976D2;
  border-radius: 2px 2px 0 0;
}

.tab-count {
  font-size: 12px;
  background: #e2e8f0;
  padding: 1px 6px;
  border-radius: 10px;
  color: #475569;
}

.content-body {
  background-color: #F7F4F2;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.55);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 16px;
}

.modal-box {
  background: #fff;
  border-radius: 20px;
  padding: 32px;
  max-width: 640px;
  width: 100%;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.25);
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
}

.modal-title {
  font-size: 22px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

.modal-close-btn {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  border: none;
  background: #1f2937;
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  flex-shrink: 0;
}

.modal-close-btn:hover {
  background: #374151;
}

.modal-form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px 24px;
}

.form-field {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-field.full-width {
  grid-column: 1 / -1;
}

.form-field label {
  font-size: 14px;
  font-weight: 600;
  color: #0f172a;
}

.form-field input,
.form-field select,
.form-field textarea {
  border: 1.5px solid #e2e8f0;
  border-radius: 10px;
  padding: 12px 14px;
  font-size: 14px;
  font-family: 'Plus Jakarta Sans', sans-serif;
  color: #0f172a;
  outline: none;
  transition: border-color 0.15s;
  resize: none;
}

.form-field input:focus,
.form-field select:focus,
.form-field textarea:focus {
  border-color: #1976D2;
}

.modal-error {
  color: #dc2626;
  font-size: 13px;
  margin-top: 12px;
  text-align: center;
}

.modal-actions {
  display: flex;
  justify-content: center;
  gap: 12px;
  margin-top: 28px;
}

.btn-save {
  background-color: #1976D2;
  color: #fff;
  font-weight: 700;
  font-size: 15px;
  padding: 12px 32px;
  border-radius: 999px;
  border: none;
  cursor: pointer;
  transition: opacity 0.2s;
}

.btn-save:hover {
  opacity: 0.9;
}

.btn-save:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-cancel {
  background-color: #e5e7eb;
  color: #4b5563;
  font-weight: 700;
  font-size: 15px;
  padding: 12px 32px;
  border-radius: 999px;
  border: none;
  cursor: pointer;
}

.btn-cancel:hover {
  background-color: #d1d5db;
}
.readonly-field {
  color: #64748b;
  cursor: not-allowed;
}
.avatar-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: all 0.2s ease;
}

.avatar-container:hover .avatar-overlay {
  background: rgba(0, 0, 0, 0.45);
  opacity: 1;
}

.avatar-uploading {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.45);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 11px;
  font-weight: 600;
}
</style>