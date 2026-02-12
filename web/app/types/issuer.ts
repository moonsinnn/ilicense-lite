export interface Issuer {
  id: number
  code: string
  name: string
  description: string
  public_key: string
  private_key: string
  key_algorithm: string
  key_size: number
  status: number
  created_at: string
  updated_at: string
}

export interface IssuerQueryBody {
  page?: number
  size?: number
}

export interface IssuerQueryData {
  total: number
  items: Issuer[]
}

export interface IssuerAddBody {
  code: string
  name: string
  description: string
}

export interface IssuerDeleteBody {
  ids: number[]
}
