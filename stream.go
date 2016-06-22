package rtsp

import (
	"github.com/nareix/av"
	"github.com/nareix/rtsp/sdp"
	"time"
)

type Stream struct {
	av.CodecData
	Sdp    sdp.Media
	client *Client

	// h264
	fuBuffer   []byte
	sps        []byte
	pps        []byte
	spsChanged bool
	ppsChanged bool

	gotpkt    bool
	pkt       av.Packet
	timestamp uint32
	firsttimestamp uint32

	lasttime time.Duration
}

