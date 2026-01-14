package chzzk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaesung9507/nvver"
)

type LiveStatus struct {
	LiveTitle                     string   `json:"liveTitle"`
	Status                        string   `json:"status"`
	ConcurrentUserCount           int      `json:"concurrentUserCount"`
	CvExposure                    bool     `json:"cvExposure"`
	AccumulateCount               int      `json:"accumulateCount"`
	PaidPromotion                 bool     `json:"paidPromotion"`
	Adult                         bool     `json:"adult"`
	KrOnlyViewing                 bool     `json:"krOnlyViewing"`
	OpenDate                      string   `json:"openDate"`
	CloseDate                     *string  `json:"closeDate"`
	ClipActive                    bool     `json:"clipActive"`
	ChatChannelID                 string   `json:"chatChannelId"`
	Tags                          []string `json:"tags"`
	CategoryType                  string   `json:"categoryType"`
	LiveCategory                  string   `json:"liveCategory"`
	LiveCategoryValue             string   `json:"liveCategoryValue"`
	LivePollingStatusJson         string   `json:"livePollingStatusJson"`
	FaultStatus                   any      `json:"faultStatus"`
	UserAdultStatus               *string  `json:"userAdultStatus"`
	AbroadCountry                 bool     `json:"abroadCountry"`
	BlindType                     any      `json:"blindType"`
	PlayerRecommendContent        any      `json:"playerRecommendContent"`
	ChatActive                    bool     `json:"chatActive"`
	ChatAvailableGroup            string   `json:"chatAvailableGroup"`
	ChatAvailableCondition        string   `json:"chatAvailableCondition"`
	MinFollowerMinute             int      `json:"minFollowerMinute"`
	AllowSubscriberInFollowerMode bool     `json:"allowSubscriberInFollowerMode"`
	ChatSlowModeSec               int      `json:"chatSlowModeSec"`
	ChatEmojiMode                 bool     `json:"chatEmojiMode"`
	ChatDonationRankingExposure   bool     `json:"chatDonationRankingExposure"`
	DropsCampaignNo               any      `json:"dropsCampaignNo"`
	LiveTokenList                 []string `json:"liveTokenList"`
	WatchPartyNo                  int      `json:"watchPartyNo"`
	WatchPartyTag                 string   `json:"watchPartyTag"`
	TimeMachineActive             bool     `json:"timeMachineActive"`
	ChannelID                     string   `json:"channelId"`
	LastAdultReleaseDate          any      `json:"lastAdultReleaseDate"`
	LastKrOnlyViewingReleaseDate  any      `json:"lastKrOnlyViewingReleaseDate"`
	LastTvAppAllowedDate          any      `json:"lastTvAppAllowedDate"`
	TvAppViewingPolicyType        string   `json:"tvAppViewingPolicyType"`
	LogPowerActive                bool     `json:"logPowerActive"`
	LogPowerRankingExposure       bool     `json:"logPowerRankingExposure"`
}

func (c *Client) GetLiveStatus(channelID string) (*LiveStatus, error) {
	url := fmt.Sprintf("https://api.chzzk.naver.com/polling/v3.1/channels/%s/live-status", channelID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", fmt.Sprintf("https://chzzk.naver.com/live/%s", channelID))

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
		Content LiveStatus `json:"content"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return &result.Content, nil
}

type LiveDetail struct {
	LiveID                        int64    `json:"liveId"`
	LiveTitle                     string   `json:"liveTitle"`
	Status                        string   `json:"status"`
	LiveImageURL                  string   `json:"liveImageUrl"`
	DefaultThumbnailImageURL      any      `json:"defaultThumbnailImageUrl"`
	ConcurrentUserCount           int      `json:"concurrentUserCount"`
	CvExposure                    bool     `json:"cvExposure"`
	AccumulateCount               int      `json:"accumulateCount"`
	OpenDate                      string   `json:"openDate"`
	CloseDate                     *string  `json:"closeDate"`
	Adult                         bool     `json:"adult"`
	KrOnlyViewing                 bool     `json:"krOnlyViewing"`
	ClipActive                    bool     `json:"clipActive"`
	Tags                          []string `json:"tags"`
	ChatChannelID                 string   `json:"chatChannelId"`
	CategoryType                  string   `json:"categoryType"`
	LiveCategory                  string   `json:"liveCategory"`
	LiveCategoryValue             string   `json:"liveCategoryValue"`
	ChatActive                    bool     `json:"chatActive"`
	ChatAvailableGroup            string   `json:"chatAvailableGroup"`
	PaidPromotion                 bool     `json:"paidPromotion"`
	ChatAvailableCondition        string   `json:"chatAvailableCondition"`
	MinFollowerMinute             int      `json:"minFollowerMinute"`
	AllowSubscriberInFollowerMode bool     `json:"allowSubscriberInFollowerMode"`
	LivePlaybackJson              string   `json:"livePlaybackJson"`
	P2PQuality                    []string `json:"p2pQuality"`
	TimeMachineActive             bool     `json:"timeMachineActive"`
	TimeMachinePlayback           bool     `json:"timeMachinePlayback"`
	Channel                       any      `json:"channel"`
	LivePollingStatusJson         string   `json:"livePollingStatusJson"`
	UserAdultStatus               *string  `json:"userAdultStatus"`
	BlindType                     any      `json:"blindType"`
	ChatDonationRankingExposure   bool     `json:"chatDonationRankingExposure"`
	LogPower                      any      `json:"logPower"`
	AdParameter                   any      `json:"adParameter"`
	DropsCampaignNo               any      `json:"dropsCampaignNo"`
	WatchPartyNo                  int      `json:"watchPartyNo"`
	WatchPartyTag                 string   `json:"watchPartyTag"`
	WatchPartyPaidProductID       any      `json:"watchPartyPaidProductId"`
	Dab                           bool     `json:"dab"`
	Earthquake                    bool     `json:"earthquake"`
	PaidProduct                   any      `json:"paidProduct"`
	TvAppViewingPolicyType        string   `json:"tvAppViewingPolicyType"`
	Party                         any      `json:"party"`
}

func (r *LiveDetail) GetLivePlayback() (*nvver.Playback, error) {
	result := &nvver.Playback{}
	if err := json.Unmarshal([]byte(r.LivePlaybackJson), result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetLiveDetail(channelID string) (*LiveDetail, error) {
	url := fmt.Sprintf("https://api.chzzk.naver.com/service/v3.2/channels/%s/live-detail", channelID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", fmt.Sprintf("https://chzzk.naver.com/live/%s", channelID))

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
		Content LiveDetail `json:"content"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return &result.Content, nil
}
