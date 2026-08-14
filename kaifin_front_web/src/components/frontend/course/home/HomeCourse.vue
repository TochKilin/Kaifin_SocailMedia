<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import NavBar from "../../navbar/NavBar.vue";

const selectedCourseCategory = ref("ALL");
const selectedCourseType = ref("ALL");
const selectedLevel = ref("ALL");
const searchQuery = ref("");
const router = useRouter();

const courseCategories = [
  {
    label: "ALL",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="7" height="7" rx="1.5"/><rect x="14" y="3" width="7" height="7" rx="1.5"/><rect x="14" y="14" width="7" height="7" rx="1.5"/><rect x="3" y="14" width="7" height="7" rx="1.5"/></svg>`
  },
  {
    label: "Development",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/></svg>`
  },
  {
    label: "Design",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h11a2 2 0 012 2v13a4 4 0 01-4 4H7z"/><path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10"/><circle cx="9.5" cy="9.5" r="1.5"/></svg>`
  },
  {
    label: "Business",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="2" y="7" width="20" height="14" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 21V5a2 2 0 00-2-2h-4a2 2 0 00-2 2v16"/></svg>`
  }
];

const courseTypes = [
  {
    label: "ALL",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="7" height="7" rx="1.5"/><rect x="14" y="3" width="7" height="7" rx="1.5"/><rect x="14" y="14" width="7" height="7" rx="1.5"/><rect x="3" y="14" width="7" height="7" rx="1.5"/></svg>`
  },
  {
    label: "Paid",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="9"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 6v12M15 9.5c0-1.4-1.34-2.5-3-2.5s-3 1.1-3 2.5 1.34 2.2 3 2.5c1.66.3 3 1.1 3 2.5s-1.34 2.5-3 2.5-3-1.1-3-2.5"/></svg>`
  },
  {
    label: "Free",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M20 12v8H4v-8M22 7H2v5h20V7zM12 7v13M12 7c-1.5-3-4-5-5.5-3.5S7 7 12 7zM12 7c1.5-3 4-5 5.5-3.5S17 7 12 7z"/></svg>`
  }
];

const courseLevels = [
  {
    label: "ALL",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="7" height="7" rx="1.5"/><rect x="14" y="3" width="7" height="7" rx="1.5"/><rect x="14" y="14" width="7" height="7" rx="1.5"/><rect x="3" y="14" width="7" height="7" rx="1.5"/></svg>`
  },
  {
    label: "Beginner",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 18h4v-4H3v4zM10 18h4v-9h-4v9zM17 18h4V6h-4v12z" opacity="0.35"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 18h4v-4H3v4z"/></svg>`
  },
  {
    label: "Intermediate",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M17 18h4V6h-4v12z" opacity="0.35"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 18h4v-4H3v4zM10 18h4v-9h-4v9z"/></svg>`
  },
  {
    label: "Expert",
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 18h4v-4H3v4zM10 18h4v-9h-4v9zM17 18h4V6h-4v12z"/></svg>`
  }
];

