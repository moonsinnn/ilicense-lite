import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Customer, CustomerAddBody } from '~/types/customer'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<CustomerAddBody>(event)

  const response = await backendFetch<ApiResponse<Customer>>(event, '/api/customer/add', {
    method: 'POST',
    body: {
      code: body?.code,
      name: body?.name,
      contact: body?.contact,
      phone: body?.phone,
      email: body?.email,
      address: body?.address
    }
  })

  return ensureApiSuccess(response, 'Customer add failed', 502)
})
