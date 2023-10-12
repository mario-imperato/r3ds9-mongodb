package keyvaluepackage

import (
	"fmt"
	"github.com/mario-imperato/r3ds9-apicommon/definitions"
	"strings"
)

func (kvp KeyValuePackage) ScopeType() (string, string, error) {
	return ScopeTypeFrom(kvp.Scope)
}

func (kvp *KeyValuePackage) IsMoreSpecificThan(another *KeyValuePackage) (bool, error) {
	return ScopeIsMoreSpecificThan(kvp.Scope, another.Scope)
}

func ScopeTypeFrom(scope string) (string, string, error) {
	if scope == "" {
		return "unknown-scope", "", fmt.Errorf("scope cannot be resolved since is missing")
	}

	if scope == definitions.RootDomain {
		return "root-scope", definitions.RootDomain, nil
	}

	if strings.Index(scope, "/") < 0 {
		return "domain-scope", strings.Join([]string{definitions.RootDomain, scope}, "/"), nil
	}

	return "site-scope", strings.Join([]string{definitions.RootDomain, scope}, "/"), nil
}

func ScopeTypeAndPathFromDomainSite(domain, site string) (string, string) {
	if domain == definitions.RootDomain {
		return "root-scope", definitions.RootDomain
	}

	if site == definitions.SiteWildCard {
		return "domain-scope", strings.Join([]string{definitions.RootDomain, domain}, "/")
	}

	return "site-scope", strings.Join([]string{definitions.RootDomain, domain, site}, "/")
}

func ScopeIsMoreSpecificThan(scope string, another string) (bool, error) {
	if scope == "" {
		return false, nil
	}

	if another == "" {
		return true, nil
	}

	_, scopePath, err := ScopeTypeFrom(scope)
	if err != nil {
		return false, err
	}

	_, anotherScopePath, err := ScopeTypeFrom(another)
	if err != nil {
		return false, err
	}

	if strings.HasPrefix(scopePath, anotherScopePath) {
		return true, nil
	}

	if !strings.HasPrefix(anotherScopePath, scopePath) {
		err = fmt.Errorf("incompatible paths compared: %s against %s", scopePath, anotherScopePath)
	}

	return false, err
}
