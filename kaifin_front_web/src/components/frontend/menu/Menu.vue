<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const menuItems = ref([])
const isLoading = ref(true)
const loadError = ref('')
const showProfileMenu = ref(false)
const profile = ref(null)
const profileError = ref('')

function slugify(title) {
  return (
    '/' +
    title
      .toLowerCase()
      .trim()
      .replace(/[^a-z0-9]+/g, '-')
      .replace(/^-+|-+$/g, '')
  )
}

function normalizeMenuItem(item) {
  return {
    label: item.title,
    path: slugify(item.title),
  }
}

async function fetchMenus() {
  isLoading.value = true
  loadError.value = ''

  const token = localStorage.getItem('token')

  try {
    const response = await axios.get(
      'http://localhost:7070/api/v1/front/menus/show',
      { headers: { Authorization: `Bearer ${token}` } }
    )

    const rawItems = response.data?.data?.menus ?? []

    menuItems.value = rawItems
      .filter(item => item.is_active)
      .sort((a, b) => a.sort_order - b.sort_order)
      .map(normalizeMenuItem)

  } catch (error) {
    console.error("Failed to load menu:", error)
    loadError.value = "Could not load menu"
  } finally {
    isLoading.value = false
  }
}

async function fetchProfile() {
  const token = localStorage.getItem('token')

  if (!token) return

  try {
    const response = await axios.get(
      "http://localhost:7070/api/v1/front/profile/show",
      {
        headers: {
          Authorization: `Bearer ${token}`
        }
      }
    )

    console.log("PROFILE RESPONSE:", response.data)

    profile.value = response.data.data

  } catch (error) {
    profileError.value = "Could not load profile"
    console.log(error.response?.data)
  }
}

function logout() {
  localStorage.removeItem('token')
  window.location.href = '/login'
}

onMounted(() => {
  fetchMenus()
  fetchProfile()
})
</script>

<template>
  <nav class="navbar">
    <ul class="menu-list">
      <li v-if="isLoading" class="menu-status">Loading...</li>
      <li v-else-if="loadError" class="menu-status">{{ loadError }}</li>

      <li v-for="item in menuItems" :key="item.path" class="menu-item">
        <router-link
          :to="item.path"
          class="menu-link"
          :class="{ active: route.path === item.path }"
        >
          {{ item.label }}
        </router-link>
      </li>
    </ul>

    <div class="nav-actions">
      <button class="icon-btn" aria-label="Messages">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <rect x="3" y="5" width="18" height="14" rx="2" />
          <path d="M3 7l9 6 9-6" />
        </svg>
      </button>

      <div class="avatar-wrapper">
        <div
          class="icon-btn avatar-btn"
          @click="showProfileMenu = !showProfileMenu"
          aria-label="Profile menu"
        >
          <img
            v-if="profile?.profile_images"
            :src="`http://localhost:7070/uploads/${profile.profile_images}`"
            alt="Profile avatar"
          />
          <svg
            v-else
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <circle cx="12" cy="8" r="4" />
            <path d="M4 20c0-4.4 3.6-8 8-8s8 3.6 8 8" />
          </svg>
        </div>

        <div v-if="showProfileMenu" class="profile-dropdown">
          <div class="profile-header">
            <div class="profile-avatar">
              <img
                v-if="profile?.profile_images"
                :src="`http://localhost:7070/uploads/${profile.profile_images}`"
                alt="Profile avatar"
              />
              <span v-else class="avatar-fallback">
                {{ (profile?.first_name?.[0] ?? '') + (profile?.last_name?.[0] ?? '') }}
              </span>
            </div>
            <div class="profile-info">
              <p class="profile-name">{{ profile?.first_name }} {{ profile?.last_name }}</p>
              <p class="profile-email">{{ profile?.email }}</p>
            </div>
          </div>

          <span v-if="profile?.role_name" class="role-badge">{{ profile.role_name }}</span>
          <p v-if="profileError" class="profile-error">{{ profileError }}</p>

          <div class="profile-actions">
            <button class="profile-btn">Edit profile</button>
            <button class="profile-btn">Settings</button>
            <button class="profile-btn danger" @click="logout">Log out</button>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  background: #fff;
  border: 3px solid #1a1a1a;
  border-radius: 999px;
  padding: 8px 10px;
  font-family: 'Baloo 2', Arial, sans-serif;
  max-width: 1200px;
  margin: 20px auto;
}

.menu-list {
  display: flex;
  align-items: center;
  list-style: none;
  margin: 0;
  padding: 0;
  flex: 1;
  overflow-x: auto;
}

.menu-status {
  padding: 12px 20px;
  font-size: 14px;
  color: #666;
}

.menu-item {
  display: flex;
  align-items: center;
  border-right: 2px solid #1a1a1a;
}

.menu-item:last-child {
  border-right: none;
}

.menu-link {
  display: block;
  padding: 12px 20px;
  font-weight: 700;
  font-size: 15px;
  color: #1a1a1a;
  text-decoration: none;
  white-space: nowrap;
  border-radius: 999px;
  transition: background 0.15s, color 0.15s;
}

.menu-link:hover {
  background: #f2f2f2;
}

.menu-link.active {
  background: #8fe3ad;
  color: #2a4bff;
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: 8px;
  padding-left: 8px;
  border-left: 2px solid #1a1a1a;
}

.icon-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border: 2px solid #1a1a1a;
  border-radius: 10px;
  background: #fff;
  color: #1a1a1a;
  cursor: pointer;
  transition: background 0.15s;
}

.icon-btn img {
    display: block;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.icon-btn:hover {
  background: #f2f2f2;
}

.avatar-wrapper {
  position: relative;
}

.avatar-btn {
  border-radius: 50%;
  overflow: hidden;
}

.profile-dropdown {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  width: 260px;
  background: #fff;
  border: 3px solid #1a1a1a;
  border-radius: 20px;
  padding: 20px;
  z-index: 100;
}

.profile-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.profile-avatar {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  background: #8fe3ad;
  border: 2px solid #1a1a1a;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 18px;
  color: #1a1a1a;
  flex-shrink: 0;
  overflow: hidden;
}

.profile-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-fallback {
  font-family: 'Baloo 2', Arial, sans-serif;
}

.profile-info {
  min-width: 0;
}

.profile-name {
  margin: 0;
  font-weight: 700;
  font-size: 16px;
  color: #1a1a1a;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.profile-email {
  margin: 2px 0 0;
  font-size: 13px;
  color: #666;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.role-badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: #8fe3ad;
  color: #173404;
  font-size: 12px;
  font-weight: 700;
  padding: 4px 12px;
  border-radius: 999px;
  margin-bottom: 16px;
}

.profile-error {
  font-size: 13px;
  color: #993c1d;
  margin: 0 0 12px;
}

.profile-actions {
  border-top: 2px solid #1a1a1a;
  padding-top: 12px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.profile-btn {
  display: flex;
  align-items: center;
  width: 100%;
  background: #fff;
  border: 2px solid #1a1a1a;
  border-radius: 12px;
  padding: 10px 12px;
  font-family: inherit;
  font-weight: 700;
  font-size: 14px;
  color: #1a1a1a;
  cursor: pointer;
  text-align: left;
  transition: background 0.15s;
}

.profile-btn:hover {
  background: #f2f2f2;
}

.profile-btn.danger {
  border-color: #d85a30;
  color: #993c1d;
}

@media (max-width: 900px) {
  .navbar {
    border-radius: 20px;
  }
  .menu-link {
    padding: 10px 14px;
    font-size: 13px;
  }
}
</style>