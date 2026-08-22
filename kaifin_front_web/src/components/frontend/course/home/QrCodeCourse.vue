<template>
  <div class="qr-code-course-page">
    <NavBar />
    
    <div class="page-wrapper">
        <p v-if="isLoading" class="state-msg">កំពុងផ្ទុក...</p>
        <p v-else-if="error" class="state-msg error">{{ error }}</p>
      <div class="qr-container">
        
        <!-- Header Section above QR Code -->
        <div class="qr-top-header">
          <svg class="qr-header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="3" width="7" height="7"></rect>
            <rect x="14" y="3" width="7" height="7"></rect>
            <rect x="14" y="14" width="7" height="7"></rect>
            <rect x="3" y="14" width="7" height="7"></rect>
          </svg>
          <span class="qr-header-title">Scan to Pay</span>
        </div>

        <!-- Top Section: QR Code Box -->
        <div class="qr-box-section">
          <div class="qr-code-wrapper">
            <!-- 100% Real, Scannable QR Code -->
            <img :src="qrCodeUrl" alt="Payment QR Code" class="real-qr-image" />
          </div>
        </div>

        <!-- Bottom Section: Order Details & Summary -->
        <div class="order-details-section">
          <div class="order-header-pill">
            <svg class="order-header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4z"></path>
              <line x1="3" y1="6" x2="21" y2="6"></line>
              <path d="M16 10a4 4 0 0 1-8 0"></path>
            </svg>
            <span>Order Summary</span>
          </div>

          <div class="order-items-list">
            <div v-for="item in cartItems" :key="item.id" class="order-item-row">
              <div class="item-thumb-box">
                <img v-if="item.thumbnail" :src="item.thumbnail" alt="Course Thumbnail" class="item-thumb-img" />
                <svg v-else class="item-thumb-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polygon points="5 3 19 12 5 21 5 3"></polygon>
                </svg>
              </div>
              <div class="item-name-pill">
                <span>{{ item.title }}</span>
                <div class="item-sub-info">
                  <div class="instructor-profile-group">
                    <img v-if="item.instructorAvatar" :src="item.instructorAvatar" alt="Instructor" class="instructor-avatar" />
                    <span class="item-instructor">{{ item.instructor }}</span>
                  </div>
                  <span class="item-level">{{ item.level }}</span>
                </div>
              </div>
              <div class="item-pricing-box">
                <span class="current-price">${{ item.currentPrice }}</span>
                <span class="original-price">${{ item.originalPrice }}</span>
              </div>
            </div>
          </div>

          <!-- Calculation Summary Block -->
          <div class="summary-calculation-block">
            <div class="calc-row">
              <div class="calc-label-pill"><span>Original Price:</span></div>
              <div class="calc-value-pill"><span>${{ calculatedOriginalTotal }}</span></div>
            </div>
            <div class="calc-row">
              <div class="calc-label-pill"><span>Discounts:</span></div>
              <div class="calc-value-pill discount-val"><span>-${{ calculatedOriginalTotal - calculatedTotal }}</span></div>
            </div>
            <div class="calc-row total-row">
              <div class="calc-label-pill"><span>Total:</span></div>
              <div class="calc-value-pill"><span>${{ calculatedTotal }}</span></div>
            </div>
          </div>

          <!-- Pay Now Button -->
          <div class="pay-action-wrapper">
            <button class="pay-now-btn" @click="handlePayment">
              <span>Pay Now</span>
            </button>
          </div>

        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import NavBar from '../../navbar/NavBar.vue'

const BASE_URL = import.meta.env.VITE_API_URL
const router = useRouter()

