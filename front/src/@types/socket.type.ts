export enum EWsMessageTypeIn {
  message = 'MESSAGE',
  idAssigned = 'IDASSIGNED',
  memberJoin = 'MEMBERJOIN',
  memberLeave = 'MEMBERLEAVE',
  broadcast = 'BROADCAST',
  roomCreated = 'ROOM_CREATED',
  roomCreatedByYou = 'ROOM_CREATED_BYYOU',
  roomMessage = 'ROOM_MESSAGE',
  connectedToRoom = 'CONNECTED_TO_ROOM',
  newConnectionToRoom = 'NEW_CONNECTION_TO_ROOM',
  disconnectedFromRoom = 'DISCONNECTED_FROM_ROOM',
  userDisconnectedFromRoom = 'USER_DISCONNECTED_FROM_ROOM',
  gameCreatedByYou = 'GAME_CREATED_BYYOU',
  gameCreated = 'GAME_CREATED',
  roomsGamesNotification = 'ROOMS_GAMES_NOTIFICATION',
  gameState = 'GAME_STATE'
}

export enum EWsMessageTypeOut {
  message = 'MESSAGE',
  broadcast = 'BROADCAST',
  connectToRoom = 'CONNECT_TO_ROOM',
  createRoom = 'CREATE_ROOM',
  sendToRoom = 'SEND_TO_ROOM',
  disconnectFromRoom = 'DISCONNECT_FROM_ROOM',
  createGame = 'CREATE_GAME',
  selectGame = 'SELECT_GAME',
}

export interface IwebSocketMessageOut {
  type: EWsMessageTypeOut
  content: Record<string, any>
}

export interface IwebSocketMessageIn {
  type: EWsMessageTypeIn
  content: IWebSocketMessageContent
}

export interface IWebSocketMessageContent {
  message: string
  userEmail: string
  userId: string
}

export interface WSMessage {
  type: string
  content: {
    message: string
    userEmail?: string
    userId?: string
  }
}

export interface IRoom {
  name: string
  id: string
}

export interface IGame {
  name: string
  id: string
  playerNumber: number
}

export interface IGameState {
  bait: IPosition
  players: [IPlayerState, IPlayerState]
}

export interface IPlayerState {
  score: number
  positions: IPosition[]
}

export interface IPosition {
  x: number
  y: number
}
