FROM alpine:3.5

RUN apk add --no-cache ca-certificates
ADD cloud_storage /
CMD ["/cloud_storage"]
