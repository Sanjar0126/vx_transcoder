package grpc

import (
	"fmt"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/genproto/content_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ContentService = "content-service"
)

type ServiceManager interface {
	MovieService() content_service.MovieServiceClient
	SeasonService() content_service.MovieSeasonServiceClient
	UploadedVideoService() content_service.UploadVideoServiceClient
}

type grpcClients struct {
	movieService         content_service.MovieServiceClient
	seasonService        content_service.MovieSeasonServiceClient
	uploadedVideoService content_service.UploadVideoServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	contentService, err := grpc.Dial(
		rpcConn(conf.ContentServiceHost, conf.ContentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		movieService:         content_service.NewMovieServiceClient(contentService),
		seasonService:        content_service.NewMovieSeasonServiceClient(contentService),
		uploadedVideoService: content_service.NewUploadVideoServiceClient(contentService),
	}, nil
}

func rpcConn(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func (g *grpcClients) MovieService() content_service.MovieServiceClient {
	return g.movieService
}

func (g *grpcClients) SeasonService() content_service.MovieSeasonServiceClient {
	return g.seasonService
}

func (g *grpcClients) UploadedVideoService() content_service.UploadVideoServiceClient {
	return g.uploadedVideoService
}
