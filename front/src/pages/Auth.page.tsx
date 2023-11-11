import { useState } from 'react'

import './styles/auth.css'
import Login from '../components/Login'
import Signup from '../components/Signup'

const AuthPage = (): JSX.Element => {
  const [isLoginNotSignup, setIsLoginNotSignup] = useState<boolean>(true)

  const toggleLoginOrSignup = (): void => {
    setIsLoginNotSignup(!isLoginNotSignup)
  }

  return (
    <div className="AuthPage">
      <div className="container">
        <div className="content left">
          {isLoginNotSignup ? <Login /> : <Signup setIsLoginNotSignup={setIsLoginNotSignup}/>}
          <div className="center">
                <p>
                  { isLoginNotSignup ? "Don't h" : 'H'}ave an account ?{' '}
                  <span
                    onClick={toggleLoginOrSignup}
                    className="color1 bold cursor"
                  >
                    { isLoginNotSignup ? 'Sign up' : 'Login'}
                  </span>
                </p>
              </div>
        </div>
        <div className="animation right">
          <div className="sunContainer top">
            <div className="sun top"></div>
          </div>
          <div className="sunContainer bottom">
            <div className="sun bottom"></div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default AuthPage
