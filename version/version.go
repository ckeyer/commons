package version

import (
	"fmt"
	"time"

	"github.com/Masterminds/semver"
)

const (
	timefmt = "2006-01-02T15:04:05Z-0700"
)

// -X github.com/ckeyer/commons/version.version=$(VERSION)
// -X github.com/ckeyer/commons/version.gitCommit=$(GIT_COMMIT)
// -X github.com/ckeyer/commons/version.buildAt=${date "+%Y-%m-%dT%H:%M:%SZ%z"}
var (
	version, gitCommit string
	buildAt            string

	v   semver.Version
	bAt time.Time
)

func init() {
	if version == "" {
		return
	}
	ver, err := semver.NewVersion(version)
	if err != nil {
		panic(fmt.Errorf("format version failed, %s", err))
	}
	v = *ver

	if gitCommit != "" {
		v, _ = ver.SetMetadata(gitCommit)
	}

	if buildAt != "" {
		t, err := time.Parse(timefmt, buildAt)
		if err != nil {
			panic(fmt.Errorf("format buildAt failed, %s", err))
		}
		bAt = t
	}
}

// GetVersion
func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch())
}

// GetGitCommit
func GetGitCommit() string {
	return gitCommit
}

// GetCompleteVersion
func GetCompleteVersion() string {
	return v.String()
}

func GetBuildAt() time.Time {
	return bAt
}

func GetBuildAtString() string {
	return bAt.Format(timefmt)
}
