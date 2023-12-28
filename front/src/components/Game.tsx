import React, { useContext, useEffect, useState } from 'react';
import { SocketContext } from '../context/socket.context';
import { SocketContextInterface } from '../@types/socketContext.type';
import { CurrencyRuble, Games, GamesRounded } from '@mui/icons-material';
import './styles/game.scss';
import { IGame, IGameConfig, IGameState, IGridDot } from '../@types/socket.type';
import { initGrid, refreshGrid } from '../utils/gameGrid';
import Scene from './board/Scene';


const Game = () => {
  const { currentGame, sendKeyCode, grid, memoPoints } = useContext(
    SocketContext
  ) as SocketContextInterface;

  useEffect(() => {
    document.addEventListener('keydown', function (event) {
      const codeArr = [39, 37];
      const index = codeArr.findIndex((code) => code === event.keyCode);
      if (index >= 0) {
        sendKeyCode(index === 0 ? -1 : 1);
      }
    });
  }, []);

  const config: IGameConfig = { size: 30 } as IGameConfig
  const iniGrid = initGrid(config);
  const state: IGameState = JSON.parse("{\"bait\":{\"x\":0,\"y\":1},\"players\":[{\"score\":0,\"positions\":[{\"x\":11,\"y\":10},{\"x\":11,\"y\":11},{\"x\":11,\"y\":12}],\"direction\":1},{\"score\":0,\"positions\":[{\"x\":19,\"y\":20},{\"x\":19,\"y\":19},{\"x\":19,\"y\":18}],\"direction\":3}],\"level\":1,\"game_config\":{\"size\":30,\"speed_ms\":1000},\"last_command\":[0,0]}")
  const fakegrid: IGridDot[][] = refreshGrid(iniGrid, state) as IGridDot[][]
  console.log("=> state : ", state)

  // useEffect(()=>{
  //   console.log("=> inside Game : ", grid)
  // }, [grid])

  return (
    <div className='Game'>
      <h3>Game </h3>
      <div className='infos'>
        <p>name : {currentGame?.name}</p>
        <p>id : {currentGame?.id}</p>
        <p>player number : {currentGame?.playerNumber}</p>
        <div>
          <span className={`point ${currentGame?.playerNumber === 0 && 'me'}`}>{memoPoints[0]} </span>
          <span>-</span>
          <span className={`point ${currentGame?.playerNumber === 1 && 'me'}`}> {memoPoints[1]}</span>
        </div>
      </div>
      {/* <ul className='grid '>
        {fakegrid.map((lines, i) => (
          <li key={'line' + i}>
            { lines.map((dot, j) => <div key={'dot' + i + j} className={'dot ' + dot.color}></div>) }
          </li>
        ))}
      </ul> */}
      <div style={{height: "600px", width: "600px", border:"1px solid black"}}>
        <Scene state={state}/>
      </div>
    </div>
  );
};

export default Game;
