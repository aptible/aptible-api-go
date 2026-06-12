package aptibleapi

import "runtime/debug"

func aptibleClientVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/aptible/aptible-api-go" {
				return "aptible/aptible-api-go/" + dep.Version
			}
		}
		if info.Main.Version != "" {
			return "aptible/aptible-api-go/" + info.Main.Version
		}
	}
	return "aptible/aptible-api-go/unknown"
}
