<script setup>
import { ref } from 'vue'

const emit = defineEmits(['open-group-join'])
</script>

<template>
<div class="left-search-contianer">
  <div class="search-box">
    <!-- SVG Search Icon -->
    <svg viewBox="0 0 24 24" class="search-svg-icon">
      <circle cx="11" cy="11" r="8" fill="none" stroke="currentColor" stroke-width="2"/>
      <path d="M21 21l-4.35-4.35" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
    </svg>

    <input
      type="search"
      placeholder="Search your friends..."
    />
  </div>
</div>
    <div class="left_container_feed">
       <div>
        <div class="left-btn font-btn btn-feed-color btn-icon">
          <span class="bg-left-side">
            <!-- SVG Feed Icon -->
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 11h9M4 16h5M4 6h16" stroke-linecap="round"/>
            </svg>
          </span>
          New Feeds
        </div>
        <div class="left-btn font-btn btn-icon left-btn-hover">
          <span class="bg-left-side">
            <!-- SVG Populars Icon (Trending/Flame) -->
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 2c0 4-4 7-4 10a4 4 0 0 0 8 0c0-3-4-6-4-10z" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M12 18v2" stroke-linecap="round"/>
            </svg>
          </span>
          Populars
        </div>
        <div class="left-btn font-btn btn-icon left-btn-hover">
          <span class="bg-left-side">
            <!-- SVG Following Icon (User Check) -->
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" stroke-linecap="round"/>
              <circle cx="9" cy="7" r="4"/>
              <polyline points="16 11 18 13 22 9" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </span>
          Following
        </div>
        <div class="left-btn font-btn btn-icon left-btn-hover">
          <span class="bg-left-side">
            <!-- SVG My Group Icon (Users) -->
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
            </svg>
          </span>
          My Group
        </div>
        <div class="btn-two-group font-btn">
          <div class="btn-sub btn-small">Group Name</div>
          <div class="btn-sub btn-small">Group Invite</div>
        </div>
        <div class="left-btn font-btn btn-icon left-btn-hover" @click="$emit('open-group-join')">
          <span class="bg-left-side">
            <!-- SVG Group Join Icon (Plus User) -->
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
  /* រក្សា Style ដើមរបស់អ្នក និងបន្ថែម CSS សម្រាប់ SVG ថ្មី */
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
    color: #000; /* ពណ៌ icon ក្នុងប៊ូតុង */
  }

  /* ពេលប៊ូតុង New Feeds មានពណ៌ខៀវ (btn-feed-color) ឱ្យ Icon ទៅជាពណ៌ស */
  .btn-feed-color .btn-icon svg,
  .btn-feed-color .bg-left-side svg {
    color: #000;
  }

  .font-btn{
    font-family: 'Nunito', sans-serif;
    /* font-weight: 700; */
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

  .btn-two-group {
    display: flex;
    flex-direction: column;
    align-items: self-start;  
    gap: 8px;              
    width: 100%;
    padding: 7px 12px;
    /* border: 1px solid #E5E7EB; */
    border-radius: 12px;
  }

  .btn-small {
  width: fit-content;   
  /* min-width: 200px;        */
  padding: 7px 12px;
    background-color: #1976D2;
  border-left:8px solid #1DD25E;
  color: #ffffff;
  border-top-right-radius: 32px;
  border-bottom-right-radius: 32px;
  /* border-radius: 32px; */
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
</style>