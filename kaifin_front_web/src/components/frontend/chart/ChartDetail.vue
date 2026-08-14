<template>
  <div class="app-layout">
    <!-- Navigation Sidebar -->
    <div class="nav-sidebar">
      <div class="nav-top">
        <div class="nav-user-profile-wrapper" :title="currentUser.name">
          <div class="nav-user-profile">
            <img :src="currentUser.avatar" :alt="currentUser.name" class="nav-user-avatar" />
            <span v-if="currentUser.isOnline" class="nav-online-dot"></span>
          </div>
        </div>

        <button class="nav-icon-btn active" title="Home">
          <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path><polyline points="9 22 9 12 15 12 15 22"></polyline></svg>
        </button>
        <button class="nav-icon-btn" title="New chat">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        </button>
      </div>
    </div>

    <!-- Chats Container (Chat List Sidebar) -->
    <div class="chats-container">
      <div class="chat-sidebar-full">
        <div class="sidebar-top">
          <div class="search-box">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
            <input type="text" placeholder="Search" />
          </div>
        </div>

        <div class="sidebar-list">
          <div
            v-for="chat in mockChats"
            :key="chat.id"
            class="chat-item"
            :class="{ active: selectedChat?.id === chat.id }"
            @click="selectChat(chat)"
          >
            <div class="chat-item-avatar-wrapper">
              <div class="chat-item-avatar">
                <img :src="chat.avatar" :alt="chat.name" />
                <span v-if="chat.online" class="chat-item-online-dot">
                    
                </span>
              </div>
              <div v-if="chat.badge" class="chat-item-badge">{{ chat.badge }}</div>
            </div>
            <div class="chat-item-info">
              <div class="chat-item-title-row">
                <span class="chat-item-title">{{ chat.name }}</span>
              </div>
              <div class="chat-item-sub">{{ chat.subtitle }}</div>
            </div>
          </div>
        </div>
      </div>
       <div v-if="showChatInfo && selectedChat" class="chat-info-panel">

          <button
    class="info-close-tab"
    title="Close"
    @click="showChatInfo = false"
  ></button>

      <div class="info-profile">
        <img :src="selectedChat.avatar" :alt="selectedChat.name" class="info-avatar" />
        <div class="info-name-row">
          <span class="info-name">{{ selectedChat.name }}</span>
          <svg v-if="selectedChat.verified" width="16" height="16" viewBox="0 0 24 24" fill="#0084ff"><path d="M12 2l2.4 2.2 3.2-.6.6 3.2L20.4 9 18.8 12l1.6 3-3.2.6-.6 3.2-3.2-.6L12 20l-2.4-2.2-3.2.6-.6-3.2L3.6 15l1.6-3-1.6-3 3.2-.6.6-3.2 3.2.6z"/><path d="M9.5 12l1.8 1.8 3.2-3.6" stroke="#fff" stroke-width="1.6" fill="none" stroke-linecap="round" stroke-linejoin="round"/></svg>
        </div>
      </div>

      <!-- <div class="info-action-list">
        <button class="info-action-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M18 8a6 6 0 0 0-12 0c0 7-3 9-3 9h18s-3-2-3-9"></path><path d="M13.73 21a2 2 0 0 1-3.46 0"></path></svg>
          <span>Notifications enabled</span>
        </button>
        <button class="info-action-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 2l2 4 4 .5-3 3 .7 4.5L12 12l-3.7 2-.6-4.5-3-3L9 6z"></path></svg>
          <span>Set chat background</span>
        </button>
        <button class="info-action-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="12 2 15 9 22 9.5 17 14.5 18.5 21.5 12 18 5.5 21.5 7 14.5 2 9.5 9 9"></polygon></svg>
          <span>Add to favorites</span>
        </button>
        <button class="info-action-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 6h18"></path><path d="M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path><path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"></path></svg>
          <span>Clear history</span>
        </button>
        <button class="info-action-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M17.94 17.94A10.94 10.94 0 0 1 12 20c-7 0-11-8-11-8a18.5 18.5 0 0 1 5.06-5.94M9.9 4.24A10.94 10.94 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path><line x1="1" y1="1" x2="23" y2="23"></line></svg>
          <span>Hide chat</span>
        </button>
        <button class="info-action-item">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"></path><path d="M10 11v6M14 11v6"></path></svg>
          <span>Remove chat</span>
        </button>
        <button class="info-action-item danger">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path><polyline points="16 17 21 12 16 7"></polyline><line x1="21" y1="12" x2="9" y2="12"></line></svg>
          <span>Leave chat</span>
        </button>
      </div> -->

      <div class="info-tabs">
        <button
          v-for="tab in infoTabs"
          :key="tab.key"
          class="info-tab-btn"
          :class="{ active: activeInfoTab === tab.key }"
          @click="activeInfoTab = tab.key"
        >
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" v-html="tab.icon"></svg>
        </button>
      </div>

     <div class="info-tab-content">

  <!-- MEDIA -->
  <template v-if="activeInfoTab === 'media'">

    <div
      v-if="mediaGroups.length"
      class="media-groups"
    >

      <div
        v-for="group in mediaGroups"
        :key="group.month"
        class="media-month-group"
      >

        <!-- Month Label -->
        <div class="media-month-label">
           
          {{ group.month }}
          
        </div>

        <!-- Grid -->
        <div class="media-grid">

          <div
            v-for="(item, index) in group.items"
            :key="index"
            class="media-grid-item"
            @click="openMedia(item)"
          >

            <!-- IMAGE -->
            <img
              v-if="item.type === 'image'"
              :src="item.image"
              alt="Media"
              class="media-grid-image"
            />

            <!-- VIDEO -->
            <div
              v-else-if="item.type === 'video'"
              class="media-video-wrapper"
            >
              <video
                :src="item.video"
                class="media-grid-video"
                muted
                preload="metadata"
              ></video>

              <!-- Video Icon -->
              <div class="media-video-icon">
                <svg
                  width="28"
                  height="28"
                  viewBox="0 0 24 24"
                  fill="none"
                >
                  <path
                    d="M8 5.5L18 12L8 18.5V5.5Z"
                    fill="white"
                  />
                </svg>
              </div>

              <!-- Duration -->
              <span
                v-if="item.duration"
                class="media-video-duration"
              >
                {{ item.duration }}
              </span>
            </div>

          </div>

        </div>

      </div>

    </div>

    <!-- Empty -->
    <div
      v-else
      class="media-empty"
    >
      <svg
        width="48"
        height="48"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
      >
        <rect
          x="2"
          y="4"
          width="20"
          height="16"
          rx="2"
        />
        <circle
          cx="9"
          cy="10"
          r="1.5"
        />
        <path
          d="M22 16l-5-5-4 4-3-3-6 6"
        />
      </svg>

      <span>No photos or videos yet</span>
    </div>

  </template>


  <!-- FILES -->
<!-- FILES -->
  <template v-else-if="activeInfoTab === 'files'">
    <div v-if="fileGroups.length > 0" class="files-list">
      <div v-for="group in fileGroups" :key="group.month" class="media-month-group">
        
        <!-- Month Label -->
        <div class="media-month-label">
          {{ group.month }}
        </div>

        <!-- Files items under this month -->
        <div class="files-month-items" style="display: flex; flex-direction: column; gap: 8px;">
          <div v-for="(file, index) in group.items" :key="index" class="file-item">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
            <div class="file-details">
              <a :href="file.fileUrl" target="_blank" class="file-name">{{ file.fileName }}</a>
              <span class="file-size">{{ file.fileSize }}</span>
            </div>
          </div>
        </div>

      </div>
    </div>

    <div v-else class="info-empty-state">
      <svg
        width="48"
        height="48"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
      >
        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
        <polyline points="14 2 14 8 20 8"/>
      </svg>

      <span>Nothing here yet</span>
    </div>
  </template>


<!-- VOICE -->
  <template v-else-if="activeInfoTab === 'voice'">
    <div v-if="voiceGroups.length > 0" class="files-list">
      <div v-for="group in voiceGroups" :key="group.month" class="media-month-group">
        
        <!-- Month Label -->
        <div class="media-month-label">
          {{ group.month }}
        </div>

        <!-- Voice items under this month -->
        <div class="files-month-items voice-items-container">
          <div v-for="(voice, index) in group.items" :key="index" class="file-item voice-item-row">
            
            <!-- Microphone Icon Button -->
            <button class="msg-action-btn voice-play-btn">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#0084ff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"/>
                <path d="M19 10v1a7 7 0 0 1-14 0v-1"/>
              </svg>
            </button>

            <div class="file-details voice-item-details">
              <span class="file-name voice-item-title">Voice message</span>
              <span class="file-size voice-item-duration">{{ voice.duration }}</span>
            </div>

          </div>
        </div>

      </div>
    </div>

    <!-- Empty State បើគ្មាន Voice -->
    <div v-else class="info-empty-state">
      <svg
        width="48"
        height="48"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
      >
        <path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"/>
        <path d="M19 10v1a7 7 0 0 1-14 0v-1"/>
      </svg>
      <span>Nothing here yet</span>
    </div>
  </template>



