// import { EWsMessageTypeIn } from "../@types/socket.type";

// export const gameStateHandler = () =>{
//     console.log("=> last message : ", lastMessage);
//       const message = JSON.parse(lastMessage.data);
//       switch (message.type) {
//         case EWsMessageTypeIn.broadcast:
//           console.log("=> message BRROADCAST received : ", message.content);
//           setBroadcastMessages((messages) => [...messages, message.content]);
//           break;
//         case EWsMessageTypeIn.roomCreatedByYou:
//           setRoom(message.content);
//           break;
//         case EWsMessageTypeIn.connectedToRoom:
//           setRoom(message.content);
//           break;
//         case EWsMessageTypeIn.roomCreated:
//           setAvailableRoomList((roomList) => [...roomList, message.content]);
//           break;
//         case EWsMessageTypeIn.roomMessage:
//           setRoomMessages((roomMessages) => [...roomMessages, message.content]);
//           break;
//         case EWsMessageTypeIn.disconnectedFromRoom:
//           setRoomMessages([]);
//           setRoom(null);
//           break;
//         case EWsMessageTypeIn.userDisconnectedFromRoom:
//           console.log("=> user disconnected ! ", message.content);
//           break;
//         //=========================
//         case EWsMessageTypeIn.gameCreatedByYou:
//           console.log("=> created successfully!");
//           setCurrentGame({
//             id: message.content.id,
//             name: message.content.name,
//             playerNumber: 0,
//           });
//           break;

//         case EWsMessageTypeIn.gameCreated:
//           console.log("=> created ", message.content);
//           setAvailableGameList((gameList) => [...gameList, message.content]);
//           break;
//         case EWsMessageTypeIn.gameConfigBroadCast:
//           console.log("=>CONFIG ", message.content);
//           // eslint-disable-next-line no-case-declarations
//           const tempCurrentGameConfig: IGameConfig = JSON.parse(
//             message.content.config
//           ) as IGameConfig;
//           setCurrentGameConfig(tempCurrentGameConfig);
//           setGrid(initGrid(tempCurrentGameConfig));

//           break;
//         //=========================
//         case EWsMessageTypeIn.gameState:
//           // setGameState(JSON.parse(message.content.state))
//           // eslint-disable-next-line no-case-declarations
//           const tempgrid = refreshGrid(grid, JSON.parse(message.content.state));
//           console.log("=> gameState : ", tempgrid);

//           setGrid(tempgrid ?? []);

//           break;
//         case EWsMessageTypeIn.roomsGamesNotification:
//           console.log("=> NOTIFICATION : ", message.content);
//           setAvailableRoomList(JSON.parse(message.content.rooms));
//           setAvailableGameList(JSON.parse(message.content.games));
//           break;
//           }
// }