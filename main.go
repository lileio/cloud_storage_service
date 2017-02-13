package main

import (
	"google.golang.org/grpc"

	log "github.com/Sirupsen/logrus"
	"github.com/lileio/cloud_storage_service/cloud_storage_service"
	"github.com/lileio/cloud_storage_service/server"
	"github.com/lileio/cloud_storage_service/storage"
	"github.com/lileio/lile"
)

func main() {
	store := &storage.GoogleCloudStorage{}
	err := store.Setup()
	if err != nil {
		log.Fatalf("storage setup error: %v", err)
	}

	s := &server.Server{Storage: store}
	impl := func(g *grpc.Server) {
		cloud_storage_service.RegisterCloudStorageServiceServer(g, s)
	}

	err = lile.NewServer(
		lile.Name("cloud_storage_service"),
		lile.Implementation(impl),
	).ListenAndServe()

	log.Fatal(err)
}
