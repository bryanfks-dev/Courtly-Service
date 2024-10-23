package config

import (
	"log"
	"main/pkg/utils"
	"strconv"
)

// ServerConfig is a struct that holds the configuration for the server.
type Server struct {
	// Host is the host name the server will listen on.
	Host string

	// Port is the port number the server will listen on.
	Port int
}

// ServerConfig is a global variable that holds the server configuration.
var ServerConfig = Server{}

// LoadData is a method that loads the server configuration from the environment variables.
func (s Server) LoadData() {
	// Get the host name from the environment variables
	host := utils.GetEnv("SERVER_HOST", "localhost")

	s.Host = host

	// Get the port number from the environment variables
	port, err := strconv.Atoi(utils.GetEnv("SERVER_PORT", "8080"))

	// Check if the port number is valid
	if err != nil {
		log.Fatal("Invalid port number")
	}

	s.Port = port

	ServerConfig = s
}
