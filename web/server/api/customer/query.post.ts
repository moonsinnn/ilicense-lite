import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { CustomerQueryBody, CustomerQueryData } from '~/types/customer'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<CustomerQueryBody>(event)

  const page = Number(body?.page) > 0 ? Number(body?.page) : 1
  const size = Number(body?.size) > 0 ? Number(body?.size) : 10

  const response = await $fetch<ApiResponse<CustomerQueryData>>(`${config.apiBase}/api/customer/query`, {
    method: 'POST',
    body: {
      page,
      size
    }
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'Customer query failed'
    })
  }

  return response
})
