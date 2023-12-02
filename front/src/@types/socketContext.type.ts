import {
  IWebSocketMessageContent,
  type IRoom,
  type IwebSocketMessageIn,
  type IwebSocketMessageOut
} from './socket.type'

export interface SocketContextInterface {
  room: IRoom | null
  lastMessage: MessageEvent<any> | null
  sendBroadcastMessage: (message: string) => void
  createRoom: (name: string) => void
  messages: Array<IwebSocketMessageIn | IwebSocketMessageOut>
  roomMessages: IWebSocketMessageContent[]
  availableRoomList: IRoom[]
  sendToRoom: (message: string) => void
  // setMessages: (messages: Array<IwebSocketMessageIn | IwebSocketMessageOut>) => void
}
