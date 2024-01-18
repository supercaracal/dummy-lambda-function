FROM golang:1.21 as builder
WORKDIR /go/src/app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -tags timetzdata -tags lambda.norpc -o bootstrap

# https://github.com/GoogleContainerTools/distroless
# https://console.cloud.google.com/gcr/images/distroless/GLOBAL
FROM gcr.io/distroless/static-debian12:nonroot-amd64
WORKDIR /var/runtime
COPY --from=builder /go/src/app/bootstrap /var/runtime/
ENTRYPOINT ["/var/runtime/bootstrap"]
