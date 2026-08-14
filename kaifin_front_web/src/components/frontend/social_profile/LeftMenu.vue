<template>
  <div class="sidebar-wrapper">
    <!-- ដុំទី ១៖ ប៊ូតុង Create Post -->
    <div class="sidebar-card border-create-post">
      <button class="post-btn" @click="handlePost">
        <div class="post-icon-wrapper">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <line x1="12" y1="5" x2="12" y2="19"></line>
            <line x1="5" y1="12" x2="19" y2="12"></line>
          </svg>
        </div>
        <span>Create Post</span>
      </button>
    </div>

    <!-- ដុំទី ២៖ មីនុយមេ (Main Menu List) -->
    <div class="sidebar-card">
      <div class="menu-list">
        <router-link to="/feed" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path><polyline points="9 22 9 12 15 12 15 22"></polyline></svg>
          <span class="truncate">Feed</span>
        </router-link>

        <router-link to="/friends" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
          <span class="truncate">Friends</span>
        </router-link>

        <router-link to="/photos" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
          <span class="truncate">Photo</span>
        </router-link>

        <router-link to="/groups" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
          <span class="truncate">Groups</span>
        </router-link>

        <router-link to="/games" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="6" y1="12" x2="18" y2="12"></line><line x1="12" y1="6" x2="12" y2="18"></line></svg>
          <span class="truncate">Games</span>
        </router-link>

        <router-link to="/gifts" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 12 20 22 4 22 4 12"></polyline><rect x="2" y="7" width="20" height="5"></rect><line x1="12" y1="22" x2="12" y2="7"></line><path d="M12 7H7.5a2.5 2.5 0 0 1 0-5C11 2 12 7 12 7z"></path><path d="M12 7h4.5a2.5 2.5 0 0 0 0-5C13 2 12 7 12 7z"></path></svg>
          <span class="truncate">Gifts</span>
        </router-link>

        <router-link to="/recommendations" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polygon points="16.24 7.76 14.12 14.12 7.76 16.24 9.88 9.88 16.24 7.76"></polygon></svg>
          <span class="truncate">Recommend</span>
        </router-link>

        <router-link to="/apps" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="7" height="7"></rect><rect x="14" y="3" width="7" height="7"></rect><rect x="14" y="14" width="7" height="7"></rect><rect x="3" y="14" width="7" height="7"></rect></svg>
          <span class="truncate">Apps</span>
        </router-link>

        <router-link to="/dating" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
          <span class="truncate">Dating</span>
        </router-link>

        <router-link to="/poster" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"></path><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path></svg>
          <span class="truncate">Poster</span>
        </router-link>

        <router-link to="/payments" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="1" y="4" width="22" height="16" rx="2" ry="2"></rect><line x1="1" y1="10" x2="23" y2="10"></line></svg>
          <span class="truncate">Payments</span>
        </router-link>

        <router-link to="/buy-oks" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8"></path><line x1="12" y1="18" x2="12" y2="22"></line><line x1="12" y1="2" x2="12" y2="6"></line></svg>
          <span class="truncate">Buy Coins</span>
        </router-link>

        <router-link to="/advertising" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="23 6 13.5 15.5 8.5 10.5 1 18"></polyline><polyline points="17 6 23 6 23 12"></polyline></svg>
          <span class="truncate">Ads Area</span>
        </router-link>

        <router-link to="/bookmarks" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path></svg>
          <span class="truncate">Bookmarks</span>
        </router-link>

        <router-link to="/holidays" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="8" r="7"></circle><polyline points="8.21 13.89 7 23 12 20 17 23 15.79 13.88"></polyline></svg>
          <span class="truncate">Holidays</span>
        </router-link>

        <router-link to="/find-friends" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="8.5" cy="7" r="4"></circle><line x1="20" y1="8" x2="20" y2="14"></line><line x1="23" y1="11" x2="17" y2="11"></line></svg>
          <span class="truncate">Find friends</span>
        </router-link>

        <router-link to="/moderator" class="menu-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path></svg>
          <span class="truncate">Moderator</span>
        </router-link>
      </div>
    </div>

    <!-- ដុំទី ៖ Invisible Mode និង Footer Links -->
    <div class="sidebar-card">
      <div class="menu-list">
        <div class="menu-item toggle-item">
          <div class="menu-left-content">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path><line x1="1" y1="1" x2="23" y2="23"></line></svg>
            <span class="truncate">Invisible mode</span>
          </div>
          <label class="switch">
            <input type="checkbox" v-model="invisibleMode" />
            <span class="slider round"></span>
          </label>
        </div>
      </div>

      <div class="divider"></div>

      <div class="sidebar-footer">
        <div class="footer-links">
          <a href="#">Ads</a>
          <span>•</span>
          <a href="#">Help</a>
          <span>•</span>
          <a href="#">News</a>
          <span>•</span>
          <a href="#">More</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const invisibleMode = ref(false)

