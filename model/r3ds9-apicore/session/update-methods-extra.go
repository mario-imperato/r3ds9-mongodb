package session

import (
	"fmt"
	"github.com/mario-imperato/r3ds9-mongodb/model/commons"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (ud *UpdateDocument) SetSysinfoModifiedAtNow() *UpdateDocument {
	mName := fmt.Sprintf(commons.SYSINFO_MODIFIEDAT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: primitive.NewDateTimeFromTime(time.Now())}
	})
	return ud
}
