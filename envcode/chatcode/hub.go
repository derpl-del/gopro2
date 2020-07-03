package chatcode

//Message a
type Message struct {
	data []byte
	room string
}

//Subscription con
type Subscription struct {
	conn *Client
	room string
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	rooms map[string]map[*Client]bool
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan Subscription

	// Unregister requests from clients.
	unregister chan Subscription
}

//StructH data
var StructH = Hub{
	broadcast:  make(chan Message),
	register:   make(chan Subscription),
	unregister: make(chan Subscription),
	rooms:      make(map[string]map[*Client]bool),
}

//Run a
func (Hreq *Hub) Run() {
	for {
		select {
		case s := <-Hreq.register:
			connections := Hreq.rooms[s.room]
			if connections == nil {
				connections = make(map[*Client]bool)
				Hreq.rooms[s.room] = connections
			}
			Hreq.rooms[s.room][s.conn] = true
		case s := <-Hreq.unregister:
			connections := Hreq.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(Hreq.rooms, s.room)
					}
				}
			}
		case m := <-Hreq.broadcast:
			connections := Hreq.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(Hreq.rooms, m.room)
					}
				}
			}
		}
	}
}
