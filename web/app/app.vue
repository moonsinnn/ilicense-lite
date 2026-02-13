<script setup lang="ts">
const colorMode = useColorMode()

const color = computed(() => colorMode.value === 'dark' ? '#1b1718' : 'white')

useHead({
  titleTemplate: (titleChunk?: string) => {
    return titleChunk ? `${titleChunk} | iLicense Lite` : 'iLicense Lite 许可证管理平台'
  },
  meta: [
    { charset: 'utf-8' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    { key: 'theme-color', name: 'theme-color', content: color }
  ],
  link: [
    { rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' }
  ],
  htmlAttrs: {
    lang: 'zh-CN'
  }
})

const title = 'iLicense Lite 许可证管理平台'
const description = '面向机构、产品、客户与许可证全流程管理的轻量级平台，支持发行、续期、激活与状态管理。'

useSeoMeta({
  title,
  description,
  ogTitle: title,
  ogDescription: description,
  ogImage: '/favicon.svg',
  twitterImage: '/favicon.svg',
  twitterCard: 'summary'
})
</script>

<template>
  <UApp>
    <NuxtLoadingIndicator />

    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </UApp>
</template>