const popularCourses = ref([
  {
    id: 1,
    tag: "New",
    tagType: "badge-new",
    contentType: "video",
    title: "Vue 3 Masterclass",
    thumbnail: "https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=500&auto=format&fit=crop&q=60",
    description: "Learn composition API with real projects",
    instructor: "John Doe",
    instructorAvatar: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&auto=format&fit=crop&q=60",
    level: "Level 10",
    rating: "4.8",
    subscribers: "10k",
    price: 30,
    originalPrice: 60,
    promoText: "Special discount ends soon",
    isFree: false
  },
  {
    id: 2,
    tag: "New",
    tagType: "badge-new",
    contentType: "article",
    title: "Advanced Node.js",
    thumbnail: "https://images.unsplash.com/photo-1627398242454-45a1465c2479?w=500&auto=format&fit=crop&q=60",
    description: "Build robust scalable backend APIs",
    instructor: "Sarah Jenkins",
    instructorAvatar: "https://media.licdn.com/dms/image/v2/D5622AQGo0R4LOKLEeg/feedshare-shrink_800/B56ZoOgcCeKAAg-/0/1761179985841?e=2147483647&v=beta&t=nfQ2OyeEGnjyvoNkIiohdBFQt1pin0HMMSUBFv4n4u4",
    level: "Level 10",
    rating: "4.9",
    subscribers: "10k",
    price: 50,
    originalPrice: 90,
    promoText: "Limited time offer available",
    isFree: false
  },
  {
    id: 3,
    tag: "New",
    tagType: "badge-new",
    contentType: "ebook",
    title: "UI/UX Figma Pro",
    thumbnail: "https://images.unsplash.com/photo-1581291518633-83b4ebd1d83e?w=500&auto=format&fit=crop&q=60",
    description: "Design clean user interfaces easily",
    instructor: "Michael V.",
    instructorAvatar: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=100&auto=format&fit=crop&q=60",
    level: "Level 10",
    rating: "4.7",
    subscribers: "10k",
    price: 20,
    originalPrice: 50,
    promoText: "Flash sale 50% discount",
    isFree: false
  },
  {
    id: 4,
    tag: "Free",
    tagType: "badge-free",
    contentType: "video",
    title: "Python Essentials",
    thumbnail: "https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5?w=500&auto=format&fit=crop&q=60",
    description: "Introduction to programming logic",
    instructor: "David Miller",
    instructorAvatar: "https://cdn.khmercoder.com/profiles/4ca64d40-8d91-459f-b327-d7e069191f41.webp",
    level: "Level 10",
    rating: "4.6",
    subscribers: "10k",
    totalLessons: 30,
    promoText: "Lifetime free access included",
    isFree: true
  },
  {
    id: 5,
    tag: "Free",
    tagType: "badge-free",
    contentType: "ebook",
    title: "Web Security 101",
    thumbnail: "https://images.unsplash.com/photo-1550751827-4bd374c3f58b?w=500&auto=format&fit=crop&q=60",
    description: "Protect your web apps from hackers",
    instructor: "Emma Watson",
    instructorAvatar: "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=100&auto=format&fit=crop&q=60",
    level: "Level 10",
    rating: "4.9",
    subscribers: "10k",
    totalVideo: "1h 30m",
    promoText: "Free certificate included",
    isFree: true
  }
]);

const latestCourses = ref([
  {
    id: 6,
    tag: "New",
    tagType: "badge-new",
    contentType: "video",
    title: "Fullstack React",
    instructorAvatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQcK4HuUt3y__WH-qjoH5AEPSNe6ThFjwdQ0qWVKcMR7Q&s",
    thumbnail: "https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=500&auto=format&fit=crop&q=60",
    description: "Next.js & PostgreSQL complete guide",
    instructor: "Alex Morgan",
    level: "Level 10",
    rating: "4.9",
    subscribers: "10k",
    price: 55,
    originalPrice: 99,
    promoText: "New release special discount",
    isFree: false
  },
  {
    id: 7,
    tag: "New",
    tagType: "badge-new",
    contentType: "article",
    title: "Tailwind CSS Mastery",
    instructorAvatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ5MZXxiaRdWjPXSFlpNaWR2Eugs0zMn0dRz5lT17_qdrXY0gjg_BBT5XzX&s=10",
    thumbnail: "https://images.unsplash.com/photo-1507238691740-187a5b1d37b8?w=500&auto=format&fit=crop&q=60",
    description: "Craft modern responsive designs fast",
    instructor: "Jessica Lee",
    level: "Level 10",
    rating: "4.8",
    subscribers: "10k",
    price: 25,
    originalPrice: 50,
    promoText: "Hot trending course today",
    isFree: false
  },
  {
    id: 8,
    tag: "Free",
    tagType: "badge-free",
    contentType: "ebook",
    title: "Git & GitHub Workflow",
    instructorAvatar: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&auto=format&fit=crop&q=60",
    thumbnail: "https://images.unsplash.com/photo-1618401471353-b98aedd04e11?w=500&auto=format&fit=crop&q=60",
    description: "Version control for professional devs",
    instructor: "Kevin Hart",
    level: "Level 10",
    rating: "4.7",
    subscribers: "10k",
    totalVideo: "12h 30m",
    promoText: "Free starter handbook",
    isFree: true
  },
  {
    id: 9,
    tag: "Free",
    tagType: "badge-free",
    contentType: "video",
    title: "Docker & DevOps Basics",
    instructorAvatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSh-1ZGJAbdPuM9UkBRxHQ91RDTFywmECt6aaupWZmNFUUGWa22dhgp3S0&s=10",
    thumbnail: "https://images.unsplash.com/photo-1605745341112-85968b19335b?w=500&auto=format&fit=crop&q=60",
    description: "Containerize applications seamlessly",
    instructor: "Chris Hemsworth",
    level: "Level 10",
    rating: "4.9",
    subscribers: "10k",
    totalLessons: 12,
    promoText: "Open source module",
    isFree: true
  },
  {
    id: 10,
    tag: "Free",
    tagType: "badge-free",
    contentType: "article",
    title: "Cloud Computing Intro",
    instructorAvatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT7rGiYQz-D9gPYSMbPZPlihNAB9UycpD8CnWUmGWnuyw&s=10",
    thumbnail: "https://images.unsplash.com/photo-1451187580459-43490279c0fa?w=500&auto=format&fit=crop&q=60",
    description: "AWS fundamentals for beginners",
    instructor: "Tom Holland",
    level: "Level 10",
    rating: "4.5",
    subscribers: "10k",
    totalContent: 30,
    promoText: "Free learning roadmap",
    isFree: true
  }
]);

