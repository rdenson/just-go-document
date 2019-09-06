package model


// maker represents a manufacturer struct
//
// a maker is the model for the constructor endpoints
//
// swagger:model maker
type Manufacturer struct {
  // the id for this maker
  // required: false
  Id int `json:"id"`
  // maker's name
  // required: true
  Name string `json:"name" binding:"required"`
  // maker's country of origin
  // required: true
  Country string `json:"country" binding:"required"`
}

type Manufacturers map[int]*Manufacturer

// swagger:response orderedMakers
type ManufacturerList []*Manufacturer
