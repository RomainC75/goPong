import {
  type PropsWithChildren,
  createContext,
  useContext,
  useReducer
} from "react";
import { type IGameState } from "../@types/socket.type";
import { type IGameContextInterface } from "../@types/game.type";
import { initialGameState } from "../utils/initialGameState";

const GameContext = createContext<IGameContextInterface | null>(null);

function gameReducer (state: IGameState, action: any): IGameState {
  console.log("=> gameReducer : ", action, " / ", state, " /")
  switch (action.type) {
    case "update":
      return action.payload;
    default:
      return action.payload;
  }
}

// writing
export const GameProvider = ({ children }: PropsWithChildren): JSX.Element => {
  const [gameState, dispatch] = useReducer(gameReducer, initialGameState);
  const { bait, players } = gameState;
  const gameSize = gameState.game_config.size;

  return (
    <GameContext.Provider
      value={{
        dispatch,
        bait,
        players,
        gameState,
        gameSize
      }}
    >
      {children}
    </GameContext.Provider>
  );
};

// reading
export const useGame = (): IGameContextInterface => {
  const context = useContext(GameContext);
  if (!context) {
    throw new Error("useGame must be used within a GameProvider");
  }
  return context;
};
