package site

import "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/commons"

func (s Site) GetAppByObjTypeAndId(objType commons.AppObjType, appId string) (commons.App, bool) {
	app := commons.App{Id: appId}
	for _, a := range s.Apps {
		if a.Id == appId {
			app = a
			if a.ObjType == string(objType) {
				return a, true
			}
		}
	}

	return app, false
}
