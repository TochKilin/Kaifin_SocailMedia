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

    <!-- Chats Container -->
    <div class="chats-container">
      <div class="chat-sidebar-full">
        <div class="sidebar-top">
          <div class="search-box">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
            <input type="text" placeholder="Search" v-model="searchQuery" />
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
                <span v-if="chat.online" class="chat-item-online-dot"></span>
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

          <div v-if="usersWithoutChat.length > 0" class="sidebar-section-title" style="padding: 14px 16px 6px; font-size: 12px; font-weight: 700; color: #8a8d91; text-transform: uppercase; letter-spacing: 0.5px;">
            All Users
          </div>

          <div
            v-for="user in usersWithoutChat"
            :key="'user-' + user.id"
            class="chat-item"
            @click="startChatWithUser(user)"
          >
            <div class="chat-item-avatar-wrapper">
              <div class="chat-item-avatar">
                <img :src="user.avatar" :alt="user.name" />
                <span v-if="user.online" class="chat-item-online-dot"></span>
              </div>
            </div>
            <div class="chat-item-info">
              <div class="chat-item-title-row">
                <span class="chat-item-title">{{ user.name }}</span>
              </div>
              <div class="chat-item-sub">Click to start chat</div>
            </div>
          </div>

          <div v-if="usersLoading" style="padding: 14px 16px; text-align: center; color: #8a8d91; font-size: 13px;">
            Loading users...
          </div>
        </div>
      </div>

      <div v-if="showChatInfo && selectedChat" class="chat-info-panel">
        <button class="info-close-tab" title="Close" @click="showChatInfo = false"></button>

        <div class="info-profile">
          <img :src="selectedChat.avatar" :alt="selectedChat.name" class="info-avatar" />
          <div class="info-name-row">
            <span class="info-name">{{ selectedChat.name }}</span>
            <svg v-if="selectedChat.verified" width="16" height="16" viewBox="0 0 24 24" fill="#0084ff"><path d="M12 2l2.4 2.2 3.2-.6.6 3.2L20.4 9 18.8 12l1.6 3-3.2.6-.6 3.2-3.2-.6L12 20l-2.4-2.2-3.2.6-.6-3.2L3.6 15l1.6-3-1.6-3 3.2-.6.6-3.2 3.2.6z"/><path d="M9.5 12l1.8 1.8 3.2-3.6" stroke="#fff" stroke-width="1.6" fill="none" stroke-linecap="round" stroke-linejoin="round"/></svg>
          </div>
        </div>

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
            <div v-if="mediaGroups.length" class="media-groups">
              <div v-for="group in mediaGroups" :key="group.month" class="media-month-group">
                <div class="media-month-label">{{ group.month }}</div>
                <div class="media-grid">
                  <div
                    v-for="(item, index) in group.items"
                    :key="index"
                    class="media-grid-item"
                    @click="openMedia(item)"
                  >
                    <img v-if="item.type === 'image'" :src="item.image" alt="Media" class="media-grid-image" />
                    <div v-else-if="item.type === 'video'" class="media-video-wrapper">
                      <video :src="item.video" class="media-grid-video" muted preload="metadata"></video>
                      <div class="media-video-icon">
                        <svg width="28" height="28" viewBox="0 0 24 24" fill="none">
                          <path d="M8 5.5L18 12L8 18.5V5.5Z" fill="white" />
                        </svg>
                      </div>
                      <span v-if="item.duration" class="media-video-duration">{{ item.duration }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="media-empty">
              <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                <rect x="2" y="4" width="20" height="16" rx="2" />
                <circle cx="9" cy="10" r="1.5" />
                <path d="M22 16l-5-5-4 4-3-3-6 6" />
              </svg>
              <span>No photos or videos yet</span>
            </div>
          </template>

          <!-- FILES -->
          <template v-else-if="activeInfoTab === 'files'">
            <div v-if="fileGroups.length > 0" class="files-list">
              <div v-for="group in fileGroups" :key="group.month" class="media-month-group">
                <div class="media-month-label">{{ group.month }}</div>
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
              <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
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
                <div class="media-month-label">{{ group.month }}</div>
                <div class="files-month-items voice-items-container">
                  <div v-for="(voice, index) in group.items" :key="index" class="file-item voice-item-row">
                    <button class="msg-action-btn voice-play-btn" @click="toggleVoicePlayback(voice)">
                      <svg v-if="playingMessageId !== voice.id" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#0084ff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"/>
                        <path d="M19 10v1a7 7 0 0 1-14 0v-1"/>
                      </svg>
                      <svg v-else width="18" height="18" viewBox="0 0 24 24" fill="none">
                        <rect x="6" y="4" width="4" height="16" fill="#0084ff"></rect>
                        <rect x="14" y="4" width="4" height="16" fill="#0084ff"></rect>
                      </svg>
                    </button>
                    <div class="file-details voice-item-details">
                      <span class="file-name voice-item-title">Voice message</span>
                      <span class="file-size voice-item-duration">{{ getVoiceDurationLabel(voice) }}</span>
                    </div>
                    <audio
                      :ref="el => setAudioRef(el, voice.id)"
                      :src="voice.voiceUrl"
                      preload="metadata"
                      @timeupdate="onAudioTimeUpdate(voice, $event)"
                      @ended="onAudioEnded(voice)"
                      style="display: none;"
                    ></audio>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="info-empty-state">
              <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                <path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"/>
                <path d="M19 10v1a7 7 0 0 1-14 0v-1"/>
              </svg>
              <span>Nothing here yet</span>
            </div>
          </template>
        </div>
      </div>
    </div>

    <!-- Main Chat Area  -->
    <div class="main-chat-area">
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
          <!-- Header -->
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
          <div class="chat-messages-body" ref="messagesBodyRef" @scroll="handleMessagesScroll">
            <template v-for="(msg, index) in selectedChat.messages" :key="msg.id ?? index">
              <div v-if="msg.timestamp" class="timestamp-divider">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="clock-icon">
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                <span>{{ msg.timestamp }}</span>
              </div>

              <div class="message-row" :class="msg.sender">
                <!-- INCOMING -->
                <template v-if="msg.sender === 'incoming'">
                  <div class="msg-avatar">
                    <img :src="selectedChat.avatar" :alt="selectedChat.name" />
                  </div>
                  <div class="message-content-wrap">
                    <div v-if="msg.type === 'text'" class="bubble text-bubble">
                      <span>{{ getDisplayText(msg) }}</span>
                      <button
                        v-if="msg.text && msg.text.length > 300"
                        class="see-more-btn"
                        @click.stop="toggleExpand(msg)"
                      >
                        {{ msg.expanded ? 'see' : ' See more' }}
                      </button>
                    </div>
                    <div v-else-if="msg.type === 'image'" class="bubble image-bubble" :class="getImageLayoutClass(msg.images)">
                      <template v-if="!msg.images || msg.images.length === 1">
                        <img :src="msg.image || msg.images[0]" class="msg-uploaded-img" alt="Image" @click="openLightbox(msg.image || msg.images[0], [msg.image || msg.images[0]], 0)" />
                      </template>
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

                    <!-- INCOMING VOICE -->
                    <div v-else-if="msg.type === 'voice'" class="voice-bubble" style="display: flex; align-items: center; gap: 12px; padding: 10px 14px; background: #ffffff; border: 1px solid #e5e7eb; border-radius: 12px; width: 420px; box-shadow: 0 1px 2px rgba(0,0,0,0.05);">
                      <button class="msg-action-btn voice-play-btn" style="background: #F0F0F0; border: none; border-radius: 12px; min-width: 36px; height: 36px; display: flex; align-items: center; justify-content: center; cursor: pointer; flex-shrink: 0;" @click="toggleVoicePlayback(msg)">
                        <svg v-if="playingMessageId !== msg.id" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#505050" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                          <polygon points="5 3 19 12 5 21 5 3" fill="#fff"></polygon>
                        </svg>
                        <svg v-else width="16" height="16" viewBox="0 0 24 24" fill="none">
                          <rect x="6" y="4" width="4" height="16" fill="#fff"></rect>
                          <rect x="14" y="4" width="4" height="16" fill="#fff"></rect>
                        </svg>
                      </button>

                      <div class="voice-waveform-wrapper" style="display: flex; flex-direction: column; gap: 4px; flex-grow: 1; overflow: hidden;">
                        <div class="voice-waveform-container" style="width: 100%; height: 28px; cursor: pointer;" @click="seekVoice(msg, $event)">
                          <div class="waveform-display" style="width: 100%; height: 100%;">
                            <svg width="100%" height="100%" viewBox="0 0 1000 100" preserveAspectRatio="none" version="1.1" xmlns="http://www.w3.org/2000/svg" style="display: block;">
                              <defs>
                                <pattern :id="'active-wave-in-' + index" x="0" y="0" width="24" height="100" patternUnits="userSpaceOnUse">
                                  <rect x="0" y="20" width="24" height="60" rx="12" fill="#1B75D2" />
                                </pattern>
                                <pattern :id="'inactive-wave-in-' + index" x="0" y="0" width="24" height="100" patternUnits="userSpaceOnUse">
                                  <rect x="0" y="20" width="24" height="60" rx="12" fill="#d1d5db" />
                                </pattern>
                              </defs>
                              <rect x="0" y="0" width="1000" height="100" :fill="'url(#inactive-wave-in-' + index + ')'" />
                              <rect x="0" y="0" :width="getActiveWaveWidth(msg)" height="100" :fill="'url(#active-wave-in-' + index + ')'" />
                            </svg>
                          </div>
                        </div>

                        <div style="display: flex; justify-content: space-between; align-items: center; font-size: 11px; color: #6b7280;">
                          <span class="file-size voice-item-duration">{{ getVoiceDurationLabel(msg) }}</span>
                          <span class="msg-time" style="display: flex; align-items: center; gap: 4px;">
                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                              <circle cx="12" cy="12" r="10"></circle>
                              <polyline points="12 6 12 12 16 14"></polyline>
                            </svg>
                            {{ msg.time || '11:11 AM' }}
                          </span>
                        </div>
                      </div>

                      <audio
                        :ref="el => setAudioRef(el, msg.id)"
                        :src="msg.voiceUrl"
                        preload="metadata"
                        @timeupdate="onAudioTimeUpdate(msg, $event)"
                        @ended="onAudioEnded(msg)"
                        style="display: none;"
                      ></audio>
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
                <template v-else>
                  <div class="message-actions-always">
                    <button class="msg-action-btn" title="Forward" @click.stop="goToForwardPanel(msg, index)">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M15 14l5-5-5-5"></path><path d="M20 9H9.5A5.5 5.5 0 0 0 4 14.5v0A5.5 5.5 0 0 0 9.5 20H13"></path></svg>
                    </button>
                  </div>
                  <div class="message-content-wrap">
                    <div v-if="msg.type === 'text'" class="bubble text-bubble outgoing-bubble">
                      <span>{{ getDisplayText(msg) }}</span>
                      <button
                        v-if="msg.text && msg.text.length > 300"
                        class="see-more-btn"
                        @click.stop="toggleExpand(msg)"
                      >
                        {{ msg.expanded ? 'ឃើញតិចជាងនេះ' : 'មើលបន្ថែម' }}
                      </button>
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

                    <div v-else-if="msg.type === 'voice'" class="voice-bubble outgoing-voice-bubble" style="display: flex; align-items: center; gap: 12px; padding: 10px 14px; background: #ffffff; border: 1px solid #e5e7eb; border-radius: 12px; width: 420px; box-shadow: 0 1px 2px rgba(0,0,0,0.05);">
                      <button
                        class="msg-action-btn voice-play-btn"
                        style="background: #F0F0F0;  border: none; border-radius: 12px; min-width: 36px; height: 36px; display: flex; align-items: center; justify-content: center; cursor: pointer; flex-shrink: 0;"
                        @click="toggleVoicePlayback(msg)"
                      >
                        <svg v-if="playingMessageId !== msg.id" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#505050" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                          <polygon points="5 3 19 12 5 21 5 3" fill="#fff"></polygon>
                        </svg>
                        <svg v-else width="16" height="16" viewBox="0 0 24 24" fill="none">
                          <rect x="6" y="4" width="4" height="16" fill="#fff"></rect>
                          <rect x="14" y="4" width="4" height="16" fill="#fff"></rect>
                        </svg>
                      </button>

                      <div class="voice-waveform-wrapper" style="display: flex; flex-direction: column; gap: 4px; flex-grow: 1; overflow: hidden;">
                        <div class="voice-waveform-container" style="width: 100%; height: 28px; cursor: pointer;" @click="seekVoice(msg, $event)">
                          <div class="waveform-display" style="width: 100%; height: 100%;">
                            <svg width="100%" height="100%" viewBox="0 0 1000 100" preserveAspectRatio="none" version="1.1" xmlns="http://www.w3.org/2000/svg" style="display: block;">
                              <defs>
                                <pattern :id="'active-wave-' + index" x="0" y="0" width="24" height="100" patternUnits="userSpaceOnUse">
                                  <rect x="0" y="20" width="24" height="60" rx="12" fill="#1B75D2" />
                                </pattern>
                                <pattern :id="'inactive-wave-' + index" x="0" y="0" width="24" height="100" patternUnits="userSpaceOnUse">
                                  <rect x="0" y="20" width="24" height="60" rx="12" fill="#d1d5db" />
                                </pattern>
                              </defs>
                              <rect x="0" y="0" width="1000" height="100" :fill="'url(#inactive-wave-' + index + ')'" />
                              <rect x="0" y="0" :width="getActiveWaveWidth(msg)" height="100" :fill="'url(#active-wave-' + index + ')'" />
                            </svg>
                          </div>
                        </div>

                        <div style="display: flex; justify-content: space-between; align-items: center; font-size: 11px; color: #6b7280;">
                          <span class="file-size voice-item-duration">{{ getVoiceDurationLabel(msg) }}</span>
                          <span class="msg-time" style="display: flex; align-items: center; gap: 4px;">
                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                              <circle cx="12" cy="12" r="10"></circle>
                              <polyline points="12 6 12 12 16 14"></polyline>
                            </svg>
                            {{ msg.time || '11:11 AM' }}
                          </span>
                        </div>
                      </div>

                      <audio
                        :ref="el => setAudioRef(el, msg.id)"
                        :src="msg.voiceUrl"
                        preload="metadata"
                        @timeupdate="onAudioTimeUpdate(msg, $event)"
                        @ended="onAudioEnded(msg)"
                        style="display: none;"
                      ></audio>
                    </div>
                  </div>

                  <div class="msg-avatar self-avatar">
                    <img :src="currentUser.avatar" alt="Kilin" />
                  </div>
                </template>
              </div>
            </template>
          </div>

          <button
            v-if="showScrollToBottomBtn"
            class="scroll-to-bottom-btn"
            @click="scrollToBottom(true); showScrollToBottomBtn = false"
          >
            ↓ New message
          </button>

          <div v-if="pendingImages.length > 0" class="image-preview-container">
            <div v-for="(img, index) in pendingImages" :key="index" class="preview-item">
              <img :src="img.url" alt="Preview" />
              <button class="preview-zoom-btn" @click="openLightbox(img.url, pendingImages.map(i => i.url), index)" title="Zoom">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line><line x1="11" y1="8" x2="11" y2="14"></line><line x1="8" y1="11" x2="14" y2="11"></line></svg>
              </button>
              <button class="preview-remove-btn" @click="removePendingImage(index)" title="Remove">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
              </button>
            </div>
          </div>

          <!-- Footer -->
          <div class="chat-input-footer">
            <template v-if="showVoice">
              <Voice
                class="full-width-voice"
                @close="showVoice = false"
                @send="handleSendVoice"
              />
            </template>

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

    <!-- Lightbox Modal -->
    <div v-if="lightboxImage" class="lightbox-overlay" @click="closeLightbox">
      <div class="lightbox-content" @click.stop>
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

        <div class="lightbox-img-wrapper">
          <img :src="lightboxImage" alt="Zoomed Preview" />
        </div>

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
import { ref, computed, onMounted, onBeforeUnmount, watch, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'
import ChatForward from './ChatForward.vue'
import ChatSticker from './ChatSticker.vue'
import Voice from './Voice.vue'

const messagesBodyRef = ref(null)
const showScrollToBottomBtn = ref(false)
const SCROLL_BOTTOM_THRESHOLD = 120

function isNearBottom() {
  const el = messagesBodyRef.value
  if (!el) return true
  const distance = el.scrollHeight - el.scrollTop - el.clientHeight
  return distance <= SCROLL_BOTTOM_THRESHOLD
}

function scrollToBottom(smooth = false) {
  nextTick(() => {
    const el = messagesBodyRef.value
    if (!el) return
    el.scrollTo({
      top: el.scrollHeight,
      behavior: smooth ? 'smooth' : 'auto',
    })
  })
}

function shouldAutoScroll() {
  return isNearBottom()
}

function handleMessagesScroll() {
  if (isNearBottom()) {
    showScrollToBottomBtn.value = false
  }
}

const route = useRoute()
const router = useRouter()
const API_BASE = import.meta.env.VITE_API_URL
const FILE_BASE = import.meta.env.VITE_API_URL

function authHeaders() {
  const token = localStorage.getItem('token')
  return { Authorization: `Bearer ${token}` }
}

function resolveUrl(path) {
  if (!path) return ''
  if (/^(https?:|blob:|data:)/.test(path)) return path
  const cleanPath = path.startsWith('/') ? path : `/${path}`
  return `${FILE_BASE}/uploads${cleanPath}`
}

const currentUser = ref({
  name: '',
  avatar: '',
  isOnline: true,
})

async function fetchProfile() {
  try {
    const res = await axios.get(`${API_BASE}/auth/profile`, {
      headers: authHeaders(),
    })
    if (res.data.success) {
      const p = res.data.data
      currentUser.value = {
        name: p.user_name,
        avatar: resolveUrl(p.profile_images),
        isOnline: true,
      }
    }
  } catch (err) {
    console.log('FETCH PROFILE ERROR:', err)
  }
}

const currentView = ref('chats')
const chatsLoading = ref(false)
const messagesLoading = ref(false)
const searchQuery = ref('')
const mockChats = ref([])

async function fetchConversations() {
  chatsLoading.value = true
  try {
    const res = await axios.get(`${API_BASE}/chats/show`, {
      headers: authHeaders(),
      params: { tab: 'all', search: searchQuery.value },
    })
    if (res.data.success) {
      mockChats.value = res.data.data.conversations.map((c) => ({
        id: c.conversation_id,
        name: c.name,
        subtitle: c.last_message || '',
        avatar: resolveUrl(c.avatar),
        online: c.online,
        isGroup: c.is_group,
        verified: false,
        badge: c.unread_count > 0 ? c.unread_count : null,
        otherUserId: c.other_user_id,
        messages: [],
        messagesLoaded: false,
        beforeIdCursor: 0,
        hasMore: true,
      }))
    }
  } catch (err) {
    console.log('FETCH CONVERSATIONS ERROR:', err)
  } finally {
    chatsLoading.value = false
  }
}

const TEXT_TRUNCATE_LIMIT = 300

function getDisplayText(msg) {
  if (!msg.text) return ''
  if (msg.expanded || msg.text.length <= TEXT_TRUNCATE_LIMIT) {
    return msg.text
  }
  return msg.text.slice(0, TEXT_TRUNCATE_LIMIT) + '...'
}

function toggleExpand(msg) {
  msg.expanded = !msg.expanded
}

const allUsers = ref([])
const usersLoading = ref(false)

async function fetchAllUsers() {
  usersLoading.value = true
  try {
    const res = await axios.get(`${API_BASE}/chats/users/search`, {
      headers: authHeaders(),
      params: { search: searchQuery.value, limit: 30 },
    })
    if (res.data.success) {
      const list = res.data.data.users || []
      allUsers.value = list.map((u) => ({
        id: u.id,
        name: (u.first_name || u.last_name)
          ? `${u.first_name || ''} ${u.last_name || ''}`.trim()
          : u.user_name,
        avatar: resolveUrl(u.profile_images),
        online: false,
      }))
    }
  } catch (err) {
    console.log('FETCH ALL USERS ERROR:', err)
  } finally {
    usersLoading.value = false
  }
}

const displayList = computed(() => {
  const existingUserIds = new Set(
    mockChats.value.map((c) => c.otherUserId).filter(Boolean)
  )
  const usersWithoutChat = allUsers.value
    .filter((u) => !existingUserIds.has(u.id))
    .map((u) => ({
      id: null,
      name: u.name,
      subtitle: 'ចាប់ផ្ដើមជជែក',
      avatar: u.avatar,
      online: u.online,
      isGroup: false,
      verified: false,
      badge: null,
      otherUserId: u.id,
      messages: [],
      messagesLoaded: false,
      beforeIdCursor: 0,
      hasMore: true,
      isNewChat: true,
    }))
  return [...mockChats.value, ...usersWithoutChat]
})

let searchTimeout = null
watch(searchQuery, () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    fetchConversations()
    fetchAllUsers()
  }, 400)
})

