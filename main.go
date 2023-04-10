package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

const (
	layout = "2006-01-02T15:04:05Z"
)

func mains(args []string) error {
	if len(args) < 2 {
		return errors.New("Usage: type-rnote USER REPO")
	}
	releases, err := getReleases(args[0], args[1], func(s ...any) {
		fmt.Fprintln(os.Stderr, s...)
	})
	if err != nil {
		return err
	}
	for _, r := range releases {
		if r.Draft {
			continue
		}
		fmt.Println(r.Name)
		fmt.Println("=======")

		dt, err := time.Parse(layout, r.PublishedAt)
		if err != nil {
			fmt.Println(r.PublishedAt)
		} else {
			fmt.Println(dt.Local().Format("Jan 2, 2006"))
		}
		fmt.Println()
		fmt.Println(r.Body)
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
