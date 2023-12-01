import { useContext } from 'react'
import { SocketContext } from '../context/socket.context'
import { SocketContextInterface } from '../@types/socketContext.type'

const RoomList = () => {
    const { availableRoomList } = useContext(
        SocketContext
        ) as SocketContextInterface
  return (
    <div>
        <h3>rooms</h3>
        <ul>
            {availableRoomList.map(availableRoom => <li key={availableRoom.id}>{availableRoom.name}</li>)}
        </ul>
    </div>
  )
}

export default RoomList