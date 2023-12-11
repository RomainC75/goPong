import React, { useContext } from 'react'
import { SocketContext } from '../context/socket.context'
import { SocketContextInterface } from '../@types/socketContext.type'
import { CurrencyRuble, Games, GamesRounded } from '@mui/icons-material'

const Game = () => {
    const { currentGame, gameState } = useContext(
    SocketContext
    ) as SocketContextInterface

  return (
    <div>
        <h3>Game </h3>
        <div>
            <p>name : {currentGame?.name}</p>
            <p>id : {currentGame?.id}</p>
            <p>player number : {currentGame?.playerNumber}</p>
        </div>
        <div>
            GAME CONTENT : {gameState?.bait.x}
        </div>
    </div>
  )
}

export default Game