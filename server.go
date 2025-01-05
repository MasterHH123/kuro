package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kuro/app/controllers"
	"kuro/db"
)

func main(){
    db.Migration()
    router := gin.Default()
    router.LoadHTMLGlob("client/*.html")

    router.GET("/", index)
    router.GET("/HTMXTest", HTMXTestRoute)
    router.GET("/db_test", db.DBTestHandler)
    router.GET("/create_doctor", createDoctor)
    router.POST("/create_doctor", controllers.CreateDoctor)
    router.GET("/search_hospitals", controllers.SearchHospital)


    router.Run("localhost:8080")
}

func index (c *gin.Context){
    response := gin.H{"message": "kuro's index page"}
    c.HTML(http.StatusOK, "index.html", response)
}

func createDoctor(c *gin.Context) {
    response := gin.H{"message": "create doctor route reached"}
    c.HTML(http.StatusOK, "create_doctor.html", response)
}

func HTMXTestRoute (c *gin.Context)  {
    response := "Welcome to kuro!!! We're using htmx to build this website!!!!"
    c.IndentedJSON(http.StatusOK, response)
}
