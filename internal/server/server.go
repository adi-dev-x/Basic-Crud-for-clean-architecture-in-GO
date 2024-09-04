package server

import "github.com/labstack/echo"

type ServerStruct struct {
	Engine *echo.Echo
}

func (s *ServerStruct) StartServer() {
	s.Engine.Start("localhost:8080")
}

func NewHTTPServer() *ServerStruct {
	e := echo.New()

	return &ServerStruct{
		Engine: e,
	}
}
