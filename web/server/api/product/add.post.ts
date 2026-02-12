import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Product, ProductAddBody } from '~/types/product'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<ProductAddBody>(event)

  const response = await $fetch<ApiResponse<Product>>(`${config.apiBase}/api/product/add`, {
    method: 'POST',
    body: {
      code: body?.code,
      name: body?.name,
      description: body?.description
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'Product add failed'
    })
  }

  return response
})
