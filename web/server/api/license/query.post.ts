import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { LicenseQueryBody, LicenseQueryData } from '~/types/license'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<LicenseQueryBody>(event)

  const page = Number(body?.page) > 0 ? Number(body?.page) : 1
  const size = Number(body?.size) > 0 ? Number(body?.size) : 10

  const response = await backendFetch<ApiResponse<LicenseQueryData>>(event, '/api/license/query', {
    method: 'POST',
    body: {
      page,
      size
    }
  })

  return ensureApiSuccess(response, 'License query failed', 502)
})
