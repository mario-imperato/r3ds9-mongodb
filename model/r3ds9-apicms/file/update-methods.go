package file

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/mario-imperato/r3ds9-mongodb/model/commons"
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
	Fn          UnsetMode
	Descr       UnsetMode
	Role        UnsetMode
	EntRef      UnsetMode
	Metadata    UnsetMode
	Vrnts       UnsetMode
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
func WithFnUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Fn = m
	}
}
func WithDescrUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Descr = m
	}
}
func WithRoleUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Role = m
	}
}
func WithEntRefUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.EntRef = m
	}
}
func WithMetadataUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Metadata = m
	}
}
func WithVrntsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Vrnts = m
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
func GetUpdateDocument(obj *File, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetFn(obj.Fn, uo.ResolveUnsetMode(uo.Fn))
	ud.setOrUnsetDescr(obj.Descr, uo.ResolveUnsetMode(uo.Descr))
	ud.setOrUnsetRole(obj.Role, uo.ResolveUnsetMode(uo.Role))
	// if !obj.EntRef.IsZero() {
	//   ud.SetEntRef ( obj.EntRef)
	// } else {
	ud.setOrUnsetEntRef(obj.EntRef, uo.ResolveUnsetMode(uo.EntRef))
	// }
	// if len(obj.Metadata) > 0 {
	//   ud.SetMetadata ( obj.Metadata)
	// } else {
	ud.setOrUnsetMetadata(obj.Metadata, uo.ResolveUnsetMode(uo.Metadata))
	// }
	// if len(obj.Vrnts) > 0 {
	//   ud.SetVrnts ( obj.Vrnts)
	// } else {
	ud.setOrUnsetVrnts(obj.Vrnts, uo.ResolveUnsetMode(uo.Vrnts))
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

//----- fn - string -  [fn]

// SetFn No Remarks
func (ud *UpdateDocument) SetFn(p string) *UpdateDocument {
	mName := fmt.Sprintf(FN)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFn No Remarks
func (ud *UpdateDocument) UnsetFn() *UpdateDocument {
	mName := fmt.Sprintf(FN)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFn No Remarks
func (ud *UpdateDocument) setOrUnsetFn(p string, um UnsetMode) {
	if p != "" {
		ud.SetFn(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFn()
		case SetData2Default:
			ud.UnsetFn()
		}
	}
}

func UpdateWithFn(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFn(p)
		} else {
			ud.UnsetFn()
		}
	}
}

//----- descr - string -  [descr]

// SetDescr No Remarks
func (ud *UpdateDocument) SetDescr(p string) *UpdateDocument {
	mName := fmt.Sprintf(DESCR)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetDescr No Remarks
func (ud *UpdateDocument) UnsetDescr() *UpdateDocument {
	mName := fmt.Sprintf(DESCR)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetDescr No Remarks
func (ud *UpdateDocument) setOrUnsetDescr(p string, um UnsetMode) {
	if p != "" {
		ud.SetDescr(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDescr()
		case SetData2Default:
			ud.UnsetDescr()
		}
	}
}

func UpdateWithDescr(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetDescr(p)
		} else {
			ud.UnsetDescr()
		}
	}
}

//----- role - string -  [role]

// SetRole No Remarks
func (ud *UpdateDocument) SetRole(p string) *UpdateDocument {
	mName := fmt.Sprintf(ROLE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRole No Remarks
func (ud *UpdateDocument) UnsetRole() *UpdateDocument {
	mName := fmt.Sprintf(ROLE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRole No Remarks
func (ud *UpdateDocument) setOrUnsetRole(p string, um UnsetMode) {
	if p != "" {
		ud.SetRole(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRole()
		case SetData2Default:
			ud.UnsetRole()
		}
	}
}

func UpdateWithRole(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRole(p)
		} else {
			ud.UnsetRole()
		}
	}
}

// ----- entRef - struct - EntRefStruct [entRef]
// SetEntRef No Remarks
func (ud *UpdateDocument) SetEntRef(p EntRefStruct) *UpdateDocument {
	mName := fmt.Sprintf(ENTREF)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEntRef No Remarks
func (ud *UpdateDocument) UnsetEntRef() *UpdateDocument {
	mName := fmt.Sprintf(ENTREF)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEntRef No Remarks - here2
func (ud *UpdateDocument) setOrUnsetEntRef(p EntRefStruct, um UnsetMode) {

	//----- struct\n

	if !p.IsZero() {
		ud.SetEntRef(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEntRef()
		case SetData2Default:
			ud.UnsetEntRef()
		}
	}
}

func UpdateWithEntRef(p EntRefStruct) UpdateOption {
	return func(ud *UpdateDocument) {

		if !p.IsZero() {
			ud.SetEntRef(p)
		} else {
			ud.UnsetEntRef()
		}
	}
}

//----- dom - string -  [entRef.dom]

// SetEntRefDom No Remarks
func (ud *UpdateDocument) SetEntRefDom(p string) *UpdateDocument {
	mName := fmt.Sprintf(ENTREF_DOM)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEntRefDom No Remarks
func (ud *UpdateDocument) UnsetEntRefDom() *UpdateDocument {
	mName := fmt.Sprintf(ENTREF_DOM)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEntRefDom No Remarks
func (ud *UpdateDocument) setOrUnsetEntRefDom(p string, um UnsetMode) {
	if p != "" {
		ud.SetEntRefDom(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEntRefDom()
		case SetData2Default:
			ud.UnsetEntRefDom()
		}
	}
}

func UpdateWithEntRefDom(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetEntRefDom(p)
		} else {
			ud.UnsetEntRefDom()
		}
	}
}

//----- ns - string -  [entRef.ns]

// SetEntRefNs No Remarks
func (ud *UpdateDocument) SetEntRefNs(p string) *UpdateDocument {
	mName := fmt.Sprintf(ENTREF_NS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEntRefNs No Remarks
func (ud *UpdateDocument) UnsetEntRefNs() *UpdateDocument {
	mName := fmt.Sprintf(ENTREF_NS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEntRefNs No Remarks
func (ud *UpdateDocument) setOrUnsetEntRefNs(p string, um UnsetMode) {
	if p != "" {
		ud.SetEntRefNs(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEntRefNs()
		case SetData2Default:
			ud.UnsetEntRefNs()
		}
	}
}

func UpdateWithEntRefNs(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetEntRefNs(p)
		} else {
			ud.UnsetEntRefNs()
		}
	}
}

//----- entType - string -  [entRef.entType]

// SetEntRefEntType No Remarks
func (ud *UpdateDocument) SetEntRefEntType(p string) *UpdateDocument {
	mName := fmt.Sprintf(ENTREF_ENTTYPE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEntRefEntType No Remarks
func (ud *UpdateDocument) UnsetEntRefEntType() *UpdateDocument {
	mName := fmt.Sprintf(ENTREF_ENTTYPE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEntRefEntType No Remarks
func (ud *UpdateDocument) setOrUnsetEntRefEntType(p string, um UnsetMode) {
	if p != "" {
		ud.SetEntRefEntType(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEntRefEntType()
		case SetData2Default:
			ud.UnsetEntRefEntType()
		}
	}
}

func UpdateWithEntRefEntType(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetEntRefEntType(p)
		} else {
			ud.UnsetEntRefEntType()
		}
	}
}

//----- entId - string -  [entRef.entId]

// SetEntRefEntId No Remarks
func (ud *UpdateDocument) SetEntRefEntId(p string) *UpdateDocument {
	mName := fmt.Sprintf(ENTREF_ENTID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEntRefEntId No Remarks
func (ud *UpdateDocument) UnsetEntRefEntId() *UpdateDocument {
	mName := fmt.Sprintf(ENTREF_ENTID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEntRefEntId No Remarks
func (ud *UpdateDocument) setOrUnsetEntRefEntId(p string, um UnsetMode) {
	if p != "" {
		ud.SetEntRefEntId(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEntRefEntId()
		case SetData2Default:
			ud.UnsetEntRefEntId()
		}
	}
}

func UpdateWithEntRefEntId(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetEntRefEntId(p)
		} else {
			ud.UnsetEntRefEntId()
		}
	}
}

//----- metadata - document -  [metadata]

// SetMetadata No Remarks
func (ud *UpdateDocument) SetMetadata(p bson.M) *UpdateDocument {
	mName := fmt.Sprintf(METADATA)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetMetadata No Remarks
func (ud *UpdateDocument) UnsetMetadata() *UpdateDocument {
	mName := fmt.Sprintf(METADATA)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetMetadata No Remarks
func (ud *UpdateDocument) setOrUnsetMetadata(p bson.M, um UnsetMode) {
	if len(p) > 0 {
		ud.SetMetadata(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetMetadata()
		case SetData2Default:
			ud.UnsetMetadata()
		}
	}
}

func UpdateWithMetadata(p bson.M) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetMetadata(p)
		} else {
			ud.UnsetMetadata()
		}
	}
}

// ----- vrnts - array -  [vrnts]
// SetVrnts No Remarks
func (ud *UpdateDocument) SetVrnts(p []Variant) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrnts No Remarks
func (ud *UpdateDocument) UnsetVrnts() *UpdateDocument {
	mName := fmt.Sprintf(VRNTS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrnts No Remarks - here2
func (ud *UpdateDocument) setOrUnsetVrnts(p []Variant, um UnsetMode) {

	//----- array\n

	if len(p) > 0 {
		ud.SetVrnts(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrnts()
		case SetData2Default:
			ud.UnsetVrnts()
		}
	}
}

func UpdateWithVrnts(p []Variant) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetVrnts(p)
		} else {
			ud.UnsetVrnts()
		}
	}
}

// ----- [] - struct - Variant [vrnts.[]]
// SetVrntsI No Remarks
func (ud *UpdateDocument) SetVrntsI(ndxI int, p Variant) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsI No Remarks
func (ud *UpdateDocument) UnsetVrntsI(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsI No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsI(ndxI int, p Variant, um UnsetMode) {
	if !p.IsZero() {
		ud.SetVrntsI(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsI(ndxI)
		case SetData2Default:
			ud.UnsetVrntsI(ndxI)
		}
	}
}

//----- ct - string -  [vrnts.[].ct vrnts.ct]

// SetVrntsICt No Remarks
func (ud *UpdateDocument) SetVrntsICt(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_CT, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsICt No Remarks
func (ud *UpdateDocument) UnsetVrntsICt(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_CT, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsICt No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsICt(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetVrntsICt(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsICt(ndxI)
		case SetData2Default:
			ud.UnsetVrntsICt(ndxI)
		}
	}
}

func UpdateWithVrntsICt(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetVrntsICt(ndxI, p)
		} else {
			ud.UnsetVrntsICt(ndxI)
		}
	}
}

//----- wd - int -  [vrnts.[].wd vrnts.wd]

// SetVrntsIWd No Remarks
func (ud *UpdateDocument) SetVrntsIWd(ndxI int, p int32) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_WD, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsIWd No Remarks
func (ud *UpdateDocument) UnsetVrntsIWd(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_WD, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsIWd No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsIWd(ndxI int, p int32, um UnsetMode) {
	if p != 0 {
		ud.SetVrntsIWd(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsIWd(ndxI)
		case SetData2Default:
			ud.UnsetVrntsIWd(ndxI)
		}
	}
}

// IncVrntsIWd No Remarks
func (ud *UpdateDocument) IncVrntsIWd(ndxI int, p int32) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_WD, ndxI)
	ud.Inc().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func UpdateWithVrntsIWd(ndxI int, p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetVrntsIWd(ndxI, p)
		} else {
			ud.UnsetVrntsIWd(ndxI)
		}
	}
}

func UpdateWithVrntsIWdIncrement(ndxI int, p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		ud.IncVrntsIWd(ndxI, p)
	}
}

//----- ht - int -  [vrnts.[].ht vrnts.ht]

// SetVrntsIHt No Remarks
func (ud *UpdateDocument) SetVrntsIHt(ndxI int, p int32) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_HT, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsIHt No Remarks
func (ud *UpdateDocument) UnsetVrntsIHt(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_HT, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsIHt No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsIHt(ndxI int, p int32, um UnsetMode) {
	if p != 0 {
		ud.SetVrntsIHt(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsIHt(ndxI)
		case SetData2Default:
			ud.UnsetVrntsIHt(ndxI)
		}
	}
}

// IncVrntsIHt No Remarks
func (ud *UpdateDocument) IncVrntsIHt(ndxI int, p int32) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_HT, ndxI)
	ud.Inc().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func UpdateWithVrntsIHt(ndxI int, p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetVrntsIHt(ndxI, p)
		} else {
			ud.UnsetVrntsIHt(ndxI)
		}
	}
}

func UpdateWithVrntsIHtIncrement(ndxI int, p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		ud.IncVrntsIHt(ndxI, p)
	}
}

//----- lks - string -  [vrnts.[].lks vrnts.lks]

// SetVrntsILks No Remarks
func (ud *UpdateDocument) SetVrntsILks(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_LKS, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsILks No Remarks
func (ud *UpdateDocument) UnsetVrntsILks(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_LKS, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsILks No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsILks(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetVrntsILks(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsILks(ndxI)
		case SetData2Default:
			ud.UnsetVrntsILks(ndxI)
		}
	}
}

func UpdateWithVrntsILks(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetVrntsILks(ndxI, p)
		} else {
			ud.UnsetVrntsILks(ndxI)
		}
	}
}

//----- bln - string -  [vrnts.[].bln vrnts.bln]

// SetVrntsIBln No Remarks
func (ud *UpdateDocument) SetVrntsIBln(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_BLN, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsIBln No Remarks
func (ud *UpdateDocument) UnsetVrntsIBln(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_BLN, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsIBln No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsIBln(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetVrntsIBln(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsIBln(ndxI)
		case SetData2Default:
			ud.UnsetVrntsIBln(ndxI)
		}
	}
}

func UpdateWithVrntsIBln(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetVrntsIBln(ndxI, p)
		} else {
			ud.UnsetVrntsIBln(ndxI)
		}
	}
}

//----- cnt - string -  [vrnts.[].cnt vrnts.cnt]

// SetVrntsICnt No Remarks
func (ud *UpdateDocument) SetVrntsICnt(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_CNT, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsICnt No Remarks
func (ud *UpdateDocument) UnsetVrntsICnt(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_CNT, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsICnt No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsICnt(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetVrntsICnt(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsICnt(ndxI)
		case SetData2Default:
			ud.UnsetVrntsICnt(ndxI)
		}
	}
}

func UpdateWithVrntsICnt(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetVrntsICnt(ndxI, p)
		} else {
			ud.UnsetVrntsICnt(ndxI)
		}
	}
}

