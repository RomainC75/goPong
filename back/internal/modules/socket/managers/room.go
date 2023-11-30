package managers


type RoomList map[*Room]bool

type Room struct{
	Name string
	Manager *Manager
	Clients ClientList
}

func NewRoom(name string, manager *Manager, client *Client) *Room{
	clients := ClientList{}
	clients[client]=true
	return &Room{
		Name: name,
		Manager: manager,
		Clients: clients,
	}
}

func (r *Room)AddClient(client *Client){
	r.Clients[client]=true
}

func (r *Room)RemoveClient(client *Client){
	delete(r.Clients,client)
}

