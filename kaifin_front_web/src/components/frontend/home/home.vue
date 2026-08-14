<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import NavBar from "../navbar/NavBar.vue";

const router = useRouter();
const searchQuery = ref("");
const isLoading = ref(true);

// 1. Top Course Data
const topCourses = ref([
  {
    id: 1,
    title: "Vue 3 Mastery",
    description: "Learn composition API with real-world projects",
    students: "400K students",
    lessons: "50 lessons",
    rating: "4.5",
    price: "$49.99",
    isPopular: false,
    image: "https://images.unsplash.com/photo-1516321318423-f06f85e504b3?w=500&auto=format&fit=crop&q=60",
    studentAvatars: [
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSW3VxjZL3_kSA2MNOF0OjZfplGwqBAhQXy7J8yrSSkMw&s=10",
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTh_k5cXIkYZ3hzJmx5T1bZ3I75UjesvsJGIrSQdhxG2g&s=10",
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRP0S6ttGS95Hyz_PrZcUpCicTSMEDWc0ygqMF342zmfw&s=10"
    ]
  },
  {
    id: 2,
    title: "UI/UX Basics",
    description: "Design clean layouts using Figma & Adobe XD",
    students: "40K students",
    lessons: "20 lessons",
    rating: "4.0",
    price: "$39.99",
    isPopular: false,
    image: "https://images.unsplash.com/photo-1522202176988-66273c2fd55f?w=500&auto=format&fit=crop&q=60",
    studentAvatars: [
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQby2dEgu1YToVDheFGglIRDCSrbbkjePalS6qJwUDEJg&s=10",
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTkYF18KgSgnBrGXMIp_-3Fh7VfKI8__LiD12p9d9FfJg&s=10",
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRWOW7kqG8TaXpCftb0Vht_sYIVGcJ9MbPqZ7Zn4gk7-g&s=10"
    ]
  },
  {
    id: 3,
    title: "Node.js API",
    description: "Build fast backends with Express & MongoDB",
    students: "30K students",
    lessons: "10 lessons",
    rating: "4.2",
    price: "$44.99",
    isPopular: false,
    image: "https://images.unsplash.com/photo-1531482615713-2afd69097998?w=500&auto=format&fit=crop&q=60",
    studentAvatars: [
      "https://api.dicebear.com/7.x/avataaars/svg?seed=node1",
      "https://api.dicebear.com/7.x/avataaars/svg?seed=node2",
      "https://api.dicebear.com/7.x/avataaars/svg?seed=node3"
    ]
  },
  {
    id: 4,
    title: "JavaScript Pro",
    description: "Advanced concepts: closures, promises, and OOP",
    students: "50K students",
    lessons: "30 lessons",
    rating: "4.8",
    price: "$59.99",
    isPopular: true,
    image: "https://images.unsplash.com/photo-1434030216411-0b793f4b4173?w=500&auto=format&fit=crop&q=60",
    studentAvatars: [
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4eqpFfPcr8emX3vrlrGgOUVgZcsL6P17FZv6IGiNa_Q&s=10",
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ8sFS-cWes5GACsxIoT-pO1m9arlRRkIf5XIQUAPo9Zw&s",
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4SPmztW6CFsWOKzJTWRS3ClZT7RWmu0JD-JSAUTYDGg&s=10"
    ]
  }
]);

// 2. Top Instructor Data
const topInstructors = ref([
  { 
    id: 1, 
    username: "Alice", 
    level: "Lvl 1", 
    students: "10k students", 
    rating: "4.8", 
    avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=Alice" 
  },
  { 
    id: 2, 
    username: "Bob", 
    level: "Lvl 2", 
    students: "12k students", 
    rating: "4.5", 
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRkITF8ud3rnoF1JleTe8zBZh1Fz2T3_3l0RQyAWrWDqg&s=10" 
  },
  { 
    id: 3, 
    username: "Charlie", 
    level: "Lvl 1", 
    students: "9k students", 
    rating: "4.2", 
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS9eBo6yUWTzp_8hRfVqs2oQNEBzHnQqkcsTTVAsZnhyA&s=10" 
  },
  { 
    id: 4, 
    username: "Diana", 
    level: "Lvl 3", 
    students: "6k students", 
    rating: "4.9", 
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTIRO0LuYaRlwSJHSJQjVSSAL_EZQO9hju9mUm3wtLSFA&s=10" 
  }
]);

