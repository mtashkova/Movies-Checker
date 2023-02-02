FROM golang:1.13 AS builder
WORKDIR /moovies/
COPY ./ ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor -o movies cmd/movies/main.go

FROM alpine:3.10
COPY --from=builder /moovies/movies /bin/movies
CMD /bin/movies

