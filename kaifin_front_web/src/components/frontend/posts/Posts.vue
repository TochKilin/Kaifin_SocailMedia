<template>
  <!-- Compost posts -->
  <div class="composer">
    <div class="pin-row">
      <svg class="pin-icon" viewBox="0 0 24 24"><path d="M12 2a1 1 0 0 1 1 1v3.6l4 4V13h-4.2l-.8 7-1-1-1 1-.8-7H6v-2.4l4-4V3a1 1 0 0 1 1-1Z"/></svg>
      <input v-model="hashtagPin" class="pin-input" type="text" placeholder="Hashtag pin" @keydown.enter.prevent />
    </div>
    <textarea ref="textareaRef" v-model="postText" class="post-text" placeholder="Whate your mind for posts?" rows="1" @input="autoGrow" >
    </textarea>
    <div class="selected-stickers" v-if="selectedStickers.length">
      <div class="selected-sticker" v-for="(s,index) in selectedStickers" :key="index">
    <img :src="s.url" :alt="s.file_name">
    <button @click="selectedStickers.splice(index,1)" >
      ✕
    </button>
  </div>
  </div>

  <!-- Code -->
  <div class="code-wrap" v-if="showCode">
    <div class="code-head">
      <span>{ } Code</span>
      <button class="code-remove" @click="showCode = false; codeText = ''" aria-label="លុប code">✕</button>
    </div>
    <textarea v-model="codeText" class="code-text" placeholder="paste or write code..." rows="4" ></textarea>
  </div>

    <div class="link-chip" v-if="linkUrl">
      <svg viewBox="0 0 24 24"><path d="M10 14a4 4 0 0 0 5.66 0l2.34-2.34a4 4 0 1 0-5.66-5.66L11 7"/><path d="M14 10a4 4 0 0 0-5.66 0L6 12.34a4 4 0 1 0 5.66 5.66L13 17"/></svg>
      <span class="link-text">{{ linkUrl }}</span>
      <button @click="linkUrl = ''" aria-label="Delete link">✕</button>
    </div>

    <div class="image-grid" v-if="images.length">
      <div class="image-tile" v-for="(img, i) in images" :key="i">
        <img :src="img.url" alt="" />
        <button class="image-remove" @click="removeImage(i)" aria-label="Delete image">✕</button>
      </div>
    </div>
    <!-- Video  -->
    <div class="video-grid" v-if="videos.length">
      <div class="video-tile"  v-for="(video,i) in videos" :key="i">
        <video :src="video.url" controls></video>
        <button class="image-remove" @click="videos.splice(i,1)" >
          ✕
        </button>
      </div>
    </div>
    <!-- topic  -->
<span class="topic-chip" v-for="(t, i) in topics" :key="i">
  <img v-if="t.stickerUrl" :src="t.stickerUrl" class="topic-sticker-badge" alt="" />
  <span v-else-if="t.icon" class="topic-icon-badge">{{ t.icon }}</span>
  #{{ t.text }}
  <button @click="topics.splice(i, 1)" aria-label="Delete topic">✕</button>
