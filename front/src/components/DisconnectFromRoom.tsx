import { useContext } from 'react'
import { SocketContext } from '../context/socket.context'
import { type SocketContextInterface } from '../@types/socketContext.type'
import { Button } from '@mui/material'

const DisconnectFromRoom = () => {
  const { disconnectFromRoom } = useContext(
      SocketContext
    ) as SocketContextInterface
  

  return (
    <Button variant="contained" onClick={disconnectFromRoom} color="error">Disconnect</Button>
  )
}

export default DisconnectFromRoom