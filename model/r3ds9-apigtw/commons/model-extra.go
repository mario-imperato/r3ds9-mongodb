package commons

type AppObjType string

const (
	AppObjTypeWWW     AppObjType = "app-www"
	AppObjTypeConsole AppObjType = "app-console"
)

// AppId ids of the available apps. The actual app is configured in the database.
type AppId string

const (
	AppIdHome AppId = "app-home"
)

var AppIdsWorld = []AppId{AppIdHome}

func IsAppIdInCatalog(n string) bool {
	for _, a := range AppIdsWorld {
		if string(a) == n {
			return true
		}
	}

	return false
}
