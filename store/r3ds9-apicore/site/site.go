package site

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/mario-imperato/r3ds9-mongodb/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Site struct {
	OId         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Code        string             `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	Domain      string             `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	ObjType     string             `json:"objType,omitempty" bson:"objType,omitempty" yaml:"objType,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Bookmark    bool               `json:"bookmark,omitempty" bson:"bookmark,omitempty" yaml:"bookmark,omitempty"`
	Langs       string             `json:"langs,omitempty" bson:"langs,omitempty" yaml:"langs,omitempty"`
	Apps        []commons.App      `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`
	SysInfo     commons.SysInfo    `json:"sysInfo,omitempty" bson:"sysInfo,omitempty" yaml:"sysInfo,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Site) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Code == "" && s.Domain == "" && s.ObjType == "" && s.Name == "" && s.Description == "" && !s.Bookmark && s.Langs == "" && len(s.Apps) == 0 && s.SysInfo.IsZero()
}

// @tpm-schematics:start-region("bottom-file-section")

func (s Site) GetAppByObjTypeAndId(objType commons.AppObjType, appId string) (commons.App, bool) {
	app := commons.App{Id: appId}
	for _, a := range s.Apps {
		if a.Id == appId {
			app = a
			if a.ObjType == string(objType) {
				return a, true
			}
		}
	}

	return app, false
}

// @tpm-schematics:end-region("bottom-file-section")
