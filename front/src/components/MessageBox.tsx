import { useContext } from 'react'
import { type IWebSocketMessageContent } from '../@types/socket.type'
import { AuthContext } from '../context/auth.context'
import { type AuthContextInterface } from '../@types/authContext.type'
import './styles/messageBox.scss'

interface IMessageBox {
  messages: IWebSocketMessageContent[]
}

const MessageBox = ({ messages }: IMessageBox) => {
  const { user } = useContext(
    AuthContext
  ) as AuthContextInterface

  return <div className='MessageBox'>
    <ul>
        {messages.map((message, i) => <li key={ message.userId + i + message.message } className={message.userEmail === user?.email ? 'sent' : 'received'}>{message.message}</li>)}
    </ul>
  </div>
}

export default MessageBox
