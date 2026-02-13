import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { License, LicenseRenewBody } from '~/types/license'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<LicenseRenewBody>(event)

  const response = await backendFetch<ApiResponse<License>>(event, '/api/license/renew', {
    method: 'POST',
    body: {
      id: body?.id,
      expire_at: body?.expire_at,
      remarks: body?.remarks
    }
  })

  return ensureApiSuccess(response, 'License renew failed', 502)
})
