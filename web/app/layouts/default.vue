<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'
import AppLogo from '~/components/AppLogo.vue'

const route = useRoute()
const toast = useToast()

const open = ref(false)

const links = [[{
  label: '首页',
  icon: 'i-lucide-house',
  to: '/',
  onSelect: () => {
    open.value = false
  }
}, {
  label: '机构',
  icon: 'i-lucide-building-2',
  to: '/issuer',
  onSelect: () => {
    open.value = false
  }
}, {
  label: '产品',
  icon: 'i-lucide-package',
  to: '/product',
  onSelect: () => {
    open.value = false
  }
}, {
  label: '客户',
  icon: 'i-lucide-user-round',
  to: '/customer',
  onSelect: () => {
    open.value = false
  }
}, {
  label: '许可证',
  icon: 'i-lucide-key-round',
  to: '/license',
  onSelect: () => {
    open.value = false
  }
}, {
  label: '设置',
  to: '/settings',
  icon: 'i-lucide-settings',
  defaultOpen: true,
  type: 'trigger',
  children: [{
    label: '通用',
    to: '/settings',
    exact: true,
    onSelect: () => {
      open.value = false
    }
  }, {
    label: '安全',
    to: '/settings/security',
    onSelect: () => {
      open.value = false
    }
  }]
}], [{
  label: '反馈',
  icon: 'i-lucide-message-circle',
  to: 'https://github.com/ebingbo/ilicense-lite#readme',
  target: '_blank'
}, {
  label: '帮助&支持',
  icon: 'i-lucide-info',
  to: 'https://github.com/ebingbo/ilicense-lite#readme',
  target: '_blank'
}]] satisfies NavigationMenuItem[][]

const _groups = computed(() => [{
  id: 'links',
  label: '前往',
  items: links.flat()
}, {
  id: 'code',
  label: '代码',
  items: [{
    id: 'source',
    label: '查看页面源码',
    icon: 'i-simple-icons-github',
    to: `https://github.com/ebingbo/ilicense-lite/blob/main/web/app/pages${route.path === '/' ? '/index' : route.path}.vue`,
    target: '_blank'
  }]
}])

onMounted(async () => {
  const cookie = useCookie('cookie-consent')
  if (cookie.value === 'accepted') {
    return
  }

  toast.add({
    title: '我们使用 Cookie 来提升你的使用体验。',
    duration: 0,
    close: false,
    actions: [{
      label: '接受',
      color: 'neutral',
      variant: 'outline',
      onClick: () => {
        cookie.value = 'accepted'
      }
    }, {
      label: '拒绝',
      color: 'neutral',
      variant: 'ghost'
    }]
  })
})
</script>

<template>
  <UDashboardGroup unit="rem">
    <UDashboardSidebar
      id="default"
      v-model:open="open"
      collapsible
      resizable
      class="bg-elevated/25"
      :ui="{ footer: 'lg:border-t lg:border-default' }"
    >
      <template #header="{ collapsed }">
        <!--        <TeamsMenu :collapsed="collapsed"/> -->
        <NuxtLink to="/" class="flex items-end gap-0.5">
          <AppLogo class="h-8 w-auto shrink-0" />
          <span v-if="!collapsed" class="text-xl font-bold text-highlighted">许可证平台</span>
        </NuxtLink>
      </template>

      <template #default="{ collapsed }">
        <!--        <UDashboardSearchButton :collapsed="collapsed" class="bg-transparent ring-default" /> -->

        <UNavigationMenu
          :collapsed="collapsed"
          :items="links[0]"
          orientation="vertical"
          tooltip
          popover
        />

        <UNavigationMenu
          :collapsed="collapsed"
          :items="links[1]"
          orientation="vertical"
          tooltip
          class="mt-auto"
        />
      </template>

      <template #footer="{ collapsed }">
        <UserMenu :collapsed="collapsed" />
      </template>
    </UDashboardSidebar>

    <!--    <UDashboardSearch :groups="groups" /> -->

    <slot />
  </UDashboardGroup>
</template>
