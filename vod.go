package nvver

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func GetVODURL(c Client, videoID, inKey string, header map[string]string) (map[string]string, error) {
	url := fmt.Sprintf("https://apis.naver.com/neonplayer/vodplay/v3/playback/%s?key=%s", videoID, inKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Accept", "application/xml")

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	data := &struct {
		Periods []struct {
			AdaptationSets []struct {
				Representations []struct {
					ID      string `xml:"id,attr"`
					BaseURL string `xml:"BaseURL"`
					M3U     string `xml:"m3u,attr"`
				} `xml:"Representation"`
			} `xml:"AdaptationSet"`
		} `xml:"Period"`
	}{}
	if err = xml.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode xml: %w", err)
	}

	result := make(map[string]string)
	for _, p := range data.Periods {
		for _, as := range p.AdaptationSets {
			for _, rep := range as.Representations {
				if rep.M3U != "" {
					result[rep.ID] = rep.M3U
				} else if rep.BaseURL != "" {
					result[rep.ID] = rep.BaseURL
				}
			}
		}
	}

	return result, nil
}
