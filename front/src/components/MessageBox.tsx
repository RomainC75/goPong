import { useContext, useEffect } from 'react'
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

  useEffect(()=>{
    console.log("=> mmessages : ", messages)
  }, [messages])

  return <div className='MessageBox'>
    <ul>
        {messages.map((message, i) => 
        <li key={message.userId + i + message.message} className={message.userEmail === user?.email ? 'sent' : 'received'}>
          {message.userEmail !== user?.email && <p className="name">{message.userEmail}</p>}
          <p className={message.userEmail === user?.email ? 'sent bubble' : 'received bubble'}>{message.message}</p>
        </li>
        )}
    </ul>
  </div>
}

export default MessageBox
