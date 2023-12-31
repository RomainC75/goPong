import React, { useContext, useEffect, useState } from 'react';
import { SocketContext } from '../context/socket.context';
import { SocketContextInterface } from '../@types/socketContext.type';
import { CurrencyRuble, Games, GamesRounded } from '@mui/icons-material';
import './styles/game.scss';
import { IGridDot } from '../@types/socket.type';
import { initGrid } from '../utils/gameGrid';


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

  useEffect(()=>{
    console.log("=> inside Game : ", grid)
  }, [grid])

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
      <ul className='grid '>
        {grid.map((lines, i) => (
          <li key={'line' + i}>
            { lines.map((dot, j) => <div key={'dot' + i + j} className={'dot ' + dot.color}></div>) }
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Game;
