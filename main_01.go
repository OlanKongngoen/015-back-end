package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  //default API Method
  router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "API SERVER STATUS OK!",
    })
  })

  //Employee API Method
  router.GET("/employee", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Get Method",
    })
  })

  router.POST("/employee ", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Post Method",
    })
  })

  router.PUT("/employee", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Put Method",
    })
  })

  router.DELETE("/employee", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Delete Method",
    })
  })

  router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}