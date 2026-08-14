<script setup>
import { ref } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

const router = useRouter();

const userTypes = ["Student"];

const userType = ref("");
const typeOpen = ref(false);

const user_name = ref("");
const password = ref("");

const showPassword = ref(false);
const remember = ref(false);

// choose user type
function selectType(t) {
  userType.value = t;
  typeOpen.value = false;
}

// login API
async function onSignIn() {
  // validate
  if (!user_name.value) {
    alert("Please enter username");
    return;
  }

  if (!password.value) {
    alert("Please enter password");
    return;
  }

  if (password.value.length < 6) {
    alert("Password must be at least 6 characters");
    return;
  }

  try {
    const payload = {
      user_name: user_name.value,
      password: password.value,
      role_id: 4
    }

    console.log("SEND DATA:", payload);

    const response = await axios.post(
      "http://localhost:7070/api/v1/front/auth/login-user",
      payload
    );

    console.log("LOGIN RESPONSE:", response.data);

    if(response.data.success){
      const token = response.data.data.auth.token;
      localStorage.setItem("token", token);
      alert("Login success");
      router.push("/home");
    }

  } catch(error) {
    console.log("ERROR:", error);
    if(error.response){
       console.log("FULL RESPONSE DATA:", error.response.data);
      alert(error.response.data.message);
    }else{
      alert("Server error");
    }
  }
}
</script>

<template>
  <div class="signup-wrapper">
    <div class="signup-card">

      <!-- Left Panel (Form Login) -->
      <div class="left-panel">
        <div class="left-inner">
          <div class="logo-row">
            <div class="logo-box">
              <img src="../../../assets/logos/kaifin_l2.png" alt="">
            </div>
            <span class="brand-word">kaifinn</span>
          </div>

          <!-- <p class="eyebrow">Student banking, uncomplicated</p> -->
          <h1 class="logo"><span class="text-stroke">Welcome back.</span><span class="let-color text-stroke-let"> Let's grow it.</span></h1>
          <p class="subtitle">Sign in to track, <span class="text-plus-stroke">plan</span>, and <span class="text-plus-stroke">manage </span>your <span class="text-plus-stroke">money</span> in one place.</p>

          <div class="form-group dropdown-wrap">
            <label>Account Type</label>
            <button type="button" class="dropdown-trigger" @click="typeOpen = !typeOpen">
              <span>{{ userType || "Choose account type" }}</span>
              <svg :class="['chevron', { open: typeOpen }]" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polyline points="6 9 12 15 18 9"></polyline>
              </svg>
            </button>
            <div v-if="typeOpen" class="dropdown-menu">
              <button
                v-for="t in userTypes"
                :key="t"
                class="dropdown-item"
                @click="selectType(t)"
              >
                {{ t }}
              </button>
            </div>
          </div>

          <!-- Email or Username Field -->
          <div class="form-group">
            <label>Email or username</label>
            <div class="input-wrap">
              <svg class="input-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M22 6l-10 7L2 6"></path>
                <path d="M2 6h20v12H2z"></path>
              </svg>
              <input
                type="text"
                v-model="user_name"
                placeholder="you@example.com"
              />
            </div>
          </div>

          <div class="form-group">
            <label>Password</label>
            <div class="password-field">
              <svg class="input-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="3" y="11" width="18" height="11" rx="2"></rect>
                <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
              </svg>
              <input
                :type="showPassword ? 'text' : 'password'"
                v-model="password"
                placeholder="Enter your password"
              />
              <button type="button" class="toggle-password" @click="showPassword = !showPassword">
                {{ showPassword ? 'Hide' : 'Show' }}
              </button>
            </div>
          </div>

          <div class="row-between">
            <label class="remember-me">
              <input type="checkbox" v-model="remember" />
              <span class="check-box"></span>
              Remember me
            </label>
            <button class="link-btn" type="button">Forgot password?</button>
          </div>

          <button class="submit-btn" @click="onSignIn">
            <span>Sign in</span>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="5" y1="12" x2="19" y2="12"></line>
              <polyline points="12 5 19 12 12 19"></polyline>
            </svg>
          </button>

          <p class="switch-text">
            New to Kaifinn? <button class="switch-link" type="button" @click="router.push('/register')">Create an account</button>
          </p>
        </div>
      </div>

      <!-- Right Panel (Full SVG / Image Area) -->
      <div class="right-panel">
        <img src="../../../assets/animate/se.svg" alt="Full Banner" class="full-svg-banner" />
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
  /* font-family: "Space Grotesk", system-ui, sans-serif; */
  font-family: "Poppins", sans-serif;
  background: var(--bg-color);
  min-height: 100vh;
}

