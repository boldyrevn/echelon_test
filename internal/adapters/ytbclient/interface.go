package ytbclient

import (
	"context"
	"echelon_task/internal/model"
)

//go:generate mockgen -source interface.go -destination interface.mock.gen.go -package ytbclient -mock_names IYoutubeClient=YoutubeClientMock
type IYoutubeClient interface {
	GetThumbnail(ctx context.Context, videoID string, quality model.ThumbnailQuality) (*model.Thumbnail, error)
}
