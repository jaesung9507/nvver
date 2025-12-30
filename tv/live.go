package tv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/jaesung9507/nvver"
)

func (c *Client) GetLivePlayback(liveNo int64) (*nvver.Playback, error) {
	url := fmt.Sprintf("https://tv.naver.com/l/%d", liveNo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	m := regexp.MustCompile(`(?s)<script[^>]*id="__NEXT_DATA__"[^>]*>(.*?)</script>`).FindSubmatch(data)
	if len(m) < 2 {
		return nil, fmt.Errorf("not found data: content-length=%d", len(data))
	}

	result := &struct {
		Props struct {
			PageProps struct {
				LiveInfo struct {
					PlaybackBody string `json:"playbackBody"`
				} `json:"liveInfo"`
			} `json:"pageProps"`
		} `json:"props"`
	}{}
	if err = json.Unmarshal(m[1], result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	playback := &nvver.Playback{}
	if err = json.Unmarshal([]byte(result.Props.PageProps.LiveInfo.PlaybackBody), playback); err != nil {
		return nil, fmt.Errorf("failed to unmarshal playback: %w", err)
	}

	return playback, nil
}
