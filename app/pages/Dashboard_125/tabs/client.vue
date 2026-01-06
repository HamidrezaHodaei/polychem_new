<template>
  <div class="space-y-6">
    <!-- Header Section -->
    <div class="flex items-center justify-between">
      <div>
        <p class="text-sm text-gray-600">Manage your clients and their information here.</p>
      </div>
      <button 
        type="button"
        @click="showAddClientModal = true"
        class="flex items-center gap-2 px-4 py-2 bg-gray-900 text-white rounded-lg hover:bg-gray-800 transition-colors"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Add Client
      </button>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg p-6 border border-gray-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Total Clients</p>
            <p class="text-3xl font-bold text-gray-900 mt-1">{{ clients.length }}</p>
          </div>
          <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg p-6 border border-gray-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Active Clients</p>
            <p class="text-3xl font-bold text-gray-900 mt-1">{{ activeClients }}</p>
          </div>
          <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg p-6 border border-gray-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Total Orders</p>
            <p class="text-3xl font-bold text-gray-900 mt-1">{{ totalOrders }}</p>
          </div>
          <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg p-6 border border-gray-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Total Revenue</p>
            <p class="text-3xl font-bold text-gray-900 mt-1">${{ totalRevenue }}</p>
          </div>
          <div class="w-12 h-12 bg-yellow-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Search and Filters -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex items-center gap-4">
        <div class="flex-1 relative">
          <svg class="w-5 h-5 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search clients..."
            class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
          />
        </div>
        <select 
          v-model="filterStatus"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
        >
          <option value="">All Status</option>
          <option value="Active">Active</option>
          <option value="Inactive">Inactive</option>
        </select>
      </div>
    </div>

    <!-- Clients Table -->
    <div class="bg-white rounded-lg border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase">
                <input type="checkbox" class="rounded border-gray-300" />
              </th>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase">Client Name</th>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase">Email</th>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase">Company</th>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase">Phone</th>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase">Status</th>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase">Total Orders</th>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase">Date Joined</th>
              <th class="text-left py-3 px-6 text-xs font-medium text-gray-500 uppercase"></th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr 
              v-for="client in filteredClients" 
              :key="client.id"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="py-4 px-6">
                <input type="checkbox" class="rounded border-gray-300" />
              </td>
              <td class="py-4 px-6">
                <div class="flex items-center gap-3">
                  <div class="w-10 h-10 rounded-full bg-gradient-to-br from-blue-400 to-purple-500 flex items-center justify-center text-white font-semibold">
                    {{ client.name.charAt(0) }}
                  </div>
                  <div>
                    <p class="font-medium text-gray-900">{{ client.name }}</p>
                  </div>
                </div>
              </td>
              <td class="py-4 px-6 text-sm text-gray-600">{{ client.email }}</td>
              <td class="py-4 px-6">
                <span class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-700">
                  {{ client.company }}
                </span>
              </td>
              <td class="py-4 px-6 text-sm text-gray-600">{{ client.phone }}</td>
              <td class="py-4 px-6">
                <span 
                  :class="[
                    'inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-medium',
                    client.status === 'Active' ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'
                  ]"
                >
                  <span class="w-1.5 h-1.5 rounded-full" :class="client.status === 'Active' ? 'bg-green-500' : 'bg-red-500'"></span>
                  {{ client.status }}
                </span>
              </td>
              <td class="py-4 px-6 text-sm text-gray-600">{{ client.totalOrders }}</td>
              <td class="py-4 px-6 text-sm text-gray-600">{{ client.dateJoined }}</td>
              <td class="py-4 px-6">
                <button 
                  type="button"
                  @click="editClient(client)"
                  class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
                >
                  <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="flex items-center justify-between px-6 py-4 border-t border-gray-200">
        <p class="text-sm text-gray-600">
          Showing {{ filteredClients.length }} of {{ clients.length }} clients
        </p>
        <div class="flex items-center gap-2">
          <button type="button" class="px-3 py-1 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors">
            Previous
          </button>
          <button type="button" class="px-3 py-1 bg-gray-900 text-white rounded-lg">1</button>
          <button type="button" class="px-3 py-1 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors">2</button>
          <button type="button" class="px-3 py-1 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors">3</button>
          <button type="button" class="px-3 py-1 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors">
            Next
          </button>
        </div>
      </div>
    </div>

    <!-- Add/Edit Client Modal -->
    <div 
      v-if="showAddClientModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      @click.self="showAddClientModal = false"
    >
      <div class="bg-white rounded-lg w-full max-w-2xl max-h-[90vh] flex flex-col">
        <!-- Modal Header -->
        <div class="flex items-center justify-between p-6 border-b border-gray-200">
          <h2 class="text-xl font-bold text-gray-900">Add New Client</h2>
          <button 
            type="button"
            @click="showAddClientModal = false"
            class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
          >
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Modal Body - Scrollable -->
        <div class="flex-1 overflow-y-auto p-6">
          <div class="space-y-6">
            <!-- Personal Information Section -->
            <div>
              <h3 class="text-sm font-semibold text-gray-900 mb-4 flex items-center gap-2">
                <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
                Personal Information
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Client Name *
                  </label>
                  <input 
                    v-model="newClient.name"
                    type="text"
                    placeholder="Enter client name"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Email Address *
                  </label>
                  <input 
                    v-model="newClient.email"
                    type="email"
                    placeholder="client@example.com"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  />
                </div>
              </div>
            </div>

            <!-- Contact Information Section -->
            <div>
              <h3 class="text-sm font-semibold text-gray-900 mb-4 flex items-center gap-2">
                <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
                Contact Information
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Phone Number *
                  </label>
                  <input 
                    v-model="newClient.phone"
                    type="tel"
                    placeholder="+1 (555) 000-0000"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Mobile Number
                  </label>
                  <input 
                    v-model="newClient.mobile"
                    type="tel"
                    placeholder="+1 (555) 000-0000"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  />
                </div>
              </div>
            </div>

            <!-- Company Information Section -->
            <div>
              <h3 class="text-sm font-semibold text-gray-900 mb-4 flex items-center gap-2">
                <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
                Company Information
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Company Name *
                  </label>
                  <input 
                    v-model="newClient.company"
                    type="text"
                    placeholder="Company name"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Industry *
                  </label>
                  <input 
                    v-model="newClient.industry"
                    type="text"
                    placeholder="Industry type"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Website
                  </label>
                  <input 
                    v-model="newClient.website"
                    type="url"
                    placeholder="https://example.com"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Status *
                  </label>
                  <select 
                    v-model="newClient.status"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  >
                    <option value="Active">Active</option>
                    <option value="Inactive">Inactive</option>
                  </select>
                </div>
              </div>
            </div>

            <!-- Address Information Section -->
            <div>
              <h3 class="text-sm font-semibold text-gray-900 mb-4 flex items-center gap-2">
                <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                Address
              </h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Street Address
                  </label>
                  <input 
                    v-model="newClient.address"
                    type="text"
                    placeholder="123 Main Street"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                  />
                </div>
                <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">
                      City
                    </label>
                    <input 
                      v-model="newClient.city"
                      type="text"
                      placeholder="City"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">
                      State/Province
                    </label>
                    <input 
                      v-model="newClient.state"
                      type="text"
                      placeholder="State"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">
                      Postal Code
                    </label>
                    <input 
                      v-model="newClient.postalCode"
                      type="text"
                      placeholder="00000"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-900"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Modal Footer -->
        <div class="flex gap-3 p-6 border-t border-gray-200">
          <button 
            type="button"
            @click="showAddClientModal = false"
            class="flex-1 px-4 py-2.5 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors font-medium"
          >
            Cancel
          </button>
          <button 
            type="button"
            @click="addClient"
            class="flex-1 px-4 py-2.5 bg-gray-900 text-white rounded-lg hover:bg-gray-800 transition-colors font-medium"
          >
            Add Client
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: '',
      filterStatus: '',
      showAddClientModal: false,
      newClient: {
        name: '',
        email: '',
        phone: '',
        mobile: '',
        company: '',
        industry: '',
        website: '',
        status: 'Active',
        address: '',
        city: '',
        state: '',
        postalCode: ''
      },
      clients: [
        {
          id: 1,
          name: 'Acme Corporation',
          email: 'contact@acme.com',
          company: 'Acme Corp',
          phone: '+1 (555) 123-4567',
          status: 'Active',
          totalOrders: 12,
          dateJoined: 'Jan 15, 2023'
        },
        {
          id: 2,
          name: 'Tech Solutions Inc',
          email: 'info@techsol.com',
          company: 'Tech Solutions',
          phone: '+1 (555) 234-5678',
          status: 'Active',
          totalOrders: 8,
          dateJoined: 'Feb 20, 2023'
        },
        {
          id: 3,
          name: 'Global Traders',
          email: 'sales@globaltraders.com',
          company: 'Global Traders',
          phone: '+1 (555) 345-6789',
          status: 'Active',
          totalOrders: 15,
          dateJoined: 'Mar 10, 2023'
        },
        {
          id: 4,
          name: 'Premium Services Ltd',
          email: 'contact@premium.com',
          company: 'Premium Services',
          phone: '+1 (555) 456-7890',
          status: 'Inactive',
          totalOrders: 5,
          dateJoined: 'Apr 5, 2023'
        },
        {
          id: 5,
          name: 'Enterprise Solutions',
          email: 'hello@enterprise.com',
          company: 'Enterprise',
          phone: '+1 (555) 567-8901',
          status: 'Active',
          totalOrders: 22,
          dateJoined: 'May 12, 2023'
        },
        {
          id: 6,
          name: 'Innovation Hub',
          email: 'team@innovationhub.com',
          company: 'Innovation Hub',
          phone: '+1 (555) 678-9012',
          status: 'Active',
          totalOrders: 9,
          dateJoined: 'Jun 8, 2023'
        }
      ]
    }
  },
  computed: {
    filteredClients() {
      return this.clients.filter(client => {
        const matchesSearch = client.name.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
                            client.email.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
                            client.company.toLowerCase().includes(this.searchQuery.toLowerCase())
        const matchesStatus = !this.filterStatus || client.status === this.filterStatus
        return matchesSearch && matchesStatus
      })
    },
    activeClients() {
      return this.clients.filter(c => c.status === 'Active').length
    },
    totalOrders() {
      return this.clients.reduce((sum, client) => sum + client.totalOrders, 0)
    },
    totalRevenue() {
      return (this.clients.length * 1250).toLocaleString()
    }
  },
  methods: {
    addClient() {
      if (!this.newClient.name || !this.newClient.email || !this.newClient.phone || 
          !this.newClient.company || !this.newClient.industry) {
        alert('Please fill in all required fields marked with *')
        return
      }

      const client = {
        id: this.clients.length + 1,
        name: this.newClient.name,
        email: this.newClient.email,
        phone: this.newClient.phone,
        company: this.newClient.company,
        status: this.newClient.status,
        totalOrders: 0,
        dateJoined: new Date().toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
      }
      this.clients.unshift(client)
      this.showAddClientModal = false
      this.newClient = {
        name: '',
        email: '',
        phone: '',
        mobile: '',
        company: '',
        industry: '',
        website: '',
        status: 'Active',
        address: '',
        city: '',
        state: '',
        postalCode: ''
      }
    },
    editClient(client) {
      console.log('Edit client:', client)
    }
  }
}
</script>
