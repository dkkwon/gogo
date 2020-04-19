package function

import "github.com/dkkwon/gogo/generate/stringset"

func InsertFunc(m stringset.StringSet) func(val string) {
	return func(val string) {
		m.Insert(val)
	}
	// Or it could be simply: return m.Insert
}
