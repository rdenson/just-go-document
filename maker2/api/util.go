package api
import (
  "time"

  "github.com/rdenson/g-serv/model"
)


//mock datastore; only accessible in this package
var ds = &datastore{
  data: model.Manufacturers{
    100: &model.Manufacturer{Id: 100, Name: "ford", Country: "usa"},
    101: &model.Manufacturer{Id: 101, Name: "lamborghini", Country: "italy"},
    102: &model.Manufacturer{Id: 102, Name: "subaru", Country: "japan"},
  },
  idx: 102,
}

// swagger:response serverMessage
type ServerMessage struct {
  Err string `json:"error,omitempty"`
  Msg string `json:"message"`
  Ts time.Time `json:"timestamp"`
}

func NewServerMessage(s string, e error) *ServerMessage {
  sm := &ServerMessage{
    Msg: s,
    Ts: time.Now(),
  }

  if e != nil {
    sm.Err = e.Error()
  }

  return sm
}
