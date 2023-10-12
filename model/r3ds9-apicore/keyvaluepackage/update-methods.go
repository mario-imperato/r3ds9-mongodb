package keyvaluepackage

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
	Name        UnsetMode
	Scope       UnsetMode
	ObjType     UnsetMode
	Category    UnsetMode
	Issystem    UnsetMode
	Description UnsetMode
	Inherited   UnsetMode
	Properties  UnsetMode
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
func WithNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Name = m
	}
}
func WithScopeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Scope = m
	}
}
func WithObjTypeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ObjType = m
	}
}
func WithCategoryUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Category = m
	}
}
func WithIssystemUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Issystem = m
	}
}
func WithDescriptionUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Description = m
	}
}
func WithInheritedUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Inherited = m
	}
}
func WithPropertiesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Properties = m
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
func GetUpdateDocument(obj *KeyValuePackage, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetName(obj.Name, uo.ResolveUnsetMode(uo.Name))
	ud.setOrUnsetScope(obj.Scope, uo.ResolveUnsetMode(uo.Scope))
	ud.setOrUnsetObjType(obj.ObjType, uo.ResolveUnsetMode(uo.ObjType))
	ud.setOrUnsetCategory(obj.Category, uo.ResolveUnsetMode(uo.Category))
	ud.setOrUnsetIssystem(obj.Issystem, uo.ResolveUnsetMode(uo.Issystem))
	ud.setOrUnsetDescription(obj.Description, uo.ResolveUnsetMode(uo.Description))
	ud.setOrUnsetInherited(obj.Inherited, uo.ResolveUnsetMode(uo.Inherited))
	// if len(obj.Properties) > 0 {
	//   ud.SetProperties ( obj.Properties)
	// } else {
	ud.setOrUnsetProperties(obj.Properties, uo.ResolveUnsetMode(uo.Properties))
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

//----- name - string -  [name]

// SetName No Remarks
func (ud *UpdateDocument) SetName(p string) *UpdateDocument {
	mName := fmt.Sprintf(NAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetName No Remarks
func (ud *UpdateDocument) UnsetName() *UpdateDocument {
	mName := fmt.Sprintf(NAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetName No Remarks
func (ud *UpdateDocument) setOrUnsetName(p string, um UnsetMode) {
	if p != "" {
		ud.SetName(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetName()
		case SetData2Default:
			ud.UnsetName()
		}
	}
}

func UpdateWithName(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetName(p)
		} else {
			ud.UnsetName()
		}
	}
}

//----- scope - string -  [scope]

// SetScope No Remarks
func (ud *UpdateDocument) SetScope(p string) *UpdateDocument {
	mName := fmt.Sprintf(SCOPE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetScope No Remarks
func (ud *UpdateDocument) UnsetScope() *UpdateDocument {
	mName := fmt.Sprintf(SCOPE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetScope No Remarks
func (ud *UpdateDocument) setOrUnsetScope(p string, um UnsetMode) {
	if p != "" {
		ud.SetScope(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetScope()
		case SetData2Default:
			ud.UnsetScope()
		}
	}
}

func UpdateWithScope(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetScope(p)
		} else {
			ud.UnsetScope()
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

//----- category - string -  [category]

// SetCategory No Remarks
func (ud *UpdateDocument) SetCategory(p string) *UpdateDocument {
	mName := fmt.Sprintf(CATEGORY)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCategory No Remarks
func (ud *UpdateDocument) UnsetCategory() *UpdateDocument {
	mName := fmt.Sprintf(CATEGORY)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCategory No Remarks
func (ud *UpdateDocument) setOrUnsetCategory(p string, um UnsetMode) {
	if p != "" {
		ud.SetCategory(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCategory()
		case SetData2Default:
			ud.UnsetCategory()
		}
	}
}

func UpdateWithCategory(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCategory(p)
		} else {
			ud.UnsetCategory()
		}
	}
}

//----- issystem - bool -  [issystem]

// SetIssystem No Remarks
func (ud *UpdateDocument) SetIssystem(p bool) *UpdateDocument {
	mName := fmt.Sprintf(ISSYSTEM)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIssystem
func (ud *UpdateDocument) UnsetIssystem() *UpdateDocument {
	mName := fmt.Sprintf(ISSYSTEM)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIssystem No Remarks
func (ud *UpdateDocument) setOrUnsetIssystem(p bool, um UnsetMode) {
	if p {
		ud.SetIssystem(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIssystem()
		case SetData2Default:
			ud.UnsetIssystem()
		}
	}
}

func UpdateWithIssystem(p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetIssystem(p)
		} else {
			ud.UnsetIssystem()
		}
	}
}

//----- description - string -  [description]

// SetDescription No Remarks
func (ud *UpdateDocument) SetDescription(p string) *UpdateDocument {
	mName := fmt.Sprintf(DESCRIPTION)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetDescription No Remarks
func (ud *UpdateDocument) UnsetDescription() *UpdateDocument {
	mName := fmt.Sprintf(DESCRIPTION)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetDescription No Remarks
func (ud *UpdateDocument) setOrUnsetDescription(p string, um UnsetMode) {
	if p != "" {
		ud.SetDescription(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDescription()
		case SetData2Default:
			ud.UnsetDescription()
		}
	}
}

func UpdateWithDescription(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetDescription(p)
		} else {
			ud.UnsetDescription()
		}
	}
}

//----- inherited - bool -  [inherited]

// SetInherited No Remarks
func (ud *UpdateDocument) SetInherited(p bool) *UpdateDocument {
	mName := fmt.Sprintf(INHERITED)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetInherited
func (ud *UpdateDocument) UnsetInherited() *UpdateDocument {
	mName := fmt.Sprintf(INHERITED)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetInherited No Remarks
func (ud *UpdateDocument) setOrUnsetInherited(p bool, um UnsetMode) {
	if p {
		ud.SetInherited(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetInherited()
		case SetData2Default:
			ud.UnsetInherited()
		}
	}
}

func UpdateWithInherited(p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetInherited(p)
		} else {
			ud.UnsetInherited()
		}
	}
}

// ----- properties - array -  [properties]
// SetProperties No Remarks
func (ud *UpdateDocument) SetProperties(p []KeyValue) *UpdateDocument {
	mName := fmt.Sprintf(PROPERTIES)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetProperties No Remarks
func (ud *UpdateDocument) UnsetProperties() *UpdateDocument {
	mName := fmt.Sprintf(PROPERTIES)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetProperties No Remarks - here2
func (ud *UpdateDocument) setOrUnsetProperties(p []KeyValue, um UnsetMode) {

	//----- array\n

	if len(p) > 0 {
		ud.SetProperties(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetProperties()
		case SetData2Default:
			ud.UnsetProperties()
		}
	}
}

func UpdateWithProperties(p []KeyValue) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetProperties(p)
		} else {
			ud.UnsetProperties()
		}
	}
}

// ----- [] - struct - KeyValue [properties.[]]
// SetPropertiesI No Remarks
func (ud *UpdateDocument) SetPropertiesI(ndxI int, p KeyValue) *UpdateDocument {
	mName := fmt.Sprintf(PROPERTIES_I, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetPropertiesI No Remarks
func (ud *UpdateDocument) UnsetPropertiesI(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(PROPERTIES_I, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetPropertiesI No Remarks
func (ud *UpdateDocument) setOrUnsetPropertiesI(ndxI int, p KeyValue, um UnsetMode) {
	if !p.IsZero() {
		ud.SetPropertiesI(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetPropertiesI(ndxI)
		case SetData2Default:
			ud.UnsetPropertiesI(ndxI)
		}
	}
}

//----- key - string -  [properties.[].key properties.key]

// SetPropertiesIKey No Remarks
func (ud *UpdateDocument) SetPropertiesIKey(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(PROPERTIES_I_KEY, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetPropertiesIKey No Remarks
func (ud *UpdateDocument) UnsetPropertiesIKey(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(PROPERTIES_I_KEY, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetPropertiesIKey No Remarks
func (ud *UpdateDocument) setOrUnsetPropertiesIKey(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetPropertiesIKey(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetPropertiesIKey(ndxI)
		case SetData2Default:
			ud.UnsetPropertiesIKey(ndxI)
		}
	}
}

func UpdateWithPropertiesIKey(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetPropertiesIKey(ndxI, p)
		} else {
			ud.UnsetPropertiesIKey(ndxI)
		}
	}
}

//----- value - string -  [properties.[].value properties.value]

// SetPropertiesIValue No Remarks
func (ud *UpdateDocument) SetPropertiesIValue(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(PROPERTIES_I_VALUE, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetPropertiesIValue No Remarks
func (ud *UpdateDocument) UnsetPropertiesIValue(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(PROPERTIES_I_VALUE, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetPropertiesIValue No Remarks
func (ud *UpdateDocument) setOrUnsetPropertiesIValue(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetPropertiesIValue(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetPropertiesIValue(ndxI)
		case SetData2Default:
			ud.UnsetPropertiesIValue(ndxI)
		}
	}
}

func UpdateWithPropertiesIValue(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetPropertiesIValue(ndxI, p)
		} else {
			ud.UnsetPropertiesIValue(ndxI)
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
