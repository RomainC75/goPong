import { type FormEvent, useState, type ChangeEvent, useContext } from 'react'
import './styles/login.scss'
import { AuthContext } from '../context/auth.context'
import { type AuthContextInterface, type LoginInterface } from '../@types/authContext.type'
import Alert from '@mui/material/Alert'
import AlertTitle from '@mui/material/AlertTitle'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

import { PurpleTextField, PurpleButton } from '../utils/mui-custom-colors'

const Login = (): JSX.Element => {
  const { authenticateUser, API_URL, storeToken } = useContext(
    AuthContext
  ) as AuthContextInterface
  const navigate = useNavigate()
  const [inputsState, setInputsState] = useState<LoginInterface>({
    email: '',
    password: ''
  })
  const [isLoginValid, setIsLoginValid] = useState<boolean>(false)
  const [isLoginError, setIsLoginError] = useState<boolean>(false)

  const handleInputs = (
    e: ChangeEvent<HTMLTextAreaElement | HTMLInputElement>
  ): void => {
    if ('value' in e.target && 'name' in e.target) {
      setIsLoginError(false)
      const newValues: LoginInterface = {
        ...inputsState,
        [e.target.name]: e.target.value
      }
      setInputsState(newValues)
      if (Object.values(newValues).includes('')) {
        setIsLoginValid(false)
      } else {
        setIsLoginValid(true)
      }
    }
  }

  const handleForm = (e: FormEvent<HTMLFormElement>): void => {
    e.preventDefault()
    axios
      .post(`${API_URL}/auth/signin`, inputsState)
      .then(ans => {
        console.log('=> ans : ', ans)
        storeToken(ans.data.token)
        authenticateUser()
        navigate('/')
      })
      .catch(() => {
        setIsLoginError(true)
      })
  }

  return (
    <div className="Login">
      <h1>Login</h1>
      <form onSubmit={handleForm}>
        <PurpleTextField
          id="email"
          name="email"
          label="email"
          variant="outlined"
          value={inputsState.email}
          onChange={handleInputs}
        />
        <PurpleTextField
          id="password"
          name="password"
          label="password"
          type="password"
          autoComplete="current-password"
          variant="outlined"
          value={inputsState.password}
          onChange={handleInputs}
        />
        <PurpleButton
          variant="contained"
          type="submit"
          disabled={!isLoginValid}
        >
          Login
        </PurpleButton>
        {isLoginError && (
          <Alert severity="error">
            <AlertTitle>Error</AlertTitle>
            wrong email or password
          </Alert>
        )}
      </form>
    </div>
  )
}

export default Login
