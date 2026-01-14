package tv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/jaesung9507/nvver"
)

type VODInfo struct {
	Clip struct {
		ClipNo      int64    `json:"clipNo"`
		Description string   `json:"description"`
		Title       string   `json:"title"`
		Hash        string   `json:"hash"`
		Tags        []string `json:"tags"`
		ChannelType string   `json:"channelType"`
		ChannelName string   `json:"channelName"`
		AuthType    string   `json:"authType"`
		ChannelURL  string   `json:"channelUrl"`
		AdultVideo  bool     `json:"adultVideo"`
		ChannelID   string   `json:"channelId"`
		VideoID     string   `json:"videoId"`
	} `json:"clip"`
	Play struct {
		InKey    string `json:"inKey"`
		Playable string `json:"playable"`
	} `json:"play"`
}

func (c *Client) GetVODInfo(clipNo int64) (*VODInfo, error) {
	url := fmt.Sprintf("https://tv.naver.com/v/%d", clipNo)
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
				VOD VODInfo `json:"vodInfo"`
			} `json:"pageProps"`
		} `json:"props"`
	}{}
	if err = json.Unmarshal(m[1], result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &result.Props.PageProps.VOD, nil
}

func (c *Client) GetVODURL(clipNo int64, videoID, inKey string) (map[string]string, error) {
	return nvver.GetVODURL(c, videoID, inKey, map[string]string{
		"Referer": fmt.Sprintf("https://tv.naver.com/v/%d", clipNo),
	})
}
