<template>
  <div class="menu-container">
    <!-- Option 1: Remove from profile -->
    <div class="menu-item" @click="handleAction('remove')">
      <div class="icon-wrapper remove-icon">
        <div class="icon-circle"></div>
        <div class="icon-circle-outline"></div>
      </div>
      <span class="menu-text">Remove from profile</span>
    </div>

    <!-- Option 2: Edit details -->
    <div class="menu-item" @click="handleAction('edit')">
      <div class="icon-wrapper edit-icon">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
          <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
        </svg>
      </div>
      <span class="menu-text">Edit details</span>
    </div>

    <!-- Option 3: Delete -->
    <div class="menu-item" @click="handleAction('delete')">
      <div class="icon-wrapper delete-icon">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
          <circle cx="9" cy="7" r="4"></circle>
          <line x1="23" y1="18" x2="17" y2="18"></line>
        </svg>
      </div>
      <span class="menu-text delete-text">Delete</span>
    </div>

    <!-- Option 4: Make private -->
    <div class="menu-item" @click="handleAction('private')">
      <div class="icon-wrapper private-icon">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
          <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
        </svg>
      </div>
      <span class="menu-text">Make private</span>
    </div>
  </div>
</template>

<script setup>
const handleAction = (action) => {
  switch (action) {
    case 'remove':
      console.log('Removing from profile...')
      break
    case 'edit':
      console.log('Editing details...')
      break
    case 'delete':
      console.log('Deleting item...')
      break
    case 'private':
      console.log('Making private...')
      break
  }
}
</script>

<style scoped>
.menu-container {
  background-color: #ffffff;
  position: absolute;
  top: 18px;
  left: 0;
  z-index: 100;
  border-radius: 12px;
  padding: 8px;
  display: flex;
  flex-direction: column;
  width: 240px;
  border: 1px solid rgba(0, 0, 0, 0.08);
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 8px 10px -6px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(8px);
}

.menu-item {
  display: flex;
  align-items: center;
  padding: 6px 12px;
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.2s ease;
}

.menu-item:hover {
  background-color: #f3f4f6; 
  transform: translateX(2px); 
}

.icon-wrapper {
  width: 25px;
  height: 25px;
  margin-right: 12px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 6px;
  background-color: #f8fafc; /* ដាក់ background ឱ្យ icon នីមួយៗមើលទៅមាន ממิติ (Dimension) */
  transition: background-color 0.2s ease;
}

.menu-item:hover .icon-wrapper {
  background-color: #ffffff;
}

.menu-text {
  color: #334155; 
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  font-size: 12px;
  font-weight: 500;
}

.icon-circle {
  width: 16px;
  height: 16px;
  border-radius: 50%;
  border: 2px solid #3b82f6;
  position: absolute;
}

.icon-circle-outline {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  border: 2px solid #3b82f6;
  position: absolute;
}

.edit-icon,
.private-icon {
  color: #64748b; 
}

.delete-icon {
  color: #ef4444; 
}

.menu-item:hover .delete-icon,
.menu-item:hover .delete-text {
  color: #dc2626;
}

.delete-text {
  color: #ef4444; 
}
</style>