package commons

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMethodsGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

type UnsetMode int64

const (
	UnSpecified     UnsetMode = 0
	KeepCurrent               = 1
	UnsetData                 = 2
	SetData2Default           = 3
)

type UnsetOption func(uopt *UnsetOptions)

type UnsetOptions struct {
	DefaultMode UnsetMode
	Sysinfo     UnsetMode
	App         UnsetMode
	Roles       UnsetMode
}

func (uo *UnsetOptions) ResolveUnsetMode(um UnsetMode) UnsetMode {
	if um == UnSpecified {
		um = uo.DefaultMode
	}

	return um
}

func WithDefaultUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.DefaultMode = m
	}
}
func WithSysinfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Sysinfo = m
	}
}
func WithAppUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.App = m
	}
}
func WithRolesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Roles = m
	}
}

type UpdateOption func(ud *UpdateDocument)
type UpdateOptions []UpdateOption

// GetUpdateDocument convenience method to create an update document from single updates instead of a whole object
func GetUpdateDocumentFromOptions(opts ...UpdateOption) UpdateDocument {
	ud := UpdateDocument{}
	for _, o := range opts {
		o(&ud)
	}

	return ud
}

// GetUpdateDocument
// Convenience method to create an Update Document from the values of the top fields of the object. The convenience is in the handling
// the unset because if I pass an empty struct to the update it generates an empty object anyway in the db. Handling the unset eliminates
// the issue and delete an existing value without creating an empty struct.
func GetUpdateDocument(obj *Commons, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	// if !obj.Sysinfo.IsZero() {
	//   ud.SetSysinfo ( obj.Sysinfo)
	// } else {
	ud.setOrUnsetSysinfo(obj.Sysinfo, uo.ResolveUnsetMode(uo.Sysinfo))
	// }
	// if !obj.App.IsZero() {
	//   ud.SetApp ( obj.App)
	// } else {
	ud.setOrUnsetApp(obj.App, uo.ResolveUnsetMode(uo.App))
	// }
	// if len(obj.Roles) > 0 {
	//   ud.SetRoles ( obj.Roles)
	// } else {
	ud.setOrUnsetRoles(obj.Roles, uo.ResolveUnsetMode(uo.Roles))
	// }

	return ud
}

// ----- sysinfo - struct - SysInfo [sysinfo]
// SetSysinfo No Remarks
func (ud *UpdateDocument) SetSysinfo(p SysInfo) *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSysinfo No Remarks
func (ud *UpdateDocument) UnsetSysinfo() *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSysinfo No Remarks - here2
func (ud *UpdateDocument) setOrUnsetSysinfo(p SysInfo, um UnsetMode) {

	//----- struct\n

	if !p.IsZero() {
		ud.SetSysinfo(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSysinfo()
		case SetData2Default:
			ud.UnsetSysinfo()
		}
	}
}

func UpdateWithSysinfo(p SysInfo) UpdateOption {
	return func(ud *UpdateDocument) {

		if !p.IsZero() {
			ud.SetSysinfo(p)
		} else {
			ud.UnsetSysinfo()
		}
	}
}

//----- status - string -  [sysinfo.status]

// SetSysinfoStatus No Remarks
func (ud *UpdateDocument) SetSysinfoStatus(p string) *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO_STATUS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSysinfoStatus No Remarks
func (ud *UpdateDocument) UnsetSysinfoStatus() *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO_STATUS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSysinfoStatus No Remarks
func (ud *UpdateDocument) setOrUnsetSysinfoStatus(p string, um UnsetMode) {
	if p != "" {
		ud.SetSysinfoStatus(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSysinfoStatus()
		case SetData2Default:
			ud.UnsetSysinfoStatus()
		}
	}
}

func UpdateWithSysinfoStatus(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetSysinfoStatus(p)
		} else {
			ud.UnsetSysinfoStatus()
		}
	}
}

