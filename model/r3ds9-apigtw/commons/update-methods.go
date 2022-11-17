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
	Apps        UnsetMode
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
func WithAppsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Apps = m
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
	// if len(obj.Apps) > 0 {
	//   ud.SetApps ( obj.Apps)
	// } else {
	ud.setOrUnsetApps(obj.Apps, uo.ResolveUnsetMode(uo.Apps))
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

// ----- apps - array -  [apps]
// SetApps No Remarks
func (ud *UpdateDocument) SetApps(p []App) *UpdateDocument {
	mName := fmt.Sprintf(APPS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetApps No Remarks
func (ud *UpdateDocument) UnsetApps() *UpdateDocument {
	mName := fmt.Sprintf(APPS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetApps No Remarks - here2
func (ud *UpdateDocument) setOrUnsetApps(p []App, um UnsetMode) {

	//----- array\n

	if len(p) > 0 {
		ud.SetApps(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetApps()
		case SetData2Default:
			ud.UnsetApps()
		}
	}
}

func UpdateWithApps(p []App) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetApps(p)
		} else {
			ud.UnsetApps()
		}
	}
}

// ----- [] - struct - App [apps.[]]
// SetAppsI No Remarks
func (ud *UpdateDocument) SetAppsI(ndxI int, p App) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppsI No Remarks
func (ud *UpdateDocument) UnsetAppsI(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppsI No Remarks
func (ud *UpdateDocument) setOrUnsetAppsI(ndxI int, p App, um UnsetMode) {
	if !p.IsZero() {
		ud.SetAppsI(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppsI(ndxI)
		case SetData2Default:
			ud.UnsetAppsI(ndxI)
		}
	}
}

//----- id - string -  [apps.[].id apps.id]

// SetAppsIId No Remarks
func (ud *UpdateDocument) SetAppsIId(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_ID, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppsIId No Remarks
func (ud *UpdateDocument) UnsetAppsIId(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_ID, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppsIId No Remarks
func (ud *UpdateDocument) setOrUnsetAppsIId(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetAppsIId(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppsIId(ndxI)
		case SetData2Default:
			ud.UnsetAppsIId(ndxI)
		}
	}
}

func UpdateWithAppsIId(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppsIId(ndxI, p)
		} else {
			ud.UnsetAppsIId(ndxI)
		}
	}
}

//----- objType - string -  [apps.[].objType apps.objType]

// SetAppsIObjType No Remarks
func (ud *UpdateDocument) SetAppsIObjType(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_OBJTYPE, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppsIObjType No Remarks
func (ud *UpdateDocument) UnsetAppsIObjType(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_OBJTYPE, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppsIObjType No Remarks
func (ud *UpdateDocument) setOrUnsetAppsIObjType(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetAppsIObjType(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppsIObjType(ndxI)
		case SetData2Default:
			ud.UnsetAppsIObjType(ndxI)
		}
	}
}

func UpdateWithAppsIObjType(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppsIObjType(ndxI, p)
		} else {
			ud.UnsetAppsIObjType(ndxI)
		}
	}
}

//----- name - string -  [apps.[].name apps.name]

// SetAppsIName No Remarks
func (ud *UpdateDocument) SetAppsIName(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_NAME, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppsIName No Remarks
func (ud *UpdateDocument) UnsetAppsIName(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_NAME, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppsIName No Remarks
func (ud *UpdateDocument) setOrUnsetAppsIName(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetAppsIName(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppsIName(ndxI)
		case SetData2Default:
			ud.UnsetAppsIName(ndxI)
		}
	}
}

func UpdateWithAppsIName(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppsIName(ndxI, p)
		} else {
			ud.UnsetAppsIName(ndxI)
		}
	}
}

//----- description - string -  [apps.[].description apps.description]

// SetAppsIDescription No Remarks
func (ud *UpdateDocument) SetAppsIDescription(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_DESCRIPTION, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppsIDescription No Remarks
func (ud *UpdateDocument) UnsetAppsIDescription(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_DESCRIPTION, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppsIDescription No Remarks
func (ud *UpdateDocument) setOrUnsetAppsIDescription(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetAppsIDescription(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppsIDescription(ndxI)
		case SetData2Default:
			ud.UnsetAppsIDescription(ndxI)
		}
	}
}

func UpdateWithAppsIDescription(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppsIDescription(ndxI, p)
		} else {
			ud.UnsetAppsIDescription(ndxI)
		}
	}
}

//----- path - string -  [apps.[].path apps.path]

// SetAppsIPath No Remarks
func (ud *UpdateDocument) SetAppsIPath(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_PATH, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppsIPath No Remarks
func (ud *UpdateDocument) UnsetAppsIPath(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_PATH, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppsIPath No Remarks
func (ud *UpdateDocument) setOrUnsetAppsIPath(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetAppsIPath(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppsIPath(ndxI)
		case SetData2Default:
			ud.UnsetAppsIPath(ndxI)
		}
	}
}

func UpdateWithAppsIPath(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAppsIPath(ndxI, p)
		} else {
			ud.UnsetAppsIPath(ndxI)
		}
	}
}

//----- roleRequired - bool -  [apps.[].roleRequired apps.roleRequired]

// SetAppsIRoleRequired No Remarks
func (ud *UpdateDocument) SetAppsIRoleRequired(ndxI int, p bool) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_ROLEREQUIRED, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAppsIRoleRequired
func (ud *UpdateDocument) UnsetAppsIRoleRequired(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(APPS_I_ROLEREQUIRED, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAppsIRoleRequired No Remarks
func (ud *UpdateDocument) setOrUnsetAppsIRoleRequired(ndxI int, p bool, um UnsetMode) {
	if p {
		ud.SetAppsIRoleRequired(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAppsIRoleRequired(ndxI)
		case SetData2Default:
			ud.UnsetAppsIRoleRequired(ndxI)
		}
	}
}

func UpdateWithAppsIRoleRequired(ndxI int, p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetAppsIRoleRequired(ndxI, p)
		} else {
			ud.UnsetAppsIRoleRequired(ndxI)
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
