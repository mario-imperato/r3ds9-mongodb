package session

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AndSessionIdEqTo No Remarks
func (ca *Criteria) AndSessionIdEqTo(sid string) *Criteria {

	const SemLogContext = "r3ds9-core/session/and-hex-oid-eq-to"
	if sid == "" {
		return ca
	}

	oId, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		log.Error().Err(err).Msg(SemLogContext)
	}

	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}