//----- createdat - date -  [sysinfo.createdat]

// SetSysinfoCreatedat No Remarks
func (ud *UpdateDocument) SetSysinfoCreatedat(p primitive.DateTime) *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO_CREATEDAT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSysinfoCreatedat No Remarks
func (ud *UpdateDocument) UnsetSysinfoCreatedat() *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO_CREATEDAT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSysinfoCreatedat No Remarks
func (ud *UpdateDocument) setOrUnsetSysinfoCreatedat(p primitive.DateTime, um UnsetMode) {
	if p != 0 {
		ud.SetSysinfoCreatedat(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSysinfoCreatedat()
		case SetData2Default:
			ud.UnsetSysinfoCreatedat()
		}
	}
}

// SetSysinfoCreatedatToCurrentDate
func (ud *UpdateDocument) SetSysinfoCreatedatToCurrentDate() *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO_CREATEDAT)
	ud.CurrentDate().Add(func() bson.E {
		return bson.E{Key: mName, Value: true}
	})
	return ud
}

func UpdateWithSysinfoCreatedat(p primitive.DateTime) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetSysinfoCreatedat(p)
		} else {
			ud.UnsetSysinfoCreatedat()
		}
	}
}

//----- modifiedat - date -  [sysinfo.modifiedat]

// SetSysinfoModifiedat No Remarks
func (ud *UpdateDocument) SetSysinfoModifiedat(p primitive.DateTime) *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO_MODIFIEDAT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSysinfoModifiedat No Remarks
func (ud *UpdateDocument) UnsetSysinfoModifiedat() *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO_MODIFIEDAT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSysinfoModifiedat No Remarks
func (ud *UpdateDocument) setOrUnsetSysinfoModifiedat(p primitive.DateTime, um UnsetMode) {
	if p != 0 {
		ud.SetSysinfoModifiedat(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSysinfoModifiedat()
		case SetData2Default:
			ud.UnsetSysinfoModifiedat()
		}
	}
}

// SetSysinfoModifiedatToCurrentDate
func (ud *UpdateDocument) SetSysinfoModifiedatToCurrentDate() *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO_MODIFIEDAT)
	ud.CurrentDate().Add(func() bson.E {
		return bson.E{Key: mName, Value: true}
	})
	return ud
}

func UpdateWithSysinfoModifiedat(p primitive.DateTime) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetSysinfoModifiedat(p)
		} else {
			ud.UnsetSysinfoModifiedat()
		}
	}
}

