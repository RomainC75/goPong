import React, { useContext, useEffect, useState } from "react";
import { SocketContext } from "../context/socket.context";
import { SocketContextInterface } from "../@types/socketContext.type";
import { CurrencyRuble, Games, GamesRounded } from "@mui/icons-material";
import "./styles/game.scss";
interface IGridDot {
  color: string | undefined;
}

const Game = () => {
  const { currentGame, currentGameConfig, gameState } = useContext(
    SocketContext
  ) as SocketContextInterface;
  const [grid, setGrid] = useState<IGridDot[][]>([]);

  useEffect(() => {
    if (currentGameConfig) {
      const tempGrid: IGridDot[][] = [];
      for (let i = 0; i < currentGameConfig?.size; i++) {
        tempGrid.push([]);
        for (let j = 0; j < currentGameConfig?.size; j++) {
          tempGrid[i].push({ color: "none" });
        }
      }
      setGrid(tempGrid);
      console.log("=> init : ", tempGrid);
    }
  }, [currentGameConfig]);

  useEffect(() => {
    if (currentGameConfig && gameState && grid.length) {
      const tempGrid = grid;
      for (let line = 0; line < currentGameConfig?.size; line++) {
        for (let column = 0; column < currentGameConfig?.size; column++) {
            console.log("=>>>>>>>>>>>>>>>>", gameState.bait, line, column)
          if (gameState?.bait.x == column && gameState.bait.y == line) {
            tempGrid[line][column] = {
              color: "red",
            };
          } else {
            tempGrid[line][column] = {
              color: "none",
            };
          }
        }
      }
      console.log("=> refresh", tempGrid, gameState, currentGameConfig);
      setGrid(tempGrid);
    }
  }, [gameState, currentGameConfig]);

  return (
    <div className="Game">
      <h3>Game </h3>
      <div>
        <p>name : {currentGame?.name}</p>
        <p>id : {currentGame?.id}</p>
        <p>player number : {currentGame?.playerNumber}</p>
      </div>
      <ul className="grid">
        {grid.map((line, i) => (
          <li key={"line" + i}>
            {line.map((dot, j) => (
              <div key={"dot" + i + j} className={"dot " + dot.color}></div>
            ))}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Game;
