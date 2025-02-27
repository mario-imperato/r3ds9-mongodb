package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)
import "github.com/mario-imperato/r3ds9-mongodb/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type User struct {
	OId            primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Nickname       string                `json:"nickname,omitempty" bson:"nickname,omitempty" yaml:"nickname,omitempty"`
	ObjType        string                `json:"objType,omitempty" bson:"objType,omitempty" yaml:"objType,omitempty"`
	Firstname      string                `json:"firstname,omitempty" bson:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname       string                `json:"lastname,omitempty" bson:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email          string                `json:"email,omitempty" bson:"email,omitempty" yaml:"email,omitempty"`
	Password       string                `json:"password,omitempty" bson:"password,omitempty" yaml:"password,omitempty"`
	Roles          []commons.UserRole    `json:"roles,omitempty" bson:"roles,omitempty" yaml:"roles,omitempty"`
	SysInfo        commons.SysInfo       `json:"sysInfo,omitempty" bson:"sysInfo,omitempty" yaml:"sysInfo,omitempty"`
	ProfilePicture commons.FileReference `json:"profilePicture,omitempty" bson:"profilePicture,omitempty" yaml:"profilePicture,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s User) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Nickname == "" && s.ObjType == "" && s.Firstname == "" && s.Lastname == "" && s.Email == "" && s.Password == "" && len(s.Roles) == 0 && s.SysInfo.IsZero() && s.ProfilePicture.IsZero()
}

// @tpm-schematics:start-region("bottom-file-section")

type QueryResult struct {
	Records int    `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []User `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

func (s User) HasRole4DomainSiteAppId(domain, site, appId string) bool {
	return AnyRole4DomainSiteAppId(s.Roles, domain, site, appId)
}

func AnyRole4DomainSiteAppId(roles []commons.UserRole, domain, site, appId string) bool {
	for _, r := range roles {

		hr := true
		if hr && r.Domain != "*" {
			if r.Domain != domain {
				hr = false
			}
		}

		if hr && r.Site != "*" {
			if r.Site != site {
				hr = false
			}
		}

		if hr && r.Apps != "*" {
			appsArray := strings.Split(r.Apps, ",")

			hr = false
			for _, a := range appsArray {

				a = strings.TrimSpace(a)
				if a == "" {
					continue
				}

				if a == "*" {
					hr = true
					break
				}

				ndx := strings.Index(a, "*")
				if ndx >= 0 {
					na := a[:ndx]
					if strings.HasPrefix(a, na) {
						hr = true
					}
				} else {
					if a == appId {
						hr = true
					}
				}

				if hr == true {
					break
				}
			}
		}

		if hr {
			return true
		}
	}

	return false
}

// @tpm-schematics:end-region("bottom-file-section")
