import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { License, LicenseAddBody } from '~/types/license'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<LicenseAddBody>(event)

  const response = await backendFetch<ApiResponse<License>>(event, '/api/license/add', {
    method: 'POST',
    body: {
      code: body?.code,
      product_id: body.product_id,
      customer_id: body.customer_id,
      issuer_id: body.issuer_id,
      expire_at: body.expire_at,
      modules: body?.modules,
      max_instances: body?.max_instances,
      remarks: body?.remarks
    }
  })

  return ensureApiSuccess(response, 'License add failed', 502)
})
