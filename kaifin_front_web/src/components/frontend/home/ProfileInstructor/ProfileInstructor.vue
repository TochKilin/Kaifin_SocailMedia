<script setup>
import { ref } from 'vue';
import NavBar from '../../navbar/NavBar.vue';
import Phea from '@/assets/team_images/phea.jpg';
import { useRoute, useRouter } from 'vue-router';
import CreateCourse from './CreateCourse.vue';
import GroupList from './GroupList.vue';
import TabsFeed from './TabsFeed.vue';
import FeedBackCourse from './FeedBackCourse.vue';

const router = useRouter();
const isCreateCourseOpen = ref(false);
const openCreateCourseModal = () => {
  isCreateCourseOpen.value = true;
};

const closeCreateCourseModal = () => {
  isCreateCourseOpen.value = false;
};

const route = useRoute();
const profileId = ref(route.params.id);
const profileType = ref(route.query.type || 'Instructor');
const profileData = ref({
  username: profileType.value === 'Instructor' ? 'Instructor Pheourn Sophea' : 'User Profile',
  profileImage: Phea,
  level: 'Lvl 5',
  ratings: '4.8 ★',
  moreInfo: 'Passionate educator & full-stack developer.',
  following: 100,
  follower: 1000,
  group: 10,
  myCourse: 100,
  myStudent: 2000,
  startUse: '01-01-2026'
});

const activeTab = ref('Article');
const tabs = [
  {
    name: 'Save',
    svg: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>'
  },
  {
    name: 'Feed',
    svg: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/><rect x="14" y="14" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/></svg>'
  },
  {
    name: 'Article',
    svg: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>'
  },
  {
    name: 'Course',
    svg: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/></svg>'
  },
  {
    name: 'Group',
    svg: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>'
  },
  {
    name: 'Like',
    svg: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"/></svg>'
  },
  {
    name: 'Feedback',
    svg: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>'
  }
];

const tabContents = {
  Feed: [
    {
      text: 'Published a new comprehensive course on Full-Stack Development and modern architecture patterns.',
      time: '01-01-2026 10:30 AM',
      likes: '21K',
      comments: '989K',
      shares: '1.8K',
      saves: '55K'
    },
    
  ],
  
  Article: [
    {
      title: 'Building Secure Microservices',
      description: 'Best practices for API security, token authentication, and data encryption.',
      subText: 'Securing enterprise backend communication lines.',
      author: 'Vuthea',
      views: '850',
      likes: '215',
      image: 'https://images.unsplash.com/photo-1509316975850-ff9c5deb0cd9?w=500&auto=format&fit=crop&q=60',
      tags: ['#Backend', '#Security', '#API']
    },
    {
      title: 'Understanding Vue 3 Composition API',
      description: 'A deep dive into script setup and reactive states for scalable frontend apps.',
      subText: 'Modern component architecture techniques.',
      author: 'Vuthea',
      views: '1.2K',
      likes: '340',
      image: 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?w=500&auto=format&fit=crop&q=60',
      tags: ['#VueJS', '#Frontend', '#JavaScript']
    }
  ],
  Course: [
    {
      title: "Go fiber Advance course",
      datetime: "12/5/2026",
      ratings: "3.9",
      description: "Master the art of building AI-powered chatbots with Go, Fiber, and OpenAI...",
      contentType: "Total content",
      contentCount: "30",
      price: "44",
      originalPrice: "99",
      image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTSPEvfuewxE6hbacMrOkvRCOpWskRWsyH9_paymd8t1w&s=10"
    },
    {
      title: "Figma Basic course",
      datetime: "11/1/2026",
      ratings: "4.8",
      description: "Course 1: Creating a Chatbot with OpenAI in Go",
      contentType: "Total Lesson",
      contentCount: "30",
      price: "20",
      originalPrice: "50",
      image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSebpk8AZpGklpMqMbMCXQpWGixkL-liIf5zJyvzhaYsg&s=10"
    },
    {
      title: "UX UI Full Course",
      datetime: "12/5/2024",
      ratings: "4.5",
      description: "and the OpenAI API for a seamless user experience.",
      contentType: "Total Video",
      contentCount: "30",
      duration: "1h 30m",
      price: "50",
      originalPrice: "100",
      image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRFlOjujaQcoveC1Db4W7ugoIAf13h_UTIO37JQWRYQTA&s=10"
    }
  ],
  Save: [
    {
      text: 'Bookmarked advanced state management patterns and clean architecture articles for future reference.',
      time: '10-01-2026 02:15 PM',
      likes: '4.5K',
      comments: '120K',
      shares: '350',
      saves: '12.4K',
      category: 'Saved Article'
    },
    {
      text: 'Master the art of building AI-powered chatbots with Go, Fiber, and OpenAI. Learn to create intelligent conversations, design responsive interfaces, and develop a complete web application that combines Go, Fiber, and the OpenAI API for a seamless user experience.',
      time: '28-01-2026 02:15 PM',
      likes: '7.5K',
      comments: '400K',
      shares: '220K',
      saves: '12.4K',
      category: 'Saved Article'
    }
  ],
  Group: [
    { name: 'Full-Stack Developers Cambodia', members: '5.2K Members', role: 'Admin / Owner' }
  ],
  Like: [
    {
      text: 'Published a new comprehensive course on Full-Stack Development and modern architecture patterns.',
      time: '01-01-2026 10:30 AM',
      author: 'Instructor Pheourn Sophea',
      avatar: Phea,
      likes: '21K',
      comments: '989K',
      shares: '1.8K',
      saves: '55K',
      groupName: 'Course Group',
      likedCount: '150'
    }
  ],
  Feedback: [
    { user: 'Dara Chan', comment: 'The teaching style is exceptional! Very clear.', rating: '5 ★' }
  ]
};