function playVideo(course) {
  alert(`Playing video: "${course.title}"`);
}

function handleSearch() {
  if (searchQuery.value.trim()) {
    alert(`Searching for: "${searchQuery.value}"`);
  }
}

function clearFilters() {
  selectedCourseCategory.value = "ALL";
  selectedCourseType.value = "ALL";
  selectedLevel.value = "ALL";
  searchQuery.value = "";
}

function enrollCourse(course) {
  router.push({
    name: 'AddCardCourse',
    params: { id: course.id },
    state: { courseData: course }
  });
}

const showCourseModal = ref(false);
const selectedCourseData = ref(null);

function closeCourseModal() {
  showCourseModal.value = false;
  selectedCourseData.value = null;
}

function goToCourseDetail(course) {
  router.push({
    name: 'CourseDetail',
    params: { id: course.id },
    state: { courseData: course }
  });
}
</script>

<template>
  <div>
    <NavBar/>
    <div class="dashboard-wrapper">
      <div class="dashboard-main-layout">

        <!-- SIDEBAR CONTAINER -->
        <div class="sidebar-wrapper">
          <aside class="filter-sidebar">
            <div class="search-box">
              <svg class="search-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
              <input
                type="text"
                v-model="searchQuery"
                placeholder="Search all courses..."
                @keyup.enter="handleSearch"
              />
            </div>
          </aside>

          <aside class="filter-sidebar">
            <!-- COURSE LIST -->
            <div class="filter-item">
              <label class="filter-label">COURSE</label>
              <ul class="filter-list">
                <li
                  v-for="cat in courseCategories"
                  :key="cat.label"
                  :class="['filter-list-item', { active: selectedCourseCategory === cat.label }]"
                  @click="selectedCourseCategory = cat.label"
                >
                  <span class="filter-icon" v-html="cat.icon"></span>
                  <span>{{ cat.label }}</span>
                </li>
              </ul>
            </div>

            <!-- COURSE TYPE LIST -->
            <div class="filter-item">
              <label class="filter-label">COURSE TYPE</label>
              <ul class="filter-list">
                <li
                  v-for="type in courseTypes"
                  :key="type.label"
                  :class="['filter-list-item', { active: selectedCourseType === type.label }]"
                  @click="selectedCourseType = type.label"
                >
                  <span class="filter-icon" v-html="type.icon"></span>
                  <span>{{ type.label }}</span>
                </li>
              </ul>
            </div>

            <!-- LEVEL LIST -->
            <div class="filter-item">
              <label class="filter-label">LEVEL</label>
              <ul class="filter-list">
                <li
                  v-for="lvl in courseLevels"
                  :key="lvl.label"
                  :class="['filter-list-item', { active: selectedLevel === lvl.label }]"
                  @click="selectedLevel = lvl.label"
                >
                  <span class="filter-icon" v-html="lvl.icon"></span>
                  <span>{{ lvl.label }}</span>
                </li>
              </ul>
            </div>

            <button class="clear-filters-btn" @click="clearFilters">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>
              Clear Filters
            </button>
          </aside>
        </div>

        <!-- MAIN CONTENT -->
        <div class="dashboard-container">

          <!-- SECTION 1: POPULAR COURSE -->
          <section class="course-section">
            <h2 class="section-title">Popular course</h2>

            <div class="courses-grid">
              <div 
                v-for="course in popularCourses" 
                :key="course.id" 
                class="course-card" 
                @click="goToCourseDetail(course)" 
                style="cursor: pointer;"
              >

                <div class="card-media-banner">
                  <img :src="course.thumbnail" alt="Thumbnail" class="banner-bg-img" />
                  <div class="banner-overlay"></div>
                  <div :class="course.tagType">{{ course.tag }}</div>

                  <button
                    v-if="course.contentType === 'video'"
                    class="content-type-icon play-icon-btn"
                    @click.stop="playVideo(course)"
                    title="Play video"
                  >
                    <svg viewBox="0 0 24 24" fill="currentColor"><path d="M8 5v14l11-7z"/></svg>
                  </button>

                  <div v-else-if="course.contentType === 'article'" class="content-type-icon" title="Article">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3" y="3" width="18" height="18" rx="2"/><circle cx="8.5" cy="8.5" r="1.5"/><path d="M21 15l-5-5L5 21"/></svg>
                  </div>

                  <div v-else-if="course.contentType === 'ebook'" class="content-type-icon" title="E-book">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1-2.5-2.5Z"/><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/></svg>
                  </div>

                  <h3 class="course-title-top">{{ course.title }}</h3>
                </div>

                <div class="card-body">
                  <p class="course-desc">{{ course.description }}</p>

                  <div class="instructor-row">
                    <div class="instructor-profile">
                      <div class="avatar-icon-wrap">
                        <img v-if="course.instructorAvatar" :src="course.instructorAvatar" :alt="course.instructor" class="instructor-avatar-img" />
                        <svg v-else class="meta-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                      </div>
                      <span class="instructor-name">{{ course.instructor }}</span>
                      <span class="level-badge">{{ course.level }}</span>
                    </div>
                  </div>

                  <div class="rating-sub-row">
                    <span class="sub-count">Sub: <strong>{{ course.subscribers }}</strong></span>
                    <div class="stars">
                      <svg class="star-svg" viewBox="0 0 24 24" fill="currentColor"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon></svg>
                      <strong>{{ course.rating }}</strong>
                    </div>
                  </div>

                  <div class="pricing-or-detail-area">
                    <template v-if="!course.isFree">
                      <div class="price-display">
                        <span class="current-price">${{ course.price }}</span>
                        <span class="old-price">${{ course.originalPrice }}</span>
                      </div>
                      <button class="action-btn buy-btn" @click.stop="enrollCourse(course)">Enroll Now</button>
                    </template>
                    <template v-else>
                      <div class="free-meta-info">
                        <span v-if="course.totalLessons">
                          <svg class="meta-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1-2.5-2.5Z"></path><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path></svg>
                          Lessons: <strong>{{ course.totalLessons }}</strong>
                        </span>
                        <span v-else-if="course.totalVideo">
                          <svg class="meta-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="23 7 16 12 23 17 23 7"></polygon><rect x="1" y="5" width="15" height="14" rx="2" ry="2"></rect></svg>
                          Video: <strong>{{ course.totalVideo }}</strong>
                        </span>
                        <span v-else>
                          <svg class="meta-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"></path><path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"></path></svg>
                          Content: <strong>30</strong>
                        </span>
                      </div>
                      <button class="action-btn free-btn" @click.stop="enrollCourse(course)">Start Free</button>
                    </template>
                  </div>
                </div>

                <div class="card-footer-promo">
                  <span>{{ course.promoText }}</span>
                </div>

              </div>
            </div>
          </section>

          <!-- SECTION 2: THE LATEST -->
          <section class="course-section">
            <h2 class="section-title">The latest</h2>

            <div class="courses-grid">
              <div 
                v-for="course in latestCourses" 
                :key="course.id" 
                class="course-card" 
                @click="goToCourseDetail(course)" 
                style="cursor: pointer;"
              >

                <div class="card-media-banner">
                  <img :src="course.thumbnail" alt="Thumbnail" class="banner-bg-img" />
                  <div class="banner-overlay"></div>
                  <div :class="course.tagType">{{ course.tag }}</div>

                  <button
                    v-if="course.contentType === 'video'"
                    class="content-type-icon play-icon-btn"
                    @click.stop="playVideo(course)"
                    title="Play video"
                  >
                    <svg viewBox="0 0 24 24" fill="currentColor"><path d="M8 5v14l11-7z"/></svg>
                  </button>

                  <div v-else-if="course.contentType === 'article'" class="content-type-icon" title="Article">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3" y="3" width="18" height="18" rx="2"/><circle cx="8.5" cy="8.5" r="1.5"/><path d="M21 15l-5-5L5 21"/></svg>
                  </div>

                  <div v-else-if="course.contentType === 'ebook'" class="content-type-icon" title="E-book">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1-2.5-2.5Z"/><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/></svg>
                  </div>

                  <h3 class="course-title-top">{{ course.title }}</h3>
                </div>

                <div class="card-body">
                  <p class="course-desc">{{ course.description }}</p>

                  <div class="instructor-row">
                    <div class="instructor-profile">
                      <div class="avatar-icon-wrap">
                        <img v-if="course.instructorAvatar" :src="course.instructorAvatar" :alt="course.instructor" class="instructor-avatar-img" />
                        <svg v-else class="meta-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                      </div>
                      <span class="instructor-name">{{ course.instructor }}</span>
                      <span class="level-badge">{{ course.level }}</span>
                    </div>
                  </div>

                  <div class="rating-sub-row">
                    <span class="sub-count">Sub: <strong>{{ course.subscribers }}</strong></span>
                    <div class="stars">
                      <svg class="star-svg" viewBox="0 0 24 24" fill="currentColor"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon></svg>
                      <strong>{{ course.rating }}</strong>
                    </div>
                  </div>

                  <div class="pricing-or-detail-area">
                    <template v-if="!course.isFree">
                      <div class="price-display">
                        <span class="current-price">${{ course.price }}</span>
                        <span class="old-price">${{ course.originalPrice }}</span>
                      </div>
                      <button class="action-btn buy-btn" @click.stop="enrollCourse(course)">Enroll Now</button>
                    </template>
                    <template v-else>
                      <div class="free-meta-info">
                        <span v-if="course.totalLessons">
                          <svg class="meta-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1-2.5-2.5Z"></path><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path></svg>
                          Lessons: <strong>{{ course.totalLessons }}</strong>
                        </span>
                        <span v-else-if="course.totalVideo">
                          <svg class="meta-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="23 7 16 12 23 17 23 7"></polygon><rect x="1" y="5" width="15" height="14" rx="2" ry="2"></rect></svg>
                          Video: <strong>{{ course.totalVideo }}</strong>
                        </span>
                        <span v-else>
                          <svg class="meta-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"></path><path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"></path></svg>
                          Content: <strong>30</strong>
                        </span>
                      </div>
                      <button class="action-btn free-btn" @click.stop="enrollCourse(course)">Start Free</button>
                    </template>
                  </div>
                </div>

                <div class="card-footer-promo">
                  <span>{{ course.promoText }}</span>
                </div>

              </div>
            </div>
          </section>

        </div>
      </div>
    </div>

    <div v-if="showCourseModal" class="modal-backdrop" @click.self="closeCourseModal">
      <AddCardCourse :courseData="selectedCourseData" />
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700&display=swap');

