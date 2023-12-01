import { useContext, useEffect } from 'react'
import './App.css'
import { SocketContext } from './context/socket.context'
import CreateRoom from './components/inputs/CreateRoom'
import { type SocketContextInterface } from './@types/socketContext.type'
import Message from './components/inputs/Message'
import { EWsMessageTypeIn } from './@types/socket.type'

function App (): JSX.Element {
  const { messages } = useContext(
    SocketContext
  ) as SocketContextInterface

  return (
    <div className='App'>
      <Message/>
      <CreateRoom />
      <ul>
        {messages.map((message, i) => (
          <li key={i} className={message.type === EWsMessageTypeIn.broadcast ? 'broadcast' : 'message' }>{message.content.message}</li>
        ))}
      </ul>
    </div>
  )
}

export default App
