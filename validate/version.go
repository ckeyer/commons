package validate

import (
	"fmt"
	"log"
)

const (
	ValidatePage  = "/validate/"
	Supported     = "[Supported, but not recommended]"
	Unsupported   = "[Unsupported]"
	Recommended   = "[Supported and recommended]"
	Unknown       = "[Unknown]"
	VersionFormat = "%d.%d%s"
	OutputFormat  = "%s: %s\n\t%s\n\n"
)

func getMajorMinor(version string) (int, int, error) {
	var major, minor int
	var ign string
	n, err := fmt.Sscanf(version, VersionFormat, &major, &minor, &ign)
	if n != 3 || err != nil {
		log.Printf("Failed to parse version for %s", version)
		return -1, -1, err
	}
	return major, minor, nil
}

func ValidateKernelVersion(version string) (string, string) {
	desc := fmt.Sprintf("Kernel version is %s. Versions >= 2.6 are supported. 3.0+ are recommended.\n", version)
	major, minor, err := getMajorMinor(version)
	if err != nil {
		desc = fmt.Sprintf("Could not parse kernel version. %s", desc)
		return Unknown, desc
	}

	if major < 2 {
		return Unsupported, desc
	}

	if major == 2 && minor < 6 {
		return Unsupported, desc
	}

	if major >= 3 {
		return Recommended, desc
	}

	return Supported, desc
}

func ValidateDockerVersion(version string) (string, string) {
	desc := fmt.Sprintf("Docker version is %s. Versions >= 1.0 are supported. 1.2+ are recommended.\n", version)
	major, minor, err := getMajorMinor(version)
	if err != nil {
		desc = fmt.Sprintf("Could not parse docker version. %s\n\t", desc)
		return Unknown, desc
	}
	if major < 1 {
		return Unsupported, desc
	}

	if major == 1 && minor < 2 {
		return Supported, desc
	}

	return Recommended, desc
}
