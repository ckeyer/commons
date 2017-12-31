package validate

import (
	"fmt"
	"math"
	"net"
	"regexp"
	"strings"
)

const dns1123LabelFmt string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
const dns1123LabelErrMsg string = "a DNS-1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character"
const DNS1123LabelMaxLength int = 63

var dns1123LabelRegexp = regexp.MustCompile("^" + dns1123LabelFmt + "$")

// IsDNS1123Label tests for a string that conforms to the definition of a label in
// DNS (RFC 1123).
func IsDNS1123Label(value string) []string {
	var errs []string
	if len(value) > DNS1123LabelMaxLength {
		errs = append(errs, MaxLenError(DNS1123LabelMaxLength))
	}
	if !dns1123LabelRegexp.MatchString(value) {
		errs = append(errs, RegexError(dns1123LabelErrMsg, dns1123LabelFmt, "my-name", "123-abc"))
	}
	return errs
}

const dns1123SubdomainFmt string = dns1123LabelFmt + "(\\." + dns1123LabelFmt + ")*"
const dns1123SubdomainErrorMsg string = "a DNS-1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character"
const DNS1123SubdomainMaxLength int = 253

var dns1123SubdomainRegexp = regexp.MustCompile("^" + dns1123SubdomainFmt + "$")

// IsDNS1123Subdomain tests for a string that conforms to the definition of a
// subdomain in DNS (RFC 1123).
func IsDNS1123Subdomain(value string) []string {
	var errs []string
	if len(value) > DNS1123SubdomainMaxLength {
		errs = append(errs, MaxLenError(DNS1123SubdomainMaxLength))
	}
	if !dns1123SubdomainRegexp.MatchString(value) {
		errs = append(errs, RegexError(dns1123SubdomainErrorMsg, dns1123SubdomainFmt, "example.com"))
	}
	return errs
}

const dns1035LabelFmt string = "[a-z]([-a-z0-9]*[a-z0-9])?"
const dns1035LabelErrMsg string = "a DNS-1035 label must consist of lower case alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character"
const DNS1035LabelMaxLength int = 63

var dns1035LabelRegexp = regexp.MustCompile("^" + dns1035LabelFmt + "$")

// IsDNS1035Label tests for a string that conforms to the definition of a label in
// DNS (RFC 1035).
func IsDNS1035Label(value string) []string {
	var errs []string
	if len(value) > DNS1035LabelMaxLength {
		errs = append(errs, MaxLenError(DNS1035LabelMaxLength))
	}
	if !dns1035LabelRegexp.MatchString(value) {
		errs = append(errs, RegexError(dns1035LabelErrMsg, dns1035LabelFmt, "my-name", "abc-123"))
	}
	return errs
}

// wildcard definition - RFC 1034 section 4.3.3.
// examples:
// - valid: *.bar.com, *.foo.bar.com
// - invalid: *.*.bar.com, *.foo.*.com, *bar.com, f*.bar.com, *
const wildcardDNS1123SubdomainFmt = "\\*\\." + dns1123SubdomainFmt
const wildcardDNS1123SubdomainErrMsg = "a wildcard DNS-1123 subdomain must start with '*.', followed by a valid DNS subdomain, which must consist of lower case alphanumeric characters, '-' or '.' and end with an alphanumeric character"

// IsWildcardDNS1123Subdomain tests for a string that conforms to the definition of a
// wildcard subdomain in DNS (RFC 1034 section 4.3.3).
func IsWildcardDNS1123Subdomain(value string) []string {
	wildcardDNS1123SubdomainRegexp := regexp.MustCompile("^" + wildcardDNS1123SubdomainFmt + "$")

	var errs []string
	if len(value) > DNS1123SubdomainMaxLength {
		errs = append(errs, MaxLenError(DNS1123SubdomainMaxLength))
	}
	if !wildcardDNS1123SubdomainRegexp.MatchString(value) {
		errs = append(errs, RegexError(wildcardDNS1123SubdomainErrMsg, wildcardDNS1123SubdomainFmt, "*.example.com"))
	}
	return errs
}

const cIdentifierFmt string = "[A-Za-z_][A-Za-z0-9_]*"
const identifierErrMsg string = "a valid C identifier must start with alphabetic character or '_', followed by a string of alphanumeric characters or '_'"

var cIdentifierRegexp = regexp.MustCompile("^" + cIdentifierFmt + "$")

// IsCIdentifier tests for a string that conforms the definition of an identifier
// in C. This checks the format, but not the length.
func IsCIdentifier(value string) []string {
	if !cIdentifierRegexp.MatchString(value) {
		return []string{RegexError(identifierErrMsg, cIdentifierFmt, "my_name", "MY_NAME", "MyName")}
	}
	return nil
}

// IsValidPortNum tests that the argument is a valid, non-zero port number.
func IsValidPortNum(port int) error {
	if 1 <= port && port <= 65535 {
		return nil
	}
	return InclusiveRangeError(1, 65535)
}

// IsInRange tests that the argument is in an inclusive range.
func IsInRange(value int, min int, max int) error {
	if value >= min && value <= max {
		return nil
	}
	return InclusiveRangeError(min, max)
}

// Now in libcontainer UID/GID limits is 0 ~ 1<<31 - 1
// TODO: once we have a type for UID/GID we should make these that type.
const (
	minUserID  = 0
	maxUserID  = math.MaxInt32
	minGroupID = 0
	maxGroupID = math.MaxInt32
)

// IsValidGID tests that the argument is a valid Unix GID.
func IsValidGID(gid int64) error {
	if minGroupID <= gid && gid <= maxGroupID {
		return nil
	}
	return InclusiveRangeError(minGroupID, maxGroupID)
}

