<template>
  <div class="space-y-4 pb-6">
    <div v-for="order in orders" :key="order.id" 
         class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center gap-4">
          <h3 class="text-sm font-semibold text-gray-700">Order <span class="text-yellow-600">{{ order.id }}</span></h3>
          <span class="text-xs text-gray-500">Order Placed: {{ order.placedDate }}</span>
        </div>
        <button class="bg-yellow-400 hover:bg-yellow-500 text-gray-900 px-6 py-2 rounded-full text-sm font-bold transition-all hover:scale-105 flex items-center gap-2">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
          </svg>
          TRACK ORDER
        </button>
      </div>

      <div class="space-y-4 mb-6">
        <div v-for="(item, idx) in order.items" :key="idx" 
             class="flex items-center gap-4 p-4 bg-white/40 rounded-2xl transition-all hover:bg-white/50">
          <div class="w-20 h-20 bg-gradient-to-br from-yellow-100 to-yellow-200 rounded-xl flex items-center justify-center flex-shrink-0">
            <svg class="w-10 h-10 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>

          <div class="flex-1">
            <h4 class="text-base font-bold text-gray-900 mb-1">{{ item.name }}</h4>
            <p class="text-xs text-gray-600 mb-2">By {{ item.designer }}</p>
            <div class="flex items-center gap-4 text-xs text-gray-700">
              <span>Size: <span class="font-semibold">{{ item.size }}</span></span>
              <span>Qty: <span class="font-semibold">{{ item.quantity }}</span></span>
              <span class="font-bold text-gray-900">Rs. {{ formatNumber(item.price) }}</span>
            </div>
          </div>

          <div class="text-right">
            <p class="text-xs text-gray-600 mb-1">Status</p>
            <p class="text-sm font-bold" :class="getStatusColor(item.status)">{{ item.status }}</p>
            <p class="text-xs text-gray-600 mt-2">Expected by</p>
            <p class="text-sm font-semibold text-gray-900">{{ order.expectedDate }}</p>
          </div>
        </div>
      </div>

      <div class="flex items-center justify-between pt-4 border-t border-gray-200">
        <button class="text-sm text-gray-700 hover:text-red-600 font-semibold flex items-center gap-2 transition-colors">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
          CANCEL ORDER
        </button>
        <div class="text-right">
          <p class="text-xs text-gray-600">Paid using credit card ending with {{ order.cardLastDigits }}</p>
          <p class="text-xl font-bold text-gray-900 mt-1">Rs. {{ formatNumber(order.total) }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'OrderComponent',
  props: {
    orders: {
      type: Array,
      required: true
    }
  },
  methods: {
    getStatusColor(status) {
      const colors = {
        'In - Transit': 'text-yellow-600',
        'Processing': 'text-blue-600',
        'Delivered': 'text-green-600',
        'Cancelled': 'text-red-600'
      }
      return colors[status] || 'text-gray-700'
    },
    formatNumber(n) {
      return n.toLocaleString()
    }
  }
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar { width: 6px }
</style>
