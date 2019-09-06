package api
import (
  "encoding/json"
  "io"
  "net/http"

  "github.com/rdenson/g-serv2/model"
)


// swagger:route GET /constructor constructor constructorGet
//
// displays all known manufacturers
//
//     Produces:
//     - application/json
//
//     Schemes: http
//     Responses:
//       200: manyConstructors
//
// swagger:operation GET /constructor constructor constructorGet
// ---
// produces:
// - application/json
// parameters:
// responses:
//   '200':
//     description: an array of manufacturers
//     schema:
//       type: array
//       items:
//         "$ref": "#/responses/manyConstructors"
func constructorGet(w http.ResponseWriter) {
  w.Header().Set("Content-Type", "application/json")
  w.Write(data.Serialize())
}

// swagger:route POST /constructor constructor constructorPost
//
// creates a new manufacturer
//
//     Produces:
//     - application/json
//
//     Schemes: http
//     Responses:
//       201: serverMessage
//
// swagger:operation POST /constructor constructor constructorPost
// ---
// produces:
// - application/json
// parameters:
// - name: maker
//   in: body
//   description: a maker object
//   required: true
//   schema:
//     type: object
//     required:
//     - name
//     - country
//     properties:
//       name:
//         type: string
//       country:
//         type: string
func constructorPost(w http.ResponseWriter, body io.Reader) {
  newMaker := new(model.Manufacturer)
  b := json.NewDecoder(body)

  newMaker.Id = data[len(data)-1].Id + 1
  if decodeErr := b.Decode(newMaker); decodeErr != nil {
    http.Error(w, "{\"m\": \"parse error\",\"e\": \"" + decodeErr.Error() + "\"}", http.StatusInternalServerError)
  }

  data = append(data, newMaker)
  w.WriteHeader(http.StatusCreated)
  w.Write(NewServerMessage("successfully created new manufacturer").Serialize())
}

func ConstructorHandler(w http.ResponseWriter, req *http.Request) {
  switch req.Method {
  case "GET":
    constructorGet(w)
  case "POST":
    constructorPost(w, req.Body)
  }
}
