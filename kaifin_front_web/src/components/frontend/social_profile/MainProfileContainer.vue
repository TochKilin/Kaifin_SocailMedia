<template>
  <div class="app-layout">
    <header class="navbar-wrapper">
      <NavBar/>
    </header>
    <div class="app-body">
      <div class="content-container">
        <aside class="sidebar-wrapper">
          <LeftMenu/>
        </aside>
        <main class="main-content-wrapper">
          <SocialProfile :userId="targetUserId" />
        </main>

        <aside class="profile-right-section">
          <RightMedia :userId="targetUserId" />
        </aside>
      </div>
    </div>
  </div>
</template>

<script setup>
import NavBar from '../navbar/NavBar.vue';
import LeftMenu from './LeftMenu.vue';
import RightMedia from './RightMedia.vue';
import SocialProfile from './SocialProfile.vue';
import { computed } from 'vue'
import { useRoute } from 'vue-router'


const route = useRoute()

function getAuthToken() {
  return localStorage.getItem('token') || ''
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

const targetUserId = computed(() => route.params?.id || getCurrentUserId())
</script>

<style>

.app-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background-color: #F7F4F2;
  font-family: 'Plus Jakarta Sans', sans-serif;
}

.navbar-wrapper {
  position: sticky;
  top: 0;
  z-index: 100;
  width: 100%;
}


.app-body {
  width: 100%;
  display: flex;
  justify-content: center;
  flex-grow: 1;
}


.content-container {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  width: 100%;
  max-width: 1251px; 

  gap: 8px;         
}


.sidebar-wrapper {
  position: sticky;
  top: 60px;
  height: calc(100vh - 60px);
  overflow-y: auto;
  flex-shrink: 0;
  width: 160px;     
  z-index: 20;

  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  scrollbar-width: none;
}


.main-content-wrapper {
  flex: 1;
  max-width: 700px; 
  min-height: calc(100vh - 60px);
  background-color: #F7F4F2;
  overflow-y: auto;
  min-width: 0;

  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  scrollbar-width: none;
}


.profile-right-section {
  position: sticky;
  top: 60px;
  width: 175px;     
  height: calc(100vh - 60px);
  overflow-y: auto;
  flex-shrink: 0;
  background-color: #f7f4f2;
  z-index: 20;

  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>