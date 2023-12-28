import React from 'react'
import { IPlayerState } from '../../@types/socket.type'

interface ISnake{
    snake: IPlayerState
    playerNumber: number
}

const Snake = ({snake, playerNumber}: ISnake) => {
  return (
    <>
        <mesh position={[snake.positions[0].x, snake.positions[0].y, 0.6]}>
            <sphereGeometry args={[0.8, 32,32]} />
            <meshStandardMaterial color={playerNumber==0 ? "red" : "green"} />
        </mesh>
        {snake.positions.slice(1).map((dot, i) => (
            <mesh key={`${i}`} position={[dot.x, dot.y, 0.7]}>
                <sphereGeometry args={[0.6, 32,32]} />
                <meshStandardMaterial color={playerNumber==0 ? "blue" : "yellow"} />
            </mesh>
        ))}
    </>
  )
}

export default Snake