//----- url - string -  [vrnts.[].url vrnts.url]

// SetVrntsIUrl No Remarks
func (ud *UpdateDocument) SetVrntsIUrl(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_URL, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsIUrl No Remarks
func (ud *UpdateDocument) UnsetVrntsIUrl(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_URL, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsIUrl No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsIUrl(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetVrntsIUrl(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsIUrl(ndxI)
		case SetData2Default:
			ud.UnsetVrntsIUrl(ndxI)
		}
	}
}

func UpdateWithVrntsIUrl(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetVrntsIUrl(ndxI, p)
		} else {
			ud.UnsetVrntsIUrl(ndxI)
		}
	}
}

//----- role - string -  [vrnts.[].role vrnts.role]

// SetVrntsIRole No Remarks
func (ud *UpdateDocument) SetVrntsIRole(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_ROLE, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrntsIRole No Remarks
func (ud *UpdateDocument) UnsetVrntsIRole(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(VRNTS_I_ROLE, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrntsIRole No Remarks
func (ud *UpdateDocument) setOrUnsetVrntsIRole(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetVrntsIRole(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrntsIRole(ndxI)
		case SetData2Default:
			ud.UnsetVrntsIRole(ndxI)
		}
	}
}

func UpdateWithVrntsIRole(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetVrntsIRole(ndxI, p)
		} else {
			ud.UnsetVrntsIRole(ndxI)
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
