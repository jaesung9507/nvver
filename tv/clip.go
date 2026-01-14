package tv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/jaesung9507/nvver"
)

func (c *Client) GetClipVideoID(clipNo int64) (string, error) {
	url := fmt.Sprintf("https://tv.naver.com/h/%d", clipNo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to new request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body: %w", err)
	}

	m := regexp.MustCompile(`(?s)<script[^>]*id="__NEXT_DATA__"[^>]*>(.*?)</script>`).FindSubmatch(data)
	if len(m) < 2 {
		return "", fmt.Errorf("not found data: content-length=%d", len(data))
	}

	result := &struct {
		Props struct {
			PageProps struct {
				ClipInfo struct {
					VideoID string `json:"videoId"`
				} `json:"clipInfo"`
			} `json:"pageProps"`
		} `json:"props"`
	}{}
	if err = json.Unmarshal(m[1], result); err != nil {
		return "", fmt.Errorf("failed to unmarshal data: %w", err)
	}

	videoID := result.Props.PageProps.ClipInfo.VideoID
	if len(videoID) <= 0 {
		return "", fmt.Errorf("not found video id: %s", string(m[1]))
	}

	return videoID, nil
}

func (c *Client) GetClipMP4URL(videoID string) (map[string]string, error) {
	return nvver.GetClipMP4URL(c, videoID, map[string]string{
		"Origin":  "https://m.naver.com",
		"Accept":  "*/*",
		"Referer": "https://m.naver.com/",
	})
}
