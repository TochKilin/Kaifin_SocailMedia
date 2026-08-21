<template>
  <Teleport to="body">
    <div class="gc-overlay" :style="{ top: navbarHeight + 'px' }" @mousedown.self="close">
      <div class="gc-wrap">
        <div class="wrapper-add">
          <button class="wrapper-add-btn" type="button" aria-label="Close" @click="close">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="3" />
                <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 1 1-4 0v-.09a1.65 1.65 0 0 0-1.08-1.51 1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 1 1 0-4h.09a1.65 1.65 0 0 0 1.51-1.08 1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 1 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 1 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z" />
            </svg>
          </button>
          <button class="wrapper-add-btn" type="button" aria-label="Close" @click="close">
            <svg viewBox="0 0 24 24"><path d="M6 6l12 12M18 6L6 18" /></svg>
           </button>
        </div>

        <div class="gc-modal">
        <!-- Header -->
        <div class="gc-header">
          <div class="gc-blob gc-blob-1"></div>
          <div class="gc-blob gc-blob-2"></div>
          <div class="gc-blob gc-blob-3"></div>
          <div class="gc-dot gc-dot-1"></div>
          <div class="gc-dot gc-dot-2"></div>
          <div class="gc-dash gc-dash-1"></div>
          <div class="gc-dash gc-dash-2"></div>
        </div>

        <!-- Body -->
        <div class="gc-body">
          <h2 class="gc-title">Public or business page</h2>
          <p class="gc-subtitle">For business, brands, media, blogs, musical artists, sports teams and many more.</p>

          <div class="gc-field">
            <label class="gc-label" for="gc-name">Name</label>
            <input
              id="gc-name"
              v-model="form.name"
              class="gc-input"
              type="text"
              placeholder=""
              maxlength="75"
            />
          </div>

          <div class="gc-field">
            <label class="gc-label" for="gc-description">Description</label>
            <textarea
              id="gc-description"
              v-model="form.description"
              class="gc-textarea"
              rows="3"
              placeholder=""
              maxlength="500"
            ></textarea>
          </div>

          <div class="gc-field">
            <label class="gc-label" for="gc-category">Category</label>
            <div class="gc-select-wrap">
              <select id="gc-category" v-model="form.category" class="gc-select">
                <option value="" disabled hidden>Not selected</option>
                <option v-for="c in categories" :key="c" :value="c">{{ c }}</option>
              </select>
              <svg class="gc-select-chevron" viewBox="0 0 24 24"><path d="M6 9l6 6 6-6" /></svg>
            </div>
          </div>

          <div class="gc-field gc-field-row">
            <span class="gc-label gc-label-access">
              Access
              <span class="gc-help" tabindex="0" :title="accessHelp">
                <svg viewBox="0 0 24 24"><circle cx="12" cy="12" r="9" /><path d="M9.5 9a2.5 2.5 0 0 1 4.9.8c0 1.7-2.4 1.9-2.4 3.5" /><path d="M12 17h.01" /></svg>
              </span>
            </span>

            <div class="gc-radio-group" role="radiogroup" aria-label="Access">
              <label v-for="opt in accessOptions" :key="opt.value" class="gc-radio">
                <input
                  type="radio"
                  name="gc-access"
                  :value="opt.value"
                  v-model="form.access"
                />
                <span class="gc-radio-dot"></span>
                <span class="gc-radio-text">{{ opt.label }}</span>
              </label>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="gc-footer">
          <p class="gc-footer-note">Please be courteous and abide by the law and the Group Content Requirements</p>
          <button
            type="button"
            class="gc-create-btn"
            :disabled="!canCreate"
            @click="createGroup"
          >
            Create
          </button>
        </div>
      </div>
    </div>
    </div>
  </Teleport>
</template>

