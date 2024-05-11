package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func queryReleases(user, repo string, logger func(...any)) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", user, repo)
	logger("Get:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

type Release struct {
	Name        string `json:"name"`
	TagName     string `json:"tag_name"`
	Draft       bool   `json:"draft"`
	PublishedAt string `json:"published_at"`
	Body        string `json:body`
}

func getReleases(name, repo string, logger func(...any)) ([]*Release, error) {
	releasesStr, err := queryReleases(name, repo, logger)
	if err != nil {
		return nil, fmt.Errorf("getReleases: %w", err)
	}
	var releases []*Release
	if err := json.Unmarshal(releasesStr, &releases); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return releases, nil
}

const layout = "2006-01-02T15:04:05Z"

var rxGitHub = regexp.MustCompile(`(?:https://github.com/)?([^/]+)/([^/+]+)`)

func userAndRepo(args []string) (string, string, error) {
	if len(args) >= 1 {
		m := rxGitHub.FindStringSubmatch(args[0])
		if m != nil {
			return m[1], m[2], nil
		}
		if len(args) >= 2 {
			return args[0], args[1], nil
		}
	}
	return "", "", errors.New("Usage: type-rnote USER REPO")
}

func mains(args []string) error {
	user, repo, err := userAndRepo(args)
	if err != nil {
		return err
	}
	releases, err := getReleases(user, repo, func(s ...any) {
		fmt.Fprintln(os.Stderr, s...)
	})
	if err != nil {
		return err
	}
	for _, r := range releases {
		if r.Draft {
			continue
		}
		if r.Name == "" {
			fmt.Println(r.TagName)
		} else {
			fmt.Println(r.Name)
		}
		fmt.Println("=======")

		dt, err := time.Parse(layout, r.PublishedAt)
		if err != nil {
			fmt.Println(r.PublishedAt)
		} else {
			fmt.Println(dt.Local().Format("Jan 2, 2006"))
		}
		fmt.Println()
		fmt.Println(strings.ReplaceAll(r.Body, "\r", ""))
		fmt.Println()
	}
	return nil
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
