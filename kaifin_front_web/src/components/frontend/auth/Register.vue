<script setup>
import { ref, reactive } from 'vue'
import axios from 'axios'
import router from '@/router'
import api from '@/api/axios'


const fileInput = ref(null)
const profileImage = ref(null)
const previewUrl = ref('')

const firstName = ref('')
const lastName = ref('')
const userName = ref('')
const email = ref('')
const password = ref('')
const rememberMe = ref(false)
const showPassword = ref(false)
const isSubmitting = ref(false)
const successMessage = ref(false)
const serverError = ref('')

const errors = reactive({
  profileImage: '',
  firstName: '',
  lastName: '',
  userName: '',
  email: '',
  password: '',
})

const emit = defineEmits(['submit'])

function handleImageChange(event) {
  const file = event.target.files[0]
  if (!file) return

  if (!file.type.startsWith('image/')) {
    errors.profileImage = 'Please choose an image file'
    return
  }

  errors.profileImage = ''
  profileImage.value = file

  if (previewUrl.value) URL.revokeObjectURL(previewUrl.value)
  previewUrl.value = URL.createObjectURL(file)
}

function validate() {
  errors.firstName = firstName.value.trim() ? '' : 'Please enter your first name'
  errors.lastName = lastName.value.trim() ? '' : 'Please enter your last name'
  errors.userName = userName.value.trim() ? '' : 'Please enter a username'
  errors.email = /\S+@\S+\.\S+/.test(email.value) ? '' : 'Please enter a valid email'
  errors.password = password.value.length >= 6 ? '' : 'Password must be at least 6 characters'

  return (
    !errors.firstName &&
    !errors.lastName &&
    !errors.userName &&
    !errors.email &&
    !errors.password
  )
}

async function handleSubmit() {
  if (!validate()) return

  isSubmitting.value = true
  serverError.value = ''

  const formData = new FormData()
  formData.append('first_name', firstName.value)
  formData.append('last_name', lastName.value)
  formData.append('user_name', userName.value)
  formData.append('email', email.value)
  formData.append('password', password.value)

  if (profileImage.value) {
    formData.append('profile_images', profileImage.value)
  }

  try {
    const response = await api.post(
      '/api/v1/front/register/create',
      formData,
      { headers: { 'Content-Type': 'multipart/form-data' } }
    )

    successMessage.value = true
    emit('submit', response.data)
    clearForm()

    setTimeout(() => {
      successMessage.value = false
      router.push("/login")
    }, 2500)
  } catch (error) {
    console.error('Registration error:', error)
    if (error.response && error.response.data) {
      serverError.value = error.response.data.message || 'Registration failed. Please try again.'
    } else {
      serverError.value = 'Cannot connect to server.'
    }
  } finally {
    isSubmitting.value = false
  }
}

function clearForm() {
  firstName.value = ''
  lastName.value = ''
  userName.value = ''
  email.value = ''
  password.value = ''
  rememberMe.value = false
  profileImage.value = null
  if (previewUrl.value) URL.revokeObjectURL(previewUrl.value)
  previewUrl.value = ''
}
</script>

