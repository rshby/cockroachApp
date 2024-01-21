package server

type Server interface {
	RunServer() error
	AddRouter()
}
