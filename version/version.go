package version

import "fmt"

var (
	Major        = 0
	Minor        = 1
	Micro        = 0
	ReleaseLevel = "dev"
)

// Version is the specification version that the package types support.
var Version = fmt.Sprintf("%d.%d.%d+%s",
	Major, Minor, Micro, ReleaseLevel)
