export interface License {
  id: number
  code: string
  product_id: number
  customer_id: number
  issuer_id: number
  activation_code: string
  issue_at: string
  expire_at: string
  modules: string
  max_instances: number
  status: number
  remarks: string
  created_at: string
  updated_at: string
}

export interface LicenseQueryBody {
  page?: number
  size?: number
}

export interface LicenseQueryData {
  total: number
  items: License[]
}

export interface LicenseAddBody {
  code: string
  product_id: number
  customer_id: number
  issuer_id: number
  expire_at: string
  modules: string
  max_instances: number
  remarks: string
}

export interface LicenseDeleteBody {
  ids: number[]
}

export interface LicenseRenewBody {
  id: number
  expire_at: string
  remarks: string
}

export interface LicenseActivateBody {
  issuer_id: number
  code: string
}

export interface LicenseActivateData {
  issuer_code: string
  issuer_name: string
  customer_code: string
  customer_name: string
  product_code: string
  product_name: string
  license_code: string
  issue_at: string
  expire_at: string
  modules: string
  max_instances: number
  ok: boolean
}
