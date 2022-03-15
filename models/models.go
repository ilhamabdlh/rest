package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataOn struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Division  DataOnBod `json:"PT Indodev Niaga Internet,omitempty" bson:"PT Indodev Niaga Internet,omitempty"`
}
type DataOnBod struct{
	DivisionTwo  []DataOnIt `json:"BOD00001. Board of Directors" bson:"BOD00001. Board of Directors"`
}
type DataOnIt struct{
	DivisionThree  []DataOnErp `json:"DVS00001. Information Technology" bson:"DVS00001. Information Technology"`
	DivisionFour  DataOnErp `json:"DVS00003. Marketing and Sales" bson:"DVS00001. Information Technology"`
	DivisionFive  []DataOnErp `json:"SPRJT0001. Special Project" bson:"DVS00001. Information Technology"`
	DivisionSix  DataOnErp `json:"division" bson:"division"`
}

type DataOnErp struct{
	DivisionSeven  string `json:"division" bson:"division"`
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
