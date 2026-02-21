package main

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var (
	rxSection = regexp.MustCompile(`^\s*\[([^\]]+)\]`)
)

func readIni(r io.Reader) (map[string]map[string]string, error) {
	secdata := map[string]string{}
	data := map[string]map[string]string{"": secdata}
	sc := bufio.NewScanner(r)
	var section string
	for sc.Scan() {
		line := sc.Text()
		if m := rxSection.FindStringSubmatch(line); m != nil {
			section = m[1]
			secdata = map[string]string{}
			data[section] = secdata
			// println("SECTION:",section)
		} else if key, val, ok := strings.Cut(line, "="); ok {
			key = strings.TrimSpace(key)
			val = strings.TrimSpace(val)
			secdata[key] = val
			// println("KEY:", key)
			// println("VALUE:", val)
		}
	}
	return data, sc.Err()
}
