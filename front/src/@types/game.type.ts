import { IGameState, IPlayerState, IPosition } from "./socket.type";

export interface IGameContextInterface {
  gameState: IGameState | null;
  bait: IPosition;
  players: [IPlayerState, IPlayerState];
  //   dispatch: (state: IGameState, action: any) => IGameState;
  dispatch: React.Dispatch<IGameState>;
}

export enum EActions {
  Update = "update",
}
