import { Canvas, useFrame, useThree } from "@react-three/fiber";
import React, { useEffect, useRef } from "react";
import { Board, type IBoard } from "./Board";
import { Vector3 } from "three";
import { Perf } from 'r3f-perf'

const Scene = () => {
    
  return (
    <Canvas
    //   camera={{
        // position: [state.game_config.size / 2, state.game_config.size / 2, 20],
        // rotation: [(180 * Math.PI) / 180, (90 * Math.PI) / 180, 10],
        // rotation: [100, 100, 0],
    //   }}
    //   orthographic
    >
        <Perf position='top-left'/>
      <Board/>
    </Canvas>
  );
};

export default Scene
