package main

import (
	"net/http"

	"kuro/db"
	"github.com/gin-gonic/gin"
)

func main(){
    router := gin.Default()
    router.GET("/", index)
    router.GET("/db_test", db.DBTestHandler)


    router.Run("localhost:8080")
}

func index (c *gin.Context){
    response := gin.H{"message": "Welcome to kuro's first route!!!"}
    c.IndentedJSON(http.StatusOK, response)
}


