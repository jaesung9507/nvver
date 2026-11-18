package tv

import (
	"fmt"

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

// Deprecated: NAVER TV service was discontinued on September 30, 2026.
// This method always returns nvver.ErrServiceDiscontinued.
func (c *Client) GetVODInfo(clipNo int64) (*VODInfo, error) {
	return nil, nvver.ErrServiceDiscontinued
}

func (c *Client) GetVODURL(clipNo int64, videoID, inKey string) (map[string]nvver.VODInfo, error) {
	return nvver.GetVODURL(c, videoID, inKey, map[string]string{
		"Referer": fmt.Sprintf("https://tv.naver.com/v/%d", clipNo),
	})
}
