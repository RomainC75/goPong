import { useContext, useEffect } from 'react';
import { SocketContext } from '../context/socket.context';
import { SocketContextInterface } from '../@types/socketContext.type';
import './styles/game.scss';
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
      <div style={{height: "800px", width: "800px", border:"1px solid black"}}>
        <Scene/>
      </div>
    </div>
  );
};

export default Game;
