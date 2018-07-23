package api

// invalidLoginTicket represents a ticket from un unsuccessful login attempt
var invalidLoginTicket = loginTicket{
	SessionID:  "N/A",
	ValidUntil: "N/A",
}

// loginTicket describes the login ticket from a successful login call
type loginTicket struct {
	SessionID  string `json:"session_id"`
	ValidUntil string `json:"valid_until"`
}

type session struct {
	ConnectionString string
	loginTicket
}
