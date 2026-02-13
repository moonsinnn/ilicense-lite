import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Issuer, IssuerAddBody } from '~/types/issuer'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<IssuerAddBody>(event)

  const response = await backendFetch<ApiResponse<Issuer>>(event, '/api/issuer/add', {
    method: 'POST',
    body: {
      code: body?.code,
      name: body?.name,
      description: body?.description
    }
  })

  return ensureApiSuccess(response, 'Issuer add failed', 502)
})
