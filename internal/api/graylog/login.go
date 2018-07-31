package graylog

const (
	loginAPI = "/system/sessions"
)

// InvalidLoginTicket represents a ticket from un unsuccessful login attempt
var InvalidLoginTicket = LoginTicket{
	SessionID:  "N/A",
	ValidUntil: "N/A",
}

// LoginTicket describes the login ticket from a successful login call
type LoginTicket struct {
	SessionID  string `json:"session_id"`
	ValidUntil string `json:"valid_until"`
}

type loginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}


