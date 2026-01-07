# NVVER
**NVVER** is an unofficial NAVER video API client for Go

## Supported Services
### CHZZK
- [x] Live
- [x] Clip
- [x] Video
### NAVER TV
- [x] Live
- [x] Clip
- [x] VOD
### NAVER Shopping LIVE
- [x] Live
- [x] Short Clip

## Example
Get HLS URL from a CHZZK live stream:
```go
package main

import (
	"fmt"

	"github.com/jaesung9507/nvver/chzzk"
)

func main() {
	client := chzzk.NewClient(nil)

	liveDetail, err := client.GetLiveDetail("CHANNEL_ID")
	if err != nil {
		panic(err)
	}

	playback, err := liveDetail.GetLivePlayback()
	if err != nil {
		panic(err)
	}

	fmt.Println("HLS URL:", playback.GetHLSPath())
}
```

## Disclaimer
This Go client uses unofficial APIs of NAVER and is not affiliated with them in any way.

NAVER may change or disable these APIs at any time, which may cause this client to stop working without notice.

This package does **not provide any backward compatibility guarantees**.

Breaking changes may be introduced at any time.

## License
MIT License
