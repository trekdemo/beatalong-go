package musicURLParser

import (
	"testing"

	"github.com/trekdemo/beatalong-go/utils"
)

func TestParseAppleMusicWithPrimaryID(t *testing.T) {
	var cases = []struct {
		url      string
		expected utils.ResourceInfo
	}{
		{
			"https://itunes.apple.com/nl/artist/kellerkind/id290891328?l=en",
			utils.ResourceInfo{"AppleMusic", "290891328", "artist", "nl"},
		},
		{
			"https://itunes.apple.com/nl/album/caracal-deluxe/id1002029534?l=en",
			utils.ResourceInfo{"AppleMusic", "1002029534", "album", "nl"},
		},
		{
			"http://itunes.apple.com/nl/album/caracal-deluxe/id1002029534?l=en",
			utils.ResourceInfo{"AppleMusic", "1002029534", "album", "nl"},
		},
		{
			"https://itunes.apple.com/nl/album/caracal-deluxe/id1002029534?i=1002030097&l=en",
			utils.ResourceInfo{"AppleMusic", "1002030097", "song", "nl"},
		},
		{
			"https://itunes.apple.com/nl/playlist/amsterdam-dance-event-amsterdam/idpl.633ff1b81d3046138a4b384c717762d9?l=en",
			utils.ResourceInfo{"AppleMusic", "pl.633ff1b81d3046138a4b384c717762d9", "playlist", "nl"},
		},
	}

	for _, testCase := range cases {
		result, _ := ParseAppleMusic(testCase.url)

		if *result != testCase.expected {
			t.Errorf(
				"Unexpected result:\n\texpected:  %+v\n\tactual: %+v",
				testCase.expected,
				result,
			)
		}
	}
}

func TestParseAppleMusicWithNonAppleMucisUrl(t *testing.T) {
	if _, err := ParseAppleMusic("https://index.hu"); err == nil {
		t.Errorf("Unexpected result: %+v", err)
	}
}
