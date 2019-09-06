// Package classification maker API 2.0.
//
// this documentation describes a trivial API for research.
// we can hopefully see some patterns for implemention here
//
// Terms Of Service:
// not present
//
//     Schemes: http
//     BasePath: /api/v1
//     Version: 2.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Host: localhost:9228
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main
import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/rdenson/g-serv/api"
)


func main() {
  router := gin.Default()

  //swagger ui
  router.StaticFS("/swaggerui", http.Dir("./swagger-ui"))

  v1 := router.Group("/api/v1")
  v1.GET("/constructor", api.GetConstructors)
  v1.GET("/constructor/:id", api.GetConstructor)
  v1.POST("/constructor", api.AddConstructor)

  router.Run(":9228");
}
