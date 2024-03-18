package site

import "github.com/mario-imperato/r3ds9-mongodb/model/commons"
import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OID         = "_id"
	CODE        = "code"
	DOMAIN      = "domain"
	OBJTYPE     = "objType"
	NAME        = "name"
	DESCRIPTION = "description"
	BOOKMARK    = "bookmark"
	LANGS       = "langs"
	APPS        = "apps"
	APPS_I      = "apps.%d"
	SYSINFO     = "sysinfo"
)

type Site struct {
	OId         primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Code        string             `json:"code,omitempty" bson:"code,omitempty"`
	Domain      string             `json:"domain,omitempty" bson:"domain,omitempty"`
	ObjType     string             `json:"objType,omitempty" bson:"objType,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Bookmark    bool               `json:"bookmark,omitempty" bson:"bookmark,omitempty"`
	Langs       string             `json:"langs,omitempty" bson:"langs,omitempty"`
	Apps        []commons.App      `json:"apps,omitempty" bson:"apps,omitempty"`
	Sysinfo     commons.SysInfo    `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
}

func (s Site) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.Code == "" {
	       return false
	   }
	   if s.Domain == "" {
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
	   if !s.Bookmark {
	       return false
	   }
	   if s.Langs == "" {
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

	return s.OId == primitive.NilObjectID && s.Code == "" && s.Domain == "" && s.ObjType == "" && s.Name == "" && s.Description == "" && !s.Bookmark && s.Langs == "" && len(s.Apps) == 0 && s.Sysinfo.IsZero()
}
