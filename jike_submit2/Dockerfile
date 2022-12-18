FROM golang:1.18.9 AS build-env
COPY main.go /main.go
WORKDIR /
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM alpine
COPY --from=build-env /main /main
WORKDIR /
ENTRYPOINT ["/main"]
