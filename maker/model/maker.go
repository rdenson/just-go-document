package model
import (
  "github.com/rdenson/g-serv2/common"
)


// manufacturer represents a maker
//
// a maker is the model for the constructor endpoints
//
// swagger:model
type Manufacturer struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Country string `json:"country"`
}
func (m Manufacturer) Serialize() []byte {
  return common.GenericSerialize(m)
}

// swagger:response manyConstructors
type Manufacturers []*Manufacturer
func (m Manufacturers) Serialize() []byte {
  mb := []byte("[")

  for i:=0; i<len(m); i++ {
    mb = append(mb, m[i].Serialize()...)
    if i < (len(m)-1) {
      mb = append(mb, ',')
    }
  }

  mb = append(mb, ']')

  return mb
}
