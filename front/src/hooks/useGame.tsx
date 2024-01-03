import { type PropsWithChildren, createContext, useContext, useReducer } from "react";
import { SocketContextInterface } from "../@types/socketContext.type";
import { type IGameState } from "../@types/socket.type";
import { IGameContextInterface } from "../@types/game.type";

const GameContext = createContext<IGameContextInterface | null>(null);

function gameReducer (state: IGameState, action: any): IGameState {
  switch (action.type) {
    case "update":
      return state;
    default :
      return state;
  };
}

// writing
export const GameProvider = ({ children }: PropsWithChildren): JSX.Element => {
  const [gameState, dispatch] = useReducer(gameReducer,null);
  const { bait, players } = gameState;

  return (
    <GameContext.Provider
      value={{
        dispatch,
        bait,
        players
      }}
    >
      {children}
    </GameContext.Provider>
  );
};

// reading
export const useGame = (): IGameContextInterface | null => {
  const context = useContext(GameContext);
  if (context === undefined) {
    throw new Error("useGame must be used within a GameProvider");
  }
  return context;
};
