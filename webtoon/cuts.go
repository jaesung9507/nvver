package webtoon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jaesung9507/nvver"
)

type CutsInfo struct {
	ServiceTicketID string `json:"serviceTicketId"`
	PageID          string `json:"pageId"`
	PageGroupID     string `json:"pageGroupId"`
	PageName        string `json:"pageName"`
	PageURL         string `json:"pageUrl"`
	Status          string `json:"status"`
	Title           string `json:"title"`
	SectionGroup    struct {
		Sections []struct {
			SectionID   string `json:"sectionId"`
			SectionType string `json:"sectionType"`
			Data        struct {
				ContentID           string `json:"contentId"`
				ThumbnailURL        string `json:"thumbnailUrl"`
				PrePlaybackImageURL string `json:"prePlaybackImageUrl"`
			} `json:"data"`
		} `json:"sections"`
	} `json:"sectionGroup"`
	ID string `json:"id"`
}

func (c *CutsInfo) AssetID() (assetID string) {
	if c != nil {
		for _, s := range c.SectionGroup.Sections {
			if s.SectionType == "VOD" && len(s.Data.ContentID) > 0 {
				assetID = s.Data.ContentID
				break
			}
		}
	}

	return assetID
}

func (c *Client) GetCutsToken(contentID string) (string, error) {
	apiURL := fmt.Sprintf("https://comic.naver.com/cuts/v?id=%s", url.QueryEscape(contentID))
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to new request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	token := resp.Header.Get("x-csrf-token")
	if len(token) <= 0 {
		return "", fmt.Errorf("not found token: status code: %d", resp.StatusCode)
	}

	return token, nil
}

func (c *Client) GetCutsInfo(contentID string) (*CutsInfo, error) {
	apiURL := fmt.Sprintf("https://comic.naver.com/cuts/api/community/v2/post/%s?displayBlindCommentAsService=false", contentID)
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", fmt.Sprintf("https://comic.naver.com/cuts/v?id=%s", url.QueryEscape(contentID)))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("wc-service-ticket-id", "cuts")
	req.Header.Set("wc-service-type", "KW")

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	result := &struct {
		Status string `json:"status"`
		Result struct {
			Post CutsInfo `json:"post"`
		} `json:"result"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return &result.Result.Post, nil
}

func (c *Client) GetCutsURL(contentID, assetID, token string) (map[string]nvver.VODInfo, error) {
	if c.client.Jar == nil {
		return nil, errors.New("cookie jar is required")
	}

	data, err := json.Marshal(map[string]string{"assetId": assetID})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	apiURL := "https://comic.naver.com/cuts/api/community/v1/video/play-kit"
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", fmt.Sprintf("https://comic.naver.com/cuts/v?id=%s", url.QueryEscape(contentID)))
	req.Header.Set("X-CSRF-Token", token)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("wc-service-ticket-id", "cuts")
	req.Header.Set("wc-service-type", "KW")

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	result := &struct {
		Status string `json:"status"`
		Result struct {
			VideoPlayKit struct {
				AssetID        string `json:"assetId"`
				VideoID        string `json:"videoId"`
				InKey          string `json:"inKey"`
				InKeyExpiresAt int64  `json:"inKeyExpiresAt"`
			} `json:"videoPlayKit"`
		} `json:"result"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return nvver.GetVODURL(c, result.Result.VideoPlayKit.VideoID, result.Result.VideoPlayKit.InKey, map[string]string{
		"Referer": fmt.Sprintf("https://comic.naver.com/cuts/v?id=%s", url.QueryEscape(contentID)),
	})
}