</div>
    </div>
     
    </div>


    

    <!-- Main Chat Area (Right Side Detail View) -->
    <div class="main-chat-area">

      <!-- ផ្ទាំង Forward (ជំនួស Modal popup, បង្ហាញនៅកន្លែងតែមួយនឹង no-chat-selected) -->
      <template v-if="currentView === 'forward'">
        <ChatForward
          :message="selectedMessagesData"
          :chats="mockChats"
          @close="closeForwardPanel"
          @forward="handleForwardSubmit"
        />
      </template>

      <template v-else-if="currentView === 'detail' && selectedChat">
        <div class="chat-detail-container">
          <!-- Header (រក្សាទម្រង់ដើមជានិច្ច) -->
          <div class="chat-header">
            <button class="action-icon-btn back-btn" title="Back" @click="currentView = 'chats'">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="15 18 9 12 15 6"></polyline></svg>
            </button>
            <div class="avatar-wrapper" style="cursor: pointer;" @click="showChatInfo = !showChatInfo">
  <div class="avatar">
    <img :src="selectedChat.avatar" :alt="selectedChat.name" class="avatar-img" />
  </div>
  <span v-if="selectedChat.online" class="online-dot"></span>
</div>
            <div class="user-meta">
              <span class="username">{{ selectedChat.name }}</span>
              <span class="user-status">{{ selectedChat.online ? 'Active now' : 'Offline' }}</span>
            </div>
            <div class="header-actions">
              <button class="action-icon-btn"><svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path></svg></button>
              <button class="action-icon-btn"><svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="23 7 16 12 23 17 23 7"></polygon><rect x="1" y="5" width="15" height="14" rx="2" ry="2"></rect></svg></button>
              <button class="action-icon-btn"><svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="12"></line><line x1="12" y1="16" x2="12.01" y2="16"></line></svg></button>
            </div>
          </div>

          <!-- Messages Body -->
          <div class="chat-messages-body">
            <template v-for="(msg, index) in selectedChat.messages" :key="index">
              <div v-if="msg.timestamp" class="timestamp-divider">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="clock-icon">
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                <span>{{ msg.timestamp }}</span>
              </div>

              <div
                class="message-row"
                :class="msg.sender"
              >
                <!-- INCOMING -->
                <template v-if="msg.sender === 'incoming'">
                  <div class="msg-avatar">
                    <img :src="selectedChat.avatar" :alt="selectedChat.name" />
                  </div>
                  <div class="message-content-wrap">
                    <div v-if="msg.type === 'text'" class="bubble text-bubble">
                      <span>{{ msg.text }}</span>
                    </div>
                    <!-- <div v-else-if="msg.type === 'image'" class="bubble image-bubble">
                      <img :src="msg.image" alt="Attachment" class="msg-uploaded-img" />
                    </div> -->
                    <!-- ប្រសិនបើសារជាប្រភេទរូបភាព (Image) -->
<div v-else-if="msg.type === 'image'" class="bubble image-bubble" :class="getImageLayoutClass(msg.images)">
  
  <!-- ១ រូប -->
<template v-if="!msg.images || msg.images.length === 1">
  <img :src="msg.image || msg.images[0]" class="msg-uploaded-img" alt="Image" @click="openLightbox(msg.image || msg.images[0], [msg.image || msg.images[0]], 0)" />
</template>

  <!-- ច្រើនរូប (២, ៣, ៤ ឬ ៦ រូបឡើងទៅ) -->
<template v-else>
  <div class="msg-image-grid" :class="'count-' + Math.min(msg.images.length, 6)">
    <template v-for="(imgUrl, imgIdx) in msg.images.slice(0, 6)" :key="imgIdx">
      <div class="grid-img-item" @click="openLightbox(imgUrl, msg.images, imgIdx)">
        <img :src="imgUrl" alt="Grid Image" />
        <div v-if="imgIdx === 5 && msg.images.length > 6" class="more-overlay">
          +{{ msg.images.length - 6 }}
        </div>
      </div>
    </template>
  </div>
</template>
</div>
                  </div>
                  
                  <div class="message-actions-always">
                    <button class="msg-action-btn" title="Forward" @click.stop="goToForwardPanel(msg, index)">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M15 14l5-5-5-5"></path><path d="M20 9H9.5A5.5 5.5 0 0 0 4 14.5v0A5.5 5.5 0 0 0 9.5 20H13"></path></svg>
                    </button>
                    <button class="msg-action-btn" title="React">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="12" cy="12" r="10"></circle><path d="M8 14s1.5 2 4 2 4-2 4-2"></path><line x1="9" y1="9" x2="9.01" y2="9"></line><line x1="15" y1="9" x2="15.01" y2="9"></line></svg>
                    </button>
                  </div>
                </template>

                <!-- OUTGOING -->
                <!-- OUTGOING -->
<!-- OUTGOING -->
<template v-else>
  <div class="message-actions-always">
    <button class="msg-action-btn" title="Forward" @click.stop="goToForwardPanel(msg, index)">
      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M15 14l5-5-5-5"></path><path d="M20 9H9.5A5.5 5.5 0 0 0 4 14.5v0A5.5 5.5 0 0 0 9.5 20H13"></path></svg>
    </button>
  </div>
  <div class="message-content-wrap">
    <div v-if="msg.type === 'text'" class="bubble text-bubble outgoing-bubble">
      <span>{{ msg.text }}</span>
    </div>

    <div v-else-if="msg.type === 'image'" class="bubble image-bubble outgoing-image-bubble" :class="{ 'multi-image-bubble': msg.images && msg.images.length > 1 }">
      <template v-if="!msg.images || msg.images.length === 1">
        <img :src="msg.image || msg.images[0]" class="msg-uploaded-img" alt="Attachment" @click="openLightbox(msg.image || msg.images[0], [msg.image || msg.images[0]], 0)" />
      </template>
      <template v-else>
        <div class="msg-image-row">
          <div v-for="(imgUrl, imgIdx) in msg.images" :key="imgIdx" class="row-img-item" @click="openLightbox(imgUrl, msg.images, imgIdx)">
            <img :src="imgUrl" alt="Row Image" />
          </div>
        </div>
      </template>
    </div>

    <div v-else-if="msg.type === 'file'" class="bubble file-bubble outgoing-file-bubble">
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
      <div class="file-info">
        <a :href="msg.fileUrl" target="_blank" class="file-name">{{ msg.fileName }}</a>
        <span class="file-size">{{ msg.fileSize }}</span>
      </div>
    </div>





<!-- 📌 Voice Template សម្រាប់ Outgoing (ប៊ូតុង Play ប្ដូរទៅ border-radius: 12px) -->
    <div v-else-if="msg.type === 'voice'" class="voice-bubble outgoing-voice-bubble" style="display: flex; align-items: center; gap: 12px; padding: 10px 14px; background: #ffffff; border: 1px solid #e5e7eb; border-radius: 12px; width: 420px; box-shadow: 0 1px 2px rgba(0,0,0,0.05);">
      
      <!-- ប៊ូតុង Play/Pause (កែ border-radius ជា 12px និងរក្សាទំហំ width/height ស្មើគ្នា) -->
      <button class="msg-action-btn voice-play-btn" style="background: #F0F0F0;  border: none; border-radius: 12px; min-width: 36px; height: 36px; display: flex; align-items: center; justify-content: center; cursor: pointer; flex-shrink: 0;">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#505050" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <polygon points="5 3 19 12 5 21 5 3" fill="#fff"></polygon>
        </svg>
      </button>

      <!-- Waveform SVG Container & Footer info (Duration + Clock with SVG Icon) -->
      <div class="voice-waveform-wrapper" style="display: flex; flex-direction: column; gap: 4px; flex-grow: 1; overflow: hidden;">
        <div class="voice-waveform-container" style="width: 100%; height: 28px;">
          <div class="waveform-display" style="width: 100%; height: 100%;">
            <svg width="100%" height="100%" viewBox="0 0 1000 100" preserveAspectRatio="none" version="1.1" xmlns="http://www.w3.org/2000/svg" style="display: block;">
              <defs>
                <!-- Pattern គ្រាប់មូលពណ៌ខៀវ (Active) ជាមួយ Dynamic ID -->
                <pattern :id="'active-wave-' + index" x="0" y="0" width="24" height="100" patternUnits="userSpaceOnUse">
                  <rect x="0" y="20" width="24" height="60" rx="12" fill="#1B75D2" />
                </pattern>

                <!-- Pattern គ្រាប់មូលពណ៌ប្រផេះ (Inactive) ជាមួយ Dynamic ID -->
                <pattern :id="'inactive-wave-' + index" x="0" y="0" width="24" height="100" patternUnits="userSpaceOnUse">
                  <rect x="0" y="20" width="24" height="60" rx="12" fill="#d1d5db" />
                </pattern>
              </defs>

              <!-- ផ្នែក Inactive (បង្ហាញពេញទំហឹងនៅខាងក្រោយ) -->
              <rect x="0" y="0" width="1000" height="100" :fill="'url(#inactive-wave-' + index + ')'" />

              <!-- ផ្នែក Active (ទទឹងរត់ប្រែប្រួលតាម activeWidth) -->
              <rect x="0" y="0" :width="msg.activeWidth || '400'" height="100" :fill="'url(#active-wave-' + index + ')'" />
            </svg>
          </div>
        </div>
        
        <!-- ផ្នែកបង្ហាញ Duration និង Badge Clock ព្រមទាំង SVG Clock Icon -->
        <div style="display: flex; justify-content: space-between; align-items: center; font-size: 11px; color: #6b7280;">
          <span class="file-size voice-item-duration">{{ msg.duration || '0:05' }}</span>
          <span class="msg-time" style="display: flex; align-items: center; gap: 4px;">
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <polyline points="12 6 12 12 16 14"></polyline>
            </svg>
            {{ msg.time || '11:11 AM' }}
          </span>
        </div>
      </div>

      <audio :src="msg.voiceUrl" preload="none" style="display: none;"></audio>
    </div>


  </div>
  
  <div class="msg-avatar self-avatar">
    <img :src="currentUser.avatar" alt="Kilin" />
  </div>
