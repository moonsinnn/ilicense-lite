import type { Issuer } from '~/types/issuer'

const issuers: Issuer[] = [{
  id: 1,
  code: 'issuer-a-code',
  name: 'issuer-a-name',
  description: 'alex.smith',
  public_key: '13683592712',
  private_key: 'alex.smith@example.com',
  key_algorithm: 'New York, USA',
  key_size: 1024,
  status: 1,
  created_at: '2026-01-01 12:12:12',
  updated_at: '2026-01-01 12:12:12'
}, {
  id: 2,
  code: 'issuer-b-code',
  name: 'issuer-b-name',
  description: 'Jordan Brown',
  public_key: '13683592712',
  private_key: 'jordan.brown@example.com',
  key_algorithm: 'jordan.brown@example.com',
  key_size: 1024,
  status: 1,
  created_at: '2026-01-01 12:12:12',
  updated_at: '2026-01-01 12:12:12'
}, {
  id: 3,
  code: 'issuer-c-code',
  name: 'issuer-c-name',
  description: 'Taylor Green',
  public_key: '13683592712',
  private_key: 'taylor.green@example.com',
  key_algorithm: 'Paris, France',
  key_size: 1024,
  status: 0,
  created_at: '2026-01-01 12:12:12',
  updated_at: '2026-01-01 12:12:12'
}]

export default eventHandler(async () => {
  return issuers
})
