package server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"google.golang.org/api/option"

	"github.com/lileio/cloud_storage_service"
	"github.com/lileio/cloud_storage_service/storage"
	"github.com/stretchr/testify/assert"
)

var s Server

func TestMain(m *testing.M) {
	keyLocation := "../google_key.json"
	o := []option.ClientOption{}

	key := os.Getenv("GOOGLE_KEY")
	if key != "" {
		err := ioutil.WriteFile(keyLocation, []byte(key), os.ModePerm)
		if err != nil {
			panic(err)
		}

		o = append(o, option.WithServiceAccountFile(keyLocation))
	}

	store := &storage.GoogleCloudStorage{Options: o}
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
	}
	res, err := s.Store(ctx, req)

	assert.Nil(t, err)
	assert.NotEmpty(t, res.Filename)
	assert.NotEmpty(t, res.Url)

	dreq := &cloud_storage_service.DeleteRequest{
		Filename: "testfile.txt",
	}
	dres, err := s.Delete(ctx, dreq)

	assert.Nil(t, err)
	assert.NotEmpty(t, dres.Filename)
}
