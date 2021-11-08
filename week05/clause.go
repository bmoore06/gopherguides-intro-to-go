package main //You will always need this at the beginning of every go program.

import ( // You will also need this at the beginning of all go programs in order to import certain packages to run.
	"fmt"
	"sort"
	"strings"
)

type Clauses map[string]interface{}

func (cls Clauses) String() string {
	lines := make([]string, 0, len(cls))

	for k, v := range cls {
		lines = append(lines, fmt.Sprintf("%q = %q", k, v))
	}

	sort.Strings(lines)
	return strings.Join(lines, " and ")
}

func (cls Clauses) Match(m Model) bool {
	for k, v := range cls {
		if m[k] != v {
			return false
		}
	}
	return true
}
