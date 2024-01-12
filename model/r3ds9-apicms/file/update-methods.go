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
	FileName    UnsetMode
	Container   UnsetMode
	Url         UnsetMode
	ContentType UnsetMode
	Width       UnsetMode
	Height      UnsetMode
	Metadata    UnsetMode
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
func WithFileNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.FileName = m
	}
}
func WithContainerUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Container = m
	}
}
func WithUrlUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Url = m
	}
}
func WithContentTypeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ContentType = m
	}
}
func WithWidthUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Width = m
	}
}
func WithHeightUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Height = m
	}
}
func WithMetadataUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Metadata = m
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
	ud.setOrUnsetFileName(obj.FileName, uo.ResolveUnsetMode(uo.FileName))
	ud.setOrUnsetContainer(obj.Container, uo.ResolveUnsetMode(uo.Container))
	ud.setOrUnsetUrl(obj.Url, uo.ResolveUnsetMode(uo.Url))
	ud.setOrUnsetContentType(obj.ContentType, uo.ResolveUnsetMode(uo.ContentType))
	ud.setOrUnsetWidth(obj.Width, uo.ResolveUnsetMode(uo.Width))
	ud.setOrUnsetHeight(obj.Height, uo.ResolveUnsetMode(uo.Height))
	// if len(obj.Metadata) > 0 {
	//   ud.SetMetadata ( obj.Metadata)
	// } else {
	ud.setOrUnsetMetadata(obj.Metadata, uo.ResolveUnsetMode(uo.Metadata))
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

//----- fileName - string -  [fileName]

// SetFileName No Remarks
func (ud *UpdateDocument) SetFileName(p string) *UpdateDocument {
	mName := fmt.Sprintf(FILENAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFileName No Remarks
func (ud *UpdateDocument) UnsetFileName() *UpdateDocument {
	mName := fmt.Sprintf(FILENAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFileName No Remarks
func (ud *UpdateDocument) setOrUnsetFileName(p string, um UnsetMode) {
	if p != "" {
		ud.SetFileName(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFileName()
		case SetData2Default:
			ud.UnsetFileName()
		}
	}
}

func UpdateWithFileName(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFileName(p)
		} else {
			ud.UnsetFileName()
		}
	}
}

//----- container - string -  [container]

// SetContainer No Remarks
func (ud *UpdateDocument) SetContainer(p string) *UpdateDocument {
	mName := fmt.Sprintf(CONTAINER)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetContainer No Remarks
func (ud *UpdateDocument) UnsetContainer() *UpdateDocument {
	mName := fmt.Sprintf(CONTAINER)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetContainer No Remarks
func (ud *UpdateDocument) setOrUnsetContainer(p string, um UnsetMode) {
	if p != "" {
		ud.SetContainer(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetContainer()
		case SetData2Default:
			ud.UnsetContainer()
		}
	}
}

func UpdateWithContainer(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetContainer(p)
		} else {
			ud.UnsetContainer()
		}
	}
}

//----- url - string -  [url]

// SetUrl No Remarks
func (ud *UpdateDocument) SetUrl(p string) *UpdateDocument {
	mName := fmt.Sprintf(URL)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetUrl No Remarks
func (ud *UpdateDocument) UnsetUrl() *UpdateDocument {
	mName := fmt.Sprintf(URL)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetUrl No Remarks
func (ud *UpdateDocument) setOrUnsetUrl(p string, um UnsetMode) {
	if p != "" {
		ud.SetUrl(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetUrl()
		case SetData2Default:
			ud.UnsetUrl()
		}
	}
}

func UpdateWithUrl(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetUrl(p)
		} else {
			ud.UnsetUrl()
		}
	}
}

//----- contentType - string -  [contentType]

// SetContentType No Remarks
func (ud *UpdateDocument) SetContentType(p string) *UpdateDocument {
	mName := fmt.Sprintf(CONTENTTYPE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetContentType No Remarks
func (ud *UpdateDocument) UnsetContentType() *UpdateDocument {
	mName := fmt.Sprintf(CONTENTTYPE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetContentType No Remarks
func (ud *UpdateDocument) setOrUnsetContentType(p string, um UnsetMode) {
	if p != "" {
		ud.SetContentType(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetContentType()
		case SetData2Default:
			ud.UnsetContentType()
		}
	}
}

func UpdateWithContentType(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetContentType(p)
		} else {
			ud.UnsetContentType()
		}
	}
}

//----- width - int -  [width]

// SetWidth No Remarks
func (ud *UpdateDocument) SetWidth(p int32) *UpdateDocument {
	mName := fmt.Sprintf(WIDTH)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetWidth No Remarks
func (ud *UpdateDocument) UnsetWidth() *UpdateDocument {
	mName := fmt.Sprintf(WIDTH)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetWidth No Remarks
func (ud *UpdateDocument) setOrUnsetWidth(p int32, um UnsetMode) {
	if p != 0 {
		ud.SetWidth(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetWidth()
		case SetData2Default:
			ud.UnsetWidth()
		}
	}
}

// IncWidth No Remarks
func (ud *UpdateDocument) IncWidth(p int32) *UpdateDocument {
	mName := fmt.Sprintf(WIDTH)
	ud.Inc().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func UpdateWithWidth(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetWidth(p)
		} else {
			ud.UnsetWidth()
		}
	}
}

func UpdateWithWidthIncrement(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		ud.IncWidth(p)
	}
}

//----- height - int -  [height]

// SetHeight No Remarks
func (ud *UpdateDocument) SetHeight(p int32) *UpdateDocument {
	mName := fmt.Sprintf(HEIGHT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetHeight No Remarks
func (ud *UpdateDocument) UnsetHeight() *UpdateDocument {
	mName := fmt.Sprintf(HEIGHT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetHeight No Remarks
func (ud *UpdateDocument) setOrUnsetHeight(p int32, um UnsetMode) {
	if p != 0 {
		ud.SetHeight(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetHeight()
		case SetData2Default:
			ud.UnsetHeight()
		}
	}
}

// IncHeight No Remarks
func (ud *UpdateDocument) IncHeight(p int32) *UpdateDocument {
	mName := fmt.Sprintf(HEIGHT)
	ud.Inc().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func UpdateWithHeight(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetHeight(p)
		} else {
			ud.UnsetHeight()
		}
	}
}

func UpdateWithHeightIncrement(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		ud.IncHeight(p)
	}
}

// SetMetadata dd
func (ud *UpdateDocument) SetMetadata(p map[string]interface{}) *UpdateDocument {
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

// setOrUnsetMetadata No Remarks - here2
func (ud *UpdateDocument) setOrUnsetMetadata(p map[string]interface{}, um UnsetMode) {

	//----- map\n

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

func UpdateWithMetadata(p map[string]interface{}) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetMetadata(p)
		} else {
			ud.UnsetMetadata()
		}
	}
}

// ----- %s - string -  [metadata.%s]
// SetMetadataS No Remarks
func (ud *UpdateDocument) SetMetadataS(keyS string, p interface{}) *UpdateDocument {
	mName := fmt.Sprintf(METADATA_S, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetMetadataS No Remarks
func (ud *UpdateDocument) UnsetMetadataS(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(METADATA_S, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetMetadataS No Remarks
func (ud *UpdateDocument) setOrUnsetMetadataS(keyS string, p interface{}, um UnsetMode) {
	// Warining.... should not get here
	if p != "" {
		ud.SetMetadataS(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetMetadataS(keyS)
		case SetData2Default:
			ud.UnsetMetadataS(keyS)
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
