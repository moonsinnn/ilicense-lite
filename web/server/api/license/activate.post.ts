import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { LicenseActivateBody, LicenseActivateData } from '~/types/license'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<LicenseActivateBody>(event)

  const response = await $fetch<ApiResponse<LicenseActivateData>>(`${config.apiBase}/api/license/activate`, {
    method: 'POST',
    body: {
      issuer_id: body?.issuer_id,
      code: body?.code
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'License activate failed'
    })
  }

  return response
})
