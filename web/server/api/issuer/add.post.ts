import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Issuer, IssuerAddBody } from '~/types/issuer'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const body = await readBody<IssuerAddBody>(event)

  const response = await $fetch<ApiResponse<Issuer>>(`${config.apiBase}/api/issuer/add`, {
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
      statusMessage: response.message || 'Issuer add failed'
    })
  }

  return response
})
