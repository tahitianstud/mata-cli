package graylog

import (
	"fmt"
	"strings"
)

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

func (t LoginTicket) GetSessionID() string {
	return t.SessionID
}

func (t LoginTicket) GetValidUntil() string {
	return t.ValidUntil
}

func (t LoginTicket) Serialized() string {
	return fmt.Sprintf("%s||%s", t.SessionID, t.ValidUntil)
}

func DeSerializeLogin(serializedLoginTicket string) LoginTicket {
	parts := strings.Split(serializedLoginTicket, "||")

	return LoginTicket{
		SessionID: parts[0],
		ValidUntil: parts[1],
	}
}

type loginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}