const isStartingChat = ref(false)

async function startChatWithUser(user) {
  if (isStartingChat.value) return
  isStartingChat.value = true
  try {
    const form = new FormData()
    form.append('target_user_id', user.id)

    const res = await axios.post(`${API_BASE}/chats/start`, form, {
      headers: { ...authHeaders(), 'Content-Type': 'multipart/form-data' },
    })
    if (res.data.success) {
      const conversationId = res.data.data.conversation_id

      let chat = mockChats.value.find((c) => c.id === conversationId)
      if (!chat) {
        await fetchConversations()
        chat = mockChats.value.find((c) => c.id === conversationId)
      }
      if (chat) {
        await selectChat(chat)
      }
    }
  } catch (err) {
    console.log('START CHAT WITH USER ERROR:', err)
  } finally {
    isStartingChat.value = false
  }
}

function openListItem(item) {
  if (item.isNewChat) {
    startChatWithUser({ id: item.otherUserId })
  } else {
    selectChat(item)
  }
}

function mapMessage(m) {
  const base = {
    id: m.id,
    sender: m.is_mine ? 'outgoing' : 'incoming',
    type: m.type,
    timestamp: formatTimestamp(m.created_at),
    time: formatTime(m.created_at),
    mediaDate: formatTimestamp(m.created_at),
    expanded: false,
  }

  if (m.type === 'text') {
    base.text = m.content
  } else if (m.type === 'image') {
    let imgs = (m.attachments || []).map((a) => resolveUrl(a.url))
    if (imgs.length === 0 && m.content) {
      imgs = [resolveUrl(m.content)]
    }
    if (imgs.length === 1) base.image = imgs[0]
    base.images = imgs
  } else if (m.type === 'file') {
    const att = (m.attachments || [])[0]
    base.fileName = att?.file_name
    base.fileSize = att?.file_size
    base.fileUrl = resolveUrl(att?.url)
  } else if (m.type === 'voice') {
    const att = (m.attachments || [])[0]
    base.voiceUrl = resolveUrl(att?.url)
    base.duration = att?.duration || '0:00'
  }

  return base
}

