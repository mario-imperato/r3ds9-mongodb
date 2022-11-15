package domain

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindByCode(collection *mongo.Collection, code string, mustFind bool, findOptions *options.FindOneOptions) (*Domain, bool, error) {

	const SemLogContext = "r3ds9-core/domain/find-by-code"
	log.Trace().Str("domain", code).Msg(SemLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := Domain{}

	f := Filter{}
	f.Or().AndCodeEqTo(code)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error().Err(err).Msg(SemLogContext)
		return &ent, false, err
	} else {
		if err != nil {
			if mustFind {
				log.Trace().Str("domain", code).Msg(SemLogContext + " document not found")
				return &ent, false, err
			}

			log.Trace().Str("ndg", code).Msg(SemLogContext + " document not found but allowed")
			ent.Code = code
			return &ent, false, nil
		} else {
			log.Trace().Str("domain", code).Msg(SemLogContext + " document found")
		}
	}

	return &ent, true, nil
}
