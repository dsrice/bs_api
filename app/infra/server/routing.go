package server

func (s *Server) routing() {
	s.echo.POST("/login", s.Login.Login)
}