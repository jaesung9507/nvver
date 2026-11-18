package tv

import (
	"github.com/jaesung9507/nvver"
)

// Deprecated: NAVER TV service was discontinued on September 30, 2026.
// This method always returns nvver.ErrServiceDiscontinued.
func (c *Client) GetLivePlayback(liveNo int64) (*nvver.Playback, error) {
	return nil, nvver.ErrServiceDiscontinued
}
