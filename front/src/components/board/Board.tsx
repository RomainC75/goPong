import { Ref, useContext, useEffect, useRef } from "react";
import { useFrame, useThree } from "@react-three/fiber";
import { Cloud, PerspectiveCamera, Sky } from "@react-three/drei";
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

const toRad = (angle: number): number => {
  return (angle * Math.PI) / 180;
};

const getNewPosition = (
  currentPosition: Vector3,
  speed: number,
  direction: number
): [number,number,number] => {
  // between 0 andd 2*Pi
  let x = currentPosition.x;
  let y = currentPosition.y;
  if (direction < 90) {
    const rad = toRad(direction);
    y -= Math.sin(rad) * speed;
    x += Math.cos(rad) * speed;
  } else if (direction >= 90 && direction <= 180) {
    const rad = toRad(180 - direction);
    y -= Math.sin(rad) * speed;
    x -= Math.cos(rad) * speed;
  } else if (direction >= 180 && direction <= 270) {
    const rad = toRad(direction - 180);
    y += Math.sin(rad) * speed;
    x -= Math.cos(rad) * speed;
  } else {
    const rad = toRad(360 - direction);
    y += Math.sin(rad) * speed;
    x += Math.cos(rad) * speed;
  }
  return [x, y, currentPosition.z];
};

const Board = () => {
  const { gameState } = useContext(SocketContext) as SocketContextInterface;
  const { camera } = useThree();
  const spot1Ref = useRef<THREE.DirectionalLight | null>(null);
  const spot2Ref = useRef<THREE.DirectionalLight | null>(null);
  const cloudRef = useRef<THREE.Group<THREE.Object3DEventMap> | null>(null);

  let cloudDirection = 315;

  useFrame((state, delta) => {
    if (cloudRef.current?.position) {
      const position = cloudRef.current.position;
      if (position.x > 30 || position.y > 30){
        cloudDirection = 135;
      }else if (position.x <= 0 || position.y <= 0){
        cloudDirection = 315;
      }
      cloudRef.current.position.set(...getNewPosition( position, 0.1, cloudDirection))
    }
  });

  const {
    camDirection,
    camPosition,
    light1Position,
    light1Intensity,
    light2Position,
    light2Intensity,
  } = useControls({
    camDirection: {
      value: [Math.PI, 0, 0],
      step: 0.1,
      min: -2 * Math.PI,
      max: 2 * Math.PI,
    },
    camPosition: {
      value: [15, 15, -40],
      step: 0.1,
      min: -40,
      max: 40,
    },
    light1Position: {
      value: [-1.5, -1.5, -3],
      step: 0.1,
      min: -40,
      max: 40,
    },
    light1Intensity: {
      value: 5,
      step: 0.1,
      min: 0,
      max: 40,
    },
    light1Color: {
      value: "#ff8686",
      onChange: (v) => {
        if (spot1Ref?.current?.color) {
          spot1Ref.current.color = new THREE.Color(v);
        }
      },
    },
    light2Position: {
      value: [35, 35, -3],
      step: 0.1,
      min: -40,
      max: 40,
    },
    light2Intensity: {
      value: 5,
      step: 0.1,
      min: 0,
      max: 40,
    },
    light2Color: {
      value: "#6e60ff",
      onChange: (v) => {
        if (spot2Ref?.current?.color) {
          spot2Ref.current.color = new THREE.Color(v);
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
          <ambientLight intensity={0.2} position={new Vector3(30, 1, -5)} />

          <directionalLight
            position={light1Position}
            ref={spot1Ref}
            intensity={light1Intensity}
          />
          <directionalLight
            position={light2Position}
            ref={spot2Ref}
            intensity={light2Intensity}
          />
          <directionalLight position={[15, 15, -10]} intensity={1} />

          <Cloud speed={0.2} segments={40} opacity={0.05} ref={cloudRef} />
          <Sky
            distance={450000}
            sunPosition={[0, 1, 0]}
            inclination={0}
            azimuth={0.25}
          />

          {displayBlocks(gameState).map((line, i) =>
            line.map((dot, j) => (
              <mesh
                key={`${i},${j}`}
                position={[i, j, 0]}
                rotation-x={Math.PI / 2}
              >
                {/* <boxGeometry args={[0.9, 0.9, 0.1]} /> */}
                <cylinderGeometry
                  args={[0.7, 0.3, 0.2, 5 + Math.random() * 3]}
                />
                <meshStandardMaterial
                  color={i === 0 && j === 0 ? "red" : "grey"}
                  metalness={0.9 + Math.random() * 0.1}
                  roughness={0.9 + Math.random() * 0.1}
                />
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
