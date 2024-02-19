import { Cloud } from "@react-three/drei";
import React, { useEffect, useRef, useState } from "react";
import { getNewPosition } from "../../utils/board.helper";
import { useFrame } from "@react-three/fiber";

interface IDecorations {
  gameSize: number;
}

const Decorations = ({ gameSize }: IDecorations): JSX.Element => {
  const [cloud1Direction, setCloud1Direction] = useState<number>(315);
  const [cloud2Direction, setCloud2Direction] = useState<number>(225);
  const cloud1Ref = useRef<THREE.Group<THREE.Object3DEventMap> | null>(null);
  const cloud2Ref = useRef<THREE.Group<THREE.Object3DEventMap> | null>(null);

  useEffect(() => {
    console.log("=> decorations ");
  }, []);

  useFrame((state, delta) => {
    if (cloud1Ref.current?.position) {
      const position = cloud1Ref.current.position;
      if (position.x > gameSize || position.y > gameSize) {
        setCloud1Direction(135);
      } else if (position.x <= 0 || position.y <= 0) {
        setCloud1Direction(315);
      }
      cloud1Ref.current.position.set(
        ...getNewPosition(position, delta, cloud1Direction)
      );
    }
    if (cloud2Ref.current?.position) {
      const position = cloud2Ref.current.position;
      if (position.x < 0 || position.y > gameSize) {
        setCloud2Direction(45);
      } else if (position.x > gameSize || position.y <= 0) {
        setCloud2Direction(225);
      }
      console.log("=> direction : ", cloud2Direction);
      cloud2Ref.current.position.set(
        ...getNewPosition(position, 1.5 * delta, cloud2Direction)
      );
    }
  });
  return (
    <>
      <Cloud speed={5} segments={40} opacity={0.05} ref={cloud1Ref} />
      <Cloud
        speed={7}
        segments={50}
        opacity={0.05}
        ref={cloud2Ref}
        position={[30, 0, -3]}
        color={"#fea837"}
      />
    </>
  );
};

export default Decorations;
