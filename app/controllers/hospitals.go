package controllers

import (
	"context"
	"fmt"
	"kuro/app/models"
	"kuro/db"

	"github.com/google/uuid"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHospitalForm(c *gin.Context) {
    c.HTML(http.StatusOK, "add_hospital_form.html", nil)
} 

func CreateHospital(c *gin.Context) {
    var input models.Hospital
    if err := c.ShouldBind(&input); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    HospitalValues := map[string]string{
        "Hospital Name": input.Name,
        "Hospital Address": input.Address,
        "City": input.City,
    }
    for k, v := range HospitalValues {
        if v == "" {
            c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Input values cannot be empty."})
            return
        }
        fmt.Println(k, v)
    }

    conn, err := db.DBConnection()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database connection unsuccessful"})
    }
    defer conn.Close(context.Background())

    HospitalID := uuid.New()
    query := `
        Insert Into Hospitals (HospitalID, Name, Address, City) Values($1, $2, $3, $4)
    `
    _, err = conn.Exec(context.Background(), query, HospitalID, input.Name, input.Address, input.City)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create hospital", "error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully created hospital", "hospital_id": HospitalID})
}

func SearchHospital(c *gin.Context) {
    conn, err := db.DBConnection()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database connection failed.", "error": err.Error()})

    }
    defer conn.Close(context.Background())

    searchQuery := c.Query("hospital")
    if searchQuery == "" {
        c.String(http.StatusBadRequest, "Hospital cannot be empty")
        return
    }
    fmt.Println("Value of searchQuery: %s\n", searchQuery)

    query := `Select HospitalID, Name
        From Hospitals
        Where Name ILIKE $1
    `
    rows, err := conn.Query(context.Background(), query, "%"+searchQuery+"%")
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to search hospitals"})
        return
    }
    defer rows.Close()

    var results []gin.H
    for rows.Next() {
        var hospitalID string
        var hospitalName string
        if err := rows.Scan(&hospitalID, &hospitalName); err != nil {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse", "error": err.Error()})
            return
        }
        results = append(results, gin.H{"id": hospitalID, "name": hospitalName})
    }

    // Flag to add a hospital
    if len(results) == 0 {
        c.HTML(http.StatusOK, "add_hospital_button.html", gin.H{
            "results": results,
            "add_hospital": true,
        })
        return
    }
    c.HTML(http.StatusOK, "hospital_results.html", gin.H{"results": results})

}

func SelectHospital(c *gin.Context) {
    hospitalID := c.Param("id")
    conn, err := db.DBConnection()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, "Database connection failed")
        return
    }
    defer conn.Close(context.Background())

    fmt.Println("HospitalID: %s", hospitalID)
    query := `Select Name From Hospitals Where HospitalID = $1`
    var hospitalName string
    err = conn.QueryRow(context.Background(), query, hospitalID).Scan(&hospitalName)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, "Hospital not found")
        return
    }

    c.HTML(http.StatusOK, "selected_hospital.html", gin.H{"hospital": hospitalName})
}

