import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserLoginBody, UserLoginData } from '~/types/user'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<UserLoginBody>(event)

  const response = await $fetch<ApiResponse<UserLoginData>>(`${config.apiBase}/api/user/login`, {
    method: 'POST',
    body: {
      username: body?.username,
      password: body?.password
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 401,
      statusMessage: response.message || 'Login failed'
    })
  }

  return response
})
