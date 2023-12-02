import { type ChangeEvent, useContext, useState } from 'react';
import { SocketContext } from '../context/socket.context';
import { type SocketContextInterface } from '../@types/socketContext.type';
import { AuthContext } from '../context/auth.context';
import { type AuthContextInterface } from '../@types/authContext.type';
import { TextField, Button } from '@mui/material';
import './styles/room.scss';
import MessageBox from './MessageBox';

export const Room = (): JSX.Element => {
  const { room, sendToRoom, roomMessages } = useContext(
    SocketContext
  ) as SocketContextInterface;
  const { user } = useContext(AuthContext) as AuthContextInterface;
  const [message, setMessage] = useState<string>('');

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    setMessage(e.target.value);
  };

  const handleSendMessage = (): void => {
    sendToRoom(message);
    setMessage('');
  };

  return (
    <div className='Room'>
      <h3>Room : {room?.name}</h3>
      <MessageBox messages={roomMessages}/>
      <div>
        <TextField
          id='outlined-basic'
          label='Outlined'
          variant='outlined'
          value={message}
          onChange={handleChange}
          size='small'
        />
        <Button variant='contained' onClick={handleSendMessage}>
          Send Message
        </Button>
      </div>
    </div>
  );
};