// 3. Top User Data
const topUsers = ref([
  { 
    id: 1, 
    username: "Evan", 
    level: "Lvl 5", 
    follower: "10M", 
    streak: "1000",
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTwRtaBBsSGeXoNNSvKOpMjJi9_In7yFCL-3q6zyagZoA&s"
  },
  { 
    id: 2, 
    username: "Fiona", 
    level: "Lvl 4", 
    follower: "16M", 
    streak: "9000",
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRg8KFf5dQ_nkTgIxZCRyeR4vBiBgci7X33PGzadC6uWQ&s=10"
  },
  { 
    id: 3, 
    username: "George", 
    level: "Lvl 6", 
    follower: "12M", 
    streak: "2000",
    avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=George"
  },
  { 
    id: 4, 
    username: "Hannah", 
    level: "Lvl 3", 
    follower: "14M", 
    streak: "3000",
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQh0q4tGldXhZ8v1R189WgOC9sZWYQw-XUqY1h6vhg8bw&s=10"
  }
]);

// 4. Sponsors Data
const sponsors = ref([
  {
    id: 1,
    title: "Wink Bank",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTXlI1mNSiydNCeS3zyw488KVjy-v_ZVsyytF6iX71E-Q&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTXlI1mNSiydNCeS3zyw488KVjy-v_ZVsyytF6iX71E-Q&s=10"
  },
  {
    id: 2,
    title: "Rupp",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyvLU7mMbM-LsYiA1C_dIvUyQb8zyr0Y7nNwiw5EO-fQ&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyvLU7mMbM-LsYiA1C_dIvUyQb8zyr0Y7nNwiw5EO-fQ&s=10"
  },
  {
    id: 3,
    title: "ACE",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVbFiigLim2Y8L3wYAXr5ZZio59rii96XEFYmEQjoBiA&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVbFiigLim2Y8L3wYAXr5ZZio59rii96XEFYmEQjoBiA&s=10"
  },
  {
    id: 4,
    title: "Chip Mong Market",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSn9NOWTme2d5Br0DDizgQAyzkfF-UfjkV72q6JxeeEsw&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSn9NOWTme2d5Br0DDizgQAyzkfF-UfjkV72q6JxeeEsw&s=10"
  },
  {
    id: 4,
    title: "Chip Mong Retail",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSKNcoZhYPHPy9T3SQGK7GDGZisXdqvy_s9A_Ie45Gfww&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSKNcoZhYPHPy9T3SQGK7GDGZisXdqvy_s9A_Ie45Gfww&s=10"
  },
  {
    id: 4,
    title: "Sastra Film",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRL2yg_MZaVkXuCIwzw9KoCMCpARkp_bXEoO0t_eFYFvg&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRL2yg_MZaVkXuCIwzw9KoCMCpARkp_bXEoO0t_eFYFvg&s=10"
  },
  {
    id: 4,
    title: "DoDo Film",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRCdILnVKJmoojtTsVPlu3k0T03vbsuszny1drVo9G1PQ&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRCdILnVKJmoojtTsVPlu3k0T03vbsuszny1drVo9G1PQ&s=10"
  },
  {
    id: 4,
    title: "Khmer Saa Film",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTg7ChEMgD-lxit-NG08tHu0z1CKjSBsPgJKUKeiSMxVw&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTg7ChEMgD-lxit-NG08tHu0z1CKjSBsPgJKUKeiSMxVw&s=10"
  },
  {
    id: 4,
    title: "Khmer Saa Film",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRPT3xHGAGHs333J04kWZJr0qXghvd_aI56CyY3atO3xg&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRPT3xHGAGHs333J04kWZJr0qXghvd_aI56CyY3atO3xg&s=10"
  },
  {
    id: 4,
    title: "Hong Meas KH",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRMPs9iCyijzUm-GssePt-_OqP7-BAOEvId9oIEy7CgOQ&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRMPs9iCyijzUm-GssePt-_OqP7-BAOEvId9oIEy7CgOQ&s=10"
  },
  {
    id: 4,
    title: "BTV KH",
    image: "https://images.seeklogo.com/logo-png/65/1/bayon-news-television-logo-png_seeklogo-654190.png",
    logo: "https://images.seeklogo.com/logo-png/65/1/bayon-news-television-logo-png_seeklogo-654190.png"
  },
  {
    id: 4,
    title: "Bayon TV",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLepv74woJ7dyjAQUTzwk3JpwD_hJdTnFxSIg757X2g&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLepv74woJ7dyjAQUTzwk3JpwD_hJdTnFxSIg757X2g&s=10"
  },
  {
    id: 4,
    title: "Chip Mong Industry",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSDG7BIKfixMXs5KfyFxOhPhUpqNjcBME1pYH1hgsSD1w&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSDG7BIKfixMXs5KfyFxOhPhUpqNjcBME1pYH1hgsSD1w&s=10"
  },
  {
    id: 4,
    title: "Khmer Saa Film",
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTg7ChEMgD-lxit-NG08tHu0z1CKjSBsPgJKUKeiSMxVw&s=10",
    logo: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTg7ChEMgD-lxit-NG08tHu0z1CKjSBsPgJKUKeiSMxVw&s=10"
  }
]);

