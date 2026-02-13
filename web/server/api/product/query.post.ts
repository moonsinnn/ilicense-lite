import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { ProductQueryBody, ProductQueryData } from '~/types/product'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<ProductQueryBody>(event)

  const page = Number(body?.page) > 0 ? Number(body?.page) : 1
  const size = Number(body?.size) > 0 ? Number(body?.size) : 10

  const response = await backendFetch<ApiResponse<ProductQueryData>>(event, '/api/product/query', {
    method: 'POST',
    body: {
      page,
      size
    }
  })

  return ensureApiSuccess(response, 'Product query failed', 502)
})
