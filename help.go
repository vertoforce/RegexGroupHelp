package regexgrouphelp

import "regexp"

// FindRegexGroups Takes a regex and returns a map of groupNames and their matches
func FindRegexGroups(regex *regexp.Regexp, data string) map[string][]string {
	matches := regex.FindAllStringSubmatch(data, -1)

	groups := map[string][]string{}
	groupNames := regex.SubexpNames()

	// Add each group
	for _, groupName := range groupNames {
		if groupName == "" {
			continue
		}
		groups[groupName] = []string{}
	}

	// For every match
	for _, match := range matches {
		// For every group in this match
		for i, groupMatch := range match {
			// Skip if match is empty
			if groupMatch == "" {
				continue
			}
			// Skip if groupName is empty
			if groupNames[i] == "" {
				continue
			}

			groups[groupNames[i]] = append(groups[groupNames[i]], groupMatch)
		}
	}

	return groups
}
