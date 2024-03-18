package keyvaluepackage_test

import (
	"fmt"

	"github.com/mario-imperato/r3ds9-mongodb/model"
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apicore/keyvaluepackage"
	"github.com/stretchr/testify/require"
	"testing"
)

type InputWanted struct {
	scope   string
	another string
	result  bool
	error   bool
}

type InputWantedDomainSite struct {
	domain string
	site   string
	scope  string
}

func TestScope(t *testing.T) {

	scopeCases := []InputWanted{
		{scope: model.RootDomain, another: model.RootDomain, result: true},
		{scope: model.RootDomain, another: "cvf", result: false},
		{scope: model.RootDomain, another: "cvf/mySite", result: false},
		{scope: "cvf", another: model.RootDomain, result: true},
		{scope: "cvf", another: "cvf/mySite", result: false},
		{scope: "cvf/mySite", another: "cvf", result: true},
		{scope: "cvf/mySite", another: "cvf/yourSite", result: false, error: true},
	}

	for i, iw := range scopeCases {
		b, err := keyvaluepackage.ScopeIsMoreSpecificThan(iw.scope, iw.another)
		if iw.error {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, iw.result, b, fmt.Sprintf("[%d] scope: %s, another: %s", i, iw.scope, iw.another))
		}
	}

	domainSiteCases := []InputWantedDomainSite{
		{domain: model.RootDomain, site: model.SiteWildCard, scope: model.RootDomain},
		{domain: "cvf", site: model.SiteWildCard, scope: "root/cvf"},
		{domain: "cvf", site: "mySite", scope: "root/cvf/mySite"},
	}
	for i, iw := range domainSiteCases {
		_, s := keyvaluepackage.ScopeTypeAndPathFromDomainSite(iw.domain, iw.site)
		require.Equal(t, iw.scope, s, fmt.Sprintf("[%d] domain: %s, site: %s, scope: %s", i, iw.domain, iw.site, iw.scope))
	}
}
