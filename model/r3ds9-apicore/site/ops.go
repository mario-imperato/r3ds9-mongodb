package site

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindByDomainAndCode(c *mongo.Collection, domain, site string) (Site, bool, error) {
	f := Filter{}
	f.Or().AndCodeEqTo(site).AndDomainEqTo(domain)

	s := Site{}
	err := c.FindOne(context.TODO(), f.Build(), nil).Decode(&s)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return s, false, err
		}

		return s, false, nil
	}

	return s, true, nil
}
