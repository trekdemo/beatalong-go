package utils

import "regexp"

// RegexpExtra adds a few helper methods to Regular Expressions
// http://blog.kamilkisiel.net/blog/2012/07/05/using-the-go-regexp-package/
type RegexpExtra struct {
	*regexp.Regexp
}

// FindStringSubmatchMap returns with a map of the matches
func (r *RegexpExtra) FindStringSubmatchMap(s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range r.SubexpNames() {
		//
		if i == 0 || name == "" {
			continue
		}
		captures[name] = match[i]

	}
	return captures
}
