import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { CustomerQueryBody, CustomerQueryData } from '~/types/customer'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<CustomerQueryBody>(event)

  const page = Number(body?.page) > 0 ? Number(body?.page) : 1
  const size = Number(body?.size) > 0 ? Number(body?.size) : 10

  const response = await backendFetch<ApiResponse<CustomerQueryData>>(event, '/api/customer/query', {
    method: 'POST',
    body: {
      page,
      size
    }
  })

  return ensureApiSuccess(response, 'Customer query failed', 502)
})
