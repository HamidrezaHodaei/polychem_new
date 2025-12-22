<template>
  <section class="bg-white p-8 rounded shadow-sm">
    <h2 class="text-3xl font-bold text-[#848484] mb-4">Send us a message</h2>
    <p class="text-sm text-[#848484] mb-6">Please fill out the form below and we'll get back to you as soon as possible.</p>

    <form @submit.prevent="handleSubmit" novalidate class="space-y-4">
      <!-- Name / Company row -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Full name *</label>
          <input v-model="form.name" type="text" class="w-full border rounded px-3 py-2" :class="{'ring-2 ring-red-300': errors.name}" placeholder="Your full name" />
          <p v-if="errors.name" class="text-xs text-red-600 mt-1">{{ errors.name }}</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Company</label>
          <input v-model="form.company" type="text" class="w-full border rounded px-3 py-2" placeholder="Company (optional)" />
        </div>
      </div>

      <!-- Email / Phone row -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Email *</label>
          <input v-model="form.email" type="email" class="w-full border rounded px-3 py-2" :class="{'ring-2 ring-red-300': errors.email}" placeholder="you@example.com" />
          <p v-if="errors.email" class="text-xs text-red-600 mt-1">{{ errors.email }}</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Phone</label>
          <input v-model="form.phone" type="tel" class="w-full border rounded px-3 py-2" placeholder="+1 555 555 5555" />
        </div>
      </div>

      <!-- Country / Subject row -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Country</label>
          <select v-model="form.country" class="w-full border rounded px-3 py-2">
            <option value="">Select country (optional)</option>
            <option>United States</option>
            <option>United Kingdom</option>
            <option>Germany</option>
            <option>France</option>
            <option>China</option>
            <option>India</option>
            <option>Other</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Subject *</label>
          <input v-model="form.subject" type="text" class="w-full border rounded px-3 py-2" :class="{'ring-2 ring-red-300': errors.subject}" placeholder="Subject of your message" />
          <p v-if="errors.subject" class="text-xs text-red-600 mt-1">{{ errors.subject }}</p>
        </div>
      </div>

      <!-- Message -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Message *</label>
        <textarea v-model="form.message" rows="6" class="w-full border rounded px-3 py-2" :class="{'ring-2 ring-red-300': errors.message}" placeholder="Describe your request, project or question..."></textarea>
        <p v-if="errors.message" class="text-xs text-red-600 mt-1">{{ errors.message }}</p>
      </div>

      <!-- File upload -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Attachment (optional)</label>
        <input @change="handleFile" type="file" class="w-full" />
        <p v-if="fileName" class="text-xs text-gray-600 mt-1">Selected: {{ fileName }}</p>
      </div>

      <!-- Preferred contact method -->
      <div class="flex items-center gap-4">
        <label class="text-sm font-medium text-gray-700">Preferred contact</label>
        <label class="flex items-center gap-2 text-sm"><input type="radio" value="email" v-model="form.contactMethod" /> Email</label>
        <label class="flex items-center gap-2 text-sm"><input type="radio" value="phone" v-model="form.contactMethod" /> Phone</label>
      </div>

      <!-- Consent -->
      <div class="flex items-start gap-3">
        <input id="consent" type="checkbox" v-model="form.consent" class="mt-1" />
        <label for="consent" class="text-sm text-gray-700">I agree to be contacted regarding my inquiry. *</label>
      </div>
      <p v-if="errors.consent" class="text-xs text-red-600 mt-1">{{ errors.consent }}</p>

      <!-- Actions -->
      <div class="flex items-center gap-4">
        <button type="submit" :disabled="loading" class="px-4 py-2 rounded bg-[#FFCD05] hover:bg-[#e6b804] disabled:opacity-60">
          <span v-if="!loading">Send message</span>
          <span v-else>Sending...</span>
        </button>
        <button type="button" @click="resetForm" class="px-4 py-2 rounded border">Reset</button>
        <div v-if="success" class="text-sm text-green-700">Message sent — thank you!</div>
        <div v-if="submitError" class="text-sm text-red-600">{{ submitError }}</div>
      </div>
    </form>
  </section>
</template>

<script setup>
import { reactive, ref } from 'vue'

const form = reactive({
  name: '',
  company: '',
  email: '',
  phone: '',
  country: '',
  subject: '',
  message: '',
  contactMethod: 'email',
  file: null,
  consent: false
})

const errors = reactive({
  name: '',
  email: '',
  subject: '',
  message: '',
  consent: ''
})

const loading = ref(false)
const success = ref(false)
const submitError = ref('')
const fileName = ref('')

const validate = () => {
  errors.name = form.name.trim() ? '' : 'Name is required.'
  const emailRe = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  errors.email = emailRe.test(form.email) ? '' : 'Please enter a valid email address.'
  errors.subject = form.subject.trim() ? '' : 'Subject is required.'
  errors.message = form.message.trim().length >= 10 ? '' : 'Message must be at least 10 characters.'
  errors.consent = form.consent ? '' : 'Consent is required.'
  return !errors.name && !errors.email && !errors.subject && !errors.message && !errors.consent
}

const handleFile = (event) => {
  const f = event.target.files && event.target.files[0]
  if (f) {
    form.file = f
    fileName.value = f.name
  } else {
    form.file = null
    fileName.value = ''
  }
}

const resetForm = () => {
  form.name = ''
  form.company = ''
  form.email = ''
  form.phone = ''
  form.country = ''
  form.subject = ''
  form.message = ''
  form.contactMethod = 'email'
  form.file = null
  form.consent = false
  fileName.value = ''
  success.value = false
  submitError.value = ''
  Object.keys(errors).forEach(k => errors[k] = '')
}

const handleSubmit = async () => {
  submitError.value = ''
  success.value = false
  if (!validate()) return

  loading.value = true
  try {
    const fd = new FormData()
    fd.append('name', form.name)
    fd.append('company', form.company)
    fd.append('email', form.email)
    fd.append('phone', form.phone)
    fd.append('country', form.country)
    fd.append('subject', form.subject)
    fd.append('message', form.message)
    fd.append('contactMethod', form.contactMethod)
    fd.append('consent', form.consent ? 'yes' : 'no')
    if (form.file) fd.append('attachment', form.file)

    // Placeholder API endpoint — replace with your backend endpoint.
    const res = await fetch('/api/contact', {
      method: 'POST',
      body: fd
    })

    if (!res.ok) {
      const text = await res.text().catch(() => null)
      throw new Error(text || 'Server error')
    }

    success.value = true
    resetForm()
  } catch (err) {
    submitError.value = typeof err === 'string' ? err : (err.message || 'Failed to send message')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* small visual adjustments for the form */
section { max-width: 900px; }
input[type="file"] { font-size: 0.9rem; }
.ring-2 { box-shadow: 0 0 0 4px rgba(255, 0, 0, 0.08); }
</style>
