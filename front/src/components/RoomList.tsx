import { useContext } from 'react'
import { SocketContext } from '../context/socket.context'
import { SocketContextInterface } from '../@types/socketContext.type'
import { IRoom } from '../@types/socket.type'

const RoomList = () => {
  const { availableRoomList, room, connectToRoom } = useContext(
      SocketContext
      ) as SocketContextInterface

  const handleConnectToRoom = (availableRoom: IRoom): boolean =>{
    connectToRoom(availableRoom.id)
    return true
  }
  
  return (
    <div>
        <h3>rooms</h3>
        <ul>
            {availableRoomList.map(availableRoom => <li key={availableRoom.id}>
              {availableRoom.id === room?.id
              ?
              <p>{availableRoom.name + ' -self'}</p>
              :
              <p onClick={() => handleConnectToRoom(availableRoom) }>{availableRoom.name}</p>
              }
              
              
            </li>)}
        </ul>
    </div>
  )
}

export default RoomList