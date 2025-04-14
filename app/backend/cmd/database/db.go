package database

import (
	"fmt"
	d "miraculous/cmd/domain"

	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{
		db: db,
	}
}

func (data *Database) Migrate() {
	data.db.Raw("CREATE EXTENSION IF NOT EXISTS 'uuid-ossp';")
	// extensão para gerar um UUID
	fmt.Println("Extensão 'uuid-ossp' adicionada ao DB . . .")
	data.db.AutoMigrate(d.User{}, d.UserAttributes{}, d.UserHistory{}, d.EnygmaAnswers{}, d.Enygma{}, d.Task{}, d.Tests{})
	fmt.Println("Migração das tabelas realizada . . .")
}
