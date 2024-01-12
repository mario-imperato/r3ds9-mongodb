package user

import (
	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apicore/commons"
	"strings"
)

func (s User) HasRole4DomainSiteAppId(domain, site, appId string) bool {
	return AnyRole4DomainSiteAppId(s.Roles, domain, site, appId)
}

func AnyRole4DomainSiteAppId(roles []commons.UserRole, domain, site, appId string) bool {
	for _, r := range roles {

		hr := true
		if hr && r.Domain != "*" {
			if r.Domain != domain {
				hr = false
			}
		}

		if hr && r.Site != "*" {
			if r.Site != site {
				hr = false
			}
		}

		if hr && r.Apps != "*" {
			appsArray := strings.Split(r.Apps, ",")

			hr = false
			for _, a := range appsArray {

				a = strings.TrimSpace(a)
				if a == "" {
					continue
				}

				if a == "*" {
					hr = true
					break
				}

				ndx := strings.Index(a, "*")
				if ndx >= 0 {
					na := a[:ndx]
					if strings.HasPrefix(a, na) {
						hr = true
					}
				} else {
					if a == appId {
						hr = true
					}
				}

				if hr == true {
					break
				}
			}
		}

		if hr {
			return true
		}
	}

	return false
}
