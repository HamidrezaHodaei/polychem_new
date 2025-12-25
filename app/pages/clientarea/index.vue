<template>
  <div class="h-screen w-screen overflow-hidden relative">
    <!-- Background Image -->
    <div class="absolute inset-0 z-0">
      <img src="/Rafcolor-1.jpg" 
           alt="Background" 
           class="w-auto h-200 object-cover" />
      <div class="absolute inset-0  "></div>
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

      <!-- Main Grid -->
      <div class="flex-1 grid grid-cols-12 gap-4">
        <!-- Solar Panel Alpha Card -->
        <div class="col-span-3 bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-3">
            <h3 class="text-base font-bold text-gray-900">Solar Panel Alpha</h3>
            <button class="w-7 h-7 bg-gray-100 rounded-lg flex items-center justify-center hover:bg-gray-200 transition-all">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
            </button>
          </div>
          <p class="text-xs text-gray-600 mb-4">Monitor the status and performance of your solar panels in real time</p>
          
          <div class="mb-3">
            <span class="inline-block bg-yellow-400 text-gray-900 text-xs font-bold px-2.5 py-1 rounded-lg">SP-001</span>
          </div>

          <!-- Circular Progress Chart -->
          <div class="relative w-36 h-36 mx-auto mb-4">
            <svg class="transform -rotate-90 w-36 h-36">
              <circle cx="72" cy="72" r="60" stroke="#ef4444" stroke-width="12" fill="transparent" stroke-dasharray="120 260" />
              <circle cx="72" cy="72" r="60" stroke="#fbbf24" stroke-width="12" fill="transparent" stroke-dasharray="140 240" stroke-dashoffset="-120" />
              <circle cx="72" cy="72" r="60" stroke="#22c55e" stroke-width="12" fill="transparent" stroke-dasharray="130 250" stroke-dashoffset="-260" />
            </svg>
            <div class="absolute inset-0 flex items-center justify-center">
              <div class="text-center">
                <p class="text-2xl font-bold text-gray-900">85%</p>
              </div>
            </div>
          </div>

          <p class="text-center text-xs text-gray-600">Last Active: 2 hours ago</p>
        </div>

        <!-- Central Image Area -->
          <!-- Central Image Area -->
        <div class="col-span-5 row-span-2 relative">
          <div class="absolute top-[400px] left-0 right-0 h-1/2 bg-white/60 backdrop-blur-xl rounded-3xl p-4 shadow-xl border border-white/40 overflow-hidden z-10">
            <div class="absolute inset-0 bg-gradient-to-br to-blue-400/10"></div>
            <div class="w-full h-full flex items-center justify-center rounded-2xl relative z-10">
              <p class="text-base font-semibold text-gray-900">Hello</p>
            </div>
          </div>
        </div>
