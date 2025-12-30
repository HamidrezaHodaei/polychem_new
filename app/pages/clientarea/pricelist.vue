<template>
  <div class="flex-1 overflow-hidden">
    <div class="grid grid-cols-12 gap-4 h-full">
      <!-- Left Side - Form -->
      <div class="col-span-8 grid grid-cols-2 gap-4 h-full">
        <!-- Select Grade -->
        <div class="col-span-2 bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
          <h3 class="text-base font-bold text-gray-900 mb-3">Select Grade</h3>
          <div class="grid grid-cols-4 gap-2">
            <button v-for="grade in grades" :key="grade"
              @click="selectedGrade = grade"
              :class="[
                'py-3 px-4 rounded-xl font-semibold text-base transition-all',
                selectedGrade === grade 
                  ? 'bg-yellow-400 text-gray-900 shadow-lg' 
                  : 'bg-white/60 text-gray-700 hover:bg-white/80'
              ]">
              {{ grade }}
            </button>
          </div>
        </div>

        <!-- Pick Your Amount -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 md:p-6 shadow-xl border border-white/40 min-h-[220px]">
          <h3 class="text-base font-bold text-gray-900 mb-3">Pick Your Amount (Tons)</h3>
          <div class="relative mb-3">
            <input 
              v-model.number="amount"
              type="number"
              min="1"
              step="0.1"
              class="w-full px-4 py-3 rounded-xl border-2 border-gray-200 focus:border-yellow-400 focus:outline-none text-xl font-bold text-gray-900 bg-white/80"
              placeholder="Enter amount">
            <span class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-500 font-semibold text-sm">tons</span>
          </div>
          <input 
            v-model.number="amount"
            type="range" 
            min="1" 
            max="100" 
            step="0.5"
            class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer accent-yellow-400">
          <div class="flex justify-between text-xs text-gray-600 mt-1">
            <span>1 ton</span>
            <span>100 tons</span>
          </div>
        </div>

        <!-- Pick Your Package -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40">
          <h3 class="text-base font-bold text-gray-900 mb-3">Pick Your Package</h3>
          <div class="grid grid-cols-2 gap-3">
            <button 
              @click="selectedPackage = 'jumbo'"
              :class="[
                'py-4 px-4 rounded-xl font-bold text-base transition-all border-2',
                selectedPackage === 'jumbo'
                  ? 'bg-yellow-400 text-gray-900 border-yellow-400 shadow-lg'
                  : 'bg-white/60 text-gray-700 border-gray-200 hover:border-yellow-400'
              ]">
              <div class="flex flex-col items-center gap-1.5">
                <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
                <span>Jumbo Bag</span>
              </div>
            </button>
            <button 
              @click="selectedPackage = '25kg'"
              :class="[
                'py-4 px-4 rounded-xl font-bold text-base transition-all border-2',
                selectedPackage === '25kg'
                  ? 'bg-yellow-400 text-gray-900 border-yellow-400 shadow-lg'
                  : 'bg-white/60 text-gray-700 border-gray-200 hover:border-yellow-400'
              ]">
              <div class="flex flex-col items-center gap-1.5">
                <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                </svg>
                <span>25 KG Bags</span>
              </div>
            </button>
          </div>
        </div>

        <!-- Payment Method -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40 col-span-2">
          <h3 class="text-base font-bold text-gray-900 mb-3">Payment Method</h3>
          <div class="grid grid-cols-5 gap-2">
            <button v-for="method in paymentMethods" :key="method.id"
              @click="selectedPayment = method.id"
              :class="[
                'py-3.5 px-4 rounded-xl font-semibold text-sm transition-all border-2 relative',
                selectedPayment === method.id
                  ? 'bg-yellow-400 text-gray-900 border-yellow-400 shadow-lg'
                  : 'bg-white/60 text-gray-700 border-gray-200 hover:border-yellow-400'
              ]">
              <div class="flex flex-col items-center gap-1">
                <span class="text-center leading-tight">{{ method.label }}</span>
                <span v-if="method.discount" class="text-[10px] bg-green-500 text-white px-1.5 py-0.5 rounded">
                  {{ method.discount }}
                </span>
              </div>
            </button>
          </div>
        </div>

        <!-- Transit Way -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-5 shadow-xl border border-white/40 col-span-2">
          <h3 class="text-base font-bold text-gray-900 mb-3">Pick Your Transit Way</h3>
          <div class="grid grid-cols-2 gap-3">
            <button v-for="transit in transitWays" :key="transit"
              @click="selectedTransit = transit"
              :class="[
                'py-4 px-4 rounded-xl font-bold text-base transition-all border-2 flex flex-col items-center gap-2 w-full',
                selectedTransit === transit
                  ? 'bg-yellow-400 text-gray-900 border-yellow-400 shadow-lg'
                  : 'bg-white/60 text-gray-700 border-gray-200 hover:border-yellow-400'
              ]">
              <svg class="w-8 h-8 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7" />
              </svg>
              <span class="text-center">{{ transit }}</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Right Side - Price Summary -->
      <div class="col-span-4 h-full">
        <div class="bg-gray-900/70 backdrop-blur-xl rounded-3xl p-6 shadow-2xl border border-gray-700/50 h-full flex flex-col">
          <div class="flex items-center gap-2 mb-4">
            <svg class="w-5 h-5 text-yellow-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
            </svg>
            <h3 class="text-lg font-bold text-white">Your Estimate</h3>
          </div>

          <div class="bg-gray-800/50 backdrop-blur-sm rounded-2xl p-5 mb-4">
            <div class="text-center mb-3">
              <p class="text-sm text-gray-400 mb-1">Total Price</p>
              <p class="text-4xl font-bold text-white">€{{ calculatedPrice.toLocaleString() }}</p>
            </div>
            <div class="h-px bg-gray-700/50 my-3"></div>
            <div class="space-y-2">
              <div class="flex justify-between items-center">
                <span class="text-xs text-gray-400">Base Price</span>
                <span class="text-sm font-semibold text-white">€{{ basePrice.toLocaleString() }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-xs text-gray-400">Packaging</span>
                <span class="text-sm font-semibold text-white">€{{ packagingCost.toLocaleString() }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-xs text-gray-400">Transit</span>
                <span class="text-sm font-semibold text-white">€{{ transitCost.toLocaleString() }}</span>
              </div>
              <div v-if="discount > 0" class="flex justify-between items-center text-green-400">
                <span class="text-xs">Discount</span>
                <span class="text-sm font-semibold">-€{{ discount.toLocaleString() }}</span>
              </div>
            </div>
          </div>

          <div class="space-y-2 mb-4 flex-1">
            <div class="bg-gray-800/50 backdrop-blur-sm rounded-xl p-3 flex items-center justify-between">
              <div class="flex items-center gap-2">
                <div class="w-8 h-8 bg-yellow-400 rounded-lg flex items-center justify-center">
                  <svg class="w-4 h-4 text-gray-900" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                  </svg>
                </div>
                <div>
                  <p class="text-[10px] text-gray-400">Grade</p>
                  <p class="text-sm font-bold text-white">{{ selectedGrade }}</p>
                </div>
              </div>
              <div class="text-right">
                <p class="text-[10px] text-gray-400">Amount</p>
                <p class="text-sm font-bold text-white">{{ amount }} tons</p>
              </div>
            </div>

            <div class="bg-gray-800/50 backdrop-blur-sm rounded-xl p-3">
              <div class="flex items-center gap-2 mb-1">
                <svg class="w-3.5 h-3.5 text-yellow-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                </svg>
                <p class="text-[10px] text-gray-400">Package</p>
              </div>
              <p class="text-sm font-bold text-white">{{ selectedPackage === 'jumbo' ? 'Jumbo Bag' : '25 KG Bags' }}</p>
            </div>

            <div class="bg-gray-800/50 backdrop-blur-sm rounded-xl p-3">
              <div class="flex items-center gap-2 mb-1">
                <svg class="w-3.5 h-3.5 text-yellow-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
                <p class="text-[10px] text-gray-400">Payment</p>
              </div>
              <p class="text-sm font-bold text-white">{{ getPaymentLabel() }}</p>
            </div>
          </div>

          <button class="w-full bg-yellow-400 hover:bg-yellow-500 text-gray-900 py-3.5 rounded-2xl font-bold text-base transition-all hover:scale-105 shadow-lg flex items-center justify-center gap-2">
            <span>Start Order</span>
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </button>

          <p class="text-[10px] text-gray-400 text-center mt-3">
            Price valid for 24 hours
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PriceList',
  data() {
    return {
      grades: ['A+', 'A', 'B+', 'B', 'C+', 'C', 'D', 'E'],
      selectedGrade: 'A+',
      amount: 10,
      selectedPackage: 'jumbo',
      selectedPayment: 'cash',
      paymentMethods: [
        { id: 'cash', label: 'Cash Payment', discount: '-5%' },
        { id: '1month', label: '1 Month Terms' },
        { id: '45days', label: '45 Days Terms' },
        { id: '3months', label: '3 Months Terms' },
        { id: 'custom', label: 'Custom Terms' }
      ],
      transitWays: ['Ex-Factory Delivery', 'Delivery to Customer'],
      selectedTransit: 'Ex-Factory Delivery'
    }
  },
  computed: {
    basePrice() {
      const gradeMultiplier = {
        'A+': 120, 'A': 110, 'B+': 100, 'B': 90,
        'C+': 80, 'C': 70, 'D': 60, 'E': 50
      }
      return Math.round(this.amount * (gradeMultiplier[this.selectedGrade] || 100))
    },
    packagingCost() {
      return this.selectedPackage === 'jumbo' ? 50 : 150
    },
    transitCost() {
      const costs = {
        'Ex-Factory Delivery': 100,
        'Delivery to Customer': 150
      }
      return costs[this.selectedTransit] || 0
    },
    discount() {
      if (this.selectedPayment === 'cash') {
        return Math.round(this.basePrice * 0.05)
      }
      return 0
    },
    calculatedPrice() {
      return this.basePrice + this.packagingCost + this.transitCost - this.discount
    }
  },
  methods: {
    getPaymentLabel() {
      const method = this.paymentMethods.find(m => m.id === this.selectedPayment)
      return method ? method.label : 'Not Selected'
    }
  }
}
</script>

<style scoped>
input[type="range"]::-webkit-slider-thumb {
  appearance: none;
  width: 20px;
  height: 20px;
  background: #fbbf24;
  border-radius: 50%;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(251, 191, 36, 0.5);
}

input[type="range"]::-moz-range-thumb {
  width: 20px;
  height: 20px;
  background: #fbbf24;
  border-radius: 50%;
  cursor: pointer;
  border: none;
  box-shadow: 0 2px 8px rgba(251, 191, 36, 0.5);
}
</style>