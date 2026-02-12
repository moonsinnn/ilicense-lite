import type { ApiMessageResponse } from '~/types/api'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const id = getRouterParam(event, 'id')

  if (!id || !/^\d+$/.test(id)) {
    throw createError({
      statusCode: 400,
      statusMessage: 'invalid issuer id'
    })
  }

  const response = await $fetch<ApiMessageResponse>(`${config.apiBase}/api/issuer/delete/${id}`, {
    method: 'POST'
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'Issuer delete failed'
    })
  }

  return response
})