function formatTimestamp(dateStr) {
  if (!dateStr) return ''
  const d = new Date(dateStr)
  return (
    d.toLocaleDateString([], { month: 'long', day: 'numeric' }) +
    ', ' +
    d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
  )
}

function formatTime(dateStr) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleTimeString([], {
    hour: '2-digit',
    minute: '2-digit',
  })
}

async function fetchMessages(chat, beforeId = 0) {
  messagesLoading.value = true
  try {
    const res = await axios.get(`${API_BASE}/chats/${chat.id}/messages`, {
      headers: authHeaders(),
      params: { before_id: beforeId, limit: 30 },
    })
    if (res.data.success) {
      const list = res.data.data.messages || []
      const msgs = list.map(mapMessage).reverse()

      if (beforeId) {
        chat.messages = [...msgs, ...chat.messages]
      } else {
        chat.messages = msgs
      }

      chat.hasMore = list.length === 30
      if (list.length > 0) {
        chat.beforeIdCursor = list[list.length - 1].id
      }
      chat.messagesLoaded = true
    }
  } catch (err) {
    console.log('FETCH MESSAGES ERROR:', err)
  } finally {
    messagesLoading.value = false
  }
}

async function loadOlderMessages() {
  if (!selectedChat.value || !selectedChat.value.hasMore || messagesLoading.value) return
  await fetchMessages(selectedChat.value, selectedChat.value.beforeIdCursor)
}

