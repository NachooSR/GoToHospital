package main

import "github.com/gin-gonic/gin"

var prefijo string ="/api/v1"

func main() {

	
  	
  router := gin.Default()
  router.GET(prefijo+"/example", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "mensaje": "Hola desde Gin GoHospital",
    })
  })
  router.Run() // listen and serve on 0.0.0.0:8080
}