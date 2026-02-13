import { readBody } from 'h3'
import type { ApiMessageResponse } from '~/types/api'
import type { UserPasswordUpdateBody } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserPasswordUpdateBody>(event)

  const response = await backendFetch<ApiMessageResponse>(event, '/api/user/password/update', {
    method: 'POST',
    body: {
      old_password: body?.old_password,
      new_password: body?.new_password
    }
  })

  return ensureApiSuccess(response, 'Update password failed', 400)
})
