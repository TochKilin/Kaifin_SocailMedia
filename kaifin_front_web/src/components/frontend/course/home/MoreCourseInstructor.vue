<template>
  <div class="page-wrapper">
    <NavBar />

    <!-- Hero Banner Section -->
    <div class="hero-banner">
      <div class="hero-inner">
        <!-- Hero Left Content -->
        <div class="hero-content">
          <h1 class="hero-title">Learn</h1>
          <p class="hero-subtitle">Learn Today, Lead Tomorrow.</p>
        </div>

        <!-- Hero Right Notes Banner -->
        <div class="hero-notes-container">
          <div class="note-card note-1">
            <div class="card-inner">
              <h2>Grow up widht Faifin</h2>
              <p>Start your learning journey today, gain valuable knowledge, and turn your ideas into real-world skills.</p>
            </div>
          </div>
          <div class="note-card note-2">
            <div class="card-inner">
              <h2>Grow up your selft</h2>
              <p>Learn at your own pace, improve your skills, and gain the confidence you need to achieve your dreams.</p>
            </div>
          </div>
          <div class="note-card note-3">
            <div class="card-inner">
              <h2>Learn More</h2>
              <p>Every course is a new opportunity to learn, grow, improve yourself, and move closer to the future you imagine.</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="more-courses-container">
      
      <!-- Section Title -->
      <h2 class="section-title">More courses from instructor</h2>

      <!-- Course Cards Grid -->
      <div class="courses-grid">
        <div v-for="(course, index) in courses" :key="index" class="course-card">
          
          <!-- Thumbnail / Image Box -->
          <div class="course-thumb-box">
            <img :src="course.thumbnail" :alt="course.title" class="course-thumb-img" />
          </div>

          <!-- Course Content Body -->
          <div class="course-content">
            <!-- Course Name -->
            <h3 class="course-name">{{ course.title }}</h3>

            <!-- Description -->
            <p class="course-desc">{{ course.description }}</p>

            <!-- Instructor Info with Unique Profile Image -->
            <div class="instructor-row">
              <span class="create-by-label">Created by:</span>
              <div class="instructor-profile">
                <img :src="course.instructorAvatar" :alt="course.instructor" class="instructor-avatar" />
                <span class="instructor-name">{{ course.instructor }}</span>
              </div>
            </div>

            <!-- Ratings & Level Info -->
            <div class="meta-row">
              <div class="rating-box">
                <span class="rating-number">{{ course.rating }}</span>
                <span class="star-icon">★</span>
                <span class="ratings-count">({{ course.ratingsCount }})</span>
              </div>
              <div class="level-box">
                {{ course.level }}
              </div>
            </div>

            <!-- Price Row -->
            <div class="price-row">
              <span class="current-price">${{ course.price }}</span>
              <span class="original-price">${{ course.originalPrice }}</span>
            </div>
          </div>

        </div>
      </div>

      <!-- Show More Button -->
      <div class="show-more-wrapper">
        <button class="show-more-btn" @click="$emit('show-more')">
          Show more <span class="arrow-down">↓</span>
        </button>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import NavBar from '../../navbar/NavBar.vue'

defineEmits(['show-more'])

