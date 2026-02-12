import type { ApiMessageResponse } from '~/types/api'

export default eventHandler(async (event) => {
  const config = useRuntimeConfig()
  const id = getRouterParam(event, 'id')

  if (!id || !/^\d+$/.test(id)) {
    throw createError({
      statusCode: 400,
      statusMessage: 'invalid license id'
    })
  }

  const response = await $fetch<ApiMessageResponse>(`${config.apiBase}/api/license/delete/${id}`, {
    method: 'POST'
  })

  if (response.code !== 0) {
    throw createError({
      statusCode: 502,
      statusMessage: response.message || 'License delete failed'
    })
  }

  return response
})
