import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { License, LicenseRenewBody } from '~/types/license'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<LicenseRenewBody>(event)

  const response = await $fetch<ApiResponse<License>>(`${config.apiBase}/api/license/renew`, {
    method: 'POST',
    body: {
      id: body?.id,
      expire_at: body?.expire_at,
      remarks: body?.remarks
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'License renew failed'
    })
  }

  return response
})
