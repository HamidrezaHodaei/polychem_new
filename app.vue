<template>
  <NuxtPage />
</template>

<script setup>
useHead({
  title: 'Polychemmb | Polymer & Chemical Solutions',

  meta: [
    {
      name: 'description',
      content:
        'Polychemmb is a leading supplier of polymer and chemical compounds for industrial applications.'
    },
    {
      name: 'robots',
      content: 'index, follow'
    }
  ],

  link: [
    {
      rel: 'icon',
      type: 'image/x-icon',
      href: '/favicon-v2.ico'
    },
    {
      rel: 'canonical',
      href: 'https://polychemmb.com/'
    }
  ],

  script: [
    {
      type: 'application/ld+json',
      children: JSON.stringify({
        "@context": "https://schema.org",
        "@type": "Organization",
        "name": "Polychemmb",
        "url": "https://polychemmb.com/",
        "logo": "https://polychemmb.com/logo.png"
      })
    }
  ]
})
</script>
