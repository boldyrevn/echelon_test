package model

type ThumbnailQuality string

const (
	LowQuality    ThumbnailQuality = "default"
	MediumQuality ThumbnailQuality = "mqdefault"
	HighQuality   ThumbnailQuality = "hqdefault"
)

func (tq ThumbnailQuality) String() string {
	return string(tq)
}

type Thumbnail struct {
	VideoID string
	Payload []byte
	Quality ThumbnailQuality
}
