ARG DIR=/app

FROM golang:1.14.4 AS build-env
LABEL author="huynhquangthao@gmail.com"

ARG DIR
WORKDIR $DIR

COPY go.mod $DIR/
RUN GO111MODULE=on go mod download
COPY . $DIR/
RUN CGO_ENABLED=0 go build -o server main.go

# second stage: copy golang binary to running image
FROM alpine:latest

ARG DIR
WORKDIR $DIR

COPY --from=build-env $DIR/server $DIR/
RUN chmod +x ./server
CMD ["sh", "-c", "./server"]




