package user

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/mario-imperato/r3ds9-mongodb/model/commons"

const (
	OID            = "_id"
	NICKNAME       = "nickname"
	OBJTYPE        = "objType"
	FIRSTNAME      = "firstname"
	LASTNAME       = "lastname"
	EMAIL          = "email"
	PASSWORD       = "password"
	ROLES          = "roles"
	ROLES_I        = "roles.%d"
	SYSINFO        = "sysinfo"
	PROFILEPICTURE = "profilePicture"
)

type User struct {
	OId            primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	Nickname       string                `json:"nickname,omitempty" bson:"nickname,omitempty"`
	ObjType        string                `json:"objType,omitempty" bson:"objType,omitempty"`
	Firstname      string                `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname       string                `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email          string                `json:"email,omitempty" bson:"email,omitempty"`
	Password       string                `json:"password,omitempty" bson:"password,omitempty"`
	Roles          []commons.UserRole    `json:"roles,omitempty" bson:"roles,omitempty"`
	Sysinfo        commons.SysInfo       `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
	ProfilePicture commons.FileReference `json:"profilePicture,omitempty" bson:"profilePicture,omitempty"`
}

func (s User) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.Nickname == "" {
	       return false
	   }
	   if s.ObjType == "" {
	       return false
	   }
	   if s.Firstname == "" {
	       return false
	   }
	   if s.Lastname == "" {
	       return false
	   }
	   if s.Email == "" {
	       return false
	   }
	   if s.Password == "" {
	       return false
	   }
	   if len(s.Roles) == 0 {
	       return false
	   }
	   if s.Sysinfo.IsZero() {
	       return false
	   }
	   if s.ProfilePicture.IsZero() {
	       return false
	   }
	       return true
	*/

	return s.OId == primitive.NilObjectID && s.Nickname == "" && s.ObjType == "" && s.Firstname == "" && s.Lastname == "" && s.Email == "" && s.Password == "" && len(s.Roles) == 0 && s.Sysinfo.IsZero() && s.ProfilePicture.IsZero()
}
