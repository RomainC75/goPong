import { useContext } from "react"
import { SocketContext } from "../context/socket.context"
import { SocketContextInterface } from "../@types/socketContext.type"
import { IGame } from "../@types/socket.type"

const GameList = (): JSX.Element => {
  const { availableGameList, selectGame, setCurrentGame } = useContext(
    SocketContext
  ) as SocketContextInterface

  const handleSelectGame = (game: IGame): boolean => {
    selectGame(game.id)
    setCurrentGame({
      ...game,
      playerNumber: 1
    })
    return true
  }

  return (
    <div className="GameList">
      <ul>
        {availableGameList.map((game, i) => (
          <li key={game.id + i} onClick={() => handleSelectGame(game)}>
            {game.name}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default GameList
