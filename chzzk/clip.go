package chzzk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaesung9507/nvver"
)

type ClipDetail struct {
	ClipUID           string `json:"clipUID"`
	VideoID           string `json:"videoId"`
	ClipTitle         string `json:"clipTitle"`
	ThumbnailImageURL string `json:"thumbnailImageUrl"`
	CategoryType      string `json:"categoryType"`
	ClipCategory      string `json:"clipCategory"`
	Duration          int    `json:"duration"`
	Adult             bool   `json:"adult"`
	BlindType         any    `json:"blindType"`
	KrOnlyViewing     bool   `json:"krOnlyViewing"`
	VodStatus         string `json:"vodStatus"`
	RecID             string `json:"recId"`
	CreatedDate       string `json:"createdDate"`
	OptionalProperty  any    `json:"optionalProperty"`
}

func (c *Client) GetClipDetail(clipID string) (*ClipDetail, error) {
	url := fmt.Sprintf("https://api.chzzk.naver.com/service/v1/clips/%s/detail?optionalProperties=COMMENT", clipID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", fmt.Sprintf("https://chzzk.naver.com/clips/%s", clipID))

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	result := &struct {
		Code    int        `json:"code"`
		Message any        `json:"message"`
		Content ClipDetail `json:"content"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return &result.Content, nil
}

func (c *Client) GetClipMP4URL(clipID, videoID string) (map[string]string, error) {
	return nvver.GetClipMP4URL(c, videoID, map[string]string{
		"Accept":  "application/json, text/plain, */*",
		"Referer": fmt.Sprintf("https://chzzk.naver.com/clips/%s", clipID),
	})
}
