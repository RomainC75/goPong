import { useContext } from "react"
import { SocketContext } from "../context/socket.context"
import { SocketContextInterface } from "../@types/socketContext.type"

const GameList = (): JSX.Element => {
  const { availableGameList, selectGame } = useContext(
    SocketContext
  ) as SocketContextInterface

  const handleSelectGame = (id: string): boolean => {
    selectGame(id)
    return true
  }

  return (
    <div className="GameList">
      <ul>
        {availableGameList.map((game, i) => (
          <li key={game.id + i} onClick={() => handleSelectGame(game.id)}>
            {game.name}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default GameList
