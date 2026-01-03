package chzzk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jaesung9507/nvver"
)

type VideoResp struct {
	Code    int `json:"code"`
	Message any `json:"message"`
	Content struct {
		VideoNo                int64    `json:"videoNo"`
		VideoId                string   `json:"videoId"`
		VideoTitle             string   `json:"videoTitle"`
		VideoType              string   `json:"videoType"`
		PublishDate            string   `json:"publishDate"`
		ThumbnailImageURL      string   `json:"thumbnailImageUrl"`
		TrailerURL             *string  `json:"trailerUrl"`
		Duration               int64    `json:"duration"`
		ReadCount              int64    `json:"readCount"`
		PublishDateAt          int64    `json:"publishDateAt"`
		CategoryType           string   `json:"categoryType"`
		VideoCategory          string   `json:"videoCategory"`
		VideoCategoryValue     string   `json:"videoCategoryValue"`
		Exposure               bool     `json:"exposure"`
		Adult                  bool     `json:"adult"`
		ClipActive             bool     `json:"clipActive"`
		LivePv                 int64    `json:"livePv"`
		Tags                   []string `json:"tags"`
		Channel                any      `json:"channel"`
		BlindType              any      `json:"blindType"`
		WatchTimeline          any      `json:"watchTimeline"`
		PaidProductId          any      `json:"paidProductId"`
		TvAppViewingPolicyType string   `json:"tvAppViewingPolicyType"`
		PaidPromotion          bool     `json:"paidPromotion"`
		InKey                  string   `json:"inKey"`
		LiveOpenDate           string   `json:"liveOpenDate"`
		VodStatus              string   `json:"vodStatus"`
		LiveRewindPlaybackJson *string  `json:"liveRewindPlaybackJson"`
		PrevVideo              any      `json:"prevVideo"`
		NextVideo              any      `json:"nextVideo"`
		UserAdultStatus        any      `json:"userAdultStatus"`
		AdParameter            any      `json:"adParameter"`
		VideoChatEnabled       bool     `json:"videoChatEnabled"`
		VideoChatChannelId     string   `json:"videoChatChannelId"`
		PaidProduct            any      `json:"paidProduct"`
		Dab                    bool     `json:"dab"`
	} `json:"content"`
}

func (r *VideoResp) GetLiveRewindPlayback() (*nvver.Playback, error) {
	if r.Content.LiveRewindPlaybackJson == nil {
		return nil, errors.New("live rewind playback is nil")
	}

	result := &nvver.Playback{}
	if err := json.Unmarshal([]byte(*r.Content.LiveRewindPlaybackJson), result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetVideo(videoNo int64) (*VideoResp, error) {
	url := fmt.Sprintf("https://api.chzzk.naver.com/service/v3/videos/%d", videoNo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", fmt.Sprintf("https://chzzk.naver.com/video/%d", videoNo))

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	result := &VideoResp{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return result, nil
}

func (c *Client) GetVideoURL(videoNo int64, videoID, inKey string) (map[string]string, error) {
	return nvver.GetVODURL(c, videoID, inKey, map[string]string{
		"Referer": fmt.Sprintf("https://chzzk.naver.com/video/%d", videoNo),
	})
}