const instructorBioData = ref({
  title: "Instructor Overview & Expertise",
  titleSvg: '<svg class="title-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>',
  description: "Dedicated full-stack software engineer and experienced instructor with a strong passion for teaching modern web and mobile development technologies to aspiring students.",
  statsInfo: [
    { id: 1, label: 'Profile Status', emoji: '🛡️', value: 'Verified' },
    { id: 2, label: 'Total Views', emoji: '👁️', value: '14.5K' },
    { id: 3, label: 'Satisfaction', emoji: '👍', value: '98%' }
  ],
  skills: [
    { name: 'Frontend', emoji: '💻' },
    { name: 'Android', emoji: '🤖' },
    { name: 'IOS', emoji: '🍏' },
    { name: 'Vue.js', emoji: '⚡' },
    { name: 'ASP.NET Core', emoji: '🌐' }
  ],
  imageUrl: "https://images.unsplash.com/photo-1522202176988-66273c2fd55f?w=500&auto=format&fit=crop&q=60"
});

function handleEdit() {
  alert('✏️ Opening Edit Profile modal...');
}

const handleCourseSubmit = (courseData) => {
  console.log('New course created:', courseData);
  closeCreateCourseModal();
};
</script>

<template>
  <div>
    <NavBar/>
    <div class="profile-page-container">
      <div class="profile-layout-grid">
        <div class="profile-left-section">
          <div class="wrap-tabs">
            <div class="profile-header-card">
              <div class="profile-avatar-large">
                <img :src="profileData.profileImage" alt="Profile Avatar" />
              </div>
              <div class="profile-user-main-info">
                <div class="username-row">
                  <h1 class="username-title">{{ profileData.username }}</h1>
                  <button class="edit-btn" @click="handleEdit">
                    <svg class="edit-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                      <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                    </svg>
                    <span>Edit Profile</span>
                  </button>
                </div>
                <div class="badges-row">
                  <span class="badge-pill level-pill">
                    <svg class="pill-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/></svg>
                    {{ profileData.level }}
                  </span>
                  <span class="badge-pill rating-pill">{{ profileData.ratings }}</span>
                </div>
                <div class="more-info-box">
                  <p>{{ profileData.moreInfo }}</p>
                </div>
              </div>
            </div>

            <div class="profile-tabs-wrapper">
              <div class="profile-tabs-row">
                <button
                  v-for="tab in tabs"
                  :key="tab.name"
                  :class="['tab-btn', { active: activeTab === tab.name }]"
                  @click="activeTab = tab.name"
                >
                  <span class="tab-icon" v-html="tab.svg"></span>
                  <span>{{ tab.name }}</span>
                </button>
              </div>
            </div>
          </div>

          <div class="profile-content-box">
            <template v-if="activeTab !== 'Article' && activeTab !== 'Feed' && activeTab !== 'Save' && activeTab !== 'Course' && activeTab !== 'Like' && activeTab !== 'Feedback'">
              <div class="instructor-bio-area">
                <div class="bio-title-wrapper">
                  <span class="title-icon-wrap" v-html="instructorBioData.titleSvg"></span>
                  <h3 class="bio-main-title">{{ instructorBioData.title }}</h3>
                </div>
                <p class="instructor-description">{{ instructorBioData.description }}</p>

                <div class="instructor-body-row">
                  <div class="instructor-left-content">
                    <div class="instructor-stats-row">
                      <div v-for="stat in instructorBioData.statsInfo" :key="stat.id" class="stat-chip">
                        <span class="chip-emoji">{{ stat.emoji }}</span>
                        <span>{{ stat.label }}:</span>
                        <strong>{{ stat.value }}</strong>
                      </div>
                    </div>
                    <div class="tech-tags-row">
                      <span v-for="(skill, index) in instructorBioData.skills" :key="index" class="tech-tag">
                        <span class="chip-emoji">{{ skill.emoji }}</span>
                        {{ skill.name }}
                      </span>
                    </div>
                  </div>
                  <div class="instructor-media-preview">
                    <img :src="instructorBioData.imageUrl" alt="Instructor preview" />
                  </div>
                </div>
              </div>
            </template>
            <div class="dynamic-tab-content-area">
              <template v-if="activeTab === 'Course'">
                <div class="course-tab-container">
                  <div class="course-top-action-bar">
                    <div class="course-filter-icons">
                      <button class="filter-icon-btn active" title="Image Course">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/></svg>
                      </button>
                      <button class="filter-icon-btn" title="Book Course">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/></svg>
                      </button>
                      <button class="filter-icon-btn" title="Video Course">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="5 3 19 12 5 21 5 3"/></svg>
                      </button>
                    </div>
                    <button class="create-course-btn" @click="openCreateCourseModal">Create new course</button>
                  </div>

                  <div class="course-items-list">
                    <div v-for="(course, index) in tabContents.Course" :key="index" class="course-card-item">
                      <div class="course-card-thumb">
                        <img :src="course.image" alt="Course Thumbnail" />
                      </div>
                      <div class="course-card-content">
                        <div class="course-header-row">
                          <h3 class="course-title">{{ course.title }}</h3>
                          <div class="course-meta-badges">
                            <span class="course-badge-outline">
                              <svg class="pill-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                              {{ course.datetime }}
                            </span>
                            <span class="course-badge-outline rating-badge">
                              <svg class="pill-svg-icon" viewBox="0 0 24 24" fill="currentColor" stroke="currentColor" stroke-width="2"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/></svg>
                              {{ course.ratings }}
                            </span>
                          </div>
                        </div>

                        <p class="course-desc">{{ course.description }}</p>

                        <div class="course-sub-info-row">
                          <div class="course-stat-pill">
                            <svg class="pill-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/></svg>
                            <span class="pill-label">{{ course.contentType }}</span>
                            <span class="pill-count">{{ course.contentCount }}</span>
                          </div>
                          <div v-if="course.duration" class="course-stat-pill">
                            <svg class="pill-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                            <span>{{ course.duration }}</span>
                          </div>
                          <button class="course-edit-btn">Edit</button>
                        </div>

                        <div class="course-pricing-row">
                          <span class="currency-symbol">$</span>
                          <span class="current-price">{{ course.price }}</span>
                          <span class="old-price">{{ course.originalPrice }}</span>
                        </div>
                      </div>
                    </div>
                  </div>

                  <div class="see-more-wrapper">
                    <button class="see-more-feedback-btn">
                      <span>{{ showAll ? 'Show less' : 'See more' }}</span>
                      <svg class="down-arrow-icon" :class="{ 'rotate-icon': showAll }" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <line x1="12" y1="5" x2="12" y2="19"></line>
                        <polyline points="19 12 12 19 5 12"></polyline>
                      </svg>
                    </button>
                  </div>
                </div>
              </template>

              <template v-else-if="activeTab === 'Feed'">
                <TabsFeed :profileData="profileData" />
              </template>

              <template v-else-if="activeTab === 'Article'">
                <div class="tab-list-container">
                  <div v-for="(article, index) in tabContents.Article" :key="index" class="article-custom-card">
                    <div class="article-top-content">
                      <div class="article-text-side">
                        <h2 class="article-main-heading">{{ article.title }}</h2>
                        <p class="article-desc-primary">{{ article.description }}</p>
                        <p class="article-desc-secondary">{{ article.subText }}</p>
                      </div>
                      <div class="article-thumbnail-box">
                        <img :src="article.image" alt="Article thumbnail" />
                      </div>
                    </div>

                    <div class="article-divider-line"></div>

                    <div class="article-footer-meta">
                      <div class="article-author-info">
                        <div class="author-avatar-circle">
                          <svg class="user-profile-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                            <circle cx="12" cy="7" r="4"></circle>
                          </svg>
                        </div>
                        <span class="author-name-text">{{ article.author }}</span>
                        <span class="meta-dot">•</span>
                        <span class="meta-stat-item">
                          <svg class="meta-icon-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
                          {{ article.views }}
                        </span>
                        <span class="meta-stat-item">
                          <svg class="meta-icon-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"/></svg>
                          {{ article.likes }}
                        </span>
                      </div>
                      <button class="article-options-btn" title="Options">
                        <svg class="options-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <circle cx="12" cy="12" r="1"></circle>
                          <circle cx="19" cy="12" r="1"></circle>
                          <circle cx="5" cy="12" r="1"></circle>
                        </svg>
                      </button>
                    </div>

                    <div class="article-tags-row">
                      <span v-for="(tag, tIdx) in article.tags" :key="tIdx" class="article-tag-badge">
                        {{ tag }}
                      </span>
                    </div>
                  </div>
                </div>
              </template>

              <template v-else-if="activeTab === 'Save'">
                <div class="tab-list-container">
                  <div v-for="(item, index) in tabContents.Save" :key="index" class="feed-item-card">
                    <div class="feed-user-header">
                      <img :src="profileData.profileImage" alt="Avatar" class="small-avatar-img" />
                      <div class="feed-user-meta">
                        <h4>{{ profileData.username }}</h4>
                        <span class="feed-time">
                          <svg class="time-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/></svg>
                          {{ item.time }}
                        </span>
                      </div>
                    </div>
                    <p class="feed-description">{{ item.text }}</p>

                    <div class="feed-interaction-bar">
                      <div class="stack-links">
                        <span class="interaction-item">
                          <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M2 12h20M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"/></svg>
                          Translate
                        </span>
                        <span class="interaction-item">
                          <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
                          {{ item.category }}
                        </span>
                      </div>
                    </div>

                    <div class="feed-footer-buttons">
                      <button class="footer-btn">
                        <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><polyline points="16 6 12 2 8 6"/><line x1="12" y1="2" x2="12" y2="15"/></svg>
                        <span>{{ item.shares }}</span>
                      </button>
                      <button class="footer-btn">
                        <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
                        <span>{{ item.comments }}</span>
                      </button>
                      <button class="footer-btn">
                        <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"/></svg>
                        <span>{{ item.likes }}</span>
                      </button>
                      <button class="footer-btn">
                        <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
                        <span>{{ item.saves }}</span>
                      </button>
                    </div>
                  </div>
                </div>
              </template>

              <template v-else-if="activeTab === 'Like'">
                <TabsFeed :profileData="profileData" />
              </template>

              <template v-else-if="activeTab === 'Group'">
                <div class="tab-list-container">
                  <GroupList />
                </div>
              </template>

              <template v-else-if="activeTab === 'Feedback'">
                <FeedBackCourse />
              </template>

              <template v-else>
                <div class="tab-list-container">
                  <h3 class="tab-category-title">📂 {{ activeTab }} List</h3>
                  <div v-for="(item, idx) in tabContents[activeTab]" :key="idx" class="generic-item-card">
                    <h4>{{ item.title || item.name || item.user }}</h4>
                    <p>{{ item.desc || item.category || item.students || item.members || item.author || item.comment }}</p>
                    <span class="item-sub-meta">{{ item.date || item.rating || item.role || item.likes || item.level }}</span>
                  </div>
                  <div v-if="tabContents[activeTab] && tabContents[activeTab].length === 0" class="no-data">
                    No data available in {{ activeTab }}.
                  </div>
                </div>
              </template>
            </div>
          </div>
        </div>

        <div class="profile-right-sidebar">
          <div class="stat-combined-card">
            <div class="stat-item">
              <span class="stat-label"><span class="chip-emoji">👥</span>Following</span>
              <span class="stat-value">{{ profileData.following }}</span>
            </div>
            <div class="stat-divider"></div>
            <div class="stat-item">
              <span class="stat-label"><span class="chip-emoji">⭐</span>Follower</span>
              <span class="stat-value">{{ profileData.follower }}</span>
            </div>
          </div>

          <div class="stat-card-single" @click="activeTab = 'Group'">
            <span class="stat-label">Group</span>
            <span class="stat-value sample-pill-style">{{ profileData.group }} <span class="badge-emoji">📁</span></span>
          </div>

          <div class="stat-card-single" @click="activeTab = 'Course'">
            <span class="stat-label">My Course</span>
            <span class="stat-value sample-pill-style">{{ profileData.myCourse }} <span class="badge-emoji">📚</span></span>
          </div>

          <div class="stat-card-single" @click="activeTab = 'Feed'">
            <span class="stat-label">My Student</span>
            <span class="stat-value sample-pill-style">{{ profileData.myStudent }} <span class="badge-emoji">🎓</span></span>
          </div>

          <div class="stat-card-single start-use-card" @click="activeTab = 'Save'">
            <span class="stat-label">Start Use</span>
            <span class="stat-value date-value sample-pill-style">{{ profileData.startUse }} <span class="badge-emoji">🚀</span></span>
          </div>
        </div>

      </div>

      <Teleport to="body">
        <Transition name="modal-fade">
          <div v-if="isCreateCourseOpen" class="modal-overlay" @click.self="closeCreateCourseModal">
            <div class="modal-box">
              <div class="modal-body">
                <CreateCourse @close="closeCreateCourseModal" @submit="handleCourseSubmit" />
              </div>
            </div>
          </div>
        </Transition>
      </Teleport>
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap');

