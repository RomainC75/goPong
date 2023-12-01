import { useState, createContext, type PropsWithChildren, useEffect } from 'react'
import { type SocketContextInterface } from '../@types/socketContext.type'
import {
  EWsMessageTypeIn,
  EWsMessageTypeOut,
  type IRoom,
  type IwebSocketMessageIn,
  type IwebSocketMessageOut
} from '../@types/socket.type'
import useWebSocket from 'react-use-websocket'

const SocketContext = createContext<SocketContextInterface | null>(null)

const SocketProviderWrapper = (props: PropsWithChildren): JSX.Element => {
  const token: string | null = localStorage.getItem('authToken')

  const { sendMessage: sendWsMessage, lastMessage } =
    useWebSocket<IwebSocketMessageIn>(`ws://localhost:5000/ws?token=${token ?? ''}`)

  const [messages, setMessages] = useState<Array<IwebSocketMessageIn | IwebSocketMessageOut>>([])
  const [room, setRoom] = useState<IRoom | null>(null)
  const [availableRoomList, setAvailableRoomList] = useState<IRoom[]>([])
  const [roomMessages, setRoomMessages] = useState<Array<IwebSocketMessageIn | IwebSocketMessageOut>>([])
  // const [broadcastmessage, setBroadcastMessage] = useState<string>('')

  const sendBroadcastMessage = (message: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.broadcast,
      content: {
        message
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const createRoom = (roomName: string): void => {
    console.log('=> click ', roomName)
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.createRoom,
      content: {
        roomName
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const sendToRoom = (message: string) => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.sendToRoom,
      content: {
        message
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  useEffect(() => {
    if (lastMessage !== null) {
      console.log('=> last message : ', lastMessage)
      const message = JSON.parse(lastMessage.data)
      switch (message.type) {
        case EWsMessageTypeIn.broadcast:
          setMessages(messages => ([...messages, message] as Array<IwebSocketMessageIn | IwebSocketMessageOut>))
          break
        case EWsMessageTypeIn.roomCreatedByYou:
          setRoom(message.content)
          break
        case EWsMessageTypeIn.roomCreated:
          setAvailableRoomList(roomList => ([...roomList, message.content]))
          break
        case EWsMessageTypeIn.roomMessage:
          setRoomMessages(roomMessages => ([...roomMessages, message.content]))
          break
      }
    }
  }, [lastMessage])

  return (
    <SocketContext.Provider
      value={{
        sendBroadcastMessage,
        lastMessage,
        createRoom,
        messages,
        room,
        availableRoomList,
        sendToRoom,
        roomMessages
      }}
    >
      {props.children}
    </SocketContext.Provider>
  )
}

export { SocketContext, SocketProviderWrapper }
