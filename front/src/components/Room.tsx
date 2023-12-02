import { type ChangeEvent, useContext, useState } from 'react'
import { SocketContext } from '../context/socket.context'
import { type SocketContextInterface } from '../@types/socketContext.type'
import { AuthContext } from '../context/auth.context'
import { type AuthContextInterface } from '../@types/authContext.type'
import { TextField, Button } from '@mui/material';
import './styles/room.scss'

export const Room = (): JSX.Element => {
  const { room, sendToRoom, roomMessages } = useContext(
    SocketContext
  ) as SocketContextInterface
  const { user } = useContext(AuthContext) as AuthContextInterface

  const [message, setMessage] = useState<string>('')

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    setMessage(e.target.value)
  }

  const handleSendMessage = () => {
    sendToRoom(message)
    setMessage('')
  }

  return (
    <div className='Room'>
      <h3>Room : {room?.name}</h3>
      <div>
      <TextField id="outlined-basic" label="Outlined" variant="outlined" value={message} onChange={handleChange}/>
        {/* <input id='input' type='text' value={message} onChange={handleChange} /> */}
        <Button variant="contained" onClick={handleSendMessage}>Send Message</Button>
        {/* <button onClick={handleSendMessage}>send messages</button> */}
      </div>
      <li>
        {roomMessages.map((message, i) => (
          <li
            key={i}
            className={message.userEmail === user?.email ? 'sent' : 'received'}
          >
            {message.message}
          </li>
        ))}
      </li>
    </div>
  )
}
