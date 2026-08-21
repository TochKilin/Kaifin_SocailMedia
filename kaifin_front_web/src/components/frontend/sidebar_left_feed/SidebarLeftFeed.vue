<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const emit = defineEmits(['open-group-join', 'open-following', 'open-populars'])
const props = defineProps({
  activeView: { type: String, default: 'feed' }
})

const BASE_URL = 'http://localhost:7070'

function authHeaders() {
  const token = localStorage.getItem('token')
  return token ? { Authorization: `Bearer ${token}` } : {}
}

function resolveImageUrl(path) {
  if (!path) return ''
  if (path.startsWith('http://') || path.startsWith('https://')) return path
  return `${BASE_URL}/uploads/${path}`
}

const showMyGroups = ref(false)
const myGroups = ref([])
const loadingMyGroups = ref(false)
const myGroupsError = ref('')

async function fetchMyGroups() {
  loadingMyGroups.value = true
  myGroupsError.value = ''
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/communities/show`, {
      params: { page: 1, perpage: 50 },
      headers: authHeaders(),
    })
    const list = res.data?.data?.communities ?? []

    myGroups.value = list
      .filter((c) => c.is_joined === true)
      .map((c) => ({
        id: c.id,
        name: c.name,
        description: c.description || 'Community',
        avatar: resolveImageUrl(c.avatar_url),
      }))
  } catch (err) {
    console.error('Failed to fetch my groups:', err)
    myGroupsError.value = 'Failed to load groups'
  } finally {
    loadingMyGroups.value = false
  }
}

function toggleMyGroups() {
  showMyGroups.value = !showMyGroups.value
  if (showMyGroups.value && myGroups.value.length === 0 && !loadingMyGroups.value) {
    fetchMyGroups()
  }
}

function goToGroup(groupId) {
  router.push({ name: 'GroupDetail', params: { id: groupId } })
}

const recommendedGroups = ref([])
const loadingRecommended = ref(false)

async function fetchRecommendedGroups() {
  loadingRecommended.value = true
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/communities/show`, {
      params: { page: 1, perpage: 10 },
      headers: authHeaders(),
    })
    const list = res.data?.data?.communities ?? []

    recommendedGroups.value = list
      .filter((c) => c.is_joined !== true)
      .slice(0, 3) 
      .map((c) => ({
        id: c.id,
        name: c.name,
        avatar: resolveImageUrl(c.avatar_url),
        isJoining: false,
      }))
  } catch (err) {
    console.error('Failed to fetch recommended groups:', err)
  } finally {
    loadingRecommended.value = false
  }
}

async function joinRecommendedGroup(group) {
  if (group.isJoining) return
  group.isJoining = true
  try {
    await axios.post(
      `${BASE_URL}/api/v1/front/communities/${group.id}/join`,
      {},
      { headers: authHeaders() }
    )
    recommendedGroups.value = recommendedGroups.value.filter((g) => g.id !== group.id)
  } catch (err) {
    console.error('Failed to join group:', err)
    group.isJoining = false
  }
}

onMounted(() => {

  fetchMyGroups()
   fetchRecommendedGroups() 
})
</script>

<template>
<div class="left-search-contianer">
  <div class="search-box">
    <svg viewBox="0 0 24 24" class="search-svg-icon">
      <circle cx="11" cy="11" r="8" fill="none" stroke="currentColor" stroke-width="2"/>
      <path d="M21 21l-4.35-4.35" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
    </svg>
    <input type="search" placeholder="Search your friends..." />
  </div>
