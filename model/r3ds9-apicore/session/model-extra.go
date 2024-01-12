package session

import "go.mongodb.org/mongo-driver/bson/primitive"

func (s Session) SessionId() string {
	if s.OId == primitive.NilObjectID {
		return ""
	}

	return s.OId.Hex()
}
