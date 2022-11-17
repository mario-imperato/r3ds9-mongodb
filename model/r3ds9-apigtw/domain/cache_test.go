package domain_test

import (
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/domain"
	"testing"
	"time"
)

const (
	targetDomain = "cvf"
)

func TestCache(t *testing.T) {

	r := domain.NewCacheResolver(collection)

	domain.NewCache(5*time.Second, 10*time.Minute)
	d, ok := domain.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("briefly sleeping....")
	time.Sleep(2 * time.Second)
	d, ok = domain.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("sleeping longer....")
	time.Sleep(10 * time.Second)

	d, ok = domain.GetFromCache(r, targetDomain)
	t.Log(d, ok)
}
