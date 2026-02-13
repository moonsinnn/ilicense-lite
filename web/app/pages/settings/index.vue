<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'

useSeoMeta({
  title: '通用设置'
})

const fileRef = ref<HTMLInputElement>()
const { user, fetchProfile, updateProfile } = useAuth()
const loading = ref(false)

const profileSchema = z.object({
  name: z.string().min(2, '至少 2 个字符'),
  email: z.string().email('邮箱格式不正确'),
  username: z.string().min(2, '至少 2 个字符'),
  avatar: z.string().optional()
})

type ProfileSchema = z.output<typeof profileSchema>

const profile = reactive<Partial<ProfileSchema>>({
  name: '',
  email: '',
  username: '',
  avatar: undefined
})
const toast = useToast()

onMounted(async () => {
  try {
    const data = await fetchProfile()
    if (data) {
      profile.name = data.name
      profile.email = data.email
      profile.username = data.username
      profile.avatar = data.avatar || undefined
    }
  } catch {
    toast.add({
      title: '获取资料失败',
      description: '请重新登录后重试',
      color: 'error'
    })
  }
})

async function onSubmit(event: FormSubmitEvent<ProfileSchema>) {
  try {
    loading.value = true
    await updateProfile({
      name: event.data.name,
      email: event.data.email,
      avatar: profile.avatar
    })
    if (user.value) {
      profile.username = user.value.username
    }
    toast.add({
      title: '保存成功',
      description: '设置已更新。',
      icon: 'i-lucide-check',
      color: 'success'
    })
  } catch (error) {
    toast.add({
      title: '保存失败',
      description: error instanceof Error ? error.message : '请求失败，请稍后重试',
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}

function onFileChange(e: Event) {
  const input = e.target as HTMLInputElement

  if (!input.files?.length) {
    return
  }

  profile.avatar = URL.createObjectURL(input.files[0]!)
}

function onFileClick() {
  fileRef.value?.click()
}
</script>

<template>
  <UForm
    id="settings"
    :schema="profileSchema"
    :state="profile"
    @submit="onSubmit"
  >
    <UPageCard
      title="个人资料"
      description="以下信息将公开展示。"
      variant="naked"
      orientation="horizontal"
      class="mb-4"
    >
      <UButton
        form="settings"
        label="保存更改"
        color="neutral"
        type="submit"
        class="w-fit lg:ms-auto"
        :loading="loading"
        :disabled="loading"
      />
    </UPageCard>

    <UPageCard variant="subtle">
      <UFormField
        name="name"
        label="姓名"
        description="将显示在收据、发票和其他通知中。"
        required
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInput
          v-model="profile.name"
          autocomplete="off"
        />
      </UFormField>
      <USeparator />
      <UFormField
        name="email"
        label="邮箱"
        description="用于登录、收据邮件和产品更新通知。"
        required
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInput
          v-model="profile.email"
          type="email"
          autocomplete="off"
        />
      </UFormField>
      <USeparator />
      <UFormField
        name="username"
        label="用户名"
        description="用于登录和个人主页链接的唯一用户名。"
        required
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInput
          v-model="profile.username"
          type="username"
          autocomplete="off"
        />
      </UFormField>
      <USeparator />
      <UFormField
        name="avatar"
        label="头像"
        description="支持 JPG、GIF、PNG，最大 1MB。"
        class="flex max-sm:flex-col justify-between sm:items-center gap-4"
      >
        <div class="flex flex-wrap items-center gap-3">
          <UAvatar
            :src="profile.avatar"
            :alt="profile.name"
            size="lg"
          />
          <UButton
            label="选择文件"
            color="neutral"
            @click="onFileClick"
          />
          <input
            ref="fileRef"
            type="file"
            class="hidden"
            accept=".jpg, .jpeg, .png, .gif"
            @change="onFileChange"
          >
        </div>
      </UFormField>
    </UPageCard>
  </UForm>
</template>
