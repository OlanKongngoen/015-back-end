package main

import (
  EmployeeController "backend/api/controller/employee"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  //Employee API Method
  router.GET("/employee", EmployeeController.GETEmployee)  //GET
  router.POST("/employee", EmployeeController.POSTEmployee) //POST
  router.PUT("/employee", EmployeeController.PUTEmployee)  //PUT
  router.DELETE("/employee", EmployeeController.DELETEEmployee) //DELETE

  router.GET("/employee/:id", EmployeeController.GETEmployeeByID)  //GET By ID
  router.POST("/employee/:id", EmployeeController.POSTEmployeeByID)  //POST By ID
  router.PUT("/employee/:id", EmployeeController.PUTEmployeeByID)  //PUT By ID
  router.DELETE("/employee/:id", EmployeeController.DELETEEmployeeByID)  //DELETE By ID

  //Customer API Method
  router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}