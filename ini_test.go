package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadIni(t *testing.T) {
	result, err := readIni(strings.NewReader(`
		[remote "origin"]
			url = git@github.com:hymkor/type-rnote.git
			fetch = +refs/heads/*:refs/remotes/origin/*
	`))

	if err != nil {
		t.Fatal(err.Error())
	}

	expect := map[string]map[string]string{
		``: map[string]string{},
		`remote "origin"`: map[string]string{
			"url":   "git@github.com:hymkor/type-rnote.git",
			"fetch": "+refs/heads/*:refs/remotes/origin/*",
		},
	}

	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Expect #%v, but %#v", expect, result)
	}
}
