<template>
  <div class="modal-backdrop">
    <div class="modal-content">
      <!-- Close Button -->
      <button class="close-btn" @click="$emit('close')">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <line x1="18" y1="6" x2="6" y2="18"></line>
          <line x1="6" y1="6" x2="18" y2="18"></line>
        </svg>
      </button>

      <!-- Bank Selection Dropdown -->
      <div class="dropdown-group">
        <div class="select-dropdown" @click="toggleDropdown">
          <div class="bank-info">
            <div class="bank-logo-placeholder">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="2" y="5" width="20" height="14" rx="2" />
                <line x1="2" y1="10" x2="22" y2="10" />
              </svg>
            </div>
            <span class="selected-text">{{ selectedBank }}</span>
          </div>
          <span class="dropdown-arrow">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="6 9 12 15 18 9"></polyline>
            </svg>
          </span>
        </div>
        <div class="dropdown-menu" v-if="isOpen">
          <div 
            v-for="bank in bankOptions" 
            :key="bank" 
            class="dropdown-item"
            @click="selectBank(bank)"
          >
            {{ bank }}
          </div>
        </div>
      </div>

      <!-- Summary Wrapper with Floating Title -->
      <div class="summary-wrapper">
        <div class="summary-title-badge">Summary</div>

        <div class="summary-card">
          <div class="summary-row">
            <span class="summary-label">Total due today:</span>
            <span class="summary-value">$100.00</span>
          </div>
          <div class="summary-row">
            <span class="summary-label">Starting on May 8, 2027:</span>
            <span class="summary-value">$120.00</span>
          </div>
          <!-- Promotion Banner inside Summary Container -->
          <div class="promo-row">
            Promotions 20% off for the first year
          </div>
        </div>
      </div>

      <!-- Pay Now Button -->
      <button class="pay-btn" @click="$emit('pay')">
        Pay Now
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const isOpen = ref(false)
const selectedBank = ref('EON Bank')
const bankOptions = ref(['EON Bank', 'ABA Bank', 'ACLEDA Bank', 'KB PRASAC Bank', 'KHQR'])

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

const selectBank = (bank) => {
  selectedBank.value = bank
  isOpen.value = false
}
</script>

<style scoped>
.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(15, 23, 42, 0.4);
  backdrop-filter: blur(5px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.modal-content {
  background-color: #ffffff; /* ពណ៌សស្អាត */
  width: 100%;
  max-width: 440px;
  border-radius: 28px;
  padding: 32px 24px 28px 24px;
  position: relative;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.15);
  box-sizing: border-box;
  border: 1.5px solid #e2e8f0;
}

/* Close Button */
.close-btn {
  position: absolute;
  top: 18px;
  right: 18px;
  background: #f1f5f9;
  border: none;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #64748b;
  cursor: pointer;
  transition: all 0.2s;
}

.close-btn:hover {
  background: #e2e8f0;
  color: #0f172a;
}

/* Dropdown */
.dropdown-group {
  position: relative;
  margin-top: 15px;
  margin-bottom: 24px;
}

.select-dropdown {
  background: #ffffff;
  border: 2px solid #38bdf8;
  border-radius: 16px;
  padding: 14px 18px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
  box-shadow: 0 4px 12px rgba(56, 189, 248, 0.1);
}

.bank-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.bank-logo-placeholder {
  width: 28px;
  height: 28px;
  background: #f0f9ff;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #0284c7;
}

.selected-text {
  font-size: 18px;
  font-weight: 700;
  color: #0f172a;
}

.dropdown-arrow {
  color: #0284c7;
  display: flex;
  align-items: center;
}

.dropdown-menu {
  position: absolute;
  top: calc(100% + 8px);
  left: 0;
  width: 100%;
  background-color: #ffffff;
  border: 2px solid #38bdf8;
  border-radius: 16px;
  overflow: hidden;
  z-index: 10;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1);
}

.dropdown-item {
  padding: 14px 18px;
  font-size: 15px;
  font-weight: 500;
  color: #334155;
  cursor: pointer;
  transition: background 0.15s;
}

.dropdown-item:hover {
  background-color: #f0f9ff;
  color: #0284c7;
}

/* Summary Wrapper */
.summary-wrapper {
  position: relative;
  margin-top: 28px;
  margin-bottom: 24px;
}

.summary-title-badge {
  position: absolute;
  top: -14px;
  left: 20px;
  background: #0f172a;
  color: #ffffff;
  font-size: 13px;
  font-weight: 700;
  padding: 4px 14px;
  border-radius: 20px;
  letter-spacing: 0.3px;
  box-shadow: 0 2px 6px rgba(15, 23, 42, 0.2);
  z-index: 2;
}

.summary-card {
  background: #f8fafc;
  border: 2px solid #38bdf8;
  border-radius: 20px;
  padding: 24px 18px 18px 18px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.summary-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #ffffff;
  padding: 12px 14px;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
}

.summary-label {
  font-size: 13px;
  font-weight: 600;
  color: #64748b;
}

.summary-value {
  font-size: 17px;
  font-weight: 800;
  color: #0f172a;
}

/* Promotion Row inside container */
.promo-row {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 12px 14px;
  text-align: center;
  font-size: 13px;
  font-weight: 700;
  color: #0284c7;
}

/* Pay Now Button */
.pay-btn {
  width: 100%;
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  color: #ffffff;
  border: none;
  border-radius: 16px;
  padding: 16px;
  font-size: 18px;
  font-weight: 700;
  cursor: pointer;
  box-shadow: 0 6px 20px rgba(234, 88, 12, 0.35);
  transition: transform 0.1s, opacity 0.2s;
}

.pay-btn:hover {
  opacity: 0.95;
}

.pay-btn:active {
  transform: scale(0.98);
}
</style>