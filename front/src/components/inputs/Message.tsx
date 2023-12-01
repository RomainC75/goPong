import React, { ChangeEvent, useContext, useState } from "react";
import { SocketContext } from "../../context/socket.context";
import { SocketContextInterface } from "../../@types/socketContext.type";


const Message = (): JSX.Element => {
  const { sendBroadcastMessage } = useContext(
      SocketContext
    ) as SocketContextInterface

  const [message, setMessage] = useState<string>('');

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    setMessage(e.target.value)
  }

  const handleSendMessage = () =>{
    sendBroadcastMessage(message)
    setMessage("")
  }

  return (
    <div>
      <input id='input' type='text' value={message} onChange={handleChange } />
      <button onClick={handleSendMessage}>send messages</button>
    </div>
  )
}

export default Message
