import { IGameConfig, IGameState, IGridDot } from "../@types/socket.type";

export const initGrid = (currentGameConfig: IGameConfig): IGridDot[][] => {
  const tempGrid: IGridDot[][] = [];
  for (let i = 0; i < currentGameConfig?.size; i++) {
    tempGrid.push([]);
    for (let j = 0; j < currentGameConfig?.size; j++) {
      tempGrid[i].push({ color: "none" });
    }
  }
  return tempGrid;
};

export const refreshGrid = (
  grid: IGridDot[][],
  gameState: IGameState
): IGridDot[][] | undefined => {
  if (gameState && grid.length) {
    const t = gameState.players.map(p => p.positions[0])
    console.log("==> refreshGrid - gameState : ", t[0], t[1]);
    
    const tempGrid = grid;
    for (let line = 0; line < gameState?.game_config.size; line++) {
      for (let column = 0; column < gameState?.game_config.size; column++) {
        if (column == gameState.bait.x && line == gameState.bait.y) {
          tempGrid[line][column] = {
            color: "red",
          };
        } else if (
          gameState.players[0].positions[0].x == column &&
          gameState.players[0].positions[0].y == line
        ) {
          tempGrid[line][column] = {
            color: "player1-head",
          };
        } else if (
          gameState.players[0].positions.find(
            (p) => p.x == column && p.y == line
          )
        ) {
          tempGrid[line][column] = {
            color: "player1",
          };
        } else if (
          gameState.players[1].positions[0].x == column &&
          gameState.players[1].positions[0].y == line
        ) {
          tempGrid[line][column] = {
            color: "player2-head",
          };
        } else if (
          gameState.players[1].positions.find(
            (p) => p.x == column && p.y == line
          )
        ) {
          tempGrid[line][column] = {
            color: "player2",
          };
        } else {
          tempGrid[line][column] = {
            color: "none",
          };
        }
      }
    }
    // console.log("=> new GRID : ", tempGrid);
    return tempGrid;
  }
};
