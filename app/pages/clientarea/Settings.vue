<template>
  <div class="h-full overflow-y-auto custom-scrollbar pb-6">
    

    <div class="grid grid-cols-12 gap-4">
      <!-- Left Column - Profile & Account -->
      <div class="col-span-8 space-y-4">
        <!-- Profile Card -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-bold text-gray-900">Profile Information</h3>
            <button v-if="!editingProfile" @click="editingProfile = true" class="text-sm text-yellow-600 hover:text-yellow-700 font-semibold transition-colors">
              Edit
            </button>
            <button v-else @click="editingProfile = false" class="text-sm text-red-600 hover:text-red-700 font-semibold transition-colors">
              Cancel
            </button>
          </div>

          <div class="flex items-center gap-6 mb-6">
            <!-- Profile Picture -->
            <div class="relative">
              <img v-if="profileImage" :src="profileImage" alt="Profile" class="w-24 h-24 rounded-2xl object-cover shadow-lg">
              <div v-else class="w-24 h-24 bg-gradient-to-br from-yellow-300 to-yellow-500 rounded-2xl flex items-center justify-center shadow-lg">
                <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>
              <button v-if="editingProfile" @click="triggerFileInput" class="absolute bottom-0 right-0 bg-yellow-400 hover:bg-yellow-500 w-7 h-7 rounded-full flex items-center justify-center shadow-lg transition-all hover:scale-110">
                <svg class="w-4 h-4 text-gray-900" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                </svg>
              </button>
              <input ref="fileInput" type="file" accept="image/*" @change="handleProfileImageChange" class="hidden">
            </div>
            
            <div class="flex-1">
              <p class="text-lg font-bold text-gray-900">{{ formData.firstName }} {{ formData.lastName }}</p>
              <p class="text-sm text-gray-600">{{ formData.email }}</p>
              <p class="text-xs text-gray-500 mt-2">Account Type: Business</p>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">First Name</label>
              <input v-model="formData.firstName" type="text" :disabled="!editingProfile" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80 disabled:opacity-60">
            </div>
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Last Name</label>
              <input v-model="formData.lastName" type="text" :disabled="!editingProfile" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80 disabled:opacity-60">
            </div>
            <div class="bg-white/40 rounded-2xl p-4 col-span-2">
              <label class="text-xs text-gray-600 block mb-2">Email Address</label>
              <input v-model="formData.email" type="email" :disabled="!editingProfile" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80 disabled:opacity-60">
            </div>
            <div class="bg-white/40 rounded-2xl p-4 col-span-2">
              <label class="text-xs text-gray-600 block mb-2">Phone Number</label>
              <input v-model="formData.phone" type="tel" :disabled="!editingProfile" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80 disabled:opacity-60">
            </div>
          </div>

          <button v-if="editingProfile" @click="saveProfile" class="w-full mt-4 bg-yellow-400 hover:bg-yellow-500 text-gray-900 py-3 rounded-2xl font-bold transition-all hover:scale-105 shadow-lg">
            Save Profile
          </button>
        </div>

        <!-- Security Card -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
          <h3 class="text-lg font-bold text-gray-900 mb-6">Security Settings</h3>

          <!-- Change Username -->
          <div class="mb-6 pb-6 border-b border-gray-200">
            <h4 class="text-base font-semibold text-gray-900 mb-4">Username</h4>
            <div class="bg-white/40 rounded-2xl p-4 mb-4">
              <label class="text-xs text-gray-600 block mb-2">Current Username</label>
              <input v-model="formData.username" type="text" :disabled="!editingUsername" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80 disabled:opacity-60">
            </div>
            <button v-if="!editingUsername" @click="editingUsername = true" class="bg-gray-900 hover:bg-gray-800 text-white px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105">
              Change Username
            </button>
            <div v-else class="flex gap-2">
              <button @click="saveUsername" class="bg-yellow-400 hover:bg-yellow-500 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105">
                Save
              </button>
              <button @click="editingUsername = false" class="bg-gray-300 hover:bg-gray-400 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105">
                Cancel
              </button>
            </div>
          </div>

          <!-- Change Password -->
          <div class="mb-6 pb-6 border-b border-gray-200">
            <h4 class="text-base font-semibold text-gray-900 mb-4">Password</h4>
            <div class="space-y-4 mb-4">
              <div class="bg-white/40 rounded-2xl p-4">
                <label class="text-xs text-gray-600 block mb-2">Current Password</label>
                <input v-model="passwordForm.current" type="password" placeholder="••••••••" :disabled="!editingPassword" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 bg-white/80 disabled:opacity-60">
              </div>
              <div class="bg-white/40 rounded-2xl p-4">
                <label class="text-xs text-gray-600 block mb-2">New Password</label>
                <input v-model="passwordForm.new" type="password" placeholder="••••••••" :disabled="!editingPassword" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 bg-white/80 disabled:opacity-60">
              </div>
              <div class="bg-white/40 rounded-2xl p-4">
                <label class="text-xs text-gray-600 block mb-2">Confirm New Password</label>
                <input v-model="passwordForm.confirm" type="password" placeholder="••••••••" :disabled="!editingPassword" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 bg-white/80 disabled:opacity-60">
              </div>
            </div>
            <div v-if="!editingPassword" class="flex gap-2">
              <button @click="editingPassword = true" class="bg-yellow-400 hover:bg-yellow-500 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105 shadow-lg">
                Change Password
              </button>
            </div>
            <div v-else class="flex gap-2">
              <button @click="savePassword" class="bg-yellow-400 hover:bg-yellow-500 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105 shadow-lg">
                Update Password
              </button>
              <button @click="editingPassword = false" class="bg-gray-300 hover:bg-gray-400 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105">
                Cancel
              </button>
            </div>
          </div>

          <!-- Two-Factor Authentication -->
          <div>
            <div class="flex items-center justify-between mb-4">
              <div>
                <h4 class="text-base font-semibold text-gray-900">Two-Factor Authentication</h4>
                <p class="text-xs text-gray-600 mt-1">Add an extra layer of security to your account</p>
              </div>
              <button class="bg-green-500 hover:bg-green-600 text-white px-5 py-2 rounded-full text-sm font-semibold transition-all hover:scale-105">
                Enable
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column - Addresses -->
      <div class="col-span-4 space-y-4">
        <!-- Main Office Address -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-base font-bold text-gray-900">Main Office</h3>
            <button v-if="!editingOffice" @click="editingOffice = true" class="text-sm text-yellow-600 hover:text-yellow-700 font-semibold">Edit</button>
            <button v-else @click="editingOffice = false" class="text-sm text-red-600 hover:text-red-700 font-semibold">Cancel</button>
          </div>

          <div v-if="editingOffice" class="space-y-4 mb-4">
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Address</label>
              <input v-model="officeData.address" type="text" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80">
            </div>
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Postal Code</label>
              <input v-model="officeData.postalCode" type="text" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80">
            </div>
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Email</label>
              <input v-model="officeData.email" type="email" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80">
            </div>
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Phone</label>
              <input v-model="officeData.phone" type="tel" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80">
            </div>
            <div class="flex gap-2">
              <button @click="saveOffice" class="bg-yellow-400 hover:bg-yellow-500 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105">
                Save
              </button>
              <button @click="editingOffice = false" class="bg-gray-300 hover:bg-gray-400 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105">
                Cancel
              </button>
            </div>
          </div>

          <div v-else class="space-y-3 mb-4">
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-yellow-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <div>
                <p class="text-xs text-gray-600 mb-1">Address</p>
                <p class="text-sm font-semibold text-gray-900">{{ officeData.address }}</p>
              </div>
            </div>
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-yellow-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
              </svg>
              <div>
                <p class="text-xs text-gray-600 mb-1">Postal Code</p>
                <p class="text-sm font-semibold text-gray-900">{{ officeData.postalCode }}</p>
              </div>
            </div>
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-yellow-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
              <div>
                <p class="text-xs text-gray-600 mb-1">Email</p>
                <p class="text-sm font-semibold text-gray-900">{{ officeData.email }}</p>
              </div>
            </div>
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-yellow-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
              </svg>
              <div>
                <p class="text-xs text-gray-600 mb-1">Phone</p>
                <p class="text-sm font-semibold text-gray-900">{{ officeData.phone }}</p>
              </div>
            </div>
          </div>

          <button v-if="!editingOffice" @click="editingOffice = true" class="w-full bg-white/80 hover:bg-white text-gray-900 border-2 border-gray-200 py-2.5 rounded-2xl font-semibold text-sm transition-all hover:scale-105">
            Update Address
          </button>
        </div>

        <!-- Factory Address -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-base font-bold text-gray-900">Factory Address</h3>
            <button v-if="!editingFactory" @click="editingFactory = true" class="text-sm text-yellow-600 hover:text-yellow-700 font-semibold">Edit</button>
            <button v-else @click="editingFactory = false" class="text-sm text-red-600 hover:text-red-700 font-semibold">Cancel</button>
          </div>

          <div v-if="editingFactory" class="space-y-4 mb-4">
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Address</label>
              <input v-model="factoryData.address" type="text" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80">
            </div>
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Postal Code</label>
              <input v-model="factoryData.postalCode" type="text" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80">
            </div>
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Email</label>
              <input v-model="factoryData.email" type="email" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80">
            </div>
            <div class="bg-white/40 rounded-2xl p-4">
              <label class="text-xs text-gray-600 block mb-2">Phone</label>
              <input v-model="factoryData.phone" type="tel" class="w-full px-3 py-2 rounded-lg border border-gray-200 focus:border-yellow-400 focus:outline-none text-gray-900 font-semibold bg-white/80">
            </div>
            <div class="flex gap-2">
              <button @click="saveFactory" class="bg-yellow-400 hover:bg-yellow-500 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105">
                Save
              </button>
              <button @click="editingFactory = false" class="bg-gray-300 hover:bg-gray-400 text-gray-900 px-6 py-2.5 rounded-full font-semibold text-sm transition-all hover:scale-105">
                Cancel
              </button>
            </div>
          </div>

          <div v-else class="space-y-3 mb-4">
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-yellow-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <div>
                <p class="text-xs text-gray-600 mb-1">Address</p>
                <p class="text-sm font-semibold text-gray-900">{{ factoryData.address }}</p>
              </div>
            </div>
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-yellow-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
              </svg>
              <div>
                <p class="text-xs text-gray-600 mb-1">Postal Code</p>
                <p class="text-sm font-semibold text-gray-900">{{ factoryData.postalCode }}</p>
              </div>
            </div>
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-yellow-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
              <div>
                <p class="text-xs text-gray-600 mb-1">Email</p>
                <p class="text-sm font-semibold text-gray-900">{{ factoryData.email }}</p>
              </div>
            </div>
            <div class="flex items-start gap-3">
              <svg class="w-5 h-5 text-yellow-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
              </svg>
              <div>
                <p class="text-xs text-gray-600 mb-1">Phone</p>
                <p class="text-sm font-semibold text-gray-900">{{ factoryData.phone }}</p>
              </div>
            </div>
          </div>

          <button v-if="!editingFactory" @click="editingFactory = true" class="w-full bg-white/80 hover:bg-white text-gray-900 border-2 border-gray-200 py-2.5 rounded-2xl font-semibold text-sm transition-all hover:scale-105">
            Update Address
          </button>
        </div>

        <!-- Notifications -->
        <div class="bg-white/60 backdrop-blur-xl rounded-3xl p-6 shadow-xl border border-white/40">
          <h3 class="text-base font-bold text-gray-900 mb-4">Notifications</h3>
          
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <label class="text-sm font-semibold text-gray-900">Email Notifications</label>
              <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-yellow-400">
                <span class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform translate-x-6"></span>
              </button>
            </div>
            <div class="flex items-center justify-between">
              <label class="text-sm font-semibold text-gray-900">SMS Notifications</label>
              <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-gray-300">
                <span class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"></span>
              </button>
            </div>
            <div class="flex items-center justify-between">
              <label class="text-sm font-semibold text-gray-900">Order Updates</label>
              <button class="relative inline-flex h-6 w-11 items-center rounded-full bg-yellow-400">
                <span class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform translate-x-6"></span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SettingsPage',
  data() {
    return {
      editingProfile: false,
      editingUsername: false,
      editingPassword: false,
      editingOffice: false,
      editingFactory: false,
      profileImage: null,
      formData: {
        firstName: 'علی',
        lastName: 'احمدی',
        email: 'ali.ahmadi@polychem.ir',
        phone: '+98 912 345 6789',
        username: 'ali.ahmadi'
      },
      passwordForm: {
        current: '',
        new: '',
        confirm: ''
      },
      officeData: {
        address: 'خیابان انقلاب، تهران',
        postalCode: '1234567',
        email: 'office@polychem.ir',
        phone: '+98 21 1234 5678'
      },
      factoryData: {
        address: 'بزرگراه تهران - قزوین، کرج',
        postalCode: '3134567',
        email: 'factory@polychem.ir',
        phone: '+98 261 4567 8910'
      }
    }
  },
  methods: {
    triggerFileInput() {
      this.$refs.fileInput.click()
    },
    handleProfileImageChange(event) {
      const file = event.target.files[0]
      if (file) {
        const reader = new FileReader()
        reader.onload = (e) => {
          this.profileImage = e.target.result
          localStorage.setItem('profileImage', this.profileImage)
        }
        reader.readAsDataURL(file)
      }
    },
    saveProfile() {
      localStorage.setItem('userProfile', JSON.stringify(this.formData))
      this.editingProfile = false
      alert('پروفایل با موفقیت ذخیره شد')
    },
    saveUsername() {
      localStorage.setItem('userProfile', JSON.stringify(this.formData))
      this.editingUsername = false
      alert('نام کاربری با موفقیت تغییر کرد')
    },
    savePassword() {
      if (this.passwordForm.new !== this.passwordForm.confirm) {
        alert('رمز عبور جدید با تأیید آن مطابقت ندارد')
        return
      }
      this.editingPassword = false
      this.passwordForm = { current: '', new: '', confirm: '' }
      alert('رمز عبور با موفقیت تغییر کرد')
    },
    saveOffice() {
      localStorage.setItem('officeData', JSON.stringify(this.officeData))
      this.editingOffice = false
      alert('آدرس دفتر مرکزی با موفقیت ذخیره شد')
    },
    saveFactory() {
      localStorage.setItem('factoryData', JSON.stringify(this.factoryData))
      this.editingFactory = false
      alert('آدرس کارخانه با موفقیت ذخیره شد')
    }
  },
  mounted() {
    const savedProfile = localStorage.getItem('userProfile')
    if (savedProfile) {
      this.formData = JSON.parse(savedProfile)
    }
    const savedOffice = localStorage.getItem('officeData')
    if (savedOffice) {
      this.officeData = JSON.parse(savedOffice)
    }
    const savedFactory = localStorage.getItem('factoryData')
    if (savedFactory) {
      this.factoryData = JSON.parse(savedFactory)
    }
    const savedImage = localStorage.getItem('profileImage')
    if (savedImage) {
      this.profileImage = savedImage
    }
  }
}
</script>

<style scoped>
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
