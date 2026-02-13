<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'

const schema = z.object({
  username: z.string().min(1, '请输入用户名'),
  password: z.string().min(1, '请输入密码')
})

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({
  username: '',
  password: ''
})
const loading = ref(false)
const toast = useToast()
const { token, login } = useAuth()

useSeoMeta({
  title: '登录'
})

if (token.value) {
  await navigateTo('/')
}

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    loading.value = true
    await login(event.data.username, event.data.password)
    toast.add({
      title: '登录成功',
      description: '欢迎回来',
      color: 'success'
    })
    await navigateTo('/')
  } catch (error) {
    toast.add({
      title: '登录失败',
      description: error instanceof Error ? error.message : '用户名或密码错误',
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-elevated/30 px-4 w-full">
    <UCard class="w-full max-w-md">
      <template #header>
        <h1 class="text-xl font-semibold text-highlighted">
          登录许可证平台
        </h1>
      </template>

      <UForm
        :schema="schema"
        :state="state"
        class="space-y-4"
        @submit="onSubmit"
      >
        <UFormField label="用户名" name="username">
          <UInput v-model="state.username" class="w-full" placeholder="请输入用户名" />
        </UFormField>
        <UFormField label="密码" name="password">
          <UInput
            v-model="state.password"
            type="password"
            class="w-full"
            placeholder="请输入密码"
          />
        </UFormField>
        <UButton
          type="submit"
          label="登录"
          color="primary"
          class="w-full justify-center"
          :loading="loading"
          :disabled="loading"
        />
      </UForm>
    </UCard>
  </div>
</template>
