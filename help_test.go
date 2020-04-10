package regexgrouphelp

import (
	"reflect"
	"regexp"
	"testing"
)

func TestFindRegexGroups(t *testing.T) {
	tests := []struct {
		regex *regexp.Regexp
		data  string
		want  map[string][]string
	}{
		{
			regexp.MustCompile(`((?P<group1>test data)|(?P<group2>test stuff))`),
			"This is some test data and test stuff",
			map[string][]string{
				"group1": {"test data"},
				"group2": {"test stuff"},
			},
		},
		{
			regexp.MustCompile(`((?P<group1>test data \d)|(?P<group2>test stuff))`),
			"This is some test data 1 and test data 2 and test stuff",
			map[string][]string{
				"group1": {"test data 1", "test data 2"},
				"group2": {"test stuff"},
			},
		},
	}

	for _, test := range tests {
		if groups := FindRegexGroups(test.regex, test.data); !reflect.DeepEqual(groups, test.want) {
			t.Errorf("Test failed: %s", test.data)
		}
	}

}
