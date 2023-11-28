import { useCallback, useEffect, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import './App.css'
import { EWsMessageType, webSocketMessage } from './@types/socket.type'

function App (): JSX.Element {
  const [message, setMessage] = useState('')
  const [inputValue, setInputValue] = useState('')
  const [messages, setMessages] = useState<string[]>([])
  const token: string | null = localStorage.getItem('authToken')

  const { sendMessage: sendWsMessage, lastMessage } =
    useWebSocket<webSocketMessage>(`ws://localhost:5000/ws?token=${token}`)

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value)
    setMessage(e.target.value)
  }, [])

  const handleClick = (): void => {
    console.log('=> click ', message)
    const msg: webSocketMessage = {
      type: EWsMessageType.message,
      content: {
        message: message,
      },
    }
    sendWsMessage(JSON.stringify(msg))
  }
  useEffect(() => {
    if (lastMessage !== null) {
      console.log('=> last message : ', lastMessage)
      const message = JSON.parse(lastMessage.data)
      setMessages((messages) => [...messages, message.content.message])
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
      <button onClick={handleClick}>Send</button>
      <ul>
        {messages.map((message, i) => (
          <li key={i}>{message}</li>
        ))}
      </ul>
    </div>
  )
}

export default App
