import { useContext } from 'react'
import './App.css'
import { SocketContext } from './context/socket.context'
import CreateRoom from './components/inputs/CreateRoom'
import { type SocketContextInterface } from './@types/socketContext.type'
import Message from './components/inputs/Message'
import { Room } from './components/Room'
import RoomList from './components/RoomList'
import MessageBox from './components/MessageBox'

function App (): JSX.Element {
  const { broadcastMessages, room, availableRoomList } = useContext(
    SocketContext
  ) as SocketContextInterface

  return (
    <div className='App'>
      <h1>Tron</h1>
      <div>
        <Message/>
        <MessageBox messages={broadcastMessages}/>
        {!room && <CreateRoom />}
      </div>
      <div>
        {room && <Room/>}
        {availableRoomList && <RoomList/>}
      </div>
    </div>
  )
}

export default App
