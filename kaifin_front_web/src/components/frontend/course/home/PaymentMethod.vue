<template>
  <div>
    <NavBar/>
    <div class="page-wrapper">
      
      <!-- ================= BANK MODAL FORM ================= -->
      <div class="modal-overlay" v-if="isModalOpen" @click.self="closeBankModal">
        <div class="modal-box">

          <!-- Thumbnail & Course Description inside Modal (Added here) -->
          <div class="modal-course-highlight-box">
            <div class="modal-course-thumb-wrapper">
              <img :src="bannerData.thumbnail" :alt="bannerData.title" class="modal-course-thumb-img" />
            </div>
            <div class="modal-course-text-content">
              Unlock unlimited access to <span class="course-bold-title">The Complete Claude Code & Claude Cowork Masterclass [2026]</span>, along with 28,000+ other courses with Personal Plan Subscription today!
            </div>
          </div>

          <!-- Type of Bank Dropdown -->
          <div class="modal-form-group">
            <label class="modal-field-label">Select Bank</label>
            <div class="select-dropdown modal-bank-select" @click="toggleModalBankDropdown">
              <div class="select-left">
                <div class="selected-bank-display" v-if="selectedBankObj">
                  <img :src="selectedBankObj.icon" class="bank-option-logo" alt="" />
                  <span class="selected-country active-bank-text">{{ selectedBankObj.name }}</span>
                </div>
                <span class="selected-country placeholder-text" v-else>Choose your bank (ABA, Acleda, KHQR...)</span>
              </div>
              <span class="dropdown-arrow">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="6 9 12 15 18 9"></polyline>
                </svg>
              </span>
            </div>

            <!-- Dropdown List inside Modal -->
            <div class="dropdown-menu-list modal-dropdown-list" v-if="isModalBankOpen">
              <div 
                v-for="bank in bankList" 
                :key="bank.name" 
                class="dropdown-item bank-dropdown-item"
                @click.stop="selectModalBank(bank)"
              >
                <img :src="bank.icon" class="bank-item-logo" :alt="bank.name" />
                <span class="bank-item-text">{{ bank.name }}</span>
              </div>
            </div>
          </div>

          <!-- Modal Plan Card with Border -->
          <div class="modal-plan-card">
            <div class="modal-plan-info">
              <span class="modal-plan-title">{{ selectedPlan === 'yearly' ? 'Yearly Access' : 'Monthly Access' }}</span>
              <span class="modal-plan-price">{{ selectedPlan === 'yearly' ? '$100.00/year' : '$8.00/mon' }}</span>
            </div>
            <div class="modal-plan-desc">
              {{ selectedPlan === 'yearly' ? 'Access to 28,000+ top-rated courses (Billed annually)' : 'Access to 28,000+ top-rated courses (Billed monthly)' }}
            </div>
          </div>

          <!-- Summary Section with Border Box -->
          <div class="modal-summary-section">
            <label class="summary-label-title">Summary</label>
            
            <div class="summary-row-box">
              <span class="summary-label">Total due today:</span>
              <span class="summary-value">$100.00</span>
            </div>

            <div class="summary-row-box">
              <span class="summary-label">Starting on May 8, 2027:</span>
              <span class="summary-value">$120.00</span>
            </div>

            <div class="promo-box">
              Promotions 20% off for the first year
            </div>
          </div>

          <!-- Pay Now Button (Short & Centered) -->
          <button class="pay-now-btn" @click="handlePayNow">
            Pay Now
          </button>

        </div>
      </div>

      <div class="payment-container">
        
        <!-- Start Learning with Kaifin Label Header -->
        <div class="demo-header-label">
          Start Learning with <span class="highlight-name">Kaifin</span>
        </div>

        <!-- Top Banner Box (Removed thumbnail from here) -->
        <div class="access-banner">
          <span class="access-text">{{ bannerData.title }}</span>
        </div>

        <!-- Subscription Plan Options -->
        <div class="plans-row">
          <div 
            class="plan-card" 
            :class="{ active: selectedPlan === 'yearly' }"
            @click="selectedPlan = 'yearly'"
          >
            <div class="plan-header">
              <div class="radio-circle" :class="{ checked: selectedPlan === 'yearly' }"></div>
              <span class="plan-title">Yearly Access</span>
            </div>
            <div class="plan-pricing">
              <span class="plan-rate">$8.00/mon</span>
              <span class="save-badge">save $200.00</span>
            </div>
            <div class="plan-subprice">
              $100.00/year <span class="old-price">$300.00 for first year</span>
            </div>
          </div>

          <div 
            class="plan-card" 
            :class="{ active: selectedPlan === 'monthly' }"
            @click="selectedPlan = 'monthly'"
          >
            <div class="plan-header">
              <div class="radio-circle" :class="{ checked: selectedPlan === 'monthly' }"></div>
              <span class="plan-title">Monthly Access</span>
            </div>
            <div class="plan-pricing">
              <span class="plan-rate">$8.00/mon</span>
            </div>
            <div class="plan-subprice billed-monthly">
              Billed monthly
            </div>
          </div>
        </div>

        <!-- What's Included Section -->
        <div class="included-section">
          <h3 class="included-title">What's included</h3>
          <div class="benefit-item" v-for="(benefit, index) in benefits" :key="index">
            <div class="check-box-icon">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="20 6 9 17 4 12"></polyline>
              </svg>
            </div>
            <span class="benefit-text">{{ benefit }}</span>
          </div>
        </div>

        <!-- Country Selection -->
        <div class="form-group">
          <label class="form-label">Country</label>
          <div class="select-dropdown" @click="toggleCountryDropdown">
            <div class="select-left">
              <span class="selected-country">{{ selectedCountry }}</span>
            </div>
            <span class="dropdown-arrow">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="18 15 12 9 6 15"></polyline>
              </svg>
            </span>
          </div>
          <div class="dropdown-menu-list" v-if="isCountryOpen">
            <div 
              v-for="country in countries" 
              :key="country" 
              class="dropdown-item"
              @click.stop="selectCountry(country)"
            >
              {{ country }}
            </div>
          </div>
        </div>

        <!-- Payment Method Section -->
        <div class="form-group">
          <label class="form-label">Payment Method</label>
          
          <div 
            class="payment-option" 
            :class="{ 'active-option': selectedPayment === 'cards' }"
            @click="selectedPayment = 'cards'"
          >
            <div class="payment-left">
              <div class="radio-circle" :class="{ checked: selectedPayment === 'cards' }"></div>
              <svg class="payment-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="1" y="4" width="22" height="16" rx="2" ry="2"></rect>
                <line x1="1" y1="10" x2="23" y2="10"></line>
              </svg>
              <span class="payment-name">Cards</span>
            </div>
            <div class="card-logos">
              <template v-for="(item, index) in cardIcons" :key="index">
                <span 
                  class="brand-badge" 
                  :class="[item.class, { 'icon-badge-wrapper': item.type === 'image' }]"
                  :style="item.bg ? { background: item.bg, color: item.color } : {}"
                >
                  <img v-if="item.type === 'image'" :src="item.url" :alt="item.alt" class="badge-img" />
                  <span v-else>{{ item.value }}</span>
                </span>
              </template>
            </div>
          </div>

          <div class="bank-dropdown-container">
            <div 
              class="payment-option" 
              :class="{ 'active-option': selectedPayment === 'bank' }"
              @click="openBankModal"
            >
              <div class="payment-left">
                <div class="radio-circle" :class="{ checked: selectedPayment === 'bank' }"></div>
                <svg class="payment-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M3 21h18M3 10h18M5 6l7-3 7 3M4 10v11M20 10v11M8 14v3M12 14v3M16 14v3"></path>
                </svg>
                <div class="selected-bank-display" v-if="selectedBankObj">
                  <img :src="selectedBankObj.icon" class="bank-option-logo" alt="" />
                  <span class="payment-name">{{ selectedBankObj.name }}</span>
                </div>
                <span class="payment-name" v-else>Bank</span>
              </div>

              <div class="bank-logos">
                <template v-for="(item, index) in bankIcons" :key="index">
                  <span 
                    class="bank-badge" 
                    :class="[item.class, { 'icon-badge-wrapper': item.type === 'image' }]"
                    :style="item.bg ? { background: item.bg, color: item.color } : {}"
                  >
                    <img v-if="item.type === 'image'" :src="item.url" :alt="item.alt" class="badge-img" />
                    <span v-else>{{ item.value }}</span>
                  </span>
                </template>
              </div>
            </div>
          </div>

        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import NavBar from '../../navbar/NavBar.vue'

