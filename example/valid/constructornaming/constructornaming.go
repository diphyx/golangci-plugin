// Package constructornaming shows a constructor named New{TypeName}.
package constructornaming

// Server is a sample server.
type Server struct {
	name string
}

// NewServer creates a new Server.
func NewServer() *Server {
	return &Server{}
}