</template>
                
              </div>
            </template>
          </div>


          <!-- 📌 បង្ហាញ Preview រូបភាពមុនពេល Send -->
  <div v-if="pendingImages.length > 0" class="image-preview-container">
    <div v-for="(img, index) in pendingImages" :key="index" class="preview-item">
      <img :src="img.url" alt="Preview" />
      
<button class="preview-zoom-btn" @click="openLightbox(img.url, pendingImages.map(i => i.url), index)" title="Zoom">
  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line><line x1="11" y1="8" x2="11" y2="14"></line><line x1="8" y1="11" x2="14" y2="11"></line></svg>
</button>

      <!-- ប៊ូតុងលុបចេញពី Preview (Close 'X' button) -->
      <button class="preview-remove-btn" @click="removePendingImage(index)" title="Remove">
        <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
      </button>
    </div>
  </div>

        <!-- Footer ជាមួយលក្ខខណ្ឌបង្ហាញ Voice ពេញប្រអប់ -->
          <div class="chat-input-footer">
            
            <!-- 📌 បើ showVoice ជា true គឺបង្ហាញតែ Voice component ពេញប្រអប់ -->
            <template v-if="showVoice">
              <Voice 
                class="full-width-voice"
                @close="showVoice = false" 
                @send="handleSendVoice" 
              />
            </template>

            <!-- 📌 បើ showVoice ជា false គឺបង្ហាញប៊ូតុង និងប្រអប់សារធម្មតា -->
            <template v-else>
              <div class="chat-input-wrapper">
                <button class="footer-icon-btn" title="Voice message" @click="showVoice = true">
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"></path><path d="M19 10v1a7 7 0 0 1-14 0v-1"></path><line x1="12" y1="19" x2="12" y2="23"></line><line x1="8" y1="23" x2="16" y2="23"></line></svg>
                </button>
              </div>

              <button class="footer-icon-btn" title="Send image" @click="triggerImagePicker">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
              </button>

              <input 
                type="file" 
                ref="fileInputRef" 
                style="display: none;" 
                accept="image/video"
                multiple 
                @change="handleImageSelected" 
              />

              <div class="chat-input-wrapper" style="position: relative;">
                <ChatSticker 
                  v-if="showStickerPicker" 
                  class="sticker-popup-menu"
                  @select-sticker="handleSelectSticker"
                />
                <button class="footer-icon-btn sticker-trigger-btn" title="Sticker" @click.stop="toggleStickerPicker">
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <circle cx="12" cy="12" r="10"></circle>
                    <path d="M8 14s1.5 2 4 2 4-2 4-2"></path>
                    <line x1="9" y1="9" x2="9.01" y2="9"></line>
                    <line x1="15" y1="9" x2="15.01" y2="9"></line>
                  </svg>
                </button>
              </div>

              <button class="footer-icon-btn gif-btn-text" title="GIF">
                <span>GIF</span>
              </button>

              <div class="message-input-box">
                <span class="input-prefix-aa">Aa</span>
                <input type="text" v-model="newMessage" @keyup.enter="sendMessage" placeholder="Aa..." />
                <button 
                  class="input-send-btn" 
                  :class="{ active: newMessage.trim().length > 0 || pendingImages.length > 0 }"
                  @click="sendMessage"
                  title="Send"
                >
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
                </button>
              </div>

              <button class="footer-icon-btn thumb-btn" title="Like">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path></svg>
              </button>
            </template>

          </div>





        </div>
      </template>

      <template v-else>
        <div class="no-chat-selected">
          <span>Select a chat to start messaging</span>
        </div>
      </template>
    </div>



  <!-- 📌 Lightbox Modal សម្រាប់បង្ហាញរូបភាពធំពេញអេក្រង់ -->
<div v-if="lightboxImage" class="lightbox-overlay" @click="closeLightbox">
  <div class="lightbox-content" @click.stop>
      
    <!-- 🟢 Lightbox Header -->
    <div class="lightbox-header">
      <span class="lightbox-counter">
        {{ lightboxIndex + 1 }} / {{ lightboxImagesList.length }}
      </span>
      <button class="lightbox-close-btn-header" @click="closeLightbox" title="Close">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <line x1="18" y1="6" x2="6" y2="18"></line>
          <line x1="6" y1="6" x2="18" y2="18"></line>
        </svg>
      </button>
    </div>

    <!-- រូបភាពធំចំកណ្ដាល -->
    <div class="lightbox-img-wrapper">
      <img :src="lightboxImage" alt="Zoomed Preview" />
    </div>

    <!-- 🟢 ប៊ូតុង Next / Prev នៅពីក្រោមរូបភាព (ប្រើប្រាស់ SVG Icon ដដែល) -->
    <div v-if="lightboxImagesList.length > 1" class="lightbox-bottom-nav">
      <button class="lightbox-nav-btn prev-btn" :disabled="lightboxIndex === 0" @click="prevImage" title="Previous">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <polyline points="15 18 9 12 15 6"></polyline>
        </svg>
      </button>
      
      <button class="lightbox-nav-btn next-btn" :disabled="lightboxIndex === lightboxImagesList.length - 1" @click="nextImage" title="Next">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <polyline points="9 18 15 12 9 6"></polyline>
        </svg>
      </button>
    </div>

    <!-- Option bar ខាងក្រោមបង្អស់ -->
    <div class="lightbox-footer">
      <button class="lightbox-action-item">
        <span>+ Save to album</span>
      </button>
      <button class="lightbox-action-item" @click="downloadImage(lightboxImage)">
        <span>↓ Save to PC</span>
      </button>
      <button class="lightbox-action-item text-danger">
        <span>ⓘ Report</span>
      </button>
    </div>
  </div>
</div>




  </div>
</template>

<script setup>
import { ref, computed,onMounted, onBeforeUnmount } from 'vue'
import ChatForward from './ChatForward.vue'
import ChatSticker from './ChatSticker.vue'
import Voice from './Voice.vue'



// 📌 State និង Function សម្រាប់ Voice Message
const showVoice = ref(false)
const isPlayingVoice = ref(false)

function toggleVoicePlay() {
  isPlayingVoice.value = !isPlayingVoice.value
}

// 📌 កែសម្រួលមុខងារ handleSendVoice ឱ្យទទួលយក audioBlob និងបង្កើតសារប្រភេទ voice ត្រឹមត្រូវ
function handleSendVoice(audioBlob) {
  if (!selectedChat.value) return
  
  // បង្កើត URL សម្រាប់ Play សំឡេងដែលបាន record
  const audioUrl = URL.createObjectURL(audioBlob)
  
  // បន្ថែមសារប្រភេទ voice ចូលទៅក្នុង messages របស់ chat បច្ចុប្បន្ន
  selectedChat.value.messages.push({
    id: Date.now(),
    sender: 'outgoing',
    type: 'voice',
    voiceUrl: audioUrl,
    duration: '0:05', // អ្នកអាចជំនួសដោយរយៈពេលពិតប្រាកដបើមាន
    timestamp: 'Just now'
  })
  
  selectedChat.value.subtitle = 'Sent a voice message'
  showVoice.value = false // បិទផ្ទាំង Voice recorder វិញក្រោយពេលផ្ញើ
}

// 📌 State សម្រាប់គ្រប់គ្រងការបើក/បិទផ្ទាំង Sticker
const showStickerPicker = ref(false)

// មុខងារពេលចុចប៊ូតុង Sticker
function toggleStickerPicker() {
  showStickerPicker.value = !showStickerPicker.value
}

// មុខងារពេលជ្រើសរើសស្ទីឃ័រ
function handleSelectSticker(sticker) {
  if (!selectedChat.value) return
  
  // បន្ថែមស្ទីឃ័រទៅក្នុងសារឆាត (ជាប្រភេទ image ឬ sticker)
  selectedChat.value.messages.push({
    id: Date.now(),
    sender: 'outgoing',
    type: 'image',
    image: sticker.url
  })
  selectedChat.value.subtitle = 'Sent a sticker'
  
  // បិទផ្ទាំង Sticker វិញក្រោយពេលជ្រើសរើសរួច
  showStickerPicker.value = false
}

// មុខងារបិទផ្ទាំង Sticker ពេលចុចក្រៅ
function handleClickOutside(event) {
  const stickerContainer = document.querySelector('.chat-sticker-container')
  const stickerBtn = document.querySelector('.sticker-trigger-btn')
  
  if (showStickerPicker.value) {
    const clickedInsideSticker = stickerContainer && stickerContainer.contains(event.target)
    const clickedInsideBtn = stickerBtn && stickerBtn.contains(event.target)
    
    if (!clickedInsideSticker && !clickedInsideBtn) {
      showStickerPicker.value = false
    }
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})




const currentUser = ref({
  name: 'Kilin',
  avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTarG4oabr9hgcas_7oq0T7QgZaiylfl2pJy-Rp7UEkzg&s=10',
  isOnline: true
})

// currentView អាចជា: 'chats' | 'detail' | 'forward'
const currentView = ref('detail')

const selectedMessageIds = ref([])

const selectedMessagesData = computed(() => {
  if (!selectedChat.value || !selectedChat.value.messages) return []
  return selectedMessageIds.value.map(idx => selectedChat.value.messages[idx])
})

// ចុច Forward លើសារណាមួយ → បើកផ្ទាំង Forward ភ្លាមតែម្ដង (គ្មាន selection-mode ទៀត)
function goToForwardPanel(msg, index) {
  selectedMessageIds.value = [index]
  currentView.value = 'forward'
}

