import { readBody } from 'h3'
import type { ApiMessageResponse } from '~/types/api'
import type { LicenseDeleteBody } from '~/types/license'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<LicenseDeleteBody>(event)
  const ids = Array.isArray(body?.ids) ? body.ids.filter(id => Number.isInteger(id)) : []

  if (!ids.length) {
    throw createError({
      statusCode: 400,
      statusMessage: 'ids is required'
    })
  }

  const response = await backendFetch<ApiMessageResponse>(event, '/api/license/delete', {
    method: 'POST',
    body: {
      ids
    }
  })

  return ensureApiSuccess(response, 'License delete failed', 502)
})
