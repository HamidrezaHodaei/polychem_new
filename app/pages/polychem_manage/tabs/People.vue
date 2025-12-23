<template>
  <section>
    <!-- Stats Bar -->
    <div class="flex items-center gap-8 mb-8">
      <div class="flex items-center gap-4 flex-1">
        <div class="flex-1 flex items-center bg-gray-900 text-white rounded-full overflow-hidden h-12">
          <div class="px-6 text-sm font-medium">Interviews 25%</div>
        </div>
        <div class="flex-1 flex items-center bg-yellow-300 rounded-full overflow-hidden h-12">
          <div class="px-6 text-sm font-medium">Hired 51%</div>
        </div>
      </div>
      <div class="flex items-center gap-4">
        <div class="px-6 py-3 bg-white/50 rounded-full border border-gray-300">
          <span class="text-sm font-medium">Project time 10%</span>
        </div>
        <div class="px-6 py-3 bg-white/50 rounded-full border border-gray-300">
          <span class="text-sm font-medium">Output 14%</span>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-4">
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Columns</option>
        </select>
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Department</option>
        </select>
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Site</option>
        </select>
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Lifecycle</option>
        </select>
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Status</option>
        </select>
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Entity</option>
        </select>
        <div class="relative">
          <input
            type="text"
            placeholder="Search"
            class="pl-10 pr-4 py-2 bg-white/70 rounded-lg border border-gray-300 w-64"
          />
          <svg class="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
      </div>
      <div class="flex items-center gap-3">
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Directory</option>
        </select>
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Org Chat</option>
        </select>
        <select class="px-4 py-2 bg-white/70 rounded-lg border border-gray-300">
          <option>Insights</option>
        </select>
      </div>
    </div>

    <!-- Action Buttons -->
    <div class="flex items-center gap-3 mb-6">
      <button class="p-2 bg-white rounded-lg hover:bg-gray-50 border border-gray-300">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
      </button>
      <button class="p-2 bg-white rounded-lg hover:bg-gray-50 border border-gray-300">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
      </button>
      <button class="px-4 py-2 bg-white rounded-lg hover:bg-gray-50 border border-gray-300 flex items-center gap-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
        </svg>
        Export
      </button>
    </div>

    <!-- Table -->
    <div class="bg-white/70 rounded-2xl overflow-hidden border border-gray-200">
      <table class="w-full">
        <thead>
          <tr class="border-b border-gray-200">
            <th class="px-6 py-4 text-left">
              <input type="checkbox" class="rounded" />
            </th>
            <th class="px-6 py-4 text-left text-sm font-medium text-gray-600">Name</th>
            <th class="px-6 py-4 text-left text-sm font-medium text-gray-600">Job title</th>
            <th class="px-6 py-4 text-left text-sm font-medium text-gray-600">Department</th>
            <th class="px-6 py-4 text-left text-sm font-medium text-gray-600">Site</th>
            <th class="px-6 py-4 text-left text-sm font-medium text-gray-600">Salary</th>
            <th class="px-6 py-4 text-left text-sm font-medium text-gray-600">Start date</th>
            <th class="px-6 py-4 text-left text-sm font-medium text-gray-600">Lifecycle</th>
            <th class="px-6 py-4 text-left text-sm font-medium text-gray-600">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="person in people" :key="person.id" :class="{'bg-yellow-100': person.highlighted}" class="border-b border-gray-200 hover:bg-gray-50">
            <td class="px-6 py-4">
              <input type="checkbox" :checked="person.checked" class="rounded" />
            </td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-3">
                <img :src="person.avatar" :alt="person.name" class="w-10 h-10 rounded-full" />
                <span class="font-medium">{{ person.name }}</span>
              </div>
            </td>
            <td class="px-6 py-4">{{ person.jobTitle }}</td>
            <td class="px-6 py-4">{{ person.department }}</td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-2">
                <span>{{ person.flag }}</span>
                <span>{{ person.site }}</span>
              </div>
            </td>
            <td class="px-6 py-4">${{ person.salary.toLocaleString() }}</td>
            <td class="px-6 py-4">{{ person.startDate }}</td>
            <td class="px-6 py-4">{{ person.lifecycle }}</td>
            <td class="px-6 py-4">
              <span :class="getStatusClass(person.status)" class="px-3 py-1 rounded-full text-sm font-medium">
                {{ person.status }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>

<script>
export default {
  name: 'PeopleTab',
  data() {
    return {
      people: [
        {
          id: 1,
          name: 'Anatoly Belik',
          avatar: 'https://i.pravatar.cc/150?img=12',
          jobTitle: 'Head of Design',
          department: 'Product',
          flag: '🇸🇪',
          site: 'Stockholm',
          salary: 1350,
          startDate: 'Mar 13, 2023',
          lifecycle: 'Hired',
          status: 'Invited',
          checked: false,
          highlighted: false
        },
        {
          id: 2,
          name: 'Ksenia Bator',
          avatar: 'https://i.pravatar.cc/150?img=5',
          jobTitle: 'Fullstack Engineer',
          department: 'Engineering',
          flag: '🇺🇸',
          site: 'Miami',
          salary: 1500,
          startDate: 'Oct 13, 2023',
          lifecycle: 'Hired',
          status: 'Absent',
          checked: true,
          highlighted: true
        },
        {
          id: 3,
          name: 'Bogdan Nikitin',
          avatar: 'https://i.pravatar.cc/150?img=33',
          jobTitle: 'Mobile Lead',
          department: 'Product',
          flag: '🇺🇦',
          site: 'Kyiv',
          salary: 2600,
          startDate: 'Nov 4, 2023',
          lifecycle: 'Employed',
          status: 'Invited',
          checked: true,
          highlighted: false
        },
        {
          id: 4,
          name: 'Arsen Yatsenko',
          avatar: 'https://i.pravatar.cc/150?img=14',
          jobTitle: 'Sales Manager',
          department: 'Operations',
          flag: '🇨🇦',
          site: 'Ottawa',
          salary: 900,
          startDate: 'Sep 4, 2021',
          lifecycle: 'Employed',
          status: 'Invited',
          checked: false,
          highlighted: false
        },
        {
          id: 5,
          name: 'Daria Yurchenko',
          avatar: 'https://i.pravatar.cc/150?img=9',
          jobTitle: 'Network engineer',
          department: 'Product',
          flag: '🇧🇷',
          site: 'Sao Paulo',
          salary: 1000,
          startDate: 'Feb 21, 2023',
          lifecycle: 'Hired',
          status: 'Invited',
          checked: false,
          highlighted: false
        },
        {
          id: 6,
          name: 'Yulia Polishchuk',
          avatar: 'https://i.pravatar.cc/150?img=20',
          jobTitle: 'Head of Design',
          department: 'Product',
          flag: '🇬🇧',
          site: 'London',
          salary: 1700,
          startDate: 'Aug 2, 2024',
          lifecycle: 'Employed',
          status: 'Absent',
          checked: true,
          highlighted: false
        }
      ]
    }
  },
  methods: {
    getStatusClass(status) {
      if (status === 'Invited') {
        return 'bg-green-100 text-green-700'
      } else if (status === 'Absent') {
        return 'bg-gray-200 text-gray-700'
      }
      return 'bg-blue-100 text-blue-700'
    }
  }
}
</script>
