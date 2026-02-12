export interface Customer {
  id: number
  code: string
  name: string
  contact: string
  phone: string
  email: string
  address: string
  status: number
  created_at: string
  updated_at: string
}

export interface CustomerQueryBody {
  page?: number
  size?: number
}

export interface CustomerQueryData {
  total: number
  items: Customer[]
}

export interface CustomerAddBody {
  code: string
  name: string
  contact?: string
  phone?: string
  email?: string
  address?: string
}

export interface CustomerDeleteBody {
  ids: number[]
}
