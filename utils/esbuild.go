package utils

import "github.com/evanw/esbuild/pkg/api"

func BuildJS(jsFile []string) api.BuildResult {
	return api.Build(api.BuildOptions{
		EntryPoints:       jsFile,
		Bundle:            true,
		MinifyWhitespace:  false,
		MinifyIdentifiers: false,
		MinifySyntax:      false,
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "58"},
			{Name: api.EngineFirefox, Version: "57"},
			{Name: api.EngineSafari, Version: "11"},
			{Name: api.EngineEdge, Version: "16"},
		},
		Write:   false,
		Outfile: "index.js",
	})
}