// ============= FUNCTIONS =============
function handleSearch() {
  if (searchQuery.value.trim()) {
    alert(`🔍 Searching for: "${searchQuery.value}"`);
  } else {
    alert("Please enter a search term");
  }
}

function handleBuy(course) {
  alert(`🛒 Added "${course.title}" to cart!\n💰 Price: ${course.price}`);
}

function handleFollow(user, type) {
  alert(`✅ You are now following ${type}: ${user.username}`);
}

function handleSponsorReadMore(sponsor) {
  alert(`📖 Reading more about sponsor: ${sponsor.title}`);
}

function viewMore(section) {
  alert(`📚 Viewing all ${section}...`);
}

function openChat() {
  alert("💬 Opening Chat Room...");
}

function openMoreOptions() {
  alert("⚙️ Opening More Options...");
}

function goToProfile(item, type) {
  router.push({
    name: 'ProfileInstructor',
    params: { id: item.id },
    query: { type: type }
  });
}

function showTooltip(event, name) {
  const tooltip = document.createElement('div');
  tooltip.className = 'custom-tooltip';
  tooltip.textContent = name;
  tooltip.style.left = event.pageX + 'px';
  tooltip.style.top = (event.pageY - 30) + 'px';
  document.body.appendChild(tooltip);
  
  setTimeout(() => {
    tooltip.remove();
  }, 1500);
}

onMounted(() => {
  setTimeout(() => {
    isLoading.value = false;
  }, 800);
});
</script>

