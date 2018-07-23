package server

// Definition describes the data used to define a server
type Definition struct {
	Connection

	User
}

// Connection describes only the connection part
type Connection struct {
	Scheme string `description:"the connection scheme" type:"select" choice:"http|https"`

	Host string `description:"the hostname"`

	Port string `description:"the port"`

	Endpoint string `description:"the API endpoint"`
}

// User defines the data used for a user
type User struct {
	Username string `description:"the username"`

	Password string `description:"the password" type:"password"`
}

// Defaults return an instance of Definition with sensible default values
func Defaults() Definition {
	return Definition{

		DefaultConnection(),

		DefaultUser(),
	}
}

// DefaultConnection creates an instance of Connection with sensible default values
func DefaultConnection() Connection {
	return Connection{

		Host: "server.domain.com",

		Port: "9000",

		Endpoint: "/api",

		Scheme: "http",
	}
}

// DefaultUser creates an instance of User with sensible default values
func DefaultUser() User {
	return User{

		Username: "user",

		Password: "pass",
	}
}
