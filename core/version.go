package core

import "fmt"

type version struct {
	Major, Minor, Patch int
	Label               string
}

var (
	// Version is the current version of tpl
	Version = version{0, 2, 1, "dev"}
	// build is the git commit hash of this build
	// it's' set during compilcation with ldflags -X
	build string
)

func (v version) String() string {
	if v.Label == "" {
		return fmt.Sprintf("%d.%d.%d (%s)", v.Major, v.Minor, v.Patch, build)
	}
	return fmt.Sprintf("%d.%d.%d-%s (%s)", v.Major, v.Minor, v.Patch, v.Label, build)
}
