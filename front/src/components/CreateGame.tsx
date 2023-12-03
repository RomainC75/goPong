import { Button, TextField } from '@mui/material'
import { ChangeEvent, useContext, useState } from 'react'
import { SocketContextInterface } from '../@types/socketContext.type'
import { SocketContext } from '../context/socket.context'

const CreateGame = () => {
    const { createGame } = useContext(
        SocketContext
    ) as SocketContextInterface
    const [name, setName] = useState<string>('')

    const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
      setName(e.target.value)
    }
    
    const handleSendMessage = (): void => {
        createGame(name)
        setName('')
    }

    return (
    <div className='CreateGame'>
        <div>
        <TextField
            id='outlined-basic'
            label='create game'
            variant='outlined'
            value={name}
            onChange={handleChange}
            size='small'
        />
        <Button variant='contained' onClick={handleSendMessage}>
            Create Game
        </Button>
        </div>
    </div>
    )
}

export default CreateGame