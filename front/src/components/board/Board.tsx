import React, { useContext, useEffect } from "react";
import { Canvas, useFrame, useThree } from "@react-three/fiber";
import {
  OrbitControls,
  OrthographicCamera,
  PerspectiveCamera,
} from "@react-three/drei";
import { IGameState } from "../../@types/socket.type";
import Snake from "./Snake";
import * as THREE from "three";
import { useControls } from "leva";
import { SocketContextInterface } from "../../@types/socketContext.type";
import { SocketContext } from "../../context/socket.context";

interface IBoard {
  state: IGameState;
}

const displayBlocks = (state: IGameState): boolean[][] => {
  const array: boolean[][] = [];
  for (let x = 0; x < state.game_config.size; x++) {
    array.push([]);
    for (let y = 0; y < state.game_config.size; y++) {
      array[x].push(true);
    }
  }
  return array;
};

const Board = () => {
  const { gameState } = useContext(SocketContext) as SocketContextInterface;

  const { camera } = useThree();

  useEffect(() => {
    camera.lookAt(5, 1, 0);
  }, []);
  
  useEffect(()=>{
    console.log("=>games state : ", gameState);
  }, [gameState])

  return (
    <>
      {gameState && (
        <>
          {/* <OrbitControls position={[state.game_config.size / 2, state.game_config.size / 2, 5]} rotation={[(180 * Math.PI) / 180, (90 * Math.PI) / 180, 10]}/> */}
          {/* <ambientLight /> */}

          <directionalLight position={[10, 10, 10]} />
          {displayBlocks(gameState).map((line, i) =>
            line.map((dot, j) => (
              <mesh key={`${i},${j}`} position={[i, j, 0]}>
                <boxGeometry args={[0.9, 0.9, 0.1]} />
                <meshStandardMaterial color="grey" />
              </mesh>
            ))
          )}

          <Snake snake={gameState.players[0]} playerNumber={0} />
          <Snake snake={gameState.players[1]} playerNumber={1} />
          <PerspectiveCamera
            makeDefault
            position={[15, 15, 35]}
            rotation={[Math.PI, Math.PI, Math.PI]}
          />
        </>
      )}
    </>
  );
};

export { type IBoard, Board };