let pollInterval = null

function startMessagePolling(chat) {
  stopMessagePolling()
  pollInterval = setInterval(async () => {
    if (!chat || !selectedChat.value || selectedChat.value.id !== chat.id) return
    try {
      const res = await axios.get(`${API_BASE}/chats/${chat.id}/messages`, {
        headers: authHeaders(),
        params: { before_id: 0, limit: 30 },
      })
      if (res.data.success) {
        const list = res.data.data.messages || []
        const msgs = list.map(mapMessage).reverse()
        const existingIds = new Set(chat.messages.map((m) => m.id))
        const newOnes = msgs.filter((m) => !existingIds.has(m.id))

        if (newOnes.length > 0) {
          const wasNearBottom = shouldAutoScroll()
          chat.messages.push(...newOnes)

          const hasIncoming = newOnes.some((m) => m.sender === 'incoming')
          if (hasIncoming && chat === selectedChat.value) {
            markAsRead(chat)
          }

          if (wasNearBottom) {
            scrollToBottom(true)
          } else {
            showScrollToBottomBtn.value = true
          }
        }
      }
    } catch (err) {
      console.log('POLL MESSAGES ERROR:', err)
    }
  }, 3000)
}

function stopMessagePolling() {
  if (pollInterval) {
    clearInterval(pollInterval)
    pollInterval = null
  }
}

