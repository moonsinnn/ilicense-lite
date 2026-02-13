import type { ApiResponse } from '~/types/api'
import type { UserProfile } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const response = await backendFetch<ApiResponse<UserProfile>>(event, '/api/user/profile', {
    method: 'GET'
  })

  return ensureApiSuccess(response, 'Get profile failed', 401)
})
