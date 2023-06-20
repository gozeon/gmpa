package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gozeon/gmpa/utils"
	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var (
	appName         = "gmpa"
	ignoreFolder    = regexp.MustCompile(`^(.git|dist|.idea|.vscode|public|node_modules)$`)
	ignoreFile      = ".gmpaignore"
	publicDir       = "public"
	indexJavascript = "main.js"
	indexCss        = "style.css"
	indexHtml       = "index.html"
	outputDir       string
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		var count int32 = 0
		workspace, err := os.Getwd()
		cobra.CheckErr(err)
		log.Info("workspace: ", workspace)

		fileInfo, err := afs.ReadDir(workspace)
		cobra.CheckErr(err)

		tempRoot := afs.GetTempDir(appName)
		destTempFolder, err := afs.TempDir(tempRoot, "build-")
		cobra.CheckErr(err)
		log.Info("Temp Folder: ", destTempFolder)

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
					dest := filepath.Join(destTempFolder, publicDir)
					cobra.CheckErr(cp.Copy(src, dest))
					log.Info("copy: ", src)
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

					dest := filepath.Join(destTempFolder, v.Name(), indexHtml)
					cobra.CheckErr(afs.SafeWriteReader(dest, strings.NewReader(htmlString)))
					log.Info("generator html: ", filepath.Join(v.Name(), indexHtml))
					atomic.AddInt32(&count, 1)
				}
			}(v)
		}
		wg.Wait()
		elapsed := time.Since(start)

		destFolder := utils.GetOutputPath(workspace, outputDir)
		exists, err := afs.Exists(destFolder)
		cobra.CheckErr(err)
		if exists {
			isDel, err := utils.PromptBool(fmt.Sprintf("`%s` is exists, delete it ? ", destFolder))
			cobra.CheckErr(err)
			if isDel {
				cobra.CheckErr(afs.RemoveAll(destFolder))
			}
		}
		// copy temp to dest
		cobra.CheckErr(cp.Copy(destTempFolder, destFolder))

		// del temp root
		cobra.CheckErr(afs.RemoveAll(tempRoot))

		log.Info("output: ", destFolder)
		log.WithFields(map[string]interface{}{
			"times": elapsed,
			"pages": count,
		}).Info("done.")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&outputDir, "output", "o", "dist", "output folder")
}
