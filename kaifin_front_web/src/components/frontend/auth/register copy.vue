<script setup>
import { ref, reactive } from 'vue'
import axios from 'axios' 
const fileInput = ref(null)
const profileImage = ref(null)
const previewUrl = ref('')

const userType = ref('')
const firstName = ref('')
const lastName = ref('')
const userName = ref('')
const email = ref('')
const password = ref('')
const rememberMe = ref(false)
const showPassword = ref(false)
const isSubmitting = ref(false)
const successMessage = ref(false)

const errors = reactive({
  profileImage: '',
  userType: '',
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
  previewUrl.value = URL.createObjectURL(file)
}

function validate() {
  errors.userType = userType.value ? '' : 'Please choose a user type'
  errors.firstName = firstName.value.trim() ? '' : 'Please enter your first name'
  errors.lastName = lastName.value.trim() ? '' : 'Please enter your last name'
  errors.userName = userName.value.trim() ? '' : 'Please enter a username'
  errors.email = /\S+@\S+\.\S+/.test(email.value) ? '' : 'Please enter a valid email'
  errors.password = password.value.length >= 6 ? '' : 'Password must be at least 6 characters'

  return (
    !errors.userType &&
    !errors.firstName &&
    !errors.lastName &&
    !errors.userName &&
    !errors.email &&
    !errors.password
  )
}

function handleSubmit() {
  if (!validate()) return

  isSubmitting.value = true

  // TODO: replace with your real signup API call
  // profileImage.value is a File object — send it as multipart/form-data, e.g.:
  // const formData = new FormData()
  // formData.append('profileImage', profileImage.value)
  emit('submit', {
    userType: userType.value,
    firstName: firstName.value,
    lastName: lastName.value,
    userName: userName.value,
    email: email.value,
    password: password.value,
    rememberMe: rememberMe.value,
    profileImage: profileImage.value,
  })

  async function handleSubmit() {
  if (!validate()) return

  isSubmitting.value = true
  serverError.value = ''

  const formData = new FormData()
  formData.append('userType', userType.value("user"))
  formData.append('firstName', firstName.value)
  formData.append('lastName', lastName.value)
  formData.append('userName', userName.value)
  formData.append('email', email.value)
  formData.append('password', password.value)
  formData.append('rememberMe', rememberMe.value ? 'true' : 'false')
  
  if (profileImage.value) {
    formData.append('profileImage', profileImage.value) 
  }

  try {
    const response = await axios.post('http://localhost:8080/api/v1/front/register/create', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })

    successMessage.value = true
    emit('submit', response.data)

    clearForm()

    setTimeout(() => {
      successMessage.value = false
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
  userType.value = ''
  firstName.value = ''
  lastName.value = ''
  userName.value = ''
  email.value = ''
  password.value = ''
  rememberMe.value = false
  profileImage.value = null
  previewUrl.value = ''
}
}
</script>
<template>
  <div class="signup-wrapper">
    <div class="signup-card">

      <div v-if="successMessage" class="toast">
        ✓ Account created — welcome aboard
      </div>

      <div class="left-panel">
        <h1 class="logo">Sign Up</h1>

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
          <p class="avatar-label">Upload profile photo</p>
          <p class="error" v-if="errors.profileImage">{{ errors.profileImage }}</p>
        </div>

        <div class="form-group">
          <select v-model="userType">
            <option value="" disabled>Choose type user</option>
            <option value="student">Student</option>
            <option value="teacher">Teacher</option>
          </select>
          <p class="error" v-if="errors.userType">{{ errors.userType }}</p>
        </div>

        <div class="form-group">
          <label>First Name</label>
          <input v-model="firstName" type="text" placeholder="Enter your first name" />
          <p class="error" v-if="errors.firstName">{{ errors.firstName }}</p>
        </div>

        <div class="form-group">
          <label>Last Name</label>
          <input v-model="lastName" type="text" placeholder="Enter your last name" />
          <p class="error" v-if="errors.lastName">{{ errors.lastName }}</p>
        </div>

        <div class="form-group">
          <label>User Name</label>
          <input v-model="userName" type="text" placeholder="Enter your username" />
          <p class="error" v-if="errors.userName">{{ errors.userName }}</p>
        </div>

        <div class="form-group">
          <label>Email address</label>
          <input v-model="email" type="email" placeholder="Enter your email..." />
          <p class="error" v-if="errors.email">{{ errors.email }}</p>
        </div>

        <div class="form-group">
          <label>Set password</label>
          <div class="password-field">
            <input
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="Create a password"
            />
            <button type="button" class="toggle-password" @click="showPassword = !showPassword">
              {{ showPassword ? 'Hide' : 'Show' }}
            </button>
          </div>
          <p class="error" v-if="errors.password">{{ errors.password }}</p>
        </div>

        <label class="remember-me">
          <input type="checkbox" v-model="rememberMe" />
          Save & remember me
        </label>

        <button class="submit-btn" :disabled="isSubmitting" @click="handleSubmit">
          {{ isSubmitting ? 'Creating account...' : 'Create account' }}
        </button>

        <p class="switch-text">
          Already have an account? <router-link to="/login">Log in</router-link>
        </p>
      </div>

      <div class="right-panel">
        <img src="../../../assets/logos/kaifin_l2.png" alt="" />
      </div>

    </div>
  </div>
</template>



<style scoped>
.signup-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 32px 16px;
  font-family: Arial, sans-serif;
}

.signup-card {
  position: relative;
  display: flex;
  width: 100%;
  max-width: 900px;
  min-height: 600px;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.left-panel {
  flex: 0 0 44%;
  background-color: #a2d6f9;
  padding: 36px 32px;
  display: flex;
  flex-direction: column;
  gap: 14px;
  color: white;
}

.right-panel {
  flex: 1;
  background-color: #cdebd6;
  display: flex;
  align-items: center;
  justify-content: center;
}

.right-panel img {
  width: 40%;
}

.logo {
  font-size: 22px;
  margin: 0 0 8px;
  color: black;
}

.avatar-group {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  margin-bottom: 6px;
}

.avatar-preview {
  width: 84px;
  height: 84px;
  border-radius: 50%;
  background-color: white;
  border: 2px dashed rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  cursor: pointer;
}

.avatar-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-placeholder {
  font-size: 28px;
  color: rgba(0, 0, 0, 0.4);
}

.hidden-input {
  display: none;
}

.avatar-label {
  font-size: 12px;
  color: black;
  margin: 0;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.form-group label {
  font-size: 13px;
  font-weight: bold;
  color: black;
}

.form-group input,
.form-group select {
  padding: 12px 14px;
  border-radius: 10px;
  border: none;
  font-size: 14px;
  outline: none;
}

.form-group select {
  background-color: white;
  border: 1px solid rgba(255, 255, 255, 0.5);
  color: black;
}

.form-group select option {
  color: black;
}

.password-field {
  position: relative;
  display: flex;
}

.password-field input {
  flex: 1;
  padding-right: 60px;
}

.toggle-password {
  position: absolute;
  right: 8px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: #333;
  font-size: 12px;
  cursor: pointer;
}

.error {
  color: red;
  font-size: 12px;
  margin: 0;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  cursor: pointer;
  color: black;
}

.submit-btn {
  margin-top: 10px;
  padding: 13px;
  border: none;
  border-radius: 10px;
  background-color: black;
  color: white;
  font-weight: bold;
  font-size: 15px;
  cursor: pointer;
}

.submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.switch-text {
  text-align: center;
  font-size: 12px;
  margin-top: 4px;
  color: black;
}

.switch-text a {
  color: black;
  font-weight: bold;
  text-decoration: none;
}

.toast {
  position: absolute;
  top: 16px;
  left: 50%;
  transform: translateX(-50%);
  background-color: #1c2033;
  color: white;
  padding: 10px 18px;
  border-radius: 8px;
  font-size: 13px;
  z-index: 10;
}

@media (max-width: 700px) {
  .signup-card {
    flex-direction: column;
    min-height: auto;
  }

  .right-panel {
    display: none;
  }
}
</style>