import {
  type IWebSocketMessageContent,
  type IRoom,
} from './socket.type'

export interface SocketContextInterface {
  room: IRoom | null
  lastMessage: MessageEvent<any> | null
  sendBroadcastMessage: (message: string) => void
  createRoom: (name: string) => void
  broadcastMessages: IWebSocketMessageContent[]
  roomMessages: IWebSocketMessageContent[]
  availableRoomList: IRoom[]
  sendToRoom: (message: string) => void
  // setMessages: (messages: Array<IwebSocketMessageIn | IwebSocketMessageOut>) => void
}
