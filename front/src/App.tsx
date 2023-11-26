import { useCallback, useEffect, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import './App.css'


function App (): JSX.Element {
  const [message, setMessage] = useState('')
  const [inputValue, setInputValue] = useState('')
  
  
  const { sendMessage: sendWsMessage, lastMessage } = useWebSocket('ws://localhost:5000/ws')

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value)
    setMessage(e.target.value)
  }, [])

  const handleClick = (): void =>{
    console.log("=> click ", message)
    const msg = {
      type: 'MESSAGE',
      content: {
        message: message
      }
    }
    sendWsMessage(JSON.stringify(msg))
    
  }

  return (
    <div className="App">
      <input id="input" type="text" value={inputValue} onChange={handleChange} />
      <button onClick={handleClick}>Send</button>
      <pre>{message}</pre>
    </div>
  )
}

export default App