<template>
  <div class="page-container">
    <NavBar />
    
    <div class="content-wrap">
      
      <!-- MAIN LAYOUT GRID -->
      <div class="main-layout-grid">
        
        <!-- 1. MAIN CONTENT LEFT -->
        <div class="sections-container" :class="{ 'fade-in': !isLoading }">
          
          <!-- SECTION 1: TOP COURSE -->
          <div class="row-section-box" :class="{ 'slide-up': !isLoading }">
            <div class="section-header">
              <h2>
                Top Course
              </h2>
              <button class="all-btn" @click="viewMore('Top Course')">View All &rarr;</button>
            </div>

            <div v-if="isLoading" class="skeleton-wrapper">
              <div v-for="n in 3" :key="'skeleton-' + n" class="skeleton-card">
                <div class="skeleton-image"></div>
                <div class="skeleton-content">
                  <div class="skeleton-line w-80"></div>
                  <div class="skeleton-line w-60"></div>
                </div>
              </div>
            </div>

            <div v-else class="horizontal-cards-scroll">
              <div v-for="course in topCourses" :key="course.id" class="item-card course-card">
                <div v-if="course.isPopular" class="popular-badge">
                  <svg width="12" height="12" viewBox="0 0 24 24" fill="currentColor"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon></svg>
                  Best Seller
                </div>
                
                <div class="img-box">
                  <img :src="course.image" :alt="course.title" loading="lazy" />
                </div>
                
                <div class="card-content-area">
                  <div class="card-info">
                    <h3>{{ course.title }}</h3>
                    <p>{{ course.description }}</p>
                    
                    <div class="students-info-row">
                      <div class="stacked-avatars">
                        <img 
                          v-for="(avatar, index) in course.studentAvatars" 
                          :key="index" 
                          :src="avatar" 
                          alt="student" 
                          class="stack-avatar"
                          @mouseenter="showTooltip($event, 'Student ' + (index + 1))"
                        />
                        <span class="more-avatars">+25</span>
                      </div>
                      <span class="tag">
                        <svg class="inline-svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
                        {{ course.students }}
                      </span>
                    </div>

                    <div class="tags-row">
                      <span class="tag lesson-tag">
                        <svg class="inline-svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"></path><path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"></path></svg>
                        {{ course.lessons }}
                      </span>
                      <span class="star-rating">
                        <svg class="inline-svg star-svg" width="12" height="12" viewBox="0 0 24 24" fill="currentColor"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon></svg>
                        {{ course.rating }}
                      </span>
                      <span class="price-tag">
                        <svg class="inline-svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="12" y1="1" x2="12" y2="23"></line><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path></svg>
                        {{ course.price }}
                      </span>
                    </div>
                  </div>
                  
                  <button class="action-btn buying-effect" @click="handleBuy(course)">
                    <svg class="btn-svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="9" cy="21" r="1"></circle><circle cx="20" cy="21" r="1"></circle><path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"></path></svg>
                    Add to Cart
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- SECTION 2: TOP INSTRUCTOR -->
          <div class="row-section-box" :class="{ 'slide-up': !isLoading }">
            <div class="section-header">
              <h2>
                Top Instructor
              </h2>
              <button class="all-btn" @click="viewMore('Top Instructor')">View All &rarr;</button>
            </div>

            <div v-if="isLoading" class="skeleton-wrapper">
              <div v-for="n in 3" :key="'skeleton-instr-' + n" class="skeleton-card">
                <div class="skeleton-image"></div>
                <div class="skeleton-content">
                  <div class="skeleton-line w-80"></div>
                </div>
              </div>
            </div>

            <div v-else class="horizontal-cards-scroll">
              <div 
                v-for="instructor in topInstructors" 
                :key="instructor.id" 
                class="item-card instructor-vertical-card clickable-card"
                @click="goToProfile(instructor, 'Instructor')"
              >
                <div class="instructor-top-box">
                  <div class="instructor-avatar-vertical">
                    <img :src="instructor.avatar" :alt="instructor.username" loading="lazy" />
                    <div class="online-status"></div>
                  </div>
                </div>
                
                <div class="card-content-area">
                  <div class="card-info instructor-text-center">
                    <h3>{{ instructor.username }}</h3>
                    <p>{{ instructor.students }} • {{ instructor.level }}</p>
                  </div>
                  
                  <button class="action-btn follow-effect-blue" @click.stop="handleFollow(instructor, 'Instructor')">
                    Add to Follow
                  </button>
                </div>
              </div>
            </div>
          </div>
          
        </div>

        <!-- 2. SIDEBAR RIGHT PANEL -->
        <div class="sidebar-right-container" :class="{ 'fade-in': !isLoading }">
          
          <!-- SEARCH BOX CARD IN SIDEBAR -->
          <div class="sidebar-search-card">
            <div class="search-input-wrap">
              <input 
                type="text" 
                v-model="searchQuery" 
                placeholder="Search..." 
                @keyup.enter="handleSearch"
                aria-label="Search"
              />
              <button class="search-icon-btn" @click="handleSearch" aria-label="Search button">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="11" cy="11" r="8"></circle>
                  <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
                </svg>
              </button>
            </div>
          </div>

          <!-- DISCOVER FEATURES CARD IN SIDEBAR -->
          <div class="sidebar-discover-card">
            <div class="hero-text">
              <h1 class="page-title">
                Discover Features
                <span class="badge-new">New</span>
              </h1>
              <p class="page-subtitle">Explore top trending courses and elite instructors in our community.</p>
            </div>
          </div>

          <!-- TOP USERS CARD IN SIDEBAR -->
          <div class="sidebar-right-panel">
            <div class="sidebar-section-header">
              <h2 class="sidebar-title">
                Top Users
              </h2>
              <button class="all-btn" @click="viewMore('Top User')">View All</button>
            </div>

            <div class="top-users-sidebar-list">
              <div 
                v-for="user in topUsers" 
                :key="user.id" 
                class="sidebar-user-card clickable-card"
                @click="goToProfile(user, 'User')"
              >
                <div class="sidebar-user-avatar">
                  <img :src="user.avatar" :alt="user.username" loading="lazy" />
                  <div class="online-status"></div>
                </div>
                
                <div class="sidebar-user-info">
                  <h4 class="sidebar-user-name">{{ user.username }}   <div class="streak-badge-sm">
                    <span>🔥</span> {{ user.streak }}
                  </div></h4>
                  <p class="sidebar-user-sub">
  <span class="user-meta-item">
    <svg
      width="13"
      height="13"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
    >
      <path d="M12 2l1.9 5.8h6.1l-4.9 3.6 1.9 5.8-5-3.6-5 3.6 1.9-5.8-4.9-3.6h6.1L12 2z"/>
    </svg>
    Lvl: {{ user.level }}
  </span>

  <span class="meta-separator">•</span>

  <span class="user-meta-item">
    <svg
      width="13"
      height="13"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
    >
      <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
      <circle cx="9" cy="7" r="4"/>
      <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
      <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
    </svg>
    {{ user.follower }}
  </span>
