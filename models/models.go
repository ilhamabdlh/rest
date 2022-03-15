package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataOn struct {
	Division  dataOnBod `json:"division,omitempty" bson:"division,omitempty"`
}

type DataOnBod struct{
	DivisionTwo  []dataOnIt `json:"division,omitempty" bson:"division,omitempty"`
}
type DataOnIt struct{
	DivisionThree  []dataOnErp `json:"division,omitempty" bson:"division,omitempty"`
}

type DataOnErp struct{
	DivisionFour  string `json:"division,omitempty" bson:"division,omitempty"`
}












//Create Struct
// type dataOn struct {
// 	IndoDev  dataOnBod `json:"PT. Indodev Niaga Internet,omitempty" bson:"PT. Indodev Niaga Internet,omitempty"`
// }

// type dataOnBod  struct {
// 	Bod []dataOnIt `json:"BOD00001 Board of Directors,omitempty" bson:"BOD00001 Board of Directors,omitempty"`
// }

// type dataOnIt  struct {
// 	It string `json:"DVS00001 Information Technology,omitempty" bson:"DVS00001 Information Technology,omitempty"`
// 	Marketing string `json:"DVS00003 Marketing and Sales,omitempty" bson:"DVS00003 Marketing and Sales,omitempty"`
// 	Purc string `json:"PURC. Purchasing,omitempty" bson:"PURC. Purchasing,omitempty"`
// 	Finance string `json:"DVS00015 Finance,omitempty" bson:"DVS00015 Finance,omitempty"`
// 	Special string `json:"SPRJT0001. Special Project,omitempty" bson:"SPRJT0001. Special Project,omitempty"`
// }

// type dataOnErp struct{}
