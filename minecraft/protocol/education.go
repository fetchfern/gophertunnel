package protocol

// EducationSharedResourceURI is an education edition feature that is used for transmitting
// education resource settings to clients. It contains a button name and a link URL.
type EducationSharedResourceURI struct {
	// ButtonName is the button name of the resource URI.
	ButtonName string
	// LinkURI is the link URI for the resource URI.
	LinkURI string
}

// Marshal reads/writes an EducationSharedResourceURI to an IO.
func (x *EducationSharedResourceURI) Marshal(r IO) {
	r.String(&x.ButtonName)
	r.String(&x.LinkURI)
}

// EducationExternalLinkSettings ...
type EducationExternalLinkSettings struct {
	// URL is the external link URL.
	URL string
	// DisplayName is the display name in game.
	DisplayName string
}

// Marshal encodes/decodes an EducationExternalLinkSettings.
func (x *EducationExternalLinkSettings) Marshal(r IO) {
	r.String(&x.URL)
	r.String(&x.DisplayName)
}
