<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import NavBar from "../navbar/NavBar.vue";

const router = useRouter();
const searchQuery = ref("");
const isLoading = ref(true);
const BASE_URL = import.meta.env.VITE_API_URL

function authHeaders() {
  const token = localStorage.getItem("token");
  return token ? { Authorization: `Bearer ${token}` } : {};
}

function resolveAvatarUrl(raw) {
  if (!raw) return "";
  if (raw.startsWith("http://") || raw.startsWith("https://")) return raw;
  return `${BASE_URL}/uploads/${raw}`;
}

const topCourses = ref([]);
const coursesError = ref("");
const fetchTopCourses = async () => {
  isLoading.value = true;
  coursesError.value = "";
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/courses/show`, {
      params: { page: 1, limit: 6 },
      headers: authHeaders(),
    });
    const data = res.data?.data ?? res.data;
    const list = data?.courses ?? [];
    topCourses.value = list.map((c) => ({
      id: c.id,
      title: c.title,
      description: c.description || c.subtitle || "",
      students: `${c.students_count ?? 0} students`,
      lessons: `${c.lectures_count ?? 0} lessons`,
      rating: (c.rating ?? 0).toFixed(1),
      price: c.is_free ? "Free" : `$${(c.current_price ?? 0).toFixed(2)}`,
      isPopular: (c.ratings_count ?? 0) >= 50,
      image: c.thumbnail || "https://via.placeholder.com/500x300?text=No+Image",
      studentAvatars: [],
    }));
  } catch (err) {
    console.error("Failed to fetch top courses:", err);
    coursesError.value = err.message || "Failed to load courses";
  } finally {
    isLoading.value = false;
  }
};

const topInstructors = ref([]);
const instructorsError = ref("");
const isLoadingInstructors = ref(true);
const TOP_INSTRUCTOR_LIMIT = 6;

async function fetchInstructorProfile(instructorId) {
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/profile/show`, {
      params: { id: instructorId },
      headers: authHeaders(),
    });
    const data = res.data?.data ?? res.data;
    return {
      username: [data.first_name, data.last_name].filter(Boolean).join(" ") || data.user_name || `Instructor #${instructorId}`,
      avatar: resolveAvatarUrl(data.profile_images) || `https://api.dicebear.com/7.x/avataaars/svg?seed=${instructorId}`,
    };
  } catch (err) {
    console.error(`Failed to load profile for instructor ${instructorId}:`, err);
    return {
      username: `Instructor #${instructorId}`,
      avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${instructorId}`,
    };
  }
}

const fetchTopInstructors = async () => {
  isLoadingInstructors.value = true;
  instructorsError.value = "";
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/courses/show`, {
      params: { page: 1, limit: 200 },
      headers: authHeaders(),
    });
    const data = res.data?.data ?? res.data;
    const courses = data?.courses ?? [];
    const grouped = {};
    courses.forEach((c) => {
      const key = c.instructor_id;
      if (!key) return;
      if (!grouped[key]) {
        grouped[key] = {
          instructorId: key,
          totalStudents: 0,
          ratingSum: 0,
          ratedCourseCount: 0,
          courseCount: 0,
        };
      }
      grouped[key].totalStudents += c.students_count ?? 0;
      if (c.rating) {
        grouped[key].ratingSum += c.rating;
        grouped[key].ratedCourseCount += 1;
      }
      grouped[key].courseCount += 1;
    });
    const ranked = Object.values(grouped)
      .map((g) => {
        const avgRating = g.ratedCourseCount ? g.ratingSum / g.ratedCourseCount : 0;
        const score =
          g.totalStudents * 0.4 +
          avgRating * 20 * 0.3 +
          g.courseCount * 5 * 0.1;
        return { ...g, avgRating, score };
      })
      .sort((a, b) => b.score - a.score)
      .slice(0, TOP_INSTRUCTOR_LIMIT);

    if (!ranked.length) {
      topInstructors.value = [];
      return;
    }
    const [profiles, followInfos] = await Promise.all([
      Promise.all(ranked.map((r) => fetchInstructorProfile(r.instructorId))),
      Promise.all(ranked.map((r) => fetchFollowInfo(r.instructorId))),
    ]);

    topInstructors.value = ranked.map((r, idx) => ({
      id: r.instructorId,
      username: profiles[idx].username,
      avatar: profiles[idx].avatar,
      level: "",
      students: `${r.totalStudents} students`,
      rating: r.avgRating.toFixed(1),
      follower: formatFollowerCount(followInfos[idx].followerCount),
      isFollowing: followInfos[idx].isFollowing,
      isFollowingLoading: false,
    }));
  } catch (err) {
    console.error("Failed to compute top instructors:", err);
    instructorsError.value = err.message || "Failed to load instructors";
  } finally {
    isLoadingInstructors.value = false;
  }
};

