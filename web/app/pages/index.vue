<script setup lang="ts">
import type { ApiResponse } from '~/types/api'
import type { IssuerQueryData } from '~/types/issuer'
import type { ProductQueryData } from '~/types/product'
import type { CustomerQueryData } from '~/types/customer'
import type { LicenseQueryData } from '~/types/license'

interface OverviewItem {
  label: string
  icon: string
  to: string
  total: number
}

useSeoMeta({
  title: '平台总览'
})

const { data: overview, status } = await useAsyncData<OverviewItem[]>('platform-overview', async () => {
  const [issuerRes, productRes, customerRes, licenseRes] = await Promise.all([
    $fetch<ApiResponse<IssuerQueryData>>('/api/issuer/query', { method: 'POST', body: { page: 1, size: 1 } }),
    $fetch<ApiResponse<ProductQueryData>>('/api/product/query', { method: 'POST', body: { page: 1, size: 1 } }),
    $fetch<ApiResponse<CustomerQueryData>>('/api/customer/query', { method: 'POST', body: { page: 1, size: 1 } }),
    $fetch<ApiResponse<LicenseQueryData>>('/api/license/query', { method: 'POST', body: { page: 1, size: 1 } })
  ])

  return [{
    label: '机构',
    icon: 'i-lucide-building-2',
    to: '/issuer',
    total: issuerRes.data.total
  }, {
    label: '产品',
    icon: 'i-lucide-package',
    to: '/product',
    total: productRes.data.total
  }, {
    label: '客户',
    icon: 'i-lucide-user-round',
    to: '/customer',
    total: customerRes.data.total
  }, {
    label: '许可证',
    icon: 'i-lucide-key-round',
    to: '/license',
    total: licenseRes.data.total
  }]
}, {
  default: () => []
})
</script>

<template>
  <UDashboardPanel id="home">
    <template #header>
      <UDashboardNavbar title="平台总览">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <UPageCard
        title="许可证管理平台"
        description="围绕机构、产品、客户与许可证进行授权全生命周期管理。"
        variant="subtle"
      >
        <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
          <NuxtLink
            v-for="item in overview"
            :key="item.label"
            :to="item.to"
            class="rounded-lg border border-default p-4 hover:bg-elevated/50 transition-colors"
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-2 text-highlighted">
                <UIcon :name="item.icon" class="size-5" />
                <span class="font-medium">{{ item.label }}</span>
              </div>
              <UIcon name="i-lucide-arrow-right" class="size-4 text-dimmed" />
            </div>
            <p class="mt-3 text-2xl font-semibold text-highlighted">
              {{ item.total }}
            </p>
            <p class="text-xs text-muted">
              当前总数
            </p>
          </NuxtLink>
        </div>

        <template #footer>
          <div class="text-sm text-muted">
            {{ status === 'pending' ? '正在加载数据...' : '点击卡片可进入对应管理页面。' }}
          </div>
        </template>
      </UPageCard>
    </template>
  </UDashboardPanel>
</template>