function handlePost() {
  alert('Post button clicked!')
}
</script>

<style scoped>
.sidebar-wrapper {
  width: 155px; /* កាត់បន្ថយ Width ឱ្យតូចជាងមុនបន្ថែមទៀត (155px) ដោយរក្សាទំហំ Font ដើម */
  font-family: 'Plus Jakarta Sans', sans-serif;
  display: flex;
  flex-direction: column;
  gap: 4px;
  box-sizing: border-box;
  margin-top: 8px;
}

/* រចនាប័ទ្មដុំនីមួយៗ (Cards) */
.sidebar-card {
  background-color: #ffffff;
  border: 1px solid #edf2f7;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.03);
  padding: 8px 6px;
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  color: #1e293b;
}

/* .border-create-post{
  border: 2px solid #1976D2;
} */

/* Post Button */
.post-btn {
  background-color: #1976D2;
  color: #ffffff;
  border: none;
  font-weight: 600;
  font-size: 13px; /* រក្សាទំហំ Font ដើម */
  padding: 8px 6px;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  width: 100%;
  transition: background-color 0.2s ease;
}

.post-btn:hover {
  background-color: #115293;
}

.post-btn svg {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  padding: 2px;
  flex-shrink: 0;
}

/* Menu List Layout */
.menu-list {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 6px;
  border-radius: 8px;
  text-decoration: none;
  color: #475569;
  font-size: 13px; /* រក្សាទំហំ Font ដើមមិនឱ្យតូចទេ */
  font-weight: 500;
  transition: background-color 0.2s ease, color 0.2s ease;
  overflow: hidden;
}

.menu-item:hover {
  background-color: #f8fafc;
  color: #0f172a;
}

.menu-item.router-link-active {
  background-color: #e3f2fd;
  color: #1976D2;
  font-weight: 600;
}

.menu-item.router-link-active svg {
  color: #1976D2;
}

.menu-item svg {
  color: #94a3b8;
  flex-shrink: 0;
  transition: color 0.2s;
}

.menu-item:hover svg {
  color: #475569;
}

.truncate {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis; /* អក្សរវែងពេកនឹងលេចចេញជាសញ្ញា ... ដោយស្វ័យប្រវត្តិ */
}

/* Toggle Switch Styling */
.toggle-item {
  justify-content: space-between;
  cursor: pointer;
  width: 100%;
}

.menu-left-content {
  display: flex;
  align-items: center;
  gap: 6px;
  overflow: hidden;
  flex-grow: 1;
}

.switch {
  position: relative;
  display: inline-block;
  width: 30px;
  height: 16px;
  flex-shrink: 0;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0; left: 0; right: 0; bottom: 0;
  background-color: #cbd5e1;
  transition: background-color .2s;
  border-radius: 20px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 12px;
  width: 12px;
  left: 2px;
  bottom: 2px;
  background-color: white;
  transition: transform .2s;
  border-radius: 50%;
}

input:checked + .slider {
  background-color: #1976D2;
}

input:checked + .slider:before {
  transform: translateX(14px);
}

/* Dividers */
.divider {
  height: 1px;
  background-color: #f1f5f9;
  margin: 8px 2px;
}

/* Footer Links */
.sidebar-footer {
  padding: 4px 2px 2px 2px;
  font-size: 11px;
  color: #94a3b8;
}

.footer-links {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 2px;
}

.footer-links a {
  color: #94a3b8;
  text-decoration: none;
  transition: color 0.2s;
}

.footer-links a:hover {
  color: #475569;
  text-decoration: underline;
}
</style>