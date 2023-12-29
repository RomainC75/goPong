import {
  useState,
  createContext,
  type PropsWithChildren,
  useEffect,
  useContext,
  useMemo,
} from "react";
import { type SocketContextInterface } from "../@types/socketContext.type";
import {
  EWsMessageTypeIn,
  EWsMessageTypeOut,
  type IWebSocketMessageContent,
  type IRoom,
  type IwebSocketMessageIn,
  type IwebSocketMessageOut,
  type IGame,
  type IGameState,
  type IGameConfig,
  type IGridDot,
} from "../@types/socket.type";
import useWebSocket from "react-use-websocket";
import { AuthContext } from "./auth.context";
import { type AuthContextInterface } from "../@types/authContext.type";
import { initGrid, refreshGrid } from "../utils/gameGrid";

const SocketContext = createContext<SocketContextInterface | null>(null);

const SocketProviderWrapper = (props: PropsWithChildren): JSX.Element => {
  const { user } = useContext(AuthContext) as AuthContextInterface;
  const token: string | null = localStorage.getItem("authToken");

  const { sendMessage: sendWsMessage, lastMessage } =
    useWebSocket<IwebSocketMessageIn>(
      `ws://localhost:5000/ws?token=${token ?? ""}`
    );

  const [broadcastMessages, setBroadcastMessages] = useState<
    IWebSocketMessageContent[]
  >([]);
  const [room, setRoom] = useState<IRoom | null>(null);
  const [currentGame, setCurrentGame] = useState<IGame | null>(null);
  const [currentGameConfig, setCurrentGameConfig] =
    useState<IGameConfig | null>(null);
  const [availableRoomList, setAvailableRoomList] = useState<IRoom[]>([]);
  const [availableGameList, setAvailableGameList] = useState<IGame[]>([]);
  const [roomMessages, setRoomMessages] = useState<IWebSocketMessageContent[]>(
    []
  );
  const [pointsState, setPointsState] = useState<[number, number]>([0, 0]);
  const memoPoints = useMemo(() => {
    return pointsState;
  }, [pointsState]);

  const [gameState, setGameState] = useState<IGameState | null>(null);
  const [grid, setGrid] = useState<IGridDot[][]>([]);

  useEffect(() => {
    const state: IGameState = JSON.parse(
      '{"bait":{"x":0,"y":1},"players":[{"score":0,"positions":[{"x":11,"y":10},{"x":11,"y":11},{"x":11,"y":12}],"direction":1},{"score":0,"positions":[{"x":19,"y":20},{"x":19,"y":19},{"x":19,"y":18}],"direction":3}],"level":1,"game_config":{"size":30,"speed_ms":1000},"last_command":[0,0]}'
    );
    setGameState(state);
    setPointsState([state.players[0].score, state.players[1].score]);
  }, []);

  const sendBroadcastMessage = (message: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.broadcast,
      content: {
        message,
      },
    };
    sendWsMessage(JSON.stringify(msg));
  };

  const connectToRoom = (roomId: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.connectToRoom,
      content: {
        roomId,
      },
    };
    sendWsMessage(JSON.stringify(msg));
  };

  const createRoom = (roomName: string): void => {
    console.log("=> click ", roomName);
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.createRoom,
      content: {
        roomName,
      },
    };
    sendWsMessage(JSON.stringify(msg));
  };

  const sendToRoom = (message: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.sendToRoom,
      content: {
        message,
      },
    };
    sendWsMessage(JSON.stringify(msg));
  };

  const disconnectFromRoom = (): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.disconnectFromRoom,
      content: {
        userId: user?.id,
        userEmail: user?.email,
      },
    };
    sendWsMessage(JSON.stringify(msg));
  };

  const createGame = (name: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.createGame,
      content: {
        gameName: name,
      },
    };
    sendWsMessage(JSON.stringify(msg));
  };

  const selectGame = (id: string): void => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.selectGame,
      content: {
        gameId: id,
      },
    };
    sendWsMessage(JSON.stringify(msg));
  };

  const sendKeyCode = (key: number) => {
    const msg: IwebSocketMessageOut = {
      type: EWsMessageTypeOut.gameCommande,
      content: {
        command: `${key}`,
      },
    };
    console.log("=> key : ", key);
    sendWsMessage(JSON.stringify(msg));
  };

  useEffect(() => {
    if (lastMessage !== null) {
      console.log("=> last message : ", lastMessage);
      const message = JSON.parse(lastMessage.data);
      switch (message.type) {
        case EWsMessageTypeIn.broadcast:
          console.log("=> message BRROADCAST received : ", message.content);
          setBroadcastMessages((messages) => [...messages, message.content]);
          break;
        case EWsMessageTypeIn.roomCreatedByYou:
          setRoom(message.content);
          break;
        case EWsMessageTypeIn.connectedToRoom:
          setRoom(message.content);
          break;
        case EWsMessageTypeIn.roomCreated:
          setAvailableRoomList((roomList) => [...roomList, message.content]);
          break;
        case EWsMessageTypeIn.roomMessage:
          setRoomMessages((roomMessages) => [...roomMessages, message.content]);
          break;
        case EWsMessageTypeIn.disconnectedFromRoom:
          setRoomMessages([]);
          setRoom(null);
          break;
        case EWsMessageTypeIn.userDisconnectedFromRoom:
          console.log("=> user disconnected ! ", message.content);
          break;
        case EWsMessageTypeIn.gameCreatedByYou:
          console.log("=> created successfully!");
          setCurrentGame({
            id: message.content.id,
            name: message.content.name,
            playerNumber: 0,
          });
          break;

        case EWsMessageTypeIn.gameCreated:
          console.log("=> created ", message.content);
          setAvailableGameList((gameList) => [...gameList, message.content]);
          break;
        case EWsMessageTypeIn.gameConfigBroadCast:
          console.log("=>CONFIG ", message.content);
          // eslint-disable-next-line no-case-declarations
          const tempCurrentGameConfig: IGameConfig = JSON.parse(
            message.content.config
          ) as IGameConfig;
          setCurrentGameConfig(tempCurrentGameConfig);
          setGrid(initGrid(tempCurrentGameConfig));
          break;
        case EWsMessageTypeIn.gameState:
          // eslint-disable-next-line no-case-declarations
          const newState = JSON.parse(message.content.state);

          setGameState(newState);
          console.log("=> state : ", newState);
          console.log("=> JSON : ", JSON.stringify(message.content.state));

          setPointsState(
            newState.players.map((p: { score: number }) => p.score)
          );

          // eslint-disable-next-line no-case-declarations
          const tempgrid = refreshGrid(grid, newState);
          console.log("=> gameState : ", tempgrid);

          setGrid(tempgrid ?? []);

          break;
        case EWsMessageTypeIn.roomsGamesNotification:
          console.log("=> NOTIFICATION : ", message.content);
          setAvailableRoomList(JSON.parse(message.content.rooms));
          setAvailableGameList(JSON.parse(message.content.games));
          break;
      }
    }
  }, [lastMessage]);

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
        setCurrentGame,
        currentGameConfig,
        sendKeyCode,
        grid,
        memoPoints,
      }}
    >
      {props.children}
    </SocketContext.Provider>
  );
};

export { SocketContext, SocketProviderWrapper };
