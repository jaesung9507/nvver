package nvver

type Playback struct {
	Meta struct {
		VideoId   string  `json:"videoId"`
		StreamSeq float64 `json:"streamSeq"`
		LiveId    string  `json:"liveId"`
		PaidLive  bool    `json:"paidLive"`
		CDNInfo   struct {
			CDNType string `json:"cdnType"`
		} `json:"cdnInfo"`
		CmcdEnabled      bool    `json:"cmcdEnabled"`
		LiveRewind       bool    `json:"liveRewind,omitempty"`
		Duration         float64 `json:"duration,omitempty"`
		PlaybackAuthType string  `json:"playbackAuthType"`
	} `json:"meta"`
	ServiceMeta *struct {
		ContentType string `json:"contentType"`
	} `json:"serviceMeta,omitempty"`
	Live *struct {
		Start       string `json:"start"`
		Open        string `json:"open"`
		TimeMachine bool   `json:"timeMachine"`
		Status      string `json:"status"`
	} `json:"live,omitempty"`
	API []struct {
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"api"`
	Media []struct {
		MediaId       string `json:"mediaId"`
		Protocol      string `json:"protocol"`
		Path          string `json:"path"`
		Latency       string `json:"latency,omitempty"`
		EncodingTrack []struct {
			EncodingTrackId    string `json:"encodingTrackId"`
			Path               string `json:"path,omitempty"`
			VideoProfile       string `json:"videoProfile,omitempty"`
			AudioProfile       string `json:"audioProfile,omitempty"`
			P2PPath            string `json:"p2pPath,omitempty"`
			P2PPathURLEncoding string `json:"p2pPathUrlEncoding,omitempty"`
			VideoCodec         string `json:"videoCodec,omitempty"`
			VideoBitRate       int64  `json:"videoBitRate,omitempty"`
			AudioCodec         string `json:"audioCodec,omitempty"`
			AudioBitRate       int64  `json:"audioBitRate"`
			AudioOnly          bool   `json:"audioOnly,omitempty"`
			VideoFrameRate     string `json:"videoFrameRate,omitempty"`
			VideoWidth         int    `json:"videoWidth,omitempty"`
			VideoHeight        int    `json:"videoHeight,omitempty"`
			AudioSamplingRate  int64  `json:"audioSamplingRate"`
			AudioChannel       int    `json:"audioChannel"`
			AvoidReencoding    bool   `json:"avoidReencoding"`
			VideoDynamicRange  string `json:"videoDynamicRange,omitempty"`
		} `json:"encodingTrack"`
	} `json:"media"`
	Thumbnail struct {
		SnapshotThumbnailTemplate string `json:"snapshotThumbnailTemplate"`
		SpriteSeekingThumbnail    struct {
			SpriteFormat struct {
				RowCount        int    `json:"rowCount"`
				ColumnCount     int    `json:"columnCount"`
				IntervalType    string `json:"intervalType"`
				Interval        int    `json:"interval"`
				ThumbnailWidth  int    `json:"thumbnailWidth"`
				ThumbnailHeight int    `json:"thumbnailHeight"`
			} `json:"spriteFormat"`
			URLTemplate       string `json:"urlTemplate"`
			ProcessingSeconds int    `json:"processingSeconds"`
		} `json:"spriteSeekingThumbnail"`
		Types []string `json:"types"`
	} `json:"thumbnail"`
	Multiview *[]struct{} `json:"multiview,omitempty"`
}

func (i *Playback) GetHLSPath() string {
	for _, media := range i.Media {
		if media.MediaId == "HLS" {
			return media.Path
		}
	}

	return ""
}
