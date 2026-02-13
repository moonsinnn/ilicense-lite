import { getHeader, readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserProfile, UserProfileUpdateBody } from '~/types/user'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const authorization = getHeader(event, 'authorization')
  const body = await readBody<UserProfileUpdateBody>(event)

  const response = await $fetch<ApiResponse<UserProfile>>(`${config.apiBase}/api/user/profile/update`, {
    method: 'POST',
    headers: authorization ? { Authorization: authorization } : {},
    body: {
      name: body?.name,
      email: body?.email,
      avatar: body?.avatar
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 400,
      statusMessage: response.message || 'Update profile failed'
    })
  }

  return response
})
