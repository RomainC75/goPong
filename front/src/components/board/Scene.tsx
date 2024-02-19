import { Canvas, useFrame, useThree } from "@react-three/fiber";
import React, { useContext, useEffect, useRef } from "react";
import { Board, type IBoard } from "./Board";
import { Vector3 } from "three";
import { Perf } from 'r3f-perf'
import { useGame } from "../../hooks/useGame";
import { SocketContext } from "../../context/socket.context";
import { SocketContextInterface } from "../../@types/socketContext.type";
import Decorations from "./Decorations";

const Scene = (): JSX.Element => {
  const { gameSize } = useContext(SocketContext) as SocketContextInterface
  return (
    <Canvas
    //   camera={{
        // position: [state.game_config.size / 2, state.game_config.size / 2, 20],
        // rotation: [(180 * Math.PI) / 180, (90 * Math.PI) / 180, 10],
        // rotation: [100, 100, 0],
    //   }}
      // orthographic
    >
        <Perf position='top-left'/>
      <Board gameSize={gameSize} />
      <Decorations gameSize={gameSize}/>
    </Canvas>
  );
};

export default Scene
