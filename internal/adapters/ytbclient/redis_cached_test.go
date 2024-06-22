package ytbclient

import (
	"context"
	"echelon_task/internal/model"
	"fmt"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func init() {
	_ = os.Setenv("REDIS_HOST", "localhost")
	_ = os.Setenv("REDIS_PORT", "6379")
}

func TestCachedClient(t *testing.T) {
	var (
		host = os.Getenv("REDIS_HOST")
		port = os.Getenv("REDIS_PORT")
	)

	redisClient := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", host, port)})
	require.NoError(t, redisClient.Ping(context.Background()).Err())

	ytbMock := NewYoutubeClientMock(gomock.NewController(t))
	client := NewCachedClient(redisClient, ytbMock)

	err := redisClient.FlushDB(context.Background()).Err()
	require.NoError(t, err)

	expectedThumb := &model.Thumbnail{
		VideoID: "some_id",
		Payload: []byte("some payload"),
		Quality: model.MediumQuality,
	}

	ytbMock.EXPECT().
		GetThumbnail(gomock.Any(), "some_id", model.MediumQuality).
		Return(expectedThumb, nil).Times(1)

	for i := 0; i < 2; i++ {
		actualThumb, err := client.GetThumbnail(context.Background(), "some_id", model.MediumQuality)
		assert.NoError(t, err)
		assert.Equal(t, expectedThumb, actualThumb)
	}
}