const courses = ref([
  {
    title: 'Advanced Vue.js 3 Masterclass',
    description: 'Learn modern composition API, Pinia, and routing best practices.',
    instructor: 'John Smith',
    instructorAvatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100&auto=format&fit=crop&q=60',
    rating: 4.8,
    ratingsCount: '4,000',
    level: 'All Levels',
    price: '9.99',
    originalPrice: '69.99',
    thumbnail: 'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=600&auto=format&fit=crop&q=60'
  },
  {
    title: 'Complete Node.js Developer Guide',
    description: 'Build scalable backend applications with Express and MongoDB.',
    instructor: 'Sarah Jenkins',
    instructorAvatar: 'https://images.unsplash.com/photo-1517841905240-472988babdf9?w=100&auto=format&fit=crop&q=60',
    rating: 4.7,
    ratingsCount: '3,500',
    level: 'Intermediate',
    price: '8.99',
    originalPrice: '59.99',
    thumbnail: 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?w=600&auto=format&fit=crop&q=60'
  },
  {
    title: 'UI/UX Design Fundamentals',
    description: 'Master Figma and design systems from scratch with real projects.',
    instructor: 'Elena Rostova',
    instructorAvatar: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=100&auto=format&fit=crop&q=60',
    rating: 4.9,
    ratingsCount: '5,200',
    level: 'Beginner',
    price: '6.99',
    originalPrice: '49.99',
    thumbnail: 'https://images.unsplash.com/photo-1581291518633-83b4ebd1d83e?w=600&auto=format&fit=crop&q=60'
  },
  {
    title: 'React & Next.js Complete Guide',
    description: 'Build high-performance server-side rendered React applications.',
    instructor: 'David Miller',
    instructorAvatar: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&auto=format&fit=crop&q=60',
    rating: 4.8,
    ratingsCount: '6,100',
    level: 'All Levels',
    price: '11.99',
    originalPrice: '84.99',
    thumbnail: 'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=600&auto=format&fit=crop&q=60'
  },
  {
    title: 'Python for Data Science & AI',
    description: 'Learn Pandas, NumPy, Machine Learning, and Neural Networks.',
    instructor: 'Dr. Alex Chen',
    instructorAvatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=100&auto=format&fit=crop&q=60',
    rating: 4.6,
    ratingsCount: '2,900',
    level: 'Intermediate',
    price: '12.99',
    originalPrice: '99.99',
    thumbnail: 'https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5?w=600&auto=format&fit=crop&q=60'
  },
  {
    title: 'Flutter Mobile App Development',
    description: 'Create cross-platform iOS and Android apps with a single codebase.',
    instructor: 'Jessica Taylor',
    instructorAvatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=100&auto=format&fit=crop&q=60',
    rating: 4.7,
    ratingsCount: '3,100',
    level: 'All Levels',
    price: '9.99',
    originalPrice: '74.99',
    thumbnail: 'https://images.unsplash.com/photo-1512941937669-90a1b58e7e9c?w=600&auto=format&fit=crop&q=60'
  },
  {
    title: 'Cybersecurity Ethical Hacking',
    description: 'Understand network security, penetration testing, and defense strategies.',
    instructor: 'Marcus Vance',
    instructorAvatar: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=100&auto=format&fit=crop&q=60',
    rating: 4.9,
    ratingsCount: '4,800',
    level: 'Advanced',
    price: '14.99',
    originalPrice: '119.99',
    thumbnail: 'https://images.unsplash.com/photo-1550751827-4bd374c3f58b?w=600&auto=format&fit=crop&q=60'
  },
  {
    title: 'Full-Stack Web Development',
    description: 'From HTML/CSS to database design and deploying cloud apps.',
    instructor: 'Michael Brown',
    instructorAvatar: 'https://images.unsplash.com/photo-1519085360753-af0119f7cbe7?w=100&auto=format&fit=crop&q=60',
    rating: 4.8,
    ratingsCount: '7,300',
    level: 'Beginner',
    price: '10.99',
    originalPrice: '89.99',
    thumbnail: 'https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=600&auto=format&fit=crop&q=60'
  }
])
</script>

<style scoped>
/* Import Google Fonts for Khmer */
@import url('https://fonts.googleapis.com/css2?family=Kantumruy+Pro:wght@400;500;600;700&display=swap');

/* Page Wrapper with #F7F4F2 Background */
.page-wrapper {
  min-height: 100vh;
  background-color: #F7F4F2;
}

/* Hero Banner Full Width Background */
.hero-banner {
  width: 100%;
  color: #ffffff;
  overflow: hidden;
  display: flex;
  justify-content: center;
}

