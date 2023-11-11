export interface UserInterface {
  id: string
  email: string
}

export interface AuthContextInterface {
  isLoggedIn: boolean
  isLoading: boolean
  user: UserInterface | null
  API_URL: string
  storeToken: (state: string) => void
  authenticateUser: () => void
  logOutUser: () => void
  removeToken: () => void
}

export interface LoginInterface {
  email: string
  password: string
}

export interface SignupInterface extends LoginInterface {
  pseudo: string
}

export interface SignupFullInterface extends SignupInterface {
  emailConf: string
  passwordConf: string
}
