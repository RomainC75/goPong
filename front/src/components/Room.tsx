import { type ChangeEvent, useContext, useState } from "react";
import { SocketContext } from "../context/socket.context";
import { type SocketContextInterface } from "../@types/socketContext.type";

export const Room = (): JSX.Element => {
  const { room, sendToRoom, roomMessages } = useContext(
    SocketContext
  ) as SocketContextInterface;

  const [message, setMessage] = useState<string>("");

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    setMessage(e.target.value);
  };

  const handleSendMessage = () => {
    sendToRoom(message);
    setMessage('');
  };

  return (
    <div>
      <h3>{room?.name}</h3>
      <div>
        <input id="input" type="text" value={message} onChange={handleChange} />
        <button onClick={handleSendMessage}>send messages</button>
      </div>
      <li>
        {roomMessages.map((messages, i) => (
          <li key={i}>{JSON.stringify(message)}</li>
        ))}
      </li>
    </div>
  );
};
