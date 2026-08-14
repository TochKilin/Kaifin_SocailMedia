<template>
  <div class="comment-section">
    <!-- SVG sprite: define once, reuse everywhere via <use> -->
    <svg width="0" height="0" style="position: absolute" aria-hidden="true">
      <defs>
        <symbol id="icon-image" viewBox="0 0 24 24">
          <rect x="3" y="4" width="18" height="16" rx="2.5" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <circle cx="8.5" cy="9.5" r="1.6" fill="currentColor"/>
          <path d="M4 16.5l5-5 3.5 3.5L17 10l3 3.5V18a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 4 18v-1.5z" fill="currentColor"/>
        </symbol>

        <symbol id="icon-at" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="9.25" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <circle cx="12" cy="12" r="3.6" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <path d="M15.6 10.6V13c0 1.4 1 2.3 2.1 2.3 1.6 0 2.4-1.7 2.4-3.6 0-1-.3-2-1-2.9" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
        </symbol>

        <symbol id="icon-smile" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="9.25" fill="none" stroke="currentColor" stroke-width="1.8"/>
          <circle cx="8.7" cy="10.2" r="1.1" fill="currentColor"/>
          <circle cx="15.3" cy="10.2" r="1.1" fill="currentColor"/>
          <path d="M7.8 14.2c1 1.5 2.6 2.3 4.2 2.3s3.2-.8 4.2-2.3" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>
        </symbol>

        <symbol id="icon-heart-outline" viewBox="0 0 24 24">
          <path d="M12 20s-7.4-4.6-9.7-9.2C.9 7.5 2.3 4 5.9 3.4c2-.3 3.9.6 5 2.3l1.1 1.6 1.1-1.6c1.1-1.7 3-2.6 5-2.3 3.6.6 5 4.1 3.6 7.4C19.4 15.4 12 20 12 20z" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round"/>
        </symbol>
        <symbol id="icon-heart-filled" viewBox="0 0 24 24">
          <path d="M12 20s-7.4-4.6-9.7-9.2C.9 7.5 2.3 4 5.9 3.4c2-.3 3.9.6 5 2.3l1.1 1.6 1.1-1.6c1.1-1.7 3-2.6 5-2.3 3.6.6 5 4.1 3.6 7.4C19.4 15.4 12 20 12 20z" fill="currentColor" stroke="currentColor" stroke-width="1.7" stroke-linejoin="round"/>
        </symbol>

        <symbol id="icon-clock" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="9.25" fill="none" stroke="currentColor" stroke-width="1.7"/>
          <path d="M12 7v5.3l3.6 2.1" fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round" stroke-linejoin="round"/>
        </symbol>

        <symbol id="icon-close" viewBox="0 0 24 24">
          <path d="M5.5 5.5l13 13m0-13l-13 13" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round"/>
        </symbol>
      </defs>
    </svg>

    <h3 class="panel-title">All comments · {{ totalCount }}</h3>

    <!-- New top-level comment composer -->
    <div class="composer-row">
      <div class="avatar" :style="{ background: me.color }">{{ me.initials }}</div>
      <div class="composer card">
        <textarea
          v-model="draft.text"
          maxlength="1000"
          placeholder="Add a comment..."
          rows="2"
          @input="autoGrow"
        ></textarea>

        <div v-if="draft.image" class="img-preview">
          <img :src="draft.image" alt="attachment preview" />
          <button @click="draft.image = null">
            <svg class="icon icon-xs"><use href="#icon-close" /></svg>
          </button>
        </div>

        <div class="composer-toolbar">
          <button class="tool-btn" title="Attach image" @click="triggerFile('draft')">
            <svg class="icon"><use href="#icon-image" /></svg>
          </button>
          <input
            type="file"
            accept="image/*"
            ref="fileInputDraft"
            style="display: none"
            @change="onFile($event, 'draft')"
          />

          <button class="tool-btn" title="Mention someone">
            <svg class="icon"><use href="#icon-at" /></svg>
          </button>

          <div class="emoji-wrap">
            <button class="tool-btn" title="Add emoji" @click="toggleEmoji('draft')">
              <svg class="icon"><use href="#icon-smile" /></svg>
            </button>
            <div class="emoji-pop" v-if="emojiOpenFor === 'draft'">
              <span v-for="e in emojis" :key="e" @click="insertEmoji(e, 'draft')">{{ e }}</span>
            </div>
          </div>

          <span class="char-count" :class="{ warn: draft.text.length > 900 }">
            {{ draft.text.length }}/1000
          </span>
          <button class="send-btn" :disabled="!draft.text.trim()" @click="postComment">
            Send
          </button>
        </div>
      </div>
    </div>

    <hr class="divider" />

    <!-- Comment list -->
    <div v-for="comment in comments" :key="comment.id">
      <div class="comment">
        <div class="avatar" :style="{ background: comment.color }">{{ comment.initials }}</div>
        <div class="comment-body">
          <div class="bubble">
            <div class="username">{{ comment.username }}</div>
            <div class="comment-text">{{ comment.text }}</div>
            <img v-if="comment.image" class="comment-img" :src="comment.image" alt="attachment" />
          </div>

          <div class="meta-row">
            <span class="meta-time">
              <svg class="icon icon-xs"><use href="#icon-clock" /></svg>
              {{ comment.time }}
            </span>
            <button class="action-btn" @click="toggleReply(comment.id)">Reply</button>
            <div class="spacer"></div>
            <button
              class="action-btn react-btn"
              :class="{ liked: comment.liked }"
              @click="toggleLike(comment)"
            >
              <svg class="icon icon-sm"><use :href="comment.liked ? '#icon-heart-filled' : '#icon-heart-outline'" /></svg>
              {{ comment.likes }}
            </button>
            <button class="action-btn react-btn" title="React">
              <svg class="icon icon-sm"><use href="#icon-smile" /></svg>
            </button>
          </div>

          <!-- Reply composer -->
          <div v-if="replyOpenFor === comment.id" class="reply-composer card">
            <textarea
              v-model="comment.replyDraft"
              maxlength="1000"
              placeholder="Write a reply..."
              rows="2"
              @input="autoGrow"
            ></textarea>

            <div v-if="comment.replyImage" class="img-preview">
              <img :src="comment.replyImage" alt="attachment preview" />
              <button @click="comment.replyImage = null">
                <svg class="icon icon-xs"><use href="#icon-close" /></svg>
              </button>
            </div>

            <div class="composer-toolbar">
              <button class="tool-btn" title="Attach image" @click="triggerReplyFile(comment.id)">
                <svg class="icon"><use href="#icon-image" /></svg>
              </button>
              <input
                type="file"
                accept="image/*"
                :ref="(el) => setReplyFileRef(el, comment.id)"
                style="display: none"
                @change="onFile($event, 'reply', comment)"
              />

              <button class="tool-btn" title="Mention someone">
                <svg class="icon"><use href="#icon-at" /></svg>
              </button>

              <div class="emoji-wrap">
                <button class="tool-btn" title="Add emoji" @click="toggleEmoji('reply-' + comment.id)">
                  <svg class="icon"><use href="#icon-smile" /></svg>
                </button>
                <div class="emoji-pop" v-if="emojiOpenFor === 'reply-' + comment.id">
                  <span
                    v-for="e in emojis"
                    :key="e"
                    @click="insertEmoji(e, 'reply', comment)"
                  >{{ e }}</span>
                </div>
              </div>

              <span class="char-count" :class="{ warn: (comment.replyDraft || '').length > 900 }">
                {{ (comment.replyDraft || '').length }}/1000
              </span>
              <button
                class="send-btn"
                :disabled="!(comment.replyDraft || '').trim()"
                @click="postReply(comment)"
              >
                Send
              </button>
            </div>
          </div>

          <!-- Replies -->
          <div v-if="comment.replies.length" class="replies">
            <div class="comment" v-for="reply in comment.replies" :key="reply.id">
              <div class="avatar" :style="{ background: reply.color }">{{ reply.initials }}</div>
              <div class="comment-body">
                <div class="bubble">
                  <div class="username">{{ reply.username }}</div>
                  <div class="comment-text">{{ reply.text }}</div>
                  <img v-if="reply.image" class="comment-img" :src="reply.image" alt="reply attachment" />
                </div>
                <div class="meta-row">
                  <span class="meta-time">
                    <svg class="icon icon-xs"><use href="#icon-clock" /></svg>
                    {{ reply.time }}
                  </span>
                  <div class="spacer"></div>
                  <button
                    class="action-btn react-btn"
                    :class="{ liked: reply.liked }"
                    @click="toggleLike(reply)"
                  >
                    <svg class="icon icon-sm"><use :href="reply.liked ? '#icon-heart-filled' : '#icon-heart-outline'" /></svg>
                    {{ reply.likes }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'

const me = { initials: 'ME', color: '#4C5FEA' }
const emojis = ['😀', '😂', '😍', '👍', '🎉', '🔥', '😮', '😢']

const emojiOpenFor = ref(null)
const replyOpenFor = ref(null)
const draft = reactive({ text: '', image: null })

const fileInputDraft = ref(null)
const replyFileRefs = {}

const comments = reactive([
  {
    id: 1,
    username: 'sokha_dev',
    initials: 'SD',
    color: '#F5A623',
    text: 'This wireframe is super clean — love how the reply box nests right under the comment.',
    time: '2h ago',
    likes: 4,
    liked: false,
    image: null,
    replyDraft: '',
    replyImage: null,
    replies: [
      {
        id: 11,
        username: 'dara_ux',
        initials: 'DU',
        color: '#2EC4B6',
        text: 'Agreed, the affordance for reply is obvious thanks to the arrow.',
        time: '1h ago',
        likes: 1,
        liked: false,
        image: null,
      },
    ],
  },
  {
    id: 2,
    username: 'lina_photo',
    initials: 'LP',
    color: '#EF476F',
    text: 'Here is a screenshot of the exported flow for reference.',
    time: '35m ago',
    likes: 2,
    liked: false,
    image: 'https://images.unsplash.com/photo-1587440871875-191322ee64b0?w=400&q=60',
    replyDraft: '',
    replyImage: null,
    replies: [],
  },
])

const totalCount = computed(() =>
  comments.reduce((sum, c) => sum + 1 + c.replies.length, 0)
)

function autoGrow(e) {
  e.target.style.height = 'auto'
  e.target.style.height = e.target.scrollHeight + 'px'
}

function setReplyFileRef(el, id) {
  if (el) replyFileRefs[id] = el
}

function triggerFile(key) {
  if (key === 'draft') fileInputDraft.value?.click()
}

function triggerReplyFile(id) {
  replyFileRefs[id]?.click()
}

function onFile(event, kind, comment) {
  const file = event.target.files[0]
  if (!file) return
  const reader = new FileReader()
  reader.onload = () => {
    if (kind === 'draft') draft.image = reader.result
    else if (kind === 'reply' && comment) comment.replyImage = reader.result
  }
  reader.readAsDataURL(file)
  event.target.value = ''
}

function toggleEmoji(key) {
  emojiOpenFor.value = emojiOpenFor.value === key ? null : key
}

function insertEmoji(e, kind, comment) {
  if (kind === 'draft') draft.text += e
  else if (kind === 'reply' && comment) comment.replyDraft = (comment.replyDraft || '') + e
  emojiOpenFor.value = null
}

function toggleReply(id) {
  replyOpenFor.value = replyOpenFor.value === id ? null : id
}

function toggleLike(item) {
  item.liked = !item.liked
  item.likes += item.liked ? 1 : -1
}

function postComment() {
  if (!draft.text.trim()) return
  comments.unshift({
    id: Date.now(),
    username: 'you',
    initials: me.initials,
    color: me.color,
    text: draft.text.trim(),
    time: 'just now',
    likes: 0,
    liked: false,
    image: draft.image,
    replyDraft: '',
    replyImage: null,
    replies: [],
  })
  draft.text = ''
  draft.image = null
}

function postReply(comment) {
  if (!(comment.replyDraft || '').trim()) return
  comment.replies.push({
    id: Date.now(),
    username: 'you',
    initials: me.initials,
    color: me.color,
    text: comment.replyDraft.trim(),
    time: 'just now',
    likes: 0,
    liked: false,
    image: comment.replyImage,
  })
  comment.replyDraft = ''
  comment.replyImage = null
  replyOpenFor.value = null
}
</script>

<style scoped>
:root {
  --ink: #1e2430;
  --ink-soft: #5b6472;
  --line: #e4e7ec;
  --paper: #ffffff;
  --canvas: #f5f6f8;
  --accent: #4c5fea;
  --accent-soft: #eef0fd;
  --heart: #ef476f;
}

.comment-section {
  font-family: 'Inter', 'Segoe UI', sans-serif;
  color: var(--ink);
  background: var(--canvas);
  padding: 28px;
  border-radius: 20px;
  max-width: 720px;
  margin: 0 auto;
}

.panel-title {
  font-family: 'Sora', 'Segoe UI', sans-serif;
  font-size: 15px;
  font-weight: 700;
  letter-spacing: 0.02em;
  text-transform: uppercase;
  margin: 0 0 16px 2px;
}

.card {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 16px;
  padding: 16px;
}

.composer-row {
  display: flex;
  gap: 12px;
  margin-bottom: 22px;
}

.avatar {
  width: 42px;
  height: 42px;
  border-radius: 50%;
  flex: none;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: 'Sora', 'Segoe UI', sans-serif;
  font-weight: 700;
  font-size: 14px;
  color: #fff;
}

.composer {
  flex: 1;
}

textarea {
  width: 100%;
  border: none;
  resize: none;
  font-family: 'Inter', 'Segoe UI', sans-serif;
  font-size: 14.5px;
  line-height: 1.5;
  color: var(--ink);
  outline: none;
  min-height: 46px;
  background: transparent;
}

textarea::placeholder {
  color: #9aa2ae;
}

.composer-toolbar {
  display: flex;
  align-items: center;
  gap: 14px;
  padding-top: 10px;
  margin-top: 8px;
  border-top: 1px solid var(--line);
}

.tool-btn {
  background: none;
  border: none;
  cursor: pointer;
  width: 30px;
  height: 30px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--ink-soft);
}

