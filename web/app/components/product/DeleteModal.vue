<script setup lang="ts">
import type { ApiMessageResponse } from '~/types/api'
import type { ProductDeleteBody } from '~/types/product'

const props = withDefaults(defineProps<{
  count?: number
  ids?: number[]
}>(), {
  count: 0,
  ids: () => []
})

const open = ref(false)
const isSubmitting = ref(false)
const toast = useToast()

async function onSubmit() {
  if (!open.value || isSubmitting.value || !props.ids.length) return

  try {
    isSubmitting.value = true
    const response = await $fetch<ApiMessageResponse>('/api/product/delete', {
      method: 'POST',
      body: { ids: props.ids } as ProductDeleteBody
    })

    if (response.code !== 0) {
      throw new Error(response.message || '删除失败')
    }

    toast.add({ title: '删除成功', description: `已删除 ${props.ids.length} 条产品`, color: 'success' })

    open.value = false
    await refreshNuxtData('product-query')
  } catch (error) {
    toast.add({
      title: '删除失败',
      description: error instanceof Error ? error.message : '请求失败，请稍后重试',
      color: 'error'
    })
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <UModal
    v-model:open="open"
    :title="`删除${props.count}个产品`"
    :description="`你确定吗, 该项目操作不可恢复.`"
  >
    <slot />

    <template #body>
      <div class="flex justify-end gap-2">
        <UButton
          label="取消"
          color="neutral"
          variant="subtle"
          @click="open = false"
        />
        <UButton
          label="删除"
          color="error"
          variant="solid"
          :loading="isSubmitting"
          :disabled="isSubmitting || !props.count"
          @click="onSubmit"
        />
      </div>
    </template>
  </UModal>
</template>