.profile-page-container {
  min-height: 100vh;
  background-color: #F7F4F2;
  font-family: 'Plus Jakarta Sans', sans-serif;
  color: #1e293b;
}

.profile-layout-grid {
  max-width: 1100px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 280px;
  gap: 12px;
}

.profile-left-section {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.wrap-tabs {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.profile-header-card {
  background: #ffffff;
  border-bottom: 1px solid #e2e8f0;
  border-left: 1px solid #e2e8f0;
  border-right: 1px solid #e2e8f0;
  padding: 20px;
  display: flex;
  gap: 20px;
  align-items: center;
}

.profile-avatar-large {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  border: 2px solid #e2e8f0;
  overflow: hidden;
  background: #f1f5f9;
  flex-shrink: 0;
}

.profile-avatar-large img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.profile-user-main-info {
  flex: 1;
}

.username-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.username-title {
  font-size: 22px;
  font-weight: 700;
  margin: 0;
  color: #0f172a;
}

.edit-btn {
  background: transparent;
  border: 1px solid #cbd5e1;
  color: #334155;
  padding: 7px 12px;
  border-radius: 32px;
  font-weight: 600;
  font-size: 13px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: all 0.2s;
}

.edit-btn:hover {
  background: #f8fafc;
  color: #0f172a;
}

.edit-icon {
  width: 16px;
  height: 16px;
  stroke: currentColor;
}

.badges-row {
  display: flex;
  gap: 10px;
  margin-bottom: 8px;
  align-items: center;
}

.badge-pill {
  background: #f1f5f9;
  border: 1px solid #cbd5e1;
  color: #475569;
  padding: 4px 12px;
  border-radius: 32px;
  font-size: 12px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 6px;
}

.pill-svg-icon {
  width: 14px;
  height: 14px;
  stroke: currentColor;
}

.rating-pill {
  color: #ffffff;
  background: #1B75D2;
  border-color: #1B75D2;
}

.more-info-box {
  padding: 6px 0;
  font-size: 13px;
  color: #64748b;
}

.profile-tabs-wrapper {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
  padding: 0 16px;
}

.profile-tabs-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  overflow-x: visible;
  border-bottom: 1px solid #e2e8f0;
  width: 100%;
}

.tab-btn {
  background: transparent;
  border: none;
  color: #64748b;
  padding: 14px 10px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  flex: 1;
  justify-content: center;
  white-space: nowrap;
  display: flex;
  align-items: center;
  gap: 8px;
  position: relative;
  transition: color 0.2s;
}

.tab-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 18px;
  height: 18px;
}

