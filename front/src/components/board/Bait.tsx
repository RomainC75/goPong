import { useContext, useRef } from "react";
import { SocketContext } from "../../context/socket.context";
import { SocketContextInterface } from "../../@types/socketContext.type";
import { IPosition } from "../../@types/socket.type";
import { useFrame } from "@react-three/fiber";
import {
  BufferGeometry,
  Material,
  Mesh,
  NormalBufferAttributes,
  Object3DEventMap,
  TorusGeometry,
} from "three";

const Bait = () => {
  const { gameState } = useContext(SocketContext) as SocketContextInterface;
  const baitRef = useRef<TorusGeometry | null>(null);
  const baitPositionRef = useRef<Mesh<
    BufferGeometry<NormalBufferAttributes>,
    Material | Material[],
    Object3DEventMap
  > | null>(null);

  useFrame((state, delta) => {
    console.log("baitRef", baitRef.current);
    baitRef.current?.rotateX(0.2 * delta);
    baitRef.current?.rotateY(0.5 * delta);

    if (gameState && baitPositionRef.current) {
      baitPositionRef.current.position.x = gameState.bait.x + 0.4;
      baitPositionRef.current.position.y = gameState.bait.y + 0.6;
    }
  });
  return (
    <mesh position={[0, 0, -3]} ref={baitPositionRef}>
      <torusGeometry args={[0.2, 0.2, 26, 10]} ref={baitRef} />
      <meshStandardMaterial color="yellow" />
    </mesh>
  );
};

export default Bait;
