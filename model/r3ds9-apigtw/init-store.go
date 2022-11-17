package r3ds9_apigtw

import (
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/domain"
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/session"
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/site"
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/user"
	"time"
)

const (
	ExpiryDuration       = 5 * time.Minute
	PurgeExpiredDuration = 10 * time.Minute
)

func InitStore() error {
	domain.NewCache(ExpiryDuration, PurgeExpiredDuration)
	site.NewCache(ExpiryDuration, PurgeExpiredDuration)
	user.NewCache(ExpiryDuration, PurgeExpiredDuration)
	session.NewCache(ExpiryDuration, PurgeExpiredDuration)
	return nil
}
