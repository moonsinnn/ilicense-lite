export interface UserProfile {
  id: number
  username: string
  name: string
  email: string
  avatar: string
  created_at: string
  updated_at: string
}

export interface UserLoginBody {
  username: string
  password: string
}

export interface UserLoginData {
  token: string
  user: UserProfile
}

export interface UserProfileUpdateBody {
  name?: string
  email?: string
  avatar?: string
}

export interface UserPasswordUpdateBody {
  old_password: string
  new_password: string
}
