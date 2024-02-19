import {
  type IGameState,
  type IPlayerState,
  type IPosition
} from "./socket.type";

export interface IGameContextInterface {
  gameState: IGameState | null;
  bait: IPosition;
  players: [IPlayerState, IPlayerState];
  gameSize: number;
  //   dispatch: (state: IGameState, action: any) => IGameState;
  dispatch: React.Dispatch<{ type: EActions, payload: IGameState }>;
}

export enum EActions {
  Update = "update",
}