.dashboard-wrapper {
  background-color: #F7F4F2;
  min-height: 100vh;
  width: 100%;
  padding: 30px 20px;
  box-sizing: border-box;
}

.dashboard-main-layout {
  width: 100%;
  max-width: 1251px;
  margin: 0 auto;
  display: flex;
  align-items: flex-start;
  gap: 24px;
}

/* SIDEBAR WRAPPER */
.sidebar-wrapper {
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 240px;
  flex-shrink: 0;
  position: sticky;
  top: 20px;
}

/* SIDEBAR FILTER */
.filter-sidebar {
  display: flex;
  flex-direction: column;
  gap: 16px;
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 20px;
  width: 100%;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.04);
}

/* SEARCH BOX */
.search-box {
  display: flex;
  align-items: center;
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 32px;
  padding: 10px 14px;
  width: 100%;
}

.search-box:focus-within {
  border-color: #9ca3af;
  background: #ffffff;
}

.search-icon {
  color: #9ca3af;
  margin-right: 8px;
  flex-shrink: 0;
}

.search-box input {
  background: transparent;
  border: none;
  color: #374151;
  font-size: 13px;
  width: 100%;
  outline: none;
  min-width: 0;
  border-radius: 32px;
}

/* FILTER LIST STYLES */
.filter-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.filter-label {
  font-size: 11px;
  font-weight: 700;
  color: #374151;
  letter-spacing: 0.5px;
}

