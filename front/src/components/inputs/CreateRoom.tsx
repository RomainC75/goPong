import React, { ChangeEvent, useContext, useState } from "react";
import { SocketContext } from "../../context/socket.context";
import { SocketContextInterface } from "../../@types/socketContext.type";
import { Button, TextField } from "@mui/material";

const CreateRoom = (): JSX.Element => {
  const { createRoom } = useContext(SocketContext) as SocketContextInterface;

  const [name, setName] = useState<string>("");

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    setName(e.target.value);
  };

  const handleCreateRoom = () => {
    createRoom(name);
  };

  return (
    <div>
      <TextField
        id="outlined-basic"
        label="Room name"
        variant="outlined"
        value={name}
        onChange={handleChange}
        size="small"
      />
      <Button variant="contained" onClick={handleCreateRoom}>
        create Room
      </Button>
    </div>
  );
};

export default CreateRoom;