/* Inner Container with max-width: 1252px matching more-courses-container */
.hero-inner {
  background: linear-gradient(135deg, #1B75D2 0%, #13559a 100%);
  max-width: 1252px;
  width: 100%;
  margin: 0 auto;
  padding: 50px 24px;
  box-sizing: border-box;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 40px;
}

.hero-content {
  max-width: 380px;
}

.hero-title {
  font-size: 48px;
  font-weight: 800;
  margin: 0 0 12px 0;
  letter-spacing: -1px;
}

.hero-subtitle {
  font-size: 18px;
  margin: 0;
  opacity: 0.9;
  line-height: 1.5;
}

/* Hero Notes Banner Container */
.hero-notes-container {
  position: relative;
  width: 600px;
  height: 340px;
  display: flex;
  align-items: center;
}

.note-card {
  position: absolute;
  width: 260px;   
  height: 310px;  
  background-color: #ffffff; 
  color: #0f172a;
  padding: 24px;
  border-radius: 8px;
  box-sizing: border-box;
  font-family: 'Kantumruy Pro', 'Khmer OS Siemreap', sans-serif;
  font-size: 15px;
  font-weight: 500;
  line-height: 1.6;
  border: 1px solid #ddd;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
  transition: transform 0.3s ease;
  
  /* ប្រើ Flexbox ເພື່ອຈັດកណ្តាលអត្ថបទខាងក្នុង */
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.card-inner {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
}

.note-card h2 {
  font-size: 22px;
  font-weight: 700;
  color: #1B75D2;
  margin: 0 0 12px 0;
}

.note-card p {
  margin: 0;
  color: #475569;
}

/* Distributed horizontally and slightly rotated */
.note-1 {
  left: 0px;
  top: 10px;
  transform: rotate(-5deg);
  z-index: 1;
}

.note-2 {
  left: 170px;
  top: -15px;
  transform: rotate(2deg);
  z-index: 2;
}

.note-3 {
  left: 340px;
  top: 10px;
  transform: rotate(-3deg);
  z-index: 3;
}

.note-card:hover {
  transform: scale(1.03) translateY(-5px);
  z-index: 10;
}

@media (max-width: 1024px) {
  .hero-inner {
    flex-direction: column;
    text-align: center;
  }
  .hero-content {
    max-width: 100%;
  }
  .hero-notes-container {
    width: 100%;
    height: auto;
    margin-top: 30px;
    flex-direction: column;
    gap: 15px;
  }
  .note-card {
    position: relative;
    left: auto !important;
    top: auto !important;
    transform: none !important;
    margin: 0 auto;
  }
}

/* More Courses Container with max-width 1252px, #ffffff Background */
.more-courses-container {
  max-width: 1252px;
  width: 100%;
  margin: 0 auto;
  padding: 40px 24px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  box-sizing: border-box;
  background-color: #ffffff;
}

/* Title Header Box */
.section-title {
  font-size: 24px;
  font-weight: 700;
  color: #0f172a;
  margin-bottom: 28px;
  letter-spacing: -0.5px;
}

/* Grid Layout set to exactly 4 columns */
.courses-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
  margin-bottom: 36px;
}

/* Responsive layout for smaller screens */
@media (max-width: 1200px) {
  .courses-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .courses-grid {
    grid-template-columns: 1fr;
  }
}

/* Course Card */
.course-card {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
}

.course-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.06);
  border-color: #cbd5e1;
}

/* Thumbnail Box */
.course-thumb-box {
  width: 100%;
  height: 160px;
  background-color: #f1f5f9;
  overflow: hidden;
}

.course-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.course-card:hover .course-thumb-img {
  transform: scale(1.05);
}

/* Course Content Padding */
.course-content {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  flex: 1;
}

/* Course Name */
.course-name {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
  line-height: 1.4;
}

/* Description */
.course-desc {
  font-size: 13px;
  color: #64748b;
  margin: 0;
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Instructor Row with Profile Avatar */
.instructor-row {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #64748b;
}

.instructor-profile {
  display: flex;
  align-items: center;
  gap: 6px;
}

.instructor-avatar {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  object-fit: cover;
}

.instructor-name {
  font-weight: 600;
  color: #334155;
}

/* Meta Row (Rating & Level) */
.meta-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 2px;
}

.rating-box {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  font-weight: 600;
  color: #0f172a;
}

.star-icon {
  color: #f59e0b;
  font-size: 13px;
}

.ratings-count {
  color: #94a3b8;
  font-weight: 400;
}

.level-box {
  font-size: 11px;
  font-weight: 600;
  color: #1B75D2;
  background-color: #f1f5f9;
  padding: 3px 8px;
  border-radius: 20px;
}

/* Price Row */
.price-row {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 2px;
  padding-top: 10px;
  border-top: 1px solid #f1f5f9;
}

.current-price {
  font-size: 16px;
  font-weight: 700;
  color: red;
}

.original-price {
  font-size: 13px;
  color: #94a3b8;
  text-decoration: line-through;
}

/* Show More Button Wrapper */
.show-more-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 10px;
}

.show-more-btn {
  background-color: rgba(15, 23, 42, 0.04);
  color: #0f172a;
  border: none;
  padding: 12px 32px;
  border-radius: 32px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: all 0.2s ease;
  box-shadow: none;
}

.show-more-btn:hover {
  background-color: rgba(15, 23, 42, 0.08);
  transform: translateY(-1px);
}

.arrow-down {
  font-size: 14px;
}
</style>