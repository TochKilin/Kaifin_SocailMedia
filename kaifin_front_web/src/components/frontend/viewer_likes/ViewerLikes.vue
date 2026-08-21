<template>
  <div
    class="engage-card">
    <div class="engage-list">
      <div
        v-for="(person, idx) in visibleUsers"
        :key="person.id"
        class="engage-row"
        :class="{ 'is-last': idx === visibleUsers.length - 1 }"
      >
        <div class="row-avatar">
          <img v-if="person.avatarUrl" :src="person.avatarUrl" alt="" />
          <svg v-else viewBox="0 0 24 24">
            <circle cx="12" cy="9" r="3.4" fill="none" stroke="currentColor" stroke-width="1.8"/>
            <path d="M5 20c0-3.9 3.1-6.5 7-6.5s7 2.6 7 6.5" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
          </svg>
        </div>
        <span class="row-name">{{ person.username }}</span>
        <button
          class="like-btn"
          :class="{ 'is-liked': likedIds.has(person.id) }"
          @click="toggleLike(person.id)"
          :aria-pressed="likedIds.has(person.id)"
          :aria-label="likedIds.has(person.id) ? 'Unlike' : 'Like'"
        >
          <svg viewBox="0 0 24 24" class="like-icon">
            <path
              d="M8 21H5.5A1.5 1.5 0 0 1 4 19.5v-7A1.5 1.5 0 0 1 5.5 11H8m0 10V11m0 10h9.1a2 2 0 0 0 2-1.7l1-6A2 2 0 0 0 18.1 11H13l.7-4.2c.2-1.2-.7-2.3-1.9-2.3-.3 0-.6.1-.8.4L8 11"
              fill="none"
              stroke="currentColor"
              stroke-width="1.8"
              stroke-linejoin="round"
              stroke-linecap="round"
            />
          </svg>
        </button>
      </div>
    </div>
    <button v-if="hasMore" class="see-more-btn" @click="showMore">
      <span>See more</span>
      <svg viewBox="0 0 24 24" class="see-more-icon">
        <path d="M6 10l6 6 6-6" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  users: {
    type: Array,
    default: () => ([
      { id: 1, username: 'Username', avatarUrl: '' },
      { id: 2, username: 'Username', avatarUrl: '' },
      { id: 3, username: 'Username', avatarUrl: '' },
      { id: 4, username: 'Username', avatarUrl: '' },
      { id: 5, username: 'Username', avatarUrl: '' },
    ]),
  },
  pageSize: {
    type: Number,
    default: 3,
  },
})

const emit = defineEmits(['like', 'keep-open', 'request-close'])

const visibleCount = ref(props.pageSize)
const likedIds = ref(new Set())

const visibleUsers = computed(() => props.users.slice(0, visibleCount.value))
const hasMore = computed(() => visibleCount.value < props.users.length)

function showMore() {
  visibleCount.value = Math.min(visibleCount.value + props.pageSize, props.users.length)
}

function toggleLike(id) {
  const next = new Set(likedIds.value)
  if (next.has(id)) {
    next.delete(id)
    emit('unlike', id)
  } else {
    next.add(id)
    emit('like', id)
  }
  likedIds.value = next
}


</script>

<style scoped>
* {
  box-sizing: border-box;
}

.engage-card {
  font-family: 'Inter', sans-serif;
  background: #ffff;
  border-radius: 8px;
  box-shadow: 0 1px 2px rgba(16, 24, 40, 0.04), 0 8px 24px rgba(16, 24, 40, 0.06);
  padding: 4px;
  width: 200px;
  
}

.engage-list {
  display: flex;
  flex-direction: column;
  pointer-events: auto;
}

.engage-row {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 4px 4px;
  border-bottom: 1px solid #F1F2F4;
  
  
}

.engage-row:hover {
  background: #F5F5F6;
  cursor: pointer;
  border-radius: 32px;
}


.engage-row.is-last {
  border-bottom: none;
}

.row-avatar {
  width: 25px;
  height: 25px;
  border-radius: 50%;
  background: #F3F4F6;
  color: #9CA3AF;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  overflow: hidden;
  border: 1.5px solid #2563EB;
}

.row-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.row-avatar svg {
  width: 12px;
  height: 12px;
}

.row-name {
  flex: 1;
  font-size: 12px;
  font-weight: 600;
  color: #1A1A1A;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.like-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 25px;
  border-radius:999px;
  border: none;
  background: #1976D2;
  color: #fff;
  cursor: pointer;
  flex-shrink: 0;
  transition: background 0.15s ease, color 0.15s ease, transform 0.15s ease;

}

.like-btn:hover {
  opacity: 0.8;
  color: #ffff;
}

.like-btn:active {
  transform: scale(0.9);
}

.like-btn.is-liked {
  background: #EEF4FF;
  color: #2563EB;
  animation: likePop 0.28s ease;
}

.like-btn.is-liked .like-icon {
  fill: #2563EB;
  stroke: #2563EB;
}

@keyframes likePop {
  0%   { transform: scale(1); }
  40%  { transform: scale(1.25); }
  100% { transform: scale(1); }
}

.like-icon {
  width: 16px;
  height: 16px;
}

.see-more-btn {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  margin-top: 4px;
  padding: 9px;
  border: none;
  border-radius: 10px;
  background: transparent;
  color: #4B5563;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.15s ease, color 0.15s ease;
}

.see-more-btn:hover {
  background: #F5F5F6;
  color: #1A1A1A;
}

.see-more-icon {
  width: 12px;
  height: 12px;
  transition: transform 0.2s ease;
}

.see-more-btn:hover .see-more-icon {
  transform: translateY(2px);
}

</style>