const topUsers = ref([]);
const isLoadingUsers = ref(true);
const usersError = ref("");

async function fetchFollowInfo(userId) {
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/followers/show`, {
      params: { user_id: userId },
      headers: authHeaders(),
    });
    const data = res.data?.data ?? res.data;
    return {
      followerCount: data?.follower_count ?? 0,
      isFollowing: data?.is_following ?? false,
    };
  } catch (err) {
    console.error(`Failed to load follow info for user ${userId}:`, err);
    return { followerCount: 0, isFollowing: false };
  }
}

function formatFollowerCount(n) {
  if (n >= 1000000) return (n / 1000000).toFixed(1).replace(/\.0$/, "") + "M";
  if (n >= 1000) return (n / 1000).toFixed(1).replace(/\.0$/, "") + "K";
  return String(n);
}

const fetchTopUsers = async () => {
  isLoadingUsers.value = true;
  usersError.value = "";
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/levels/leaderboard`, {
      params: { limit: 6 },
      headers: authHeaders(),
    });
    const list = res.data?.data ?? [];
    const followInfos = await Promise.all(
      list.map((u) => fetchFollowInfo(u.user_id))
    );

    topUsers.value = list.map((u, idx) => ({
      id: u.user_id,
      username: u.user_name || `User #${u.user_id}`,
      level: `Lvl ${Math.max(1, Math.floor(u.longest_streak / 3) + 1)}`,
      streak: String(u.current_streak),
      avatar: resolveAvatarUrl(u.profile_images) || `https://api.dicebear.com/7.x/avataaars/svg?seed=${u.user_id}`,
      follower: formatFollowerCount(followInfos[idx].followerCount),
      isFollowing: followInfos[idx].isFollowing,  
    }));
  } catch (err) {
    console.error("Failed to fetch top users:", err);
    usersError.value = err.message || "Failed to load users";
  } finally {
    isLoadingUsers.value = false;
  }
};


const sponsors = ref([]);
const sponsorsError = ref("");
const isLoadingSponsors = ref(true);

function resolveSponsorLogoUrl(raw) {
  if (!raw) return "";
  if (raw.startsWith("http://") || raw.startsWith("https://")) return raw;
  return `${BASE_URL}/uploads/${raw}`;
}

