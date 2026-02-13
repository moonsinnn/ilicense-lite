import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { LicenseActivateBody, LicenseActivateData } from '~/types/license'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<LicenseActivateBody>(event)

  const response = await backendFetch<ApiResponse<LicenseActivateData>>(event, '/api/license/activate', {
    method: 'POST',
    body: {
      issuer_id: body?.issuer_id,
      code: body?.code
    }
  })

  return ensureApiSuccess(response, 'License activate failed', 502)
})
