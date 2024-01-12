package user

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindByNickname(c *mongo.Collection, nickName string) (User, bool, error) {
	f := Filter{}
	f.Or().AndNicknameEqTo(nickName)

	u := User{}
	err := c.FindOne(context.TODO(), f.Build(), nil).Decode(&u)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return u, false, err
		}

		return u, false, nil
	}

	return u, true, nil
}
