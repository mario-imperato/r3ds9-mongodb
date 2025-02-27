package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type UserRole struct {
	Domain string `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site   string `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Apps   string `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s UserRole) IsZero() bool {
	return s.Domain == "" && s.Site == "" && s.Apps == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
