<template>
  <NuxtPage />
</template>

<script setup>
const siteUrl = 'https://polychemmb.com'
const siteName = 'Polychemmb | Advanced Polymer Compounds & Masterbatch Solutions'
const siteDescription = 'POLYCHEM is a leading manufacturer of advanced engineering polymers, compounds, and masterbatch solutions in Aras Free Zone, Iran. We specialize in color masterbatch, filler masterbatch, additive masterbatch, polymer blends, and custom compounds for industrial applications since 2015.'
const siteLogo = `${siteUrl}/logo.png`
const ogImage = `${siteUrl}/og-image.jpg`

useHead({
  htmlAttrs: {
    lang: 'en',
    dir: 'ltr'
  },
  
  title: siteName,

  meta: [
    // Basic Meta Tags
    { name: 'description', content: siteDescription },
    { name: 'keywords', content: 'polymer compounds, masterbatch, engineering polymers, color masterbatch, filler masterbatch, additive masterbatch, polymer blends, polypropylene talc compound, anti-blocking masterbatch, flame retardant masterbatch, white masterbatch, calcium carbonate masterbatch, Aras Free Zone, Iran polymer industry, petrochemical compounds' },
    { name: 'author', content: 'Polychem (Polychemmb)' },
    { name: 'robots', content: 'index, follow, max-image-preview:large, max-snippet:-1, max-video-preview:-1' },
    { name: 'googlebot', content: 'index, follow' },
    { name: 'googlebot-image', content: 'index, follow' },
    
    // Google Image Search
    { itemprop: 'name', content: siteName },
    { itemprop: 'description', content: siteDescription },
    { itemprop: 'image', content: siteLogo },
    
    // Open Graph (Facebook, LinkedIn, WhatsApp, Telegram)
    { property: 'og:type', content: 'website' },
    { property: 'og:url', content: siteUrl },
    { property: 'og:title', content: siteName },
    { property: 'og:description', content: siteDescription },
    { property: 'og:image', content: ogImage },
    { property: 'og:image:secure_url', content: ogImage },
    { property: 'og:image:type', content: 'image/jpeg' },
    { property: 'og:image:width', content: '1200' },
    { property: 'og:image:height', content: '630' },
    { property: 'og:image:alt', content: 'Polychemmb - Advanced Polymer Compounds & Masterbatch Solutions' },
    { property: 'og:site_name', content: 'Polychemmb' },
    { property: 'og:locale', content: 'en_US' },
    { property: 'og:locale:alternate', content: 'fa_IR' },
    
    // Twitter Card (Twitter/X)
    { name: 'twitter:card', content: 'summary_large_image' },
    { name: 'twitter:site', content: '@polychemmb' },
    { name: 'twitter:creator', content: '@polychemmb' },
    { name: 'twitter:url', content: siteUrl },
    { name: 'twitter:title', content: siteName },
    { name: 'twitter:description', content: siteDescription },
    { name: 'twitter:image', content: ogImage },
    { name: 'twitter:image:alt', content: 'Polychemmb - Leading Polymer & Masterbatch Manufacturer' },
    
    // LinkedIn Specific
    { property: 'article:publisher', content: 'https://www.linkedin.com/company/polychemmb' },
    { property: 'article:author', content: 'Polychem Corporation' },
    
    // Additional SEO Meta Tags
    { name: 'geo.region', content: 'IR-03' },
    { name: 'geo.placename', content: 'Aras Free Zone, Jolfa' },
    { name: 'geo.position', content: '38.930278;45.633889' },
    { name: 'ICBM', content: '38.930278, 45.633889' }
  ],

  link: [
    { rel: 'icon', type: 'image/x-icon', href: '/favicon-v2.ico' },
    { rel: 'apple-touch-icon', sizes: '180x180', href: '/logo.png' },
    { rel: 'icon', type: 'image/png', sizes: '32x32', href: '/logo.png' },
    { rel: 'icon', type: 'image/png', sizes: '16x16', href: '/logo.png' },
    { rel: 'canonical', href: siteUrl },
    { rel: 'alternate', hreflang: 'en', href: `${siteUrl}/en/` },
    { rel: 'alternate', hreflang: 'fa', href: `${siteUrl}/fa/` },
    { rel: 'alternate', hreflang: 'x-default', href: siteUrl }
  ],

  script: [
    {
      type: 'application/ld+json',
      children: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'Organization',
        '@id': `${siteUrl}/#organization`,
        name: 'Polychem',
        legalName: 'Polychem Corporation (Polychemmb)',
        alternateName: ['Polychemmb', 'Polychem MB', 'پلی‌کم', 'پلیمر شیمی ارس'],
        url: siteUrl,
        logo: {
          '@type': 'ImageObject',
          '@id': `${siteUrl}/#logo`,
          url: siteLogo,
          contentUrl: siteLogo,
          caption: 'Polychemmb Logo',
          inLanguage: 'en',
          width: 250,
          height: 250
        },
        image: {
          '@type': 'ImageObject',
          '@id': `${siteUrl}/#primaryImage`,
          url: siteLogo,
          contentUrl: siteLogo,
          caption: 'Polychemmb - Polymer and Chemical Solutions',
          inLanguage: 'en'
        },
        description: 'Manufacturer of advanced engineering polymers, compounds and masterbatch solutions for industrial applications',
        foundingDate: '2015',
        slogan: 'Polymer & Chemical Solutions',
        
        // Contact Information
        telephone: '+98-921-22898980',
        fax: '+98-921-22898980',
        email: 'info@polychemmb.com',
        
        // Address
        address: {
          '@type': 'PostalAddress',
          streetAddress: 'Aras Free Trade-Industrial Zone',
          addressLocality: 'Jolfa',
          addressRegion: 'East Azerbaijan Province',
          postalCode: '5481',
          addressCountry: 'IR'
        },
        
        // Location
        geo: {
          '@type': 'GeoCoordinates',
          latitude: '38.930278',
          longitude: '45.633889'
        },
        
        // Area Served
        areaServed: [
          {
            '@type': 'Country',
            name: 'Iran'
          },
          {
            '@type': 'Place',
            name: 'Middle East'
          },
          {
            '@type': 'Place',
            name: 'Caucasus Region'
          }
        ],
        
        // Products & Services
        hasOfferCatalog: {
          '@type': 'OfferCatalog',
          name: 'Polymer Products',
          itemListElement: [
            {
              '@type': 'Offer',
              itemOffered: {
                '@type': 'Product',
                name: 'Color Masterbatch',
                description: 'Concentrated mixture of pigments and polymers for coloring plastics'
              }
            },
            {
              '@type': 'Offer',
              itemOffered: {
                '@type': 'Product',
                name: 'Filler Masterbatch',
                description: 'CaCO3-based masterbatch for cost reduction and property enhancement'
              }
            },
            {
              '@type': 'Offer',
              itemOffered: {
                '@type': 'Product',
                name: 'Additive Masterbatch',
                description: 'Specialized additives including anti-blocking, flame retardant, and UV protection'
              }
            },
            {
              '@type': 'Offer',
              itemOffered: {
                '@type': 'Product',
                name: 'Polymer Compounds',
                description: 'Custom formulated polymer compounds including PP-Talc compounds'
              }
            },
            {
              '@type': 'Offer',
              itemOffered: {
                '@type': 'Product',
                name: 'Polymer Blends',
                description: 'Advanced engineering polymer blends for specialized applications'
              }
            }
          ]
        },
        
        // Contact Points
        contactPoint: [
          {
            '@type': 'ContactPoint',
            telephone: '+98-921-22898980',
            contactType: 'Sales',
            availableLanguage: ['English', 'Persian', 'Turkish'],
            areaServed: 'IR'
          },
          {
            '@type': 'ContactPoint',
            telephone: '+98-921-22898980',
            contactType: 'Technical Support',
            availableLanguage: ['English', 'Persian'],
            areaServed: 'IR'
          }
        ],
        
        // Social Media & External Links
        sameAs: [
          'https://www.linkedin.com/company/polychemmb',
          'https://twitter.com/polychemmb',
          siteUrl
        ]
      })
    },
    {
      type: 'application/ld+json',
      children: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'WebSite',
        '@id': `${siteUrl}/#website`,
        url: siteUrl,
        name: 'Polychemmb',
        description: siteDescription,
        publisher: {
          '@id': `${siteUrl}/#organization`
        },
        inLanguage: ['en', 'fa'],
        potentialAction: {
          '@type': 'SearchAction',
          target: {
            '@type': 'EntryPoint',
            urlTemplate: `${siteUrl}/search?q={search_term_string}`
          },
          'query-input': 'required name=search_term_string'
        }
      })
    },
    {
      type: 'application/ld+json',
      children: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'BreadcrumbList',
        itemListElement: [
          {
            '@type': 'ListItem',
            position: 1,
            name: 'Home',
            item: siteUrl
          }
        ]
      })
    }
  ]
})
</script>