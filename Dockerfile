FROM alpine:3.5

COPY . /src/github.com/lileio/cloud_storage_service
RUN apk add --no-cache ca-certificates
ADD build/cloud_storage /bin
CMD ["cloud_storage", "server"]
