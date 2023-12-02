import { type ChangeEvent, useContext, useState } from 'react'
import { SocketContext } from '../../context/socket.context'
import { SocketContextInterface } from '../../@types/socketContext.type'
import { Button, TextField } from '@mui/material'

const Message = (): JSX.Element => {
  const { sendBroadcastMessage } = useContext(
    SocketContext
  ) as SocketContextInterface

  const [message, setMessage] = useState<string>('')

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    setMessage(e.target.value)
  }

  const handleSendMessage = (): void => {
    sendBroadcastMessage(message)
    setMessage('')
  }

  return (
    <div>
      <TextField id="outlined-basic" label="Broadcast" variant="outlined" value={message} onChange={handleChange} size="small"/>
      <Button variant="contained" onClick={handleSendMessage}>send</Button>
    </div>
  )
}

export default Message
