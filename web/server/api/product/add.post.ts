import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Product, ProductAddBody } from '~/types/product'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<ProductAddBody>(event)

  const response = await backendFetch<ApiResponse<Product>>(event, '/api/product/add', {
    method: 'POST',
    body: {
      code: body?.code,
      name: body?.name,
      description: body?.description
    }
  })

  return ensureApiSuccess(response, 'Product add failed', 502)
})
