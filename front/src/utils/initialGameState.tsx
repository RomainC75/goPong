import { type IGameState } from "../@types/socket.type";

export const initialGameState: IGameState = {
  bait: { x: 0, y: 0 },
  players: [
    {
      score: 0,
      direction: 1,
      positions: [
        { x: 11, y: 14 },
        { x: 11, y: 15 },
        { x: 11, y: 16 }
      ]
    },
    {
      score: 0,
      direction: 3,
      positions: [
        { x: 19, y: 16 },
        { x: 19, y: 15 },
        { x: 19, y: 14 }
      ]
    }
  ],
  lever: 0,
  game_config: {
    size: 30,
    speed_ms: 0
  },
  lastCommands: [],
};
