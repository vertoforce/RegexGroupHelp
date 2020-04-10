package regexgrouphelp

import "regexp"

// FindRegexGroups Takes a regex and returns a map of groupNames and their matches
func FindRegexGroups(regex *regexp.Regexp, data string) map[string][]string {
	matches := regex.FindAllStringSubmatch(data, -1)

	groups := map[string][]string{}

	groupNames := regex.SubexpNames()

	// Add each group
	for _, group := range groupNames {
		groups[group] = []string{}
	}

	// For every match
	for _, match := range matches {
		// For every group in this match
		for i, group := range match {
			groups[groupNames[i]] = append(groups[groupNames[i]], group)
		}
	}

	return groups
}