<template>
  <div class="signup-wrapper">
    <div class="signup-card">

      <transition name="toast-drop">
        <div v-if="successMessage" class="toast">
          <span class="toast-dot"></span>
          Account created — welcome aboard
        </div>
      </transition>

      <div class="left-panel">
        <div class="left-inner">
          
          <h1 class="logo">
            <span class="start">Start your first </span>

            <span class="first-tide"> with kaifinn</span>
          </h1>
          <p class="eyebrow">Join Kaifinn</p>

          <div class="avatar-group">
            <div class="avatar-preview" @click="fileInput.click()">
              <img v-if="previewUrl" :src="previewUrl" alt="Profile preview" />
              <span v-else class="avatar-placeholder">+</span>
            </div>
            <input
              ref="fileInput"
              type="file"
              accept="image/*"
              class="hidden-input"
              @change="handleImageChange"
            />
            <p class="avatar-label">Upload a profile photo</p>
            <p class="error" v-if="errors.profileImage">{{ errors.profileImage }}</p>
          </div>

          <div class="name-row">
            <div class="form-group floating">
              <div class="input-wrapper">
                <svg class="input-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                <input v-model="firstName" type="text" placeholder=" " />
                <label>First name*</label>
                <button v-if="firstName" type="button" class="clear-input-btn" @click="firstName = ''">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"><path d="M12 10.586l4.95-4.95 1.414 1.414-4.95 4.95 4.95 4.95-1.414 1.414-4.95-4.95-4.95 4.95-1.414-1.414 4.95-4.95-4.95-4.95L7.05 5.636z"/></svg>
                </button>
              </div>
              <p class="error" v-if="errors.firstName">{{ errors.firstName }}</p>
            </div>

            <div class="form-group floating">
              <div class="input-wrapper">
                <svg class="input-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                <input v-model="lastName" type="text" placeholder=" " />
                <label>Last name*</label>
                <button v-if="lastName" type="button" class="clear-input-btn" @click="lastName = ''">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"><path d="M12 10.586l4.95-4.95 1.414 1.414-4.95 4.95 4.95 4.95-1.414 1.414-4.95-4.95-4.95 4.95-1.414-1.414 4.95-4.95-4.95-4.95L7.05 5.636z"/></svg>
                </button>
              </div>
              <p class="error" v-if="errors.lastName">{{ errors.lastName }}</p>
            </div>
          </div>

          <div class="form-group floating">
            <div class="input-wrapper">
              <svg class="input-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path><polyline points="22,6 12,13 2,6"></polyline></svg>
              <input v-model="email" type="email" placeholder=" " />
              <label>Email address*</label>
              <button v-if="email" type="button" class="clear-input-btn" @click="email = ''">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"><path d="M12 10.586l4.95-4.95 1.414 1.414-4.95 4.95 4.95 4.95-1.414 1.414-4.95-4.95-4.95 4.95-1.414-1.414 4.95-4.95-4.95-4.95L7.05 5.636z"/></svg>
              </button>
            </div>
            <p class="error" v-if="errors.email">{{ errors.email }}</p>
          </div>

          <div class="form-group floating">
            <div class="input-wrapper">
              <svg class="input-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="8.5" cy="7" r="4"></circle><line x1="20" y1="8" x2="20" y2="14"></line><line x1="23" y1="11" x2="17" y2="11"></line></svg>
              <input v-model="userName" type="text" placeholder=" " />
              <label>Username*</label>
              <button v-if="userName" type="button" class="clear-input-btn" @click="userName = ''">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"><path d="M12 10.586l4.95-4.95 1.414 1.414-4.95 4.95 4.95 4.95-1.414 1.414-4.95-4.95-4.95 4.95-1.414-1.414 4.95-4.95-4.95-4.95L7.05 5.636z"/></svg>
              </button>
            </div>
            <p class="error" v-if="errors.userName">{{ errors.userName }}</p>
          </div>

          <div class="form-group floating">
            <div class="input-wrapper password-field">
              <svg class="input-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
              <input
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                placeholder=" "
              />
              <label>Password*</label>
              <button type="button" class="toggle-password" @click="showPassword = !showPassword">
                {{ showPassword ? 'Hide' : 'Show' }}
              </button>
            </div>
            <p class="error" v-if="errors.password">{{ errors.password }}</p>
          </div>

          <label class="remember-me">
            <input type="checkbox" v-model="rememberMe" />
            <span class="check-box"></span>
            Save &amp; remember me
          </label>

          <p class="error server-error" v-if="serverError">{{ serverError }}</p>

          <button class="submit-btn" :disabled="isSubmitting" @click="handleSubmit">
            <span v-if="isSubmitting" class="spinner"></span>
            <span>{{ isSubmitting ? 'Creating account...' : 'Create account' }}</span>
          </button>

          <p class="switch-text">
            Already have an account? <router-link to="/login">Log in</router-link>
          </p>
        </div>
      </div>

      <div class="right-panel">
        <div class="brand-mark">
          <img src="../../../assets/logos/kaifin_l2.png" alt="" />
        </div>

        <div class="right-caption">
          <p class="right-kicker">Every account starts small</p>
          <p class="right-heading">Grow with confidence</p>
          <p class="right-copy">Track, plan, and manage your money in one place.</p>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,500;9..144,600&family=Poppins:wght@400;500;600&family=Space+Grotesk:wght@400;500;600&display=swap');

