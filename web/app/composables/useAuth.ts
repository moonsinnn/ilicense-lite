import type { ApiMessageResponse, ApiResponse } from '~/types/api'
import type { UserLoginData, UserProfile, UserProfileUpdateBody, UserPasswordUpdateBody } from '~/types/user'

export function useAuth() {
  const token = useCookie<string | null>('auth_token', {
    default: () => null,
    sameSite: 'lax'
  })
  const user = useState<UserProfile | null>('auth_user', () => null)

  async function login(username: string, password: string) {
    const response = await $fetch<ApiResponse<UserLoginData>>('/api/user/login', {
      method: 'POST',
      body: { username, password }
    })
    token.value = response.data.token
    user.value = response.data.user
    return response.data
  }

  function logout() {
    token.value = null
    user.value = null
  }

  async function fetchProfile() {
    if (!token.value) {
      user.value = null
      return null
    }
    const response = await $fetch<ApiResponse<UserProfile>>('/api/user/profile', {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    })
    user.value = response.data
    return user.value
  }

  async function updateProfile(body: UserProfileUpdateBody) {
    const response = await $fetch<ApiResponse<UserProfile>>('/api/user/profile/update', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token.value || ''}`
      },
      body
    })
    user.value = response.data
    return response.data
  }

  async function updatePassword(body: UserPasswordUpdateBody) {
    return $fetch<ApiMessageResponse>('/api/user/password/update', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token.value || ''}`
      },
      body
    })
  }

  return {
    token,
    user,
    login,
    logout,
    fetchProfile,
    updateProfile,
    updatePassword
  }
}
