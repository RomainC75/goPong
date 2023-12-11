import { useState, createContext, type PropsWithChildren, useEffect, useContext } from 'react'
import { type SocketContextInterface } from '../@types/socketContext.type'
import {
  EWsMessageTypeIn,
  EWsMessageTypeOut,
  IWebSocketMessageContent,
  type IRoom,
  type IwebSocketMessageIn,
  type IwebSocketMessageOut,
  IGame,
  IGameState
} from '../@types/socket.type'
import useWebSocket from 'react-use-websocket'
import { AuthContext } from './auth.context'
import { AuthContextInterface } from '../@types/authContext.type'

const SocketContext = createContext<SocketContextInterface | null>(null)

const SocketProviderWrapper = (props: PropsWithChildren): JSX.Element => {
  const { user } = useContext(
    AuthContext
  ) as AuthContextInterface
  const token: string | null = localStorage.getItem('authToken')

  const { sendMessage: sendWsMessage, lastMessage } =
    useWebSocket<IwebSocketMessageIn>(`ws://localhost:5000/ws?token=${token ?? ''}`)

  const [broadcastMessages, setBroadcastMessages] = useState<IWebSocketMessageContent[]>([])
  const [room, setRoom] = useState<IRoom | null>(null)
  const [currentGame, setCurrentGame] = useState<IGame | null>(null)
  const [availableRoomList, setAvailableRoomList] = useState<IRoom[]>([])
  const [availableGameList, setAvailableGameList] = useState<IGame[]>([])
  const [roomMessages, setRoomMessages] = useState<IWebSocketMessageContent[]>([])
  const [gameState, setGameState] = useState<IGameState | null>(null)
  // const [broadcastmessage, setBroadcastMessage] = useState<string>('')

  const sendBroadcastMessage = (message: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.broadcast,
      content: {
        message
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const connectToRoom = (roomId: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.connectToRoom,
      content: {
        roomId
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const createRoom = (roomName: string): void => {
    console.log('=> click ', roomName)
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.createRoom,
      content: {
        roomName
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const sendToRoom = (message: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.sendToRoom,
      content: {
        message
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const disconnectFromRoom = (): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.disconnectFromRoom,
      content:{
        userId: user?.id,
        userEmail: user?.email
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const createGame = (name: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.createGame,
      content: {
        gameName: name
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  const selectGame = (id: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.selectGame,
      content: {
        gameId: id
      }
    }
    sendWsMessage(JSON.stringify(msg))
  }

  useEffect(() => {
    if (lastMessage !== null) {
      console.log('=> last message : ', lastMessage)
      const message = JSON.parse(lastMessage.data)
      switch (message.type) {
        case EWsMessageTypeIn.broadcast:
          console.log("=> message BRROADCAST received : ", message.content)
          setBroadcastMessages(messages => ([...messages, message.content]))
          break
        case EWsMessageTypeIn.roomCreatedByYou:
          setRoom(message.content)
          break
        case EWsMessageTypeIn.connectedToRoom:
          setRoom(message.content)
          break
        case EWsMessageTypeIn.roomCreated:
          setAvailableRoomList(roomList => ([...roomList, message.content]))
          break
        case EWsMessageTypeIn.roomMessage:
          setRoomMessages(roomMessages => ([...roomMessages, message.content]))
          break
        case EWsMessageTypeIn.disconnectedFromRoom:
          setRoomMessages([])
          setRoom(null)
          break
        case EWsMessageTypeIn.userDisconnectedFromRoom:
          console.log('=> user disconnected ! ', message.content)
          break
        case EWsMessageTypeIn.gameCreatedByYou:
          console.log('=> created successfully!')
          setCurrentGame({
            id: message.content.id,
            name: message.content.name,
            playerNumber: 0
          })
          break
        case EWsMessageTypeIn.gameCreated:
          console.log('=> created ', message.content)
          setAvailableGameList(gameList => [...gameList, message.content])
          break
        case EWsMessageTypeIn.gameState:
          console.log('=> gameState : ', JSON.parse(message.content.state))
          setGameState(JSON.parse(message.content.state))
          break
        case EWsMessageTypeIn.roomsGamesNotification:
          console.log('=> NOTIFICATION : ', message.content)
          setAvailableRoomList(JSON.parse(message.content.rooms))
          setAvailableGameList(JSON.parse(message.content.games))
          break
      }
    }
  }, [lastMessage])

  return (
    <SocketContext.Provider
      value={{
        sendBroadcastMessage,
        lastMessage,
        createRoom,
        broadcastMessages,
        room,
        availableRoomList,
        availableGameList,
        sendToRoom,
        roomMessages,
        connectToRoom,
        disconnectFromRoom,
        createGame,
        selectGame,
        currentGame,
        gameState,
        setCurrentGame
      }}
    >
      {props.children}
    </SocketContext.Provider>
  )
}

export { SocketContext, SocketProviderWrapper }
