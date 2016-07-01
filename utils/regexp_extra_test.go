package utils

import (
	"reflect"
	"regexp"
	"testing"
)

func TestParse(t *testing.T) {
	var myExp = RegexpExtra{regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)}

	actualCaptures := myExp.FindStringSubmatchMap("1234.5678.9")
	expectedCaptures := map[string]string{
		"first":  "1234",
		"second": "9",
	}

	if !reflect.DeepEqual(actualCaptures, expectedCaptures) {
		t.Errorf("Expected captures: %+v\nActual captures: %+v", expectedCaptures, actualCaptures)
	}
}
