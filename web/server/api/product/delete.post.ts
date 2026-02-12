import { readBody } from 'h3'
import type { ApiMessageResponse } from '~/types/api'
import type { ProductDeleteBody } from '~/types/product'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<ProductDeleteBody>(event)
  const ids = Array.isArray(body?.ids) ? body.ids.filter(id => Number.isInteger(id)) : []

  if (!ids.length) {
    throw createError({
      statusCode: 400,
      statusMessage: 'ids is required'
    })
  }

  const response = await $fetch<ApiMessageResponse>(`${config.apiBase}/api/product/delete`, {
    method: 'POST',
    body: {
      ids
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'Product delete failed'
    })
  }

  return response
})
