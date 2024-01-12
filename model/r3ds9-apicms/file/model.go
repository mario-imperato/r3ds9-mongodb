package file

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/mario-imperato/r3ds9-mongodb/model/commons"

const (
	OID         = "_id"
	FILENAME    = "fileName"
	CONTAINER   = "container"
	URL         = "url"
	CONTENTTYPE = "contentType"
	WIDTH       = "width"
	HEIGHT      = "height"
	METADATA    = "metadata"
	METADATA_S  = "metadata.%s"
	SYSINFO     = "sysinfo"
)

type File struct {
	OId         primitive.ObjectID     `json:"-" bson:"_id,omitempty"`
	FileName    string                 `json:"fileName,omitempty" bson:"fileName,omitempty"`
	Container   string                 `json:"container,omitempty" bson:"container,omitempty"`
	Url         string                 `json:"url,omitempty" bson:"url,omitempty"`
	ContentType string                 `json:"contentType,omitempty" bson:"contentType,omitempty"`
	Width       int32                  `json:"width,omitempty" bson:"width,omitempty"`
	Height      int32                  `json:"height,omitempty" bson:"height,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Sysinfo     commons.SysInfo        `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
}

func (s File) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.FileName == "" {
	       return false
	   }
	   if s.Container == "" {
	       return false
	   }
	   if s.Url == "" {
	       return false
	   }
	   if s.ContentType == "" {
	       return false
	   }
	   if s.Width == 0 {
	       return false
	   }
	   if s.Height == 0 {
	       return false
	   }
	   if len(s.Metadata) == 0 {
	       return false
	   }
	   if s.Sysinfo.IsZero() {
	       return false
	   }
	       return true
	*/

	return s.OId == primitive.NilObjectID && s.FileName == "" && s.Container == "" && s.Url == "" && s.ContentType == "" && s.Width == 0 && s.Height == 0 && len(s.Metadata) == 0 && s.Sysinfo.IsZero()
}
