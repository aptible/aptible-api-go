package aptibleapi

import "runtime/debug"

func aptibleClientVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/aptible/aptible-api-go" {
				return dep.Version
			}
		}
		if info.Main.Path == "github.com/aptible/aptible-api-go" {
			return info.Main.Version
		}
	}
	return "unknown"
}
