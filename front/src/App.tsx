import { useContext, useEffect } from 'react'
import './App.css'
import { SocketContext } from './context/socket.context'
import CreateRoom from './components/inputs/CreateRoom'
import { type SocketContextInterface } from './@types/socketContext.type'
import Message from './components/inputs/Message'
import { Room } from './components/Room'
import RoomList from './components/RoomList'
import MessageBox from './components/MessageBox'
import DisconnectFromRoom from './components/DisconnectFromRoom'
import CreateGame from './components/CreateGame'
import GameList from './components/GameList'
import Game from './components/Game'

function App (): JSX.Element {
  const { broadcastMessages, room, availableRoomList, currentGame, currentGameConfig } = useContext(
    SocketContext
  ) as SocketContextInterface
  return (
    <div className='App'>
      <h1>Tron</h1>
      <div>
        <CreateGame/>
        <GameList/>
        {currentGame && currentGameConfig && <Game/>}
        <Message/>
        <MessageBox messages={broadcastMessages}/>
        {room ? <DisconnectFromRoom/> : <CreateRoom />}
      </div>
      <div>
        {room && <Room/>}
        {availableRoomList && <RoomList/>}
      </div>
    </div>
  )
}

export default App