.tab-icon :deep(svg) {
  width: 100%;
  height: 100%;
  stroke: currentColor;
}

.tab-btn:hover {
  color: #1B75D2;
}

.tab-btn.active {
  color: #1B75D2;
}

.tab-btn.active::after {
  content: '';
  position: absolute;
  bottom: -1px;
  left: 0;
  right: 0;
  height: 3px;
  background-color: #1B75D2;
  border-radius: 3px 3px 0 0;
}

.profile-content-box {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  padding: 20px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.course-tab-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.course-top-action-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-radius: 12px;
}

.course-filter-icons {
  display: flex;
  gap: 8px;
}

.filter-icon-btn {
  background-color: transparent;
  border: 1px solid #cbd5e1;
  color: #64748b;
  width: 30px;
  height: 30px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
}

.filter-icon-btn svg {
  width: 20px;
  height: 20px;
  stroke: currentColor;
}

.create-course-btn {
  background: #1B75D2;
  border: none;
  color: #ffffff;
  padding: 10px 20px;
  border-radius: 32px;
  font-weight: 700;
  font-size: 14px;
  cursor: pointer;
  transition: opacity 0.2s;
}

.create-course-btn:hover {
  opacity: 0.9;
}

.course-items-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.course-card-item {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  padding: 16px;
  display: flex;
  gap: 16px;
  align-items: center;
  box-shadow: 0 4px 12px rgba(227, 7, 7, 0.03);
}

