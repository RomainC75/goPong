import { useContext, useEffect, useRef, useState } from "react";
import { IPlayerState } from "../../@types/socket.type";
import { useFrame } from "@react-three/fiber";
import {
  BufferGeometry,
  Material,
  Mesh,
  NormalBufferAttributes,
  Object3DEventMap,
  SphereGeometry,
} from "three";
import { SocketContext } from "../../context/socket.context";
import { SocketContextInterface } from "../../@types/socketContext.type";

interface ISnake {
  playerNumber: number;
}

// const getSize = (normalSize: number, index: number, shift: number) => {
//   // console.log()
//   return normalSize + Math.cos((index + shift) * 4) * 0.1;
// };

const Snake = ({ playerNumber }: ISnake) => {
  const { gameState } = useContext(SocketContext) as SocketContextInterface;
  const shiftRef = useRef<SphereGeometry>(null);
  const icosaRef = useRef<Mesh<
    BufferGeometry<NormalBufferAttributes>,
    Material | Material[],
    Object3DEventMap
  > | null>(null);

  // const refArray: Array<Mesh<BufferGeometry<NormalBufferAttributes>> | null> =
  //   [];
  // const buffRef = useRef(null);

  

  useFrame((state, delta) => {
    // setShift(state.clock.getElapsedTime());
    if (shiftRef.current) {
      console.log("==> ", shiftRef.current);
      // shift.current.radius =
    }

    if (icosaRef.current == null) return;
    icosaRef.current.rotation.x += 0.2 * delta;
    icosaRef.current.rotation.y += 0.05 * delta;

    if (gameState) {
      icosaRef.current.position.x = gameState.players[playerNumber].positions[0].x;
      icosaRef.current.position.y = gameState.players[playerNumber].positions[0].y;
    }
  });
  return (
    <>
      <mesh
        position={[0, 0, -0.6]}
        ref={icosaRef}
      >
        {/* <sphereGeometry args={[0.8, 5, 5]} /> */}
        <icosahedronGeometry args={[0.8, 0]} />
        <meshStandardMaterial
          color={playerNumber == 0 ? "red" : "green"}
          metalness={1}
          roughness={1}
        />
      </mesh>

      
    </>
  );
};

export default Snake;


// {snake.positions.slice(1).map((dot, i) => (
//   <mesh key={`${i}`} position={[dot.x, dot.y, -0.7]} ref={buffRef[i + 1]}>
//     {/* <sphereGeometry args={[getSize(0.6, i, shift), 5, 5]} ref={shiftRef} /> */}
//     <sphereGeometry args={[0.6, 5, 5]} ref={shiftRef} />
//     <meshStandardMaterial
//       color={playerNumber == 0 ? "blue" : "yellow"}
//       metalness={1}
//       roughness={1}
//     />
//   </mesh>
// ))}