.filter-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.filter-list-item {
  font-size: 13px;
  color: #4b5563;
  padding: 6px 10px;
  border-radius: 32px;
  cursor: pointer;
  transition: all 0.2s ease;
  background: #f9fafb;

  display: flex;          /* ➕ បន្ថែម */
  align-items: center;    /* ➕ បន្ថែម */
  gap: 8px;   
}


.filter-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 16px;
  height: 16px;
  flex-shrink: 0;
  color: inherit;
}

.filter-icon svg {
  width: 100%;
  height: 100%;
}

.filter-list-item:hover {
  background: #f3f4f6;
  color: #111827;
}

.filter-list-item.active {
  background: #1976D2;
  color: #ffffff;
  font-weight: 600;
  border-radius: 32px;
}

.clear-filters-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  background: #f3f4f6;
  border: 1px solid #d1d5db;
  color: #374151;
  font-size: 12px;
  font-weight: 600;
  padding: 8px 12px;
  border-radius: 32px;
  cursor: pointer;
  transition: background 0.2s;
  margin-top: 4px;
}

.clear-filters-btn:hover {
  background: #e5e7eb;
}

/* MAIN CONTENT AREA */
.dashboard-container {
  flex: 1;
  min-width: 0;
  color: #1f2937;
  font-family: 'Plus Jakarta Sans', sans-serif;
}