<script>
export default {
  name: 'GroupCreateFrom',

  props: {
    categories: {
      type: Array,
      default: () => [
        'Buy & sell',
        'Community',
        'Custom',
        'Gaming',
        'Learning',
        'Parenting',
        'Social club',
        'Support',
      ],
    },
    navbarHeight: {
      type: Number,
      default: 64,
    },
  },

  emits: ['close', 'create'],

  data() {
    return {
      form: {
        name: '',
        description: '',
        category: '',
        access: 'public',
      },
      accessOptions: [
        { value: 'public', label: 'Public' },
        { value: 'closed', label: 'Private' },
        { value: 'secret', label: 'Only me' },
      ],
      accessHelp: 'Public: anyone can see who\u2019s in the group and what they post. Closed: anyone can find the group, only members see posts. Secret: only members can find and see the group.',
    };
  },

  computed: {
    canCreate() {
      return this.form.name.trim().length > 0 && !!this.form.category;
    },
  },

  mounted() {
    this.prevBodyOverflow = document.body.style.overflow;
    document.body.style.overflow = 'hidden';
  },

  beforeUnmount() {
    document.body.style.overflow = this.prevBodyOverflow || '';
  },

  methods: {
    close() {
      this.$emit('close');
    },
    createGroup() {
      if (!this.canCreate) return;
      this.$emit('create', { ...this.form });
    },
  },
};
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.gc-overlay {
  position: fixed;
  top: 64px; 
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(15, 15, 20, 0.55);
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding: 12px;
  overflow-y: auto;
  z-index: 1000;
  font-family: 'Nunito', 'Inter', sans-serif;
}

.gc-wrap {
  position: relative;
  width: 100%;
  max-width: 620px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.gc-modal {
  width: 100%; 
  max-width: 820px;
  max-height: 85vh;
  background: #fff;
  border-radius: 12px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  box-shadow: 0 24px 60px rgba(0, 0, 0, 0.35);
}

.gc-header {
  position: relative;
  height: 120px;
  flex-shrink: 0;
  background: linear-gradient(135deg, #1B75D2 0%, #1B75D2 45%, #1B75D2 100%);
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

.gc-blob {
  position: absolute;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.16);
}

.gc-blob-1 {
  width: 190px;
  height: 90px;
  top: -40px;
  left: -60px;
  transform: rotate(-28deg);
}

.gc-blob-2 {
  width: 240px;
  height: 130px;
  top: -30px;
  right: -80px;
  border-radius: 50%;
}

.gc-blob-3 {
  width: 210px;
  height: 100px;
  bottom: -50px;
  right: -40px;
  border-radius: 999px;
  transform: rotate(-20deg);
}

.gc-dot {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.5);
}

.gc-dot-1 {
  width: 16px;
  height: 16px;
  top: 62px;
  left: 34%;
}

.gc-dot-2 {
  width: 10px;
  height: 10px;
  top: 42px;
  left: 46%;
}

.gc-dash {
  position: absolute;
  width: 26px;
  height: 9px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.5);
}

.gc-dash-1 {
  top: 32px;
  left: 40%;
  transform: rotate(-40deg);
}

.gc-dash-2 {
  top: 62px;
  left: 20%;
  transform: rotate(-40deg);
}

.gc-icon-card {
  position: relative;
  width: 108px;
  height: 96px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 14px 26px rgba(60, 0, 90, 0.28);
  overflow: hidden;
  z-index: 2;
}

.gc-icon-tab {
  height: 16px;
  background: #f4c752;
}

.gc-icon-body {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 12px;
}

.gc-icon-star {
  width: 36px;
  height: 36px;
  flex-shrink: 0;
  border-radius: 8px;
  background: #6a4bf0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.gc-icon-star svg {
  width: 20px;
  height: 20px;
  fill: #f4c752;
}

.gc-icon-lines {
  display: flex;
  flex-direction: column;
  gap: 6px;
  flex: 1;
}

.gc-icon-lines span {
  display: block;
  height: 4px;
  border-radius: 2px;
  background: #d9d9de;
}

.gc-icon-lines span:nth-child(2) {
  width: 70%;
}

.gc-icon-lines span:nth-child(3) {
  width: 45%;
}

.gc-body {
  padding: 24px 28px 8px;
  overflow-y: auto;
}

.gc-title {
  margin: 0 0 6px;
  text-align: center;
  font-size: 21px;
  font-weight: 800;
  color: #1a1a1a;
}

.gc-subtitle {
  margin: 0 0 22px;
  text-align: center;
  font-size: 13.5px;
  color: #6a6a6e;
  font-family: 'Inter', sans-serif;
  line-height: 1.5;
}

.gc-field {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 18px;
}

.gc-field-row {
  align-items: flex-start;
}

.gc-label {
  width: 92px;
  flex-shrink: 0;
  font-size: 12px;
  font-weight: 700;
  color: #1a1a1a;
  font-weight: 500;

}

.gc-label-access {
  display: flex;
  align-items: center;
  gap: 6px;
  padding-top: 8px;
}

