import { Ref, useContext, useEffect, useRef } from "react";
import { useThree } from "@react-three/fiber";
import { PerspectiveCamera } from "@react-three/drei";
import { IGameState } from "../../@types/socket.type";
import Snake from "./Snake";
import { SocketContextInterface } from "../../@types/socketContext.type";
import { SocketContext } from "../../context/socket.context";
import { useControls } from "leva";
import { Vector3 } from "three";
import * as THREE from "three";

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
  const spot1Ref = useRef<THREE.DirectionalLight | null>(null);
  const spot2Ref = useRef<THREE.DirectionalLight | null>(null);

  const { 
    camDirection, 
    camPosition, 
    light1Position, 
    light1Intensity, 
    light2Position, 
    light2Intensity
  } = useControls({
    camDirection: {
      value: [Math.PI, 0, 0],
      step: 0.1,
      min: -2*Math.PI,
      max: 2*Math.PI,
    },
    camPosition: {
      value: [15, 15, -40],
      step: 0.1,
      min: -40,
      max: 40,
    },
    light1Position: {
      value: [5, 10, -40],
      step: 0.1,
      min: -40,
      max: 40,
    },
    light1Intensity: {
      value: 1,
      step: 0.1,
      min: 0,
      max: 40,
    },
    light1Color: {
      value: 'white',
      onChange: (v) => {
        if(spot1Ref?.current?.color){
          spot1Ref.current.color = new THREE.Color(v)
        }
      },
    },
    light2Position: {
      value: [5, 10, -40],
      step: 0.1,
      min: -40,
      max: 40,
    },
    light2Intensity: {
      value: 1,
      step: 0.1,
      min: 0,
      max: 40,
    },
    light2Color: {
      value: 'white',
      onChange: (v) => {
        if(spot2Ref?.current?.color){
          spot2Ref.current.color = new THREE.Color(v)
        }
      },
    },
  });

  useEffect(() => {
    camera.lookAt(5, 1, 0);
  }, []);

  return (
    <>
      {gameState && (
        <>
          {/* <OrbitControls position={[state.game_config.size / 2, state.game_config.size / 2, 5]} rotation={[(180 * Math.PI) / 180, (90 * Math.PI) / 180, 10]}/> */}
          <ambientLight intensity={0.2} position={new Vector3(5,5,-30)}/>

          <directionalLight position={light1Position} ref={spot1Ref} intensity={light1Intensity}/>
          <directionalLight position={light2Position} ref={spot2Ref} intensity={light2Intensity}/>
          <directionalLight position={[15,15,-10]}  intensity={1}/>

          {displayBlocks(gameState).map((line, i) =>
            line.map((dot, j) => (
              <mesh key={`${i},${j}`} position={[i, j, 0]}>
                <boxGeometry args={[0.9, 0.9, 0.1]} />
                <meshStandardMaterial color={i===0 && j===0 ? "red" : "grey"} metalness={1} roughness={1}/>
              </mesh>
            ))
          )}

          <Snake snake={gameState.players[0]} playerNumber={0} />
          <Snake snake={gameState.players[1]} playerNumber={1} />
          <PerspectiveCamera
            makeDefault
            // position={[15, 15, 35]}
            // rotation={[Math.PI, Math.PI, Math.PI]}
            position={camPosition}
            rotation={camDirection}
          />
        </>
      )}
    </>
  );
};

export { type IBoard, Board };
