package commons

import "go.mongodb.org/mongo-driver/bson/primitive"

type SysInfo struct {
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	Createdat  primitive.DateTime `json:"createdat,omitempty" bson:"createdat,omitempty"`
	Modifiedat primitive.DateTime `json:"modifiedat,omitempty" bson:"modifiedat,omitempty"`
}

func (s SysInfo) IsZero() bool {
	/*
	   if s.Status == "" {
	       return false
	   }
	   if s.Createdat == 0 {
	       return false
	   }
	   if s.Modifiedat == 0 {
	       return false
	   }
	       return true
	*/
	return s.Status == "" && s.Createdat == 0 && s.Modifiedat == 0
}
