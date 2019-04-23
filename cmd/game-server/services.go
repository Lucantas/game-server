package main

// Hub type represents the hub of connections
type Hub struct {
	ID             string `json:"id"`
	Enabled        bool   `json:"enabled,string"`
	Clients        []Client
	MaxConnections int `json:"maxConnections,string"`
}

//
