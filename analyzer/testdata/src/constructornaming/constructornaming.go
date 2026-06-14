package constructornaming

type Server struct{}

func NewServer() *Server { return &Server{} }

func BuildServer() *Server { return &Server{} }

func ProvideServer() *Server { return &Server{} } // want "should be named 'NewServer'"
