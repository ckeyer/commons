package version

import (
	"testing"

	"github.com/Masterminds/semver"
)

func TestV(t *testing.T) {
	v, err := semver.NewVersion("v1.2.3-beta.1+build345")
	if err != nil {
		t.Error(err)
		return
	}
	v2, _ := v.SetMetadata("asdfwer")

	t.Logf("%#v", v)
	t.Logf("%#v", v2)
}
