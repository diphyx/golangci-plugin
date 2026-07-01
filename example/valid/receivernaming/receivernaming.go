// Package receivernaming shows a receiver using the camelCase full type name.
package receivernaming

// Server is a sample server.
type Server struct {
	name string
}

// Start starts the server.
func (server *Server) Start() {
	_ = server.name
}
