package cmd

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gozeon/gmpa/utils"
	cp "github.com/otiai10/copy"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	ignoreFolder    = regexp.MustCompile(`^(.git|dist|.idea|.vscode|asserts)$`)
	ignoreFile      = ".gmpaignore"
	outputDir       = "dist"
	assertsDir      = "asserts"
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

		afs := afero.Afero{afero.NewOsFs()}
		fileInfo, err := afs.ReadDir(workspace)
		cobra.CheckErr(err)
		for _, v := range fileInfo {
			log.Debug(v)
			if v.IsDir() {
				if v.Name() == assertsDir {
					src := filepath.Join(workspace, v.Name())
					dest := filepath.Join(workspace, outputDir)
					err := cp.Copy(src, dest)
					cobra.CheckErr(err)
				}
				if !ignoreFolder.MatchString(v.Name()) {
					isIgnore, err := afs.Exists(filepath.Join(workspace, v.Name(), ignoreFile))
					cobra.CheckErr(err)
					if isIgnore {
						continue
					}

					html := utils.HtmlHelper{}

					jsFilePath := filepath.Join(workspace, v.Name(), indexJavascript)
					jsFileExists, err := afs.Exists(jsFilePath)
					cobra.CheckErr(err)
					if jsFileExists {
						jsFile, err := afs.ReadFile(jsFilePath)
						cobra.CheckErr(err)
						html.SetJs(string(jsFile))
					}

					cssFilePath := filepath.Join(workspace, v.Name(), indexCss)
					cssFileExists, err := afs.Exists(cssFilePath)
					cobra.CheckErr(err)
					if cssFileExists {
						jsFile, err := afs.ReadFile(cssFilePath)
						cobra.CheckErr(err)
						html.SetCss(string(jsFile))
					}

					if !jsFileExists && !cssFileExists {
						continue
					}

					log.Debug(html)

					dest := filepath.Join(workspace, outputDir, v.Name(), indexHtml)
					indexHtmlExists, err := afs.Exists(dest)
					if indexHtmlExists {
						err := afs.Remove(dest)
						cobra.CheckErr(err)
					}
					err = afs.SafeWriteReader(dest, strings.NewReader(html.GetHtml()))
					cobra.CheckErr(err)
					log.Info("generator html ", dest)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringP("output", "o", "", "output folder")
}
