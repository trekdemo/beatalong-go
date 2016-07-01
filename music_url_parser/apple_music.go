package musicURLParser

import (
	"errors"
	"regexp"

	"github.com/trekdemo/beatalong-go/utils"
)

var appleMusicPrimaryAtributes = utils.RegexpExtra{
	regexp.MustCompile(`(?i)https?://itunes.apple.com/(?P<country_code>[a-z]{2})/(?P<kind>artist|album|playlist)(?P<description>/.+)?/id(?P<id>(pl\.)?\w+)`),
}

var appleMusicsecondaryAttributes = utils.RegexpExtra{regexp.MustCompile(`(?i)i=(?P<id>\w+)`)}

// ParseAppleMusic ...
func ParseAppleMusic(url string) (*utils.ResourceInfo, error) {
	matches := appleMusicPrimaryAtributes.FindStringSubmatchMap(url)
	secondaryMatches := appleMusicsecondaryAttributes.FindStringSubmatchMap(url)
	var id string

	if secondaryMatches["id"] != "" {
		id = secondaryMatches["id"]
	} else {
		id = matches["id"]
	}

	if id != "" {
		return &utils.ResourceInfo{
			Provider:    "AppleMusic",
			ID:          id,
			Kind:        matches["kind"],
			CountryCode: matches["country_code"],
		}, nil
	}
	return nil, errors.New("Not an AppleMusic URL")
}