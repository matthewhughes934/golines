package internal

//go:generate sh -c "git tag --list --sort=-version:refname | head --lines 1 > ./version.txt"

import (
	_ "embed"
	"strings"
)

//go:embed version.txt
var version string

func Version() string {
	return strings.TrimSuffix(version, "\n")
}