.signup-card {
  position: relative;
  display: flex;
  width: 100%;
  max-width: 1110px;
  min-height: 640px;
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

.logo-row {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 2px;
}

.logo-box {
  width: 36px;
  height: 36px;
  border-radius: 8px;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f3f4f6;
}

.logo-box img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.brand-word {
  font-family: "CodePro", sans-serif;
  font-weight: 600;
  font-size: 17px;
  letter-spacing: 0.02em;
  color: var(--ink);
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
  font-family: "CodePro", sans-serif;
  font-weight: 900;
  font-size: 40px;
  line-height: 1.15;
  margin: 0;
  color: var(--ink);

}

.text-stroke{
    -webkit-text-stroke: 1.5px var(--ink);
}

.text-plus-stroke{
  -webkit-text-stroke: 0.8px var(--ink);
}

.let-color{
  color: #1976D2;
  -webkit-text-stroke: 1.5px #1976D2;
}

.subtitle {
  font-size: 13px;
  line-height: 1.5;
  color: var(--ink-soft);
  margin: 0 0 4px;
  max-width: 38ch;
  font-family: "CodePro", sans-serif;

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

.dropdown-trigger {
  width: 100%;
  padding: 11px 14px;
  border-radius: 10px;
  border: 1px solid var(--border-color);
  /* background: #f9fafb; */
  background-color: transparent;
  color: var(--ink);
  font-size: 14px;
  font-family: inherit;
  outline: none;
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
  font-weight: 500;
  transition: background 0.15s, border-color 0.15s;
}

.dropdown-trigger:hover {
  background: #f3f4f6;
  border-color: #d1d5db;
}

.chevron {
  transition: transform 0.15s;
  color: var(--ink-soft);
}

.chevron.open {
  transform: rotate(180deg);
}

.dropdown-menu {
  position: absolute;
  z-index: 10;
  margin-top: 6px;
  width: 100%;
  background: #ffffff;
  border: 1px solid var(--border-color);
  border-radius: 10px;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.dropdown-item {
  width: 100%;
  text-align: left;
  padding: 10px 16px;
  font-size: 14px;
  color: var(--ink);
  background: transparent;
  border: none;
  cursor: pointer;
}

.dropdown-item:hover {
  background: #f3f4f6;
}

/* Fix Input & Icon Overlapping */
.input-wrap {
  position: relative;
  display: flex;
  align-items: center;
}

.input-wrap input {
  width: 100%;
  padding: 11px 14px 11px 40px; /* បន្ថែម padding ខាងឆ្វេងកុំให้อក្សរជាន់ icon */
  border-radius: 10px;
  border: 1px solid var(--border-color);
  background: #f9fafb;
  color: var(--ink);
  font-size: 14px;
  font-family: inherit;
  outline: none;
  transition: background 0.15s, border-color 0.15s;
}

.input-wrap input::placeholder {
  color: #9ca3af;
}

.input-wrap input:focus {
  background: #ffffff;
  border-color: var(--accent);
}

.input-icon {
  position: absolute;
  left: 14px;
  color: var(--ink-soft);
  pointer-events: none;
  z-index: 2;
}

.password-field {
  position: relative;
  display: flex;
  align-items: center;
}

.password-field input {
  width: 100%;
  padding: 11px 56px 11px 40px;
  border-radius: 10px;
  border: 1px solid var(--border-color);
  background: #f9fafb;
  color: var(--ink);
  font-size: 14px;
  font-family: inherit;
  outline: none;
  transition: background 0.15s, border-color 0.15s;
}

.password-field input:focus {
  background: #ffffff;
  border-color: var(--accent);
}

.password-field .input-icon {
  left: 14px;
  z-index: 2;
}

.toggle-password {
  position: absolute;
  right: 12px;
  background: none;
  border: none;
  color: var(--accent);
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
}

.row-between {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 2px 0;
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

.link-btn {
  background: none;
  border: none;
  font-size: 13px;
  font-weight: 500;
  color: var(--accent);
  cursor: pointer;
  padding: 0;
}

.link-btn:hover {
  text-decoration: underline;
}

.submit-btn {
  margin-top: 6px;
  padding: 13px;
  border: none;
  border-radius: 10px;
  background-color: #1976D2;
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

.submit-btn:hover {
  transform: translateY(-1px);
  opacity: 0.8;
}

.switch-text {
  text-align: center;
  font-size: 12px;
  margin-top: 2px;
  color: var(--ink-soft);
}

.switch-link {
  background: none;
  border: none;
  color: var(--accent);
  font-weight: 600;
  cursor: pointer;
  padding: 0;
  font-size: 12px;
  font-family: inherit;
}

.switch-link:hover {
  text-decoration: underline;
}

/* ---------- Right (Full SVG / Image Area) ---------- */
.right-panel {
  flex: 1;
  position: relative;
  background: #e5e7eb;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  padding: 0;
}

.full-svg-banner {
  width: 100%;
  height: 100%;
  object-fit: cover;
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
}
</style>