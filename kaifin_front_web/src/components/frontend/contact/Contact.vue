<script setup>
import { ref, onMounted } from 'vue'
import NavBar from '../navbar/NavBar.vue'

const form = ref({
  name: '',
  email: '',
  message: ''
})

const isSubmitting = ref(false)
const successMessage = ref('')

// Typing Animation States
const fullText = "Let’s Get In Touch"
const typedText = ref('')
let isDeleting = false // សម្រាប់ឆែកថាកំពុងវាយ ឬកំពុងលុប

const typeWriter = () => {
  const currentTextLength = typedText.value.length

  if (!isDeleting && currentTextLength < fullText.length) {
    // ដំណាក់កាលកំពុងវាយអក្សរបញ្ជូល (Typing)
    typedText.value = fullText.substring(0, currentTextLength + 1)
    setTimeout(typeWriter, 150) // ល្បឿនវាយ
  } 
  else if (isDeleting && currentTextLength > 0) {
    // ដំណាក់កាលកំពុងលុបអក្សរ (Deleting)
    typedText.value = fullText.substring(0, currentTextLength - 1)
    setTimeout(typeWriter, 50) // ល្បឿនលុប (លឿនជាងវាយបន្តិច)
  } 
  else {
    // ប្ដូរស្ថានភាពរវាងការវាយ និងការលុប ពេលវាដើរដល់ចុងបញ្ចប់
    isDeleting = !isDeleting
    
    // បើវាយពេញហើយ ផ្អាក 2 វិនាទីសិនមុនលុបវិញ, បើលុបអស់ហើយ ផ្អាក 0.5 វិនាទីមុនវាយវិញ
    const delay = isDeleting ? 2000 : 500 
    setTimeout(typeWriter, delay)
  }
}

onMounted(() => {
  typeWriter()
})

const handleSubmit = () => {
  if (!form.value.name || !form.value.email || !form.value.message) {
    alert('please input info!')
    return
  }

  isSubmitting.value = true
  
  setTimeout(() => {
    isSubmitting.value = false
    successMessage.value = 'sucess!'
    form.value = { name: '', email: '', message: '' }
    
    setTimeout(() => {
      successMessage.value = ''
    }, 4000)
  }, 1000)
}
</script>

<template>
  <div class="page-container">
    <NavBar/>
    <section class="contact-wrapper">
      <div class="contact-card">
        
        <!-- Main Content Grid -->
        <div class="contact-grid">
          
          <!-- Left Side: Creative Intro & Info -->
          <div class="contact-info-side">
            <div class="badge-tag">
              <span class="icon-circle">
                <svg class="send-icon" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
              </span> Get in Touch
            </div>
            
            <!-- Typing Text Target -->
            <h2 class="main-title">
              {{ typedText }}<span class="cursor">|</span>
            </h2>

            <p class="description-text">
              Do you have any question or join? Please send me my team reply.
            </p>

            <div class="creative-shapes">
              <div class="shape-circle"></div>
              <div class="shape-line"></div>
            </div>
          </div>

          <!-- Right Side: Simple & Clean Form -->
          <div class="contact-form-side">
            <form @submit.prevent="handleSubmit" class="form-container">
              
              <div v-if="successMessage" class="alert-success">
                {{ successMessage }}
              </div>

              <div class="input-group">
                <label for="name">Your Name</label>
                <input 
                  id="name"
                  type="text" 
                  v-model="form.name" 
                  placeholder="input your name" 
                  required
                />
              </div>

              <div class="input-group">
                <label for="email">Your Email</label>
                <input 
                  id="email"
                  type="email" 
                  v-model="form.email" 
                  placeholder="input your email" 
                  required
                />
              </div>

              <div class="input-group">
                <label for="message">Your Message</label>
                <textarea 
                  id="message"
                  v-model="form.message" 
                  rows="4" 
                  placeholder="input your message" 
                  required
                ></textarea>
              </div>

              <button type="submit" class="submit-btn" :disabled="isSubmitting">
                <span v-if="isSubmitting">Submitting...</span>
                <span v-else>Submit</span>
              </button>
            </form>
          </div>

        </div>

        <!-- Footer Bar: Visit Us & Socials -->
        <div class="contact-footer-bar">
          
          <div class="visit-us-section">
            <h3 class="footer-heading">Visit Us</h3>
            <div class="info-items">
              <div class="info-row">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path><circle cx="12" cy="10" r="3"></circle></svg>
                <span>#1,3,5A, St.1, Borey Piphup Thmey Chamkar Doung, Phnom Penh</span>
              </div>
              <div class="info-row">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path></svg>
                <span>096 738 5255</span>
              </div>
            </div>
          </div>

          <div class="follow-us-section">
            <span class="follow-label">Follow Us</span>
            <div class="social-icons">
              <a href="#" class="social-btn facebook" title="Facebook">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor"><path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"></path></svg>
              </a>
              <a href="#" class="social-btn telegram" title="Telegram">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor"><path d="M21.5 3.5L2 11l6.5 2.5L20 6l-9 10.5v4l3.5-3 3.5 2.5L21.5 3.5z"></path></svg>
              </a>
            </div>
          </div>

        </div>

      </div>
    </section>
  </div>