// ត្រឡប់ពី ផ្ទាំង Forward មកកាន់ chat detail វិញ (ដូច Cancel)
function closeForwardPanel() {
  currentView.value = 'detail'
  selectedMessageIds.value = []
}

function handleForwardSubmit(payload) {
  const { targetChats, messages } = payload
  if (!targetChats || !messages) {
    currentView.value = 'detail'
    return
  }

  targetChats.forEach(targetChat => {
    messages.forEach(message => {
      targetChat.messages.push({
        timestamp: 'Just now',
        sender: 'outgoing',
        type: message.type,
        text: message.text || undefined,
        image: message.image || undefined
      })
      targetChat.subtitle = message.type === 'image' ? 'Sent an image' : message.text
    })
  })

  currentView.value = 'detail'
  selectedMessageIds.value = []
}

const mockChats = ref([
  {
    id: 1,
    name: 'Phorn Vuthea',
    subtitle: 'Sent an image',
    avatar: 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?w=100&h=100&fit=crop',
    online: true,
    badge: 2,
    type: 'file',
    messages: [
      { timestamp: 'Today 10:00 AM', sender: 'incoming', type: 'text', text: 'Hey Kilin, did you check the backend repository?' },
      { timestamp: 'Today 10:02 AM', sender: 'outgoing', type: 'text', text: 'Yes, I am reviewing it right now. Looks good!' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?w=500&h=300&fit=crop' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTGLZTkHKpjUPXhYJFe5ie_DU3Ci67TTpIv15fmdO7aeQ&s=10' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQIEyds_IeFiVxPXxAOL-W31dpRT6d3QvtqnSwHhKoWRA&s=10' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSndnvgoNrs_HTdNtRE4fxtdq0s1KcUpwuS-qIo6FqkmA&s=10' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROGnVJM5jnQDYKnjbdyrgO8XHpQj0fi2BDQuj3SISLgg&s=10' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRu6htPsLnM6kccirwTxeC6NRtlwWnS8n_6XK62mlqX9pMk2TsVDZeTMAs&s=10' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTc0x1c3jqIi3ogXyXnF1e8T-_Raj7GAMqanKiG-XJnFQ&s=10' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTL4k5F8SaRaOyKH7xYK5GqjjDm92-6ZnEG6QLibgAd1w&s=10' },
      { timestamp: 'Today 10:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQXhncA-ST7sUyiOa7VrQAptRfcGNfwlsb2XFgxJ2_n3A&s=10' },
      { timestamp: 'Yesterday 11:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRdPpB55zG8Hpj4Q5iE7fwkZrO5IBkL9Nnk8jxz5TEUJw&s=10' },
      { timestamp: 'Yesterday 1:05 AM', sender: 'incoming', type: 'image', image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQPqYeunhZSNupL1DG_qjRKZgKvo0H_OhIJRUHuJEERWw&s=10' },
      { timestamp: 'Today 1:05 AM', sender: 'incoming', type: 'file', fileName: 'Project_Report.pdf', fileSize: '2.4 MB', fileUrl: '#' },
{ timestamp: 'Yesterday 1:05 AM', sender: 'incoming', type: 'file', fileName: 'Source_Code.zip', fileSize: '15.8 MB', fileUrl: '#' },
{ timestamp: 'Yesterday 1:05 AM', sender: 'incoming', type: 'file', fileName: 'Database_Schema.png', fileSize: '850 KB', fileUrl: '#' },
{ timestamp: 'Yesterday 1:05 AM', sender: 'incoming', type: 'file', fileName: 'UI_Wireframes.pdf', fileSize: '4.1 MB', fileUrl: '#' },
{ timestamp: 'Yesterday 1:05 AM', sender: 'incoming', type: 'file', fileName: 'API_Documentation.docx', fileSize: '520 KB', fileUrl: '#' },
{ timestamp: 'Lastyear 8:05 AM', sender: 'incoming', type: 'file', fileName: 'Old_Archive.rar', fileSize: '45.0 MB', fileUrl: '#' },
{ timestamp: 'Today 2:10 PM', sender: 'incoming', type: 'voice', duration: '0:15', voiceUrl: '#' },
{ timestamp: 'Yesterday 4:20 PM', sender: 'incoming', type: 'voice', duration: '0:32', voiceUrl: '#' }
      
    ]
  },
  {
    id: 2,
    name: 'Pisey Soy',
    subtitle: 'See you tomorrow!',
    avatar: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=100&h=100&fit=crop',
    online: false,
    badge: null,
    type: 'file',
    messages: [
      { timestamp: 'Yesterday 3:30 PM', sender: 'incoming', type: 'text', text: 'Are we still meeting for the project discussion?' },
      { timestamp: 'Yesterday 3:32 PM', sender: 'outgoing', type: 'text', text: 'Yes, 9 AM sharp at the library.' },
      { timestamp: 'Yesterday 3:35 PM', sender: 'incoming', type: 'text', text: 'See you tomorrow!' }
    ]
  },
  {
    id: 3,
    name: 'Sokha Chan',
    subtitle: 'Can you send me the source code?',
    avatar: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&h=100&fit=crop',
    online: true,
    badge: 1,
    messages: [
      { timestamp: 'Yesterday 1:15 PM', sender: 'incoming', type: 'text', text: 'Hello Kilin! How is the web application coming along?' },
      { timestamp: 'Yesterday 1:20 PM', sender: 'outgoing', type: 'text', text: 'Almost done, just fixing a few minor bugs.' },
      { timestamp: 'Yesterday 1:22 PM', sender: 'incoming', type: 'text', text: 'Can you send me the source code?' }
    ]
  },
  {
    id: 4,
    name: 'Dara Chea',
    subtitle: 'Thanks for the update.',
    avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=100&h=100&fit=crop',
    online: false,
    badge: null,
    messages: [
      { timestamp: 'May 12, 11:00 AM', sender: 'incoming', type: 'text', text: 'Did you submit the assignment?' },
      { timestamp: 'May 12, 11:05 AM', sender: 'outgoing', type: 'text', text: 'Yes, submitted it this morning.' },
      { timestamp: 'May 12, 11:06 AM', sender: 'incoming', type: 'text', text: 'Thanks for the update.' }
    ]
  },
  {
    id: 5,
    name: 'Leakhena Meas',
    subtitle: 'Let’s grab coffee later.',
    avatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=100&h=100&fit=crop',
    online: true,
    badge: 3,
    messages: [
      { timestamp: 'May 10, 4:20 PM', sender: 'incoming', type: 'text', text: 'Are you free this afternoon?' },
      { timestamp: 'May 10, 4:25 PM', sender: 'outgoing', type: 'text', text: 'Working on code right now, maybe around 5?' },
      { timestamp: 'May 10, 4:30 PM', sender: 'incoming', type: 'text', text: 'Let’s grab coffee later.' }
    ]
  },
  {
    id: 6,
    name: 'Rithy Nov',
    subtitle: 'Sent an image',
    avatar: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=100&h=100&fit=crop',
    online: false,
    badge: null,
    messages: [
      { timestamp: 'May 08, 9:10 AM', sender: 'incoming', type: 'text', text: 'Check out this UI wireframe layout.' },
      { timestamp: 'May 08, 9:15 AM', sender: 'outgoing', type: 'text', text: 'Looks very clean and intuitive!' },
      { timestamp: 'May 08, 9:16 AM', sender: 'incoming', type: 'image', image: 'https://images.unsplash.com/photo-1507238691740-187a5b1d37b8?w=500&h=300&fit=crop' }
    ]
  },
  {
    id: 7,
    name: 'Bopha Kim',
    subtitle: 'Happy birthday!',
    avatar: 'https://images.unsplash.com/photo-1544005313-94ddf0286df2?w=100&h=100&fit=crop',
    online: true,
    badge: null,
    messages: [
      { timestamp: 'May 05, 8:00 AM', sender: 'incoming', type: 'text', text: 'Happy birthday! Wishing you success in everything!' },
      { timestamp: 'May 05, 8:10 AM', sender: 'outgoing', type: 'text', text: 'Thank you so much Bopha!' }
    ]
  },
  {
    id: 8,
    name: 'Vanna Heng',
    subtitle: 'See you at the meeting.',
    avatar: 'https://images.unsplash.com/photo-1519085360753-af0119f7cbe7?w=100&h=100&fit=crop',
    online: false,
    badge: null,
    messages: [
      { timestamp: 'May 01, 2:00 PM', sender: 'incoming', type: 'text', text: 'The meeting is moved to room B2.' },
      { timestamp: 'May 01, 2:05 PM', sender: 'outgoing', type: 'text', text: 'Got it. Thanks for letting me know.' },
      { timestamp: 'May 01, 2:06 PM', sender: 'incoming', type: 'text', text: 'See you at the meeting.' }
    ]
  },
  {
    id: 9,
    name: 'Srey Nich',
    subtitle: 'Ok sounds good!',
    avatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=100&h=100&fit=crop',
    online: true,
    badge: null,
    messages: [
      { timestamp: 'Apr 28, 6:45 PM', sender: 'incoming', type: 'text', text: 'Should we postpone the project review?' },
      { timestamp: 'Apr 28, 6:50 PM', sender: 'outgoing', type: 'text', text: 'Yes, let us do it tomorrow instead.' },
      { timestamp: 'Apr 28, 6:52 PM', sender: 'incoming', type: 'text', text: 'Ok sounds good!' }
    ]
  },
  {
    id: 10,
    name: 'Sophea Prum',
    subtitle: 'Great work on the release.',
    avatar: 'https://images.unsplash.com/photo-1522075469751-3a6694fb2f61?w=100&h=100&fit=crop',
    online: false,
    badge: null,
    messages: [
      { timestamp: 'Apr 25, 10:30 AM', sender: 'incoming', type: 'text', text: 'Have you deployed the latest build?' },
      { timestamp: 'Apr 25, 10:35 AM', sender: 'outgoing', type: 'text', text: 'Yes, production deployment is complete.' },
      { timestamp: 'Apr 25, 10:36 AM', sender: 'incoming', type: 'text', text: 'Great work on the release.' }
    ]
  },
  {
    id: 11,
    name: 'Chanthy Oum',
    subtitle: 'Let me check my schedule.',
    avatar: 'https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?w=100&h=100&fit=crop',
    online: true,
    badge: null,
    messages: [
      { timestamp: 'Apr 22, 1:10 PM', sender: 'incoming', type: 'text', text: 'Are you available for a quick sync call?' },
      { timestamp: 'Apr 22, 1:15 PM', sender: 'outgoing', type: 'text', text: 'Give me 10 minutes.' },
      { timestamp: 'Apr 22, 1:16 PM', sender: 'incoming', type: 'text', text: 'Let me check my schedule.' }
    ]
  },
  {
    id: 12,
    name: 'Kanha Rath',
    subtitle: 'Sent an image',
    avatar: 'https://images.unsplash.com/photo-1517841905240-472988babdf9?w=100&h=100&fit=crop',
    online: false,
    badge: null,
    messages: [
      { timestamp: 'Apr 20, 5:00 PM', sender: 'incoming', type: 'text', text: 'Here is the screenshot of the bug.' },
      { timestamp: 'Apr 20, 5:02 PM', sender: 'outgoing', type: 'text', text: 'Thanks, I will patch it right away.' },
      { timestamp: 'Apr 20, 5:03 PM', sender: 'incoming', type: 'image', image: 'https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5?w=500&h=300&fit=crop' }
    ]
  }
])

const selectedChat = ref(mockChats.value[0])
const newMessage = ref('')

function selectChat(chat) {
  selectedChat.value = chat
  currentView.value = 'detail'
  selectedMessageIds.value = []
   showChatInfo.value = false 
}

// Variable សម្រាប់ផ្ទុកសារថ្មី និងរូបភាពបណ្តោះអាសន្ន


// មុខងារសម្រាប់ផ្ញើសារ (ទាំង Text និង Image)
function sendMessage() {
  // បើគ្មានទាំងអត្ថបទ និងគ្មានរូបភាពទេ មិនបាច់ធ្វើអ្វីទេ
  if (!newMessage.value.trim() && pendingImages.value.length === 0) return
  if (!selectedChat.value) return

  // 📌 យក URL ទាំងអស់ក្នុង pendingImages មកដាក់បញ្ចូលគ្នាក្នុង Array តែមួយ
  if (pendingImages.value.length > 0) {
    const imageUrls = pendingImages.value.map(img => img.url)
    
    selectedChat.value.messages.push({
      id: Date.now(),
      sender: 'outgoing',
      type: 'image',
      images: imageUrls // បញ្ជូនជា Array តែមួយ ដើម្បីវាបង្ហាញជា Grid ជាមួយគ្នា
    })
    
    selectedChat.value.subtitle = 'Sent images'
  }

  // បើមានអត្ថបទ (Text)
  if (newMessage.value.trim() !== '') {
    selectedChat.value.messages.push({
      id: Date.now() + 1,
      sender: 'outgoing',
      type: 'text',
      text: newMessage.value
    })
    selectedChat.value.subtitle = newMessage.value
  }

  // សម្អាតទិន្នន័យក្នុងប្រអប់ Preview ចោលក្រោយពេលផ្ញើ
  newMessage.value = ''
  pendingImages.value = []
}

function sendMockImage() {
  selectedChat.value.messages.push({
    timestamp: 'Just now',
    sender: 'outgoing',
    type: 'image',
    image: 'https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=500&h=300&fit=crop'
  })
  selectedChat.value.subtitle = 'Sent an image'
}



// បន្ថែមក្រោម currentView
const showChatInfo = ref(false)
const activeInfoTab = ref('media')

const infoTabs = [
  { key: 'media', icon: '<rect x="2" y="4" width="20" height="16" rx="2"/><circle cx="9" cy="10" r="1.5"/><path d="M22 16l-5-5-4 4-3-3-6 6"/>' },
//   { key: 'links', icon: '<path d="M10 13a5 5 0 0 0 7.07 0l1.93-1.93a5 5 0 0 0-7.07-7.07L10 5.86"/><path d="M14 11a5 5 0 0 0-7.07 0L4.93 12.93a5 5 0 0 0 7.07 7.07L14 18.14"/>' },
  { key: 'files', icon: '<path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/>' },
  { key: 'voice', icon: '<path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"/><path d="M19 10v1a7 7 0 0 1-14 0v-1"/>' },
//   { key: 'music', icon: '<path d="M9 18V5l12-2v13"/><circle cx="6" cy="18" r="3"/><circle cx="18" cy="16" r="3"/>' },
]



const mediaGroups = computed(() => {

  if (!selectedChat.value?.messages) {
    return []
  }

  const mediaMessages = selectedChat.value.messages.filter(
    msg =>
      msg.type === 'image' ||
      msg.type === 'video'
  )

  const groups = {}

  mediaMessages.forEach(msg => {

    // ប្រើ mediaDate បើមាន
    // បើអត់ ប្រើ timestamp
    const date = msg.mediaDate || msg.timestamp || 'Unknown date'

    let month = date

    // Example:
    // "May 08, 9:16 AM"
    // -> "May 2026"
    const match = date.match(
      /^(January|February|March|April|May|June|July|August|September|October|November|December)/
    )

    if (match) {
      month = `${match[1]} 2026`
    }

    if (!groups[month]) {
      groups[month] = []
    }

    groups[month].push(msg)
  })

  return Object.keys(groups).map(month => ({
    month,
    items: groups[month]
  }))
})

function openMedia(item) {
  console.log('Open media:', item)

  // ពេលក្រោយអាចបើក fullscreen viewer
}




const fileGroups = computed(() => {
  if (!selectedChat.value || !selectedChat.value.messages) return []
  
  const fileMessages = selectedChat.value.messages.filter(msg => msg.type === 'file')
  const groups = {}

  fileMessages.forEach(msg => {
    const date = msg.timestamp || 'Unknown date'
    let month = date

    const match = date.match(
      /^(January|February|March|April|May|June|July|August|September|October|November|December)/
    )

    if (match) {
      month = `${match[1]} 2026`
    }

    if (!groups[month]) {
      groups[month] = []
    }

    groups[month].push({
      fileName: msg.fileName || 'Unnamed File',
      fileSize: msg.fileSize || 'Unknown size',
      fileUrl: msg.fileUrl || '#'
    })
  })

  return Object.keys(groups).map(month => ({
    month,
    items: groups[month]
  }))
})



// បន្ថែម computed property សម្រាប់ Voice Groups
const voiceGroups = computed(() => {
  if (!selectedChat.value || !selectedChat.value.messages) return []
  
  const voiceMessages = selectedChat.value.messages.filter(msg => msg.type === 'voice')
  const groups = {}

  voiceMessages.forEach(msg => {
    const date = msg.timestamp || 'Unknown date'
    let month = date

    const match = date.match(
      /^(January|February|March|April|May|June|July|August|September|October|November|December)/
    )

    if (match) {
      month = `${match[1]} 2026`
    }

    if (!groups[month]) {
      groups[month] = []
    }

    groups[month].push({
      duration: msg.duration || '0:00',
      voiceUrl: msg.voiceUrl || '#',
      timestamp: msg.timestamp
    })
  })

  return Object.keys(groups).map(month => ({
    month,
    items: groups[month]
  }))
})



// Reference ទៅកាន់ input file
const fileInputRef = ref(null)

// មុខងារពេលចុចលើប៊ូតុងរូបភាព ឱ្យបើក Image Picker
function triggerImagePicker() {
  if (fileInputRef.value) {
    fileInputRef.value.click()
  }
}

// មុខងារពេលអ្នកប្រើប្រាស់ជ្រើសរើសរូបភាពច្រើនរួច
function handleImageSelected(event) {
  const files = event.target.files
  if (!files || files.length === 0) return

  Array.from(files).forEach(file => {
    const imageUrl = URL.createObjectURL(file)
    pendingImages.value.push({
      file: file,
      url: imageUrl
    })
  })

  event.target.value = '' // Reset input
}


// Variable សម្រាប់ផ្ទុកcompressed/selected files បណ្តោះអាសន្នមុនផ្ញើ
const pendingImages = ref([])



// មុខងារលុបរូបភាពណាមួយចេញពី Preview វិញ
function removePendingImage(index) {
  pendingImages.value.splice(index, 1)
}

// មុខងារពេលចុចបញ្ជូនសារ (Send) គឺទើបយករូបភាពពី pendingImages ទៅដាក់ក្នុង Chat
function sendMessagesWithImages() {
  if (!selectedChat.value) return

  // បញ្ជូនរូបភាពទាំងអស់ដែលបាន Preview រួច
  pendingImages.value.forEach(img => {
    selectedChat.value.messages.push({
      timestamp: 'Just now',
      sender: 'outgoing',
      type: 'image',
      image: img.url
    })
  })

  if (pendingImages.value.length > 0) {
    selectedChat.value.subtitle = 'Sent an image'
  }

  // សម្អាតចោលក្រោយពេលផ្ញើ
  pendingImages.value = []
}



const lightboxImagesList = ref([])
const lightboxIndex = ref(0)


// Variable សម្រាប់ផ្ទុក URL រូបភាពដែលត្រូវបង្ហាញក្នុង Lightbox (បើ null គឺអត់បង្ហាញទេ)
const lightboxImage = ref(null)

// បើក Lightbox
// function openLightbox(url) {
//   lightboxImage.value = url
// }

function openLightbox(imgUrl, list = [], index = 0) {
  lightboxImage.value = imgUrl
  lightboxImagesList.value = list.length > 0 ? list : [imgUrl]
  
  lightboxIndex.value = index
}
function closeLightbox() {
  lightboxImage.value = null
  lightboxImagesList.value = []
  lightboxIndex.value = 0
}
// 🟢 មុខងារសម្រាប់រូបភាពបន្ទាប់ (Next)
function nextImage() {
  if (lightboxIndex.value < lightboxImagesList.value.length - 1) {
    lightboxIndex.value++
    lightboxImage.value = lightboxImagesList.value[lightboxIndex.value]
  }
}

// 🟢 មុខងារសម្រាប់រូបភាពមុន (Previous)
function prevImage() {
  if (lightboxIndex.value > 0) {
    lightboxIndex.value--
    lightboxImage.value = lightboxImagesList.value[lightboxIndex.value]
  }
}

// មុខងារទាញយករូបភាពចុះក្រោម (Save to PC)
function downloadImage(url) {
  const a = document.createElement('a')
  a.href = url
  a.download = 'image.jpg'
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
}
function getImageLayoutClass(images) {
  if (!images || images.length <= 1) return ''
  return `layout-${images.length}`
}





</script>

<style scoped>
button:focus,
button:active,
.msg-action-btn:focus,
.action-icon-btn:focus,
.footer-icon-btn:focus,
.nav-icon-btn:focus {
  outline: none !important;
  box-shadow: none !important;
}

.app-layout {
  width: 100%;
  height: 100%;
  min-height: 0;
  display: flex;
  padding: 0;
  flex: 1;
  flex-direction: row;
  justify-content: flex-end;
  overflow: hidden;
  background-color: #ffffff;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
  box-sizing: border-box;
}

/* Nav Sidebar */
.nav-sidebar {
  width: 70px;
  min-width: 70px;
  height: 100%;
  min-height: 0;
  background-color: #ffffff;
  border-right: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 12px;
  gap: 8px;
  flex-shrink: 0;
  box-sizing: border-box;
}

.nav-top {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  width: 100%;
}

.nav-user-profile-wrapper {
  position: relative;
  margin-bottom: 6px;
  cursor: pointer;
}

.nav-user-profile {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid #e5e7eb;
  position: relative;
}

.nav-user-avatar {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.nav-online-dot {
  position: absolute;
  bottom: 0px;
  right: 0px;
  width: 12px;
  height: 12px;
  background-color: #22c55e;
  border: 2px solid #ffffff;
  border-radius: 50%;
}

.nav-icon-btn {
  background: transparent;
  border: none;
  cursor: pointer;
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #4b5563;
}

.nav-icon-btn.active {
  color: #1B75D2;
}

/* Chats List Container (Left Side) */
/* .chats-container {
  width: 370px;
  min-width: 370px;
  height: 100%;
  min-height: 0;
  border-right: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  flex-shrink: 0;
  box-sizing: border-box;
  overflow: hidden;
} */
.chats-container {
  width: 370px;
  min-width: 370px;
  height: 100%;
  min-height: 0;
  border-right: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  flex-shrink: 0;
  box-sizing: border-box;
  overflow: hidden;

  /* ⭐ important */
  position: relative;
}

.chat-sidebar-full {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  overflow: hidden;
}

.sidebar-top {
  display: flex;
  align-items: center;
  height: 65px;
  padding: 0 16px;
  box-sizing: border-box;
  flex-shrink: 0;
}

.search-box {
  flex-grow: 1;
  display: flex;
  align-items: center;
  background-color: #f3f4f6;
  border-radius: 20px;
  padding: 6px 12px;
  gap: 8px;
}
.search-box input {
  background: transparent;
  border: none;
  outline: none;
  font-size: 14px;
  width: 100%;
}

.sidebar-list {
  flex: 1;
  overflow-y: auto;
  padding: 6px;
  min-height: 0;
}

.chat-item {
  display: flex;
  align-items: center;
  padding: 10px 12px;
  border-radius: 10px;
  cursor: pointer;
  gap: 12px;
  transition: background-color 0.15s;
}
.chat-item:hover { background-color: #f9fafb; }
.chat-item.active { background-color: #f3f4f6; }

.chat-item-avatar-wrapper {
  position: relative;
  flex-shrink: 0;
}

.chat-item-avatar {
  width: 52px;
  height: 52px;
  border-radius: 50%;
  position: relative;
}
.chat-item-avatar img { width: 100%; height: 100%; border-radius: 50%; object-fit: cover; }

.chat-item-online-dot {
  position: absolute;
  bottom: 0px;
  right: 0px;
  width: 13px;
  height: 13px;
  background-color: #22c55e;
  border: 2px solid #ffffff;
  border-radius: 50%;
}

.chat-item-badge {
  position: absolute;
  top: -2px;
  left: -2px;
  background-color: #0077ff;
  color: #ffffff;
  font-size: 11px;
  font-weight: 700;
  height: 20px;
  min-width: 18px;
  padding: 0 4px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
  z-index: 2;
}

.chat-item-info { flex-grow: 1; min-width: 0; }
.chat-item-title { font-size: 16px; font-weight: 700; color: #111827; }
.chat-item-sub { font-size: 14px; color: #6b7280; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; margin-top: 2px; }

/* Main Chat Area (Right Side) */
.main-chat-area {
  flex: 1;
  height: 100%;
  min-height: 0;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  position: relative;
  overflow: hidden;
}

.chat-detail-container {
  width: 100%;
  height: 100%;
  min-height: 0;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  overflow: hidden;
}

.chat-header {
  display: flex;
  align-items: center;
  height: 65px;
  padding: 0 16px;
  background-color: #ffffff;
  border-bottom: 1px solid #e5e7eb;
  color: #111827;
  gap: 12px;
  box-sizing: border-box;
  flex-shrink: 0;
}

.back-btn {
  display: none;
}

@media (max-width: 768px) {
  .back-btn { display: flex; }
}

.avatar-wrapper { position: relative; flex-shrink: 0; }
.avatar { width: 40px; height: 40px; border-radius: 50%; background-color: #e5e7eb; overflow: hidden; display: flex; align-items: center; justify-content: center; }
.avatar-img { width: 100%; height: 100%; object-fit: cover; }
.online-dot { position: absolute; bottom: 0; right: 0; width: 10px; height: 10px; background-color: #10b981; border: 2px solid #ffffff; border-radius: 50%; }

.user-meta { display: flex; flex-direction: column; flex-grow: 1; min-width: 0; }
.username { font-size: 15px; font-weight: 600; color: #111827; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.user-status { font-size: 12px; color: #6b7280; }

.header-actions { display: flex; gap: 4px; flex-shrink: 0; }
.action-icon-btn { background: transparent; border: none; color: #4b5563; cursor: pointer; padding: 8px; border-radius: 50%; display: flex; align-items: center; justify-content: center; }
.action-icon-btn:hover { background-color: #f3f4f6; }

.chat-messages-body {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  background: #f5f8fc;
  position: relative;
}

.timestamp-divider {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  margin: 8px 0;
}
.timestamp-divider span {
  font-size: 12px;
  color: #9ca3af;
  font-weight: 500;
}
.clock-icon {
  color: #9ca3af;
  flex-shrink: 0;
}

.message-row {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  position: relative;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 6px;
  transition: background-color 0.15s;
}
.message-row.incoming { justify-content: flex-start; }
.message-row.outgoing { justify-content: flex-end; }
/* កំណត់ទទឹងសរុបរបស់ Bubble រូបភាពផ្ញើចេញឱ្យល្មមដាក់បាន ៣ រូបក្នុងមួយជួរ */
.bubble.image-bubble {
  max-width: 320px;
  padding: 4px;
  border-radius: 12px;
  overflow: hidden;
}
.msg-image-grid {
  display: grid;
  gap: 3px;
  width: 100%;
}
/* ២ រូប: បញ្ឈរ ២ ស្មើគ្នា (1 ជួរ ២ រូប) */
.msg-image-grid.count-2 {
  grid-template-columns: repeat(2, 1fr);
  grid-auto-rows: 140px;
}

/* 3 រូប: ខាងឆ្វេងរូបធំ ១ ខាងស្ដាំរូបតូច ២ តម្រៀបលើក្រោម */
.msg-image-grid.count-3 {
  grid-template-columns: 2fr 1fr;
  grid-template-rows: repeat(2, 70px);
  height: 143px;
}
.msg-image-grid.count-3 .grid-img-item:nth-child(1) {
  grid-row: span 2; /* រូបទី១ យកកម្ពស់ពេញ ២ ជួរខាងឆ្វេង */
}

/* 4 រូប: ការ៉េ 2x2 */
.msg-image-grid.count-4 {
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: repeat(2, 110px);
}

/* 6 រូប: 3x2 (៣ រូបជួរលើ ៣ រូបជួរក្រោម) */
.msg-image-grid.count-6 {
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(2, 90px);
}

/* ─── រចនាប័ទ្មរូបភាពខាងក្នុង ─── */
.grid-img-item {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background: #e4e6eb;
  cursor: pointer;
}

.grid-img-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.2s ease;
}

.grid-img-item:hover img {
  transform: scale(1.03);
}

/* Overlay សម្រាប់បង្ហាញចំនួនរូបលើស (+N) */
.more-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.45);
  color: #fff;
  font-size: 18px;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* រចនាសម្ព័ន្ធ Grid តូចៗសម្រាប់រូបភាព outgoing */
.message-row.outgoing .msg-image-row {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
  width: 100%;
}

/* บังคับឱ្យរូបភាពនីមួយៗមានទំហំស្មើគ្នា ៣ ក្នុងមួយជួរ */
.message-row.outgoing .row-img-item {
  width: calc(33.333% - 3px);
  aspect-ratio: 1 / 1;
  border-radius: 6px;
  overflow: hidden;
  background: #f0f2f5;
  cursor: pointer;
}

.message-row.outgoing .row-img-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.message-actions-always {
  display: flex;
  align-items: center;
  gap: 4px;
}

.msg-action-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 48px;                 /* ទំហំទទឹងប្រអប់ប៊ូតុង */
  height: 32px;                /* ទំហំកម្ពស់ប្រអប់ប៊ូតុង */
  border-radius: 32px;          /* រាងមូល (អាចប្តូរជា 8px បើចង់បានជ្រុងទ្រវែង) */
  background-color: #f0f2f5;   /* ពណ៌ផ្ទៃខាងក្រោយ (Background) */
  border: 1px solid #e4e6eb;   /* ពណ៌ស៊ុម (Border) */
  cursor: pointer;
  color: #65676b;              /* ពណ៌របស់ Icon SVG ខាងក្នុង */
  transition: all 0.2s ease;
}

.msg-action-btn:hover {
  background-color: #e4e6eb;
  color: #050505;
}


.msg-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  overflow: hidden;
  flex-shrink: 0;
}
.msg-avatar img { width: 100%; height: 100%; object-fit: cover; }
.self-avatar { background-color: #3b82f6; }

.bubble {
  border-radius: 0;
  padding: 10px 14px;
  max-width: 300px;
  word-break: break-word;
  font-size: 14px;
  line-height: 1.4;
}

.text-bubble {
  background-color: #f0f2f5;
  color: #050505;
  border-left: 4px solid #0084ff;
}

.outgoing-bubble {
  background-color: #D9601C;
  color: #ffffff;
  border-left: none;
  border-right: 4px solid #1B75D2;
}

.image-bubble {
  background: transparent;
  padding: 0;
  overflow: hidden;
  border-radius: 12px;
  max-width: 220px;
}
.msg-uploaded-img {
  width: 100%;
  height: auto;
  border-radius: 12px;
  display: block;
  object-fit: cover;
}

.chat-input-footer {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  background-color: #ffffff;
  border-top: 1px solid #e5e7eb;
  gap: 8px;
  flex-shrink: 0;
}

.footer-icon-btn {
  background-color: #1B75D2;
  border: none;
  color: #ffffff;
  cursor: pointer;
  width: 36px;
  height: 36px;
  border-radius: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  position: relative;
}
.footer-icon-btn:hover {
  background-color: #1f63a8;
}

.gif-btn-text {
  background-color: #297CD4;
  font-weight: 700;
  font-size: 12px;
  color: #ffffff;
  border: none;
  border-radius: 50px;
  width: auto;
  padding: 0 12px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.gif-btn-text:hover {
  background-color: #1f63a8;
}

.message-input-box {
  flex-grow: 1;
  background-color: #f3f4f6;
  border-radius: 24px;
  padding: 8px 16px;
  display: flex;
  align-items: center;
  gap: 10px;
  min-width: 0;
}

.input-prefix-aa {
  color: #4b5563;
  font-weight: 700;
  font-size: 15px;
  flex-shrink: 0;
}

.message-input-box input {
  background: transparent;
  border: none;
  outline: none;
  color: #111827;
  font-size: 14px;
  width: 100%;
}
.message-input-box input::placeholder {
  color: #9ca3af;
}

.no-chat-selected {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #9ca3af;
  font-size: 15px;
  background-color: #f9fafb;
}



.chat-info-panel {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  min-width: 0;

  background: #ffffff;

  display: flex;
  flex-direction: column;

  overflow-y: auto;
  box-sizing: border-box;

  z-index: 20;
}

.info-close-row {
  height: 65px;
  min-height: 65px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 12px;
  border-bottom: 1px solid #e5e7eb;
  box-sizing: border-box;
}
.info-profile {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 8px 16px 20px;
  background-color: #ffffff;

}

.info-avatar {
  width: 96px;
  height: 96px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 12px;
}

.info-name-row {
  display: flex;
  align-items: center;
  gap: 6px;
}

.info-name {
  font-size: 17px;
  font-weight: 700;
  color: #111827;
}

.info-action-list {
  display: flex;
  flex-direction: column;
  padding: 8px;
  border-bottom: 1px solid #e5e7eb;
}

.info-action-item {
  display: flex;
  align-items: center;
  gap: 14px;
  background: transparent;
  border: none;
  padding: 12px 10px;
  border-radius: 8px;
  cursor: pointer;
  color: #111827;
  font-size: 15px;
  text-align: left;
}

.info-action-item:hover {
  background-color: #f3f4f6;
}

.info-action-item.danger {
  color: #ef4444;
}

.info-tabs {
  display: flex;
  /* padding: 10px 8px 0; */
  /* gap: 4px; */
  /* width: 100%; */

}

.info-tab-btn {
  flex: 1;
  background: transparent;
  border: none;
  padding: 12px 0;
  border-top: 1px solid #e5e7eb;

    border-bottom: 1px solid #e5e7eb;
  color: #6b7280;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #F7F7F7;
}

.info-tab-btn.active {
  color: #1B75D2;
  border-bottom-color: #1B75D2;
}

.info-tab-content {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  color: #9ca3af;
  font-size: 14px;
  /* padding: 18px 14px 30px; */
  box-sizing: border-box;
}

/* ================================
   MEDIA GALLERY
================================ */

.media-groups {
  width: 100%;
  background-color: #F7F7F7;
}

.media-month-group {
  width: 100%;
  /* margin-bottom: 24px; */
  /* padding: 12px; */
  /* background-color: #0000000b; */

}

/* May 2026 */
.media-month-label {
  display: flex;
  align-items: center;
  gap: 2px;
  padding-left: 8px;
    padding-right: 8px;
    padding-top: 14px;

  /* margin: 12px 12px 10px; */

  color: #4b5563;
  font-size: 11px;
  font-weight: 600;
  line-height: 1;

  white-space: nowrap;
}

.clock-month-label {
  width: 24px;
  height: 24px;

  display: flex;
  align-items: center;
  justify-content: center;

  flex-shrink: 0;





  border-radius: 50%;

  box-sizing: border-box;
}

.clock-month-label svg {
  width: 13px;
  height: 13px;
}


/* 4 columns */
.media-grid {
  width: 100%;
    padding-top: 14px;
  display: grid;

  grid-template-columns: repeat(4, minmax(0, 1fr));

  gap: 2px;
  padding-left: 8px;
  padding-right: 8px;
}


/* Individual item */
.media-grid-item {
  position: relative;

  width: 100%;
  aspect-ratio: 1 / 1;

  overflow: hidden;

  background: #f3f4f6;

  border-radius: 4px;

  cursor: pointer;

  transition:
    transform 0.18s ease,
    box-shadow 0.18s ease;
}

.media-grid-item:hover {

  box-shadow:
    0 5px 15px rgba(0, 0, 0, 0.12);
}


/* Image */
.media-grid-image {
  width: 100%;
  height: 100%;

  display: block;

  object-fit: cover;

  transition: transform 0.25s ease;
}

.media-grid-item:hover .media-grid-image {
  opacity: 0.8;
}


/* ================================
   VIDEO
================================ */

.media-video-wrapper {
  position: relative;

  width: 100%;
  height: 100%;

  background: #111827;
}

.media-grid-video {
  width: 100%;
  height: 100%;

  display: block;

  object-fit: cover;
}


/* Dark overlay */
.media-video-wrapper::after {
  content: "";

  position: absolute;

  inset: 0;

  background: rgba(0, 0, 0, 0.08);

  pointer-events: none;
}


/* Play button */
.media-video-icon {
  position: absolute;

  top: 50%;
  left: 50%;

  width: 44px;
  height: 44px;

  transform: translate(-50%, -50%);

  display: flex;
  align-items: center;
  justify-content: center;

  border-radius: 50%;

  background: rgba(0, 0, 0, 0.55);

  z-index: 2;
}


/* Duration */
.media-video-duration {
  position: absolute;

  right: 8px;
  bottom: 7px;

  padding: 3px 6px;

  border-radius: 5px;

  background: rgba(0, 0, 0, 0.65);

  color: #ffffff;

  font-size: 11px;
  font-weight: 600;

  z-index: 3;
}


/* Empty */
.media-empty {
  width: 100%;
  height: 100%;

  min-height: 220px;

  display: flex;

  flex-direction: column;

  align-items: center;
  justify-content: center;

  gap: 12px;

  color: #9ca3af;
}


/* Other tabs empty state */
.info-empty-state {
  width: 100%;
  height: 100%;

  min-height: 220px;

  display: flex;

  flex-direction: column;

  align-items: center;
  justify-content: center;

  gap: 12px;

  color: #9ca3af;
  background-color: #F7F7F7;
}

.info-empty-icon {
  opacity: 0.6;
}

.info-close-tab {
  position: absolute;
  top: 50%;
  right: 0;

  width: 8px;
  height: 42px;

  padding: 0;
  margin: 0;

  border: none;
  border-radius: 5px 0 0 5px;

  background: #ef4444;
  cursor: pointer;

  z-index: 30;

  transform: translateY(-50%);

  transition: all 0.18s ease;
}



.info-close-tab:hover {
  width: 12px;
  background: #dc2626;
}



.bubble.file-bubble {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 14px;
  background-color: #f1f3f5;
  color: #333;
  border-radius: 12px;
  max-width: 280px;
}

.outgoing-file-bubble {
  background-color: #0084ff; /* ពណ៌សម្រាប់សារចេញរបស់អ្នក */
  color: #fff;
}

.outgoing-file-bubble .file-name {
  color: #000;
}

.outgoing-file-bubble .file-size {
  color: rgba(255, 255, 255, 0.8);
}

.file-info {
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.file-name {
  font-weight: 500;
  font-size: 14px;
  text-decoration: none;
  color: #4f4e4e;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.file-name:hover {
  text-decoration: underline;
}

.file-size {
  font-size: 12px;
  color: #6c757d;
}
.files-list {
  padding-top: 12px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  overflow-y: auto;
}

.file-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px;
  margin-top: 10px;
  /* border-radius: 8px; */
  /* background: rgba(0, 0, 0, 0.02); */
}

.file-details {
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.file-item svg {
  color: #3b82f6; 
}

/* ឬប្តូរពណ៌ពេលយក鼠标ដាក់លើ (Hover) */
.file-item:hover {
  background-color: rgba(0, 0, 0, 0.047);
}

/* .file-name {
  font-size: 13px;
  font-weight: 500;
  color: inherit;
  text-decoration: none;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
} */


.voice-items-container {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.voice-item-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 12px;
}

.voice-play-btn {
  width: 36px;
  height: 36px;
  min-width: 36px;
  border-radius: 50%;
  border: none;
  /* background: #0084ff; */
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.voice-play-btn svg{
    color: #ffffff;
}

.voice-item-details {
  display: flex;
  flex-direction: column;
  flex: 1;
}

.voice-item-title {
  font-weight: 500;
  font-size: 13px;
  color: #050505;
}

.voice-item-duration {
  font-size: 11px;
  color: #65676b;
}


.image-preview-container {
  display: flex;
  gap: 10px;
  padding: 8px 12px;
  background: #ffffff;
  border-top: 1px solid #e4e6eb;
  overflow-x: auto;
}

.preview-item {
  position: relative;
  width: 64px;
  height: 64px;
  border: 1px solid #ced0d4;
  border-radius: 8px;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f7f8f9;
}

.preview-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* ប៊ូតុង Zoom (កណ្តាលរូប) */
.preview-zoom-btn {
  position: absolute;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.4);
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
  cursor: pointer;
  opacity: 0;
  transition: opacity 0.2s;
}

.preview-item:hover .preview-zoom-btn {
  opacity: 1;
}

/* ប៊ូតុង Close (X) ខាងស្តាំលើ */
.preview-remove-btn {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: #e4e6eb;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #050505;
  cursor: pointer;
  box-shadow: 0 1px 2px rgba(0,0,0,0.2);
}

.preview-remove-btn:hover {
  background: #d8dadf;
}


/* Lightbox Background Overlay */
.lightbox-overlay {
  position: fixed;
  top: 12px;
  left: 0;
  /* width: 100vw; */
  /* height: 100vh; */
  /* background: rgba(0, 0, 0, 0.8); */
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}

/* ផ្ទាំងកុងតឺន័រធំ */
/* ផ្ទាំងកុងតឺន័រធំ */
.lightbox-content {
  position: relative;
  width: 90vw;
  max-width: 1000px;
  height: 85vh;
  max-height: 800px; /* កំណត់កម្ពស់អតិបរមាដើម្បីការពារកុំឱ្យហៀរផុតអេក្រង់ */
  background: #212121;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
  /* align-items: center; */
}

/* កន្លែងបង្ហាញរូបភាពកណ្តាល */
.lightbox-img-wrapper {
  flex: none;
  justify-content: center;
  align-items: center;
 padding: 10px 20px 0px 20px;
  min-height: 0; 
}

.lightbox-img-wrapper img {
  max-width: 100%;
 max-height: 100%;
  object-fit: contain;
}

/* Footer bar ពណ៌សនៅខាងក្រោម */
/* កន្លែងបង្ហាញរូបភាពកណ្តាល */
.lightbox-img-wrapper {
  flex: 1;
  min-height: 0; /* សំខាន់ណាស់ ដើម្បីការពារកុំឱ្យរូបភាពុបបាំង Footer */
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  overflow: hidden;
}

/* ប៊ូតុង Close (X) ខាងលើ */
.lightbox-close-btn {
  position: absolute;
  top: 15px;
  right: 15px;
  background: rgba(0, 0, 0, 0.3);
  border: none;
  color: #ffffff;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 10;
}

.lightbox-close-btn:hover {
  background: rgba(255, 255, 255, 0.2);
}

/* កន្លែងបង្ហាញរូបភាពកណ្តាល */
.lightbox-img-wrapper {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  overflow: hidden;
}

.lightbox-img-wrapper img {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}

/* Footer bar ពណ៌សនៅខាងក្រោម (ធានាថាមិនបាត់អក្សរ ឬធ្លាក់ក្រោម) */
.lightbox-footer {
  min-height: 44px;
  background: #ffffff;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding: 0 20px;
  gap: 15px;
  flex-shrink: 0;
}

.lightbox-action-item {
  background: none;
  border: none;
  font-size: 13px;
  font-weight: 500;
  color: #242526;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 12px;
  border-radius: 6px;
}

.lightbox-action-item:hover {
  background: #f0f2f5;
}

.lightbox-action-item.text-danger {
  color: #e41e3f;
}
/* រចនាសម្ព័ន្ធប៊ូតុង Send ក្នុងប្រអប់ Input */
.input-send-btn {
  display: none; 
  background-color: transparent;
  border: none;
  cursor: pointer;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;

  border-radius: 50%;
  color: #1B75D2; /* ពណ៌ខៀវ */
  flex-shrink: 0;
  padding: 0;
  transition: background-color 0.2s ease;

}

/* ពេល Active (មានអក្សរ ឬមានរូបភាព preview) */
.input-send-btn.active {
 display: flex;
}

.input-send-btn.active:hover {
  background-color: rgba(0, 132, 255, 0.1);
}


/* កំណត់ទំហំ Bubble ពេលមានរូបភាពច្រើន */
.bubble.image-bubble.multi-image-bubble {
  max-width: 320px;
  padding: 4px;
  background: transparent;
}

/* រចនាសម្ព័ន្ធតម្រៀបគ្នាបែប Flex Row (មិនមែន Grid ដូច Messenger ទេ) */
.msg-image-row {
  display: flex;
  flex-wrap: wrap; /* ពេលហៀរជួរ វាធ្លាក់មកក្រោមស្វ័យប្រវត្តិ */
  gap: 6px;
  width: 100%;
}

.row-img-item {
  width: 90px;
  height: 90px;
  border-radius: 8px;
  overflow: hidden;
  background: #f0f2f5;
  cursor: pointer;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  transition: transform 0.15s ease;
}

.row-img-item:hover {
  transform: scale(1.02);
}

.row-img-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}
.lightbox-header {
  position: relative;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 4px 24px;
  z-index: 10;
  background-color: #ffffff;
  /* background: linear-gradient(to bottom, rgba(0, 0, 0, 0.6), transparent); */
}

.lightbox-counter {
  color: #000;
  font-size: 15px;
  font-weight: 500;
}

.lightbox-close-btn-header {
  background: rgba(255, 255, 255, 0.15);
  color: #000;
  border: none;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.2s ease;
}

.lightbox-close-btn-header:hover {
  background: rgba(255, 255, 255, 0.3);
}
.lightbox-nav-btn {
  position: static;
  top: 50%;
  transform: translateY(-50%);
  background: rgba(0, 0, 0, 0.5);
  color: #fff;
  border: none;
  width: 44px;
  height: 44px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 10;
  transition: background 0.2s ease, transform 0.2s ease;
}

.lightbox-nav-btn:hover {
  background: rgba(0, 0, 0, 0.8);
  transform: translateY(-50%) scale(1.1);
}

.prev-btn {
  left: 20px;
}

.next-btn {
  right: 20px;
}

.lightbox-bottom-nav {
  display: flex;
  gap: 20px;
  /* padding: 10px 0; */
  justify-content: center;
  align-items: center;
  width: 100%;
 padding: 0px 0 6px 0;
 margin-top: -10px;
}
.lightbox-nav-btn {
  position: static !important; /* បំបាត់ absolute ចោលដាច់ខាត */
  transform: none !important;
  background: rgba(255, 255, 255, 0.15);
  color: #fff;
  border: none;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.2s ease, transform 0.2s ease;
}

.lightbox-nav-btn:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.3);
  transform: scale(1.05) !important;
}

.lightbox-nav-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}




.chat-input-wrapper {
  /* position: relative; */
  /* background-color: red; */
  display: inline-block;
  

}
 

.sticker-popup-menu {
  position: absolute;
  bottom: 50px; 
  left: -100px;
  /* transform: translateX(-50%); */

  z-index: 1000;


} 


</style>