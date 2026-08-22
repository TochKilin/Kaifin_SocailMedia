<template>
  <div class="shopping-cart-page">
    <NavBar />
    <div class="page-wrapper">
      <p v-if="isLoading" class="state-msg">please wait...</p>
      <p v-else-if="error" class="state-msg error">{{ error }}</p>

      <div v-else class="cart-container">
        <div class="cart-items-section">
          <div class="cart-header">
            <h1 class="cart-title">Shopping Cart</h1>
            <div class="courses-count-badge">
              <svg class="badge-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="9" cy="21" r="1"></circle>
                <circle cx="20" cy="21" r="1"></circle>
                <path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"></path>
              </svg>
              <span>{{ cartItems.length }} Courses Selected</span>
            </div>
          </div>

          <p v-if="!cartItems.length" class="state-msg">Your cart is empty</p>

          <div v-else class="cart-items-list">
            <!-- Dynamic Course Items with Image / Avatar Support -->
            <div v-for="item in cartItems" :key="item.id" class="cart-item-card">
              <div class="item-thumb-box">
                <img v-if="item.thumbnail" :src="item.thumbnail" alt="Course Thumbnail" class="item-thumb-img" />
                <svg v-else class="item-thumb-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polygon points="5 3 19 12 5 21 5 3"></polygon>
                </svg>
              </div>
              <div class="item-details">
                <h3 class="item-title">{{ item.title }}</h3>
                <div class="instructor-row">
                  <div class="instructor-avatar">
                    <img v-if="item.instructorAvatar" :src="item.instructorAvatar" alt="Instructor" class="instructor-avatar-img" />
                    <svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                      <circle cx="12" cy="7" r="4"></circle>
                    </svg>
                  </div>
                  <span class="instructor-name">{{ item.instructor }}</span>
                </div>
                <div class="meta-row">
                  <!-- Custom Enamel Pin Badge Shape-->
                  <div class="pin-badge">
                    <div class="pin-left-circle">
                      <svg class="pin-eagle-icon" viewBox="0 0 24 24" fill="currentColor">
                        <path d="M12 2L9 6H5L3 10L6 14L4 22L12 19L20 22L18 14L21 10L19 6H15L12 2Z" />
                      </svg>
                    </div>
                    <div class="pin-text-content">
                      <span>LEVEL</span>
                      <span class="pin-level-num">{{ item.level }}</span>
                    </div>
                  </div>
                  <div class="rating-box">
                    <span class="rating-label">Ratings</span>
                    <div class="stars">
                      <svg v-for="i in 5" :key="i" viewBox="0 0 24 24" :fill="i <= Math.floor(item.rating) ? '#fbbf24' : 'none'" :stroke="i <= Math.floor(item.rating) ? '#fbbf24' : '#94a3b8'" stroke-width="1.5">
                        <polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon>
                      </svg>
                    </div>
                  </div>
                </div>
              </div>
              <div class="item-pricing">
                <span class="currency">$</span>
                <span class="current-price">{{ item.currentPrice }}</span>
                <span class="original-price">{{ item.originalPrice }}</span>
              </div>
              <button class="remove-item-btn" @click="removeFromCart(item.id)" title="Remove">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <line x1="18" y1="6" x2="6" y2="18"></line>
                  <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
              </button>
            </div>
          </div>
        </div>

        <!-- Right -->
        <div class="cart-summary-section">
          <div class="summary-box-card">
            <div class="summary-title-box">
              <span>Summary total</span>
            </div>

            <div class="summary-calculation-block">
              <div class="calc-row">
                <span class="calc-label">Subtotal</span>
                <span class="calc-value">${{ calculatedOriginalTotal }}</span>
              </div>
              <div class="calc-row discount-row">
                <span class="calc-label">Discount</span>
                <span class="calc-value">-${{ calculatedOriginalTotal - calculatedTotal }}</span>
              </div>
              <div class="divider-line"></div>
              <div class="calc-row total-row">
                <span class="calc-label">Total</span>
                <span class="calc-value">${{ calculatedTotal }}</span>
              </div>
            </div>

            <button class="buying-btn" @click="handleCheckout" :disabled="!cartItems.length">
              <span>Buying</span>
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