import chipIcon from '@/assets/bank_card/visa.png'
import crownIcon from '@/assets/bank_card/kb.png'

import abaLogo from '@/assets/bank_card/kb.png'
import acledaLogo from '@/assets/bank_card/wink.png'
import kbLogo from '@/assets/bank_card/aba.png'
import khqrLogo from '@/assets/bank_card/kh.png'
import Eon from '@/assets/bank_card/eon.png' 
import Ac from '@/assets/bank_card/ac.png' 
import Amk from '@/assets/bank_card/amk.png'

const bannerData = ref({
  title: 'Access to 28,000+ top-rated courses',
  thumbnail: 'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=600&auto=format&fit=crop&q=60'
})

const selectedPlan = ref('yearly')
const selectedPayment = ref('cards')

const benefits = ref([
  'Access lessons anytime and learn at a comfortable pace that works for you.',
  'Gain useful skills through real-world examples, exercises, and hands-on projects.',
  'Learn from experienced instructors with clear and easy-to-follow lessons.',
  'Keep track of your learning progress and see how far you have come.',
  'Complete your courses and earn certificates to showcase your new skills.',
  'Access your courses from anywhere and continue learning whenever you have time.',
])

const selectedCountry = ref('Cambodia')
const isCountryOpen = ref(false)
const countries = ref(['Cambodia', 'Thailand', 'United States', 'Singapore', 'Vietnam'])

