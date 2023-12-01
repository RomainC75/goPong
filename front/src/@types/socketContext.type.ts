import { IwebSocketMessageIn, IwebSocketMessageOut } from "./socket.type"

export interface SocketContextInterface {
  lastMessage: MessageEvent<any> | null
  sendBroadcastMessage: (message: string) => void
  createRoom: (name: string) => void
  messages: Array<IwebSocketMessageIn | IwebSocketMessageOut>
  // setMessages: (messages: Array<IwebSocketMessageIn | IwebSocketMessageOut>) => void
}
