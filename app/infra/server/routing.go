package server

func (s *Server) routing() {
	s.echo.POST("/login", s.Login.Login)

	s.echo.GET("/user", s.User.GetUser)
	s.echo.POST("/user", s.User.RegistUser)
}