package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"os"
)

//go:embed config.json
var configFile []byte

type Configuration struct {
	AccessToken string `json:"access_token"`
	RepoOwner   string `json:"repo_owner"`
	RepoName    string `json:"repo_name"`
}

func main() {
	var config Configuration

	err := json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	release, _, err := client.Repositories.GetLatestRelease(ctx, config.RepoOwner, config.RepoName)
	if err != nil {
		panic(err)
	}
	asset, _, err := client.Repositories.ListReleaseAssets(ctx, config.RepoOwner, config.RepoName, release.GetID(), &github.ListOptions{})
	if err != nil {
		panic(err)
	}
	_, url, err := client.Repositories.DownloadReleaseAsset(ctx, config.RepoOwner, config.RepoName, asset[0].GetID())
	err = DownloadFile(asset[0].GetName(), url)
	if err != nil {
		panic(err)
	}
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
