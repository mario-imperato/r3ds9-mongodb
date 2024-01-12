package site_test

import (
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apicore/site"
	"testing"
	"time"
)

const (
	targetDomain = "cvf"
	targetSite   = "champ42"
)

func TestCache(t *testing.T) {

	r := site.NewCacheResolver(collection)

	site.NewCache(5*time.Second, 10*time.Minute)
	d, ok := site.GetFromCache(r, targetDomain, targetSite)
	t.Log(d, ok)

	t.Log("briefly sleeping....")
	time.Sleep(2 * time.Second)
	d, ok = site.GetFromCache(r, targetDomain, targetSite)
	t.Log(d, ok)

	t.Log("sleeping longer....")
	time.Sleep(10 * time.Second)

	d, ok = site.GetFromCache(r, targetDomain, targetSite)
	t.Log(d, ok)
}