</template>

<style scoped>
/* Page & Container Theme */
.page-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #F7F4F2;
}

.page-container > :deep(nav), 
.page-container > div:first-child {
  position: sticky;
  top: 0;
  z-index: 1000;
  width: 100%;
}

.contact-wrapper {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: stretch;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.contact-card {
  width: 100%;
  max-width: 1100px;
  background: #ffffff;
  overflow: hidden;
  border: 1px solid #edf2f7;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

/* Grid Layout */
.contact-grid {
  display: grid;
  grid-template-columns: 1fr 1.25fr;
  background: #ffffff;
  padding: 50px;
  gap: 50px;
  flex: 1;
  align-items: center;
}

@media (max-width: 768px) {
  .contact-grid {
    grid-template-columns: 1fr;
    padding: 30px 20px;
  }
}

/* Left Side (Creative Minimalist) */
.contact-info-side {
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: relative;
}

.badge-tag {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: #3b82d2;
  color: #ffffff;
  padding: 6px 14px;
  border-radius: 30px;
  font-size: 12px;
  font-weight: 600;
  text-transform: uppercase;
  margin-bottom: 20px;
  width: fit-content;
  border: 1px solid #3b82f6;
}

.icon-circle {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 23px;
  height: 23px;
  background-color: rgba(255, 255, 255, 0.112); 
  border-radius: 50%;
  backdrop-filter: blur(4px); 
}

.send-icon {
  color: #ffffff;
}

.main-title {
  font-size: 36px;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 16px;
  letter-spacing: -0.5px;
  min-height: 48px; /* ជួយទប់កុំឱ្យ Layout រំកិលពេលអក្សរកំពុងលេចឡើង */
}

/* Typing Cursor Effect */
.cursor {
  font-weight: 400;
  color: #3b82f6;
  animation: blink 0.8s infinite;
}

@keyframes blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}

.description-text {
  color: #64748b;
  font-size: 15px;
  line-height: 1.7;
  margin-bottom: 30px;
}

/* Creative Decorative Elements */
.creative-shapes {
  display: flex;
  align-items: center;
  gap: 12px;
}

.shape-circle {
  width: 12px;
  height: 12px;
  background: #8ac926;
  border-radius: 50%;
}

.shape-line {
  width: 60px;
  height: 3px;
  background: #8ac926;
  border-radius: 2px;
}

/* Right Side (Form UI) */
.contact-form-side {
  padding: 35px;
}

.form-container {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.input-group label {
  font-size: 13px;
  font-weight: 600;
  color: #000;
}

.input-group input,
.input-group textarea {
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  padding: 12px 16px;
  color: #0f172a;
  font-size: 14px;
  outline: none;
  transition: all 0.2s ease;
}

.input-group input:focus,
.input-group textarea:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.1);
}

.input-group input::placeholder,
.input-group textarea::placeholder {
  color: #94a3b8;
}

.submit-btn {
  background: #1976d2;
  color: white;
  border: none;
  border-radius: 32px;
  padding: 14px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s, transform 0.1s;
  margin-top: 10px;
}

.submit-btn:hover {
  opacity: 0.8;
}

.submit-btn:active {
  transform: scale(0.98);
}

.alert-success {
  background: #dcfce7;
  color: #166534;
  padding: 10px 14px;
  border-radius: 10px;
  font-size: 13px;
  border: 1px solid #bbf7d0;
}

/* Footer Bottom Bar (Simple & Clean) */
.contact-footer-bar {
  background: #ffffff;
  padding: 30px 50px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 24px;
  border-top: 1px solid #edf2f7;
}

.footer-heading {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  margin-bottom: 8px;
}

.info-items {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.info-row {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #64748b;
  font-size: 13px;
}

.info-row svg {
  color: #0f172a;
  flex-shrink: 0;
}

.follow-us-section {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 8px;
}

.follow-label {
  font-size: 13px;
  font-weight: 700;
  color: #0f172a;
}

.social-icons {
  display: flex;
  gap: 10px;
}

.social-btn {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  text-decoration: none;
  transition: transform 0.2s, opacity 0.2s;
}

.social-btn:hover {
  transform: translateY(-2px);
  opacity: 0.9;
}

.social-btn.facebook {
  background-color: #1877f2;
}

.social-btn.telegram {
  background-color: #229ed9;
}

@media (max-width: 768px) {
  .contact-footer-bar {
    flex-direction: column;
    align-items: flex-start;
    padding: 25px 20px;
  }
}
</style>