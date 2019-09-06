package api
import (
  "github.com/rdenson/g-serv/model"
)


type datastore struct {
  data model.Manufacturers
  idx int
}
func (s *datastore) add(name, origin string) {
  newMaker := &model.Manufacturer{
    Name: name,
    Country: origin,
  }

  s.idx++
  newMaker.Id = s.idx
  s.data[s.idx] = newMaker
}
