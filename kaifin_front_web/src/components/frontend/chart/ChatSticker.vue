<template>
  <div class="chat-sticker-container">
    
    <!-- 📌 1. Category Icons Header (ដាក់នៅខាងលើជា Header) -->
    <div class="sticker-top-tabs">
      <button 
        v-for="cat in categories" 
        :key="cat.id"
        class="category-icon-btn"
        :class="{ 
          'avatar-cat-btn': cat.id === 'cat',
          'active': currentCategory === cat.id 
        }"
        @click="selectCategory(cat.id)"
        :title="cat.title"
      >
        <template v-if="cat.id === 'cat'">
          <img :src="cat.icon" class="cat-avatar" alt="Cat" />
        </template>
        <template v-else>
          <span v-html="cat.icon"></span>
        </template>
      </button>
    </div>

    <!-- 📌 2. Main Content Wrapper -->
    <div class="main-content-wrapper">
      
      <!-- Grid រូបភាពស្ទីឃ័រ និងប៊ូតុង + -->
      <div class="sticker-grid-content">
        <div class="sticker-grid">
          
          <!-- ប៊ូតុងសញ្ញា (+) -->
          <div class="sticker-item add-sticker-grid-btn" @click="handleAddStickerClick" title="Add Sticker">
            <div class="add-icon-wrapper">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <line x1="12" y1="5" x2="12" y2="19"></line>
                <line x1="5" y1="12" x2="19" y2="12"></line>
              </svg>
            </div>
          </div>

          <!-- បញ្ជីរាយនាមស្ទីឃ័រតាម Category នីមួយៗ -->
          <div 
            v-for="(sticker, index) in currentStickers" 
            :key="index"
            class="sticker-item"
            @click="selectSticker(sticker)"
          >
            <img :src="sticker.url" :alt="'Sticker ' + index" />
          </div>

        </div>
      </div>

      <!-- Footer Tabs -->
      <div class="sticker-tabs-footer">
        <button 
          v-for="(tab, index) in tabs" 
          :key="index"
          class="sticker-tab-btn"
          :class="{ active: currentTab === tab.id }"
          @click="currentTab = tab.id"
        >
          {{ tab.name }}
        </button>
      </div>

    </div>

  </div>
</template>

<script>
import St1 from "@/assets/sticker_chat/st1.png"
import St2 from "@/assets/sticker_chat/st2.png"
import St3 from "@/assets/sticker_chat/st3.png"
import St4 from "@/assets/sticker_chat/st4.png"
import St5 from "@/assets/sticker_chat/st5.png"
import St6 from "@/assets/sticker_chat/st6.png"
import St7 from "@/assets/sticker_chat/st7.png"
import St8 from "@/assets/sticker_chat/st8.png"
import St9 from "@/assets/sticker_chat/st9.png"
import St10 from "@/assets/sticker_chat/st10.png"
import St11 from "@/assets/sticker_chat/st11.png"
import St12 from "@/assets/sticker_chat/st12.png"
import St13 from "@/assets/sticker_chat/st3.png"
import St14 from "@/assets/sticker_chat/st14.png"
import St15 from "@/assets/sticker_chat/st15.png"
import St16 from "@/assets/sticker_chat/st16.png"
import St17 from "@/assets/sticker_chat/st17.png"
import St18 from "@/assets/sticker_chat/st18.png"
import St19 from "@/assets/sticker_chat/st19.png"
import St20 from "@/assets/sticker_chat/st20.png"
import St21 from "@/assets/sticker_chat/st21.png"
import St22 from "@/assets/sticker_chat/st22.png"

