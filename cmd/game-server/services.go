package main

// Hub type represents the hub of connections
type Hub struct {
	ID             string `json:"id"`
	Enabled        bool   `json:"enabled,string"`
	Clients        []Client
	MaxConnections int64 `json:"maxConnections,string"`
	GameMaxPlayers int
}

func newHub(c ...Client) Hub {
	return Hub{Clients: c}
}

func (h Hub) add(c Client) {
	h.Clients = append(h.Clients, c)
	if len(h.Clients) >= h.GameMaxPlayers {
		// enough players to start a match
		games = append(games, newMatch(h.Clients))
	}
}

// Match type represents the game between two or more players
type Match struct {
	ID      int64 `json:"id"`
	Clients []Client
	Status  string
}

func newMatch(c []Client) Match {
	return Match{Clients: c}
}
