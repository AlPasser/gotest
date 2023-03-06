FROM golang:1.18.9 AS build-env
RUN mkdir -p /go/src/jike_submit5/metrics
COPY main.go /go/src/jike_submit5/main.go
COPY metrics/ /go/src/jike_submit5/metrics/
WORKDIR /go/src/jike_submit5/
RUN go mod init && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM alpine
COPY --from=build-env /go/src/jike_submit5/main /main
WORKDIR /
ENTRYPOINT ["/main"]