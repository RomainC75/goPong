import { useContext } from 'react'
import { AuthContext } from '../../context/auth.context'
import { type AuthContextInterface } from '../../@types/authContext.type'
import { Navigate } from 'react-router-dom'

interface IsPrivateInterface {
  children: JSX.Element
}

const IsPrivate = ({ children }: IsPrivateInterface): JSX.Element => {
  const { isLoading, isLoggedIn } = useContext(
    AuthContext
  ) as AuthContextInterface
  if (isLoading) return <p>Loading...</p>
  if (isLoggedIn) {
    return <>{children}</>
  } else {
    return <Navigate to="/auth" />
  }
}

export default IsPrivate
