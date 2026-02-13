import { getHeader, readBody } from 'h3'
import type { ApiMessageResponse } from '~/types/api'
import type { UserPasswordUpdateBody } from '~/types/user'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const authorization = getHeader(event, 'authorization')
  const body = await readBody<UserPasswordUpdateBody>(event)

  const response = await $fetch<ApiMessageResponse>(`${config.apiBase}/api/user/password/update`, {
    method: 'POST',
    headers: authorization ? { Authorization: authorization } : {},
    body: {
      old_password: body?.old_password,
      new_password: body?.new_password
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 400,
      statusMessage: response.message || 'Update password failed'
    })
  }

  return response
})
