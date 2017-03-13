package cmd

import (
	"log"

	"github.com/lileio/cloud_storage_service"
	"github.com/lileio/cloud_storage_service/server"
	"github.com/lileio/cloud_storage_service/storage"
	"github.com/lileio/lile"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
