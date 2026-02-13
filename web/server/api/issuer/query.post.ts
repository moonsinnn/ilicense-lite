import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { IssuerQueryBody, IssuerQueryData } from '~/types/issuer'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<IssuerQueryBody>(event)

  const page = Number(body?.page) > 0 ? Number(body?.page) : 1
  const size = Number(body?.size) > 0 ? Number(body?.size) : 10

  const response = await backendFetch<ApiResponse<IssuerQueryData>>(event, '/api/issuer/query', {
    method: 'POST',
    body: {
      page,
      size
    }
  })

  return ensureApiSuccess(response, 'Issuer query failed', 502)
})
