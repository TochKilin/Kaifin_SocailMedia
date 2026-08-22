<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";

const router = useRouter();
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

const TABS = [
  { 
    key: "everyone", 
    label: "Everyone",
    icon: `<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>` 
  },
  { 
    key: "online", 
    label: "Online",
    icon: `<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>` 
  },
  { 
    key: "requests", 
    label: "Requests",
    icon: `<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="8.5" cy="7" r="4"></circle><line x1="20" y1="8" x2="20" y2="14"></line><line x1="23" y1="11" x2="17" y2="11"></line></svg>` 
  },
  { 
    key: "suggested", 
    label: "People you may know",
    icon: `<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="8.5" cy="7" r="4"></circle><line x1="20" y1="8" x2="20" y2="14"></line></svg>` 
  },
  { 
    key: "more", 
    label: "More",
    icon: `` 
  },
];
const activeTab = ref("more");
const searchQuery = ref("");
const invisibleModeEnabled = ref(false);
const isTogglingInvisible = ref(false);

async function toggleInvisibleMode() {
  if (isTogglingInvisible.value) return;
  isTogglingInvisible.value = true;
  try {
    const res = await axios.post(
      `${BASE_URL}/api/v1/front/settings/invisible-mode`,
      { enabled: !invisibleModeEnabled.value },
      { headers: authHeaders() }
    );
    const data = res.data?.data ?? res.data;
    invisibleModeEnabled.value = data?.enabled ?? !invisibleModeEnabled.value;
  } catch (err) {
    console.error("Failed to toggle invisible mode:", err);
    alert("Failed to update invisible mode");
  } finally {
    isTogglingInvisible.value = false;
  }
}

function openSchoolFriends() {
  router.push({ name: "SchoolFriends" });
}

function openImportVk() {
  alert("🔗 Opening Kaifin import...");
}

async function shareProfile() {
  const url = `${window.location.origin}/profile/me`;
  try {
    await navigator.clipboard.writeText(url);
    alert("🔗 Profile link copied to clipboard!");
  } catch (err) {
    console.error("Clipboard write failed:", err);
    alert(url);
  }
}

const sentRequests = ref([]);
const isLoadingRequests = ref(true);
const requestsError = ref("");

const MOCK_SENT_REQUESTS = [
  {
    id: 1,
    username: "Sokha Chan",
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTSkXGSq73H3pWLk34gkNyfJTcaf2F7rbcVe-l3ot3S-w&s",
    isOnline: true,
    isCancelling: false,
  },
  {
    id: 2,
    username: "Dara Pich",
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRHQPxY-hq9Az2qOCVRcptkosJ7p3a4mV34f1z9-nmHIQ&s=10",
    isOnline: false,
    isCancelling: false,
  },
  {
    id: 3,
    username: "Sreymom Lay",
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ09JjRcIesU-6Z5o6M5MZL31mlyZ0YQRWD2OMMHW_0QA&s=10",
    isOnline: true,
    isCancelling: false,
  },
  {
    id: 4,
    username: "Vibol Heng",
    avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTMUj15ZfwZ9MSB_YPqQBnY5SJoFLdL7u0HghjrFk2-Jg&s=10",
    isOnline: false,
    isCancelling: false,
  },
];

const USE_MOCK = true;

const fetchSentRequests = async () => {
  isLoadingRequests.value = true;
  requestsError.value = "";

  if (USE_MOCK) {
    await new Promise((resolve) => setTimeout(resolve, 500));
    sentRequests.value = MOCK_SENT_REQUESTS;
    isLoadingRequests.value = false;
    return;
  }

  try {
    const res = await axios.get(`${BASE_URL}/api/v1/front/friends/sent-requests`, {
      params: { page: 1, limit: 50 },
      headers: authHeaders(),
    });
    const data = res.data?.data ?? res.data;
    const list = data?.requests ?? data ?? [];

    sentRequests.value = list.map((r) => ({
      id: r.id ?? r.user_id,
      username: r.user_name || [r.first_name, r.last_name].filter(Boolean).join(" ") || `User #${r.user_id}`,
      avatar: resolveAvatarUrl(r.profile_images) || `https://api.dicebear.com/7.x/avataaars/svg?seed=${r.user_id}`,
      isOnline: r.is_online ?? false,
      isCancelling: false,
    }));
  } catch (err) {
    console.error("Failed to fetch sent requests:", err);
    requestsError.value = err.message || "Failed to load sent requests";
  } finally {
    isLoadingRequests.value = false;
  }
};

