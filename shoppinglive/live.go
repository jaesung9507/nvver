package shoppinglive

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/jaesung9507/nvver"
)

func (c *Client) GetLivePlayback(broadcastID int64) (*nvver.Playback, error) {
	url := fmt.Sprintf("https://view.shoppinglive.naver.com/lives/%d", broadcastID)
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

	m := regexp.MustCompile(`window\.__viewerConfig\.broadcast\s*=\s*'([^']*)'`).FindSubmatch(data)
	if len(m) < 2 {
		return nil, fmt.Errorf("not found broadcast: content-length=%d", len(data))
	}

	var raw string
	if err = json.Unmarshal(fmt.Appendf(nil, `"%s"`, m[1]), &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal raw: %w", err)
	}

	result := &struct {
		Playback string `json:"playback"`
	}{}
	if err = json.Unmarshal([]byte(raw), result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	playback := &nvver.Playback{}
	if err = json.Unmarshal([]byte(result.Playback), playback); err != nil {
		return nil, fmt.Errorf("failed to unmarshal playback: %w", err)
	}

	return playback, nil
}
