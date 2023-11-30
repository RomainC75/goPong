export enum EWsMessageTypeReceiver {
  message = 'MESSAGE',
  idAssigned = 'IDASSIGNED',
  memberJoin = 'MEMBERJOIN',
  memberLeave = 'MEMBERLEAVE',
  broadcast = 'BROADCAST',
  roomCreated = 'ROOM_CREATED',
}
export enum EWsMessageTypeSender {
  message = 'MESSAGE',
  broadcast = 'BRAODCAST',
  connectToRoom = 'CONNECT_TO_ROOM',
  createRoom = 'CREATE_ROOM'
}



export interface webSocketMessage {
  type: EWsMessageType
  content: Record<string, any>
}


