package server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/lileio/cloud_storage_service/cloud_storage_service"
	"github.com/lileio/cloud_storage_service/storage"
	"github.com/stretchr/testify/assert"
)

var s Server

func TestMain(m *testing.M) {
	store := &storage.GoogleCloudStorage{}
	err := store.Setup()
	if err != nil {
		log.Fatalf("storage setup error: %v", err)
	}

	s = Server{Storage: store}
	os.Exit(m.Run())
}

func TestStoreDelete(t *testing.T) {
	b, err := ioutil.ReadFile("../test/testfile.txt")
	assert.Nil(t, err)

	ctx := context.Background()
	req := &cloud_storage_service.StoreRequest{
		Filename: "testfile.txt",
		Data:     b,
		Master:   true,
	}
	res, err := s.Store(ctx, req)

	assert.Nil(t, err)
	assert.NotEmpty(t, res.Filename)

	dreq := &cloud_storage_service.DeleteRequest{
		Filename: "testfile.txt",
	}
	dres, err := s.Delete(ctx, dreq)

	assert.Nil(t, err)
	assert.NotEmpty(t, dres.Filename)
}
