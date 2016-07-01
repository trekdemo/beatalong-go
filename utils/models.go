// Package utils provides ...
package utils

// ResourceInfo ...
type ResourceInfo struct {
	Provider    string `json:"provider,omitempty"`
	ID          string `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}
