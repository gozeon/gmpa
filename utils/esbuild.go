package utils

import "github.com/evanw/esbuild/pkg/api"

func BuildJS(js string) api.TransformResult {
	return api.Transform(js, api.TransformOptions{
		Target: api.ES2015,
		Format: api.FormatIIFE,
	})
}
