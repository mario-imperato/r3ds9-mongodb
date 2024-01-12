package site

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindByDomainAndCode(collection *mongo.Collection, domain, code string, mustFind bool, findOptions *options.FindOneOptions) (*Site, bool, error) {

	const SemLogContext = "r3ds9-core/site/find-by-domain-and-code"
	log.Trace().Str("domain", code).Msg(SemLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := Site{}

	f := Filter{}
	f.Or().AndCodeEqTo(code).AndDomainEqTo(domain)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error().Err(err).Msg("site find operation")
		return &ent, false, err
	} else {
		if err != nil {
			if mustFind {
				log.Trace().Str("site", code).Str("domain", domain).Msg(SemLogContext + " document not found")
				return &ent, false, err
			}

			log.Trace().Str("site", code).Str("domain", domain).Msg(SemLogContext + " document not found but allowed")
			ent.Code = code
			return &ent, false, nil
		} else {
			log.Trace().Str("site", code).Str("domain", domain).Msg(SemLogContext + " document found")
		}
	}

	return &ent, true, nil
}