function authHeaders() {
  const token = localStorage.getItem('token')
  return token ? { Authorization: `Bearer ${token}` } : {}
}
function resolveImageUrl(url) {
  if (!url) return ''
  if (/^https?:\/\//i.test(url) || url.startsWith('data:image/')) return url
  const path = url.startsWith('/') ? url : `/uploads/${url}`
  return `${BASE_URL}${path}`
}

const cartItems = ref([])
const isLoading = ref(false)
const error = ref(null)
const isPaying = ref(false)

function mapCartItem(it) {
  return {
    id: it.course_id,
    title: it.title,
    instructor: it.instructor_name || 'Unknown',
    instructorAvatar: resolveImageUrl(it.instructor_avatar),
    thumbnail: resolveImageUrl(it.thumbnail),
    level: it.level_id || 1,
    currentPrice: it.current_price,
    originalPrice: it.original_price,
  }
}

async function fetchCart() {
  isLoading.value = true
  error.value = null
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/cart/show`, {
      headers: { ...authHeaders() },
    })
    const raw = await res.text()
    let data = null
    if (raw) {
      try { data = JSON.parse(raw) } catch {
        throw new Error(`Server returned non-JSON response (status ${res.status})`)
      }
    }
    if (!res.ok) {
      throw new Error(data?.message || `Request failed with status ${res.status}`)
    }
    cartItems.value = (data?.data?.items || []).map(mapCartItem)
  } catch (e) {
    console.error('Failed to fetch cart', e)
    error.value = 'មិនអាចទាញ Cart បានទេ'
  } finally {
    isLoading.value = false
  }
}

onMounted(fetchCart)

const calculatedTotal = computed(() => {
  return cartItems.value.reduce((sum, item) => sum + item.currentPrice, 0)
})

const calculatedOriginalTotal = computed(() => {
  return cartItems.value.reduce((sum, item) => sum + item.originalPrice, 0)
})

const qrCodeUrl = computed(() => {
  const checkoutData = `checkout://pay?amount=${calculatedTotal.value}&items=${cartItems.value.length}`
  return `https://api.qrserver.com/v1/create-qr-code/?size=250x250&data=${encodeURIComponent(checkoutData)}&margin=10`
})

