package version

import "runtime"

var (
	gitVersion = "none"
	gitCommit  = "none"
	buildDate  = "none"
)

var ver = Version{
	GoVersion:  runtime.Version(),
	GoOs:       runtime.GOOS,
	GoArch:     runtime.GOARCH,
	GitVersion: gitVersion,
	GitCommit:  gitCommit,
	BuildDate:  buildDate,
}

type Version struct {
	GoVersion  string
	GoOs       string
	GoArch     string
	GitVersion string
	GitCommit  string
	BuildDate  string
}

func Get() Version {
	return ver
}
