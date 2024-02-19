import { useEffect, useRef, useState } from "react";
import { useFrame, useThree } from "@react-three/fiber";
import { Cloud, PerspectiveCamera, Sky } from "@react-three/drei";
import { Vector3 } from "three";
import type * as THREE from "three";
import Snake from "./Snake";
import Bait from "./Bait";
import { IBoardSettings, generateBoardSetings, getNewPosition } from "../../utils/board.helper";

interface IBoard {
  gameSize: number
}


const Board = ({ gameSize }: IBoard): JSX.Element => {

  const { camera } = useThree();
  const spot1Ref = useRef<THREE.DirectionalLight | null>(null);
  const spot2Ref = useRef<THREE.DirectionalLight | null>(null);

  const [boardSettings, setBoardSettings] = useState<IBoardSettings[][]>([]);


  useEffect(() => {
    setBoardSettings(generateBoardSetings(gameSize))
    console.log("=> gamesize : ", gameSize)
  }, [])

  useEffect(() => {
    console.log('App rendered : gameSize');
  }, [gameSize]);

  useEffect(() => {
    console.log("==> GAMESIZE : ", gameSize)
  }, [gameSize])


  //   light1Intensity,
  //   light2Position,
  //   light2Intensity,
  // } = useControls({
  //   camDirection: {
  //     value: [Math.PI, 0, 0],
  //     step: 0.1,
  //     min: -2 * Math.PI,
  //     max: 2 * Math.PI,
  //   },
  //   camPosition: {
  //     value: [15, 15, -36],
  //     step: 0.1,
  //     min: -40,
  //     max: 40,
  //   },
  //   light1Position: {
  //     value: [-1.5, -1.5, -3],
  //     step: 0.1,
  //     min: -40,
  //     max: 40,
  //   },
  //   light1Intensity: {
  //     value: 5,
  //     step: 0.1,
  //     min: 0,
  //     max: 40,
  //   },
  //   light1Color: {
  //     value: "#ff8686",
  //     onChange: (v) => {
  //       if (spot1Ref?.current?.color) {
  //         spot1Ref.current.color = new THREE.Color(v);
  //       }
  //     },
  //   },
  //   light2Position: {
  //     value: [35, 35, -3],
  //     step: 0.1,
  //     min: -40,>
  //     max: 40,
  //   },
  //   light2Intensity: {
  //     value: 5,
  //     step: 0.1,
  //     min: 0,
  //     max: 40,
  //   },
  //   light2Color: {
  //     value: "#6e60ff",
  //     onChange: (v) => {
  //       if (spot2Ref?.current?.color) {
  //         spot2Ref.current.color = new THREE.Color(v);
  //       }
  //     },
  //   },
  // });

  useEffect(() => {
    camera.lookAt(5, 1, 0);
  }, []);

  return (
    <>
      {/* <OrbitControls position={[state.game_config.size / 2, state.game_config.size / 2, 5]} rotation={[(180 * Math.PI) / 180, (90 * Math.PI) / 180, 10]}/> */}
      <ambientLight intensity={0.2} position={new Vector3(30, 1, -5)} />
      <directionalLight
        // position={light1Position}
        position={[-1.5, -1.5, -3]}
        ref={spot1Ref}
        intensity={5}
        color={"#ff8686"}
      />
      <directionalLight
        // position={light2Position}
        position={[35, 35, -3]}
        ref={spot2Ref}
        intensity={5}
      />
      <directionalLight position={[15, 15, -10]} intensity={1} />
      
      <Sky
        distance={450000}
        sunPosition={[0, 1, 0]}
        inclination={0}
        azimuth={0.25}
      />
      // TODO
      { boardSettings.map((line, i) =>
        line.map((tile, j) => (
          <mesh key={`${i},${j}`} position={[i, j, 0]} rotation-x={Math.PI / 2}>
            {/* <boxGeometry args={[0.9, 0.9, 0.1]} /> */}
            <cylinderGeometry args={[0.7, 0.3, 0.2, tile.cylinder.radialSegment]} />
            <meshStandardMaterial
              color={i === 0 && j === 0 ? "red" : "grey"}
              metalness={tile.mesh.metalness}
              roughness={tile.mesh.roughness}
            />
          </mesh>
        ))
      )}

      <Snake playerNumber={0} />
      <Snake playerNumber={1} />
      <Bait /> */

      <PerspectiveCamera
        makeDefault
        position={[15, 15, -36]}
        rotation={[Math.PI, 0, 0]}
        // position={camPosition}
        // rotation={camDirection}
      />
    </>
  );
};

export { type IBoard, Board };
