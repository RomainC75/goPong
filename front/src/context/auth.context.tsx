import {
  useState,
  createContext,
  useEffect,
  type PropsWithChildren,
} from 'react'
import {
  type AuthContextInterface,
  type UserInterface,
} from '../@types/authContext.type'
import { type NavigateFunction, useNavigate } from 'react-router-dom'
import axios from 'axios'

const API_URL: string = process.env.REACT_APP_API_URL || 'http://localhost:5000'

const AuthContext = createContext<AuthContextInterface | null>(null)

const AuthProviderWrapper = (props: PropsWithChildren): JSX.Element => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false)
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const [user, setUser] = useState<UserInterface | null>(null)
  const navigate: NavigateFunction = useNavigate()

  const storeToken = (token: string): void => {
    localStorage.setItem('authToken', token)
  }

  const removeToken = () => {
    localStorage.removeItem('authToken')
    setUser(null)
    navigate('/auth')
  }

  const logOutUser = () => {
    setUser(null)
  }

  const authenticateUser = (): void => {
    const storedToken: string | null = localStorage.getItem('authToken')

    if (!storedToken) {
      setIsLoggedIn(false)
      setIsLoading(false)
      setUser(null)
      return
    }
    setIsLoading(true)
    console.log("=> VERIRY : ", storedToken)
    axios
      .get(`${API_URL}/auth/verify`, {
        headers: {
          Authorization: `Bearer ${storedToken}`,
        },
      })
      .then(ans => {
        setIsLoggedIn(true)
        setIsLoading(false)
        setUser(ans.data)
      })
      .catch(() => {
        setIsLoading(false)
        setIsLoading(false)
        removeToken()
        navigate('/')
      })
  }

  useEffect(() => {
    authenticateUser()
  }, [])

  return (
    <AuthContext.Provider
      value={{
        isLoggedIn,
        isLoading,
        user,
        storeToken,
        authenticateUser,
        logOutUser,
        API_URL,
        removeToken,
      }}
    >
      {props.children}
    </AuthContext.Provider>
  )
}

export { AuthContext, AuthProviderWrapper }