const isModalOpen = ref(false)
const isModalBankOpen = ref(false)
const selectedBankObj = ref(null)

const bankList = ref([
  { name: 'ABA Bank', icon: abaLogo },
  { name: 'ACLEDA Bank', icon: acledaLogo },
  { name: 'KB PRASAC Bank', icon: kbLogo },
  { name: 'KHQR', icon: khqrLogo },
  { name: 'EON Bank', icon: Eon },
  { name: 'ACLIDA Bank', icon: Ac },
  { name: 'AMK Bank', icon: Amk }
])

const cardIcons = ref([
  { type: 'text', value: 'VISA', class: 'visa' },
  { type: 'image', url: chipIcon, alt: 'Chip', class: 'icon-badge' },
  { type: 'text', value: 'DISCOVER', class: 'discover' },
  { type: 'text', value: 'MC', class: 'mastercard' },
  { type: 'image', url: chipIcon, alt: 'Chip', class: 'icon-badge' }
])

const bankIcons = ref([
  { type: 'image', url: crownIcon, alt: 'Crown', class: 'icon-badge' },
  { type: 'text', value: 'KHQR', class: 'khqr' },
  { type: 'text', value: 'ABA', class: 'aba' }
])

watch([isCountryOpen, isModalOpen, isModalBankOpen], ([countryOpen, modalOpen, modalBankOpen]) => {
  if (countryOpen || modalOpen || modalBankOpen) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
})

const toggleCountryDropdown = () => {
  isCountryOpen.value = !isCountryOpen.value
}

const selectCountry = (country) => {
  selectedCountry.value = country
  isCountryOpen.value = false
}

const openBankModal = () => {
  selectedPayment.value = 'bank'
  isModalOpen.value = true
}

const closeBankModal = () => {
  isModalOpen.value = false
  isModalBankOpen.value = false
}

const toggleModalBankDropdown = () => {
  isModalBankOpen.value = !isModalBankOpen.value
}

const selectModalBank = (bank) => {
  selectedBankObj.value = bank
  isModalBankOpen.value = false
}

const handlePayNow = () => {
  alert('Processing payment with ' + (selectedBankObj.value ? selectedBankObj.value.name : 'Bank'))
  closeBankModal()
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Kantumruy+Pro:wght@400;500;600;700&display=swap');

.page-wrapper {
  min-height: calc(100vh - 65px);
  background-color: #F7F4F2;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  box-sizing: border-box;
}

.payment-container {
  max-width: 790px;
  width: 100%;
  margin: 0 auto;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  padding: 20px;
  box-sizing: border-box;
}

.demo-header-label {
  font-size: 24px;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 16px;
}

.highlight-name {
  color: #1B75D2;
}

.access-banner {
  background: transparent;
  border-radius: 0;
  border-bottom: 2px solid #e2e8f0;
  padding: 16px 4px;
  display: flex;
  align-items: center;
  color: #0f172a;
  margin-bottom: 24px;
}

.access-text {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}

.plans-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 24px;
}

.plan-card {
  border: 6px solid #e2e8f0;
  padding: 16px;
  background-color: #ffffff;
  cursor: pointer;
  transition: all 0.2s ease;
}

.plan-card.active {
  border-color: #1B75D2;
  background-color: #f8fafc;
}

.plan-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 12px;
}

