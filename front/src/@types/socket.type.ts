export enum EWsMessageType {
  message = 'MESSAGE',
  idAssigned = 'IDASSIGNED',
  memberJoin = 'MEMBERJOIN',
  memberLeave = 'MEMBERLEAVE',
  broadcast = 'BROADCAST'

}

export interface webSocketMessage {
  type: EWsMessageType
  content: Record<string, any>
}


