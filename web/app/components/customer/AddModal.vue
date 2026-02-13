<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import type { ApiResponse } from '~/types/api'
import type { Customer, CustomerAddBody } from '~/types/customer'

const schema = z.object({
  code: z.string().min(1, '编码不能为空'),
  name: z.string().min(1, '名称不能为空'),
  contact: z.string().min(1, '联系人不能为空'),
  phone: z
    .string()
    .min(1, '手机号不能为空')
    .regex(/^1[3-9]\d{9}$/, '手机号格式不正确'),
  email: z.email('邮箱格式不正确'),
  address: z.string().optional()
})
const open = ref(false)
const isSubmitting = ref(false)

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({
  code: '',
  name: '',
  contact: '',
  phone: '',
  email: '',
  address: ''
})

const toast = useToast()

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    isSubmitting.value = true
    const response = await $fetch<ApiResponse<Customer>>('/api/customer/add', {
      method: 'POST',
      body: event.data as CustomerAddBody
    })

    if (response.code !== 0) {
      throw new Error(response.message || '创建客户失败')
    }

    toast.add({ title: '成功', description: `新的客户 ${event.data.name} 已添加`, color: 'success' })
    open.value = false
    state.code = ''
    state.name = ''
    state.contact = ''
    state.phone = ''
    state.email = ''
    state.address = ''
    await refreshNuxtData('customer-query')
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
  <UModal v-model:open="open" title="新建客户" description="添加新的客户到数据库">
    <UButton label="新建客户" icon="i-lucide-plus" />

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
        <UFormField label="名称" placeholder="请输入客户名称" name="name">
          <UInput v-model="state.name" class="w-full" />
        </UFormField>
        <UFormField label="联系人" placeholder="请输入联系人" name="contact">
          <UInput v-model="state.contact" class="w-full" />
        </UFormField>
        <UFormField label="联系电话" placeholder="请输入联系电话" name="phone">
          <UInput v-model="state.phone" class="w-full" />
        </UFormField>
        <UFormField label="联系邮箱" placeholder="请输入联系邮箱" name="email">
          <UInput v-model="state.email" class="w-full" />
        </UFormField>
        <UFormField label="联系地址" placeholder="请输入联系地址" name="address">
          <UInput v-model="state.address" class="w-full" />
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
