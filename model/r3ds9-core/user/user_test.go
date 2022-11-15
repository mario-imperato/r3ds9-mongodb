package user_test

import (
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-core/commons"
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-core/user"
	"github.com/stretchr/testify/require"
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

type roleTestCases struct {
	domain   string
	site     string
	appId    string
	roles    []commons.UserRole
	expected bool
}

func TestHasRole(t *testing.T) {

	cases := []roleTestCases{
		{"dom1", "site1", "app-1", []commons.UserRole{}, false},
		{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "*", Site: "*", Apps: "appero,app-2"}}, false},
		{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "*", Site: "*", Apps: "appero, app-1"}}, true},
		{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "*", Site: "*", Apps: "appero,app-*"}}, true},
		{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "*", Site: "site2", Apps: "app-*"}}, false},
		{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "dom2", Site: "*", Apps: "app-*"}}, false},
		{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "dom1", Site: "*", Apps: "app-*"}}, true},
		{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "dom1", Site: "site1", Apps: "app-*"}, {Domain: "dom2", Site: "site2", Apps: "app-2"}}, true},
	}

	for _, c := range cases {
		ok := user.AnyRole4DomainSiteAppId(c.roles, c.domain, c.site, c.appId)
		require.Equal(t, ok, c.expected, "actual %t, wanted: %t")
	}
}