<!-- Activity Log Card -->
<div class="col-span-4 bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
  <div class="flex items-center justify-between mb-3">
    <h3 class="text-base font-bold text-gray-700">Activity Log</h3>
    <button class="w-7 h-7 bg-gray-100 rounded-lg flex items-center justify-center hover:bg-gray-200 transition-all">
      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
      </svg>
    </button>
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

        <!-- Current Power Card -->
        <div class="col-span-3 bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-base font-bold text-gray-900">Current Power</h3>
            <button class="w-7 h-7 bg-gray-100 rounded-lg flex items-center justify-center hover:bg-gray-200 transition-all">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
            </button>
          </div>

          <!-- Power Stats -->
          <div class="grid grid-cols-3 gap-3 mb-4">
            <div v-for="power in powerStats" :key="power.label" class="text-center">
              <p class="text-xl font-bold text-gray-900">{{ power.value }}</p>
              <p class="text-[10px] text-gray-500">{{ power.unit }}</p>
              <p class="text-[10px] font-medium text-gray-600 mt-0.5">{{ power.label }}</p>
            </div>
          </div>

          <!-- Power Chart -->
          <div class="flex items-end justify-between h-24 gap-0.5">
            <div v-for="(bar, idx) in powerBars" :key="idx" 
                 class="flex-1 rounded-t transition-all hover:opacity-80"
                 :style="{ height: bar.height + '%', backgroundColor: bar.color }"></div>
          </div>
        </div>

        <!-- Overview Card -->
        <div class="col-span-2 bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-base font-bold text-gray-900">Overview</h3>
            <button class="w-7 h-7 bg-gray-100 rounded-lg flex items-center justify-center hover:bg-gray-200 transition-all">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
            </button>
          </div>

          <div class="space-y-4">
            <!-- RPM -->
            <div>
              <div class="flex items-center justify-between mb-1">
                <span class="text-2xl font-bold text-gray-900">17</span>
                <span class="text-xs text-gray-600">RPM</span>
              </div>
              <p class="text-[10px] text-gray-600 mb-2">Rotor RPM</p>

              <!-- Gauge Chart -->
              <div class="relative w-24 h-24 mx-auto">
                <svg class="transform -rotate-90 w-24 h-24">
                  <circle cx="48" cy="48" r="38" stroke="#fecaca" stroke-width="6" fill="transparent" />
                  <circle cx="48" cy="48" r="38" stroke="#ef4444" stroke-width="6" fill="transparent" 
                          stroke-dasharray="150 238" class="transition-all" />
                </svg>
                <div class="absolute inset-0 flex flex-col items-center justify-center">
                  <p class="text-[9px] text-gray-500">Power</p>
                  <p class="text-lg font-bold text-gray-900">30</p>
                  <p class="text-[9px] text-gray-500">kwh</p>
                </div>
              </div>
            </div>

            <!-- Utilization Rate -->
            <div class="pt-3 border-t border-gray-200">
              <p class="text-2xl font-bold text-gray-900 mb-0.5">91.6%</p>
              <p class="text-[10px] text-gray-600">Utilization Rate</p>
            </div>

            <!-- Operational Hour -->
            <div>
              <div class="flex items-center justify-between">
                <p class="text-2xl font-bold text-gray-900">22</p>
                <span class="text-xs text-gray-600">Hr</span>
              </div>
              <p class="text-[10px] text-gray-600">Operational Hour</p>
            </div>
          </div>
        </div>

        <!-- Earning Card -->
        <div class="col-span-2 bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-base font-bold text-gray-900">Earning</h3>
            <button class="w-7 h-7 bg-gray-100 rounded-lg flex items-center justify-center hover:bg-gray-200 transition-all">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
            </button>
          </div>

          <div class="flex items-center gap-2 mb-4">
            <div class="w-10 h-10 bg-orange-400 rounded-xl flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div>
              <p class="text-2xl font-bold text-gray-900">34.60</p>
              <p class="text-xs text-gray-600">AUD</p>
            </div>
          </div>
          <p class="text-[10px] text-gray-500 mb-5">September 2025</p>

          <!-- CO2 Savings -->
          <div class="pt-4 border-t border-gray-200">
            <div class="flex items-center justify-between mb-2">
              <h4 class="text-xs font-semibold text-gray-900">CO2 Savings Total</h4>
              <button class="w-5 h-5 bg-gray-100 rounded-lg flex items-center justify-center">
                <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                </svg>
              </button>
            </div>
            <div class="flex items-center gap-2">
              <div class="w-10 h-10 bg-yellow-400 rounded-xl flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21v-4m0 0V5a2 2 0 012-2h6.5l1 1H21l-3 6 3 6h-8.5l-1-1H5a2 2 0 00-2 2zm9-13.5V9" />
                </svg>
              </div>
              <div>
                <p class="text-2xl font-bold text-gray-900">5.160</p>
                <p class="text-xs text-gray-600">km</p>
              </div>
            </div>
            <p class="text-[10px] text-gray-500 mt-1">September 2025</p>
          </div>
        </div>

        <!-- CMOS Solar Panel Card -->
        
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SolarDashboard',
  data() {
    return {
      activeTab: 'Dashboard',
      tabs: ['Dashboard', 'Overview', 'Energy', 'Reports'],
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
/* Ensure no scrolling */
* {
  box-sizing: border-box;
}

/* Additional animations */
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
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