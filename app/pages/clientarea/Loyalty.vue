<template>
  <div class="h-full overflow-y-auto custom-scrollbar pb-6">
    <!-- Page Header -->
   

    <div class="grid grid-cols-12 gap-4">
      <!-- Left Column -->
      <div class="col-span-8 space-y-4">
        <!-- Current Level Card -->
        <div class="bg-gradient-to-br from-gray-400 to-gray-600 rounded-3xl p-6 shadow-xl border border-white/40 relative overflow-hidden">
          <!-- Decorative Elements -->
          <div class="absolute top-0 right-0 w-32 h-32 bg-white/10 rounded-full -mr-16 -mt-16"></div>
          <div class="absolute bottom-0 left-0 w-24 h-24 bg-white/10 rounded-full -ml-12 -mb-12"></div>
          
          <div class="relative z-10">
            <p class="text-xs text-white/80 mb-2">Loyalty level</p>
            <h3 class="text-2xl font-bold text-white mb-1">Nova</h3>
            <p class="text-sm text-white/90">-15% fee</p>
            
            <!-- Diamond Icon -->
            <div class="w-20 h-20 bg-white/20 backdrop-blur-sm rounded-2xl flex items-center justify-center mt-4 mx-auto">
              <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Stats Cards -->
        <div class="grid grid-cols-2 gap-4">
          <!-- USDT Earned -->
          <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
            <div class="flex items-center gap-2 mb-3">
              <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <span class="inline-flex items-center gap-1 bg-yellow-400 text-gray-900 text-xs font-bold px-2 py-1 rounded-lg">
                <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                </svg>
                Boosted
              </span>
            </div>
            <p class="text-xs text-gray-600 mb-2">USDT earned</p>
            <p class="text-3xl font-bold text-gray-900 mb-3">7,500</p>
            <button class="w-full bg-gray-900 hover:bg-gray-800 text-white text-sm font-bold py-2.5 rounded-full transition-all hover:scale-105">
              Withdraw
            </button>
          </div>

          <!-- 90d USDT Volume -->
          <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
            <div class="flex items-center gap-2 mb-3">
              <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
            </div>
            <p class="text-xs text-gray-600 mb-2">90d USDT volume</p>
            <p class="text-3xl font-bold text-gray-900 mb-3">47,335,235</p>
            <button class="w-full bg-white/80 hover:bg-white text-blue-600 text-sm font-bold py-2.5 rounded-full transition-all hover:scale-105 border border-blue-200">
              View graph
            </button>
          </div>
        </div>

        <!-- All Levels Overview -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
          <h3 class="text-lg font-bold text-gray-900 mb-4">All Loyalty Levels</h3>
          
          <div class="space-y-3">
            <div v-for="level in loyaltyLevels" :key="level.name" 
                 class="flex items-center justify-between p-4 rounded-2xl transition-all hover:bg-white/50"
                 :class="level.current ? 'bg-yellow-50 border-2 border-yellow-400' : 'bg-white/40'">
              <div class="flex items-center gap-3">
                <div class="w-12 h-12 rounded-xl flex items-center justify-center"
                     :style="{ background: level.gradient }">
                  <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z" />
                  </svg>
                </div>
                <div>
                  <p class="text-sm font-bold text-gray-900">{{ level.name }}</p>
                  <p class="text-xs text-gray-600">{{ level.volume }}</p>
                </div>
              </div>
              <div class="text-right">
                <p class="text-lg font-bold text-gray-900">{{ level.discount }}</p>
                <p class="text-xs text-gray-600">Fee discount</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column -->
      <div class="col-span-4 space-y-4">
        <!-- Next Level Card -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40 relative overflow-hidden">
          <h3 class="text-base font-bold text-gray-900 mb-6">Next Level</h3>
          
          <!-- Gold Trophy with Circle Progress -->
          <div class="relative w-48 h-48 mx-auto mb-6">
            <!-- Circle Progress -->
            <svg class="transform -rotate-90 w-48 h-48 absolute inset-0">
              <circle cx="96" cy="96" r="85" stroke="#fef3c7" stroke-width="14" fill="transparent" />
              <circle cx="96" cy="96" r="85" stroke="#fbbf24" stroke-width="14" fill="transparent" 
                      stroke-dasharray="380 534" class="transition-all duration-1000" stroke-linecap="round" />
            </svg>
            
            <!-- Center Content -->
            <div class="absolute inset-0 flex items-center justify-center">
              <div class="text-center">
                <!-- Gold Trophy Icon -->
                <div class="w-20 h-20 bg-gradient-to-br from-yellow-300 via-yellow-400 to-yellow-600 rounded-2xl flex items-center justify-center mb-3 mx-auto shadow-lg transform rotate-12 hover:rotate-0 transition-transform duration-300">
                  <svg class="w-12 h-12 text-white transform -rotate-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z" />
                  </svg>
                </div>
                
                <p class="text-sm font-bold text-yellow-600 mb-1">Sirius</p>
                <div class="mb-1">
                  <span class="text-2xl font-bold text-gray-900">47.3</span>
                  <span class="text-sm text-gray-600">/60K</span>
                </div>
                <span class="inline-block bg-yellow-100 text-yellow-700 text-xs font-semibold px-3 py-1 rounded-full">
                  Cashback 0.2%
                </span>
              </div>
            </div>
          </div>

          <!-- Fee Discount -->
          <div class="flex items-center justify-between pt-4 border-t border-gray-200">
            <p class="text-sm text-gray-600">Fee discount:</p>
            <p class="text-3xl font-bold text-gray-900">20%</p>
          </div>
        </div>

        <!-- Benefits Card -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
          <h3 class="text-base font-bold text-gray-900 mb-4">Your Benefits</h3>
          
          <div class="space-y-3">
            <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-2xl p-4 border border-green-200">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-green-500 rounded-xl flex items-center justify-center flex-shrink-0">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12zm-7 4h14M5 12a2 2 0 110-4h14a2 2 0 110 4M5 12v7a2 2 0 002 2h10a2 2 0 002-2v-7" />
                  </svg>
                </div>
                <div>
                  <p class="text-xl font-bold text-gray-900">2,340</p>
                  <p class="text-xs text-gray-600">Reward Points</p>
                </div>
              </div>
            </div>

            <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-2xl p-4 border border-purple-200">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-purple-500 rounded-xl flex items-center justify-center flex-shrink-0">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                  </svg>
                </div>
                <div>
                  <p class="text-xl font-bold text-gray-900">12</p>
                  <p class="text-xs text-gray-600">Exclusive Offers</p>
                </div>
              </div>
            </div>

            <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-2xl p-4 border border-blue-200">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500 rounded-xl flex items-center justify-center flex-shrink-0">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                </div>
                <div>
                  <p class="text-xl font-bold text-gray-900">⚡ Fast</p>
                  <p class="text-xs text-gray-600">Priority Support</p>
                </div>
              </div>
            </div>

            <div class="bg-gradient-to-br from-orange-50 to-orange-100 rounded-2xl p-4 border border-orange-200">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-orange-500 rounded-xl flex items-center justify-center flex-shrink-0">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div>
                  <p class="text-xl font-bold text-gray-900">24/7</p>
                  <p class="text-xs text-gray-600">Access Anytime</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoyaltyPage',
  data() {
    return {
      loyaltyLevels: [
        {
          name: 'Meteor',
          volume: '0 - 10K USDT',
          discount: '0%',
          gradient: 'linear-gradient(135deg, #cd7f32 0%, #8b5a2b 100%)',
          current: false
        },
        {
          name: 'Nova ',
          volume: '10K - 50K USDT',
          discount: '15%',
          gradient: 'linear-gradient(135deg, #c0c0c0 0%, #808080 100%)',
          current: true
        },
        {
          name: 'Sirius ',
          volume: '50K - 100K USDT',
          discount: '20%',
          gradient: 'linear-gradient(135deg, #fbbf24 0%, #f59e0b 100%)',
          current: false
        },
        {
          name: 'Polaris ',
          volume: '100K - 500K USDT',
          discount: '25%',
          gradient: 'linear-gradient(135deg, #e5e7eb 0%, #9ca3af 100%)',
          current: false
        }
      ]
    }
  }
}
</script>

<style scoped>
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.space-y-3 > * {
  animation: fadeIn 0.3s ease-out;
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