import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { License, LicenseAddBody } from '~/types/license'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<LicenseAddBody>(event)

  const response = await $fetch<ApiResponse<License>>(`${config.apiBase}/api/license/add`, {
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

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'License add failed'
    })
  }

  return response
})