* {
  box-sizing: border-box;
}

.signup-wrapper {
  --ink: #111827;
  --bg-color: #f9fafb;
  --card-bg: #ffffff;
  --border-color: #d1d5db;
  --accent: #1B76D2;

  display: flex;
  justify-content: center;
  align-items: center;
  padding: 16px;
  font-family: "Space Grotesk", system-ui, sans-serif;
  background: var(--bg-color);
  min-height: 100vh;
}

.signup-card {
  position: relative;
  display: flex;
  width: 100%;
  max-width: 1110px;
  background: var(--card-bg);
  border-radius: 24px;
  overflow: hidden;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.05), 0 8px 10px -6px rgba(0, 0, 0, 0.05);
  border: 1px solid var(--border-color);
}

/* ---------- Left ---------- */
.left-panel {
  flex: 0 0 50%;
  background: #ffffff;
  color: var(--ink);
  padding: 24px 32px;
  display: flex;
  position: relative;
  max-height: 88vh;
  overflow-y: auto;
}

.left-inner {
  position: relative;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.eyebrow {
  font-family: "CodePro", sans-serif;
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  color: var(--accent);
  margin: 0;
}

.logo {
  font-family: "CodePro", sans-serif;
  font-weight: 500;
  font-size: 20px;
  line-height: 1.15;
  margin: 0;
}

.start {
  -webkit-text-stroke: 1.5px #000;
  color: var(--ink);
}

.x-mark {
  -webkit-text-stroke: 1.5px #111827;
  color: #111827;
  font-weight: 900;
}

.first-tide {
  -webkit-text-stroke: 1.5px #1B76D2;
  color: #1B76D2;
}

.avatar-group {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  margin: 2px 0;
}

.avatar-preview {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background-color: #f3f4f6;
  border: 2px dashed #d1d5db;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  cursor: pointer;
  transition: border-color 0.15s, background 0.15s;
}

.avatar-preview:hover {
  border-color: var(--accent);
  background: #e5e7eb;
}

.avatar-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-placeholder {
  font-size: 22px;
  color: #9ca3af;
}

.hidden-input {
  display: none;
}

.avatar-label {
  font-size: 11px;
  color: #4b5563;
  margin: 0;
}

.name-row {
  display: flex;
  gap: 12px;
}

.name-row .form-group {
  flex: 1;
  min-width: 0;
}

/* ---------- Floating Label Styles ---------- */
.form-group.floating {
  position: relative;
  margin-top: 6px;
  font-family: 'Poppins', sans-serif;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  border-bottom: 2px solid var(--border-color);
  transition: border-color 0.15s;
}

.input-icon {
  width: 18px;
  height: 18px;
  color: #9ca3af;
  margin-right: 10px;
  flex-shrink: 0;
  transition: color 0.15s;
}

.form-group.floating input {
  width: 100%;
  padding: 14px 0px 6px 0px;
  border: none;
  background: transparent;
  color: var(--ink);
  font-size: 14px;
  font-family: 'Poppins', sans-serif;
  font-weight: 500;
  outline: none;
}

.form-group.floating label {
  position: absolute;
  left: 28px;
  top: 12px;
  font-size: 14px;
  color: #9ca3af;
  pointer-events: none;
  transition: 0.2s ease all;
  font-family: 'Poppins', sans-serif;
}

.input-wrapper:focus-within {
  border-color: var(--accent);
}

.input-wrapper:focus-within .input-icon {
  color: var(--accent);
}

/* លោតឡើងលើពេល Focus ឬមានអត្ថបទ */
.form-group.floating input:focus ~ label,
.form-group.floating input:not(:placeholder-shown) ~ label {
  top: -8px;
  left: 0px;
  font-size: 11px;
  color: var(--accent);
  font-weight: 600;
}

.clear-input-btn {
  background: #111827;
  border: none;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  padding: 0;
  margin-left: 8px;
  flex-shrink: 0;
  transition: background 0.15s;
}

.clear-input-btn svg {
  width: 12px;
  height: 12px;
  fill: #ffffff;
}

.clear-input-btn:hover {
  background: #1f2937;
}

.password-field {
  position: relative;
  display: flex;
  width: 100%;
}

.password-field input {
  flex: 1;
  padding-right: 50px;
}

.toggle-password {
  position: absolute;
  right: 0px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: var(--accent);
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  z-index: 2;
  font-family: 'Poppins', sans-serif;
}

.error {
  color: #ef4444;
  font-size: 11px;
  margin-top: 2px;
}

.server-error {
  text-align: center;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  cursor: pointer;
  color: #4b5563;
  font-weight: 500;
  margin-top: 8px;
  font-family: 'Poppins', sans-serif;
}

.remember-me input {
  position: absolute;
  opacity: 0;
  width: 0;
  height: 0;
}

.check-box {
  width: 15px;
  height: 15px;
  border-radius: 4px;
  border: 1px solid #9ca3af;
  display: inline-block;
  position: relative;
  flex-shrink: 0;
  transition: background 0.15s, border-color 0.15s;
}

.remember-me input:checked + .check-box {
  background: var(--accent);
  border-color: var(--accent);
}

.remember-me input:checked + .check-box::after {
  content: "";
  position: absolute;
  left: 4px;
  top: 1px;
  width: 4px;
  height: 7px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.submit-btn {
  margin-top: 8px;
  padding: 12px;
  border: none;
  border-radius: 32px;
  background-color: #1B76D2;
  color: #ffffff;
  font-weight: 600;
  font-size: 14px;
  font-family: inherit;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  transition: background 0.15s, transform 0.15s;
}

.submit-btn:hover:not(:disabled) {
  background: #155fa0;
  transform: translateY(-1px);
}

.submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.spinner {
  width: 14px;
  height: 14px;
  border-radius: 50%;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: #ffffff;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.switch-text {
  text-align: center;
  font-size: 12px;
  margin-top: 4px;
  color: #4b5563;
}

.switch-text a {
  color: var(--accent);
  font-weight: 600;
  text-decoration: none;
}

.switch-text a:hover {
  text-decoration: underline;
}

.toast {
  position: absolute;
  top: 16px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: var(--ink);
  color: #ffffff;
  padding: 10px 18px;
  border-radius: 999px;
  font-size: 13px;
  z-index: 10;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
}

.toast-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--accent);
  flex-shrink: 0;
}