</span>

    <div class="group-choose-wrap">
      <button class="group-choose" @click="showGroupPicker = !showGroupPicker">
        <svg viewBox="0 0 24 24"><rect x="3" y="3" width="18" height="18" rx="3"/><circle cx="8.5" cy="9.5" r="1.7"/><path d="m4 17 5-5 4 4 3-3 4 4"/></svg>
        <span>{{ selectedGroup ? selectedGroup.name : 'Choose Group' }}</span>
        <svg class="chevron" viewBox="0 0 24 24"><path d="M6 9l6 6 6-6"/></svg>
      </button>

      <div class="group-picker" v-if="showGroupPicker" @click.stop @mousedown.stop>
        <button v-for="g in groups" :key="g.id" class="group-option" :class="{ selected: selectedGroup?.id === g.id }" @click="selectedGroup = g; showGroupPicker = false">
          {{ g.name }}
        </button>
      </div>
    </div>

    <div class="action-bar">
      <div class="action-left">
        <div class="chip-wrap" ref="pickerRef">
          <button class="chip" :class="{ active: showEmoji }" @click="showEmoji = !showEmoji">
            <svg viewBox="0 0 24 24"><circle cx="12" cy="12" r="9"/><path d="M8.5 10h.01M15.5 10h.01"/><path d="M8 14.5c1 1.3 2.4 2 4 2s3-.7 4-2"/></svg>
            Emoji
          </button>

         

     
             <div class="picker-panel" v-if="showEmoji" @mousedown.stop @click.stop>
            <div class="picker-tabs">
              <button type="button" class="picker-tab" :class="{ active: pickerTab === 'stickers' }" @click.stop="pickerTab = 'stickers'" >
              <!-- Stickers -->
              <svg viewBox="0 0 64 64" xmlns="http://www.w3.org/2000/svg" width="64" height="64">
                <path d="M12 8 H44 A8 8 0 0 1 52 16 V44 L38 58 H12 A8 8 0 0 1 4 50 V16 A8 8 0 0 1 12 8 Z" 
                      fill="#FAC775" stroke="#BA7517" stroke-width="1.5"/>
                <path d="M52 44 H42 A4 4 0 0 0 38 48 V58 Z" 
                      fill="#EF9F27" stroke="#BA7517" stroke-width="1.5" stroke-linejoin="round"/>
                <circle cx="20" cy="26" r="3" fill="#412402"/>
                <circle cx="36" cy="26" r="3" fill="#412402"/>
                <path d="M18 36 Q28 46 38 36" stroke="#412402" stroke-width="3" fill="none" stroke-linecap="round"/>
              </svg>
             </button>

              <!-- Emoji -->
              <button type="button" class="picker-tab" :class="{ active: pickerTab === 'emojis' }" @click.stop="pickerTab = 'emojis'" >
             <svg viewBox="0 0 64 64" xmlns="http://www.w3.org/2000/svg" width="64" height="64">
                <circle cx="32" cy="32" r="28" fill="#FAC775" stroke="#BA7517" stroke-width="2"/>
                <circle cx="22" cy="26" r="3.5" fill="#412402"/>
                <circle cx="42" cy="26" r="3.5" fill="#412402"/>
                <path d="M18 38 Q32 52 46 38" stroke="#412402" stroke-width="3.5" fill="none" stroke-linecap="round"/>
              </svg>
            </button>
            </div>
            <div class="picker-body">

            <template v-if="pickerTab === 'stickers'">
            <!-- All -->
            <template v-if="stickerFilterTab === 'all'">
              <div class="sticker-grid" v-if="filteredStickers.length">
                <button type="button" class="sticker-item" v-for="s in filteredStickers" :key="s.id" @click.stop="addSticker(s)">
                  <img :src="s.url" :alt="s.file_name" />
                </button>
              </div>
              <div class="picker-empty" v-else>
                No sticker Sticker...
              </div>
            </template>

            <template v-else-if="stickerFilterTab === 'mine'">
              <div class="sticker-set-list" v-if="mySetsStickers.length">
                <div class="sticker-set-row" v-for="set in mySetsStickers" :key="set.id">
                  <button type="button" class="set-drag-handle" aria-label="ផ្លាស់ទីលំដាប់">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
                      <line x1="4" y1="7" x2="20" y2="7"/>
                      <line x1="4" y1="12" x2="20" y2="12"/>
                      <line x1="4" y1="17" x2="20" y2="17"/>
                    </svg>
                  </button>

                  <div class="set-stickers-scroll">
                    <button type="button" class="set-sticker-item" v-for="s in set.stickers" :key="s.id" @click.stop="addSticker(s)">
                      <img :src="s.url" alt="" />
                    </button>
                  </div>

                  <button type="button" class="set-delete-btn" @click.stop="deleteStickerSet(set.id)">
                    Delete Set
                  </button>
                </div>
              </div>
              <div class="picker-empty" v-else>
                No sticker set...
              </div>
            </template>

          <template v-else-if="stickerFilterTab === 'animated'">
            <div class="pack-grid" v-if="!showPackStickers && currentPackList.length">
              <div class="pack-card" v-for="pack in currentPackList" :key="pack.id" @click="openPack(pack)">
                <div class="pack-thumb">
                  <img :src="pack.thumbnail_url" :alt="pack.name" />
                  <span class="pack-count-badge">{{ pack.sticker_count }}</span>
                </div>
                <div class="pack-footer">
                  <span class="pack-name">{{ pack.name }}</span>
                  <button type="button" class="pack-add-btn" @click.stop="addPackToCollection(pack)">ADD</button>
                </div>
              </div>
            </div>
            <div class="picker-empty" v-else-if="!showPackStickers">
              No sticker pack...
            </div>
            <template v-else>
              <div class="pack-detail-head">
                <button type="button" class="pack-back-btn" @click.stop="backToPacks">‹ Back</button>
                <span class="pack-detail-title">{{ activePack?.name }}</span>
              </div>
              <div class="sticker-grid" v-if="stickers.length">
                <button type="button" class="sticker-item" v-for="s in stickers" :key="s.id" @click.stop="addSticker(s)">
                  <img :src="s.url" :alt="s.file_name" />
                </button>
              </div>
              <div class="picker-empty" v-else>
                No sticker Sticker...
              </div>
            </template>

          </template>

          <template v-else-if="stickerFilterTab === 'create'">
            <div class="create-sticker-panel">
              <button type="button" class="add-photo-circle" @click.stop="addStickerPack">
                <img src="../../../assets/animate/gallary_1.svg" alt="svg">
              </button>
              <button type="button" class="add-photo-btn" @click="addStickerPack">
                Add Stickers
              </button>
              <input ref="stickerFileInputRef" type="file" accept="image/*" hidden @change="onStickerFilePicked" />
            </div>
          </template>
          <template v-else-if="stickerFilterTab === 'create'">

          <div class="create-sticker-panel">
            <button type="button" class="add-photo-circle" @click.stop="addStickerPack">
              <img src="../../../assets/animate/gallary_1.svg" alt="svg">
            </button>
            <button type="button" class="add-photo-btn" @click.stop="addStickerPack">
              Add Stickers
            </button>
            <input ref="stickerFileInputRef" type="file" accept="image/*" hidden @change="onStickerFilePicked" />
          </div>
        </template>

        </template>
        <!-- Emoji tab -->
        <template v-else-if="pickerTab === 'emojis'">
          <div class="emoji-grid">
            <button type="button" v-for="e in emojiList" :key="e" @click="addEmoji(e)">{{ e }}</button>
          </div>
        </template>
      </div>



      <div class="picker-footer">
        <div class="sticker-filter-row">
      <button type="button" class="filter-chip" :class="{ active: stickerFilterTab === 'all' }" @click.stop="stickerFilterTab = 'all'">
        <!-- <img src="../../../assets/animate/more.svg" alt=""> -->
         <svg class="filter-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <rect x="3" y="3" width="7" height="7"/>
        <rect x="14" y="3" width="7" height="7"/>
        <rect x="14" y="14" width="7" height="7"/>
        <rect x="3" y="14" width="7" height="7"/>
      </svg>
        <span>All</span>
      </button>
      <button type="button" class="filter-chip" :class="{ active: stickerFilterTab === 'animated' }" @click.stop="stickerFilterTab = 'animated'">
        <svg class="filter-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <polygon points="5 3 19 12 5 21 5 3"/>
      </svg>
        Animated
      </button>
      <button type="button" class="filter-chip" :class="{ active: stickerFilterTab === 'mine' }" @click.stop="stickerFilterTab = 'mine'">
       <svg class="filter-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/>
      </svg>
      <span>My stickers</span>
      </button>
      <button type="button" class="filter-chip filter-chip-create" :class="{ active: stickerFilterTab === 'create' }" @click.stop="stickerFilterTab = 'create'">
       <svg class="filter-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <line x1="12" y1="5" x2="12" y2="19"/>
        <line x1="5" y1="12" x2="19" y2="12"/>
      </svg>
      <span>Create Sticker</span>
        <!-- <input ref="stickerFileInputRef" type="file" accept="image/*" hidden @change="onStickerFilePicked" /> -->
      </button>
      </div>  
      </div>
      </div>
      </div>
      <div class="chip-wrap" ref="pickerRef">
        <button class="chip" :class="{ active: showPictureMenu }" @click="showPictureMenu = !showPictureMenu">
          <svg viewBox="0 0 24 24"><rect x="3" y="4" width="18" height="16" rx="2"/><circle cx="9" cy="10" r="1.6"/><path d="m4 18 5-5 4 4 3-3 4 4"/></svg>
          Picture
          <svg class="chevron-mini" viewBox="0 0 24 24"><path d="M6 9l6 6 6-6"/></svg>
        </button>

        <div class="picture-dropdown" v-if="showPictureMenu" @click.stop @mousedown.stop>
          <button type="button" class="picture-dropdown-item" @click="handlePictureSelect">
            <svg class="picture-dropdown-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
              <rect x="3" y="4" width="18" height="16" rx="2.5"/>
              <circle cx="9" cy="10" r="1.8"/>
              <path d="m4 18 5-5 4 4 3-3 4 4"/>
            </svg>
            <span>Picture</span>
          </button>
       <!-- <input ref="fileInputRef" type="file" accept="image/*,video/*" multiple hidden @change="onFilesPicked" /> -->
        <button type="button" class="picture-dropdown-item" @click="handleTemplateSelect">
          <svg class="picture-dropdown-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="3" width="7.5" height="7.5" rx="1.5"/>
            <rect x="13.5" y="3" width="7.5" height="7.5" rx="1.5"/>
            <rect x="3" y="13.5" width="7.5" height="7.5" rx="1.5"/>
            <rect x="13.5" y="13.5" width="7.5" height="7.5" rx="1.5"/>
          </svg>
          <span>Template</span>
        </button>
        </div>
        </div>
      <input ref="fileInputRef" type="file" accept="image/*,video/*" multiple hidden @change="onFilesPicked" />
        <div class="chip-wrap">
          <button class="chip" :class="{ active: showLinkInput }" @click="showLinkInput = !showLinkInput">
            <svg viewBox="0 0 24 24"><path d="M10 14a4 4 0 0 0 5.66 0l2.34-2.34a4 4 0 1 0-5.66-5.66L11 7"/><path d="M14 10a4 4 0 0 0-5.66 0L6 12.34a4 4 0 1 0 5.66 5.66L13 17"/></svg>
            Link
          </button>
          <div class="mini-input" v-if="showLinkInput" @click.stop @mousedown.stop>
            <input v-model="linkDraft" type="text" placeholder="https://..." @keydown.enter.prevent="confirmLink"/>
            <button type="button" @click="confirmLink">✓</button>
          </div>
        </div>
        <div class="chip-wrap">
          <button class="chip" :class="{ active: showTopicInput }" @click="showTopicInput = !showTopicInput">
            <span class="hash-glyph">#</span>
            Topic
          </button>
          <div class="mini-input" v-if="showTopicInput" @click.stop @mousedown.stop>
            <!-- Emoji -->
            <button type="button" class="topic-emoji-btn" @click.stop="showTopicEmojiPicker = !showTopicEmojiPicker; showTopicStickerPicker = false">
              <span v-if="topicEmojiDraft">{{ topicEmojiDraft }}</span>
              <svg v-else viewBox="0 0 24 24"><circle cx="12" cy="12" r="9"/><path d="M8.5 10h.01M15.5 10h.01"/><path d="M8 14.5c1 1.3 2.4 2 4 2s3-.7 4-2"/></svg>
            </button>
            <button type="button" class="topic-emoji-btn" @click.stop="showTopicStickerPicker = !showTopicStickerPicker; showTopicEmojiPicker = false">
              <img v-if="topicStickerDraft" :src="topicStickerDraft.url" class="topic-sticker-badge" alt="" />
              <svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="3" y="3" width="18" height="18" rx="3"/>
                <circle cx="8.5" cy="8.5" r="1.5"/>
                <path d="m21 15-5-5L5 21"/>
              </svg>
            </button>

            <input v-model="topicDraft" type="text" placeholder="Topic name" @keydown.enter.prevent="confirmTopic"/>
            <button type="button" @click="confirmTopic">✓</button>

            <div class="topic-emoji-popover" v-if="showTopicEmojiPicker" @click.stop @mousedown.stop>
              <button type="button" v-for="e in emojiList" :key="e" @click="topicEmojiDraft = e; topicStickerDraft = null; showTopicEmojiPicker = false">{{ e }}</button>
            </div>
            <div class="topic-emoji-popover topic-sticker-popover" v-if="showTopicStickerPicker" @click.stop @mousedown.stop>
              <button type="button" v-for="s in stickers" :key="s.id" @click="topicStickerDraft = s; topicEmojiDraft = ''; showTopicStickerPicker = false">
                <img :src="s.url" :alt="s.file_name" />
              </button>
            </div>
          </div>
        </div>

        <button class="chip" :class="{ active: showCode }" @click="showCode = !showCode">
          <svg viewBox="0 0 24 24"><path d="m8 6-5 6 5 6M16 6l5 6-5 6"/></svg>
          Code
        </button>
      </div>

      <button class="post-btn" :disabled="!canPost" @click="submitPost">
        Posts
        <svg width="512" height="512" viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M470.832 41.168C459.712 30.048 441.792 27.2 427.392 34.112L49.152 216.512C33.632 223.936 27.648 242.816 35.584 258.304C38.656 264.384 43.424 269.152 49.504 272.224L165.76 330.368L220.16 461.376C226.752 477.248 245.504 484.416 261.376 477.824C268.416 474.88 274.304 469.728 278.208 463.168L477.888 84.608C484.8 70.208 481.952 52.288 470.832 41.168ZM201.76 317.056L102.464 267.328L425.472 73.088L201.76 317.056ZM244.608 438.336L201.76 335.744L307.2 222.72C313.6 215.872 313.28 204.992 306.432 198.592C299.584 192.192 288.704 192.512 282.304 199.36L169.6 318.528L78.688 273.664L445.696 102.72L244.608 438.336Z" fill="#FFFFFF"/>
        </svg>
      </button>
    </div>

    <Teleport to="body">
      <div class="template-modal-overlay" v-if="showTemplateModal" @click.self="showTemplateModal = false">
        <div class="template-modal">
          <div class="template-modal-head">
            <span class="template-modal-title">{{ selectedTemplate ? 'Put your images' : 'Choose Template' }}</span>
            <button type="button" class="template-modal-close" @click="showTemplateModal = false">✕</button>
          </div>

          <div class="template-modal-body">
            <div class="template-grid" v-if="!selectedTemplate && templateList.length">
              <button type="button" class="template-item" v-for="tpl in templateList" :key="tpl.id" @click="selectTemplatePreview(tpl)">
                <img :src="tpl.thumbnail_url" :alt="tpl.name" />
                <span class="template-item-name">{{ tpl.name }}</span>
              </button>
            </div>
            <div class="picker-empty" v-else-if="!selectedTemplate">
              No Template
            </div>
            <div class="frame-preview-wrap" v-else>
              <canvas
                ref="frameCanvasRef"
                class="frame-canvas"
                :class="{ draggable: userPhotoImg }"
                @mousedown="onCanvasPointerDown"
                @mousemove="onCanvasPointerMove"
                @mouseup="onCanvasPointerUp"
                @mouseleave="onCanvasPointerUp"
                @touchstart.prevent="onCanvasPointerDown"
                @touchmove.prevent="onCanvasPointerMove"
                @touchend="onCanvasPointerUp"
              ></canvas>

              <div class="scale-control" v-if="userPhotoImg">
                <button type="button" class="scale-reset-btn" @click="resetPhotoPosition">Reset</button>
                <input
                  type="range"
                  min="0.3"
                  max="3"
                  step="0.01"
                  v-model.number="photoScale"
                  class="scale-slider"
                />
                <span class="scale-label">{{ Math.round(photoScale * 100) }}%</span>
              </div>

                <button type="button" class="add-photo-btn" @click="triggerUserPhotoInput">
                  {{ userPhotoImg ? 'Change images' : 'Choose your images' }}
                </button>
                <input ref="userPhotoInputRef" type="file" accept="image/*" hidden @change="onUserPhotoPicked" />
              </div>
             </div>
              <div class="template-modal-footer">
                <button type="button" class="template-cancel-btn" @click="showTemplateModal = false">Cancel</button>
                <button type="button" class="template-apply-btn" :disabled="!userPhotoImg" @click="applyTemplate">
                  Use Template
                </button>
              </div>
        </div>
      </div>
    </Teleport>

  </div>
