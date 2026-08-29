package version

import "fmt"

// Info holds semver-style build metadata.
type Info struct {
	Major int
	Minor int
	Patch int
	Build string
}

func (i Info) String() string {
	if i.Build == "" {
		return fmt.Sprintf("v%d.%d.%d", i.Major, i.Minor, i.Patch)
	}
	return fmt.Sprintf("v%d.%d.%d+%s", i.Major, i.Minor, i.Patch, i.Build)
}

var Current = Info{Major: 0, Minor: 100, Patch: 0, Build: "day92"}
