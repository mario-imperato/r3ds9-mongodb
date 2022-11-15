package user_test

import (
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-core/user"
	"testing"
	"time"
)

const (
	targetDomain = "cvf"
)

func TestCache(t *testing.T) {

	r := user.NewCacheResolver(collection)

	user.NewCache(5*time.Second, 10*time.Minute)
	d, ok := user.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("briefly sleeping....")
	time.Sleep(2 * time.Second)
	d, ok = user.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("sleeping longer....")
	time.Sleep(10 * time.Second)

	d, ok = user.GetFromCache(r, targetDomain)
	t.Log(d, ok)
}