</template>

<script setup>
import { ref, computed, nextTick, onMounted, onBeforeUnmount, Teleport, watch } from 'vue'
const emit = defineEmits(['post'])
const hashtagPin = ref('')
const postText = ref('')
const textareaRef = ref(null)
const images = ref([])
const videos = ref([])
const fileInputRef = ref(null)
const linkUrl = ref('')
const linkDraft = ref('')
const showLinkInput = ref(false)
const topics = ref([])
const topicDraft = ref('')
const showTopicInput = ref(false)
const showCode = ref(false)
const codeText = ref('')
const showEmoji = ref(false)
const pickerRef = ref(null)
const pickerTab = ref('stickers') 

// =======
const topicEmojiDraft = ref('')
const showTopicEmojiPicker = ref(false)
function splitTagIcon(tag) {
  const m = tag.match(/^(\p{Extended_Pictographic}\uFE0F?)/u)
  if (m) return { icon: m[1], text: tag.slice(m[1].length) }
  return { icon: '', text: tag }
}

const topicStickerDraft = ref(null)
const showTopicStickerPicker = ref(false)

// const emojiList = ['😀', '😂', '😍', '🥳', '😎', '🤔', '😢', '😡', '👍', '🙏', '🔥', '🎉', '❤️', '👏', '💡', '✅']
const emojiList = [
  // 😀 Faces
  '😀', '😃', '😄', '😁', '😆',
  '😅', '😂', '🤣', '😊', '😇',
  '😖', '😫', '😩', '🥺', '😢',
  '😭', '😤', '😠', '😡', '🤬',

  // ❤️ Hearts & Love
  '❤️', '🧡', '💛', '💚', '💙',
  '💜', '🖤', '🤍', '🤎', '💔',,

  // 👍 Hands
  '👍', '👎', '👌', '✌️', '🤞',

  // 🎉 Fun & Objects
  '🔥', '🎉', '🎊', '✨', '⭐',
  '🌟', '💫', '💥', '💯', '🎯',
  '🎁', '🎈', '🚀', '💡', '✅',
  '❌', '⚡', '🌈', '☀️', '🌙'
];
const stickers = ref([])
const selectedStickers = ref([])
const stickerFileInputRef = ref(null)
const isUploadingSticker = ref(false)
const newPackName = ref('')

const showTemplateModal = ref(false)
const selectedTemplate = ref(null)

const templateList = ref([])
const frameCanvasRef = ref(null)
const userPhotoInputRef = ref(null)
const userPhotoImg = ref(null)

const photoScale = ref(1)
const photoOffsetX = ref(0)
const photoOffsetY = ref(0)
const baseScale = ref(1)

let frameNaturalWidth = 0
let frameNaturalHeight = 0
const isDragging = ref(false)
let dragStartX = 0
let dragStartY = 0
let dragStartOffsetX = 0
let dragStartOffsetY = 0