function getAuthToken() {
  return localStorage.getItem('token') || ''
}
function authHeaders() {
  const token = getAuthToken()
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
function mapCartItem(it) {
  return {
    id: it.course_id,
    title: it.title,
    instructor: it.instructor_name || 'Unknown',
    instructorAvatar: resolveImageUrl(it.instructor_avatar),
    thumbnail: resolveImageUrl(it.thumbnail),
    level: it.level_id || 1,
    rating: Number(it.rating) || 0,
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

async function removeFromCart(courseId) {
  try {
    const res = await fetch(`${BASE_URL}/api/v1/front/cart/remove/${courseId}`, {
      method: 'DELETE',
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
    cartItems.value = cartItems.value.filter(item => item.id !== courseId)
  } catch (e) {
    console.error('Failed to remove item from cart', e)
    alert(e.message || 'មិនអាចលុប item នេះបានទេ')
  }
}

onMounted(fetchCart)

const calculatedTotal = computed(() => {
  return cartItems.value.reduce((sum, item) => sum + item.currentPrice, 0)
})

const calculatedOriginalTotal = computed(() => {
  return cartItems.value.reduce((sum, item) => sum + item.originalPrice, 0)
})

const discountPercentage = computed(() => {
  if (calculatedOriginalTotal.value === 0) return 0
  const savings = calculatedOriginalTotal.value - calculatedTotal.value
  return Math.round((savings / calculatedOriginalTotal.value) * 100)
})

const handleCheckout = () => {
  router.push({ name: 'QrCodeCourse' })
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@500;600;700;800&display=swap');

.shopping-cart-page {
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

.cart-container {
  display: flex;
  width: 1050px;
  max-width: 100%;
  height: 100%;
  gap: 12px;
  box-sizing: border-box;
  padding: 0;
}

.cart-items-section {
  flex: 1;
  border: 1px solid #e2e8f0;
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  background-color: #ffffff;
  overflow-y: auto;
}

.cart-header {
  display: flex;
  flex-direction: column;
  gap: 10px;
  border-radius: 16px;
  padding: 18px 24px;
}

.cart-title {
  font-size: 28px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}

.courses-count-badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border-radius: 12px;
  padding: 6px 14px;
  font-size: 13px;
  font-weight: 700;
  color: #64748b;
  width: fit-content;
  background: none;
  border: none;
}

.badge-icon {
  width: 16px;
  height: 16px;
  stroke: #64748b;
}

.cart-items-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.cart-item-card {
  border: 1px solid #f1f5f9;
  border-radius: 20px;
  padding: 16px;
  display: flex;
  align-items: center;
  gap: 16px;
  background-color: #ffffff;
  transition: all 0.2s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.02);
}

.cart-item-card:hover {
  border-color: #cbd5e1;
}

.item-thumb-box {
  width: 80px;
  height: 80px;
  border: 1px solid #e2e8f0;
  background-color: #f8fafc;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  overflow: hidden;
}

.item-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.item-thumb-icon {
  width: 36px;
  height: 36px;
  stroke: #1b75d2;
}

.item-details {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.item-title {
  font-size: 16px;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
  border-radius: 10px;
  padding: 6px 12px;
  width: fit-content;
}

.instructor-row {
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #f1f5f9;
  background-color: #f8fafc;
  border-radius: 10px;
  padding: 4px 10px;
  width: fit-content;
}

.instructor-avatar {
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  overflow: hidden;
}

.instructor-avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.instructor-avatar svg {
  width: 16px;
  height: 16px;
  stroke: #64748b;
}

.instructor-name {
  font-size: 12px;
  font-weight: 700;
  color: #64748b;
}

.meta-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.pin-badge {
  display: inline-flex;
  align-items: center;
  background-color: #ffffff;
  border: 2px solid #1B75D2;
  border-radius: 20px;
  padding: 2px 10px 2px 2px;
  box-shadow: 0 2px 5px rgba(27, 117, 210, 0.2);
  gap: 6px;
}

.pin-left-circle {
  width: 24px;
  height: 24px;
  background-color: #ffffff;
  border: 1.5px solid #1B75D2;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.pin-eagle-icon {
  width: 14px;
  height: 14px;
  color: #1B75D2;
}

.pin-text-content {
  display: flex;
  flex-direction: column;
  line-height: 1;
}

.pin-text-content span:first-child {
  font-size: 7px;
  font-weight: 800;
  color: #1B75D2;
  letter-spacing: 0.5px;
}

.pin-level-num {
  font-size: 11px;
  font-weight: 800;
  color: #0f172a;
}

.rating-box {
  display: flex;
  align-items: center;
  gap: 6px;
  border-radius: 8px;
  padding: 2px 8px;
}

.rating-label {
  font-size: 10px;
  font-weight: 700;
  color: #64748b;
}

.stars {
  display: flex;
  gap: 2px;
}

.stars svg {
  width: 12px;
  height: 12px;
}

.item-pricing {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #F64242;
  padding: 8px 12px;
  min-width: 25px;
}

.currency {
  font-size: 14px;
  font-weight: 700;
  color: #ffffff;
}

.current-price {
  font-size: 22px;
  font-weight: 800;
  color: #ffffff;
}

.original-price {
  font-size: 13px;
  font-weight: 700;
  color: #94a3b8;
  text-decoration: line-through;
  margin-top: 2px;
}

.cart-summary-section {
  width: 320px;
  border: 1px solid #e2e8f0;
  padding: 12px;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.03);
  box-sizing: border-box;
}

.summary-box-card {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: stretch;
  gap: 10px;
  height: 100%;
}

.summary-title-box {
  border-radius: 12px;
  padding: 8px 12px;
  text-align: center;
  font-size: 18px;
  font-weight: 800;
  color: #64748b;
}

.summary-calculation-block {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding: 16px;
  border-radius: 12px;
  background-color: #f8fafc;
}

.calc-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.calc-label {
  font-size: 16px;
  font-weight: 700;
  color: #64748b;
}

.calc-value {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}

.discount-row .calc-value {
  color: #10b981;
}

.total-row .calc-label {
  font-size: 18px;
  font-weight: 800;
  color: #0f172a;
}

.total-row .calc-value {
  font-size: 24px;
  font-weight: 800;
}

.divider-line {
  height: 1px;
  background-color: #e2e8f0;
  width: 100%;
}

.buying-btn {
  margin-top: 10px;
  border: none;
  border-radius: 32px;
  padding: 8px 12px;
  background-color: #1b75d2;
  cursor: pointer;
  transition: all 0.2s ease;
  text-align: center;
  box-shadow: none;
}

.buying-btn span {
  font-size: 18px;
  font-weight: 800;
  color: #ffffff;
}

.buying-btn:hover {
  background-color: #155fa0;
  box-shadow: none;
}

.buying-btn:active {
  transform: scale(0.98);
}

.remove-item-btn {
  background: transparent;
  border: none;
  color: #94a3b8;
  cursor: pointer;
  padding: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.2s;
}
.remove-item-btn:hover {
  color: #dc2626;
}
.remove-item-btn svg {
  width: 18px;
  height: 18px;
}
</style>