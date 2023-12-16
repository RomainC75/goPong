import {
  type IWebSocketMessageContent,
  type IRoom,
  IGame,
  IGameState,
  IGameConfig,
  IGridDot,
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
  currentGame: IGame | null
  gameState: IGameState | null
  setCurrentGame: (currentGame: IGame) => void
  currentGameConfig: IGameConfig | null
  // setMessages: (messages: Array<IwebSocketMessageIn | IwebSocketMessageOut>) => void
  sendKeyCode: (code: number) => void
  grid: IGridDot[][];
}
