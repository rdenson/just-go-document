package api
import (
  "time"

  "github.com/rdenson/g-serv2/common"
  "github.com/rdenson/g-serv2/model"
)


//mock datastore; only accessible in this package
var data = model.Manufacturers{
  &model.Manufacturer{Id: 100, Name: "ford", Country: "usa"},
  &model.Manufacturer{Id: 101, Name: "lamborghini", Country: "italy"},
  &model.Manufacturer{Id: 102, Name: "subaru", Country: "japan"},
}

// swagger:response serverMessage
type ServerMessage struct {
  Msg string `json:"message"`
  Ts time.Time `json:"timestamp"`
}
func (sm *ServerMessage) Serialize() []byte {
  return common.GenericSerialize(sm)
}
func NewServerMessage(s string) *ServerMessage {
  sm := &ServerMessage{
    Msg: s,
    Ts: time.Now(),
  }

  return sm
}
