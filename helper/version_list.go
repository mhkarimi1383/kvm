package helper

import (
	"context"
	"github.com/google/go-github/v48/github"
)

func GetVersions(page int) ([]string, error) {
	gh := github.NewClient(nil)
	tags, _, err := gh.Repositories.ListTags(context.TODO(), "kubernetes", "kubernetes", &github.ListOptions{
		Page:    page,
		PerPage: 100,
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
