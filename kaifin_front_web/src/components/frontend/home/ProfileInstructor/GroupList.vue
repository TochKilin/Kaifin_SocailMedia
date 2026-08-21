<script setup>
import { ref } from 'vue';
import Star from "@/assets/images_course/idea.svg"

const groups = ref([
  {
    id: 1,
    name: 'Banking And Financial',
    members: '40K member',
    activity: '6 post a day',
    friendsInfo: '3 friends are member',
    joined: false,
    image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTzp95DKUCiLga_EeNYn0-DnvaZXhglhd8Z6tMvr0r3KjpjPcZhoWWqjlk&s=10',
   badgeImage: Star, 
    friendAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSTuNFXXEBno2JzMPycH9q7PzKk6MefqwMYkDdDXSXdxA&s=10',
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa7Dn53kJS46GLwFB9iEj8mnNgUEvBSM1_E2__EJM7wQ&s=10',
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRIECpusF1z5qUjMS8QH0QHby_-5ToKx3sLs4kHss9O8xW-ufxxEexBQaPc&s=10'
    ]
  },
  {
    id: 2,
    name: 'UI/UX Designers Hub',
    members: '400K member',
    activity: '5 post a month',
    friendsInfo: '12 friends are member',
    joined: false,
    image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa69mMLBGGUmjGokgD5N4f6K3aQV5-Ssq9SSIh8AEvhA&s',
    badgeImage: Star,
    friendAvatars: [
      'https://www.vodkhmer.news/wp-content/uploads/2022/07/%E1%9E%98%E1%9F%89%E1%9E%B6%E1%9E%9A%E1%9F%89%E1%9E%B6-500x322.jpg',
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQe5wrdpYOC5gnvHuJP7LH4d_91WGTpVVw9Gr_5O2oA_Q&s=10'
    ]
  },
  {
    id: 3,
    name: 'Fullstack Developers',
    members: '40K member',
    activity: '10+ post a day',
    friendsInfo: '5 friends are member',
    joined: false,
    image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRl8s14p-FhHkepqpMrcIVF3BAMe4ZqguqbjxFkPuGpmQ&s=10',
    badgeImage: Star,
    friendAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ7aTWpOkPUwqAh7SslvXazN_-jiTCuzmLDdn4BSl9nWA&s=10',
      'https://images.unsplash.com/photo-1522075469751-3a6694fb2f61?w=100&auto=format&fit=crop&q=60'
    ]
  },
  {
    id: 4,
    name: 'Vue.js Cambodia',
    members: '4K member',
    activity: '4 post a day',
    friendsInfo: '8 friends are member',
    joined: false,
    image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSbWYtsW2H_EK309Mh71aUCENIP2CUWCQXxD-tqgaKYLg&s=10',
    badgeImage: Star,
    friendAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROsz-7aXNP4xqY8qs-aU0kNiMD1rEzDDxEAzSXnrBwQ8ni4yRFqpoVsmM&s=10',
      'https://images.unsplash.com/photo-1539571696357-5a69c17a67c6?w=100&auto=format&fit=crop&q=60'
    ]
  },
  {
    id: 5,
    name: 'Startup Founders',
    members: '20K member',
    activity: '8 post a month',
    friendsInfo: '2 friends are member',
    joined: false,
    image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS9nZMi3VcrRTjY04KoEzfALnYYUOLGJOaeYbhE3xSJBg&s=10',
    badgeImage: Star,
    friendAvatars: [
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROsz-7aXNP4xqY8qs-aU0kNiMD1rEzDDxEAzSXnrBwQ8ni4yRFqpoVsmM&s=10',
      'https://images.unsplash.com/photo-1501196354995-cbb51c65aaea?w=100&auto=format&fit=crop&q=60'
    ]
  }
]);

const handleJoin = (group) => {
  group.joined = !group.joined;
};
</script>

<template>
  <div class="groups-container">
    <div class="groups-header">
      <h3 class="section-title">Discover Groups</h3>
      <p class="section-subtitle">Find and join groups that match your interests</p>
    </div>
    <div class="groups-grid">
      <div v-for="group in groups" :key="group.id" class="group-card">
        <div class="badge-icon">
          <img :src="group.badgeImage" alt="Badge" class="badge-img" />
        </div>
        <div class="group-banner">
          <img :src="group.image" :alt="group.name" class="banner-img" />
        </div>
        <div class="group-content">
          <h4 class="group-name">{{ group.name }}</h4>

          <div class="group-meta-row">
            <span class="meta-badge">{{ group.members }}</span>
            <span class="meta-badge">{{ group.activity }}</span>
          </div>

          <div class="friends-row">
            <div class="avatar-group">
              <img 
                v-for="(avatar, index) in group.friendAvatars" 
                :key="index" 
                :src="avatar" 
                alt="Friend Avatar" 
                class="mini-avatar" 
              />
            </div>
            <span class="friends-text">{{ group.friendsInfo }}</span>
          </div>
          <button 
            class="join-group-btn" 
            :class="{ 'joined': group.joined }"
            type="button" 
            @click="handleJoin(group)"
          >
            {{ group.joined ? 'Joined' : 'Join Group' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700&display=swap');

.groups-container {
  font-family: 'Plus Jakarta Sans', sans-serif;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 16px;
  box-sizing: border-box;
}

.groups-header {
  display: flex;
  flex-direction: column;
  gap: 2px;
  border-bottom: 1px solid #e2e8f0;
  padding-bottom: 10px;
}

.section-title {
  font-size: 16px;
  font-weight: 700;
  color: #1e293b;
  margin: 0;
}

.section-subtitle {
  font-size: 12px;
  color: #64748b;
  margin: 0;
  font-weight: 500;
}

.groups-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
  gap: 20px;
}

.group-card {
  border: 1px solid #cbd5e1;
  border-radius: 16px;
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  position: relative;
  box-shadow: 0 4px 12px rgba(227, 7, 7, 0.03);
}

.group-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 10px 20px -5px rgba(0, 0, 0, 0.08);
}

.group-banner {
  height: 110px;
  position: relative;
  border-top-left-radius: 16px;
  border-top-right-radius: 16px;
  overflow: hidden;
}

.banner-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.badge-icon {
  position: absolute;
  top: -8px;
  right: -8px;
  width: 28px;
  height: 28px;
  background-color: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 50%;
  z-index: 20;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.badge-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.group-content {
  padding: 14px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.group-name {
  font-size: 15px;
  font-weight: 700;
  color: #1e293b;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.group-meta-row {
  display: flex;
  gap: 8px;
}

.meta-badge {
  font-size: 11px;
  font-weight: 600;
  color: #475569;
  background-color: #f1f5f9;
  border: 1px solid #cbd5e1;
  padding: 4px 8px;
  border-radius: 8px;
}

.friends-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.avatar-group {
  display: flex;
  align-items: center;
}

.mini-avatar {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #ffffff;
  margin-left: -8px;
}

.mini-avatar:first-child {
  margin-left: 0;
}

.friends-text {
  font-size: 11px;
  color: #64748b;
  font-weight: 500;
}

.join-group-btn {
  background-color: #1C75D2;
  border: none;
  color: #ffffff;
  padding: 10px;
  border-radius: 12px;
  font-weight: 700;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.2s ease;
  text-align: center;
  margin-top: 4px;
}

.join-group-btn:hover {
  opacity: 0.9;
}

.join-group-btn.joined {
  background-color: #f1f5f9;
  color: #64748b;
  border: 1px solid #cbd5e1;
}
</style>