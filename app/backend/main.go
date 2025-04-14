package main

import (
	"miraculous/cmd/database"
	"miraculous/cmd/server/server"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	mux := http.NewServeMux()
	dsn := "host=localhost port=5432 user=miraculous password=1234 dbname=miraculous_db sslmode=disable"
	data, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("erro ao conectar ao database")
	}
	db := database.NewDatabase(data)
	db.Migrate()
	server := server.NewHttpServer(":3030", mux)
	server.AddHandlers(db)
	server.Serve()
}
