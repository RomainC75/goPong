import { useContext, useRef } from "react";
import { SocketContext } from "../../context/socket.context";
import { SocketContextInterface } from "../../@types/socketContext.type";
import { IPosition } from "../../@types/socket.type";
import { useFrame } from "@react-three/fiber";
import { TorusGeometry } from "three";

interface IBait {
    position: IPosition;
}

const Bait = ({position}: IBait) =>{
    const baitRef = useRef< TorusGeometry| null>(null)
    useFrame((state, delta)=>{
        console.log("baitRef", baitRef.current)
        baitRef.current?.rotateX(0.2 * delta)
        baitRef.current?.rotateY(0.5 * delta)
    })
    return (
        <mesh position={[position.x+0.4, position.y+0.6, -3]}>
            <torusGeometry args={[0.2, 0.2, 26, 10]} ref={baitRef}/>
            <meshStandardMaterial color="yellow" />
        </mesh>
    );
}

export default Bait;