</p>
                
                </div>

                <button class="user-add-btn" @click.stop="handleFollow(user, 'User')" aria-label="Follow user">
                  Follow
                </button>
              </div>
            </div>
          </div>

          <!-- SPONSORS PROFILE CARD -->
<div class="sidebar-right-panel sponsors-sidebar-card">
  <div class="sidebar-section-header">
    <h2 class="sidebar-title">
      Sponsors
    </h2>

    <button class="all-btn" @click="viewMore('Sponsors')">
      See All
    </button>
  </div>

  <div class="sponsor-profile-grid">
    <div
      v-for="sponsor in sponsors"
      :key="sponsor.id"
      class="sponsor-profile-item"
      @click="handleSponsorReadMore(sponsor)"
    >
      <!-- Avatar -->
      <div class="sponsor-avatar-wrap">
        <img
          :src="sponsor.logo || sponsor.image"
          :alt="sponsor.title"
          class="sponsor-avatar"
          loading="lazy"
        />

        <!-- Online / Active dot -->
        <span class="sponsor-active-dot"></span>
      </div>

      <!-- Name -->
      <span class="sponsor-profile-name">
        {{ sponsor.title }}
      </span>
    </div>
  </div>
</div>

        </div>

      </div>

    </div>

    <!-- FLOATING ACTIONS -->
    <div class="floating-sidebar-right">
      <div class="sidebar-card chat-card" @click="openChat">
        <span class="chat-badge">3</span>
        <div class="chat-icon-container">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"></path>
          </svg>
        </div>
        <span class="sidebar-label">Chat</span>
      </div>

      <div class="sidebar-card more-card" @click="openMoreOptions">
        <div class="more-dots-row">
          <span class="dot"></span>
          <span class="dot"></span>
          <span class="dot"></span>
        </div>
        <span class="sidebar-label">More</span>
      </div>
    </div>

  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap');

.page-container {
  min-height: 100vh;
  font-family: 'Plus Jakarta Sans', system-ui, -apple-system, sans-serif;
  color: #1f2937;

}

