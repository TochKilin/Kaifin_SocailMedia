<script setup lang="ts">
import { ref } from 'vue';
import SidebarRightFeed from '../sidebar_right_feed/SidebarRightFeed.vue';
import SidebarLeftFeed from '../sidebar_left_feed/SidebarLeftFeed.vue';
import SidebarContainerFeed from '../sidebar_container_feed/SidebarContainerFeed.vue';
import Posts from '../posts/Posts.vue';
import PostsCard from '../postcards/PostsCard.vue';
import Navbar from '../navbar/NavBar.vue';
import GroupFeed from '../../../components/frontend/sidebar_left_feed/GoupeFeed.vue';
import FeedFollowing from '../sidebar_left_feed/FeedFollowing.vue';

const currentView = ref('feed');
const handleOpenGroupJoin = () => {
  currentView.value = 'group-join';
};

const handleBackToFeed = () => {
  currentView.value = 'feed';
};

const handleOpenFollowing = () => {
  currentView.value = 'following';
};

const handleOpenPopulars = () => {
  currentView.value = 'populars';
}


</script>

<template>
    <div class="main-container">
       <Navbar/>
    <div class="container">
    <div class="row">
      <aside class="col-12 col-md-3 col-lg-2.5 mb-3">
        <SidebarLeftFeed 
         :active-view="currentView"
          @open-group-join="handleOpenGroupJoin" 
          @open-following="handleOpenFollowing"
          @open-news-feed="handleBackToFeed" 
          @open-populars="handleOpenPopulars"
        />
      </aside>
      <main class="col-12 col-md-6 col-lg-6.5 mb-3 bd-main">
        <template v-if="currentView === 'group-join'">
            <GroupFeed />
        </template>
        <template v-else-if="currentView === 'following'">
            <FeedFollowing />
        </template>
        <template v-else-if="currentView === 'populars'">
            <PostsCard mode="popular" />
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