/* SECTION TITLE */
.course-section {
  margin-bottom: 40px;
}

.section-title {
  font-size: 20px;
  font-weight: 700;
  color: #111827;
  margin-bottom: 20px;
  border-left: 4px solid #4b5563;
  padding-left: 12px;
}

/* GRID */
.courses-grid {
  display: grid;
    grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

/* COURSE CARD */
.course-card {
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.course-card:hover {
 opacity: 0.8;
}

/* VERTICAL BANNER */
.card-media-banner {
  position: relative;
  background: #4b5563;
  height: 180px;
  padding: 14px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: flex-start;
  overflow: hidden;
}

.banner-bg-img {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  z-index: 1;
}

.banner-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;

  z-index: 2;
}

.course-title-top {
  position: relative;
  z-index: 3;
  font-size: 16px;
  font-weight: 700;
  color: #ffffff;
  margin: 0;
  line-height: 1.3;
}

.badge-new,
.badge-free {
  position: relative;
  z-index: 3;
  color: #ffffff;
  font-size: 10px;
  font-weight: 700;
  padding: 4px 8px;
  border-radius: 4px;
  align-self: flex-start;
}

.badge-new {
  background: #1976D2;
}

.badge-free {
  background: #06e86f;
}

/* BODY */
.card-body {
  padding: 14px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  flex-grow: 1;
}

.course-desc {
  font-size: 12px;
  color: #6b7280;
  margin: 0;
  height: 32px;
  overflow: hidden;
}

.instructor-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-top: 1px solid #f3f4f6;
  padding-top: 8px;
}

