// Package classification maker API.
//
// this documentation describes a trivial API for research.
// we can hopefully see some patterns for implemention here
//
// Terms Of Service:
// not present
//
//     Schemes: http
//     BasePath: /api/v1
//     Version: 1.0.0
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
  "fmt"
  "net/http"
  //"net/http/httputil"

  "github.com/rdenson/g-serv2/api"
)


func main() {
  const SERVER_PORT = "9228"

  suiHandler := http.FileServer(http.Dir("./swagger-ui"))
  http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", suiHandler))

  http.HandleFunc("/api/v1/constructor", api.ConstructorHandler)
  fmt.Printf("server listening on :%s\n", SERVER_PORT)
  if err := http.ListenAndServe("127.0.0.1:" + SERVER_PORT, nil); err != nil {
    fmt.Println(err.Error())
  }
}
