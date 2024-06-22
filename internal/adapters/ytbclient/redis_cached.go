package ytbclient

import (
	"context"
	"echelon_task/internal/model"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type cachedClient struct {
	innerClient IYoutubeClient
	redisConn   *redis.Client
}

func getThumbnailID(videoID, quality string) string {
	return fmt.Sprintf("%s:%s", videoID, quality)
}

func (c *cachedClient) cacheThumbnailPayload(ctx context.Context, key string, p []byte) error {
	return c.redisConn.Set(ctx, key, p, 0).Err()
}

func (c *cachedClient) GetThumbnail(ctx context.Context, videoID string, quality model.ThumbnailQuality) (*model.Thumbnail, error) {
	res, err := c.redisConn.Get(ctx, getThumbnailID(videoID, quality.String())).Bytes()
	if err == nil {
		return &model.Thumbnail{
			VideoID: videoID,
			Payload: res,
			Quality: quality,
		}, nil
	}

	thumb, err := c.innerClient.GetThumbnail(ctx, videoID, quality)
	return thumb, c.cacheThumbnailPayload(ctx, getThumbnailID(videoID, quality.String()), thumb.Payload)
}

func (c *cachedClient) Close() error {
	return c.redisConn.Close()
}

func NewCachedClient(redis *redis.Client, client IYoutubeClient) IYoutubeClient {
	return &cachedClient{
		innerClient: client,
		redisConn:   redis,
	}
}
