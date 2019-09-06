package api
import (
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/rdenson/g-serv/model"
)


// swagger:route GET /constructor/{id} constructor GetConstructor
//
// displays a manufacturer referenced by id
//
//     Produces:
//     - application/json
//
//     Schemes: http
//     Responses:
//       200: maker
//       400: serverMessage
//
// swagger:operation GET /constructor/{id} constructor GetConstructor
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: id of the maker
//   type: string
//   required: true
// responses:
//   '200':
//     description: a maker
//     schema:
//       "$ref": "#/definitions/maker"
//   '400':
//     description: error message from the server
//     schema:
//       "$ref": "#/responses/serverMessage"
func GetConstructor(c *gin.Context) {
  cid, _ := strconv.Atoi(c.Param("id"))
  if cid < 100 || cid > ds.idx {
    c.JSON(http.StatusBadRequest, NewServerMessage("no data could be found with id", nil))
    return
  }

  c.JSON(http.StatusOK, ds.data[cid])
}

// swagger:route GET /constructor constructor GetConstructors
//
// returns an ordered list (by id) of known makers
//
//     Produces:
//     - application/json
//
//     Schemes: http
//     Responses:
//       200: orderedMakers
//
// swagger:operation GET /constructor constructor GetConstructors
// ---
// produces:
// - application/json
// parameters:
// responses:
//   '200':
//     description: an ordered list of makers
//     schema:
//       type: array
//       items:
//         "$ref": "#/responses/orderedMakers"
func GetConstructors(c *gin.Context) {
  var constructors model.ManufacturerList
  //output constructors in id order
  for i:=100; i<=ds.idx; i++ {
    constructors = append(constructors, ds.data[i])
  }

  c.JSON(http.StatusOK, constructors)
}

// swagger:route POST /constructor constructor AddConstructor
//
// creates a new maker
//
//     Produces:
//     - application/json
//
//     Schemes: http
//     Responses:
//       201: serverMessage
//       400: serverMessage
//
// swagger:operation POST /constructor constructor AddConstructor
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
// responses:
//   '201':
//     description: success message
//     schema:
//       "$ref": "#/responses/serverMessage"
//   '400':
//     description: error message from the server
//     schema:
//       "$ref": "#/responses/serverMessage"
func AddConstructor(c *gin.Context) {
  var newMaker model.Manufacturer

  if err := c.ShouldBindJSON(&newMaker); err != nil {
    c.JSON(http.StatusBadRequest, NewServerMessage("error while attempting to add a new constructor", err))
    return
  }

  ds.add(newMaker.Name, newMaker.Country)
  c.JSON(http.StatusCreated, NewServerMessage("successfully created new manufacturer", nil))
}
