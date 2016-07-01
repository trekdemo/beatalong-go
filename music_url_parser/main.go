// Package musicURLParser provides ...
package musicURLParser

import "github.com/trekdemo/beatalong-go/utils"

var urlParsers = []func(string) (*utils.ResourceInfo, error){
	ParseAppleMusic,
	ParseSporitfy,
}

// Parse ...
func Parse(url string) (*utils.ResourceInfo, error) {
	var resourceInfo *utils.ResourceInfo
	var err error

	for _, parser := range urlParsers {
		if resourceInfo, err = parser(url); err == nil {
			break
		}
	}
	return resourceInfo, nil
}
