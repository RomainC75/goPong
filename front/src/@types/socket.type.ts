export enum EWsMessageTypeIn {
  message = 'MESSAGE',
  idAssigned = 'IDASSIGNED',
  memberJoin = 'MEMBERJOIN',
  memberLeave = 'MEMBERLEAVE',
  broadcast = 'BROADCAST',
  roomCreated = 'ROOM_CREATED',
}

export enum EWsMessageTypeOut {
  message = 'MESSAGE',
  broadcast = 'BROADCAST',
  connectToRoom = 'CONNECT_TO_ROOM',
  createRoom = 'CREATE_ROOM',
}

export interface IwebSocketMessageOut {
  type: EWsMessageTypeOut
  content: Record<string, any>
}

export interface IwebSocketMessageIn {
  type: EWsMessageTypeIn
  content: Record<string, any>
}

export interface WSMessage {
  type: string
  content: {
    message: string
    userEmail?: string
    userId?: string
  }
}
