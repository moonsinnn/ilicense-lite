import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { LicenseQueryBody, LicenseQueryData } from '~/types/license'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<LicenseQueryBody>(event)

  const page = Number(body?.page) > 0 ? Number(body?.page) : 1
  const size = Number(body?.size) > 0 ? Number(body?.size) : 10

  const response = await $fetch<ApiResponse<LicenseQueryData>>(`${config.apiBase}/api/license/query`, {
    method: 'POST',
    body: {
      page,
      size
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'License query failed'
    })
  }

  return response
})