.content-wrap {
  max-width: 1251px;
  margin: 0 auto;
  /* padding: 32px 12px; */
  /* background-color: red; */
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes slideUp {
  from { opacity: 0; transform: translateY(40px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

.fade-in {
  animation: fadeIn 0.6s ease forwards;
}

.slide-up {
  animation: slideUp 0.5s ease forwards;
}

.main-layout-grid {
  width: 1251px;
  max-width: 100%;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 320px;
  gap: 16px;
  position: relative;
  align-items: start;
  justify-content: center;
}

/* Sidebar Discover Features Card Style */
.sidebar-discover-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 20px;
}

.hero-text {
  display: flex;
  flex-direction: column;
  flex: 1;
}

.page-title {
  font-size: 18px;
  font-weight: 800;
  margin: 0 0 6px 0;
  color: #111827;
  letter-spacing: -0.5px;
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.title-sparkle-icon {
  color: #f59e0b;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.badge-new {
  background: linear-gradient(135deg, #f59e0b, #f97316);
  color: white;
  font-size: 9px;
  font-weight: 700;
  padding: 2px 8px;
  border-radius: 20px;
  letter-spacing: 0.5px;
  text-transform: uppercase;
}

.page-subtitle {
  font-size: 12px;
  color: #4b5563;
  margin: 0;
  line-height: 1.4;
}

/* Sidebar Right Container */
.sidebar-right-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
  position: sticky;
  top: 24px;
  margin-top: 12px;
}

.sidebar-search-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 16px;
}

.search-input-wrap {
  position: relative;
  display: flex;
  align-items: center;
  width: 100%;
  background: transparent;
  border: 1px solid #e2e8f0;
  border-radius: 32px;
  /* transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1); */
}

.search-input-wrap:focus-within {
  border-color: #1B75D2;

}

.search-input-wrap input {
  width: 100%;
  padding: 10px 16px;
  border: none;
  background: transparent;
  font-size: 13px;
  color: #1f2937;
  outline: none;
  font-family: inherit;
}

.search-input-wrap input::placeholder {
  color: #9ca3af;
}

.search-icon-btn {
  background: transparent;
  border: none;
  padding: 0 14px;
  color: #6b7280;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.search-icon-btn:hover {
  color: #1B75D2;
  transform: scale(1.1);
}

/* Sidebar Right Panel for Top Users & Sponsors */
.sidebar-right-panel {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.sidebar-section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.sidebar-title {
  font-size: 18px;
  font-weight: 700;
  color: #111827;
  margin: 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.top-users-sidebar-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.sidebar-user-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 4px;
  border-radius: 12px;
  transition: all 0.2s ease;
  /* background: #f8fafc; */
  /* border: 1px solid #f1f5f9; */
}

.sidebar-user-card:hover {
  background: #0000000e;
  border-radius: 32px;
  cursor: cell;
}

.sidebar-user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50px;
  overflow: hidden;
  position: relative;
  flex-shrink: 0;
  background: #e2e8f0;
}

.sidebar-user-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.sidebar-user-info {
  flex: 1;
  min-width: 0;
}

.sidebar-user-name {
  font-size: 12px;
  font-weight: 700;
  color: #111827;
  margin: 0 0 2px 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.sidebar-user-sub {
  display: flex;
  align-items: center;
  gap: 6px;

  margin: 4px 0 0;

  font-size: 11.5px;
  color: #8a94a6;
}

.user-meta-item {
  display: inline-flex;
  align-items: center;
  gap: 3px;
  white-space: nowrap;
}

.user-meta-item svg {
  flex-shrink: 0;
  color: #8a94a6;
}

.meta-separator {
  color: #c5ccd6;
  font-size: 10px;
}

.streak-badge-sm {
  display: inline-flex;
  align-items: center;
  gap: 2px;
  font-size: 10px;
  font-weight: 700;
  color: #ffffff;
  background: linear-gradient(135deg, #f59e0b, #f97316);
  padding: 1px 6px;
  border-radius: 10px;
  font-style: italic;
}

.user-add-btn {
  background: #1B75D2;
  color: #ffffff;
  border: none;
  /* width: 32px;
  height: 32px; */
  border-radius: 32px;
  padding: 4px 8px;
  font-size: 12px;
  font-weight: 400;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  flex-shrink: 0;
}

.user-add-btn:hover {
  transform: scale(1.1);
  background: #155bb5;
}

/* Sponsors Sidebar Card Specific Styles */
.sponsors-sidebar-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.sponsor-item-card {
  position: relative;
  height: 100px;
  border-radius: 12px;
  overflow: hidden;
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

.sponsor-item-card:hover {
  border-color: #1B75D2;
}

.sponsor-img-bg {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
}

.sponsor-img-bg img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.35;
  transition: transform 0.5s ease;
}

.sponsor-item-card:hover .sponsor-img-bg img {
  transform: scale(1.05);
}

.sponsor-content-overlay {
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 14px;
  height: 100%;
  background: linear-gradient(90deg, rgba(255,255,255,0.92) 0%, rgba(255,255,255,0.7) 100%);
}

.sponsor-info-text h4 {
  font-size: 15px;
  font-weight: 700;
  color: #111827;
  margin: 0 0 8px 0;
}

.sponsor-read-btn {
  background: #ffffff;
  border: 1px solid #cbd5e1;
  color: #1f2937;
  font-size: 11px;
  font-weight: 700;
  padding: 4px 10px;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.2s;
}

.sponsor-read-btn:hover {
  background: #1B75D2;
  color: #ffffff;
  border-color: #1B75D2;
}

.sponsor-logo-badge {
  width: 44px;
  height: 44px;
  border-radius: 10px;
  overflow: hidden;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  box-shadow: 0 2px 6px rgba(0,0,0,0.06);
  flex-shrink: 0;
}

.sponsor-logo-badge img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Common Layout Components */
.sections-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 12px;
  min-width: 0;
}

.floating-sidebar-right {
  display: flex;
  flex-direction: column;
  gap: 10px;
  position: fixed;
  right: 20px;
  bottom: 20px;
  width: 70px;
  z-index: 100;
}

.sidebar-card {
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  padding: 8px 6px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.25s ease;
}

.more-card {
  background-color: #1B75D2;
}

.chat-card {
  background-color: #EE8207;
  position: relative;
  border: 2.1px solid #1B75D2;
  transform: rotate(-5deg);
}

.chat-card:hover {
  transform: rotate(0deg) translateY(-3px);
}

.chat-badge {
  position: absolute;
  top: -8px;
  right: -14px;
  min-width: 26px;
  height: 26px;
  border-radius: 50%;
  background: #EE8207;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 700;
  border: 2px solid #fff;
  animation: badgePulse 1.5s infinite;
}

@keyframes badgePulse {
  0%, 100% { transform: scale(1) rotate(0deg); }
  20% { transform: scale(1.2) rotate(-10deg); }
  40% { transform: scale(1.25) rotate(10deg); }
  60% { transform: scale(1.1) rotate(-5deg); }
}

.sidebar-card:hover {
  border-color: #1B75D2;
  opacity: 0.8;
  transform: translateY(-3px);
  box-shadow: 0 6px 15px rgba(27, 117, 210, 0.15);
}

.chat-icon-container {
  color: #1B75D2;
  margin-bottom: 6px;
  background: #eff6ff;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid #dbeafe;
}

.more-dots-row {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 3px;
  background: #f1f5f9;
  border: 1px solid #e2e8f0;
  padding: 8px 10px;
  border-radius: 9999px;
  margin-bottom: 6px;
}

.dot {
  width: 4px;
  height: 4px;
  background: #475569;
  border-radius: 50%;
}

.sidebar-label {
  color: #ffffff;
  font-size: 11px;
  font-weight: 700;
  text-align: center;
}

.row-section-box {
  background: #ffffff;
  border-left: 1px solid #e2e8f0;
  border-right: 1px solid #e2e8f0;
  /* border: 1px solid #e2e8f0; */
  /* border-radius: 12px; */
  padding: 16px 20px;
  transition: all 0.3s ease;
}

.row-section-box:hover {
  border-color: #cbd5e1;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.section-header h2 {
  font-size: 20px;
  font-weight: 700;
  margin: 0;
  color: #111827;
  display: flex;
  align-items: center;
  gap: 10px;
}

.section-count {
  background: hsl(226, 100%, 97%);
  color: #1B75D2;
  font-size: 12px;
  font-weight: 600;
  padding: 2px 10px;
  border-radius: 20px;
}

.section-icon {
  color: #1B75D2;
  flex-shrink: 0;
}

.all-btn {
  background: transparent;
  border: none;
  color: #1B75D2;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
  padding: 6px 12px;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.all-btn:hover {
  background: #eff6ff;
  transform: translateX(-4px);
}

/* Updated to Grid to remove Scroll X */
.horizontal-cards-scroll {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 20px;
  padding: 4px 2px 8px 2px;
  overflow-x: visible;
}

.item-card {
  width: 100%;
  flex: unset;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}

.item-card:hover {
  border-color: #1B75D2;
  transform: translateY(-6px);
}

.popular-badge {
  position: absolute;
  top: 12px;
  right: 12px;
  background: linear-gradient(135deg, #f59e0b, #f97316);
  color: white;
  font-size: 10px;
  font-weight: 700;
  padding: 4px 12px;
  border-radius: 20px;
  z-index: 10;
  letter-spacing: 0.5px;
  animation: pulse 2s infinite;
  display: flex;
  align-items: center;
  gap: 4px;
}

.img-box {
  width: 100%;
  height: 160px;
  background: #e2e8f0;
  overflow: hidden;
  position: relative;
}

.img-box img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.5s ease;
}

.item-card:hover .img-box img {
  transform: scale(1.05);
}

.online-status {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 10px;
  height: 10px;
  background: #22c55e;
  border-radius: 50%;
  border: 2px solid white;
}

.card-content-area {
  padding: 14px 16px 16px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex: 1;
}

.card-info h3 {
  font-size: 16px;
  font-weight: 700;
  margin: 0 0 4px 0;
  color: #111827;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-info p {
  font-size: 13px;
  color: #6b7280;
  margin: 0 0 12px 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.students-info-row {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 10px;
  margin-bottom: 10px;
}

.stacked-avatars {
  display: flex;
  align-items: center;
}

.stack-avatar {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #ffffff;
  margin-left: -10px;
  background: #e2e8f0;
  cursor: pointer;
  transition: all 0.3s ease;
}

.stack-avatar:first-child {
  margin-left: 0;
}

.stack-avatar:hover {
  transform: translateY(-4px) scale(1.1);
  border-color: #1B75D2;
  z-index: 10;
}

.more-avatars {
  font-size: 10px;
  font-weight: 600;
  color: #6b7280;
  margin-left: 4px;
}

.tags-row {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  margin-bottom: 12px;
  align-items: center;
  gap: 6px;
  flex-wrap: wrap;
}

.tag {
  background: #f1f5f9;
  color: #475569;
  padding: 3px 10px;
  border-radius: 6px;
  font-weight: 500;
  font-size: 11px;
  white-space: nowrap;
  display: inline-flex;
  align-items: center;
  gap: 4px;
}

.lesson-tag {
  background: #1B75D2;
  color: #fff;
  font-weight: 600;
}

.star-rating {
  color: #f59e0b;
  font-weight: 700;
  display: inline-flex;
  align-items: center;
  gap: 3px;
}

.star-svg {
  fill: #f59e0b;
}

.price-tag {
  background: #03CF60;
  color: #ffffff;
  font-weight: 700;
  padding: 3px 10px;
  border-radius: 6px;
  font-size: 11px;
  display: inline-flex;
  align-items: center;
  gap: 3px;
}

.action-btn {
  width: 100%;
  border: none;
  padding: 10px 14px;
  border-radius: 32px;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  letter-spacing: 0.3px;
}

.buying-effect {
  background: linear-gradient(135deg, #1B75D2, #1B75D2);
  color: #ffffff;
  margin-top: 6px;
}

.buying-effect:hover {
  transform: translateY(-2px);
}

.follow-effect-blue {
  background: linear-gradient(135deg, #1B75D2, #1B75D2);
  color: #ffffff;
  padding: 8px 12px;
  font-size: 12px;
  border-radius: 32px;
}

.follow-effect-blue:hover {
  transform: translateY(-2px);
}

/* Updated Skeleton Wrapper to Grid */
.skeleton-wrapper {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 20px;
  padding: 8px 4px 12px 4px;
  overflow-x: visible;
}

.skeleton-card {
  width: 100%;
  flex: unset;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  overflow: hidden;
}

.skeleton-image {
  width: 100%;
  height: 160px;
  background: #e2e8f0;
}

.skeleton-content {
  padding: 12px;
}

.skeleton-line {
  height: 12px;
  background: #e2e8f0;
  border-radius: 6px;
  margin-bottom: 8px;
}

.instructor-vertical-card {
  flex: unset;
}

.instructor-top-box {
  padding: 20px 16px 0 16px;
  display: flex;
  justify-content: center;
}

.instructor-avatar-vertical {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  overflow: hidden;
  background: #e2e8f0;
  position: relative;
  border: 2px solid #f1f5f9;
}

.instructor-avatar-vertical img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.instructor-text-center {
  text-align: center;
  margin-bottom: 12px;
}


/* =========================================
   SPONSOR PROFILE GRID
========================================= */

.sponsors-sidebar-card {
  width: 100%;
  background: #fff;
  border: 1px solid #e8edf3;
  border-radius: 12px;
  padding: 16px;
  box-sizing: border-box;
}

/* Header */
.sponsors-sidebar-card .sidebar-section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 4px;
}

.sponsors-sidebar-card .sidebar-title {
  display: flex;
  align-items: center;
  gap: 6px;

  margin: 0;
  font-size: 15px;
  font-weight: 800;
  color: #172033;
}

/* =========================================
   GRID
========================================= */

.sponsor-profile-grid {
  display: grid;

  /* 4 profiles per row in sidebar */
  grid-template-columns: repeat(7, 1fr);

  gap: 16px 10px;

  width: 100%;
}

/* =========================================
   PROFILE ITEM
========================================= */

.sponsor-profile-item {
  display: flex;
  flex-direction: column;
  align-items: center;

  min-width: 0;

  cursor: pointer;

  transition:
    transform 0.2s ease,
    opacity 0.2s ease;
}

/* =========================================
   AVATAR
========================================= */

.sponsor-avatar-wrap {
  position: relative;

  width: 32px;
  height: 32px;

  flex-shrink: 0;
}

/* Circle image */
.sponsor-avatar {
  width: 32px;
  height: 32px;

  display: block;

  object-fit: cover;

  border-radius: 8px;

  border: 2px solid #fff;

  box-shadow:
    0 0 0 1px #dfe5ec,
    0 3px 8px rgba(0, 0, 0, 0.08);

  background: #f1f4f8;

  transition:
    transform 0.2s ease,
    box-shadow 0.2s ease;
}

.sponsor-profile-item:hover .sponsor-avatar {
  transform: scale(1.08);

  box-shadow:
    0 0 0 2px #1b75d2,
    0 4px 12px rgba(27, 117, 210, 0.18);
}

/* =========================================
   ACTIVE DOT
========================================= */

.sponsor-active-dot {
  position: absolute;

  right: 1px;
  bottom: 1px;

  width: 10px;
  height: 10px;

  border-radius: 50%;

  background: #22c55e;

  border: 2px solid #fff;

  box-sizing: border-box;
}

/* =========================================
   NAME
========================================= */

.sponsor-profile-name {
  width: 100%;

  margin-top: 7px;

  font-size: 10.5px;
  font-weight: 600;

  color: #596579;

  text-align: center;

  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;

  line-height: 1.3;
}

/* Hover name */
.sponsor-profile-item:hover .sponsor-profile-name {
  color: #1b75d2;
}
</style>