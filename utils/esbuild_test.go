package utils_test

import (
	"testing"

	"github.com/gozeon/gmpa/utils"
)

func TestBuildJS(t *testing.T) {
	source := `
	class A {
		constructor(a) {
			this.a = a
			this.a()
		}
	
		async a() {
			console.log(this?.b)
		}
	}
	const a = new A()
	`
	result := utils.BuildJS(source)
	if result.Errors != nil {
		t.Fatal(result.Errors)
	}
}