async function fetchTemplates() {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/templates/show`, {
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    })
    const data = await res.json()
    const list = data?.data?.templates || []
    templateList.value = await Promise.all(
      list.map(async t => ({ ...t, thumbnail_url: await loadStickerImage(`${API_BASE}${t.thumbnail_url}`) }))
    )
  } catch (err) {
    console.error('Failed to load templates', err)
  }
}

function selectTemplatePreview(tpl) {
  selectedTemplate.value = tpl
  userPhotoImg.value = null
  photoScale.value = 1
  photoOffsetX.value = 0
  photoOffsetY.value = 0
}

function triggerUserPhotoInput() {
  userPhotoInputRef.value?.click()
}

function onUserPhotoPicked(e) {
  const file = e.target.files?.[0]
  if (!file) return
  e.target.value = ''

  const img = new Image()
  img.onload = () => {
    userPhotoImg.value = img
    photoScale.value = 1
    photoOffsetX.value = 0
    photoOffsetY.value = 0
    nextTick(() => drawComposite())
  }
  img.src = URL.createObjectURL(file)
}

async function drawComposite() {
  if (!selectedTemplate.value || !userPhotoImg.value || !frameCanvasRef.value) return

  const canvas = frameCanvasRef.value
  const ctx = canvas.getContext('2d')

  if (!frameNaturalWidth) {
    const frameImg = new Image()
    frameImg.crossOrigin = 'anonymous'
    await new Promise((resolve, reject) => {
      frameImg.onload = resolve
      frameImg.onerror = reject
      frameImg.src = selectedTemplate.value.thumbnail_url
    })
    frameNaturalWidth = frameImg.naturalWidth
    frameNaturalHeight = frameImg.naturalHeight
  }

  const MAX_DIMENSION = 1080
  const frameRatio = frameNaturalWidth / frameNaturalHeight

  let canvasWidth, canvasHeight
  if (frameRatio >= 1) {
    canvasWidth = MAX_DIMENSION
    canvasHeight = MAX_DIMENSION / frameRatio
  } else {
    canvasHeight = MAX_DIMENSION
    canvasWidth = MAX_DIMENSION * frameRatio
  }

  canvas.width = canvasWidth
  canvas.height = canvasHeight

  ctx.fillStyle = '#ffffff'
  ctx.fillRect(0, 0, canvas.width, canvas.height)

  const userImg = userPhotoImg.value
  baseScale.value = Math.max(canvas.width / userImg.width, canvas.height / userImg.height)

  const finalScale = baseScale.value * photoScale.value
  const drawWidth = userImg.width * finalScale
  const drawHeight = userImg.height * finalScale

  const x = (canvas.width - drawWidth) / 2 + photoOffsetX.value
  const y = (canvas.height - drawHeight) / 2 + photoOffsetY.value

  ctx.drawImage(userImg, x, y, drawWidth, drawHeight)

  const frameImg2 = new Image()
  frameImg2.crossOrigin = 'anonymous'
  frameImg2.src = selectedTemplate.value.thumbnail_url
  if (frameImg2.complete) {
    ctx.drawImage(frameImg2, 0, 0, canvas.width, canvas.height)
  } else {
    await new Promise(resolve => { frameImg2.onload = resolve })
    ctx.drawImage(frameImg2, 0, 0, canvas.width, canvas.height)
  }
}

watch([photoScale, photoOffsetX, photoOffsetY], () => {
  drawComposite()
})

function onCanvasPointerDown(e) {
  if (!userPhotoImg.value) return
  isDragging.value = true
  const point = e.touches ? e.touches[0] : e
  dragStartX = point.clientX
  dragStartY = point.clientY
  dragStartOffsetX = photoOffsetX.value
  dragStartOffsetY = photoOffsetY.value
}

function onCanvasPointerMove(e) {
  if (!isDragging.value) return
  const point = e.touches ? e.touches[0] : e
  const canvas = frameCanvasRef.value
  const rect = canvas.getBoundingClientRect()
  const scaleFactor = canvas.width / rect.width

  const dx = (point.clientX - dragStartX) * scaleFactor
  const dy = (point.clientY - dragStartY) * scaleFactor

  photoOffsetX.value = dragStartOffsetX + dx
  photoOffsetY.value = dragStartOffsetY + dy
}

function onCanvasPointerUp() {
  isDragging.value = false
}

function resetPhotoPosition() {
  photoScale.value = 1
  photoOffsetX.value = 0
  photoOffsetY.value = 0
}

async function applyTemplate() {
  const canvas = frameCanvasRef.value
  if (!canvas || !userPhotoImg.value) return

  canvas.toBlob(blob => {
    const file = new File([blob], `template-${Date.now()}.png`, { type: 'image/png' })
    images.value.push({ file, url: URL.createObjectURL(file) })
    showTemplateModal.value = false
    selectedTemplate.value = null
    userPhotoImg.value = null
    photoScale.value = 1
    photoOffsetX.value = 0
    photoOffsetY.value = 0
    frameNaturalWidth = 0
    frameNaturalHeight = 0
  }, 'image/png')
}

const showPictureMenu = ref(false)

function handlePictureSelect() {
  showPictureMenu.value = false
  triggerFileInput()
}

function handleTemplateSelect() {
  showPictureMenu.value = false
  selectedTemplate.value = null
  userPhotoImg.value = null
  showTemplateModal.value = true
}

const mySetsStickers = ref([])
const showPackStickers = ref(false)
const activePack = ref(null)
const imageCache = new Map()

async function loadStickerImage(url) {
  if (imageCache.has(url)) return imageCache.get(url)
  const token = localStorage.getItem('token')
  const res = await fetch(url, {
    headers: { Authorization: `Bearer ${token}` },
  })
  if (!res.ok) return ''
  const blob = await res.blob()
  const objectUrl = URL.createObjectURL(blob)
  imageCache.set(url, objectUrl)
  return objectUrl
}

async function fetchMySets() {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/stickers/my-sets`, {
      headers: { Authorization: `Bearer ${token}` },
    })
    const data = await res.json()
    const sets = data?.data?.sets || []

    mySetsStickers.value = await Promise.all(
      sets.map(async s => ({
        id: s.pack_id,
        name: s.pack_name,
        stickers: await Promise.all(
          (s.stickers || []).map(async st => ({
            id: st.id,
            url: await loadStickerImage(`${API_BASE}${st.url}`),
          }))
        ),
      }))
    )
  } catch (err) {
    console.error('Failed to load my sticker sets', err)
  }
}

async function openPack(pack) {
  activePack.value = pack
  showPackStickers.value = true
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/stickers/show?pack_id=${pack.id}`, {
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    })
    const data = await res.json()
    const list = data?.data?.stickers || []
    stickers.value = await Promise.all(
      list.map(async s => ({ ...s, url: await loadStickerImage(`${API_BASE}${s.url}`) }))
    )
  } catch (err) {
    console.error('Failed to open pack', err)
  }
}

function backToPacks() {
  showPackStickers.value = false
  activePack.value = null
}

const animatedPacks = ref([])
async function fetchPacks() {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/stickers/packs`, {
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    })
    const data = await res.json()
    const packs = data?.data?.packs || []
    animatedPacks.value = await Promise.all(
      packs.map(async p => ({
        id: p.id,
        name: p.name,
        sticker_count: p.sticker_count,
        thumbnail_url: p.thumbnail_url ? await loadStickerImage(`${API_BASE}${p.thumbnail_url}`) : '',
        added: false,
      }))
    )
  } catch (err) {
    console.error('Failed to load sticker packs', err)
  }
}
const currentPackList = computed(() => animatedPacks.value)

function addStickerPack() {
  stickerFileInputRef.value?.click()
}

async function onStickerFilePicked(e) {
  const file = e.target.files?.[0]
  if (!file) return
  e.target.value = ''

  isUploadingSticker.value = true
  try {
    const token = localStorage.getItem('token')
    const formData = new FormData()
    formData.append('pack_id', '1')        
    formData.append('file', file)

    const res = await fetch(`${API_BASE}/api/v1/front/stickers/create`, {
      method: 'POST',
      headers: token ? { Authorization: `Bearer ${token}` } : {},
      body: formData,
    })

    const data = await res.json()
    if (!res.ok) {
      console.error('Create sticker failed:', res.status, data)
      return
    }
    await fetchStickers()
  } catch (err) {
    console.error('Sticker upload error:', err)
  } finally {
    isUploadingSticker.value = false
  }
}

function addSticker(sticker) {
  selectedStickers.value.push(sticker)
  showEmoji.value = false
}

async function fetchStickers() {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(
      `${API_BASE}/api/v1/front/stickers/show?pack_id=1`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: "application/json"
        }
      }
    )
    const text = await res.text()
    const data = JSON.parse(text)
    stickers.value = data.data.stickers || []
  } catch (err) {
    console.error("Sticker Error:", err)
  }
}

async function deleteStickerSet(packId) {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/stickers/my-sets/${packId}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${token}` },
    })
    if (!res.ok) throw new Error(`Remove failed (${res.status})`)
    mySetsStickers.value = mySetsStickers.value.filter(s => s.id !== packId)
  } catch (err) {
    console.error('Remove sticker set failed', err)
  }
}

async function addPackToCollection(pack) {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/stickers/my-sets/${pack.id}`, {
      method: 'POST',
      headers: { Authorization: `Bearer ${token}` },
    })
    if (!res.ok) throw new Error(`Add failed (${res.status})`)
    pack.added = true
    await fetchMySets()
  } catch (err) {
    console.error('Add pack failed', err)
  }
}

