package controllers

import (
	"context"
	"fmt"
	"kuro/app/models"
	"kuro/db"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateDoctor(c *gin.Context){
    var input models.Doctor
    if err := c.ShouldBind(&input); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    conn, err := db.DBConnection()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database connection unsuccessful", "error": err.Error()})
    }
    defer conn.Close(context.Background())
    
    DoctorID := uuid.New()
    DoctorValues := map[string] string {
        "Doctor Name": input.Name,
        "Doctor Last Name": input.LastName,
        "Hospital": input.Hospital,
    }
    for k, v := range DoctorValues {
        fmt.Printf("%s: %s\n", k, v)
    }
    var HospitalID string
    hospitalIDQuery := `Select HospitalID
        From Hospitals
        Where Name = $1
        `
    err = conn.QueryRow(context.Background(), hospitalIDQuery, input.Hospital).Scan(&HospitalID)
    if err != nil {
        c.String(http.StatusInternalServerError, "error\n", err.Error())
        fmt.Println(http.StatusInternalServerError, "error", err.Error())
    }
    fmt.Println("HospitalID: ", HospitalID)
    query := `
        Insert into doctors (DoctorID, Name, LastName, Hospital) Values($1, $2, $3, $4)
        `
        _, err = conn.Exec(context.Background(), query, DoctorID, input.Name, input.LastName, HospitalID)
        if err != nil {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create doctor account", "error": err.Error()})
            return
        }

        /*
        c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully created account", "doctor_id": DoctorID})
        */
        c.HTML(http.StatusCreated, "doctor_created.html", gin.H{"message": "Successfully created account", "doctor_id": DoctorID})
}

func GetDoctorByID(c *gin.Context){
    idParam := c.Param("id")
    id, err := strconv.Atoi((idParam))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
        return
    }

    conn, err := db.DBConnection()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database connection unsuccessful"})
    }

    query := `Select *
    From Doctors
    Where ID = $1
    `
    var doctor models.Doctor 
    err = conn.QueryRow(context.Background(), query, id).Scan(&doctor.DoctorID, &doctor.Name, &doctor.LastName, &doctor.Hospital)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }


    c.IndentedJSON(http.StatusOK, doctor)

}
