package ytbclient

import (
    "context"
    "echelon_task/internal/model"
    "errors"
    "fmt"
    "io"
    "net/http"
)

type youtubeClient struct {
    client *http.Client
}

func NewYoutubeClient(client *http.Client) IYoutubeClient {
    return &youtubeClient{client: client}
}

func (c *youtubeClient) GetThumbnail(
    ctx context.Context,
    videoID string,
    quality model.ThumbnailQuality,
) (*model.Thumbnail, error) {
    baseURL := fmt.Sprintf("https://img.youtube.com/vi/%s/", videoID)
    switch quality {
    case model.LowQuality, model.MediumQuality, model.HighQuality:
        baseURL += fmt.Sprintf("%s.jpg", quality)
    default:
        return nil, errors.New("invalid thumbnail quality")
    }

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to do request: %w", err)
    }

    payload, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    return &model.Thumbnail{
        VideoID: videoID,
        Payload: payload,
        Quality: quality,
    }, nil
}
