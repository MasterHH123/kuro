package kuro

import (
	"kuro/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDoctorByID(c *gin.Context){
    idParam := c.Param("id")
    id, err := strconv.Atoi((idParam))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
        return
    }

    var doctor models.Doctor
    if err := db.DB.First(&doctor, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
        return
    }

    c.JSON(http.StatusOK, doctor)

}
