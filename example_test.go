package regexgrouphelp

import (
	"fmt"
	"regexp"
)

func Example() {
	regex := regexp.MustCompile(`(?P<testGroup>test) data`)

	foundGroups := FindRegexGroups(regex, "this is test data")
	fmt.Printf("%v\n", foundGroups)

	// Output: map[:[test data] testGroup:[test]]
}
