package server

import "cockroachApp/app/handler"

type Server interface {
	RunServer() error
	AddRouter(handler handler.CockroachHandler)
}
