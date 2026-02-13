<script setup lang="ts">
useSeoMeta({
  title: '通知设置'
})

const state = reactive<{ [key: string]: boolean }>({
  email: true,
  desktop: false,
  product_updates: true,
  weekly_digest: false,
  important_updates: true
})

const sections = [{
  title: '通知渠道',
  description: '你希望通过哪些方式接收通知？',
  fields: [{
    name: 'email',
    label: '邮件',
    description: '接收每日邮件摘要。'
  }, {
    name: 'desktop',
    label: '桌面通知',
    description: '接收桌面通知。'
  }]
}, {
  title: '账号更新',
  description: '接收系统功能与账号相关更新。',
  fields: [{
    name: 'weekly_digest',
    label: '每周摘要',
    description: '每周接收一次动态摘要。'
  }, {
    name: 'product_updates',
    label: '产品更新',
    description: '每月接收新功能与更新汇总邮件。'
  }, {
    name: 'important_updates',
    label: '重要通知',
    description: '接收安全修复、维护等重要更新邮件。'
  }]
}]

async function onChange() {
  // Do something with data
  console.log(state)
}
</script>

<template>
  <div v-for="(section, index) in sections" :key="index">
    <UPageCard
      :title="section.title"
      :description="section.description"
      variant="naked"
      class="mb-4"
    />

    <UPageCard variant="subtle" :ui="{ container: 'divide-y divide-default' }">
      <UFormField
        v-for="field in section.fields"
        :key="field.name"
        :name="field.name"
        :label="field.label"
        :description="field.description"
        class="flex items-center justify-between not-last:pb-4 gap-2"
      >
        <USwitch
          v-model="state[field.name]"
          @update:model-value="onChange"
        />
      </UFormField>
    </UPageCard>
  </div>
</template>