export default {
  name: "ChatSticker",
  data() {
    return {
      currentTab: "stickers",
      tabs: [
        { id: "stickers", name: "Stickers" },
        { id: "cards", name: "Cards" },
        { id: "emoticons", name: "Emoticons" },
        { id: "constructor", name: "Constructor" }
      ],
      currentCategory: "liked",
      categories: [
        {
          id: "history",
          title: "History",
          icon: '<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>'
        },
        {
          id: "liked",
          title: "Favorites",
          icon: '<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path></svg>'
        },
        {
          id: "smile",
          title: "Smiles",
          icon: '<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><path d="M8 14s1.5 2 4 2 4-2 4-2"></path><line x1="9" y1="9" x2="9.01" y2="9"></line><line x1="15" y1="9" x2="15.01" y2="9"></line></svg>'
        },
        {
          id: "cat",
          title: "Collection",
          icon: "https://images.unsplash.com/photo-1543852786-1cf6624b9987?w=100&h=100&fit=crop"
        },
        {
          id: "cool",
          title: "Trending",
          icon: '<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon></svg>'
        }
      ],
      stickerDatabase: {
        history: [
          { url: St1 }, { url: St2 }, { url: St3 }, { url: St4 },
          { url: St5 }, { url: St6 }, { url: St7 }, { url: St8 },
          { url: St9 }, { url: St10 }, { url: St12 }, { url: St11 },
          { url: St12 }, { url: St13 }, { url: St14 }, { url: St15 }, { url: St16 }
        ],
        liked: [
          { url: "https://images.unsplash.com/photo-1518709268805-4e9042af9f23?w=150&h=150&fit=crop" },
          { url: "https://images.unsplash.com/photo-1534447677768-be436bb09401?w=150&h=150&fit=crop" },
          { url: "https://images.unsplash.com/photo-1579783902614-a3fb3927b675?w=150&h=150&fit=crop" },
          { url: "https://images.unsplash.com/photo-1541701494587-cb58502866ab?w=150&h=150&fit=crop" }
        ],
        smile: [
          { url: "https://images.unsplash.com/photo-1534447677768-be436bb09401?w=150&h=150&fit=crop" },
          { url: "https://images.unsplash.com/photo-1518709268805-4e9042af9f23?w=150&h=150&fit=crop" }
        ],
        cat: [
          { url: "https://images.unsplash.com/photo-1543852786-1cf6624b9987?w=100&h=100&fit=crop" },
          { url: "https://images.unsplash.com/photo-1579783902614-a3fb3927b675?w=100&h=100&fit=crop" },
          { url: "https://images.unsplash.com/photo-1563089145-599997674d42?w=100&h=100&fit=crop" }
        ],
        cool: [
          { url: "https://images.unsplash.com/photo-1509198397868-475647b2a1e5?w=150&h=150&fit=crop" },
          { url: "https://images.unsplash.com/photo-1541701494587-cb58502866ab?w=150&h=150&fit=crop" }
        ]
      }
    };
  },
  computed: {
    currentStickers() {
      return this.stickerDatabase[this.currentCategory] || [];
    }
  },
  methods: {
    selectCategory(catId) {
      this.currentCategory = catId;
    },
    selectSticker(sticker) {
      this.$emit("select-sticker", sticker);
    },
    handleAddStickerClick() {
      console.log("Add new sticker clicked!");
    }
  }
};
</script>

<style scoped>
.chat-sticker-container {
  width: 640px; 
  height: 370px;
  position: absolute;
  right: 0;
  background-color: #ffffff;
  border: 1px solid #e4e6eb;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
  box-sizing: border-box;
}

/* 📌 Category Icons Header (ដាក់នៅខាងលើជា Header) */
.sticker-top-tabs {
  width: 100%;
  height: 52px;
  background-color: #ffffff;
  border-bottom: 1px solid #e4e6eb;
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 0 12px;
  gap: 8px;
  flex-shrink: 0;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
}

.category-icon-btn {
  border: none;
  background: transparent;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #65676b;
  cursor: pointer;
  transition: all 0.15s;
}

.category-icon-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
  color: #050505;
}

.category-icon-btn.active {
  background-color: rgba(0, 132, 255, 0.1);
  color: #0084ff;
}

.avatar-cat-btn {
  padding: 0;
  overflow: hidden;
}

.cat-avatar {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  object-fit: cover;
  display: block;
}

/* 📌 Main Content Wrapper */
.main-content-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background-color: #f9fafb;
}

/* 📌 Grid Content */
.sticker-grid-content {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
}

.sticker-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr); /* 📌 បន្ថែមចំនួនຖັນ (Columns) ពី 5 ទៅ 6 ដើម្បីឱ្យសមាមាត្រនឹង Width ដែលធំជាងមុន */
  gap: 12px;
}

.sticker-item {
  aspect-ratio: 1 / 1;
  background: #FFFFFF;
  border-radius: 8px;
  padding: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid transparent;
  transition: all 0.15s ease;
  box-shadow: 0 1px 2px rgba(0,0,0,0.04);
}

.sticker-item:hover {
  opacity: 0.8;
}

.sticker-item img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  display: block;
}

/* 📌 ប៊ូតុងសញ្ញា (+) */
.add-sticker-grid-btn {
  color: #000;
  background-color: rgba(0, 0, 0, 0.054);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.add-sticker-grid-btn:hover {
  opacity: 0.8;
}

.add-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 📌 Footer Tabs */
.sticker-tabs-footer {
  display: flex;
  border-top: 1px solid #e4e6eb;
  padding: 0 12px;
  background-color: #ffffff;
  flex-shrink: 0;
  border-bottom-left-radius: 12px;
  border-bottom-right-radius: 12px;
}

.sticker-tab-btn {
  background: transparent;
  border: none;
  padding: 14px 16px;
  font-size: 15px;
  font-weight: 500;
  color: #65676b;
  cursor: pointer;
  position: relative;
  transition: color 0.2s;
}

.sticker-tab-btn:hover {
  color: #050505;
}

.sticker-tab-btn.active {
  color: #0084ff;
  font-weight: 600;
}

.sticker-tab-btn.active::after {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background-color: #0084ff;
  border-radius: 0 0 3px 3px;
}
</style>