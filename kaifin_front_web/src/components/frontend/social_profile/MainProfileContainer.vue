<template>
  <div class="app-layout">
    <!-- Navbar នៅផ្នែកខាងលើគេបង្អស់ -->
    <header class="navbar-wrapper">
      <NavBar/>
    </header>

    <!-- Body Layout: រៀបចំឱ្យនៅចំកណ្តាល និងជិតគ្នា -->
    <div class="app-body">
      <div class="content-container">
        <!-- Sidebar នៅខាងឆ្វេង (កាត់បន្ថយទំហំ width ឱ្យចូលមកជិត) -->
        <aside class="sidebar-wrapper">
          <LeftMenu/>
        </aside>

        <!-- Profile Page នៅកណ្តាល (រក្សាទំហំធំទូលាយស្អាត) -->
        <main class="main-content-wrapper">
          <SocialProfile :userId="targetUserId" />
        </main>

        <!-- Right Media នៅខាងស្តាំ -->
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
/* កំណត់ទម្រង់ទំព័រទាំងមូល */
.app-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background-color: #F7F4F2;
  font-family: 'Plus Jakarta Sans', sans-serif;
}

/* រៀបចំ Navbar ให้อยជាប់នៅខាងលើគេ */
.navbar-wrapper {
  position: sticky;
  top: 0;
  z-index: 100;
  width: 100%;
}

/* កន្លែងសម្រាប់ទប់ទំហំកណ្តាលអេក្រង់ */
.app-body {
  width: 100%;
  display: flex;
  justify-content: center;
  flex-grow: 1;
}

/* Container រួម */
.content-container {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  width: 100%;
  max-width: 1251px; 
  /* background-color: red; */
  gap: 8px;         /* រក្សាគលាតរវាងផ្ទាំងនីមួយៗឱ្យជិតគ្នាជាប់ល្អ */
}

/* រៀបចំ Sidebar ឱ្យនៅខាងឆ្វេង */
.sidebar-wrapper {
  position: sticky;
  top: 60px;
  height: calc(100vh - 60px);
  overflow-y: auto;
  flex-shrink: 0;
  width: 160px;     /* បន្ថយទំហំ Width ខាងឆ្វេងមកត្រឹម 160px ដើម្បីឱ្យវាទាញចូលមកជិត Profile (អាចសារ៉េ 150px-165px បាន) */
  /* background-color: #f7f4f2; */
  z-index: 20;

  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  scrollbar-width: none;
}

/* រៀបចំ Profile Page នៅកណ្តាល (រក្សាទំហំធំទូលាយស្អាតពេញចិត្ត) */
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

/* រៀបចំ Right Media ឱ្យនៅខាងស្តាំ */
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