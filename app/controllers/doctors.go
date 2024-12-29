package controllers

import (
	"context"
	"kuro/db"
	"kuro/app/models"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

    query := `Select HospitalID
        From Hospitals
        Where Name = $1
    `
    var HospitalID int
    err = conn.QueryRow(context.Background(), query, input.Hospital).Scan(&HospitalID)
    if err != nil {
        //hospital doesn't exists so I create it.
        query = `Insert into Hospital (name) Values ($1) Returning id`
        err = conn.QueryRow(context.Background(), query, input.Hospital).Scan(&input.Hospital)
        if err != nil {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hospital", "Reason": err.Error()})
            return
        }
    }
    query = `
        Insert into doctors (Name, LastName, Hospital) Values($1, $2, $3) Returning id`
        var doctorID int
        err = conn.QueryRow(context.Background(), query, input.Name, input.LastName, input.Hospital).Scan(&doctorID)
        if err != nil {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create doctor account", "error": err.Error()})
            return
        }

        c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully created account", "doctor_id": doctorID})
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
