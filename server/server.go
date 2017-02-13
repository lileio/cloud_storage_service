package server

import (
	"strconv"

	context "golang.org/x/net/context"

	"github.com/lileio/cloud_storage_service/cloud_storage_service"
	"github.com/lileio/cloud_storage_service/storage"
)

type Server struct {
	cloud_storage_service.CloudStorageServiceServer
	Storage storage.Storage
}

func (s Server) Store(ctx context.Context, r *cloud_storage_service.StoreRequest) (*cloud_storage_service.StorageObject, error) {
	metadata := map[string]string{
		"master": strconv.FormatBool(r.Master),
	}

	err := s.Storage.Store(
		ctx,
		r.Filename,
		r.Data,
		metadata,
	)

	if err != nil {
		return nil, err
	}

	return &cloud_storage_service.StorageObject{
		Filename: r.Filename,
		Master:   r.Master,
	}, nil
}
