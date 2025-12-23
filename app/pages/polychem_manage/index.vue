<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 to-yellow-50 overflow-hidden">
    <div class="w-full h-full">
      <!-- Header -->
      <header class="bg-white/50 backdrop-blur-sm px-8 py-4 flex items-center justify-between border-b border-gray-200/50">
        <div class="flex items-center gap-8">
            <span class="text-xl font-semibold">
                <img src="/polychem wall B.png" alt="" class="h-10">
            </span>
          <nav class="flex gap-4">
            <button @click="setActiveTab('Dashboard')" :class="{'px-4 py-2 rounded-lg text-gray-600 hover:bg-gray-100': true, 'bg-gray-900 text-white': activeTab === 'Dashboard'}">Dashboard</button>
            <button @click="setActiveTab('People')" :class="{'px-4 py-2 rounded-lg text-gray-600 hover:bg-gray-100': true, 'bg-gray-900 text-white': activeTab === 'People'}">People</button>
            <button @click="setActiveTab('Hiring')" :class="{'px-4 py-2 rounded-lg text-gray-600 hover:bg-gray-100': true, 'bg-gray-900 text-white': activeTab === 'Hiring'}">Hiring</button>
            <button @click="setActiveTab('Devices')" :class="{'px-4 py-2 rounded-lg text-gray-600 hover:bg-gray-100': true, 'bg-gray-900 text-white': activeTab === 'Devices'}">Devices</button>
            <button @click="setActiveTab('Apps')" :class="{'px-4 py-2 rounded-lg text-gray-600 hover:bg-gray-100': true, 'bg-gray-900 text-white': activeTab === 'Apps'}">Apps</button>
            <button @click="setActiveTab('Salary')" :class="{'px-4 py-2 rounded-lg text-gray-600 hover:bg-gray-100': true, 'bg-gray-900 text-white': activeTab === 'Salary'}">Salary</button>
            <button @click="setActiveTab('Calendar')" :class="{'px-4 py-2 rounded-lg text-gray-600 hover:bg-gray-100': true, 'bg-gray-900 text-white': activeTab === 'Calendar'}">Calendar</button>
            <button @click="setActiveTab('Reviews')" :class="{'px-4 py-2 rounded-lg text-gray-600 hover:bg-gray-100': true, 'bg-gray-900 text-white': activeTab === 'Reviews'}">Reviews</button>
          </nav>
        </div>
        <div class="flex items-center gap-4">
          <button class="p-2 hover:bg-gray-100 rounded-lg">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </button>
          <button class="p-2 hover:bg-gray-100 rounded-lg">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
          </button>
          <button class="p-2 hover:bg-gray-100 rounded-lg">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
            </svg>
          </button>
        </div>
      </header>

      <!-- Main Content -->
      <main class="p-8">
        <div class="flex items-center justify-between mb-6">
          <h1 class="text-4xl font-bold">{{ activeTab }}</h1>
          <div class="flex items-center gap-2">
            <template v-for="tab in tabs" :key="tab">
              <button
                @click="setActiveTab(tab)"
                :class="[ 'px-3 py-1 rounded-md text-sm', activeTab === tab ? 'bg-gray-900 text-white' : 'bg-white/70 text-gray-700 border border-gray-200' ]"
              >
                {{ tab }}
              </button>
            </template>
          </div>
        </div>

        <!-- Tab content: dynamic components loaded from ./tabs/ -->
        <div>
          <component :is="activeTab" />
        </div>
      </main>
    </div>
  </div>
</template>

<script>
import Dashboard from './tabs/Dashboard.vue'
import People from './tabs/People.vue'
import Hiring from './tabs/Hiring.vue'
import Devices from './tabs/Devices.vue'
import Apps from './tabs/Apps.vue'
import Salary from './tabs/Salary.vue'
import Calendar from './tabs/Calendar.vue'
import Reviews from './tabs/Reviews.vue'

export default {
  components: { Dashboard, People, Hiring, Devices, Apps, Salary, Calendar, Reviews },
  data() {
    return {
      tabs: ['Dashboard','People','Hiring','Devices','Apps','Salary','Calendar','Reviews'],
      activeTab: 'People'
    }
  },
  methods: {
    setActiveTab(tab) {
      this.activeTab = tab
    }
  }
}
</script>