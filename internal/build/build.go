package build

import "os"

// Defines a Buildkite build
type BuildInfo struct {
	Number   string
	URL      string
	Pipeline string
	Author   string
	Branch   string
	Message  string
}

// GetBuildInfo retrieves the build info from the environment
func GetBuildInfo() BuildInfo {
	build := BuildInfo{
		Number:   os.Getenv("BUILDKITE_BUILD_NUMBER"),
		URL:      os.Getenv("BUILDKITE_BUILD_URL"),
		Pipeline: os.Getenv("BUILDKITE_PIPELINE_NAME"),
		Author:   os.Getenv("BUILDKITE_BUILD_CREATOR"),
		Branch:   os.Getenv("BUILDKITE_BRANCH"),
		Message:  os.Getenv("BUILDKITE_MESSAGE"),
	}

	return build
}