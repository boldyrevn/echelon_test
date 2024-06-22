package grpc

import (
    "context"
    "echelon_task/internal/model"
    "echelon_task/pkg/proto"
    "errors"
    "fmt"
    "golang.org/x/sync/errgroup"
    "net/url"
    "os"
    "path"
)

const videoIDQueryParam = "v"

func saveImage(t *model.Thumbnail, dir string) error {
    if t.Payload == nil {
        return errors.New("payload must be not nil")
    }

    fileName := fmt.Sprintf("%s_%s", t.VideoID, t.Quality)
    file, err := os.OpenFile(path.Join(dir, fileName), os.O_CREATE|os.O_WRONLY, 0660)
    if err != nil {
        return err
    }

    return err
}

func parseVideoID(videoURL string) (string, error) {
    parsedURL, err := url.Parse(videoURL)
    if err != nil {
        return "", err
    }

    videoID := parsedURL.Query().Get(videoIDQueryParam)
    if videoID == "" {
        return videoID, errors.New("no video id query param")
    }
    return videoID, nil
}

func DownloadFilesAsynchronously(
    ctx context.Context,
    c proto.ThumbnailServiceClient,
    quality model.ThumbnailQuality,
    dir string,
    videoURLs ...string,
) error {
    var wg errgroup.Group
    for _, videoURL := range videoURLs {
        wg.Go(func() error {
            return DownloadFiles(ctx, c, quality, dir, videoURL)
        })
    }

    return wg.Wait()
}

func DownloadFiles(
    ctx context.Context,
    c proto.ThumbnailServiceClient,
    quality model.ThumbnailQuality,
    dir string,
    videoURLs ...string,
) error {
    qualityProto, err := model.ThumbnailQualityToProto(quality)
    if err != nil {
        return err
    }

    for _, videoID := range videoURLs {
        thumbProto, err := c.GetThumbnail(ctx, &proto.GetThumbnailRequest{
            VideoId: videoID,
            Quality: qualityProto,
        })
        if err != nil {
            return fmt.Errorf("failed to download thumbnail: %w", err)
        }

        thumb, err := model.ThumbnailFromProto(thumbProto)
        if err != nil {
            return fmt.Errorf("failed to convert thumbnail: %w", err)
        }

        if err = saveImage(thumb, dir); err != nil {
            return fmt.Errorf("failed to save thumbnail: %w", err)
        }
    }
    return nil
}
