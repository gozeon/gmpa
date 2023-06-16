package utils

import "github.com/spf13/afero"

var Afs = afero.Afero{Fs: afero.NewOsFs()}
