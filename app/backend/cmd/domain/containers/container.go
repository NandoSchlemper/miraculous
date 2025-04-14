package containers

import (
	"fmt"
	"miraculous/cmd/database"
	"miraculous/cmd/server/handlers"
)

func AuthContainer(db *database.Database) handlers.IAuthHandler {
	authHandler := handlers.NewAuthHandler()
	fmt.Println("Handler de AUTENTICAÇÃO cadastrado . . .")
	return authHandler
}
