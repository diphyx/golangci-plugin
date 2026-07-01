// Package constructornaming contains a constructor with a non-New name.
package constructornaming

// Server is a sample server.
type Server struct{}

// ProvideServer returns a *Server but is not named NewServer.
func ProvideServer() *Server {
	return &Server{}
}
