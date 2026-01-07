package shoppinglive

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/jaesung9507/nvver"
)

type ClipInfo struct {
	ShortClipID int64  `json:"shortclipId"`
	Status      string `json:"status"`
	VID         string `json:"vid"`
	VODMediaURL string `json:"vodMediaUrl"`
}

func (c *Client) GetShortClipInfo(shortClipID int64) (*ClipInfo, error) {
	url := fmt.Sprintf("https://view.shoppinglive.naver.com/shortclips/%d", shortClipID)
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

	m := regexp.MustCompile(`window\.__shortclip\s*=\s*'([^']*)'`).FindSubmatch(data)
	if len(m) < 2 {
		return nil, fmt.Errorf("not found broadcast: content-length=%d", len(data))
	}

	var raw string
	if err = json.Unmarshal(fmt.Appendf(nil, `"%s"`, m[1]), &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal raw: %w", err)
	}

	result := &ClipInfo{}
	if err = json.Unmarshal([]byte(raw), result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return result, nil
}

func (c *Client) GetShortClipURL(shortClipID int64, vodMediaURL string) (map[string]string, error) {
	u, err := url.Parse(vodMediaURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse vod media url: %w", err)
	}

	inKey, ok := u.User.Password()
	if !ok {
		return nil, fmt.Errorf("not found key: %s", vodMediaURL)
	}

	return nvver.GetVODURL(c, u.Host, inKey, map[string]string{
		"Referer": fmt.Sprintf("https://view.shoppinglive.naver.com/shortclips/%d", shortClipID),
	})
}
