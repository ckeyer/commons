package version

import (
	"testing"

	"github.com/Masterminds/semver"
	check "gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestV(c *check.C) {
	v, err := semver.NewVersion("v1.2.3-beta.1+build345")
	c.Assert(err, check.IsNil)
	c.Assert(v.Major(), check.Equals, int64(1))
	c.Assert(v.Minor(), check.Equals, int64(2))
	c.Assert(v.Patch(), check.Equals, int64(3))
	c.Assert(v.Metadata(), check.Equals, "build345")

	v2, err := v.SetMetadata("asdfwer")
	c.Assert(err, check.IsNil)
	c.Assert(v2.Metadata(), check.Equals, "asdfwer")
}