// ----- app - struct - App [app]
// SetApp No Remarks
func (ud *UpdateDocument) SetApp(p App) *UpdateDocument {
	mName := fmt.Sprintf(APP)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetApp No Remarks
func (ud *UpdateDocument) UnsetApp() *UpdateDocument {
	mName := fmt.Sprintf(APP)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetApp No Remarks - here2
func (ud *UpdateDocument) setOrUnsetApp(p App, um UnsetMode) {

	//----- struct\n

	if !p.IsZero() {
		ud.SetApp(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetApp()
		case SetData2Default:
			ud.UnsetApp()
		}
	}
}

func UpdateWithApp(p App) UpdateOption {
	return func(ud *UpdateDocument) {

		if !p.IsZero() {
			ud.SetApp(p)
		} else {
			ud.UnsetApp()
		}
	}
}

//----- id - string -  [app.id]

// SetAppId No Remarks
func (ud *UpdateDocument) SetAppId(p string) *UpdateDocument {
	mName := fmt.Sprintf(APP_ID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppId No Remarks
func (ud *UpdateDocument) UnsetAppId() *UpdateDocument {
	mName := fmt.Sprintf(APP_ID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppId No Remarks
func (ud *UpdateDocument) setOrUnsetAppId(p string, um UnsetMode) {
	if p != "" {
		ud.SetAppId(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppId()
		case SetData2Default:
			ud.UnsetAppId()
		}
	}
}

func UpdateWithAppId(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppId(p)
		} else {
			ud.UnsetAppId()
		}
	}
}

//----- objType - string -  [app.objType]

// SetAppObjType No Remarks
func (ud *UpdateDocument) SetAppObjType(p string) *UpdateDocument {
	mName := fmt.Sprintf(APP_OBJTYPE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppObjType No Remarks
func (ud *UpdateDocument) UnsetAppObjType() *UpdateDocument {
	mName := fmt.Sprintf(APP_OBJTYPE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppObjType No Remarks
func (ud *UpdateDocument) setOrUnsetAppObjType(p string, um UnsetMode) {
	if p != "" {
		ud.SetAppObjType(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppObjType()
		case SetData2Default:
			ud.UnsetAppObjType()
		}
	}
}

func UpdateWithAppObjType(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppObjType(p)
		} else {
			ud.UnsetAppObjType()
		}
	}
}

//----- name - string -  [app.name]

// SetAppName No Remarks
func (ud *UpdateDocument) SetAppName(p string) *UpdateDocument {
	mName := fmt.Sprintf(APP_NAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppName No Remarks
func (ud *UpdateDocument) UnsetAppName() *UpdateDocument {
	mName := fmt.Sprintf(APP_NAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppName No Remarks
func (ud *UpdateDocument) setOrUnsetAppName(p string, um UnsetMode) {
	if p != "" {
		ud.SetAppName(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppName()
		case SetData2Default:
			ud.UnsetAppName()
		}
	}
}

func UpdateWithAppName(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppName(p)
		} else {
			ud.UnsetAppName()
		}
	}
}

//----- description - string -  [app.description]

// SetAppDescription No Remarks
func (ud *UpdateDocument) SetAppDescription(p string) *UpdateDocument {
	mName := fmt.Sprintf(APP_DESCRIPTION)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppDescription No Remarks
func (ud *UpdateDocument) UnsetAppDescription() *UpdateDocument {
	mName := fmt.Sprintf(APP_DESCRIPTION)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppDescription No Remarks
func (ud *UpdateDocument) setOrUnsetAppDescription(p string, um UnsetMode) {
	if p != "" {
		ud.SetAppDescription(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppDescription()
		case SetData2Default:
			ud.UnsetAppDescription()
		}
	}
}

func UpdateWithAppDescription(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppDescription(p)
		} else {
			ud.UnsetAppDescription()
		}
	}
}

//----- path - string -  [app.path]

// SetAppPath No Remarks
func (ud *UpdateDocument) SetAppPath(p string) *UpdateDocument {
	mName := fmt.Sprintf(APP_PATH)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppPath No Remarks
func (ud *UpdateDocument) UnsetAppPath() *UpdateDocument {
	mName := fmt.Sprintf(APP_PATH)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppPath No Remarks
func (ud *UpdateDocument) setOrUnsetAppPath(p string, um UnsetMode) {
	if p != "" {
		ud.SetAppPath(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppPath()
		case SetData2Default:
			ud.UnsetAppPath()
		}
	}
}

func UpdateWithAppPath(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppPath(p)
		} else {
			ud.UnsetAppPath()
		}
	}
}

//----- roleRequired - bool -  [app.roleRequired]

// SetAppRoleRequired No Remarks
func (ud *UpdateDocument) SetAppRoleRequired(p bool) *UpdateDocument {
	mName := fmt.Sprintf(APP_ROLEREQUIRED)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppRoleRequired
func (ud *UpdateDocument) UnsetAppRoleRequired() *UpdateDocument {
	mName := fmt.Sprintf(APP_ROLEREQUIRED)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppRoleRequired No Remarks
func (ud *UpdateDocument) setOrUnsetAppRoleRequired(p bool, um UnsetMode) {
	if p {
		ud.SetAppRoleRequired(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppRoleRequired()
		case SetData2Default:
			ud.UnsetAppRoleRequired()
		}
	}
}

func UpdateWithAppRoleRequired(p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetAppRoleRequired(p)
		} else {
			ud.UnsetAppRoleRequired()
		}
	}
}

// ----- roles - array -  [roles]
// SetRoles No Remarks
func (ud *UpdateDocument) SetRoles(p []UserRole) *UpdateDocument {
	mName := fmt.Sprintf(ROLES)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRoles No Remarks
func (ud *UpdateDocument) UnsetRoles() *UpdateDocument {
	mName := fmt.Sprintf(ROLES)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRoles No Remarks - here2
func (ud *UpdateDocument) setOrUnsetRoles(p []UserRole, um UnsetMode) {

	//----- array\n

	if len(p) > 0 {
		ud.SetRoles(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRoles()
		case SetData2Default:
			ud.UnsetRoles()
		}
	}
}

func UpdateWithRoles(p []UserRole) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetRoles(p)
		} else {
			ud.UnsetRoles()
		}
	}
}

// ----- [] - struct - UserRole [roles.[]]
// SetRolesI No Remarks
func (ud *UpdateDocument) SetRolesI(ndxI int, p UserRole) *UpdateDocument {
	mName := fmt.Sprintf(ROLES_I, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRolesI No Remarks
func (ud *UpdateDocument) UnsetRolesI(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(ROLES_I, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRolesI No Remarks
func (ud *UpdateDocument) setOrUnsetRolesI(ndxI int, p UserRole, um UnsetMode) {
	if !p.IsZero() {
		ud.SetRolesI(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRolesI(ndxI)
		case SetData2Default:
			ud.UnsetRolesI(ndxI)
		}
	}
}

//----- domain - string -  [roles.[].domain roles.domain]

// SetRolesIDomain No Remarks
func (ud *UpdateDocument) SetRolesIDomain(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(ROLES_I_DOMAIN, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRolesIDomain No Remarks
func (ud *UpdateDocument) UnsetRolesIDomain(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(ROLES_I_DOMAIN, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRolesIDomain No Remarks
func (ud *UpdateDocument) setOrUnsetRolesIDomain(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetRolesIDomain(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRolesIDomain(ndxI)
		case SetData2Default:
			ud.UnsetRolesIDomain(ndxI)
		}
	}
}

func UpdateWithRolesIDomain(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRolesIDomain(ndxI, p)
		} else {
			ud.UnsetRolesIDomain(ndxI)
		}
	}
}

//----- site - string -  [roles.[].site roles.site]

// SetRolesISite No Remarks
func (ud *UpdateDocument) SetRolesISite(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(ROLES_I_SITE, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRolesISite No Remarks
func (ud *UpdateDocument) UnsetRolesISite(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(ROLES_I_SITE, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRolesISite No Remarks
func (ud *UpdateDocument) setOrUnsetRolesISite(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetRolesISite(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRolesISite(ndxI)
		case SetData2Default:
			ud.UnsetRolesISite(ndxI)
		}
	}
}

func UpdateWithRolesISite(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRolesISite(ndxI, p)
		} else {
			ud.UnsetRolesISite(ndxI)
		}
	}
}

//----- apps - string -  [roles.[].apps roles.apps]

// SetRolesIApps No Remarks
func (ud *UpdateDocument) SetRolesIApps(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(ROLES_I_APPS, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRolesIApps No Remarks
func (ud *UpdateDocument) UnsetRolesIApps(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(ROLES_I_APPS, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRolesIApps No Remarks
func (ud *UpdateDocument) setOrUnsetRolesIApps(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetRolesIApps(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRolesIApps(ndxI)
		case SetData2Default:
			ud.UnsetRolesIApps(ndxI)
		}
	}
}

func UpdateWithRolesIApps(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRolesIApps(ndxI, p)
		} else {
			ud.UnsetRolesIApps(ndxI)
		}
	}
}
