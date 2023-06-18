package utils_test

import (
	"fmt"
	"testing"

	"github.com/gozeon/gmpa/utils"
)

func TestBuildJS(t *testing.T) {

	result := utils.BuildJS([]string{"fixtures/index.js"})

	if result.Errors != nil {
		t.Fatal(result.Errors)
	}

	for _, out := range result.OutputFiles {
		fmt.Println(string(out.Contents))
	}
}
