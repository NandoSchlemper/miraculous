package server

import (
	"fmt"
	"miraculous/cmd/database"
	"miraculous/cmd/domain/containers"
	"net/http"
)

type HttpServer struct {
	addr   string
	server *http.ServeMux
}

func NewHttpServer(addr string, server *http.ServeMux) *HttpServer {
	return &HttpServer{
		addr:   addr,
		server: server,
	}
}

func (h *HttpServer) AddHandlers(db *database.Database) {
	authHandler := containers.AuthContainer(db)
	h.server.HandleFunc("POST /api/auth/register", authHandler.Register)
	fmt.Printf("Handlers configurados com sucesso . . .\n")
}

func (h *HttpServer) Serve() {
	fmt.Printf("Server rodando na porta:[ %s ] 👀🚀\n", h.addr)
	http.ListenAndServe(h.addr, h.server)
}
