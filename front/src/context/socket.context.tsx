import { useState, createContext, type PropsWithChildren } from 'react';
import { type SocketContextInterface } from '../@types/socketContext.type';
import {
  EWsMessageTypeOut,
  type webSocketMessageOut,
  type EWsMessageTypeIn,
  type webSocketMessageIn,
} from '../@types/socket.type';
import useWebSocket from 'react-use-websocket';

const SocketContext = createContext<SocketContextInterface | null>(null);

const SocketProviderWrapper = (props: PropsWithChildren): JSX.Element => {
  const token: string | null = localStorage.getItem('authToken')

  const { sendMessage: sendWsMessage, lastMessage } =
    useWebSocket<webSocketMessageIn>(`ws://localhost:5000/ws?token=${token ?? ''}`)

  // const [broadcastmessage, setBroadcastMessage] = useState<string>('')

  const sendBroadcastMessage = (message: string): void => {
    console.log('=> broadcast ', message)
    const msg: webSocketMessageOut = {
      type: EWsMessageTypeOut.broadcast,
      content: {
        message
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const createRoom = (roomName: string): void => {
    console.log('=> click ', roomName)
    const msg: webSocketMessageOut = {
      type: EWsMessageTypeOut.createRoom,
      content: {
        roomName
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  // useEffect(() => {
  //   //
  // }, [])

  return (
    <SocketContext.Provider
      value={{
        sendBroadcastMessage,
        lastMessage,
        createRoom
      }}
    >
      {props.children}
    </SocketContext.Provider>
  );
};

export { SocketContext, SocketProviderWrapper };
