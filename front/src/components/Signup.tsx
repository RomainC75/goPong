import { useState, useContext } from 'react'
import type { ChangeEvent, FormEvent } from 'react'
import { AuthContext } from '../context/auth.context'
import axios from 'axios'
import Alert from '@mui/material/Alert'
import AlertTitle from '@mui/material/AlertTitle'
import { PurpleButton, PurpleTextField } from '../utils/mui-custom-colors'
import {
  detailsAboutNeededCharactersInPass,
  isEmailValidFn,
  isPasswordValidFn
} from '../utils/signugFieldsTests'

import type { SignupFullInterface, AuthContextInterface } from '../@types/authContext.type'

import './styles/signup.scss'

interface SignupComponentInterface {
  setIsLoginNotSignup: (isLoginNotSignup: boolean) => void
}

const Signup = ({
  setIsLoginNotSignup
}: SignupComponentInterface): JSX.Element => {
  const { API_URL } = useContext(AuthContext) as AuthContextInterface
  const [inputsState, setInputsState] = useState<SignupFullInterface>({
    email: '',
    password: '',
    pseudo: '',
    emailConf: '',
    passwordConf: ''
  })
  const [isSignupError, setIsSignupError] = useState<boolean>(false)

  const [isFirstNameValid, setIsFirstNameValid] = useState<boolean>(true)
  const [isLastNameValid] = useState<boolean>(true)

  const [isEmailValid, setIsEmailValid] = useState<boolean>(true)
  const [isPasswordValid, setIsPasswordValid] = useState<boolean>(true)
  const [isPasswordsEquals, setIsPasswordsEquals] = useState<boolean>(true)
  const [isEmailsEquals, setIsEmailsEquals] = useState<boolean>(true)

  const handleInputs = (
    e: ChangeEvent<HTMLTextAreaElement | HTMLInputElement>
  ): void => {
    if ('value' in e.target && 'name' in e.target) {
      setIsSignupError(false)
      const newValues: SignupFullInterface = {
        ...inputsState,
        [e.target.name]: e.target.value
      }
      setInputsState(newValues)

      setIsFirstNameValid(newValues.pseudo.length > 0)

      setIsEmailValid(isEmailValidFn(newValues.email))
      setIsEmailsEquals(newValues.email === newValues.emailConf)

      setIsPasswordValid(isPasswordValidFn(newValues.password))
      setIsPasswordsEquals(newValues.password === newValues.passwordConf)
    }
  }

  const handleForm = (e: FormEvent<HTMLFormElement>): void => {
    e.preventDefault()
    axios
      .post(`${API_URL}/auth/signup`, inputsState)
      .then(ans => {
        console.log('signup', ans.data)
        setIsLoginNotSignup(true)
      })
      .catch(err => {
        setIsSignupError(true)
        console.log('err : ', err)
      })
  }

  return (
    <div className="Signup">
      <h1>Signup</h1>
      <form onSubmit={handleForm}>
        <PurpleTextField
          id="pseudo"
          name="pseudo"
          label="pseudo"
          variant="outlined"
          value={inputsState.pseudo}
          onChange={handleInputs}
          helperText={!isLastNameValid && 'need a last name'}
          error={!isLastNameValid}
        />
        <PurpleTextField
          id="email"
          name="email"
          label="email"
          variant="outlined"
          value={inputsState.email}
          onChange={handleInputs}
          helperText={!isEmailValid && 'need a valid email'}
          error={!isEmailValid}
        />
        <PurpleTextField
          id="emailConf"
          name="emailConf"
          label="email confirmation"
          variant="outlined"
          value={inputsState.emailConf}
          onChange={handleInputs}
          helperText={!isEmailsEquals && 'need the same email'}
          error={!isEmailsEquals}
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
          helperText={
            !isPasswordValid &&
            detailsAboutNeededCharactersInPass(inputsState.password)
          }
          error={!isPasswordValid}
        />
        <PurpleTextField
          id="passwordConf"
          name="passwordConf"
          label="password confirmation"
          type="password"
          autoComplete="current-passwordConf"
          variant="outlined"
          value={inputsState.passwordConf}
          onChange={handleInputs}
          helperText={!isPasswordsEquals && 'need the same password'}
          error={!isPasswordsEquals}
        />
        <PurpleButton
          variant="contained"
          type="submit"
          disabled={
            !isFirstNameValid ||
            !isLastNameValid ||
            !isEmailValid ||
            !isEmailsEquals ||
            !isPasswordValid ||
            !isPasswordsEquals
          }
        >
          Login
        </PurpleButton>
        {isSignupError && (
          <Alert severity="error">
            <AlertTitle>Error</AlertTitle>
            wrong email or password
          </Alert>
        )}
      </form>
    </div>
  )
}

export default Signup
