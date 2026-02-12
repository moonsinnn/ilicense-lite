import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { ProductQueryBody, ProductQueryData } from '~/types/product'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<ProductQueryBody>(event)

  const page = Number(body?.page) > 0 ? Number(body?.page) : 1
  const size = Number(body?.size) > 0 ? Number(body?.size) : 10

  const response = await $fetch<ApiResponse<ProductQueryData>>(`${config.apiBase}/api/product/query`, {
    method: 'POST',
    body: {
      page,
      size
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'Product query failed'
    })
  }

  return response
})
