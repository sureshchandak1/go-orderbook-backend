package database

import "github.com/sureshchandak1/go-orderbook-backend/internal/database/tables"

func SyncDatabase() {
	DB.AutoMigrate(&tables.User{})
}
