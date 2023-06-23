package utils

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

var Afs = afero.Afero{Fs: afero.NewOsFs()}

func GetWorkspacePath(target string) (string, error) {
	if filepath.IsAbs(target) {
		return target, nil
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return filepath.Join(cwd, target), nil
	}
}

func GetOutputPath(cwd string, output string) string {
	if filepath.IsAbs(output) {
		return output
	} else {
		return filepath.Join(cwd, output)
	}
}
