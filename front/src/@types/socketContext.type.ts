export interface SocketContextInterface {
  broadcastmessage: string
  lastMessage: MessageEvent<any> | null
  setBroadcastMessage: (message: string) => void
  sendBroadcastMessage: () => void
  createRoom: (name: string) => void
}
