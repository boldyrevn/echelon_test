package grpc

import (
	"context"
	httpclient "echelon_task/internal/adapters/ytbclient"
	"echelon_task/internal/model"
	"echelon_task/pkg/proto"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestServer_GetThumbnail(t *testing.T) {
	mockClient := httpclient.NewYoutubeClientMock(gomock.NewController(t))
	s := &Server{
		client: mockClient,
		logger: slog.New(slog.NewJSONHandler(io.Discard, nil)),
	}

	t.Run("ValidRequest", func(t *testing.T) {
		thumbnail := &model.Thumbnail{
			VideoID: "valid_id",
			Payload: []byte("some thumbnail"),
			Quality: model.MediumQuality,
		}

		thumbProto, err := model.ThumbnailToProto(thumbnail)
		require.NoError(t, err)

		mockClient.EXPECT().
			GetThumbnail(gomock.Any(), gomock.Eq("valid_id"), gomock.Eq(model.MediumQuality)).
			Return(thumbnail, nil)

		resp, err := s.GetThumbnail(
			context.Background(),
			&proto.GetThumbnailRequest{
				VideoId: "valid_id",
				Quality: proto.ThumbnailQuality_MEDIUM,
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, thumbProto, resp)
	})

	t.Run("InvalidQuality", func(t *testing.T) {
		resp, err := s.GetThumbnail(context.Background(), &proto.GetThumbnailRequest{
			VideoId: "valid_id",
			Quality: -1,
		})

		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("ClientError", func(t *testing.T) {
		mockClient.EXPECT().
			GetThumbnail(gomock.Any(), gomock.Eq("invalid_id"), gomock.Eq(model.MediumQuality)).
			Return(nil, errors.New("requested resource can't be found"))

		resp, err := s.GetThumbnail(
			context.Background(),
			&proto.GetThumbnailRequest{
				VideoId: "invalid_id",
				Quality: proto.ThumbnailQuality_MEDIUM,
			},
		)
		assert.Error(t, err)
		assert.Nil(t, nil, resp)
	})
}
