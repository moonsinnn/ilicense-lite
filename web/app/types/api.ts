export interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

export interface ApiMessageResponse {
  code: number
  message: string
}
