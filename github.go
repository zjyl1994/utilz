package utilz

import (
	"fmt"
	"time"
)

type GithubReleaseAsset struct {
	Name        string    `json:"name"`
	DownloadURL string    `json:"browser_download_url"`
	ContentType string    `json:"content_type"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GithubReleaseItem struct {
	Name        string               `json:"name"`
	TagName     string               `json:"tag_name"`
	CreatedAt   time.Time            `json:"created_at"`
	PublishedAt time.Time            `json:"published_at"`
	Assets      []GithubReleaseAsset `json:"assets"`
}

func ListGithubRelease(owner, repo string) ([]GithubReleaseItem, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)
	var resp []GithubReleaseItem
	bResp, err := HttpGet(apiURL)
	if err != nil {
		return nil, err
	}
	err = FromJSON(bResp, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetGithubLastestRelease(owner, repo, filename string) (tag, url string, err error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
	var resp GithubReleaseItem
	bResp, err := HttpGet(apiURL)
	if err != nil {
		return "", "", err
	}
	err = FromJSON(bResp, &resp)
	if err != nil {
		return "", "", err
	}
	for _, asset := range resp.Assets {
		if asset.Name == filename {
			return resp.TagName, asset.DownloadURL, nil
		}
	}
	return "", "", fmt.Errorf("%s not found", filename)
}
