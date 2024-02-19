import { type Vector3 } from "three";

export interface IBoardSettings {
  cylinder: {
    radialSegment: number;
  };
  mesh: {
    metalness: number;
    roughness: number;
  };
}

// const displayBlocks = (size: number): boolean[][] => {
//   const array: boolean[][] = [];
//   for (let x = 0; x < size; x++) {
//     array.push([]);
//     for (let y = 0; y < size; y++) {
//       array[x].push(true);
//     }
//   }
//   return array;
// };

const toRad = (angle: number): number => {
  return (angle * Math.PI) / 180;
};

export const getNewPosition = (
  currentPosition: Vector3,
  speed: number,
  direction: number
): [number, number, number] => {
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

export const generateBoardSetings = (gameSize: number): IBoardSettings[][] => {
  console.log("=> generate : ", gameSize);
  const res = Array(gameSize)
    .fill(false)
    .map((line, i) =>
      Array(gameSize)
        .fill(false)
        .map((dot, j) => ({
          cylinder: {
            radialSegment: 5 + Math.random() * 3
          },
          mesh: {
            metalness: 0.85 + Math.random() * 0.15,
            roughness: 0.85 + Math.random() * 0.15
          }
        }))
    );
  console.log("=> generate : ", gameSize, res);
  return res;
};
