package version

import "fmt"

type Info struct {
	Major, Minor, Patch int
}

func (i Info) String() string {
	return fmt.Sprintf("v%d.%d.%d", i.Major, i.Minor, i.Patch)
}

var Current = Info{Major: 1, Minor: 0, Patch: 0}
