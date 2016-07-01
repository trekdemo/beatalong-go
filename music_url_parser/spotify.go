package musicURLParser

import (
	"errors"
	"regexp"

	"github.com/trekdemo/beatalong-go/utils"
)

var spotifyPrimaryAtributes = utils.RegexpExtra{
	regexp.MustCompile(`x`),
}

// ParseSporitfy ...
func ParseSporitfy(url string) (*utils.ResourceInfo, error) {
	return nil, errors.New("Not a Spotify URL")
}
