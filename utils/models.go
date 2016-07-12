// Package utils provides ...
package utils

type ResourceData struct {
}

// ResourceInfo is extracted from a URL which is being shared
type ResourceInfo struct {
	Provider    string `json:"provider,omitempty"`
	ID          string `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}
