import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Customer, CustomerAddBody } from '~/types/customer'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<CustomerAddBody>(event)

  const response = await $fetch<ApiResponse<Customer>>(`${config.apiBase}/api/customer/add`, {
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

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'Customer add failed'
    })
  }

  return response
})
