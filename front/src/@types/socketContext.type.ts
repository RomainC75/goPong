import {
  type IWebSocketMessageContent,
  type IRoom,
  IGame,
} from './socket.type'

export interface SocketContextInterface {
  room: IRoom | null
  lastMessage: MessageEvent<any> | null
  sendBroadcastMessage: (message: string) => void
  createRoom: (name: string) => void
  broadcastMessages: IWebSocketMessageContent[]
  roomMessages: IWebSocketMessageContent[]
  availableRoomList: IRoom[]
  availableGameList: IGame[]
  sendToRoom: (message: string) => void
  connectToRoom: (roomId: string) => void
  disconnectFromRoom: () => void
  createGame: (name: string) => void
  selectGame: (id: string) => void
  // setMessages: (messages: Array<IwebSocketMessageIn | IwebSocketMessageOut>) => void
}
