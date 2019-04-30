package main

// Payload type represents the default object to be sent
// and received along the server's life
type Payload struct {
	Method  string                 `json:"method"`
	MatchID int64                  `json:"matchId,string"`
	Params  map[string]interface{} `json:"params,string"`
}
