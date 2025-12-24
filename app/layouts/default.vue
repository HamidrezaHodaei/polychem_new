<template>
  <div class="flex h-screen bg-gray-100">
    <!-- Sidebar Navigation (Navbar سمت چپ) -->
    <aside class="w-64 bg-white shadow-lg">
      <!-- Logo Section -->
      <div class="p-6 border-b">
        <div class="flex items-center space-x-2">
          <div class="w-8 h-8 bg-teal-500 rounded-lg flex items-center justify-center">
            <span class="text-white font-bold text-lg">M</span>
          </div>
          <span class="font-semibold text-gray-800">Medisight</span>
        </div>
      </div>

      <!-- Navigation Menu -->
      <nav class="p-4 space-y-1">
        <NuxtLink
          to="/polychem_manage/dashboard"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
          :class="isActive('/polychem_manage/dashboard') ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          <span class="font-medium">Dashboard</span>
        </NuxtLink>

        <NuxtLink
          to="/polychem_manage/task-manager"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
          :class="isActive('/polychem_manage/task-manager') ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
          </svg>
          <span class="font-medium">Task Manager</span>
        </NuxtLink>

        <NuxtLink
          to="/polychem_manage/product-management"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
          :class="isActive('/polychem_manage/product-management') ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
          <span class="font-medium">Product Management</span>
        </NuxtLink>

        <NuxtLink
          to="/polychem_manage/order-management"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
          :class="isActive('/polychem_manage/order-management') ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
          </svg>
          <span class="font-medium">Order Management</span>
        </NuxtLink>

        <NuxtLink
          to="/polychem_manage/user-management"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
          :class="isActive('/polychem_manage/user-management') ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          <span class="font-medium">User Management</span>
        </NuxtLink>

        <NuxtLink
          to="/polychem_manage/content-management"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
          :class="isActive('/polychem_manage/content-management') ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          <span class="font-medium">Site Content</span>
        </NuxtLink>

        <NuxtLink
          to="/polychem_manage/inbox"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
          :class="isActive('/polychem_manage/inbox') ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
          </svg>
          <span class="font-medium">Inbox</span>
        </NuxtLink>

        <NuxtLink
          to="/polychem_manage/chat"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
          :class="isActive('/polychem_manage/chat') ? 'bg-gray-100 text-gray-900' : 'text-gray-600 hover:bg-gray-50'"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
          </svg>
          <span class="font-medium">Chat</span>
        </NuxtLink>
      </nav>

      <!-- Bottom Navigation (Settings & Logout) -->
      <div class="absolute bottom-0 w-64 p-4 border-t bg-white space-y-1">
        <NuxtLink
          to="/polychem_manage/settings"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span class="font-medium">Settings</span>
        </NuxtLink>

        <button
          @click="handleLogout"
          class="flex items-center space-x-3 px-4 py-3 rounded-lg text-red-600 hover:bg-red-50 transition-colors w-full"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          <span class="font-medium">Logout</span>
        </button>
      </div>
    </aside>

    <!-- Main Content Area -->
    <main class="flex-1 overflow-y-auto">
      <slot />
    </main>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()

const isActive = (path) => {
  return route.path === path
}

const handleLogout = () => {
  if (confirm('Are you sure you want to logout?')) {
    // Add logout logic here
    console.log('Logging out...')
  }
}
</script>