// ⚠️ TODO: replace with a real POST /checkout call once the orders API is ready
const handlePayment = async () => {
  if (isPaying.value || cartItems.value.length === 0) return
  isPaying.value = true
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/course-enrollments/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...authHeaders(),
      },
      body: JSON.stringify({
        course_ids: cartItems.value.map(item => item.id),
      }),
    })
    const raw = await res.text()
    let data = null
    if (raw) {
      try { data = JSON.parse(raw) } catch {
        throw new Error(`Server returned non-JSON response (status ${res.status})`)
      }
    }
    if (!res.ok) {
      throw new Error(data?.message || `Request failed with status ${res.status}`)
    }

    alert(`ការទូទាត់ជោគជ័យ! ចំនួនទឹកប្រាក់: $${calculatedTotal.value}`)
    router.push('/my-learning') // ឬ route ណាមួយសម្រាប់ enrolled courses
  } catch (e) {
    console.error('Payment/enrollment failed', e)
    error.value = 'ការទូទាត់មិនជោគជ័យ សូមព្យាយាមម្តងទៀត'
  } finally {
    isPaying.value = false
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@500;600;700;800&display=swap');

.qr-code-course-page {
  width: 100vw;
  height: 100vh;
  background-color: #F7F4F2;
  font-family: 'Plus Jakarta Sans', sans-serif;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  padding: 0;
}

.page-wrapper {
  width: 100%;
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  box-sizing: border-box;
  overflow: hidden;
}

.qr-container {
  display: flex;
  flex-direction: column;
  width: 749px;
  max-width: 100%;
  height: 100%;
  background-color: #ffffff;
  border: none;
  padding: 16px 24px;
  box-sizing: border-box;
  gap: 12px;
  position: relative;
  box-shadow: none;
  overflow-y: auto;
}

/* QR Top Header */
.qr-top-header {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 2px 0;
}

.qr-header-icon {
  width: 18px;
  height: 18px;
  stroke: #1b75d2;
}

.qr-header-title {
  font-size: 16px;
  font-weight: 800;
  color: #0f172a;
}

.qr-box-section {
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  padding: 2px 0;
  background-color: #ffffff;
}

.qr-code-wrapper {
  width: 180px;
  height: 180px;
  background-color: #ffffff;
  border-radius: 12px;
  padding: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  border: 2px solid #1B75D2; 
}

.real-qr-image {
  width: 100%;
  height: 100%;
  object-fit: contain;
  border-radius: 8px;
}

.order-details-section {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.order-header-pill {
  border: none;
  padding: 4px 0px;
  width: fit-content;
  background-color: transparent;
  display: flex;
  align-items: center;
  gap: 8px;
}

.order-header-icon {
  width: 18px;
  height: 18px;
  stroke: #1b75d2;
}

.order-header-pill span {
  font-size: 16px;
  font-weight: 800;
  color: #0f172a;
}

.order-items-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.order-item-row {
  display: flex;
  align-items: center;
  gap: 12px;
  background-color: #ffffff;
  border: none;
  padding: 4px 0px;
}

.item-thumb-box {
  width: 38px;
  height: 38px;
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f8fafc;
  flex-shrink: 0;
}

.item-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 8px;
}

.item-thumb-icon {
  width: 18px;
  height: 18px;
  stroke: #1b75d2;
}

.item-name-pill {
  flex: 1;
  padding: 2px 4px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.item-name-pill span:first-child {
  font-size: 13px;
  font-weight: 700;
  color: #0f172a;
}

.item-sub-info {
  display: flex;
  gap: 10px;
  align-items: center;
}

.instructor-profile-group {
  display: flex;
  align-items: center;
  gap: 6px;
}

.instructor-avatar {
  width: 16px;
  height: 16px;
  border-radius: 50%;
  object-fit: cover;
}

.item-instructor {
  font-size: 11px;
  font-weight: 600;
  color: #64748b;
}

.item-level {
  font-size: 10px;
  font-weight: 700;
  background-color: #f1f5f9;
  color: #1b75d2;
  padding: 1px 6px;
  border-radius: 4px;
}

.item-pricing-box {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  justify-content: center;
  min-width: 50px;
  padding-right: 2px;
}

.current-price {
  font-size: 13px;
  font-weight: 800;
  color: #0f172a;
}

.original-price {
  font-size: 11px;
  font-weight: 700;
  color: #94a3b8;
  text-decoration: line-through;
}

/* Summary Calculation Block */
.summary-calculation-block {
  display: flex;
  flex-direction: column;
  gap: 6px;
  background-color: transparent;
  padding: 8px 0px;
}

.calc-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.calc-label-pill span {
  font-size: 13px;
  font-weight: 700;
  color: #64748b;
}

.calc-value-pill span {
  font-size: 13px;
  font-weight: 800;
  color: #0f172a;
}

.discount-val span {
  color: #10b981;
}

.total-row {
  border-top: 1px solid #e2e8f0;
  padding-top: 6px;
  margin-top: 2px;
}

.total-row .calc-label-pill span {
  font-size: 14px;
  font-weight: 800;
  color: #0f172a;
}

.total-row .calc-value-pill span {
  font-size: 16px;
  font-weight: 800;
  color: #1b75d2;
}

/* Pay Now Action */
.pay-action-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 2px;
}

.pay-now-btn {
  width: 100%;
  border: none;
  border-radius: 28px;
  padding: 10px 20px;
  background-color: #1b75d2;
  cursor: pointer;
  transition: all 0.2s ease;
  text-align: center;
}

.pay-now-btn span {
  font-size: 15px;
  font-weight: 800;
  color: #ffffff;
}

.pay-now-btn:hover {
  background-color: #155fa0;
}

.pay-now-btn:active {
  transform: scale(0.98);
}
</style>