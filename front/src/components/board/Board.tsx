import React, { useEffect } from "react";
import { Canvas, useFrame, useThree } from "@react-three/fiber";
import { OrbitControls, OrthographicCamera, PerspectiveCamera } from "@react-three/drei";
import { IGameState } from "../../@types/socket.type";
import Snake from "./Snake";
import * as THREE from "three";
import { useControls } from "leva";

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

const Board = ({ state }: IBoard) => {
  const { camera } = useThree();

  useEffect(() => {
    console.log("=>",state.game_config.size/2)
    // camera.position.set(state.game_config.size / 2, state.game_config.size / 2, 1);
    // camera.position.set(state.game_config.size / 2, state.game_config.size / 2, 10);
    camera.lookAt(5,1,0);
  }, []);

  return (
    <>
      {/* <OrbitControls position={[state.game_config.size / 2, state.game_config.size / 2, 5]} rotation={[(180 * Math.PI) / 180, (90 * Math.PI) / 180, 10]}/> */}
      {/* <ambientLight /> */}

      <directionalLight position={[10, 10, 10]} />
      {displayBlocks(state).map((line, i) =>
        line.map((dot, j) => (
          <mesh key={`${i},${j}`} position={[i, j, 0]}>
            <boxGeometry args={[0.9, 0.9, 0.1]} />
            <meshStandardMaterial color="grey" />
          </mesh>
        ))
      )}

      <Snake snake={state.players[0]} playerNumber={0} />
      <Snake snake={state.players[1]} playerNumber={1}/>
      <PerspectiveCamera makeDefault position={[15,15,35]} rotation={[
        Math.PI,
        Math.PI,
        Math.PI,
      ]}/>
    </>
  );
};

export { type IBoard, Board };
