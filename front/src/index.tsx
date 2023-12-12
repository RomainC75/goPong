import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import reportWebVitals from './reportWebVitals'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { AuthProviderWrapper } from './context/auth.context'
import AuthPage from './pages/Auth.page'
import IsPrivate from './components/helper/isPrivate'
import { SocketProviderWrapper } from './context/socket.context'

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
)
root.render(
  // <React.StrictMode>
    <BrowserRouter>
      <AuthProviderWrapper>
        <Routes>
          <Route path='/auth' element={<AuthPage/>} />
          <Route path='/' element={
            <IsPrivate>
              <SocketProviderWrapper>
                <App />
              </SocketProviderWrapper>
            </IsPrivate>
          }/>
        </Routes>
      </AuthProviderWrapper>
    </BrowserRouter>
  // </React.StrictMode>
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