// IsValidUID tests that the argument is a valid Unix UID.
func IsValidUID(uid int64) error {
	if minUserID <= uid && uid <= maxUserID {
		return nil
	}
	return InclusiveRangeError(minUserID, maxUserID)
}

var portNameCharsetRegex = regexp.MustCompile("^[-a-z0-9]+$")
var portNameOneLetterRegexp = regexp.MustCompile("[a-z]")

// IsValidPortName check that the argument is valid syntax. It must be
// non-empty and no more than 15 characters long. It may contain only [-a-z0-9]
// and must contain at least one letter [a-z]. It must not start or end with a
// hyphen, nor contain adjacent hyphens.
//
// Note: We only allow lower-case characters, even though RFC 6335 is case
// insensitive.
func IsValidPortName(port string) []string {
	var errs []string
	if len(port) > 15 {
		errs = append(errs, MaxLenError(15))
	}
	if !portNameCharsetRegex.MatchString(port) {
		errs = append(errs, "must contain only alpha-numeric characters (a-z, 0-9), and hyphens (-)")
	}
	if !portNameOneLetterRegexp.MatchString(port) {
		errs = append(errs, "must contain at least one letter or number (a-z, 0-9)")
	}
	if strings.Contains(port, "--") {
		errs = append(errs, "must not contain consecutive hyphens")
	}
	if len(port) > 0 && (port[0] == '-' || port[len(port)-1] == '-') {
		errs = append(errs, "must not begin or end with a hyphen")
	}
	return errs
}

// IsValidIP tests that the argument is a valid IP address.
func IsValidIP(value string) []string {
	if net.ParseIP(value) == nil {
		return []string{"must be a valid IP address, (e.g. 10.9.8.7)"}
	}
	return nil
}

const percentFmt string = "[0-9]+%"
const percentErrMsg string = "a valid percent string must be a numeric string followed by an ending '%'"

var percentRegexp = regexp.MustCompile("^" + percentFmt + "$")

func IsValidPercent(percent string) []string {
	if !percentRegexp.MatchString(percent) {
		return []string{RegexError(percentErrMsg, percentFmt, "1%", "93%")}
	}
	return nil
}

const httpHeaderNameFmt string = "[-A-Za-z0-9]+"
const httpHeaderNameErrMsg string = "a valid HTTP header must consist of alphanumeric characters or '-'"

var httpHeaderNameRegexp = regexp.MustCompile("^" + httpHeaderNameFmt + "$")

// IsHTTPHeaderName checks that a string conforms to the Go HTTP library's
// definition of a valid header field name (a stricter subset than RFC7230).
func IsHTTPHeaderName(value string) []string {
	if !httpHeaderNameRegexp.MatchString(value) {
		return []string{RegexError(httpHeaderNameErrMsg, httpHeaderNameFmt, "X-Header-Name")}
	}
	return nil
}

const envVarNameFmt = "[-._a-zA-Z][-._a-zA-Z0-9]*"
const envVarNameFmtErrMsg string = "a valid environment variable name must consist of alphabetic characters, digits, '_', '-', or '.', and must not start with a digit"

var envVarNameRegexp = regexp.MustCompile("^" + envVarNameFmt + "$")

// IsEnvVarName tests if a string is a valid environment variable name.
func IsEnvVarName(value string) []string {
	var errs []string
	if !envVarNameRegexp.MatchString(value) {
		errs = append(errs, RegexError(envVarNameFmtErrMsg, envVarNameFmt, "my.env-name", "MY_ENV.NAME", "MyEnvName1"))
	}

	errs = append(errs, hasChDirPrefix(value)...)
	return errs
}

const mapKeyFmt = `[-._a-zA-Z0-9]+`
const mapKeyErrMsg string = "a valid config key must consist of alphanumeric characters, '-', '_' or '.'"

var mapKeyRegexp = regexp.MustCompile("^" + mapKeyFmt + "$")

// IsMapKeyName tests for a string that is a valid key for a Map or Secret
func IsMapKeyName(value string) []string {
	var errs []string
	if len(value) > DNS1123SubdomainMaxLength {
		errs = append(errs, MaxLenError(DNS1123SubdomainMaxLength))
	}
	if !mapKeyRegexp.MatchString(value) {
		errs = append(errs, RegexError(mapKeyErrMsg, mapKeyFmt, "key.name", "KEY_NAME", "key-name"))
	}
	errs = append(errs, hasChDirPrefix(value)...)
	return errs
}

// MaxLenError returns a string explanation of a "string too long" validation
// failure.
func MaxLenError(length int) string {
	return fmt.Sprintf("must be no more than %d characters", length)
}

// RegexError returns a string explanation of a regex validation failure.
func RegexError(msg string, fmt string, examples ...string) string {
	if len(examples) == 0 {
		return msg + " (regex used for validation is '" + fmt + "')"
	}
	msg += " (e.g. "
	for i := range examples {
		if i > 0 {
			msg += " or "
		}
		msg += "'" + examples[i] + "', "
	}
	msg += "regex used for validation is '" + fmt + "')"
	return msg
}

func prefixEach(msgs []string, prefix string) []string {
	for i := range msgs {
		msgs[i] = prefix + msgs[i]
	}
	return msgs
}

// InclusiveRangeError returns a string explanation of a numeric "must be
// between" validation failure.
func InclusiveRangeError(lo, hi int) error {
	return fmt.Errorf(`must be between %d and %d, inclusive`, lo, hi)
}

func hasChDirPrefix(value string) []string {
	var errs []string
	switch {
	case value == ".":
		errs = append(errs, `must not be '.'`)
	case value == "..":
		errs = append(errs, `must not be '..'`)
	case strings.HasPrefix(value, ".."):
		errs = append(errs, `must not start with '..'`)
	}
	return errs
}
