package db

import (
    "fmt"
    "log"

    "kuro/app/models"
	"github.com/jackc/pgx/v5"
)

func Migration() {
    err := godotenv.Load(".env")
    if err != nil {
        return nil, fmt.Errorf("Error loading .env file: %v", err)
    }

    db_URL := os.Getenv("DB_URL")
    //define table relationships
}
