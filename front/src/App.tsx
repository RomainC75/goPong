import { useContext, useEffect } from 'react'
import './App.css'
import { SocketContext } from './context/socket.context'
import CreateRoom from './components/inputs/CreateRoom'
import { type SocketContextInterface } from './@types/socketContext.type'
import Message from './components/inputs/Message'

function App (): JSX.Element {
  const { lastMessage } = useContext(
    SocketContext
  ) as SocketContextInterface

  useEffect(() => {
    if (lastMessage !== null) {
      console.log('=> last message : ', lastMessage)
      // const message = JSON.parse(lastMessage.data)
      // setMessages((messages) => [...messages, message])
    }
  }, [lastMessage])

  return (
    <div className='App'>
      <Message/>
      <CreateRoom />
      {/* <ul>
        {messages.map((message, i) => (
          <li key={i} className={message.type === EWsMessageTypeIn.broadcast ? 'broadcast' : 'message' }>{message.content.message}</li>
        ))}
      </ul> */}
    </div>
  )
}

export default App
