<!-- components/FormInput.vue -->
<template>
    <div class="relative">
      <input
        :type="type"
        :id="id"
        class="w-full px-0 py-2 text-sm font-light border-b border-gray-800 transition-colors focus:border-primary focus:outline-none peer"
        :class="{ 'filled': modelValue }"
        placeholder=" "
        @input="updateValue"
        required
      />
      <label
        :for="id"
        class="absolute left-0 -top-3.5 text-sm text-gray-400 transition-all peer-focus:text-primary peer-placeholder-shown:top-2 peer-placeholder-shown:text-base peer-placeholder-shown:font-light"
      >
        {{ label }}
      </label>
    </div>
  </template>
  
  <script setup lang="ts">
  const props = defineProps<{
    id: string
    label: string
    type: string
    modelValue: string
  }>()
  
  const emit = defineEmits<{
    (e: 'update:modelValue', value: string): void
  }>()
  
  const localValue = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
  })
  
  const updateValue = (e: Event) => {
    localValue.value = (e.target as HTMLInputElement).value
  }
  </script>
  
  <style scoped>
  input::placeholder {
    color: transparent;
  }
  </style>