package session

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/commons"

const (
	OID        = "_id"
	NICKNAME   = "nickname"
	REMOTEADDR = "remoteaddr"
	FLAGS      = "flags"
	SYSINFO    = "sysinfo"
)

type Session struct {
	OId        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Nickname   string             `json:"nickname,omitempty" bson:"nickname,omitempty"`
	Remoteaddr string             `json:"remoteaddr,omitempty" bson:"remoteaddr,omitempty"`
	Flags      string             `json:"flags,omitempty" bson:"flags,omitempty"`
	Sysinfo    commons.SysInfo    `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
}

func (s Session) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.Nickname == "" {
	       return false
	   }
	   if s.Remoteaddr == "" {
	       return false
	   }
	   if s.Flags == "" {
	       return false
	   }
	   if s.Sysinfo.IsZero() {
	       return false
	   }
	       return true
	*/

	return s.OId == primitive.NilObjectID && s.Nickname == "" && s.Remoteaddr == "" && s.Flags == "" && s.Sysinfo.IsZero()
}