.instructor-profile {
  display: flex;
  align-items: center;
  gap: 6px;
}

.avatar-icon-wrap {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  overflow: hidden;
  flex-shrink: 0;
  background: #f3f4f6;
}

.instructor-avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.meta-svg-icon {
  width: 13px;
  height: 13px;
  color: #6b7280;
  stroke-width: 2.2;
}

.instructor-name {
  font-size: 11px;
  font-weight: 600;
  color: #374151;
}

.level-badge {
  font-size: 9px;
  background: #f3f4f6;
  color: #4b5563;
  padding: 2px 5px;
  border-radius: 4px;
  font-weight: 700;
  margin-left: 4px;
}

.rating-sub-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #f3f4f6;
  padding-bottom: 8px;
  font-size: 10px;
  color: #6b7280;
}

.stars {
  display: flex;
  align-items: center;
  gap: 3px;
  color: #d97706;
}

.star-svg {
  width: 12px;
  height: 12px;
}

/* PRICE AREA */
.pricing-or-detail-area {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price-display {
  display: flex;
  align-items: baseline;
  gap: 6px;
}

.current-price {
  font-size: 16px;
  font-weight: 800;
  color: red;
}

.old-price {
  font-size: 11px;
  color: #9ca3af;
  text-decoration: line-through;
}

.free-meta-info {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 10px;
  color: #059669;
  font-weight: 600;
}

.free-meta-info span {
  display: flex;
  align-items: center;
  gap: 3px;
}

.free-meta-info .meta-svg-icon {
  color: #059669;
}

.action-btn {
  padding: 6px 14px;
  border-radius: 6px;
  font-size: 11px;
  font-weight: 700;
  cursor: pointer;
  border: 1px solid #d1d5db;
  transition: background 0.2s;
}

.action-btn:hover {
  background: #f9fafb;
}

.buy-btn {
  background: #1976D2;
  color: white;
  border-color: #1976D2;
  border-radius: 32px;
}

.buy-btn:hover {
  background: #1976D2;
  opacity: 0.8;
}

.free-btn {
  background: #06e86f;
  color: white;
  border-color: #06e86f;
  border-radius: 32px;
}

.free-btn:hover {
  background: #06e86f;
  opacity: 0.8;
}

/* FOOTER PROMO */
.card-footer-promo {
  background: #f9fafb;
  padding: 8px 14px;
  font-size: 10px;
  color: #6b7280;
  border-top: 1px solid #f3f4f6;
  text-align: center;
}


.content-type-icon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 3;
  width: 52px;
  height: 52px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.92);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #1976D2;
  box-shadow: 0 4px 14px rgba(0, 0, 0, 0.25);
  border: none;
}

.content-type-icon svg {
  width: 24px;
  height: 24px;
}

.play-icon-btn {
  cursor: pointer;
  transition: transform 0.15s ease, background 0.15s ease;
}

.play-icon-btn:hover {
  transform: translate(-50%, -50%) scale(1.08);
  background: #ffffff;
}

.play-icon-btn svg {
  /* play icon ត្រូវ shift បន្តិចទៅស្តាំដើម្បីមើលទៅចំកណ្តាលដោយភ្នែក (visual balance) */
  margin-left: 3px;
}


.content-type-icon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 3;
  width: 52px;
  height: 52px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.92);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #1976D2;
  box-shadow: 0 4px 14px rgba(0, 0, 0, 0.25);
  border: none;
}

.content-type-icon svg {
  width: 24px;
  height: 24px;
}

.play-icon-btn {
  cursor: pointer;
  transition: transform 0.15s ease, background 0.15s ease;
}

.play-icon-btn:hover {
  transform: translate(-50%, -50%) scale(1.08);
  background: #ffffff;
}

.play-icon-btn svg {
  margin-left: 3px; /* Visual balance សម្រាប់ play triangle */
}



.modal-backdrop {
  position: fixed;
  inset: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}
/* RESPONSIVE */
@media (max-width: 900px) {
  .dashboard-main-layout {
    flex-direction: column;
  }
  .sidebar-wrapper {
    width: 100%;
    position: static;
  }
}
</style>