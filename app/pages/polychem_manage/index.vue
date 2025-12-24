<template>
  <div class="flex min-h-screen bg-gradient-to-br from-gray-50 to-yellow-50">
    <!-- Sidebar (pinned fixed, fixed height) -->
    <aside class="fixed left-0 top-0 h-screen w-64 bg-white border-r border-gray-200 flex flex-col z-40 pointer-events-auto">
      <!-- Logo -->
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-center gap-2">
          
          <img src="/polychem wall B.png" alt="PolyChem Logo" class="w-auto h-12" />
          
        </div>
      </div>

      <!-- Navigation -->
      <nav class="flex-1 p-4 overflow-auto">
        <div class="space-y-1">
          <button
            v-for="item in navItems"
            :key="item.name"
            type="button"
            @click.stop="setActiveTab(item.component, item.name)"
            :class="[
              'w-full flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-colors',
              activeComponent === item.component
                ? 'bg-gray-900 text-white'
                : 'text-gray-700 hover:bg-gray-100'
            ]"
          >
            <img :src="item.icon" :alt="item.name" class="w-5 h-5" />
            <span>{{ item.name }}</span>
          </button>
        </div>

        <!-- Tools Section -->
        <div class="mt-8 pt-8 border-t border-gray-200">
          <p class="px-4 text-xs font-semibold text-gray-400 uppercase mb-2">
            Tools
          </p>
          <div class="space-y-1">
            <button type="button" class="w-full flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-100">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <span>Settings</span>
            </button>
          </div>
        </div>
      </nav>

      <!-- User Profile -->
      <div class="p-4 border-t border-gray-200">
        <button type="button" class="w-full flex items-center gap-3 px-2 py-2 rounded-lg hover:bg-gray-50" @click.stop>
           <div class="w-10 h-10 bg-gray-300 rounded-full flex items-center justify-center">
             <svg class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
             </svg>
           </div>
           <div class="flex-1 text-left">
             <p class="text-sm font-medium text-gray-900">ali</p>
             <p class="text-xs text-gray-500">alifox@polychemmb.com</p>
           </div>
           <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
           </svg>
         </button>
       </div>
     </aside>

    <!-- Main Content (leave space for fixed sidebar) -->
    <div class="flex-1 flex flex-col ml-64 relative z-0">
      <!-- Header -->
      <header class="bg-white/50 backdrop-blur-sm px-8 py-4 flex items-center justify-between border-b border-gray-200/50">
        <h1 class="text-4xl font-bold text-gray-900">{{ activeLabel }}</h1>
        <div class="flex items-center gap-4">
          <button type="button" class="p-2 hover:bg-gray-100 rounded-lg transition-colors">
            <svg class="w-5 h-5 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </button>
          <button type="button" class="p-2 hover:bg-gray-100 rounded-lg transition-colors">
            <svg class="w-5 h-5 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
          </button>
          <button type="button" class="p-2 hover:bg-gray-100 rounded-lg transition-colors">
            <svg class="w-5 h-5 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
            </svg>
          </button>
        </div>
      </header>

      <!-- Main Content Area -->
      <main class="flex-1 p-8 overflow-auto">
        <component :is="activeComponent" />
      </main>
    </div>
  </div>
</template>

<script>
import Dashboard from './tabs/Dashboard.vue'
import Task from './tabs/Task.vue'
import UserManagment from './tabs/UserManagment.vue'
import Pricing from './tabs/Pricingmanagment.vue'
import orders from './tabs/orders.vue' 
import inbox from './tabs/inbox.vue'
import Calendar from './tabs/Calendar.vue'
import Client from './tabs/client.vue'

export default {
  components: { Dashboard, Task, UserManagment,Pricing, orders, inbox, Calendar, Client },
  data() {
    return {
      activeComponent: 'Task',
      activeLabel: 'Task',
      navItems: [
        { name: 'Dashboard', component: 'Dashboard', icon: 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"%3E%3Crect x="3" y="3" width="7" height="7"%3E%3C/rect%3E%3Crect x="14" y="3" width="7" height="7"%3E%3C/rect%3E%3Crect x="14" y="14" width="7" height="7"%3E%3C/rect%3E%3Crect x="3" y="14" width="7" height="7"%3E%3C/rect%3E%3C/svg%3E' 
        },
        { 
          name: 'Task', 
          component: 'Task',
          icon: 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"%3E%3Cpath d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"%3E%3C/path%3E%3Ccircle cx="9" cy="7" r="4"%3E%3C/circle%3E%3Cpath d="M22 21v-2a4 4 0 0 0-3-3.87"%3E%3C/path%3E%3Cpath d="M16 3.13a4 4 0 0 1 0 7.75"%3E%3C/path%3E%3C/svg%3E' 
        },
        { 
          name: 'User Managment', 
          component: 'UserManagment',
          icon: 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"%3E%3Cpath d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"%3E%3C/path%3E%3Ccircle cx="9" cy="7" r="4"%3E%3C/circle%3E%3Cline x1="19" y1="8" x2="19" y2="14"%3E%3C/line%3E%3Cline x1="22" y1="11" x2="16" y2="11"%3E%3C/line%3E%3C/svg%3E' 
        },
         { 
          name: 'Client Management', 
          component: 'Client Management',
          icon: 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"%3E%3Cpolygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"%3E%3C/polygon%3E%3C/svg%3E' 
        },
        { 
          name: 'Customer Pricing', 
          component: 'Customer Pricing',
          icon: 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"%3E%3Crect x="2" y="3" width="20" height="14" rx="2" ry="2"%3E%3C/rect%3E%3Cline x1="8" y1="21" x2="16" y2="21"%3E%3C/line%3E%3Cline x1="12" y1="17" x2="12" y2="21"%3E%3C/line%3E%3C/svg%3E' 
        },

        { 
          name: 'Orders', 
          component: 'orders',
          icon: 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"%3E%3Crect x="3" y="3" width="7" height="7"%3E%3C/rect%3E%3Crect x="14" y="3" width="7" height="7"%3E%3C/rect%3E%3Crect x="14" y="14" width="7" height="7"%3E%3C/rect%3E%3Crect x="3" y="14" width="7" height="7"%3E%3C/rect%3E%3C/svg%3E' 
        },

        { 
          name: 'Inbox', 
          component: 'inbox',
          icon: 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"%3E%3Cline x1="12" y1="1" x2="12" y2="23"%3E%3C/line%3E%3Cpath d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"%3E%3C/path%3E%3C/svg%3E' 
        },
        { 
          name: 'Calendar', 
          component: 'Calendar',
          icon: 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"%3E%3Crect x="3" y="4" width="18" height="18" rx="2" ry="2"%3E%3C/rect%3E%3Cline x1="16" y1="2" x2="16" y2="6"%3E%3C/line%3E%3Cline x1="8" y1="2" x2="8" y2="6"%3E%3C/line%3E%3Cline x1="3" y1="10" x2="21" y2="10"%3E%3C/line%3E%3C/svg%3E' 
        }
       
      ]
    }
  },
  methods: {
    setActiveTab(componentKey, label) {
      this.activeComponent = componentKey
      this.activeLabel = label
    }
  }
}
</script>

<style scoped>
/* Additional custom styles if needed */
</style>