.toast-drop-enter-active,
.toast-drop-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.toast-drop-enter-from,
.toast-drop-leave-to {
  opacity: 0;
  transform: translate(-50%, -8px);
}

/* ---------- Right ---------- */
.right-panel {
  flex: 1;
  position: relative;
  background: #e5e7eb;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.brand-mark {
  position: relative;
  z-index: 2;
  width: 96px;
  height: 96px;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 28px;
}

.brand-mark img {
  width: 60%;
  height: 60%;
  object-fit: contain;
}

.right-caption {
  position: relative;
  z-index: 2;
  text-align: center;
  color: var(--ink);
  padding: 0 40px;
}

.right-kicker {
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.1em;
  text-transform: uppercase;
  color: var(--accent);
  margin: 0 0 8px;
}

.right-heading {
  font-family: "CodePro", sans-serif;
  font-weight: 900;
  font-size: 22px;
  margin: 0 0 6px;
  -webkit-text-stroke: 1.5px var(--ink);
}

.right-copy {
  font-size: 14px;
  color: var(--ink-soft);
  margin: 0;
}

@media (prefers-reduced-motion: reduce) {
  .spinner {
    animation: none;
  }
}

@media (max-width: 700px) {
  .signup-card {
    flex-direction: column;
    min-height: auto;
  }

  .left-panel {
    flex: 1;
    max-height: none;
  }

  .right-panel {
    display: none;
  }

  .name-row {
    flex-direction: column;
  }
}
</style>