import { useContext, useEffect } from 'react'
import './App.css'
import { SocketContext } from './context/socket.context'
import CreateRoom from './components/inputs/CreateRoom'
import { type SocketContextInterface } from './@types/socketContext.type'
import Message from './components/inputs/Message'
import { EWsMessageTypeIn } from './@types/socket.type'
import { Room } from './components/Room'
import RoomList from './components/RoomList'

function App (): JSX.Element {
  const { messages, room, availableRoomList } = useContext(
    SocketContext
  ) as SocketContextInterface

  return (
    <div className='App'>
      <h1>Tron</h1>
      <div>
        <Message/>
        <ul>
          {messages.map((message, i) => (
            <li key={i} className={message.type === EWsMessageTypeIn.broadcast ? 'broadcast' : 'message' }>{message.content.message}</li>
          ))}
        </ul>
        <CreateRoom />
      </div>
      <div>
        {room && <Room/>}
        {availableRoomList && <RoomList/>}
      </div>
    </div>
  )
}

export default App
