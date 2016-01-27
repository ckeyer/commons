package format

import (
	"testing"
)

type stringEqualer func(string) string
type stringCsphereEqualer func(string, string) string

func TestCamel(t *testing.T) {
	stringEquale(t, FormatToCamel, "I", "I")
	stringEquale(t, FormatToCamel, "i", "I")
	stringEquale(t, FormatToCamel, "im", "IM")
	stringEquale(t, FormatToCamel, "iM", "I_M")
	stringEquale(t, FormatToCamel, "IM", "IM")
	stringEquale(t, FormatToCamel, "Im", "IM")
	stringEquale(t, FormatToCamel, "Name", "NAME")
	stringEquale(t, FormatToCamel, "Instance_name", "INSTANCE__NAME")
	stringEquale(t, FormatToCamel, "NAME", "NAME")
	stringEquale(t, FormatToCamel, "CPUNums", "CPU_NUMS")

	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "I", "CSPHERE_I")
	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "i", "CSPHERE_I")
	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "im", "CSPHERE_IM")
	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "iM", "CSPHERE_I_M")
	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "IM", "CSPHERE_IM")
	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "Im", "CSPHERE_IM")
	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "Name", "CSPHERE_NAME")
	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "NAME", "CSPHERE_NAME")
	stringEqualeCSPHERE(t, FormatToCamelWithPrefix, "CPUNums", "CSPHERE_CPU_NUMS")
}

func stringEquale(t *testing.T, f stringEqualer, str string, should string) {
	out := f(str)
	if should != out {
		t.Errorf(`%s be formated, should be %s ,but %s `, str, should, out)
	}
}

func stringEqualeCSPHERE(t *testing.T, f stringCsphereEqualer, str string, should string) {
	out := f("CSPHERE", str)
	if should != out {
		t.Errorf(`%s be formated, should be %s ,but %s `, str, should, out)
	}
}
