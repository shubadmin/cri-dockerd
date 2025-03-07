package version

var (
	// Version of the product
	Version = "0.2.1"
	// PreRelease is set during the build
	PreRelease = ""
	// GitCommit is set during the build
	GitCommit = "HEAD"
	// BuildTime is set during the build
	BuildTime = "<unknown>"
)

const (
	// PlatformName of the product
	PlatformName = "cri-dockerd"
)

// FullVersion returns the formatted "$Version[-$PreRelease] ($GitCommit)"
func FullVersion() string {
	if PreRelease != "" {
		return Version + "-" + PreRelease + " (" + GitCommit + ")"
	}
	return Version + " (" + GitCommit + ")"
}

// TagVersion returns "$Version[-$PreRelease]" without the git commit
func TagVersion() string {
	if PreRelease != "" {
		return Version + "-" + PreRelease
	}
	return Version
}
