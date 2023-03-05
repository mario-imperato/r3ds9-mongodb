package commons

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	SYSINFO            = "sysinfo"
	SYSINFO_STATUS     = "sysinfo.status"
	SYSINFO_CREATEDAT  = "sysinfo.createdat"
	SYSINFO_MODIFIEDAT = "sysinfo.modifiedat"
	APP                = "app"
	APP_ID             = "app.id"
	APP_OBJTYPE        = "app.objType"
	APP_NAME           = "app.name"
	APP_DESCRIPTION    = "app.description"
	APP_PATH           = "app.path"
	APP_ROLEREQUIRED   = "app.roleRequired"
	ROLES              = "roles"
	ROLES_I            = "roles.%d"
	ROLES_I_DOMAIN     = "roles.%d.domain"
	ROLES_DOMAIN       = "roles.domain"
	ROLES_I_SITE       = "roles.%d.site"
	ROLES_SITE         = "roles.site"
	ROLES_I_APPS       = "roles.%d.apps"
	ROLES_APPS         = "roles.apps"
)

type Commons struct {
	Sysinfo SysInfo    `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
	App     App        `json:"app,omitempty" bson:"app,omitempty"`
	Roles   []UserRole `json:"roles,omitempty" bson:"roles,omitempty"`
}

func (s Commons) IsZero() bool {
	/*
	   if s.Sysinfo.IsZero() {
	       return false
	   }
	   if s.App.IsZero() {
	       return false
	   }
	   if len(s.Roles) == 0 {
	       return false
	   }
	       return true
	*/

	return s.Sysinfo.IsZero() && s.App.IsZero() && len(s.Roles) == 0
}

type SysInfo struct {
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	Createdat  primitive.DateTime `json:"createdat,omitempty" bson:"createdat,omitempty"`
	Modifiedat primitive.DateTime `json:"modifiedat,omitempty" bson:"modifiedat,omitempty"`
}

func (s SysInfo) IsZero() bool {
	/*
	   if s.Status == "" {
	       return false
	   }
	   if s.Createdat == 0 {
	       return false
	   }
	   if s.Modifiedat == 0 {
	       return false
	   }
	       return true
	*/
	return s.Status == "" && s.Createdat == 0 && s.Modifiedat == 0
}

type App struct {
	Id           string `json:"id,omitempty" bson:"id,omitempty"`
	ObjType      string `json:"objType,omitempty" bson:"objType,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Description  string `json:"description,omitempty" bson:"description,omitempty"`
	Path         string `json:"path,omitempty" bson:"path,omitempty"`
	RoleRequired bool   `json:"roleRequired,omitempty" bson:"roleRequired,omitempty"`
}

func (s App) IsZero() bool {
	/*
	   if s.Id == "" {
	       return false
	   }
	   if s.ObjType == "" {
	       return false
	   }
	   if s.Name == "" {
	       return false
	   }
	   if s.Description == "" {
	       return false
	   }
	   if s.Path == "" {
	       return false
	   }
	   if !s.RoleRequired {
	       return false
	   }
	       return true
	*/
	return s.Id == "" && s.ObjType == "" && s.Name == "" && s.Description == "" && s.Path == "" && !s.RoleRequired
}

type UserRole struct {
	Domain string `json:"domain,omitempty" bson:"domain,omitempty"`
	Site   string `json:"site,omitempty" bson:"site,omitempty"`
	Apps   string `json:"apps,omitempty" bson:"apps,omitempty"`
}

func (s UserRole) IsZero() bool {
	/*
	   if s.Domain == "" {
	       return false
	   }
	   if s.Site == "" {
	       return false
	   }
	   if s.Apps == "" {
	       return false
	   }
	       return true
	*/
	return s.Domain == "" && s.Site == "" && s.Apps == ""
}
