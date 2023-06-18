package cmd

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/gozeon/gmpa/utils"
	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var (
	ignoreFolder    = regexp.MustCompile(`^(.git|dist|.idea|.vscode|public)$`)
	ignoreFile      = ".gmpaignore"
	outputDir       = "dist"
	publicDir       = "public"
	indexJavascript = "main.js"
	indexCss        = "style.css"
	indexHtml       = "index.html"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		workspace, err := os.Getwd()
		cobra.CheckErr(err)
		log.Info("workspace: ", workspace)

		fileInfo, err := afs.ReadDir(workspace)
		cobra.CheckErr(err)
		var wg sync.WaitGroup
		for _, v := range fileInfo {
			wg.Add(1)
			log.Debug(v)
			go func(v fs.FileInfo) {
				defer wg.Done()
				if !v.IsDir() {
					return
				}
				if v.Name() == publicDir {
					src := filepath.Join(workspace, v.Name())
					dest := filepath.Join(workspace, outputDir, publicDir)
					err := cp.Copy(src, dest)
					cobra.CheckErr(err)
				}
				if !ignoreFolder.MatchString(v.Name()) {
					isIgnore, err := afs.Exists(filepath.Join(workspace, v.Name(), ignoreFile))
					cobra.CheckErr(err)
					if isIgnore {
						return
					}

					html := utils.HtmlHelper{}

					jsFilePath := filepath.Join(workspace, v.Name(), indexJavascript)
					jsFileExists, err := afs.Exists(jsFilePath)
					cobra.CheckErr(err)
					if jsFileExists {
						r := utils.BuildJS([]string{jsFilePath})
						if len(r.Errors) > 0 {
							cobra.CheckErr(r.Errors)
						}
						for _, out := range r.OutputFiles {
							html.SetJs(string(out.Contents))
						}
					}

					cssFilePath := filepath.Join(workspace, v.Name(), indexCss)
					cssFileExists, err := afs.Exists(cssFilePath)
					cobra.CheckErr(err)
					if cssFileExists {
						c, err := afs.ReadFile(cssFilePath)
						cobra.CheckErr(err)
						html.SetCss(string(c))
					}

					if !jsFileExists && !cssFileExists {
						return
					}

					tempHtml := filepath.Join(workspace, v.Name(), indexHtml)
					tpl, err := utils.GetTemplate(tempHtml)
					cobra.CheckErr(err)

					html.SetTemplate(tpl)
					htmlString, err := html.GetHtml()
					cobra.CheckErr(err)

					dest := filepath.Join(workspace, outputDir, v.Name(), indexHtml)
					afs.Remove(dest)
					err = afs.SafeWriteReader(dest, strings.NewReader(htmlString))
					cobra.CheckErr(err)
					log.Info("generator html: ", dest)
				}
			}(v)
		}
		wg.Wait()
		log.Info("done.")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringP("output", "o", "", "output folder")
}
