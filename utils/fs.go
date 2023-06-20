package utils

import (
	"github.com/spf13/afero"
	"path/filepath"
)

var Afs = afero.Afero{Fs: afero.NewOsFs()}

func GetOutputPath(cwd string, output string) string {
	if filepath.IsAbs(output) {
		return output
	} else {
		return filepath.Join(cwd, output)
	}
}