.course-card-thumb {
  width: 130px;
  height: 110px;
  background: #f1f5f9;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border: 1px solid #e2e8f0;
  flex-shrink: 0;
}

.course-card-thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.course-card-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.course-header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.course-title {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

.course-meta-badges {
  display: flex;
  gap: 8px;
}

.course-badge-outline {
  background: #f8fafc;
  border: 1px solid #cbd5e1;
  color: #475569;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  gap: 6px;
  display: flex;
  align-items: center;
}

.course-badge-outline .pill-svg-icon {
  width: 14px;
  height: 14px;
  stroke: currentColor;
}
.rating-badge {
  background: #ffffff;
  color: #1B75D2;
  border-color: #1B75D2;
}

.course-desc {
  font-size: 13px;
  color: #64748b;
  margin: 0;
}

.course-sub-info-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.course-stat-pill {
  background: #f8fafc;
  border: 1px solid #cbd5e1;
  color: #334155;
  padding: 4px 12px;
  border-radius: 8px;
  font-size: 12px;
  display: flex;
  align-items: center;
  gap: 6px;
}

.course-edit-btn {
  margin-left: auto;
  background: #1B75D2;
  border: none;
  color: #ffffff;
  padding: 6px 20px;
  border-radius: 20px;
  font-weight: 600;
  font-size: 13px;
  cursor: pointer;
  transition: opacity 0.2s;
}

.course-edit-btn:hover {
  opacity: 0.9;
}

.course-pricing-row {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-top: 4px;
}

.currency-symbol {
  font-weight: 700;
  color: #16a34a;
  font-size: 15px;
}

.current-price {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}

.old-price {
  color: #94a3b8;
  text-decoration: line-through;
  font-size: 13px;
  font-weight: 600;
  margin-left: 4px;
}


.see-more-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 5px;
}

