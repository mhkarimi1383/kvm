package helper

import (
	"context"
	"github.com/google/go-github/v48/github"
)

func GetVersions() ([]string, error) {
	gh := github.NewClient(nil)
	tags, _, err := gh.Repositories.ListTags(context.TODO(), "kubernetes", "kubernetes", &github.ListOptions{
		Page:    1,
		PerPage: 15,
	})
	if err != nil {
		return nil, err
	}
	var versions []string
	for _, tag := range tags {
		versions = append(versions, *tag.Name)
	}
	return versions, nil
}
