<script setup>
import { ref, reactive } from 'vue'
import axios from 'axios'
import router from '@/router'

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
    const response = await axios.post(
      'http://localhost:7070/api/v1/front/register/create',
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
          <p class="eyebrow">Join Kaifinn</p>
          <h1 class="logo">Start your<br />first tide.</h1>
          <!-- <p class="subtitle">One account, every tool you need to manage student money well.</p> -->

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
            <div class="form-group">
              <label>First name</label>
              <input v-model="firstName" type="text" placeholder="Sokha" />
              <p class="error" v-if="errors.firstName">{{ errors.firstName }}</p>
            </div>

            <div class="form-group">
              <label>Last name</label>
              <input v-model="lastName" type="text" placeholder="Chan" />
              <p class="error" v-if="errors.lastName">{{ errors.lastName }}</p>
            </div>
          </div>

          <div class="form-group">
            <label>Username</label>
            <input v-model="userName" type="text" placeholder="Choose a username" />
            <p class="error" v-if="errors.userName">{{ errors.userName }}</p>
          </div>

          <div class="form-group">
            <label>Email address</label>
            <input v-model="email" type="email" placeholder="you@example.com" />
            <p class="error" v-if="errors.email">{{ errors.email }}</p>
          </div>

          <div class="form-group">
            <label>Set password</label>
            <div class="password-field">
              <input
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                placeholder="At least 6 characters"
              />
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
        <div class="tide-signature">
          <div class="tide-bar" style="--h: 30%; --d: 0s"></div>
          <div class="tide-bar" style="--h: 46%; --d: .15s"></div>
          <div class="tide-bar" style="--h: 38%; --d: .3s"></div>
          <div class="tide-bar" style="--h: 62%; --d: .45s"></div>
          <div class="tide-bar" style="--h: 50%; --d: .6s"></div>
          <div class="tide-bar" style="--h: 78%; --d: .75s"></div>
          <div class="tide-line"></div>
        </div>

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
@import url('https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,500;9..144,600&family=Space+Grotesk:wght@400;500;600&display=swap');

* {
  box-sizing: border-box;
}

.signup-wrapper {
  --ink: #111827;
  --ink-soft: #4b5563;
  --bg-color: #f9fafb;
  --card-bg: #ffffff;
  --border-color: #e5e7eb;
  --accent: #2563eb;

  display: flex;
  justify-content: center;
  align-items: center;
  padding: 32px 16px;
  font-family: "Space Grotesk", system-ui, sans-serif;
  background: var(--bg-color);
  height: 100vh;
}

.signup-card {
  position: relative;
  display: flex;
  width: 100%;
  max-width: 1110px;
  /* min-height: 640px; */
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
  padding: 44px 40px;
  display: flex;
  position: relative;
  overflow-y: auto;
}

.left-inner {
  position: relative;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.eyebrow {
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  color: var(--accent);
  margin: 0;
}

.logo {
  font-family: "Fraunces", serif;
  font-weight: 500;
  font-size: 30px;
  line-height: 1.15;
  margin: 0;
  color: var(--ink);
}

.subtitle {
  font-size: 13px;
  line-height: 1.5;
  color: var(--ink-soft);
  margin: 0 0 4px;
  max-width: 38ch;
}

.avatar-group {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  margin: 4px 0 6px;
}

.avatar-preview {
  width: 76px;
  height: 76px;
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
  font-size: 26px;
  color: #9ca3af;
}

.hidden-input {
  display: none;
}

.avatar-label {
  font-size: 12px;
  color: var(--ink-soft);
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

.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.form-group label {
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.04em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.form-group input,
.form-group select {
  padding: 11px 14px;
  border-radius: 10px;
  border: 1px solid var(--border-color);
  background: #f9fafb;
  color: var(--ink);
  font-size: 14px;
  font-family: inherit;
  outline: none;
  transition: background 0.15s, border-color 0.15s;
}

.form-group input::placeholder {
  color: #9ca3af;
}

.form-group input:focus,
.form-group select:focus {
  background: #ffffff;
  border-color: var(--accent);
}

.form-group select option {
  color: var(--ink);
}

.password-field {
  position: relative;
  display: flex;
}

.password-field input {
  flex: 1;
  padding-right: 56px;
}

.toggle-password {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: var(--accent);
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
}

.error {
  color: #ef4444;
  font-size: 12px;
  margin: 0;
}

.server-error {
  text-align: center;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  cursor: pointer;
  color: var(--ink-soft);
}

.remember-me input {
  position: absolute;
  opacity: 0;
  width: 0;
  height: 0;
}

.check-box {
  width: 16px;
  height: 16px;
  border-radius: 4px;
  border: 1px solid #d1d5db;
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
  left: 5px;
  top: 2px;
  width: 4px;
  height: 8px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.submit-btn {
  margin-top: 6px;
  padding: 13px;
  border: none;
  border-radius: 10px;
  background-color: var(--ink);
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
  background: #1f2937;
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
  margin-top: 2px;
  color: var(--ink-soft);
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

.tide-signature {
  position: absolute;
  left: 50%;
  bottom: 110px;
  transform: translateX(-50%);
  width: 200px;
  height: 110px;
  display: flex;
  align-items: flex-end;
  gap: 10px;
  z-index: 1;
  opacity: 0.5;
}

.tide-bar {
  flex: 1;
  height: var(--h);
  border-radius: 6px 6px 2px 2px;
  background: linear-gradient(180deg, #60a5fa 0%, #2563eb 100%);
  animation: rise 3.2s ease-in-out infinite;
  animation-delay: var(--d);
  transform-origin: bottom;
}

.tide-line {
  position: absolute;
  left: -20px;
  right: -20px;
  bottom: -8px;
  height: 1px;
  background: rgba(0, 0, 0, 0.15);
}

@keyframes rise {
  0%, 100% { transform: scaleY(1); }
  50% { transform: scaleY(1.08); }
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
  font-family: "Fraunces", serif;
  font-weight: 500;
  font-size: 22px;
  margin: 0 0 6px;
}

.right-copy {
  font-size: 14px;
  color: var(--ink-soft);
  margin: 0;
}

@media (prefers-reduced-motion: reduce) {
  .tide-bar {
    animation: none;
  }
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
  }

  .right-panel {
    display: none;
  }

  .name-row {
    flex-direction: column;
  }
}
</style>