package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindByCode(c *mongo.Collection, domain string) (Domain, bool, error) {
	f := Filter{}
	f.Or().AndCodeEqTo(domain)

	s := Domain{}
	err := c.FindOne(context.TODO(), f.Build(), nil).Decode(&s)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return s, false, err
		}

		return s, false, nil
	}

	return s, true, nil
}