</div>
    <div class="left_container_feed">
       <div>
        <div class="left-btn font-btn btn-feed-color btn-icon" style="cursor: pointer;" @click="$emit('open-news-feed')"
          :class="{ 'left-btn-active': activeView === 'feed' }"
        >
          <span class="bg-left-side">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 11h9M4 16h5M4 6h16" stroke-linecap="round"/>
            </svg>
          </span>
          New Feeds
        </div>
        <div class="left-btn font-btn btn-icon left-btn-hover" style="cursor: pointer;" @click="$emit('open-populars')"
        :class="{ 'left-btn-active': activeView === 'populars' }"
        >
          <span class="bg-left-side">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 2c0 4-4 7-4 10a4 4 0 0 0 8 0c0-3-4-6-4-10z" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M12 18v2" stroke-linecap="round"/>
            </svg>
          </span>
          Populars
        </div>
        <div class="left-btn font-btn btn-icon left-btn-hover" style="cursor: pointer;" @click="$emit('open-following')"
        :class="{ 'left-btn-active': activeView === 'following' }"
        >
          <span class="bg-left-side">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" stroke-linecap="round"/>
              <circle cx="9" cy="7" r="4"/>
              <polyline points="16 11 18 13 22 9" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </span>
          Following
        </div>
        <div class="left-btn font-btn btn-icon left-btn-hover" style="cursor: pointer;" @click="toggleMyGroups">
          <span class="bg-left-side">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
            </svg>
          </span>
          My Group
          <svg
            class="chevron-icon"
            :class="{ rotated: showMyGroups }"
            viewBox="0 0 24 24"
            width="16"
            height="16"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <polyline points="6 9 12 15 18 9" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </div>

        <div v-if="showMyGroups" class="my-groups-list">
          <div v-if="loadingMyGroups" class="group-skeleton">
            <div class="group-item-skeleton" v-for="n in 2" :key="n"></div>
          </div>

          <div v-else-if="myGroupsError" class="group-empty">
            {{ myGroupsError }}
          </div>

          <div v-else-if="myGroups.length === 0" class="group-empty">
            You haven't joined any group yet
          </div>

          <div
            v-else
            v-for="group in myGroups"
            :key="group.id"
            class="group-item"
            @click="goToGroup(group.id)"
          >
            <span class="group-icon">
              <img v-if="group.avatar" :src="group.avatar" :alt="group.name" />
              <svg v-else viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2.5">
                <path d="M23 4v6h-6" stroke-linecap="round" stroke-linejoin="round"/>
                <path d="M20.49 15a9 9 0 1 1-2.12-9.36L23 10" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </span>
            <div class="group-info">
              <span class="group-name">{{ group.name }}</span>
              <span class="group-desc">{{ group.description }}</span>
            </div>
          </div>
        </div>

        <div class="group-recommend-section" v-if="recommendedGroups.length > 0 || loadingRecommended">
        <div class="recommend-title font-btn">Group Recommend</div>

        <div v-if="loadingRecommended" class="group-skeleton">
          <div class="recommend-item-skeleton" v-for="n in 2" :key="n"></div>
        </div>

          <div v-else class="recommend-list">
            <div v-for="group in recommendedGroups" :key="group.id" class="recommend-item">
              <div class="recommend-avatar">
                <img v-if="group.avatar" :src="group.avatar" :alt="group.name" />
                <svg v-else viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
                  <circle cx="9" cy="7" r="4"/>
                </svg>
              </div>
              <span class="recommend-name">{{ group.name }}</span>
              <button
                class="recommend-join-btn"
                :disabled="group.isJoining"
                @click="joinRecommendedGroup(group)"
              >
                {{ group.isJoining ? '...' : 'Join' }}
              </button>
            </div>
          </div>
        </div>
        <div class="left-btn font-btn btn-icon left-btn-hover" @click="$emit('open-group-join')"
        :class="{ 'left-btn-active': activeView === 'group-join' }"
        >
          <span class="bg-left-side">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <line x1="19" y1="8" x2="19" y2="14" stroke-linecap="round"/>
              <line x1="22" y1="11" x2="16" y2="11" stroke-linecap="round"/>
            </svg>
          </span>
          Group Join
        </div>
       </div>
       <div class="view-this-posts">
          <span>Views Of this posts</span>
       </div>
    </div>
</template>

<style scoped>
  .search-svg-icon {
    width: 18px;
    height: 18px;
    color: #7d8794;
    flex-shrink: 0;
  }

  .btn-icon svg {
    width: 18px;
    height: 18px;
    display: inline-block;
    vertical-align: middle;
    color: #000;
  }

  .btn-feed-color .btn-icon svg,
  .btn-feed-color .bg-left-side svg {
    color: #000;
  }

  .font-btn{
    font-family: 'Nunito', sans-serif;
    color: #2b2b2b;
  }

.left-search-contianer{
    background: #fff;
    border: 1px solid #E5E7EB;
    border-radius: 12px;
    padding: 12px;
    margin-top: 12px;
}

.search-box{
    display: flex;
    align-items: center;
    gap: 12px;
    width: 100%;
    height: 35px;
    background: #ffffff;
    border: 1px solid #dfe3ea;
    border-radius: 32px;
    padding: 0 16px;
    transition: .25s ease;
}

.search-box:focus-within{
    border-color: #1976D2;
    background: #fff;
}

.search-box input{
    width: 100%;
    height: 100%;
    border: none;
    outline: none;
    background: transparent;
    font-size: 15px;
    color: #1f2937;
}

