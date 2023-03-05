package domain

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apicore/commons"

const (
	OID               = "_id"
	CODE              = "code"
	OBJTYPE           = "objType"
	NAME              = "name"
	DESCRIPTION       = "description"
	LANGS             = "langs"
	MEMBERS           = "members"
	MEMBERS_I         = "members.%d"
	MEMBERS_I_CODE    = "members.%d.code"
	MEMBERS_CODE      = "members.code"
	MEMBERS_I_OBJTYPE = "members.%d.objType"
	MEMBERS_OBJTYPE   = "members.objType"
	APPS              = "apps"
	APPS_I            = "apps.%d"
	SYSINFO           = "sysinfo"
)

type Domain struct {
	OId         primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Code        string             `json:"code,omitempty" bson:"code,omitempty"`
	ObjType     string             `json:"objType,omitempty" bson:"objType,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Langs       string             `json:"langs,omitempty" bson:"langs,omitempty"`
	Members     []DomainMember     `json:"members,omitempty" bson:"members,omitempty"`
	Apps        []commons.App      `json:"apps,omitempty" bson:"apps,omitempty"`
	Sysinfo     commons.SysInfo    `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
}

func (s Domain) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.Code == "" {
	       return false
	   }
	   if s.ObjType == "" {
	       return false
	   }
	   if s.Name == "" {
	       return false
	   }
	   if s.Description == "" {
	       return false
	   }
	   if s.Langs == "" {
	       return false
	   }
	   if len(s.Members) == 0 {
	       return false
	   }
	   if len(s.Apps) == 0 {
	       return false
	   }
	   if s.Sysinfo.IsZero() {
	       return false
	   }
	       return true
	*/

	return s.OId == primitive.NilObjectID && s.Code == "" && s.ObjType == "" && s.Name == "" && s.Description == "" && s.Langs == "" && len(s.Members) == 0 && len(s.Apps) == 0 && s.Sysinfo.IsZero()
}

type DomainMember struct {
	Code    string `json:"code,omitempty" bson:"code,omitempty"`
	ObjType string `json:"objType,omitempty" bson:"objType,omitempty"`
}

func (s DomainMember) IsZero() bool {
	/*
	   if s.Code == "" {
	       return false
	   }
	   if s.ObjType == "" {
	       return false
	   }
	       return true
	*/
	return s.Code == "" && s.ObjType == ""
}
