import React, { ChangeEvent, useContext, useState } from "react";
import { SocketContext } from "../../context/socket.context";
import { SocketContextInterface } from "../../@types/socketContext.type";


const CreateRoom = (): JSX.Element => {
  const { createRoom } = useContext(
      SocketContext
    ) as SocketContextInterface

  const [name, setName] = useState<string>('')

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    setName(e.target.value)
  }

  const handleCreateRoom = () => {
    createRoom(name)
  }

  return (
    <div>
      <input id='input' type='text' value={name} onChange={handleChange } />
      <button onClick={handleCreateRoom}>create room</button>
    </div>
  )
}

export default CreateRoom