<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import type { ApiResponse } from '~/types/api'
import type { Issuer, IssuerAddBody } from '~/types/issuer'

const schema = z.object({
  code: z.string(),
  name: z.string(),
  description: z.string()
})
const open = ref(false)
const isSubmitting = ref(false)

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({
  code: '',
  name: '',
  description: ''
})

const toast = useToast()
async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    isSubmitting.value = true
    const response = await $fetch<ApiResponse<Issuer>>('/api/issuer/add', {
      method: 'POST',
      body: event.data as IssuerAddBody
    })

    if (response.code !== 0) {
      throw new Error(response.message || '创建机构失败')
    }

    toast.add({ title: 'Success', description: `新的机构 ${event.data.name} 被添加`, color: 'success' })
    open.value = false
    state.code = ''
    state.name = ''
    state.description = ''
    await refreshNuxtData('issuer-query')
  } catch (error) {
    toast.add({
      title: '创建失败',
      description: error instanceof Error ? error.message : '请求失败，请稍后重试',
      color: 'error'
    })
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <UModal v-model:open="open" title="新建机构" description="添加新的机构到数据库">
    <UButton label="新建机构" icon="i-lucide-plus" />

    <template #body>
      <UForm
        :schema="schema"
        :state="state"
        class="space-y-4"
        @submit="onSubmit"
      >
        <UFormField label="编码" placeholder="iss-xxx" name="code">
          <UInput v-model="state.code" class="w-full" />
        </UFormField>
        <UFormField label="名称" placeholder="John Doe" name="name">
          <UInput v-model="state.name" class="w-full" />
        </UFormField>
        <UFormField label="描述" placeholder="描述是什么样的机构" name="description">
          <UTextarea v-model="state.description" class="w-full" />
        </UFormField>
        <div class="flex justify-end gap-2">
          <UButton
            label="取消"
            color="neutral"
            variant="subtle"
            @click="open = false"
          />
          <UButton
            label="创建"
            color="primary"
            variant="solid"
            type="submit"
            :loading="isSubmitting"
            :disabled="isSubmitting"
          />
        </div>
      </UForm>
    </template>
  </UModal>
</template>
