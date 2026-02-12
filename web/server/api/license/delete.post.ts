import { readBody } from 'h3'
import type { ApiMessageResponse } from '~/types/api'
import type { LicenseDeleteBody } from '~/types/license'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<LicenseDeleteBody>(event)
  const ids = Array.isArray(body?.ids) ? body.ids.filter(id => Number.isInteger(id)) : []

  if (!ids.length) {
    throw createError({
      statusCode: 400,
      statusMessage: 'ids is required'
    })
  }

  const response = await $fetch<ApiMessageResponse>(`${config.apiBase}/api/license/delete`, {
    method: 'POST',
    body: {
      ids
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'License delete failed'
    })
  }

  return response
})
