export interface SocketContextInterface {
  lastMessage: MessageEvent<any> | null
  sendBroadcastMessage: (message: string) => void
  createRoom: (name: string) => void
}
