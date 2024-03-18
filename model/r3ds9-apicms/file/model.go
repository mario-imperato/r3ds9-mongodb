package file

//const (
//	OID               = "_id"
//	FN                = "fn"
//	DESCR             = "descr"
//	ROLE              = "role"
//	ENTREFS           = "entRefs"
//	ENTREFS_I         = "entRefs.%d"
//	ENTREFS_I_DOM     = "entRefs.%d.dom"
//	ENTREFS_DOM       = "entRefs.dom"
//	ENTREFS_I_NS      = "entRefs.%d.ns"
//	ENTREFS_NS        = "entRefs.ns"
//	ENTREFS_I_ENTTYPE = "entRefs.%d.entType"
//	ENTREFS_ENTTYPE   = "entRefs.entType"
//	ENTREFS_I_ENTID   = "entRefs.%d.entId"
//	ENTREFS_ENTID     = "entRefs.entId"
//	METADATA          = "metadata"
//	VRNTS             = "vrnts"
//	VRNTS_I           = "vrnts.%d"
//	VRNTS_I_CT        = "vrnts.%d.ct"
//	VRNTS_CT          = "vrnts.ct"
//	VRNTS_I_WD        = "vrnts.%d.wd"
//	VRNTS_WD          = "vrnts.wd"
//	VRNTS_I_HT        = "vrnts.%d.ht"
//	VRNTS_HT          = "vrnts.ht"
//	VRNTS_I_LKS       = "vrnts.%d.lks"
//	VRNTS_LKS         = "vrnts.lks"
//	VRNTS_I_BLN       = "vrnts.%d.bln"
//	VRNTS_BLN         = "vrnts.bln"
//	VRNTS_I_CNT       = "vrnts.%d.cnt"
//	VRNTS_CNT         = "vrnts.cnt"
//	VRNTS_I_URL       = "vrnts.%d.url"
//	VRNTS_URL         = "vrnts.url"
//	VRNTS_I_ROLE      = "vrnts.%d.role"
//	VRNTS_ROLE        = "vrnts.role"
//	SYSINFO           = "sysinfo"
//)
//
//type File struct {
//	OId      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
//	Fn       string             `json:"fn,omitempty" bson:"fn,omitempty"`
//	Descr    string             `json:"descr,omitempty" bson:"descr,omitempty"`
//	Role     string             `json:"role,omitempty" bson:"role,omitempty"`
//	EntRefs  []EntRefsStruct    `json:"entRefs,omitempty" bson:"entRefs,omitempty"`
//	Metadata bson.M             `json:"metadata,omitempty" bson:"metadata,omitempty"`
//	Vrnts    []Variant          `json:"vrnts,omitempty" bson:"vrnts,omitempty"`
//	Sysinfo  commons.SysInfo    `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
//}
//
//func (s File) IsZero() bool {
//	/*
//	   if s.OId == primitive.NilObjectID {
//	       return false
//	   }
//	   if s.Fn == "" {
//	       return false
//	   }
//	   if s.Descr == "" {
//	       return false
//	   }
//	   if s.Role == "" {
//	       return false
//	   }
//	   if len(s.EntRefs) == 0 {
//	       return false
//	   }
//	   if len(s.Metadata) == 0 {
//	       return false
//	   }
//	   if len(s.Vrnts) == 0 {
//	       return false
//	   }
//	   if s.Sysinfo.IsZero() {
//	       return false
//	   }
//	       return true
//	*/
//
//	return s.OId == primitive.NilObjectID && s.Fn == "" && s.Descr == "" && s.Role == "" && len(s.EntRefs) == 0 && len(s.Metadata) == 0 && len(s.Vrnts) == 0 && s.Sysinfo.IsZero()
//}
//
//type EntRefsStruct struct {
//	Dom     string `json:"dom,omitempty" bson:"dom,omitempty"`
//	Ns      string `json:"ns,omitempty" bson:"ns,omitempty"`
//	EntType string `json:"entType,omitempty" bson:"entType,omitempty"`
//	EntId   string `json:"entId,omitempty" bson:"entId,omitempty"`
//}
//
//func (s EntRefsStruct) IsZero() bool {
//	/*
//	   if s.Dom == "" {
//	       return false
//	   }
//	   if s.Ns == "" {
//	       return false
//	   }
//	   if s.EntType == "" {
//	       return false
//	   }
//	   if s.EntId == "" {
//	       return false
//	   }
//	       return true
//	*/
//	return s.Dom == "" && s.Ns == "" && s.EntType == "" && s.EntId == ""
//}
//
//type Variant struct {
//	Ct   string `json:"ct,omitempty" bson:"ct,omitempty"`
//	Wd   int32  `json:"wd,omitempty" bson:"wd,omitempty"`
//	Ht   int32  `json:"ht,omitempty" bson:"ht,omitempty"`
//	Lks  string `json:"lks,omitempty" bson:"lks,omitempty"`
//	Bln  string `json:"bln,omitempty" bson:"bln,omitempty"`
//	Cnt  string `json:"cnt,omitempty" bson:"cnt,omitempty"`
//	Url  string `json:"url,omitempty" bson:"url,omitempty"`
//	Role string `json:"role,omitempty" bson:"role,omitempty"`
//}
//
//func (s Variant) IsZero() bool {
//	/*
//	   if s.Ct == "" {
//	       return false
//	   }
//	   if s.Wd == 0 {
//	       return false
//	   }
//	   if s.Ht == 0 {
//	       return false
//	   }
//	   if s.Lks == "" {
//	       return false
//	   }
//	   if s.Bln == "" {
//	       return false
//	   }
//	   if s.Cnt == "" {
//	       return false
//	   }
//	   if s.Url == "" {
//	       return false
//	   }
//	   if s.Role == "" {
//	       return false
//	   }
//	       return true
//	*/
//	return s.Ct == "" && s.Wd == 0 && s.Ht == 0 && s.Lks == "" && s.Bln == "" && s.Cnt == "" && s.Url == "" && s.Role == ""
//}
