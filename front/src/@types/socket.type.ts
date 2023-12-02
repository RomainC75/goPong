export enum EWsMessageTypeIn {
  message = 'MESSAGE',
  idAssigned = 'IDASSIGNED',
  memberJoin = 'MEMBERJOIN',
  memberLeave = 'MEMBERLEAVE',
  broadcast = 'BROADCAST',
  roomCreated = 'ROOM_CREATED',
  roomCreatedByYou = 'ROOM_CREATED_BYYOU',
  roomMessage = 'ROOM_MESSAGE'
}

export enum EWsMessageTypeOut {
  message = 'MESSAGE',
  broadcast = 'BROADCAST',
  connectToRoom = 'CONNECT_TO_ROOM',
  createRoom = 'CREATE_ROOM',
  sendToRoom = 'SEND_TO_ROOM',
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
