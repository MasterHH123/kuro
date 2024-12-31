package controllers

import (
	"context"
	"fmt"
	"kuro/db"

	//"kuro/app/models"

	"net/http"

	"github.com/gin-gonic/gin"
)


func SearchHospital(c *gin.Context) {
    conn, err := db.DBConnection()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database connection failed.", "error": err.Error()})

    }
    defer conn.Close(context.Background())

    query := `Select Name
        From Hospitals
    `
    rows, err := conn.Query(context.Background(), query)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "query failed"})
        return
    }
    defer rows.Close()

    var hospitals []string
    for rows.Next() {
        var hospitalName string
        if err := rows.Scan(&hospitalName); err != nil {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse hospital name", "error": err.Error()})
            return
        }
        hospitals = append(hospitals, hospitalName)
    }

    if rows.Err() != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed during iteration"})
        return
    }

    var htmlResponse string
    for _, name := range hospitals {
		htmlResponse += fmt.Sprintf(`<div class="hospital-option">%s</div>`, name)
	}

    c.Header("Content-Type", "text/html")
    c.String(http.StatusOK, htmlResponse)
}
