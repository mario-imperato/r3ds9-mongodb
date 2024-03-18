package domain

import "github.com/mario-imperato/r3ds9-mongodb/model/commons"

// GetAppByObjTypeAndId search for a match of App.Id and App.ObjType. If not found it returns an empty App
// or the last configured App with same Id (and different objType). It works in the assumption to redirect to
// the counter-part of the defined app.
func (s Domain) GetAppByObjTypeAndId(objType commons.AppObjType, appId string) (commons.App, bool) {

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
