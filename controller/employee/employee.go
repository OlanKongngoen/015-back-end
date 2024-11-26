package employee

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

//GET
func GETEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}

func POSTEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Post Method",
    })
  }

func PUTEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Put Method",
    })
  }

func DELETEEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Employee Delete Method",
    })
  }

  //GET By ID
func GETEmployeeByID(c *gin.Context) {
  id := c.Param("id")
  c.JSON(http.StatusOK, gin.H{
  "GET By ID": id,
  })
}

func POSTEmployeeByID(c *gin.Context) {
  id := c.Param("id")
  c.JSON(http.StatusOK, gin.H{
  "POST By ID": id,
  })
}

func PUTEmployeeByID(c *gin.Context) {
  id := c.Param("id")
  c.JSON(http.StatusOK, gin.H{
  "PUT By ID": id,
  })
}

func DELETEEmployeeByID(c *gin.Context) {
  id := c.Param("id")
  c.JSON(http.StatusOK, gin.H{
  "DELETE By ID": id,
  })
}