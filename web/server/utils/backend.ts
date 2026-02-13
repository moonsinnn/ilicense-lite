import type { H3Event } from 'h3'
import { getCookie, getHeader } from 'h3'

function getAuthorization(event: H3Event): string | undefined {
  const headerAuth = getHeader(event, 'authorization')
  if (headerAuth) return headerAuth

  const token = getCookie(event, 'auth_token')
  if (!token) return undefined

  return `Bearer ${token}`
}

type BackendMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'

interface BackendFetchOptions {
  method?: BackendMethod
  headers?: Record<string, string>
  body?: BodyInit | Record<string, unknown> | null
}

interface BackendApiResponse {
  code: number
  message?: string
}

export async function backendFetch<T>(event: H3Event, path: string, options: BackendFetchOptions = {}): Promise<T> {
  const config = useRuntimeConfig(event)
  const authorization = getAuthorization(event)
  const headers = {
    ...(options.headers || {}),
    ...(authorization ? { Authorization: authorization } : {})
  }
  const fetchOptions = {
    ...options,
    headers
  }

  try {
    const response = await $fetch(`${config.apiBase}${path}`, fetchOptions as never)
    return response as T
  } catch (error: unknown) {
    const err = error as {
      response?: { status?: number, _data?: { message?: string } }
      statusCode?: number
      statusMessage?: string
    }
    const statusCode = Number(err.response?.status || err.statusCode || 500)
    const statusMessage = String(err.response?._data?.message || err.statusMessage || 'backend request failed')
    throw createError({
      statusCode,
      statusMessage
    })
  }
}

export function ensureApiSuccess<T extends BackendApiResponse>(
  response: T,
  fallbackMessage: string,
  statusCode = 502
): T {
  if (response.code === 0) return response
  throw createError({
    statusCode,
    statusMessage: response.message || fallbackMessage
  })
}
