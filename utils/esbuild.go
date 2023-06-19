package utils

import "github.com/evanw/esbuild/pkg/api"

func BuildJS(jsFile []string) api.BuildResult {
	return api.Build(api.BuildOptions{
		EntryPoints:       jsFile,
		Bundle:            true,
		MinifyWhitespace:  false,
		MinifyIdentifiers: false,
		MinifySyntax:      false,
		JSXFactory:        "h",
		JSXFragment:       "Fragment",
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "58"},
			{Name: api.EngineFirefox, Version: "57"},
			{Name: api.EngineSafari, Version: "11"},
			// https://github.com/evanw/esbuild/issues/988#issuecomment-801315072
			{Name: api.EngineEdge, Version: "18"},
		},
		Loader: map[string]api.Loader{
			".js":   api.LoaderJSX,
			".png":  api.LoaderDataURL,
			".jpg":  api.LoaderDataURL,
			".jpeg": api.LoaderDataURL,
			".gif":  api.LoaderDataURL,
			".bmp":  api.LoaderDataURL,
			".svg":  api.LoaderDataURL,
		},
		Plugins: []api.Plugin{StyleInlinePlugin},
		Write:   false,
		Outfile: "index.js",
	})
}

func BuildCss(cssFile []string) api.BuildResult {
	return api.Build(api.BuildOptions{
		EntryPoints: cssFile,
		Bundle:      true,
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "58"},
			{Name: api.EngineFirefox, Version: "57"},
			{Name: api.EngineSafari, Version: "11"},
			// https://github.com/evanw/esbuild/issues/988#issuecomment-801315072
			{Name: api.EngineEdge, Version: "18"},
		},
		Loader: map[string]api.Loader{
			".js":   api.LoaderJSX,
			".png":  api.LoaderDataURL,
			".jpg":  api.LoaderDataURL,
			".jpeg": api.LoaderDataURL,
			".gif":  api.LoaderDataURL,
			".bmp":  api.LoaderDataURL,
			".svg":  api.LoaderDataURL,
		},
		Write:   false,
		Outfile: "style.css",
	})
}
