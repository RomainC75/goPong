import { useCallback, useEffect, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import './App.css'
import { EWsMessageTypeIn, EWsMessageTypeOut, webSocketMessageIn, webSocketMessageOut } from './@types/socket.type'

interface WSMessage {
  type: string
  content: {
    message: string
    userEmail?: string
    userId?: string
  }
}

function App (): JSX.Element {
  const [message, setMessage] = useState('')
  const [inputValue, setInputValue] = useState('')
  const [messages, setMessages] = useState<WSMessage[]>([])
  const token: string | null = localStorage.getItem('authToken')

  const { sendMessage: sendWsMessage, lastMessage } =
    useWebSocket<webSocketMessageIn>(`ws://localhost:5000/ws?token=${token}`)

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value)
    setMessage(e.target.value)
  }, [])

  const handleBroadcast = (): void => {
    console.log('=> click ', message)
    const msg: webSocketMessageOut = {
      type: EWsMessageTypeOut.broadcast,
      content: {
        message
      },
    }
    sendWsMessage(JSON.stringify(msg))
  }
  useEffect(() => {
    if (lastMessage !== null) {
      console.log('=> last message : ', lastMessage)
      const message = JSON.parse(lastMessage.data)
      setMessages((messages) => [...messages, message])
    }
  }, [lastMessage])

  return (
    <div className='App'>
      <input
        id='input'
        type='text'
        value={inputValue}
        onChange={handleChange}
      />
      <button onClick={handleBroadcast}>Send</button>
      <ul>
        {messages.map((message, i) => (
          <li key={i} className={message.type === EWsMessageTypeIn.broadcast ? 'broadcast' : 'message' }>{message.content.message}</li>
        ))}
      </ul>
    </div>
  )
}

export default App
