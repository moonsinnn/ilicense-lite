<script setup lang="ts">
import * as z from 'zod'
import dayjs from 'dayjs'
import type { FormSubmitEvent } from '@nuxt/ui'
import type { ApiResponse } from '~/types/api'
import type { License, LicenseAddBody } from '~/types/license'
import type { Product, ProductQueryData } from '~/types/product'
import type { Customer, CustomerQueryData } from '~/types/customer'
import type { Issuer, IssuerQueryData } from '~/types/issuer'

const schema = z.object({
  code: z.string().optional(),
  product_id: z.number().int().positive('请选择产品'),
  customer_id: z.number().int().positive('请选择客户'),
  issuer_id: z.number().int().positive('请选择机构'),
  expire_at: z.string(),
  modules: z.string().optional(),
  max_instances: z.number().optional(),
  remarks: z.string().optional()
})
const open = ref(false)
const isSubmitting = ref(false)
const isOptionsLoading = ref(false)
const optionsLoaded = ref(false)

type Schema = z.output<typeof schema>
type OptionItem = { label: string, value: number }

const state = reactive<Partial<Schema>>({
  code: '',
  product_id: undefined,
  customer_id: undefined,
  issuer_id: undefined,
  expire_at: '',
  modules: '',
  max_instances: undefined,
  remarks: ''
})

const toast = useToast()
const issuerOptions = ref<OptionItem[]>([])
const productOptions = ref<OptionItem[]>([])
const customerOptions = ref<OptionItem[]>([])

async function loadOptions() {
  if (isOptionsLoading.value) return

  try {
    isOptionsLoading.value = true
    const [issuerRes, productRes, customerRes] = await Promise.all([
      $fetch<ApiResponse<IssuerQueryData>>('/api/issuer/query', {
        method: 'POST',
        body: { page: 1, size: 1000 }
      }),
      $fetch<ApiResponse<ProductQueryData>>('/api/product/query', {
        method: 'POST',
        body: { page: 1, size: 1000 }
      }),
      $fetch<ApiResponse<CustomerQueryData>>('/api/customer/query', {
        method: 'POST',
        body: { page: 1, size: 1000 }
      })
    ])

    issuerOptions.value = issuerRes.data.items.map((item: Issuer) => ({
      label: `${item.name} (${item.code})`,
      value: item.id
    }))
    productOptions.value = productRes.data.items.map((item: Product) => ({
      label: `${item.name} (${item.code})`,
      value: item.id
    }))
    customerOptions.value = customerRes.data.items.map((item: Customer) => ({
      label: `${item.name} (${item.code})`,
      value: item.id
    }))
    optionsLoaded.value = true
  } catch {
    toast.add({
      title: '列表加载失败',
      description: '产品/客户/机构列表加载失败，请稍后重试',
      color: 'error'
    })
  } finally {
    isOptionsLoading.value = false
  }
}

watch(open, async (value) => {
  if (value && !optionsLoaded.value) {
    await loadOptions()
  }
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  if (!event.data.expire_at) {
    toast.add({ title: '创建失败', description: '请选择过期日期', color: 'error' })
    return
  }

  try {
    isSubmitting.value = true
    const body: LicenseAddBody = {
      code: event.data.code || '',
      product_id: event.data.product_id,
      customer_id: event.data.customer_id,
      issuer_id: event.data.issuer_id,
      expire_at: dayjs(event.data.expire_at).format('YYYY-MM-DD HH:mm:ss'),
      modules: event.data.modules || '',
      max_instances: event.data.max_instances || 0,
      remarks: event.data.remarks || ''
    }

    const response = await $fetch<ApiResponse<License>>('/api/license/add', {
      method: 'POST',
      body
    })

    if (response.code !== 0) {
      throw new Error(response.message || '创建许可证失败')
    }

    toast.add({ title: 'Success', description: `新的许可证 ${response.data.code} 被添加`, color: 'success' })
    open.value = false
    state.code = ''
    state.product_id = undefined
    state.customer_id = undefined
    state.issuer_id = undefined
    state.expire_at = ''
    state.modules = ''
    state.max_instances = undefined
    state.remarks = ''
    await refreshNuxtData('license-query')
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
  <UModal v-model:open="open" title="新建许可证" description="添加新的许可证到数据库">
    <UButton label="新建许可证" icon="i-lucide-plus" />

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
        <UFormField label="机构" name="issuer_id">
          <USelect
            v-model="state.issuer_id"
            :items="issuerOptions"
            class="w-full"
            placeholder="请选择机构"
            :loading="isOptionsLoading"
            :disabled="isOptionsLoading || !issuerOptions.length"
          />
        </UFormField>
        <UFormField label="产品" name="product_id">
          <USelect
            v-model="state.product_id"
            :items="productOptions"
            class="w-full"
            placeholder="请选择产品"
            :loading="isOptionsLoading"
            :disabled="isOptionsLoading || !productOptions.length"
          />
        </UFormField>
        <UFormField label="客户" name="customer_id">
          <USelect
            v-model="state.customer_id"
            :items="customerOptions"
            class="w-full"
            placeholder="请选择客户"
            :loading="isOptionsLoading"
            :disabled="isOptionsLoading || !customerOptions.length"
          />
        </UFormField>
        <UFormField label="过期日期" name="expire_at">
          <UInput v-model="state.expire_at" type="datetime-local" class="w-full" />
        </UFormField>
        <UFormField label="模块" name="modules">
          <UInput v-model="state.modules" class="w-full" />
        </UFormField>
        <UFormField label="备注" placeholder="添加备注" name="remarks">
          <UTextarea v-model="state.remarks" class="w-full" />
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
