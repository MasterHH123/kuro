package db

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DBConnection() (*pgx.Conn, error) {
    err := godotenv.Load(".env")
    if err != nil {
        return nil, fmt.Errorf("Error loading .env file: %v", err)
    }

    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        return nil, fmt.Errorf("Database URL is not set correctly.\n")
    }

    conn, err := pgx.Connect(context.Background(), dbURL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
    }

    return conn, nil
}

func DBTestHandler(c *gin.Context) {
    conn, err := DBConnection()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Database connection successful!", "Connection": conn})

}



