<script setup lang="ts">
import * as z from 'zod'
import dayjs from 'dayjs'
import type { FormSubmitEvent } from '@nuxt/ui'
import type { ApiResponse } from '~/types/api'
import type { License, LicenseRenewBody } from '~/types/license'

const props = defineProps<{
  open: boolean
  license?: License | null
}>()

const emit = defineEmits<{
  'update:open': [boolean]
  'success': []
}>()

const schema = z.object({
  expire_at: z.string().min(1, '请选择过期时间'),
  remarks: z.string().optional()
})

type Schema = z.output<typeof schema>

const isSubmitting = ref(false)
const toast = useToast()
const state = reactive<Partial<Schema>>({
  expire_at: '',
  remarks: ''
})

watch(() => props.open, (isOpen) => {
  if (!isOpen) return
  state.expire_at = props.license?.expire_at
    ? dayjs(props.license.expire_at).format('YYYY-MM-DDTHH:mm:ss')
    : ''
  state.remarks = props.license?.remarks || ''
})

function closeModal() {
  emit('update:open', false)
}

async function onSubmit(event: FormSubmitEvent<Schema>) {
  if (!props.license?.id) return

  try {
    isSubmitting.value = true
    const body: LicenseRenewBody = {
      id: props.license.id,
      expire_at: dayjs(event.data.expire_at).format('YYYY-MM-DD HH:mm:ss'),
      remarks: event.data.remarks || ''
    }

    const response = await $fetch<ApiResponse<License>>('/api/license/renew', {
      method: 'POST',
      body
    })

    if (response.code !== 0) {
      throw new Error(response.message || '续期失败')
    }

    toast.add({
      title: '续期成功',
      description: `许可证 ${response.data.code} 已续期`,
      color: 'success'
    })
    emit('success')
    closeModal()
  } catch (error) {
    toast.add({
      title: '续期失败',
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
    :open="open"
    :title="`续期许可证 ${license?.code || ''}`"
    description="请设置新的过期时间和备注信息"
    @update:open="(value) => emit('update:open', value)"
  >
    <template #body>
      <UForm
        :schema="schema"
        :state="state"
        class="space-y-4"
        @submit="onSubmit"
      >
        <UFormField label="过期时间" name="expire_at">
          <UInput v-model="state.expire_at" type="datetime-local" class="w-full" />
        </UFormField>
        <UFormField label="备注" name="remarks">
          <UTextarea v-model="state.remarks" class="w-full" />
        </UFormField>
        <div class="flex justify-end gap-2">
          <UButton
            label="取消"
            color="neutral"
            variant="subtle"
            @click="closeModal"
          />
          <UButton
            label="续期"
            color="primary"
            variant="solid"
            type="submit"
            :loading="isSubmitting"
            :disabled="isSubmitting || !license"
          />
        </div>
      </UForm>
    </template>
  </UModal>
</template>