.search-box input::placeholder{
    color: #9ca3af;
}

  .left_container_feed {
    background: #ffffff;
    border: 1px solid #E5E7EB;
    padding: 12px;
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 12px;
    border-radius: 12px;
    margin-top: 10px;
  }

  .left-btn{
    width: 100%;
    margin: 8px 0;
    padding: 7px 12px 7px 12px;
    border-radius: 32px;
    font-size: 14px;
    font-weight: 600;
    display: flex;     
    align-items: center; 
    gap: 12px;
    cursor: pointer;
  }

  .left-btn-hover:hover{
    background-color: #0000001c;
  }

  .btn-feed-color{
    color: #000;
  }

  .chevron-icon {
    margin-left: auto;
    transition: transform 0.2s ease;
    color: #6b7280;
  }

  .chevron-icon.rotated {
    transform: rotate(180deg);
  }

  .my-groups-list {
    display: flex;
    flex-direction: column;
    gap: 6px;
    padding: 4px 4px 4px 44px;
    margin-top: -4px;
  }

  .group-item {
    display: flex;
    align-items: center;
    gap: 10px;
    background-color: #ECEAE4;
    border-radius: 8px;
    padding: 10px 14px;
    cursor: pointer;
    transition: background-color 0.2s ease;
  }

  .group-item:hover {
    opacity: 0.8;;
  }

  .group-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 24px;
    height: 24px;
    border-radius: 50%;
    background-color: #1976D2;
    color: #fff;
    flex-shrink: 0;
    overflow: hidden;
  }

  .group-icon img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .group-info {
    display: flex;
    flex-direction: column;
    min-width: 0;
  }

  .group-name {
    font-family: 'Nunito', sans-serif;
    font-size: 14px;
    font-weight: 700;
    color: #1976D2;
  }

  .group-desc {
    font-family: 'Nunito', sans-serif;
    font-size: 12px;
    color: #9ca3af;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    max-width: 180px;
  }

  .group-empty {
    font-family: 'Nunito', sans-serif;
    font-size: 12px;
    color: #9ca3af;
    padding: 8px 14px;
  }

  .group-item-skeleton {
    height: 44px;
    border-radius: 12px;
    background: #f0f0f0;
    margin-bottom: 6px;
    animation: pulse 1.2s infinite ease-in-out;
  }

  @keyframes pulse {
    0%, 100% { opacity: 1; }
    50% { opacity: 0.5; }
  }

  .btn-two-group {
    display: flex;
    flex-direction: column;
    align-items: self-start;  
    gap: 8px;              
    width: 100%;
    padding: 7px 12px;
    border-radius: 12px;
  }

  .btn-small {
  width: fit-content;   
  padding: 7px 12px;
    background-color: #BBDEFB;
  border-left: 6px solid #1976D2;
  border-right: 6px solid #1976D2;
  color: #000;
  border-radius: 4px;
  text-align: center;
  cursor: pointer;
  font-size: 12px;
  font-weight: 700;
}

.btn-small:hover {
  opacity: 0.8;
}

.bg-left-side{
  border-radius: 50%;
  padding: 3px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.view-this-posts{
  width: 160px;
  padding: 6px 8px 6px 8px;
  text-align: center;
  background-color: #1976d2;
  color: #ffff;
  border-radius: 32px;
  margin: auto;
  margin-top: 20px;
  font-size: 14px;
}
.group-recommend-section {
  width: 100%;
  padding: 4px 4px 4px 4px;
}

.recommend-title {
  font-size: 12px;
  font-weight: 700;
  color: #6b7280;
  padding: 4px 8px 8px;
}

.recommend-list {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.recommend-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 7px 10px;
  border-radius: 32px;
  transition: background-color 0.2s ease;
}

.recommend-item:hover {
  background-color: #0000001c;
}

.recommend-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: #BBDEFB;
  color: #1976D2;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  flex-shrink: 0;
}

.recommend-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.recommend-name {
  flex: 1;
  font-family: 'Nunito', sans-serif;
  font-size: 12px;
  font-weight: 600;
  color: #2b2b2b;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.recommend-join-btn {
  background-color: #F4F1EC;
  color: #000;
  border: none;
  border-radius: 32px;
  padding: 5px 12px;
  font-size: 12px;
  font-weight: 700;
  cursor: pointer;
  flex-shrink: 0;
  transition: opacity 0.2s ease;
}

.recommend-join-btn:hover {
  opacity: 0.85;
}

.recommend-join-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.recommend-item-skeleton {
  height: 40px;
  border-radius: 32px;
  background: #f0f0f0;
  margin-bottom: 6px;
  animation: pulse 1.2s infinite ease-in-out;
}

.left-btn-active {
  /* background-color: #1976D2; */
  color: #1976D2;
}

.left-btn-active .bg-left-side svg,
.left-btn-active svg {
  color: #000;
  stroke: #000;
}
</style>