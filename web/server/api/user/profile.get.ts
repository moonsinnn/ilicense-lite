import { getHeader } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserProfile } from '~/types/user'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const authorization = getHeader(event, 'authorization')

  const response = await $fetch<ApiResponse<UserProfile>>(`${config.apiBase}/api/user/profile`, {
    method: 'GET',
    headers: authorization ? { Authorization: authorization } : {}
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 401,
      statusMessage: response.message || 'Get profile failed'
    })
  }

  return response
})
