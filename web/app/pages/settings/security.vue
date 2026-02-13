<script setup lang="ts">
import * as z from 'zod'
import type { FormError, FormSubmitEvent } from '@nuxt/ui'

useSeoMeta({
  title: '安全设置'
})

const { updatePassword } = useAuth()
const loading = ref(false)
const toast = useToast()

const passwordSchema = z.object({
  current: z.string().min(8, '至少 8 位字符'),
  new: z.string().min(8, '至少 8 位字符')
})

type PasswordSchema = z.output<typeof passwordSchema>

const password = reactive<Partial<PasswordSchema>>({
  current: '',
  new: ''
})

const validate = (state: Partial<PasswordSchema>): FormError[] => {
  const errors: FormError[] = []
  if (state.current && state.new && state.current === state.new) {
    errors.push({ name: 'new', message: '新旧密码不能相同' })
  }
  return errors
}

async function onSubmit(event: FormSubmitEvent<PasswordSchema>) {
  try {
    loading.value = true
    await updatePassword({
      old_password: event.data.current,
      new_password: event.data.new
    })
    password.current = ''
    password.new = ''
    toast.add({
      title: '更新成功',
      description: '密码已更新',
      color: 'success'
    })
  } catch (error) {
    toast.add({
      title: '更新失败',
      description: error instanceof Error ? error.message : '请求失败，请稍后重试',
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <UPageCard
    title="密码设置"
    description="设置新密码前请先确认当前密码。"
    variant="subtle"
  >
    <UForm
      :schema="passwordSchema"
      :state="password"
      :validate="validate"
      class="flex flex-col gap-4 max-w-xs"
      @submit="onSubmit"
    >
      <UFormField name="current">
        <UInput
          v-model="password.current"
          type="password"
          placeholder="当前密码"
          class="w-full"
        />
      </UFormField>

      <UFormField name="new">
        <UInput
          v-model="password.new"
          type="password"
          placeholder="新密码"
          class="w-full"
        />
      </UFormField>

      <UButton
        label="更新密码"
        class="w-fit"
        type="submit"
        :loading="loading"
        :disabled="loading"
      />
    </UForm>
  </UPageCard>

  <UPageCard
    title="账号管理"
    description="如果你不再使用本服务，可在此删除账号。该操作不可恢复，所有相关数据将被永久删除。"
    class="bg-gradient-to-tl from-error/10 from-5% to-default"
  >
    <template #footer>
      <UButton label="删除账号" color="error" />
    </template>
  </UPageCard>
</template>
