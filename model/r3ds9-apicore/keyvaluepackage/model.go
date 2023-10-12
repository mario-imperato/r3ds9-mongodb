package keyvaluepackage

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/commons"

const (
	OID                = "_id"
	DOMAIN             = "domain"
	SITE               = "site"
	OBJTYPE            = "objType"
	CATEGORY           = "category"
	ISSYSTEM           = "issystem"
	DESCRIPTION        = "description"
	INHERIT            = "inherit"
	PROPERTIES         = "properties"
	PROPERTIES_I       = "properties.%d"
	PROPERTIES_I_KEY   = "properties.%d.key"
	PROPERTIES_KEY     = "properties.key"
	PROPERTIES_I_VALUE = "properties.%d.value"
	PROPERTIES_VALUE   = "properties.value"
	SYSINFO            = "sysinfo"
)

type KeyValuePackage struct {
	OId         primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Domain      string             `json:"domain,omitempty" bson:"domain,omitempty"`
	Site        string             `json:"site,omitempty" bson:"site,omitempty"`
	ObjType     string             `json:"objType,omitempty" bson:"objType,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Issystem    bool               `json:"issystem,omitempty" bson:"issystem,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Inherit     bool               `json:"inherit,omitempty" bson:"inherit,omitempty"`
	Properties  []KeyValue         `json:"properties,omitempty" bson:"properties,omitempty"`
	Sysinfo     commons.SysInfo    `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
}

func (s KeyValuePackage) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.Domain == "" {
	       return false
	   }
	   if s.Site == "" {
	       return false
	   }
	   if s.ObjType == "" {
	       return false
	   }
	   if s.Category == "" {
	       return false
	   }
	   if !s.Issystem {
	       return false
	   }
	   if s.Description == "" {
	       return false
	   }
	   if !s.Inherit {
	       return false
	   }
	   if len(s.Properties) == 0 {
	       return false
	   }
	   if s.Sysinfo.IsZero() {
	       return false
	   }
	       return true
	*/

	return s.OId == primitive.NilObjectID && s.Domain == "" && s.Site == "" && s.ObjType == "" && s.Category == "" && !s.Issystem && s.Description == "" && !s.Inherit && len(s.Properties) == 0 && s.Sysinfo.IsZero()
}

type KeyValue struct {
	Key   string `json:"key,omitempty" bson:"key,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

func (s KeyValue) IsZero() bool {
	/*
	   if s.Key == "" {
	       return false
	   }
	   if s.Value == "" {
	       return false
	   }
	       return true
	*/
	return s.Key == "" && s.Value == ""
}
