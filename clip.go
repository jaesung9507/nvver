package nvver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetClipMP4URL(c Client, videoID string, header map[string]string) (map[string]string, error) {
	url := fmt.Sprintf("https://api-videohub.naver.com/shortformhub/feeds/v8/card?serviceType=%s&seedMediaId=%s&mediaType=VOD", c.GetServiceType(), videoID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	data := &struct {
		Card struct {
			Content struct {
				VOD struct {
					Playback struct {
						MPD []struct {
							Period []struct {
								AdaptationSet []struct {
									Representation []struct {
										BaseURL  []string `json:"BaseURL"`
										ID       string   `json:"@id"`
										MIMEType string   `json:"@mimeType"`
									} `json:"Representation"`
								} `json:"AdaptationSet"`
							} `json:"Period"`
						} `json:"MPD"`
					} `json:"playback"`
				} `json:"vod"`
			} `json:"content"`
		} `json:"card"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	result := make(map[string]string)
	for _, mpd := range data.Card.Content.VOD.Playback.MPD {
		for _, p := range mpd.Period {
			for _, as := range p.AdaptationSet {
				for _, rep := range as.Representation {
					if rep.MIMEType == "video/mp4" && len(rep.BaseURL) > 0 {
						result[rep.ID] = rep.BaseURL[0]
					}
				}
			}
		}
	}

	return result, nil
}