.gc-help {
  width: 16px;
  height: 16px;
  color: #9a9a9e;
  display: inline-flex;
  cursor: help;
}

.gc-help svg {
  width: 100%;
  height: 100%;
  stroke: currentColor;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
}

.gc-input,
.gc-select,
.gc-textarea {
  width: 100%;
  border: 1px solid #d9d9de;
  border-radius: 12px;
  padding: 11px 14px;
  font-size: 14.5px;
  font-family: 'Inter', sans-serif;
  color: #000;
  outline: none;
  background: #fff;
  transition: border-color 0.15s ease;
}

.gc-textarea {
  resize: vertical;
  min-height: 70px;
  line-height: 1.5;
  align-self: flex-start;
}

.gc-field:has(.gc-textarea) {
  align-items: flex-start;
}

.gc-field:has(.gc-textarea) .gc-label {
  padding-top: 10px;
}

.gc-input:focus,
.gc-select:focus,
.gc-textarea:focus {
  border-color: #1B75D2;
}

.gc-select-wrap {
  position: relative;
  flex: 1;
}

.gc-select {
  appearance: none;
  -webkit-appearance: none;
  cursor: pointer;
  padding-right: 36px;
  color: #1a1a1a;
}

.gc-select:invalid {
  color: #9a9a9e;
}

.gc-select-chevron {
  position: absolute;
  top: 50%;
  right: 12px;
  transform: translateY(-50%);
  width: 16px;
  height: 16px;
  stroke: #6a6a6e;
  fill: none;
  stroke-width: 2;
  stroke-linecap: round;
  stroke-linejoin: round;
  pointer-events: none;
}

.gc-radio-group {
  display: flex;
  align-items: center;
  gap: 26px;
  flex-wrap: wrap;
  padding-top: 4px;
}

.gc-radio {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  font-size: 15px;
  color: #1a1a1a;
}

.gc-radio input {
  position: absolute;
  opacity: 0;
  width: 0;
  height: 0;
}

.gc-radio-dot {
  width: 20px;
  height: 20px;
  border-radius: 6px;
  border: 1px solid #d9d9de;
  position: relative;
  flex-shrink: 0;
  transition: border-color 0.15s ease;
}

.gc-radio-dot::after {
  content: '';
  position: absolute;
  inset: 3px;
  border-radius: 3px;
  background: #1B75D2;
  transform: scale(0);
  transition: transform 0.15s ease;
}

.gc-radio input:checked + .gc-radio-dot {
  border-color: #1B75D2;
}

.gc-radio input:checked + .gc-radio-dot::after {
  transform: scale(1);
}

.gc-radio input:focus-visible + .gc-radio-dot {
  outline: 2px solid #1B75D2;
  outline-offset: 2px;
}

.gc-footer {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 16px 28px;
  border-top: 1px solid #ececef;
}

.gc-footer-note {
  margin: 0;
  font-size: 12.5px;
  line-height: 1.4;
  color: #8a8a8e;
  font-family: 'Inter', sans-serif;
}

.gc-create-btn {
  flex-shrink: 0;
  border: none;
  border-radius: 32px;
  padding: 7px 12px;
  font-size: 15px;
  font-weight: 800;
  color: #fff;
  background: #1B75D2;
  cursor: pointer;
  transition: background 0.15s ease, opacity 0.15s ease;
}

.gc-create-btn:hover:not(:disabled) {
  background: #155ea8;
}

.gc-create-btn:disabled {
  background: #e7e7e7;
  color: #b5b5b8;
  cursor: not-allowed;
}

.wrapper-add {
  position: absolute;
  bottom: 2px;
  right: -34px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  z-index: 1001;
}

.wrapper-add-btn {
  width: 32px;
  height: 32px;
  border: none;
  background-color: transparent;
  color: #CCBFBD;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.wrapper-add-btn svg {
  width: 18px;
  height: 18px;
  stroke: currentColor;
  fill: none;
  stroke-width: 2.2;
  stroke-linecap: round;
}

/* ---------- Responsive ---------- */
@media (max-width: 480px) {
  .gc-body {
    padding: 20px 18px 4px;
  }
  .gc-footer {
    flex-direction: column;
    align-items: stretch;
    padding: 16px 18px;
  }
  .gc-create-btn {
    width: 100%;
  }
  .gc-field {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
  .gc-label {
    width: auto;
  }
  .gc-radio-group {
    gap: 18px;
  }
}
</style>