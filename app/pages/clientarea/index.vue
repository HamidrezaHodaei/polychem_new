<template>
  <div class="h-screen w-screen overflow-hidden relative">
    <!-- Background Image -->
    <div class="absolute inset-0 z-0">
      <img src="/Rafcolor-1.jpg" 
           alt="Background" 
           class="w-full h-full object-cover grayscale" />
      <div class="absolute inset-0 bg-gradient-to-t from-yellow-400/20 to-transparent"></div>
    </div>

    <!-- Content -->
    <div class="relative z-10 h-full flex flex-col p-6">
      <!-- Header -->
      <header class="bg-gray-200/50 backdrop-blur-lg rounded-3xl px-6 py-3 mb-4 flex items-center justify-between shadow-lg border border-white/30">
        <div class="flex items-center gap-2">
          <img src="/english logo W1.png" alt="" class="h-8">
        </div>
        
        <nav class="flex items-center gap-6">
          <button v-for="tab in tabs" :key="tab" 
            :class="['px-4 py-2 font-medium transition-all text-sm relative group']"
            @click="activeTab = tab">
            <span :class="activeTab === tab ? 'text-white' : 'text-white'">{{ tab }}</span>
            
            <!-- خط زیر برای حالت اکتیو -->
            <span v-if="activeTab === tab" 
              class="absolute bottom-0 left-0 right-0 h-0.5 bg-yellow-400 rounded-full"></span>
            
            <!-- خط زیر برای hover -->
            <span v-else 
              class="absolute bottom-0 left-1/2 right-1/2 h-0.5 bg-yellow-400/50 rounded-full group-hover:left-0 group-hover:right-0 transition-all duration-300"></span>
          </button>
        </nav>

        <div class="flex items-center gap-2">
          <button class="w-9 h-9 bg-white/40 hover:bg-white/70 rounded-full flex items-center justify-center transition-all hover:scale-105">
            <svg class="w-5 h-5 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
          </button>
          <button class="w-9 h-9 bg-white/40 hover:bg-white/70 rounded-full flex items-center justify-center transition-all hover:scale-105">
            <svg class="w-5 h-5 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
            </svg>
          </button>
          <img src="https://i.pravatar.cc/150?img=12" alt="User" class="w-9 h-9 rounded-full border-2 border-white shadow-sm" />
        </div>
      </header>

      <!-- Content Container -->
      <div class="flex-1 overflow-hidden h-full flex flex-col">
        <!-- Orders Component -->
        <OrdersPage v-if="activeTab === 'Orders'" />

        <!-- Price List Component -->
        <PriceList v-else-if="activeTab === 'Price List'" />

        <!-- Dashboard Content -->
        <div v-else-if="activeTab === 'Dashboard'" class="grid grid-cols-12 gap-4 h-full overflow-y-auto custom-scrollbar pb-6">
        <!-- Recent Orders Summary -->
        <div class="col-span-6 bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-bold text-gray-900">Recent Orders</h3>
            <button @click="activeTab = 'Orders'" class="text-sm text-yellow-600 hover:text-yellow-700 font-semibold">View All</button>
          </div>
          
          <div class="space-y-3">
            <div v-for="order in recentOrders" :key="order.id" class="flex items-center justify-between p-4 bg-white/40 rounded-2xl hover:bg-white/50 transition-all">
              <div class="flex-1">
                <p class="font-semibold text-gray-900">{{ order.id }}</p>
                <p class="text-xs text-gray-600">{{ order.placedDate }}</p>
              </div>
              <div class="text-right">
                <p class="font-bold text-gray-900">Rs. {{ order.total.toLocaleString() }}</p>
                <span class="text-xs font-semibold px-2.5 py-1 rounded-lg"
                      :class="order.status === 'In - Transit' ? 'bg-yellow-100 text-yellow-700' : order.status === 'Processing' ? 'bg-blue-100 text-blue-700' : 'bg-green-100 text-green-700'">
                  {{ order.status }}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Loyalty Status Summary -->
        <div class="col-span-6 bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-bold text-gray-900">Loyalty Status</h3>
            <button @click="activeTab = 'Loyalty'" class="text-sm text-yellow-600 hover:text-yellow-700 font-semibold">View All</button>
          </div>
          
          <div class="space-y-4">
            <div class="bg-gradient-to-br from-gray-400 to-gray-600 rounded-2xl p-4 text-white">
              <p class="text-xs text-white/80 mb-1">Current Level</p>
              <h4 class="text-xl font-bold mb-2">Nova</h4>
              <p class="text-sm">Fee Discount: <span class="font-bold">-15%</span></p>
            </div>
            
            <div class="grid grid-cols-2 gap-3">
              <div class="bg-green-50 rounded-xl p-3 border border-green-200">
                <p class="text-xs text-gray-600 mb-1">USDT Earned</p>
                <p class="text-2xl font-bold text-gray-900">7,500</p>
              </div>
              <div class="bg-blue-50 rounded-xl p-3 border border-blue-200">
                <p class="text-xs text-gray-600 mb-1">90d Volume</p>
                <p class="text-2xl font-bold text-gray-900">47M</p>
              </div>
            </div>
          </div>
        </div>


        <!-- Current Power Card -->
        <div class="col-span-3 bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-base font-bold text-gray-900">Activity Log</h3>
          </div>
          <p class="text-xs text-gray-600 mb-4">Recent login attempts and user activities</p>
          
          <div class="h-[230px] overflow-y-auto pr-2 custom-scrollbar space-y-2">
            <div v-for="(log, idx) in activityLogs" :key="idx" 
                 class="flex items-start justify-between p-2.5 rounded-xl transition-all hover:bg-white/70 flex-shrink-0"
                 :class="log.success ? 'bg-green-50/60' : 'bg-red-50/60'">
              <div class="flex items-start gap-2 flex-1">
                <div class="w-7 h-7 rounded-lg flex items-center justify-center flex-shrink-0"
                     :class="log.success ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'">
                  <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path v-if="log.success" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-xs font-semibold text-gray-900">{{ log.action }}</p>
                  <p class="text-[10px] text-gray-500 mb-1">{{ log.time }}</p>
                  <div class="flex items-center gap-1.5 text-[9px] text-gray-600">
                    <span class="flex items-center gap-0.5">
                      <svg class="w-2.5 h-2.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                      </svg>
                      {{ log.location }}
                    </span>
                    <span class="text-gray-400">•</span>
                    <span class="font-mono">{{ log.ip }}</span>
                  </div>
                </div>
              </div>
              <span class="text-[9px] font-semibold px-2 py-0.5 rounded-lg whitespace-nowrap ml-2"
                    :class="log.success ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'">
                {{ log.status }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Placeholder for other tabs -->
      <div v-else-if="activeTab === 'Loyalty'" class="flex-1 overflow-hidden">
        <LoyaltyPage />
      </div>

      <div v-else-if="activeTab === 'Settings'" class="flex-1 overflow-hidden">
        <SettingsPage />
      </div>
      </div>
    </div>
  </div>
</template>

<script>
import OrdersPage from './order.vue'
import PriceList from './pricelist.vue'
import LoyaltyPage from './Loyalty.vue'
import SettingsPage from './Settings.vue'

export default {
  name: 'SolarDashboard',
  components: {
    OrdersPage,
    PriceList,
    LoyaltyPage,
    SettingsPage
  },
  data() {
    return {
      activeTab: 'Dashboard',
      tabs: ['Dashboard', 'Orders', 'Price List', 'Loyalty','Settings'],
      recentOrders: [
        {
          id: '#RO374915036',
          placedDate: 'Thu, 17th Nov 16',
          total: 3010,
          status: 'In - Transit'
        },
        {
          id: '#RO374915037',
          placedDate: 'Mon, 21st Nov 16',
          total: 4200,
          status: 'Processing'
        },
        {
          id: '#RO374915038',
          placedDate: 'Wed, 23rd Nov 16',
          total: 890,
          status: 'Delivered'
        }
      ],
      activityLogs: [
        { action: 'Login Successful', time: '2 mins ago', status: 'Success', success: true, location: 'Tehran, Iran', ip: '185.143.234.12' },
        { action: 'Login Failed', time: '15 mins ago', status: 'Failed', success: false, location: 'Unknown', ip: '192.168.1.45' },
        { action: 'Login Successful', time: '1 hour ago', status: 'Success', success: true, location: 'Isfahan, Iran', ip: '185.143.234.12' },
        { action: 'Password Changed', time: '3 hours ago', status: 'Success', success: true, location: 'Tehran, Iran', ip: '185.143.234.12' },
        { action: 'Login Failed', time: '5 hours ago', status: 'Failed', success: false, location: 'Unknown', ip: '103.245.67.89' },
        { action: 'Login Successful', time: '8 hours ago', status: 'Success', success: true, location: 'Shiraz, Iran', ip: '185.143.234.12' },
      ],
      powerStats: [
        { value: '4.43', unit: 'kW', label: 'Sun' },
        { value: '4.43', unit: 'kW', label: 'Returns' },
        { value: '4.43', unit: 'kW', label: 'Electrics' }
      ],
      powerBars: Array.from({ length: 40 }, (_, i) => ({
        height: Math.random() * 100,
        color: ['#fbbf24', '#ef4444', '#22c55e'][i % 3]
      }))
    }
  }
}
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: rgba(243, 244, 246, 0.5);
  border-radius: 10px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: linear-gradient(180deg, #f8cf48 0%, #f8cf48 100%);
  border-radius: 10px;
  transition: background 0.3s ease;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(180deg, #f8cf48 0%, #f8cf48 100%);
}
</style>