.icon {
  width: 18px;
  height: 18px;
  flex: none;
  display: block;
}

.icon-sm {
  width: 16px;
  height: 16px;
}

.icon-xs {
  width: 13px;
  height: 13px;
}

.tool-btn:hover {
  background: var(--accent-soft);
  color: var(--accent);
}

.char-count {
  margin-left: auto;
  font-size: 12.5px;
  color: #b0b6c0;
  font-variant-numeric: tabular-nums;
}

.char-count.warn {
  color: var(--heart);
}

.send-btn {
  background: var(--accent);
  color: #fff;
  border: none;
  font-family: 'Sora', 'Segoe UI', sans-serif;
  font-weight: 600;
  font-size: 13.5px;
  padding: 8px 18px;
  border-radius: 999px;
  cursor: pointer;
  transition: transform 0.12s ease, box-shadow 0.12s ease;
}

.send-btn:hover:not(:disabled) {
  box-shadow: 0 4px 14px rgba(76, 95, 234, 0.35);
  transform: translateY(-1px);
}

.send-btn:disabled {
  background: #c9cedA;
  cursor: not-allowed;
}

.img-preview {
  margin-top: 10px;
  position: relative;
  display: inline-block;
}

.img-preview img {
  max-height: 90px;
  border-radius: 10px;
  display: block;
}

