package model

import (
    "echelon_task/pkg/proto"
    "errors"
)

func ThumbnailQualityFromProto(q proto.ThumbnailQuality) (ThumbnailQuality, error) {
    switch q {
    case proto.ThumbnailQuality_LOW:
        return LowQuality, nil
    case proto.ThumbnailQuality_MEDIUM:
        return MediumQuality, nil
    case proto.ThumbnailQuality_HIGH:
        return HighQuality, nil
    default:
        return "", errors.New("invalid quality")
    }
}

func ThumbnailQualityToProto(t ThumbnailQuality) (proto.ThumbnailQuality, error) {
    switch t {
    case LowQuality:
        return proto.ThumbnailQuality_LOW, nil
    case MediumQuality:
        return proto.ThumbnailQuality_MEDIUM, nil
    case HighQuality:
        return proto.ThumbnailQuality_HIGH, nil
    default:
        return -1, errors.New("invalid quality")
    }
}

func ThumbnailToProto(t *Thumbnail) (*proto.Thumbnail, error) {
    quality, err := ThumbnailQualityToProto(t.Quality)
    if err != nil {
        return nil, err
    }

    return &proto.Thumbnail{
        VideoId: t.VideoID,
        Payload: t.Payload,
        Quality: quality,
    }, nil
}

func ThumbnailFromProto(t *proto.Thumbnail) (*Thumbnail, error) {
    quality, err := ThumbnailQualityFromProto(t.Quality)
    if err != nil {
        return nil, err
    }

    return &Thumbnail{
        VideoID: t.VideoId,
        Payload: t.Payload,
        Quality: quality,
    }, nil
}
