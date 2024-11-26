package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.GET("/employee", getEmployee)
  router.POST("/employee", postEmployee)
  router.PUT("/employee", putEmployee)
  router.DELETE("/employee", deleteEmployee)

  router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "Employee GET Method!",
    })
  }

func postEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Post Method",
    })
  }

func putEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Put Method",
    })
  }

func deleteEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Delete Method",
    })
  }
