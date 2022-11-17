package user

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apigtw/commons"
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
	OId         UnsetMode
	Nickname    UnsetMode
	ObjType     UnsetMode
	Firstname   UnsetMode
	Lastnamw    UnsetMode
	Password    UnsetMode
	Roles       UnsetMode
	Sysinfo     UnsetMode
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
func WithOIdUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.OId = m
	}
}
func WithNicknameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Nickname = m
	}
}
func WithObjTypeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ObjType = m
	}
}
func WithFirstnameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Firstname = m
	}
}
func WithLastnamwUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Lastnamw = m
	}
}
func WithPasswordUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Password = m
	}
}
func WithRolesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Roles = m
	}
}
func WithSysinfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Sysinfo = m
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
func GetUpdateDocument(obj *User, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetNickname(obj.Nickname, uo.ResolveUnsetMode(uo.Nickname))
	ud.setOrUnsetObjType(obj.ObjType, uo.ResolveUnsetMode(uo.ObjType))
	ud.setOrUnsetFirstname(obj.Firstname, uo.ResolveUnsetMode(uo.Firstname))
	ud.setOrUnsetLastnamw(obj.Lastnamw, uo.ResolveUnsetMode(uo.Lastnamw))
	ud.setOrUnsetPassword(obj.Password, uo.ResolveUnsetMode(uo.Password))
	// if len(obj.Roles) > 0 {
	//   ud.SetRoles ( obj.Roles)
	// } else {
	ud.setOrUnsetRoles(obj.Roles, uo.ResolveUnsetMode(uo.Roles))
	// }
	// if !obj.Sysinfo.IsZero() {
	//   ud.SetSysinfo ( obj.Sysinfo)
	// } else {
	ud.setOrUnsetSysinfo(obj.Sysinfo, uo.ResolveUnsetMode(uo.Sysinfo))
	// }

	return ud
}

//----- oId - object-id -  [oId]

// SetOId No Remarks
func (ud *UpdateDocument) SetOId(p primitive.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetOId No Remarks
func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetOId No Remarks
func (ud *UpdateDocument) setOrUnsetOId(p primitive.ObjectID, um UnsetMode) {
	if !p.IsZero() {
		ud.SetOId(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetOId()
		case SetData2Default:
			ud.UnsetOId()
		}
	}
}

//----- nickname - string -  [nickname]

// SetNickname No Remarks
func (ud *UpdateDocument) SetNickname(p string) *UpdateDocument {
	mName := fmt.Sprintf(NICKNAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetNickname No Remarks
func (ud *UpdateDocument) UnsetNickname() *UpdateDocument {
	mName := fmt.Sprintf(NICKNAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetNickname No Remarks
func (ud *UpdateDocument) setOrUnsetNickname(p string, um UnsetMode) {
	if p != "" {
		ud.SetNickname(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetNickname()
		case SetData2Default:
			ud.UnsetNickname()
		}
	}
}

func UpdateWithNickname(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetNickname(p)
		} else {
			ud.UnsetNickname()
		}
	}
}

//----- objType - string -  [objType]

// SetObjType No Remarks
func (ud *UpdateDocument) SetObjType(p string) *UpdateDocument {
	mName := fmt.Sprintf(OBJTYPE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetObjType No Remarks
func (ud *UpdateDocument) UnsetObjType() *UpdateDocument {
	mName := fmt.Sprintf(OBJTYPE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetObjType No Remarks
func (ud *UpdateDocument) setOrUnsetObjType(p string, um UnsetMode) {
	if p != "" {
		ud.SetObjType(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetObjType()
		case SetData2Default:
			ud.UnsetObjType()
		}
	}
}

func UpdateWithObjType(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetObjType(p)
		} else {
			ud.UnsetObjType()
		}
	}
}

//----- firstname - string -  [firstname]

// SetFirstname No Remarks
func (ud *UpdateDocument) SetFirstname(p string) *UpdateDocument {
	mName := fmt.Sprintf(FIRSTNAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFirstname No Remarks
func (ud *UpdateDocument) UnsetFirstname() *UpdateDocument {
	mName := fmt.Sprintf(FIRSTNAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFirstname No Remarks
func (ud *UpdateDocument) setOrUnsetFirstname(p string, um UnsetMode) {
	if p != "" {
		ud.SetFirstname(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFirstname()
		case SetData2Default:
			ud.UnsetFirstname()
		}
	}
}

func UpdateWithFirstname(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFirstname(p)
		} else {
			ud.UnsetFirstname()
		}
	}
}

//----- lastnamw - string -  [lastnamw]

// SetLastnamw No Remarks
func (ud *UpdateDocument) SetLastnamw(p string) *UpdateDocument {
	mName := fmt.Sprintf(LASTNAMW)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLastnamw No Remarks
func (ud *UpdateDocument) UnsetLastnamw() *UpdateDocument {
	mName := fmt.Sprintf(LASTNAMW)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLastnamw No Remarks
func (ud *UpdateDocument) setOrUnsetLastnamw(p string, um UnsetMode) {
	if p != "" {
		ud.SetLastnamw(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLastnamw()
		case SetData2Default:
			ud.UnsetLastnamw()
		}
	}
}

func UpdateWithLastnamw(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLastnamw(p)
		} else {
			ud.UnsetLastnamw()
		}
	}
}

//----- password - string -  [password]

// SetPassword No Remarks
func (ud *UpdateDocument) SetPassword(p string) *UpdateDocument {
	mName := fmt.Sprintf(PASSWORD)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetPassword No Remarks
func (ud *UpdateDocument) UnsetPassword() *UpdateDocument {
	mName := fmt.Sprintf(PASSWORD)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetPassword No Remarks
func (ud *UpdateDocument) setOrUnsetPassword(p string, um UnsetMode) {
	if p != "" {
		ud.SetPassword(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetPassword()
		case SetData2Default:
			ud.UnsetPassword()
		}
	}
}

func UpdateWithPassword(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetPassword(p)
		} else {
			ud.UnsetPassword()
		}
	}
}

// ----- roles - array -  [roles]
// SetRoles No Remarks
func (ud *UpdateDocument) SetRoles(p []commons.UserRole) *UpdateDocument {
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
func (ud *UpdateDocument) setOrUnsetRoles(p []commons.UserRole, um UnsetMode) {

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

func UpdateWithRoles(p []commons.UserRole) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetRoles(p)
		} else {
			ud.UnsetRoles()
		}
	}
}

// ----- [] - ref-struct -  [roles.[]]
// SetRolesI No Remarks
func (ud *UpdateDocument) SetRolesI(ndxI int, p commons.UserRole) *UpdateDocument {
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
func (ud *UpdateDocument) setOrUnsetRolesI(ndxI int, p commons.UserRole, um UnsetMode) {
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

// ----- sysinfo - ref-struct -  [sysinfo]
// SetSysinfo No Remarks
func (ud *UpdateDocument) SetSysinfo(p commons.SysInfo) *UpdateDocument {
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
func (ud *UpdateDocument) setOrUnsetSysinfo(p commons.SysInfo, um UnsetMode) {

	//----- ref-struct\n

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

func UpdateWithSysinfo(p commons.SysInfo) UpdateOption {
	return func(ud *UpdateDocument) {

		if !p.IsZero() {
			ud.SetSysinfo(p)
		} else {
			ud.UnsetSysinfo()
		}
	}
}