const selectedChat = ref(null)
const newMessage = ref('')

async function selectChat(chat) {
  selectedChat.value = chat
  currentView.value = 'detail'
  selectedMessageIds.value = []
  showChatInfo.value = false
  showScrollToBottomBtn.value = false

  if (chat.id && route.query.open !== String(chat.id)) {
    router.replace({ query: { ...route.query, open: chat.id } })
  }

  await fetchMessages(chat)
  markAsRead(chat)
  scrollToBottom()
  startMessagePolling(chat)
}

function goBackToChats() {
  stopMessagePolling()
  currentView.value = 'chats'
  selectedChat.value = null
  const { open, ...rest } = route.query
  router.replace({ query: rest })
}

async function markAsRead(chat) {
  if (!chat.badge) return
  try {
    await axios.post(
      `${API_BASE}/chats/${chat.id}/read`,
      {},
      { headers: authHeaders() }
    )
    chat.badge = null
  } catch (err) {
    console.log('MARK AS READ ERROR:', err)
  }
}

async function sendMessage() {
  if (!newMessage.value.trim() && pendingImages.value.length === 0) return
  if (!selectedChat.value) return

  if (pendingImages.value.length > 0) {
    const form = new FormData()
    form.append('conversation_id', selectedChat.value.id)
    form.append('type', 'image')
    pendingImages.value.forEach((img) => {
      form.append('attachments', img.file)
    })

    try {
      const res = await axios.post(`${API_BASE}/chats/send`, form, {
        headers: { ...authHeaders(), 'Content-Type': 'multipart/form-data' },
      })
      if (res.data.success) {
        const mapped = mapMessage(res.data.data)
        mapped.sender = 'outgoing' 
        selectedChat.value.messages.push(mapped)
        selectedChat.value.subtitle = 'Sent images'
        scrollToBottom(true)
      }
    } catch (err) {
      console.log('SEND IMAGE ERROR:', err)
    }
    pendingImages.value = []
  }

  if (newMessage.value.trim() !== '') {
    const form = new FormData()
    form.append('conversation_id', selectedChat.value.id)
    form.append('content', newMessage.value.trim())
    form.append('type', 'text')

    try {
      const res = await axios.post(`${API_BASE}/chats/send`, form, {
        headers: { ...authHeaders(), 'Content-Type': 'multipart/form-data' },
      })
      if (res.data.success) {
        const mapped = mapMessage(res.data.data)
        mapped.sender = 'outgoing'
        selectedChat.value.messages.push(mapped) 
        selectedChat.value.subtitle = newMessage.value
        scrollToBottom(true) 
      }
    } catch (err) {
      console.log('SEND TEXT ERROR:', err)
    }
  }

  newMessage.value = ''
}

