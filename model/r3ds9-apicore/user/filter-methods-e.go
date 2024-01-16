package user

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ca *Criteria) AndHexOIdEqTo(oId string) *Criteria {
	const semLogContext = "mongo-user::and-hex-oid-eq-to"
	if oId == "" {
		return ca
	}

	objId, err := primitive.ObjectIDFromHex(oId)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return ca
	}
	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: objId} }
	*ca = append(*ca, c)
	return ca
}
