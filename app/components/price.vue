<template>
  <section class="p-6 bg-white/60 rounded-2xl border border-gray-200">
    <h2 class="text-xl font-semibold mb-4">Price List</h2>
    <p class="text-sm text-gray-600 mb-4">قیمت‌های محصولات و خدمات در این بخش نمایش داده می‌شود.</p>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div v-for="(p, idx) in prices" :key="idx" class="p-4 bg-white/40 rounded-lg">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-base font-bold text-gray-900">{{ p.name }}</h3>
            <p class="text-xs text-gray-600">{{ p.description }}</p>
          </div>
          <div class="text-right">
            <p class="text-sm font-semibold text-gray-900">Rs. {{ formatNumber(p.price) }}</p>
            <p class="text-xs text-gray-500">/ {{ p.unit }}</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'PriceComponent',
  props: {
    prices: {
      type: Array,
      default: () => ([]) 
    }
  },
  methods: {
    formatNumber(n) { return n.toLocaleString() }
  }
}
</script>

<style scoped>
</style>
