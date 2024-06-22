package grpc

import (
	"context"
	"echelon_task/internal/adapters/ytbclient"
	"echelon_task/internal/model"
	"echelon_task/pkg/proto"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedThumbnailServiceServer
	client ytbclient.IYoutubeClient
	logger *slog.Logger
	redis  *redis.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) initLogger() {
	s.logger = slog.Default()
}

func (s *Server) initClient() {
	s.client = ytbclient.NewYoutubeClient(http.DefaultClient)
}

func (s *Server) initRedis() error {
	var (
		host = os.Getenv("REDIS_HOST")
		port = os.Getenv("REDIS_PORT")
	)

	s.redis = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", host, port)})
	if err := s.redis.Ping(context.Background()).Err(); err != nil {
		return err
	}

	s.client = ytbclient.NewCachedClient(s.redis, s.client)
	return nil
}

func (s *Server) Init() error {
	s.initLogger()
	s.initClient()
	return s.initRedis()
}

func (s *Server) Shutdown() error {
	return s.redis.Close()
}

func (s *Server) GetThumbnail(ctx context.Context, req *proto.GetThumbnailRequest) (*proto.Thumbnail, error) {
	quality, err := model.ThumbnailQualityFromProto(req.Quality)
	if err != nil {
		s.logger.Info("invalid quality argument", slog.Any("request", req))
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	thumb, err := s.client.GetThumbnail(ctx, req.VideoId, quality)
	if err != nil {
		s.logger.Info("failed to get thumbnail", slog.String("error", err.Error()))
		return nil, status.New(codes.Unknown, err.Error()).Err()
	}

	resp, err := model.ThumbnailToProto(thumb)
	if err != nil {
		s.logger.Info("failed to convert thumbnail to proto", slog.Any("thumbnail", thumb))
		return nil, status.New(codes.Unknown, err.Error()).Err()
	}
	return resp, nil
}
