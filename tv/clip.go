package tv

import (
	"github.com/jaesung9507/nvver"
)

// Deprecated: NAVER TV service was discontinued on September 30, 2026.
// This method always returns nvver.ErrServiceDiscontinued.
func (c *Client) GetClipVideoID(clipNo int64) (string, error) {
	return "", nvver.ErrServiceDiscontinued
}

func (c *Client) GetClipMP4URL(videoID string) (map[string]string, error) {
	return nvver.GetClipMP4URL(c, videoID, map[string]string{
		"Origin":  "https://m.naver.com",
		"Accept":  "*/*",
		"Referer": "https://m.naver.com/",
	})
}
