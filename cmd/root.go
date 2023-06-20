package cmd

import (
	"os"

	"github.com/gozeon/gmpa/utils"
	"github.com/spf13/cobra"
)

var log = utils.Log
var afs = utils.Afs
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gmpa",
	Short: "Build Multi-Page Apps Tools",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		if debug {
			utils.SetLogInfo()
		}
	})
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug log output")
}
