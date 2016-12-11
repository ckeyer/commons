package mongo

import (
	"testing"
)

func TestURL2Database(t *testing.T) {
	tb := map[string]string{
		"mongodb://user:pass@host1:1234/database?options": "database",
		"mongodb://user:pass@host1:1234/a?options":        "a",
		"mongodb://user:pass@host1:1234/?options":         "",
		"mongodb://user:pass@host1:1234?options":          "",

		"mongodb://user:pass@host1:1234/database": "database",
		"mongodb://user:pass@host1:1234/a":        "a",
		"mongodb://user:pass@host1:1234/":         "",
		"mongodb://user:pass@host1:1234":          "",

		"user:pass@host1:1234/database": "database",
		"user:pass@host1:1234/a":        "a",
		"user:pass@host1:1234/":         "",
		"user:pass@host1:1234":          "",
	}

	for k, v := range tb {
		cur := getDbNameFromURL(k)
		if v != cur {
			t.Errorf("get from %s, should be: %s, but %s", k, v, cur)
		}
	}
}
