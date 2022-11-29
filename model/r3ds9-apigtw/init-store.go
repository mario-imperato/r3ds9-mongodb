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

func InitStore(expiryDuration time.Duration, purgeExpiryDuration time.Duration) error {

	if expiryDuration == 0 {
		expiryDuration = ExpiryDuration
	}

	if purgeExpiryDuration == 0 {
		purgeExpiryDuration = PurgeExpiredDuration
	}

	domain.NewCache(expiryDuration, purgeExpiryDuration)
	site.NewCache(expiryDuration, purgeExpiryDuration)
	user.NewCache(expiryDuration, purgeExpiryDuration)
	session.NewCache(expiryDuration, purgeExpiryDuration)
	return nil
}
