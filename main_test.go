package main

import (
	"reflect"
	"testing"
)

func TestRxGitUrl(t *testing.T) {
	source := `git@github.com:hymkor/type-rnote.git`
	expect := []string{
		`git@github.com:hymkor/type-rnote.git`,
		`hymkor`,
		`type-rnote`,
	}
	result := rxGitUrl.FindStringSubmatch(source)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("expect %#v,\nbut %#v", expect, result)
	}
}
