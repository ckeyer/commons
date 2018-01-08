package crypto

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	check "gopkg.in/check.v1"
)

type S struct{}

var _ = check.Suite(&S{})

// Test main
func Test(t *testing.T) {
	check.TestingT(t)
}

// TestBcrypt test
func (s *S) TestBcrypt(c *check.C) {
	defPass := []byte("abc")

	for i := 0; i < 10; i++ {
		pass, err := bcrypt.GenerateFromPassword(defPass, 11)
		c.Check(err, check.IsNil)
		c.Logf("%s", pass)

		err = bcrypt.CompareHashAndPassword(pass, defPass)
		c.Check(err, check.IsNil)
	}
}

// Banckmark ...
func BenchmarkBcrypt11(c *testing.B) {
	defPass := []byte("abc")
	pass, err := bcrypt.GenerateFromPassword(defPass, 11)
	for i := 0; i < c.N; i++ {
		err = bcrypt.CompareHashAndPassword(pass, defPass)
		if err != nil {
			c.Error(err)
			return
		}

	}
}

// Banckmark ...
func BenchmarkBcrypt13(c *testing.B) {
	defPass := []byte("abc")
	pass, err := bcrypt.GenerateFromPassword(defPass, 13)
	for i := 0; i < c.N; i++ {
		err = bcrypt.CompareHashAndPassword(pass, defPass)
		if err != nil {
			c.Error(err)
			return
		}

	}
}

// Banckmark ...
func BenchmarkBcrypt19(c *testing.B) {
	defPass := []byte("abc")
	pass, err := bcrypt.GenerateFromPassword(defPass, 9)
	for i := 0; i < c.N; i++ {
		err = bcrypt.CompareHashAndPassword(pass, defPass)
		if err != nil {
			c.Error(err)
			return
		}
	}
}
