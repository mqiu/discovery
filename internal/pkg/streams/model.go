package streams

import (
	"strings"
)

type Stream struct {
	ID string `json:id,omitempty`
	StreamUrl string `json:streamUrl,omitempty`
	Captions Caption `json:captions`
}

type Caption struct {
	VTT CaptionType `json:vtt`
	SCC CaptionType `json:scc`
}

type CaptionType struct {
	EN string `json:en`
}

type Ads struct {
	BreakOffSets []*OffSet `json:breakOffsets`
	Breaks []*Breaks `json:breaks`
}

type OffSet struct {
	Index int `json:index`
	Time float64 `json:timeOffset`
}

type Break struct {
	Ads []*AdsDetail `json:ads`
	BreakID string `json:breakId`
	Duration int `json:duration`
	Events []*Event `json:events`
	Position string `json:position`
	TimeOffSet float64 `json:timeOffset`
	Type string `json:type`
}

type AdsDetail struct {
	Creative string `json:creative`
	Duration int `json:duration`
	Events []*Event `json:events`
}
type Event struct {
	Impressions []string `json:impressions`
}