function resolveUrl(url) {
  if (!url) return ''
  if (url.startsWith('http://') || url.startsWith('https://') || url.startsWith('data:')) {
    return url
  }
  return `${API_BASE}${url}`
}

const groups = ref([])
const selectedGroup = ref(null)
const showGroupPicker = ref(false)
const isSubmitting = ref(false)
const submitError = ref('')
const API_BASE = 'http://localhost:7070'

async function fetchGroups() {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/communities/show`, {
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    })
    const data = await res.json()
    groups.value = data?.data?.communities || []
  } catch (err) {
    console.error('Failed to load communities', err)
  }
}

//  объединение onMounted ទាំងអស់ចូលគ្នាតែមួយ
onMounted(() => {
  fetchGroups()
  fetchStickers() 
  fetchPacks()
  fetchMySets()
  fetchTemplates()
  fetchMyProfile() 
  window.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  window.removeEventListener('click', handleClickOutside)
})

// បន្ថែម topics.value.length ទីនេះ ដើម្បីឱ្យអាច Post បានពេលមាន Topic
const canPost = computed(() => {
  return (
    postText.value.trim() ||
    images.value.length ||
    videos.value.length ||
    linkUrl.value ||
    codeText.value.trim() ||
    selectedStickers.value.length ||
    topics.value.length
  )
})

function autoGrow(e) {
  const el = e.target
  el.style.height = 'auto'
  el.style.height = el.scrollHeight + 'px'
}

function addEmoji(e) {
  postText.value += e
  showEmoji.value = false
  nextTick(() => textareaRef.value?.focus())
}

function triggerFileInput() {
  fileInputRef.value?.click()
}

function onFilesPicked(e) {
  const files = Array.from(e.target.files || [])
  files.forEach(file => {
    if (file.type.startsWith('image/')) {
      images.value.push({ file, url: URL.createObjectURL(file) })
    }
    if (file.type.startsWith('video/')) {
      videos.value.push({ file, url: URL.createObjectURL(file) })
    }
  })
  e.target.value = ''
}

function removeImage(i) {
  URL.revokeObjectURL(images.value[i].url)
  images.value.splice(i, 1)
}

function confirmLink() {
  const url = linkDraft.value.trim()
  if (!url) return
  linkUrl.value = url
  linkDraft.value = ''
  showLinkInput.value = false
}

function confirmTopic() {
  const t = topicDraft.value.trim().replace(/^#/, '')
  if (!t) return
  
  const tagObj = {
    text: t,
    icon: topicStickerDraft.value ? '' : topicEmojiDraft.value,
    stickerUrl: topicStickerDraft.value ? topicStickerDraft.value.url : '',
    stickerId: topicStickerDraft.value ? topicStickerDraft.value.id : null,
  }
  
  if (!topics.value.some(x => x.text === tagObj.text)) {
    topics.value.push(tagObj)
  }
  
  topicDraft.value = ''
  topicEmojiDraft.value = ''
  topicStickerDraft.value = null
  showTopicInput.value = false
  showTopicEmojiPicker.value = false
  showTopicStickerPicker.value = false
}
//====

function resolvePostType() {
  if (codeText.value.trim()) return 'code'
  if (linkUrl.value) return 'link'
  if (videos.value.length) return 'video'
  if (images.value.length) return 'image'
  return 'text'
}

async function submitPost() {
   
  if (!canPost.value || isSubmitting.value) return
  isSubmitting.value = true
  submitError.value = ''

  const formData = new FormData()
  if (selectedGroup.value?.id) {
    formData.append('community_id', String(selectedGroup.value.id))
  }

  formData.append('caption', postText.value)
  formData.append('post_type', resolvePostType())

  if (codeText.value.trim()) formData.append('code_content', codeText.value)
  if (linkUrl.value) formData.append('link_url', linkUrl.value)
  images.value.forEach(img => formData.append('image', img.file))
  videos.value.forEach(video => formData.append('video', video.file))

  selectedStickers.value.forEach(sticker => {
    formData.append("sticker_ids[]", String(sticker.id))
  })

topics.value.forEach(tag => {
    const fullTagName = tag.icon ? tag.icon + tag.text : tag.text
    formData.append("tag_name", fullTagName)
    if (tag.stickerId !== null && tag.stickerId !== undefined && tag.stickerId !== '') {
      formData.append("tag_sticker_ids[]", String(tag.stickerId))
    } else {
      formData.append("tag_sticker_ids[]", "") 
    }
  })

  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/posts/create`, {
      method: 'POST',
      headers: token ? { Authorization: `Bearer ${token}` } : {},
      body: formData,
    })
    let data = null
    try {
      data = await res.json()
    } catch {}

    if (!res.ok) {
      submitError.value = data?.message || `Post failed (${res.status})`
      return
    }

    triggerStreakCheckIn()
    const currentUserIdVal = getCurrentUserId()
    const postType = resolvePostType()

    const optimisticPost = {
      id: 'temp-' + Date.now(),
      userId: currentUserIdVal,
      avatarUrl: myAvatarUrl.value, 
      username: myUsername.value, 
        postType: postType, 
      datetime: new Date().toLocaleString('km-KH', {
        year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit',
      }),
      description:
        postType === 'code' ? codeText.value :
        postType === 'link' ? linkUrl.value :
        postText.value,
      translatedText: '',
      photos: images.value.map(img => img.url),
      photosExpanded: false,
      tags: [...topics.value],
      videoPath: videos.value[0]?.url || null,
      videoThumbnail: null,
      videoDuration: 0,
      videoCurrentTime: 0,
      videoProgress: 0,
      isSeeking: false,
      videoExpanded: false,
      isPlaying: false,
      showFullText: false,
      showVideoControls: false,
      isMuted: false,
      isMenuOpen: false,
      playbackRate: 1,
      views: 0,
      shareCount: 0,
      commentCount: 0,
      likeCount: 0,
      isLiked: false,
      isBookmarked: false,
      bookmarkCount: 0,
      isFollowing: false,
      postToGroup: true,
      showTranslated: false,
      showMore: false,
      likedByAvatars: [],
      reaction: null,
      showReactions: false,
      showComments: false,
      showLikers: false,
      stickerIds: selectedStickers.value.map(s => s.id),
      stickers: [...selectedStickers.value],
      showSharePicker: false,
      isRepost: false,
      repostedByUserId: null,
      repostedByUsername: null,
      repostedAt: '',
      repostedByAvatar: '',
    }

    // emit('post', data)
    emit('post', optimisticPost)

    hashtagPin.value = ''
    postText.value = ''
    images.value = []
    videos.value = []
    linkUrl.value = ''
    topics.value = []
    codeText.value = ''
    showCode.value = false
    selectedGroup.value = null
    selectedStickers.value = []
    nextTick(() => {
      if (textareaRef.value) textareaRef.value.style.height = 'auto'
    })
  } catch (err) {
    console.error(err)
    submitError.value = 'Network error'
  } finally {
    isSubmitting.value = false
  }
}

// ✅ Helper function ថ្មី — ត្រូវបន្ថែមផងដែរ (មិនទាន់មានក្នុង Post.vue)
function getCurrentUserId() {
  const token = localStorage.getItem('token')
  if (!token) return null
  try {
    const payload = token.split('.')[1]
    const decoded = JSON.parse(
      decodeURIComponent(
        atob(payload.replace(/-/g, '+').replace(/_/g, '/'))
          .split('')
          .map((c) => '%' + c.charCodeAt(0).toString(16).padStart(2, '0'))
          .join('')
      )
    )
    return decoded.user_id ?? decoded.uid ?? decoded.sub ?? decoded.id ?? null
  } catch {
    return null
  }
}

const stickerFilterTab = ref('all')

const filteredStickers = computed(() => {
  switch (stickerFilterTab.value) {
    case 'animated':
      return stickers.value.filter(s => s.is_animated === true)
    case 'all':
    default:
      return stickers.value
  }
})


