package utils

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"path/filepath"
)

// @link: https://github.com/hyrious/esbuild-plugin-style
var StyleInlinePlugin = api.Plugin{
	Name: "style-inline",
	Setup: func(build api.PluginBuild) {
		// Intercept import paths called "env" so esbuild doesn't attempt
		// to map them to a file system location. Tag them with the "env-ns"
		// namespace to reserve them for this plugin.
		build.OnResolve(api.OnResolveOptions{Filter: `.css$`, Namespace: "file"},
			func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				return api.OnResolveResult{
					Path:       filepath.Join(args.ResolveDir, args.Path),
					Namespace:  "style-stub",
					PluginData: args,
				}, nil
			})

		build.OnResolve(api.OnResolveOptions{Filter: `.css$`, Namespace: "style-stub"},
			func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				return api.OnResolveResult{
					Path:      args.Importer,
					Namespace: "style-content",
				}, nil
			})

		build.OnResolve(api.OnResolveOptions{Filter: `^__style_helper__$`, Namespace: "style-stub"}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
			return api.OnResolveResult{
				Path:        args.Path,
				Namespace:   "style-helper",
				SideEffects: api.SideEffectsFalse,
			}, nil
		})

		// Load paths tagged with the "env-ns" namespace and behave as if
		// they point to a JSON file containing the environment variables.
		build.OnLoad(api.OnLoadOptions{Filter: `.*`, Namespace: "style-helper"},
			func(args api.OnLoadArgs) (api.OnLoadResult, error) {
				contents := `
						export function injectStyle(text) {
							if (typeof document !== 'undefined') {
							  var style = document.createElement('style')
							  var node = document.createTextNode(text)
							  style.appendChild(node)
							  document.head.appendChild(style)
							}
						  }
					`
				return api.OnLoadResult{
					Contents: &contents,
				}, nil
			})

		build.OnLoad(api.OnLoadOptions{Filter: `.*`, Namespace: "style-stub"},
			func(args api.OnLoadArgs) (api.OnLoadResult, error) {
				contents := fmt.Sprintf(`
						import { injectStyle } from "__style_helper__"
						import css from "%s"
						injectStyle(css)
					`, args.Path)
				return api.OnLoadResult{
					Contents: &contents,
				}, nil
			})

		build.OnLoad(api.OnLoadOptions{Filter: `.*`, Namespace: "style-content"},
			func(args api.OnLoadArgs) (api.OnLoadResult, error) {
				result := BuildCss([]string{args.Path})
				content := string(result.OutputFiles[0].Contents)
				return api.OnLoadResult{
					Errors:   result.Errors,
					Warnings: result.Warnings,
					Contents: &content,
					Loader:   api.LoaderText,
				}, nil
			})
	},
}
