package server

func (s *Server) routing() {
	s.echo.POST("/login", s.Login.Login)

	s.echo.POST("/user", s.User.RegistUser)
}