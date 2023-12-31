package server

func (s *Server) routing() {
	s.echo.GET("/", s.Login.Login)
}