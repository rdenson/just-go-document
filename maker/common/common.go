package common
import (
  "encoding/json"
  "fmt"
)


func GenericSerialize(v interface{}) []byte {
  b, marshalErr := json.Marshal(v)
  if marshalErr != nil {
    fmt.Printf("%s\n", marshalErr.Error())
  }

  return b
}