.radio-circle {
  width: 18px;
  height: 18px;
  border: 2px solid #cbd5e1;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.radio-circle.checked {
  border-color: #1B75D2;
  background-color: #1B75D2;
  box-shadow: inset 0 0 0 3px #ffffff;
}

.plan-title {
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
}

.plan-pricing {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 6px;
}

.plan-rate {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}

.save-badge {
  font-size: 11px;
  font-weight: 600;
  color: #00D262;
  padding: 2px 6px;
  border-radius: 4px;
}

.plan-subprice {
  font-size: 12px;
  color: #64748b;
}

.old-price {
  font-size: 11px;
  color: #ef4444;
  text-decoration: line-through;
  display: block;
  margin-top: 2px;
}

.billed-monthly {
  color: #64748b;
  margin-top: 4px;
}

.included-section {
  padding: 16px;
  margin-bottom: 24px;
}

.included-title {
  font-size: 14px;
  font-weight: 700;
  color: #0f172a;
  margin: 0 0 12px 0;
}

.benefit-item {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 15px;
  color: #0f172a;
  font-weight: 500;
  margin-bottom: 12px;
}

.benefit-item:last-child {
  margin-bottom: 0;
}

.check-box-icon {
  background-color: #00D262;
  color: #ffffff;
  width: 24px;
  height: 24px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.form-group {
  margin-bottom: 24px;
  position: relative;
}

.form-label {
  display: block;
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
  margin-bottom: 8px;
}

.select-dropdown {
  border: none;
  border-bottom: 1.5px solid #64748b;
  border-radius: 0;
  padding: 10px 4px 12px 4px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: transparent;
  cursor: pointer;
}

.select-left {
  display: flex;
  align-items: center;
}

.selected-country {
  font-size: 18px;
  font-weight: 400;
  color: #64748b;
}

.dropdown-arrow {
  color: #0f172a;
  display: flex;
  align-items: center;
}

.dropdown-menu-list {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  margin-top: 4px;
  z-index: 10;
}

.dropdown-item {
  padding: 10px 14px;
  font-size: 15px;
  color: #334155;
  cursor: pointer;
  background: #ffffff;
  transition: background 0.2s;
}

.dropdown-item:hover {
  background-color: #f1f5f9;
}

.bank-dropdown-item {
  display: flex;
  align-items: center;
  gap: 10px;
}

.bank-item-logo {
  width: 24px;
  height: 24px;
  object-fit: contain;
  border-radius: 4px;
}

.selected-bank-display {
  display: flex;
  align-items: center;
  gap: 10px;
}

.bank-option-logo {
  width: 22px;
  height: 22px;
  object-fit: contain;
  border-radius: 4px;
}

.payment-option {
  border: none;
  border-bottom: 1.5px solid #64748b;
  border-radius: 0;
  padding: 12px 4px 14px 4px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: transparent;
  cursor: pointer;
  margin-bottom: 12px;
  width: 100%;
  box-sizing: border-box;
  transition: all 0.2s ease;
}

.bank-dropdown-container {
  position: relative;
  margin-bottom: 12px;
}

.payment-option.active-option {
  border-bottom-color: #1B75D2;
  background-color: transparent;
}

.payment-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.payment-icon {
  width: 20px;
  height: 20px;
  color: #64748b;
}

.payment-name {
  font-size: 15px;
  font-weight: 600;
  color: #0f172a;
}

.card-logos, .bank-logos {
  display: flex;
  align-items: center;
  gap: 6px;
}

.brand-badge, .bank-badge {
  font-size: 10px;
  font-weight: 700;
  padding: 4px 8px;
  border-radius: 6px;
  background: #e2e8f0;
  color: #334155;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 28px;
  box-sizing: border-box;
}

.icon-badge-wrapper {
  background: transparent !important;
  border: none !important;
  box-shadow: none !important;
  padding: 0 !important;
  height: 28px;
  display: inline-flex;
  align-items: center;
}

.badge-img {
  width: 36px;
  height: 28px;
  border-radius: 6px;
  object-fit: cover;
  display: block;
}

.brand-badge.visa {
  background: #1a1f71;
  color: #ffffff;
}

.brand-badge.discover {
  background: #ff6000;
  color: #ffffff;
}

.brand-badge.mastercard {
  background: #eb001b;
  color: #ffffff;
}

.bank-badge.khqr {
  background: #DB1718;
  color: #ffffff;
}

.bank-badge.aba {
  background: #045A76;
  color: #ffffff;
}

/* ================= STYLES FOR THE MODAL ================= */
.modal-overlay {
  position: fixed;
  top: 65px;
  left: 0;
  width: 100vw;
  height: calc(100vh - 65px);
  background-color: rgba(15, 23, 42, 0.4);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: flex-start;
  justify-content: center;
  z-index: 1000;
  overflow-y: auto;
}

.modal-box {
  width: 790px;
  max-width: 95%;
  height: 100%;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  position: relative;
  box-sizing: border-box;
  padding: 20px;
  padding-top: 30px;
}

/* Modal Course Highlight Box Styles */
.modal-course-highlight-box {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 20px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e2e8f0;
}

.modal-course-thumb-wrapper {
  width: 130px;
  height: 80px;
  border-radius: 8px;
  overflow: hidden;
  flex-shrink: 0;
  border: 1px solid #e2e8f0;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
}

.modal-course-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.modal-course-text-content {
  font-size: 15px;
  color: #334155;
  line-height: 1.5;
}

.course-bold-title {
  font-weight: 700;
  color: #0f172a;
}

.modal-form-group {
  margin-top: 10px;
  margin-bottom: 18px;
  position: relative;
}

.modal-field-label {
  display: block;
  font-size: 14px;
  font-weight: 700;
  color: #0f172a;
  margin-bottom: 6px;
}

.modal-bank-select {
  background: transparent;
  border: none;
  border-bottom: 1.5px solid #64748b;
  border-radius: 0;
  padding: 10px 4px 12px 4px;
  box-shadow: none;
}

.placeholder-text {
  color: #64748b;
  font-size: 15px;
  font-weight: 500;
}

.active-bank-text {
  color: #0f172a !important;
  font-weight: 600;
  font-size: 15px;
}

.modal-bank-select .dropdown-arrow {
  color: #0f172a;
}

.modal-dropdown-list {
  background-color: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 10px;
  top: 100%;
  margin-top: 6px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.06);
}

