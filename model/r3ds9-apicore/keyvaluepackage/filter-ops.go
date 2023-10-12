package keyvaluepackage

import (
	"context"
	"fmt"
	"github.com/mario-imperato/r3ds9-apicommon/definitions"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindByDomainSiteName(c *mongo.Collection, domain, site, pkgName string) (KeyValuePackage, bool, error) {
	f := Filter{}
	criteria := f.Or().AndNameEqTo(pkgName)

	if domain == definitions.RootDomain {
		criteria.AndScopeEqTo(definitions.RootDomain)
	} else {
		if site == definitions.SiteWildCard {
			criteria.AndScopeIn([]string{definitions.RootDomain, domain})
		} else {
			criteria.AndScopeIn([]string{definitions.RootDomain, domain, fmt.Sprintf("%s/%s", domain, site)})
		}
	}

	kvp := KeyValuePackage{}

	crs, err := c.Find(context.TODO(), f.Build(), nil)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return kvp, false, err
		}

		return kvp, false, nil
	}

	for crs.Next(context.TODO()) {
		var kvp1 KeyValuePackage
		if err = crs.Decode(&kvp1); err != nil {
			return kvp, false, err
		}

		if ok, err := kvp1.IsMoreSpecificThan(&kvp); ok {
			kvp = kvp1
		} else if err != nil {
			return kvp, false, err
		}
	}

	if err = crs.Err(); err != nil {
		return kvp, false, err
	}

	if !kvp.IsZero() {
		scopeType, _ := ScopeTypeAndPathFromDomainSite(domain, site)
		kvpScopeType, _, _ := kvp.ScopeType()
		if scopeType != kvpScopeType {
			kvp.Inherited = true
		}
	}

	return kvp, !kvp.IsZero(), nil
}