async function cancelRequest(request) {
  if (request.isCancelling) return;
  request.isCancelling = true;

  if (USE_MOCK) {
    await new Promise((resolve) => setTimeout(resolve, 400));
    sentRequests.value = sentRequests.value.filter((r) => r.id !== request.id);
    return;
  }

  try {
    await axios.post(
      `${BASE_URL}/api/v1/front/friends/cancel-request`,
      { user_id: request.id },
      { headers: authHeaders() }
    );
    sentRequests.value = sentRequests.value.filter((r) => r.id !== request.id);
  } catch (err) {
    console.error("Failed to cancel request:", err);
    alert("Failed to cancel request");
    request.isCancelling = false;
  }
}

const filteredRequests = computed(() => {
  const q = searchQuery.value.trim().toLowerCase();
  if (!q) return sentRequests.value;
  return sentRequests.value.filter((r) => r.username.toLowerCase().includes(q));
});

onMounted(() => {
  fetchSentRequests();
});
</script>

<template>
  <div class="page-container">
    <div class="content-wrap">
      <!-- TABS BAR -->
      <div class="tabs-search-card card">
        <div class="tabs-row">
          <button
            v-for="tab in TABS"
            :key="tab.key"
            class="tab-btn"
            :class="{ active: activeTab === tab.key }"
            @click="activeTab = tab.key"
          >
            <span class="tab-icon" v-html="tab.icon"></span>
            <span>{{ tab.label }}</span>
            <span v-if="tab.key === activeTab" class="tab-underline"></span>
          </button>
        </div>
      </div>

      <!-- MORE PANEL -->
      <div v-if="activeTab === 'more'" class="more-panel-card more-panel-card-raduis card">
        <div class="more-cards-grid">
          <div class="more-card school-card" @click="openSchoolFriends">
            <div class="more-card-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M22 10v6M2 10l10-5 10 5-10 5-10-5z"></path>
                <path d="M6 12v5c3 3 9 3 12 0v-5"></path>
              </svg>
            </div>
            <h3>School friends</h3>
            <p>Search for your school acquaintances and friends</p>
          </div>

          <div class="more-card vk-card" @click="openImportVk">
            <div class="more-card-icon">
              <img src="@/assets/logos/kaifin_l2.png" alt="Kaifin" width="24" height="24" style="border-radius: 50%; object-fit: cover;" />
            </div>
            <h3>Import from Kaifin</h3>
            <p>Add friends who are already on Kaifin</p>
          </div>

          <div class="more-card share-card" @click="shareProfile">
            <div class="more-card-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="18" cy="5" r="3"></circle>
                <circle cx="6" cy="12" r="3"></circle>
                <circle cx="18" cy="19" r="3"></circle>
                <line x1="8.59" y1="13.51" x2="15.42" y2="17.49"></line>
                <line x1="15.41" y1="6.51" x2="8.59" y2="10.49"></line>
              </svg>
            </div>
            <h3>Share profile</h3>
            <p>Send a link to friends on other social networks</p>
          </div>

          <div class="more-card invisible-card">
            <button
              class="enable-btn"
              :class="{ enabled: invisibleModeEnabled }"
              :disabled="isTogglingInvisible"
              @click="toggleInvisibleMode"
            >
              {{ invisibleModeEnabled ? 'Enabled' : 'Enable' }}
            </button>
            <div class="more-card-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M12 2a5 5 0 0 0-5 5v3H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2h-2V7a5 5 0 0 0-5-5z"></path>
              </svg>
            </div>
            <h3>Invisible mode</h3>
            <p>Visit anyone you like invisibly</p>
          </div>
        </div>
      </div>

      <!-- sent -->
      <div class="sent-requests-card card">
        <div class="sent-requests-header">
          <h2 class="sent-requests-title">Sent requests</h2>
        </div>

        <div v-if="isLoadingRequests" class="skeleton-requests">
          <div v-for="n in 3" :key="'skeleton-req-' + n" class="skeleton-request-row">
            <div class="skeleton-avatar"></div>
            <div class="skeleton-line w-40"></div>
          </div>
        </div>

        <div v-else-if="requestsError" class="state-box error">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="12"></line><line x1="12" y1="16" x2="12.01" y2="16"></line></svg>
          <p>Can't load sent requests — {{ requestsError }}</p>
          <button class="all-btn" @click="fetchSentRequests">Try again</button>
        </div>

        <div v-else-if="filteredRequests.length === 0" class="state-box">
          <svg width="36" height="36" viewBox="0 0 24 24" fill="none" stroke="#9ca3af" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path><polyline points="22 4 12 14.01 9 11.01"></polyline></svg>
          <p>No sent requests</p>
        </div>

        <div v-else class="requests-list">
          <div v-for="request in filteredRequests" :key="request.id" class="request-row">
            <div class="request-avatar">
              <img :src="request.avatar" :alt="request.username" loading="lazy" />
            </div>

            <span class="request-name">{{ request.username }}</span>

            <button
              class="cancel-btn"
              :disabled="request.isCancelling"
              @click="cancelRequest(request)"
            >
              {{ request.isCancelling ? 'Cancelling...' : 'Cancel request' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap');

.page-container {
  width: 100%;
  min-height: 100vh;
  font-family: 'Plus Jakarta Sans', system-ui, -apple-system, sans-serif;
  color: #1f2937;
}

.content-wrap {
  width: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 12px;
}

.tabs-search-card {
  background: #ffffff;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
  padding: 12px 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  flex-wrap: wrap;
}

.tabs-row {
  display: flex;
  align-items: center;
  gap: 20px;
  flex-wrap: wrap;
  width: 100%;
}

.tab-btn {
  position: relative;
  background: transparent;
  border: none;
  font-family: inherit;
  font-size: 15px;
  font-weight: 500;
  color: #9ca3af;
  cursor: pointer;
  transition: color 0.2s ease;
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 0;
}

.tab-icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.tab-btn.active {
  color: #111827;
  font-weight: 700;
}

.tab-btn.active .tab-icon {
  color: #1B75D2;
}

.tab-underline {
  position: absolute;
  left: 0;
  right: 0;
  bottom: -2px;
  height: 3px;
  border-radius: 3px;
  background: #1B75D2;
}

/* ===== MORE PANEL ===== */
.more-panel-card {
  background: #ffffff;
  padding: 20px;
}

.more-panel-card-raduis {
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.more-cards-grid {
  display: grid;
  /* grid-template-columns: repeat(auto-fit, minmax(220px, 1fr)); */
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.more-card {
  position: relative;
  border-radius: 12px;
  padding: 12px;
  cursor: pointer;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.more-card:hover {
  opacity: 0.8;
}

.school-card { background-color: #ef44441f; }
.vk-card { background-color: #1b77d23b;}
.share-card { background-color: #b453092a;}
.invisible-card { cursor: default; background-color: #7c3aed36; }

.more-card-icon {
  width: 28px;
  height: 28px;
  border-radius: 4px;
  background: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 14px;
}

.school-card .more-card-icon { color: #ef4444; }
.vk-card .more-card-icon { color: #2563eb; }
.share-card .more-card-icon { color: #b45309; }
.invisible-card .more-card-icon { color: #7c3aed; }

.more-card h3 {
  font-size: 12px;
  font-weight: 700;
  color: #111827;
  margin: 0 0 6px 0;
}

.more-card p {
  font-size: 12px;
  color: #6b7280;
  margin: 0;
  line-height: 1.4;
}

.enable-btn {
  position: absolute;
  top: 16px;
  right: 16px;
  background: #00000026;
  border: none;
  border-radius: 32px;
  padding: 4px 6px;
  font-family: inherit;
  font-size: 11px;
  font-weight: 700;
  color: #000;
  cursor: pointer;
  transition: all 0.2s ease;
}

.enable-btn:hover {
  background: #f3f4f6;
}

.enable-btn.enabled {
  background: #1B75D2;
  color: #ffffff;
}

.enable-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.sent-requests-card {
  background: #ffffff;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
  padding: 24px;
}

.sent-requests-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 20px;
  color: #1B75D2;
}

.sent-requests-title {
  font-size: 16px;
  font-weight: 700;
  color: #111827;
  margin: 0;
}

.requests-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.request-row {
  display: flex;
  align-items: center;
  gap: 16px;
}

.request-avatar {
  position: relative;
  width: 94px;
  height: 94px;
  border-radius: 8px;
  overflow: hidden;
  background: #e2e8f0;
  flex-shrink: 0;
}

.request-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.request-name {
  flex: 1;
  font-size: 15px;
  font-weight: 700;
  color: #111827;
}

.cancel-btn {
  background: #f4f1ec;
  border: none;
  border-radius: 32px;
  padding: 7px 12px;
  font-family: inherit;
  font-size: 12px;
  font-weight: 600;
  color: #374151;
  cursor: pointer;
  transition: all 0.2s ease;
  white-space: nowrap;
}

.cancel-btn:hover {
  background: #e5e1d8;
}

.cancel-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.state-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  text-align: center;
  padding: 30px;
  color: #6b7280;
  font-size: 13px;
}

.state-box.error {
  color: #dc3545;
}

.all-btn {
  background: transparent;
  border: none;
  color: #1B75D2;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  padding: 6px 12px;
}

.skeleton-requests {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.skeleton-request-row {
  display: flex;
  align-items: center;
  gap: 16px;
}

.skeleton-avatar {
  width: 52px;
  height: 52px;
  border-radius: 12px;
  background: #e2e8f0;
  flex-shrink: 0;
}

.skeleton-line {
  height: 12px;
  background: #e2e8f0;
  border-radius: 6px;
}

.w-40 { width: 40%; }
</style>