function sendMockImage() {
  console.log('sendMockImage is a dev helper only')
}

const showVoice = ref(false)

async function handleSendVoice(audioBlob) {
  if (!selectedChat.value) return

  const form = new FormData()
  form.append('conversation_id', selectedChat.value.id)
  form.append('type', 'voice')
  form.append('attachments', audioBlob, 'voice.webm')

  try {
    const res = await axios.post(`${API_BASE}/chats/send`, form, {
      headers: { ...authHeaders(), 'Content-Type': 'multipart/form-data' },
    })
    if (res.data.success) {
      const mapped = mapMessage(res.data.data)
      mapped.sender = 'outgoing'
      selectedChat.value.messages.push(mapped)
      selectedChat.value.subtitle = 'Sent a voice message'
      scrollToBottom(true)
    }
  } catch (err) {
    console.log('SEND VOICE ERROR:', err)
  }

  showVoice.value = false
}

const playingMessageId = ref(null)
const audioElements = {}
const audioProgress = ref({}) 

function setAudioRef(el, msgId) {
  if (el) audioElements[msgId] = el
}

function toggleVoicePlayback(msg) {
  const audio = audioElements[msg.id]
  if (!audio) return

  if (playingMessageId.value && playingMessageId.value !== msg.id) {
    const prev = audioElements[playingMessageId.value]
    if (prev) prev.pause()
  }

  if (playingMessageId.value === msg.id) {
    audio.pause()
    playingMessageId.value = null
  } else {
    audio.play()
    playingMessageId.value = msg.id
  }
}

function onAudioTimeUpdate(msg, event) {
  const audio = event.target
  audioProgress.value[msg.id] = {
    currentTime: audio.currentTime,
    duration: audio.duration || 0,
  }
}

function onAudioEnded(msg) {
  playingMessageId.value = null
  const prev = audioProgress.value[msg.id]
  audioProgress.value[msg.id] = { currentTime: 0, duration: prev?.duration || 0 }
}

function seekVoice(msg, event) {
  const audio = audioElements[msg.id]
  if (!audio || !audio.duration) return
  const rect = event.currentTarget.getBoundingClientRect()
  const ratio = (event.clientX - rect.left) / rect.width
  audio.currentTime = Math.max(0, Math.min(1, ratio)) * audio.duration
}

