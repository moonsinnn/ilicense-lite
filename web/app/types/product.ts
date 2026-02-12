export interface Product {
  id: number
  code: string
  name: string
  description: string
  status: number
  created_at: string
  updated_at: string
}

export interface ProductQueryBody {
  page?: number
  size?: number
}

export interface ProductQueryData {
  total: number
  items: Product[]
}

export interface ProductAddBody {
  code: string
  name: string
  description: string
}

export interface ProductDeleteBody {
  ids: number[]
}
