package domain

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName                = "_id"
	CodeFieldName               = "code"
	ObjTypeFieldName            = "objType"
	NameFieldName               = "name"
	DescriptionFieldName        = "description"
	LangsFieldName              = "langs"
	MembersFieldName            = "members"
	Members_IFieldName          = "members.%d"
	AppsFieldName               = "apps"
	Apps_IFieldName             = "apps.%d"
	SysInfoFieldName            = "sysInfo"
	SysInfo_StatusFieldName     = "sysInfo.status"
	SysInfo_CreatedatFieldName  = "sysInfo.createdat"
	SysInfo_ModifiedatFieldName = "sysInfo.modifiedat"
)

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
