import { useRef } from "react";
import { useFrame } from "@react-three/fiber";
import {
  BufferGeometry,
  Material,
  Mesh,
  NormalBufferAttributes,
  Object3DEventMap,
  TorusGeometry,
} from "three";
import { useGame } from "../../hooks/useGame";

const Bait = () => {
  const { bait } = useGame();
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

    if (baitPositionRef.current) {
      baitPositionRef.current.position.x = bait.x + 0.4;
      baitPositionRef.current.position.y = bait.y + 0.6;
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