.modal-dropdown-list .dropdown-item {
  color: #334155;
  background-color: #ffffff;
  padding: 10px 14px;
  font-size: 14px;
  font-weight: 500;
}

.modal-dropdown-list .dropdown-item:hover {
  background-color: #f8fafc;
  color: #0f172a;
}

/* ================= MODAL PLAN CARD STYLES ================= */
.modal-plan-card {
  background-color: #ffffff;
  padding: 16px;
  margin-top: 15px;
  margin-bottom: 20px;
  box-sizing: border-box;
}

.modal-plan-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 6px;
}

.modal-plan-title {
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
}

.modal-plan-price {
  font-size: 15px;
  font-weight: 700;
  color: #eb001b;
}

.modal-plan-desc {
  font-size: 13px;
  color: #64748b;
}

/* ================= MODAL SUMMARY SECTION STYLES ================= */
.modal-summary-section {
  position: relative;
  margin-top: 15px;
  margin-bottom: 24px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 16px;
  background-color: #ffffff;
}

.summary-label-title {
  display: block;
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
  margin-bottom: 8px;
}

.summary-row-box {
  background: transparent;
  border: none;
  border-bottom: 1.5px solid #e2e8f0;
  border-radius: 0;
  padding: 14px 4px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 4px;
}

.summary-label {
  font-size: 15px;
  font-weight: 500;
  color: #475569;
}

.summary-value {
  font-size: 16px;
  color: #0f172a;
  font-weight: 700;
}

.promo-box {
  background: transparent;
  border: none;
  border-bottom: 1.5px solid #e2e8f0;
  border-radius: 0;
  padding: 14px 4px;
  text-align: left;
  color: #475569;
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 4px;
}

.modal-summary-section .summary-row-box:last-of-type,
.modal-summary-section .promo-box:last-child {
  border-bottom: none;
  margin-bottom: 0;
  padding-bottom: 0;
}

/* Pay Now Button (Short & Centered with border-radius 32px) */
.pay-now-btn {
  display: block;
  width: 100%;
  max-width: 680px;
  margin: 0 auto;
  background: #1B75D2;
  border: none;
  color: #ffffff;
  padding: 14px;
  border-radius: 32px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  text-align: center;
  transition: transform 0.1s, opacity 0.2s;
  box-shadow: 0 4px 12px rgba(27, 117, 210, 0.2);
}

.pay-now-btn:hover {
  opacity: 0.9;
}

.pay-now-btn:active {
  transform: scale(0.99);
}

@media (max-width: 768px) {
  .plans-row {
    grid-template-columns: 1fr;
  }
}
</style>