function handleClickOutside(event) {
  if (showTopicInput.value) {
    const topicEl = event.target.closest('.chip-wrap')
    if (!topicEl) {
      showTopicInput.value = false
    }
  }

  if (showLinkInput.value) {
    const linkEl = event.target.closest('.chip-wrap')
    if (!linkEl) {
      showLinkInput.value = false
    }
  }

  if (showEmoji.value) {
    const emojiEl = event.target.closest('.chip-wrap')
    if (!emojiEl) {
      showEmoji.value = false
    }
  }
}

const myAvatarUrl = ref('')
const myUsername = ref('You')

async function fetchMyProfile() {
  try {
    const currentUserIdVal = getCurrentUserId()
    if (!currentUserIdVal) return
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_BASE}/api/v1/front/profile/show?id=${currentUserIdVal}`, {
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    })
    if (!res.ok) return
    const json = await res.json()
    const data = json?.data ?? json
    myUsername.value = data.user_name || 'You'
    const rawAvatar = data.profile_images
    myAvatarUrl.value = rawAvatar
      ? (rawAvatar.startsWith('http') ? rawAvatar : `${API_BASE}/uploads/${rawAvatar}`)
      : ''
  } catch (e) {
    console.error('Failed to fetch my profile', e)
  }
}

async function triggerStreakCheckIn() {
  try {
    const token = localStorage.getItem('token')
    await fetch(`${API_BASE}/api/v1/front/levels/checkin`, {
      method: 'POST',
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    })
  } catch (err) {
    console.error('Streak check-in failed', err)

  }
}
</script>

<style scoped>
.emoji-picker button {
  pointer-events: auto;
}

* {
  box-sizing: border-box;
}

.composer {
  background: #fff;
  border-radius: 16px;
  padding: 16px 18px;
  font-family: 'Inter', sans-serif;
  max-width: 640px;
  margin-top: 14px;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.04);
  /* border: 1px solid#E5E7EB; */
}

.pin-row {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  color: #000;
  background: #ECEAE4;
  border-radius: 999px;
  padding: 6px 14px 6px 10px;
  margin-bottom: 12px;
}

.pin-icon {
  width: 15px;
  height: 15px;
  fill: #1976D2;
}

.pin-input {
  border: none;
  background: transparent;
  outline: none;
  font-size: 13px;
  font-weight: 700;
  color: #000;
  width: 130px;
  font-family: 'Nunito', sans-serif;
}

.pin-input::placeholder {
  color: #000;
}

.post-text {
  width: 100%;

  border: none;
  outline: none;
  resize: none;
  background: #e5e5e561;

  font-family: 'Inter', sans-serif;
  font-size: 15px;
  color: #2B2B2B;
  /* min-height: 84px; */
  min-height: 120px;
  overflow-y: auto;
  display: block;
  padding: 7px 12px;
  border-radius: 2px;

}

.post-text::placeholder {
  color: #9A9A9E;
}

.code-wrap {
  background: #F7F7F8;
  border: 1.5px solid #E7E7E7;
  border-radius: 12px;
  margin: 8px 0;
  overflow: hidden;
}

.code-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 7px 12px;
  background: #EDEDEF;
  font-size: 12px;
  font-weight: 700;
  color: #4A4A4E;
  font-family: 'Nunito', sans-serif;
}

.code-remove {
  border: none;
  background: transparent;
  color: #8A8A8E;
  cursor: pointer;
  font-size: 13px;
}

.code-text {
  width: 100%;
  border: none;
  outline: none;
  resize: vertical;
  background: transparent;
  padding: 10px 12px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 13px;
  color: #2B2B2B;
}

.link-chip {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #EFF6FB;
  border: 1.5px solid #CFE6F5;
  border-radius: 12px;
  padding: 8px 12px;
  margin: 8px 0;
  font-size: 13px;
}

.link-chip svg {
  width: 16px;
  height: 16px;
  stroke: #1E6E9C;
  fill: none;
  stroke-width: 2;
  stroke-linecap: round;
  flex-shrink: 0;
}

.link-text {
  flex: 1;
  color: #1E6E9C;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.link-chip button, .image-remove, .topic-chip button {
  border: none;
  background: transparent;
  cursor: pointer;
  color: #ffffff;
  font-size: 12px;
  flex-shrink: 0;
}

.image-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(90px, 1fr));
  gap: 8px;
  margin: 8px 0;
}

.image-tile {
  position: relative;
  aspect-ratio: 1;
  border-radius: 10px;
  overflow: hidden;
  border: 1.5px solid #E7E7E7;
}

.image-tile img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.image-remove {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: rgba(0, 0, 0, .55);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
}

.topic-row {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin: 8px 0;
}

.topic-chip {
  display: inline-flex;
  align-items: center;
  gap: 0px;
  background: #2B58F7;
   color: #ffffff;
  font-size: 13px;
  font-weight: 700;
  padding: 2px 8px 2px 2px;
  border-radius: 4px;
  font-family: 'Nunito', sans-serif;
  margin-top: 12px;
}

.group-choose-wrap {
  position: relative;
  margin: 10px 0 14px;
}

.group-choose {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  background: #ECEAE4;
  color: #000;
  border: none;
  font-weight: 700;
  font-size: 12px;
  padding: 8px 14px;
  border-radius: 999px;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
}

.group-choose svg {
  width: 15px;
  height: 15px;
  stroke: currentColor;
  fill: none;
  stroke-width: 2;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.group-choose .chevron {
  width: 12px;
  height: 12px;
  margin-left: 2px;
}

.group-picker {
  position: absolute;
  top: 40px;
  left: 0;
  background: #fff;
  border: 1px solid #E7E7E7;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .1);
  padding: 6px;
  min-width: 200px;
  z-index: 10;
}

.group-option {
  display: block;
  width: 100%;
  text-align: left;
  border: none;
  background: transparent;
  padding: 9px 12px;
  border-radius: 8px;
  font-size: 13.5px;
  color: #2B2B2B;
  cursor: pointer;
}

.group-option:hover {
  background: #F2F2F3;
}

.group-option.selected {
  background: #bbdefb;
  color: #4A4A4E;
  font-weight: 700;
}

.action-bar {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
  border-top: 1px solid #EFE2D3;
  padding-top: 14px;
}

.action-left {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
  flex: 1;
}

.chip-wrap {
  position: relative;
  display: inline-flex;
  align-items: center;
}

.chip {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  border: 2px solid #1976D2;
  background: #1976D2;
  font-weight: 700;
  font-size: 12px;
  padding: 7px 13px;
  border-radius: 999px;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
  color: #ffffff;
  transition: all 0.2s ease;
}

.chip svg {
  width: 16px;
  height: 16px;
  stroke: currentColor;
  fill: none;
  stroke-width: 1.8;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.hash-glyph {
  font-weight: 800;
}

.chip:hover {
  background: #1976D2;
  color: #ffff;
   transform: scale(1.05);
}

.chip.active {
  background: #bbdefb;
  color: #fff;
  border-color: #bbdefb;
}

.emoji-picker {
  position: absolute;
  bottom: 44px;
  left: 0;
  background: #fff;
  border: 1px solid #E7E7E7;
  border-radius: 14px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .12);
  padding: 8px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 2px;
  z-index: 10;
  width: 180px;
}

.emoji-picker button {
  border: none;
  background: transparent;
  font-size: 20px;
  padding: 6px;
  border-radius: 8px;
  cursor: pointer;
}

.emoji-picker button:hover {
  background: #F2F2F3;
}

.mini-input {
  position: absolute;
  bottom: 44px;
  left: 0;

  display: flex;
  align-items: center;
  gap: 6px;

  background: #fff;
  border: 1px solid #E7E7E7;
  border-radius: 999px;

  padding: 6px 6px 6px 14px;
  min-height: 42px;

  box-shadow: 0 8px 24px rgba(0, 0, 0, .12);
  z-index: 50;
}

.mini-input::after {
  content: "";
  position: absolute;
  top: 100%;
  left: 20%;
  transform: translateX(-50%);
  border-width: 6px;
  border-style: solid;
  border-color: #1976D2 transparent transparent transparent;
}

.mini-input input {
    pointer-events: auto;
  border: none;
  outline: none;
  font-size: 13px;
  width: 170px;
  font-family: 'Inter', sans-serif;
}

.mini-input button {
  width: 26px;
  height: 26px;
  border-radius: 50%;
  border: none;
  background: #1976D2;
  color: #fff;
  cursor: pointer;
  flex-shrink: 0;
  pointer-events: auto;
}

.post-btn {
  background: #1976D2;
  /* background: #bbdefb; */
  border: 2px solid #1976D2;
  font-weight: 700;
  font-size: 14px;
  padding: 9px 22px;
  border-radius: 999px;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
  margin-left: auto;
  color: #ffffff;
}

.post-btn svg {
  width: 18px;
  height: 18px;
}

.post-btn:hover:not(:disabled) {
  background:  #1976D2;
}

.post-btn:disabled {
  background: #E7E7E7;
  color: #B5B5B8;
  border-color: #E7E7E7;
  cursor: not-allowed;
}



.picker-panel {
  position: absolute;   
      top: 44px;  
  left: 0;
  background: #fff;
  border: 1.5px solid #E7E7E7;
  border-radius: 16px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.15);
   width: 580px;  
    max-width: 90vw; 
  max-height: 80vh;
  display: flex;
  flex-direction: column;
  z-index: 1001;            
  overflow: hidden;
  pointer-events: auto;
}

.picker-panel::before {
  display: none;
}

.picker-tabs {
  
  display: flex;
  background: #F2F2F3;
  border-top-left-radius: 16px;
  /* padding: 6px; */
  gap: 6px;
  border-bottom: 1px solid #E7E7E7;
  border-top-right-radius: 16px;
  pointer-events: auto;
}

.picker-tabs::-webkit-scrollbar {
  height: 0;
}

.picker-tab {
  flex: 1;
  border: none;
  background: transparent;
  padding: 3px;
  font-family: 'Nunito', sans-serif;
  font-weight: 700;
  font-size: 13.5px;
  color: #6A6A6E;
  cursor: pointer;
  border-radius: 10px;
  transition: all 0.2s;
  pointer-events: auto;
}

.picker-tab svg{
  width: 30px;
  height: 30px;
}

.picker-tab.active {
  background: transparent;
  color: #1976D2;
  border-bottom: 2px solid #1976D2;
  border-bottom-left-radius: 0;
  border-bottom-right-radius: 0;
  border-top-right-radius: 0;
}


.picker-body {
  flex: 1;
  padding: 12px;
  overflow-y: auto;
  max-height: 600px;
  
}

.picker-footer {
  min-height: 50px;
  border-top: 1px solid #E7E7E7;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
  border-bottom-left-radius: 16px;
  border-bottom-right-radius: 16px;
  padding: 6px;
}

.add-sticker-btn {
  width: 40%;
  height: 34px;
  border: none;
  border-radius: 999px;
  background: #1976D2;
  color: white;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  margin-right: 6px;
  pointer-events: auto;
  
}

.add-sticker-btn img{
  width: 20px;
  height: 20px;
  border-radius: 50%;
  margin-right: 2px;
}



.add-sticker-btn:hover {
  opacity: .9;
}

.sticker-grid {
  display: grid;
  grid-template-columns: repeat(10, 1fr);
  gap: 8px;
}

.sticker-item {
  border: none;
  background: #1976d247; 
  border-radius: 12px;
  padding: 6px;
  cursor: pointer;
  aspect-ratio: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.15s;
}

.sticker-item:hover {
  transform: scale(1.08);
  background: #FCE4E4;
}

.sticker-grid .sticker-item {
  width: 64px;  /* ពីមុនអាចជា 48px ឬ 50px - កែទំហំតាមតម្រូវការ */
  height: 64px;
}

.sticker-item img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.emoji-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  gap: 6px;
}

.emoji-grid button {
  border: none;
  background: transparent;
  font-size: 24px;
  padding: 6px;
  cursor: pointer;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.emoji-grid button:hover {
  background: #F2F2F3;
}

.picker-empty {
  text-align: center;
  color: #9A9A9E;
  font-size: 13px;
  padding: 20px 0;
  font-family: 'Inter', sans-serif;
}

.sticker-item img {
  width: 40px;
  height: 40px;
  display: block;
  object-fit: contain;
}

.sticker-item {
  pointer-events: auto !important;
}

.sticker-item img {
  pointer-events: none;
}

.selected-stickers {
  display:flex;
  gap:8px;
  margin:10px 0;
  flex-wrap:wrap;
}

.selected-sticker {
  position:relative;
  width:80px;
  height:80px;
  background:#f5f5f5;
  border-radius:12px;
  display:flex;
  align-items:center;
  justify-content:center;
}

.selected-sticker img {
  width:65px;
  height:65px;
  object-fit:contain;
}

.selected-sticker button {
  position:absolute;
  top:-5px;
  right:-5px;
  width:20px;
  height:20px;
  border:none;
  border-radius:50%;
  cursor:pointer;
}

.video-grid{
  display:grid;
  grid-template-columns:repeat(auto-fill,minmax(150px,1fr));
  gap:8px;
  margin:8px 0;
}


.video-tile{
  position:relative;
  border-radius:12px;
  overflow:hidden;
  border:1px solid #ddd;
}


.video-tile video{
  width:100%;
  height:150px;
  object-fit:cover;
}

.sticker-filter-row {
  display: flex;
  gap: 6px;
  overflow-x: auto;
  scrollbar-width: none;
  justify-content: center;
  width: 100%;
}

.sticker-filter-row::-webkit-scrollbar {
  display: none;
}

.sticker-filter-row::-webkit-scrollbar {
  display: none;
}

.filter-chip {
  flex-shrink: 0;
  border: none;
  /* border: 1.5px solid #1976D2; */
  background: transparent;
  color: #1976D2;
  font-family: 'Nunito', sans-serif;
  font-weight: 400;
  font-size: 12px;
  padding: 4px 4px;
  border-radius: 0;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.15s ease;
  pointer-events: auto;
  margin-bottom: -1.5px; 
  border-bottom: 2px solid transparent; 
  gap: 5px;
  display: inline-flex;
}

.filter-chip img {
  width: 16px;
  height: 16px;
  object-fit: contain;
  display: block;
  flex-shrink: 0;
}

/* .filter-chip:hover {
  background: #EFF6FB;
} */

.filter-chip.active {
  background: transparent;
  color: #1976D2;
  border-bottom: 4px solid #1976D2;
}

.filter-chip-create {
  border-color: transparent;   
  background: transparent;
  color: #8A6D1D;
}

.filter-chip-create.active {
  color: #B8901E;
  border-bottom: 2px solid #F2C744;
}

.create-sticker-panel {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 20px;
  padding: 40px 20px;
}

.add-photo-circle {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  border: none;
  background: #EFF6FB;
  color: #7FA8C9;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.15s ease;
}

.add-photo-circle img{
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.add-photo-circle:hover {
  background: #DCEBF7;
}

.add-photo-circle svg {
  width: 42px;
  height: 42px;
}

.add-photo-btn {
  border: 1.5px solid #1976D2;
  background: #1976D2;
  color: #ffff;
  font-family: 'Nunito', sans-serif;
  font-weight: 500;
  font-size: 14px;
  padding: 4px 6px;
  border-radius: 32px;
  cursor: pointer;
  transition: background 0.15s ease;
}

.add-photo-btn:hover {
  border: 1.5px solid #D95F1C;
  background-color: transparent;
  color: #D95F1C;
}

.pack-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
}

.pack-card {
  background: #fff;
  border: 1px solid #1976D2;
  border-radius: 16px;
  overflow: hidden;
  cursor: pointer;
  transition: transform 0.15s ease;
}

.pack-card:hover {
  transform: translateY(-2px);
}

.pack-thumb {
  position: relative;
  background: #FDEBAE;
  aspect-ratio: 1.3;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pack-thumb img {
  width: 70%;
  height: 70%;
  object-fit: contain;
}

.pack-count-badge {
  position: absolute;
  bottom: 8px;
  right: 8px;
  background: #1976D2;
  color: #ffff;
  border: 1px solid #1976D2;
  border-radius: 32px;
  padding: 2px 10px;
  font-weight: 800;
  font-size: 14px;
  font-family: 'Nunito', sans-serif;
}

.pack-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 10px;
  gap: 8px;
}

.pack-name {
  font-size: 12.5px;
  font-weight: 700;
  font-family: 'Inter', sans-serif;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #000;
}

.pack-add-btn {
  flex-shrink: 0;
  background: #1976D2;
  border: 1px solid #1976D2;
  color: #ffff;
  border-radius: 8px;
  font-weight: 800;
  font-size: 12px;
  padding: 4px 12px;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
}

.pack-add-btn:hover {
  background: #D95F1C;
  border: 1px solid #D95F1C;
}

.pack-detail-head {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

.pack-back-btn {
  border: none;
  background: transparent;
  font-weight: 700;
  font-size: 13px;
  color: #1976D2;
  cursor: pointer;
}

.pack-detail-title {
  font-weight: 800;
  font-size: 14px;
  font-family: 'Nunito', sans-serif;
}

.sticker-set-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.sticker-set-row {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 4px;
}

.set-drag-handle {
  flex-shrink: 0;
  width: 30px;
  height: 30px;
  border: none;
  background: #1976D2;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffff;
  cursor: grab;
}

.set-drag-handle svg {
  width: 18px;
  height: 18px;
}

.set-stickers-scroll {
  flex: 1;
  display: flex;
  gap: 6px;
  overflow-x: auto;
  scrollbar-width: none;
}

.set-stickers-scroll::-webkit-scrollbar {
  display: none;
}

.set-sticker-item {
  flex-shrink: 0;
  width: 46px;
  height: 46px;
  border: none;
  border-radius: 10px;
  /* background: #FDF0F0; */
  padding: 4px;
  cursor: pointer;
  transition: transform 0.15s;
  background-color: transparent;
  border: 1px solid #1976D2;
}

.set-sticker-item:hover {
  transform: scale(1.08);
}

.set-sticker-item img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  pointer-events: none;
  
}

.set-delete-btn {
  flex-shrink: 0;
  border: 1px solid #B5B5B8;
  background: #fff;
  color: #4A4A4E;
  font-family: 'Nunito', sans-serif;
  font-weight: 500;
  font-size: 12px;
  padding: 4px 4px;
  border-radius: 32px;
  cursor: pointer;
  white-space: nowrap;
  transition: background 0.15s ease;
}

.set-delete-btn:hover {
  background: #E0E0E0;
  color: #C6402E;
  border-color: #E7A399;
}

.chevron-mini {
  width: 10px;
  height: 10px;
  stroke: currentColor;
  fill: none;
  stroke-width: 2.2;
  margin-left: 2px;
  pointer-events: auto;
}

.dropdown-menu {
  position: absolute;
  bottom: 44px;
  left: 0;
  background: #fff;
  border: 1px solid #E7E7E7;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, .12);
  padding: 6px;
  min-width: 150px;
  z-index: 50;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
  text-align: left;
  border: none;
  background: transparent;
  padding: 8px 10px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 600;
  color: #2B2B2B;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
}

.dropdown-item svg {
  width: 16px;
  height: 16px;
  stroke: currentColor;
  fill: none;
  stroke-width: 1.8;
  flex-shrink: 0;
}

.picture-dropdown {
  position: absolute;
  top: 44px;
  left: 0;
  background: #ffff;
  border-radius: 8px;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.18);
  padding: 4px;
  min-width: 200px;
  z-index: 100;
  border: 1px solid #EFE2D3;
}

.picture-dropdown-item {
  display: flex;
  align-items: center;
  gap: 14px;
  width: 100%;
  text-align: left;
  border: none;
  background: transparent;
  padding: 6px 6px;
  border-radius: 14px;
  font-size: 12px;
  font-weight: 600;
  color: #1A1A1A;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
  transition: background 0.15s ease;
  pointer-events: auto;
}

.picture-dropdown-item:hover {
  background: rgba(0, 0, 0, 0.06);
}

.picture-dropdown-icon {
  width: 24px;
  height: 24px;
  flex-shrink: 0;
  color: #4A4A4E;
}

.dropdown-item:hover {
  background: #F2F2F3;
}

.template-modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.45);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.template-modal {
  background: #fff;
  border-radius: 12px;
  width: 90%;
  max-width: 420px;
  max-height: 80vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.25);
}

.template-modal-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 12px;
  border-bottom: 1px solid #E7E7E7;
}

.template-modal-title {
  font-weight: 800;
  font-size: 14px;
  font-family: 'Nunito', sans-serif;
  color: #1A1A1A;
}

.template-modal-close {
  border: none;
  background: transparent;
  font-size: 16px;
  color: #8A8A8E;
  cursor: pointer;
  padding: 4px;
}

.template-modal-body {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
}

.template-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.template-item {
  border: 2px solid #E7E7E7;
  border-radius: 14px;
  overflow: hidden;
  cursor: pointer;
  background: #fff;
  padding: 0;
  display: flex;
  flex-direction: column;
  transition: border-color 0.15s ease;
}

.template-item.selected {
  border-color: #1976D2;
}

.template-item img {
  width: 100%;
  aspect-ratio: 1;
  object-fit: cover;
  display: block;
}

.template-item-name {
  font-size: 12px;
  font-weight: 700;
  padding: 8px;
  font-family: 'Nunito', sans-serif;
  color: #2B2B2B;
  text-align: center;
}

.template-modal-footer {
  display: flex;
  gap: 10px;
  padding: 14px 18px;
  border-top: 1px solid #E7E7E7;
}

.template-cancel-btn,
.template-apply-btn {
  flex: 1;
  border-radius: 999px;
  padding: 10px;
  font-weight: 700;
  font-size: 14px;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
  pointer-events: auto;
}

.template-cancel-btn {
  border: 1.5px solid #E7E7E7;
  background: #fff;
  color: #4A4A4E;
}

.template-apply-btn {
  border: none;
  background: #1976D2;
  color: #fff;
}

.template-apply-btn:disabled {
  background: #E7E7E7;
  color: #B5B5B8;
  cursor: not-allowed;
}

.frame-preview-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 14px;
}

.frame-canvas {
  width: 100%;
  max-width: 320px;
  height: auto;
  aspect-ratio: 1;
  border-radius: 14px;
  background: #F2F2F3;
}

.frame-canvas.draggable {
  cursor: grab;
  touch-action: none;
}

.frame-canvas.draggable:active {
  cursor: grabbing;
}

.scale-control {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  max-width: 320px;
  padding: 8px 4px;
}

.scale-slider {
  flex: 1;
  cursor: pointer;
}

.scale-label {
  font-size: 12px;
  font-weight: 700;
  color: #4A4A4E;
  font-family: 'Nunito', sans-serif;
  min-width: 42px;
  text-align: right;
}

.scale-reset-btn {
  border: 1.5px solid #E7E7E7;
  background: #fff;
  color: #4A4A4E;
  font-size: 11.5px;
  font-weight: 600;
  padding: 5px 10px;
  border-radius: 999px;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
  flex-shrink: 0;
}

.scale-reset-btn:hover {
  background: #F2F2F3;
}

.filter-chip {
  display: flex;
  align-items: center;
  gap: 6px; 
}

.filter-icon {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

.topic-emoji-btn {
  border: none;
  background: transparent;
  width: 22px;
  height: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 15px;
  cursor: pointer;
  flex-shrink: 0;
  color: #1976D2;
}
.topic-emoji-btn svg { width: 18px; height: 18px; stroke: currentColor; fill: none; }

.topic-emoji-popover {
  position: absolute;
  bottom: 44px;
  left: 0;
  background: #fff;
  border: 1px solid #E7E7E7;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0,0,0,.12);
  padding: 6px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 2px;
  z-index: 60;
  width: 288px;
  height: 200px;
  overflow-y: auto;
}
.topic-emoji-popover button {
  border: none;
  background: transparent;
  font-size: 18px;
  padding: 4px;
  border-radius: 8px;
  cursor: pointer;
}
.topic-emoji-popover button:hover { background: #F2F2F3; }

.topic-icon-badge { 
  margin-right: 2px; 
  font-size: 13px; 
  width: 22px;
  height: 22px;
  background-color: #ffffff;
  padding: 4px;
}

.topic-sticker-popover {
  grid-template-columns: repeat(6, 1fr);
}
.topic-sticker-popover img {
  width: 28px;
  height: 28px;
  object-fit: contain;
}
.topic-sticker-badge {
  width: 16px;
  height: 16px;
  object-fit: contain;
  margin-right: 3px;
  vertical-align: middle;
}

</style>