function formatDuration(seconds) {
  if (!seconds || isNaN(seconds)) return '0:00'
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m}:${s.toString().padStart(2, '0')}`
}

function getActiveWaveWidth(msg) {
  const p = audioProgress.value[msg.id]
  if (!p || !p.duration) return 0
  return Math.min(1000, (p.currentTime / p.duration) * 1000)
}

function getVoiceDurationLabel(msg) {
  const p = audioProgress.value[msg.id]
  if (p && p.duration) {
    const remaining = p.duration - p.currentTime
    return formatDuration(playingMessageId.value === msg.id ? remaining : p.duration)
  }
  return msg.duration || '0:00'
}


async function toggleReaction(msg, reactionTypeId = 1) {
  try {
    const res = await axios.post(
      `${API_BASE}/chats/${msg.id}/react`,
      { reaction_type_id: reactionTypeId },
      { headers: authHeaders() }
    )
    if (res.data.success) {
      return res.data.data.reacted
    }
  } catch (err) {
    console.log('TOGGLE REACTION ERROR:', err)
  }
}

const showStickerPicker = ref(false)

function toggleStickerPicker() {
  showStickerPicker.value = !showStickerPicker.value
}

async function handleSelectSticker(sticker) {
  if (!selectedChat.value) return

  try {
    const res = await axios.post(
      `${API_BASE}/chats/send`,
      (() => {
        const form = new FormData()
        form.append('conversation_id', selectedChat.value.id)
        form.append('type', 'image')
        form.append('content', sticker.url)
        return form
      })(),
      { headers: { ...authHeaders(), 'Content-Type': 'multipart/form-data' } }
    )
    if (res.data.success) {
      const mapped = mapMessage(res.data.data)
      mapped.sender = 'outgoing'
      mapped.image = sticker.url
      mapped.images = [sticker.url]
      selectedChat.value.messages.push(mapped)
    } else {
      selectedChat.value.messages.push({
        id: Date.now(),
        sender: 'outgoing',
        type: 'image',
        image: sticker.url,
        images: [sticker.url],
      })
    }
  } catch (err) {
    console.log('SEND STICKER ERROR:', err)
    selectedChat.value.messages.push({
      id: Date.now(),
      sender: 'outgoing',
      type: 'image',
      image: sticker.url,
      images: [sticker.url],
    })
  }

  selectedChat.value.subtitle = 'Sent a sticker'
  showStickerPicker.value = false
  scrollToBottom(true)
}

function handleClickOutside(event) {
  const stickerContainer = document.querySelector('.chat-sticker-container')
  const stickerBtn = document.querySelector('.sticker-trigger-btn')

  if (showStickerPicker.value) {
    const clickedInsideSticker =
      stickerContainer && stickerContainer.contains(event.target)
    const clickedInsideBtn = stickerBtn && stickerBtn.contains(event.target)

    if (!clickedInsideSticker && !clickedInsideBtn) {
      showStickerPicker.value = false
    }
  }
}

const selectedMessageIds = ref([])

const selectedMessagesData = computed(() => {
  if (!selectedChat.value || !selectedChat.value.messages) return []
  return selectedMessageIds.value.map((idx) => selectedChat.value.messages[idx])
})

function goToForwardPanel(msg, index) {
  selectedMessageIds.value = [index]
  currentView.value = 'forward'
}

function closeForwardPanel() {
  currentView.value = 'detail'
  selectedMessageIds.value = []
}

async function handleForwardSubmit(payload) {
  const { targetChats, messages } = payload
  if (!targetChats || !messages) {
    currentView.value = 'detail'
    return
  }

  for (const targetChat of targetChats) {
    for (const message of messages) {
      const form = new FormData()
      form.append('conversation_id', targetChat.id)
      form.append('type', message.type)
      if (message.text) form.append('content', message.text)
      if (message.id) form.append('forwarded_from_id', message.id)

      try {
        const res = await axios.post(`${API_BASE}/chats/send`, form, {
          headers: { ...authHeaders(), 'Content-Type': 'multipart/form-data' },
        })
        if (res.data.success) {
          targetChat.messages = targetChat.messages || []
          const mapped = mapMessage(res.data.data)
          mapped.sender = 'outgoing'
          targetChat.messages.push(mapped)
          targetChat.subtitle =
            message.type === 'image' ? 'Sent an image' : message.text
        }
      } catch (err) {
        console.log('FORWARD MESSAGE ERROR:', err)
      }
    }
  }

  currentView.value = 'detail'
  selectedMessageIds.value = []
}

const showChatInfo = ref(false)
const activeInfoTab = ref('media')

const infoTabs = [
  {
    key: 'media',
    icon: '<rect x="2" y="4" width="20" height="16" rx="2"/><circle cx="9" cy="10" r="1.5"/><path d="M22 16l-5-5-4 4-3-3-6 6"/>',
  },
  {
    key: 'files',
    icon: '<path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/>',
  },
  {
    key: 'voice',
    icon: '<path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"/><path d="M19 10v1a7 7 0 0 1-14 0v-1"/>',
  },
]

function groupByMonth(items, dateFn) {
  const groups = {}
  items.forEach((item) => {
    const date = dateFn(item) || 'Unknown date'
    let month = date
    const match = date.match(
      /^(January|February|March|April|May|June|July|August|September|October|November|December)/
    )
    if (match) {
      const yearMatch = date.match(/\b(20\d{2})\b/)
      month = `${match[1]} ${yearMatch ? yearMatch[1] : ''}`.trim()
    }
    if (!groups[month]) groups[month] = []
    groups[month].push(item)
  })
  return Object.keys(groups).map((month) => ({ month, items: groups[month] }))
}

const mediaGroups = computed(() => {
  if (!selectedChat.value?.messages) return []
  const mediaMessages = selectedChat.value.messages.filter(
    (msg) => msg.type === 'image' || msg.type === 'video'
  )
  return groupByMonth(mediaMessages, (msg) => msg.mediaDate || msg.timestamp)
})

function openMedia(item) {
  console.log('Open media:', item)
}

const fileGroups = computed(() => {
  if (!selectedChat.value?.messages) return []
  const fileMessages = selectedChat.value.messages.filter(
    (msg) => msg.type === 'file'
  )
  return groupByMonth(fileMessages, (msg) => msg.timestamp).map((g) => ({
    month: g.month,
    items: g.items.map((msg) => ({
      fileName: msg.fileName || 'Unnamed File',
      fileSize: msg.fileSize || 'Unknown size',
      fileUrl: msg.fileUrl || '#',
    })),
  }))
})

const voiceGroups = computed(() => {
  if (!selectedChat.value?.messages) return []
  const voiceMessages = selectedChat.value.messages.filter(
    (msg) => msg.type === 'voice'
  )
  return groupByMonth(voiceMessages, (msg) => msg.timestamp).map((g) => ({
    month: g.month,
    items: g.items.map((msg) => ({
      id: msg.id,
      duration: msg.duration || '0:00',
      voiceUrl: msg.voiceUrl || '#',
      timestamp: msg.timestamp,
    })),
  }))
})

const fileInputRef = ref(null)
const pendingImages = ref([])

function triggerImagePicker() {
  if (fileInputRef.value) {
    fileInputRef.value.click()
  }
}

function handleImageSelected(event) {
  const files = event.target.files
  if (!files || files.length === 0) return

  Array.from(files).forEach((file) => {
    const imageUrl = URL.createObjectURL(file)
    pendingImages.value.push({ file, url: imageUrl })
  })

  event.target.value = ''
}

function removePendingImage(index) {
  pendingImages.value.splice(index, 1)
}

function getImageLayoutClass(images) {
  if (!images || images.length <= 1) return ''
  return `layout-${images.length}`
}

const lightboxImage = ref(null)
const lightboxImagesList = ref([])
const lightboxIndex = ref(0)

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

function nextImage() {
  if (lightboxIndex.value < lightboxImagesList.value.length - 1) {
    lightboxIndex.value++
    lightboxImage.value = lightboxImagesList.value[lightboxIndex.value]
  }
}

function prevImage() {
  if (lightboxIndex.value > 0) {
    lightboxIndex.value--
    lightboxImage.value = lightboxImagesList.value[lightboxIndex.value]
  }
}

function downloadImage(url) {
  const a = document.createElement('a')
  a.href = url
  a.download = 'image.jpg'
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
}

const usersWithoutChat = computed(() => {
  const existingUserIds = new Set(
    mockChats.value.map((c) => c.otherUserId).filter(Boolean)
  )
  return allUsers.value.filter((u) => !existingUserIds.has(u.id))
})

async function openChatFromQuery() {
  const targetId = Number(route.query.open)
  if (!targetId) return

  let chat = mockChats.value.find((c) => c.id === targetId)

  if (!chat) {
    await fetchConversations()
    chat = mockChats.value.find((c) => c.id === targetId)
  }

  if (chat) {
    await selectChat(chat)
  }
}

watch(() => route.query.open, () => {
  openChatFromQuery()
})

onMounted(async () => {
  document.addEventListener('click', handleClickOutside)
  fetchProfile()
  fetchConversations()
  fetchAllUsers()
  await openChatFromQuery()
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
  stopMessagePolling()
  Object.values(audioElements).forEach((audio) => audio && audio.pause())
})
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

.see-more-btn {
  display: block;
  margin-top: 4px;
  background: none;
  border: none;
  padding: 0;
  color: #0084ff;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
}
.see-more-btn:hover {
  text-decoration: underline;
}

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
  position: relative;
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
.msg-image-grid.count-2 {
  grid-template-columns: repeat(2, 1fr);
  grid-auto-rows: 140px;
}
.msg-image-grid.count-3 {
  grid-template-columns: 2fr 1fr;
  grid-template-rows: repeat(2, 70px);
  height: 143px;
}
.msg-image-grid.count-3 .grid-img-item:nth-child(1) {
  grid-row: span 2;
}
.msg-image-grid.count-4 {
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: repeat(2, 110px);
}
.msg-image-grid.count-6 {
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(2, 90px);
}

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

.message-row.outgoing .msg-image-row {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
  width: 100%;
}

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
  width: 48px;
  height: 32px;
  border-radius: 32px;
  background-color: transparent;
  border: none;
  cursor: pointer;
  color: #65676b;
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
  background-color: #e5e6e7;
  color: #050505;
  border-left: 4px solid #0084ff;
  border-top-left-radius: 12px;
  /* border-bottom-right-radius: 12px; */
}

.outgoing-bubble {
  background-color: #297cd416;
  color: #000;
  border-left: none;
  border-right: 4px solid #1B75D2;
  border-top-left-radius: 12px;
  border-top-right-radius: 12px;
}

.image-bubble {
  background: transparent;
  padding: 0;
  overflow: hidden;
  border-radius: 12px;
  max-width: 120px;
  height: 120px;
}
.msg-uploaded-img {
  width: 100%;
  height: 100%;
  border-radius: 12px;
  display: block;
  object-fit: cover;
  background-color: transparent;
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
  box-sizing: border-box;
}

.media-groups {
  width: 100%;
  background-color: #F7F7F7;
}

.media-month-group {
  width: 100%;
}

.media-month-label {
  display: flex;
  align-items: center;
  gap: 2px;
  padding-left: 8px;
  padding-right: 8px;
  padding-top: 14px;
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

.media-grid {
  width: 100%;
  padding-top: 14px;
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 2px;
  padding-left: 8px;
  padding-right: 8px;
}

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
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.12);
}

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

.media-video-wrapper::after {
  content: "";
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.08);
  pointer-events: none;
}

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
  background-color: #0084ff;
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
}

.file-details {
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.file-item svg {
  color: #3b82f6;
}

.file-item:hover {
  background-color: rgba(0, 0, 0, 0.047);
}

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

.lightbox-overlay {
  position: fixed;
  top: 12px;
  left: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}

.lightbox-content {
  position: relative;
  width: 90vw;
  max-width: 1000px;
  height: 85vh;
  max-height: 800px;
  background: #212121;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
}

.lightbox-img-wrapper {
  flex: 1;
  min-height: 0;
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
  color: #1B75D2;
  flex-shrink: 0;
  padding: 0;
  transition: background-color 0.2s ease;
}

.input-send-btn.active {
  display: flex;
}

.input-send-btn.active:hover {
  background-color: rgba(0, 132, 255, 0.1);
}

.bubble.image-bubble.multi-image-bubble {
  max-width: 320px;
  padding: 4px;
  background: transparent;
}

.msg-image-row {
  display: flex;
  flex-wrap: wrap;
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
  position: static !important;
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

.lightbox-bottom-nav {
  display: flex;
  gap: 20px;
  justify-content: center;
  align-items: center;
  width: 100%;
  padding: 0px 0 6px 0;
  margin-top: -10px;
}

.chat-input-wrapper {
  display: inline-block;
}

.sticker-popup-menu {
  position: absolute;
  bottom: 50px;
  left: -100px;
  z-index: 1000;
}

.scroll-to-bottom-btn {
  position: absolute;
  bottom: 90px;
  left: 50%;
  transform: translateX(-50%);
  background: #0084ff;
  color: #fff;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(0,0,0,0.2);
  z-index: 10;
}
</style>