# Cloud Storage Service
[![Build Status](https://travis-ci.org/lileio/cloud_storage_service.svg?branch=master)](https://travis-ci.org/lileio/cloud_storage_service)

A gRPC service made with the [Lile generator](https://github.com/lileio/lile) for storing object, files etc in cloud storage like Google Cloud Storage, AWS S3 or similar (PR's welcome!)

``` protobuf
service CloudStorageService {
  rpc Store(StoreRequest) returns (StorageObject) {}
  rpc Delete(DeleteRequest) returns (DeleteResponse) {}
}
```

## Details

The cloud storage service is for storing objects that other services may require or be used later by clients. An example would be uploading a photo along with a user account. In this scenario you would use the `Store` method to store the object in cloud storage, which will return you a filename and url. You can then store this for later use by a browser or similar.

Be aware that for this url to work the bucket in question must be public.

Specific object ACL's and URL generation are coming soon.

## Docker

Builds (based on Alpine) of master (after test runs) are available at

```
docker pull lileio/cloud_storage_service
```

## Setup

### Google Cloud Storage

The service will create the cloud storage bucket on first run if it doesn't exist

```
GOOGLE_STORAGE_BUCKET=some-bucket
GOOGLE_STORAGE_PROJECT_ID=googleprojectid-20142
GOOGLE_STORAGE_LOCATION=eu
```

## Development

PR's are welcome. Testing currently is done end to end with communication with the services in question.
