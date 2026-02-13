import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserLoginBody, UserLoginData } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserLoginBody>(event)

  const response = await backendFetch<ApiResponse<UserLoginData>>(event, '/api/user/login', {
    method: 'POST',
    body: {
      username: body?.username,
      password: body?.password
    }
  })

  return ensureApiSuccess(response, 'Login failed', 401)
})
