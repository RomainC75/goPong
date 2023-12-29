import React, { useRef, useState } from "react";
import { IPlayerState } from "../../@types/socket.type";
import { useFrame } from "@react-three/fiber";
import { Mesh } from "three";

interface ISnake {
  snake: IPlayerState;
  playerNumber: number;
}

const getSize = (normalSize: number, index: number, shift: number) => {
    // console.log()
    return normalSize+Math.cos(index+shift)*0.1
}

const Snake = ({ snake, playerNumber }: ISnake) => {
  const [shift, setShift] = useState<number>(0);

  useFrame((state, delta) => {
    setShift(state.clock.getElapsedTime());
  });
  return (
    <>
      <mesh position={[snake.positions[0].x, snake.positions[0].y, -0.6]}>
        <sphereGeometry args={[0.8, 32, 32]} />
        <meshStandardMaterial color={playerNumber == 0 ? "red" : "green"} metalness={1} roughness={1}/>
      </mesh>

      {snake.positions.slice(1).map((dot, i) => (
        <mesh key={`${i}`} position={[dot.x, dot.y, -0.7]}>
          <sphereGeometry args={[getSize(0.6, i, shift), 32, 32]} />
          <meshStandardMaterial color={playerNumber == 0 ? "blue" : "yellow"} metalness={1} roughness={1}/>
        </mesh>
      ))}
    </>
  );
};

export default Snake;