const fetchSponsors = async () => {
  isLoadingSponsors.value = true;
  sponsorsError.value = "";
  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/sponsors/show`, {
      params: { page: 1, per_page: 20 },
      headers: authHeaders(),
    });
    const data = res.data?.data ?? res.data;
    const list = data?.sponsors ?? data?.Sponsors ?? [];

    sponsors.value = list.map((s) => ({
      id: s.id,
      title: s.name,
      image: resolveSponsorLogoUrl(s.logo_image),
      logo: resolveSponsorLogoUrl(s.logo_image),
      websiteUrl: s.website_url || "",
      isVerified: s.is_verified ?? false,
    }));
  } catch (err) {
    console.error("Failed to fetch sponsors:", err);
    sponsorsError.value = err.message || "Failed to load sponsors";
  } finally {
    isLoadingSponsors.value = false;
  }
};

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

async function handleFollow(item, type) {
  if (type !== 'User' && type !== 'Instructor') {
    alert(`You are now following ${type}: ${item.username}`);
    return;
  }

  if (item.isFollowingLoading) return;
  item.isFollowingLoading = true;
  try {
    const res = await axios.post(
      `${BASE_URL}/api/v1/front/followers/create`,
      { user_id: item.id },
      { headers: authHeaders() }
    );
    const data = res.data?.data ?? res.data;

    const list = type === 'User' ? topUsers.value : topInstructors.value;
    const target = list.find((u) => u.id === item.id);
    if (target) {
      target.isFollowing = data?.is_following ?? false;
      target.follower = formatFollowerCount(data?.follower_count ?? 0);
    }
  } catch (err) {
    console.error('Follow toggle failed', err);
    alert('Failed to update follow status');
  } finally {
    item.isFollowingLoading = false;
  }
}

function handleSponsorReadMore(sponsor) {
  alert(`Reading more about sponsor: ${sponsor.title}`);
}

function viewMore(section) {
  alert(`Viewing all ${section}...`);
}

function openChat() {
  alert("Opening Chat Room...");
}

function openMoreOptions() {
  alert("Opening More Options...");
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
  fetchTopCourses();
  fetchTopInstructors();
  fetchTopUsers(); 
   fetchSponsors();  
});
</script>

<template>
  <div class="page-container">
    <NavBar />
    <div class="content-wrap">
      <!-- Main layout-->
      <div class="main-layout-grid">
        <!-- Main content -->
        <div class="sections-container" :class="{ 'fade-in': !isLoading }">
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
                      <div class="stacked-avatars" v-if="course.studentAvatars.length">
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

                    <div v-if="!isLoading && coursesError" class="state-box error" style="text-align:center; padding: 20px; color: #dc3545;">
                      <p>Can't load courses — {{ coursesError }}</p>
                      <button class="all-btn" @click="fetchTopCourses">Try again</button>
                    </div>
                    <div v-else-if="!isLoading && topCourses.length === 0" class="state-box" style="text-align:center; padding: 20px; color: #6b7280;">
                      <p>No courses yet</p>
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

          <div class="row-section-box" :class="{ 'slide-up': !isLoading }">
            <div class="section-header">
              <h2>
                Top Instructor
              </h2>
              <button class="all-btn" @click="viewMore('Top Instructor')">View All &rarr;</button>
            </div>

            <div v-if="isLoadingInstructors" class="skeleton-wrapper">
              <div v-for="n in 3" :key="'skeleton-instr-' + n" class="skeleton-card">
                <div class="skeleton-image"></div>
                <div class="skeleton-content">
                  <div class="skeleton-line w-80"></div>
                </div>
              </div>
            </div>

            <div v-else-if="instructorsError" class="state-box error" style="text-align:center; padding: 20px; color: #dc3545;">
              <p>Can't load instructors — {{ instructorsError }}</p>
              <button class="all-btn" @click="fetchTopInstructors">Try again</button>
            </div>

            <div v-else-if="topInstructors.length === 0" class="state-box" style="text-align:center; padding: 20px; color: #6b7280;">
              <p>No instructors with courses yet</p>
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
                    <p>{{ instructor.students }} • ⭐ {{ instructor.rating }}</p>
                  </div>

                  <button
                    class="action-btn follow-effect-blue"
                    :class="{ 'is-following': instructor.isFollowing }"
                    :disabled="instructor.isFollowingLoading"
                    @click.stop="handleFollow(instructor, 'Instructor')"
                  >
                    <svg
                      v-if="instructor.isFollowing"
                      width="12" height="12" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"
                    >
                      <polyline points="20 6 9 17 4 12"></polyline>
                    </svg>
                    {{ instructor.isFollowing ? 'Following' : 'Follow' }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="sidebar-right-container" :class="{ 'fade-in': !isLoading }">
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

          <div class="sidebar-discover-card">
            <div class="hero-text">
              <h1 class="page-title">
                Discover Features
                <span class="badge-new">New</span>
              </h1>
              <p class="page-subtitle">Explore top trending courses and elite instructors in our community.</p>
            </div>
          </div>

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
                <div
                  class="sidebar-user-avatar"
                  :class="{ 'is-following': user.isFollowing }"
                >
                  <img :src="user.avatar" :alt="user.username" loading="lazy" />
                  <div class="online-status"></div>
                  <div v-if="user.isFollowing" class="follow-check-badge">
                    <svg width="9" height="9" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round">
                      <polyline points="20 6 9 17 4 12"></polyline>
                    </svg>
                  </div>
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

                <button
                  class="user-add-btn"
                  :class="{ 'is-following': user.isFollowing }"
                  :disabled="user.isFollowingLoading"
                  @click.stop="handleFollow(user, 'User')"
                  aria-label="Follow user"
                >
                  <svg
                    v-if="user.isFollowing"
                    width="12"
                    height="12"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="3"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline>
                  </svg>
                  {{ user.isFollowing ? 'Following' : 'Follow' }}
                </button>
              </div>
            </div>
          </div>

          <div class="sidebar-right-panel sponsors-sidebar-card">
            <div class="sidebar-section-header">
              <h2 class="sidebar-title">Sponsors</h2>
              <button class="all-btn" @click="viewMore('Sponsors')">See All</button>
            </div>

            <div v-if="isLoadingSponsors" class="state-box" style="text-align:center; padding: 16px; color: #6b7280; font-size: 12px;">
              Loading...
            </div>
            <div v-else-if="sponsorsError" class="state-box error" style="text-align:center; padding: 16px; color: #dc3545; font-size: 12px;">
              Can't load sponsors
              <button class="all-btn" @click="fetchSponsors">Try again</button>
            </div>
            <div v-else-if="sponsors.length === 0" class="state-box" style="text-align:center; padding: 16px; color: #6b7280; font-size: 12px;">
              No sponsors yet
            </div>

            <div v-else class="sponsor-profile-grid">
              <div
                v-for="sponsor in sponsors"
                :key="sponsor.id"
                class="sponsor-profile-item"
                @click="handleSponsorReadMore(sponsor)"
              >
                <div class="sponsor-avatar-wrap">
                  <img
                    :src="sponsor.logo || sponsor.image"
                    :alt="sponsor.title"
                    class="sponsor-avatar"
                    loading="lazy"
                  />
                  <span v-if="sponsor.isVerified" class="sponsor-active-dot"></span>
                </div>

                <span class="sponsor-profile-name">
                  {{ sponsor.title }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

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
  border: 2px solid transparent;
  transition: border-color 0.2s ease;
}

.sidebar-user-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.sidebar-user-avatar.is-following {
  border-color: #1B75D2;
}

.follow-check-badge {
  position: absolute;
  bottom: -2px;
  right: -2px;
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: #1B75D2;
  border: 2px solid #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  z-index: 2;
}

.user-add-btn.is-following {
  background: transparent;
  color: #1B75D2;
  border: 1.5px solid #1B75D2;
  padding: 3px 10px;     
}


.user-add-btn.is-following:hover {
  background: transparent;
}

.user-add-btn.is-following:hover {
  opacity: 0.8;
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
  border-radius: 32px;
  padding: 4px 8px;
  font-size: 12px;
  font-weight: 400;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;              /* <-- បន្ថែម gap រវាង icon និង text */
  transition: all 0.2s;
  flex-shrink: 0;
}

.user-add-btn:hover {

  background: #155bb5;
}

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
  background-color: #ffffff;
}

.chat-card {
  background-color: #ffffff;
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
  background: #1B75D2;
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
  color: #000;
  font-size: 11px;
  font-weight: 700;
  text-align: center;
}

.row-section-box {
  background: #ffffff;
  border-left: 1px solid #e2e8f0;
  border-right: 1px solid #e2e8f0;
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
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
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

.sponsors-sidebar-card {
  width: 100%;
  background: #fff;
  border: 1px solid #e8edf3;
  border-radius: 12px;
  padding: 16px;
  box-sizing: border-box;
}

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

.sponsor-profile-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 16px 10px;
  width: 100%;
}

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

.sponsor-avatar-wrap {
  position: relative;

  width: 32px;
  height: 32px;

  flex-shrink: 0;
}

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

.sponsor-profile-item:hover .sponsor-profile-name {
  color: #1b75d2;
}

.user-add-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.follow-effect-blue.is-following {
  background: transparent;
  color: #1B75D2;
  border: 1.5px solid #1B75D2;
}

.follow-effect-blue:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}
</style>