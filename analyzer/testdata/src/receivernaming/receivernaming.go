package receivernaming

type Server struct{}

func (server *Server) Start() {}

func (s *Server) Stop() {} // want "receiver 's' for type 'Server' should be 'server'"
