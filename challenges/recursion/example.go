package main

type Server struct {
	Address string
	Name    string
}

func NewServer(arr ...Option) *Server {
	s := &Server{}
	for _, v := range arr {
		v(s)
	}
	return s
}

type Option func(*Server)

func WithAddress(addr string) Option {
	return func(s *Server) {
		s.Address = addr
	}
}
