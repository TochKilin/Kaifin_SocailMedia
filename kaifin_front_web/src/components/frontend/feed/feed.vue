<script setup lang="ts">
import { ref } from 'vue';
import SidebarRightFeed from '../sidebar_right_feed/SidebarRightFeed.vue';
import SidebarLeftFeed from '../sidebar_left_feed/SidebarLeftFeed.vue';
import SidebarContainerFeed from '../sidebar_container_feed/SidebarContainerFeed.vue';
import Posts from '../posts/Posts.vue';
import PostsCard from '../postcards/PostsCard.vue';
import Navbar from '../navbar/NavBar.vue';
import GroupFeed from '../../../components/frontend/sidebar_left_feed/GoupeFeed.vue';

// State សម្រាប់គ្រប់គ្រងការបង្ហាញ View ('feed' សម្រាប់ទំព័រដើម, 'group-join' សម្រាប់ទំព័រក្រុម)
const currentView = ref('feed');

// មុខងារពេលចុចលើប៊ូតុង Group Join ពី Sidebar
const handleOpenGroupJoin = () => {
  currentView.value = 'group-join';
};

// មុខងារពេលចុចលើ New Feeds ដើម្បីត្រឡប់មក Feed ដើមវិញ
const handleBackToFeed = () => {
  currentView.value = 'feed';
};
</script>

<template>
    <div class="main-container">
       <Navbar/>
    <div class="container">
    <div class="row">
      <!-- ភ្ជាប់ Event ទាំងពីរនៅទីនេះ -->
      <aside class="col-12 col-md-3 col-lg-2.5 mb-3">
        <SidebarLeftFeed 
          @open-group-join="handleOpenGroupJoin" 
          @open-news-feed="handleBackToFeed" 
        />
      </aside>

      <main class="col-12 col-md-6 col-lg-6.5 mb-3 bd-main">
        <!-- ប្តូរមាតិកាកណ្តាលឱ្យដាច់ពីគ្នាផ្អែកលើ currentView -->
        <template v-if="currentView === 'group-join'">
            <GroupFeed />
        </template>
        <template v-else>
            <SidebarContainerFeed />
            <!-- <Posts /> -->
            <PostsCard />
        </template>
      </main>

      <aside class="col-12 col-md-3 col-lg-3 mb-3">
        <SidebarRightFeed />
      </aside>
    </div>
    <div class="row">
        <aside class="col-12 col-md-3 col-lg-2.5 mb-3">
        </aside>
    </div>
  </div>
  <div>
  </div>
    </div>
</template>

<style scoped>
body{
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
}

.container{
    width: 85%;
    margin: auto;
}

.bd-main{
  /* border: 0.5px solid #EFE2D3; */
  /* background-color: #FDFDFD; */
  /* padding: 0; */
}

.row > aside {
  position: sticky;
  top: 60px;
  align-self: flex-start;
  max-height: calc(100vh - 100px);
  
  overflow-y: auto;
  overflow-x: hidden;
  scrollbar-width: none;
  box-sizing: border-box;
}

.row > aside * {
  box-sizing: border-box;
}

.row > aside::-webkit-scrollbar {
  width: 5px;
}
.row > aside::-webkit-scrollbar-thumb {
  background: #E5E7EB;
  border-radius: 10px;
}
</style>