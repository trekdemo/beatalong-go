package musicURLParser

import (
	"testing"

	"github.com/trekdemo/beatalong-go//utils"
)

func TestParseAppleMusicWithPrimaryID(t *testing.T) {
	url := "https://itunes.apple.com/nl/artist/kellerkind/id290891328?l=en"
	expectedResourceInfo := utils.ResourceInfo{
		Provider:    "AppleMusic",
		ID:          "290891328",
		Kind:        "artist",
		CountryCode: "nl",
	}

	result, _ := ParseAppleMusic(url)

	if *result != expectedResourceInfo {
		t.Errorf("Unexpected result: %v+", result)
	}
}

func TestParseAppleMusicWithSecondaryID(t *testing.T) {
	url := "https://itunes.apple.com/nl/artist/kellerkind/id290891328?l=en&i=123"
	expectedResourceInfo := utils.ResourceInfo{
		Provider:    "AppleMusic",
		ID:          "123",
		Kind:        "artist",
		CountryCode: "nl",
	}

	result, _ := ParseAppleMusic(url)

	if *result != expectedResourceInfo {
		t.Errorf("Unexpected result: %v+", result)
	}
}

func TestParseAppleMusicWithNonAppleMucisUrl(t *testing.T) {
	if _, err := ParseAppleMusic("https://index.hu"); err == nil {
		t.Errorf("Unexpected result: %+v", err)
	}
}
