import type { ApiMessageResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const id = getRouterParam(event, 'id')

  if (!id || !/^\d+$/.test(id)) {
    throw createError({
      statusCode: 400,
      statusMessage: 'invalid product id'
    })
  }

  const response = await backendFetch<ApiMessageResponse>(event, `/api/product/delete/${id}`, {
    method: 'POST'
  })

  return ensureApiSuccess(response, 'Product delete failed', 502)
})