.img-preview button {
  position: absolute;
  top: -7px;
  right: -7px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: var(--ink);
  color: #fff;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
}

.emoji-wrap {
  position: relative;
}

.emoji-pop {
  position: absolute;
  bottom: 36px;
  left: 0;
  background: #fff;
  border: 1px solid var(--line);
  border-radius: 12px;
  padding: 8px;
  display: flex;
  gap: 4px;
  box-shadow: 0 8px 24px rgba(20, 24, 40, 0.12);
  z-index: 5;
}

.emoji-pop span {
  cursor: pointer;
  font-size: 18px;
  padding: 4px;
  border-radius: 6px;
}

.emoji-pop span:hover {
  background: var(--accent-soft);
}

.divider {
  border: none;
  border-top: 1px solid var(--line);
  margin: 22px 0;
}

.comment {
  display: flex;
  gap: 12px;
  margin-bottom: 22px;
}

.comment-body {
  flex: 1;
  min-width: 0;
}

.bubble {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 14px;
  padding: 12px 14px;
}

.username {
  font-family: 'Sora', 'Segoe UI', sans-serif;
  font-weight: 700;
  font-size: 13.5px;
  margin-bottom: 4px;
}

.comment-text {
  font-size: 14px;
  line-height: 1.5;
  color: var(--ink);
}

.comment-img {
  margin-top: 8px;
  max-width: 220px;
  border-radius: 10px;
  display: block;
}

.meta-row {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-top: 8px;
  padding-left: 4px;
}

.meta-time {
  font-size: 12px;
  color: #9aa2ae;
  display: inline-flex;
  align-items: center;
  gap: 5px;
}

.action-btn {
  background: none;
  border: none;
  cursor: pointer;
  font-family: 'Inter', 'Segoe UI', sans-serif;
  font-weight: 600;
  font-size: 12.5px;
  color: var(--ink-soft);
  padding: 2px 4px;
  display: inline-flex;
  align-items: center;
  gap: 5px;
}

.action-btn:hover {
  color: var(--accent);
}

.action-btn.liked {
  color: var(--heart);
}

.spacer {
  flex: 1;
}

.react-btn {
  font-size: 15px;
}

.reply-composer {
  margin-top: 12px;
}

.replies {
  margin-top: 14px;
  padding-left: 20px;
  border-left: 2px solid var(--line);
}

.replies .comment {
  margin-bottom: 16px;
}

.replies .avatar {
  width: 34px;
  height: 34px;
  font-size: 12px;
}
</style>