.see-more-feedback-btn {
  background-color: #9c9a9a41;
  border: none;
  border-radius: 32px;
  padding: 10px 30px;
  font-size: 14px;
  font-weight: 600;
  color: #000;
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
}

.see-more-feedback-btn:hover {
  background-color: #f9fafb;
  border-color: #3b82f6;
  color: #2563eb;
  transform: translateY(-2px);
}

.down-arrow-icon {
  width: 16px;
  height: 16px;
  transition: transform 0.3s ease;
}

.rotate-icon {
  transform: rotate(180deg);
}

/* Common Styles */
.instructor-bio-area {
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.bio-title-wrapper {
  display: flex;
  align-items: center;
  gap: 8px;
  border-bottom: 1px solid #e2e8f0;
  padding-bottom: 8px;
}

.bio-main-title {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

.instructor-description {
  font-size: 13px;
  color: #475569;
  margin: 0;
  line-height: 1.5;
}

.instructor-body-row {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 16px;
}

.instructor-left-content {
  display: flex;
  flex-direction: column;
  gap: 10px;
  flex: 1;
}

.instructor-media-preview {
  width: 120px;
  height: 120px;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #e2e8f0;
  flex-shrink: 0;
}

.instructor-media-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.instructor-stats-row, .tech-tags-row {
  display: flex;
  gap: 8px;
  flex-wrap: nowrap;
  overflow-x: auto;
  padding-bottom: 2px;
}

.stat-chip, .tech-tag {
  border: 1px solid #cbd5e1;
  background: #ffffff;
  color: #334155;
  padding: 6px 12px;
  border-radius: 32px;
  font-size: 12px;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 6px;
  white-space: nowrap;
  flex-shrink: 0;
}

.chip-emoji {
  font-size: 14px;
  line-height: 1;
}

.stat-chip strong {
  color: #0f172a;
  margin-left: 2px;
}

.article-custom-card {
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  padding: 20px;
  background: #ffffff;
  display: flex;
  flex-direction: column;
  gap: 14px;
  box-shadow: 0 4px 12px rgba(227, 7, 7, 0.03);
}

.article-top-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 20px;
}

.article-text-side {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.article-main-heading {
  font-size: 20px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

.article-desc-primary {
  font-size: 14px;
  color: #334155;
  margin: 0;
  line-height: 1.4;
}

.article-desc-secondary {
  font-size: 14px;
  color: #64748b;
  margin: 0;
  line-height: 1.4;
}

.article-thumbnail-box {
  width: 100px;
  height: 80px;
  border-radius: 12px;
  overflow: hidden;
  border: 1px solid #e2e8f0;
  flex-shrink: 0;
}

.article-thumbnail-box img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.article-divider-line {
  height: 1px;
  background-color: #f1f5f9;
  width: 100%;
}

.article-footer-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.article-author-info {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #475569;
}

.author-avatar-circle {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: #f1f5f9;
  border: 1px solid #cbd5e1;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #64748b;
}

.user-profile-svg {
  width: 18px;
  height: 18px;
  stroke: currentColor;
}

.author-name-text {
  font-weight: 700;
  color: #0f172a;
}

.meta-dot {
  color: #cbd5e1;
}

.meta-stat-item {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #64748b;
  font-size: 13px;
}

.meta-icon-svg {
  width: 15px;
  height: 15px;
  stroke: currentColor;
}

.article-options-btn {
  background: transparent;
  border: 1px solid #e2e8f0;
  color: #64748b;
  cursor: pointer;
  padding: 0;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 24px;
  transition: all 0.2s;
}

.options-svg-icon {
  width: 16px;
  height: 16px;
  stroke: currentColor;
  fill: currentColor;
}

.article-options-btn:hover {
  background: #f1f5f9;
  color: #0f172a;
  border-color: #cbd5e1;
}

.article-tags-row {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.article-tag-badge {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  color: #475569;
  padding: 4px 10px;
  border-radius: 32px;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.article-tag-badge:hover {
  background: #f1f5f9;
  color: #1B75D2;
  border-color: #cbd5e1;
}

.feed-item-card {
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 14px;
  background: #ffffff;
  display: flex;
  flex-direction: column;
  gap: 10px;
  box-shadow: 0 4px 12px rgba(227, 7, 7, 0.03);
}

.feed-user-header {
  display: flex;
  align-items: center;
  gap: 10px;
}

.small-avatar-img {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: 1px solid #cbd5e1;
  object-fit: cover;
}

.feed-user-meta h4 {
  margin: 0;
  font-size: 14px;
  color: #0f172a;
}

.feed-time {
  font-size: 11px;
  color: #94a3b8;
  display: flex;
  align-items: center;
  gap: 4px;
}

.time-svg-icon {
  width: 12px;
  height: 12px;
  stroke: currentColor;
}

.feed-description {
  font-size: 13px;
  color: #334155;
  margin: 0;
  border-top: 1px solid #f1f5f9;
  padding: 8px 0;
}

.feed-interaction-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
  color: #64748b;
}

.stack-links {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.interaction-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

.liked-by-container {
  display: flex;
  align-items: center;
  gap: 8px;
}

.liked-avatars-group {
  display: flex;
  align-items: center;
}

.liked-avatar {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  border: 2px solid #ffffff;
  object-fit: cover;
  background: #e2e8f0;
  margin-left: -8px;
}

.liked-count-text {
  font-size: 12px;
  font-weight: 600;
  color: #475569;
}

.svg-icon {
  width: 18px;
  height: 18px;
  stroke: currentColor;
  flex-shrink: 0;
}

.feed-footer-buttons {
  display: flex;
  justify-content: flex-end;
  border-top: 1px solid #f1f5f9;
  padding-top: 10px;
  gap: 8px;
}

.footer-btn {
  background: #1B75D2;
  border: none;
  color: #ffffff;
  font-size: 14px;
  cursor: pointer;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 10px 20px;
  border-radius: 32px;
  transition: opacity 0.2s;
}

.footer-btn:hover {
  opacity: 0.9;
  color: #ffffff;
}

.tab-list-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.tab-category-title {
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
  margin: 0 0 4px 0;
}

.generic-item-card {
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  padding: 14px;
  background: #f8fafc;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.no-data {
  font-size: 13px;
  color: #94a3b8;
  text-align: center;
  padding: 20px;
}

.profile-right-sidebar {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  padding: 14px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  height: fit-content;
  margin-top: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.stat-combined-card {
  border: none;
  border-radius: 10px;
  background: #ffffff;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 8px;
}

.stat-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
}

.stat-divider {
  width: 1px;
  height: 24px;
  background-color: #cbd5e1;
  margin: 0 4px;
}

.stat-card-single {
  border: none;
  border-radius: 10px;
  padding: 8px 10px;
  background: transparent;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  transition: background-color 0.2s ease;
}

.stat-label {
  font-size: 12px;
  font-weight: 600;
  color: #64748b;
  display: flex;
  align-items: center;
  gap: 6px;
}

.stat-value {
  font-size: 12px;
  font-weight: 700;
  color: #0f172a;
}

.sample-pill-style {
  background-color: #1B75D2;
  color: #ffffff;
  padding: 4px 8px;
  border-radius: 50px;
  font-weight: 400;
  font-size: 12px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
}

.stat-card-single:hover {
  background-color: rgba(241, 245, 249, 0.8);
  cursor: pointer;
  border-radius: 32px;
}

.badge-emoji {
  font-size: 16px;
}

.start-use-card {
  border: none !important;
}

.date-value {
  font-size: 14px;
}

.modal-overlay {
  position: fixed;
  top: 64px;
  left: 0;
  width: 100vw;
  height: calc(100vh - 64px); 
  background-color: #97999a8b; 
  display: flex;
  justify-content: center;
  align-items: flex-start; 
  overflow-y: auto; 
  z-index: 1000;
  /* padding: 30px 20px; */
  padding-top: 12px;
  box-sizing: border-box;
}

.modal-box.custom-form-box,
.modal-box {
  background: #ffffff;
  border-radius: 12px;
  width: 100%;
  max-width: 800px;
  margin-bottom: 40px;
  box-sizing: border-box;
}

.modal-box {
  background: #ffffff;
  border-radius: 12px;
  width: 100%;
  max-width: 890px; 
  border: 1px solid #e2e8f0;
  margin-bottom: 50px;
  box-sizing: border-box;
  overflow: hidden;
  /* margin-top: 290px; */
  /* margin-right: 200px; */
  margin-right: 0;
}

.modal-box.custom-form-box {
  background: #ffffff;
  border-radius: 16px;
  width: 100%;
  max-width: 800px;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 8px 10px -6px rgba(0, 0, 0, 0.1);
  margin-bottom: 40px;
  box-sizing: border-box;
}
.modal-body {
  padding: 24px;
}

@keyframes modalScaleUp {
  from {
    opacity: 0;
    transform: translateY(20px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.25s ease;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 18px 24px;
  border-bottom: 1px solid #e2e8f0;
  flex-shrink: 0;
}

.modal-header h2 {
  color: #0f172a;
  font-size: 18px;
  font-weight: 700;
  margin: 0;
}

.modal-close-btn {
  background: transparent;
  border: 1px solid #e2e8f0;
  color: #475569;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background 0.2s;
}

.modal-close-btn:hover {
  background: #f1f5f9;
}

.modal-close-btn svg {
  width: 18px;
  height: 18px;
  stroke: currentColor;
}

.modal-body {
  padding: 16px 20px;
}

.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.25s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
.modal-fade-enter-active .modal-box,
.modal-fade-leave-active .modal-box {
  transition: transform 0.25s ease;
}
.modal-fade-enter-from .modal-box,
.modal-fade-leave-to .modal-box {
  transform: scale(0.95) translateY(10px);
}
</style>